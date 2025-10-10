# Naming Functions in Go

## Overview
Good function names are crucial for code readability and maintainability. Go has specific conventions and best practices for naming functions that make code more intuitive and easier to understand. This guide covers everything you need to know about naming functions in Go.

## General Naming Conventions

### Use CamelCase
Go uses camelCase (mixed caps) for function names, not snake_case.

```go
// Good
func calculateTotal() int { return 0 }
func getUserProfile() *User { return nil }
func validateEmailAddress() bool { return true }

// Bad
func calculate_total() int { return 0 }
func get_user_profile() *User { return nil }
func validate_email_address() bool { return true }
```

### Use Descriptive Verbs
Functions perform actions, so use verbs that describe what they do.

```go
// Good - action verbs
func parseJSON(data []byte) error
func validateInput(s string) bool
func sendEmail(to, subject, body string) error
func createUser(name, email string) (*User, error)
func deleteAccount(id int) error
func updateSettings(cfg Config) error

// Avoid - nouns or unclear names
func json(data []byte) error        // What does this do?
func input(s string) bool           // Validate? Parse? What?
func email(to, subject string)      // Send? Read? Format?
func user(name, email string)       // Create? Update? Get?
```

### Keep Names Concise but Clear
Balance brevity with clarity. Names should be as short as possible while remaining clear.

```go
// Good
func findUserByID(id int) (*User, error)
func encodeJSON(v interface{}) ([]byte, error)
func startServer(port int) error

// Too verbose
func findUserByIDAndReturnUserPointerOrError(id int) (*User, error)
func encodeDataStructureToJSONByteArray(v interface{}) ([]byte, error)

// Too short - unclear
func find(id int) (*User, error)    // Find what?
func enc(v interface{}) ([]byte, error)  // Encode how?
func start(port int) error          // Start what?
```

## Exported vs Unexported Functions

### Exported Functions (Public)
Start with an uppercase letter. These are accessible from other packages.

```go
package user

// Exported - can be called from other packages
func CreateUser(name, email string) (*User, error) {
    return &User{Name: name, Email: email}, nil
}

func ValidateEmail(email string) bool {
    return strings.Contains(email, "@")
}

func GetUserByID(id int) (*User, error) {
    // Implementation
    return nil, nil
}
```

Usage from another package:
```go
package main

import "myapp/user"

func main() {
    u, err := user.CreateUser("Alice", "alice@example.com")
    if err != nil {
        // Handle error
    }
    _ = u
}
```

### Unexported Functions (Private)
Start with a lowercase letter. Only accessible within the same package.

```go
package user

// Unexported - only usable within user package
func validatePassword(password string) bool {
    return len(password) >= 8
}

func hashPassword(password string) (string, error) {
    // Implementation
    return "", nil
}

func sendWelcomeEmail(email string) error {
    // Implementation
    return nil
}

// Exported function using unexported helpers
func CreateUser(name, email, password string) (*User, error) {
    if !validatePassword(password) {
        return nil, errors.New("password too weak")
    }
    
    hash, err := hashPassword(password)
    if err != nil {
        return nil, err
    }
    
    user := &User{Name: name, Email: email, PasswordHash: hash}
    
    if err := sendWelcomeEmail(email); err != nil {
        // Log but don't fail
        log.Printf("failed to send welcome email: %v", err)
    }
    
    return user, nil
}
```

### When to Export vs Keep Private

**Export when**:
- The function is part of your package's public API
- Other packages need to call the function
- It's a core feature of your package

**Keep private when**:
- The function is an implementation detail
- It's a helper function used internally
- It might change frequently
- It's not stable enough for public use

```go
package calculator

// Exported - main API
func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

// Unexported - internal helper
func validate(n int) error {
    if n < 0 {
        return errors.New("negative numbers not allowed")
    }
    return nil
}

// Unexported - implementation detail
func cache(key string, value int) {
    // Cache implementation
}
```

## Common Naming Patterns

### Getters - No "Get" Prefix
In Go, don't use "Get" prefix for simple getters.

```go
type User struct {
    name  string
    email string
    age   int
}

// Good - no "Get" prefix
func (u *User) Name() string {
    return u.name
}

func (u *User) Email() string {
    return u.email
}

func (u *User) Age() int {
    return u.age
}

// Bad - unnecessary "Get" prefix
func (u *User) GetName() string {
    return u.name
}

func (u *User) GetEmail() string {
    return u.email
}
```

**Exception**: Use "Get" when the operation is not trivial:
```go
// OK - involves computation or I/O
func (u *User) GetProfile() (*Profile, error) {
    // Fetches from database, makes API call, etc.
    return fetchProfileFromDB(u.ID)
}

func (c *Cache) GetOrCompute(key string) (string, error) {
    // Complex operation
    return "", nil
}
```

### Setters - Use "Set" Prefix
Use "Set" prefix for setter methods.

```go
type Config struct {
    timeout time.Duration
    retries int
}

func (c *Config) SetTimeout(d time.Duration) {
    c.timeout = d
}

func (c *Config) SetRetries(n int) {
    c.retries = n
}
```

### Boolean Functions - Use Question Form
Name boolean functions to read like a question.

```go
// Good - reads naturally
func (u *User) IsActive() bool
func (u *User) HasPermission(p Permission) bool
func (f *File) Exists() bool
func (s *Service) IsRunning() bool
func (c *Connection) IsClosed() bool

// Usage reads like English
if user.IsActive() && user.HasPermission(AdminPermission) {
    // ...
}

// Avoid
func (u *User) Active() bool        // Unclear
func (u *User) Permission() bool    // Confusing
func (f *File) FileExists() bool    // Redundant
```

### Constructor Functions - Use "New" Prefix
Constructor functions typically start with "New".

```go
// Simple constructor
func NewUser(name, email string) *User {
    return &User{
        Name:  name,
        Email: email,
    }
}

// Constructor with error handling
func NewConnection(url string) (*Connection, error) {
    conn, err := dial(url)
    if err != nil {
        return nil, fmt.Errorf("connection failed: %w", err)
    }
    return &Connection{conn: conn}, nil
}

// Multiple constructors
func NewDefaultConfig() *Config {
    return &Config{
        Timeout: 30 * time.Second,
        Retries: 3,
    }
}

func NewConfig(timeout time.Duration, retries int) *Config {
    return &Config{
        Timeout: timeout,
        Retries: retries,
    }
}
```

### Conversion Functions - Use "To" or type name

```go
// Converting between types
func (u *User) ToJSON() ([]byte, error) {
    return json.Marshal(u)
}

func (s *Session) String() string {  // Implements fmt.Stringer
    return fmt.Sprintf("Session{ID: %s}", s.ID)
}

// Type conversions
func IntToString(n int) string {
    return strconv.Itoa(n)
}

func StringToInt(s string) (int, error) {
    return strconv.Atoi(s)
}
```

## Function Naming by Purpose

### CRUD Operations

```go
// Create
func CreateUser(u *User) error
func InsertRecord(r *Record) error
func AddItem(item *Item) error

// Read
func GetUser(id int) (*User, error)
func FindUsers(filter Filter) ([]*User, error)
func ListUsers(limit, offset int) ([]*User, error)
func QueryUsers(query string) ([]*User, error)

// Update
func UpdateUser(u *User) error
func ModifySettings(s *Settings) error
func SaveUser(u *User) error

// Delete
func DeleteUser(id int) error
func RemoveItem(id int) error
func DestroySession(token string) error
```

### Validation

```go
func ValidateEmail(email string) error
func ValidatePassword(password string) error
func IsValid(input string) bool
func CheckFormat(s string) error
func VerifySignature(data, signature []byte) bool
```

### Processing

```go
func ProcessPayment(amount float64) error
func HandleRequest(r *Request) (*Response, error)
func ParseInput(input string) (*Result, error)
func TransformData(data []byte) ([]byte, error)
func EncodeMessage(msg *Message) ([]byte, error)
func DecodeMessage(data []byte) (*Message, error)
```

### I/O Operations

```go
func ReadFile(path string) ([]byte, error)
func WriteFile(path string, data []byte) error
func LoadConfig(path string) (*Config, error)
func SaveConfig(path string, cfg *Config) error
func FetchData(url string) ([]byte, error)
func SendRequest(req *Request) (*Response, error)
```

## Method Names

Methods follow the same rules but are shorter since they have context from the receiver.

```go
type Database struct {
    conn *sql.DB
}

// Good - receiver provides context
func (db *Database) Connect() error
func (db *Database) Query(sql string) (*Rows, error)
func (db *Database) Close() error

// Redundant - receiver already provides context
func (db *Database) ConnectToDatabase() error  // "Database" is redundant
func (db *Database) QueryDatabase(sql string) (*Rows, error)
func (db *Database) CloseDatabase() error
```

### Method Receiver Naming

Use short, consistent receiver names (usually 1-2 letters).

```go
// Good
func (u *User) Save() error
func (c *Client) Connect() error
func (s *Server) Start() error
func (db *Database) Query() error

// Common receiver names:
// s for Server, Service, Session
// c for Client, Config, Connection
// u for User
// db for Database
// r for Request
// w for Writer
// h for Handler

// Avoid
func (user *User) Save() error       // Too long
func (this *User) Save() error       // Not Go style
func (self *User) Save() error       // Not Go style
func (u1 *User) Save() error         // Confusing
```

## Package-Qualified Names

Remember that function names will be qualified by package name when used, so avoid redundancy.

```go
// Package: user
package user

// Good - package name provides context
func Create(u *User) error      // Called as user.Create()
func Find(id int) (*User, error) // Called as user.Find()
func List() ([]*User, error)     // Called as user.List()

// Redundant - "user" appears twice
func CreateUser(u *User) error      // Called as user.CreateUser()
func FindUser(id int) (*User, error) // Called as user.FindUser()
func ListUsers() ([]*User, error)    // Called as user.ListUsers()
```

However, exported types don't have this issue:
```go
package user

type User struct { ... }  // user.User is fine

// This is OK because it helps when reading
func NewUser() *User     // user.NewUser() - clear
func CreateUser() *User  // user.CreateUser() - also clear
```

## Avoid Common Anti-Patterns

### Don't Use Generic Names

```go
// Bad
func DoStuff(data []byte) error
func Process(x interface{}) error
func Handle(v interface{}) error
func Manage(item interface{}) error

// Good
func ParseJSON(data []byte) (*Config, error)
func ValidateInput(s string) error
func HandleHTTPRequest(r *http.Request) error
func UpdateInventory(item *Item) error
```

### Avoid Unclear Abbreviations

```go
// Bad
func ProcReq(r *Req) (*Resp, error)  // What?
func GetUsrInf(id int) (*UI, error)  // Unclear
func ChkAuth(t string) bool          // Hard to read

// Good
func ProcessRequest(r *Request) (*Response, error)
func GetUserInfo(id int) (*UserInfo, error)
func CheckAuthentication(token string) bool
```

### Don't Mix Naming Styles

```go
// Inconsistent
func getUser(id int) *User
func FindUser(name string) *User
func user_list() []*User
func CreateNewUserAccount(email string) *User

// Consistent
func Get(id int) *User
func Find(name string) *User
func List() []*User
func Create(email string) *User
```

## Special Cases

### Interface Implementation Methods
When implementing standard interfaces, use the expected method names.

```go
// io.Reader interface
func (r *MyReader) Read(p []byte) (n int, err error) {
    // Must be named "Read"
}

// io.Writer interface
func (w *MyWriter) Write(p []byte) (n int, err error) {
    // Must be named "Write"
}

// fmt.Stringer interface
func (u *User) String() string {
    // Must be named "String"
}

// error interface
func (e *MyError) Error() string {
    // Must be named "Error"
}
```

### Test Functions
Test function names must start with "Test" and take *testing.T.

```go
func TestUserCreation(t *testing.T) {
    // Test implementation
}

func TestValidateEmail(t *testing.T) {
    // Test implementation
}

func TestHTTPHandler(t *testing.T) {
    // Test implementation
}
```

### Benchmark Functions
Benchmark names must start with "Benchmark".

```go
func BenchmarkSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sort.Ints(data)
    }
}
```

### Example Functions
Example names must start with "Example".

```go
func ExampleUser_Create() {
    user := NewUser("Alice", "alice@example.com")
    fmt.Println(user.Name)
    // Output: Alice
}
```

## Real-World Examples

### HTTP Handler Functions

```go
package handlers

import "net/http"

// Good naming for HTTP handlers
func HandleHome(w http.ResponseWriter, r *http.Request)
func HandleLogin(w http.ResponseWriter, r *http.Request)
func HandleLogout(w http.ResponseWriter, r *http.Request)
func HandleUserProfile(w http.ResponseWriter, r *http.Request)

// Or simpler when context is clear
func Home(w http.ResponseWriter, r *http.Request)
func Login(w http.ResponseWriter, r *http.Request)
func Logout(w http.ResponseWriter, r *http.Request)
func UserProfile(w http.ResponseWriter, r *http.Request)
```

### Database Operations

```go
package repository

type UserRepository struct {
    db *sql.DB
}

// Clear CRUD operations
func (r *UserRepository) Create(u *User) error
func (r *UserRepository) FindByID(id int) (*User, error)
func (r *UserRepository) FindByEmail(email string) (*User, error)
func (r *UserRepository) Update(u *User) error
func (r *UserRepository) Delete(id int) error
func (r *UserRepository) List(limit, offset int) ([]*User, error)
```

### Service Layer

```go
package service

type EmailService struct {
    client *smtp.Client
}

func (s *EmailService) Send(to, subject, body string) error
func (s *EmailService) SendBatch(emails []*Email) error
func (s *EmailService) Validate(email string) bool
```

## Documentation

Always document exported functions.

```go
// CreateUser creates a new user with the given name and email.
// It returns an error if the email is invalid or already exists.
func CreateUser(name, email string) (*User, error) {
    // Implementation
    return nil, nil
}

// IsValid reports whether the email address is valid.
func IsValid(email string) bool {
    return strings.Contains(email, "@")
}
```

## Summary

### Naming Checklist

- [ ] Use camelCase, not snake_case
- [ ] Start with verb (e.g., Create, Get, Update, Delete)
- [ ] Export only what needs to be public (uppercase)
- [ ] Keep names concise but clear
- [ ] Avoid redundancy with package names
- [ ] Use "New" for constructors
- [ ] No "Get" prefix for simple getters
- [ ] Use "Is", "Has", "Can" for booleans
- [ ] Document all exported functions
- [ ] Be consistent across your codebase

### Quick Reference

| Purpose | Pattern | Example |
|---------|---------|---------|
| Constructor | New* | NewUser, NewConfig |
| Getter | PropertyName | Name(), Email() |
| Setter | Set* | SetName(), SetEmail() |
| Boolean | Is*, Has*, Can* | IsActive(), HasPermission() |
| CRUD | Create, Get, Update, Delete | CreateUser(), GetUser() |
| Validation | Validate*, Is*, Check* | ValidateEmail(), IsValid() |
| Conversion | To*, String | ToJSON(), String() |

## Resources

- [Effective Go - Names](https://go.dev/doc/effective_go#names)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#naming)
- [Go Blog: Package names](https://go.dev/blog/package-names)

## Next Steps

- Review your existing code for naming consistency
- Refactor poorly named functions
- Document all exported functions
- Practice naming with new code
- Study well-named packages in the standard library
