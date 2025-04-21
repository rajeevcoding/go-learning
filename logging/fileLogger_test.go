package logging

import (
	"os"
	"strings"
	"testing"
)

func TestFileLogger_Info(t *testing.T) {
	fileName := "testlog.log"
	os.Remove(fileName)

	logger, err := NewFileLogger(fileName, false)

	if err != nil {
		t.Fatalf("Failed to create logger at: %s", fileName)
	}

	defer os.Remove(fileName)
	defer logger.Close()

	testMsg := "This is a test message"
	logger.Info(testMsg)

	content, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read log file at:  %s", fileName)
	}

	if !strings.Contains(string(content), testMsg) {
		t.Errorf("Expected log message \"%s\" not found in the logfile: %s", testMsg, fileName)
	}
}
