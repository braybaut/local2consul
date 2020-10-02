### LOCAL2CONSUL

This tool allows developers to put large amounts of key/value in consul

## Build 

`make build` 

## Usage 

```
./local2consul put --consulUrl http://consul.cloud:8500 --domain configurations/microservice1/qa/settings --file values-qa.txt
```


## Considerations 

the file must be contain the keyvalues in format **key=value** example:

```
FOO=bar
MAX_CONNECTIONS=2
```

## Requeriments

go >= 1.14


## Credits 

Made with ❤  By [Braybaut](https://twitter.com/braybaut)
