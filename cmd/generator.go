package cmd

import (
	"bytes"
	"os"

	"github.com/olekukonko/tablewriter"
)

type tool struct {
	apis   []*api
	output *os.File
}

type api struct {
	Method  string
	Path    string
	Doc     string
	Request *message
	Obj     map[string]*message
	Reply   *message
	Input   string
	Output  string
}

type message struct {
	Name   string
	Fields []field
	Doc    string
}

type field struct {
	Name string
	Type string
	Note string
	Doc  string
}

func newGenerator() *tool {
	t := &tool{
		apis: []*api{},
	}

	return t
}

// P forwards to g.gen.P, which prints output.
func (t *tool) P(args ...string) {
	for _, v := range args {
		t.output.WriteString(v)
	}
	t.output.WriteString("\n")
}

func (t *tool) generateDoc(file string) {
	// open output file
	var err error
	t.output, err = os.Create(file)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := t.output.Close(); err != nil {
			panic(err)
		}
	}()

	var objKeys []string
	for _, api := range t.apis {
		t.P("## ", api.Path)
		t.P()
		t.P(api.Doc)
		t.P()
		t.P("### Method")
		t.P()
		t.P(api.Method)
		t.P()
		t.P("### Request")

		rows := make([][]string, 0, len(api.Request.Fields))
		for _, message := range api.Request.Fields {
			rows = append(rows, []string{message.Name, message.Type, message.Note, message.Doc})
		}
		t.P()

		buf := new(bytes.Buffer)
		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"参数名", "类型", "说明", "是否必须"})
		table.SetBorders(tablewriter.Border{Left: true})
		table.SetCenterSeparator("|")
		table.SetColWidth(72)
		table.AppendBulk(rows)
		table.Render()
		t.P(buf.String())

		t.P()

		for k, message := range api.Obj {
			isPrint := false
			for _, v := range objKeys { // 避免重复输出相同对象
				if k == v {
					isPrint = true
				}
			}
			if !isPrint {
				t.P("### ", k)
				rows = make([][]string, 0, len(message.Fields))
				for _, fs := range message.Fields {
					rows = append(rows, []string{fs.Name, fs.Type, fs.Note, fs.Doc})
				}

				buf = new(bytes.Buffer)
				table = tablewriter.NewWriter(buf)
				table.SetHeader([]string{"参数名", "类型", "说明", "是否必须"})
				table.SetBorders(tablewriter.Border{Left: true})
				table.SetCenterSeparator("|")
				table.SetColWidth(72)
				table.AppendBulk(rows)
				table.Render()
				t.P(buf.String())

				t.P()
				objKeys = append(objKeys, k)
			}
		}

		t.P()
		t.P("### Reply")
		rows = make([][]string, 0, len(api.Reply.Fields))
		for _, message := range api.Reply.Fields {
			rows = append(rows, []string{message.Name, message.Type, message.Note, ""})
		}
		t.P()

		buf = new(bytes.Buffer)
		table = tablewriter.NewWriter(buf)
		table.SetHeader([]string{"参数名", "类型", "说明", "是否必须"})
		table.SetBorders(tablewriter.Border{Left: true})
		table.SetCenterSeparator("|")
		table.SetColWidth(72)
		table.AppendBulk(rows)
		table.Render()
		t.P(buf.String())
	}
}
