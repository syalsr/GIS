create table if not exists stop
(
	stop_id serial primary key,
	name varchar(128) not null unique,
	longitude int,
    latitude int
);
create table if not exists bus
(
	bus_id serial primary key,
	name varchar(128) not null unique,
    is_roundtrip boolean
);
create table if not exists bus_stop
(
	stop_name varchar(128),
	bus_name varchar(128)
);
create table if not exists curvature
(
	stop_id_from int references stop(stop_id),
	stop_id_to int references stop(stop_id),
	length int
);