package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func main() {
	//zerolog.TimeFieldFormat = zerolog.TimeFieldFormat
	// log.Print("hello world")
	// log.Debug().Str("Scale", "asdfasdf").Msg("debug print")

	err := errors.New("first error")
	err2 := fmt.Errorf("second error:%w", err)
	fmt.Printf("err2: %v\n", err2)
	err3 := errors.Wrapf(err2, "this is err3 %s", "zhouli")
	fmt.Printf("err3: %v\n", err3.Error())

	logFile, err4 := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, (os.ModeAppend)|(os.ModePerm))
	if err4 != nil {
		fmt.Printf("open log file errerr: %v\n", err4)
		return
	}

	multi := zerolog.MultiLevelWriter(logFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	for i := 0; i < 100; i++ {
		logger.Info().Msgf("Hello world! %d", i)
		time.Sleep(time.Second)
	}

}
