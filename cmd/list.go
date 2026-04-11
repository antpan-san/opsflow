package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "列出所有可用场景",
    Long:  `列出所有已注册的诊断场景。`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("=== 可用场景 ===")
        fmt.Println("web    - Web服务诊断 (DNS + TCP + HTTP)")
        fmt.Println("k8s    - Kubernetes诊断 (开发中)")
        fmt.Println("database - 数据库诊断 (开发中)")
        
        fmt.Println("\n使用示例:")
        fmt.Println("  ops run web example.com")
        fmt.Println("  ops run web example.com --port 443 --protocol https")
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}