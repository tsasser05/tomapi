# tomapi

Basic plan:

Database
1) Create database and prepare for simple use   DONE (sqlite3)
2) Put into container
3) Run in Docker
4) Automate creation of Docker instance

Webform
1) Create simple webform to submit data using API
2) Create simple webform to query data using API
3) Have it run in Docker
4) Automate creation of Docker instance

API
1) Create simple API
2) POST for webform submit data to database
3) GET for webform query data from database
4) Run in Docker
5) Automate creation of Docker instance

Kafka
1) Create kafka instance
2) Somehow get the webform to interact with it
3) Run in Docker
4) Automate creation of Docker instance


h2. Database

h3. Installation
brew install sqllite sqlite-utils

h3. Create Example DB
sqlite3 names.db
sqlite> CREATE TABLE [names] (
   ...> id INTEGER NOT NULL PRIMARY KEY,
   ...> first_name TEXT NOT NULL,
   ...> last_name TEXT NOT NULL
   ...> );

sqlite> insert into names values
   ...> (NULL, "Tom","Sasser");
sqlite> insert into names values
   ...> (NULL, "Foo","Bar");

h3. Verify
sqlite-utils names.db "select * from names" --table
  id  first_name    last_name
----  ------------  -----------
   1  Tom           Sasser
   2  Foo           Bar






