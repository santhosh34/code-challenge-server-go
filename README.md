# Pragmateam code challenge server (GO)

## Framework & languages
This project uses
* Golang 1.20
* Optimized the code by reducing number of requests to backend from UI
* Instead of React Front End calling golang backend for each product
  Backend takes care of parallel processing of getting temperature for all products
 using go routines by hitting aws lambda service and consolidate the result and send to front end
  This also gives the flexibility of adding new containers in future in a new admin module to add more containers or beer types.
* The business logic related to weather the temperate is within the range business logic also moved to backend for optimal performance and parallel execution. 
  This gives us better performance as we used go concurrency model using go channels for all types of beers

* 80 % + test coverage for core logic i.e excluding main and external mocking calls which can be covered in e2e tests.

### Run project

- `go run main.go` - Start the server (http://localhost:8081)
- `go test ./... --cover -v` - To run the tests with coverage 


Request : http://localhost:8081/products

Response: 

```
    [
        {
            "id": "4",
            "name": "Stout",
            "temperature": 5,
            "minTemperature": 6,
            "maxTemperature": 8,
            "tempRangeStatus": "Temperature is outside the correct range -  Its lower than desired"
        },
        {
            "id": "6",
            "name": "Pale Ale",
            "temperature": -2,
            "minTemperature": 4,
            "maxTemperature": 6,
            "tempRangeStatus": "Temperature is outside the correct range -  Its lower than desired"
        },
        {
            "id": "3",
            "name": "Lager",
            "temperature": 0,
            "minTemperature": 4,
            "maxTemperature": 7,
            "tempRangeStatus": "Temperature is outside the correct range -  Its lower than desired"
        },
        {
            "id": "2",
            "name": "IPA",
            "temperature": 0,
            "minTemperature": 5,
            "maxTemperature": 6,
            "tempRangeStatus": "Temperature is outside the correct range -  Its lower than desired"
        },
        {
            "id": "1",
            "name": "Pilsner",
            "temperature": 11,
            "minTemperature": 4,
            "maxTemperature": 6,
            "tempRangeStatus": "Temperature is outside the correct range -  Its higher than desired"
        },
        {
            "id": "5",
            "name": "Witbier",
            "temperature": 2,
            "minTemperature": 3,
            "maxTemperature": 5,
            "tempRangeStatus": "Temperature is outside the correct range -  Its lower than desired"
        }
    ]

```

```
Coverage Report:
$ go test ./... --cover
?   	server	[no test files]
ok  	server/handlers	(cached)	coverage: 80.0% of statements
ok  	server/repository	(cached)	coverage: 82.4% of statements
```
