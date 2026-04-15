package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/opsflow/internal/plugins"
)

var removeCmd = &cobra.Command{
	Use:   "remove <plugin>",
	Short: "卸载插件",
	Long: `卸载指定的插件。

示例：
  ops remove k8s
  ops remove database`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pluginName := args[0]

		// 创建插件管理器
		pluginDir := cmd.Flag("plugin-dir").Value.String()
		pm := plugins.NewPluginManager(pluginDir)

		// 卸载插件
		fmt.Printf("正在卸载插件: %s\n", pluginName)
		err := pm.RemovePlugin(pluginName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "卸载插件失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ 插件卸载成功: %s\n", pluginName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// 添加标志
	removeCmd.Flags().String("plugin-dir", "", "插件目录")
}