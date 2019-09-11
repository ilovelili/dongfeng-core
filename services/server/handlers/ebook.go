package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetEbooks get ebooks
func (f *Facade) GetEbooks(ctx context.Context, req *proto.GetEbooksRequest, rsp *proto.GetEbooksResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	ebookcontroller := controllers.NewEbookController()
	ebooks, err := ebookcontroller.GetEbooks(req.GetYear(), req.GetClass(), req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetEbook)
	}

	_ebooks := []*proto.Ebook{}
	for _, ebook := range ebooks {
		_ebooks = append(_ebooks, &proto.Ebook{
			Year:  ebook.Year,
			Class: ebook.Class,
			Name:  ebook.Name,
			Dates: ebook.Dates,
		})
	}
	rsp.Ebooks = _ebooks
	return nil
}

// UpdateEbook update ebook
func (f *Facade) UpdateEbook(ctx context.Context, req *proto.UpdateEbookRequest, rsp *proto.UpdateEbookResponse) error {
	pid := req.GetPid()
	userinfo, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	var user *models.User
	err = json.Unmarshal(userinfo, &user)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	_, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	ebookcontroller := controllers.NewEbookController()
	err = ebookcontroller.SaveEbook(&models.Ebook{
		Year:   req.GetYear(),
		Class:  req.GetClass(),
		Name:   req.GetName(),
		Date:   req.GetDate(),
		HTML:   req.GetHtml(),
		CSS:    req.GetCss(),
		Images: req.GetImages(),
	})

	return err
}
