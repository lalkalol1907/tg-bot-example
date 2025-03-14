create table good
(
    id       uuid primary key not null default uuid_generate_v4(),
    name     varchar          not null,
    owner_id bigint           not null
);

create table tag
(
    id      uuid primary key not null default uuid_generate_v4(),
    text    varchar          not null,
    good_id uuid             not null,

    foreign key (good_id) references good (id)
);

create table chat
(
    id        bigint primary key not null,
    owner_id  bigint             not null,
);

create index concurrently if not exists good_owner_idx on good (owner_id);
