include makefile.conf
NAME=main

main: $(NAME).c $(STARTUP)
	$(CC) $^ $(CFLAGS) -o $(NAME)-$(CORE).axf

debug: $(NAME).c $(STARTUP)
	$(CC) $^ $(DEBUG_FLAGS) $(CFLAGS) -o $(NAME)-$(CORE).axf

clean: 
	rm -f $(NAME)-$(CORE).axf $(NAME).map

openocd:
	openocd -f ek-tm4c123gxl.cfg

gdb:
	arm-none-eabi-gdb $(NAME)-$(CORE).axf

flash:
	openocd -f ek-tm4c123gxl.cfg -c "program $(NAME)-$(CORE).axf verify reset exit"
