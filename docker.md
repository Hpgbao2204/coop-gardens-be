# How to run docker postresql step-by-step

**Requirements**

1. Docker-compose: v3.12 
2. psql: (PostgreSQL) 17.2 (Ubuntu 17.2-1.pgdg24.04+1)

U should focus `docker-compose.yml`. Don't care `Dockerfile`

**Command:**
```
   bash
   docker-compose down
   docker-compose up -d
```

Go on: [pgadmi4-Website](http://localhost:5050/browser/)

---

**Follow**:

   1. Add new server
   2. In *General* tab -> Name: type **coop_gardens_db**
   3. In *Connection* tab:
      1. *Host name/address*: type **Ip address host**
      2. *Maintenance database* + *Username*: type  **coop_gardens**
      3. *Password*: type **Password**
   4. Save

---

**How to work remote with database**:

```bash
docker-compose up -d
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' coop_gardens_db
```

Output: Ip address who run this container

Connect
