# Компонент кода

## Описание

Компонент позволяет выполнять пользовательский код. 

## Данные компонента

```
data: {
    code: string
}
```

Поле `code` содержит выполняемый код. 
Документация по языку находится 
[здесь](https://github.com/botscubes/bql/blob/main/docs/lang.md).

## Выходы 

```
outputs: {
    nextComponentId: number,
    idIfError: number
}
```
