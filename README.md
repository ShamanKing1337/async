# async

Добавим функцию отчетов. Создадим микросервис, который   принимает от главного сервиса запрос на создание отчета. Сервис отчетов должен быть запущен минимум в 2 инстансах, при этом каждый отчёт должен быть обработан единожды. Нужно, чтобы сервис отчетов справлялся с нагрузкой в 1000 запросов в секунду. Бизнес логику создания отчета можно заменить time.sleep(10sec)
