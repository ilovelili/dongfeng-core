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

// GetProcurements get procurement
func (f *Facade) GetProcurements(ctx context.Context, req *proto.GetProcurementRequest, rsp *proto.GetProcurementResponse) error {
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
		return utils.NewError(errorcode.NutritionNoUser)
	}

	procurementcontroller := controllers.NewProcurementController()
	from, to := req.GetFrom(), req.GetTo()

	procurements, err := procurementcontroller.GetProcurements(from, to)
	if err != nil {
		return utils.NewError(errorcode.NutritionFailedToGetProcurement)
	}

	rsp.Procurements = procurements
	return nil
}

// UpdateProcurement update procurement
func (f *Facade) UpdateProcurement(ctx context.Context, req *proto.UpdateProcurementRequest, rsp *proto.UpdateProcurementResponse) error {
	return nil
}
