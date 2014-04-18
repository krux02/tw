CODE_DIR = ./AntTweakBar/src

.PHONY: ant_tweak_bar

ant_tweak_bar:
	$(MAKE) -C $(CODE_DIR)

clean:
	$(MAKE) -C $(CODE_DIR) clean
