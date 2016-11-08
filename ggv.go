package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type (
	config struct {
		Path     string
		VarName  string
		Branch   string
		Hash     string
		Hostname string
		Now      string
	}
)

const (
	tmplSrc = "package main\n\nvar {{.VarName}} = `{{.Hash}} ({{.Branch}}) at {{.Hostname}} ({{.Now}})`\n"
)

func main() {
	tmpl := template.Must(template.New("package").Parse(tmplSrc))

	var cfg config
	checkError(cfg.load())

	w, err := os.Create(cfg.Path)
	checkError(err)
	defer w.Close()

	tmpl.Execute(w, cfg)
}

func (c *config) load() error {
	var err error
	var dir, fname string
	if dir, err = os.Getwd(); err != nil {
		return err
	}

	flag.StringVar(&fname, "file", "ggver_build_version.go",
		"Filename of generated go file")
	flag.StringVar(&c.VarName, "var", "ggverBuildVersion",
		"Variable to hold generated build version")

	flag.Parse()

	c.Path = filepath.Join(dir, fname)
	c.Now = formatNow()

	c.Branch = gitBranch()
	c.Hash = gitHash()
	c.Hostname = hostname()

	return nil
}

func formatNow() string {
	return time.Now().Format(time.RFC822Z)
}

func hostname() string {
	s, err := os.Hostname()
	if err != nil {
		fmt.Printf("Hostname error: %s", err)
		return "<unknown>"
	}
	return s
}

func gitBranch() string {
	out, err := runExternal("git", "branch", "--no-color")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "<no branch>"
	}
	for _, s := range strings.Split(out, "\n") {
		if strings.HasPrefix(s, "* ") {
			return s[2:]
		}
	}
	return "<no branch>"
}

func gitHash() string {
	out, err := runExternal("git", "log", "-1", "--no-color")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "<no hash>"
	}
	s := strings.Split(out, "\n")[0]
	return strings.Split(s, " ")[1]
}

func runExternal(name string, arg ...string) (string, error) {
	b, err := exec.Command(name, arg...).Output()
	if err != nil {
		err = fmt.Errorf(`command "%s %s" error: %v`, name, strings.Join(arg, " "), err)
	}
	return string(b), err
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
