alter table gobudget.envelopes drop column envelope_type;
alter table gobudget.envelopes add column envelope_interval smallint not null; --this fails if you have any envelopes set up
drop type envelope_type;

alter table gobudget.envelopes drop column bank_account_id cascade;
alter table gobudget.envelopes add column user_id uuid not null references gobudget.users(id);
alter table gobudget.transactions add column bank_account_id uuid not null references gobudget.bank_accounts(id);
alter table gobudget.bank_accounts add column current_balance decimal not null default 0;