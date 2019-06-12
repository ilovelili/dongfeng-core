package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetEbooks get ebooks
func (f *Facade) GetEbooks(ctx context.Context, req *proto.GetEbooksRequest, rsp *proto.GetEbooksResponse) error {
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
		})
	}
	rsp.Ebooks = _ebooks
	return nil
}

// UpdateEbook update ebook
func (f *Facade) UpdateEbook(ctx context.Context, req *proto.UpdateEbookRequest, rsp *proto.UpdateEbookResponse) error {
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
	_, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	profilecontroller := controllers.NewEbookController()
	err = profilecontroller.SaveEbook(&models.Ebook{
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
