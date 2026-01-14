package processing

// ProcessLog takes a raw log line and adds a prefix.
// It returns the processed string.
func ProcessLog(line string) string {
	return "Processed: " + line
}
