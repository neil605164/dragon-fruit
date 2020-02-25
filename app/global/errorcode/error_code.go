package errorcode

var errorCode = map[int]string{
	/** 共同 [1001XXX] **/
	1001001: "SUCCESS",                            // 呼叫API成功
	1001002: "PERMISSION DENIED",                  // 權限不足
	1001003: "CREATE DIR ERROR",                   // 建立資料夾失敗
	1001004: "STRUCT TO MAP JSON MARSHAL ERROR",   // json encode 失敗
	1001005: "STRUCT TO MAP JSON UNMARSHAL ERROR", // json decode 失敗
	1001006: "PARSE TIME ERROR",                   // 時間格式轉換錯誤
	1001007: "CRYPTION ERROR",                     // 密碼加密錯誤
	1001008: "GET TIME ZONE ERROR",                // 取當前時區錯誤
	1001009: "LOG ID NOT EXIST",                   // Log 身份證
	1001010: "MASTER DB CONNECT ERROR",            // Master DB連線失敗
	1001011: "SLAVER DB CONNECT ERROR",            // Master DB連線失敗
	1001012: "REDIS CONNECT ERROR",                // Redis連線失敗
	1001013: "REDIS INSERT ERROR",                 // Redis寫入失敗
	1001014: "REDIS HASH INSERT ERROR",            // Redis寫入失敗
	1001015: "REDIS DELETE ERROR",                 // Redis刪除失敗
	1001016: "REDIS APPEND ERROR",                 // Redis增加值失敗
	1001017: "REDIS SET EXPIRE ERROR",             // Redis設定過期時間失敗
	1001018: "REDIS CHECK EXIST ERROR",            // 檢查Redis值是否存在時發生錯誤
	1001019: "REDIS PING ERROR",                   // Redis Ping 錯誤
	1001020: "REDIS GET VALUE ERROR",              // Redis 取值錯誤
	1001021: "REDIS GET HASH VALUE ERROR",         // Redis 取值錯誤
	1001022: "CURL GET METHOD CREATE FAIL",        // CURL GET　METHOD 建立失敗
	1001023: "CURL POST METHOD CREATE FAIL",       // CURL POST 建立失敗
	1001024: "CURL PUT METHOD CREATE FAIL",        // CURL PUT 建立失敗
	1001025: "GET METHOD API CONNECT ERROR",       // 對外連線失敗
	1001026: "POST METHOD API CONNECT ERROR",      // 對外連線失敗
	1001027: "PUT METHOD API CONNECT ERROR",       // 對外連線失敗
	1001028: "CURL GET FAIL",                      // CURL GET 失敗
	1001029: "CURL POST FAIL",                     // 取API失敗
	1001030: "CURL POST FAIL",                     // 取API失敗
	1001031: "GET METHOD API STATUS ERROR",        // 對外連線回傳code異常
	1001032: "POST METHOD API STATUS ERROR",       // 對外連線回傳code異常
	1001033: "PUT METHOD API STATUS ERROR",        // 對外連線回傳code異常
	1001034: "DB TABLE NOT EXIST",                 // 資料庫表不存在

	/** Redis 錯誤 [1003XXX] **/
	1003001: "REDIS PUBLISH ERROR",           // Redis publish 失敗
	1003002: "REDIS SUBSCRIBE CONNECT ERROR", // Redis subscribe connect 失敗
	1003003: "REDIS SUBSCRIBE Receive ERROR", // Redis subscribe receive 失敗

	/** WebSocket 錯誤 [1004XXX] **/
	1004001: "WEBSOCKET CONNECT ERROR", // Websocket Connect Error
}
