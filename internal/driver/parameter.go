package driver

// BaseParam Both `GET` and `SET` command are required information
type BaseParam struct {
	// todo
}

type GetParam struct {
	BaseParam
}

type SetParam struct {
	BaseParam
	Value interface{}
}

func newBaseParam(attr map[string]interface{}) (base BaseParam, err error) {
	// todo from attr to BaseParam
	return BaseParam{}, nil
}

func NewGetParam(attr map[string]interface{}) (*GetParam, error) {
	base, err := newBaseParam(attr)

	return &GetParam{
		BaseParam: base,
	}, err
}

func NewSetParam(attr map[string]interface{}, value interface{}) (*SetParam, error) {
	base, err := newBaseParam(attr)

	return &SetParam{
		BaseParam: base,
		Value:     value,
	}, err
}
