package math

import "errors"

type Calculator struct{}

func (calculator Calculator) Sum(value1, value2 float64) float64{
	return value1 + value2
}

func (calculator Calculator) Division(value1, value2 float64) (float64, error){
	if value2 == 0{
		return 0, errors.New("Divison by Zero")
	}

	return value1 / value2, nil

}