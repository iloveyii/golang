Go Tutorial
==========

Go is a faster C++ like language which can be either interpretted or compiled.
We will teach you Go language in seven days (chapters).

# DAY 1 
## Installations
- A go play ground [link](https://play.golang.org/)
- We will use docker container for the run time enviroment.
- Write a docker file from [official image](https://hub.docker.com/_/golang)
- Build image `docker build -t golang-app1 .`
- Run it `docker run -it --rm -v /home/alex/projects/golang/day1/:/go/src/app/day1/ --name golang-running-app1 golang-app1`
- You will in the container shell type `go version`
- Check go environment variables `go env`

## The basics
- Hello world - 
### Data types
- These are the data types in go
```
string
bool
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint84
byte - alias for uint8
rune - alias for int32
float32 float64
complex64 complex128

```
### Import
- import keyword is used to import a package
### Function
- Similar to other programming languages
- Can pass by value or ref
### Arrays
- Collections of similar data types, 0 indexed
- Can slice like in python arr[start: stop]
- Its syntax is different `rolls := [2]int{1,2}`
### Conditions if else
- Similar to other programming languages
### Loops
- For loop can be used as either while or for
### Maps
- It is like an associative array
- Iterate over using `range` keyword
- Its syntax is very different ` emails := map[string]string{"name":"em@il.com"`}

# DAY 2
### Pointer

### Repository, modules and packages 
- [docs](https://golang.org/doc/code)
- [packages](https://pkg.go.dev/)
### Struct
### Web

# DAY 3
### Database

# DAY 4
# DAY 5
# DAY 6
# DAY 7
## Workshops
### FizzBuzz
### A simple web app with Database
### Serve React app
### Google API app