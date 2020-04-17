"""
Generates a plantuml gantt chart of the milestone timelines in clubhouse
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
            completed=epic.get('completed', False)
            if completed:
               epic['start_tag'] = "<s>"
               epic['end_tag'] = "</s>"
            else:
               epic['start_tag'] = ""
               epic['end_tag'] = ""
            print "***_ %(start_tag)s%(name)s%(end_tag)s" % epic


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
response = requests.get(URL)
if response.status_code == 200:
    milestones = response.json()
    today = datetime.now()
    last = today + timedelta(LAST_DAYS)
    print "@startmindmap"
    proj_start = determine_project_start(milestones)
    start_date = proj_start.strftime("%Y/%m/%d")
    #print "project starts %s" % start_date
    #print "printscale weekly"
    #print "[Today] happens at %s" % today.strftime('%Y/%m/%d')
    print "* Milestones"
    half = len(milestones) / 2
    switch_sides = True
    count = 0
    for milestone in milestones:
        count += 1
        if milestone.get('completed', False):  continue
        if not milestone.get('started', False):
            continue
        print "** %(name)s" % milestone
        show_epics(milestone)
        if count > half and switch_sides:
            switch_sides = False
            print "\nleft side\n"
    print "@endmindmap"
