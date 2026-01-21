package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// TokenFileName Token 文件名
	TokenFileName = ".akshare_token"
	// TokenErrMsg Token 错误提示信息
	TokenErrMsg = "请先设置 Token，使用 SetToken 函数设置"
)

// getTokenFilePath 获取 Token 文件路径
func getTokenFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("获取用户目录失败: %w", err)
	}
	return filepath.Join(homeDir, TokenFileName), nil
}

// SetToken 设置 Token
// token: API Token 字符串
func SetToken(token string) error {
	fp, err := getTokenFilePath()
	if err != nil {
		return err
	}

	// 创建或覆盖文件
	file, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("创建 Token 文件失败: %w", err)
	}
	defer file.Close()

	// 写入 token
	_, err = file.WriteString(token)
	if err != nil {
		return fmt.Errorf("写入 Token 失败: %w", err)
	}

	return nil
}

// GetToken 获取 Token
// 返回存储的 Token 字符串，如果文件不存在返回空字符串和错误
func GetToken() (string, error) {
	fp, err := getTokenFilePath()
	if err != nil {
		return "", err
	}

	// 检查文件是否存在
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return "", fmt.Errorf(TokenErrMsg)
	}

	// 读取文件
	file, err := os.Open(fp)
	if err != nil {
		return "", fmt.Errorf("打开 Token 文件失败: %w", err)
	}
	defer file.Close()

	// 读取第一行
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("读取 Token 失败: %w", err)
	}

	return "", fmt.Errorf(TokenErrMsg)
}

// MustGetToken 获取 Token，失败返回空字符串
func MustGetToken() string {
	token, _ := GetToken()
	return token
}

// HasToken 检查是否已设置 Token
func HasToken() bool {
	fp, err := getTokenFilePath()
	if err != nil {
		return false
	}
	_, err = os.Stat(fp)
	return err == nil
}

// DeleteToken 删除 Token 文件
func DeleteToken() error {
	fp, err := getTokenFilePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return nil // 文件不存在，无需删除
	}

	return os.Remove(fp)
}
