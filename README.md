# Creating network for services

```$ docker network create proyekspt```

# Creating UI service

```$ cd ui-service; docker build --tag ui-service .;docker run -d --name ui-service --network proyekspt ui-service```

# Creating mysql service

```$ docker run -d --name spt-mysql -e MYSQL_ROOT_PASSWORD=proyekspt  --network proyekspt mysql:latest```

```$ docker exec -it spt-mysql mysql -u root -p'proyekspt'```

```$ > CREATE DATABASE proyekspt;```

```$ > exit;```

# Staff-service

```$ cd staff-service```

```$ docker build --tag staff-service .```

```$ docker run -d --name staff-service --network proyekspt staff-service```

# Reservation service

```$ cd reservation-service/app```

```$ docker build --tag res-service .```

```$ docker run -d --name res-service --network proyekspt res-service```

# Creating web server

```$ cd nginx```

```$ docker build --tag webserver .```

```$ docker run -d -p 443:443 --name webserver --network proyekspt webserver```

