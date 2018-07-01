
# http://ymsir.com/papers/pmds-iciss.pdf
# Pmds: Permission-based malware detection system

import json
from os import listdir
from os.path import isfile, join
from constants import PERMISSIONS
from sklearn import svm
from sklearn import tree
from sklearn import linear_model
from sklearn.naive_bayes import GaussianNB
from sklearn.cluster import KMeans
from sklearn.model_selection import KFold
from sklearn.tree import DecisionTreeClassifier
from sklearn.ensemble import RandomForestClassifier

max_count = 10000

# dl: 50 minutes for perms, filesize, and malicious for 94,000 apks, 12.9KB/s
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

def get_mal_split_benign(apks):
	mals = []
	benign = []
	for apk in apks:
		if apk['Malicious'] == 'true': # and len(mals) < max_count:
			mals.append(apk)
		else:
			benign.append(apk)
		# elif len(benign) < max_count:
		# 	benign.append(apk)

		# if len(mals) > max_count and len(benign) > max_count:
		# 	break

	return mals, benign

def main():
	models = {}
	models["SimpleLogistic"] = linear_model.LogisticRegression() # simple logistic
	# models["SMO"] = svm.SVC() # SMO
	# models["NaiveBayes"] = GaussianNB() # naive bayes
	models["RandomTree"] = DecisionTreeClassifier(random_state=0) # random tree
	models["RandomForest10"] = RandomForestClassifier(max_depth=10, random_state=0) # random forest, max depth: 10
	models["RandomForest50"] = RandomForestClassifier(max_depth=50, random_state=0) # random forest, max depth: 50
	models["RandomForest100"] = RandomForestClassifier(max_depth=100, random_state=0) # random forest, max depth: 100

	apks = get_apk_json("permsfilesize.json")
	feature_vector = []
	target_vector = []
	apks = apks['data']
	counts = [100, 500, 1000, 5000, 10000, 50000, len(apks)]

	mals, benigns = get_mal_split_benign(apks)

	# for i in range(1, 10):
	# 	print("===============================")
	# 	mal_idx = int(float(float(i)/float(10))*float(max_count))
	# 	ben_idx = int(float(abs(float(i)-float(10))/float(10))*float(max_count))
	# 	print("Ratio: (" + str(mal_idx) + "/" + str(ben_idx) + ")")
	# 	temp_mals = mals[:mal_idx]
	# 	temp_benigns = benigns[:ben_idx]
	# 	round_apks = temp_mals
	# 	for apk in temp_benigns:
	# 		round_apks.append(apk)
	# 	for model_name, model in models.iteritems():
	# 		kf = KFold(10, True, None)
	# 		kf.get_n_splits(round_apks)
	# 		train_apks = []
	# 		predict_apks = []
	# 		for train, test in kf.split(round_apks):
	# 			train_apks = train
	# 			predict_apks = test
	# 			break
	# 		for idx in train_apks:
	# 			feature_vector.append(get_permissions(round_apks[idx]))
	# 			target_type = 1 if round_apks[idx]['Malicious'] == 'true' else 0
	# 			target_vector.append(target_type)
	# 		clf = model
	# 		clf.fit(feature_vector, target_vector)

	# 		total_test = len(predict_apks)
	# 		number_correct = 0
	# 		for idx in predict_apks:
	# 				test_feature_vector = get_permissions(round_apks[idx])
	# 				result = clf.predict([test_feature_vector])
	# 				mal_status = 1 if round_apks[idx]['Malicious'] == 'true' else 0
	# 				if result == mal_status:
	# 						number_correct += 1
	# 		percent = (float(number_correct)/float(total_test))*float(100)
	# 		print("Model: " + model_name)
	# 		print("Accuracy: " + str(percent) + "%")

	# print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	for count in counts:
		print("===========================================")
		print("Count: " + str(count))
		mal_idx = int(float(float(1)/float(10))*float(count))
		ben_idx = int(float(abs(float(9)/float(10))*float(count)))
		print("Ratio: (" + str(mal_idx) + "/" + str(ben_idx) + ")")
		temp_mals = mals[:mal_idx]
		temp_benigns = benigns[:ben_idx]
		round_apks = temp_mals
		for apk in temp_benigns:
	 		round_apks.append(apk)
		for model_name, model in models.iteritems():
			kf = KFold(10, True, None)
			kf.get_n_splits(round_apks)
			train_apks = []
			predict_apks = []
			for train, test in kf.split(round_apks):
				train_apks = train
				predict_apks = test
				break
			for idx in train_apks:
				feature_vector.append(get_permissions(round_apks[idx]))
				target_type = 1 if round_apks[idx]['Malicious'] == 'true' else 0
				target_vector.append(target_type)
			clf = model
			clf.fit(feature_vector, target_vector)

			total_test = len(predict_apks)
			number_correct = 0
			for idx in predict_apks:
					test_feature_vector = get_permissions(round_apks[idx])
					result = clf.predict([test_feature_vector])
					mal_status = 1 if round_apks[idx]['Malicious'] == 'true' else 0
					if result == mal_status:
							number_correct += 1
			percent = (float(number_correct)/float(total_test))*float(100)
			print("Model: " + model_name)
			print("Accuracy: " + str(percent) + "%")

main()
