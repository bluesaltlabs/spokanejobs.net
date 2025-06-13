create table locations (
  id uuid not null primary key default uuid_generate_v4(),
  name varchar(255) not null,
  identifier varchar(255) not null,
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

-- todo: set up a relationship for locations, organizations, events, etc to reference a location.
