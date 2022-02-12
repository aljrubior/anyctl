package advisors

func NewUnauthorizedAdviser(err error) *UnauthorizedAdviser {
	return &UnauthorizedAdviser{
		err: err,
	}
}

type UnauthorizedAdviser struct {
	err error
}

func (this UnauthorizedAdviser) Advise() {
	println("ERROR: Unauthorized. Try with 'anyctl config refresh-token'")
}
