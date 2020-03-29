#include "tm4c123ge6pm.h"
#include "tm4c_utils.c"
#include <stdio.h>
#include <stdlib.h>

#ifndef __NO_SYSTEM_INIT
void SystemInit() {}
#endif


#define alarm PA2
#define alert PA3

#define alarmButton PE1

_Bool alarmButtonPrevState = 0;
_Bool allClosed = 0;

int main(void) {
        // activate clock
        SYSCTL_RCGC2_R |= SYSCTL_RCGC2_GPIOA | SYSCTL_RCGC2_GPIOE | SYSCTL_RCGC2_GPIOF;

        GPIO_PORTA_DEN_R |= BIT2 | BIT3; // enable ports
        GPIO_PORTA_DIR_R |= BIT2 | BIT3; // output diretion

        GPIO_PORTE_DEN_R |= BIT1 | BIT2 | BIT3;    // enable ports
        GPIO_PORTE_DIR_R &= ~(BIT1 | BIT2 | BIT3); // output diretion

        while (1) {
                // turn on/off alarm
                if (!alarmButton && !alarmButtonPrevState) {
                        alarmButtonPrevState = 1;
                }
                if (alarmButton && alarmButtonPrevState) {
                        alarmButtonPrevState = 0;

                        alarm = ~alarm; // toggle alarm
                        alert = 0;      // turn off alert
                }

                // start alert in case of either door or window is opened
                allClosed = PE2 && PE3;
                if (!alarmButton && alarm && !allClosed) {
                        alert = ~alert;
                        DelayMS(100);
                }
        }
}
