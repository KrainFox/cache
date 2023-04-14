#Cache
##Описание RU 
###Данная библиотека создана в качестве домашнего задания для курса Golang-Ninja. В ней реализованны следубщие функции:
####_cache.New()_
Функция конструктор, возвращает объект кеш хранилица
####_cache.Set(key,value)_
Функция установки значения в хранилище, принимает ключ и значение. Ключ имеет тип строки, значение может быть любым типом. Вызов функции происходит от объекта хранилища.
####_cache.Get(key)_
Функция получения значения по ключу. Возвращает устоновленное значение, если значение было удалено, то вернет пустой указатель nil.
####_cahe.Delete(key)_
Функция удаления значения по ключу. Удаляет выбранное значение из хранилища.
##Description EN
###This library was created as a homework assignment for the Golang-Ninja course. It implements the following functions:
####_cache.New()_
Constructor function, returns a cache object
####_cache.Set(key,value)_
The function of setting a value in the storage, accepts a key and a value. The key has a string type, the value can be any type. The function call comes from the storage object.
####_cache.Get(key)_
The function of getting the value by the key. Returns a fixed value, if the value was deleted, it will return an empty nil pointer.
####_cahe.Delete(key)_
The function of deleting a value by key. Deletes the selected value from the storage.
