# go-modify-attributes

![go-modify-attributes](https://user-images.githubusercontent.com/33835461/112892143-a0845300-90e1-11eb-8247-78d5ecae8875.gif)

## IMPORTANT
```
This project not finished yet. There is a lot of jobs that will be done. 
A lot of functions will be created maybe some of them can be changed. We will see soon.
There will be code refactoring.
```

## Flags (added)
- ``` -file: target file. ```
- ``` -type: convert type. ex: for json -type json. for json and xml -type json,xml. ```
- ``` -override: if the struct field has a tag and you want to change it use it else not. ```
- ``` -write-file:  if set writes output to given -file tag file else it will print stdout. ```

There will be other flag options will be but they are not ready.

Some of them:
<br><br>
Now code converts all struct in the given file. You will have an option to just change given the named struct.
<br>
Like that: go run main.go -file example_modified.go -struct NamedStruct -type json, xml -write-file 
<br><br>
And also camelcase, written struct field names will be transformed like snakecase.
## Examples

``` go run main.go -file example_modified.go -type json ```
<br>
This will print output to stdout with all structs field tags set with JSON tag option.

``` go run main.go -file example_modified.go -type json, xml -write-file ```
<br>
This will write a file to all structs field tags set with JSON and XML tag options.

## Notes

There are similar things like this tool for all programming languages that but I want to create my own tool that will have other features I will add them one by one.
<br><br>
Especially I am creating this tool for one big project that will have a lot of network requests, response, and some calculations processes on JSON,  XML, etc., and also comment lines.
