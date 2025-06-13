create table locations (
  id uuid not null primary key,
  name varchar(255) not null,
  identifier varchar(255) not null,
  company_id bigint not null references companies(id),
  address varchar(255),
  city varchar(255),
  state us_state,
  postal_code varchar(32),
  latitude float,
  longitude float,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
);
alter table locations enable row level security;

-- make the locations table data publicly readable
create policy "public can read locations"
on public.locations
for select to anon
using (true);
