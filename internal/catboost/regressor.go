package catboost

/*
#cgo linux LDFLAGS: -lcatboostmodel
#cgo darwin LDFLAGS: -lcatboostmodel
#include <stdlib.h>
#include <stdbool.h>
#include <model_calcer_wrapper.h>
*/
import "C"
import (
	"fmt"
)

// regressor is a wrapper over ModelCalcerHandler. It is primarily used to represent CBRegressor
type regressor struct {
	model
}

// NewRegressor loads model from file and returns Model interface
func NewRegressor(filename string) (Model, error) {
	m := &regressor{}
	m.Handler = C.ModelCalcerCreate()
	if !C.LoadFullModelFromFile(m.Handler, C.CString(filename)) {
		return nil, getError()
	}
	return m, nil
}

// Predict returns raw predictions for specified data points
func (r *regressor) Predict(floats [][]float32, floatLength int, cats [][]string, catLength int) ([]float64, error) {
	preds, err := r.model.Predict(floats, floatLength, cats, catLength)
	if err != nil {
		return nil, fmt.Errorf("failed to get predictions for regressor: %w", err)
	}

	return preds, nil
}
