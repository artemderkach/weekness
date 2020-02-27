#include "tm4c123ge6pm.h"

#ifndef __NO_SYSTEM_INIT
void SystemInit()
{}
#endif


// #define SYS_CTRL_RCGC2  (*((volatile unsigned long *)0x400FE108))   //offset of RCGC2 register is 0x108
// #define CLK_GPIOF   0x20
// 
// //---GPIO-F REGISTER---//
// #define PORTF_DATA  (*((volatile unsigned long *)0x40025038))   //offset of DATA register for PF1, PF2, PF3 is 0x38 [PF7:PF0::9:2]
// #define PORTF_DIR   (*((volatile unsigned long *)0x40025400))   //offset of DIR register is 0x400
// #define PORTF_DEN   (*((volatile unsigned long *)0x4002551C))   //offset of DEN register is 0x51C
// 
// //---PORT-F I/O---//
// #define PF0 0x01
// #define PF1 0x02
// #define PF2 0x04
// #define PF3 0x08
// #define PF4 0x10
// 
// unsigned long In;
// unsigned long Out;
// 
// int main(void)
// {
//    SYS_CTRL_RCGC2 |= CLK_GPIOF;
//    PORTF_DIR |= 0x0000000E;    //set PF1, PF2, PF3 as output
//    PORTF_DEN |= 0x0000001F;    //enable PF1, PF2, PF3
//    PORTF_DATA = 0;
//    while(1) {
//        In = PORTF_DATA&0x10;   // read PF4 into Sw1
//        In = In>>2;                    // shift into position PF2
//        Out = PORTF_DATA;
//        Out = Out&0xFB;
//        Out = Out|In;
//        PORTF_DATA = Out;        // output
//    }
// }


unsigned long In;  // input from PF4

unsigned long Out; // output to PF2 (blue LED)

//   Function Prototypes
void PortF_Init(void);
// 3. Subroutines Section
// MAIN: Mandatory for a C Program to be executable
int main(void){    // initialize PF0 and PF4 and make them inputs
  PortF_Init();    // make PF3-1 out (PF3-1 built-in LEDs)
  while(1){
    In = (~GPIO_PORTF_DATA_R)&0x10;   // read PF4 into Sw1
    In = In>>2;                    // shift into position PF2
    Out = GPIO_PORTF_DATA_R;
    Out = Out&0xFB;
    Out = Out|In;
    GPIO_PORTF_DATA_R = Out;        // output
  }
}
// Subroutine to initialize port F pins for input and output
// PF4 is input SW1 and PF2 is output Blue LED
// Inputs: None
// Outputs: None
// Notes: ...
//
volatile unsigned long delay;

void PortF_Init(void){ 
  SYSCTL_RCGC2_R |= 0x00000020;     // 1) activate clock for Port F
  delay = SYSCTL_RCGC2_R;           // allow time for clock to start
  GPIO_PORTF_LOCK_R = 0x4C4F434B;   // 2) unlock GPIO Port F
  GPIO_PORTF_CR_R = 0x1F;           // allow changes to PF4-0
  // only PF0 needs to be unlocked, other bits can't be locked
  GPIO_PORTF_AMSEL_R = 0x00;        // 3) disable analog on PF
  GPIO_PORTF_PCTL_R = 0x00000000;   // 4) PCTL GPIO on PF4-0
  GPIO_PORTF_DIR_R = 0x0E;          // 5) PF4,PF0 in, PF3-1 out
  GPIO_PORTF_AFSEL_R = 0x00;        // 6) disable alt funct on PF7-0
  GPIO_PORTF_PUR_R = 0x11;          // enable pull-up on PF0 and PF4
  GPIO_PORTF_DEN_R = 0x1F;          // 7) enable digital I/O on PF4-0
}
