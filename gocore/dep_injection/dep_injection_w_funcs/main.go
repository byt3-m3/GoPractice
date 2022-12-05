package main

func main() {
	// With builder function

	input := NewServiceInput{modelGetter: GetModelFromDB}
	svc := NewService(input)
	_ = svc.modelGetter("test_id")

}

type DBTable struct {
}

type service struct {
	modelGetter func(id string) interface{}
}

type NewServiceInput struct {
	modelGetter func(id string) interface{}
}

func NewService(input NewServiceInput) *service {
	return &service{modelGetter: input.modelGetter}
}

func GetModelFromDB(id string) interface{} {
	return nil
}
