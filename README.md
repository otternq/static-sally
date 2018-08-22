# Static Sally

`static-sally` generates a set of static webpages supporting custom Golang import paths.

## Docker

### Building the docker image

```
$ GOOS=linux GOARCH=arm go build -o ./bin/static-sally .
$ docker build -t static-sally .
```

### Using the docker image

Within a directory containing a [sally.yml](https://github.com/uber-go/sally/blob/master/sally.yaml)
file, run the following script to generate a static site in a `./source` directory.

> the input and output can be altered by flags

```
$ docker run --rm -v "$PWD:/tmp" -w /tmp static-sally
```

## Background

While researching how to support custom Golang imports, I came across
[uber-go/sally](https://github.com/uber-go/sally) which uses a yaml file to support
a small Golang server for routing the imports.

Since I want to limit costs by using Google Cloud Storage, I need to support
Golang imports with a static site. Since sally isn't going to work, I've copied
the code for the underlying yaml and templates to generate the static site.
