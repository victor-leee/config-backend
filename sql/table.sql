create table service_tab (
                             id int auto_increment primary key comment 'auto increment pk, no usage',
                             name varchar(256) unique comment 'service name, must be unique'
)