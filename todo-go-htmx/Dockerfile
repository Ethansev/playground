FROM postgres:16

COPY initdb/init.sql /docker-entrypoint-initdb.d

# would not do this in production
ENV POSTGRES_USER=ethan
ENV POSTGRES_PASSWORD=todo
ENV POSTGRES_DB=todo

CMD ["docker-entrypoint.sh", "postgres"]