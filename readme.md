# Weather API
Реализация простого API клиента для OpenWeather

## Запрос
GET http://localhost:8080/ask
```json
{
    "location": "Самара"
}
```

## Ответ
200 OK
```json
{
    "Location": "Samara",
    "Temperature": 9.59,
    "WeatherType": "Clouds",
    "Pressure": 1023,
    "Humidity": 53,
    "WindSpeed": 1.88
}
```
