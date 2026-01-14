package processing

import (
	"strings"
)


func ProcessLog(line string) string {
	return "Processed: " + line
}

func GetLogLevel(line string) string{
	if(strings.Contains(line,"Error")){
		return "Critical"
	}else{
		return "Info"
	}
}
