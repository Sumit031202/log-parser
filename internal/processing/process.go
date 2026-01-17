package processing

import (
	"strings"
)

const (
	ErrorKey="ERROR"
	LevelCrit="CRITICAL"
	LevelInfo="INFO"
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

func CountCritical(fileContent string) int{
	count:=0

	// split the lines
	lines:=strings.Split(fileContent, "\n")

	for _, line:= range lines {
		if line=="" {
			continue
		}
		level:=GetLogLevel(line)

		if(level==LevelCrit){
			count++
		}
	}
	return count
}
