-- name: CreateUser :one
insert into users (username,email,password)
values ($1,$2,$3)
returning id,username,email,created_at;

-- name: GetUserById :one
select * from users
where id = $1
limit 1;

-- name: GetUsers :many
select * from users
order by id
limit $1
offset $2;