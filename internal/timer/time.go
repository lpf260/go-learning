package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string)( time.Time, error ){
	// ParseDuration方法用于从字符串中解析出duration（持续时间）
	// 其支持的有效单位有ns、us （或μs）、ms、s、m和h，例如，″300ms″，″-1.5h″ or ″2h45m″
	duration, err := time.ParseDuration(d)
	if err != nil{
		return time.Time{},err
	}

	return currentTimer.Add(duration), nil
}
