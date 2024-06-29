-- Crear la base de datos
DROP DATABASE IF EXISTS gymbro_db;
CREATE DATABASE IF NOT EXISTS gymbro_db;

-- Utilizar la base de datos recién creada
USE gymbro_db;
create table core_modules
(
    id          varchar(36)  not null,
    name        varchar(255) not null,
    description varchar(255) not null,
    icon        varchar(255) not null,
    enable      tinyint(1)   not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null,
    constraint core_modules_pk
        unique (id)
);

create table core_roles
(
    id          varchar(36)  not null
        primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    enable      tinyint(1)   not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null
)
    comment 'roles of user';

create table core_users
(
    id         varchar(36)  not null
        primary key,
    username   varchar(255) not null,
    password   varchar(255) not null,
    enable     tinyint(1)   not null,
    created_at timestamp    not null,
    deleted_at timestamp    null
);

create table core_user_roles
(
    id      varchar(36) not null
        primary key,
    user_id varchar(36) not null,
    role_id varchar(36) null,
    constraint user_roles_roles_id_fk
        foreign key (role_id) references core_roles (id),
    constraint user_roles_users_id_fk
        foreign key (user_id) references core_users (id)
)
    comment 'roles por usuario';

create table core_views
(
    id          varchar(36)  not null
        primary key,
    module_id   varchar(36)  not null,
    name        varchar(255) not null,
    description varchar(255) not null,
    icon        varchar(255) not null,
    path        varchar(255) null,
    enable      tinyint(1)   not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null,
    constraint core_views_core_modules_id_fk
        foreign key (module_id) references core_modules (id)
);

create table core_roles_views
(
    id      varchar(36) not null
        primary key,
    role_id varchar(36) not null,
    view_id varchar(36) not null,
    constraint roles_views_roles_id_fk
        foreign key (role_id) references core_roles (id),
    constraint roles_views_views_id_fk
        foreign key (view_id) references core_views (id)
);

INSERT INTO core_users (id, username, password, enable, created_at)
VALUES ('e29716df-3652-11ef-9217-0242ac130006', 'admin', 'admin', 1, NOW());
INSERT INTO core_users (id, username, password, enable, created_at)
VALUES ('e29716df-3652-11ef-9217-0242ac130007', 'user', 'user', 1, NOW());

