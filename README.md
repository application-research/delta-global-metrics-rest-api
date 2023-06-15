# Delta Metrics Rest
Rest API for Delta Global Metrics

## Set up the .env file
```
DB_NAME=deltadb_metrics
DB_HOST=<delta-metrics connection uri>
DB_USER=deltadb_metrics_user
DB_PASS=<password>
DB_PORT=5432
```

## Build the binary
```
make dmr
```

## Run the binary
```
./dmr
```

## Check the live totals info here
```
https://global.delta.store/open/stats/totals/info
```

## Global Stats available
- total deals attempted
- total e2e deals attempted
- total import deals attempted
- total deals succeeded
- total e2e deals succeeded
- total import deals succeeded
- total deals failed
- total e2e deals failed
- total import deals failed
- total deals active
- total e2e deals active
- total import deals active
- total number of sps
- total number of delta nodes

## Storage stats available
- total storage consumed by deals attempted
- total storage consumed by e2e deals attempted
- total storage consumed by import deals attempted
- total storage consumed by deals succeeded
- total storage consumed by e2e deals succeeded
- total storage consumed by import deals succeeded
- total storage consumed by deals failed
- total storage consumed by e2e deals failed
- total storage consumed by import deals failed
- total storage consumed by deals active
- total storage consumed by e2e deals active
- total storage consumed by import deals active
- total storage consumed by all deals
- total storage consumed by all e2e deals

## SP
- list of SP
- list of SP location
- list of SP deals attempted
- list of SP deals succeeded
- list of SP deals failed
- list of SP deals active
- list of SP deals attempted by e2e
- list of SP deals succeeded by e2e
- list of SP deals failed by e2e
- list of SP deals active by e2e
- list of SP deals attempted by import
- list of SP deals succeeded by import
- list of SP deals failed by import


# Author
Protocol Labs Outercore Engineering.
