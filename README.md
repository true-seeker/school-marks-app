### Лабораторная работа по курсу "Анализ уязвимостей программного обеспечения"

#### Запуск приложения

1. Docker:
   ```
   docker-compose -f deployments/docker-compose.yaml up
   ```
2. Сборка golang:
    - [Установить](https://go.dev/doc/install) go
    - Перейти в папку с точкой входа
      ```
      cd cmd/school-marks-app
      ```
    - Запустить сборку
      ```
      go build
      ```
