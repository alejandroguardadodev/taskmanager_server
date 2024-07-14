package customtypes

type CustomLogType string

const (
	LOG_TYPE_ACCOUNT CustomLogType = "ACCOUNT"
	LOG_TYPE_CREATE  CustomLogType = "CREATE"
	LOG_TYPE_UPDATE  CustomLogType = "UPDATE"
	LOG_TYPE_DELETE  CustomLogType = "DELETE"
)

func (cm *CustomLogType) Scan(value interface{}) error {

	if val, ok := value.(string); ok {
		*cm = CustomLogType(val)
	} else {
		*cm = CustomLogType(value.([]byte))
	}

	return nil
}

func (cmt CustomLogType) String() string {
	return string(cmt)
}
