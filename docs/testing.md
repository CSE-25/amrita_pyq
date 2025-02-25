# Running unit tests for Amrita PYQ CLI Tool

## Run Unit Test

To run all the unit tests for this project use the following command:

```bash
go test ./... -v
```

You will be able to see the output and status for each unit test available.

To run a unit test specific to a package, go to the package location and run:

```bash
go test -run "PACKAGE_NAME"
```

## Run Lint check

```bash
golangci-lint run
```

**Note:** This command will return an output only when it finds an error. Otherwise, no visible change occurs.

Good Luck developing and testing the application!!
