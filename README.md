# tracking
[![Build Status](https://travis-ci.org/joaosoft/tracking.svg?branch=master)](https://travis-ci.org/joaosoft/tracking) | [![codecov](https://codecov.io/gh/joaosoft/tracking/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/tracking) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/tracking)](https://goreportcard.com/report/github.com/joaosoft/tracking) | [![GoDoc](https://godoc.org/github.com/joaosoft/tracking?status.svg)](https://godoc.org/github.com/joaosoft/tracking)

A simple tracking and counting tool for site events. This is to be used for auditing.

## Support for 
> Http

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

## Save tracking

> Saving location by street

Method: ```POST``` 
Route: ```http://localhost:8001/api/v1/tracking/event```
Body:
```
{
	"category": "category",
	"action": "action",
	"label": "label",
	"value": 1,
	"street": "rua particular de monsanto",
	"meta_data": {
        "teste_1": "teste",
        "teste_2": 1,
        "teste_3": 1.1
    }
}
```

> Saving location by latitude, longitude

Method: ```POST``` 
Route: ```http://localhost:8001/api/v1/tracking/event```
Body:
```
{
	"category": "category",
	"action": "action",
	"label": "label",
	"value": 1,
	"latitude": 41.1718238,
	"longitude": -8.6186277,
	"meta_data": {
        "teste_1": "teste",
        "teste_2": 1,
        "teste_3": 1.1
    }
}
```

>### Go
```
go get github.com/joaosoft/tracking
```

## Usage 
This examples are available in the project at [tracking/examples](https://github.com/joaosoft/tracking/tree/master/examples)

```go
func main() {
	m, err := tracking.NewTracking()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
