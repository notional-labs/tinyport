# Tinyport CLI Action
This action makes the `tinyport` CLI available as a Github Action.

## Quick start

Add a new workflow to your repo:

```yml
on: push

jobs:
  help:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Print Help 
        uses: notional-labs/tinyport/actions/cli@develop
        with:
          args: -h 
```
