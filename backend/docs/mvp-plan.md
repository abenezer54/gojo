# **🚀 Gojo MVP Backend – Step-by-Step Detailed Plan**

---

## **🧩 Phase 0: Initial Setup (Day 1–2)**

### **✅ Step 0.1: Set up GitHub Monorepo**

* Create one repo (e.g., `gojo-backend`) under either of your GitHub accounts.

* Add the other person as a **collaborator**.

* Clone it locally on both machines.

bash  
CopyEdit  
`mkdir gojo-backend && cd gojo-backend`  
`git init`

### **✅ Step 0.2: Set Up Folder Structure**

bash  
CopyEdit  
`mkdir user-service property-service booking-service`  
`touch docker-compose.yml .env README.md`

Inside each service:

bash  
CopyEdit  
`cd user-service`  
`go mod init github.com/yourusername/gojo-backend/user-service`

Repeat for the others.

---

## **🛠️ Phase 1: Environment \+ Tooling (Day 2–3)**

### **✅ Step 1.1: Set up Docker \+ Docker Compose**

* Write Dockerfiles for each service

* Create `docker-compose.yml` that:

  * Builds the services

  * Spins up individual PostgreSQL containers

  * Includes Redis (if needed later)

✅ Let me know if you want me to generate full working Dockerfile and docker-compose code.

### **✅ Step 1.2: Add CI/CD (Optional at this stage)**

Use GitHub Actions or skip for now. We can add this later.

---

## **🔑 Phase 2: user-service (Day 4–6)**

### **✅ Step 2.1: Setup Clean Architecture Folders**

Inside `user-service/internal/`:

bash  
CopyEdit  
`controller/   # HTTP handlers`  
`service/      # Business logic`  
`repository/   # DB access`  
`model/        # Request/response structs`  
`config/       # ENV, DB config`  
`middleware/   # JWT, logging, etc`

Create `cmd/main.go` with Gin setup.

### **✅ Step 2.2: Set up PostgreSQL Schema**

Write migration script (e.g., using [golang-migrate](https://github.com/golang-migrate/migrate)):

sql  
CopyEdit  
`CREATE TABLE users (`  
    `id UUID PRIMARY KEY,`  
    `full_name VARCHAR(100),`  
    `email VARCHAR(100) UNIQUE NOT NULL,`  
    `password_hash TEXT NOT NULL,`  
    `role VARCHAR(10) CHECK (role IN ('tenant', 'landlord', 'admin')),`  
    `created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP`  
`);`

### **✅ Step 2.3: Implement Features**

* **Signup**

  * Accepts name, email, password

  * Hash password (use `bcrypt`)

  * Store in DB

* **Login**

  * Verify email/password

  * Return JWT

* **JWT Middleware**

  * Validates token, adds user ID to request context

* **Get Profile**

  * Protected route

  * Returns user info

✅ Want me to generate all user-service code for this?

---

## **🏠 Phase 3: property-service (Day 7–9)**

### **✅ Step 3.1: Folder structure (same as above)**

### **✅ Step 3.2: PostgreSQL Schema**

sql  
CopyEdit  
`CREATE TABLE properties (`  
    `id UUID PRIMARY KEY,`  
    `landlord_id UUID NOT NULL,`  
    `title TEXT NOT NULL,`  
    `description TEXT,`  
    `location TEXT NOT NULL,`  
    `price_per_month NUMERIC NOT NULL,`  
    `type VARCHAR(20) CHECK (type IN ('apartment', 'house', 'studio')),`  
    `created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,`  
    `FOREIGN KEY (landlord_id) REFERENCES users(id)`  
`);`

### **✅ Step 3.3: Endpoints**

* **Add Property** (landlord-only, use JWT)

* **Get All Properties** (open)

* **Get Single Property**

* **Filter by type/location/price**

Use GORM or raw SQL.

---

## **📆 Phase 4: booking-service (Day 10–12)**

### **✅ Step 4.1: Setup as previous**

### **✅ Step 4.2: PostgreSQL Schema**

sql  
CopyEdit  
`CREATE TABLE bookings (`  
    `id UUID PRIMARY KEY,`  
    `tenant_id UUID NOT NULL,`  
    `property_id UUID NOT NULL,`  
    `start_date DATE NOT NULL,`  
    `end_date DATE NOT NULL,`  
    `status VARCHAR(20) CHECK (status IN ('pending', 'confirmed', 'cancelled')) DEFAULT 'pending',`  
    `created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,`  
    `FOREIGN KEY (tenant_id) REFERENCES users(id),`  
    `FOREIGN KEY (property_id) REFERENCES properties(id)`  
`);`

### **✅ Step 4.3: Endpoints**

* **Create Booking** (tenant)

* **View My Bookings**

* **Cancel Booking**

* (Optional) Admin/landlord: Confirm/reject booking

---

## **🔌 Phase 5: Service Communication (Day 13–14)**

### **✅ Option 1: REST Calls Between Services**

From booking → user & property:

* Use internal REST calls with service discovery via `docker-compose` (e.g., call `http://user-service:8000/api/users/:id`)

### **✅ Option 2: Use Message Queues (Optional for MVP)**

If time allows, use Redis Streams or Kafka for async booking notifications.

---

## **🔒 Phase 6: Security and Finishing Touches (Day 15–17)**

* Protect all routes using JWT middleware

* Add role-based access (e.g., only landlords can create properties)

* Add request validation (e.g., with [go-playground/validator](https://github.com/go-playground/validator))

* Use HTTPS in production (via reverse proxy)

* Log errors properly

* Add 404 and error handlers

---

## **📦 Bonus Improvements (Post-MVP)**

* E-signatures & Stripe (automated contracts \+ payments)

* AI-based recommendations (Phase 2\)

* Email notifications

* Deployment (K8s)

---

## **✅ Summary Table**

| Phase | Task | Est. Days |
| ----- | ----- | ----- |
| Phase 0 | GitHub \+ Folder \+ Init Modules | 1–2 |
| Phase 1 | Docker \+ Compose \+ DB Setup | 1–2 |
| Phase 2 | user-service (Auth \+ JWT \+ Profile) | 2–3 |
| Phase 3 | property-service (CRUD) | 2–3 |
| Phase 4 | booking-service (Create, Cancel, View) | 2–3 |
| Phase 5 | Service-to-service communication | 1–2 |
| Phase 6 | Security, Validation, Error Handling | 2–3 |
| **Total** | **Full MVP** | **14–17** |

