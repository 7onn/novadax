# novadax

## TL;DR;
```bash
dep ensure && \
    go build -o bid src/* && \
    ./bid 0.001
```

## Build
```bash
dep ensure # fetching dependencies
go build -o bid src/* # build the src content into an executable called 'bid'
./bid 0.001 # runs the executable with the desired BTC amount to buy at average price
```

## Dev
```bash
air . # golang live reload
```


## Description

this project is meant to cancel old orders and place new ones to updated prices;

ideally it runs every 2 hours
