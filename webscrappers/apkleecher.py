#!/usr/bin/env python

import sys
from bs4 import BeautifulSoup
import requests
import hashlib
import time
import shutil
from datetime import date

dl_loc = ""

def get_soup(url):
	res = requests.get(url)
	return BeautifulSoup(res.content, "html.parser")

def download(url):
	try:
		response = requests.get(url, stream=True)
		response.raise_for_status()
		with open('temp2.apk', 'wb') as handle:
			for block in response.iter_content(1024):
				handle.write(block)
		filehash = hashlib.sha256(open("temp2.apk", 'rb').read()).hexdigest()
		shutil.move("temp2.apk", dl_loc + "/" + filehash + ".apk")
		print("Downloaded: " + str(url))
		time.sleep(10)
	except:
		return

def crawl_site(url):
	soup = get_soup(url)
	a_tags = soup.findAll("a", href=True)
	for link in a_tags:
		if 'Parent' in str(link):
			continue
		elif '.apk' in str(link):
			to_append = link.get('href')
			download(url + to_append)
		else:
			next_link = link.get('href')
			print("Moving to: " + url + next_link)
			crawl_site(url + next_link)

def main():
	base_url = "http://apkleecher.com/apps/2018/"
	if(len(sys.argv) < 2):
		print("Usage: apk-downloaders <download dir> <today>")
		sys.exit(1)
	global dl_loc
	dl_loc = sys.argv[1]
	try:
		access = sys.argv[2]
		today = date.today()
		if today.month < 10:
			month = "0" + str(today.month) + "/"
		else:
			month = str(today.month) + "/"
		if today.day < 10:
			day = "0" + str(today.day) + "/"
		else:
			day = str(today.day) + "/"
		base_url += month + day
		crawl_site(base_url)
	except:
		crawl_site(base_url)
main()
