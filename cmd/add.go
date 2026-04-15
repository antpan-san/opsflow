package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/opsflow/internal/plugins"
)

var addCmd = &cobra.Command{
	Use:   "add <plugin>",
	Short: "安装插件",
	Long: `安装指定的插件。

可用插件：
  k8s      - Kubernetes 诊断插件
  database - 数据库诊断插件
  ci-cd    - CI/CD 诊断插件

示例：
  ops add k8s
  ops add database`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pluginName := args[0]

		// 创建插件管理器
		pluginDir := cmd.Flag("plugin-dir").Value.String()
		pm := plugins.NewPluginManager(pluginDir)

		// 安装插件
		fmt.Printf("正在安装插件: %s\n", pluginName)
		err := pm.InstallPlugin(pluginName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "安装插件失败: %v\n", err)
			os.Exit(1)
		}

		// 加载插件
		err = pm.LoadPlugin(pluginName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "加载插件失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ 插件安装成功: %s\n", pluginName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// 添加标志
	addCmd.Flags().String("plugin-dir", "", "插件目录")
}