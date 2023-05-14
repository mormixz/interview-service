# Interview Service

Service using Go and MongoDB

## Get started
```bash
docker-compose up -d
```

## APIs
### Get Interview  
- api for get interview using interview_id
```bash
GET http://{{host}}/interview/{{interview_id}} HTTP/1.1
Content-Type: application/json
```
### Get List Interview
-  api for get list interview with querystring
    * limit - using for limit number of list interview 
    * status - "In Progress", "To Do", "Done"
```bash
GET http://{{host}}/interview/all?limit={{limit}}&status{{status}} HTTP/1.1
Content-Type: application/json
```

### Update Interview Comment
- api for update interview comment
```bash
PUT http://{{host}}/interview/{{interview_id}}/comment HTTP/1.1
Content-Type: application/json
{
   "message":"comment",
   "created_by":{
      "id":"645f2b47e0931117713e74ca",
      "name":"โรบินฮู้ด",
      "email":"user1@robinhood.co.th"
   }
}
```
### Update Interview Detail
- api for update interview status
```bash
PUT http://{{host}}/interview/{{interview_id}}/update HTTP/1.1
Content-Type: application/json

{
   "id":"645f4882bcc2918b5e4605b5",
   "description":"description",
   "status":"In Progress",
   "created_by":{
      "id":"645f2b47e0931117713e74ca",
      "name":"โรบินฮู้ด",
      "email":"user1@robinhood.co.th"
   },
   "created_at":"2023-01-04T10:00:00.000Z"
}
```

## Library
* [Fiber](https://github.com/gofiber/fiber) - Go web framework
* [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - The MongoDB supported driver for Go
* [Viper](https://github.com/spf13/viper) - For Read Configuration
* [Testify](https://github.com/stretchr/testify) - For Testing

