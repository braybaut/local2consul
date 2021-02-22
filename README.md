### bulk-upload-to-consul

This tool allows developers to upload large amounts of key/value to consul

## Build 

`make build` 

## Usage 

Using a **.txt** file 

```
./bulk-upload-to-consul put --consulUrl http://consul.cloud:8500 --domain configurations/microservice1/qa/settings --file values-qa.txt
```

Using  a **.json**  file you can define the environments you need 

```
./bulk-upload-to-consul put  --domain configurations/microservice1/qa/settings --file values-qa.json
```

The json file must have the format like this:

```
[
    {
        "environment:":  "qa1",
        "ConsulAddress": "consul-qa1-cloud:8500",
        "keyvalues": [    
            {
                "key": "FOO",
                "value": "bar"
            }
    ]
},
{
        "environment:":  "qa2",
        "ConsulAddress": "consul-qa2.cloud:8500",
        "keyvalues": [    
            {
                "key": "FOO",
                "value": "bar"
            }          
        ]
    }
]

```

## Assumptions 

the .txt file must be contain the keyvalues in format **key=value** example:

```
FOO=bar
MAX_CONNECTIONS=2
```

## Requirements

go >= 1.14


## Credits 

Made with ❤  By [Braybaut](https://twitter.com/braybaut)
