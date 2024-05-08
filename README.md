# helperBot

## Структура проекта:

## /cmd

Основная логика работы cli-приложения для запуска и конфигурации бота

## /configs

Конфигурации приложения

## /handlers

Обработка всех входящих запросов

Здесь определяется сущность, ответсвенная за обработку запросов (Бот). 

В данной директории нет логики работы бота и просто передает проверенные запросы соответствующим службам и возвращает ответ клиенту.

## /logs

Сюда сохраняются логи

## /models

Структуры и типы, которые необходимы для работы приложения

## /pkg

Библиотеки или функции, которые помогают поддерживать и улучшать приложение.

## /repositories

Базы данных и их функциональность

## /services

Здесь существует основная бизнес-логика, она работает как мост между каталогами «обработчиков» и «репозиториев».