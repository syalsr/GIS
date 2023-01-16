create table if not exists stop
(
	id serial primary key,
	name varchar(128) not null unique,
	longitude int,
    latitude int
);
create table if not exists bus
(
	id serial primary key,
	name varchar(128) not null unique,
    is_roundtrip boolean
);
create table if not exists bus_trip
(
	id serial primary key,
	stop_id int references stop(id),
	bus_id int references bus(id)
);
create table if not exists curvature
(
	stop_id_from int references stop(id),
	stop_id_to int references stop(id),
	length int
);