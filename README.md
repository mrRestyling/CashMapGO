# CashMapGO.
Реализация кеш-хранилища

## Обзор
Реализация кеш-хранилища на Go, предоставляющая базовую систему кеша с функцией TTL (время жизни).

## Интерфейс кеша
Интерфейс Cache определяет методы Get и Set для операций с кешем.

- Метод установки
Метод Set устанавливает элемент кеша с заданным ключом, значением и временем жизни TTL.

- Метод получения
Метод Get получает элемент кеша по ключу, возвращая значение и логическое значение, указывающее, существует ли элемент и не истекло ли время его жизни.

# Документация 
## Интерфейс
- Get(key string) (value any, ok bool): Получает элемент кеша по ключу, возвращая значение и логическое значение, указывающее, существует ли элемент и не истекло ли время его жизни.
- Set(key string, value any, ttl time.Duration): Устанавливает элемент кеша с заданным ключом, значением и временем жизни TTL.
## Структура хранилища
- List map[string]interface{}: Основная мапа, хранящая элементы кеша.
- mu sync.Mutex: Mutex для синхронизации доступа к кешу.
## Структура элемента
- Value interface{}: Хранимое значение.
- Ttl time.Time: Метка времени TTL для хранимого элемента.
## Инициализация 
New(m map[string]any) Storage: Создает новое хранилище кеша из заданной мапы.
