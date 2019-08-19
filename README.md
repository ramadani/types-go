# types-go

Go custom data type

## Installation

```bash
$ go get github.com/garavan/types-go
```

## Usage

Import types-go to your go project

```go
import types "github.com/garavan/types-go"
```

### Date and Null Date

```go
type User struct {
	Name        string   `json:"name"`
	Dob         Date     `json:"dob"`
	DeletedDate NullDate `json:"deleted_date"`
}
```

**Marshalling json**
```go
dobTime, _ := time.Parse("2006-01-02", "1990-06-03")
eko := User{
	Name:        "Eko Syamsudin",
	Dob:         Date{Time: dobTime},
	DeletedDate: NullDate{Time: time.Now(), Valid: true},
}

bytes, _ := json.Marshal(&eko)

// json:
// {"name":"Eko Syamsudin","dob":"1990-06-03" "deleted_date":"2019-08-19"}
```

if deleted date is null
```go
dobTime, _ := time.Parse("2006-01-02", "1990-06-03")
eko := User{
	Name:        "Eko Syamsudin",
	Dob:         Date{Time: dobTime},
	DeletedDate: NullDate{},
}

bytes, _ := json.Marshal(&eko)

// json:
// {"name":"Eko Syamsudin","dob":"1990-06-03" "deleted_date":"null"}
```

**Unmarshalling json**
```go
dani := User{}
err := json.Unmarshal([]byte("{\"name\":\"Ramadani\",\"dob\":\"1992-07-25\"}"), &dani)

// dani.DeletedDate is equal NullDate{}
```

## Tests

```bash
go test
```