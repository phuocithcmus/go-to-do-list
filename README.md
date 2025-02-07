# Project to_do_project

Project Description
The "Build Your Own Simple To-Do List API" project aims to develop a straightforward yet effective API that enables users to manage their daily tasks efficiently. This API will allow users to sign up for an account, log in, and create, edit, and delete tasks in their to-do list. Users will also be able to view their list of tasks, providing a streamlined way to keep track of their responsibilities and deadlines.

Detailed Description
In a busy world, keeping track of tasks and responsibilities is crucial for productivity and organization. This project will create an API that supports all fundamental functionalities required for a basic to-do list application. Here’s a detailed look at how users will interact with the app:

User Registration and Authentication:

Sign Up: New users can create an account by providing a username, email, and password. Upon registration, a confirmation email will be sent to verify the account.

Login: Registered users can log in using their email and password. For enhanced security, the API will support multi-factor authentication (MFA).

Task Management:

Create Task: Users can add new tasks to their to-do list by providing a title, description, due date, and priority level (e.g., low, medium, high). This helps users organize and prioritize their tasks effectively.

Edit Task: Users can update the details of their tasks, such as changing the title, description, due date, or priority. This feature ensures that users can adjust their tasks as needed.

Delete Task: Users can remove tasks from their to-do list. This helps users keep their list current and relevant.

View Tasks: Users can view their list of tasks, with the ability to filter and sort by due date, priority, and status (e.g., completed, pending).

Task Completion:

Mark Task as Complete: Users can mark tasks as completed once they have finished them. This provides a sense of accomplishment and helps users track their progress.

View Completed Tasks: Users can view a list of their completed tasks, allowing them to review their achievements and maintain a record of their work.

Real-World Example
Consider Sarah, a busy professional juggling multiple projects and personal commitments. Sarah signs up for an account on the to-do list app using her email and a secure password. She logs in and starts adding tasks to her list:

"Finish project report" due by Friday, marked as high priority.

"Grocery shopping" set for Saturday, marked as medium priority.

"Book dentist appointment" with no specific due date, marked as low priority.

Throughout the week, Sarah updates her tasks as needed. She completes the "Finish project report" task on Thursday and marks it as complete. Over the weekend, she deletes the "Grocery shopping" task after completing it and books her dentist appointment. Sarah frequently checks her to-do list to stay organized and ensure she doesn't miss any important tasks.

Introduction
The "Build Your Own Simple To-Do List API" project aims to develop a user-friendly and efficient API that supports essential functionalities for managing a to-do list. Users will be able to sign up for an account, log in, create, edit, and delete tasks, and view their list of tasks.

Objectives
Allow users to sign up, log in, and manage their accounts.

Enable users to create, read, update, and delete tasks.

Provide features to mark tasks as complete and view completed tasks.

Ensure secure and efficient handling of user data and tasks.

Functional Requirements
User Management
Sign Up: Users can create an account by providing a username, email, and password.

Login: Users can log in using their email and password.

Profile Management: Users can update their profile information.

Task Management
Create Task: Users can create a new task with a title, description, due date, and priority level.

Read Task: Users can view their tasks, with options to filter and sort them.

Update Task: Users can edit the details of their tasks.

Delete Task: Users can delete tasks from their list.

Mark Task as Complete: Users can mark tasks as completed.

View Completed Tasks: Users can view a list of completed tasks.

Non-Functional Requirements
Scalability: The API should handle a growing number of users and tasks.

Performance: The API should have a fast response time and handle concurrent requests efficiently.

Security: Implement authentication and authorization mechanisms to protect user data.

Reliability: The API should be highly available and handle failures gracefully.

Usability: The API should be easy to use and well-documented.

Use Cases
User Sign Up and Login: New users sign up and existing users log in.

Create and Manage Tasks: Users create, update, and delete their tasks.

Mark Task as Complete: Users mark tasks as completed.

View Tasks: Users view their list of tasks and completed tasks.

User Stories
As a user, I want to sign up for an account so that I can manage my tasks.

As a user, I want to log in to my account so that I can access my to-do list.

As a user, I want to create a new task so that I can add it to my to-do list.

As a user, I want to edit a task so that I can update its details.

As a user, I want to delete a task so that I can remove it from my list.

As a user, I want to mark a task as complete so that I can track my progress.

As a user, I want to view my completed tasks so that I can see what I have accomplished.

Technical Requirements
Programming Language: Choose an appropriate backend language (e.g., Node.js, Python, Ruby).

Database: Use a database to store user and task data (e.g., PostgreSQL, MongoDB).

Authentication: Implement JWT for secure user authentication.

API Documentation: Use Swagger or similar tools for API documentation.

API Endpoints
User Management
POST /signup: Register a new user.

POST /login: Authenticate a user.

GET /profile: Get user profile details.

PUT /profile: Update user profile.

Task Management
POST /tasks: Create a new task.

GET /tasks: Retrieve a list of tasks.

GET /tasks/{id}: Retrieve a single task by ID.

PUT /tasks/{id}: Update a task by ID.

DELETE /tasks/{id}: Delete a task by ID.

PUT /tasks/{id}/complete: Mark a task as complete.

GET /tasks/completed: Retrieve a list of completed tasks.

Security
Use HTTPS to encrypt data in transit.

Implement input validation and sanitization to prevent SQL injection and XSS attacks.

Use strong password hashing algorithms like bcrypt.

Performance
Implement caching strategies to improve response times.

Optimize database queries to handle large datasets efficiently.

Use load balancing to distribute traffic evenly across servers.

Documentation
Provide comprehensive API documentation using tools like Swagger.

Create user guides and developer documentation to assist with integration and usage.

Glossary
API: Application Programming Interface.

JWT: JSON Web Token.

CRUD: Create, Read, Update, Delete.

Appendix
Include any relevant diagrams, data models, and additional references.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
