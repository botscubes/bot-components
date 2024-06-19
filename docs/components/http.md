# Компонент http

## Описание

Компонент позволяет обмениваться данными по сети с помощью протокола HTTP.




## Данные компонента

```
data: {
    url: string,
    method: string,
    header: string,
    body: string
}
```

Поля `url`, `body` перед отправкой [форматируются](./format.md#управляющие-символы-и-конструкции), если там имеются вставки данных из контекста.


Поле `header` представляет собой объект JSON и имеет следующую структуру:
```
{
    "<name>":"<value>",
    ...
}
```

Пример:
```
{
    "Content-Type": "application/json"
}
```


## Выходы 

```
outputs: {
    nextComponentId: number,
    idIfError: number
}
```