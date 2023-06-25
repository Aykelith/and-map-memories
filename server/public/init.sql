CREATE TABLE pins (
    id SERIAL4 NOT NULL,
	name varchar(128) NOT NULL UNIQUE,
	country varchar(128) NOT NULL,
	county varchar(128) NOT NULL,
	location_name varchar(128) NOT NULL,
	location_address varchar(256) NOT NULL,
	pin_type INT2 NOT NULL,
	lat_integral INT4 NOT NULL,
	lat_fractional INT8 NOT NULL,
	lng_integral INT4 NOT NULL,
	lng_fractional INT8 NOT NULL,
	CONSTRAINT pk_pins_id PRIMARY KEY (id)
);

CREATE TABLE pins_visits (
    id SERIAL4 NOT NULL,
	pin_id INT4 NOT NULL,
	visit_date timestamp NOT NULL,
	photos_dir varchar(128) NULL,
	CONSTRAINT pk_pins_visits_id PRIMARY KEY (id),
	CONSTRAINT fk_pins_visits_pinid FOREIGN KEY (pin_id) REFERENCES pins(id)
);