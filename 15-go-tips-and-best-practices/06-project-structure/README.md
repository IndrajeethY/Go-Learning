# Project Structure

## Recommended Project Structure

Here is a recommended directory structure for your Go projects:

```
myproject/
├── cmd/
│   └── myproject/
│       └── main.go
├── internal/
│   ├── mymodule/
│   │   └── mymodule.go
│   └── myothermodule/
│       └── myothermodule.go
├── pkg/
│   └── mypackage/
│       └── mypackage.go
├── api/
│   └── openapi.yaml
├── web/
│   ├── static/
│   └── templates/
├── test/
│   └── myproject_test.go
├── go.mod
└── README.md
```

## Explanation of Each Directory

- **cmd/**: This directory holds the main applications for your project. Each application should be in a separate subdirectory.
- **internal/**: This directory contains private application and library code. This code is not meant to be imported by other projects.
- **pkg/**: This directory is for code that is intended to be used by external applications (e.g., `/pkg/mypackage`).
- **api/**: This directory is for API definitions, such as OpenAPI specifications.
- **web/**: This directory contains web-related files, such as static assets and HTML templates.
- **test/**: This directory is for additional test files.
- **go.mod**: This file defines the module's properties, including dependencies.
- **README.md**: This file provides an overview of the project.

## Further Reading & References

- Official Go Blog: [Go Blog](https://blog.golang.org)
- Go Documentation: [Go Docs](https://golang.org/doc/)
- Go Community: [Go Community](https://golang.org/help/)
- Go Best Practices: [Go Best Practices](https://github.com/golang/go/wiki/CodeReview)
