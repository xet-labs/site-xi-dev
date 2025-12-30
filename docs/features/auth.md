# Authentication

XetIndustries uses a secure, stateless authentication system based on **JWT (JSON Web Tokens)**.

## Features
*   **Stateless:** No server-side session storage required.
*   **Dual Tokens:**
    *   **Access Token:** Short-lived, used for API authorization. stored in memory/localStorage.
    *   **Refresh Token:** Long-lived, used to obtain new access tokens. Stored in a secure, HTTP-only cookie.
*   **Auto-Login:** Users are automatically logged in upon signup.
*   **Security:** Passwords are hashed using `bcrypt` before storage.

## Usage
*   **Signup:** `POST /api/auth/signup` - Creates a user and returns tokens.
*   **Login:** `POST /api/auth/login` - Authenticates credentials and returns tokens.
*   **Refresh:** `POST /api/auth/refresh` - Uses the cookie to get a new access token.

## Configuration
Configure JWT secret and expiry in `config/config.json`:
```json
"auth": {
  "jwt_secret": "your-secret-key"
}
```
