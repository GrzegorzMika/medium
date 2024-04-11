CREATE OR REPLACE FUNCTION fn_new_signals() RETURNS TRIGGER AS 
$$
BEGIN
    PERFORM pg_notify(
        'new_signals',
        to_json(NEW)::TEXT
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER new_signals
AFTER INSERT ON signals
FOR EACH ROW EXECUTE PROCEDURE fn_new_signals();