package params
import (
	"net/http"
	"github.com/mholt/binding"
)

type InfoReq struct {
    Name string
    Age   int
    From    string
    Sex int
    Memo    string
    Id int
}

type IDReq struct {
	Id int
}

func (data *InfoReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Name: binding.Field{
			Form:         "nick",
			Required:     true,
			ErrorMessage: "1",
		},
		&data.Age: binding.Field{
			Form:         "age",
			Required:     true,
			ErrorMessage: "2",
		},
		&data.From: binding.Field{
			Form:         "from",
			Required:     true,
			ErrorMessage: "3",
		},
		&data.Sex: binding.Field{
			Form:         "sex",
			Required:     true,
			ErrorMessage: "4",
		},
		&data.Memo: binding.Field{
			Form:         "subject",
			Required:     true,
			ErrorMessage: "5",
		},
		&data.Id: binding.Field{
			Form:         "id",
		},
	}
}

func (data *IDReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Id: binding.Field{
			Form:         "id",
			Required:     true,
			ErrorMessage: "invaldate",
		},
	}
}



