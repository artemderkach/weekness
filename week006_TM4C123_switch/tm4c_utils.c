#include "tm4c_utils.h"

void DelayMS(unsigned long time) {
        unsigned long i;
        while (time > 0) {
                i = 1333;
                while (i > 0) {
                        i -= 1;
                }
                time -= 1;
        }
}

