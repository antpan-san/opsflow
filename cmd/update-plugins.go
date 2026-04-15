package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/opsflow/internal/plugins"
)

var updatePluginsCmd = &cobra.Command{
	Use:   "update-plugins",
	Short: "更新所有插件",
	Long: `更新所有已安装的插件。

示例：
  ops update-plugins`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建插件管理器
		pluginDir := cmd.Flag("plugin-dir").Value.String()
		pm := plugins.NewPluginManager(pluginDir)

		// 加载所有插件
		err := pm.LoadAllPlugins()
		if err != nil {
			fmt.Fprintf(os.Stderr, "加载插件失败: %v\n", err)
			os.Exit(1)
		}

		// 获取所有插件
		installedPlugins := pm.ListPlugins()
		if len(installedPlugins) == 0 {
			fmt.Println("未安装任何插件")
			return
		}

		fmt.Println("正在更新插件...")
		var updatedCount int
		for _, plugin := range installedPlugins {
			pluginName := plugin.Name()
			fmt.Printf("正在更新: %s\n", pluginName)

			err := pm.UpdatePlugin(pluginName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "更新插件 %s 失败: %v\n", pluginName, err)
				continue
			}

			updatedCount++
			fmt.Printf("✅ 插件更新成功: %s\n", pluginName)
		}

		fmt.Printf("\n更新完成: %d/%d 个插件更新成功\n", updatedCount, len(installedPlugins))
	},
}

func init() {
	rootCmd.AddCommand(updatePluginsCmd)

	// 添加标志
	updatePluginsCmd.Flags().String("plugin-dir", "", "插件目录")
}