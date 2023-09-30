package test

import (
	"io"
	"os"
	"testing"

	"github.com/carneirofc/go-studies/azdevops"
)

func compareFmtCommand(t *testing.T, inp string, expect string, f func(string)) {

	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	t.Log("replacing original stdout with pipe...")
	os.Stdout = w
	if err != nil {
		t.Error("could not create pipe\n")
		return
	}

	f("command type log")

	err = w.Close()
	if err != nil {
		t.Error("failed to close w pipe")
	}

	log, _ := io.ReadAll(r)
	sLog := string(log)
	if sLog != expect {
		t.Errorf("result '%s' differs from expected '%s'\n", sLog, expect)
	}

	err = r.Close()
	if err != nil {
		t.Error("failed to close r pipe")
	}

	os.Stdout = originalStdout
	t.Log("restoring original stdout\n")
}

func Test_etc(t *testing.T) {
	compareFmtCommand(t, "command type log", "##[command]command type log\n", azdevops.LogFmtCommand)
}
