#include "tm4c123ge6pm.h"
#include "tm4c_utils.c"
#include <stdlib.h>
#include <stdio.h>

#ifndef __NO_SYSTEM_INIT
void SystemInit() {}
#endif

unsigned long previousState = 0x0;
int main(void){
	PortF_Init();

	previousState = SW1;
	while (1) {
		if (!SW1) {
			LED_GREEN = 0;

			previousState = SW1;
		}

		if (SW1 && previousState) {
			LED_GREEN = PF_LED_GREEN;

			previousState = SW1;
		}

		if (SW1 && !previousState) {
			DelayMS(250);
			LED_RED = PF_LED_RED;
			DelayMS(250);
			LED_RED = 0;

			previousState = SW1;
		}
	}
}
