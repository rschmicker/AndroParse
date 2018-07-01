#!/usr/bin/env python

import os
import sys
import time

if len(sys.argv) < 2:
	print("gplay.py <file of package names> <download directory>")
	sys.exit(1)

dl_loc = sys.argv[2]
package_file = sys.argv[1]
package_list = []

file = open(package_file, "r")
for line in file:
	package_list.append(line.strip("\n"))

counter = 0
for package_name in package_list:
	cmd = "gplaycli -d " + package_name + " -f " + dl_loc + " -p"
	print(cmd)
	os.system(cmd)
	print("Downloaded: " + package_name)
	counter += 1
	print(str(float(counter)/float(len(package_list)) * 100.0) + "%")
	#time.sleep(5)