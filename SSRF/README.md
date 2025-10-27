# Server-Side Request Forgery (SSRF)

This directory contains a simple Go web application that is intentionally vulnerable to Server-Side Request Forgery (SSRF).

## Purpose

This application is for demonstration purposes. It's part of a larger collection of vulnerable apps to demonstrate common security flaws.

## The Vulnerability

The application runs on `http://localhost:8080` and has a single endpoint: `/fetch`.

This endpoint takes a `url` query parameter and makes a GET request to whatever URL is provided. It does not perform any validation, filtering, or sanitization on the user-supplied URL. This allows an attacker to make the server send requests to arbitrary destinations.

The vulnerability is in `main.go`, where `http.Get()` is called directly with user-supplied input:

```go
urlToFetch := r.URL.Query().Get("url")
// ...
resp, err := http.Get(urlToFetch) // VULNERABLE LINE
```

# Impact Varies by HTTP Implementation

The specific impact of an SSRF vulnerability is heavily dependent on the HTTP client library being used and its configuration.Go's default http.Get() does not support the file:/// protocol, preventing this vulnerability from being directly escalated to Local File Inclusion (LFI). Other implementations often do support it.

# Other Important Aspects

- This Go client does follow redirects by default. If a weak allow-list was in place (e.g., "URL must start with https://example.com"), an attacker could host a redirect on that domain to bypass the check and still hit internal targets (http://example.com/redirect -> 302 Found: Location: http://10.10.10.10/).

- This application returns detailed errors (e.g., "connection refused"), which enables effective error-based port scanning. A "blind" SSRF would occur if the application returned a generic "Error" for all failed requests, forcing an attacker to use slower, time-based inference to find open ports.