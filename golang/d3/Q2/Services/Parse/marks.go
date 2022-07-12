package Parse

func ParseMarksList(requestJSON map[string]string) (string, error) {
	marksList, found := requestJSON["marksList"]
	if !found {
		return "", nil
	}

	return marksList, nil
}
