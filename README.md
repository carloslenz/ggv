```ggv``` - go+git version
==========================

Embed git version information into go binaries.

Installation
------------

```git``` command must be installed and reachable via ```$PATH```.

```sh
go get github.com/carloslenz/ggv
```

Usage
-----

Insert the line below in any go source file and call ```go generate``` before build:

```go
//go:generate ggv
```

Customization parameters:

* ```-file```: defaults to ```ggver_build_version.go```
* ```-var:``` defaults to ```ggverBuildVersion```

License
-------

MIT