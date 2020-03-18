#include "tm4c123ge6pm.h"
#include "tm4c_utils.c"
#include <stdlib.h>
#include <stdio.h>

#ifndef __NO_SYSTEM_INIT
void SystemInit() {}
#endif

int main(void){    // initialize PF0 and PF4 and make them inputs
	PortF_Init();    // make PF3-1 out (PF3-1 built-in LEDs)

	while (1) {
		if (!SW1) {
			LED_GREEN = PF_LED_GREEN;
			continue;
		}

		LED_GREEN = 0;	
	}
}
