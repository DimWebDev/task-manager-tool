# Task Manager API Documentation

## Table of Contents

1. [Overview](#overview)
   - [Prerequisites](#prerequisites)
   - [API-Driven Development Approach](#api-driven-development-approach)
2. [API Endpoints](#api-endpoints)
   - [List Tasks](#list-tasks)
   - [Create a New Task](#create-a-new-task)
   - [Get a Task by ID](#get-a-task-by-id)
   - [Update a Task](#update-a-task)
   - [Delete a Task](#delete-a-task)
3. [Schemas](#schemas)
   - [Task](#task)
   - [ErrorResponse](#errorresponse)
4. [Database Schema Definition](#database-schema-definition)
   - [Database Setup on macOS](#database-setup-on-macos)
   - [Tasks Table Structure](#tasks-table-structure)
5. [Data Access Layer (DAL) Implementation](#data-access-layer-dal-implementation)
6. [Presentation Layer](#presentation-layer)
   - [Running the Handlers Locally with Postman](#running-the-handlers-locally-with-postman)
7. [Unit Testing](#unit-testing)
   - [Repository Tests](#repository-tests)
   - [Handlers Tests](#handlers-tests)
8. [Deployment Guidelines](#deployment-guidelines)
9. [Security Considerations](#security-considerations)

## Overview

### Prerequisites

**Before running the Task Manager Tool locally, please ensure you have the following installed:**

- **GoLang**: Refer to the `go.mod` file for the required Go version. Use the command `go mod download` to install the required dependencies.
- **PostgreSQL**: Version 16 or higher. Ensure you have configured the PostgreSQL user with the appropriate privileges for the `task_manager` database.

Note: The `go.mod` file in the project's root directory specifies the exact version of Go required to run this application. For PostgreSQL, please install the specified version or higher for compatibility.

This project is centered around building a comprehensive Task Manager tool using GoLang for backend development. The core of the application is an API that provides operations to create, retrieve, update, and delete tasks, following RESTful principles. This API offers a coherent and user-friendly set of endpoints for client-side interactions. By adhering to RESTful design, the application ensures a high level of compatibility with various clients and simplifies integrations.

### API-Driven Development Approach

The development of the Task Manager tool is guided by API-driven development, a decision made for several key reasons:

- **Integration of Frontend and Backend**: By designing the API first, we ensure seamless integration with the frontend parts of the broader application. This approach allows different teams to work independently on frontend and backend components, reducing dependencies and speeding up the development process.
- **Client Flexibility**: The API-first approach enables various clients, including web interfaces, mobile applications, and potential third-party integrations, to access the backend. This makes it easier to support multiple platforms without making significant changes to the backend.
- **Clear Contract**: Establishing the API early in the development process creates a clear contract between frontend and backend development teams. This contract defines the interactions and expectations, minimizing miscommunication and making collaboration more efficient. It also helps in building mock clients for testing purposes early in the project lifecycle.
- **Scalability and Maintainability**: A well-defined API is easier to scale and maintain, providing a solid foundation for the application to grow over time. As the application expands, having a consistent and modular API helps ensure new features can be added without disrupting existing functionality.
- **Focused Development**: Backend developers can concentrate on building a robust API while frontend developers create a responsive interface. This separation of concerns not only improves efficiency but also ensures that each part of the application is optimized for its specific responsibilities.
- **Future Proofing**: With an API-first approach, it becomes much easier to future-proof the application. For instance, if new devices or technologies need to be integrated, the existing API can serve as a universal point of access, requiring fewer changes across the ecosystem.

Version: 1.0.0

## API Endpoints

### Task Operations

#### List Tasks

**`GET /tasks`**

Retrieves a list of tasks. This endpoint allows clients to query all tasks currently stored in the system, enabling visibility into all task data available.

**Responses:**

- **`200 OK`**: Successfully retrieved list of tasks.

  **Example Response:**
  ```json
  [
    {
      "id": 1,
      "title": "Sample Task",
      "description": "This is a sample task.",
      "dueDate": "2023-12-31",
      "priority": "High",
      "status": "Open"
    },
    {
      "id": 2,
      "title": "Another Task",
      "description": "Details about another task.",
      "dueDate": "2024-01-15",
      "priority": "Medium",
      "status": "In Progress"
    }
  ]
  ```
- **`500 Internal Server Error`**: Failed to retrieve tasks due to a server error. This can occur if there is an issue with the database connection or if an unexpected condition was encountered.

#### Create a New Task

**`POST /tasks`**

Allows the creation of a new task with necessary details. Clients can use this endpoint to add new tasks to the system. It requires at least a `title` and can include other optional details.

**Request Body:**

- **Task**: JSON object containing `title` (required), `description` (optional), `dueDate` (optional), `priority` (optional), `status` (optional).

**Example Request:**

```json
{
  "title": "New Task",
  "description": "Details about the new task",
  "dueDate": "2024-01-01",
  "priority": "Medium",
  "status": "In Progress"
}
```

**Responses:**

- **`201 Created`**: Successfully created a new task.

  **Example Response:**
  ```json
  {
    "id": 2,
    "title": "New Task",
    "description": "Details about the new task",
    "dueDate": "2024-01-01",
    "priority": "Medium",
    "status": "In Progress"
  }
  ```
- **`400 Bad Request`**: Failure due to invalid input (e.g., missing `title` or providing invalid data types for fields). This error is returned when mandatory fields are not provided, or when the provided data does not match the expected format.
- **`500 Internal Server Error`**: Failed to create a task due to a server error. This might happen if there is an issue in connecting to the database or in handling the request internally.

#### Get a Task by ID

**`GET /tasks/{id}`**

Retrieves a single task based on the provided ID. This endpoint is used to fetch specific task details, allowing clients to access and display information about a given task.

**Parameters:**

- **`id`**: (integer) ID of the desired task. The ID is used to uniquely identify the task that the client wishes to retrieve.

**Responses:**

- **`200 OK`**: Successfully retrieved the task.

  **Example Response:**
  ```json
  {
    "id": 1,
    "title": "Sample Task",
    "description": "This is a sample task.",
    "dueDate": "2023-12-31",
    "priority": "High",
    "status": "Open"
  }
  ```
- **`404 Not Found`**: Task with given ID does not exist. This response indicates that no task with the specified ID could be found in the system.
- **`500 Internal Server Error`**: Failed to retrieve the task due to a server error. This might be caused by issues with the database or other unexpected internal errors.

#### Update a Task

**`PUT /tasks/{id}`**

Updates the details of an existing task. This endpoint allows clients to modify the attributes of a specific task, such as changing the title, updating the description, or marking it as completed.

**Parameters:**

- **`id`**: (integer) ID of the task to update. The ID uniquely identifies which task will be updated.

**Request Body:**

- **Task**: JSON object containing updated `title`, `description`, etc.

**Example Request:**

```json
{
  "title": "Updated Task Title",
  "description": "This is an updated description for the task.",
  "dueDate": "2024-12-31",
  "priority": "Low",
  "status": "Completed"
}
```

**Responses:**

- **`200 OK`**: Successfully updated the task.
- **`400 Bad Request`**: Failure due to invalid input. This response is returned if the data provided does not meet the validation criteria, such as incorrect data types or missing required fields.
- **`404 Not Found`**: Task with given ID does not exist. This indicates that the task to be updated could not be found in the system.
- **`500 Internal Server Error`**: Failed to update the task due to a server error. This could result from internal system errors or issues interacting with the database.

#### Delete a Task

**`DELETE /tasks/{id}`**

Deletes the specified task. This endpoint is used when a client wants to remove a task permanently from the system.

**Parameters:**

- **`id`**: (integer) ID of the task to delete. The ID specifies the task that should be removed.

**Responses:**

- **`204 No Content`**: Successfully deleted the task. This response means the task was successfully removed from the system, and there is no additional content to return.
- **`404 Not Found`**: Task with given ID does not exist. This response indicates that no task with the specified ID could be found in the system to delete.
- **`500 Internal Server Error`**: Failed to delete the task due to a server error. This error might happen if there are internal issues preventing the task from being deleted.

## Schemas

### Task

Represents a task within the Task Manager application.

- `id` (integer): Unique identifier for the task.
- `title` (string): Title of the task.
- `description` (string): A detailed description of the task.
- `dueDate` (string, optional): Due date of the task in YYYY-MM-DD format.
- `priority` (string, optional): Task priority level, such as `High`, `Medium`, or `Low`.
- `status` (string, optional): Current status of the task, such as `Pending`, `In Progress`, or `Completed`.

### ErrorResponse

Represents an error response when operations fail.

- `message` (string): A human-readable message providing more details about the error. This helps the client understand what went wrong and provides guidance for resolving the issue.

## Database Schema Definition

### Database Setup on macOS

PostgreSQL was chosen for its robust features and compatibility with GoLang. Install and initiate the PostgreSQL server on macOS, and create a dedicated database named `task_manager` to store the application data.

**Create Database:**

```sql
CREATE DATABASE task_manager;
```

### Tasks Table Structure

The `tasks` table is structured as follows:

- `id`: A unique identifier for each task. It is an auto-incrementing integer and is the primary key of the table.
- `title`: A string that holds the title of the task. This field is mandatory.
- `description`: A text field to store detailed information about the task. This field can accommodate long descriptions, allowing users to provide all necessary details.
- `dueDate`: A date field capturing the due date for task completion. This helps in setting deadlines for task management.
- `priority`: A string indicating the priority of the task, such as `High`, `Medium`, or `Low`. This field can be used to prioritize tasks for better productivity.
- `status`: A string representing the current status of the task, such as `pending`, `in progress`, or `completed`. It provides insight into the progress and helps users track their workflow.

**Schema Creation Command:**

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

**Verify Table Creation:**

- List all tables:

  ```sql
  \dt
  ```

- Describe the structure of the `tasks` table:

  ```sql
  \d tasks
  ```

## Data Access Layer (DAL) Implementation

The DAL, located in `taskrepo.go` within the `internal/repo` directory, provides an abstraction for database operations. It enables the application to perform CRUD operations without directly interacting with SQL queries. This abstraction is crucial for maintainability and for decoupling business logic from database-specific details, making the system more adaptable to future changes in the database.

The DAL also includes methods for querying, inserting, updating, and deleting tasks. These methods are wrapped in well-defined functions that return appropriate values and error messages to the higher-level application logic.

## Presentation Layer

The presentation layer handles HTTP requests for CRUD operations, adhering to RESTful design principles. Handlers are located in `internal/api/handlers` and routing is managed by `gorilla/mux`. The presentation layer is responsible for parsing client requests, invoking the appropriate business logic, and sending responses back to clients.

The routing is configured in `router.go`, mapping each endpoint to its respective handler. The `gorilla/mux` package helps ensure that routes are properly defined and requests are processed efficiently.

### Running the Handlers Locally with Postman

1. **Clone the Repository**

   Clone the repository to your local machine using Git to begin the setup process.

2. **Start the Application**

   ```sh
   cd task-manager-tool
   go run cmd/main.go
   ```

   The server will start locally, typically listening on `http://localhost:8080`. Ensure all dependencies are installed beforehand by using `go mod download`.

3. **Interact via Postman**

   Use Postman to test the various API endpoints:

   - **Create a Task**: Use `POST http://localhost:8080/tasks` with the example JSON body to create a new task.
   - **List All Tasks**: Use `GET http://localhost:8080/tasks` to retrieve a list of all tasks in the system.
   - **Get a Task by ID**: Use `GET http://localhost:8080/tasks/{id}` to fetch a specific task using its unique ID.
   - **Update a Task**: Use `PUT http://localhost:8080/tasks/{id}` with the updated JSON body to modify an existing task.
   - **Delete a Task**: Use `DELETE http://localhost:8080/tasks/{id}` to remove a task from the system.

## Unit Testing

### Repository Tests

Repository tests validate interactions with the database, ensuring successful data retrieval and error handling. By testing the DAL, we verify that database queries are working correctly and that errors are handled gracefully. Tests are created using the `testing` package in Go, and mock database connections are established to isolate the unit tests.

### Handlers Tests

Handlers tests simulate HTTP requests and verify that each endpoint returns the correct response and status code. These tests are crucial for ensuring that the application logic is correctly processing requests and generating appropriate responses, even in edge cases. The handlers are tested with mock data to ensure there is no dependency on the actual database.

### Running the Tests

To run the unit tests, navigate to the top-level directory and use the command:

```sh
go test ./...
```

This command will recursively run all tests defined across the different packages in the project, ensuring that both DAL and handler layers are properly validated.


