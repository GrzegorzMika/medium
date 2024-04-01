psql -U postgres -d signals -f "/init_queries/signals_table.sql"
psql -U postgres -d signals -f "/init_queries/notification.sql"