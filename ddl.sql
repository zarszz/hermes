create table product (
	id SERIAL primary KEY,
	sku varchar(16) not null unique,
	name varchar(128) not null,
	display varchar(128) not null
)

create table users (
	id SERIAL primary key,
	username varchar(48),
	email varchar(128) unique,
	password varchar(128)
)
