# Authentication

Stateless authentication using **JWT (JSON Web Tokens)**.

## Mechanism
*   **Access Token:** Short-lived, stored in client memory/localStorage. Used for API access.
*   **Refresh Token:** Long-lived, stored in HTTP-only Secure cookie. Used to rotate Access tokens.
*   **Auto-Login:** Signup endpoint returns tokens immediately, logging the user in.

## Endpoints
*   `POST /api/auth/signup`: Create user, return tokens.
*   `POST /api/auth/login`: Validate credentials, return tokens.
*   `POST /api/auth/refresh`: Rotate access token.

## Config
Set `jwt_secret` in `config.json`.
