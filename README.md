# Experimental library to convert arbitrary files to gocode

DO NOT USE!

If you still want to use it, you could try:
````
go get github.com/pscn/file2go
````

And add something like this to your go file:
````
//go:generate file2go -verbose -prefix Index templates/index.html
````

Finally call:
````
go generate
````

This will generate a file ````templates/index.go```` that contains ````templates.index.html````.  To get the content use ````templates.IndexContent()````.

Content is GZIPed and BASE64 encoded.

This is likely to change without further notice so => DO NOT USE ;)