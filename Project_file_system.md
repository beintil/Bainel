# Project file system

* ./cmd/myapp -- Basic main file. For start to server
* ./cfg -- Configuration file
* ./internal/app/authentication -- Authentication users. (registration, authorization etc.)
* ./internal/app/client -- User action. (search user by id)
* ./internal/app/database -- Database MongoDB connection
* ./internal/app/handler -- Handler 
* ./internal/app/home -- Home page
* ./internal/pkg/myapp -- The main code for starting the server
* ./pkg/error_handler/client -- Error handler user (not Fatal error)
* ./pkg/error_handler/server -- Error handler server (Fatal error)
* ./repository/user -- Structure users and authorization 
* ./tests -- Tests
* Dockerfile -- Dockerfile
* Makefile -- Starting unit tests and lints