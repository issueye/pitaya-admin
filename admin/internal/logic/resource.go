package logic

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/global"
	"github.com/issueye/pitaya_admin/internal/repository"
	"github.com/issueye/pitaya_admin/internal/service"
	"github.com/issueye/pitaya_admin/pkg/utils"
)

var Rsse = make(map[string]*ResourceSSE)

type ResourceSSE struct {
	Id string `json:"id"` // 编码
}

type Resource struct{}

func (Resource) Get(req *repository.QueryResource) ([]*model.ResourceInfo, error) {
	return service.NewResource().Query(req)
}

func (Resource) GetById(id string) (*model.ResourceInfo, error) {
	return service.NewResource().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Resource) Modify(req *repository.ModifyResource) error {
	ResourceService := service.NewResource()
	return ResourceService.Modify(req)
}

// Modify
// 修改信息 不包含状态
func (Resource) ModifyByMap(id string, datas map[string]any) error {
	ResourceService := service.NewResource()
	return ResourceService.ModifyByMap(id, datas)
}

// Create
// 创建数据
func (Resource) Create(req *repository.CreateResource) error {
	ResourceService := service.NewResource()

	// 创建数据
	err := ResourceService.Create(req)
	if err != nil {
		return fmt.Errorf("创建资源失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Resource) Del(id string) error {
	ResourceService := service.NewResource()
	return ResourceService.Del(id)
}

// Upload
// 上传文件
func (Resource) Upload(data *repository.UploadData) (map[string]string, error) {

	ext := filepath.Ext(data.UploadKey.Filename)
	filename := strings.TrimSuffix(path.Base(data.UploadKey.Filename), ext)
	// 生成一个sha字符串
	filename = utils.Sha256_2Str(fmt.Sprintf("%s-%s", filename, utils.GetUUID()))

	Rsse[filename] = &ResourceSSE{Id: filename}

	return map[string]string{
		"name": filename,
		"ext":  ext,
	}, nil
}

type Progress struct {
	Code     int    `json:"code"`     // 状态码
	Message  string `json:"message"`  // 消息
	Progress int    `json:"progress"` // 进度
}

func getProgress(code int, progress int, msg string) string {
	data := &Progress{
		Code:     code,
		Progress: progress,
		Message:  msg,
	}

	return utils.Struct2Json(data)
}

type Send struct {
	event string // 编码
}

func NewSend(event string) *Send {
	return &Send{
		event: event,
	}
}

func (s *Send) Info(progress int, msg string) {
	global.SSE.SendEventMessage(getProgress(1, progress, msg), s.event, "")
	time.Sleep(1 * time.Second)
}

func (s *Send) Fail(progress int, msg string) {
	global.SSE.SendEventMessage(getProgress(0, progress, msg), s.event, "")
}

func (s *Send) Failf(progress int, formatStr string, args ...any) {
	global.SSE.SendEventMessage(getProgress(0, progress, fmt.Sprintf(formatStr, args...)), s.event, "")
}

func (Resource) runGzip(sourceDir string, destDir string) error {
	// 遍历源文件夹并压缩文件
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 获取文件的相对路径
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}

			// 创建目标文件路径
			destPath := filepath.Join(destDir, relPath+".gz")

			// 创建目标文件夹
			destDir := filepath.Dir(destPath)
			err = os.MkdirAll(destDir, 0755)
			if err != nil {
				return err
			}

			// 读取文件内容
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// 创建 Gzip writer
			var b bytes.Buffer
			w := gzip.NewWriter(&b)
			_, err = w.Write(data)
			if err != nil {
				return err
			}
			err = w.Close()
			if err != nil {
				return err
			}

			// 保存压缩后的文件
			err = os.WriteFile(destPath, b.Bytes(), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
