
import json
with open('/home/mkleina/ceph_dump.json') as data_file:
    j = json.load(data_file)

def find(element, JSON, path, all_paths):
  if element in JSON:
    path = path + element + ' = ' + str(JSON[element]).encode('utf-8')
    print path
    all_paths.append(path)
  for key in JSON:
    if isinstance(JSON[key], dict):
      find(element, JSON[key],path + key + '.',all_paths)

all_paths = []
print j
find('msgr_created_connections', j, '', all_paths)
