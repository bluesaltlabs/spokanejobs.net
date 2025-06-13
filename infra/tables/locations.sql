create table locations (
  id uuid not null primary key,
  name text not null,
  identifier text not null,
  company_id uuid not null references companies(id),
  address text,
  city text,
  state us_state,
  postal_code text,
  latitude float,
  longitude float,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
alter table locations enable row level security;

-- make the locations table data publicly readable
create policy "public can read locations"
on public.locations
for select to anon
using (true);
