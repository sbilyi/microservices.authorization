```
## create user 
$ curl -X POST -d '{"name":"sergii", "username":"kartanis", "email":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signup

## login
$ curl -X POST -d '{"usernameOrEmail":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signin

## load token
$ curl -s -X POST -d '{"usernameOrEmail":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signin | python3 -c "import sys, json; print('Bearer ', json.load(sys.stdin)['accessToken'])"

## current user info
$ curl -X GET \
    -H "Content-Type: application/json"\
    -H "Authorization: Bearer  eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNTY1MTM1NTM5LCJleHAiOjE1NjU3NDAzMzl9.2wPwR-VTzyowADERZQJrmeg4vuLzT0DL7xZiW5_5lOwm1rTBT4Gzo82nsJ6cT2Pkj-64wQVlbtUL0SyFv6pkuA" \
    localhost:8080/api/user/me


``` 

curl -s -X POST -d '{"usernameOrEmail":"kartanisoid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signin | python3 -c "import sys, json; print('Bearer ', json.load(sys.stdin)['accessToken'])"

curl -X GET \
    -H "Content-Type: application/json"\
    -H "Authorization: Bearer  eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNTY1MTM1MTA5LCJleHAiOjE1NjU3Mzk5MDl9.DpHDYG8OVZhXk32Y5LmFskPaGOVjyxqPyumeBG3LI3bCz7vjHEQCEFyksMK4D_Hc6LItStZwIzofH95HLPS7SA" \
    localhost:8080/api/user/me
    
    curl -X POST -d '{"name":"sergii", "username":"karta2is", "email":"kartani2oid@gmail.com", "password": "secret123"}' -H "Content-Type: application/json" localhost:8080/api/auth/signup