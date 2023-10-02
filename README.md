# hltvgo - This is an unofficial library of hltv.org (Under development)

Table of contents

- [Usage](#usage)
- [SDKs](#sdks)
- [Go to Python](#go-to-python)
    - [Generate package Python](#generate-package-python)
    - [Generate dist](#generate-dist)
    - [Install local (test)](#install-local-test)
    - [Uninstall local (test)](#uninstall-local-test)

## Usage

A library written in Go with the aim of allowing the extraction of information from the https://www.hltv.org/ website.

:warning: **WARNING:** This is an unofficial library and if it is used in a way that breaks hltv.org rules, your account may be banned or your IP blocked by Cloudflare protection.

Our intention is to write in Go and create an SDK for other programming languages, the first will be python as it has an easy way to convert to a Python library using gopy.

## SDKs

- [Python SDK](hltv_sdk/README.md)

## Go to Python3

#### Generate package Python

```sh
gopy pkg -name=hltvsdk -author="Rich Ramalho" -email="richelton14@gmail.com" -desc="The unofficial HLTV Python API" -url="https://github.com/richecr/hltvgo" -output=hltv_sdk -vm=python3 github.com/richecr/hltvgo github.com/richecr/hltvgo/lib/api  github.com/richecr/hltvgo/lib/operations github.com/richecr/hltvgo/lib/entity
```

#### Generate dist

```sh
python setup.py bdist_wheel
```

To publish to PyPi you need to use this command:

```sh
python -m build
```

#### Install local (test)

```sh
wheel_file=$(ls dist/*.whl | head -n1); pip install $wheel_file
```

#### Uninstall local (test)

```sh
wheel_file=$(ls dist/*.whl | head -n1); pip uninstall $wheel_file
```
