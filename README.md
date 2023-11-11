A basic Event Management System backend build with go pramming language. Is has a very few usecases and 3 modules. 

1. Events module.
2. Workshops module.
3. Reservations module.


Usecase_

![UseCase](static/app-usecase.png)

1. There can have multiple events. 
2. A event can have multiple workshops. 
3. A workshop can have multiple reservations. 


Database Schema_

![!DBSchema](static/db-diagram.png)



EndPoints_

1. Get all events

```{domain}/api/v1/events/list?perPage={PerPageIndicentCount}&currentPage={CurrentPageNumber}```

Example:

```localhost:8081/api/v1/events/list?perPage=10&currentPage=1```


Request Method: GET

Dummy Response_




2. Get Event Details

```{domain}/api/v1/events/details?eventId={EventID}```

Example:
```localhost:8081/api/v1/events/details?eventId=1```

Dummy Response_


3. Get Workshop List by Event

```{domain}/api/v1/workshops/list?eventId={EventID}```

Example:
```localhost:8081/api/v1/workshops/list?eventId=1```



4, Get Workshop Details

```{domain}/api/v1/workshops/details?workshopId={WorkshopID}```


Example:
```localhost:8081/api/v1/workshops/details?workshopId=13```



5. Create Reservation in Workshops


```{domain}/api/v1/reservations/```

Example
```localhost:8081/api/v1/reservations/```


Payload:
{
    "name": "mahin",
    "email": "mahifffn@gmail.com",
    "workshop_id": 13
}


Response_


