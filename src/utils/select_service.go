package utils

import "InfobaeAPI/src/constants"

func SelectService(topic string) (string, error) {
	url, ok := constants.URL[topic]
	if !ok {
		return constants.URL["init"], nil
	}
	return url, nil
}
