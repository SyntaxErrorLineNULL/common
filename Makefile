# This Makefile defines common tasks for a Go project, including linting the code,
# running tests, and generating mocks using mockery.

# Target: lint
# Description: Run the Go linter using golangci-lint.
# This checks the code for stylistic issues, potential bugs, and other improvements.
lint: ## Run linter
	golangci-lint run

# Target: test
# Description: Run all tests in the project.
# The -v flag makes the test output verbose, showing detailed information about each test.
test: ## Run tests
	go test -v ./...

# Target: bump
# Description: Increment the project's semantic version by bumping the patch version.
# Fetches the latest tags from the remote repository to ensure the versioning is up to date,
# then uses semver-cli to increment the patch version while maintaining the 'v' prefix.
bump: semver-cli
	@printf "\033[36m%s\033[0m\n" "Bumping version..."
	@git fetch --all --tags
	@semver-cli tags bump -t patch -p v