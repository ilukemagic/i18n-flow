# 前端使用指南

i18n Flow 管理界面是一个基于 React 开发的现代化前端应用，提供友好的用户界面来管理您的国际化内容。本指南将帮助您了解如何使用管理界面的各项功能。

## 登录系统

首次访问 i18n Flow 管理界面时，您需要登录：

1. 访问管理界面（默认为 <http://localhost> 或生产环境中的部署地址）
2. 输入您的用户名和密码
   - 默认管理员账号：`admin`
   - 默认密码：`admin123`
3. 点击"登录"按钮

::: warning 注意
首次登录后，请立即修改默认密码以确保系统安全！
:::

## 仪表板概览

登录成功后，您将看到主仪表板，它提供了系统的整体统计信息：

- 项目总数
- 语言总数
- 翻译键总数
- 翻译完成率
- 最近活动

## 项目管理

### 创建新项目

1. 在侧边栏菜单中点击"项目"
2. 点击"创建项目"按钮
3. 填写项目信息：
   - 项目名称（必填）
   - 项目描述（可选）
   - 项目标识符（slug，用于 API 和 CLI 调用）
4. 点击"保存"按钮

### 查看项目列表

项目列表页面显示所有项目及其基本信息：

- 项目名称
- 创建日期
- 语言数量
- 翻译键数量
- 完成率

### 编辑项目

1. 在项目列表中找到要编辑的项目
2. 点击"编辑"按钮
3. 修改项目信息
4. 点击"保存"按钮

### 删除项目

1. 在项目列表中找到要删除的项目
2. 点击"删除"按钮
3. 在确认对话框中点击"确认"

::: danger 警告
删除项目将永久移除该项目的所有翻译数据，此操作无法撤销！
:::

## 语言管理

### 添加语言

1. 在侧边栏菜单中点击"语言"
2. 点击"添加语言"按钮
3. 填写语言信息：
   - 语言名称（例如：English）
   - 语言代码（例如：en）
   - 语言标识符（例如：en-US）
   - RTL 支持（从右到左文字方向）
4. 点击"保存"按钮

### 编辑语言

1. 在语言列表中找到要编辑的语言
2. 点击"编辑"按钮
3. 修改语言信息
4. 点击"保存"按钮

## 翻译管理

### 翻译界面

1. 在侧边栏菜单中点击"项目"
2. 点击要管理翻译的项目名称
3. 在项目详情页面，切换到"翻译"选项卡

翻译界面采用表格形式显示翻译内容，每一行代表一个翻译键，每一列代表一种语言。

### 添加翻译键

1. 在翻译界面中点击"添加键"按钮
2. 填写以下信息：
   - 翻译键（例如：common.buttons.save）
   - 描述（可选，帮助翻译人员理解上下文）
   - 默认语言的翻译文本
3. 点击"保存"按钮

### 编辑翻译

1. 在翻译表格中找到要编辑的翻译键
2. 点击对应语言列的单元格
3. 输入翻译文本
4. 点击单元格外部或按回车键保存

### 批量导入翻译

1. 在翻译界面中点击"导入"按钮
2. 选择导入格式（JSON、Excel、CSV）
3. 上传文件或粘贴内容
4. 选择导入选项（合并或覆盖）
5. 点击"导入"按钮

### 导出翻译

1. 在翻译界面中点击"导出"按钮
2. 选择导出格式（JSON、Excel、CSV）
3. 选择要导出的语言
4. 点击"导出"按钮

## 用户管理

### 创建用户

1. 在侧边栏菜单中点击"用户"
2. 点击"创建用户"按钮
3. 填写用户信息：
   - 用户名
   - 电子邮件
   - 密码
   - 角色（管理员或翻译员）
4. 点击"保存"按钮

### 编辑用户

1. 在用户列表中找到要编辑的用户
2. 点击"编辑"按钮
3. 修改用户信息
4. 点击"保存"按钮

## 设置

### 个人资料设置

1. 点击右上角的用户头像
2. 选择"个人资料"
3. 修改您的个人信息
4. 点击"保存"按钮

### 修改密码

1. 点击右上角的用户头像
2. 选择"修改密码"
3. 输入当前密码和新密码
4. 点击"保存"按钮

## 常见问题

### 无法登录

- 确认用户名和密码是否正确
- 检查后端服务是否正常运行
- 清除浏览器缓存，然后重试

### 翻译无法保存

- 确认您有编辑该项目的权限
- 检查网络连接
- 查看浏览器控制台是否有错误信息
