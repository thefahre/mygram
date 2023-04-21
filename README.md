# mygram
Final Project for Hacktiv8 FGA Kominfo: Golang Scalable Web Microservices. Mentored by Ayat Maulana.

## Requirements 
* Go
* PostgreSQL

## How to run 
1. Create a PostgreSQL WITH at least a password credentials (md5)
2. Clone this repo
3. Fill the following env info needed, either with .env, setting it on Railway if deployed there, or else:
```
DB_HOST=(PostgreSQL host address)
DB_USER=(PostgreSQL username)
DB_PASSWORD=(PostgreSQL password)
DB_PORT=(PostgreSQL port)
DB_NAME=(PostgreSQL database name)
SECRET_KEY=(Secret key for password hashing)
```
4. ```go run main.go```
