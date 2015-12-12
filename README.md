# Natural Language Processing: The Longest Compound Word

## What is it?:
Given a file with a list of words like "word.list", this go program will find the longest, case-insensitive, compound-word in the list, which is also a concatenation of other sub-words that exist in the list.

No 3rd party libraries are used.

### Time complexity: O(nlogn) 

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
 ./GoCompoundWords -file DATA-FILE
```
Where DATA_FILE is the path to a file with a new-line separated list of words much like ***word.list***. Just like running from go code, if the **-file** flag is not provided, then this executable will default to using any local file named ***word.list***.

