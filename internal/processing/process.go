package processing

import (
	"strings"
)

const (
	ErrorKey="ERROR"
	LevelCrit="CRITICAL"
	LevelInfo="INFO"
)

type LogEntry struct{
	Time string
	Level string
	Message string
}

func ParseLogLine(line string) LogEntry{
	// divide the line into parts
	// 2026-01-15 10:00:00 [INFO] System starting up
	parts:=strings.Split(line, " ")

	if len(parts)<4 {
		return LogEntry{}
	}
	timeStamp:= parts[0] + " " + parts[1]
	level:= strings.Trim(parts[2],"[]")
	message:= strings.Join(parts[3:]," ")

	return LogEntry{
		Time: timeStamp,
		Level: level,
		Message: message,
	}
}

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
