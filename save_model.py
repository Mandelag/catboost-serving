from catboost import CatBoostRegressor
# Initialize data

train_data = [["a", 1, 4, 5, 6],
              ["a", 4.1, 5, 6, 7],
              ["b", 30, 40, 50, 60]]

train_labels = [10, 20, 30]

# Initialize CatBoostRegressor
model = CatBoostRegressor(iterations=2,
                          learning_rate=1,
                          cat_features=[0],
                          depth=2)
# Fit model
model.fit(train_data, train_labels)

eval_data = [["a", 2, 4, 6, 8],
             ["b", 1, 4, 50, 60]]

# Get predictions
preds = model.predict(eval_data, prediction_type="Probability")
print(preds)


model.save_model("model-regressor")
