package lets

import (
	"github.com/kataras/golog"
)

var (
	Log  = golog.New()
	LogD = Log.Debugf
	LogI = Log.Infof
	LogW = Log.Warnf
	LogE = Log.Errorf
	LogF = Log.Fatalf
	Logf = Log.Logf
)

// Log JSON
func LogJ(arg interface{}) {
	data := ToJson(arg)
	LogD("%s", data)
}

// Log as JSON with Indent
func LogJI(arg interface{}) {
	data := ToJsonIndent(arg)
	LogD("%s", data)
}

// Log Error
func LogErr(err error) {
	Log.Error(err)
}
