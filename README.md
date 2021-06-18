# datetime-service
***
### Ручки:
+ Проверка работоспособности сервиса:  
  Request: **[GET] /alive**

Response:
<pre>
200
</pre>

+ Получение актуального времени (с учетом коррекции) в формате F64:  
  Request: **[GET] /time/now**  

Response:
<pre>
200
{
    "time": 210618.135307869
}
</pre>

+ Получение строкового представления времени формата F64:  
  Request: **[GET] /time/string**
  
Body:
<pre>
{
  "time":210617.114229255
}
</pre>

Response:
<pre>
200
{
    "str": "2021-06-17 11:42:29.255 +0000 UTC"
}

400
invalid time format: parsing time "000017114229.255": month out of range
</pre>

+ Получение времени в формате F64 со смещением на дельту в формате F64 
  (формат дельты схож с форматом даты, с тем исключением, что можно указывать только 
  дни и время, и количество дней ограничено 99):  
  Request: **[GET] /time/string**

Body:
<pre>
{
    "time":210617.114229691,
    "delta":-1.114229
}
</pre>

Response:
<pre>
200
{
    "time": 210616.000000691
}

400
invalid delta format: days are limited to 99
</pre>

+ Скорректировать время сервера:  
  Request: **[POST] /time/correct**

Body:
<pre>
{
    "time":210618.133000
}
</pre>

Response:
<pre>
200

400
invalid time format: parsing time "000018133000.000": month out of range
</pre>

***
### Запуск:
<pre>
docker-compose up
</pre>