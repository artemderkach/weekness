# loadTest
cli for testing http services. loadTest will create tons of requests for given URL

## run
run main `$ go run main.go --help`  
or install and then run  
`$ go install`  
`$ loadtest --help`  

## help
help output
```
$ loadtest --help

Usage:
  main [OPTIONS] URL

Application Options:
  -p, --protocol= http or https protocol (default: http)
  -m, --method=   http method (default: GET)
  -d, --duration= requests duration (in seconds) (default: 3)
  -r, --rps=      requests per second (default: 100)

Help Options:
  -h, --help      Show this help message
```
