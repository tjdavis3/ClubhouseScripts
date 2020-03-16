import requests
from environs import Env

env = Env()
env.read_env()

APITOKEN=env("APITOKEN")


URL="https://api.clubhouse.io/api/v3/search/stories?token=%s" % APITOKEN
MEMURL="https://api.clubhouse.io/api/v3/members?token=%s" % APITOKEN

members = {}

def run_query(url=URL, query={}):
    data = []
    response = requests.get(url, query)
    if response.status_code == 200:
        return response.json()

def load_members():
    members = {}
    for data in run_query(url=MEMURL):
        member = data.get("profile",  {})
        members[data.get('id', None)] = member.get('name','')
    return members

def get_stories(query):
    stories = run_query(url=URL, query=query)
    data = stories.get("data", [])
    for story in data:
        story['owner'] = members.get(story.get('owner_ids',[])[0], 'none')
        print "  * %(id)3d: %(name)-40s - %(updated_at)s / %(owner)s" % story

members = load_members()

print "## Stories In Development"
query = {"query": "state:\"In Development\""}
get_stories(query)
print "\n## Stories Awaiting Review"
query = {"query": "state:\"Ready for Review\""}
get_stories(query)
print "\n## Stories Waiting to be Deployed"
query = {"query": "state:\"Ready for Deploy\""}
get_stories(query)
