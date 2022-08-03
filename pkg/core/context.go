package core

import (
	"io"
	"os"
)

type Context struct {
	configPath string
	writer     io.Writer
}

func NewDefault() *Context {
	return &Context{
		configPath: GetDefaultConfigPath(),
		writer:     os.Stdout,
	}
}

func (c *Context) ConfigPath() string {
	return c.configPath
}

func (c *Context) Writer() io.Writer {
	return c.writer
}
