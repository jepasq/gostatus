.PHONY: doc

doc:
	jsdoc --destination doc/js/ static/*
	cd src&&make doc         # goland doc
	@echo "Please open doc/js/index.html"
