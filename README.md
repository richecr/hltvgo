# hltvgo


### Go to Python3

#### Generate package Python

```sh
gopy pkg -name=hltvsdk -author="Rich Ramalho" -email="richelton14@gmail.com" -desc="The unofficial HLTV Python API" -url="https://github.com/richecr/hltvgo" -output=hltv_sdk -vm=python3 github.com/richecr/hltvgo github.com/richecr/hltvgo/lib/api  github.com/richecr/hltvgo/lib/operations github.com/richecr/hltvgo/lib/entity
```

#### Generate dist

```sh
python setup.py bdist_wheel
```

#### Install local (test)

```sh
wheel_file=$(ls dist/*.whl | head -n1); pip install $wheel_file
```

#### Uninstall local (test)

```sh
wheel_file=$(ls dist/*.whl | head -n1); pip uninstall $wheel_file
```
