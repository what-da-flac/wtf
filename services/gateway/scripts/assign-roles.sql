delete from user_role ur
using users u
where u.id = ur.user_id
and u.email in ('mauricio.leyzaola@gmail.com', 'ru.leyzaola@gmail.com');

insert into user_role (user_id, role_id)
select u.id, r.id
from users u, roles r
where u.email in ('mauricio.leyzaola@gmail.com', 'ru.leyzaola@gmail.com');