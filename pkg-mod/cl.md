### go mod init <module-name>
    Initializes a new module, creating a `go.mod` file with the specified module name.
Ex: go mod init github.com/awangelo/my-app

### go mod tidy
    Removes unused dependencies and adds any missing necessary dependencies to the `go.mod` file.


### go get <package>
    Downloads and installs the specified package, and updates the `go.mod` and `go.sum` files with the new dependency.
Ex: go get github.com/gin-gonic/gin

### go get -u <package>
    Updates the package and its dependencies to the latest compatible versions.