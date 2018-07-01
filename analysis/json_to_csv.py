#!/usr/bin/env python
import sys
import json
from constants import PERMISSIONS

def get_perm_file():
		pfile = open("permissions.txt", "r")
		data = pfile.readlines()
		for i in range(len(data)):
				data[i] = data[i].replace('\n', '')
		return data

def get_permissions(apk):
		perms = []
		for permission in get_perm_file():
				status = "TRUE" if permission in apk['permissions'] else "FALSE"
				perms.append(status)
		return perms

if(len(sys.argv) < 3):
		print("Pass the filename and out as parameter")
		sys.exit(1)

data = {}
with open(sys.argv[1]) as json_data:
    data = json.load(json_data)

out = open(sys.argv[2], "w")
const_perms = ','.join(get_perm_file())
out.write(const_perms + ",Malicious\n")
data = data["data"]

for apk in data:
		perms = get_permissions(apk)
		mal = apk['Malicious']
		perms = ','.join(perms)
		out.write(perms + "," + mal + "\n")
out.close()

