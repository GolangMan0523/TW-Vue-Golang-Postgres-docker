# Installation - Running locally 💻

Here you will understand how to run and setup the development environment for twitterclone

> NOTE: To run this app locally you need to have an AWS S3 Bucket available so that uploading images will work. (Might add Cloudinary and on-disk implementation in the future)

## 1. Using docker-compose

***Prerequisites***
- [Docker](https://docker.com/)
- [AWS Access Key](https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html) and [S3 Bucket](https://aws.amazon.com/s3/)

Create .env file in `configs/` directory by copying `configs/.env.example` and setup the environment variables

- Running the containers with `docker-compose up -d`
- To stop the containers `docker-compose stop`

Login into the postgres service container's terminal and set a password for the postgres user that matches the DB_PASSWORD env variable. You can do this by accessing docker's terminal on the db service, then running:

- `psql -U postgres` 
- `ALTER USER postgres WITH PASSWORD 'postgres';`

Restart your `rest` docker's service.

If the `web` service is not working simply run the vite server directly:

Navigate to `/web`

- Create an `.env` file based on the `.env.example` file located at `/web/.env.example`
- Run `yarn` to install the dependencies
- Run `yarn dev` to run the frontend

NOTE: Proven to work with node version: `14.18.1` use `nvm` if you need to install this version directly. [NVM](https://github.com/nvm-sh/nvm)

## 2. Manually

***Prerequisites***
- [Golang](golang.org)
- [PostgreSQL](postgresql.org)
- [Redis](redis.io)
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [AWS Access Key](https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html) and [S3 Bucket](https://aws.amazon.com/s3/)
- [air (optional hot-reload)](https://github.com/cosmtrek/air)

###  Local Rest API Server Development

#### PostgreSQL
Install PostgreSQL:
- **macOS**: Run `brew install postgresql`.
- **Windows**: Follow [this](https://www.postgresqltutorial.com/install-postgresql/) guide.
- **Linux**: Follow [this](https://www.postgresqltutorial.com/install-postgresql-linux/) guide.


Create a database named `twitterclone`

```sh
$ psql

postgres=# CREATE DATABASE twitterclone;
```

Install golang-migrate [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Run the migrations.

```sh
$ ./scripts/run_migrations.sh Your_PostgreSQL_URL

# Example PostgreSQL url: postgres://postgres:postgres@127.0.0.1:5432/twitterclone\?sslmode=disable
```

#### Redis
Install Redis:
- **macOS**: Run `brew install redis`.
- **Windows**: Follow [this](https://redis.io/download#installation) guide.
- **Linux**: Follow [this](https://redis.io/download#installation) guide.

#### Golang
Install Golang:
- **macOS**: Run `brew install golang`.
- **Windows**: Follow [this](https://golang.org/dl/) guide.
- **Linux**: Follow [this](https://golang.org/dl/) guide.


Navigate to `/configs` and set the following environment variables:

```
APP_NAME=twitter-clone
APP_HOST=127.0.0.1
APP_PORT=8000
APP_DOMAIN=localhost
DEBUG=false

ACCESS_TOKEN_SECRET=
REFRESH_TOKEN_SECRET=
# Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
# 15m -> 15 minutes
ACCESS_TOKEN_DURATION=15m
# 168h -> 7 days
REFRESH_TOKEN_DURATION=168h

DEFAULT_AVATAR_URL=

# 2.5 mb in bytes
MAX_UPLOAD_SIZE=2621440

# For attaching images in tweets
# 32 mb in bytes
MAX_TWEET_ATTACHMENT_SIZE=33554432

DB_USER=postgres
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=twitterclone

REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=
REDIS_KEY_DELIMITER=::

AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_S3_BUCKET_NAME=

# [Optional] can leave empty here
AWS_SESSION_TOKEN=
```

**Running the server**

```sh
$ go run cmd/rest/main.go --prefork

# Using hot-reloading (Optional must have air installed)
$ make dev
```

### Local Frontend Development

Navigate to `/web`

- Create an `.env` file based on the `.env.example` file located at `/web/.env.example`
- Run `yarn` to install the dependencies
- Run `yarn dev` to run the frontend


# Technology Stack 🛠

- [Golang](golang.org)
- [Fiber HTTP framework](https://github.com/gofiber/fiber)
- [PostgreSQL](postgresql.org)
- [Redis](redis.io)
- [NodeJS](https://nodejs.org/en/)
- [TypeScript](https://www.typescriptlang.org/)
- [migrate](https://github.com/golang-migrate/migrate)
- [air](https://github.com/cosmtrek/air)
- [Amazon Web Service S3](https://aws.amazon.com/s3/)
- [Vue 3](https://v3.vuejs.org/)
- [Vite 2.0](https://vitejs.dev/)
- [Vuex 4](https://next.vuex.vuejs.org)
- [Vue Router 4](https://next.router.vuejs.org)
- [TailwindCSS](http://tailwindcs.com/)

# Resources & references used

- https://github.com/shuber/postgres-twitter
- [Build a twitter clone using vue.js and tailwind css! (by: this.stephie)](https://www.youtube.com/watch?v=bQU-jPyQJ4A)
- [Is SELECT * Expensive? (by: Hussein Nasser)](https://www.youtube.com/watch?v=QQVNVOneZNg)
- [SELECT COUNT (*) can impact your Backend Application Performance, here is why (by: Hussein Nasser)](https://www.youtube.com/watch?v=8xKS7QQKgzk)
- [Full Text Search PostgreSQL (by: Ben Awad)](https://www.youtube.com/watch?v=szfUbzsKvtE)
- https://www.postgresql.org/message-id/20050810133157.GA46247@winnie.fuhr.org
- https://dev.to/shubhadip/vue-3-vuex-4-modules-typescript-2i2o
- https://dev.to/3vilarthas/vuex-typescript-m4j
- [Today i learned golang live reload for development using docker compose air (by: Iman Tumorang)](https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076)

# Features from original project ✨

- Modular Architecture
- Database migration tool using [migrate](https://github.com/golang-migrate/migrate)
- Golang Hot-reloading using [air](https://github.com/cosmtrek/air)
- Supports dark-mode and light-mode with [TailwindCSS](http://tailwindcs.com/)
- Database seeding script using NodeJS
- Authentication using JWT Refresh token flow and Redis for token blacklisting
- Strongly typed Vuex store
- List Tweets feed
- Create Tweets with images
- Retweets
- Reply to Tweets or reply to another reply!
- Like Tweets
- Follow users
- Images & Media uploads stored in AWS S3 Buckets
- Up to 4 images in a single tweet with the same layout as Twitter
- Crop profile image
- Edit Profile Details
- Edit Profile Image
- See who a user is following and see their followers