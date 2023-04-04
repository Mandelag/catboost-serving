package main

// #include <model_calcer_wrapper.h>
// #include <utility.h>
// #include <stdlib.h>
// #include <stdio.h>
//
// void debug(char*  str) {
//    printf("%s\n", str);
// }
//
//
// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/lib -lcatboostmodel
import "C"

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var modelPath string
	flag.StringVar(&modelPath, "m", "ranking-model", "path to model file")
	flag.Parse()

	ref, err := LoadModel(modelPath)
	if err != nil {
		log.Panicln(err)
	}
	defer ref.Close()

	ref.SetPredictionType(PredictionTypeProbability)

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
