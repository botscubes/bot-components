# Компонент условия

## Описание

Компонент проверяет переменную в контексте на истинность или ложь.

## Данные компонента

```
data: {
    expression: string
}
```
Поле `expression` может хранить значение `true`, `false` или путь к переменной в 
[контексте](../context.md).

Если значение в `expression` или переменная в ней содержит `true`, то переход происходит 
по ветке `nextComponentId`, иначе `idIfFalse`.


## Выходы 

```
outputs: {
    nextComponentId: number,
    idIfFalse: number,
    idIfError: number
}
```
