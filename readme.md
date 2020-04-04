# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Helpers
[![](https://img.shields.io/github/v/release/codemodify/systemkit-helpers?style=flat-square)](https://github.com/codemodify/systemkit-helpers/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-helpers?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-helpers?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-helpers/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-helpers?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-helpers?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-helpers)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-helpers)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-helpers?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-helpers?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-helpers?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-helpers?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-helpers?style=flat-square)


#### A Swiss-Knife collection of helpers for all cases of coding life
#### Supported: Linux, Raspberry Pi, FreeBSD, Mac OS, Windows, Solaris

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-helpers
```
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;																| &nbsp;
---     															| ---
ZipFiles() | Fully qualified path to where the new zip file
AddFileToZip() | Add files to zip
ToNanoSeconds() | To nano seconds
ToMicroSeconds() | To micro seconds
ToMilliSeconds() | To milli seconds
StartOfMonth() | Start of month
EndOfMonth() | End of month
DeltaDays() | Delta days
MillisecondsToDateTime () | Milliseconds to date time
MillisecondsToDate() | Milliseconds to date
isInDictionary() | DSA utils
isInArray() | DSA utils
Is() | Errors
getOSC() | Get the underlying OS command shell
getOSE() | Get the shell/command startup option to execute commands
ExecutableExists() | ExecutableExists
Exec() | Executable
ExecInFolder() | Executable in folder
ExecWithArgs() | Executable with arguments
FileOrFolderExists() | Check File or folder exists
IsFolder() | IsFolder
IsFile() | IsFile
MakeDirTree() | Make directory tree
FileChangedEventHandler() | File changed event handler
WatchFile() | WatchFile
ReadFileAsBytes() | Read file as bytes
ReadFileAsString() | Read file as string
ReadFileAsLines() | Read file as lines
ReadFileAsCSV() | Read file as CSV
ListFolderContent() | List folder content
NewGUID() | NewGUID
NewGUIDWithLength() | NewGUIDWithLength
AsJSONString() | Takes an object and outputs a JSON formatted string
AsJSONStringWithIndentation() | Takes an object and outputs a JSON formatted string
AsJSONStringWithCustomIndentation() | Takes an object and outputs a JSON formatted string
DoPOSTWithJSON() | DoPOSTWithJSON
NewTCPForwarder() | NewTCPForwarder
handleConnection() | Handle connection
Random() | Numbers
ReadNumericFromInterface() | Read numeric from interface
ParseFloat() | Parse float
InitMainPackagePath() | InitMainPackagePath - to be called in `main()`
GetPathForThisPackage() | Get path for this package
GetStructName() | Returns struct name
GetPackageName() | Returns package name
GetPackageNameWithStruct() | Returns package name with struct
GetCallingFuncName() | Returns package + function name at runtime
GetCallStackWithFileAndLineNumber() | Traces a call with line number
CallStack() | Returns function name at runtime
PackageName() | Returns function name at runtime
Trace() | Traces a call with line number
getFuncName() | Return function name
GetLineNumber() | Return line number
GetThisFuncName() | Return getFuncName()
GetThisFuncNameSkip1() | Return getFuncName()
GetThisFuncNameSkip2() | Return getFuncName()
GetThisFuncNameSkip3() | Return getFuncName()
GetThisFuncNameSkip4() | Return getFuncName()
GetThisFuncNameSkip5() | Return getFuncName()
GetThisFuncNameWithLine() | Return function name + line number
GetStackTrace() | Get stack trace
GetStackTrace2() | Get stack trace
NewSSHClient() | New SSH client
RemoteUploadFiles() | Remote upload files
RemoteDownload() | Remote download
RemoteExists() | Remote exists
IsNullOrEmpty() | IsNullOrEmpty
RandomString() | RandomString
Contains() | Contains
ArrayContains() |  ArrayContains
StringToFloat64() |  StringToFloat64
StringToInt() |  StringToInt
StringToInt64() |  StringToInt64
IsRoot() | IsRoot
HomeDir() | HomeDir
UserName() | UserName


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
```go
package main

import (
	"fmt"
	"testing"

	helpersReflect "github.com/codemodify/systemkit-helpers"
)
// Example with GetThisFuncName()
func main() {
	var thisFuncName = fmt.Sprintf("THIS FUNC: %s", helpersReflect.GetThisFuncName())

	t.Log(thisFuncName)
	fmt.Println(thisFuncName)
}
```
> #### For examples see `tests` 