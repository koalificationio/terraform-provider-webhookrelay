# terraform-provider-webhookrelay

[![Build Status](https://travis-ci.org/koalificationio/terraform-provider-webhookrelay.svg?branch=master)](https://travis-ci.org/koalificationio/terraform-provider-webhookrelay)

Terraform provider for [Webhookrelay](https://webhookrelay.com/)

## Using the Provider

For installation instructions and resource documentation check [provider page on terraform registry](https://registry.terraform.io/providers/koalificationio/webhookrelay/latest).

## Developing the Provider

Clone repository to: `$HOME/development/koalificationio/`

```sh
$ mkdir -p $HOME/development/koalificationio/; cd $HOME/development/koalificationio/
$ git clone https://github.com/koalificationio/terraform-provider-webhookrelay
...
```

Enter the provider directory and run `make tools`. This will install the needed tools for the provider.

```sh
$ make tools
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-webhookrelay
...
```

## Testing the Provider

In order to test the provider, you can run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note*: Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc TESTARGS='-run='
```

## Releasing the Provider

```shell
$ fingerprint=$(gpg --with-colons --list-key <key@email> | awk -F: '$1 == "fpr" {print $10;}' | head -n 1)
$ export GPG_FINGERPRINT="${fingerprint}"
$ export GITHUB_TOKEN=xxx
$ git tag v0.x.x -s
$ goreleaser release --rm-dist
```
