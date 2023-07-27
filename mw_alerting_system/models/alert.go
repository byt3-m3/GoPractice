package models

type NewAlertOpt func(alert *Alert)

type AlertType string

var (
	AlertTypePrometheus AlertType = "prometheus"
	AlertTypeGCP        AlertType = "gcp"
	AlertTypeOpsGenie   AlertType = "ops_genie"

	AlertTypeAll AlertType = "all"
)

func WithAlertName(name string) NewAlertOpt {

	return func(alert *Alert) {
		alert.Name = name

	}
}

func WithAlertExpr(expr string) NewAlertOpt {

	return func(alert *Alert) {
		alert.Expr = expr

	}
}

func WithID(id string) NewAlertOpt {
	return func(alert *Alert) {
		alert.ID = id
	}
}

func WithTestData() NewAlertOpt {
	return func(alert *Alert) {
		alert.Name = "test-alert"
		alert.Expr = "test-Expr"
		alert.ID = "test-ID"
	}
}

type Alert struct {
	ID        string
	Name      string
	Expr      string
	AlertType AlertType
}

func NewPrometheusAlert(opts ...NewAlertOpt) *Alert {
	alert := &Alert{}
	alert.AlertType = AlertTypePrometheus
	applyOpts(alert, opts...)

	return alert

}

func NewGCPAlert(opts ...NewAlertOpt) *Alert {
	alert := &Alert{}
	alert.AlertType = AlertTypeGCP
	applyOpts(alert, opts...)

	return alert

}

func NewOpsGenieAlert(opts ...NewAlertOpt) *Alert {
	alert := &Alert{}
	alert.AlertType = AlertTypeOpsGenie
	applyOpts(alert, opts...)

	return alert

}

func applyOpts(alert *Alert, opts ...NewAlertOpt) {
	for _, opt := range opts {
		opt(alert)
	}
}
