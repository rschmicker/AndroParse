#!/usr/bin/env python

import sys
from bs4 import BeautifulSoup
import requests
import hashlib
import time
import shutil
import re
from datetime import date

dl_loc = ""

def get_soup(url):
	res = requests.get(url)
	return BeautifulSoup(res.content, "html.parser")

def download(url):
	try:
		response = requests.get(url, stream=True)
		response.raise_for_status()
		with open('temp10.apk', 'wb') as handle:
			for block in response.iter_content(1024):
				handle.write(block)
		filehash = hashlib.sha256(open("temp10.apk", 'rb').read()).hexdigest()
		shutil.move("temp10.apk", dl_loc + "/" + filehash + ".apk")
		print("Downloaded: " + str(url))
		time.sleep(10)
	except:
		return

def get_links(url, page_count):
	if page_count != 1:
		url += "/" + str(page_count) + "/"
	soup = get_soup(url)
	a_tags = soup.findAll("a", href=True)
	links = []
	for link in a_tags:
		if re.match(r'\/en\/packages\/[a-zA-Z]+[0-9a-zA-Z\.]+', link.get('href')):
			links.append(link.get('href'))
	return links

def get_all_links(url):
	link_builder = get_links(url, 1)
	for i in range(2, 47):
		temp = get_links(url, i)
		for link in temp:
			link_builder.append(link)
	time.sleep(2)
	return link_builder

def download_links(links):
	for link in links:
		url = "https://f-droid.org" + link
		soup = get_soup(url)
		a_tags = soup.findAll("a", href=True)
		for entry in a_tags:
			to_dl = entry.get('href')
			to_dl = to_dl.strip()
			if '.apk' in to_dl:
				if '.asc' not in to_dl:
					if 'https://f-droid.org/FDroid.apk' != to_dl:
						print("Downloading: " + to_dl)
						download(to_dl)

def main():
	base_url = "https://f-droid.org/en/packages/"
	if(len(sys.argv) < 2):
		print("Usage: fdroid.py <download dir>")
		sys.exit(1)
	global dl_loc
	dl_loc = sys.argv[1]
	links = get_all_links(base_url)
	download_links(links)
main()
