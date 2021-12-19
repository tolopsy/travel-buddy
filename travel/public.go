package travel

type PublicViewership interface {
	Public() interface{}
}

func Public(o interface{}) interface{} {
	if p, canViewInPublic := o.(PublicViewership); canViewInPublic {
		return p.Public()
	}

	return o
}