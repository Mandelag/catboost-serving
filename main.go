package main

// #include <model_calcer_wrapper.h>
// #include <utility.h>
// #include <stdlib.h>
// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L/home/keenan/catboost/lib -lcatboostmodel
import "C"

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	model, err := C.ModelCalcerCreate()
	if err != nil {
		log.Panicln("create model handle", err)
	}

	_, err = C.LoadFullModelFromFile(model, C.CString("model-regressor"))
	if err != nil {
		log.Panicln("create model handle", err)
	}

	floatFeatureSize := C.ulong(4)
	docSize := C.ulong(2)
	features := C.FloatFeatures(floatFeatureSize, docSize)

	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 0, 2)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 1, 4)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 2, 8)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 0), 3, 9)

	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 0, 1)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 1, 4)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 2, 50)
	C.FloatFeaturesWrite(C.GetFloatFeature(features, floatFeatureSize, 1), 3, 60)

	C.FloatFeaturesDebug(features, floatFeatureSize, docSize)

	featuresCollection := C.FloatFeaturesCollection(features, floatFeatureSize, docSize)

	categoricalFeatureSize := C.ulong(1)
	categoricalFeatures := C.CategoricalFeatures(categoricalFeatureSize, docSize)
	a := C.CString("a")
	b := C.CString("b")
	C.CategoricalFeaturesWrite(C.GetCategoricalFeature(categoricalFeatures, categoricalFeatureSize, 0), 0, a) // TODO: free the string as well
	C.CategoricalFeaturesWrite(C.GetCategoricalFeature(categoricalFeatures, categoricalFeatureSize, 1), 0, b) // TODO: free the string as well

	C.CategoricalFeaturesDebug(categoricalFeatures, categoricalFeatureSize, docSize)

	categoricalFeaturesCollection := C.CategoricalFeaturesCollection(categoricalFeatures, categoricalFeatureSize, docSize)

	resultSize := docSize // since not a multi class model
	result := C.Result(resultSize)

	C.CalcModelPrediction(model, docSize, featuresCollection, floatFeatureSize, categoricalFeaturesCollection, categoricalFeatureSize, result, resultSize)

	v1 := float64(C.GetResult(result, 0))
	v2 := float64(C.GetResult(result, 1))

	fmt.Println(v1, v2)

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
