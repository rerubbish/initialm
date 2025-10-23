package template

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"mao/assets"
	"mao/cmd/rules"
	"path"
	"strings"
)

type DefaultTemplate struct {
	name string
	data any
	rule rules.RuleData
}

func NewTempate(name string, data any, rule rules.RuleData) DefaultTemplate {
	return DefaultTemplate{
		name: name,
		data: data,
		rule: rule,
	}
}

func (t DefaultTemplate) applyVariableRule(fileData []byte, relPath string) ([]byte, error) {
	for _, variable := range t.rule.Variable {
		if strings.Contains(relPath, variable.FilePath) {
			tpl, err := template.New(variable.FilePath).Parse(string(fileData))
			if err != nil {
				return fileData, err
			}
			var result strings.Builder
			if err := tpl.Execute(&result, t.data); err != nil {
				return fileData, err
			}
			return []byte(result.String()), nil
		}
	}
	return fileData, nil
}

func (t DefaultTemplate) applyPathRule(relPath string) string {
	for _, path := range t.rule.Path {
		if relPath == path.Source {
			return path.Target
		}
	}
	return relPath
}

func (t DefaultTemplate) GenZip() ([]byte, error) {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	defer zw.Close()
	err := fs.WalkDir(assets.StaticFs, fmt.Sprintf("mao/assets/template/%s", t.name), func(filePath string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relPath := strings.TrimPrefix(filePath, "mao/assets/template/")
		if d.IsDir() {
			// 跳过根目录的创建（relPath为空时不创建）
			if relPath != "" {
				// 创建目录条目（zip目录需要以/结尾）
				_, err := zw.Create(path.Clean(t.applyPathRule(relPath) + "/"))
				return err
			}
			return nil
		}

		fileData, err := assets.StaticFs.ReadFile(filePath)
		if err != nil {
			return err
		}
		fileData, err = t.applyVariableRule(fileData, relPath)
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}

		// 创建文件头
		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		fh.Name = relPath       // 设置文件在zip中的路径
		fh.Method = zip.Deflate // 使用默认压缩算法
		// 写入文件内容
		fileWriter, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}
		_, err = fileWriter.Write(fileData)
		return err
	})

	if err != nil {
		return nil, err
	}

	if err := zw.Flush(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t DefaultTemplate) B64Zip() (string, error) {
	zipByte, err := t.GenZip()
	if err != nil {
		return "", err
	}
	b64Data := base64.StdEncoding.EncodeToString(zipByte)
	return b64Data, nil
}
