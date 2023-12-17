# go-consul-instance-manager

Библиотека для отслеживания состояния экземпляров какого-либо сервиса через Hashicorp Consul.

## Пример использования

см. в директории [/example](/example/)

## Консистентное хэширование

В основе механизма определения держателя ключа (метод `iman.GetDataOwner(key)`) лежит механизм консистентного хэширования на основе хэш-кольца. Подробнее можно почитать [ЗДЕСЬ](https://habr.com/ru/companies/mygames/articles/669390/) и [ЗДЕСЬ](https://habr.com/ru/companies/timeweb/articles/691506/).

## Диаграмма классов

![Диаграмма классов](./docs/class_diagram.png)

## instance status

![instance_status_diagram](/docs/instance_status_diagram.png)
