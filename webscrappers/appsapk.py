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
import re
import random

ignore_links = [
	"http://www.appsapk.com/latest-android-apps/",
	"http://www.appsapk.com/best-android-apps/",
	"http://www.appsapk.com/recently-updated/",
	"http://www.appsapk.com/android/all-apps/#sidr-main",
	"http://www.appsapk.com/",
	"http://www.appsapk.com/android/all-apps/",
	"http://www.appsapk.com/android/fun-games/",
	"http://www.appsapk.com/android/themes/",
	"http://www.appsapk.com/android/live-wallpaper/",
	"http://www.appsapk.com/android/launchers/",
	"http://www.appsapk.com/android-wallpapers/",
	"https://www.facebook.com/AppsApkcom",
	"https://www.twitter.com/AppsApkcom",
	"https://plus.google.com/+Appsapk",
	"http://www.appsapk.com/contact/",
	"http://www.appsapk.com/privacy-policy/",
	"http://www.appsapk.com/dmca/",
	"http://www.appsapk.com/terms-of-use/",
	"http://www.appsapk.com/submit-app/",
	"http://www.appsapk.com/sitemap_index.xml",
	"http://www.appsapk.com/apps/notes/",
	"http://www.appsapk.com/android-wallpapers/",
	"http://www.appsapk.com/android-wallpaper/",
]

dl_loc = ""

def init_webdriver():
	driver = webdriver.PhantomJS()
	driver.set_window_size(1120, 550)
	return driver
	#return webdriver.Chrome()

def good_link(link):
	global ignore_links
	good = True
	if not link.startswith("http://www.appsapk.com/"):
		good = False
	if link in ignore_links:
		good = False
	if link.find("/all-apps/#") != -1:
		good = False
	if link.find("/android/") != -1:
		good = False
	if link.find("/apps/") != -1:
		good = False
	return good	

def get_current_page_links(driver, url):
	global ignore_links
	link_list = []
	driver.get(url)
	time.sleep(2)
	print("Received link")
	main_div = driver.find_element_by_class_name('main-box-inside')
	temp_links = main_div.find_elements_by_xpath("//a[@href]")
	print("Links found: " + str(len(temp_links)))
	for link in temp_links:
		str_link = link.get_attribute("href")
		print("Current link: " + str_link)
		if good_link(str_link):
			print("Appending link: " + str_link)
			link_list.append(str_link)
			
	print("Good links: " + str(len(link_list)))
	return list(set(link_list))

def get_all_links(driver, url):
	print("Moving to: " + url)
	link_builder = get_current_page_links(driver, url)
	for i in range(2, 224):
		append_url = url + "page/" + str(i) + "/"
		print("Moving to: " + append_url)
		tmp_links = get_current_page_links(driver, append_url)
		for tmp in tmp_links:
			link_builder.append(tmp)
		time.sleep(5)
		print(link_builder)
		time.sleep(random.randint(5, 10))
	return link_builder

def download(url):
	global dl_loc
	response = requests.get(url, stream=True)
	response.raise_for_status()
	with open('temp5.apk', 'wb') as handle:
		for block in response.iter_content(1024):
			handle.write(block)
	filehash = hashlib.sha256(open("temp5.apk", 'rb').read()).hexdigest()
	shutil.move("temp5.apk", dl_loc + "/" + filehash + ".apk")
	print("Downloaded: " + str(url))
	time.sleep(10)

def iterate_links(driver, links):
	for link in links:
		driver.get(link)
		if str(requests.get(link).status_code) == '404' or str(requests.get(link).status_code) == '301':
			continue
		print("Currently at: " + link)
		time.sleep(5)
		try:
			dl_button = driver.find_element_by_class_name("download")
		except:
			dl_button = driver.find_element_by_class_name("apna-download")
		dl_link = dl_button.get_attribute("href")
		print("Download Btn: " + dl_link)
		if str(requests.get(link).status_code) == '404' or str(requests.get(link).status_code) == '301' or str(requests.get(link).status_code) == '502':
			continue
		driver.get(dl_link)
		time.sleep(5)
		try:
			dl_apk_btn = driver.find_element_by_id("download-file")
			dl_apk_link = dl_apk_btn.get_attribute("href")
			print("Real link: " + dl_apk_link)
			download(dl_apk_link)
		except:
			print("Real link: " + dl_link)
			download(dl_link)

def main():
	global dl_loc
	if len(sys.argv) < 2:
		print("appsapk.py <download directory>")
		sys.exit(1)

	dl_loc = sys.argv[1]
	base_url = "http://www.appsapk.com/android/all-apps/"
	driver = init_webdriver()
	all_links = get_all_links(driver, base_url)
	iterate_links(driver, all_links)
	driver.quit()

main()
