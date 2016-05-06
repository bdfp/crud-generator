# Crud-generator (WIP)

This is a cmd tool to generate APis for CRUD operations  

This can be used to generate such CRUD operations for any languages and any package structure  
by implementing the required interface

## Architecture

There are three main interfaces  
*  **Reader**  defines how the program will read the given models 
*  **Parser**  defines how the syntax of templates are understood
*  **Writer**  defines how the program writes to the file system

The contents of `template/[languageName]/*` define the syntax of target contents generated

## Extending

### Implementing for different package structure
Modify the `Reader` and `Writer` interface as required

### Implementing for diffrent language
* Modify the `Parser` interface according to syntax of the language
* Modify the `Reader` and `Writer` interface as per the package strucutre

## Default Usage
 
 * Install the generator.go file
 * Run `generator [project_directory]`
 
The program reads all the struct present in the model/*.go files  
and generates the corresponding db and api Controller functions as per following

* `[proj_dir]/db/[modelName].go` Contains all the db operations code
* `[proj_dir]/rest/[modelName]handler.go` Contains all the REST API handlers 
* `[proj_dir]/rest/utils.go` utility functions used for API handlers
*  `[proj_dir]/model/[modelName.go]` Contains all the representational code

 

## Target
* Generate coress front end code
* For Web React-Redux-Universal Application
* For Android some MVP based App
* Add Permission in routes
* This can be used to generate routes alowed only after login 
*  Add corres JWT authentication
*  Use a config file to specify which routes are pemitted to all and which are not