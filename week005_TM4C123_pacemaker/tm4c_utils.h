// Color    LED(s) PortF
// dark     ---    0
// red      R--    0x02
// blue     --B    0x04
// green    -G-    0x08
// yellow   RG-    0x0A
// sky blue -GB    0x0C
// white    RGB    0x0E


// bit specific mapping for LED's and switches


#ifndef TM4C_UTILS
#define TM4C_UTILS

// delay 1 millisecond based on clock cicles
void DelayMS(unsigned long);

// variable needed for port initialization
volatile unsigned long delay;

void PortF_init(void);

#define SW2		(*((volatile uint32_t *)0x40025004)) //	PF0
#define LED_RED 	(*((volatile uint32_t *)0x40025008)) // PF1
#define LED_BLUE 	(*((volatile uint32_t *)0x40025010)) // PF2
#define LED_GREEN 	(*((volatile uint32_t *)0x40025020)) // PF3
#define SW1 		(*((volatile uint32_t *)0x40025040)) // PF4

#define PF_SW2 		0x01
#define PF_LED_RED 	0x02
#define PF_LED_BLUE 	0x04
#define PF_LED_GREEN 	0x08
#define PF_SW1 		0x10


#endif // TM4C_UTILS
