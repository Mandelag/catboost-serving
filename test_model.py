from catboost import CatBoostRegressor, Pool

model_from_file = CatBoostRegressor()
model_from_file.load_model("model-regressor")


eval_data = [["a", 2, 4, 6, 8],
             ["b", 1, 4, 50, 60]]

# preds_class = model_from_file.predict(eval_data)
# print(preds_class)

# Get predicted probabilities for each class
# preds_proba = model_from_file.predict_proba(eval_data)
# print(preds_proba)

# Get predicted RawFormulaVal
preds_raw = model_from_file.predict(eval_data, prediction_type='RawFormulaVal')
print(preds_raw)