select md5(concat(random()::text, clock_timestamp()::text));