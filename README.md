This is a WIP

# Create the model file

```sh
# assuming you have python3 ready with 'pip install catboost'
python3 save_model.py
```

Test if the model working
```sh
python3 test_model.py
```

They're based on example here [here](https://catboost.ai/en/docs/concepts/python-usages-examples#regression)


# Build the go code

```sh
# of course you need to install golang, for me I use go1.19
go build .
```

To run, don't forget to 

```sh
export LD_LIBRARY_PATH=<./lib folder in this repo>
./catboost-serving
```

so the linker know how to get the catboost static library.
The official distribution is [here](https://github.com/catboost/catboost/releases/tag/v1.1.1)
