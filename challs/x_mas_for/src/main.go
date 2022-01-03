package main

import (
	"fmt"
	"io"
	golog "log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/iver-wharf/wharf-core/pkg/ginutil"
	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/logger/consolepretty"
)

var log = logger.NewScoped("main")

func main() {
	logCfg := consolepretty.DefaultConfig
	logCfg.CallerMaxLength = 16
	logCfg.CallerMinLength = 16
	logger.AddOutput(logger.LevelDebug, consolepretty.New(logCfg))

	cfg, err := loadConfig()
	if err != nil {
		log.Error().WithError(err).Message("Failed to load config.")
		os.Exit(1)
	}

	gin.DefaultWriter = ginutil.DefaultLoggerWriter
	gin.DefaultErrorWriter = ginutil.DefaultLoggerWriter
	golog.SetOutput(newEndlessLoggerWriter(logger.NewScoped("endless"), logger.LevelInfo))
	golog.SetFlags(0)

	r := gin.New()
	r.Use(
		ginutil.DefaultLoggerHandler,
		ginutil.RecoverProblem,
		ipBlockerMiddleware{cfg.cidrs}.filterMiddleware,
	)

	r.LoadHTMLGlob("assets/html/*.html")
	r.GET("/", indexHandler{cfg.flag}.handleIndex)
	r.NoRoute(handle404)

	if err := endless.ListenAndServe(cfg.BindAddress, r); err != nil {
		log.Error().WithError(err).Message("Failed to start web server.")
	}
}

type ipBlockerMiddleware struct {
	nets []*net.IPNet
}

func (m ipBlockerMiddleware) filterMiddleware(c *gin.Context) {
	if m.shouldBlockIP(c) {
		c.HTML(http.StatusForbidden, "403-ip-block.html", struct {
			Nets []*net.IPNet
		}{
			Nets: m.nets,
		})
		c.Abort()
	}
}

func (m ipBlockerMiddleware) shouldBlockIP(c *gin.Context) bool {
	ipStr := c.ClientIP()
	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Warn().WithString("ip", ipStr).Message("Failed to parse client IP.")
		return true
	}
	for _, n := range m.nets {
		if n.Contains(ip) {
			return false
		}
	}
	return true
}

type indexHandler struct {
	flag string
}

func (h indexHandler) handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", struct {
		Flag string
	}{
		Flag: h.flag,
	})
}

func handle404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

type endlessLoggerWriter struct {
	logger       logger.Logger
	defaultLevel logger.Level
}

func newEndlessLoggerWriter(log logger.Logger, defaultLevel logger.Level) io.Writer {
	return endlessLoggerWriter{log, defaultLevel}
}

func (w endlessLoggerWriter) Write(p []byte) (n int, err error) {
	str := strings.TrimSpace(string(p))
	event := logger.NewEventFromLogger(w.logger, w.defaultLevel)

	var pidStr string
	var pid int
	if _, err := fmt.Sscan(str, &pidStr); err == nil {
		if pid, err = strconv.Atoi(pidStr); err == nil {
			event = event.WithInt("pid", pid)
			str = str[len(pidStr)+1:]
		}
	}
	event.Message(str)
	return len(p), nil
}
