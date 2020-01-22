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

// GetEbooks get ebooks
func (f *Facade) GetEbooks(ctx context.Context, req *proto.GetEbooksRequest, rsp *proto.GetEbooksResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
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
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
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

	f.syslog(notification.EbookUpdated(user.ID))
	return err
}
