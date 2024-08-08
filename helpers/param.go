package helpers

import "strconv"

func ReadIDFromRequest(id_string string) (int, error) {
	id, err := strconv.Atoi(id_string)
	if err != nil {
		return -1, err
	}

	return id, nil
}
