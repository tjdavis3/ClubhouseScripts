"""
Generates a plantuml gantt chart of the milestone timelines in clubhouse
"""


import requests
import dateutil.parser
from datetime import datetime, timedelta
from environs import Env
import sys

env = Env()
env.read_env()

APITOKEN=env("APITOKEN")

MILESTONES="https://api.clubhouse.io/api/v3/milestones"
tokenurl = "?token=%s" % APITOKEN
EPICS=MILESTONES + "/%d/epics"
LAST_DAYS = 365

def get_date(milestone, key):
    dt = milestone.get(key, None)
    if dt == None or dt == 'None' or dt == '' or dt == 'null':
        dt = milestone.get(key + '_override', '')
        if dt == None or dt == 'None' or dt == 'null' or dt == '':
            dt = None
    return dt

def show_epics(milestone):
    epic_response = requests.get((EPICS % milestone.get('id')) + tokenurl)
    if response.status_code == 200:
        epics = epic_response.json()
        for epic in epics:
            epic['starts'] = get_date(epic, 'started_at')
            epic['ends'] = get_date(epic, 'completed_at')
            if epic.get('ends') == None:
                epic['ends'] = get_date(epic, 'deadline')
            show_dates(epic)
            determine_color(epic)


def show_dates(milestone):
    start = get_date(milestone, 'started_at')
  
    if start == None:
        if milestone.get('starts', None) != None:
            start = dateutil.parser.parse(milestone.get('starts'))
    else:
        start = dateutil.parser.parse(start)
    if start != None:
        milestone['start_date'] = start
        start_date = start.strftime("%Y/%m/%d")
        stag = etag = ""
        if milestone.get("completed", False):
            stag = "<s>"
            etag = "</s>"
        if start.date() <= last.date():
            print "[%s%s%s] as [%s_%s] starts %s " % (stag,milestone.get('name'),etag, milestone.get('entity_type'), milestone.get('id'), start_date)
            end = None
            if milestone.get('ends', None) != None:
                end = dateutil.parser.parse(milestone.get('ends'))
            elif milestone.get('deadline', None) != None:
                end = dateutil.parser.parse(milestone.get('deadline'))
            if end != None:
                milestone['end_date'] = end
                if not milestone.get("completed", False) and end.date() < today.date():
                    end = today.date()
                print "[%s%s%s] ends %s" % (stag,milestone.get('name'), etag, end.strftime('%Y/%m/%d'))
            else:
                stats = milestone.get('stats',{})
                duration = stats.get('num_points',0)
                if duration > 0:
                    print "[%s%s%s] lasts %d days" % (stag,milestone.get('name'), etag, duration)

def determine_color(milestone):
    stag = etag = ""
    if milestone.get("completed", False):
        stag = "<s>"
        etag = "</s>"
    if milestone.get("completed", False):
        color = "LightGrey"
    elif milestone.get('started', False):
        color = "LightGreen"
    elif (milestone.get('start_date', today).date() - today.date()).days < 0:
        color = "Pink"
    else:
        color = "LightBlue"
    # Check to see if the milestone is scheduled to be complete, but isn't
    end_date = milestone.get('end_date', None)
    if not milestone.get("completed", False) and end_date is not None and end_date.date() < today.date():
        color = "Magenta"

    print "[%s%s%s] as [%s_%s] is colored in %s\n" % (stag,milestone.get('name'),etag, milestone.get('entity_type'), milestone.get('id'), color)

def insert(fname):
    try:
        with open('gantt_%s.inc' % fname) as fd:
            print(fd.read())
    except:
        pass

def determine_project_start(milestones):
    start_date = today.date()
    for milestone in milestones:
        msstart = get_date(milestone, 'started_at')
        if msstart == None:
            continue
        start = dateutil.parser.parse(msstart)
        if start.date() < today.date():
            if milestone.get('completed', False):
                continue
        start_date = start
        break
    return start_date
        
   

single = False
if len(sys.argv) == 2:
    URL = MILESTONES + "/%s" % sys.argv[1] + tokenurl
    single = True
else:
    URL = MILESTONES + tokenurl
response = requests.get(URL)
if response.status_code == 200:
    milestones = response.json()
    if single:  milestones = [milestones]
    today = datetime.now()
    last = today + timedelta(LAST_DAYS)
    print "@startgantt"
    print "saturday are closed"
    print "sunday are closed"
    proj_start = determine_project_start(milestones)
    start_date = proj_start.strftime("%Y/%m/%d")
    print "project starts %s" % start_date
    print "printscale weekly"
    print "[Today] happens at %s" % today.strftime('%Y/%m/%d')
    insert('start')
    for milestone in milestones:
        milestone['starts'] = get_date(milestone, 'started_at')
        milestone['ends'] = get_date(milestone, 'completed_at')
        show_dates(milestone)
        determine_color(milestone)
        if single: show_epics(milestone)
        show_epics(milestone)
    insert('end')
    print "@endgantt"
