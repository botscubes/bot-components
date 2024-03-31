# Значение контекста

## Описание

Значение [контекста](./context.md) может принимать любое значение, поэтому для приведения его к нужному типу служит объект Value.

## Доступные методы

Методы для приведения неопределенного значения к требуему типу:

```golang
func (v *Value) ToString() (string, error) 

func (v *Value) ToInt64() (int64, error) 

func (v *Value) ToInt() (int, error) 

func (v *Value) ToBool() (bool, error) 

```

