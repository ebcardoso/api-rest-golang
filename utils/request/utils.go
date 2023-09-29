package request

import (
	"errors"
	"strconv"
)

func ParsePageNumber(paramPage string) (int, error) {
	if paramPage == "" {
		return 1, nil
	}

	page, err := strconv.Atoi(paramPage)
	if err != nil {
		return 0, err
	}
	if page < 1 {
		return 0, errors.New("invalid page")
	}
	return page, nil
}
