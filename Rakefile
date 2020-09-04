# frozen_string_literal: true

desc 'Runs docker-compose to spin up dependencies'
task :up do
  sh 'docker-compose up --build -d'
end

desc 'Stops docker so it does not consume your computer resources'
task :down do
  sh 'docker-compose down'
end

namespace :app do
  desc 'Builds the app'
  task :build do
    sh 'go build'
  end

  desc 'Run built app'
  task :run do
    sh './eng_metrics'
  end

  desc 'Build and run app'
  task build_and_run: ['app:build', 'app:run']
end

desc 'Build and run app'
task app: ["app:build_and_run"]

namespace :db do
  namespace :migrate do
    desc 'Creates a new migration. Usage: `rake migrate:create[my_migration_name]`'
    task :create, [:name] do |t, args|
      name = args.name
      sh "tern new #{name} -m migrate"
    end

    desc 'Prints information about migration status'
    task :status do
      sh 'tern status -c migrate/tern.conf -m migrate'
    end

    desc 'Migrate up to latest version. Usage: `rake migrate:up (or simply: rake migrate)`'
    task :up do
      sh 'tern migrate -c migrate/tern.conf -m migrate'
    end

    desc 'Rolls back latest migration. Usage: `rake migrate:rollback'
    task :rollback do
      sh 'tern migrate -c migrate/tern.conf -m migrate -d -1'
    end
  end

  desc 'Same as db:migrate:up'
  task migrate: ['db:migrate:up']

  desc 'Open psql console'
  task :console do
    sh "docker-compose exec postgres bash -c 'psql $POSTGRES_DB -U $POSTGRES_USER'"
  end

  desc 'Nukes DB data'
  task :nuke do
    sh "docker-compose exec postgres bash -c 'psql $POSTGRES_DB -U $POSTGRES_USER -c \"delete from issues; delete from pull_requests; delete from repositories;\"'"
  end
end


task default: ["app"]
