# PrivX SDK Style Guide

The are a few guidelines to follow when developing or contributing to the PrivX Go SDK.
This documentation should help to point out those guidelines and read before opening a pull request in order to keep the SDK consistent.

# Table of content
- [Creating And Initializing New Client](#creating-and-initializing-new-client)
- [Adding New Handlers And Models](#adding-new-handlers-and-models)
	- [Naming Conventions for Handler Functions](#naming-conventions-for-handler-functions)
	- [Parameter Ordering in Handler Functions](#parameter-ordering-in-handler-functions)
	- [Return Types for Handler Functions](#return-types-for-handler-functions)

## Creating And Initializing New Client

When adding a new client (service) to the SDK, this should be done in the `api` directory.
The package name should be the service name that is new and is not yet defined in the `api` directory.
Create a new directory inside `api` with the name of the service and the corresponding files `model.go` which should contain any struct definitions and a `client.go` file containing the API methods and the client constructor.

The client should be initialized as follow:
```go
type Client struct {
	api restapi.Connector
}

func New(api restapi.Connector) *Client {
	return &Client{api: api}
}
```

There is no need to add the service name to any struct or constructor.
It will be verbose enough when the client is eventually imported and initialized.

```go
// initializing new client
client := rolestore.New(api)
```

## Adding New Handlers And Models

### Naming Conventions for Handler Functions

Handler function names should follow a consistent structure:
* Start with a verb describing the action
* Follow the verb with the subject (the source of the endpoint)

Example:
```
	func (c *HostStore) GetHosts() {...} // (get multiple entries)
	func (c *HostStore) GetHost() {...} // (get one entry)
	func (c *HostStore) CreateHost() {...}
	func (c *HostStore) DeleteHost() {...}
	func (c *HostStore) UpdateHost() {...}
	func (c *HostStore) SearchHosts() {...}
```

### Parameter Ordering in Handler Functions

When a function requires parameters, they should be ordered as follows:
* The ID (or multiple IDs, if applicable) of the source must always come first
* Any additional parameters with basic types (e.g., string, int) should follow the IDs
* If required, a structured request body should come after basic-type parameters
* End with variadic options, e.g., opts ...filters.Option

Example:
```
	func (c *HostStore) UpdateHost(hostID string, host *Host, opts ...filters.Option) {...}
```

### Return Types for Handler Functions

For handlers returning multiple entries, always use the generic `ResultSet` type from the `response` package which has the Count and Item fields defined.

Example:
```
func (c *HostStore) GetHosts(...) (*response.ResultSet[Host], error)
```
This standardizes response behavior across all handlers, addressing inconsistencies in earlier SDK versions. Avoid returning raw slices like `[]Host`.

For handlers that return an ID (e.g., when creating entries), use the `Identifier` type from the `response` package.
