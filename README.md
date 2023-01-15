# CATBOOST-SERVING

This project allows you to serve catboost models in Go via gRPC endpoints.


#### Docker Usage
```
1. make docker
2. docker run -p 9090:9090 --rm catboost-serving ./catboost-serving

If you want to use a different port, update the command as

docker run -p <port>:<port> --rm catboost-serving ./catboost-serving -port 9091
```

#### Build Executable
```
1. git clone https://github.com/catboost/catboost.git

2. make build. The output files are generated in build folder.

3. if you are using m1, make build-m1

You will be asked to enter your user password.
```