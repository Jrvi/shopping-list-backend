# Shopping List Backend

This is a backend service for managing a shopping list. It provides a RESTful API to manage products and categories.

## Requirements

- Go 1.23.4 or later
- MySQL database

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/jrvi/shopping-list-backend.git
    cd shopping-list-backend
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Configure the database connection in [database.go](http://_vscodecontentref_/1):
    ```go
    dsn := "shopping:salainen@tcp(localhost:3306)/shopping_list_db"
    ```

4. Initialize the database:
    ```sh
    go run main.go
    ```

## API Endpoints

### Products

- **GET /product**
    - Retrieves a list of all products.
    - Response:
        ```json
        [
            {
                "id": "1",
                "title": "Milk",
                "category_id": "1"
            },
            ...
        ]
        ```

- **POST /product**
    - Adds a new product.
    - Request body:
        ```json
        {
            "id": "1",
            "title": "Milk",
            "category_id": "1"
        }
        ```

### Categories

- **GET /categories**
    - Retrieves a list of all categories.
    - Response:
        ```json
        [
            {
                "id": "1",
                "title": "Dairy"
            },
            ...
        ]
        ```

- **POST /categories**
    - Adds a new category.
    - Request body:
        ```json
        {
            "id": "1",
            "title": "Dairy"
        }
        ```

## Running the Server

To run the server, execute the following command:
```sh
go run main.go