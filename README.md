# GOSPLAT preprocessor

The preprocessor for GOSPLAT's model. It parses code made in the language GoLang and forms a JSON file that the training script can parse and extract tokens from. 

An example JSON file may look something like this:
```
{"database GetUser"}
{"server Handler"}
...
```
In this example, the the two lines represented are two different packages, `database` and `server`, with ties to the function-headers `GetUser` and `Handler`. This is done to train the model on relationship between words, in this example, package to function relationships. 

The preprocessor works for multiple kinds of models, not only FastText used in the implementation seen in GOSPLAT.
