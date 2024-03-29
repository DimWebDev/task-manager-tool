# Task Manager API Documentation

## Overview

### Prerequisites

**Before running the Task Manager Tool locally, please ensure you have the following installed:**

- **GoLang**: Refer to the `go.mod` file for the required Go version.
- **PostgreSQL**: Version 16.

- Note: The `go.mod` file in the project's root directory specifies the exact version of Go required to run this application. For PostgreSQL, please install the specified version or higher for compatibility.

This project is centered around building a comprehensive Task Manager tool, utilizing GoLang for backend development. In order to proceed with it and run it locally, you must first download Golang and PostgreSQL. The core of the application is an API that provides the necessary operations to create, retrieve, update, and delete tasks. Following RESTful principles, this API offers a coherent and user-friendly set of endpoints for client-side interactions.

### API-Driven Development Approach

The development of the Task Manager tool is guided by the principles of API-driven development, a decision made for several key reasons:

- **Integration of Frontend and Backend**: By designing the API first, we ensure that this backend project can seamlessly integrate the frontend parts of a broader application. This clear delineation of responsibilities allows for efficient parallel development.

- **Client Flexibility**: The API-first approach enables the Task Manager tool to be accessible to a variety of clients, including web interfaces, mobile applications, and potential third-party integrations.

- **Clear Contract**: Establishing the API early in the development process creates a clear "contract". This guides both frontend and backend development teams, ensuring that both sides have a common understanding of the data flow and business logic.

- **Scalability and Maintainability**: A well-defined API is easier to scale and maintain, providing a solid foundation for the application to evolve and adapt over time.

- **Focused Development**: This methodology allows for a focused development process, where backend developers can concentrate on building a robust and efficient API, and frontend developers can create a responsive and user-friendly interface.

The Task Manager tool, built with GoLang and following an API-driven development approach, aims to be versatile, scalable, and maintainable. It's designed to provide a solid foundation for future growth and enhancements, making the application robust and adaptable for various user needs.

Version: 1.0.0

## API Endpoints

### Task Operations

#### List Tasks

`GET /tasks`

Retrieves a list of tasks.

**Responses:**

- `200 OK`: Successfully retrieved list of tasks.
- `500 Internal Server Error`: Failed to retrieve tasks due to a server error.

#### Create a New Task

`POST /tasks`

Allows the creation of a new task with necessary details.

**Request Body:**

- `Task`: JSON object containing `title` and `description`.

**Responses:**

- `201 Created`: Successfully created a new task.
- `400 Bad Request`: Failure due to invalid input.
- `500 Internal Server Error`: Failed to create a task due to a server error.

#### Get a Task by ID

`GET /tasks/{id}`

Retrieves a single task based on the provided ID.

**Parameters:**

- `id`: ID of the desired task.

**Responses:**

- `200 OK`: Successfully retrieved the task.
- `404 Not Found`: Task with given ID does not exist.
- `500 Internal Server Error`: Failed to retrieve the task due to a server error.

#### Update a Task

`PUT /tasks/{id}`

Updates the details of an existing task.

**Parameters:**

- `id`: ID of the task to update.

**Request Body:**

- `Task`: JSON object containing updated `title`, `description`, etc.

**Responses:**

- `200 OK`: Successfully updated the task.
- `400 Bad Request`: Failure due to invalid input.
- `404 Not Found`: Task with given ID does not exist.
- `500 Internal Server Error`: Failed to update the task due to a server error.

#### Delete a Task

`DELETE /tasks/{id}`

Deletes the specified task.

**Parameters:**

- `id`: ID of the task to delete.

**Responses:**

- `204 No Content`: Successfully deleted the task.
- `404 Not Found`: Task with given ID does not exist.
- `500 Internal Server Error`: Failed to delete the task due to a server error.

## Schemas

### Task

Represents a task within the Task Manager application.

- `id` (integer): Unique identifier for the task.
- `title` (string): Title of the task.
- `description` (string): A detailed description of the task.
- `dueDate` (string, optional): Due date of the task in YYYY-MM-DD format.
- `priority` (string, optional): Task priority level.
- `status` (string, optional): Current status of the task.

### ErrorResponse

Represents an error response when operations fail.

- `message` (string): A human-readable message providing more details about the error.

## Database Schema Definition

This project uses PostgreSQL 16.
As part of the backend setup for the Task Manager tool, a PostgreSQL database schema has been defined to store and manage the data related to tasks.

### Database Setup on macOS

PostgreSQL was chosen as the database system, leveraging its robust features and compatibility with GoLang. The database server was installed and initiated on macOS, and a dedicated database named `task_manager` was created to store the application data.

### Tasks Table Structure

The central component of the database schema is the `tasks` table, which is structured as follows:

- `id`: A unique identifier for each task. It is an auto-incrementing integer and is the primary key of the table.
- `title`: A string that holds the title of the task. This field is mandatory for each task entry.
- `description`: A text field to store detailed information about the task. It can accommodate longer text entries.
- `dueDate`: A date field that captures the due date for the task completion.
- `priority`: A string that indicates the priority of the task. This could be expanded to use an enumeration for predefined priority levels.
- `status`: A string that represents the current status of the task, such as 'pending' or 'completed'. Like priority, this could also be implemented as an enum in the future.

### Schema Creation Command

1. Create Database:
   Create a new database called `task_manager`:

   ```sql
   CREATE DATABASE task_manager;
   ```

2. The following SQL command was used to create the `tasks` table:

   ```sql
   CREATE TABLE tasks (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       dueDate DATE,
       priority VARCHAR(50),
       status VARCHAR(50)
   );
   ```

This command was executed in the PostgreSQL terminal interface, after connecting to the task_manager database.

### Verify Table Creation

After creating the table, you can verify its existence and structure using PostgreSQL commands. These commands help ensure that the table has been created with the correct structure.

- To list all tables in the current database:

  ```sql
  \dt

  ```

- To describe the structure of the tasks table specifically:

  ```sql
  \d tasks
  ```

Executing this command will display the detailed structure of the tasks table. It shows all column names, their respective data types, and any constraints applied, such as the PRIMARY KEY.

By following these verification steps, you can confidently proceed with the development, knowing that your database schema is set up correctly.

## Data Access Layer (DAL) Implementation

The Data Access Layer (DAL) for the Task Manager application has been implemented to provide an abstraction layer for database operations. This layer, contained within the `taskrepo.go` file in the `internal/repo` directory, enables the application to perform CRUD operations on tasks without direct interaction with the underlying SQL queries.

### Commitment to Industry Standards

To adhere to industry standards and best practices, I have incorporated unit testing into the development process. These tests ensure that each function in the DAL operates correctly. By isolating each part of the application and testing it in isolation, we can identify and resolve issues promptly, leading to a more stable and reliable application.

### Unit Testing with SQLMock

For our unit tests, I employ `go-sqlmock`, a mock SQL driver that simulates database operations. This tool allows us to test our DAL functions without the need for a real database connection, thus ensuring that the tests are fast, reliable, and do not have side effects on actual databases.

## Presentation Layer

### Overview

The presentation layer of the Task Manager Tool is engineered to ensure a seamless and interactive user experience. It functions as the primary interface for all API requests, effectively directing these requests to the appropriate business logic and data access layers.

### Design & Implementation

Adhering to RESTful design principles, the presentation layer is composed of HTTP handlers responsible for the CRUD operations related to task management. Located within the internal/api/handlers directory, these handlers are pivotal in maintaining a clear separation of concerns, each dedicated to handling specific aspects of the application's functionality.

### Routing

Our routing configuration, established in router.go, utilizes the gorilla/mux package to associate HTTP methods and URL paths with their designated handlers. This setup ensures a robust and intuitive routing mechanism across the application.

### Running the Application's Handlers

To interact with the Task Manager Tool's functionality, you can locally execute the handlers using Postman, which simulates the client-side interaction with the API.

#### Prerequisites

1. **Clone the Repository**
2. **Start the Application**

   Within the root directory of the cloned repository, initiate the application:

   ```sh
   cd task-manager-tool
   go run cmd/main.go
   ```

   This will start the local server, typically listening on http://localhost:8080.

#### Running the `CreateTaskHandler` handler with Postman

**Steps**

1. **Open Postman on your local machine**
2. **Configure the request**

   - Set the HTTP method to "POST".
   - Enter the request URL: `http://localhost:8080/tasks`
   - In the "Headers" section, add a header with Content-Type as the key and application/json as the value.

3. **Specify task details**

   - Navigate to the "Body" tab, opt for "raw" data input, and select "JSON" format.
   - Input the details of the task you want to create. For example:

   ```sh
   {
     "title": "Sample Task",
     "description": "This is a sample task to test the CreateTaskHandler",
     "dueDate": "2023-12-31T00:00:00Z",
     "priority": "High",
     "status": "Open"
   }
   ```

4. **Sumbit the request by clicking the "Send" button**

#### Running the `GetAllTasks` handler with Postman

**Steps**

1. **Open Postman on your local machine**
2. **Configure the request**

   - Set the HTTP method to "GET".
   - Enter the request URL to fetch all tasks: `http://localhost:8080/tasks`

3. **Sumbit the request by clicking the "Send" button**
   - A successful request will result in a 200 OK status code, and the response body will contain a JSON array of tasks. Each task in the response should adhere to the structure of your Task model. In the case of no available tasks, an empty array ([]) might be returned, depending on your API's implementation.

#### Running the `GetTaskByID` handler with Postman

**Steps**

1. **Open Postman on your local machine**
2. **Configure the request**

   - Set the HTTP method to "GET".
   - Enter the request URL, including the ID of the task you want to retrieve: `http://localhost:8080/tasks/1`. Replace 1 with the actual ID of the task you're interested in. Our databse uses an integer as aut0-increment primary key.

3. **Sumbit the request by clicking the "Send" button**

#### Running the `UpdateTask` handler with Postman

**Steps**

1. **Launch Postman on your local machine**
2. **Configure the request**

   - Set the HTTP method to "PUT".
   - Enter the request URL, including the ID of the task you want to retrieve: `http://localhost:8080/tasks/1`. Replace 1 with the actual ID of the task you're interested in. Our databse uses an integer as aut0-increment primary key.
   - Navigate to the "Headers" tab below the URL input.
   - Enter Content-Type as the key and application/json as the value. This header informs the server that the request body contains JSON.
   - Prepare the JSON Payload. Choose the "raw" radio button. Select JSON from the dropdown menu that appears next to the radio buttons. Input the JSON data that corresponds to the Task struct in your Go application. For example:

   ```sh
   {
    "title": "Updated Task Title",
    "description": "This is an updated description for the task.",
    "dueDate": "2024-12-31T23:59:59Z",
    "priority": "Medium",
    "status": "In Progress"``
   }
   ```

3. **Sumbit the request by clicking the "Send" button**

#### Running the `DeleteTask` handler with Postman

**Steps**

1. **Launch Postman on your local machine**
2. **Configure the request**

   - Set the HTTP method to "DELETE".
   - Enter the request URL, including the ID of the task you want to delete: `http://localhost:8080/tasks/1`. Replace 1 with the actual ID of the task you're targeting. The database uses an integer as an auto-increment primary key.

3. **Sumbit the request by clicking the "Send" button**
   - A successful deletion will result in a 204 No Content response if your API is configured to respond with this status code.
   - If the specified task is not found, you should see a 404 Not Found response.
   - Any server errors will typically result in a 500 Internal Server Error.

## Unit Testing

This application ensures the robustness of its features through comprehensive unit testing. I have developed tests for both the repository (repo) layer and the HTTP handlers (handlers) to verify that our business logic and API endpoints function as expected.

### Repository Tests

The repository tests validate the data access layer, ensuring that all interactions with the database operate correctly. These tests cover a range of scenarios, including successful data retrieval and handling of error cases.

### Handlers Tests

The handlers tests check the application's API layer. They simulate HTTP requests and assert that the responses and HTTP status codes returned by each endpoint are correct. These tests are crucial in confirming that our API behaves correctly under various conditions.

### Running the Tests

To run the unit tests, navigate to the top-level directory of the repository and use the following command:

```sh
go test ./...
```

This command will recursively run all tests in the project, including both the repository and handler tests. If you wish to run tests for a specific package, navigate to the package's directory and execute go test
