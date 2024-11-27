
-- -- datatimestr to timestamp时间转化
-- local function stringToTimestamp(dateString)
-- 	if string.len(dateString) == 14 then
-- 		local year = tonumber(string.sub(dateString, 1, 4))
--     	local month = tonumber(string.sub(dateString, 5, 6))
--     	local day = tonumber(string.sub(dateString, 7, 8))
--     	local hour = tonumber(string.sub(dateString, 9, 10))
--     	local min = tonumber(string.sub(dateString, 11, 12))
--     	local sec = tonumber(string.sub(dateString, 13, 14))
--     	return tostring(os.time({year=year, month=month, day=day, hour=hour, min=min, sec=sec}))
-- 	end
-- 	return ""
-- end
-- -- timestamp to datatimestr
-- local function timestampToDatatimeStr(unixTime)
--     unixTime = tonumber(unixTime)
-- 	if unixTime and unixTime >= 0 then
-- 		local tb = {}
--         tb.year = tonumber(os.date("%Y",unixTime))
--         tb.month =tonumber(os.date("%m",unixTime))
--         tb.day = tonumber(os.date("%d",unixTime))
--         tb.hour = tonumber(os.date("%H",unixTime))
--         tb.minute = tonumber(os.date("%M",unixTime))
--         tb.second = tonumber(os.date("%S",unixTime))
-- 		local datatimeStr = string.format("%04d%02d%02d%02d%02d%02d",tb.year, tb.month,tb.day, tb.hour,tb.minute,tb.second)
--         return datatimeStr
--     end
-- 	return ""
-- end

-- local dateStr = "20240515235900" -- 20240515235900

-- -- -- 日期转时间戳
-- -- local arg_start12 = stringToTimestamp(dateStr)
-- -- print(dateStr .. "--stringToTimestamp--" .. arg_start12)
-- -- if string.len(arg_start12) <= 0 then
-- --     print( "arg_start12 == nil ")
-- -- end

-- local arg_seek_start = "20240515235900"
-- local arg_seek_end = "" 
-- local arg_seek_begin_backtime = "100"

-- local arg_start = stringToTimestamp(arg_seek_start)
-- local arg_end = stringToTimestamp(arg_seek_end)

-- if (#arg_start == 0 or #arg_end == 0) and (#arg_seek_begin_backtime > 0) then
-- 	arg_end = 10000
-- 	arg_start = 10000-tonumber(arg_seek_begin_backtime)
-- end
-- print("start=".. arg_start)
-- print("end  ="..arg_end)
-- if string.len(arg_start) <= 0 then
-- 	print("start len =".. string.len(arg_start))

-- end


-- local req_url = "/migu/sx/test.m3u8"
-- 	local url_pos = string.find(req_url,"?")
--     local stream_path = ""
-- 	if url_pos == nil then
--         --ngx.log(ngx.ERR, "[get_request_param] url not find ? .url[",req_url,"]\n")
-- 		print("url_pos == nil")
-- 	end
--     stream_path = string.sub(req_url,2,url_pos-1)

-- 	if string.len(stream_path) <= 0 then
--         --ngx.log(ngx.ERR, "[get_request_param] stream_path is nil\n")
-- 		print("string.len(stream_path) <= 0")
-- 	end
-- print("ok--stream_path=", stream_path)


local uri = "/hbstatic/fdfsdf/test.jpg"
local match_path_1 = string.match(uri, "^/hbstatic/")
if match_path_1 ~= nil then
print(match_path_1)
else 
	print("-----")
end