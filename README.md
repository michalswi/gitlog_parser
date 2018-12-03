
## GIT LOG PARSER
Written in golang is [here](./go/go_parse.go).  
Written in python is [here](./py/py_parse.py). TODO

Two ways to run:  
**1. 'git log' already added to some local file for example 'logHUGE.log'**
```sh
$ go run go/go_parse.go
go run go/go_parse.go | jq
```

**2. point to .git in specific location as a first parameter**
```sh
$ go run go/go_parse.go /<full_path>/.git
go run go/go_parse.go /<full_path>/.git | jq
```

TODO:
- display logs as a json in webui
- save output to some database