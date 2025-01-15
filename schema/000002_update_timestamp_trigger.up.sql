CREATE OR REPLACE FUNCTION update_user_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = current_timestamp;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_user_updated_at();