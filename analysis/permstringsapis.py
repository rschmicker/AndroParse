import json

def get_apk_json(filepath):
	d = {}
	with open(filepath) as json_data:
		d = json.load(json_data)
	return d

def get_permissions(apk):
	perms = []
	for permission in PERMISSIONS:
		status = 1 if permission in apk['Permissions'] else 0
		perms.append(status)
	return perms

def get_sys_cmds(apk):
	cmds = []
	for cmd in SYSTEM_COMMANDS:
		status = 1 if cmd in apk['Strings'] else 0
		cmds.append(status)
	return cmds

feature_vector = []
target_vector = []
apks = get_apk_json("permsfilesize.json")
apks = apks['data']
for apk in apks:
	vector = []
	vector.append(get_permissions(apk))
	vector.append(apk['Apis'])
	vector.append(get_sys_cmds(apk))
	feature_vector.append(vector)
	target_type = 1 if apk['Malicious'] == 'true' else 0
	target_vector.append(target_type)
