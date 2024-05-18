package main

import (
	"fmt"

	"github.com/nacimux/go-edi/edi"
)

func main() {
	x12Message := "ISA*00*          *00*          *ZZ*ABCDEFGHIJKLM  *01*000000000000000*120723*1253*U*00401*000000001*0*P*>~GS*PO*4405197800*4405197800*20120723*1701*1*X*004010~ST*850*000000001~BEG*00*SA*08292233294**20120723~REF*IA*08292233294~REF*DP*038~REF*BL*BL382648~DTM*002*20120723~N1*ST*Company Name*92*0000000198~PO1*1*144*EA*3.95*CB*IN*VP*000000000000123456~CTT*1~SE*10*000000001~GE*1*1~IEA*1*000000001~"
	edifactMessage := "UNA:+.? 'UNB+UNOC:3+1234567890123:14+1234567890123:14+200429:1907+1++1'UNH+1+ORDERS:D:96A:UN'BGM+220+BKOD99+9'DTM+137:20040129:102'FTX+AAI+++DOCUMENTATION'UNS+S'UNT+9+1'"

	x12Parser := edi.NewParser(edi.X12)
	edifactParser := edi.NewParser(edi.EDIFACT)

	parsedX12, err := x12Parser.Parse(x12Message)
	if err != nil {
		fmt.Println("Error parsing X12 message:", err)
		return
	}
	fmt.Println("Parsed X12 message:", parsedX12)

	serializedX12, err := x12Parser.Serialize(parsedX12)
	if err != nil {
		fmt.Println("Error serializing X12 message:", err)
		return
	}
	fmt.Println("Serialized X12 message:", serializedX12)

	parsedEDIFACT, err := edifactParser.Parse(edifactMessage)
	if err != nil {
		fmt.Println("Error parsing EDIFACT message:", err)
		return
	}
	fmt.Println("Parsed EDIFACT message:", parsedEDIFACT)

	serializedEDIFACT, err := edifactParser.Serialize(parsedEDIFACT)
	if err != nil {
		fmt.Println("Error serializing EDIFACT message:", err)
		return
	}
	fmt.Println("Serialized EDIFACT message:", serializedEDIFACT)
}
