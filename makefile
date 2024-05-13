# Set the virtual environment directory
MODS_DIR = test

# Install dependencies
install:
	go build main.go
	go install daniel-le97/go-cli

# Activate the virtual environment
mkdirs:
	make clean
	mkdir -p $(MODS_DIR)
	mkdir -p $(MODS_DIR)/1
	cd $(MODS_DIR)/1 && bun init --yes
	cd ..
	mkdir -p $(MODS_DIR)/2
	cd $(MODS_DIR)/2 && bun init --yes

run:
	go run main.go

# Clean up
clean:
	rm -rf $(MODS_DIR)
