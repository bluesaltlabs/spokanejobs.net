create table categories (
  id bigint not null primary key,
  name varchar(255) not null unique,
  slug varchar(255) not null unique,
  description text,
  parent_id int references categories(id),
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
);
alter table categories enable row level security;

-- make the categories table data publicly readable
create policy "public can read categories"
on public.categories
for select to anon
using (true);
