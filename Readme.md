# Student Management System - Developer Skill Test

A comprehensive full-stack web application for managing school operations including students, staff, classes, notices, and leave management. This project serves as a skill assessment platform for **Frontend**, **Backend**, and **Blockchain** developers.

## 🏗️ Project Architecture

```
skill-test/
├── frontend/           # React + TypeScript + Material-UI
├── backend/            # Node.js + Express + PostgreSQL
├── go-service/         # Golang microservice for PDF reports
├── seed_db/           # Database schema and seed data
└── README.md          # This file
```

## 🚀 Quick Start

### Prerequisites
- Node.js (v16 or higher)
- PostgreSQL (v12 or higher)
- npm or yarn

### 1. Backend Setup
```bash
cd backend
npm install
cp .env.example .env  # Configure your environment variables
npm start
```

### 2. Frontend Setup
```bash
cd frontend
npm install
npm run dev
```

### 3. Access the Application
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:5007
- **Demo Credentials**: 
  - Email: `admin@school-admin.com`
  - Password: `3OU4zn3q6Zh9`

### ** Database Setup **
```bash
# Create PostgreSQL database
createdb school_mgmt

# Run database migrations
psql -d school_mgmt -f seed_db/tables.sql
psql -d school_mgmt -f seed_db/seed-db.sql
```

## 🎯 Skill Test Problems

### **Problem 1: Frontend Developer Challenge**
**Fix "Add New Notice" Page**
- **Location**: `/app/notices/add`
- **Issue**: When clicking the 'Save' button, the 'description' field doesn't get saved
- **Skills Tested**: React, Form handling, State management, API integration
- **Expected Fix**: Ensure description field is properly bound and submitted

### **Problem 2: Backend Developer Challenge**
**Complete CRUD Operations in Student Management**
- **Location**: `/src/modules/students/students-controller.js`
- **Issue**: Implement missing CRUD operations for student management
- **Skills Tested**: Node.js, Express, PostgreSQL, API design, Error handling
- **Expected Implementation**: Full Create, Read, Update, Delete operations

### **Problem 3: Blockchain Developer Challenge**
**Implement Certificate Verification System**
- **Objective**: Add blockchain-based certificate verification for student achievements
- **Skills Tested**: Smart contracts, Web3 integration, Ethereum/Polygon
- **Requirements**:
  - Create smart contract for certificate issuance and verification
  - Integrate Web3 wallet connection in frontend
  - Add certificate management in admin panel
  - Implement IPFS for certificate metadata storage

### **Problem 4: Golang Developer Challenge**
**Build PDF Report Generation Microservice via API Integration**
- **Objective**: Create a standalone microservice in Go to generate PDF reports for students by consuming the existing Node.js backend API.
- **Location**: A new `go-service/` directory at the root of the project.
- **Description**: This service will connect to the existing Node.js backend's `/api/v1/students/:id` endpoint to fetch student data, and then use the returned JSON to generate a downloadable PDF report.
- **Skills Tested**: Golang, REST API consumption, JSON parsing, file generation, microservice integration.
- **Requirements**:
  - Create a new endpoint `GET /api/v1/students/:id/report` in the Go service.
  - The Go service must not connect directly to the database; it must fetch data from the Node.js API.
  - The developer **must** have the PostgreSQL database and the Node.js backend running to complete this task.

### **Problem 5: DevOps Engineer Challenge**
**Containerize the Full Application Stack**
- **Objective**: Create a multi-container setup to run the entire application stack (Frontend, Backend, Database) using Docker and Docker Compose.
- **Location**: `Dockerfile` in the `frontend` and `backend` directories, and a `docker-compose.yml` file at the project root.
- **Description**: The goal is to make the entire development environment reproducible and easy to launch with a single command. The candidate must ensure all services can communicate with each other inside the Docker network.
- **Skills Tested**: Docker, Docker Compose, container networking, database seeding in a container, environment variable management.
- **Requirements**:
  - Write a `Dockerfile` for the `frontend` service.
  - Write a `Dockerfile` for the `backend` service.
  - Create a `docker-compose.yml` at the root to define and link the `frontend`, `backend`, and `postgres` services.
  - The `postgres` service must be automatically seeded with the data from the `seed_db/` directory on its first run.
  - The entire application should be launchable with `docker-compose up`.

## 🛠️ Technology Stack

### Frontend
- **Framework**: React 18 + TypeScript
- **UI Library**: Material-UI (MUI) v6
- **State Management**: Redux Toolkit + RTK Query
- **Form Handling**: React Hook Form + Zod validation
- **Build Tool**: Vite
- **Code Quality**: ESLint, Prettier, Husky

### Backend
- **Runtime**: Node.js
- **Framework**: Express.js
- **Database**: PostgreSQL
- **Authentication**: JWT + CSRF protection
- **Password Hashing**: Argon2
- **Email Service**: Resend API
- **Validation**: Zod

### Database
- **Primary DB**: PostgreSQL
- **Schema**: Comprehensive school management schema
- **Features**: Role-based access control, Leave management, Notice system

## 📋 Features

### Core Functionality
- **Dashboard**: User statistics, notices, birthday celebrations, leave requests
- **User Management**: Multi-role system (Admin, Student, Teacher, Custom roles)
- **Academic Management**: Classes, sections, students, class teachers
- **Leave Management**: Policy definition, request submission, approval workflow
- **Notice System**: Create, approve, and distribute notices
- **Staff Management**: Employee profiles, departments, role assignments
- **Access Control**: Granular permissions system

### Security Features
- JWT-based authentication with refresh tokens
- CSRF protection
- Role-based access control (RBAC)
- Password reset and email verification
- Secure cookie handling

## 🔧 Development Guidelines

### Code Standards
- **File Naming**: kebab-case for consistency across OS
- **Import Style**: Absolute imports for cleaner code
- **Code Formatting**: Prettier with consistent configuration
- **Git Hooks**: Husky for pre-commit quality checks

### Project Structure
```
frontend/src/
├── api/           # API configuration and base setup
├── assets/        # Static assets (images, styles)
├── components/    # Shared/reusable components
├── domains/       # Feature-based modules
│   ├── auth/      # Authentication module
│   ├── students/  # Student management
│   ├── notices/   # Notice system
│   └── ...
├── hooks/         # Custom React hooks
├── routes/        # Application routing
├── store/         # Redux store configuration
├── theme/         # MUI theme customization
└── utils/         # Utility functions
```

```
backend/src/
├── config/        # Database and app configuration
├── middlewares/   # Express middlewares
├── modules/       # Feature-based API modules
│   ├── auth/      # Authentication endpoints
│   ├── students/  # Student CRUD operations
│   ├── notices/   # Notice management
│   └── ...
├── routes/        # API route definitions
├── shared/        # Shared utilities and repositories
├── templates/     # Email templates
└── utils/         # Helper functions
```

## 🧪 Testing Instructions

### For Frontend Developers
1. Navigate to the notices section
2. Try to create a new notice with description
3. Verify the description is saved correctly
4. Test form validation and error handling

### For Backend Developers
1. Test all student CRUD endpoints using Postman/curl
2. Verify proper error handling and validation
3. Check database constraints and relationships
4. Test authentication and authorization

### For Blockchain Developers
1. Set up local blockchain environment (Hardhat/Ganache)
2. Deploy certificate smart contract
3. Integrate Web3 wallet connection
4. Test certificate issuance and verification flow

### For Golang Developers
1. Set up the PostgreSQL database using `seed_db/` files.
2. Set up and run the Node.js backend by following its setup instructions.
3. Run the Go service.
4. Use a tool like `curl` or Postman to make a GET request to the Go service's `/api/v1/students/:id/report` endpoint.
5. Verify that the Go service correctly calls the Node.js backend and that a PDF file is successfully generated.
6. Check the contents of the PDF for correctness.

### For DevOps Engineers
1. Ensure Docker and Docker Compose are installed on your machine.
2. From the project root, run the command `docker-compose up --build`.
3. Wait for all services to build and start.
4. Access the frontend at `http://localhost:5173` and verify the application is running.
5. Log in with the demo credentials to confirm that the frontend, backend, and database are all communicating correctly.

## 📚 API Documentation

### Authentication Endpoints
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout
- `GET /api/v1/auth/refresh` - Refresh access token

### Student Management
- `GET /api/v1/students` - List all students
- `POST /api/v1/students` - Create new student
- `PUT /api/v1/students/:id` - Update student
- `DELETE /api/v1/students/:id` - Delete student

### Notice Management
- `GET /api/v1/notices` - List notices
- `POST /api/v1/notices` - Create notice
- `PUT /api/v1/notices/:id` - Update notice
- `DELETE /api/v1/notices/:id` - Delete notice

### PDF Generation Service (Go)
- `GET /api/v1/students/:id/report` - Generate and download a PDF report for a specific student.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For questions and support:
- Create an issue in the repository
- Check existing documentation in `/frontend/README.md` and `/backend/README.md`
- Review the database schema in `/seed_db/tables.sql`

---

**Happy Coding! 🚀**