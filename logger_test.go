package logger

import (
	"os"
	"testing"

	"github.com/bukerdevid/logger"
)

var PKG_APP_LOGGER_WR_DFILE_PATH = "salair-back.re.log"

func TestWriteLog(t *testing.T) {
	logger.InitCastomLogger(&logger.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logger.TraceLevel, true, true)
	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logger.Entry)
			logger.WithFields(logger.Fields{
				"omg":         true,
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
			}).Error("The ice breaks!")
		}
		logger.require.True(t, true)
	}()

	for count := 0; count < 10; count++ {
		logger.Info("Random data - ", count)
		logger.Debug("Random data - ", count)
		logger.Warn("Random data - ", count)
		logger.Error("Random data - ", count)
	}
	logger.Panic("Check pnic")
}

func TestAddOut(t *testing.T) {
	logger.InitCastomLogger(&logger.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logger.TraceLevel, true, true)
	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logger.Entry)
			logger.WithFields(logger.Fields{
				"omg":         true,
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
			}).Error("The ice breaks!")
		}
		logger.require.True(t, true)
	}()

	for count := 0; count < 10; count++ {
		logger.Info("Random data - ", count)
		logger.Debug("Random data - ", count)
		logger.Warn("Random data - ", count)
		logger.Error("Random data - ", count)
		if count == 5 {
			file, err := os.OpenFile(PKG_APP_LOGGER_WR_DFILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				logger.Fatal(err.Error())
				logger.require.True(t, true)
			}
			logger.AddOut(file)
		}
	}
	logger.Panic("Check pnic")
}
