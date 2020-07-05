package logging

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"uit_payment/lib/env"
	libtime "uit_payment/lib/lib_time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type GrayLogLogging struct {
	Logging logrus.FieldLogger
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	response   []byte
}

func InitLoggerGrayLog() GrayLogLogging {
	graylog := GrayLogLogging{}

	return graylog
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK, []byte{}}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(data []byte) (int, error) {
	lrw.response = data
	int, err := lrw.ResponseWriter.Write(data)

	return int, err
}

func NewGrayLogWritter(statusCode int) io.Writer {
	rotateLog := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./log/%v_payment.log", libtime.GetTime()),
		MaxSize:    env.GetFileMaxSize(),
		MaxBackups: env.GetFileBackups(),
		MaxAge:     env.GetFileMaxAge(),
		Compress:   false,
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	mw := io.MultiWriter(os.Stdout, rotateLog)
	log.SetOutput(mw)
	logrus.SetOutput(mw)

	// graylogAddr := env.GetGraylogURL()
	// if graylogAddr == "" {
	// 	mw := io.MultiWriter(os.Stdout, rotateLog)
	// 	log.SetOutput(mw)
	// 	log.Fatalf("Graylog address is undenfined")
	// 	return nil
	// }

	// gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	// if err != nil {
	// 	mw := io.MultiWriter(os.Stdout, rotateLog)
	// 	log.SetOutput(mw)
	// 	logrus.SetOutput(mw)
	// 	log.Fatalf("gelf.NewUDPWriter: %s", err)
	// } else {
	// 	gelfWriter.Facility = fmt.Sprintf("%s-%s", env.GetFacility(), env.GetEnvironment())
	// 	mw := io.MultiWriter(os.Stdout, rotateLog, gelfWriter)
	// 	logrus.SetOutput(mw)
	// 	log.SetOutput(io.MultiWriter(os.Stdout, rotateLog))
	// }

	logLevel, err := logrus.ParseLevel(env.GetLogLevel())
	if err != nil {
		logrus.WithError(err).Panicln("LogLevel")
	}
	logrus.SetLevel(logLevel)

	return mw
}

func (l *GrayLogLogging) Logger(next http.Handler) http.Handler {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		r.Body = rdr2

		path := r.URL.Path
		start := time.Now()
		method := r.Method

		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := lrw.statusCode
		clientIP := r.RemoteAddr
		clientUserAgent := r.UserAgent()
		referer := r.Referer()

		dataLength := 100
		response := string(lrw.response)

		if dataLength < 0 {
			dataLength = 0
		}

		newLog := logrus.New()

		mWriter := NewGrayLogWritter(statusCode)
		newLog.SetOutput(io.MultiWriter(mWriter))
		l.Logging = newLog

		entry := l.Logging.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
			"request":    readBody(rdr1),
			"response":   response,
		})

		if statusCode > 499 {
			entry.Error()
		} else if statusCode > 399 {
			entry.Warn()
		} else {
			entry.Info()
		}
	})
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func WithError(err error) *logrus.Entry {
	return logrus.WithField(logrus.ErrorKey, err)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		log.Println("DEBUG", args)
	}
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	log.Println("INFO", args)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}
