#Cache
##Описание RU 
###Данная библиотека создана в качестве домашнего задания для курса Golang-Ninja. 
Это вторая версия библиотеки. Отличия от первой в том что у кадого объекта кеша есть время жизни и устанавливается оно при вставки значения в кеш. В анонимной функции в отдельной го-рутине зачищаются те значения у которых закончилось время жизни. Так же все функции теперь имеют защиту от состояния гонки. При чтении или записи лочится мютекс. Было произведено тестирование с помощью флага -race при запуске приложения.
###Реализованны следующие функции:
####_cache.New()_
Функция конструктор, возвращает объект кеш хранилица
####_cache.Set(key,value,timeLive)_
Функция установки значения в хранилище, принимает ключ, значение и время жизни. Ключ имеет тип строки, значение может быть любым типом. Время жизни имеет тип _time.Duration_ Вызов функции происходит от объекта хранилища.
####_cache.Get(key)_
Функция получения значения по ключу. Возвращает устоновленное значение, если значение было удалено, то вернет пустой указатель nil.
####_cahe.Delete(key)_
Функция удаления значения по ключу. Удаляет выбранное значение из хранилища.
##Description EN
###This library was created as a homework assignment for the Golang-Ninja course. 
This is the second version of the library. The difference from the first one is that each cache object has a lifetime and it is set when inserting a value into the cache. In an anonymous function, those values whose lifetime has ended are cleaned in a separate routine. Also, all functions now have protection from the race condition. When reading or writing, the mutex is locked. Testing was performed using the -race flag when launching the application.
###The following functions are implemented:
####_cache.New()_
Constructor function, returns a cache object
####_cache.Set(key,value,timeLive)_
The function of setting the value in the storage, accepts the key, value and lifetime. The key has a string type, the value can be any type. The lifetime is of type _time.Duration_ The function call comes from the storage object.
####_cache.Get(key)_
The function of getting the value by the key. Returns a fixed value, if the value was deleted, it will return an empty nil pointer.
####_cahe.Delete(key)_
The function of deleting a value by key. Deletes the selected value from the storage.
