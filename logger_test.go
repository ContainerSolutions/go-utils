package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

var logFile = "test.log"

func cleanup() {
	os.Remove(logFile)
}

func TestMain(m *testing.M) {
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestLoggerWithFile(t *testing.T) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatal(err)
	}

	logger := NewLogger(file, LogLevel.DEBUG)
	if err != nil {
		t.Fatal(err)
	}

	logger.Debug("hello")
	b, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatal(err)
	}

	dump := bytes.NewBuffer(b).String()
	if !regexp.MustCompile("hello").MatchString(dump) {
		t.Errorf("trace message not found in string:\n%s", dump)
	}
}

func TestPrint(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.TRACE)
	logger.Trace("print this message")
	dump := b.String()
	if !regexp.MustCompile("print this message").MatchString(dump) {
		t.Errorf("trace message not found in string:\n%s", dump)
	}
}

func TestTrace(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.TRACE)
	logger.Trace("some trace message")
	dump := b.String()
	if !regexp.MustCompile("some trace message").MatchString(dump) {
		t.Errorf("trace message not found in string:\n%s", dump)
	}
}

func TestDebug(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.DEBUG)
	logger.Debug("some debug message")
	dump := b.String()
	if !regexp.MustCompile("some debug message").MatchString(dump) {
		t.Errorf("debug message not found in string:\n%s", dump)
	}
}

func TestInfo(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.INFO)
	logger.Info("some info message")
	dump := b.String()
	if !regexp.MustCompile("some info message").MatchString(dump) {
		t.Errorf("info message not found in string:\n%s", dump)
	}
}

func TestWarn(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.WARN)
	logger.Warn("some warn message")
	dump := b.String()
	if !regexp.MustCompile("some warn message").MatchString(dump) {
		t.Errorf("warn message not found in string:\n%s", dump)
	}
}

func TestError(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.ERROR)
	logger.Error("some error message")
	dump := b.String()
	if !regexp.MustCompile("some error message").MatchString(dump) {
		t.Errorf("error message not found in string:\n%s", dump)
	}
}

func TestFatal(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.FATAL)
	logger.Fatal("some fatal message")
	dump := b.String()
	if !regexp.MustCompile("some fatal message").MatchString(dump) {
		t.Errorf("fatal message not found in string:\n%s", dump)
	}
}

func TestLogLevel(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b, LogLevel.TRACE)

	logger.Level = LogLevel.DEBUG
	logger.Trace("trace me!")
	dump := b.String()
	if regexp.MustCompile("trace me!").MatchString(dump) {
		t.Errorf("expected empty; got %s", dump)
	}

	logger.Level = LogLevel.INFO
	logger.Debug("debugging")
	dump = b.String()
	if regexp.MustCompile("debugging").MatchString(dump) {
		t.Errorf("expected empty; got %s", dump)
	}

	logger.Level = LogLevel.WARN
	logger.Info("some info")
	dump = b.String()
	if regexp.MustCompile("some info").MatchString(dump) {
		t.Errorf("expected empty; got %s", dump)
	}

	logger.Level = LogLevel.ERROR
	logger.Warn("warning!")
	dump = b.String()
	if regexp.MustCompile("warning!").MatchString(dump) {
		t.Errorf("expected empty; got %s", dump)
	}

	logger.Level = LogLevel.FATAL
	logger.Error("some error")
	dump = b.String()
	if regexp.MustCompile("some error").MatchString(dump) {
		t.Errorf("expected empty; got %s", dump)
	}
}
