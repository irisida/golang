![](https://github.com/irisida/golang/blob/master/assets/gopher.jpeg)

## Creating and importing a custom package

when creating a custom package, we adhere to the:

- always keep all Go source code files of a package in the same folder.
  With this new folder we can run a `go build` and `go install` and that will place a compiled copy that can be imported into our program. Note to repsect the path.

When importing we can remove the `GOPATH` values and start insoide the pkg folder, so typically that will leave `github.com/<username>/project_or_lirary_name/package_name`

#### Packages in Go

- All source code files of a package should be in the same directory.
- All files in the folder should belong to the same package.
- A file can belong to only 1 package

## Packages

- all files in a package need to be in the same folder.
- each file in the package should declare its package name
- packages come in two types
  - executable (has a main)
  - library or reusable packages
- each project will have a package main because a main is required in an executable
  - within a main package will be a file with a `main()` function which is the automatic entry point of a program.
