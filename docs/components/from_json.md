# Компонент получения переменных из JSON

## Описание

Компонент получает новые переменные из JSON.

## Путь назначения

Полученные переменные сохраняется в переменную по пути в поле `path`.

После преобразования, к переменным можно обращаться следующим образом:
```
<path>.<field> // Если JSON содержал объект
<path>[<index>] // Если JSON содержал массив
```

## Данные компонента

```
data: {
    json: string
}
```

Поле json - текст, содержащий JSON. Перед обработкой [форматируется](./format.md#управляющие-символы-и-конструкции), если там имеются вставки данных из контекста.


## Выходы 

```
outputs: {
    nextComponentId: number,
    idIfError: number
}
```