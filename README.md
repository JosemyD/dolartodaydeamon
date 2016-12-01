# DolarToday Deamon

Un simple programa en Golang, que se encarga de revisar DolarToday por cambios en los precios y almacena el historial en MongoDB.

##Despliegue

Se requiere tener instalado [glide](https://github.com/Masterminds/glide), y una instancia de MongoDB en ejecución. En el archivo *conf/config.ini* se definen algunos parametros de configuración.

```
>>> git clone https://github.com/JosemyD/dolartodaydeamon.git
>>> cd dolartodaydeamon
>>> glide install
>>> go run main.go
```

