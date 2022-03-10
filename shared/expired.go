package shared

import (
	"errors"
	"fmt"
	"time"
)

func ExpiredVerify(timeParam string) error {
	layoutFormat := "20060102150405"
	t := time.Now()
	now := t.Format(layoutFormat)

	date, err := time.Parse(layoutFormat, timeParam)
	if err != nil {
		return err
	}
	dateNow, _ := time.Parse(layoutFormat, now)

	createdAt := dateNow.Add(1 * time.Hour)
	expiresAt := dateNow.Add(2 * time.Hour)

	expired := expiresAt.Sub(createdAt)
	fmt.Println(expired)

	created := dateNow.Sub(date)
	fmt.Println(created)

	if created > expired {
		err = errors.New("your activation is expired")
		return err
	}
	return nil
}
