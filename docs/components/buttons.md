# Компонент кнопок

## Описание

Выводит кнопки для пользователя. Вывод кнопок происходит через [интерфейс ввода-вывода](../io.md).

## Данные компонента

```

type Button = {
    text: string;
}

data: {
    text: string,
    buttons: {
        <numeric field name>: Button,
        ...
    }
}
```

Поле text - текст для опроса пользователя.
Поле buttons - кнопки, которые являются возможным выбором пользователя. Каждая кнопка имеет текст. Ключ, который идентифицирует кнопку, должен иметь числовое значение в виде строки.

## Выходы 

```
outputs: {
    <numeric field name>: number,
    ...
}
```

Имена выходов совпадают с элементами поля buttons в data.

