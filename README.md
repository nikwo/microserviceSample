#Microservice sample packed into docker container
##Getting started
```bash
docker build -t application .
docker container run --name application -dit -p 8080:8090 application
```
After success start:
```bash
curl http://localhost:8080/api/v1/greeting
#response is {"message": "Hello"}
```