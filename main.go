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
// #cgo LDFLAGS: -L/home/keenan/catboost/lib -lcatboostmodel
import "C"

import (
	"flag"
	"fmt"
	"log"
	"unsafe"
)

func main() {
	var modelPath string
	flag.StringVar(&modelPath, "m", "model", "path to model file")
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
}

func main2() {
	var modelPath string
	flag.StringVar(&modelPath, "m", "model", "path to model file")
	flag.Parse()
	model, err := C.ModelCalcerCreate()
	if err != nil {
		log.Panicln("create model handle gg", err)
	}

	_path := C.CString(modelPath)
	C.debug(_path)
	ok, err := C.LoadFullModelFromFile(model, _path)
	if !ok {
		log.Panicln("its not ok", ok, err)
	}

	ref, err := LoadModel(modelPath)
	if err != nil {
		log.Panicln(err)
	}
	defer ref.Close()
	// time.Sleep(1 * time.Minute)

	floatFeatureSize := C.ulong(4 * 2) // featureSize * docSize
	docSize := C.ulong(2)
	features := C.FloatFeatures(floatFeatureSize)

	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 0, 2)

	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 1, 4)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 2, 8)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 3, 9)

	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 0, 1)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 1, 4)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 2, 50)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 3, 60)
	// ref, err = LoadModel(modelPath)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	C.FloatFeaturesDebug(features, C.ulong(4), docSize)

	featuresCollection := C.FloatFeaturesCollection(features, floatFeatureSize, docSize)

	categoricalFeatureSize := C.ulong(1 * 2)
	categoricalFeatures := C.CategoricalFeatures(categoricalFeatureSize)
	a := C.CString("a")
	b := C.CString("b")
	C.CategoricalFeaturesWrite(C.GetCategoricalFeature(categoricalFeatures, categoricalFeatureSize, 0), 0, a) // TODO: free the string as well
	C.CategoricalFeaturesWrite(C.GetCategoricalFeature(categoricalFeatures, categoricalFeatureSize, 1), 0, b) // TODO: free the string as well

	C.CategoricalFeaturesDebug(categoricalFeatures, C.ulong(1), docSize)

	categoricalFeaturesCollection := C.CategoricalFeaturesCollection(categoricalFeatures, categoricalFeatureSize, docSize)

	resultSize := docSize // since not a multi class model
	result := C.Result(resultSize)

	C.CalcModelPrediction(model, docSize, featuresCollection, floatFeatureSize, categoricalFeaturesCollection, categoricalFeatureSize, result, resultSize)
	v1 := float64(C.GetResult(result, 0))
	v2 := float64(C.GetResult(result, 1))
	v3 := float64(C.GetResult(result, 2))
	v4 := float64(C.GetResult(result, 3))
	v5 := float64(C.GetResult(result, 4))
	v6 := float64(C.GetResult(result, 5))
	fmt.Println("raw", v1, v2, v3, v4, v5, v6)

	C.SetPredictionType(model, C.APT_PROBABILITY)
	C.CalcModelPrediction(model, docSize, featuresCollection, floatFeatureSize, categoricalFeaturesCollection, categoricalFeatureSize, result, resultSize)

	v1 = float64(C.GetResult(result, 0))
	v2 = float64(C.GetResult(result, 1))
	v3 = float64(C.GetResult(result, 2))
	v4 = float64(C.GetResult(result, 3))
	v5 = float64(C.GetResult(result, 4))
	v6 = float64(C.GetResult(result, 5))
	fmt.Println("prob", v1, v2, v3, v4, v5, v6)

	C.FreeFloatFeatures(features)
	C.FreeCategoricalFeatures(categoricalFeatures)
	C.FreeResult(result)
	C.free(unsafe.Pointer(a))
	C.free(unsafe.Pointer(b))

}

type SinglePredictor struct {
	modelRef unsafe.Pointer
}

func predictSingleWrapper(modelRef unsafe.Pointer, floatFeatures []float32, categoricalFeatures []string) {

}
