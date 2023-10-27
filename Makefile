# Define subdirectories
SUBDIRS = libprintHelloWorld use-libprintHelloWorld zgoprintHelloWorld

# Define the default target
all: $(SUBDIRS)

# Rule to run 'make' in each subdirectory
$(SUBDIRS):
	$(MAKE) -C $@

# Rule to clean in each subdirectory
clean:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir clean; \
	done

.PHONY: all $(SUBDIRS) clean
