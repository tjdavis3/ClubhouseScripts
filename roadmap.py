"""
Prints the clubhouse milestones in markdown by quarter for the current year
"""


import requests
import dateutil.parser
from datetime import datetime, timedelta
from environs import Env

env = Env()
env.read_env()

APITOKEN=env("APITOKEN")


MILESTONES="https://api.clubhouse.io/api/v3/milestones"
tokenurl = "?token=%s" % APITOKEN
EPICS=MILESTONES + "/%d/epics"

year = datetime.now().year
roadmap = dict(q1=[], q2=[], q3=[], q4=[], other=[])

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
    if milestone.get('starts', None) != None:
        start = dateutil.parser.parse(milestone.get('starts'))
        milestone['start_date'] = start
        start_date = start.strftime("%Y/%m/%d")
        if start.date() <= last.date():
            print "[%s] as [%s_%s] starts %s " % (milestone.get('name'), milestone.get('entity_type'), milestone.get('id'), start_date)
            if milestone.get('ends', None) != None:
                end = dateutil.parser.parse(milestone.get('ends'))
                milestone['end_date'] = end
                if not milestone.get("completed", False) and end.date() < today.date():
                    end = today.date()
                print "[%s] ends %s" % (milestone.get('name'), end.strftime('%Y/%m/%d'))

def determine_color(milestone):
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

    print "[%s] as [%s_%s] is colored in %s\n" % (milestone.get('name'), milestone.get('entity_type'), milestone.get('id'), color)

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



URL = MILESTONES + tokenurl
print "# %s Roadmap" % year
response = requests.get(URL)
if response.status_code == 200:
    milestones = response.json()
    today = datetime.now()
    for milestone in milestones:
        ends = get_date(milestone, 'completed_at')
        quarter = None
        if ends != None:
            end_date = dateutil.parser.parse(ends)
            if end_date.year == year:
                if end_date.month < 4:
                    quarter = 'q1'
                elif end_date.month < 7:
                    quarter = 'q2'
                elif end_date.month < 11:
                    quarter = 'q3'
                else:
                    quarter = 'q4'
        else:
            quarter = 'other'
        if quarter != None:
            roadmap[quarter].append(milestone.get('name', 'unknown'))
    quarters = roadmap.keys()
    quarters.sort()
    for quarter in quarters:
        print "## %s" % quarter
        for item in roadmap.get(quarter, []):
            print "   * %s" % item
        print ""

        # show_epics(milestone)
