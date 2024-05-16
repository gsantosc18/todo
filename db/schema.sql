CREATE TABLE IF NOT EXISTS todo (
    id varchar(50),
    name varchar(250) not null,
    description text,
    done boolean
);