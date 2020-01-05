## Setup



##  Docker commands
```
  docker version

  docker info

  docker run -it alpine sh

  docker images

  docker run -d -p 80:80 nginx   # run as deamon, 映射端口

  docker kill <container id>  
```

## Elastic
```
  docker run -d -p 9200:9200 elasticsearch:6.5.0
```


create item
```
PUT http://localhost:9200/imooc/course/1
content-type: application/json

{
  "name":"test"
}
```
/index/type/id
index = database
type = table

Search all
```
PUT http://localhost:9200/imooc/course/_search
```

## elastic go client
```
  go get -v gopkg.in/olivere/elastic.v6
```