
### GIT LOG PARSER
Written in golang is [here](./go/go_parse.go).  
Written in python is [here](./py/py_parse.py). TODO

#### RUN
There are two ways to run it:
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

#### OUTPUT
After script successfully run it displays json to stdout and runs simple webui with log contents.
```sh
$ curl http://localhost:5000/api/v1/log
```

TODO:
- save output to some database