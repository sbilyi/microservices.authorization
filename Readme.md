```
## create user 
$ curl -X POST -d '{"name":"sergii", "username":"kartanis", "email":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signup

## login
$ curl -X POST -d '{"usernameOrEmail":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signin

## current user info
$ curl -X GET \
    -H "Content-Type: application/json"\
    -H "Authorization: Bearer eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNTY1MTI3OTU4LCJleHAiOjE1NjU3MzI3NTh9.Pe0RvhLzXByJQlWqESRXe3_86ui3SVTj9w8pfL_0e-rnlUeuOfCOVRWhtfDwATijdsvZaau1tvniyffPE9dXsw" \
    localhost:8080/api/user/me

eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNTY1MTI3OTU4LCJleHAiOjE1NjU3MzI3NTh9.Pe0RvhLzXByJQlWqESRXe3_86ui3SVTj9w8pfL_0e-rnlUeuOfCOVRWhtfDwATijdsvZaau1tvniyffPE9dXsw"
``` 
