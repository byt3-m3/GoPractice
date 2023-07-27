package main

type Repository interface {
	GetModel(id string) interface{}
}

type Service interface {
	DoStuff() error
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

func (s service) DoStuff() error {
	return nil
}

type ServiceBuilder struct {
	repo Repository
}

func (b *ServiceBuilder) WithRepository(repo Repository) *ServiceBuilder {
	b.repo = repo
	return b
}
func (b ServiceBuilder) Build() Service {
	s := service{repo: b.repo}
	return s

}

func NewService(repoType string, builder func(repoType string) Repository) *service {
	repo := builder(repoType)

	return &service{repo: repo}
}
