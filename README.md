# Task Manager API Documentation

## Overview

This project is centered around building a comprehensive Task Manager tool, utilizing GoLang for backend development. In order to proceed with it and run it locally, you must first download Golang and PostgreSQL The core of the application is an API that provides the necessary operations to create, retrieve, update, and delete tasks. Following RESTful principles, this API offers a coherent and user-friendly set of endpoints for client-side interactions.

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

I have implemented the Data Access Layer (DAL) for our Task Manager application. The DAL is a crucial component that provides an abstraction over the actual database operations, such as creating, retrieving, updating, and deleting tasks. This layer is implemented in the taskrepo.go file within the internal/repo directory.

The purpose of the DAL is to interact directly with the database through SQL queries, ensuring that the rest of our application can perform data operations without being concerned with the underlying database specifics. This separation of concerns leads to cleaner, more maintainable, and scalable code.

### Commitment to Industry Standards

As part of my commitment to best practices and industry standards, unit testing is introduced. Unit tests are automated tests written to ensure that a section of our application (known as the "unit") meets its design and behaves as intended.

In the context of our DAL, unit testing involves testing each function in taskrepo.go to validate their expected behavior against a controlled test database. This process helps us identify and fix issues early in the development cycle, paving the way for a stable and reliable application.

### Setting Up a Test Database

To facilitate effective unit testing, we will set up a test database that mirrors the structure of our production database. This test database allows us to replicate real-world scenarios and test our DAL's interactions with the database. By doing so, we can simulate the behavior of our application in a production-like environment, ensuring that our unit tests provide us with accurate and meaningful results.

The test database will be used exclusively for running our tests and will not contain any real user data. It is configured to be reset before each test to maintain a consistent starting state for every test run.

## Unit Testing

To ensure the reliability and robustness of the Data Access Layer (DAL) in our Task Manager application, we adhere to industry best practices by implementing comprehensive unit tests. These tests are designed to validate that each function within our DAL performs as expected under various conditions.

### Setting Up a Test Database

To run the tests locally, you will need to set up a test database that mirrors the structure of the production database. Follow these steps to create and configure your test database:

1. **Connect to PostgreSQL**:
   Open your PostgreSQL command line tool and connect to your PostgreSQL server:

   ```sql
   \c postgres
   ```

2. **Create Test Database**:
   Create a new test database called `task_manager_test`:

   ```sql
   CREATE DATABASE task_manager_test;
   ```

3. **Set Up Test Database Schema**:
   Connect to the test database and create the necessary tables:

   ```sql
   \c task_manager_test

   CREATE TABLE tasks (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       dueDate DATE,
       priority VARCHAR(50),
       status VARCHAR(50)
   );
   ```

4. **Insert Seed Data** (Optional):
   Optionally, you can insert some seed data for testing purposes:
   ```sql
   INSERT INTO tasks (title, description, dueDate, priority, status)
   VALUES ('Test Task 1', 'Description for test task 1', '2024-01-01', 'High', 'Pending');
   ```

### Running the Tests

Once your test database is set up, you can run the unit tests using the Go command:

```sh
go test ./...
```
