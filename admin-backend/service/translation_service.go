package service

import (
	"errors"
	"i18n-flow/model"
	"i18n-flow/model/db"
	"strings"

	"gorm.io/gorm"
)

// TranslationService 翻译服务
type TranslationService struct{}

// TranslationRequest 翻译请求参数
type TranslationRequest struct {
	ProjectID  uint   `json:"project_id" binding:"required" example:"1"`             // 项目ID
	KeyName    string `json:"key_name" binding:"required" example:"welcome_message"` // 翻译键名
	Context    string `json:"context" example:"主页欢迎消息"`                              // 上下文说明
	LanguageID uint   `json:"language_id" binding:"required" example:"2"`            // 语言ID
	Value      string `json:"value" example:"欢迎使用"`                                  // 翻译值
}

// BatchTranslationRequest 批量翻译请求
type BatchTranslationRequest struct {
	ProjectID    uint              `json:"project_id" binding:"required" example:"1"`             // 项目ID
	KeyName      string            `json:"key_name" binding:"required" example:"welcome_message"` // 翻译键名
	Context      string            `json:"context" example:"欢迎消息"`                                // 上下文说明
	Translations map[string]string `json:"translations" example:"{'zh-CN':'欢迎','en':'Welcome'}"`  // 语言代码 -> 翻译值
}

// TranslationResponse 翻译响应
type TranslationResponse struct {
	ID           uint   `json:"id" example:"1"`                           // 翻译ID
	ProjectID    uint   `json:"project_id" example:"1"`                   // 项目ID
	KeyName      string `json:"key_name" example:"welcome_message"`       // 翻译键名
	Context      string `json:"context" example:"欢迎消息"`                   // 上下文说明
	LanguageID   uint   `json:"language_id" example:"2"`                  // 语言ID
	Value        string `json:"value" example:"欢迎使用"`                     // 翻译值
	Status       string `json:"status" example:"active"`                  // 状态
	CreatedAt    string `json:"created_at" example:"2023-04-01 12:00:00"` // 创建时间
	UpdatedAt    string `json:"updated_at" example:"2023-04-01 12:00:00"` // 更新时间
	ProjectName  string `json:"project_name" example:"网站翻译"`              // 项目名称
	LanguageCode string `json:"language_code" example:"zh-CN"`            // 语言代码
	LanguageName string `json:"language_name" example:"简体中文"`             // 语言名称
}

// CreateTranslation 创建翻译
func (s *TranslationService) CreateTranslation(req TranslationRequest) (*TranslationResponse, error) {
	// 检查项目是否存在
	var project model.Project
	if err := db.DB.First(&project, req.ProjectID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("项目不存在")
		}
		return nil, err
	}

	// 检查语言是否存在
	var language model.Language
	if err := db.DB.First(&language, req.LanguageID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("语言不存在")
		}
		return nil, err
	}

	// 检查是否已存在相同的键名和语言
	var existingTranslation model.Translation
	if err := db.DB.Where("project_id = ? AND key_name = ? AND language_id = ?",
		req.ProjectID, req.KeyName, req.LanguageID).First(&existingTranslation).Error; err == nil {
		return nil, errors.New("该项目下已存在相同键名和语言的翻译")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	translation := model.Translation{
		ProjectID:  req.ProjectID,
		KeyName:    req.KeyName,
		Context:    req.Context,
		LanguageID: req.LanguageID,
		Value:      req.Value,
		Status:     "active",
	}

	if err := db.DB.Create(&translation).Error; err != nil {
		return nil, err
	}

	return &TranslationResponse{
		ID:           translation.ID,
		ProjectID:    translation.ProjectID,
		KeyName:      translation.KeyName,
		Context:      translation.Context,
		LanguageID:   translation.LanguageID,
		Value:        translation.Value,
		Status:       translation.Status,
		CreatedAt:    translation.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    translation.UpdatedAt.Format("2006-01-02 15:04:05"),
		ProjectName:  project.Name,
		LanguageCode: language.Code,
		LanguageName: language.Name,
	}, nil
}

// BatchCreateTranslations 批量创建翻译
func (s *TranslationService) BatchCreateTranslations(req BatchTranslationRequest) ([]TranslationResponse, error) {
	// 检查项目是否存在
	var project model.Project
	if err := db.DB.First(&project, req.ProjectID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("项目不存在")
		}
		return nil, err
	}

	// 获取所有语言的映射
	var languages []model.Language
	if err := db.DB.Find(&languages).Error; err != nil {
		return nil, err
	}

	langMap := make(map[string]model.Language)
	for _, lang := range languages {
		langMap[lang.Code] = lang
	}

	// 批量创建翻译
	var translations []model.Translation
	var responses []TranslationResponse

	cleanedTranslations := make(map[string]string)
	for code, value := range req.Translations {
		// 清理语言代码
		cleanCode := strings.Trim(code, "\"{}\\")
		// 清理翻译值
		cleanValue := strings.Trim(value, "\"{}\\")
		cleanedTranslations[cleanCode] = cleanValue
	}

	for langCode, value := range cleanedTranslations {
		lang, exists := langMap[langCode]
		if !exists {
			return nil, errors.New("语言代码不存在: " + langCode)
		}

		// 检查是否已存在相同的键名和语言
		var existingTranslation model.Translation
		if err := db.DB.Where("project_id = ? AND key_name = ? AND language_id = ?",
			req.ProjectID, req.KeyName, lang.ID).First(&existingTranslation).Error; err == nil {
			// 如果存在，更新值
			existingTranslation.Value = value
			if err := db.DB.Save(&existingTranslation).Error; err != nil {
				return nil, err
			}

			responses = append(responses, TranslationResponse{
				ID:           existingTranslation.ID,
				ProjectID:    existingTranslation.ProjectID,
				KeyName:      existingTranslation.KeyName,
				Context:      existingTranslation.Context,
				LanguageID:   existingTranslation.LanguageID,
				Value:        existingTranslation.Value,
				Status:       existingTranslation.Status,
				CreatedAt:    existingTranslation.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    existingTranslation.UpdatedAt.Format("2006-01-02 15:04:05"),
				ProjectName:  project.Name,
				LanguageCode: lang.Code,
				LanguageName: lang.Name,
			})
		} else if err == gorm.ErrRecordNotFound {
			// 如果不存在，创建新的
			translation := model.Translation{
				ProjectID:  req.ProjectID,
				KeyName:    req.KeyName,
				Context:    req.Context,
				LanguageID: lang.ID,
				Value:      value,
				Status:     "active",
			}

			translations = append(translations, translation)
		} else {
			return nil, err
		}
	}

	// 批量插入新的翻译
	if len(translations) > 0 {
		if err := db.DB.CreateInBatches(translations, len(translations)).Error; err != nil {
			return nil, err
		}

		// 构建响应
		for _, t := range translations {
			lang := langMap[getLanguageCodeByID(languages, t.LanguageID)]
			responses = append(responses, TranslationResponse{
				ID:           t.ID,
				ProjectID:    t.ProjectID,
				KeyName:      t.KeyName,
				Context:      t.Context,
				LanguageID:   t.LanguageID,
				Value:        t.Value,
				Status:       t.Status,
				CreatedAt:    t.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    t.UpdatedAt.Format("2006-01-02 15:04:05"),
				ProjectName:  project.Name,
				LanguageCode: lang.Code,
				LanguageName: lang.Name,
			})
		}
	}

	return responses, nil
}

// 根据语言ID获取语言代码
func getLanguageCodeByID(languages []model.Language, id uint) string {
	for _, lang := range languages {
		if lang.ID == id {
			return lang.Code
		}
	}
	return ""
}

// GetTranslationsByProject 获取项目的所有翻译
func (s *TranslationService) GetTranslationsByProject(projectID uint, page, pageSize int, keyword string) ([]TranslationResponse, int64, error) {
	var translations []model.Translation
	var total int64

	query := db.DB.Model(&model.Translation{}).Where("project_id = ?", projectID)

	// 如果有关键字，添加搜索条件
	if keyword != "" {
		query = query.Where("key_name LIKE ? OR value LIKE ? OR context LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询，预加载关联对象
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).
		Preload("Project").Preload("Language").
		Order("key_name, language_id").
		Find(&translations).Error; err != nil {
		return nil, 0, err
	}

	// 格式化响应
	var response []TranslationResponse
	for _, t := range translations {
		response = append(response, TranslationResponse{
			ID:           t.ID,
			ProjectID:    t.ProjectID,
			KeyName:      t.KeyName,
			Context:      t.Context,
			LanguageID:   t.LanguageID,
			Value:        t.Value,
			Status:       t.Status,
			CreatedAt:    t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    t.UpdatedAt.Format("2006-01-02 15:04:05"),
			ProjectName:  t.Project.Name,
			LanguageCode: t.Language.Code,
			LanguageName: t.Language.Name,
		})
	}

	return response, total, nil
}

// GetTranslationByID 根据ID获取翻译
func (s *TranslationService) GetTranslationByID(id uint) (*TranslationResponse, error) {
	var translation model.Translation

	if err := db.DB.Preload("Project").Preload("Language").First(&translation, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("翻译不存在")
		}
		return nil, err
	}

	return &TranslationResponse{
		ID:           translation.ID,
		ProjectID:    translation.ProjectID,
		KeyName:      translation.KeyName,
		Context:      translation.Context,
		LanguageID:   translation.LanguageID,
		Value:        translation.Value,
		Status:       translation.Status,
		CreatedAt:    translation.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    translation.UpdatedAt.Format("2006-01-02 15:04:05"),
		ProjectName:  translation.Project.Name,
		LanguageCode: translation.Language.Code,
		LanguageName: translation.Language.Name,
	}, nil
}

// UpdateTranslation 更新翻译
func (s *TranslationService) UpdateTranslation(id uint, req TranslationRequest) (*TranslationResponse, error) {
	var translation model.Translation

	if err := db.DB.Preload("Project").Preload("Language").First(&translation, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("翻译不存在")
		}
		return nil, err
	}

	// 更新翻译
	translation.Value = req.Value
	if req.Context != "" {
		translation.Context = req.Context
	}

	if err := db.DB.Save(&translation).Error; err != nil {
		return nil, err
	}

	return &TranslationResponse{
		ID:           translation.ID,
		ProjectID:    translation.ProjectID,
		KeyName:      translation.KeyName,
		Context:      translation.Context,
		LanguageID:   translation.LanguageID,
		Value:        translation.Value,
		Status:       translation.Status,
		CreatedAt:    translation.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    translation.UpdatedAt.Format("2006-01-02 15:04:05"),
		ProjectName:  translation.Project.Name,
		LanguageCode: translation.Language.Code,
		LanguageName: translation.Language.Name,
	}, nil
}

// DeleteTranslation 删除翻译
func (s *TranslationService) DeleteTranslation(id uint) error {
	if err := db.DB.Delete(&model.Translation{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ExportTranslations 导出项目的翻译
func (s *TranslationService) ExportTranslations(projectID uint, format string) (interface{}, error) {
	var translations []model.Translation

	if err := db.DB.Where("project_id = ?", projectID).
		Preload("Language").Find(&translations).Error; err != nil {
		return nil, err
	}

	// 按语言分组
	langMap := make(map[string]map[string]string)
	for _, t := range translations {
		langCode := t.Language.Code
		if _, exists := langMap[langCode]; !exists {
			langMap[langCode] = make(map[string]string)
		}
		langMap[langCode][t.KeyName] = t.Value
	}

	// 根据格式返回不同的结构
	switch strings.ToLower(format) {
	case "json":
		return langMap, nil
	case "flat":
		// 扁平化格式
		flatMap := make(map[string]map[string]string)
		for lang, values := range langMap {
			flatMap[lang] = values
		}
		return flatMap, nil
	default:
		return langMap, nil
	}
}

// ImportTranslations 导入项目的翻译
func (s *TranslationService) ImportTranslations(projectID uint, data map[string]map[string]string) (int, error) {
	// 检查项目是否存在
	var project model.Project
	if err := db.DB.First(&project, projectID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.New("项目不存在")
		}
		return 0, err
	}

	// 获取所有语言的映射
	var languages []model.Language
	if err := db.DB.Find(&languages).Error; err != nil {
		return 0, err
	}

	langMap := make(map[string]model.Language)
	for _, lang := range languages {
		langMap[lang.Code] = lang
	}

	// 批量创建或更新翻译
	var translations []model.Translation
	count := 0

	// 开始事务
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	for langCode, entries := range data {
		lang, exists := langMap[langCode]
		if !exists {
			// 如果语言不存在，创建新语言
			newLang := model.Language{
				Code:   langCode,
				Name:   langCode, // 使用代码作为名称的占位符
				Status: "active",
			}

			if err := tx.Create(&newLang).Error; err != nil {
				tx.Rollback()
				return 0, err
			}

			lang = newLang
			langMap[langCode] = lang
		}

		for key, value := range entries {
			// 检查是否已存在相同的键名和语言
			var existingTranslation model.Translation
			if err := tx.Where("project_id = ? AND key_name = ? AND language_id = ?",
				projectID, key, lang.ID).First(&existingTranslation).Error; err == nil {
				// 存在则更新
				existingTranslation.Value = value
				if err := tx.Save(&existingTranslation).Error; err != nil {
					tx.Rollback()
					return 0, err
				}
			} else if err == gorm.ErrRecordNotFound {
				// 不存在则创建
				translation := model.Translation{
					ProjectID:  projectID,
					KeyName:    key,
					Context:    "",
					LanguageID: lang.ID,
					Value:      value,
					Status:     "active",
				}

				translations = append(translations, translation)
			} else {
				tx.Rollback()
				return 0, err
			}

			count++
		}
	}

	// 批量插入新的翻译
	if len(translations) > 0 {
		if err := tx.CreateInBatches(translations, 100).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return count, nil
}

// BatchDeleteTranslations 批量删除翻译
func (s *TranslationService) BatchDeleteTranslations(ids []uint) (int, error) {
	if len(ids) == 0 {
		return 0, errors.New("未提供待删除的翻译ID")
	}

	// 开始事务
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	// 执行批量删除操作
	result := tx.Delete(&model.Translation{}, ids)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	// 返回删除的记录数
	return int(result.RowsAffected), nil
}
