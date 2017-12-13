gody
=======

DynamoDB Command Line Interface

## Installation

```
$ go get github.com/watarukura/gody
```

## Usage

### Get

```
$ gody get --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --format <ssv|csv|tsv|json> \
     --header
```

### Query (PartitionKey, SortKey)

```
$ gody query --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --format <ssv|csv|tsv|json> \
     --header \
     --limit 10
```

### Query (Global Secondary Index)

```
$ gody query --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --index <IndexName> \
     --format <ssv|csv|tsv|json> \
     --header \
     --limit 10

```

### Update

```
$ gody update --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --set {}
```

### Put

```
$ gody put --table <TableName> \
     --format <ssv|csv|tsv|json> \
     --file <FilePath>
```

```
$ cat put.csv
jan,name,price
4515438304003,茶こし共柄,500
4571277751224,スパイダージェル　500ml,_
$ cat put.csv |
> gody put --table item --format csv
$ gody get --table item --format csv --pkey 4515438304003 --header --field jan,name,price
jan,name,price
4515438304003,茶こし共柄,500
$ gody get --table item --format csv --pkey 4571277751224 --header --field jan,name,price
jan,name,price
4571277751224,スパイダージェル　500ml,_
```

### Delete

```
$ gody delete --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey>
```

### Scan

```
$ gody scan --table <TableName> \
     --format <ssv|csv|tsv|json> \
     --header \
     --limit 10
```
