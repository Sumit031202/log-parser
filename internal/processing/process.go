package processing

import (
	"strings"
)

const (
	ErrorKey="Error"
	LevelCrit="Critical"
	LevelInfo="Info"
)

func ProcessLog(line string) string {
	return "Processed: " + line
}

func GetLogLevel(line string) string{
	if(strings.Contains(line,ErrorKey)){
		return LevelCrit
	}else{
		return LevelInfo
	}
}
