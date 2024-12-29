# Pulse API

Welcome to **Pulse**, a dynamic social network where users can create their own blogs, write engaging posts on a variety of topics, and share their thoughts with the world.

This application was developed as part of the technical challenge for the PROD '23 contest. For more details, you can find the task description and the OpenAPI specification in the `docs` directory.

## Topics covered

- **Authorization and Authentication**: Securely register and log in to user accounts.

- **RESTful API Design**: Interact with profiles, friends, and posts through a structured API.

- **JSON Responses**: Receive data in a standardized JSON format for easy parsing.

- **Makefile**: Automate tasks and manage builds efficiently. 

- **Database Migrations**: Manage schema changes with simple commands like `make up` and `make down`.

## Technologies Used

- **Go**: The programming language used for the backend.

- **Gin**: A web framework for building the API.

- **PostgreSQL**: A relational database for data storage.

<!-- - **Swagger**: API documentation for easy reference. -->

- **JWT**: For authorization and authentication.

## Getting Started

Follow these steps to run the application locally:

1. **Clone the Repository**:
```bash
git clone https://github.com/ursulgwopp/pulse-api
```

2. **Install Required Go Packages:**
```bash
go mod tidy
```

3. **Run the Application:**
```bash
make up && make run
```

<!-- 4. **Access the API:**
The API will be available at http://localhost:2024/swagger/index.html. -->

## TODO
- Make code cleaner

- Add TXs to repository layer

## Contributing

Contributions are welcome! If you have suggestions for improvements or want to report a bug, please open an issue or submit a pull request.
