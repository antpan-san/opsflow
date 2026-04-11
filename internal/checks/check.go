package checks

import "github.com/yourusername/opsflow/internal/types"

// Check 定义检测接口
type Check interface {
    Name() string
    Run(input types.Input) types.Result
}

// BaseCheck 基础检测结构体
type BaseCheck struct {
    Name string
}

func (b *BaseCheck) GetName() string {
    return b.Name
}