# MyCatsApiChallenge

**Before we begin, keep in mind this project was developed in MacOS for Docker usage. It may have different behavior or may not work on Windows or \*Unix like systems**

# SRE Challenge - The Cat API challenge
Develop an application in the programming language of your preference to collect the following information from The Cat API (https://thecatapi.com/):
- For each one of the available breeds, store the origin, temperament, and description in a database;
- For each one of the available breeds from the topic above, store 3 images URLs in a database;
- Store 3 images URLs of cats wearing hats;
- Store 3 images URLs of cats wearing glasses.

### Why GoLang?
Python was a possibility but GoLang has been gaining more and more space in my heart and I've enjoying to learn GoLang.

### Why MySQL?
Familiarity.

## Usage
Before you deploy this application, be sure to configure it first. There are two configuration files:
- ./GoLang/api_configuration.toml
- ./RestApi/api_configuration.toml

Update these two files with the following pattern:
```
Apikey = "MY_API_KEY"
```
The key must be enclosed by double-quotes.

Great, now that the API key is configured, just run
```
docker-compose up --build --detach
```
to build, run, and detach the application STDOUT from your terminal.

## The Cats
The cats_requests.go application will use the configured API key to request all the needed information from The Cat API, will treat only the desired fields, and store it on MySQL DB. 

You can access the MySQL server by running (you might need mysql-client installed on the host server)
```
docker exec -it mysql_db mysql -uroot -p
```
**You can find the password on the docker-compose file**

The data are being stored in the **cats_api** database in the tables:
- **cats_breeds** (It stores the breed id, breed name, temperament, origin, and description)
- **cats_images** (It stores the images URLs)
- **stylish_cats_images** (it stores the images URLs of cats with style (hats/glasses))

***If you don't run docker-compose down to destroy the DB, the application won't force delete the breeds table if it already has data!!***

## API
You can access the API through localhost:8000.

```
/breeds
```
- It returns all the information from all the breeds available on The Cat API;

```
/breeds?id={breed_id}
```
- It returns all the breeds that the name matches the breed_id (For example **/breeds?id=brit* will return both British Longhair and British Shorthair breeds);

```
/temperament/{temperament}
```
- It returns all the breeds that have the temperament requested;

```
/origin/{origin}
```
- It returns all the breeds from the requested origin.

## TODO
- Architecture 
- Logging
- Dashboards