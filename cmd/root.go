package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ops",
    Short: "OpsFlow - 运维诊断与自动化执行引擎",
    Long: `OpsFlow 是一个场景驱动的运维诊断与自动化执行引擎。

核心理念：把运维经验变成一条可执行命令。

使用示例：
  ops run web example.com
  ops list
  ops generate nginx --domain example.com`,
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.PersistentFlags().StringP("output", "o", "text", "输出格式 (text|json)")
}