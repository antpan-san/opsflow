package plugins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yourusername/opsflow/internal/checks"
	"github.com/yourusername/opsflow/internal/types"
)

// PluginManager 插件管理器
type PluginManager struct {
	pluginDir string
	plugins   map[string]Plugin
}

// NewPluginManager 创建插件管理器
func NewPluginManager(pluginDir string) *PluginManager {
	if pluginDir == "" {
		pluginDir = filepath.Join(os.Getenv("HOME"), ".ops", "plugins")
	}

	return &PluginManager{
		pluginDir: pluginDir,
		plugins:   make(map[string]Plugin),
	}
}

// GetPluginDir 获取插件目录
func (pm *PluginManager) GetPluginDir() string {
	return pm.pluginDir
}

// LoadPlugin 加载指定插件
func (pm *PluginManager) LoadPlugin(pluginName string) error {
	pluginPath := filepath.Join(pm.pluginDir, pluginName)

	// 检查插件目录是否存在
	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		return fmt.Errorf("插件不存在: %s", pluginName)
	}

	// TODO: 实现动态加载逻辑
	// 这里可以使用 Go 的 plugin 包实现动态加载
	// 目前先实现占位逻辑

	plugin := &SamplePlugin{
		BasePlugin: NewBasePlugin(pluginName, "1.0.0", fmt.Sprintf("%s 插件", pluginName)),
	}

	err := plugin.Load()
	if err != nil {
		return fmt.Errorf("加载插件失败: %s, 错误: %v", pluginName, err)
	}

	pm.plugins[pluginName] = plugin
	return nil
}

// UnloadPlugin 卸载插件
func (pm *PluginManager) UnloadPlugin(pluginName string) error {
	plugin, exists := pm.plugins[pluginName]
	if !exists {
		return fmt.Errorf("插件未加载: %s", pluginName)
	}

	err := plugin.Unload()
	if err != nil {
		return fmt.Errorf("卸载插件失败: %s, 错误: %v", pluginName, err)
	}

	delete(pm.plugins, pluginName)
	return nil
}

// ListPlugins 列出所有插件
func (pm *PluginManager) ListPlugins() []Plugin {
	plugins := make([]Plugin, 0, len(pm.plugins))
	for _, plugin := range pm.plugins {
		plugins = append(plugins, plugin)
	}
	return plugins
}

// GetPlugin 获取插件
func (pm *PluginManager) GetPlugin(pluginName string) (Plugin, error) {
	plugin, exists := pm.plugins[pluginName]
	if !exists {
		return nil, fmt.Errorf("插件未加载: %s", pluginName)
	}
	return plugin, nil
}

// GetScenarios 获取所有插件的场景
func (pm *PluginManager) GetScenarios() []types.Scenario {
	var scenarios []types.Scenario
	for _, plugin := range pm.plugins {
		scenarios = append(scenarios, plugin.GetScenarios()...)
	}
	return scenarios
}

// GetChecks 获取所有插件的检测
func (pm *PluginManager) GetChecks() []checks.Check {
	var checkList []checks.Check
	for _, plugin := range pm.plugins {
		checkList = append(checkList, plugin.GetChecks()...)
	}
	return checkList
}

// InstallPlugin 安装插件
func (pm *PluginManager) InstallPlugin(pluginName string) error {
	// 创建插件目录
	pluginPath := filepath.Join(pm.pluginDir, pluginName)
	if err := os.MkdirAll(pluginPath, 0755); err != nil {
		return fmt.Errorf("创建插件目录失败: %v", err)
	}

	// TODO: 实现插件下载和安装逻辑
	// 这里可以实现从远程仓库下载插件

	return nil
}

// RemovePlugin 移除插件
func (pm *PluginManager) RemovePlugin(pluginName string) error {
	// 如果插件已加载，先卸载
	if _, exists := pm.plugins[pluginName]; exists {
		if err := pm.UnloadPlugin(pluginName); err != nil {
			return err
		}
	}

	// 删除插件目录
	pluginPath := filepath.Join(pm.pluginDir, pluginName)
	if err := os.RemoveAll(pluginPath); err != nil {
		return fmt.Errorf("删除插件目录失败: %v", err)
	}

	return nil
}

// UpdatePlugin 更新插件
func (pm *PluginManager) UpdatePlugin(pluginName string) error {
	// 先卸载
	if _, exists := pm.plugins[pluginName]; exists {
		if err := pm.UnloadPlugin(pluginName); err != nil {
			return err
		}
	}

	// TODO: 实现插件更新逻辑
	// 这里可以实现从远程仓库更新插件

	// 重新加载
	return pm.LoadPlugin(pluginName)
}

// ScanPlugins 扫描插件目录
func (pm *PluginManager) ScanPlugins() ([]string, error) {
	entries, err := os.ReadDir(pm.pluginDir)
	if err != nil {
		return nil, fmt.Errorf("读取插件目录失败: %v", err)
	}

	var plugins []string
	for _, entry := range entries {
		if entry.IsDir() {
			plugins = append(plugins, entry.Name())
		}
	}

	return plugins, nil
}

// LoadAllPlugins 加载所有插件
func (pm *PluginManager) LoadAllPlugins() error {
	plugins, err := pm.ScanPlugins()
	if err != nil {
		return err
	}

	var errors []string
	for _, pluginName := range plugins {
		if err := pm.LoadPlugin(pluginName); err != nil {
			errors = append(errors, fmt.Sprintf("%s: %v", pluginName, err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("加载插件失败: %s", strings.Join(errors, "; "))
	}

	return nil
}