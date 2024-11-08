package queries

const MonitorRunCreateQuery = `INSERT INTO monitor_runs (
                          monitor_id, 
                          status, 
                          dns_started_at, 
                          dns_ended_at, 	
                          tls_handshake_started_at, 
                          tls_handshake_ended_at, 
                          connection_started_at, 
                          connection_ended_at, 
                          latency,  
                          ran_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

const MonitorRunGetTimeSeries90dQuery = `WITH dates as ( 
    SELECT date_trunc('day', dd)::timestamp as key
    from generate_series(now() - '90 days'::interval, now(), '1 day') dd
), actives as (
    SELECT date_trunc('day', ran_at)::timestamp as key, count(*) as value FROM monitor_runs
    WHERE monitor_id = $1 and ran_at > (SELECT min(key) from dates) and status = 'active' group by 1
), totals as (
    SELECT date_trunc('day', ran_at)::timestamp as key, count(*) as value  FROM monitor_runs
    WHERE monitor_id = $1 and ran_at > (SELECT min(key) from dates) group by 1
)
SELECT d.key, COALESCE(actives.value, 0) as ok, COALESCE(totals.value, 0) as total, COALESCE(ROUND((actives.value::numeric / totals.value::numeric) * 10000) / 100, 0) as uptime FROM dates d
LEFT JOIN actives ON d.key = actives.key
LEFT JOIN totals ON d.key = totals.key
`
