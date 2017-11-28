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
     --skey <SortKey>
```

### Query (PartitionKey, SortKey)

```
$ gody query --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --limit 10
```

### Query (Global Secondary Index)

```
$ gody query --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --index <IndexName> \
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
     --pkey <PartitionKey> \
     --skey <SortKey> \
     --set {}
       
```

### Delete

```
$ gody delete --table <TableName> \
     --pkey <PartitionKey> \
     --skey <SortKey>
```

### scan

```
$ gody scan --table <TableName> \
     --limit 10
```
