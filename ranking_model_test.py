from catboost import CatBoostRegressor, Pool

model_name = "ranking-model"

model_from_file = CatBoostRegressor()
model_from_file.load_model(model_name)

test_data = [
             ["cat1", 1.5, 4, 5, 6],
             ["cat2", 9.5, 18, 1, 8],
             ["cat3", 33, 50, 65, 60],
        ]

eval_pool = Pool(
    data = test_data,
    # group_id = ["a", "a", "a"],
    cat_features = [0]
)

# # # Get predictions
preds = model_from_file.predict(eval_pool, prediction_type="Probability")

print(preds)
print(model_from_file.feature_names_)
