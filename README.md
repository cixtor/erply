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
- API authentication is based on [HTTP Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication). The username and password are stored in the database. However, the password field is not really a user password but an API key in UUID format, for uniqueness
