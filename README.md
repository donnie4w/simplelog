## simplelog  Fast, simple logging in go

------------

#### Log level printing：

######  Debug()，Info()，Warn(), Error() ,Fatal()

## Set the log print format：

##### e.g. SetFormat(FORMAT_SHORTFILENAME|FORMAT_DATE|FORMAT_TIME)

###### FORMAT_SHORTFILENAME|FORMAT_DATE|FORMAT_TIME   default format

	only log content is printed(not formatted)	FORMAT_NANO	 	e.g. SetFormat(FORMAT_NANO)
	long file name and number of lines		FORMAT_LONGFILENAME	e.g.	/etc/log/logging_test.go:10	
	short file name and line number			FORMAT_SHORTFILENAME	e.g.  logging_test.go:10
	to date						FORMAT_DATE		e.g.	2023/02/14
	to the second					FORMAT_TIME		e.g.	01:33:27
	to microseconds					FORMAT_MICROSECNDS	e.g.	01:33:27.123456

##### printed result e.g.

###### [DEBUG]2023/02/14 01:33:27 logging_test.go:10: 11111111111111

### Log level

#####  DEBUG < INFO < WARN < ERROR < FATAL

###### Close all logs : SetLevel(OFF)

#### Direction for use：

	when:	SetLevel(INFO)
	then:	Debug("*********")   will not be executed

	by parity of reasoning:
	when:	SetLevel(ERROR)
	then:	Debug() and Info() and Warn()  will not be executed
	
	when:	SetLevel(OFF) 
	All logs will not be printed


#### If logs need to be written to a file, set the log file

	When the global log object is used, the setting method is called directly：

	SetRollingDaily()		Split log files by date
	SetRollingByTime()		Log file can be cut by hour, day, and month
	SetRollingFile()		Split log files by file size
	SetRollingFileLoop()		Split log files by file size and retain a maximum number of log files
	SetGzipOn(true)			Compress the split log file

#### Multiple instance object：

	log1 := NewLogger()
	log1.SetRollingDaily("", "logMonitor.log")
	 
	log12:= NewLogger()
	log2.SetRollingDaily("", "logBusiness.log")

#### 1. Split log files by date

	log.SetRollingDaily("d:/foldTest", "log.txt")
	Split the log file in this format every day: log_20221015.txt
	if log_20221015.txt exists， log_20221015.1.txt will be generated
	
	log.SetRollingByTime("d:/foldTest", "log.txt",MODE_MONTH)
	Split logs by month, and keep the logs of the previous month when cross-month, such as:
		log_202210.txt
		log_202211.txt
		log_202212.txt
	
	log.SetRollingByTime("d:/foldTest", "log.txt",MODE_HOUR)
	Split the log by hour, such as:
		log_2022101506.txt
		log_2022101507.txt
		log_2022101508.txt

#### 2. Split log files by file size

	log.SetRollingFile("d:/foldTest", "log.txt", 300, MB)
	If the file size exceeds 300MB, back up the file in the log.1.txt or log.2.txt format
	the directory parameter can be empty, the current directory is used by default.
	
	log.SetRollingFileLoop(`d:/foldTest`, "log.txt", 300, MB, 50) 
	Set the file size to a maximum of 300 MB and reserve only the latest 50 log files


#### FileOption

###### FileOption is an interface. There are two implementation objects FileSizeMode and FileTimeMode

###### FileTimeMode

	Filename   path of Log file 
	Timemode   slice by the hour, the day, the month：MODE_HOUR，MODE_DAY，MODE_MONTH
	Maxbuckup  Maximum number of backup log files
	IsCompress  Whether backup files are compressed

###### FileSizeMode

	Filename   path of Log file 
	Maxsize    The maximum log file size. If the log file size exceeds the maximum, the log file will be backed up
	Maxbuckup  Maximum number of backup log files
	IsCompress  Whether backup files are compressed

##### SetOption Example 1

	SetOption(&Option{Level: LEVEL_DEBUG, Console: false, FileOption: &FileTimeMode{Filename: "testlogtime.log", Maxbuckup: 10, IsCompress: true, Timemode: MODE_DAY}})

##### SetOption Example 2

	SetOption(&Option{Level: LEVEL_DEBUG, Console: false, FileOption: &FileSizeMode{Filename: "testlog.log", Maxsize: 1<<30, Maxbuckup: 10, IsCompress: true}})

------------

#### Console log Settings

	SetConsole(on bool)  default:true

------------

### Example：

	//SetRollingFile("", "log.txt", 1000, KB)  Log files are cut every 1000 KB
	//SetRollingFileLoop(``, "log.txt", 300, MB, 50)   Set the log file size to 300M. A maximum of 50 recent log files can be reserved
	//SetRollingByTime(``, "log.txt", MODE_MONTH) Split logs by month
	//SetRollingByTime(``, "log.txt", MODE_HOUR)  Split logs by hour
	//SetRollingByTime(``, "log.txt", MODE_DAY)   Split logs by date ,the same as SetRollingDaily("", "log.txt")
	
	
	
	//The console does not print logs
	//SetConsole(false)
	
	Debug("00000000000")
	//default format：[DEBUG]2023/07/10 18:40:49 logging_test.go:12: 00000000000

	SetFormat(FORMAT_NANO) //Set unformatted：
	Debug("111111111111")
	//111111111111

	SetFormat(FORMAT_LONGFILENAME) 
	Info("22222222")
	//Print:	[INFO]/usr/log/logging/logging_test.go:14: 22222222

	SetFormat(FORMAT_DATE | FORMAT_SHORTFILENAME) 
	Warn("333333333")
	//Print:	[WARN]2023/07/10 logging_test.go:16: 333333333

	SetFormat(FORMAT_DATE | FORMAT_TIME) /
	Error("444444444")
	//Print:	[ERROR]2023/07/10 18:35:19 444444444
	
	SetFormat(FORMAT_SHORTFILENAME)
	Fatal("5555555555")
	//Print:	[FATAL]logging_test.go:21: 5555555555

	SetFormat(FORMAT_TIME)
	Fatal("66666666666")
	//Print:	[FATAL]18:35:19 66666666666

#### Corrected print time

	//The log print time can be corrected by correcting the TIME_DEVIATION value (Unit nanosecond)
	TIME_DEVIATION 
