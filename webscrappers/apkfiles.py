#!/usr/bin/env python

import sys
from bs4 import BeautifulSoup
import requests
import hashlib
import time
import shutil
import random
import selenium
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
import urllib2

dl_loc = ""
base_url = "https://www.apkfiles.com"

def init_webdriver():
        driver = webdriver.PhantomJS()
        driver.set_window_size(1120, 550)
        return driver

def get_soup(url):
	res = requests.get(url)
	return BeautifulSoup(res.content, "html.parser")

def download(url):
	f = urllib2.urlopen(url)
	if f.geturl().find("/storage/") == -1:
		driver = init_webdriver()
		driver.get(url)
		temp_links = driver.find_elements_by_xpath("//a[@href]")
		for link in temp_links:
			if link.get_attribute("href").find("/download/") != -1:
				download(link.get_attribute("href"))
	else:
		response = requests.get(url, stream=True)
		response.raise_for_status()
		with open('temp6.apk', 'wb') as handle:
			for block in response.iter_content(1024):
				handle.write(block)
		filehash = hashlib.sha256(open("temp6.apk", 'rb').read()).hexdigest()
		shutil.move("temp6.apk", dl_loc + "/" + filehash + ".apk")
		print("Downloaded: " + str(url))
		time.sleep(10)

def crawl_site(url):
	time.sleep(random.randint(5, 10))
	global base_url
	soup = get_soup(url)
	a_tags = soup.findAll("a", href=True)
	for link in a_tags:
		link = str(link.get("href"))
		if link.startswith("/download/"):
			print("Download: " + base_url + link)
			download(base_url + link)
			return
		elif link.startswith("/cat/applications/") and len(link) > 18:
			link = link[18:]
			print("Cat App: " + url + link)
			crawl_site(url + link)
		elif link.startswith("/cat/games/") and len(link) > 11:
			link = link[11:]
			print("Cat Game: " + url + link)
                        crawl_site(url + link)
		elif link.startswith("/cat/"):
			link = link[5:]
			print("Cat: " + url + link)
                        crawl_site(url + link)
		elif link.startswith("/apk-"):
			print("Apk: " + base_url + link)
			crawl_site(base_url + link)

def main():
	base_url = "https://www.apkfiles.com/cat/"
	if(len(sys.argv) < 2):
		print("Usage: apkfiles.py <download dir>")
		sys.exit(1)
	global dl_loc
	dl_loc = sys.argv[1]
	crawl_site(base_url)
main()
