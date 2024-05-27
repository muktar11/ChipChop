# ChipChop
Welcome Developers!
This test is designed to assess your comprehension of REST API design, Golang
implementation, hexagonal architecture principles, and adherence to best practices.
Scenario:
Develop a robust REST API in Golang that facilitates price estimations for users. The API
should adhere to the CRUD (Create, Read, Update, Delete) paradigm and leverage a
Postgresql database for persistent storage.
REST Endpoints:
● GET /estimations: Retrieves a paginated list of all price estimations (consider
implementing pagination using query parameters like limit and offset).
● GET /estimations/{id}: Retrieves a specific price estimation by its unique identifier (ID).
● POST /estimations: Creates a new price estimation. The request body should contain
essential details such as:
○ product_name: Name of the product being estimated (string)
○ price: Estimated price (float64)
○ location: User-specified location data (consider using a struct or custom type
to represent location details)
○ user_id: Unique identifier of the user creating the estimation (string or integer,
depending on your user ID format)

● PATCH /estimations/{id}: Updates an existing price estimation. The request body can
include any of the following fields to be updated (partial updates):
○ price
Technical Specifications:
● Hexagonal Architecture: Employ hexagonal architecture to promote loose coupling and
testability. This means separating your application domain logic (core services) from the
infrastructure layer (database access, HTTP routing, etc.) using interfaces. For quick
reference check out this repo (https://github.com/iDevoid/stygis)
● Database: Use Postgresql as the database for persistence. Implement database access
logic using SQLC (https://github.com/sqlc-dev/sqlc) to generate type-safe queries and
avoid manual string concatenation, reducing the risk of SQL injection vulnerabilities. As
of now SQLC still doesn’t support dynamic query so you will need to use additional sql
builder for list API. If possible use goqu(https://github.com/doug-martin/goqu) but you are
free to use whatever you are comfortable with but no ORM allowed 🙂
● Web Framework: Utilize Gin (https://gin-gonic.com/) as the HTTP web framework

● Error Handling: Implement robust error handling throughout the API layers. Return
appropriate HTTP status codes and descriptive error messages to aid in client-side
debugging.
Optional
● Write integration tests to validate the interaction between API endpoints and the
database. Consider using a framework like Ginkgo (https://onsi.github.io/ginkgo/) for
BDD-style testing.
