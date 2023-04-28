package gateways

/*
type UserGatewayInterface interface {
	SaveUser(user entity.User)
}

type UserGateway struct {
	model      *models.UserModel
	repository *repository.UserRepository
}

func NewOrderGateway(model *models.UserModel, repository *repository.UserRepository) *UserGateway {
	return &UserGateway{model, repository}
}

func (g *UserGateway) SaveUser(user *entity.User) (*entity.User, error) {
	userModel := g.userEntityToModel(user)
	usedSaved, err := g.repository.Save(userModel)
	if err != nil {
		return nil, err
	}
	userEntity := g.userModelToEntity(usedSaved)
	return userEntity, nil
}

func (g *UserGateway) userEntityToModel(user *entity.User) *models.UserModel {
	//g.model.ID = order.Id
	g.model.IdUser = user.IdUser
	g.model.Rut = user.Rut
	g.model.FirstName = user.FirstName
	g.model.LastName = user.LastName
	g.model.Email = user.Email
	g.model.UserName = user.UserName
	g.model.Password = user.Password

	return g.model
}

func (g *UserGateway) userModelToEntity(userModel *models.UserModel) *entity.User {
	userEntity := entity.User{}
	userEntity.IdUser = userModel.IdUser
	userEntity.Rut = userModel.Rut
	userEntity.FirstName = userModel.FirstName
	userEntity.LastName = userModel.LastName
	userEntity.Email = userModel.Email
	userEntity.Username = userModel.UserName
	userEntity.Password = userModel.Password

	return &userEntity
}
*/
