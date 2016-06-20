# fileReader
Golang File Reader will read file and look for commented out delimiters and remove them.

### Created By:
[Krysler Pinto](https://github.com/PintoKrysler)

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

## Usage
In a file wrap commented out single line with (: :).

No multi-line comment!

##### Good Go comment:
```
// (:This will be read and deleted:)
```
##### Bad Go Comment:
```
// (:This will be read
// and deleted:)
```

### Run Script
Run in Shell
```
$ fileReader /path/to/file
```
### Example response will be:
* ![Image of response after command](https://github.com/kwtucker/fileReader/blob/master/examples/commandRes.png)
### [How To Video](https://asciinema.org/a/49570)


## To Contribute
Fork Repo and Submit a pull request.
