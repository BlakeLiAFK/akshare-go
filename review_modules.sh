#!/bin/bash

# 代码审查自动化脚本
# 对比 Python 和 Go 实现，生成完整报告

OUTPUT_FILE="code_review_report.md"

echo "# akshare-go 代码审查报告" > $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "> 生成时间: $(date '+%Y-%m-%d %H:%M:%S')" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "## 审查范围" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "逐个模块对比 Python 源代码和 Go 实现的完整性：" >> $OUTPUT_FILE
echo "- 函数数量对比" >> $OUTPUT_FILE
echo "- 测试覆盖率" >> $OUTPUT_FILE
echo "- 缺失函数清单" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "---" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE

# 要审查的模块列表（排除空模块）
MODULES=(
  "air" "article" "bank" "bond" "cal" "crypto" "currency"
  "economic" "energy" "event" "forex" "fortune" "fund"
  "futures" "futures_derivative" "fx" "hf" "index" "interest_rate"
  "movie" "news" "nlp" "option" "other" "pro" "qdii"
  "qhkc" "qhkc_web" "rate" "reits" "spot"
  "stock" "stock_a" "stock_feature" "stock_fundamental" "tool"
)

total_modules=0
complete_modules=0
incomplete_modules=0

for module in "${MODULES[@]}"; do
  total_modules=$((total_modules + 1))

  echo "## 模块: $module" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE

  # 检查模块目录是否存在
  if [ ! -d "$module" ]; then
    echo "⚠️ **状态: Go 模块目录不存在**" >> $OUTPUT_FILE
    echo "" >> $OUTPUT_FILE
    incomplete_modules=$((incomplete_modules + 1))
    continue
  fi

  if [ ! -d "_akshare_source/akshare/$module" ]; then
    echo "⚠️ **状态: Python 模块目录不存在**" >> $OUTPUT_FILE
    echo "" >> $OUTPUT_FILE
    continue
  fi

  # 统计 Python 函数数量（排除 _ 开头的内部函数和 __init__.py）
  py_func_count=$(find "_akshare_source/akshare/$module" -name "*.py" ! -name "__init__.py" ! -name "cons.py" -exec grep -h "^def [a-z]" {} \; 2>/dev/null | grep -v "^def _" | wc -l | tr -d ' ')

  # 统计 Go 函数数量（只统计公开函数，以大写字母开头）
  go_func_count=$(find "$module" -name "*.go" ! -name "*_test.go" -exec grep -h "^func [A-Z]" {} \; 2>/dev/null | wc -l | tr -d ' ')

  # 统计测试函数数量
  test_count=$(find "$module" -name "*_test.go" -exec grep -h "^func Test" {} \; 2>/dev/null | wc -l | tr -d ' ')

  echo "### 统计" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE
  echo "| 项目 | 数量 |" >> $OUTPUT_FILE
  echo "|------|------|" >> $OUTPUT_FILE
  echo "| Python 函数 | $py_func_count |" >> $OUTPUT_FILE
  echo "| Go 函数 | $go_func_count |" >> $OUTPUT_FILE
  echo "| 测试函数 | $test_count |" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE

  # 判断状态
  if [ "$go_func_count" -ge "$py_func_count" ] && [ "$test_count" -ge "$go_func_count" ]; then
    echo "✅ **状态: 完整**（函数: $go_func_count/$py_func_count, 测试: $test_count/$go_func_count）" >> $OUTPUT_FILE
    complete_modules=$((complete_modules + 1))
  elif [ "$go_func_count" -ge "$py_func_count" ] && [ "$test_count" -lt "$go_func_count" ]; then
    echo "⚠️ **状态: 测试不足**（函数: $go_func_count/$py_func_count, 测试: $test_count/$go_func_count）" >> $OUTPUT_FILE
    incomplete_modules=$((incomplete_modules + 1))
  else
    echo "❌ **状态: 函数缺失**（函数: $go_func_count/$py_func_count, 测试: $test_count/$go_func_count）" >> $OUTPUT_FILE
    incomplete_modules=$((incomplete_modules + 1))

    # 列出 Python 函数清单
    echo "" >> $OUTPUT_FILE
    echo "### Python 函数清单" >> $OUTPUT_FILE
    echo "\`\`\`" >> $OUTPUT_FILE
    find "_akshare_source/akshare/$module" -name "*.py" ! -name "__init__.py" ! -name "cons.py" -exec grep -h "^def [a-z]" {} \; 2>/dev/null | grep -v "^def _" | sed 's/def \([^(]*\).*/\1/' >> $OUTPUT_FILE
    echo "\`\`\`" >> $OUTPUT_FILE

    # 列出 Go 函数清单
    echo "" >> $OUTPUT_FILE
    echo "### Go 函数清单" >> $OUTPUT_FILE
    echo "\`\`\`" >> $OUTPUT_FILE
    find "$module" -name "*.go" ! -name "*_test.go" -exec grep -h "^func [A-Z]" {} \; 2>/dev/null | sed 's/func \([^(]*\).*/\1/' >> $OUTPUT_FILE
    echo "\`\`\`" >> $OUTPUT_FILE
  fi

  echo "" >> $OUTPUT_FILE
  echo "---" >> $OUTPUT_FILE
  echo "" >> $OUTPUT_FILE
done

# 生成汇总统计
echo "## 汇总统计" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "| 项目 | 数量 | 百分比 |" >> $OUTPUT_FILE
echo "|------|------|--------|" >> $OUTPUT_FILE
echo "| 总模块数 | $total_modules | 100% |" >> $OUTPUT_FILE
echo "| 完整模块 | $complete_modules | $((complete_modules * 100 / total_modules))% |" >> $OUTPUT_FILE
echo "| 不完整模块 | $incomplete_modules | $((incomplete_modules * 100 / total_modules))% |" >> $OUTPUT_FILE
echo "" >> $OUTPUT_FILE

echo "审查报告已生成: $OUTPUT_FILE"
