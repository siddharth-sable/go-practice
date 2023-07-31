```
.
├── cache-project
│   ├── go.mod
│   └── main.go
├── ecommerce-go
│   ├── controllers
│   │   ├── address.go
│   │   ├── cart.go
│   │   └── controllers.go
│   ├── database
│   │   ├── cart.go
│   │   └── database-setup.go
│   ├── docker-compose.yml
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middleware
│   │   └── middleware.go
│   ├── models
│   │   └── models.go
│   ├── routes
│   │   └── routes.go
│   └── tokens
│       └── token-gen.go
├── email-checker
│   └── main.go
├── go-bookstore
│   ├── cmd
│   │   └── main
│   │       ├── main.exe
│   │       └── main.go
│   ├── go.mod
│   ├── go.sum
│   └── pkg
│       ├── config
│       │   └── app.go
│       ├── controllers
│       │   └── book-controller.go
│       ├── models
│       │   └── book.go
│       ├── routes
│       │   └── bookstore-routes.go
│       └── utils
│           └── utils.go
├── go-csrf
│   ├── db
│   │   ├── db.go
│   │   └── models
│   │       └── models.go
│   ├── go.mod
│   ├── go.sum
│   ├── keys
│   │   ├── app.rsa
│   │   └── app.rsa.pub
│   ├── main.go
│   ├── randomstrings
│   │   └── randomStrings.go
│   └── server
│       ├── middleware
│       │   ├── middleware.go
│       │   └── myJwt
│       │       └── myJwt.go
│       ├── server.go
│       └── templates
│           ├── templateFiles
│           │   ├── login.tmpl
│           │   ├── register.tmpl
│           │   └── restricted.tmpl
│           └── templates.go
├── go-docker
│   ├── Dockerfile
│   └── main.go
├── go-dynamodb-crud
│   ├── cmd
│   │   └── app
│   │       └── main.go
│   ├── config
│   │   └── config.go
│   ├── internal
│   │   ├── controllers
│   │   ├── entities
│   │   ├── handlers
│   │   ├── repository
│   │   │   ├── adapter
│   │   │   │   └── adapter.go
│   │   │   └── instance
│   │   │       └── instance.go
│   │   ├── routes
│   │   │   ├── config.go
│   │   │   └── routes.go
│   │   └── rules
│   └── utils
│       ├── env
│       ├── http
│       └── logger
├── go-fiber-crm-basic
│   ├── SIDDO
│   ├── database
│   │   └── database.go
│   ├── go.mod
│   ├── go.sum
│   ├── lead
│   │   └── lead.go
│   └── main.go
├── go-fiber-mongo-hrms
│   ├── go.mod
│   ├── go.sum
│   ├── main.exe
│   └── main.go
├── go-graphql-done
│   ├── database
│   │   └── database.go
│   ├── go.mod
│   ├── go.sum
│   ├── gqlgen.yml
│   ├── graph
│   │   ├── generated.go
│   │   ├── model
│   │   │   └── models_gen.go
│   │   ├── resolver.go
│   │   ├── schema.graphqls
│   │   └── schema.resolvers.go
│   ├── server.go
│   └── tools.go
├── go-grpc-X
│   └── proto
│       └── greet.proto
├── go-grpc-done
│   ├── client
│   │   ├── bi_stream.go
│   │   ├── client_stream.go
│   │   ├── main.go
│   │   ├── server_stream.go
│   │   └── unary.go
│   ├── go.mod
│   ├── go.sum
│   ├── proto
│   │   ├── greet.pb.go
│   │   ├── greet.proto
│   │   └── greet_grpc.pb.go
│   └── server
│       ├── bi_stream.go
│       ├── client_stream.go
│       ├── main.go
│       ├── server_stream.go
│       └── unary.go
├── go-jwt-project-done
│   ├── controllers
│   │   └── userController.go
│   ├── database
│   │   └── databaseConnection.go
│   ├── go.mod
│   ├── go.sum
│   ├── helpers
│   │   ├── authHelper.go
│   │   └── tokenHelper.go
│   ├── main.go
│   ├── middleware
│   │   └── authMiddleware.go
│   ├── models
│   │   └── userModel.go
│   └── routes
│       ├── authRouter.go
│       └── userRouter.go
├── go-mongo
│   ├── controllers
│   │   └── user.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── models
│       └── user.go
├── go-movies-crud
│   ├── go-movies-crud.exe
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── go-pexels-api
│   ├── go.mod
│   └── main.go
├── go-postgres
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middleware
│   │   └── handlers.go
│   ├── models
│   │   └── models.go
│   └── router
│       └── router.go
├── go-server
│   ├── go-server.exe
│   ├── go.mod
│   ├── main.go
│   └── static
│       ├── form.html
│       └── index.html
├── hex
│   ├── cmd
│   │   └── main.go
│   ├── go.mod
│   └── internal
│       ├── adapters
│       │   ├── app
│       │   │   └── api
│       │   │       └── api.go
│       │   ├── core
│       │   │   └── arithmetic
│       │   │       └── arithmetic.go
│       │   └── framework
│       │       └── right
│       │           └── db
│       │               └── db.go
│       └── ports
│           ├── app.go
│           ├── core.go
│           └── framework_right.go
├── restraurant-management-done
│   ├── controllers
│   │   ├── foodController.go
│   │   ├── invoiceController.go
│   │   ├── menuController.go
│   │   ├── orderController.go
│   │   ├── orderItemsController.go
│   │   ├── tableController.go
│   │   └── userController.go
│   ├── database
│   │   └── databaseConnection.go
│   ├── go.mod
│   ├── go.sum
│   ├── helpers
│   │   └── tokenHelper.go
│   ├── main.go
│   ├── middleware
│   │   └── authMiddleware.go
│   ├── models
│   │   ├── foodModel.go
│   │   ├── invoiceModel.go
│   │   ├── menuModel.go
│   │   ├── noteModels.go
│   │   ├── orderItemModel.go
│   │   ├── orderModel.go
│   │   ├── tableModel.go
│   │   └── userModel.go
│   └── routes
│       ├── foodRouter.go
│       ├── invoiceRouter.go
│       ├── menuRouter.go
│       ├── orderItemRouter.go
│       ├── orderRouter.go
│       ├── tableRouter.go
│       └── userRouter.go
├── shorten-url-fiber-redis-go-done
│   ├── api
│   │   ├── Dockerfile
│   │   ├── database
│   │   │   └── database.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── helpers
│   │   │   └── helpers.go
│   │   ├── main.go
│   │   └── routes
│   │       ├── resolve.go
│   │       └── shorten.go
│   ├── data
│   ├── db
│   │   └── Dockerfile
│   └── docker-compose.yml
├── slack-age-bot
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── slack-age-bot.exe
└── slack-file-bot
    ├── CM4109_Java Programming-II.pdf
    ├── go.mod
    ├── go.sum
    ├── index.php
    └── main.go

106 directories, 173 files

```
