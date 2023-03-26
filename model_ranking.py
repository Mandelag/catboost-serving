from catboost import CatBoost, Pool, CatBoostRanker

# Initialize data

x_train = [["a", 1.5, 4, 5, 6],
              ["a", 4, 5, 6, 7],
              ["a", 9.5, 18, 1, 8],
              ["b", 30, 45.5, 50, 59.5],
              ["b", 33, 50, 65, 90],
              ["b", 45, 30, 29.3, 77.1],
              ["c", 100, 120, 150, 190],
              ["c", 150, 220, 150, 200],]

queries_train = ["a", "a", "a", "b", "b", "b", "c", "c"]
y_train = [1, 0, 0, 1, 1, 0, 1, 1]

train_pool = Pool(
    data=x_train,
    label=y_train,
    group_id=queries_train,
    cat_features=[0]
)

# Initialize CatBoostRanker
model = CatBoost(
    {"iterations": 2000, "verbose": False, "loss_function": 'QueryRMSE'}
                        #   learning_rate=1,
                        #   cat_features=[0],
                        #   depth=2,
                        #   loss_function='QueryRMSE', 
                        # verbose=False,
                        )
# Fit model
model.fit(train_pool)

eval_data = [
             ["a", 1.5, 4, 5, 6],
             ["b", 9.5, 18, 1, 8],
             ["z", 33, 50, 65, 90],
             ]


# # Get predictions
preds = model.predict(eval_data, prediction_type="Probability")
print(preds)


model.save_model("model")
