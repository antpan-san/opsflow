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

var runCmd = &cobra.Command{
    Use:   "run <scenario> <target>",
    Short: "运行诊断场景",
    Long: `运行指定的诊断场景。

可用场景：
  web    - Web服务诊断
  k8s    - Kubernetes诊断（开发中）

示例：
  ops run web example.com
  ops run web example.com --port 443 --protocol https`,
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        startTime := time.Now()
        scenarioName := args[0]
        target := args[1]

        // 获取输出格式
        output, _ := cmd.Flags().GetString("output")
        colorOutput, _ := cmd.Flags().GetBool("color")

        // 创建引擎
        eng := engine.NewEngine()

        // 注册场景
        scenarios.RegisterWebScenario(eng)
        scenarios.RegisterWebDetailedScenario(eng)

        // 准备输入参数
        input := types.Input{
            Target: target,
            Params: make(map[string]string),
        }

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

        // 计算耗时
        duration := time.Since(startTime)
        ctx.Duration = duration

        // 输出结果
        var result string
        if output == "json" {
            result = eng.FormatOutputJSON(ctx)
        } else if colorOutput {
            result = eng.FormatOutputColor(ctx)
        } else {
            result = eng.FormatOutput(ctx, output)
        }
        fmt.Println(result)

        // 如果诊断发现异常，退出码为1
        if ruleResult, exists := ctx.Results["rule"]; exists {
            if conclusion, ok := ruleResult.Data["conclusion"].(string); ok {
                if conclusion != "Web服务正常" {
                    os.Exit(1)
                }
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(runCmd)

    // 添加标志
    runCmd.Flags().StringP("port", "p", "", "指定端口（默认80）")
    runCmd.Flags().StringP("protocol", "P", "", "指定协议（http/https）")
    runCmd.Flags().StringP("output", "o", "text", "输出格式 (text|json)")
    runCmd.Flags().BoolP("color", "c", false, "启用彩色输出")
}