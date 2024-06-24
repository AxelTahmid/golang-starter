# TODO: make a makefile command to build & run this image, using env values
FROM migrate/migrate as migrate

# Copy all db files
COPY ./database/migrations /migrations

ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]

CMD ["postgresql://${DB_USER}:${DB_PASSWORD}@db:5432/${DB_NAME}?sslmode=disable -verbose up"]