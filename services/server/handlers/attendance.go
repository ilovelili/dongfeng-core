package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// UpdateAttendance update attendance
func (f *Facade) UpdateAttendance(ctx context.Context, req *proto.AttendanceRequest, rsp *proto.AttendanceResponse) error {
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
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	attendances := make([]*models.Attendance, 0)
	for _, attendance := range req.Attendances {
		for _, day := range attendance.GetAttendances() {
			attendances = append(attendances, &models.Attendance{
				CreatedBy: user.Email,
				Date:      fmt.Sprintf("%d-%d-%d", attendance.GetYear(), attendance.GetMonth(), day),
				Class:     attendance.GetClass(),
				Name:      attendance.GetName(),
			})
		}
	}

	attendancecontroller := controllers.NewAttendanceController()
	err = attendancecontroller.Save(attendances)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveAttendance)
	}

	return nil
}
