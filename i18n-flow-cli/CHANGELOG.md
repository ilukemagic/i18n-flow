# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [0.0.6](https://github.com/ilukemagic/i18n-flow/compare/v0.0.5...v0.0.6) (2025-05-04)


### Features

* Refactor API client initialization and connection testing to support dynamic configuration updates; enhance user input handling in init command with Chinese comments for better clarity. ([8c6528d](https://github.com/ilukemagic/i18n-flow/commit/8c6528d5b3b08e0b1d88219482d2ef29f1b10d46))

### [0.0.5](https://github.com/ilukemagic/i18n-flow/compare/v0.0.4...v0.0.5) (2025-05-04)

### 0.0.4 (2025-05-04)


### Features

* Add Chinese and English contributing guidelines and README documentation to enhance project accessibility and support for international contributors. ([b54a98e](https://github.com/ilukemagic/i18n-flow/commit/b54a98e64b68d3391a77df92b08c971b83e0b053))
* Add Chinese README documentation for i18n-flow CLI to improve accessibility and support for Chinese-speaking users. ([734b0cf](https://github.com/ilukemagic/i18n-flow/commit/734b0cf0d3a5b50ad5a7c698fb9597b6a9e342c5))
* Add CLI functionality with new API endpoints for translation management. Implement CLIController to handle requests for checking API keys, retrieving translations, and pushing new keys. Update configuration and middleware for API key validation. Enhance documentation and Swagger specifications for new CLI routes. ([95ef054](https://github.com/ilukemagic/i18n-flow/commit/95ef054e78b8bf4a75e063279dfb658723354e25))
* Add comprehensive English and Chinese translations for login, dashboard, project management, and translation management features, enhancing internationalization support across the application. ([bb6e1b8](https://github.com/ilukemagic/i18n-flow/commit/bb6e1b84242cd0c592c8c5b1304d14bb0fbdeda2))
* Add comprehensive README documentation for i18n-flow backend and frontend, detailing features, installation instructions, API endpoints, and project structure to assist developers in understanding and utilizing the system effectively. ([d89a86e](https://github.com/ilukemagic/i18n-flow/commit/d89a86e2848d52640d92ff7a0995141ab55abf4e))
* Add Docker configuration and environment setup for i18n-flow application, including backend and frontend services, database configuration, and example environment variables for seamless deployment. ([f63b586](https://github.com/ilukemagic/i18n-flow/commit/f63b586cdba07945eebfccc63d908d7978a52784))
* Add endpoint and functionality to retrieve translation matrix by project, including pagination and search support. Update related documentation and frontend components to accommodate new data structure. ([2a84129](https://github.com/ilukemagic/i18n-flow/commit/2a841298ea62cf2e07f5a74cc5fe0e908af32ad5))
* Add README documentation for i18n-flow CLI, detailing features, installation, commands, configuration, and usage examples to assist users in managing translations effectively. ([39601b9](https://github.com/ilukemagic/i18n-flow/commit/39601b95d44dc2f6ff4cf3c05642ade330cc1e9f))
* Add Zustand for state management and implement language store for improved language handling in the application ([32358ab](https://github.com/ilukemagic/i18n-flow/commit/32358ab85d0f97cbeb0dea6de0179a5e681215aa))
* Enhance BatchTranslationModal to set form values based on paginated translation matrix. Remove unnecessary setTimeout in TranslationManagement for modal visibility. Update DashboardLayout to remove user management navigation. ([6dbfd27](https://github.com/ilukemagic/i18n-flow/commit/6dbfd27166239df6456bc2ca251424b2b75a6c00))
* Enhance README with detailed Docker deployment instructions, including prerequisites, quick setup steps, service access, and troubleshooting tips for i18n-flow application. ([1b57c24](https://github.com/ilukemagic/i18n-flow/commit/1b57c2437d2b426a08740f0070eb571d1ec70d71))
* Enhance translation modals with Chinese comments for better clarity and understanding of functionality ([594955a](https://github.com/ilukemagic/i18n-flow/commit/594955a42370b276239a51a45da9e0c42be08bda))
* Expand documentation with comprehensive English and Chinese guides for installation, usage, and best practices, enhancing accessibility and support for international users. ([9b7339d](https://github.com/ilukemagic/i18n-flow/commit/9b7339d11a71b188532b0832ceadcfee0ad9a371))
* Expand README documentation to include CLI tool integration, setup instructions, and detailed command usage for improved developer guidance on managing translations within i18n-flow. ([29f1670](https://github.com/ilukemagic/i18n-flow/commit/29f16705241e1f3f101a90dc5180a4e394a1e337))
* Implement ColumnSelector component for language column selection in translation management. Enhance TranslationTable and TranslationToolbar to support dynamic column visibility based on user preferences. Add localStorage utilities for saving and loading selected columns. ([3e3d572](https://github.com/ilukemagic/i18n-flow/commit/3e3d572ec1caab176bba193d8cb6ebd84fede546))
* Implement dashboard statistics feature with new DashboardController and service. Add API endpoint for retrieving dashboard stats and update frontend components to display statistics. Enhance documentation for new API routes and data structures. ([85d201b](https://github.com/ilukemagic/i18n-flow/commit/85d201b9fcd8c313d62e54a9395aea88bd4e2f46))
* Initialize i18n-flow documentation with comprehensive guides, API references, and deployment instructions for backend and frontend setups ([9c742ea](https://github.com/ilukemagic/i18n-flow/commit/9c742eaa8c631b6324d4fd7b698c7a93fe391428))
* Integrate i18next for internationalization support, adding English and Chinese translations, and implement language selection functionality in the frontend application. ([44e0c91](https://github.com/ilukemagic/i18n-flow/commit/44e0c9140d4a0b157ffbdd1b0efa7190f1f57cb7))
* Integrate lodash for debounced search functionality in TranslationToolbar. Enhance input handling with local state management and improve search performance by implementing a debounce mechanism on keyword changes. ([ae2a68e](https://github.com/ilukemagic/i18n-flow/commit/ae2a68ed70054a184e7346de95e2eeb7de01d6f4))
* Introduce documentation site for i18n Flow with Docker integration, including Nginx configuration and bilingual support in English and Chinese, enhancing accessibility for users. ([4daf047](https://github.com/ilukemagic/i18n-flow/commit/4daf047c277ecd33f26aa1d520e64d30fbdac57f))
* Introduce i18n-flow CLI for translation management with commands for initialization, syncing, pushing keys, and checking status. Implement core API interactions, configuration handling, and file scanning for translation keys. Enhance logging and error handling throughout the CLI tool. ([78c8bec](https://github.com/ilukemagic/i18n-flow/commit/78c8bec80e1adddb2ec05f5f729d34373d3c1029))
* Refactor translation management to support new data structure with language info, enhance table column generation, and improve import/export functionality for translations. ([ffda1da](https://github.com/ilukemagic/i18n-flow/commit/ffda1da1064a1297278714f886df53e4ec9dc613))
* Revamp DashboardLayout and Login components for improved user experience. Enhance DashboardLayout with user avatar and dropdown menu for profile and logout options. Update Login page design with a gradient background, secure login section, and improved button styles. ([137bfb0](https://github.com/ilukemagic/i18n-flow/commit/137bfb02b4ac12e7440fe668efde5e1b00d6aaa8))
* Update API base URL in axios instance to use environment variable VITE_API_URL for better configuration management; add environment variable to docker-compose.yml for local development. ([e618f40](https://github.com/ilukemagic/i18n-flow/commit/e618f403a572bbfafdc66c43c2fdcae738f8be4d))
* Update frontend guide to reflect correct default admin interface URL and configure VitePress to ignore dead link checks for improved documentation accuracy. ([1792737](https://github.com/ilukemagic/i18n-flow/commit/1792737442696150843d820da1a0acb172c60f4c))
* Update translation management UI to utilize i18next for dynamic text rendering, enhance modal forms with internationalized labels and placeholders, and improve overall user experience in translation workflows. ([86ce5f9](https://github.com/ilukemagic/i18n-flow/commit/86ce5f9bff713ae0ac93cda11ef7da28030d3a29))


### Bug Fixes

* Update button icon and label in TranslationTable for clarity. Change 'Batch add' button to 'Edit' with an EditOutlined icon to better reflect its functionality. ([75c3fdd](https://github.com/ilukemagic/i18n-flow/commit/75c3fdd05f24be62cb261fe2ea3fb431f88963e3))
* Update npm version badge in README files to correct package name for i18n-flow CLI ([5b6729b](https://github.com/ilukemagic/i18n-flow/commit/5b6729ba615ba81b44bf1d834d0c61bd19bbc1b3))
* Update token retrieval and error handling in API utility to use consistent naming for localStorage items and improve user feedback on authentication errors. ([b6eacbb](https://github.com/ilukemagic/i18n-flow/commit/b6eacbb22faba6a564f8737e3ad85c473061d958))
* Update translation management to include empty string values for language fields, ensuring all translations are captured in the request. ([e9b55ec](https://github.com/ilukemagic/i18n-flow/commit/e9b55ec9b5ba8649a68c56b3ad75cdc317facbbd))
