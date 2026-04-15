package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/opsflow/internal/plugins"
)

var listPluginsCmd = &cobra.Command{
	Use:   "list-plugins",
	Short: "列出已安装的插件",
	Long: `列出所有已安装的插件。

示例：
  ops list-plugins`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建插件管理器
		pluginDir := cmd.Flag("plugin-dir").Value.String()
		pm := plugins.NewPluginManager(pluginDir)

		// 扫描插件
		pluginNames, err := pm.ScanPlugins()
		if err != nil {
			fmt.Fprintf(os.Stderr, "扫描插件失败: %v\n", err)
			os.Exit(1)
		}

		if len(pluginNames) == 0 {
			fmt.Println("未安装任何插件")
			return
		}

		fmt.Println("已安装的插件:")
		for _, name := range pluginNames {
			fmt.Printf("  - %s\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listPluginsCmd)

	// 添加标志
	listPluginsCmd.Flags().String("plugin-dir", "", "插件目录")
}