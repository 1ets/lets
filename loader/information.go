package loader

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func Launching() {
	var printer = launcher{
		Width: 100,
	}
	printer.init()

	printer.hr()
	printer.printTitle("Let's Go Framework", "Framework is maintained by github.com/dhutapratama")

	if InfoIdentService != nil {
		printer.printStruct(*InfoIdentService)
	}

	if InfoIdentSource != nil {
		printer.printStruct(*InfoIdentSource)
	}

	// printer.printData("Microservice Name", "Dhuta Pratama")
	// printer.printData("Description", "Dhuta Pratama")
	// printer.printData("Service code/dns/module name", "Dhuta Pratama")
	// printer.printData("Microservice Layer", "Dhuta Pratama")
	// printer.printData("Owner/Author", "Dhuta Pratama (dhutapratama@gmail.com)")
	// printer.hr()
	// printer.printData("Repository", "Dhuta Pratama")
	// printer.printData("Documentation", "Dhuta Pratama")
	// printer.hr()
	// printer.printData("Environment", "Dhuta Pratama")
	// printer.printData("Debug", "Dhuta Pratama")
	// printer.hr()

	// printer.printHeading("HTTP Server")
	// // printer.printHeading("GIN - github.com/gin/gin")
	// printer.hr2()
	// printer.printData("Port", "80")
	// printer.printData("Paths", "")
	// printer.hr()
}

type launcher struct {
	Width int

	writer     io.Writer
	separator  string
	separator2 string
}

func (l *launcher) init() {
	l.separator = strings.Repeat("/", l.Width)
	l.separator += "\n"
	l.separator2 = "// " + strings.Repeat("-", l.Width-6) + " //"
	l.separator2 += "\n"

	l.writer = os.Stdout
}

func (l *launcher) printTitle(title string, maintener string) {
	var format = "// %-40s %53s //\n"

	fmt.Fprintf(l.writer, format, title, maintener)
	l.hr()
}

func (l *launcher) printHeading(title string) {
	lenTitle := len(title)
	spaceL := (l.Width - 6 - lenTitle) / 2
	spaceR := spaceL

	if (lenTitle % 2) == 1 {
		spaceR++
	}

	var format = "// %-" + fmt.Sprintf("%v", spaceL) + "s%s%-" + fmt.Sprintf("%v", spaceR) + "s //\n"

	fmt.Fprintf(l.writer, format, "", title, "")
}

func (l *launcher) printData(field string, value any) {
	var format = "// %-30s : %-61s //\n"

	fmt.Fprintf(l.writer, format, field, value)
}

func (l *launcher) printStruct(info any) {
	rt := reflect.TypeOf(info)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}

	v := reflect.ValueOf(info)
	for i := 0; i < v.NumField(); i++ {
		f := rt.Field(i)

		name := f.Name
		if f.Tag.Get("desc") != "" {
			name = strings.Split(f.Tag.Get("desc"), ",")[0]
		}

		if v.Field(i).Interface() != "" {
			l.printData(name, v.Field(i).Interface())
		}
	}
	l.hr()
}

func (l *launcher) hr() {
	fmt.Fprintf(l.writer, "%s", l.separator)
}

func (l *launcher) hr2() {
	fmt.Fprintf(l.writer, "%s", l.separator2)
}
