package types

import "time"

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02"`, string(data))
	*d = Date(parsedTime)
	return err
}
