# Quiz API

This is a simple API for a quiz application, where users can answer quiz questions and submit their answers.

## Getting Started

To get started, ensure you have [Go](https://golang.org/doc/install) installed on your machine. Follow the steps below to run the application.

### Prerequisites

- Go 1.18 or later

### Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/Bellzebuth/SuperSimpleQuizz.git
cd SuperSimpleQuizz
```

### Run the server

Server run on port 8080

```
go run main.go serve
```

### Get a quizz

```
go run main.go quiz
```

### Submit answers

```
go run main.go submit "username" 1-A 2-B 3-C
```

### View ranking

```
go run main.go ranking

```
