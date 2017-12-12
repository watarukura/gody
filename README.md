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
