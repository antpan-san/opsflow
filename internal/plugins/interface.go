package plugins

import (
	"github.com/yourusername/opsflow/internal/checks"
	"github.com/yourusername/opsflow/internal/types"
)

// Plugin 插件接口
type Plugin interface {
	// 基本信息
	Name() string
	Version() string
	Description() string

	// 生命周期
	Load() error
	Unload() error

	// 获取插件内容
	GetScenarios() []types.Scenario
	GetChecks() []checks.Check

	// 插件状态
	IsLoaded() bool
	GetError() error
}

// BasePlugin 基础插件结构
type BasePlugin struct {
	name        string
	version     string
	description string
	loaded      bool
	err         error
}

// NewBasePlugin 创建基础插件
func NewBasePlugin(name, version, description string) BasePlugin {
	return BasePlugin{
		name:        name,
		version:     version,
		description: description,
		loaded:      false,
	}
}

// Name 获取插件名称
func (p *BasePlugin) Name() string {
	return p.name
}

// Version 获取插件版本
func (p *BasePlugin) Version() string {
	return p.version
}

// Description 获取插件描述
func (p *BasePlugin) Description() string {
	return p.description
}

// IsLoaded 检查插件是否已加载
func (p *BasePlugin) IsLoaded() bool {
	return p.loaded
}

// GetError 获取插件错误
func (p *BasePlugin) GetError() error {
	return p.err
}

// Load 加载插件
func (p *BasePlugin) Load() error {
	p.loaded = true
	return nil
}

// Unload 卸载插件
func (p *BasePlugin) Unload() error {
	p.loaded = false
	return nil
}