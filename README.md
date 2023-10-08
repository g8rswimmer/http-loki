# HTTP-Loki
HTTP-Loki is a mocking service.  One can define request-response pairs for endpoints.  This is useful in testing against third party systems.

## Varables
This mocking service allows for varables in the request and response.  This allows for validtion of fields that may be generated from the service.  If the field is not assocaited with a varaible, then a strict comparison will be done. 

Variables are denoted in the body `{{ var:args }}`, where `var` is the _function_ that will be excuted with the `args`.  The agruments are seprated by a pipe `|`.

For example, to validate that a field is an number with the range of -10 to 10 the variable would look like `{{ intRange:-10|10  }}`.  As the service grows, more variable functions will be added.

### Request
Variables for request bodies will validate the body that is being received.

| Name | Description | Arguments | Prefix/Suffix Supported | Example |
|------|-------------|-------------|-----------------------|-----------|
| `uuid` | Validates if the field is UUIDv4 format | none | Y | `{{ uuid }}` |
| `intRange` | Validates that the field is an int between the range | min and max | N | `{{ intRange:-10|10 }}` |
| `ignore` | Will ignore that value when comparing | none | Y | `{{ ignore }}` |
| `regex` | Validates against a regex expersion | expersion | Y | `{{ regex:p([a-z]+)ch }}` |

### Response
Varables for the response will replace the body that will be responded.

| Name | Description | Arguments | Prefix/Suffix Supported | Example |
|------|-------------|-------------|-----------------------|-----------|
| `uuid`| Will generate a UUIDv4 ID | none | Y | `{{ uuid }}` |
| `path` | The request value of the field will be used | JSON path | Y | `{{ path:json.path }}` |
| `currTime` | Will add the current time with a layout.  Supports RFC3339 or Go [formatted layout](https://pkg.go.dev/time).| layout `RFC3339 or 02 Jan 06` | Y | `{{ currTime:RFC3339 }}` |

## Examples
Examples and a postman collection are included under the `_examples` directory.

To run and example, simply use the make command
```
make [example-target]
```

| Make Target | Loki Server Port | File Directory | Collection Folder |
|-------------|-------------------|--------------|---------------------|
| `example-basic` | 8000 | `_examples/bacic/mock-files` | Basic |
