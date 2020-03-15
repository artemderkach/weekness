include makefile.conf
NAME=sos
STARTUP_DEFS=-D__STARTUP_CLEAR_BSS -D__START=main

LDSCRIPTS=-L. -L$(BASE)/ldscripts -T nokeep.ld
LFLAGS=$(USE_NANO) $(USE_NOHOST) $(LDSCRIPTS) $(GC) $(MAP)

$(NAME)-$(CORE).axf: $(NAME).c $(STARTUP)
	$(CC) $^ -g -ggdb3 -O0 -Wall $(CFLAGS) $(LFLAGS) -o $@

clean: 
	rm -f $(NAME)*.axf *.map *.o

openocd:
	openocd -f ek-tm4c123gxl.cfg

gdb:
	arm-none-eabi-gdb sos-CM4.axf

flash:
	openocd -f ek-tm4c123gxl.cfg -c "program sos-CM4.axf verify reset exit"
