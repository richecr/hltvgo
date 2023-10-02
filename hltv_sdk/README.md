# HLTV Python SDK

The unofficial HLTV Python SDK

Table of contents

- [Installation](#installation)
- [Operations](#operations)
    - [GetMatches](#getmatches)

## Installation

```sh
pip install hltvsdk
```

## Operations

### GetMatches

```python
from hltvsdk import operations

ms = operations.GetMatches()
print(ms[0].Id) # '2367046'
```
