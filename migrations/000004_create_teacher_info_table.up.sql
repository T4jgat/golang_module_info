CREATE TABLE teacher_info (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp(0) with time zone not null  default now(),
    updated_at timestamp(0) with time zone not null  default now(),
    name varchar(50) not null,
    surname varchar(50) not null ,
    email varchar(100) not null unique ,
    module_info_fk int,
    foreign key (module_info_fk) references module_info(id)
)