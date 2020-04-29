package log

type Fields map[string]interface{}

type F = Fields // for short

// merge src into dst
func MergeFields(dst, src Fields) Fields {
	if len(src) > 0 {
		if dst == nil {
			dst = Fields{}
		}
		for key, val := range src {
			dst[key] = val
		}
	}
	return dst
}

func ExtractFields(in []interface{}) (fields Fields, out []interface{}) {
	for _, v := range in {
		if _fields, ok := v.(Fields); ok {
			fields = MergeFields(fields, _fields)
		} else {
			out = append(out, v)
		}
	}
	return
}
