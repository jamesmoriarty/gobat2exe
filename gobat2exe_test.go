package gobat2exe

import (
	"os"
	"testing"
	"time"
)

func TestBat2Exe(t *testing.T) {
	Bat2Exe("test/test.bat")

	time.Sleep(1 * time.Second)

	if _, err := os.Stat("test/test.bat.exe"); os.IsNotExist(err) {
		t.Errorf(err.Error())
	}
}
