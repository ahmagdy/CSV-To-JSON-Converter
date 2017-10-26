# CSV-To-JSON-Converter
A little script to convert CSV Files to JSON in easy and flexible way using GoLang.

**The Reason** for this tool is i was in charge to convert a couple of csv files to JSON and 
i didn't find simple thing give me exactly what i expect from such a tool,
and i looked for golang examples but i found most of it using struct for single case, so i built this to be flexible enough for any file.
## How to use this tool:
* After Downloading the go file you can run
`go run main.go -path=C:\\TheFile.csv` or `go run main.go -path C:\\TheFile.csv` or after getting the executable file `myfile.exe -path=C:\\TheFile.csv`
* you will have a new file in the same root as the csv one but with JSON extension.


## Example:
If i have a csv file contains a people data **people.csv**:
```csv
Id,Name,Age
1,Ahmad,21
2,Ali,50
```
and we need to convert it to json file:

`go run main.go -path=C:\Users\User\Desktop\data\src\people.csv`

after writing this command you will get another file in the same directory called **people.json** with the data in the new format:
```json
[
  {
    "Id": 1,
    "Name": "Ahmad",
    "Age": 21
  },
  {
    "Id": 2,
    "Name": "Ali",
    "Age": 50
  }
]
```

and that's it.

## License:
[The MIT License](https://github.com/Ahmad-Magdy/CSV-To-JSON-Converter/blob/master/LICENSE)
