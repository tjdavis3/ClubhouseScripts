import requests
import sys
from environs import Env


env = Env()
env.read_env()

show_detail = env.bool("DETAIL",True)
APITOKEN=env("APITOKEN")

MILESTONES="https://api.clubhouse.io/api/v3/milestones"
URL=MILESTONES + "?token=%s" % APITOKEN
tokenurl = "?token=%s" % APITOKEN
EPICS=MILESTONES + "/%d/epics"
EPIC_URL = "https://api.clubhouse.io/api/v3/epic/%d"

def get_date(milestone, key):
    dt = milestone.get(key, None)
    if dt == None or dt == 'None':
        dt = milestone.get(key + '_override', '')
        if dt == None or dt == 'None':
            dt = 'unscheduled'
        else:
            dt = '(projected) ' + dt
    return dt

def show_epics(milestone):
    epic_response = requests.get((EPICS % milestone.get('id')) + tokenurl)
    if response.status_code == 200:
        epics = epic_response.json()
        if len(epics) > 0:
            print "#### Epics"
        for epic in epics:
            epic['starts'] = get_date(epic, 'started_at')
            epic['ends'] = get_date(epic, 'completed_at')
            if epic.get('ends') == None:
                epic['ends'] = get_date(epic, 'deadline')
            epic_name = "  * "
            if epic.get('completed', False):
                epic_name += "~~"
            epic_name += epic.get('name', '')
            if epic.get('completed', False):
                epic_name += "~~"
            print epic_name
        print ""
        for m_epic in epics:
            epicresp = requests.get(EPIC_URL % m_epic.get('id', 0) + APITOKEN)
            if epicresp.status_code == 200:
                epic = epicresp.json()
                epic['description'] = epic.get('description', '')
                print "##### %(name)s\n\n%(description)s" % epic
                print "\n"
    print ""
    with open('epics.json', 'w') as f:
        f.write(response.text)

def print_milestone(milestone):
    milestone['summary'] = milestone.get("description", "").split('\n')[0]
    if show_detail:
        print "### %(name)s " % milestone
        print "**start %s**" % get_date(milestone, 'started_at')
    #print "  > %s\n" % milestone.get("description", "").split('\n')[0]
        print "\n  %s\n" % milestone.get("description", "")
    #print "| %(name)s | %(summary)s |" % milestone
        show_epics(milestone)
    else:
        print "%(name)s,\"%(summary)s\"" % milestone

def insert(fname):
    try:
        with open('%s.inc' % fname) as fd:
            print(fd.read())
    except:
        pass

def print_heading(heading):
    print "## %s" % heading
    #print "| Project | Summary |"
    #print "| ---------| ----------- |"
    if not show_detail:
        print "Project, Description"

only = "0"
if len(sys.argv) > 1:
    only = sys.argv[1]
response = requests.get(URL)
if response.status_code == 200:
    milestones = response.json()
    if show_detail:
        insert('milestones')
    print_heading("Current Projects")
    for milestone in milestones:
        if only > "0" and milestone.get('id', 0) != int(only):
            continue
        if milestone.get("state", "") == "in progress":
            print_milestone(milestone)
    print "\n"
    print_heading("Backlogged Projects")
    for milestone in milestones:
        if only > "0" and milestone.get('id', 0) != int(only):
            continue
        if milestone.get("state", "") == "to do":
            print_milestone(milestone)
with open('milestones.json', 'w') as f:
    f.write(response.text)
