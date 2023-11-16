package json



type J map[string]any

func Ok(msg string) map[string]any {
	return map[string]any{"success": true, "message": msg}
}

func Fail(msg string) map[string]any {
	return map[string]any{"success": false, "message": msg}
}

func FailUnknown() map[string]any {
	return Fail("something went wrong")
}
