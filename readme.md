# URL Shortening Service 🚀

## Overview

This is a simple URL shortening service built using **Golang** and **Redis**. The service provides two main APIs:

1. **Shorten URL API** (`POST /api`) - Accepts a long URL and returns a shortened URL along with expiry time and access count.
2. **Resolve URL API** (`GET /:shortenedURL`) - Redirects users to the original URL. If the shortened URL does not exist, it returns an error.

The service also tracks how many times a shortened URL has been accessed and supports **rate limiting** to prevent abuse.

## Features

- **Generate short URLs** with an expiration time.
- **Redirect short URLs** to their original URLs.
- **Track access count** for each shortened URL.
- **Rate limiting** to prevent excessive requests.
- **Redis database** for fast and efficient storage.
- **Docker & Docker Compose** setup for easy deployment.

---

## API Endpoints

### 1. **Shorten a URL**

- **Endpoint:** `POST /api`
- **Request Body:**
  ```json
  {
    "url": "https://www.youtube.com/watch?v=3ExDEeSnyvE"
  }
  ```
- **Response:**
  ```json
  {
    "url": "https://www.youtube.com/watch?v=3ExDEeSnyvE",
    "short": "localhost:3000/1b14b82",
    "expiry": 24,
    "AccessCount": 0
  }
  ```

### 2. **Resolve a Shortened URL**

- **Endpoint:** `GET /:shortenedURL`
- **Behavior:** Redirects to the original URL.
- **Error Handling:** If the shortened URL does not exist, returns an error.

---

## Deployment and Running the Service

### **Prerequisites**

- **Docker** installed on your machine.

### **Setup Instructions**

1. Clone the repository and navigate into the project directory.
2. Create a **.env** file with the following contents:

   ```ini
   # DB VARS
   DB_ADDRESS="db:6379"

   # API VARS
   API_PORT=3000
   API_HOST="localhost"
   API_RATE_LIMIT_THRESHOLD=10
   API_DOMAIN="localhost:3000"
   ```

3. Ensure that **Docker** is running on your machine.
4. Run the following command to build and start the service:
   ```sh
   docker-compose up --build -d
   ```
   - The `--build` flag ensures the images are rebuilt if needed.
   - The `-d` flag runs the containers in detached mode.

### **Docker Containers**

- **Backend API** container runs the Go application.
- **Redis Database** container stores the shortened URLs.

> **Note:** In the `.env` file, the `DB_ADDRESS` should be set to `db` because Redis will be created with the name `db` in the **Docker network**.

---

## Testing the Service

Once the service is running:

- Use `POST /api` to shorten a URL.
- Use `GET /:shortenedURL` to access the original URL.

Example:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.example.com"}' http://localhost:3000/api
```

This should return a response with a shortened URL that you can use to redirect.

---

## Notes

- The rate limit is controlled by `API_RATE_LIMIT_THRESHOLD` in the `.env` file.
- URLs expire after a set period (default: **24 hours**).
