# TaskFlow

TaskFlow is a beginner-level CRUD operations todo app built using Golang for the backend and React for the frontend. This project serves as a hands-on learning experience for creating, reading, updating, and deleting todo items.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Backend (Golang)](#backend-golang)
  - [Frontend (React)](#frontend-react)
- [Running the Application](#running-the-application)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

<a name="features"></a>
## Features
- Tech Stack: Go, React, TypeScript, MongoDB, TanStack Query, ChakraUI
- CRUD functionality for todos
- Responsive design for various screen sizes
- Stylish UI components with ChakraUI

<a name="prerequisites"></a>
## Prerequisites
- Golang (version 1.22.4)
- Node.js (version 18.3.1)
- MongoDB (Atlas or local instance)

<a name="installation"></a>
## Installation

<a name="backend-golang"></a>
### Backend (Golang)
1. **Clone the repository:**
   ```sh
   git clone https://github.com/DipeshBartaula/TaskFlow_GoReact.git
   cd TaskFlow_GoReact
   
2. **Install dependencies:**
    ```sh
    go mod tidy
3. **Create a `.env` file and add your MongoDB URI, and environment:**
    ```sh
    MONGODB_URI=your_mongodb_uri
    ENV=development
    PORT=5000

4. **To ensure that your React frontend can communicate with your Golang backend without any issues, you need to enable CORS in your Golang server**
    ```sh
    Uncomment line 69-72 in main.go for CORS

<a name="frontend-react"></a>
### Frontend(React)
1. **Navigate to the frontend directory:**
    ```sh
    cd client
    
2. **Install dependencies:**
    ```sh
    npm install
  
<a name="running-the-application"></a>
## Running the Application

### Backend (Golang)
1. **Start the Golang Server:**
    ```sh
    go run main.go
    
### Frontend (React)
1. **Start the React development server:**
    ```sh
    cd client
    npm run dev
    
The React app will start and opein in your default web browser, typically at `http://localhost:5173`.

<a name="usage"></a>
## Usage
1. **Open the React app in your browser.**
2. **Use the interface to add, view, update, and delete todo items.**
3. **The backend server handles all API requests for CRUD operations.**
