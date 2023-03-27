from catboost import CatBoost, Pool, CatBoostRanker

# Initialize data

x_train = [["cat1", 1.5, 4, 5, 6],
              ["cat2", 4, 5, 6, 7],
              ["cat3", 9.5, 18, 1, 8],
              ["cat3", 30, 45.5, 50, 59.5],
              ["cat5", 33, 50, 65, 90],
              ["cat1", 45, 30, 29.3, 77.1],
              ["cat2", 100, 120, 150, 190],
              ["cat3", 150, 220, 150, 200]]

queries_train = ["a", "a", "a", "a", "a", "b", "b", "b"]
y_train = [1, 0, 1, 1, 1, 0, 0, 1]

train_pool = Pool(
    data=x_train,
    label=y_train,
    group_id=queries_train,
    cat_features=[0]
)

# Initialize CatBoostRanker
model = CatBoost(
    {"iterations": 2000, "verbose": False, "loss_function": 'QueryRMSE'}
)

# Fit model
model.fit(train_pool)

eval = [
             ["cat1", 1.5, 4, 5, 6],
             ["cat2", 9.5, 18, 1, 8],
             ["cat3", 33, 50, 65, 90],
        ]

eval_pool = Pool(
    data = eval,
    group_id = ["a", "a", "a"],
    cat_features = [0]
)

# # # Get predictions
preds = model.predict(eval_pool, prediction_type="Probability")
print(preds)


model.save_model("model")
