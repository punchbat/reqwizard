package utils

import "time"

func GetISOTimeFromInt64(int64Time int64) string {
	if int64Time == 0 {
		return ""
	}

	// Преобразование int64 в значение типа time.Time
	timeValue := time.UnixMilli(int64Time)

	// Преобразование time.Time в строку формата ISO 8601
	return timeValue.Format(time.RFC3339)
}

func GetInt64FromISOTime(isoTime string) (int64, error) {
	if len(isoTime) == 0 {
		return 0, nil
	}

	// Преобразование строки формата ISO 8601 в значение типа time.Time
	timeValue, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		return 0, err
	}

	// Преобразование time.Time в int64 (временную метку Unix)
	return timeValue.UnixMilli(), nil
}

func GetTimeFromInt64(int64Time int64) time.Time {
	// Если int64Time равно нулю, вернуть нулевое значение time.Time
	if int64Time == 0 {
		return time.Time{}
	}

	// Преобразование int64 в значение типа time.Time (время в миллисекундах)
	return time.Unix(0, int64Time*int64(time.Millisecond))
}

func GetInt64FromTime(t time.Time) int64 {
	// Если t является нулевым значением time.Time, вернуть ноль
	if t.IsZero() {
		return 0
	}

	// Преобразование time.Time в int64 (временную метку Unix в миллисекундах)
	return t.UnixNano() / int64(time.Millisecond)
}

func GetStringFromTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(time.RFC3339)
}

func GetTimeFromString(strTime string) (time.Time, error) {
	if len(strTime) == 0 {
		return time.Time{}, nil
	}

	timeValue, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		return time.Time{}, err
	}

	return timeValue, nil
}