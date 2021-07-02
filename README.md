<img align="right" width="130" height="130" src="https://cdn.discordapp.com/attachments/349217421082755075/860577315029647380/input-onlinepngtools.png">
<br><br>

# brevity                    

A time management app written in Go.

Intent
---
Build an application that can display different **timetables** by storing **"scheduable" events** and the **dependency relationship** between them.

Motivation
---
The motivation behind this project was to apply the **design patterns** in the book [Design Patterns: Elements of Reusable Object-Oriented Software](https://www.amazon.com/Design-Patterns-Elements-Reusable-Object-Oriented/dp/0201633612) in a context where they have to **interact** with each other.
The book left me with very many questions regarding how the different patterns would **interact in a more complex system**.

Progress
---
The project is **retired**. The **UI layer** of the application has **not been yet completed**.

Usage
---
In order to try this project you have to clone the repository using 
```
git clone https://github.com/MxHonesty/brevity.git
```
After cloning the repository use the `go run` command inside the repository
```
go run brevity
```

Code Examples
---
### Service
The service acts as a *Facade* used by the UI layer for interacting with the underlying app logic.
There are two types of Services inside the app. There is the *LocalService* and the *ProxyService*.
This separation must be made in order for the application to work both in the offline mode and online mode. 
In either mode, the UI does not know which type of service it is using, thus they are interchangeable.
This is achieved by using a common interface for them. 
```go
// Interface for ScheduableService. Declares all common service operations.
type AbsScheduableService interface {
	AddContainer(startYear int, startMonth time.Month, startDay, startHour, startMin int,
		endYear int, endMonth time.Month, endDay, endHour, endMin int)
	RemoveContainer(id uint64) bool
	GetAllContainers() []task.Container
}

// Interface for DependencyService. Declares all common service operations.
type AbsDependencyService interface {
	AddDependency(dependentId uint64, dependentOnId ...uint64) error
	RemoveDependency(id uint64) bool
	GetAllDependencies() []dependency.Dependency
}
```
Now, here is a comparison between a LocalService and ProxyService.
#### ProxyService
```go

// Every method creates a command.Command instance.
// Has a Sender member that is responsible for encoding
// and sending the command.Command to the server.
type ProxyDependencyService struct {
	client *client.Client  // Used for sending the commands to the server.
	commandFactory command.AbstractCommandFactory
}

// Create a new instance of ProxyDependencyService. Gets pointer to client.Client
// as argument.
func NewProxyDependencyService(client *client.Client) *ProxyDependencyService {
	return &ProxyDependencyService{client: client,
		commandFactory: command.Factory{}}
}

func (p ProxyDependencyService) AddDependency(dependentId uint64, dependentOnId ...uint64) error {
	com := p.commandFactory.AddDependency(dependentId, dependentOnId)
	resp, _ := p.client.SendCommand(com)
	return resp.Data.(error)  // We know that the return type is an error.
}
```

#### LocalService
```go
// Struct responsible for serving dependency.Dependency
// related functionalities.
type DependencyService struct {
	taskRepo repository.TaskRepository
	depRepo repository.DependencyRepository
	currentId uint64  // Used for storing the id of the next dependency.Dependency
}

// Create a new DependencyService from a repository.Factory.
func NewDependencyService(depRepo repository.DependencyRepository,
	taskRepo repository.TaskRepository) *DependencyService {
	return &DependencyService{taskRepo: taskRepo, depRepo: depRepo, currentId: 0}
}

// Adds a dependency between a task.Scheduable and a list of task.Scheduable
// elements. Returns an error depending on the success of the operation.
//
// Params:
// 		dependentId - the id of the dependent item
// 		dependentOnId - a list of id's for the items that dependent
// 		depends on
//
// Errors:
//		returns errors if either the dependent is not found or if any of the
//		dependentOn are not found. The operation is not done if such error
// 		occurs. Otherwise returns nil.
func (srv *DependencyService) AddDependency(dependentId uint64, dependentOnId ...uint64) error {
	builder := dependency.NewConcreteBuilder()

	// Find dependentId
	// Find a list of all dependentsOn
	dependent, err := srv.taskRepo.Retrieve(dependentId)
	if err != nil {
		return errors.New("could not find item for dependentId")
	} else {
		builder.SetDependent(dependent)
	}

	for _, itemId := range dependentOnId {
		tsk, err := srv.taskRepo.Retrieve(itemId)
		if err != nil {
			return errors.New(fmt.Sprintf("could not find item for dependentOnId %d", itemId))
		} else {
			builder.AddDependentOn(tsk)
		}
	}


	dep := builder.GetResult(srv.currentId)
	srv.currentId++
	srv.depRepo.Add(dep)
	return nil
}
```

The main difference is that the ProxyService uses a Client instance to forward a request to the server, where an instance of LocalSerivce is waiting.
The server LocalService will execute the operation and then return a Response instance.

### Commands and Responses

The Client and the Server communicate using Commands and Responses.

The following is the **interface** of all Commands.
```go
// Common interface for all Command instances. A command instance stores the
// necessary data for executing it's action.
//
// The execute() method:
// 		Command implements an execute() method.
//		The execute method must connect to a service
// 		once it reaches the back end. A pointer to
//		that service will be provided as an argument.
//		In this case we don't know which service will be used
//		so we provide the whole Session as an argument.
//		Execute() returns a server.Response instance.
//
// More on encoding an interface:
// https://golang.org/src/encoding/gob/example_interface_test.go
type Command interface {
	Execute(session *sessions.Session) response.Response
}
```

All they do is define a method that takes a Session as an argument.
An instance of Command is sent to the Server using the `encoding/gob` package.
```go
// Sends the command.Command instance to the server using gob encoding. Returns a
// non-nil error if the send was not completed successfully
func (c *Client) SendCommand(com command.Command) (response.Response, error) {
	if c.connected {
		errSend := gob.NewEncoder(c.connection).Encode(com)
		if errSend != nil {
			return response.Response{}, errors.New("could not send command")
		} else {
			var data response.Response
			errReceive := gob.NewDecoder(c.connection).Decode(&data)
			if errReceive != nil {
				return response.Response{}, errors.New("decoding error")
			} else {
				return data, nil
			}
		}
	}
	return response.Response{}, errors.New("no connection started")
}
```

The Execute method returns a Response that is sent back to the client.
```go
// A Response type is sent by the server to the client. The data field contains
// the data that the operation returns. This data will be cast to the appropriate
// type. This can be done because every request will know it's response return
// type.
type Response struct {
	Data interface{}
}

// Create a new Response.
func NewResponse(data interface{}) Response {
	return Response{Data: data}
}
```

### Dependency Builder
The Dependency has the following structure
```go
type Dependency struct {
	dependentOn []task.Scheduable
	dependent task.Scheduable
	id uint64  // Unique id of this Dependency
}
```
It maps a "many-to-one" relationship where one event is dependent on a series of other events.
Thus an instance of Dependency can be quite difficult to create. To solve this issue, I implemented a Builder.
Now an instance of Dependency can be created step by step.
```go
// Common interface for all Dependency Builders
type Builder interface {
	SetDependent(scheduable task.Scheduable)
	AddDependentOn(scheduable task.Scheduable)
	RemoveDependentOn(id uint64)
}

type ConcreteBuilder struct {
	built Dependency
}

func (c *ConcreteBuilder) SetDependent(scheduable task.Scheduable) {
	c.built.dependent = scheduable
}

func (c *ConcreteBuilder) AddDependentOn(scheduable task.Scheduable) {
	dependentOn := c.built.DependentOn()
	dependentOn = append(dependentOn, scheduable)
	c.built.SetDependentOn(dependentOn)
}
```

### Repository Factory
The application needs to store events and dependencies between them. Serving this functionality, I created a **Repository**. It is responsible for storing and retrieving items. There a **multiple types** of repositories. There could have an InMemoryRepository, FileRepository, SQLiteRepository, and many more.
Similarly to **how the UI interacted with the logic layer**, the module that uses a Repository does not know what type it's using.
This crates a problem with creating an instance of repository. A problem that can be solved using the Factory pattern.
```go
package repository

// Abstract factory for creating Repositories.
type Factory interface {
	CreateTaskRepository() TaskRepository
	CreateDependencyRepository() DependencyRepository
}

// Implements the Abstract Factory.
// Creates in memory versions of the repositories.
type InMemoryRepositoryFactory struct {}

// Creates a new InMemoryRepositoryFactory.
func NewInMemoryRepositoryFactory() *InMemoryRepositoryFactory {
	return &InMemoryRepositoryFactory{}
}

// Create an in memory version of the TaskRepository.
func (fac *InMemoryRepositoryFactory) CreateTaskRepository() TaskRepository {
	repo := NewScheduleRepository()
	return repo
}

// Create a in memory version of the DependencyRepository.
func (fac *InMemoryRepositoryFactory) CreateDependencyRepository() DependencyRepository {
	repo := NewDepRepository()
	return repo
}
```

External Dependencies
---
* The UI has been written using the [Fyne](https://github.com/fyne-io/fyne) ui toolkit.
* This project uses data structures from the [GoDS](https://github.com/emirpasic/gods) library.
