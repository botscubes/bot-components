# Интерфейс ввода-вывода

## Описание

Интерфейс ввода-вывода используется при [выполнении компонентов](./executor.md). Он позволяет взаимодействовать со средой, в которой выполняются компоненты. 

## Структура интерфейса

Интерфейс ввода-вывода имеет следюущий вид:

```golang

type IO interface {
	PrintText(text string) error
	PrintButtons(text string, buttons []*ButtonData) error
	PrintPhoto(file []byte, name string) error
	ReadText() *string
}
```

Окружение, которое будет использовать компоненты, должна реализовать данный интерфейс.

Интерфейс состоит из следующих методов:
- PrintText - Вывод текста пользователю.
- PrintButtons - Вывод кнопок пользователю с указанием текста опроса.
- PrintPhoto - Вывод картинки пользователю.
- ReadText - Считывание текста от пользователя, если текст не введен, возвращается нулевое значение - nil.

