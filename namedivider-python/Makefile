
.PHONY: download-csv

download-csv:
	@mkdir -p work
	@pip3 install --no-deps -r requirements.txt -t work
	@cp -R work/namedivider/assets/* assets
	@rm -rR work
