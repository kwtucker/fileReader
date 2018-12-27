# fileReader
Golang File Reader will read file and look for commented out delimiters. It will then return a slice of all comments inside delimiters. It will leave the comments and remove the delimiters.

### Created By:
[Krysler Tucker](https://github.com/PintoKrysler)

[Kevin Tucker](https://github.com/kwtucker)

## Getting Started
* Install Go
* Set up go file structure -> [go file structure](https://golang.org/doc/code.html)
* Grab project through go get

### File Structure
```
+-- src
|   +-- github.com
```

#### Go Get Repo
This will place the project in you github.com directory.
```
go get github.com/kwtucker/fileReader
```
---

#### Good Go comment:
```
// (:This will be read and deleted:)
```

```
// (:This will be read
// and deleted:)
```

### Usage
```go
 import (
   "github.com/kwtucker/fileReader"
 )

 func Example() {
   dataSlice := fileReader.ReadFile("~/Desktop/example.go")
   fmt.Println(dataSlice)
 }
```

#### Fmt Println Response
```
[comment, comment2]
```

## To Contribute
Fork Repo and Submit a pull request.
