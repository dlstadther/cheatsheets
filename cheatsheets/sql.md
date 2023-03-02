# SQL

Style Guide
* https://github.com/dlstadther/sql-style-guide/

## Order of Operations

| Operation | Description                       |
|-----------|-----------------------------------|
| from      | source locations to get base data |
| where     | filters base data                 |
| group by  | aggregates base data              |
| having    | filters aggregated data           |
| select    | returns final data                |
| order by  | sorts final data                  |
| limit     | reduces final data to a row count |

## General


### Create daily (or DOM) snapshots of change data

Given a dataset which generates at a non-daily rate, create a view which displays the state as of each end-of-month (EOM) date.

Use of `generate_series()` is specific to postgres and various sql dialects may or may not have similar functionality.

```postgresql
with
-- list of all dates between range
static_dates as (
    SELECT date_trunc('day', dd):: date as the_date
    FROM generate_series
             ( '2022-01-01'::timestamp
             , '2022-08-31'::timestamp
             , '1 day'::interval) dd
)
-- list of EOM dates from "static_dates"
, eom_dates as (
    select cast(the_date - interval '1 day' as date) as the_date from static_dates where extract('day' from the_date) = 1
)
-- fake data to mock a dataset which exists multi-times per month, or missing months
, data as (
    select 'abc123' as client_id, 'x' as row_id, '1' as val, cast('2022-03-01' as date) as exec_date union all -- Mar, Apr
    select 'abc123' as client_id, 'y' as row_id, '2' as val, cast('2022-05-01' as date) as exec_date union all -- none
    select 'abc123' as client_id, 'z' as row_id, '3' as val, cast('2022-05-02' as date) as exec_date union all -- May
    select 'abc123' as client_id, 'a' as row_id, '4' as val, cast('2022-06-30' as date) as exec_date           -- Jun
)
-- create date range for when "data" rows are relevant
, data_range as (
    select
        client_id, row_id, val, exec_date
        , lead(exec_date) over(partition by client_id order by exec_date asc) as next_date
    from
        data
)
-- ensure that all "data_range" rows include a "next_date" (the last record per id will have a null "next_date")
, data_range_notnull as (
    select
        client_id, row_id, val, exec_date,
        case when next_date is null then current_date else next_date end as next_date
    from
        data_range
)
-- join data range to all dates between "exec_date" and "next_date"
-- filter to only EOM dates
select
    *
from
    data_range_notnull as d
left join
    static_dates as sd
on
    d.exec_date <= sd.the_date
    and d.next_date > sd.the_date
where
    sd.the_date in (select the_date from eom_dates)
```
