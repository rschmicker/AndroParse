#!/usr/bin/env python
import urllib
import sys
import hashlib
import os
import time

if len(sys.argv) < 2:
		print("apkmirror.py <download location>")
		sys.exit(1)

dl_loc = sys.argv[1] + "/"
base_url = "https://www.apkmirror.com/wp-content/themes/APKMirror/download.php?id="
for i in range(400000):
		time.sleep(5)
		urllib.urlretrieve (base_url + str(i), dl_loc + "temp.apk")
		filehash = hashlib.sha256(open(dl_loc + "temp.apk", 'rb').read()).hexdigest()
		if filehash == "2f0040910db221dd7ca82931d2c5b0a27877d4badb4e57b41a4849ebbde0085a":
				continue
		os.rename(dl_loc + "temp.apk", dl_loc + filehash + ".apk")
		print("Downloaded: " + filehash + ".apk")
		time.sleep(1)
