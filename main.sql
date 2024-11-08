DROP TABLE IF EXISTS monitor_runs;
DROP TABLE IF EXISTS incidents;
DROP TABLE IF EXISTS monitor_status_pages;
DROP TABLE IF EXISTS status_pages;
DROP TABLE IF EXISTS monitor_notifications;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS monitors;

CREATE TABLE monitors(
                         id bigserial primary key not null,
                         title varchar not null,
                         description varchar,
                         status varchar not null default 'active',
                         periodicity varchar not null,
                         url varchar not null,
                         method varchar not null,
                         headers varchar,
                         body varchar,
                         degraded_after integer,
                         created_at timestamp not null default now(),
                         updated_at timestamp not null default now()
);

CREATE TABLE monitor_runs(
                             id bigserial primary key not null,
                             monitor_id bigint not null,
                             status varchar not null,
                             dns_started_at timestamp not null,
                             dns_ended_at timestamp not null,
                             tls_handshake_started_at timestamp not null,
                             tls_handshake_ended_at timestamp not null,
                             connection_started_at timestamp not null,
                             connection_ended_at timestamp not null,
                             latency bigint not null,
                             created_at timestamp not null default now(),
                             ran_at timestamp not null default now(),
                             foreign key (monitor_id) references monitors(id)
);

CREATE TABLE incidents (
                           id bigserial primary key not null,
                           monitor_id bigint not null,
                           title varchar not null,
                           description varchar,
                           status varchar not null,
                           started_at timestamp not null default now(),
                           resolved_at timestamp,
                           resolved_by varchar,
                           acknowledged_at timestamp,
                           acknowledged_by varchar,
                           foreign key (monitor_id) references monitors(id)
);

CREATE TABLE notifications (
                               id bigserial primary key not null,
                               title varchar not null,
                               description varchar,
                               provider varchar not null,
                               provider_data jsonb,
                               created_at timestamp not null default now(),
                               updated_at timestamp not null default now()
);

CREATE TABLE monitor_notifications(
                                      monitor_id bigint not null,
                                      notification_id bigint not null,
                                      assigned_at timestamp not null default now(),
                                      foreign key (monitor_id) references monitors(id),
                                      foreign key (notification_id) references notifications(id),
                                      constraint monitor_notifications_pk primary key (monitor_id, notification_id)
);

CREATE TABLE status_pages (
                              id bigserial primary key not null,
                              title varchar not null,
                              description varchar,
                              created_at timestamp not null default now(),
                              updated_at timestamp not null default now()
);

CREATE TABLE monitor_status_pages (
                                      monitor_id bigint not null,
                                      status_page_id bigint not null,
                                      assigned_at timestamp not null default now(),
                                      foreign key (monitor_id) references monitors(id),
                                      foreign key (status_page_id) references status_pages(id),
                                      constraint monitor_status_pages_pk primary key (monitor_id, status_page_id)
)