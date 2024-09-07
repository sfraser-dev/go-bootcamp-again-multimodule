# Go programming

Having a single module (in the project) root directory is the **standard approach.**
Having multiple modules in a project is generally not recommended (unless perhaps a very huge project)

Review of the review.

## Multi-Module Project With Go.work in Root Dir

There is a multi module setup in this project.

Note, having just one top-level go.mod file in an entire Go project is generally the recommended approach, especially if using Go modules (standard way to manage dependencies in Go development).

### Reasons for a single go.mod file

- Simplicity: It keeps dependency management straightforward. All the dependencies for your entire project are tracked in one place.
- Consistent Versioning: All packages within your project share the same dependency versions, ensuring compatibility and avoiding conflicts.
- Easier Builds: You can build your entire project with a single command (go build), and Go will automatically resolve all the necessary dependencies.
- Maintainability: Having a single go.mod file makes it easier to update dependencies and keep your project organized.

### Reasons for a multiple go.mod files

There are some specific scenarios using multiple modules within a larger project:

- Large-Scale Projects: If the project is exceptionally large and complex, breaking it down into smaller modules can help with organization and maintainability.
- Reusable Libraries: If there is a library or component that I intend to use across multiple projects, creating a separate module for it can make it easier to version and share.
- Independent Versioning: If different parts of the project need to be versioned independently, using multiple modules can provide this flexibility.

Important Considerations:

- Increased Complexity: Managing multiple modules adds complexity to the project setup and dependency management.
- Potential Conflicts: Need to be careful to avoid dependency conflicts between modules, especially if they rely on different versions of the same package.
- Module Replacement: If multiple modules used and need to reference local copies of certain modules for development, can use the replace directive in go.mod files.

Best Practice:

For most projects, a single go.mod file at the root level is the
recommended best practice. Can always consider splitting your project into multiple
modules later if the need arises.

### Issues with this multiple go.mod file structure

I have a lot of redunancy in my project. cards/executable/main.go should be my root level main.go.
I should get rid of odd_even/executable/main.go, structs/executable/main.go too. Having all these requires
a go.work file in my root dir to tell MY machine where the modules are. go.work is NOT put into version
control.
