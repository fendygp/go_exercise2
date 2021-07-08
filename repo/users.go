package repo

//UsersRepoStruct your function for query db, consume another service cant define here
type UsersRepoStruct struct {
}

type UsersInterface interface {
	CreateUsers() error
}

func UsersTokenRepoImp() UsersInterface {
	return &UsersRepoStruct{}
}

func (g *UsersRepoStruct) CreateUsers() error {

	return nil
}
