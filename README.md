# Hackathon_backend

api endpoints

signUp : POST
https://hackathonbackend-production.up.railway.app/api/v1/user/SignUp/
 
```json
Sign Up json paylod
{
    "name" : "Vishwajeet",
    "email":"Vishvajeet878@gmail.com",
    "password":"password",
    "phone":"8707405904",
    "user_type":"normal"
}

response 

{
    "data": {
        "user_id": "191f99c6-9ee0-4740-8a53-43134dbda273",
        "name": "Vishwajeet",
        "email": "Vishvajeet878@gmail.com",
        "password": "$2a$10$qFTgZQ8Fhk6HKQjbiK5.cuCL48bN7b8qBfLRRuKc4x9CNA60CkMbu",
        "phone": "8707405904",
        "user_type": "normal"
    },
    "internal_response_code": 0,
    "message": "Success"
}

```

login : POST https://hackathonbackend-production.up.railway.app/api/v1/user/login/
```json
Login Payload
{
    "email":"Vishvajeet878@gmail.com",
    "password":"password"
}

Response

{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxOTFmOTljNi05ZWUwLTQ3NDAtOGE1My00MzEzNGRiZGEyNzMiLCJ1c2VyVHlwZSI6Im5vcm1hbCIsImV4cCI6MTY3MDE2NTUyOH0.fdEuqdJO6NVbXI2NSRqDaV-sEH50-fQhGb2FHTgp2B4",
    "internal_response_code": 0,
    "message": "Success"
}
```
