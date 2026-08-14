# GDPS Tools

A collection of utilities designed to simplify diagnostics, configuration, and maintenance of GDPS servers. The project provides modular tools for checking server health, generating configuration files, analyzing logs, and validating network tunnels.

---

## English Version

## Overview
GDPS Tools is a modular toolkit intended for GDPS server administrators and developers. It automates routine tasks, helps identify configuration issues, and provides structured diagnostics for both local and remote GDPS environments.

---

## Features

### 1. GDPS Diagnostic Scanner
A diagnostic module for checking the operational state of a GDPS server:
- MySQL availability
- FTP availability
- Port accessibility
- UDP/TCP tunnel testing
- Connection stability checks
- Detailed diagnostic reports

### 2. Config Builder
A configuration generator for GDPS and Globed:
- Automatic creation of `config.toml`
- Template-based configuration generation
- Error validation
- Export of ready-to-use configuration files

### 3. Playit.gg Tunnel Tester
A network testing module:
- Port 4342 (UDP) validation
- Port 5351 (TCP/UDP) validation
- Latency measurement
- Tunnel stability analysis

### 4. Log Analyzer
A log analysis module:
- Error detection
- Categorization of issues
- Statistical summaries
- Recommendations for resolving common problems

---

## Planned Modules
- GDPS API Wrapper (Python/TypeScript SDK)
- Automated GDPS deployment tool
- GDPS Dashboard (web-based management panel)

---

## Installation
Installation instructions will be added later.

---

## Documentation
Each module will include a dedicated README with usage examples and technical details.

---

## Contributing
Contributions are welcome.  
To propose improvements or report issues, please open an Issue.

---

## License
MIT License.

---

## Русская версия

## Обзор
GDPS Tools — это набор модульных инструментов для администраторов и разработчиков GDPS. Он автоматизирует рутинные задачи, помогает выявлять ошибки конфигурации и предоставляет структурированную диагностику для локальных и удалённых GDPS‑серверов.

---

## Возможности

### 1. GDPS Diagnostic Scanner
Модуль диагностики состояния GDPS‑сервера:
- доступность MySQL
- доступность FTP
- проверка портов
- тестирование UDP/TCP‑туннелей
- проверка стабильности соединения
- подробные отчёты о состоянии сервера

### 2. Config Builder
Генератор конфигураций для GDPS и Globed:
- автоматическое создание `config.toml`
- генерация конфигов по шаблонам
- проверка ошибок
- экспорт готовых файлов

### 3. Playit.gg Tunnel Tester
Модуль проверки туннелей:
- проверка порта 4342 (UDP)
- проверка порта 5351 (TCP/UDP)
- измерение задержки
- анализ стабильности туннеля

### 4. Log Analyzer
Модуль анализа логов:
- поиск ошибок
- классификация проблем
- статистика
- рекомендации по устранению распространённых ошибок

---

## Планируемые модули
- GDPS API Wrapper (Python/TypeScript SDK)
- Автоматический деплой GDPS
- GDPS Dashboard (веб‑панель управления)

---

## Установка
Инструкция по установке будет добавлена позже.

---

## Документация
Каждый модуль будет иметь отдельный README с примерами использования.

---

## Вклад
Pull‑requests приветствуются.  
Для предложений и отчётов об ошибках создавайте Issue.

---

## Лицензия
MIT License.
