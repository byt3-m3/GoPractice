package models

import "log"

type ServiceDetails struct {
	Name             string
	Env              string
	Regions          []string
	PromAlerts       []*Alert
	GCPAlerts        []*Alert
	OpsGenieAlerts   []*Alert
	deployContextMap map[AlertType]DeployContext
}

type DeployContext struct {
	IsComplete bool
	IsSuccess  bool
}

type NewServiceDetailsInput struct {
	Name           string
	Env            string
	Regions        []string
	GCPAlerts      []*Alert
	PromAlerts     []*Alert
	OpsGenieAlerts []*Alert
}

func NewServiceDetails(input *NewServiceDetailsInput) *ServiceDetails {
	m := make(map[AlertType]DeployContext)
	return &ServiceDetails{
		Name:             input.Name,
		Env:              input.Env,
		Regions:          input.Regions,
		PromAlerts:       input.PromAlerts,
		GCPAlerts:        input.GCPAlerts,
		OpsGenieAlerts:   input.OpsGenieAlerts,
		deployContextMap: m,
	}
}

func (s *ServiceDetails) SetSourceContext(source AlertType, deployCtx DeployContext) {

	s.deployContextMap[source] = deployCtx
}

func (s *ServiceDetails) GetSourceContext(source AlertType) *DeployContext {
	deployCtx, ok := s.deployContextMap[source]
	if !ok {
		log.Println("context not found")

		return nil
	}

	return &deployCtx
}

func (s *ServiceDetails) GetDeployContextMap() map[AlertType]DeployContext {
	return s.deployContextMap
}
