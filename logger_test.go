package loggolog

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var PKG_APP_LOGGER_WR_DFILE_PATH = "salair-back.re.log"

func TestWriteLog(t *testing.T) {
	InitCastomLogger(&logrus.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logrus.TraceLevel, true, true)
	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			logrus.WithFields(logrus.Fields{
				"omg":         true,
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
			}).Error("The ice breaks!")
		}
		require.True(t, true)
	}()

	for count := 0; count < 10; count++ {
		logrus.Info("Random data - ", count)
		logrus.Debug("Random data - ", count)
		logrus.Warn("Random data - ", count)
		logrus.Error("Random data - ", count)
	}
	logrus.Panic("Check pnic")
}

func TestAddOut(t *testing.T) {
	InitCastomLogger(&logrus.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logrus.TraceLevel, true, true)
	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			logrus.WithFields(logrus.Fields{
				"omg":         true,
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
			}).Error("The ice breaks!")
		}
		require.True(t, true)
	}()

	for count := 0; count < 10; count++ {
		logrus.Info("Random data - ", count)
		logrus.Debug("Random data - ", count)
		logrus.Warn("Random data - ", count)
		logrus.Error("Random data - ", count)
		if count == 5 {
			file, err := os.OpenFile(PKG_APP_LOGGER_WR_DFILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				logrus.Fatal(err.Error())
				require.True(t, true)
			}
			AddOut(file)
		}
	}
	logrus.Panic("Check pnic")
}
