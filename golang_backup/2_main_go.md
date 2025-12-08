Packages are a way to organize code in meaningful way in golang.

package main  -> Tells that this code will be treated as executable

Whenever we are accessing anything from a different package it will be in capital form like fmt.Println

Go support UTF-8, some languages only support ASCII [a standard to represent characters]

golang code can be compiled into various OS like windows
`GOOS=windows GOARCH=amd64 go build -o executables.exe` which can then be run without installing go, as it is compiled.

Build is a bit slow process because build doesn't cache anything while go run cache the details.

Package name should be small, because it will be referred in your code with the same name.

We will have a folder, which will be called package. All the .go files inside it will have the package name as 
package folder_name [convention]


If we want to define a variable or function that will be used by other package then, its name will start from Capital Letter [Encapsulation]
