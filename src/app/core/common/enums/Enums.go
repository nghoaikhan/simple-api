package enums

type httpMethod struct {
	GET    string
	POST   string
	PUT    string
	OPTION string
	PATCH  string
	DELETE string
}

func GetHttpMethod() httpMethod {
	return httpMethod{
		GET:    "GET",
		POST:   "POST",
		PUT:    "PUT",
		OPTION: "OPTION",
		PATCH:  "PATCH",
		DELETE: "DELETE",
	}
}
