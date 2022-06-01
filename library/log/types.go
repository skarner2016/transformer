package log

const logKeyPrefix = "log"

type Conf struct {
	Level     string
	Path      string	// 文件保存路径
	MaxSize   int		// 对文件切割之前，日志文件的最大大小（MB）
	MaxAge    int		// 保留文件的最大天数
	MaxBackup int		// 保留文件的最大个数
	LocalTime bool
	Compress  bool		// 是否压缩/归档旧文件
}
