
## GIT LOG PARSER
Written in golang is [here](./go/go_parse.go).  
Written in python is [here](./py/py_parse.py).  

### RUN

**GOLANG**

There are two ways to run it:  

**git log output was added earlier to local file `./gitlog.log*`**:
```sh
$ go run go/go_parse.go
```
\* `gitlog.log` by default is hardcoded in go script

**provide full path to the directory where `.git` is located as a first parameter**:
```sh
$ go run go/go_parse.go <path_to_repo>
# for example
go run go/go_parse.go /tmp/myrepo
```

After script successfully run it will display json to stdout and run simple web server with git log content.
```sh
$ curl http://localhost:5000/api/v1/log | jq
```

**PYTHON**  

There is no webserver.  
```sh
$ python py_parse.py -p <path_to_repo> | jq
```


TODO:
- save output to some database