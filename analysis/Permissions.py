import urllib2

def get_arff():
	to_return = "@RELATION AndroidPermissions\n"
	to_return += "@ATTRIBUTE HASPERMISSION {TRUE,FALSE}\n"
	header, data = get_solr()
	permissions = get_permissions()
        to_return += "@ATTRIBUTE Malicious {TRUE,FALSE}\n"
        to_return += "@DATA\n"
        for i in range(len(permissions)):
                permissions[i] = permissions[i].strip("\n")
        for apk in data:
                apk_data = apk.split(",")
                is_mal = 1 if apk_data[0] == "true" else 0
                apk_data = apk_data[1:]
                for perm in permissions:
                        if perm in apk_data:
                                to_return += "TRUE,"
                        else:
                                to_return += "FALSE,"
                if is_mal == 1:
                        to_return += "TRUE\n"
                else:
                        to_return += "FALSE\n"

        return to_return

def get_solr():
        url = "https://64.251.61.74/?q=*:*&fl=Malicious,Permissions&rows=10000&wt=csv"
        req = urllib2.Request(url)
        myssl = ssl.create_default_context();
	myssl.check_hostname=False
	myssl.verify_mode=ssl.CERT_NONE
	res = urllib2.urlopen(req, context=myssl)
        data = res.read()
        data = data.split("\n")
        header = data[0]
        data = data[1:]
        return (header, data)

def get_permissions():
        pfile = open("permissions-all.txt", "r")
        return pfile.readlines()

print get_arff()
