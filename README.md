# Golang Multi Module Gradle Style

## Download

```shell
$ git clone http://github.com/anon-org/pet
$ cd pet
```

## Run

in one terminal run

```shell
$ make service
```

in another terminal run

```shell
$ make rpc
```

## Example

See RPC usage example [here](rpc/example/main.go)

### service

```shell
$ make service
listening at :8000
2022/04/21 21:54:52 storing satyr
2022/04/21 21:54:52 fetching...
2022/04/21 21:54:54 storing bat
2022/04/21 21:54:54 fetching...
2022/04/21 21:54:55 storing hound
2022/04/21 21:54:55 fetching...
```

### rpc

```shell
$ make rpc
[
  {
    "ID": "c559b192-d428-451f-8ef7-4c36e3576068",
    "Name": "satyr"
  }
]

$ make rpc
[
  {
    "ID": "c559b192-d428-451f-8ef7-4c36e3576068",
    "Name": "satyr"
  },
  {
    "ID": "9cfdca9f-99e5-46d1-ab2f-34899804f59f",
    "Name": "bat"
  }
]

$ make rpc
[
  {
    "ID": "c559b192-d428-451f-8ef7-4c36e3576068",
    "Name": "satyr"
  },
  {
    "ID": "9cfdca9f-99e5-46d1-ab2f-34899804f59f",
    "Name": "bat"
  },
  {
    "ID": "50a3c89d-bf91-416b-b200-53130a91778b",
    "Name": "hound"
  }
]
```