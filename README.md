# Description

This is a simple http client which can abstract the usage of an http package

# Usage

## Run all tests:

- go test -v | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

## Run single test:

- cd [*DIRECTORY*]
- go test -v -run=Test[*NAME OF FUNCTION*] | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
  > Example:
  >
  > - cd ./httpgo
  > - go test -v -run=TestGetRequestHeaders | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
