package loggolog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
)

const (
	PKG_APP_LOGGER_WR_FILE_PATH = "default.logger.log"
)

var (
	COLORReset  = "\033[0m"
	COLORRed    = "\033[31m"
	COLORYellow = "\033[33m"
	COLORPurple = "\033[35m"
	COLORCyan   = "\033[36m"
	COLORWhite  = "\033[37m"
	BOLDText    = "\033[1m"
)

var FINAL_WRITER_LOG io.Writer

type OutConsole struct {
	Data JSONLogFormat
}

type JSONLogFormat struct {
	Level string `json:"level"`
	Mgs   string `json:"msg"`
	Time  string `json:"time"`
}

// OutConsole - Console function for write data
func (cl *OutConsole) OutData() {
	head := fmt.Sprintf("%s%s%s", BOLDText, strings.ToUpper(cl.Data.Level), COLORReset)
	switch cl.Data.Level {
	case "info":
		{
			head = fmt.Sprintf("%s%s", COLORWhite, head)
			break
		}
	case "trace", "debug":
		{
			head = fmt.Sprintf("%s%s", COLORCyan, head)
			break
		}
	case "warning":
		{
			head = fmt.Sprintf("%s%s", COLORYellow, head)
			break
		}
	case "fatal", "error", "panic":
		{
			head = fmt.Sprintf("%s%s", COLORRed, head)
			break
		}
	default:
		{
			head = fmt.Sprintf("%s%s", COLORPurple, head)
		}
	}
	fmt.Printf("%s\t[%s] \t%s\n", head, cl.Data.Time, cl.Data.Mgs)
}

func (ou *OutConsole) Write(p []byte) (n int, err error) {
	err = json.Unmarshal(p, &ou.Data)
	if err != nil {
		return 0, err
	}
	ou.OutData()
	return len(p), nil
}

func InitCastomLogger(format logrus.Formatter, level logrus.Level, file bool, std bool) {
	logrus.SetFormatter(format)
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)

	if file {
		file, err := os.OpenFile(PKG_APP_LOGGER_WR_FILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logrus.Fatal(err.Error())
			panic(err)
		}
		FINAL_WRITER_LOG = io.MultiWriter(file)
	}
	if std {
		if FINAL_WRITER_LOG != nil {
			FINAL_WRITER_LOG = io.MultiWriter(FINAL_WRITER_LOG, &OutConsole{})
		} else {
			FINAL_WRITER_LOG = &OutConsole{}
		}
	}
	if FINAL_WRITER_LOG != nil {
		logrus.SetOutput(FINAL_WRITER_LOG)
	}
}

func AddOut(writer io.Writer) {
	if writer != nil {
		FINAL_WRITER_LOG = io.MultiWriter(FINAL_WRITER_LOG, writer)
		logrus.SetOutput(FINAL_WRITER_LOG)
	}
}
