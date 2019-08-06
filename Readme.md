```
## create user 
$ curl -X POST -d '{"name":"sergii", "username":"kartanis", "email":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signup

## login
$ curl -X POST -d '{"usernameOrEmail":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signin
``` 
