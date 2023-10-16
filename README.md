# tomapi

## A simple web API example written in Go.


## Implementations

### Go
Simply install go 1.17.

### Docker Build

cd to 'tomapi' directory
docker build -t tomapi-dev .
docker run -it --rm -p 8084:8084 -v $PWD:/Users/tom/tomapi tomapi-dev


### Tools
I recommend you use Postman to test the API, though curl commands are included below.

### Database

#### Installation
brew install sqllite sqlite-utils

#### Create Example DB
```
sqlite3 names.db
sqlite> CREATE TABLE [names] (
   ...> id INTEGER NOT NULL PRIMARY KEY,
   ...> first_name TEXT NOT NULL,
   ...> last_name TEXT NOT NULL
   ...> );

sqlite> insert into names values
   ...> (NULL, "Peter","Parker");
sqlite> insert into names values
   ...> (NULL, "Foo","Bar");

h3. Verify
sqlite-utils names.db "select * from names" --table
  id  first_name    last_name
----  ------------  -----------
   1  Peter         Parker
   2  Foo           Bar
```

### Go run
If you want to run the app on the CLI to test it, then run:
go run main.go

## API

### Endpoints

#### http://localhost:8084/names

| HTTP Op | Description |
| --- | --- |
| GET | Get a list of all names, returned in JSON format |
| POST | Add a name from request data sent as JSON |

#### http://localhost:8084/names/:first_name

| HTTP Op | Description |
| --- | --- |
| GET | Get a name by first name, returning data in JSON format |

#### http://localhost:8084/names/:id

| HTTP Op | Description |
| --- | --- |
| PATCH | Edit a name by database ID, first_name and/or last_name, submitted in JSON format |
| DELETE | Delete a name by database ID |

## Examples

### Get all names:
```
curl --location --request GET 'http://localhost:8084/names'
```

### Add a new name:
```
curl --location --request POST 'http://localhost:8084/names' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Baz",
    "last_name": "Boo"

}'
```

### Get name by searching for first name:
```
curl --location --request GET 'http://localhost:8084/names/Foo'
```

### Update a name by record ID.  Do a GET first to find the ID.
Do a GET first to find the ID.
```
curl --location --request PATCH 'http://localhost:8084/names/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name":  "Baz",
    "last_name": "Boo"
}'
```

### Delete a name by record ID.
Do a GET first to find the ID.
```
curl --location --request DELETE 'http://localhost:8084/names/2'
```
