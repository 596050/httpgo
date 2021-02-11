# Description

This is a simple http client for abstracting usage of an http package. This improves the process of testing endpoints, allowing for the client to be mocked simply.

# Usage

## Testing

One test per return statement

### Run all tests:

- ```bash
  go test ./... -v | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
  ```

### Run single test:

- ```bash
  cd [*DIRECTORY*]
  ```

- ```bash
  go test -v -run=Test[*NAME OF FUNCTION*] | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
  ```

> Example:
>
> - ```bash
>   cd ./httpgo
>   ```

> - ```bash
>   go test -v -run=TestGetRequestHeaders | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
>   ```

### Run subset of tests:

- ```bash
  go test ./httpgo/... -v | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
  ```
