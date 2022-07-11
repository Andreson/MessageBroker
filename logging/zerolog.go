package log

import (
	"github.com/rs/zerolog/log"
)

func init() {
	//	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

}

type Zerolog struct {
}

func (z Zerolog) Info(msg string) {
	log.Info().Msg(msg)
}

func (z Zerolog) Infof(format string, v ...interface{}) {
	log.Info().Msgf(format, v)
}

func (z Zerolog) Error(msg string) {
	log.Error().Msg(msg)
}
func (z Zerolog) Errorf(msg string, v ...interface{}) {
	log.Error().Msgf(msg, v)
}
func (z Zerolog) Warn(msg string) {
	log.Warn().Msg(msg)
}

func (z Zerolog) Warnf(msg string, v ...interface{}) {
	log.Warn().Msgf(msg)
}

func (z Zerolog) Debug(msg string) {
	log.Debug().Msg(msg)
}

func (z Zerolog) Trace(msg string) {
	log.Trace().Msg(msg)
}
