// Color    LED(s) PortF
// dark     ---    0
// red      R--    0x02
// blue     --B    0x04
// green    -G-    0x08
// yellow   RG-    0x0A
// sky blue -GB    0x0C
// white    RGB    0x0E

#ifndef TM4C_UTILS
#define TM4C_UTILS

// delay 1 millisecond based on clock cicles
void DelayMS(unsigned long);

// port bit specific mask
#define BIT0		0x01
#define BIT1		0x02
#define BIT2		0x04
#define BIT3		0x08
#define BIT4		0x10
#define BIT5		0x20
#define BIT6		0x48
#define BIT7		0x80

// port bit specific mapping

#define PA0			(*((volatile uint32_t *)0x40004004))
#define PA1 			(*((volatile uint32_t *)0x40004008))
#define PA2 			(*((volatile uint32_t *)0x40004010))
#define PA3 			(*((volatile uint32_t *)0x40004020))
#define PA4			(*((volatile uint32_t *)0x40004040))
#define PA5			(*((volatile uint32_t *)0x40004080))
#define PA6 			(*((volatile uint32_t *)0x40004100))
#define PA7			(*((volatile uint32_t *)0x40004200))

#define PE0			(*((volatile uint32_t *)0x40024004))
#define PE1 			(*((volatile uint32_t *)0x40024008))
#define PE2 			(*((volatile uint32_t *)0x40024010))
#define PE3 			(*((volatile uint32_t *)0x40024020))
#define PE4			(*((volatile uint32_t *)0x40024040))
#define PE5			(*((volatile uint32_t *)0x40024080))
#define PE6 			(*((volatile uint32_t *)0x40024100))
#define PE7			(*((volatile uint32_t *)0x40024200))

#define PF0			(*((volatile uint32_t *)0x40025004))
#define PF1 			(*((volatile uint32_t *)0x40025008))
#define PF2 			(*((volatile uint32_t *)0x40025010))
#define PF3 			(*((volatile uint32_t *)0x40025020))
#define PF4			(*((volatile uint32_t *)0x40025040))
#define PF5			(*((volatile uint32_t *)0x40025080))
#define PF6 			(*((volatile uint32_t *)0x40025100))
#define PF7			(*((volatile uint32_t *)0x40025200))


#define SW2			(*((volatile uint32_t *)0x40025004)) //	PF0
#define LED_RED 		(*((volatile uint32_t *)0x40025008)) // PF1
#define LED_BLUE 		(*((volatile uint32_t *)0x40025010)) // PF2
#define LED_GREEN 		(*((volatile uint32_t *)0x40025020)) // PF3
#define SW1 			(*((volatile uint32_t *)0x40025040)) // PF4

#define SW2_BIT			0x01 // PF0
#define LED_RED_BIT 		0x02 // PF1
#define LED_BLUE_BIT 		0x04 // PF2
#define LED_GREEN_BIT	 	0x08 // PF3
#define SW1_BIT 		0x10 // PF4

#endif // TM4C_UTILS
