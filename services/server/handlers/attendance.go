package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetAttendance get attendance
func (f *Facade) GetAttendance(ctx context.Context, req *proto.GetAttendanceRequest, rsp *proto.GetAttendanceResponse) error {
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

	attendancecontroller := controllers.NewAttendanceController()
	attendances, err := attendancecontroller.Get(req.GetFrom(), req.GetTo(), req.GetClass(), req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetAttendance)
	}

	_attendances := make([]*proto.Attendance, 0)
	for _, attendance := range attendances {
		_attendances = append(_attendances, &proto.Attendance{
			Id:    attendance.ID,
			Date:  attendance.Date,
			Class: attendance.Class,
			Name:  attendance.Name,
		})
	}

	rsp.Attendances = _attendances
	return nil
}

// UpdateAttendance update attendance
func (f *Facade) UpdateAttendance(ctx context.Context, req *proto.UpdateAttendanceRequest, rsp *proto.UpdateAttendanceResponse) error {
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
			class, err := resolveClass(attendance.GetClass())
			if err != nil {
				utils.NewError(errorcode.CoreFailedToSaveAttendance)
			}

			attendances = append(attendances, &models.Attendance{
				CreatedBy: user.Email,
				Date:      fmt.Sprintf("%04d-%02d-%02d", attendance.GetYear(), attendance.GetMonth(), day),
				Class:     class,
				Name:      attendance.GetName(),
			})
		}
	}

	attendancecontroller := controllers.NewAttendanceController()
	err = attendancecontroller.Save(attendances)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveAttendance)
	}

	f.syslog(notification.AttendanceUpdated(user.ID))
	return nil
}

func resolveClass(class string) (classid string, err error) {
	switch class {
	case "小一班":
		fallthrough
	case "小1班":
		fallthrough
	case "小1":
		fallthrough
	case "小一":
		classid = "01"
		return

	case "小二班":
		fallthrough
	case "小2班":
		fallthrough
	case "小2":
		fallthrough
	case "小二":
		classid = "02"
		return

	case "小三班":
		fallthrough
	case "小3班":
		fallthrough
	case "小3":
		fallthrough
	case "小三":
		classid = "03"
		return

	case "小四班":
		fallthrough
	case "小4班":
		fallthrough
	case "小4":
		fallthrough
	case "小四":
		classid = "04"
		return

	case "中一班":
		fallthrough
	case "中1班":
		fallthrough
	case "中一":
		fallthrough
	case "中1":
		classid = "11"
		return

	case "中二班":
		fallthrough
	case "中2班":
		fallthrough
	case "中二":
		fallthrough
	case "中2":
		classid = "12"
		return

	case "中三班":
		fallthrough
	case "中3班":
		fallthrough
	case "中三":
		fallthrough
	case "中3":
		classid = "13"
		return

	case "大一班":
		fallthrough
	case "大1班":
		fallthrough
	case "大一":
		fallthrough
	case "大1":
		classid = "21"
		return

	case "大二班":
		fallthrough
	case "大2班":
		fallthrough
	case "大二":
		fallthrough
	case "大2":
		classid = "22"
		return

	case "大三班":
		fallthrough
	case "大3班":
		fallthrough
	case "大三":
		fallthrough
	case "大3":
		classid = "23"
		return
	}

	err = fmt.Errorf("invalid class")
	return
}
