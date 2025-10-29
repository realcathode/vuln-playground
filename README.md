# Go Vulnerable Lab

This repository is a collection of simple, intentionally vulnerable Go applications. Its purpose is to provide a hands-on environment to learn, test, and demonstrate common web application vulnerabilities.

Each vulnerability is contained in its own directory, with a minimal `main.go` file and a specific `README.md` explaining the flaw.

> These applications are designed to be insecure. Never deploy them in a production or publicly accessible environment. Use them in a sandboxed, isolated network.

## General Usage

Each vulnerability is self-contained. To run one:

1.  Navigate to the vulnerability's directory:
    ```bash
    cd <vulnerability-directory>
    ```
2.  Run the application:
    ```bash
    go run main.go
    ```
    and then the poc script:
    ```bash
    bash <vulnerability_poc>.sh
    ```
3.  The server will start (usually on `http://localhost:8080`). Follow the specific `README.md` in that directory for detailed exploitation instructions.

---

## Vulnerability Catalog

This section details the vulnerabilities available in this lab.

### 1. Server-Side Request Forgery (SSRF)

* **Directory:** ./SSRF/

**Description:**
This vulnerability allows an attacker to coerce the server-side application to make HTTP requests to an unintended destination. The request originates from the server's trusted IP address, allowing an attacker to scan internal networks, access internal-only services (like admin panels on `localhost`), or query cloud metadata endpoints to steal credentials.

**The Flaw (`/SSRF/main.go`):**
The application takes a URL from a query parameter (`?url=...`) and passes it directly into the `http.Get()` function. There is no validation, allow-listing, or deny-listing, allowing an attacker to provide any URL they choose.

---

### 2. Server-Side Template Injection (SSTI)

* **Directory:** ./SSTI/

**Description:**
This vulnerability occurs when user-supplied input is parsed by the server as a template, rather than just being passed to the template as data. This allows an attacker to execute template logic on the server.

This high-severity example demonstrates a critical mistake: passing the entire web framework's context object to the template. This gives the attacker access to all of the context's powerful, built-in methods, leading to Arbitrary File Read.

**The Flaw (`/SSTI/main.go`):** 
The application makes two critical errors. First, it concatenates the user's name parameter directly into the template string. Second, it passes the entire echo.Context (c) to the template's Execute method.

* Arbitrary File Read: `curl "http://localhost:8080/hello?name={{ .File "/etc/passwd" }}"`
