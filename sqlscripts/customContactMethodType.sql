DO
$$
BEGIN
  IF NOT EXISTS (SELECT *
                        FROM pg_type typ
                             INNER JOIN pg_namespace nsp
                                        ON nsp.oid = typ.typnamespace
                        WHERE nsp.nspname = current_schema()
                              AND typ.typname = 'custom_contact_method_type') THEN
    CREATE TYPE custom_contact_method_type AS ENUM ('EMAIL', 'PHONENUMBER');
  END IF;
END;
$$
LANGUAGE plpgsql;