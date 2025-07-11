# product-se
product-se



## Running Application on Docker Container

This application you can run on docker container.

### Prerequisite :
* Docker
* OS Linux Docker Image, 
* Golang version 1.20 or latest

### Environment Variable available


### build docker image
```bash
$ docker build -t {IMAGE_NAME} -f deployment/dockerfile .
# example
$ docker build -t product-se -f deployment/dockerfile .
```

### run docker container after build the image
```bash
# example run http serve:
$ docker run -i --name product-se -p 8081:8081 -t product-se

```


## Running on your local machine

Linux or MacOS

## Installation guide
#### 1. install go version 1.17
```bash
# please read this link installation guide of go
# https://go.dev/doc/install
```

#### 2. Create directory workspace    
```bash
run command below: 
mkdir $HOME/go
mkdir $HOME/go/src
mkdir $HOME/go/pkg
mkdir $HOME/go/bin
mkdir -p $HOME/go/src/product-se
chmod -R 775 $HOME/go
cd $HOME/go/src/product-se
export GOPATH=$HOME/go
```    
```bash
# edit bash profile in $HOME/.bash_profile        
# add below to new line in file .bash_profile         
    PATH=$PATH:$HOME/bin:$HOME/go/bin
    export PATH  
    export GOPATH=$HOME/go 
# run command :
source $HOME/.bash_profile
```

#### 3. Build the application    
```bash
# run command :
cd $HOME/go/src/product-se
git clone -b development https://product-se .
cd $HOME/go/src/product-se
go mod tidy && go mod download && go mod vendor
cp config/app.yaml.dist config/app.yaml
# edit config/app.yaml with server environtment
go build

# run application after build or create on supervisord 
./product-se http
```


### 4. Health check Route PATH
```bash
{{host}}/liveness
```


#### Postman Collection
```go
```

### Database Migration
migration up
```bash
go run main.go db:migrate up
```

migration down
```bash
go run main.go db:migrate down
```

migration reset
```bash
go run main.go db:migrate reset
```

migration reset
```bash
go run main.go db:migrate reset
```

migration redo
```bash
go run main.go db:migrate redo
```

migration status
```bash
go run main.go db:migrate status
```

create migration table
```bash
go run main.go db:migrate create {table-name} sql

# example
go run main.go db:migrate create users sql
```

to show all command
```bash
go run main.go db:migrate
```

## run docker compose on your local machine
```bash
docker-compose -f deployment/docker-compose.yaml --project-directory . up -d --build
```

## run docker on your local machine with life reload when modify all file *.go or modify config app.yaml
```bash
docker build -t product-se -f deployment/dockerfile-local .
docker run -i --name product-se -p 8081:8081 -v $PWD:/go/src/product-se/product-se product-se
```

### generate automatic CRUD
#### generate entity, presentation, repositories, use case, dto
make sure already have database and table
```bash
go run main.go gen:all {tableName}


# example
go run main.go gen:entity users
```

#### generate entity only
make sure already have database and table
```bash
go run main.go gen:entity {tableName}

# example
go run main.go gen:entity users
```

#### generate use case only
no needed database
```bash
go run main.go gen:logic {packageName} {fileName}

# example
go run main.go gen:entity users user_create
```