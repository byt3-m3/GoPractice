package main

func main() {
	repo := baseRepo{dbtable: DBTable{}}
	svc := NewService(repo)
	_ = svc.repo.GetModel("test_id")

}

type Repository interface {
	GetModel(id string) interface{}
}

type DBTable struct {
}

type baseRepo struct {
	dbtable DBTable
}

func (receiver baseRepo) GetModel(id string) interface{} {
	return nil
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {

	return &service{repo: repo}
}
