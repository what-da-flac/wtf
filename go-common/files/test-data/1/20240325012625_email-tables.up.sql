create table tenant (
    id text primary key ,
    name text not null,
    created timestamp with time zone not null
);

create table recipient (
    id text primary key ,
    created timestamp with time zone not null,
    tenant_id text not null,
    email text not null,
    enabled boolean not null
);

alter table recipient add constraint ix_recipient_tenant_email_unique unique (tenant_id, email);

create table message (
    id text primary key ,
    tenant_id text not null,
    sender text not null,
    recipient_id text not null,
    subject text not null,
    body text not null,
    html_body text not null,
    bcc jsonb not null,
    cc jsonb not null,
    created timestamp with time zone not null,
    status text not null
);

alter table message add constraint fk_message_tenant foreign key (tenant_id) references tenant(id);
alter table message add constraint fk_message_recipient foreign key (recipient_id) references recipient(id);