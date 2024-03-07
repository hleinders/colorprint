package AnsiTerm

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	at "github.com/hleinders/AnsiTerm"
)

var Out *bufio.Writer = bufio.NewWriter(os.Stdout)

//
// Little Printer: ------------------------------------------------
//

func hline(lchar string, n int) string {
	if n == 0 {
		n, _, _ = at.GetSize()
	}

	return strings.Repeat(lchar, n)
}

// Printer is a little shortcut for verbose and debug output
type Printer struct {
	flagVerbose, flagDebug, flagSilent bool
}

func NewPrinter() *Printer {
	var p Printer

	// Detect locale for printing:
	if runtime.GOOS == "windows" {
		at.AsciiChars()
	}

	return &p
}

// Management functions for Printer
func (l *Printer) SetDebug(b bool) {
	l.flagDebug = b
}

func (l *Printer) SetVerbose(b bool) {
	l.flagVerbose = b
}

func (l *Printer) SetSilent(b bool) {
	l.flagSilent = b
}

// Helper functions for Printer
func (l Printer) Frame(str string) string {
	sl := len(str)
	rh := at.FrameOpenL + strings.Repeat(at.FrameHLine, sl+2) + at.FrameOpenR
	rt := at.FrameCloseL + strings.Repeat(at.FrameHLine, sl+2) + at.FrameCloseR
	return fmt.Sprintf("%s\n%s %s %s\n%s\n", rh, at.FrameVLine, str, at.FrameVLine, rt)
}

func (l Printer) OFrame(str string) string {
	sl := len(str)
	rh := at.FrameOOpenL + strings.Repeat(at.FrameOHLine, sl+2) + at.FrameOOpenR
	rt := at.FrameOCloseL + strings.Repeat(at.FrameOHLine, sl+2) + at.FrameOCloseR
	return fmt.Sprintf("\n%s\n%s %s %s\n%s\n", rh, at.FrameOVLine, str, at.FrameOVLine, rt)
}

func (l Printer) Underlines(row []string) []string {
	anonRow := make([]string, len(row))
	for i, v := range row {
		anonRow[i] = strings.Repeat(at.FrameHLine, len(v))
	}
	return anonRow
}

func (l Printer) WriteOut(fmtString string, args ...interface{}) {
	if !l.flagSilent {
		fmt.Printf(fmtString, args...)
	}
}

func (l Printer) WriteAny(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString, args...)
}

// Print functions for logger
func (l Printer) Banner(fmtString string, args ...interface{}) {
	rStr := fmt.Sprintf(fmtString, args...)
	str := l.Frame(rStr)
	l.WriteOut(at.Bold(at.Green(str)))
}

// Print functions for logger
func (l Printer) OBanner(fmtString string, args ...interface{}) {
	rStr := fmt.Sprintf(fmtString, args...)
	str := l.OFrame(rStr)
	l.WriteOut(at.Bold(at.Green(str)))
}

func (p Printer) ModuleHeading(subPage bool, modName, fmtString string, args ...interface{}) {
	var eStr string
	if subPage {
		fmt.Println("\n" + hline(at.ThinHLine, 80))
		eStr = fmt.Sprintf("\nModule %-10s   %s\n", modName+":", fmt.Sprintf(fmtString, args...))
	} else {
		eStr = fmt.Sprintf("\nUsage:   %s\n", fmt.Sprintf(fmtString, args...))
	}
	p.WriteAny(at.Yellow(eStr))
}

func (l Printer) Verbose(fmtString string, args ...interface{}) {
	if l.flagVerbose {
		l.WriteOut(fmtString, args...)
	}
}

func (l Printer) Verboseln(fmtString string, args ...interface{}) {
	l.Verbose(fmtString+"\n", args...)
}

func (l Printer) VerboseInfo(fmtString string, args ...interface{}) {
	if l.flagVerbose {
		l.WriteOut(at.Green(fmtString), args...)
	}
}

func (l Printer) VerboseInfoln(fmtString string, args ...interface{}) {
	l.VerboseInfo(fmtString+"\n", args...)
}

func (l Printer) VerboseBold(fmtString string, args ...interface{}) {
	if l.flagVerbose {
		l.WriteOut(at.Bold(fmtString), args...)
	}
}

func (l Printer) VerboseBoldln(fmtString string, args ...interface{}) {
	l.VerboseBold(fmtString+"\n", args...)
}

func (l Printer) Debug(fmtString string, args ...interface{}) {
	if l.flagDebug {
		fs := "*** DEB: " + fmtString
		l.WriteOut(at.Red(fs), args...)
	}
}

func (l Printer) Debugln(fmtString string, args ...interface{}) {
	l.Debug(fmtString+"\n", args...)
}

func (l Printer) Warning(fmtString string, args ...interface{}) {
	fs := "*** WARN: " + fmtString
	l.WriteAny(at.Yellow(fs), args...)
}

func (l Printer) Warningln(fmtString string, args ...interface{}) {
	l.Warning(fmtString+"\n", args...)
}

func (l Printer) Error(fmtString string, args ...interface{}) {
	fs := "*** ERR: " + fmt.Sprintf(fmtString, args...)
	l.WriteAny(at.Red(fs))
}

func (l Printer) Errorln(fmtString string, args ...interface{}) {
	l.Error(fmtString+"\n", args...)
}
