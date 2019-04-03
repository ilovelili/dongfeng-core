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

// GetTeacherlist get teacher list master data
func (f *Facade) GetTeacherlist(ctx context.Context, req *proto.GetTeacherlistRequest, rsp *proto.GetTeacherlistResponse) error {
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

	teacherlistcontroller := controllers.NewTeacherlistController()
	teacherlists, err := teacherlistcontroller.GetTeacherlists(req.GetClass(), req.GetYear())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetTeacherlist)
	}

	itemmap := make(map[string] /*year*/ []*proto.TeacherItem)
	for _, teacherlist := range teacherlists {
		key := teacherlist.Year
		if items, ok := itemmap[key]; ok {
			itemmap[key] = append(items, &proto.TeacherItem{
				Id:    teacherlist.ID,
				Name:  teacherlist.Name,
				Class: teacherlist.Class,
				Email: teacherlist.Email,
				Role:  teacherlist.Role,
			})
		} else {
			itemmap[key] = []*proto.TeacherItem{&proto.TeacherItem{
				Id:    teacherlist.ID,
				Name:  teacherlist.Name,
				Class: teacherlist.Class,
				Email: teacherlist.Email,
				Role:  teacherlist.Role,
			}}
		}
	}

	items := make([]*proto.TeacherlistItem, 0)
	for k, v := range itemmap {
		items = append(items, &proto.TeacherlistItem{
			Year:  k,
			Items: v,
		})
	}

	rsp.Items = items
	return nil
}

// UpdateTeacherlist update teacher list master data
func (f *Facade) UpdateTeacherlist(ctx context.Context, req *proto.UpdateTeacherlistRequest, rsp *proto.UpdateTeacherlistResponse) error {
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

	teacherlists := req.GetItems()
	for _, teacherlist := range teacherlists {
		teacherlist.CreatedBy = exsitinguser.Email
	}

	teacherlistcontroller := controllers.NewTeacherlistController()
	err = teacherlistcontroller.UpdateTeacherlists(teacherlists)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateTeacherlist)
	}

	f.syslog(notification.TeacherlistUpdated(user.ID))
	return nil
}
