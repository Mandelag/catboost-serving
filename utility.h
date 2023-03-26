#pragma once

#include <utility.h>

// FloatFeatures allocates memory to create a collection of float features (feature for all doc)
float* FloatFeatures(size_t float_feature_size, size_t doc_count);

// FreeFloatFeatures deallocates memory for float features
void FreeFloatFeatures(float* float_features);

// GetFloatFeature get specific float feature of a doc
float* GetFloatFeature(float* features, size_t float_feature_size, size_t doc_index);

// FloatFeatures get references for all of the float features 
float** FloatFeaturesCollection(float* features, size_t float_feature_size, size_t doc_count);

// FloatFeaturesDebug print out float features
void FloatFeaturesDebug(float* features, size_t float_feature_size, size_t doc_count);

// FloatFeaturesWrite write to a particular feature
void FloatFeaturesWrite(float* features, size_t index, float value);


// CategoricalFeatures allocates memory for categorical features
char** CategoricalFeatures(size_t categorical_feature_size, size_t doc_count);

// FreeCategoricalFeatures deallocates memory for float features
void FreeCategoricalFeatures(char** categorical_features);


// Write to a feature
void CategoricalFeaturesWrite(char** features, size_t index, char* value);

// GetCategoricalFeature feature to write
char** GetCategoricalFeature(char** features, size_t categorical_feature_size, size_t doc_index);

// CategoricalFeaturesCollection allocates memory for float features
char*** CategoricalFeaturesCollection(char** features, size_t categorical_feature_size, size_t doc_count);

void CategoricalFeaturesDebug(char** features, size_t categorical_feature_size, size_t doc_count);


// Result allocates memory to store prediction result 
double* Result(size_t size);

// FreeResult deallocates result
void FreeResult(double* float_features);

// GetResult is the getter for individual prediction result
double GetResult(double* result, size_t index);

