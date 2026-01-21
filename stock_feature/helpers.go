package stock_feature

import "fmt"

// getStringFeature 从map中安全获取字符串
func getStringFeature(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if v == nil {
			return ""
		}
		switch val := v.(type) {
		case string:
			return val
		case float64:
			return fmt.Sprintf("%v", val)
		case int64:
			return fmt.Sprintf("%d", val)
		default:
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}
