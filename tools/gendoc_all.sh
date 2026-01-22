#!/bin/bash

# 为所有模块生成文档

# 需要排除的目录
exclude_dirs="examples plans data file_fold utils internal .git .claude docs tools coverage"

# 获取所有模块目录
for dir in */; do
    # 移除尾部斜杠
    module=$(basename "$dir")

    # 检查是否在排除列表中
    skip=0
    for exclude in $exclude_dirs; do
        if [ "$module" = "$exclude" ]; then
            skip=1
            break
        fi
    done

    if [ $skip -eq 1 ]; then
        continue
    fi

    # 检查是否包含Go文件
    if ls "$dir"/*.go 2>/dev/null | grep -v "_test.go" > /dev/null; then
        echo "生成 $module 模块文档..."
        go run tools/gendoc.go "$module"
    fi
done

echo ""
echo "文档生成完成！"
echo "生成的文档数量："
ls -1 docs/*.md | wc -l
