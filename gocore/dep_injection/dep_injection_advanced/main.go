package main

type (
	ReadModel interface {
		GetAggregate(id string, dbFetcher func(id string) interface{}) aggregate
	}

	Repo interface {
		GetAggregate(id string) aggregate
	}
)

type (
	aggregate struct {
		ID string
	}
)
type basicRepo struct {
	readModel ReadModel
	fetcher   func(id string) interface{}
}
type NewBasicRepoInput struct {
	modelType        string
	readModelBuilder func(modelType string) ReadModel
	fetcher          func(id string) interface{}
}

func NewBasicRepo(input NewBasicRepoInput) Repo {
	rModel := input.readModelBuilder(input.modelType)
	return basicRepo{
		readModel: rModel,
		fetcher:   DefaultDBFetcher,
	}
}

func (r basicRepo) GetAggregate(id string) aggregate {
	return r.readModel.GetAggregate(id, r.fetcher)
}

type basicReadModel struct {
}

func (b basicReadModel) GetAggregate(id string, dbFetcher func(id string) interface{}) aggregate {
	result := dbFetcher(id).(aggregate)
	result.ID = "1"
	return result
}

func DefaultDBFetcher(id string) interface{} {
	return aggregate{}
}

func ReadModelBuilder(modelType string) ReadModel {
	if modelType == "basic" {
		return basicReadModel{}
	}

	return nil
}
