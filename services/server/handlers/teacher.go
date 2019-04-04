package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetTeachers get teachers
func (f *Facade) GetTeachers(ctx context.Context, req *proto.GetTeacherRequest, rsp *proto.GetTeacherResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	_, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	teachercontroller := controllers.NewTeacherController()
	teachers, err := teachercontroller.GetTeachers(req.GetClass(), req.GetYear())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetTeachers)
	}

	items := []*proto.Teacher{}
	for _, teacher := range teachers {
		items = append(items, &proto.Teacher{
			Id:        teacher.ID,
			Year:      teacher.Year,
			Name:      teacher.Name,
			Class:     teacher.Class,
			Email:     teacher.Email,
			Role:      teacher.Role,
			CreatedBy: teacher.CreatedBy,
		})
	}

	rsp.Teachers = items
	return nil
}

// UpdateTeachers update teachers
func (f *Facade) UpdateTeachers(ctx context.Context, req *proto.UpdateTeacherRequest, rsp *proto.UpdateTeacherResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	claims, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// Unmarshal user info
	userinfo, _ := json.Marshal(claims)
	var user *models.User
	err = json.Unmarshal(userinfo, &user)

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	teachers := req.GetTeachers()
	for _, teacher := range teachers {
		teacher.CreatedBy = exsitinguser.Email
	}

	teachercontroller := controllers.NewTeacherController()
	err = teachercontroller.UpdateTeachers(teachers)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateTeachers)
	}

	f.syslog(notification.TeacherlistUpdated(user.ID))
	return nil
}
