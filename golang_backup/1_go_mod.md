<!-- Questions to MySelf for revision -->

what is a package in go? and what is a module? 
what are different types of programs in go?
what are the components we. write in go.mod file
what is the meaning of each component.
Try to think of the syntax, as well.





















Module is a collection of go packages.

There are two types of go programs
- Executables, that we run
- Library/Package, that is imported in other projects [fmt package]


// name of the package/project

```
module golang
// name of the package/project

go 1.23.3
// version of go on which go project depends on, not the one which is installed
```

go.mod will have also have all the dependencies on which the go project depends on.

go mod init example/mymodule
Example is a place from where the module can be downloaded from

Three files
- module fullPlace-WhereThis-PackageExist
- go versionNumber on which our module/package is dependent
- dependecies

require (
    github.com/aws/aws-sdk-go v1.41.141
    .
    .
    .
)



<!-- Questions to MySelf for revision -->
