===Endpoint===
Index  : GET    http://localhost:8080/v1/orders
Show   : GET    http://localhost:8080/v1/orders/{id}
Create : POST   http://localhost:8080/v1/orders
Update : PUT    http://localhost:8080/v1/orders/{id}
Delete : DELETE http://localhost:8080/v1/orders/{id}


===Request Body===
{
  "orderedAt": "2019-11-09T21:21:46+00:00",
  "customerName": "Tomas & friends 15",
  "items": [
    {
      "itemCode": "123",
      "description": "IPhone 10X",
      "quantity": 1
    }
  ]
}