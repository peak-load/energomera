# Energomera protocol encoder / decoder package                                                                                                                                                                                                                            
                                                                                                                                                                                                                                                                            
This Go package is fully based on Python code which is used to encode and decode Energomera protocol messages,                                                                                                                                                              
which can be used to communicate with counter over rs-485 protocol.                                                                                                                                                                                                         
                                                                                                                                                                                                                                                                            
Cli example application is available at https://github.com/peak-load/energomera-cli                                                                                                                                                                                                                 
                                                                                                                                                                                                                                                                            
# Usage  / available functions  

```go
import (
     "github.com/peak-load/energomera"
)
```

Package provides two functions:
* energomera.DataDecode(sdata []byte) (msg map[string]string)                                                                                                                                                                                                                 
* energomera.DataEncode(msg map[string]string) (sdata []byte)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        
                                                                                                                                                                                                                                                                            
# Credits 
## Original Python code: 
* https://support.wirenboard.com/t/schityvanie-pokazanij-i-programmirovanie-elektroschetchika-energomera-se102m-po-rs-485/212                                                                                                                                               
* https://github.com/sj-asm/energomera        

## Documentation / resources
* GOST-R MEK 61107-2001 (RU) https://standartgost.ru/g/ГОСТ_Р_МЭК_61107-2001
* Manufacturer website (RU) http://www.energomera.ru
* Power meter users manual (RU) http://www.energomera.ru/documentations/ce102m_full_re.pdf
* Power meter basic setup guide (RU) https://shop.energomera.kharkov.ua/DOC/ASKUE-485/meter_settings_network_RS485.pdf

# License
MIT License, see [LICENSE](https://github.com/peak-load/energomera/blob/main/LICENSE)
