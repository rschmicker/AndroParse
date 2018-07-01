#!/usr/bin/env python3
import selenium
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
import time
import sys
import urllib
import hashlib
import os
import subprocess
import requests
import shutil

if len(sys.argv) < 2:
	print("evozi.py <file of package names> <download directory>")
	sys.exit(1)

dl_loc = sys.argv[2]
package_file = sys.argv[1]
package_list = []

file = open(package_file, "r")
for line in file:
	package_list.append(line)

url = "https://apps.evozi.com/apk-downloader/"

profile = webdriver.FirefoxProfile()
profile.set_preference('browser.download.folderList', 2) # custom location
profile.set_preference('browser.download.manager.showWhenStarting', False)
profile.set_preference('browser.download.dir', '/tmp/')
profile.set_preference('browser.helperApps.neverAsk.saveToDisk', '*')

#driver = webdriver.Firefox(profile)
driver = webdriver.PhantomJS()
driver.set_window_size(1120, 550)
counter = 0
for package_name in package_list:
	try:
		driver.get(url)
		time.sleep(30)
		print("Url: " + url)
		package_name_box = driver.find_element_by_class_name("form-control")
		package_name_box.send_keys(package_name)
		print("Package: " + package_name)
		submit_button = driver.find_element_by_class_name("btn").click()
		time.sleep(5)
		dl_button = driver.find_element_by_class_name("btn-success")
		link = dl_button.get_attribute("href")
		print("Link: " + link)
		response = requests.get(link, stream=True)
		response.raise_for_status()
		with open('temp.apk', 'wb') as handle:
		    for block in response.iter_content(1024):
		        handle.write(block)
		stats = os.stat("temp.apk")
		if (stats.st_size < (20 * 1024)):
		    time.sleep(30)
		    continue
		filehash = hashlib.sha256(open("temp.apk", 'rb').read()).hexdigest()
		shutil.move("temp.apk", dl_loc + "/" + filehash + ".apk")
		print("Downloaded: " + package_name)
		time.sleep(30)
		counter += 1
		print(str(float(counter)/float(len(package_list)) * 100.0) + "%")
	except:
		counter += 1
		continue
driver.quit()
