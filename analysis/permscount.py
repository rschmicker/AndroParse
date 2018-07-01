import json

def get_apk_json(filepath):
	d = {}
	with open(filepath) as json_data:
		d = json.load(json_data)
	return d

def get_permissions(apk):
	perms = []
	for permission in PERMISSIONS:
		status = 1 if permission in apk['permissions'] else 0
		perms.append(status)
	return perms

permission_count = 0
apks = get_apk_json("permsfilesize.json")
apks = apks['data']
for apk in apks:
	perms = get_permissions(apk)
	permission_count += perms
