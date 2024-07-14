package models

type Dictionary map[string]interface{}

func DictionarySetup(values map[string]string) Dictionary {

	dic := Dictionary{}

	for key, value := range values {
		dic[key] = value
	}

	return dic
}

func (dic Dictionary) Add(values Dictionary) Dictionary {
	for key, value := range values {
		dic[key] = value
	}

	return dic
}
