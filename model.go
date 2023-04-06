package catboost

// #include <utility.h>
// #include <stdlib.h>
// #include <model_calcer_wrapper.h>
// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/lib -lcatboostmodel
import "C"

import (
	"errors"
	"unsafe"
)

type PredictionType int

const (
	PredictionTypeRawFormulaVal       PredictionType = C.APT_RAW_FORMULA_VAL
	PredictionTypeExponent            PredictionType = C.APT_EXPONENT
	PredictionTypeRMSEWithUncertainty PredictionType = C.APT_RMSE_WITH_UNCERTAINTY
	PredictionTypeProbability         PredictionType = C.APT_PROBABILITY
	PredictionTypeClass               PredictionType = C.APT_CLASS
)

var (
	ErrFailedToLoadModel  = errors.New("failed to load model")
	ErrInvalidFeatureSize = errors.New("invalid feature size")
)

type Model struct {
	modelRef unsafe.Pointer
}

// LoadModel creates load a catboost model by specifying their path
func LoadModel(path string) (model *Model, err error) {
	modelRef, err := C.ModelCalcerCreate()
	if err != nil {
		return nil, err
	}

	_path := C.CString(path)
	defer C.free(unsafe.Pointer(_path))

	ok, err := C.LoadFullModelFromFile(modelRef, _path)
	// somehow even if ok, err may contain value
	// so we only validate on ok
	if !ok {
		return nil, err
	}

	return &Model{
		modelRef: modelRef,
	}, nil
}

// SetPredictionType
func (m *Model) SetPredictionType(t PredictionType) (err error) {
	ok, err := C.SetPredictionType(m.modelRef, C.enum_EApiPredictionType(t))
	if !ok {
		return err
	}
	return nil
}

// Predict single document
func (m *Model) Predict(features []float32, featuresSize int, categoricalFeatures []string, categoricalFeaturesSize int, docSize int) (prediction []float32, err error) {
	_features, close1, err := copyFeatures(features, featuresSize, docSize)
	if err != nil {
		return prediction, err
	}
	defer close1()

	_categoricalFeatures, close2, err := copyCategoricalFeatures(categoricalFeatures, categoricalFeaturesSize, docSize)
	if err != nil {
		return prediction, err
	}
	defer close2()

	prediction, err = m.predict(_features, featuresSize, _categoricalFeatures, categoricalFeaturesSize, docSize)
	if err != nil {
		return prediction, err
	}

	return prediction, nil
}

func (m *Model) predict(features **C.float, featuresSize int, categoricalFeatures ***C.char, categoricalFeaturesSize int, docSize int) (result []float32, err error) {
	resultSize := C.ulong(docSize) // since not a multi class model
	_result := C.Result(resultSize)
	defer C.free(unsafe.Pointer(_result))

	ok, err := C.CalcModelPrediction(m.modelRef, resultSize, features, C.ulong(featuresSize), categoricalFeatures, C.ulong(categoricalFeaturesSize), _result, resultSize)
	if !ok {
		return nil, err
	}

	result = make([]float32, docSize)
	for i := 0; i < docSize; i++ {
		result[i] = float32(C.GetResult(_result, C.ulong(i)))
	}

	return result, nil
}

// Close release resource(s) used by the model
func (m *Model) Close() {
	C.ModelCalcerDelete(m.modelRef)
}

// GetFeatures get the model's feature column names (if any)
func (m *Model) Info() (result ModelInfo, err error) {
	nFloatFeatures := C.GetFloatFeaturesCount(m.modelRef)
	nCatFeatures := C.GetCatFeaturesCount(m.modelRef)
	nTextFeatures := C.GetTextFeaturesCount(m.modelRef)
	nEmbeddingFeatures := C.GetEmbeddingFeaturesCount(m.modelRef)
	nTree := C.GetTreeCount(m.modelRef)
	nDimensions := C.GetDimensionsCount(m.modelRef)
	nPredictionDimensions := C.GetPredictionDimensionsCount(m.modelRef)

	result = ModelInfo{
		NFloatFeatures:        int(nFloatFeatures),
		NCatFeatures:          int(nCatFeatures),
		NTextFeatures:         int(nTextFeatures),
		NEmbeddingFeatures:    int(nEmbeddingFeatures),
		NTree:                 int(nTree),
		NDimensions:           int(nDimensions),
		NPredictionDimensions: int(nPredictionDimensions),
	}

	_resultRef := C.AllocateStringArrayRef()
	var nFeatures C.ulong
	ok, err := C.GetModelUsedFeaturesNames(m.modelRef, _resultRef, &nFeatures)
	if !ok {
		return result, err
	}

	result.NFeatures = int(nFeatures)

	// Allocates, copy, free

	features := unsafe.Slice(*_resultRef, int(nFeatures))
	for i := 0; i < int(nFeatures); i++ {
		result.Features = append(result.Features, C.GoString(features[i]))
		C.free(unsafe.Pointer(features[i]))
	}
	C.FreeStringArrayRef(_resultRef)
	return result, nil
}

type ModelInfo struct {
	NFloatFeatures        int
	NCatFeatures          int
	NTextFeatures         int
	NEmbeddingFeatures    int
	NTree                 int
	NDimensions           int
	NPredictionDimensions int
	Features              []string
	NFeatures             int
}

func copyFeatures(features []float32, featureSize, docSize int) (featuresRef **C.float, close func(), err error) {
	if len(features) != featureSize*docSize {
		return nil, nil, ErrInvalidFeatureSize
	}

	_features := C.FloatFeatures(C.ulong(len(features)))
	for i := 0; i < len(features); i++ {
		C.FloatFeaturesWrite(_features, C.ulong(i), C.float(features[i]))
	}

	_featuresCollection := C.FloatFeaturesCollection(_features, C.ulong(featureSize), C.ulong(docSize))

	close = func() {
		C.free(unsafe.Pointer(_featuresCollection))
		C.free(unsafe.Pointer(_features))
	}

	return _featuresCollection, close, nil
}

// todo: more generic (not only categorical features, but allocate char )
func copyCategoricalFeatures(features []string, featureSize, docSize int) (featuresRef ***C.char, close func(), err error) {
	if len(features) != featureSize*docSize {
		return nil, nil, ErrInvalidFeatureSize
	}

	_freeList := make([]*C.char, len(features))
	_features := C.CategoricalFeatures(C.ulong(len(features)))
	for i := 0; i < len(features); i++ {
		_freeList[i] = C.CString(features[i])
		C.CategoricalFeaturesWrite(_features, C.ulong(i), _freeList[i])
	}

	_featuresCollection := C.CategoricalFeaturesCollection(_features, C.ulong(featureSize), C.ulong(docSize))

	close = func() {
		for i := range _freeList {
			C.free(unsafe.Pointer(_freeList[i]))
		}
		C.free(unsafe.Pointer(_featuresCollection))
	}

	return _featuresCollection, close, nil
}
