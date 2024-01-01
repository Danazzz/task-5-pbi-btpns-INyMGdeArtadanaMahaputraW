<h1>BTPN Syariah x Rakamin Final Project</h1>

<h3>Description</h3>

This project serves as the final task for the Project-Based Internship conducted in collaboration between BTPN Syariah and Rakamin. The project is a Full Stack application that provides APIs for managing users and photos. The APIs include functionalities such as user registration, login, updating user information, deleting users, creating photos, fetching photos, updating photo information, and deleting photos. The technologies used include the Gin Gonic Framework, Gorm, JWT Go, and Go Validator. The application utilizes a relational database PostgreSQL for storing user and photo data.

<h3>Project Structure</h3>

The project adheres to a modular structure to enhance organization and maintainability:

<b>app</b>: Contains the main struct definition<br />
<b>controllers</b>: Hosts the logic for handling HTTP requests related to users and photos.<br />
<b>database</b>: Manages the configuration, connection, and migration of the database.<br />
<b>helpers</b>: Includes helper functions for JWT, bcrypt, formatter, and environment<br />
<b>middleware</b>: Implements middleware functions, such as JWT authentication.<br />
<b>models</b>: Defines the database models for users and photos using Gorm.<br />
<b>router</b>: Configures the API routes and endpoints using the Gin Gonic router.<br />
<b>go.mod</b>: Manages the project's dependencies and packages.<br />

<h3>Tools Used</h3>

Gin Gonic Framework: A web framework for Go.<br />
Gorm: A powerful ORM library for Go.<br />
JWT Go: JSON Web Token library for Go.<br />
Go Validator: A package of validators and sanitizers for strings, numbers, slices, and structs.

<h3>Setup</h3>

Install Go dependencies using `go mod tidy`.<br />
Configure and set up the database by executing the migration scripts in the database directory.<br />
Customize the database connection details in the database/config.go file.<br />
Run the main application using `go run main.go`.<br />

<h3>API Endpoints</h3>

<h4>User Endpoints</h4>
<b>POST</b> /users/register: Register a new user with the provided information.<br />
Attributes:<br />
- ID (primary key, required)<br />
- Username (required)<br />
- Email (unique & required)<br />
- Password (required & minlength 6)<br />
- Created At (timestamp)<br />
- Updated At (timestamp)<br />
<b>POST</b> /users/login: Login with email and password.<br />
Attributes:<br />
- Email & Password (required)<br />
<b>PUT</b> /users/:userId: Update user information.<br />
Parameters:<br />
- userId (user ID to update)<br />
<b>DELETE</b> /users/:userId: Delete a user.<br />
Parameters:<br />
- userId (user ID to delete)<br />

<h4>Photos Endpoints</h4>
<b>POST</b> /photos: Create a new photo with the provided information.<br />
Attributes:<br />
- ID<br />
- Title<br />
- Caption<br />
- PhotoUrl<br />
- UserID<br />
<h4>Relationships:</h4>
User (constraint cascade)<br />
<b>GET</b> /photos: Fetch all photos.<br />
<b>PUT</b> /photos/:photoId: Update photo information.<br />
Parameters:<br />
photoId (photo ID to update)<br />
<b>DELETE</b> /photos/:photoId: Delete a photo.<br />
Parameters:<br />
photoId (photo ID to delete)<br />

<h3>Authorization</h3>

Authorization is implemented using JWT (JSON Web Token). Ensure proper authentication headers are included in requests to protected endpoints.
