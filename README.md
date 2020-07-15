# go-bookstore
Simple Bookstore backend services, built with Golang

----------------------------
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Work Progress](#work-progress)
----------------------------
## Features
* Book Catalog
* Transaction
* JWT Authentication dan Authorization
* Registration
* Searching
* Best Seller

## Requirements
  * Go 1.10 or higher
  * MySQL (4.1+)

## Installation
1. Clone this Repo
2. Rename .env.example to .env and change the variable value
3. Run this command in terminal / shell
```bash
#Build Project
go build .

##Migrate DB
./go-bookstore -migrate

#Start Server
./go-bookstore
```

## Work Progress
* :white_check_mark: Init Project
* :white_check_mark: Setting up env
* :white_check_mark: Create Database Migration
* :black_square_button: Create Database Seed
* :black_square_button: Create Router
