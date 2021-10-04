create schema gobudget;
create extension "uuid-ossp";

create table gobudget.users(
    id UUID primary key default uuid_generate_v4(),
    username varchar(255) not null,
    email text not null unique,
    password text not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
    -- todo all user preferences
);

create table gobudget.bank_accounts(
    id UUID primary key default uuid_generate_v4(),
    user_id UUID not null references gobudget.users(id),
    name text not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create type envelope_type as enum('goal', 'monthly', 'yearly');

create table gobudget.envelopes(
    id UUID primary key default uuid_generate_v4(),
    bank_account_id UUID not null references gobudget.bank_accounts(id),
    name varchar(255) not null,
    "envelope_type" envelope_type not null,
    budgeted_amount decimal not null,
    current_amount decimal not null default 0,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create table gobudget.transactions(
    id UUID primary key default uuid_generate_v4(),
    payee varchar(255) not null,
    note varchar(1023),
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create table gobudget.transaction_envelopes(
    id UUID primary key default uuid_generate_v4(),
    transaction_id UUID references gobudget.transactions(id) not null,
    envelope_id UUID references gobudget.envelopes(id) not null,
    amount decimal not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);