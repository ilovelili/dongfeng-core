package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetNamelist get name list master data
func (f *Facade) GetNamelist(ctx context.Context, req *proto.GetNamelistRequest, rsp *proto.GetNamelistResponse) error {
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

	namelistcontroller := controllers.NewNamelistController()
	namelists, err := namelistcontroller.GetNamelists(req.GetClass(), req.GetYear())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetNamelist)
	}

	itemmap := make(map[string] /*year_classid*/ []string)
	for _, namelist := range namelists {
		key := fmt.Sprintf("%s_%s", namelist.Year, namelist.Class)
		if names, ok := itemmap[key]; ok {
			itemmap[key] = append(names, namelist.Name)
		} else {
			itemmap[key] = []string{namelist.Name}
		}
	}

	items := make([]*proto.NamelistItem, 0)
	for k, v := range itemmap {
		year, class := strings.Split(k, "_")[0], strings.Split(k, "_")[1]
		items = append(items, &proto.NamelistItem{
			Year:  year,
			Class: class,
			Names: v,
		})
	}

	rsp.Items = items
	return nil
}
