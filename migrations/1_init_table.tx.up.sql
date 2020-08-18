start transaction;

create extension if not exists "uuid-ossp";

create table wallets
(
    id           uuid primary key default uuid_generate_v4(),
    cur          varchar(3) default 'btc',
    balance      float default 0
);

create table transactions
(
    id           uuid primary key default uuid_generate_v4(),
    sender       uuid not null references wallets (id),
    receiver     uuid not null references wallets (id),
    currency     varchar(3) not null,
    amount       float not null,
    commission   float not null,
    created_at   timestamptz default now()
);

/* generating sample data (wallets with balances) */
insert into wallets(cur, balance) select (
    case (random() * 1)::int
      WHEN 0 THEN 'btc'
      WHEN 1 THEN 'eth'
    end
  ), random() from generate_series(1, 100) seq;

commit;