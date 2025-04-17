package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {}

func (l Logger) logTime() (hour int, minute int, second int, millisecond int) {
	now := time.Now()
	millis := now.Nanosecond() / int(time.Millisecond)
	return now.Hour(), now.Minute(), now.Second(), millis
}

func (l Logger) logTimeString() string {
	hour, minute, second, millis := l.logTime()
	hourStr := fmt.Sprintf("%02d", hour)
	minuteStr := fmt.Sprintf("%02d", minute)
	secondStr := fmt.Sprintf("%02d", second)
	millisStr := fmt.Sprintf("%03d", millis)

	lTimeString := "[" + color.GreenString(hourStr) + ":" + color.GreenString(minuteStr) + ":" + color.GreenString(secondStr) + "." + color.GreenString(millisStr) + "]"

	return lTimeString
}

func (l Logger)	DefaultSeperator() string {
	return  color.RGB(50, 50, 50).Sprint("--")
}

func (l Logger) Print(m string) {
	prefix := l.logTimeString()
	fmt.Printf("%s %s %s\n", prefix, l.DefaultSeperator(), m)
}

func (l Logger) Printf(f string, a ...any) {
	prefix := l.logTimeString()
	m := fmt.Sprintf("%s %s %s\n", prefix, l.DefaultSeperator(), f)
	fmt.Printf(m, a...)
}

func (l Logger) Info(m string) {
	prefix := "[" + color.RGB(0, 106, 79).Sprint("INFO") + "]  " + l.logTimeString()
	fmt.Printf("%s %s %s\n", prefix, l.DefaultSeperator(), m)
}

func (l Logger) Infof(f string, a ...any) {
	prefix := "[" + color.RGB(0, 106, 79).Sprint("INFO") + "]  " + l.logTimeString()
	m := fmt.Sprintf("%s %s %s\n", prefix, l.DefaultSeperator(), f)
	fmt.Printf(m, a...)
}

func (l Logger) Warn(m string) {
	prefix := "[" + color.RGB(166, 107, 0).Sprint("WARN") + "]  " + l.logTimeString()
	fmt.Printf("%s %s %s\n", prefix, l.DefaultSeperator(), m)
}

func (l Logger) Warnf(f string, a ...any) {
	prefix := "[" + color.RGB(166, 107, 0).Sprint("WARN") + "]  " + l.logTimeString()
	m := fmt.Sprintf("%s %s %s\n", prefix, l.DefaultSeperator(), f)
	fmt.Printf(m, a...)
}

func (l Logger) Fatal(m string, err error) {
	prefix := "[" + color.RGB(137, 0, 21).Sprint("FATAL") + "] " + l.logTimeString()
	fmt.Printf("%s %s %s\n", prefix, l.DefaultSeperator(), m)
	fmt.Printf("%s %s %s\n", prefix, l.DefaultSeperator(), err.Error())
	os.Exit(1)
}

func (l Logger) Fatalf(f string, err error, a ...any) {
	prefix := "[" + color.RGB(137, 0, 21).Sprint("FATAL") + "] " + l.logTimeString()
	m := fmt.Sprintf("%s %s %s\n", prefix, l.DefaultSeperator(), f)
	fmt.Printf(m, a...)
	fmt.Printf("%s %s %s\n", prefix, l.DefaultSeperator(), err.Error())
}
