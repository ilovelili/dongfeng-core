package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetAttendances get attendances
func (f *Facade) GetAttendances(ctx context.Context, req *proto.GetAttendanceRequest, rsp *proto.GetAttendanceResponse) error {
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
	attendances, err := attendancecontroller.SelectAttendances(req.GetYear(), req.GetFrom(), req.GetTo(), req.GetClass(), req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetAttendances)
	}

	attendancemap := make(map[string] /*year_class_date*/ [][]string)
	for _, attendance := range attendances.Attendances {
		key := fmt.Sprintf("%s_%s_%s", attendance.Year, attendance.Class, attendance.Date)
		if v, ok := attendancemap[key]; ok {
			if attendance.AttendanceFlag {
				attendancemap[key][0] = append(v[0], attendance.Name)
			} else {
				attendancemap[key][1] = append(v[1], attendance.Name)
			}
		} else {
			if attendance.AttendanceFlag {
				attendancesubslice := []string{attendance.Name}
				attendancemap[key] = [][]string{attendancesubslice, []string{}}
			} else {
				absencesubslice := []string{attendance.Name}
				attendancemap[key] = [][]string{[]string{}, absencesubslice}
			}
		}
	}

	_attendances := []*proto.Attendance{}
	for k, v := range attendancemap {
		segments := strings.Split(k, "_")
		if len(segments) != 3 {
			return utils.NewError(errorcode.CoreFailedToGetAttendances)
		}

		year, class, date := segments[0], segments[1], segments[2]
		_attendances = append(_attendances, &proto.Attendance{
			Year:        year,
			Class:       class,
			Date:        date,
			Attendances: v[0],
			Absences:    v[1],
		})
	}

	holidays := []*proto.Holiday{}
	for _, h := range attendances.Holidays {
		holidays = append(holidays, &proto.Holiday{
			Date: h.Date,
			Type: h.Type,
		})
	}

	rsp.Attendances = _attendances
	rsp.Holidays = holidays
	return nil
}

// UpdateAttendance update single attendance
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

	absences := []*models.Absence{}
	for _, attendance := range req.Attendances {
		names := attendance.GetAbsences()
		for _, name := range names {
			absences = append(absences, &models.Absence{
				Year:      attendance.GetYear(),
				Date:      attendance.GetDate(),
				Class:     attendance.GetClass(),
				Name:      name,
				CreatedBy: user.Email,
			})
		}
	}

	attendances := []*models.Absence{}
	for _, attendance := range req.Attendances {
		names := attendance.GetAttendances()
		for _, name := range names {
			attendances = append(attendances, &models.Absence{
				Year:      attendance.GetYear(),
				Date:      attendance.GetDate(),
				Class:     attendance.GetClass(),
				Name:      name,
				CreatedBy: user.Email,
			})
		}
	}

	attendancecontroller := controllers.NewAttendanceController()
	if err := attendancecontroller.UpdateAbsence(absences, attendances); err != nil {
		err = utils.NewError(errorcode.CoreFailedToUpdateAttendances)
	}

	f.syslog(notification.AttendanceUpdated(user.ID))
	return nil
}

// UpdateAttendances update attendances
func (f *Facade) UpdateAttendances(ctx context.Context, req *proto.UpdateAttendanceRequest, rsp *proto.UpdateAttendanceResponse) error {
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

	absences := []*models.Absence{}
	for _, attendance := range req.Attendances {
		names := attendance.GetAbsences()
		for _, name := range names {
			absences = append(absences, &models.Absence{
				Year:      attendance.GetYear(),
				Date:      attendance.GetDate(),
				Class:     attendance.GetClass(),
				Name:      name,
				CreatedBy: user.Email,
			})
		}
	}

	attendancecontroller := controllers.NewAttendanceController()
	if err := attendancecontroller.UpdateAbsences(absences); err != nil {
		err = utils.NewError(errorcode.CoreFailedToUpdateAttendances)
	}

	f.syslog(notification.AttendanceUpdated(user.ID))
	return nil
}
