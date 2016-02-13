# Goremark

Simple script to consume a [Remarkjs](http://remarkjs.com/) compliant
markdown file and serve it up as a presentation. This way you can keep
your documentation in git, let Github or Gitlab render it, but easily
present it as a slide deck when needed.

## Installation

```
go get github.com/jmcfarlane/goremark
```

## Usage

```
goremark
```

This will serve up `README.md` and assume all static content is inside
the `docs` directory.

## Other markdown files

```
goremark -f foo.md
```

This will make the same assumption about static content as before.

## Other markdown files in subdirectories

```
goremark -f docs/design.md -p .
```

This time we serve up the design document from **within** the docs
directory. Here we need to pass the `-p` flag to specify the project
directory as the current working directory, else it'll look for the
`docs` directory inside the `docs` folder.

Any valid path will do. For example the following would be a perfectly
valid thing to do:

```
cd /tmp
goremark -f ~/my/repo/docs/best-practices.md -p ~/my/repo
```

## Static content located elsewhere

The default static content folder is `docs`. If you happen to keep
yours in another folder, provide the name via the `-s` flag (without
any leading slashes):

```
goremark -s static
```
