
### https://feik.herokuapp.com/v1/about

# FEIK - is a "fake-data" rest api based for free usage

Feik is based on the [go-faker](https://pkg.go.dev/github.com/bxcodec/faker) package in order to provide a plain REST interface that can be used for development/testing/etc when some "non-sense" data is needed. 

All endpoints so far:


Users: 
```
GET /v1/users/fake
GET /v1/users/list
GET /v1/users/find?email=email@example.com
POST /v1/users/create
GET /v1/users/delete?email=email@example.com
```

Persons: 
```
GET /v1/persons/fake
GET /v1/persons/list
POST /v1/persons/create
```

Work. In. Progress......
