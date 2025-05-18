# User CRUD API with Go and MongoDB

![Image](https://github.com/user-attachments/assets/6d17e4e9-8943-47ac-9a13-3e99cd13d9e9)


## Project Overview

This project implements a simple CRUD REST API to manage `TravelUser` data using:

- Go (Golang)
- MongoDB (version 8.0)
- Docker & Docker Compose for containerization

---

## How to Run

Run this project with Docker Compose:

```bash
docker-compose up --build

```

MongoDB runs on port  ``27017``
Go API server runs on port ``8080``

Access the API at
```
http://localhost:8080
```

## End points

#### Get all users (GET)
```
curl http://localhost:8080/travelusers
```

#### Get a TravelUser by ID (GET)
```
curl http://localhost:8080/travelusers/<id>
```

#### Create a TravelUser (POST)
```
curl -X POST http://localhost:8080/travelusers \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "isEmailVerified": false,
    "contactNumber": "1234567890",
    "isContactNumberVerified": false,
    "address": "123 Main St",
    "country": "USA",
    "age": 30
  }'

```

#### Update a User by ID (PUT)

```
curl -X PUT http://localhost:8080/travelusers/<id> \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johnupdated",
    "email": "johnupdated@example.com",
    "isEmailVerified": true,
    "contactNumber": "0987654321",
    "isContactNumberVerified": true,
    "address": "456 Another St",
    "country": "USA",
    "age": 31
  }'

```

#### Delete a User by ID (DELETE)

```
curl -X DELETE http://localhost:8080/travelusers/<id>
```


## Notes

- MongoDB data is persisted in Docker volume mongodb_data.

- API logs its connection status on startup.

- Use the ID from the create response for fetching, updating, or deleting users.



