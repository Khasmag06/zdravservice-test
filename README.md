Сервис разработан с использованием современных технологий и следует принципам Clean Architecture,
что обеспечивает легкость расширения функционала и тестирования. Также был реализован Graceful Shutdown
для корректного завершения работы сервиса.

# Getting Started


Для запуска сервиса, вам нужно всего лишь заполнить файл .env в корневой директории, следуя примеру файла .env.example,
где установлены значения для запуска через Docker.

# Usage

Запустить сервис можно с помощью команды `make compose-up`.

Документацию после запуска сервиса можно посмотреть по адресу `http://localhost:8080/swagger/index.html`
с портом 8080 по умолчанию.

Для запуска тестов необходимо выполнить команду `make test`, для запуска тестов с покрытием `make cover` и `make cover-html` для получения отчёта в html формате.

# Decisions <a name="decisions"></a>

В ходе разработки был сомнения по тем или иным вопросам, которые были решены следующим образом:

1. Как выводить уникальные идентификаторы товаров с несколькими свойствами?
> При разработке функциональности для запроса товаров, возник вопрос о том, как правильно выводить уникальные идентификаторы товаров,
которые имеют несколько свойств. Рассматривалась идея выводить все свойства, привязанные к уникальному идентификатору,
но в конечном итоге от неё было решено отказаться. Это решение было принято во избежание возможных проблем с корректностью пагинации.
