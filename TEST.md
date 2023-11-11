# Testing APIs

### Get all events

```sh
curl -X GET "http://localhost:8081/api/v1/events/list?perPage=10&currentPage=1"
```


### Get Event Details

```sh
curl -X GET "http://localhost:8081/api/v1/events/details?eventId=1"
```


### Get Workshop List by Event

```sh
curl -X GET "http://localhost:8081/api/v1/workshops/list?eventId=5"
```


### Get Workshop Details

```sh
curl -X GET "http://localhost:8081/api/v1/workshops/details?workshopId=3"
```


### Create Reservation in Workshops

```sh
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "test",
    "email": "test@gmail.com",
    "workshop_id": 2
}' http://localhost:8081/api/v1/reservations/
```
