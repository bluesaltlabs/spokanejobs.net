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

-- todo: cascade delete when job is deleted
--
