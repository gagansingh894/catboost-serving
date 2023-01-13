package modelmanager

import (
	"errors"
	"fmt"
	"github.com/gagansingh894/catboost-serving/go/internal/catboost"
	"os"
	"strings"
)

type ModelManager interface {
	Initialize() error
	GetPredictions(modelName string, floats [][]float32, cats [][]string) ([]float64, error)
}

type modelManager struct {
	artefactsPath string
	models        map[string]catboost.Model
}

func NewModelManager(artefactsPath string) ModelManager {
	return &modelManager{
		artefactsPath: artefactsPath,
		models:        make(map[string]catboost.Model),
	}
}

func (m *modelManager) Initialize() error {
	var model catboost.Model

	files, err := os.ReadDir(m.artefactsPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %s", err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "reg_") {
			model, err = catboost.NewRegressor(m.artefactsPath + "/" + file.Name())
			if err != nil {
				fmt.Printf("failed to load model: %s", err)
			}
		}
		if strings.HasPrefix(file.Name(), "bc_") {
			model, err = catboost.NewClassifier(m.artefactsPath + "/" + file.Name())
			if err != nil {
				fmt.Printf("failed to load model: %s", err)
			}
		}

		modelName := strings.TrimPrefix(strings.TrimSuffix(file.Name(), ".cbm"), "reg_")
		m.models[modelName] = model
	}
	fmt.Println("number of models loaded: ", len(m.models))
	return nil
}

func (m *modelManager) GetPredictions(modelName string, floats [][]float32, cats [][]string) ([]float64, error) {
	model, ok := m.models[modelName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("failed to find model: %s", modelName))
	}

	preds, err := model.Predict(floats, cats) // todo: make it more dynamic. Need to train a model with category features and then test
	if err != nil {
		fmt.Printf("failed to predicti: %s", err)
	}

	return preds, nil
}
