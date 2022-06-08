# tomapi

## A simple web API example written in Go.


## Implementations

### Go
Simply install go 1.17.
### Database

#### Installation
brew install sqllite sqlite-utils

### Create Example DB
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
   1  Tom           Sasser
   2  Foo           Bar
```
## API

### Endpoints

####/names

| HTTP Op | Description |
| --- | --- |
| GET | Get a list of all names, returned in JSON format |
| POST | Add a name from request data sent as JSON |

#### /names/:first_name

| HTTP Op | Description |
| --- | --- |
| GET | Get a name by first name, returning data in JSON format |

#### /names/:id

| HTTP Op | Description |
| --- | --- |
| PATCH | Edit a name by database ID, first_name and/or last_name, submitted in JSON format |
| DELETE | Delete a name by database ID |

## Examples

### Get all names:
```
curl --location --request GET 'http://localhost:8080/names'
```

### Add a new name:
```
curl --location --request POST 'http://localhost:8080/names' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Baz",
    "last_name": "Boo"

}'
```

### Get name by searching for first name:
```
curl --location --request GET 'http://localhost:8080/names/Foo'
```

### Update a name by record ID.  Do a GET first to find the ID.
```
curl --location --request PATCH 'http://localhost:8080/names/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name":  "Baz",
    "last_name": "Boo"
}'
```

### Delete a name by record ID.  Do a GET first to find the ID.
```
curl --location --request DELETE 'http://localhost:8080/names/2'
```