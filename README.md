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

The following SQL command was used to create the `tasks` table:

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
