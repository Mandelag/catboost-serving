This is a WIP

Currently only supported for linux distribution.

# Create the model file

```sh
# assuming you have python3 ready with 'pip install catboost'
python3 ranking_model_train.py
```

Test if the model working
```sh
python3 ranking_model_test.py
```

They're based on example here [here](https://catboost.ai/en/docs/concepts/python-usages-examples#regression)


# Build the go code

```sh
# of course you need to install golang, for me I use go1.19
go build .
```

Then run:

```sh
export LD_LIBRARY_PATH=lib
LD_LIBRARY_PATH=lib ./catboost-serving -m model-ranking
```

or 

```sh
export LD_LIBRARY_PATH=lib
./catboost-serving -m model-ranking
```

`LD_LIBRARY_PATH` so the linker know how to get the catboost static library.

The official distribution is [here](https://github.com/catboost/catboost/releases/tag/v1.1.1)
Don't forgeet to verify their checksum!
