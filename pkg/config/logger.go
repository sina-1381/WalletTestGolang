package config

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/robfig/cron/v3"
)

type Logger struct {
	Log    *slog.Logger
	File   *os.File
	folder string
	buf    *bufio.Writer
	cron   *cron.Cron
}

func NewLogger(folder string, location *time.Location) (*Logger, error) {
	log := &Logger{folder: folder, cron: cron.New(cron.WithLocation(location))}
	if err := log.openFile(); err != nil {
		return nil, err
	}
	log.cron.AddFunc("0 0 * * *", func() { log.openFile() })
	log.cron.Start()
	return log, nil
}

func (l *Logger) openFile() error {
	l.closeFile()
	path := l.path()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.buf = bufio.NewWriter(file)
	l.File = file
	l.Log = slog.New(slog.NewJSONHandler(l.buf, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return nil
}

func (l *Logger) Close() {
	l.cron.Stop()
	l.closeFile()
}

func (l *Logger) path() string {
	return fmt.Sprintf("%s%s-%s.log", l.folder, time.Now().Format("2006-01-02"), time.Now().Format("1504"))
}

func (l *Logger) closeFile() {
	if l.buf != nil {
		l.buf.Flush()
	}
	if l.File != nil {
		l.File.Close()
	}
}

// ============ STARTUP/SHUTDOWN ============

func (l *Logger) Startup(version, port string) {
	l.Log.Info("server_starting",
		"type", "system",
		"event", "startup",
		"version", version,
		"port", port,
		"time", time.Now().Format(time.RFC3339),
	)
}

func (l *Logger) Shutdown(reason string) {
	l.Log.Info("server_shutdown",
		"type", "system",
		"event", "shutdown",
		"reason", reason,
	)
}
