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
	"math"
)

// classifier is a wrapper over ModelCalcerHandler. It is primarily used to represent CBClassifier
type classifier struct {
	model
}

// NewClassifier loads model from file and returns Model interface
func NewClassifier(filename string) (Model, error) {
	m := &classifier{}
	m.Handler = C.ModelCalcerCreate()
	if !C.LoadFullModelFromFile(m.Handler, C.CString(filename)) {
		return nil, getError()
	}
	return m, nil
}

func sigmoid(probit float64) float64 {
	return 1.0 / (1.0 + math.Exp(-probit))
}

// Predict returns raw predictions for specified data points
func (c *classifier) Predict(floats [][]float32, floatLength int, cats [][]string, catLength int) ([]float64, error) {
	preds, err := c.model.Predict(floats, floatLength, cats, catLength)
	if err != nil {
		return nil, fmt.Errorf("failed to get predictions for classifier: %w", err)
	}

	for i := range preds {
		preds[i] = sigmoid(preds[i])
	}

	return preds, nil
}
