# To Do App

## Getting Started

Use the below command to run the app. **MUST** be run from the root of the project

`go run ./cmd/todo/.` use one of the following flags to choose which app to run.

`--cli` cli to do app using in memory store  
`--web` web app using in memory store  
`--cli2` basic cli using in memory store (WIP)

Access the web app at [http://localhost:8081/todos](http://localhost:8081/todos).

## Run tests

`go run ./...` from the root.

## Fail

Obviously having 3 copies of `view.html` is awful. Wasted too much time trying to resolve issues of importing when running tests / the app itself :sob:
