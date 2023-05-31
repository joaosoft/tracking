# geo-location
[![Build Status](https://travis-ci.org/joaosoft/geo-location.svg?branch=master)](https://travis-ci.org/joaosoft/geo-location) | [![codecov](https://codecov.io/gh/joaosoft/geo-location/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/geo-location) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/geo-location)](https://goreportcard.com/report/github.com/joaosoft/geo-location) | [![GoDoc](https://godoc.org/github.com/joaosoft/geo-location?status.svg)](https://godoc.org/github.com/joaosoft/geo-location)

A simple location finder.

## Support for 
> Search
> Reverse

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/geo-location
```

## Usage 
This examples are available in the project at [geo-location/examples](https://github.com/joaosoft/geo-location/tree/master/examples)

```go
var geo, _ = geolocation.NewGeoLocation()

func main() {

	// document create
	fmt.Println(":: SEARCH BY: STREET")
	search("rua particular de monsanto")

	fmt.Println(":: REVERSE BY: LATITUDE/LONGITUDE")
	reverse(41.1718238, -8.6186277)
}

func search(street string) {
	result, err := geo.NewSearch().
		Street(street).
		Search()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nsearch by street %s: %s\n", street, result[0].Name)
	}
}

func reverse(latitude float64, longitude float64) {
	result, err := geo.NewSearch().
		Latitude(latitude).
		Longitude(longitude).
		Reverse()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nsearch by latitude %f, longitude: %f: %s\n", latitude, longitude, result[0].Name)
	}
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
