create table service_tab (
                                 id int auto_increment primary key comment 'auto inc id',
                                 name varchar(256) comment 'service name, can duplicate under different hierarchy',
                                 service_key varchar(767) comment 'service auth key',
                                 is_service bool comment 'false indicates a directory entry',
                                 parent_id int comment 'refers to auto inc id',
                                 complete_path varchar(767) unique key comment 'the fully qualified service name',
                                 git_repo varchar(767) comment 'git repo url',
                                 build_file_rel_path varchar(767) comment 'build file(Dockerfile) relative path',
                                 app_type varchar(256) comment 'app type, can be scrpc / http',
                                 custom_port int comment 'custom port for the http app',
                                 prefix_mapping varchar(256) unique comment 'the prefix for the http app'
) comment 'rpc service meta tab'

create table service_config_key_tab (
    `id` int auto_increment primary key comment 'auto inc id',
    `service_id` varchar(767) comment 'refers to service_tab.complete_path',
    `key` varchar(767) comment 'config key'
) comment 'store config keys for each service'