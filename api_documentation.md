# Loan Management API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. **Get All Loans**

- **Endpoint:** `GET /loans`
- **Description:** Retrieves a list of all loans.
- **Request Example:**
    ```
    GET http://localhost:8080/loans
    ```
- **Response:**
    ```json
    [
        {
            "_id": "66cde2edfa7016c2ebe65098",
            "amount": 600,
            "category": "Personal Loans",
            "description": "hello there",
            "status": "pending",
            "created_at": "2024-08-27T14:00:00Z",
            "updated_at": "2024-08-27T14:00:00Z"
        },
        {
            "_id": "66cde334fa7016c2ebe65099",
            "amount": 1500,
            "category": "Auto Loans",
            "description": "Car loan for Honda Civic",
            "status": "completed",
            "created_at": "2024-08-27T14:10:00Z",
            "updated_at": "2024-08-27T14:20:00Z"
        }
    ]
    ```

### 2. **Get Loan By ID**

- **Endpoint:** `GET /loans/:id`
- **Description:** Retrieves details of a specific loan using its ID.
- **Request Example:**
    ```
    GET http://localhost:8080/loans/66cde2edfa7016c2ebe65098
    ```
- **Response:**
    ```json
    {
        "_id": "66cde2edfa7016c2ebe65098",
        "amount": 600,
        "category": "Personal Loans",
        "description": "hello there",
        "status": "pending",
        "created_at": "2024-08-27T14:00:00Z",
        "updated_at": "2024-08-27T14:00:00Z"
    }
    ```

### 3. **Create Loan**

- **Endpoint:** `POST /loans`
- **Description:** Creates a new loan.
- **Request Body:**
    ```json
    {
        "amount": 600,
        "category": "Personal Loans",
        "description": "hello there",
        "status": "pending"
    }
    ```
- **Request Example:**
    ```
    POST http://localhost:8080/loans
    ```
- **Response:**
    ```json
    {
        "_id": "66cde2edfa7016c2ebe65098",
        "amount": 600,
        "category": "Personal Loans",
        "description": "hello there",
        "status": "pending",
        "created_at": "2024-08-27T14:00:00Z",
        "updated_at": "2024-08-27T14:00:00Z"
    }
    ```

### 4. **Delete Loan**

- **Endpoint:** `DELETE /loans/:id`
- **Description:** Deletes a loan by its ID.
- **Request Example:**
    ```
    DELETE http://localhost:8080/loans/66cde2edfa7016c2ebe65098
    ```
- **Response:**
    ```json
    {
        "message": "Loan successfully deleted"
    }
    ```

### 5. **Approve/Reject Loan**

- **Endpoint:** `PATCH /loans/:id/status`
- **Description:** Updates the status of a loan (e.g., approve or reject).
- **Request Body:**
    ```json
    {
        "status": "rejected"
    }
    ```
- **Request Example:**
    ```
    PATCH http://localhost:8080/loans/66cde334fa7016c2ebe65099/status
    ```
- **Response:**
    ```json
    {
        "_id": "66cde334fa7016c2ebe65099",
        "amount": 1500,
        "category": "Auto Loans",
        "description": "Car loan for Honda Civic",
        "status": "rejected",
        "created_at": "2024-08-27T14:10:00Z",
        "updated_at": "2024-08-27T14:30:00Z"
    }
    ```

### 6. **Get System Logs**

- **Endpoint:** `GET /loans/logs`
- **Description:** Retrieves system logs for activities such as login attempts, loan application submissions, status updates, and password reset actions.
- **Request Example:**
    ```
    GET http://localhost:8080/loans/logs
    ```
- **Response:**
    ```json
    {
        "logs": [
            {
                "timestamp": "2024-08-27T14:00:00Z",
                "event_type": "Login Attempt",
                "description": "User login successful",
                "user_id": "60c72b2f5f1b2c001c8e4df3"
            },
            {
                "timestamp": "2024-08-27T14:05:00Z",
                "event_type": "Loan Application Submission",
                "description": "User submitted a new loan application",
                "user_id": "60c72b2f5f1b2c001c8e4df3",
                "loan_id": "66cde2edfa7016c2ebe65098"
            },
            {
                "timestamp": "2024-08-27T14:30:00Z",
                "event_type": "Loan Application Status Update",
                "description": "Admin rejected loan application",
                "loan_id": "66cde334fa7016c2ebe65099"
            }
        ]
    }
    ```

---



##POSTMAN DOCUMENTATION: [https://blogposts.postman.co/documentation/37289771-565ab4d9-44d2-485c-a693-0bba23b6eed1/publish?workspaceId=35b4fba7-9877-4892-990f-6e1297a0924d]