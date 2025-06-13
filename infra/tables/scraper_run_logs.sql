create table scraper_run_logs (
  id uuid not null default uuid_generate_v4(),
  job_count int not null default 0,
  status scraper_status not null default 'Pending',
  error_message text,
  output_message text,
  run_at timestamp not null default now(),
  completed_at timestamp
);
alter table scraper_run_logs enable row level security;

-- make the scraper_run_logs table data publicly readable
create policy "public can read scraper_run_logs"
on public.scraper_run_logs
for select to anon
using (true);
