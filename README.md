# HTTP-Loki
HTTP-Loki is a mocking service.  One can define request-response pairs for endpoints.  This is useful in testing against third party systems.

## Varables
This mocking service allows for varables in the request and response.  This allows for validtion of fields that may be generated from the service.  If the field is not assocaited with a varaible, then a strict comparison will be done. 

The variables are defined in the mock file for the fields.  The following object will be defined.

| Name | Required | Description | Example |
|------|----------|-------------|---------|
| `json_path` | Y | The JSON path of the field | `address.street1` |
| `func` | Y | The variable function that will be run. | `uuid` |
| `args` | N | Any agruments that are used by the function.  This is a string array | `["-10", "10" ]` |
| `prefix` | N | A prefix string that will be removed from the request body.  Will not be used in the function. | `test prefix` |
| `suffix` | N | A suffix string that will be removed from the request body.  Will not be used in the funciton. | `test suffix` |

### Example
```json
{
    "json_path": "num_apples",
    "func": "intRange",
    "args":["0","100"]
}
```

### Request
Request body variables will validate the field assocaited with them.  In the request the field that will be validated needs to contain the string `{{ validation }}`. 

| Name | Description | Arguments | Prefix/Suffix Supported |
|------|-------------|-------------|-----------------------|
| `uuid` | Validates if the field is UUIDv4 format | none | Y |
| `intRange` | Validates that the field is an int between the range | min and max | N |
| `ignore` | Will ignore that value when comparing | none | Y |
| `regex` | Validates against a regex expersion | expersion | Y |

#### Example
```json
"request":{
    "body":{
        "id": "{{ validation }}",
        "first_name": "John",
        "last_name": "Funster",
        "preferred_name": "Jonny"
    },
    "body_validations":[
        {
            "json_path": "id",
            "func": "uuid",
            "prefix": "test|"
        }
    ]
}
```

### Response
Varables for the response will replace the body that will be responded.

| Name | Description | Arguments | Prefix/Suffix Supported |
|------|-------------|-------------|-----------------------|
| `uuid`| Will generate a UUIDv4 ID | none | Y |
| `path` | The request value of the field will be used | JSON path | Y |
| `currTime` | Will add the current time with a layout.  Supports RFC3339 or Go [formatted layout](https://pkg.go.dev/time).| layout `RFC3339 or 02 Jan 06` | Y | 

#### Example
```json
"response":{
    "status_code": 200,
    "body": {
        "id": "{{ replacememt }}",
        "first_name": "Jon",
        "last_name": "Doe",
        "language": "en"
    },
    "body_replacements":[
        {
            "json_path": "id",
            "func": "path",
            "prefix": "test|",
            "args": ["id"]
        }
    ]
}
```

## Examples
Examples and a postman collection are included under the `_examples` directory.

To run and example, simply use the make command
```
make [example-target]
```

| Make Target | Loki Server Port | File Directory | Collection Folder |
|-------------|-------------------|--------------|---------------------|
| `example-basic` | 8000 | `_examples/basic/mock-files` | Basic |
| `example-req-validation` | 8000 | `_examples/request-validation/mock-files` | Request Valiation |
| `example-resp-replacement` | 8000 | `_examples/response-replacement/mock-files` | Response Replacement |
| `example-advanced` | 8000 | `_examples/advanced/mock-files` | Advanced |
