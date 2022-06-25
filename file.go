package ilysa

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

const (
	goCmd    = "go"
	cacheDir = ".ilysa_build"
)

var (
	//go:embed template/ilysamain.go
	mainTemplate []byte
)

type Pkg struct {
	GoFiles []string
}

func Invoke(path string, args ...string) error {
	goFiles, err := goList(path)
	if err != nil {
		return err
	}

	cacheDirAbsPath, err := filepath.Abs(cacheDir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(cacheDirAbsPath, 0755); err != nil {
		return err
	}

	tmpDir, err := os.MkdirTemp(cacheDirAbsPath, "")
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	for _, f := range goFiles {
		if err := copyFile(filepath.Join(tmpDir, f), filepath.Join(path, f)); err != nil {
			return err
		}
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "ilysamain.go"), mainTemplate, 0644); err != nil {
		return err
	}

	goFiles = append(goFiles, "ilysamain.go")

	outputFilename := exeName()
	outputPath := filepath.Join(tmpDir, outputFilename)

	if err := Compile(tmpDir, outputPath, goFiles); err != nil {
		return err
	}

	if err := RunCompiled(outputPath, args...); err != nil {
		return err
	}

	return nil
}

func Compile(path, compileTo string, goFiles []string) error {
	args := []string{"build", "-o", compileTo}
	args = append(args, goFiles...)

	stdErrBuf := &bytes.Buffer{}
	cmd := exec.Command(goCmd, args...)
	cmd.Dir = path
	cmd.Stderr = stdErrBuf

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error compiling: %s - %s", err, stdErrBuf.String())
	}
	return nil
}

func RunCompiled(exePath string, args ...string) error {
	cmd := exec.Command(exePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func goList(path string) ([]string, error) {
	args := []string{"list", "-e", "-json"}

	cmd := exec.Command(goCmd, args...)
	cmd.Dir = path

	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var pkg Pkg
	if err := json.Unmarshal(b, &pkg); err != nil {
		return nil, err
	}

	return pkg.GoFiles, nil
}

func exeName() string {
	filename := time.Now().Format("ilysa_20060102150405")
	if runtime.GOOS == "windows" {
		filename += ".exe"
	}
	return filename
}

// copyFile robustly copies the source file to the destination, overwriting the destination if necessary.
func copyFile(dst string, src string) error {
	from, err := os.Open(src)
	if err != nil {
		return fmt.Errorf(`can't copy %s: %v`, src, err)
	}
	defer from.Close()
	finfo, err := from.Stat()
	if err != nil {
		return fmt.Errorf(`can't stat %s: %v`, src, err)
	}
	to, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, finfo.Mode())
	if err != nil {
		return fmt.Errorf(`can't copy to %s: %v`, dst, err)
	}
	defer to.Close()
	_, err = io.Copy(to, from)
	if err != nil {
		return fmt.Errorf(`error copying %s to %s: %v`, src, dst, err)
	}
	return nil
}
