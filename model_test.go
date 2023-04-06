package catboost_test

import (
	"flag"
	"fmt"
	"log"

	"github.com/mandelag/catboost-serving/catboost"
)

func main() {
	var modelPath string
	flag.StringVar(&modelPath, "m", "ranking-model", "path to model file")
	flag.Parse()

	ref, err := catboost.LoadModel(modelPath)
	if err != nil {
		log.Panicln(err)
	}
	defer ref.Close()

	ref.SetPredictionType(catboost.PredictionTypeProbability)

	result, err := ref.Predict([]float32{1.5, 4, 5, 6, 9.5, 18, 1, 8, 33, 50, 65, 60}, 4, []string{"cat1", "cat2", "cat3"}, 1, 3)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Result:", result)

	info, err := ref.Info()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Info: %+v\n", info)
}
