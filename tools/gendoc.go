package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FuncInfo 函数信息
type FuncInfo struct {
	Name    string  // 函数名
	Comment string  // 注释
	URL     string  // 数据源URL
	Params  []Param // 参数列表
	File    string  // 所在文件
}

// Param 参数信息
type Param struct {
	Name string
	Type string
}

// ModuleDoc 模块文档
type ModuleDoc struct {
	Name  string
	Funcs []FuncInfo
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run gendoc.go <模块目录>")
		os.Exit(1)
	}

	moduleDir := os.Args[1]
	moduleName := filepath.Base(moduleDir)

	// 扫描模块目录
	funcs, err := scanModule(moduleDir)
	if err != nil {
		fmt.Printf("扫描模块失败: %v\n", err)
		os.Exit(1)
	}

	// 生成文档
	doc := generateDoc(moduleName, funcs)

	// 输出到docs目录
	outputPath := filepath.Join("docs", moduleName+".md")
	if err := os.WriteFile(outputPath, []byte(doc), 0644); err != nil {
		fmt.Printf("写入文档失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("文档已生成: %s (%d个接口)\n", outputPath, len(funcs))
}

// scanModule 扫描模块目录下的所有Go文件
func scanModule(dir string) ([]FuncInfo, error) {
	var funcs []FuncInfo

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理.go文件，排除_test.go
		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			fileFuncs, err := parseFile(path)
			if err != nil {
				fmt.Printf("警告: 解析文件 %s 失败: %v\n", path, err)
				return nil
			}
			funcs = append(funcs, fileFuncs...)
		}

		return nil
	})

	// 按函数名排序
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].Name < funcs[j].Name
	})

	return funcs, err
}

// parseFile 解析单个Go文件
func parseFile(filename string) ([]FuncInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var funcs []FuncInfo

	// 遍历所有声明
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || funcDecl.Name == nil {
			continue
		}

		// 只处理导出函数（首字母大写）
		if !funcDecl.Name.IsExported() {
			continue
		}

		funcInfo := FuncInfo{
			Name: funcDecl.Name.Name,
			File: filename,
		}

		// 提取注释
		if funcDecl.Doc != nil {
			comment := funcDecl.Doc.Text()
			funcInfo.Comment, funcInfo.URL = parseComment(comment)
		}

		// 提取参数
		if funcDecl.Type.Params != nil {
			for _, field := range funcDecl.Type.Params.List {
				paramType := exprToString(field.Type)
				for _, name := range field.Names {
					funcInfo.Params = append(funcInfo.Params, Param{
						Name: name.Name,
						Type: paramType,
					})
				}
			}
		}

		funcs = append(funcs, funcInfo)
	}

	return funcs, nil
}

// parseComment 解析注释，提取描述和URL
func parseComment(comment string) (desc, url string) {
	lines := strings.Split(strings.TrimSpace(comment), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			url = line
		} else if line != "" && desc == "" {
			desc = line
		}
	}

	return
}

// exprToString 将AST表达式转换为字符串
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.MapType:
		return "map[" + exprToString(t.Key) + "]" + exprToString(t.Value)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return "unknown"
	}
}

// generateDoc 生成markdown文档
func generateDoc(moduleName string, funcs []FuncInfo) string {
	var sb strings.Builder

	// 文档标题
	sb.WriteString(fmt.Sprintf("# %s 模块\n\n", moduleName))
	sb.WriteString(fmt.Sprintf("> 本模块共有 %d 个接口\n\n", len(funcs)))

	// 目录
	sb.WriteString("## 目录\n\n")
	for _, f := range funcs {
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", f.Name, strings.ToLower(f.Name)))
	}
	sb.WriteString("\n---\n\n")

	// 每个函数的详细文档
	for _, f := range funcs {
		sb.WriteString(generateFuncDoc(f))
		sb.WriteString("\n---\n\n")
	}

	return sb.String()
}

// generateFuncDoc 生成单个函数的文档
func generateFuncDoc(f FuncInfo) string {
	var sb strings.Builder

	// 函数名
	sb.WriteString(fmt.Sprintf("## %s\n\n", f.Name))

	// 描述
	if f.Comment != "" {
		sb.WriteString(fmt.Sprintf("**描述**: %s\n\n", f.Comment))
	}

	// 数据源
	if f.URL != "" {
		sb.WriteString(fmt.Sprintf("**数据源**: %s\n\n", f.URL))
	}

	// 输入参数
	sb.WriteString("### 输入参数\n\n")
	if len(f.Params) > 0 {
		sb.WriteString("| 参数名 | 类型 | 说明 |\n")
		sb.WriteString("|------|------|------|\n")
		for _, p := range f.Params {
			paramDesc := getParamDesc(f.Name, p.Name)
			sb.WriteString(fmt.Sprintf("| %s | %s | %s |\n", p.Name, p.Type, paramDesc))
		}
	} else {
		sb.WriteString("无参数\n")
	}
	sb.WriteString("\n")

	// 输出参数
	sb.WriteString("### 输出参数\n\n")
	sb.WriteString("返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：\n\n")
	sb.WriteString("具体字段请参考示例输出。\n\n")

	// 代码示例
	sb.WriteString("### 代码示例\n\n")
	sb.WriteString("```go\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"fmt\"\n")
	sb.WriteString(fmt.Sprintf("\t\"github.com/BlakeLiAFK/akshare/%s\"\n", getModuleName(f.File)))
	sb.WriteString(")\n\n")
	sb.WriteString("func main() {\n")

	// 生成函数调用
	callStr := generateFuncCall(f)
	sb.WriteString(fmt.Sprintf("\tdata, err := %s\n", callStr))
	sb.WriteString("\tif err != nil {\n")
	sb.WriteString("\t\tpanic(err)\n")
	sb.WriteString("\t}\n\n")
	sb.WriteString("\tfmt.Println(data)\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	return sb.String()
}

// getModuleName 从文件路径获取模块名
func getModuleName(filePath string) string {
	// 获取文件所在目录
	dir := filepath.Dir(filePath)
	// 返回最后一个目录名
	return filepath.Base(dir)
}

// generateFuncCall 生成函数调用示例
func generateFuncCall(f FuncInfo) string {
	moduleName := getModuleName(f.File)

	if len(f.Params) == 0 {
		return fmt.Sprintf("%s.%s()", moduleName, f.Name)
	}

	// 生成参数示例值
	params := make([]string, 0, len(f.Params))
	for _, p := range f.Params {
		params = append(params, getExampleValue(p))
	}

	return fmt.Sprintf("%s.%s(%s)", moduleName, f.Name, strings.Join(params, ", "))
}

// getExampleValue 获取参数的示例值
func getExampleValue(p Param) string {
	switch p.Type {
	case "string":
		// 根据参数名推断示例值
		switch {
		case strings.Contains(strings.ToLower(p.Name), "symbol"), strings.Contains(strings.ToLower(p.Name), "code"):
			return `"000001"`
		case strings.Contains(strings.ToLower(p.Name), "date"):
			return `"20230101"`
		case strings.Contains(strings.ToLower(p.Name), "period"):
			return `"daily"`
		case strings.Contains(strings.ToLower(p.Name), "adjust"):
			return `"qfq"`
		default:
			return `""`
		}
	case "int", "int64":
		return "1"
	case "float64":
		return "1.0"
	case "bool":
		return "false"
	default:
		return "nil"
	}
}

// getParamDesc 获取参数描述
func getParamDesc(funcName, paramName string) string {
	paramLower := strings.ToLower(paramName)

	switch {
	case strings.Contains(paramLower, "symbol"), strings.Contains(paramLower, "code"):
		return "股票/基金代码"
	case strings.Contains(paramLower, "start") && strings.Contains(paramLower, "date"):
		return "开始日期，格式：YYYYMMDD"
	case strings.Contains(paramLower, "end") && strings.Contains(paramLower, "date"):
		return "结束日期，格式：YYYYMMDD"
	case strings.Contains(paramLower, "period"):
		return "周期：daily/weekly/monthly"
	case strings.Contains(paramLower, "adjust"):
		return "复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权)"
	case strings.Contains(paramLower, "indicator"):
		return "指标类型"
	case strings.Contains(paramLower, "date"):
		return "日期，格式：YYYYMMDD"
	default:
		return "-"
	}
}
