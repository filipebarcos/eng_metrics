-- Write your migrate up statements here

update teams set team_label = 'team/void-crew' where name = 'Void Crew';
update teams set team_label = 'team/enablers' where name = 'Enablers';
update teams set team_label = 'team/the-a-team' where name = 'The A-Team';
update teams set team_label = 'team/qa' where name = 'QA';
update teams set team_label = 'team/platinum-falcons' where name = 'Platinum Falcons';
update teams set team_label = 'team/growth' where name = 'Growth';
---- create above / drop below ----

update teams set team_label = null where 1 = 1;
