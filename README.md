Timelogger
==========
[![Build Status](https://travis-ci.org/minhajuddin/timelogger.png)](https://travis-ci.org/minhajuddin/timelogger)

## Logger

  1. You can log a task simply by running `timelogger task description .....`


## Watcher
  1. Allows you to hook up a cron task which sends an email with all the
     analysis

## Analyzer

 Analyzes your timelogs and emails you the statistics at the end of
 every week/work. For this to work, we make a few assumptions:

  1. The first word /[a-zA-Z-_]/ is treated as the project name, e.g. `gl gocms`
     implies that gocms is the name of your project

  2. The second word /[a-zA-Z-_]/ is treated as the task in that project.
     e.g. `gl learn code-reading` implies a subproject of code-reading

  3. A task with two stars '**' or two colons '::' is considered to be a non
     work task. These tasks too can have a project and subproject.
     e.g. `gl break lunch`

  4. A task with '+' parses the projects between '+'
     e.g. `gl break + sleep` assigns the time to both the projects?
     Try to avoid these

  5. You are assumed to log *after* the task is finished, i.e. once you have
     spent time working on something. e.g. At the end of working on a coding
     session on timelogger you would run `gl timelogger dev`

## Statistics / Analysis
  1. Projects, dates and hours

        Project | Time spent in hours | Date
        TL      | 8.33                | 2012-10-23
        TL      | 6.33                | 2012-10-24
        Learn   | 8.3                 | 2012-10-24
        Break   | 3                   | 2012-10-25

  2. Projects, tasks, hours and dates
      
        Project | Task         | Time spent in hours | Date
        TL      | dev          | 5.33                | 2012-10-23
        TL      | dev          | 3.33                | 2012-10-24
        TL      | design       | 3.33                | 2012-10-24
        Learn   | code-reading | 8.3                 | 2012-10-25

  3. Projects, tasks and hours

        Project | Task         | Time spent in hours
        TL      | dev          | 10.33
        TL      | design       | 3.33
        Learn   | code-reading | 8.3
        Break   | lunch        | 3

  4. Projects and hours

        Project | Time spent in hours
        TL      | 13.33
        Learn   | 8.3
        Break   | 3

## Filters
  We need to be able to filter the output by:
  
    1. Number of lines/or number of tasks: Use the -n flag
        `timelogger -n 33` shows the last 33 tasks
    2. Date
    3. Task regexp
    4. Project/Task/Subtask

  We should be able to combine these filters. We should also be able to select 
  the output independently
  
## Bash helpers
  You can add these in your ~/.bashrc to help with logging

~~~bash
  #shortcut to timelogger
  alias tl='timelogger'
  #to allow editing of your log file from your editor
  alias tle='vi ~/.timelog.txt'
  #helper to extend last logged task:
  #TODO: this needs to be moved to the main program
  tla()
  {
    echo "$(date "+%Y-%m-%d %H:%M"): $(tail -1 ~/.timelog.txt | sed -e 's/^[0-9 :-]*//g')" >> ~/.timelog.txt
  }
~~~

## Syncing across computers
 Use Dropbox to sync this file across multiple computers
    
~~~bash
    $ touch $HOME/Dropbox/timelog.txt
    $ ln -s $HOME/Dropbox/timelog.txt $HOME/.timelog.txt
~~~
