create table accounting_payment_media
(
    id         varchar(36)  not null
        primary key,
    name       varchar(255) not null,
    created_at timestamp    not null,
    deleted_at timestamp    null
);

create table accounting_payment_methods
(
    id          varchar(36)  not null
        primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null
);

create table accounting_payment_types
(
    id         varchar(36)  not null
        primary key,
    name       varchar(255) not null,
    created_at timestamp    not null,
    deleted_at timestamp    null
);

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
    id         varchar(36) not null
        primary key,
    user_id    varchar(36) not null,
    role_id    varchar(36) null,
    created_at timestamp   null,
    deleted_at timestamp   null,
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

create table core_role_views
(
    id         varchar(36) not null
        primary key,
    role_id    varchar(36) not null,
    view_id    varchar(36) not null,
    created_at timestamp   not null,
    deleted_at timestamp   null,
    constraint roles_views_roles_id_fk
        foreign key (role_id) references core_roles (id),
    constraint roles_views_views_id_fk
        foreign key (view_id) references core_views (id)
);

create table membership_attendance_statuses
(
    id          varchar(36)  not null
        primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    code        varchar(36)  not null,
    enable      tinyint(1)   not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null
);

create table membership_attendance_type
(
    id          varchar(36)  not null
        primary key,
    name        varchar(255) not null,
    description varchar(36)  not null,
    enable      tinyint(1)   not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null
);

create table membership_clients
(
    id         varchar(36)  not null
        primary key,
    names      varchar(255) not null,
    surname    varchar(255) not null,
    age        int          not null,
    gender     varchar(255) null,
    phone      varchar(255) null,
    email      varchar(36)  null,
    enable     tinyint(1)   not null comment 'Activo o inactivo',
    created_at timestamp    not null,
    deleted_at timestamp    null
);

create table membership_attendances
(
    id             varchar(36) not null
        primary key,
    client_id      varchar(36) not null,
    status_id      varchar(36) not null,
    date           date        not null,
    check_in_time  time        not null,
    check_out_time time        null,
    created_at     timestamp   not null,
    deleted_at     timestamp   null,
    constraint membership_attendances_membership_attendance_statuses_id_fk
        foreign key (status_id) references membership_attendance_statuses (id),
    constraint membership_attendances_membership_clients_id_fk
        foreign key (client_id) references membership_clients (id)
);

create table membership_durations
(
    id         varchar(36)  not null
        primary key,
    name       varchar(255) not null,
    duration   int          not null,
    enable     tinyint(1)   not null,
    created_at timestamp    not null,
    deleted_at timestamp    null
);

create table membership_promotions
(
    id              varchar(36)    not null
        primary key,
    name            varchar(255)   not null,
    discount_amount decimal(10, 2) not null,
    enable          tinyint(1)     not null,
    created_at      timestamp      not null,
    deleted_at      timestamp      null
);

create table membership_statuses
(
    id          varchar(36)  not null
        primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    code        varchar(255) not null,
    enable      tinyint(1)   not null,
    created_at  timestamp    not null,
    deleted_at  timestamp    null
);

create table membership_types
(
    id          varchar(36)    not null
        primary key,
    name        varchar(255)   not null,
    description varchar(255)   not null,
    price       decimal(10, 2) not null,
    enable      tinyint        not null,
    created_at  timestamp      not null,
    deleted_at  timestamp      null
);

create table memberships
(
    id                   varchar(36)    not null
        primary key,
    client_id            varchar(36)    not null comment 'Cliente',
    membership_type_id   varchar(36)    not null comment 'Tipo de membresia (dance, machines, etc)',
    attendance_type_id   varchar(36)    not null comment 'Tipo de asistencia (daily, every other day, weekly)',
    membership_status_id varchar(36)    not null comment 'Estado de membresia (PRESENT, ETC)',
    promotion_id         varchar(36)    not null comment 'Promocion (descuento)',
    duration_id          varchar(36)    not null comment 'Duracion',
    start_date           timestamp      not null,
    end_date             timestamp      not null,
    amount               decimal(10, 2) not null,
    total                decimal(10, 2) not null,
    discount_amount      decimal(10, 2) not null,
    created_at           timestamp      not null,
    deleted_at           timestamp      null,
    constraint membership_membership_membership_attendance_type_id_fk
        foreign key (attendance_type_id) references membership_attendance_type (id),
    constraint membership_membership_membership_clients_id_fk
        foreign key (client_id) references membership_clients (id),
    constraint membership_membership_membership_promotions_id_fk
        foreign key (promotion_id) references membership_promotions (id),
    constraint membership_membership_membership_statuses_id_fk
        foreign key (membership_status_id) references membership_statuses (id),
    constraint membership_membership_membership_types_id_fk
        foreign key (membership_type_id) references membership_types (id),
    constraint memberships_membership_durations_id_fk
        foreign key (duration_id) references membership_durations (id)
);

create table membership_payments
(
    id                varchar(36)    not null
        primary key,
    membeship_id      varchar(36)    not null,
    payment_media_id  varchar(36)    not null,
    payment_type_id   varchar(36)    not null,
    payment_method_id varchar(36)    null,
    date              datetime       not null,
    amount            decimal(10, 2) not null,
    discount_amount   decimal(10, 2) not null,
    created_at        timestamp      not null,
    deleted_at        timestamp      null,
    constraint membership_payments_accounting_payment_media_id_fk
        foreign key (payment_media_id) references accounting_payment_media (id),
    constraint membership_payments_accounting_payment_methods_id_fk
        foreign key (payment_method_id) references accounting_payment_methods (id),
    constraint membership_payments_accounting_payment_types_id_fk
        foreign key (payment_type_id) references accounting_payment_types (id),
    constraint membership_payments_memberships_id_fk
        foreign key (membeship_id) references memberships (id)
);

