CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

insert into roles (id, name, description)
values (uuid_generate_v4()::text, 'administrators', 'Elevated privileges');

insert into roles (id, name, description)
values (uuid_generate_v4()::text, 'users', 'Regular authenticated users');