-- name: CreateUser :one
insert into users (username,email,password)
values ($1,$2,$3)
returning *;

-- name: GetUserByEmail :one
select * from users
where email = $1
limit 1;

-- name: GetUserById :one
select * from users
where id = $1
limit 1;


-- name: GetUsers :many
select * from users
order by id
limit $1
offset $2;