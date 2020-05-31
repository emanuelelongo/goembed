# goembed

Embed static resources in a Go program

## Usage

``` bash
goembed [options] file|directory
```

### Options

  * `-name`   
    The variable name to assign the resource (file mode only)
  * `-noheader`  
    Don't output headers (useful to append to existing file
  * `-package`  
  The package to output in headers (default "main")
  * `-prefix`  
  A prefix for variables (ignored if name is set) (default "resource")

## Example

``` bash
$ goembed --prefix statusIcon myIcons > myFile.go
```
Will write into `myFile.go`:
``` go
// auto-generated

package mypackage


var statusIconRedPng = []byte{0x89, 0x50, 0x4e, 0x47, 0xd, ...
var statusIconOrangePng = []byte{0x89, 0x50, 0x4e, 0x47, 0xd, ...
var statusIconGreenPng = []byte{0x89, 0x50, 0x4e, 0x47, 0xd, ...

