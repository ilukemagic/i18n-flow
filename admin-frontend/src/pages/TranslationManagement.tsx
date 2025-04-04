import React, { useState, useEffect } from "react";
import { Empty, Form, message, Spin } from "antd";
import { useParams } from "react-router-dom";

// Custom hooks and components
import { useTranslationData } from "../hooks/useTranslationData";
import TranslationTable, {
  generateTableColumns,
} from "../components/translation/TranslationTable";
import TranslationToolbar from "../components/translation/TranslationToolbar";
import {
  CreateTranslationModal,
  BatchTranslationModal,
  EditTranslationModal,
  ImportTranslationModal,
  ExcelImportModal,
} from "../components/translation/TranslationModal";
import {
  parseExcelFile,
  formatExcelDataForImport,
  autoMapLanguageColumns,
  ExcelData,
} from "../components/translation/ExcelUtils";

// Types
import {
  TranslationResponse,
  BatchTranslationRequest,
} from "../types/translation";

const TranslationManagement: React.FC = () => {
  const { projectId } = useParams<{ projectId: string }>();

  // Use the custom hook to handle translation data
  const {
    translations,
    projects,
    languages,
    columns,
    selectedProject,
    loading,
    keyword,
    paginatedMatrix,
    localPagination,
    selectedKeys,
    selectedTranslations,
    batchDeleteLoading,

    setColumns,
    setSelectedProject,
    setKeyword,
    handleTableChange,
    fetchTranslations,
    createTranslation,
    batchCreateTranslations,
    updateTranslation,
    deleteTranslation,
    exportTranslations,
    importTranslationsFromJson,
    batchDeleteTranslations,
    handleRowSelectionChange,
  } = useTranslationData(projectId);

  // Modal state
  const [createModalVisible, setCreateModalVisible] = useState<boolean>(false);
  const [batchModalVisible, setBatchModalVisible] = useState<boolean>(false);
  const [editModalVisible, setEditModalVisible] = useState<boolean>(false);
  const [importModalVisible, setImportModalVisible] = useState<boolean>(false);
  const [currentTranslation, setCurrentTranslation] =
    useState<TranslationResponse | null>(null);

  // Excel import state
  const [excelImportModalVisible, setExcelImportModalVisible] =
    useState<boolean>(false);
  const [excelData, setExcelData] = useState<any[]>([]);
  const [excelPreviewColumns, setExcelPreviewColumns] = useState<any[]>([]);
  const [excelPreviewData, setExcelPreviewData] = useState<any[]>([]);
  const [selectedLanguages, setSelectedLanguages] = useState<{
    [key: string]: string;
  }>({});
  const [excelImportLoading, setExcelImportLoading] = useState<boolean>(false);

  // Forms
  const [singleForm] = Form.useForm();
  const [batchForm] = Form.useForm();
  const [editForm] = Form.useForm();

  // Effect to generate table columns when languages are loaded
  useEffect(() => {
    if (languages.length > 0) {
      const generatedColumns = generateTableColumns(
        languages,
        translations,
        showEditModal,
        handleDeleteTranslation,
        handleAddTranslation,
        showBatchAddModal
      );
      setColumns(generatedColumns);
    }
  }, [languages, translations]);

  // Handlers for translation operations
  const handleAddTranslation = (keyName: string, languageId: number) => {
    if (!selectedProject) {
      message.warning("请先选择项目");
      return;
    }

    // Get context info for the key name
    const context = paginatedMatrix.find(
      (m) => m.key_name === keyName
    )?.context;

    singleForm.setFieldsValue({
      project_id: selectedProject,
      key_name: keyName,
      context: context,
      language_id: languageId,
    });
    setCreateModalVisible(true);
  };

  const handleCreateTranslation = async () => {
    try {
      const values = await singleForm.validateFields();
      const success = await createTranslation(values);
      if (success) {
        setCreateModalVisible(false);
        singleForm.resetFields();
      }
    } catch (error) {
      console.error("创建翻译失败:", error);
    }
  };

  const showBatchAddModal = async (keyName: string, context?: string) => {
    if (!selectedProject) {
      message.warning("请先选择项目");
      return;
    }

    try {
      // Clear previous form data
      batchForm.resetFields();

      // Set basic info
      batchForm.setFieldsValue({
        project_id: selectedProject,
        key_name: keyName,
        context: context,
      });

      // Find all existing translations for the key name
      const existingTranslations = translations.filter(
        (t) => t.key_name === keyName
      );

      // Set values for each language
      languages.forEach((lang) => {
        const translation = existingTranslations.find(
          (t) => t.language_code === lang.code
        );
        if (translation) {
          batchForm.setFieldsValue({
            [`lang_${lang.code}`]: translation.value,
          });
        }
      });

      // Make sure form values are set before showing modal
      setTimeout(() => {
        setBatchModalVisible(true);
      }, 100);
    } catch (error) {
      console.error("加载现有翻译失败:", error);
      message.error("加载现有翻译失败");
    }
  };

  const handleBatchCreateTranslations = async () => {
    try {
      const values = await batchForm.validateFields();

      // Build request data
      const request: BatchTranslationRequest = {
        project_id: selectedProject!,
        key_name: values.key_name,
        context: values.context,
        translations: {},
      };

      // Process language fields in the form
      Object.keys(values).forEach((key) => {
        if (key.startsWith("lang_") && values[key]) {
          const langCode = key.replace("lang_", "");
          request.translations[langCode] = values[key];
        }
      });

      // Only send request if there are values
      if (Object.keys(request.translations).length > 0) {
        const success = await batchCreateTranslations(request);
        if (success) {
          setBatchModalVisible(false);
          batchForm.resetFields();
        }
      } else {
        message.warning("请至少添加一种语言的翻译");
      }
    } catch (error) {
      console.error("批量创建翻译失败:", error);
    }
  };

  const showEditModal = (translation: TranslationResponse) => {
    setCurrentTranslation(translation);
    editForm.setFieldsValue({
      project_id: translation.project_id,
      key_name: translation.key_name,
      context: translation.context,
      language_id: translation.language_id,
      value: translation.value,
    });
    setEditModalVisible(true);
  };

  const handleEditTranslation = async () => {
    if (!currentTranslation) return;

    try {
      const values = await editForm.validateFields();
      const success = await updateTranslation(currentTranslation.id, values);
      if (success) {
        setEditModalVisible(false);
      }
    } catch (error) {
      console.error("更新翻译失败:", error);
    }
  };

  const handleDeleteTranslation = async (id: number) => {
    await deleteTranslation(id);
  };

  // Export translations
  const handleExportTranslations = async (format: string = "json") => {
    const data = await exportTranslations(format);
    if (data) {
      // Create and download file
      const fileName = `translations_${selectedProject}_${new Date().toISOString()}.json`;
      const jsonStr = JSON.stringify(data, null, 2);
      const blob = new Blob([jsonStr], { type: "application/json" });
      const href = URL.createObjectURL(blob);

      const link = document.createElement("a");
      link.href = href;
      link.download = fileName;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      URL.revokeObjectURL(href);
    }
  };

  // Import translations from JSON
  const handleImportTranslations = (file: File) => {
    if (!selectedProject) {
      message.warning("请先选择项目");
      return false;
    }

    try {
      const reader = new FileReader();
      reader.onload = async (e) => {
        try {
          const content = e.target?.result as string;
          const data = JSON.parse(content);

          const success = await importTranslationsFromJson(data);
          if (success) {
            setImportModalVisible(false);
          }
        } catch (error) {
          console.error("解析文件失败:", error);
          message.error("导入失败：文件格式错误");
        }
      };
      reader.readAsText(file);
      return false; // Prevent default upload behavior
    } catch (error) {
      console.error("导入翻译失败:", error);
      message.error("导入翻译失败");
      return false;
    }
  };

  // Excel processing functions
  const handleExcelFile = (file: File) => {
    parseExcelFile(
      file,
      languages,
      (data: ExcelData | null) => {
        if (data) {
          setExcelData(data.jsonData);
          setExcelPreviewColumns(data.columns);
          setExcelPreviewData(data.previewData);

          // Auto-map columns to languages
          const initialMapping = autoMapLanguageColumns(
            data.columns,
            languages
          );
          setSelectedLanguages(initialMapping);

          setExcelImportModalVisible(true);
        }
      },
      setExcelImportLoading
    );
    return false; // Prevent default upload behavior
  };

  const handleLanguageSelect = (columnKey: string, languageCode: string) => {
    setSelectedLanguages((prev) => ({
      ...prev,
      [columnKey]: languageCode,
    }));
  };

  const handleExcelImport = async () => {
    if (!selectedProject) {
      message.warning("请先选择项目");
      return;
    }

    if (Object.keys(selectedLanguages).length === 0) {
      message.warning("请至少选择一种语言映射");
      return;
    }

    try {
      setExcelImportLoading(true);

      // Format Excel data for import
      const importData = formatExcelDataForImport(excelData, selectedLanguages);

      // Call existing import API
      const success = await importTranslationsFromJson(importData);
      if (success) {
        setExcelImportModalVisible(false);
      }
    } catch (error) {
      console.error("导入Excel翻译失败:", error);
      message.error(
        "导入Excel翻译失败: " +
          (error instanceof Error ? error.message : "未知错误")
      );
    } finally {
      setExcelImportLoading(false);
    }
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow">
      <TranslationToolbar
        projects={projects}
        selectedProject={selectedProject}
        keyword={keyword}
        onProjectChange={setSelectedProject}
        onKeywordChange={setKeyword}
        onSearch={fetchTranslations}
        onAddTranslation={() => {
          if (!selectedProject) {
            message.warning("请先选择项目");
            return;
          }
          singleForm.setFieldsValue({ project_id: selectedProject });
          setCreateModalVisible(true);
        }}
        onBatchAddTranslation={() => {
          if (!selectedProject) {
            message.warning("请先选择项目");
            return;
          }
          batchForm.setFieldsValue({ project_id: selectedProject });
          setBatchModalVisible(true);
        }}
        onImportJsonClick={() => setImportModalVisible(true)}
        onExportClick={() => handleExportTranslations()}
        onExcelFileUpload={handleExcelFile}
        selectedTranslations={selectedTranslations}
        onBatchDelete={batchDeleteTranslations}
        batchDeleteLoading={batchDeleteLoading}
      />

      {loading ? (
        <div className="flex justify-center items-center py-12">
          <Spin size="large" tip="加载中..." />
        </div>
      ) : paginatedMatrix.length > 0 ? (
        <TranslationTable
          loading={loading}
          paginatedMatrix={paginatedMatrix}
          columns={columns}
          translations={translations}
          languages={languages}
          selectedKeys={selectedKeys}
          pagination={localPagination}
          onTableChange={handleTableChange}
          onRowSelectionChange={handleRowSelectionChange}
          onEditTranslation={showEditModal}
          onDeleteTranslation={handleDeleteTranslation}
          onAddTranslation={handleAddTranslation}
          onShowBatchAddModal={showBatchAddModal}
        />
      ) : (
        <Empty description="暂无翻译数据" />
      )}

      {/* Modals */}
      <CreateTranslationModal
        visible={createModalVisible}
        onCancel={() => setCreateModalVisible(false)}
        onOk={handleCreateTranslation}
        form={singleForm}
        projects={projects}
        languages={languages}
        selectedProject={selectedProject}
      />

      <BatchTranslationModal
        visible={batchModalVisible}
        onCancel={() => setBatchModalVisible(false)}
        onOk={handleBatchCreateTranslations}
        form={batchForm}
        projects={projects}
        languages={languages}
        selectedProject={selectedProject}
        translations={translations}
      />

      <EditTranslationModal
        visible={editModalVisible}
        onCancel={() => setEditModalVisible(false)}
        onOk={handleEditTranslation}
        form={editForm}
        projects={projects}
        languages={languages}
        selectedProject={selectedProject}
      />

      <ImportTranslationModal
        visible={importModalVisible}
        onCancel={() => setImportModalVisible(false)}
        onImport={handleImportTranslations}
      />

      <ExcelImportModal
        visible={excelImportModalVisible}
        onCancel={() => setExcelImportModalVisible(false)}
        onOk={handleExcelImport}
        excelPreviewColumns={excelPreviewColumns}
        excelPreviewData={excelPreviewData}
        selectedLanguages={selectedLanguages}
        onLanguageSelect={handleLanguageSelect}
        languages={languages}
        loading={excelImportLoading}
      />
    </div>
  );
};

export default TranslationManagement;
