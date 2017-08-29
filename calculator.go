package math

type Calculator struct{}

func (calculator Calculator) Sum(value1, value2 float64) float64{
	return value1 + value2
}