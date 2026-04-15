package plugins

import (
	"github.com/yourusername/opsflow/internal/checks"
	"github.com/yourusername/opsflow/internal/types"
)

// SamplePlugin 示例插件
type SamplePlugin struct {
	BasePlugin
}

// GetScenarios 获取插件场景
func (p *SamplePlugin) GetScenarios() []types.Scenario {
	return []types.Scenario{
		{
			Name:        p.name + "-sample",
			Input:       "target",
			Checks:      []string{"sample-check"},
			Description: p.name + " 示例场景",
		},
	}
}

// GetChecks 获取插件检测
func (p *SamplePlugin) GetChecks() []checks.Check {
	return []checks.Check{
		&SampleCheck{},
	}
}

// SampleCheck 示例检测
type SampleCheck struct {
	BaseCheck
}

// BaseCheck 基础检测结构
type BaseCheck struct {
	name string
}

// Name 获取检测名称
func (c *BaseCheck) Name() string {
	return c.name
}

// SampleCheck.Name 实现
func (c *SampleCheck) Name() string {
	return "sample-check"
}

// SampleCheck.Run 实现
func (c *SampleCheck) Run(input types.Input) types.Result {
	return types.Result{
		Name:    "sample-check",
		Success: true,
		Message: "示例检测通过",
		Data:    map[string]interface{}{"input": input.Target},
	}
}