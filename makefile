APP       := NewsAD
TARGET    := StateReportServer
MFLAGS    :=
DFLAGS    :=
CONFIG    := client
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/tars/makefile.taf))

upload2test:
	/usr/local/app/dpatch --env test --moduleType taf --application $(APP) --server $(TARGET) --tgz $(TARGET).tgz --user tafa
upload:
	/usr/local/app/dpatch --env pre --moduleType taf --application $(APP) --server $(TARGET) --tgz $(TARGET).tgz --user tafa


upload2test:
	'/home/svn/rootzhang/dpatch' --env test --moduleType taf --application $(APP) --server $(TARGET) --tgz $(TARGET).tgz --user tafa
upload:
	'/home/svn/rootzhang/dpatch' --env pre --moduleType taf --application $(APP) --server $(TARGET) --tgz $(TARGET).tgz --user tafa
