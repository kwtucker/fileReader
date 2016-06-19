# fileReader
Golang File Reader will read file and look for commented out delimiters and remove them.

## Getting Started
* Install Go
* Set up go file structure
* Install project

---

## Usage
In a file wrap commented out single line with (: :). Example Go comment:
```
// (:This will be read and deleted:)
```
### Run Script
Run in Shell
```
$ fileReader /path/to/file
```
### response will be:
![Image of response after command](https://github.com/kwtucker/fileReader/blob/master/examples/commandRes.png)
