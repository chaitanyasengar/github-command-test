1. For the first time, user hits "/login" Api to get token, the parameters expected are :

    Request Type : POST
    Request Body : 

    {
        username : admin,
        password : admin
    }

    Response :

    {
        "code": 200,
        "expire": "2019-07-11T23:42:00Z",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI4ODg1MjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2Mjg4NDkyMH0.OcNyPukHj4XyB0VYjMYopdBOG8H2iZObKzE_96HYT_w"
    }

2. User saves this token, by default token is valid for an hour, if token gets expired, user gets an invalid token message in response : 

    If token is invalid, then response will be 

    {
        "code": 401,
        "message": "signature is invalid"
    }

    If token gets expired, then response will be

    {
        "code": 401,
        "message": "token is expired"
    }

    Once any user gets this error response, user needs to call "/auth/refresh_token" Api
    Request Type : POST
    
    Request Headers : 

    {
        AUTH_KEY : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI4ODg1MjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2Mjg4NDkyMH0.OcNyPukHj4XyB0VYjMYopdBOG8H2iZObKzE_96HYT_w
    }

    Success response of this Api will be :

    {
        "code": 200,
        "expire": "2019-07-12T00:54:32Z",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI4OTI4NzIsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2Mjg4OTI3Mn0.4vVvl66HMceL0vE5JvBCltYztx9N0nMW7GehEBOZdc8"
    }

3. To access any route, every request needs to pass token in the request Header, along with the default parameters that any Api requires

    {
        AUTH_KEY : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI4ODg1MjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2Mjg4NDkyMH0.OcNyPukHj4XyB0VYjMYopdBOG8H2iZObKzE_96HYT_w
    }

WIP : Token Flag for Local and Staging Environment(as discussed over the call)
Blockage : None