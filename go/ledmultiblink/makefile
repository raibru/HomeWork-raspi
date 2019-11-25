#
# Simple build and lint environment
#

.SUFFIXES	:
.SUFFIXES	:	.go

#
# print colored output
RESET_COLOR    = \033[0m
make_std_color = \033[3$1m      # defined for 1 through 7
make_color     = \033[38;5;$1m  # defined for 1 through 255
OK_COLOR       = $(strip $(call make_std_color,2))
WRN_COLOR      = $(strip $(call make_std_color,3))
ERR_COLOR      = $(strip $(call make_std_color,1))
STD_COLOR      = $(strip $(call make_color,8))

COLOR_OUTPUT = 2>&1 |                                        \
    while IFS='' read -r line; do                            \
        if  [[ $$line == FAIL* ]]; then                      \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == *:[\ ]FAIL:* ]]; then              \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == [\-][\-][\-][\ ]FAIL:* ]]; then    \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == WARN* ]]; then                     \
            echo -e "$(WRN_COLOR)$${line}$(RESET_COLOR)";    \
        elif [[ $$line == PASS ]]; then                       \
            echo -e "$(OK_COLOR)$${line}$(RESET_COLOR)";     \
        elif [[ $$line == [\-][\-][\-][\ ]PASS:* ]]; then    \
            echo -e "$(OK_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == ok* ]]; then                       \
            echo -e "$(OK_COLOR)$${line}$(RESET_COLOR)";     \
        else                                                 \
            echo -e "$(STD_COLOR)$${line}$(RESET_COLOR)";    \
        fi;                                                  \
    done; exit $${PIPESTATUS[0]};

.DEFAULT: $(help)

SHELL           = /bin/bash
#GO              = /usr/local/go/bin/go
GO              = go
#GO_PREFIX       = GOCACHE=off
GO_PREFIX       = 
GO_FLAGS        = -v

OUTFILE         = mblink
MAIN_FILE       = $(OUTFILE).go
VER_FILE        = cmd/version.go
VER_FILE_TMP    = cmd/version.tmp
TEST_FILES      = $(wildcard *_test.go)
TEST_FILES     += $(wildcard test/*.go)
SRCS            = $(filter-out $(wildcard *_test.go), $(wildcard *.go))
#SRCS_TEST       = $(wildcard *_test.go)
SRCS_TEST       = ./...

CLEAN_FILES 		=                     \
									tags                \
									$(wildcard ./tmp/*) \
									$(VER_FILE_TMP)     \
									$(OUTFILE)

help:
	-@echo "Makefile with following options (make <option>):"
	-@echo "	clean"
	-@echo "	tdd"
	-@echo "	test"
	-@echo "	test_cache"
	-@echo "	test_verbose"
	-@echo "	test_cover (needs path ./tmp)"
	-@echo "	ctags"
	-@echo "	build (curent os)"
	-@echo "	build_win"
	-@echo "	build_linux"
	-@echo "	build_rpi"
	-@echo "	deploy"
	-@echo "	run"
	-@echo "    (*) not implemented"
	-@echo ""

print:
	-@echo "SRCS       ==> [$(SRCS)]"
	-@echo "SRCS_TEST  ==> [$(SRCS_TEST)]"
	-@echo "TEST_FILES ==> [$(TEST_FILES)]"

clean:
	$(GO) clean
	rm -f $(CLEAN_FILES)

tdd:
	@$(GO) test ./... $(COLOR_OUTPUT)

test: $(TEST_FILES) $(SRCS)
	@$(GO) test ./... $(COLOR_OUTPUT)

test_cache: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test $(SRCS_TEST) $(COLOR_OUTPUT)

test_verbose: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -v $(SRCS_TEST) $(COLOR_OUTPUT)

test_trace_view: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -trace ./tmp/trace.out $(COLOR_OUTPUT)
	@$(GO) tool trace ./tmp/trace.out

test_cover: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -coverprofile=./tmp/coverage.out ./... $(COLOR_OUTPUT)

run:
	@$(GO) run $(MAIN_FILE)

build:
	@$(GO) build $(MAIN_FILE)

build_win:
	@GOOS=windows GOARCH=amd64 $(GO) build $(MAIN_FILE)

build_linux:
	@GOOS=linux GOARCH=amd64 $(GO) build $(MAIN_FILE)

build_rpi:
	@sed -r \
				-e 's/(.*)(patch := ")([0-9]+)(".*$$)/echo "\1\2\\"$$((\3+1))\\"\4"/ge' \
				-e 's/(.*)(bdate := ")([0-9][0-9-]+)(".*$$)/echo "\1\2$$(date +\\"%Y-%m-%d-%H-%M-%S\\")\4"/ge' \
				$(VER_FILE) > $(VER_FILE_TMP)
	@cp -f $(VER_FILE_TMP) $(VER_FILE)
	@GOOS=linux GOARCH=arm GOARM=7 $(GO) build $(MAIN_FILE)

deploy:	build_rpi
	scp $(OUTFILE) pi@192.168.1.16:/opt/projects
	@ssh pi@192.168.1.16 /opt/projects/$(OUTFILE) -v

remote_check:
	@ssh pi@192.168.1.16 /opt/projects/$(OUTFILE) -v

remote_shutdown:
	ssh pi@192.168.1.16 sudo shutdown now >/dev/null

ctags:
	ctags -RV .

# EOF
