package customtypes

type CustomContactMethodType string

const (
	CUSTOM_CONTACT_METHOD_TYPE_EMAIL       CustomContactMethodType = "EMAIL"
	CUSTOM_CONTACT_METHOD_TYPE_PHONENUMBER CustomContactMethodType = "PHONENUMBER"
)

func (cm *CustomContactMethodType) Scan(value interface{}) error {

	if val, ok := value.(string); ok {
		*cm = CustomContactMethodType(val)
	} else {
		*cm = CustomContactMethodType(value.([]byte))
	}

	return nil
}

func (cmt CustomContactMethodType) String() string {
	return string(cmt)
}
