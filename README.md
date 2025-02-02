# arctic-wolf-risk-manager
Welcome to Arctic Wolf Risk Manager 
Arctic Wolf Risk manager is a Web based Server that can create and keep tracks of risks at the organiztion.

To Run the Server Please Download the go modules first by running the below command
```
go mod tidy
```

Once All your Modules are downloaded the code can be Run using the command 
```
go run cmd\main.go
```

Arctic Wolf Risk manager serves 2 endpoints

### Create Risk

Create Risk Endpoint helps to create new risks and stores them in memory with Mutex locks to ensure concurrent access do not cause in consistency in the data

```
curl --location --request POST 'http://localhost:8080/v1/risks' \
--header 'Content-Type: application/json' \
--data '{
    "state":"investigating",
    "title": "random title",
    "description":"random description"
}'
```
State of a Risk can only be in [open, closed, accepted, investigating] and is a mandatory parameter. Once a risk is created the Risk is associated with a unique id , which can be used further to enquire about a risk

### Get Risk

Create Risk Endpoint helps to create new risks and stores them in memory with Mutex locks to ensure concurrent access do not cause in consistency in the data

#### Get Entire Risk List
```
curl --location --request GET 'http://localhost:8080/v1/risks'
```

#### Get a Risk by Id
```
curl --location --request GET 'http://localhost:8080/v1/risks/{id}'
```
To fetch a Risk Details by its id please replace `{id}` with the id returned after risk creation

