## Goals
- Based on the HTTP interface gin, automated build markdown interface documentation

## Usages

- Based on the ast implementation, through the ast parsing source code structure.
  Extract the parameter name, type, comment, and other fields and generate them into the md document.

- You need to add the necessary comments `@request`, `@response`, and exposed interface names before you
    build the document. Please refer to [examples/api.go](https://github.com/lanshipeng/gin-gen-markdown/blob/master/examples/api.go) for details.

- The directory tree path can be specified with `mark_down` command line argument,
  and `prefix` command line argument specifies the file prefix or suffix to traverse,
  `domain` set the url path. If you do not set these variables, they will use the default values.


## Quick start
```bash
cd $GOPATH/src/github.com/lanshipeng/gin-gen-markdown
go install

# check installation by running:
gin-gen-markdown -h

# generate doc
gin-gen-markdown doc --mark_down=. --prefix=api.go --domain=http://127.0.0.1/
```
