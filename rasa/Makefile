run:
	rasa run -vv --enable-api --log-file=/tmp/rasa.log --cors="*"

train:	
	rasa train -v
	# Copy the trained model to GCS.
	# N.B currently models need to be stored in the root of the folder
	# https://rasa.com/docs/rasa/model-storage#load-model-from-cloud
	# TODO(jlewi): We could use an init container to stage the model to 
	# get around that limitation.
	gsutil cp models/* gs://issue-label-bot-dev_jlewi/