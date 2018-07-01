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

feature_vector = []
target_vector = []
apks = get_apk_json("permsfilesize.json")
apks = apks['data']
for apk in apks:
	feature_vector.append(get_permissions(apk))
	target_type = 1 if round_apks[idx]['Malicious'] == 'true' else 0
	target_vector.append(target_type)
