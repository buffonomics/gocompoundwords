# NodePrime Challenge


## Requirement:
Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution. Please send us a link to your repo.

Obviously, we are looking for a fresh solution, not based on others' code.

## Requirements
* Go 1.4.x

No 3rd party libraries are used in this project.

## How to Run
**You do not need to change your GOPATH to be able to run these commands.**

To quickly run the code, use
```
go run main.go -file DATA_FILE
```
Where DATA_FILE is the path to a file with a new-line separated list of words much like ***word.list***. If the **-file** flag is not provided, then it will default to using any local file named ***word.list***.



## To build a stand-alone executable

Simply run the command in the project directory
```
go build
```
and then you can take the resulting executable anywhere and run it with 
```
 ./quiz -file DATA-FILE
```
Where DATA_FILE is the path to a file with a new-line separated list of words much like ***word.list***. Just like running from go code, if the **-file** flag is not provided, then this executable will default to using any local file named ***word.list***.



## Testing Hints
To help you score the algorithmic integrity of this submission (and possibly others), here are some thinks to look out for.

* Should work well with unsorted input data.
* Should be case-insensitive, allowing for more flexible real-world processing.

