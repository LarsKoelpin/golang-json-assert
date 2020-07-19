# Golang Json Assert
![golang-json-assert](https://github.com/LarsKoelpin/golang-json-assert/workflows/golang-json-assert/badge.svg)

A Go library for testing JSON strings.

## Mission
I could not find any library implementing the "Conform behavior", which checks,
 if JSON objects are downwards-compatible, therefore testing if there is a semantic versioning major bump change.  
 
 
## Usage

For full specification of the library see `pkg/jsonassert_test.go`.

In a Nutshell, given to strings, the library can check for *strict* equality:
``` go
actual := `
{
  "name": "Lars",
  "age": 12
}
`

Expect(actual).ToEqual(actual) // returns true
```


But, It can also check, if a JSON conforms with another. This means, that in theory
actual and expected value are schematically compatible. This can be used to ensure that 
diffrent JSON objects provide some compability.

``` go
actual := `
{
  "name": "Lars",
  "age": 12
}
`

expect := `
{
  "age": 12
}
`

Expect(actual).ToConform(expect) // returns true
```

## Contribution
If you feel like there are missing some assertions / specifications, feel free to open an issue / Submit a PR.

## Change log

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Testing

``` bash
$ go test
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

