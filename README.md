# ERPLY — Go Challenge

Objective is to create a simple Go (golang) based API server that is able to read and write data from database. Data structure that is being read and written can be chosen freely. The project should fill following requirements:

- Read/Write capability over API
- Simple authentication of requests
- Use database for storage
- JSON - format for communication
- Includes simple documentation (README.md)
- Include unit-tests

Preferences:

- Project is delivered to us via Git (Zip is alternative)
- As efficient API communication as possible

## Architecture Notes

- Using [cixtor/middleware](https://github.com/cixtor/middleware) as an alternative to Go’s DefaultServeMux. This project is lightweight, and contains basic methods to write a web API service very fast
- SQLite was chosen as the database engine for its flexibility and portability. However, keep in mind that this is just a “toy project”, and a production version would require a more robust database like PostgreSQL
- Data structure describes a contact in a phonebook, including fields like: first name, last name, phone number, address, and email
- An additional endpoint allows you to export one or more contacts in [vCard 3.0](https://en.wikipedia.org/wiki/VCard) format, hinting the possibility of adding support for photos

## Authentication

API authentication is based on [HTTP Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication). The username and password are stored in the database. However, the password field is not really a user password but an API key in UUID format, for uniqueness. The project includes two API keys for testing, they are defined in the `migration.sql` file, and are stored in plain text in the database. Because “security” is out of the scope of this project, the keys are not salted nor hashed.

Here is an example request:

```
GET /contact?id=2 HTTP/1.1
User-Agent: paw
Authorization: Basic am9obkBleGFtcGxlLmNvbTo4NUVDNDk2Qi03RUM0LTQ0NzgtQjI3Qi05NEIzODFCNDAzMEY=
Host: localhost:3000
Connection: close
```

Here is an example of a response:

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 23 Mar 2019 06:39:54 GMT
Content-Length: 160
Connection: close

{
    "ok":true,
    "data": {
        "id":2,
        "firstname":"Alice",
        "lastname":"Smith",
        "phone":"6045551234",
        "address":"350 W Georgia St, Vancouver, BC",
        "email":"alice@example.com"
    }
}
```

## Unit Tests

All the endpoints are accompanied of their corresponding unit-test.


```
go test -v
```
