# Hotel_booking_api
Hotel booking using REST API

# Run 
Navigate to directory where `main.go` is and run `go run .` command.
Which will run the server on `localhost:8080`.

# Endpoints
1. /tables - GET request will return all the tables, POST request with new table information will add new table.
2. /tables/:id - GET request with table ID will give that table information.
3. /available_tables - GET request will return available tables to book.
4. /booked_tables - GET request will return booked tables.
5. /book/:id - POST request with person information will book the table supplied in the Url parameters.

