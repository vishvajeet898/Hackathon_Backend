# Hackathon_backend

#Features
- [x] Signup
- [x] Login
- [x] Add Basic Info
- [x] Get Basic Info
- [x] Update Basic Info
- [x] Add Measurements
- [x] Get Measurements
- [ ] Update Measurements
- [ ] Add Visits
- [ ] Update Visits
- [ ] Generate Health Card


api endpoints

# Auth

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

----

# Basic Info

Add users Basic Info : POST https://hackathonbackend-production.up.railway.app/api/v1/user/basicInfo/
```json
Request 
Auth type : Breare token (JWT)
{
    "age":21,
    "height":"5'12",
    "sex":"M",
    "blood_group":"A+",
    "weight":72,
    "allergies":["pollen","rubber"]
}

Response

{
    "data": {
        "basic_id": "e212783e-ccd0-4a37-8c48-7117f3372f12",
        "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
        "age": 21,
        "height": "5'12",
        "sex": "M",
        "blood_group": "A+",
        "weight": 72,
        "allergies": [
            "pollen",
            "rubber"
        ]
    },
    "internal_response_code": 0,
    "message": "Success"
}
```

Get users Basic Info : Get https://hackathonbackend-production.up.railway.app/api/v1/user/basicInfo/
```json
Auth type : Breare token (JWT)

Response
{
    "data": {
        "basic_id": "e212783e-ccd0-4a37-8c48-7117f3372f12",
        "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
        "age": 21,
        "height": "5'12",
        "sex": "M",
        "blood_group": "A+",
        "weight": 72,
        "allergies": [
            "pollen",
            "rubber"
        ],
        "diseases": [
            "ED",
            "Insomania"
        ]
    },
    "internal_response_code": 0,
    "message": "Success"
}
```


Edit users Basic Info : PUT https://hackathonbackend-production.up.railway.app/api/v1/user/basicInfo/
```json
Request 
Auth type : Breare token (JWT)
//Chages
{
    "age":21,
    "height":"5'12",
    "sex":"M",
    "blood_group":"A+",
    "weight":72,
    "allergies":["pollen","rubber","Peanuts","sugar"]
}

Response

{
    "data": {
        "basic_id": "e212783e-ccd0-4a37-8c48-7117f3372f12",
        "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
        "age": 21,
        "height": "5'12",
        "sex": "M",
        "blood_group": "A+",
        "weight": 72,
        "allergies": [
            "pollen",
            "rubber",
            "Peanuts",
            "sugar"
        ]
    },
    "internal_response_code": 0,
    "message": "Success"
}
```

----

# Measurements


Add user's Measurement(Temp, glucose) : POST https://hackathonbackend-production.up.railway.app/api/v1/user/measurement
```json
Request 
Auth type : Breare token (JWT)
{
    "type" : "Temperature",
    "x_value" : "31",
    "y_value" : "",
    "name" :"Temperature",
    "taken_by" :"Self"
}
```


Get user's Measurement(Temp, glucose) : GET https://hackathonbackend-production.up.railway.app/api/v1/user/measurement
```json
Request 
Auth type : Breare token (JWT)
Response

{
    "data": {
        "Results": [
            {
                "measurement_id": "faf0da34-f099-4624-9b25-9ca22987631a",
                "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
                "type": "Temperature",
                "x_value": "32",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-05T02:29:13.696988+05:30",
                "taken_by": "Self"
            },
            {
                "measurement_id": "5e0dc58a-4161-49a2-971a-c17de6cf11a8",
                "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
                "type": "Temperature",
                "x_value": "48",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-05T02:29:37.738282+05:30",
                "taken_by": "Self"
            },
            {
                "measurement_id": "2cc91f4d-1e2a-41e7-9e34-afc9237fcfee",
                "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
                "type": "Temperature",
                "x_value": "52",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-05T02:29:44.1121+05:30",
                "taken_by": "Self"
            },
            {
                "measurement_id": "ff374801-fbdd-438d-ac9b-2d7fb47aa825",
                "user_id": "0f47607c-a818-404d-9b30-4a633c8318cb",
                "type": "Temperature",
                "x_value": "31",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-05T02:29:54.434255+05:30",
                "taken_by": "Self"
            }
        ],
        "Page": 1,
        "Limit": 20
    },
    "internal_response_code": 0,
    "message": "Success"
}
```


Get All of user's Measurement from date to curr date (Temp, glucose) : GET https://hackathonbackend-production.up.railway.app/api/v1/user/allMeasurement/2022-12-05
```json
Request 
Auth type : Breare token (JWT)
Response
{
    "data": {
        "Results": [
            {
                "measurement_id": "e435b1ad-0b65-4ee0-b944-f6975b3e1fc8",
                "user_id": "94d3bb8e-24ac-4bea-9279-4f1beae5ca55",
                "type": "Temperature",
                "x_value": "31",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-06T19:12:53.66439Z",
                "taken_by": "Self"
            },
            {
                "measurement_id": "8bfea080-ceb0-446b-a231-622a0a4186cc",
                "user_id": "94d3bb8e-24ac-4bea-9279-4f1beae5ca55",
                "type": "Temperature",
                "x_value": "21",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-06T19:13:02.067682Z",
                "taken_by": "Self"
            },
            {
                "measurement_id": "40d299ae-90fb-40a3-a486-482aef9b4f1d",
                "user_id": "94d3bb8e-24ac-4bea-9279-4f1beae5ca55",
                "type": "Temperature",
                "x_value": "51",
                "y_value": "",
                "name": "Temperature",
                "created_at": "2022-12-06T19:13:05.57253Z",
                "taken_by": "Self"
            },
            {
                "measurement_id": "964c23e9-0abf-4f7e-a310-adbff51040a9",
                "user_id": "94d3bb8e-24ac-4bea-9279-4f1beae5ca55",
                "type": "Gluco",
                "x_value": "12",
                "y_value": "",
                "name": "Glucose",
                "created_at": "2022-12-06T19:13:19.215627Z",
                "taken_by": "Self"
            },
            {
                "measurement_id": "343c8c9b-17ca-487e-9651-7cc8aedb0285",
                "user_id": "94d3bb8e-24ac-4bea-9279-4f1beae5ca55",
                "type": "Gluco",
                "x_value": "122",
                "y_value": "",
                "name": "Glucose",
                "created_at": "2022-12-06T19:13:22.091894Z",
                "taken_by": "Self"
            },
            {
                "measurement_id": "f99f328d-8d96-48ec-9581-208530331ac1",
                "user_id": "94d3bb8e-24ac-4bea-9279-4f1beae5ca55",
                "type": "Gluco",
                "x_value": "52",
                "y_value": "",
                "name": "Glucose",
                "created_at": "2022-12-06T19:13:28.225408Z",
                "taken_by": "Self"
            }
        ],
        "Page": 1,
        "Limit": 100
    },
    "internal_response_code": 0,
    "message": "Success"
}
```

