// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

import (
	"context"
	"github.com/gogf/gf/internal/intlog"
	"io"

	"github.com/gogf/gf/os/gfile"
)

// Ctx is a chaining function,
// which sets the context for current logging.
func (l *Logger) E上下文(ctx context.Context, keys ...interface{}) *Logger {
	if ctx == nil {
		return l
	}
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.ctx = ctx
	if len(keys) > 0 {
		logger.SetCtxKeys(keys...)
	}
	return logger
}

// To is a chaining function,
// which redirects current logging content output to the specified <writer>.
func (l *Logger) E重定向(writer io.Writer) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetWriter(writer)
	return logger
}

// Path is a chaining function,
// which sets the directory path to <path> for current logging content output.
//
// Note that the parameter <path> is a directory path, not a file path.
func (l *Logger) E路径(path string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if path != "" {
		if err := logger.SetPath(path); err != nil {
			// panic(err)
			intlog.Error(err)
		}
	}
	return logger
}

// Cat is a chaining function,
// which sets the category to <category> for current logging content output.
// Param <category> can be hierarchical, eg: module/user.
func (l *Logger) E分类(category string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if logger.config.Path != "" {
		if err := logger.SetPath(gfile.Join(logger.config.Path, category)); err != nil {
			// panic(err)
			intlog.Error(err)
		}
	}
	return logger
}

// File is a chaining function,
// which sets file name <pattern> for the current logging content output.
func (l *Logger) E文件(file string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetFile(file)
	return logger
}

// Level is a chaining function,
// which sets logging level for the current logging content output.
func (l *Logger) E级别(level int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetLevel(level)
	return logger
}

// LevelStr is a chaining function,
// which sets logging level for the current logging content output using level string.
func (l *Logger) E级别文本(levelStr string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if err := logger.SetLevelStr(levelStr); err != nil {
		// panic(err)
		intlog.Error(err)
	}
	return logger
}

// Skip is a chaining function,
// which sets stack skip for the current logging content output.
// It also affects the caller file path checks when line number printing enabled.
func (l *Logger) E跳过(skip int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStackSkip(skip)
	return logger
}

// Stack is a chaining function,
// which sets stack options for the current logging content output .
func (l *Logger) E堆栈(enabled bool, skip ...int) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStack(enabled)
	if len(skip) > 0 {
		logger.SetStackSkip(skip[0])
	}
	return logger
}

// StackWithFilter is a chaining function,
// which sets stack filter for the current logging content output .
func (l *Logger) E堆栈过滤(filter string) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStack(true)
	logger.SetStackFilter(filter)
	return logger
}

// Stdout is a chaining function,
// which enables/disables stdout for the current logging content output.
// It's enabled in default.
func (l *Logger) E控制台(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// stdout printing is enabled if <enabled> is not passed.
	if len(enabled) > 0 && !enabled[0] {
		logger.config.StdoutPrint = false
	} else {
		logger.config.StdoutPrint = true
	}
	return logger
}

// Header is a chaining function,
// which enables/disables log header for the current logging content output.
// It's enabled in default.
func (l *Logger) E头(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// header is enabled if <enabled> is not passed.
	if len(enabled) > 0 && !enabled[0] {
		logger.SetHeaderPrint(false)
	} else {
		logger.SetHeaderPrint(true)
	}
	return logger
}

// Line is a chaining function,
// which enables/disables printing its caller file path along with its line number.
// The parameter <long> specified whether print the long absolute file path, eg: /a/b/c/d.go:23,
// or else short one: d.go:23.
func (l *Logger) E行号(long ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if len(long) > 0 && long[0] {
		logger.config.Flags |= F_FILE_LONG
	} else {
		logger.config.Flags |= F_FILE_SHORT
	}
	return logger
}

// Async is a chaining function,
// which enables/disables async logging output feature.
func (l *Logger) E异步(enabled ...bool) *Logger {
	logger := (*Logger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// async feature is enabled if <enabled> is not passed.
	if len(enabled) > 0 && !enabled[0] {
		logger.SetAsync(false)
	} else {
		logger.SetAsync(true)
	}
	return logger
}
