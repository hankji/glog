package glog

// SetThreshold set log threshold level
func SetThreshold(level string) error {
	return stderrThreshold.Set(level)
}

// SetLogtostderr set whether to stderr
func SetLogDefaultToStderr() {
	toStderr = true
}