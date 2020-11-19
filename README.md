Date
[![Go Report Card](https://goreportcard.com/badge/github.com/ganigeorgiev/date)](https://goreportcard.com/report/github.com/ganigeorgiev/date)
[![GoDoc](https://godoc.org/github.com/ganigeorgiev/date?status.svg)](https://pkg.go.dev/github.com/ganigeorgiev/date)
================================================================================

Date is a small Go package that defines a `date.Date` struct, representing a date as in **ISO 8601** (eg. `2006-01-02`).

The package was created primarily to scan SQL `DATE` type and other date-only values.

Under the hood, each date is stored as `time.Time` instant with _zero time part_ in UTC (eg. `2006-01-02 00:00:00:00 UTC`).
As a result, many of the `date.Date` methods are implemented using the corresponding methods of `time.Time`.

> There is an active proposal for something similar to be implemented in the standard library - [golang/go#19700](https://github.com/golang/go/issues/19700)


## Installation

```
go get github.com/ganigeorgiev/date
```

Example usage:

```go
import github.com/ganigeorgiev/date

type User {
    JoinDate date.Date
}
```

See the [package documentation](@todo) for more details and examples.
