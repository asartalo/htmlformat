HTML Format
===========

A simple pretty printer for HTML, written in Go. This is a fork of [github.com/hermanschaaf/prettyprint](https://github.com/hermanschaaf/prettyprint).

## Installation

```shell
go get github.com/asartalo/htmlformat
```

## Known issues

 - Poor support for self-closing tags, like HTML5 `<br>`. For now, only explicitly self-closing tags like `<br/>` will print correctly.
 - No special cases for `<pre>` and `<code>` tags
 - Comments may or may not work, they are more or less untested
