create table core_modules
(
    id          varchar(36)  not null,
    name        varchar(255) not null,
    description varchar(255) not null,
    number      int          not null,
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

create table core_role_modules
(
    id        varchar(36) not null
        primary key,
    role_id   varchar(36) not null,
    module_id varchar(36) not null,
    constraint core_role_modules_core_modules_id_fk
        foreign key (module_id) references core_modules (id),
    constraint core_role_modules_core_roles_id_fk
        foreign key (role_id) references core_roles (id)
);

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
    number      int          not null,
    icon        varchar(255) not null,
    url         varchar(255) not null,
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


-- Insertar datos de prueba iniciales
INSERT INTO core_users (id, username, password, enable, created_at)VALUES ('94d3dca2-f9da-11ee-811d-0242ac130002', 'afit@gym.com', '00000000', 1, NOW());
INSERT INTO core_user_roles (id, user_id, role_id) VALUES ('f4e169aa-080b-11ef-80f9-0242ac130002', '94d3dca2-f9da-11ee-811d-0242ac130002', '7f0463e0-bf21-42d7-b324-234cf60162ec');
INSERT INTO core_modules (id, name, description, number, icon, enable, created_at, deleted_at) VALUES ('0d9fe896-069b-11ef-a4ad-0242ac130002', 'Core', 'MOdulo 1', 1, 'ri ri-table', 1, '2024-04-29 21:42:05', null);
INSERT INTO core_roles (id, name, description, enable, created_at, deleted_at) VALUES ('7f0463e0-bf21-42d7-b324-234cf60162ec', 'jefe de gym', 'jefe de gym', 1, '2024-04-29 02:35:37', null);
INSERT INTO core_role_modules (id, role_id, module_id) VALUES ('155ea847-080c-11ef-80f9-0242ac130002', '7f0463e0-bf21-42d7-b324-234cf60162ec', '0d9fe896-069b-11ef-a4ad-0242ac130002');
INSERT INTO gymbro_db.core_views (id, module_id, name, description, number, icon, url, enable, created_at, deleted_at) VALUES ('3a208652-080c-11ef-80f9-0242ac130002', '0d9fe896-069b-11ef-a4ad-0242ac130002', 'Usuarios', 'Vista prueba', 1, 'fa fa-table', '/gym/users', 1, '2024-05-01 17:43:34', null);
INSERT INTO gymbro_db.core_views (id, module_id, name, description, number, icon, url, enable, created_at, deleted_at) VALUES ('4fc690cd-080c-11ef-80f9-0242ac130002', '0d9fe896-069b-11ef-a4ad-0242ac130002', 'Roles', 'Vista  2', 2, 'fa fa-table', '/gym/roles', 1, '2024-05-01 17:44:22', null);
