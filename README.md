# Basic Go Microservices

This project consists of two simple microservices written in Go. Each service runs on a separate port and provides a JSON response when accessed.

## Features
- **Service A**: Runs on port `8080` and returns a JSON message.
- **Service B**: Runs on port `8081` and returns a JSON message.
- Uses `net/http` for handling requests.
- Simple, lightweight, and easy to extend.

## Installation & Running

### Prerequisites
- Install [Go](https://go.dev/dl/)

### Steps to Run
```sh
# Clone the repository
git clone https://github.com/mfahadqureshi786/go-microservice
cd gomicroservice

# Run the microservices
go run main.go
```

### Access the Services
- **Service A**: [http://localhost:8080/service-a](http://localhost:8080/service-a)
- **Service B**: [http://localhost:8081/service-b](http://localhost:8081/service-b)

## Testing the Services
You can test the endpoints using `curl` or on Browser:
```sh
curl http://localhost:8080/service-a
curl http://localhost:8081/service-b
```

### Expected Response
For **Service A**:
```json
{
  "message": "Hello from Service A!"
}
```

For **Service B**:
```json
{
  "message": "Hello from Service B!"
}
```

## License
This project is open-source and available under the MIT License.

