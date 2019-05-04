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

// GetPhysiques get physiques
func (f *Facade) GetPhysiques(ctx context.Context, req *proto.GetPhysiqueRequest, rsp *proto.GetPhysiqueResponse) error {
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

	physiquecontroller := controllers.NewPhysiqueController()
	physiques, err := physiquecontroller.GetPhysiques(req.GetClass(), req.GetYear(), req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiques)
	}

	items := []*proto.Physique{}
	for _, physique := range physiques {
		items = append(items, &proto.Physique{
			Id:            physique.ID,
			Year:          physique.Year,
			Class:         physique.Class,
			Name:          physique.Name,
			Gender:        proto.Physique_Gender(physique.Gender),
			BirthDate:     physique.BirthDate,
			ExamDate:      physique.ExamDate,
			Height:        physique.Height,
			Weight:        physique.Weight,
			Age:           physique.Age,
			AgeCmp:        physique.AgeComparison,
			HeightP:       physique.HeightP,
			WeightP:       physique.WeightP,
			HeightWeightP: physique.HeightToWeightP,
			Bmi:           physique.BMI,
			FatCofficient: physique.FatCofficient,
			Conclusion:    physique.Conclusion,
			CreatedBy:     physique.CreatedBy,
		})
	}

	rsp.Physiques = items
	return nil
}

// UpdatePhysique update physique
func (f *Facade) UpdatePhysique(ctx context.Context, req *proto.UpdatePhysiqueRequest, rsp *proto.UpdatePhysiqueResponse) error {
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

	physiques := req.GetPhysiques()
	if len(physiques) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdatePhysiques)
	}

	physique := physiques[0]
	physique.CreatedBy = exsitinguser.Email

	physiquecontroller := controllers.NewPhysiqueController()
	err = physiquecontroller.UpdatePhysique(&models.Physique{
		ID:        physique.GetId(),
		Year:      physique.GetYear(),
		Name:      physique.GetName(),
		Class:     physique.GetClass(),
		Gender:    int64(physique.GetGender()),
		BirthDate: physique.GetBirthDate(),
		ExamDate:  physique.GetExamDate(),
		Height:    physique.GetHeight(),
		Weight:    physique.GetWeight(),
		CreatedBy: physique.GetCreatedBy(),
	})

	f.syslog(notification.PhysiqueUpdated(exsitinguser.ID))
	return err
}

// UpdatePhysiques update physiques
func (f *Facade) UpdatePhysiques(ctx context.Context, req *proto.UpdatePhysiqueRequest, rsp *proto.UpdatePhysiqueResponse) error {
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

	physiques := []*models.Physique{}
	for _, physique := range req.GetPhysiques() {
		physiques = append(physiques, &models.Physique{
			ID:        physique.GetId(),
			Year:      physique.GetYear(),
			Name:      physique.GetName(),
			Class:     physique.GetClass(),
			Gender:    int64(physique.GetGender()),
			BirthDate: physique.GetBirthDate(),
			ExamDate:  physique.GetExamDate(),
			Height:    physique.GetHeight(),
			Weight:    physique.GetWeight(),
			CreatedBy: exsitinguser.Email,
		})
	}

	physiquecontroller := controllers.NewPhysiqueController()
	err = physiquecontroller.UpdatePhysiques(physiques)

	f.syslog(notification.PhysiqueUpdated(exsitinguser.ID))
	return err
}
