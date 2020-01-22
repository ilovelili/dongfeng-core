package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetTeachers get teachers
func (f *Facade) GetTeachers(ctx context.Context, req *proto.GetTeacherRequest, rsp *proto.GetTeacherResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
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
			Role:      *teacher.Role,
			CreatedBy: teacher.CreatedBy,
		})
	}

	rsp.Teachers = items
	return nil
}

// UpdateTeacher update teacher
func (f *Facade) UpdateTeacher(ctx context.Context, req *proto.UpdateTeacherRequest, rsp *proto.UpdateTeacherResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	teachers := req.GetTeachers()
	if len(teachers) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdateTeachers)
	}

	teacher := teachers[0]
	teacher.CreatedBy = user.Email
	id, email, role, name, class := teacher.GetId(), teacher.GetEmail(), teacher.GetRole(), teacher.GetName(), teacher.GetClass()
	teachercontroller := controllers.NewTeacherController()
	err = teachercontroller.UpdateTeacher(&models.Teacher{
		ID:    id,
		Name:  name,
		Class: class,
		Role:  &role,
		Email: email,
	})

	f.syslog(notification.NamelistUpdated(user.ID))
	return err
}

// UpdateTeachers update teachers
func (f *Facade) UpdateTeachers(ctx context.Context, req *proto.UpdateTeacherRequest, rsp *proto.UpdateTeacherResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	teachers := req.GetTeachers()
	for _, teacher := range teachers {
		teacher.CreatedBy = user.Email
	}

	teachercontroller := controllers.NewTeacherController()
	err = teachercontroller.UpdateTeachers(teachers)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateTeachers)
	}

	f.syslog(notification.TeacherlistUpdated(user.ID))
	return nil
}
