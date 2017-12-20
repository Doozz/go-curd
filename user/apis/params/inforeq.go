package params
import (
	"net/http"
	"github.com/mholt/binding"
)

type RegReq struct {
    Name string
    Pass   string
    Comfirm string
    Mobile   string
    Email string
}

type IndexReq struct {
	Id int
}

type LoginReq struct {
	Name string
	Pass string
}

type UpdataReq struct {
    Name string
    Mobile   string
    Email string
}

func (data *RegReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Name: binding.Field{
			Form:         "name",
			Required:     true,
			ErrorMessage: "1",
		},
		&data.Pass: binding.Field{
			Form:         "pass",
			Required:     true,
			ErrorMessage: "2",
		},
		&data.Comfirm: binding.Field{
			Form:         "confirm",
			Required:     true,
			ErrorMessage: "3",
		},
		&data.Mobile: binding.Field{
			Form:         "mobile",
			Required:     true,
			ErrorMessage: "4",
		},
		&data.Email: binding.Field{
			Form:         "email",
			Required:     true,
			ErrorMessage: "5",
		},
	}
}

func (data *IndexReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Id: binding.Field{
			Form:         "id",
			Required:     true,
			ErrorMessage: "invaldate",
		},
	}
}

func (data *LoginReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Name: binding.Field{
			Form:         "username",
			Required:     true,
			ErrorMessage: "invaldate",
		},
		&data.Pass: binding.Field{
			Form:         "pass",
			Required:     true,
			ErrorMessage: "invaldate",
		},
	}
}

func (data *UpdataReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Name: binding.Field{
			Form:         "name",
			Required:     true,
			ErrorMessage: "invaldate",
		},
		&data.Email: binding.Field{
			Form:         "email",
			Required:     true,
			ErrorMessage: "invaldate",
		},
		&data.Mobile: binding.Field{
			Form:         "mobile",
			Required:     true,
			ErrorMessage: "invaldate",
		},
	}
}



