# filterAlarms

filterAlarms is an Iskratel PON utility program for generating Alarm Filters from the Alarm table in CLI.
A typical example, shown, is where ports are activated but SFP modules are not present. This is common for the GE Ethernet ports (1/1-6) which are not critical to the operation of the PON. Since Ethernet ports 1/5 and 1/6 are combo SFP/Electrical, if no device is connected to the RJ45 interface, the "Ethernet link fault" alarm will display. Another common example is if the PON SFP is connected but has no active subscribers. This may mean there is no fiber connected to the SFP but it is inserted and active.
Once these alarm filters are entered in the "diagnostics" menu, the Alarm will not show. It is important to remember to remove the Alarm Filter once a previously inactive port starts carrying services, else not use the filters at all.

| Flag | Description |
| ------ | ------ |
| -h | Show this help |
| -i | File to read alarm entries from (default "alarmList.txt") |
| -o | File to write alarm filters to (default "filterList.txt") |
| -a | Overwrite entries if file exists (default false) |
| -stdin | Read alarm entries from stdin |
| -stdout | Write alarm filters to stdout |

# Example Input
```sh
Alarm Code   Severity      Alarm Description                             DM DC Object Identity
------------ ------------- --------------------------------------------- -- -- -------------------------
   100230    Warning       Autoconfiguration disabled                    13 0  spd01.Board1           
   900240    Critical      Ethernet link fault                           13 0  SPD01.Slot1/Port6      
   900240    Critical      Ethernet link fault                           13 0  SPD01.Slot1/Port5      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot1/Port4      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot1/Port3      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot1/Port2      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot1/Port1      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port16     
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port15      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port14     
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port13      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port12     
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port11      
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port10     
  1001130    Critical      SFP module not present                        13 0  SPD01.Slot0/Port9      
  3610020    Critical      PON link loss of signal                       13 0  SPD01.Slot0/Port3      
  3610020    Critical      PON link loss of signal                       13 0  SPD01.Slot0/Port4     
--------------------------------------------------------------------------------------------------------
17 total alarms:    16 critical, 0 major, 0 minor, 1 warning(s), 0 indeterminate
--------------------------------------------------------------------------------------------------------
Legend: DM = Diagnostic module, DC = Diagnostic component
```

# Example Output
```sh
diagnostics
set alarm-filter code 100230 object spd01.Board1
set alarm-filter code 900240 object SPD01.Slot1/Port6
set alarm-filter code 900240 object SPD01.Slot1/Port5
set alarm-filter code 1001130 object SPD01.Slot1/Port4
set alarm-filter code 1001130 object SPD01.Slot1/Port3
set alarm-filter code 1001130 object SPD01.Slot1/Port2
set alarm-filter code 1001130 object SPD01.Slot1/Port1
set alarm-filter code 1001130 object SPD01.Slot0/Port16
set alarm-filter code 1001130 object SPD01.Slot0/Port15
set alarm-filter code 1001130 object SPD01.Slot0/Port14
set alarm-filter code 1001130 object SPD01.Slot0/Port13
set alarm-filter code 1001130 object SPD01.Slot0/Port12
set alarm-filter code 1001130 object SPD01.Slot0/Port11
set alarm-filter code 1001130 object SPD01.Slot0/Port10
set alarm-filter code 1001130 object SPD01.Slot0/Port9
set alarm-filter code 3610020 object SPD01.Slot0/Port3
set alarm-filter code 3610020 object SPD01.Slot0/Port4
exit
```
