# Task Manager

## Описание

Task Manager — это программа на языке Go, которая генерирует задачи, обрабатывает их и периодически выводит результаты. Программа использует каналы и горутины для асинхронной обработки задач, а также логгирует все действия в файл `app.log`.

## Функциональность

1. **Генерация задач**: Программа генерирует задачи в течение 10 секунд и отправляет их в канал `taskChan`.
2. **Обработка задач**: Задачи из канала `taskChan` обрабатываются и сортируются в каналы `doneTasks` и `undoneTasks`.
3. **Периодическая печать результатов**: Программа периодически печатает результаты задач каждые 3 секунды.
4. **Логгирование**: Все действия логгируются в файл `app.log`.
5. **Грациозное завершение**: Программа завершает свою работу после нажатия клавиши Enter пользователем.

## Структура проекта

- `main.go`: Основной файл программы, содержащий всю логику.
- `app.log`: Файл для логгирования действий программы.

## Как запустить

1. Убедитесь, что у вас установлен Go.
2. Склонируйте репозиторий.
3. Перейдите в директорию проекта.
4. Запустите программу командой:

    ```sh
    go run main.go
    ```

## Пример использования

После запуска программы вы увидите сообщение о том, что файл журнала создан, и обратный отсчет перед началом выполнения основной логики. Программа будет генерировать и обрабатывать задачи, периодически выводя результаты. Для завершения программы нажмите клавишу Enter.

## Лицензия

Этот проект лицензирован под лицензией MIT. Подробности см. в файле LICENSE.
