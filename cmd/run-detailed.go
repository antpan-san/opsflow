package cmd

import (
    "fmt"
    "os"
    "time"

    "github.com/spf13/cobra"
    "github.com/yourusername/opsflow/internal/engine"
    "github.com/yourusername/opsflow/internal/scenarios"
    "github.com/yourusername/opsflow/internal/types"
)

var runDetailedCmd = &cobra.Command{
    Use:   "run-detailed <scenario> <target>",
    Short: "运行详细诊断场景",
    Long: `运行指定的详细诊断场景，提供更丰富的诊断信息。

可用场景：
  web-detailed - Web服务详细诊断（包含网络、DNS、HTTP等详细信息）

示例：
  ops run-detailed web-detailed example.com
  ops run-detailed web-detailed example.com --port 443 --protocol https`,
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        scenarioName := args[0]
        target := args[1]

        // 获取输出格式
        output, _ := cmd.Flags().GetString("output")

        // 创建引擎
        eng := engine.NewEngine()

        // 注册详细场景
        scenarios.RegisterWebDetailedScenario(eng)

        // 准备输入参数
        input := types.Input{
            Target: target,
            Params: make(map[string]string),
        }

        // 添加时间戳
        input.Params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")

        // 获取可选参数
        port, _ := cmd.Flags().GetString("port")
        if port != "" {
            input.Params["port"] = port
        }

        protocol, _ := cmd.Flags().GetString("protocol")
        if protocol != "" {
            input.Params["protocol"] = protocol
        }

        // 运行诊断
        ctx, err := eng.Run(scenarioName, input)
        if err != nil {
            fmt.Fprintf(os.Stderr, "错误: %v\n", err)
            os.Exit(1)
        }

        // 输出详细结果
        result := eng.FormatOutputDetailed(ctx)
        fmt.Println(result)

        // 如果诊断发现异常，退出码为1
        if ruleResult, exists := ctx.Results["rule"]; exists {
            if conclusion, ok := ruleResult.Data["conclusion"].(string); ok {
                if conclusion != "Web服务完全正常" {
                    os.Exit(1)
                }
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(runDetailedCmd)

    // 添加标志
    runDetailedCmd.Flags().StringP("port", "p", "", "指定端口（默认80）")
    runDetailedCmd.Flags().StringP("protocol", "P", "", "指定协议（http/https）")
}