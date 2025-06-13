create table contacts (
  id uuid not null primary key default uuid_generate_v4(),
  first_name varchar(255) not null,
  last_name varchar(255) not null,
  email varchar(255) not null,
  phone varchar(255) not null,
  location_id uuid,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  deleted_at timestamp,
  primary key (id)
);
alter table contacts enable row level security;

-- make the contacts table data publicly readable
create policy "public can read contacts"
on public.contacts
for select to anon
using (true);
