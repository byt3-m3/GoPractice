package main

func main() {
	// With builder function
	svc := NewService("base", repoBuilder)
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

func NewService(repoType string, builder func(repoType string) Repository) *service {
	repo := builder(repoType)

	return &service{repo: repo}
}

func repoBuilder(repoType string) Repository {
	if repoType == "base" {
		table := DBTable{}
		return baseRepo{dbtable: table}
	}
	return nil
}
