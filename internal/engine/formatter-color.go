package engine

import (
	"fmt"
	"github.com/fatih/color"
)

// ColorFormatter 彩色格式化器
type ColorFormatter struct {
	successColor *color.Color
	errorColor   *color.Color
	warnColor    *color.Color
	infoColor    *color.Color
}

// NewColorFormatter 创建彩色格式化器
func NewColorFormatter() *ColorFormatter {
	return &ColorFormatter{
		successColor: color.New(color.FgGreen).Add(color.Bold),
		errorColor:   color.New(color.FgRed).Add(color.Bold),
		warnColor:    color.New(color.FgYellow).Add(color.Bold),
		infoColor:    color.New(color.FgBlue).Add(color.Bold),
	}
}

// FormatSuccess 格式化成功信息
func (f *ColorFormatter) FormatSuccess(message string) string {
	return f.successColor.Sprint("✅ " + message)
}

// FormatError 格式化错误信息
func (f *ColorFormatter) FormatError(message string) string {
	return f.errorColor.Sprint("❌ " + message)
}

// FormatWarning 格式化警告信息
func (f *ColorFormatter) FormatWarning(message string) string {
	return f.warnColor.Sprint("⚠️  " + message)
}

// FormatInfo 格式化信息
func (f *ColorFormatter) FormatInfo(message string) string {
	return f.infoColor.Sprint("ℹ️  " + message)
}

// FormatTitle 格式化标题
func (f *ColorFormatter) FormatTitle(title string) string {
	return color.New(color.FgCyan).Add(color.Bold).Sprint(title)
}

// FormatSeparator 格式化分隔符
func (f *ColorFormatter) FormatSeparator() string {
	return color.New(color.FgWhite).Sprint("───────────────────────────────────────────")
}

// FormatResult 格式化检测结果
func (f *ColorFormatter) FormatResult(name string, success bool, message string, duration string) string {
	if success {
		return fmt.Sprintf("%s %s (%s)", f.FormatSuccess(name), message, duration)
	}
	return fmt.Sprintf("%s %s (%s)", f.FormatError(name), message, duration)
}