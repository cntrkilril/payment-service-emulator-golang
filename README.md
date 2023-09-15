# Эмулятор платежного сервиса

## Для запуска приложения

```sh
make build && make run
```

## Описание

### Методы

---

#### Транзакции

- ``POST`` ``body`` ``/api/transaction`` ``Создание транзакции``

| Имя       | Тип     | Описание                   | Дополнительное     |
|-----------|---------|----------------------------|--------------------|
| userID    | string  | идентификатор пользователя | required, >0       |
| userEmail | string  | почта пользователя         | required, is email |
| amount    | decimal | сумма                      | required, >=0      |
| currency  | string  | валюта                     | required, len>0    |

**Request**

```
{
    "userID": 1,
    "userEmail": "a@mail.ru",
    "amount": 1.0,
    "currency": "RUB"
}
```

**Response**

```
{
    "data": {
        "transaction": {
            "id": 2,
            "userID": 1,
            "userEmail": "a@mail.ru",
            "amount": "1",
            "currency": "RUB",
            "createdAt": "2023-09-15T13:47:25.343693Z",
            "updatedAt": "2023-09-15T13:47:25.343693Z",
            "status": "New"
        }
    },
    "metadata": {
        "success": true
    }
}
```

- ``PATCH`` ``body && params`` ``/api/transaction/status/{id}`` ``Обновление статуса транзакции``

АВТОРИЗАЦИЯ:
Header: Authorization = {systemPaymentID} (id системы оплаты)

P.S авторизация моковая, можно изучить в коде, используется через middleware

| Имя           | Тип    | Описание                 | Дополнительное                       |
|---------------|--------|--------------------------|--------------------------------------|
| id (params)   | int    | идентификатор транзакции | required, >0                         |
| status (body) | string | статус транзакции        | required, =="Success" or =="Failure" |

**Request**

```
{
    "status": "Failure"
}
```

**Response**

```
{
    "data": {
        "transaction": {
            "id": 3,
            "userID": 1,
            "userEmail": "a@mail.ru",
            "amount": "1",
            "currency": "RUB",
            "createdAt": "2023-09-15T13:57:34.583812Z",
            "updatedAt": "2023-09-15T13:57:39.53181Z",
            "status": "Failure"
        }
    },
    "metadata": {
        "success": true
    }
}
```

- ``GET`` ``params`` ``/api/transaction/status/{id}`` ``Получение статуса транзакции``

| Имя | Тип | Описание                 | Дополнительное |
|-----|-----|--------------------------|----------------|
| id  | int | идентификатор транзакции | required, >0   |

**Request**

```
GET /api/transaction/status/1
```

**Response**

```
{
    "data": {
        "transaction": {
            "id": 1,
            "status": "Error"
        }
    },
    "metadata": {
        "success": true
    }
}
```

- ``GET`` ``query`` ``/api/transaction/userID?userID={}&limit={}&offset={}`` ``Получение транзакций пользователя``

| Имя    | Тип | Описание                   | Дополнительное     |
|--------|-----|----------------------------|--------------------|
| userID | int | идентификатор пользователя | required, >0       |
| limit  | int | limit пагинации            | required, >0, <101 |
| offset | int | offset пагинации           | required, >=0      |

**Request**

```
GET api/transaction/user-id?userID=1&limit=2&offset=0
```

**Response**

```
{
    "data": {
        "transactions": [
            {
                "id": 1,
                "userID": 1,
                "userEmail": "a@mail.ru",
                "amount": "1",
                "currency": "RUB",
                "createdAt": "2023-09-15T13:47:18.972345Z",
                "updatedAt": "2023-09-15T13:47:18.972345Z",
                "status": "Error"
            },
            {
                "id": 2,
                "userID": 1,
                "userEmail": "a@mail.ru",
                "amount": "1",
                "currency": "RUB",
                "createdAt": "2023-09-15T13:47:25.343693Z",
                "updatedAt": "2023-09-15T13:47:35.864202Z",
                "status": "Failure"
            }
        ]
    },
    "metadata": {
        "count": 3,
        "success": true
    }
}
```

- ``GET`` ``query`` ``/api/transaction/email?email={}&limit={}&offset={}`` ``Получение транзакций по почте``

| Имя    | Тип    | Описание           | Дополнительное     |
|--------|--------|--------------------|--------------------|
| email  | string | почта пользователя | required, is email |
| limit  | int    | limit пагинации    | required, >0, <101 |
| offset | int    | offset пагинации   | required, >=0      |

**Request**

```
GET api/transaction/email?email=a@mail.ru&limit=1&offset=0
```

**Response**

```
{
    "data": {
        "transactions": [
            {
                "id": 1,
                "userID": 1,
                "userEmail": "a@mail.ru",
                "amount": "1",
                "currency": "RUB",
                "createdAt": "2023-09-15T13:47:18.972345Z",
                "updatedAt": "2023-09-15T13:47:18.972345Z",
                "status": "Error"
            }
        ]
    },
    "metadata": {
        "count": 3,
        "success": true
    }
}
```

- ``GET`` ``params`` ``/api/transaction/cancel/{id}`` ``Отмена транзакции``

| Имя | Тип | Описание                 | Дополнительное |
|-----|-----|--------------------------|----------------|
| id  | int | идентификатор транзакции | required, >0   |

**Request**

```
GET api/transaction/cancel/4
```

**Response**

```
{
    "data": {
        "transaction": {
            "id": 4,
            "userID": 1,
            "userEmail": "a@mail.ru",
            "amount": "1",
            "currency": "RUB",
            "createdAt": "2023-09-15T14:06:08.530611Z",
            "updatedAt": "2023-09-15T14:06:13.18729Z",
            "status": "Cancel"
        }
    },
    "metadata": {
        "success": true
    }
}
```
