package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	url_path = "http://127.0.0.1/"
)

var (
	mdPath      string
	prefix_file string
)

func init() {
	Cmd.Flags().StringVar(&mdPath, "mark_down", ".", "mark_down path")
	Cmd.Flags().StringVar(&prefix_file, "prefix", "api.go", "prefix file")
}

// Cmd run version
var Cmd = &cobra.Command{
	Use:   "doc",
	Short: "Run doc",
	Long:  `Run doc`,
	Run: func(cmd *cobra.Command, args []string) {
		var files []string
		root := "."
		t := newGenerator()

		err := filepath.Walk(root, scanFile(&files))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if strings.HasSuffix(file, prefix_file) {
				t.apis = scanStruct(file)
				mdName := strings.Replace(file, ".go", ".md", -1)
				if mdPath != "." {
					mdName = mdPath + "/" + mdName
				}
				t.generateDoc(mdName)
			}
		}
	},
}

func scanFile(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func scanStruct(path string) (apis []*api) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// ast.Print(fset, f)
	ap := &api{}
	ap.Method = "GET/POST"
	ap.Obj = make(map[string]*message)
	for _, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
			for _, s := range gen.Specs {
				if do, ok := s.(*ast.TypeSpec); ok && do.Doc != nil {
					j := strings.Index(do.Doc.List[0].Text, "@")

					if len(do.Doc.List) > 1 {
						doc := do.Doc.List[0].Text[:j]
						doc = strings.TrimSpace(strings.Replace(doc, "//", "", -1))
						ap.Doc = doc
						ap.Path = url_path + do.Doc.List[0].Text[j+1:]

						k := strings.Index(do.Doc.List[1].Text, "@")
						if strings.TrimSpace(do.Doc.List[1].Text[k+1:]) == "request" {
							m := &message{}
							parseStruct(do, m, ap.Obj, false, "")
							ap.Request = m
						}
					} else if strings.TrimSpace(do.Doc.List[0].Text[j+1:]) == "response" {
						m := &message{}
						parseStruct(do, m, ap.Obj, false, "")
						ap.Reply = m
					}
					if ap.Reply != nil && ap.Request != nil {
						apis = append(apis, ap)
						ap = &api{}
						ap.Method = "GET/POST"
						ap.Obj = make(map[string]*message)
					}
				}
			}

		}
	}
	return

}

// parseStruct 递归解析结构体参数 TODO: 请求参数复杂类型
func parseStruct(d *ast.TypeSpec, m *message, objs map[string]*message, recursive bool, typeName string) {
	me := &message{}
	if st, ok := d.Type.(*ast.StructType); ok {
		if len(st.Fields.List) > 0 {
			for _, v := range st.Fields.List {
				f := field{}
				if v.Tag != nil {
					f.Name = v.Tag.Value
					if strings.Contains(f.Name, "required") {
						i := strings.Index(f.Name, "binding")
						if i > -1 {
							f.Name = f.Name[:i]
						}
						f.Doc = "Y"
					} else {
						f.Doc = "N"
					}
					s := strings.Trim(f.Name, "`")
					f.Name = strings.TrimSpace(strings.Replace(s, "json:", "", -1))
					f.Name = strings.Replace(f.Name, "\"", "", -1)
					//if recursive {
					//	f.Name = "\t" + f.Name
					//}
				}
				if t, ok := v.Type.(*ast.Ident); ok {
					f.Type = t.Name
				} else if t, ok := v.Type.(*ast.ArrayType); ok {
					if tt, ok := t.Elt.(*ast.Ident); ok {
						f.Type = "[]" + tt.Name
					}
				}

				if v.Comment != nil && len(v.Comment.List) > 0 {
					f.Note = v.Comment.List[0].Text
					f.Note = strings.Replace(f.Note, "/", "", -1)
				}

				if recursive {
					me.Fields = append(me.Fields, f)
				} else {
					m.Fields = append(m.Fields, f)
				}
				if t, ok := v.Type.(*ast.Ident); ok {
					if t.Obj != nil {
						if ot, ok := t.Obj.Decl.(*ast.TypeSpec); ok {
							parseStruct(ot, m, objs, true, t.Obj.Name)
						}
					}
				} else if t, ok := v.Type.(*ast.ArrayType); ok {
					if tt, ok := t.Elt.(*ast.Ident); ok {
						if tt.Obj != nil {
							if ott, ok := tt.Obj.Decl.(*ast.TypeSpec); ok {
								parseStruct(ott, m, objs, true, tt.Obj.Name)
							}
						}
					}
				}
			}
			if recursive {
				if _, ok := objs[typeName]; !ok {
					objs[typeName] = me
				}
			}
		}
	}
}

// Error print error and exit
func Error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("error:", s)
	os.Exit(1)
}
