# lib-management

## Introduction
Creating a Backend library management system which can be used in real life with Go and PostgreSQL. Using chi for router and pgx for communication between  DB and serrver. For postgresql I am using a postgresql in a docker cotainer.

## Installation
  ### Requirements 
  - go version 1.16 +
  - postgreSQL
  
  ## Steps
  - create a database and add it's url in .env as DB_URL
  - run `make run`

## Tutorial
  ### Users
    - POST 
      `
      
      `

## TODO
- [x] Create Books
  - [x] Add author if not exsists
  - [x] Check Book if exists
  - [x] API handlers
- [x] Create Instances of book
  - [x] Link instances to book
  - [x] loan instances to members
- [x] Users/Admin Authentication
- [ ] User Authorization
- [ ] Loan books
  - [x] create new loan
  - [ ] delete old loan
  - [x] change status of instance
  - [ ] check if member active
  - [ ] late fees functionality
- [ ] frontend
