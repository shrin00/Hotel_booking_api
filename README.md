# Hotel_booking_api
Hotel booking using REST API

# Run 
Navigate to directory where `main.go` is and run `go run .` command.
Which will run the server on `localhost:8080`.

# Endpoints
/tables - GET request will return all the tables, POST request with new table information will add new table.
/tables/:id - GET request with table ID will give that table information.
/available_tables - GET request will return available tables to book.
/booked_tables - GET request will return booked tables.
/book/:id - POST request with person information will book the table supplied in the Url parameters.

