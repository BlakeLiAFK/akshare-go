package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ToFloat64 将任意值转换为float64
func ToFloat64(v interface{}) (float64, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case string:
		return ParseFloat(val)
	case json.Number:
		return val.Float64()
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("无法将 %T 转换为 float64", v)
	}
}

// ToInt64 将任意值转换为int64
func ToInt64(v interface{}) (int64, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case int64:
		return val, nil
	case int:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case string:
		return ParseInt(val)
	case json.Number:
		return val.Int64()
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("无法将 %T 转换为 int64", v)
	}
}

// ToString 将任意值转换为string
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case bool:
		return strconv.FormatBool(val)
	case json.Number:
		return val.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// ToBool 将任意值转换为bool
func ToBool(v interface{}) bool {
	if v == nil {
		return false
	}

	switch val := v.(type) {
	case bool:
		return val
	case int, int64, int32:
		return reflect.ValueOf(val).Int() != 0
	case float64:
		return val != 0
	case string:
		b, _ := strconv.ParseBool(val)
		return b
	default:
		return false
	}
}

// MustFloat64 将任意值转换为float64，失败返回0
func MustFloat64(v interface{}) float64 {
	f, _ := ToFloat64(v)
	return f
}

// MustInt64 将任意值转换为int64，失败返回0
func MustInt64(v interface{}) int64 {
	i, _ := ToInt64(v)
	return i
}

// MustInt 将任意值转换为int，失败返回0
func MustInt(v interface{}) int {
	i, _ := ToInt64(v)
	return int(i)
}

// MapToStruct 将map转换为struct
func MapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// StructToMap 将struct转换为map
func StructToMap(v interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}
