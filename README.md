# Welcome to the Golang MySQL Inventory Project - [KodeKloud](https://kodekloud.com/lessons/api-development-project/)

This project utilizes Go programming language (Golang) to interact with a MySQL database. The main feature of this project is a simple CRUD (Create, Read, Update, Delete) operation on a single "Inventory" table with four columns: `id`, `name`, `quantity`, and `price`.

## Features

* **CRUD operations**: This application allows you to perform basic CRUD operations on the `inventory` table using a user interface or through API calls.
* **User Interface**: You can access the UI by going to <http://localhost:10000> in your web browser. From there, you can add new items to the inventory or modify existing ones.
* **API Endpoints**: To integrate this functionality into other applications or services, you can use the provided REST APIs. These endpoints accept HTTP requests containing data in JSON format and return results as plain text or JSON objects. Available endpoints include:
  * GET /products: Retrieve all items from the inventory
  * POST /product: Add a new item to the inventory
  * PUT /product/{id}: Modify an existing item
  * DELETE /product/{id}: Remove an item from the inventory

## Technologies Used

* **Go Programming Language (Golang)**: For back-end development
* **MySQL Database**: For storing and retrieving data

## Usage

To start the server and begin testing, follow these steps:

1. Clone the repository or download it locally: `git clone repo`
2. Navigate to the project directory in the terminal: `cd path/to/project`
3. Run `go build` command to compile the code. If everything goes well, no errors should appear.
4. Start the server with the following command: `go run .` or `./my-inventory` for Unix | `.\my-inventory.exe` for Windows
5. Open your preferred browser and go to <http://localhost:10000>
6. Explore the available API endpoints listed above and feel free to test them out using tools like Postman, curl, etc.
