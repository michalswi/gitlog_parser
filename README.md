
### GIT LOG PARSER
Written in golang is [here](./go/go_parse.go).  
Written in python is [here](./py/py_parse.py). TODO

#### RUN
There are two ways to run it:  

**`git log output` was added earlier local file `./gitlog.log*`**:
```sh
$ go run go/go_parse.go
```
\* `gitlog.log` by default is hardcoded in go script

**provide full path to the directory where `.git` is located as a first parameter**:
```sh
$ go run go/go_parse.go <full_path>
# for example
go run go/go_parse.go /tmp/myrepo
```

#### OUTPUT
After script successfully run it will display json to stdout and run simple webui with git log content.
```sh
$ curl http://localhost:5000/api/v1/log | jq
```

TODO:
- save output to some database