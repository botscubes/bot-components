# Контекст

## Описание

Передача данных между компонентами происходит через Context.
Так как структура данных контекста изначально не определена, внутри 
Context находятся данные неопределенного типа.


## Создание 

Объект Context создаётся из JSON следующим методом:

```golang
func NewContextFromJSON(jsonData []byte) (*Context, error) 
```

## Вставка данных

Для вставки данных в Context служит следующий метод:

```golang

func (ctx *Context) SetValue(path string, value *any) error

```

## Обращение к данным

Для обращения к данным у структуры Context реализован следующий метод:

```golang
func (d *Context) GetValue(path string) (*Value, error)
```

Функция возвращает значение типа [Value](./value.md).


Параметр путь к данным (path) имеет следующий синтаксис:

- Для обращения к полю Context
```
<property name>
```

- Для обращения к свойству объекта:
```
.<property name>
```

- Для обращения к данным массива:
```
[<index>]
```

- Также возможна их кобинация:

Пример:
```
users[1].phoneNumbers[2]
```
```
users[2].age
```

- Неявное обращение к элементу массива:

Пример:
```
users[userIndex].phoneNumbers[phoneNumberIndex]
```

- Неявное обращение к свойствy объекта:
```
.[<path to value>]
```

Пример:
```
[objects[objectIndex].propertyName].[properyNameStoringNameOfArray][1]
```

В случае, если синтаксис пути будет неверен, или там не будет находиться нужная структура или массив, будет возвращена ошибка.
