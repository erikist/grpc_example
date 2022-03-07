### Usage
Run the `main.go` file 

### Call the method: 
```
grpcurl -plaintext -d '{"message":"parrot this"}' localhost:9000 grpc_example.echo.Echo/Echo  
```