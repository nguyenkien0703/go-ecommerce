package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age:%d", "tipgo", 18) // like fmt.Printf(format, a)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "tipgo"), zap.Int("age", 18)) // like fmt.Println(a)
	encoder := getEncoderLog()
	sync := getWriterSync()

	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))



}


// /format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()


	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.TimeKey = "time"

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)

}

// 
func getWriterSync() zapcore.WriteSyncer {
	// Tạo thư mục nếu chưa tồn tại
    if err := os.MkdirAll("./log", os.ModePerm); err != nil {
        panic(err)
    }



	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)


	

}



