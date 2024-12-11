Project Idea: Task Management System

Description

Build a Task Management System where users can:
1. Create an account and log in (authentication).
2. Add, view, update, and delete tasks.
3. Assign priorities, deadlines, and statuses to tasks.
4. Group tasks into categories or projects.
5. View statistics such as completed tasks, pending tasks, etc.
Features
1. User Authentication
• Sign up and log in using secure password hashing (e.g., bcrypt).
2. Task Management
• CRUD operations for tasks (Create, Read, Update, Delete).
• Filter tasks by status, priority, or deadlines.
3. Project and Category Management
• Group tasks into projects or categories.
• Fetch tasks by project or category.
4. Database Design
• Use PostgreSQL for relational data storage.
• Ensure normalization and proper indexing.
5. RESTful API
• Implement a RESTful API with Gin for all operations.
6. Logging
• Add logging for debugging and monitoring.
7. Testing
• Unit tests for key functionalities using Go’s testing package.

- sample tables
- Users (id,username,email, password,status,last_modified)
- projects(id,user_id,name,description, status, last_modified)
- tasks(id,user_id,title,description, priority(low,medium,high),status(pending,in-progress,completed),due_date,last_modified)
- project_tasks(id,project_id, task_id)