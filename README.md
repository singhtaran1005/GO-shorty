# Go - shorty
A Golang implementation of URL shortener just like : [bitly](https://bitly.com/)

## Tech stack 💻
![goland-image](https://github.com/singhtaran1005/GO-shorty/assets/53126276/a481b986-b83a-416e-a4f2-2564b095ca05)
![go-fiber](https://github.com/singhtaran1005/GO-shorty/assets/53126276/979530c2-1b23-4093-ad57-ab903fa28026)
![redis](https://github.com/singhtaran1005/GO-shorty/assets/53126276/28417e8d-709b-452a-a859-e2199524c2c8)
![docker](https://github.com/singhtaran1005/GO-shorty/assets/53126276/3e3f1840-e895-478f-bc3a-0bfe978712a9)

## Features included 📋

Rate limiting concept added along with expiry option to maintain service reliability.

Certain checks to ensure correct URL/domain creation.

Integrated Redis for scalable and persistent data storage.

Dockerized the whole application using Docker and Docker-compose

## Project Setup 📎

- Clone the repository using `git clone <repoURL>`

- Go to the project working directory.

- RUN `go mod tidy` to get all Golang dependencies

- Install docker and docker-compose -> [Install](https://docs.docker.com/engine/install/)
(Check docs for help during installation)

- Create a `.env` file in the main working directory
  - Example `.env` file:

```
APP_PORT:(add your port)
DB_ADDR:(add your database port), used Redis mostly locally should be `:6379`
DB_PASS:(add your database password)
DOMAIN:(add your domain where the project will be hosted along with the port)
API_QUOTA:(set to 10 by default, you can choose some other number)
```
- Run `docker-compose up -d` to spin docker containers for both Go-fiber and Redis.

- Eventually, you can also set up a makefile to shorthand docker commands

- YOUR Project is up NOW !!

- To test it call API on [Postman](https://www.postman.com)
- POST request
    - (`localhost:3000/api/v1`) 
```json
{
    "url": "URL_TO_SHORT"
    "short": "CUSTOM_URL_ID"
}
```

  where you will get a response :
```json
{
    "url": "URL_TO_SHORT",
    "short": "SHORTEN_URL",
    "expiry": "SET IT TO A NUMBER say 30",
    "rate_limit": "{how many more number of times API can be called, which in our case will be 9}"
    "rate_limit_reset": "After how much time rate will reset"
}
```

- GET request
  - (`localhost:3000/:SHORTEN_URL`) will redirect to the original URL.
