```ggv``` - go+git version
==========================

Embedding build information (`git`, `hostname`, date and time) into go binaries.

Add to your shell's init script (bash-compatible): 

```sh
function ggv {
    (cd `go list -f {{.Dir}} $1` && echo `git rev-parse HEAD` \(`git rev-parse --abbrev-ref HEAD`\) at `hostname -f` \(`date -R`\))
}

function gogv {
    OPERATION=$1
    shift
    VARIABLE=$1
    shift
    PACKAGE=$1
    shift
    echo \> ggv $PACKAGE
    GGV=`ggv $PACKAGE`
    COMMAND=`echo go $OPERATION -ldflags \'-X \"$VARIABLE=$GGV\"\' $PACKAGE $*`
    echo \> $COMMAND
    eval $COMMAND
}
```


Usage
-----

Define a string variable to hold version information:

```go
package main

var BuildVersion string
```

To build or install your package:

```sh
gogv install main.BuildVersion main.BuildVersion [package]
```

WARNING
-------

THERE IS NO WARRANTY OF ANY KIND, as declared in the license below.

Shell scripts like the one above are prone to break with even slight variations (e.g, in command line parameters, shell version, `$PATH`, installed commands, etc) and some scenarios might damage your system.

So evaluate and use at your own risk. You've been warned!  

LICENSE
-------

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>
