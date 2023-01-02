# quizON

# MVP

# Пользовательские сценарии

## [`Залогиниться как админ`](#Залогиниться-как-админ)

## [`Создать игру`](#Создать-игру)

## [`Отредактировать игру`](#Отредактировать-игру)

## [`Получить список игр`](#Получить-список-игр)

## [`Зарегистрировать команду на игру`](#Зарегистрировать-команду-на-игру)

### Залогиниться как админ

**Путь:** `POST /login`

**Запрос:**

```json lines
{
  // login - логин
  "login": "slukash",
  // password - пароль
  "password": "1234567890"
}
```

**Ответ:**

- `200` - Удачный логин, ставится кука, которая позволяет делать действия админа
- `403` - Неверный логин или пароль

### Создать игру 
(для выполнения запроса нужно залогиниться)

**Путь:** `POST /game`

**Запрос:**

```json lines
{
  // name - название игры
  "name": "Квизон #2.5",
  // description - описание игры
  "description": "Игра по Гарри Поттеру",
  // date - время начала игры в UTC
  "date": "2018-12-10T13:49:51.141Z",
  // teams_amount - количество команд на игре
  "teams_amount": 20,
  // registered_teams - количество зарегистрированных команд
  "registered_teams": 17,
  // price_per_person - цена за человека в рублях
  "price_per_person": 500,
  // location - местро проведения игры
  "location": "345 аудитория"
}
```

**Ответ**:

- `200` - Игра создалась

```json lines
{
  // id - уникальный идентификатор игры (проставляется на бэкэнде)
  "id": 1,
  // name - название игры
  "name": "Квизон #2.5",
  // description - описание игры
  "description": "Игра по Гарри Поттеру",
  // date - время начала игры в UTC
  "date": "2018-12-10T13:49:51.141Z",
  // teams_amount - количество команд на игре
  "teams_amount": 20,
  // registered_teams - количество зарегистрированных команд
  "registered_teams": 17,
  // price_per_person - цена за человека в рублях
  "price_per_person": 500,
  // location - местро проведения игры
  "location": "345 аудитория"
}
```

- `403` - Действие сделано неавторизованным пользователем (на самом деле у таких пользователей не должно быть видно админку, но на всякий на бэке будем проверять куку)

## Отредактировать игру

**Путь:** `PATCH /game`

**Запрос:** (нужно передать все поля, даже не измененные, т.к. на бэке перезаписываются все поля)

```json lines
{
  // id - уникальный идентификатор игры (проставляется на бэкэнде)
  "id": 1,
  // name - название игры
  "name": "Квизон #2.5",
  // description - описание игры
  "description": "Игра по Гарри Поттеру",
  // date - время начала игры в UTC
  "date": "2018-12-10T13:49:51.141Z",
  // teams_amount - количество команд на игре
  "teams_amount": 20,
  // price_per_person - цена за человека в рублях
  "price_per_person": 500,
  // location - местро проведения игры
  "location": "345 аудитория",
  // registeredTeamsAmount - количество зарегестрированных команд
  "registeredTeamsAmount": 12
}
```

**Ответ:**

- `200` - Игра успешно изменена

- `403` - Действие выполнено неавторизованным пользователем

- `404` - Игра не найдена

## Получить список игр

**Путь:** `GET/games`

**Запрос:**(на mvp просто игры отсортированные по дате проведения)

```json lines
{
  // page - страница пагинации (начиная с 1)
  "page": 1,
  // perPage - сколько элементов на странице
  "perPage": 10
}
```

**Ответ:**

- `200` - все ок игры найдены (если игр нет -> пустой массив)

```json lines
{
  // games - массив игр отсортированный по дате (от новых к старым)
  "games": [
    {
      "id": 2,
      "name": "Квизон #2.5",
      "description": "Игра по Гарри Поттеру",
      "date": "2022-12-10T13:49:51.141Z",
      "teams_amount": 20,
      "price_per_person": 500,
      "location": "345 аудитория"
    },
    {
      "id": 1,
      "name": "Квизон #2.5",
      "description": "Игра по Гарри Поттеру",
      "date": "2010-12-10T13:49:51.141Z",
      "teams_amount": 20,
      "price_per_person": 500,
      "location": "345 аудитория"
    }
  ]
}
```

## Зарегистрировать команду на игру

**Путь:** `POST /registerTeam`

**Запрос:**

```json lines
{
  // gameID - id игры на которую нужно зарегаться
  "gameID": 1,
  // tramName - навзание команды
  "teamName": "Ураган Донам",
  // captainName - имя капитана
  "captainName": "Постя Коляков",
  // phone - номер телефона
  "phone": "8800553535",
  // telega - телега
  "telega": "@SomeNick",
  // amount - количество человек в команде
  "amount": 5
}
```

- `200` Успешная регистрация на игру