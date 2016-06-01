#CMD DEFS
CP = cp
CP_R = cp -r

#Component name
COMP_NAME = asicd

#Set default build targets
BUILD_TARGET = cel_redstone

#Set paths for use during install
PARAMSDIR = $(DESTDIR)/params
DESTDIR = $(SR_CODE_BASE)/snaproute/src/out/bin

#IPC related vars
IPC_GEN_CMD = thrift
IPC_SRCS = rpc/asicd.thrift
IPC_SVC_NAME = asicd
GENERATED_IPC = $(SR_CODE_BASE)/generated/src

#Determine lib/bin paths based on build target
ifeq ($(BUILD_TARGET), cel_redstone)
	BCM_KMODS = bin/bcm/cel_redstone/kmod/*.ko
	BCM_LIBS = bin/bcm/cel_redstone/lib/libbcm.so.1
else ifeq ($(BUILD_TARGET), accton_as5712)
	BCM_KMODS = bin/bcm/accton_as5712/kmod/*.ko
	BCM_LIBS = bin/bcm/accton_as5712/lib/libbcm.so.1
else ifeq ($(BUILD_TARGET), accton_wedge)
	BCM_KMODS = bin/bcm/accton_wedge/kmod/*.ko
	BCM_LIBS = bin/bcm/accton_wedge/lib/libbcm.so.1
else
	#Default case use cel_redstone as bcm target
	BCM_KMODS = bin/bcm/cel_redstone/kmod/*.ko
	BCM_LIBS = bin/bcm/cel_redstone/lib/libbcm.so.1
endif
ifeq ($(BUILD_TARGET), mlnx)
	SAI_LIBS = bin/sai/mlnx/libs/*
else
	#Default case use mlnx as sai target
	SAI_LIBS = bin/sai/mlnx/libs/*
endif
ifeq ($(BUILD_TARGET), cel_redstone)
	ASICD_BIN = bin/bcm/cel_redstone/asicd
else ifeq ($(BUILD_TARGET), accton_as5712)
	ASICD_BIN = bin/bcm/accton_as5712/asicd
else ifeq ($(BUILD_TARGET), accton_wedge)
	ASICD_BIN = bin/bcm/accton_wedge/asicd
else ifeq ($(BUILD_TARGET), mlnx)
	ASICD_BIN = bin/sai/mlnx/asicd
endif

#TARGETS
all:ipc

exe:
	echo "ASICd - precompiled binaries available"
	$(CP) $(ASICD_BIN) $(DESTDIR)/$(COMP_NAME)

ipc:
	$(IPC_GEN_CMD) -r --gen go -out $(GENERATED_IPC) $(IPC_SRCS)

install:
	install params/asicd.conf $(PARAMSDIR)/
	$(CP_R) $(BCM_KMODS) $(DESTDIR)/kmod/
	$(CP_R) $(BCM_LIBS) $(DESTDIR)/sharedlib/
	$(CP_R) $(SAI_LIBS) $(DESTDIR)/sharedlib/

clean:
	echo "No-op"
