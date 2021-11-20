# REST Servers in Go: Part 1 - standard library

- [Eli Bendersky](https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library/)
- [Code](https://github.com/eliben/code-for-blog/tree/master/2021/go-rest-servers/stdlib-basic)

## The task

In our case, the server is a simple backend for a task management application; it presents the following REST API to clients:

| Request | Path                   | Description |
| ------- | ---------------------- | ----------- |
| POST    | /task/                 | create a task, returns ID |
| GET     | /task/\<taskid>        | returns a single task by ID |
| GET     | /task/                 | returns all tasks |
| DELETE  | /task/\<taskid>        | delete a task by ID |
| GET     | /tag/\<tagname>        | returns list of tasks with this tag |
| GET     | /due/\<yy>/\<mm>/\<dd> | returns list of tasks due by this date |

Our server supports GET, POST and DELETE requests, some of them with several potential paths.
The parts between angle brackets `<...>` denote parameters that the client supplies as part of the request; for example, `GET /task/42` is a request to fetch the task with ID 42, etc.
Tasks are uniquely identified by IDs.

The data encoding is JSON. In `POST /task/` the client will send a JSON representation of the task to create.
Similarly, everywhere it says the server "returns" something, the returned data is encoded as JSON in the body of the HTTP response.

## The model

The `taskstore` package is the model (or the "data layer") for our server.
This is a simple abstraction representing a database of tasks; here is its API:

```go
type Task struct {
  Id   int       `json:"id"`
  Text string    `json:"text"`
  Tags []string  `json:"tags"`
  Due  time.Time `json:"due"`
}

func New() *TaskStore

// CreateTask creates a new task in the store.
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int

// GetTask retrieves a task from the store, by id. If no such id exists, an error is returned.
func (ts *TaskStore) GetTask(id int) (Task, error)

// DeleteTask deletes the task with the given id. If no such id exists, an error is returned.
func (ts *TaskStore) DeleteTask(id int) error

// DeleteAllTasks deletes all tasks in the store.
func (ts *TaskStore) DeleteAllTasks() error

// GetAllTasks returns all the tasks in the store, in arbitrary order.
func (ts *TaskStore) GetAllTasks() []Task

// GetTasksByTag returns all the tasks that have the given tag, in arbitrary order.
func (ts *TaskStore) GetTasksByTag(tag string) []Task

// GetTasksByDueDate returns all the tasks that have the given due date, in arbitrary order.
func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task
```

```go
```

```go
```