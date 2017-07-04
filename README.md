CMDR
=====

A predefined shell commands with web interface.

## How to use

1. Start CMDR. CMDR will use configuration from ```cmdr.json``` file.
2. Rest API Request requires a header named ```api-key``` with a key configured in ```api_token``` key in file ```cmdr.json```.
3. Send Json request in POST body as below. CMDR will execute each command sequentially and return console output as array corresponding to the request.

    ```json
    {
        "commands": [
            "stop WMSvc",
            "start WMSvc"
        ]
    }
    ```

## Unix-based/OSX configuration example

```json
{
    "api_token": "XXXXXXX",
    "ip": "0.0.0.0",
    "port": "33487",
    "commands": {
        "ls": "sh -c 'ls %[1]s'"
    }
}
```

## Windows-based configuration example

```json
{
    "api_token": "XXXXXXX",
    "ip": "0.0.0.0",
    "port": "33487",
    "commands": {
        "stop": "net stop %[1]s",
        "start": "net start %[1]s"
    }
}
```

## Alias

Alias will be declared under ```commands``` key as key-value. Key is the first token in the request command. Parameters in the request command will be treated as nth-index array started from 1 and can be referred in alias as %[n]s parameter.

### Example

From windows-based configuration example, a command ```stop WMSvc``` is equivalent to ```net stop WMSvc```.

## License

MIT: <http://chonla.mit-license.org/>