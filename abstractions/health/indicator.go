package health

type Indicator interface {
	Health() ComponentStatus
}
