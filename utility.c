#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include "utility.h"


float* FloatFeatures(size_t float_feature_size) {
    float* features = (float*) malloc(float_feature_size * sizeof(float));
    return features;
}

// // FloatFeatures allocates memory for float features
// float* FloatFeatures(size_t float_feature_size, size_t doc_count) {
//     float* features = (float*) malloc(float_feature_size * doc_count * sizeof(float));
//     // printf("%p\n", (void *) &features[0]); // DEBUG
//     return features;
// }

// Write to a feature
void FloatFeaturesWrite(float* features, size_t index, float value) {
    *(features + index) = value;
}

void FreeFloatFeatures(float* float_features) {
    free(float_features);
}

// Get float feature to write
float* GetFloatFeature(float* features, size_t float_feature_size, size_t doc_index) {
    return features + doc_index * float_feature_size;
}

// FloatFeatures allocates memory for float features
float** FloatFeaturesCollection(float* features, size_t float_feature_size, size_t doc_count) {
    float** collection = (float**) malloc(doc_count * sizeof(float*));
    for (size_t i = 0; i < doc_count; i++) {
        *(collection + i) = features + i * float_feature_size;         //  GetFloatFeature(features, float_feature_size, i);
        // printf("%p\n", (void *) features + i * float_feature_size); // DEBUG
    }
    return collection;
}

void FloatFeaturesDebug(float* features, size_t float_feature_size, size_t doc_count) {
    printf("---float features---\n");
    for (size_t i = 0; i < doc_count * float_feature_size; i++) {
        printf("%f ", *(features + i) );
    }
    printf("\n");
    printf("---\n");
}

char** CategoricalFeatures(size_t categorical_feature_size) {
    char** features = (char**) malloc(categorical_feature_size * sizeof(char*));
    // printf("%s\n", (void *) &features[0]); // DEBUG
    return features;
}

void FreeCategoricalFeatures(char** categorical_features) {
    free(categorical_features);
}

// Write to a feature
void CategoricalFeaturesWrite(char** features, size_t index, char* value) {
    *(features + index) = value;
}

// GetCategoricalFeature feature to write
char** GetCategoricalFeature(char** features, size_t categorical_feature_size, size_t doc_index) {
    return features + doc_index * categorical_feature_size;
}


// CategoricalFeaturesCollection allocates memory for float features
char*** CategoricalFeaturesCollection(char** features, size_t categorical_feature_size, size_t doc_count) {
    char*** collection = (char***) malloc(doc_count * sizeof(char**));
    for (size_t i = 0; i < doc_count; i++) {
        *(collection + i) = features + i * categorical_feature_size;         //  GetCategoricalFeature(features, categorical_feature_size, i);
        // printf("%p\n", (void *) features + i * categorical_feature_size); // DEBUG
    }
    return collection;
}

char*** AllocateStringArrayRef() {
    char*** collection = (char***) malloc(1 * sizeof(char**));
    return collection;
}

void FreeStringArrayRef(char*** ref) {
    free(ref);
}


void CategoricalFeaturesDebug(char** features, size_t categorical_feature_size, size_t doc_count) {
    printf("---categorical features---\n");
    for (size_t i = 0; i < doc_count * categorical_feature_size; i++) {
        printf("%s ", *(features + i) );
    }
    printf("\n");
    printf("---\n");
}

double GetResult(double* result, size_t index) {
    // printf("%p\n", (void *) result + index); // DEBUG
    return *(result + index);
}

double* Result(size_t size) {
    double* features = (double*) malloc(size * sizeof(double));
    return features;
}

void FreeResult(double* result) {
    free(result);
}