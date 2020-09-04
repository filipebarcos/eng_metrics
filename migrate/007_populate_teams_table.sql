-- Write your migrate up statements here

insert into teams (name, members)
    values ('Platinum Falcons', '{"chohmann", "pytlesk4", "XVincentX", "karol-maciaszek", "rainum"}'),
            ('The A-Team', '{"casserni", "johnrazmus", "ssspear", "cappslock"}'),
            ('Void Crew', '{"wmhilton", "P0lip", "mallachari", "marcelltoth"}'),
            ('Enablers', '{"collinbachi", "zee-hussain"}'),
            ('Growth', '{"paulatulis", "falsaffa"}'),
            ('11Sigma', '{"karol-maciaszek", "P0lip", "mallachari", "marcelltoth", "chris-miaskowski", "mmiask"}'),
            ('QA', '{"StoplightDeb", "chris-miaskowski", "mmiask"}');
---- create above / drop below ----

delete from teams;
