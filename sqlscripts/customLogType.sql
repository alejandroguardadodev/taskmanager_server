DO
$$
BEGIN
  IF NOT EXISTS (SELECT *
                        FROM pg_type typ
                             INNER JOIN pg_namespace nsp
                                        ON nsp.oid = typ.typnamespace
                        WHERE nsp.nspname = current_schema()
                              AND typ.typname = 'custom_log_type') THEN
    CREATE TYPE custom_log_type AS ENUM ('ACCOUNT', 'CREATE', 'UPDATE', 'DELETE');
  END IF;
END;
$$
LANGUAGE plpgsql;