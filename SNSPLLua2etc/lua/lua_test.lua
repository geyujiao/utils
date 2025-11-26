
-- local request_uri = "/migu/sx/test.jpg?t=100"
-- local url = string.match(request_uri, "(.*)?.*") or ""
-- local _,_,data = string.find(url, "(.*/).-$")

-- print("data=" .. data)



-- local _,_,path = string.find("/migu/sx/test.jpg?t=100", "([^?]+)")
-- local url = "host" .. (path or "/migu/sx/test.jpg")

-- local _,_,data = string.find(url, "(.*/).-$")

-- print("data=" .. data)


-- local sha2 = require "sha2"
 
-- local input = "zzspz7s0966soivs6pj3tbie20240425102719/plugin-platform-manager/config/792c87eba5c946a88ff84b99265696f4.zip"
-- local digest = sha2.sha256(input)
 
-- print("SHA256: " .. digest)

-- local domain_list = {
--     ["honorboard-drcn.hihonorcdn.com"]={directory = {["/"]=2592000,["/blue/serviceContent/sportMatchInfo"]=10,["/blue/contentInstance/"]=60,["/blue/info-stream/infoContent"]=86400}},
--     ["appmarket-test-drcn.hihonorcdn.com"]={directory = {["/"]=2592000}},
--     ["apppkg-p01-drcn.hihonorcdn.com"]={directory = {["/"]=2592000},suffix = {["php"]=0,["jsp"]=0,["asp"]=0,["aspx"]=0}}
-- -- }
-- local domain_list = {
--     ["honorboard-drcn.hihonorcdn.com"]={directory={{["/"]=2592000},{["/blue/serviceContent/sportMatchInfo"]=10},{["/blue/contentInstance/"]=60},{["/blue/info-stream/infoContent"]=86400}}}
-- }

-- local domain = "honorboard-drcn.hihonorcdn.com"
-- local uri = "/blue/serviceContent/sportMatchInfo"
-- local domainInfo = domain_list[domain]
-- if domainInfo ~= nil then 
--     domain_exist = true
--     if domainInfo.directory ~= nil then
--         for i,item in pairs(domainInfo.directory) do
--             -- print("i=" .. i )
--             for k,v in pairs(item) do 
--                 print("k=".. k)
--                 if item["/"] ~= nil then
--                     cache_control_value = item["/"]
--                 end
--                 if string.find(uri, k, 1, true) then
--                     cache_control_value = v
--                 end
--             end
--         end
--     end
--     -- if domainInfo.suffix ~= nil and domainInfo.suffix[file_suffix] ~= nil then
--     --     cache_control_value = domainInfo.suffix[file_suffix]
--     -- end
-- end

-- local ccc = -1
-- --                                                                         set ture    false
-- uri1 = "https://planck-drcn.hihonorcdn.com/blue/serviceContent/sportMatchInfo" -- 2       2
-- uri2 = "https://planck-drcn.hihonorcdn.com/blue/contentInstance/a.test"      -- 1       3
-- uri3 = "https://planck-drcn.hihonorcdn.com/blue/info-stream/infoContent"     -- 1       1
-- a = {
--     {["/"] = 1},
--     {["/blue/serviceContent/sportMatchInfo"] = 2},
--     {["/blue/contentInstance/"] = 3},
--     {["/blue/info-stream/infoContent"] = 4}
-- }
-- function is_array(t)
--     for k, v in pairs(t) do
--         if type(k) ~= "number" then
--             return false
--         end
--     end
--     return true
-- end

-- -- if is_array(a) then
--     for i, item in ipairs(a) do
--         for key, val in pairs(item) do
--             if key == "/" then
--                 ccc = val
--             end
--             if string.find(uri1, key, 1, true) then
--                 ccc = val
--             end
--         end
--     end
-- -- else
--     -- for key, val in pairs(item) do
--     --     if key == "/" then
--     --         ccc = val
--     --     end
--     --     if string.find(uri1, key, 1, true) then
--     --         ccc = val
--     --     end
--     -- end
-- -- end





-- print(ccc)
-- local uri = "/log/test.txt"
-- local cache_time_tmp =0
-- local cache_url = {"/picture","/log","/vedio","/doc","/unity","/other","/resourc"}
--  for k,v in pairs(cache_url) do
--     print(v)
--         if uri:find(v) then
--             cache_time_tmp = 604800
--         end
--     end

--     print("kkk", cache_time_tmp)

-- local nowtime = os.time()
-- local datatime = "20240425102719"
-- print(nowtime)
-- print(tonumber(datatime+100))
-- if tonumber(nowtime) > tonumber(datatime+100) then
--     --ngx.log(ngx.ERR, "expired")
--     print("---")
-- end

-- function func1(url, mark) 
--     local captureUrl_str = ""
--     if "" ~= url then
-- 		url = "http://"..ts.ctx["req_host"]..url 
-- 		if mark == 0 then
-- 			captureUrl_str = string.match(url, "(.*/[^/]+.[^/]+)?.*") or ""
-- 		elseif mark == 1 then
-- 			captureUrl_str = string.match(url, "(.*/)[^/]+.[^/]+?.*") or ""
-- 		elseif mark == 2 then
-- 			captureUrl_str = string.match(url, ".*/[^/]+.[^/]+?(.*)") or ""
-- 		elseif mark == 3 then
-- 			captureUrl_str = string.match(url, "http[s]?%://[^/]+(.*/[^/]+%.[^/]+)?.*") or ""
-- 		else
-- 			-- captureUrl_str = ts.client_request.get_url()
-- 		end
-- 	end
-- end


--     print("captureUrl_str" .. func1("http://", 0))


-- local str = "bbvaryaa"
-- if string.match(str, ".*vary.*") then
--     print("---")
-- end


-- local function stringToTimestamp(dateString)
--     local year = tonumber(string.sub(dateString, 1, 4))
--     local month = tonumber(string.sub(dateString, 5, 6))
--     local day = tonumber(string.sub(dateString, 7, 8))
--     local hour = tonumber(string.sub(dateString, 9, 10))
--     local min = tonumber(string.sub(dateString, 11, 12))
--     local sec = tonumber(string.sub(dateString, 13, 14))
--     return os.time({year=year, month=month, day=day, hour=hour, min=min, sec=sec})
-- end
-- local str = "20240515235900"
-- str = stringToTimestamp("20240515235900")
-- print(str) -- 1715788740

-- local function format_unix_time_to_date_string(unixTime)

--     if unixTime and unixTime >= 0 then
--     	local tb = {}
--         tb.year = os.date("%Y",unixTime)
--         tb.month =os.date("%m",unixTime)
--         tb.day = os.date("%d",unixTime)
--         tb.hour = os.date("%H",unixTime)
--         tb.minute = os.date("%M",unixTime)
--         tb.second = os.date("%S",unixTime)
--         return true,tb
--     end
-- 	return false,nil
-- end
-- --   return string.format("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
-- -- %04d%02d%02d%02d%02d%02d
-- local _, datatime = format_unix_time_to_date_string(1715788740)
-- -- local str = string.format("%04d-%02d-%02d %02d:%02d:%02d", datatime.year, datatime.month, datatime.day, datatime.hour, datatime.min, datatime.sec)
-- local str = string.format("%04d%02d%02d%02d%02d%02d",datatime.year, datatime.month,datatime.day, datatime.hour,datatime.minute,datatime.second)

-- print("datatime="..  str)




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

-- local arg_start1 = "" -- 20240515235900

-- local arg_start12 = stringToTimestamp(arg_start1)
-- print(arg_start1 .. "--stringToTimestamp--" .. arg_start12)
-- if string.len(arg_start12) <= 0 then
--     print( "arg_start12 == nil ")
-- end
-- local arg_start21 = "1715788740" -- 1715788740 -- 20240515235900
-- local arg_start22 = timestampToDatatimeStr(arg_start21)

-- print(arg_start21 .. "--timestampToDatatimeStr--" .. arg_start22)


-- print("diff" .. os.difftime(end12,start1))

-- 假设你有两个日期时间，使用os.time获取它们的时间戳
-- local t1 = os.time({year=2023, month=4, day=15, hour=0, min=0, sec=0})
-- local t2 = os.time({year=2023, month=4, day=16, hour=0, min=0, sec=0})
 
-- -- 计算差异
-- local diff = os.difftime(t2, t1)
 
-- print("diff=" .. diff)
-- -- 将差异转换为可读的格式（小时、分钟或秒）
-- local hours = diff / 3600
-- local minutes = diff / 60 % 60
-- local seconds = diff % 60
 
-- print("时间差为：" .. hours .. "小时 " .. minutes .. "分钟 " .. seconds .. "秒")

-- url = 'http://fsdfsdfs/bytedance/live/stream01/1715792399.ts?uuid=59ec7921-23b3-4877-9c64-5472e6d6fdad&storage_alias=source-hebei-ceph'
-- -- captureUrl_str = string.match(url, ".*/[^/]+.[^/]+?(.*)") or ""
-- captureUrl_str = string.match(url, "http[s]?%://[^/]+(.*/[^/]+%.[^/]+)?.*") or ""

-- print(captureUrl_str)

-- content-range:bytes 10-20/-1
-- range:bytes=10-20
-- local str = "bytes 10-20/-1"

-- local strnew = string.sub(str,7,#str-3)
-- print("range:".. "bytes=" ..strnew)


-- local function captureUrl(url, mark)
--     local captureUrl_str = ""
--     if "" ~= url then
--         url = "http://" .. "hostname" .. url
--         if mark == 0 then
--             captureUrl_str = string.match(url, "(.*/[^/]+.[^/]+)?.*") or ""
--         elseif mark == 1 then
--             captureUrl_str = string.match(url, "(.*/)[^/]+.[^/]+?.*") or ""
--         elseif mark == 2 then
--             captureUrl_str = string.match(url, ".*/[^/]+.[^/]+?(.*)") or ""
--         elseif mark == 3 then
--             captureUrl_str = string.match(url, "http[s]?%://[^/]+(.*/[^/]+%.[^/]+)?.*") or ""
--         else
--             captureUrl_str = "other"
--         end
--     end
--     -- if mark < 2 and "" == captureUrl_str then
--     --     captureUrl_str = ts.client_request.get_url_host() .. ts.client_request.get_uri()
--     -- end
--     -- if mark == 2 and "" == captureUrl_str then
--     --     captureUrl_str = ts.client_request.get_uri_args()
--     -- end
--     -- if mark == 3 and "" == captureUrl_str then
--     --     captureUrl_str = ts.client_request.get_uri()
--     -- end
--     return captureUrl_str
-- end

-- local cache_key = captureUrl("/migu/sx/test.ts?start=1&end=2", 2)
-- print("cache_key=" .. cache_key)


-- function getFileSize(filePath)
--     local file = io.open(filePath, "r")
--     if not file then
--         return nil, "Unable to open file: " .. filePath
--     end
--     local fileSize = file:seek("end")
--     file:close()
--     return fileSize
-- end
 
-- -- 使用函数
-- local filePath = "example.txt"
-- local size, err = getFileSize(filePath)
-- if size then
--     print("File size: " .. size .. " bytes")
-- else
--     print(err)
-- end

-- local url = "/migu/sx/a/b/test.jpg?t=1&m=2&encrypt=3&n=4"
-- local md5str = (string.match(url, "(.*)&encrypt.*"))
-- print("md5str11="..md5str)
-- local md5str, md5str2 = string.match(md5str, "/(.*)/.*/.*?(.*)")
-- print("md5str="..md5str .. "/" .. md5str2) -- local url = "migu/sx/a/t=1&m=2"

-- local url = "/migu/sx/a/b/test.jpg"
-- local args = "t=1&m=2&encrypt=3&n=4"
-- local url1 = (string.match(url, "/(.*)/.*/.*$"))
-- print("url1="..url1)

-- local args2 = string.match(args, "(.*)&encrypt.*")
-- print("args2="..args2)

-- print("md5str="..url1 .. "/" .. args2) -- local url = "migu/sx/a/t=1&m=2"


-- local tmp_body_ts_relative_uri = "migu/sx/a/b/test.jpg"
-- local dir = tmp_body_ts_relative_uri:split("/")
-- print("ts_path_pre=" ..dir[1])
-- local l_location = "http://baidu.com/migu|/sx/test.jpg?a=123456"
-- -- local l_scheme, l_domain, port_t = string.match( l_location, "^(%w+)://([^:/]+):*(%d*)(.*)$") 
-- -- if l_scheme == nil then
-- --     print("l_scheme is nil")
-- -- end

-- local function urlEncode(s)
--     s =  string.gsub(s, "|", function(c) return string.format("%%%02X", string.byte(c)) end)
--     return s
-- end

-- print( urlEncode(l_location) )
-- local host = "fs:fsf:fsf:fsdf"
-- local s = string.find(host, ":",1, true) 
-- 	if s then
--         print("exist")
--         else
--             print("no exist")
-- 	end

-- local r = ngx_re_find(host, [[^\d+\.\d+\.\d+\.\d+$]], "ijo")
-- if r then
--     print("exist")
--     else
--         print("no exist")
-- end


-- function getlast(list)
--     if not list then
--         return "-"
--     else
--         local i = 0
--         local index = 1
--         while true do
--             i = string.find(list,"[,:] ",i+1)
--             if not i then
--                 break
--             end
--             index = i + 2
--         end
--         return string.sub(list,index,#list)
--     end
-- end
-- 获取最后一个upstream地址
-- local upstream_addr = "10.12.7.145:8686"
-- local upstream_addr = "[2001:202:27:7::12]:8686"

-- local upstream_addr_internal = ""
-- local addr = getlast(upstream_addr) or "-"
-- print("---addr---" .. addr)
-- local index = string.find(addr,":")
-- if addr ~= "-" and index then
--     addr = string.sub(addr,0,index - 1)
-- end
-- if nil == addr or "-" == addr or "" == addr then
--     addr =  getlast(upstream_addr_internal) or "-"
--     local index = string.find(addr,":")
--     if addr ~= "-" and index then
--         addr = string.sub(addr,0,index - 1)
--     end
-- end

-- if addr == "" or "-" ==  addr then
--     addr = "NULL"
-- end
-- local cache_addr = addr
-- print("--cache_addr---" .. addr)


-- local m, _ = ngx_re_match("[2001:202:27:7::12]:8686", [[^(http[s]?)://([^/]+\])(?::(\d+))?(.*)]], "ijo")
-- local m, _ = ngx_re_match("[2001:202:27:7::12]:8686", [[([^/]+\])(?::(\d+))?(.*)]], "ijo")
-- if not m then
--     addr = "-"
-- end
    
-- print("addr="..addr)
-- local uri_ipv6_pattern = "([%w:%]+:[%w:%]+)"; -- IPv6地址正则表达式
-- local ngx_re_match = ngx.re.match;
 
-- -- 假设请求的URI中包含IPv6地址
-- local uri = "http://[2001:202:27:7::12]:8686";
 
-- -- 使用Lua的ngx.re.match进行正则匹配
-- local match = ngx_re_match(uri, "http[s]?://[" .. uri_ipv6_pattern .. "]");
 
-- if match then
--     -- 如果匹配成功，则可以通过match[1]获取整个IPv6地址
--     local ipv6_address = match[1];
--     ngx.say("IPv6 address found: ", ipv6_address);
-- else
--     ngx.say("No IPv6 address found in URI.");
-- end

-- function find_last_colon(str)
--     local reversed_str = string.reverse(str)
--     local first_colon_pos, _ = string.find(reversed_str, ":")
--     if first_colon_pos then
--         return string.len(str) - first_colon_pos + 1
--     end
--     return nil
-- end
 
-- local uri = "[fe80::95eb:7779:216c:e367]:80" --"http://[2001:202:27:7::12]:8686";
-- local last_colon_pos = find_last_colon(uri)
-- if last_colon_pos then
--     ngx.log(ngx.ERR, "testlog---uri-last_colon_pos", last_colon_pos)
--     ngx.log(ngx.ERR, "testlog---uri-ip", string.sub(uri,0,last_colon_pos - 1))
-- end
--
-- print("ipv4 1   " .. gethostFromUpstream("10.12.7.141"))
-- print("ipv4 2  " .. gethostFromUpstream("10.12.7.141:80"))
-- print("ipv6 3  " .. gethostFromUpstream("[fe80::95eb:7779:216c:e367]:80"))
-- print("ipv6 4  " .. gethostFromUpstream("[fe80::95eb:7779:216c:e367]"))
-- print("ipv6 5  " .. gethostFromUpstream("fe80::95eb:7779:216c:e367"))
-- print("- 5  " .. gethostFromUpstream("-"))
-- print(" 5  " .. gethostFromUpstream(""))

-- local atslog = "TCP_MISS|MISS|127.0.0.1|GET|HTTP/1.1|v-cm.pstatp.com|/video/tos/cn/tos-cn-ve-15/owfmwF9yA5DlxU4ERjFjUI9EAeCcEgB0IXQCnA/|text/html; charset=utf-8,gbk|200|FIN|Cache|-|-|0|1812|-|20240722T061618.532Z|20240722T061618.535Z|20240722T061618.532Z|-"
-- -- 2. 解析header 值
-- local addr_tmp = "huiyuan ip"
-- local status_tmp = "huiyuan status"

-- -- 3. 解析atslog值，替换22和28字段
-- local atsloglist = strSplit(atslog, "|")
--  for i, part in ipairs(atsloglist) do
--         if i == 3 then
--             -- 回源ip
--             -- part = addr_tmp
--             print("333==" .. part)
--         end
--         if i == 9 then
--             -- 源站返回的状态status
--             -- part = status_tmp
--             print("999==" .. part)
--         end

--     end 
-- atsloglist[3] = addr_tmp
-- atsloglist[9] = status_tmp
-- local atslog2 =  table.concat(atsloglist, "|")
-- print("atslog2 result==" .. atslog2)

-- function getlast(list)
--     if not list then
--         return "-"
--     else
--         local i = 0
--         local index = 1
--         while true do
--             i = string.find(list,"[,:] ",i+1)
--             if not i then
--                 break
--             end
--             index = i + 2
--         end
--         return string.sub(list,index,#list)
--     end
-- end
-- local function getUpstreamList(str)
--     local pattern =  "[,:] "  -- "[,:]%s*"
--     local parts = {}
--     if not str then
--         return parts
--     end
--     local from = 1
--     local delim_from, delim_to = string.find(str, pattern, from)
--     while delim_from do
--         table.insert(parts, string.sub(str, from, delim_from - 1))
--         from = delim_to + 1
--         delim_from, delim_to = string.find(str, pattern, from)
--     end
--     table.insert(parts, string.sub(str, from))
--     return parts
-- end

-- local function gethostFromUpstream(str)
--     if not str then
--         return "-"
--     end
--     if string.find(str, ".",1, true) then
--         -- ipv4
--         local index = string.find(str,":",1, true)
--         if index then
--             str = string.sub(str,0,index - 1)
--             return str
--         end
--     else
--         --ipv6
--         local index = string.find(str,"]",1, true)
--         if index then
--             str = string.sub(str,0,index)
--             return str
--         end
--     end
--     return str
-- end

-- local function getOneAddrStatus(addrlistStr, statuslistStr)
--     local addrParts = getUpstreamList(addrlistStr)
--     local statusParts = getUpstreamList(statuslistStr)

--     local addr2xx = ""
--     local status2xx = ""
--     local addr3xx = ""
--     local status3xx = ""
--     local addr4xx = ""
--     local status4xx = ""
--     local addr5xx = ""
--     local status5xx = ""
--     -- 4xx,5xx 取最后一个
--     -- 2xx,3xx 取第一个(3xx再3xx，取第一个3xx忽略后面的)
--     for i, part_tmp in ipairs(statusParts) do
--         local part = tonumber(part_tmp) or 0
--         if part >= 500 and part < 600 then
--             addr5xx = addrParts[i] or addrParts[#addrParts]
--             status5xx = part
--         elseif part >= 400 and part < 500  then
--             addr4xx = addrParts[i] or addrParts[#addrParts]
--             status4xx = part
--         elseif part >= 300 and part < 400 and (status3xx == nil or status3xx == "")   then
--             addr3xx = addrParts[i] or addrParts[#addrParts]
--             status3xx = part
--         elseif part >= 200 and part < 300 and (status2xx == nil or status2xx == "") then
--             addr2xx = addrParts[i] or addrParts[#addrParts]
--             status2xx = part
--         end
--     end
--     if status3xx ~= nil and status3xx ~= "" then
--         return addr3xx,status3xx
--     end
--     if status2xx ~= nil and status2xx ~= "" then
--         return addr2xx,status2xx
--     end
--     if status4xx ~= nil and status4xx ~= "" then
--         return addr4xx,status4xx
--     end
--     if status5xx ~= nil and status5xx ~= "" then
--         return addr5xx,status5xx
--     end
-- end

-- [2001:202:27:7::12]:80 : [2001:202:27:7::12]:443
-- local upstream_addr = "[2001:202:27:7::12]:80 : [2001:202:27:7::13]:443"
-- local upstream_addr = "[2001:202:27:7::12]:80 , [2001:202:27:7::13]:443"

-- local upstream_addr_list = strSplit(upstream_addr, "[,:] ")

-- for i, part in ipairs(upstream_addr_list) do
--     print("part=" .. part)

-- end

-- print( getlast(upstream_addr))


-- 4xx,5xx 取最后一个
-- 2xx,3xx 取第一个(3xx再3xx，取第一个3xx忽略后面的)
-- local addrlistStr = "[2001:202:27:7::11]:80 : [2001:202:27:7::12]:443 : [2001:202:27:7::13]:443 : [2001:202:27:7::14]:443" -- "hello, world:this is a test,with commas and colons: example"
-- local statuslistStr = "200 : 400 : 206 : 202" -- "hello, world:this is a test,with commas and colons: example"

--  -- 2. 解析header 值
--  local addr_tmp,status_tmp = getOneAddrStatus(addrlistStr, statuslistStr)
--  addr_tmp = gethostFromUpstream(addr_tmp)
--  if addr_tmp == nil or addr_tmp == "" then
--      addr_tmp = "-"
--  end
--  if status_tmp == nil or status_tmp == "" then
--     status_tmp = "-"
-- end
-- print(addr_tmp .. "|" .. status_tmp )



-- local function sortArgs(uri_parts)
--     -- local uri_parts = ngx.req.get_uri_args()  -- 获取请求参数
--     -- 将字典的键保存到一个数组中
--     local keysArray = {}
--     for key, _ in pairs(uri_parts) do
--         table.insert(keysArray, key)
--     end
--     -- 对每个参数的值进行字典排序
--     table.sort(keysArray)
--     -- 重新拼接
--     local newArgs = ""
--     for _, key in ipairs(keysArray) do
--         if newArgs == "" then
--             newArgs = string.format("?%s=%s", key, uri_parts[key])
--         else
--             newArgs = string.format("%s&%s=%s", newArgs, key, uri_parts[key])
--         end
--     end

--     return newArgs
-- end

-- local uri_parts = "ch=0&cr=9&dr=0&er=0&lr=default&cd=0%7C1%7C0%7C9&cv=1&br=2440&bt=2440&cs=0&ds=3&ft=Eat_JBaffPdG2~NN1VNvAPXDy37LdU~Eg_rv~ug3..scgVQ9H6eM&mime_type=video_mp4&qs=0&rc=O2Q3aGU7Zmc1O2RkZWRoZEBpamZtO285cjVwdDMzNDYzM0BhLWE0YTEuNTMxLzQuMTFgYSNkMHJlMmRjaGFgLS1kMC9zcw%3D%3D&btag=c0000e000a0000&cdn_type=2&dy_q=1721627548&feature_id=59cb2766d89ae6284516c6a254e9fb61&l=20240722135227412DD29456A63E74A050&pwid=182&req_cdn_type=r&a=1"
-- local newArgs = sortArgs(uri_parts)
-- print("---".. newArgs)

-- local function tokbnum(limit_rate_t)
--     local rate_num = 0
--     if limit_rate_t == nil or limit_rate_t == "0"  then
--         return rate_num
--     end
--     local limit_rate_t = string.lower(limit_rate_t)
--     if string.find(limit_rate_t, "k", 1,true) then
--         local new_str = string.gsub(limit_rate_t,"k","")
--         rate_num = tonumber(new_str)
--     elseif string.find(limit_rate_t, "m", 1,true) then
--         local new_str = string.gsub(limit_rate_t,"m","")
--         rate_num = tonumber(new_str)*1024
--     elseif string.find(limit_rate_t, "g", 1,true) then
--         local new_str = string.gsub(limit_rate_t,"g","")
--         rate_num = tonumber(new_str)*1024*1024
--     elseif  tonumber(limit_rate_t) ~= nil then
--         rate_num = math.floor(tonumber(limit_rate_t)/1024) -- // 1
--     end
--     return rate_num
-- end

-- local rate_num_1 = tokbnum("100k")
-- print(rate_num_1)
-- local rate_num_2 = tokbnum("1025")
-- print(rate_num_2)

-- result = string.format("%dk",rate_num_1)
-- if rate_num_1 == 0 then 
--     result = string.format("%dk",rate_num_2)
-- elseif rate_num_2 == 0 then 
--     result = string.format("%dk",rate_num_1)
-- elseif rate_num_1 > rate_num_2 then
--     result = string.format("%dk",rate_num_2)
-- end

-- print("result="..result)

-- local uri = "https://download.aifast.komect.com:443/migu/sx/test.apk?fsfsf=111"
-- uri = string.gsub(uri, ":443/", ":80/")

-- print("uri====" .. uri)


-- local hostname = ""
-- local node_name = string.match(string.upper(hostname), "[^-]+")


-- local node = "SC03-CCS-009-CMCD-CMG"
-- local hostnamex = "SC03-CCS-006-CMCD-CMG" -- SC04-SNS-025-CMCD-CMG

-- node = string.upper(string.match(hostnamex, "[^-]+"))
-- hostnamex = string.upper(string.match(hostnamex, "[^-]+"))
-- if node == hostnamex then
--     print("same----" .. node  .. "||" .. hostnamex)
-- else
--     print("diff----".. node  .. "||" .. hostnamex)
-- end
-- local need_cache = {"js", "css", "jpg", "jpeg", "png", "gif", "woff"}
-- local filename = "example.file.with.many.dots.lua"
-- local extension = filename:match("^.*%.(.*)$")
-- print(extension) -- 输出: lua

-- function file_need_cache()
--     local num = #need_cache
--     -- local uri = ts.client_request.get_uri() or "/"
--     -- local uri = "http://h5.nf.migu.cn/app/v3/public/pkg/libs/zt_libs.729775fc.js"
--     local uri = "http://h5.nf.migu.cn/app/v3/public/pkg/libs/729775fc.js"

--     -- local suffix = string.match(uri, "^[^%.]*%.(.*)$")
--     local suffix = string.match(uri, "^.*%.(.*)$")
--     print("fffffff" .. suffix)
--     local ret = 0
--     if  suffix then
--        for i = 1, num, 1 do
--           if suffix == need_cache[i] then
--               ret = 1
--               break
--           end
--         end
--     end

--     return ret
-- end
-- local url = "http://h5.nf.migu.cn/app/v3/public/pkg/libs/zt_libs.729775fc.js"

-- local r = file_need_cache()
-- if r ~= 1 then
--     print("0000000")
-- else
--     print("1111")
-- end


-- local str = "Hfsdfsdforld"
-- if (str:match("^H.*d$")) then  -- ^Hello.*World$  -- ^%l+%l$
--     print("11111111111。")
-- else
--     print("2222222222")
-- end


-- local str = "a123d"
-- local pattern = "^a.*d$"
 
-- local match_result = string.match(str, pattern)
 
-- if match_result then
--     print("匹配成功：", match_result)
-- else
--     print("匹配失败")
-- end

-- local ua_domain_list = {["apppkg-p01-test-drcn.hihonorcdn.com"]={ua = {["Fuzz Faster U Fool"]=1,["Nmap Scripting Engine"]=1,["masscan/1.0"]=1,["zgrab/0.x"]=1}},
--     ["appmarket-drcn-test.hihonorcdn.com"]={ua = {["Fuzz Faster U Fool"]=1,["Nmap Scripting Engine"]=1,["masscan/1.0"]=1,["zgrab/0.x"]=1}},
--     ["appmarket-test-drcn.hihonorcdn.com"]={ua = {["Fuzz Faster U Fool"]=1,["Nmap Scripting Engine"]=1,["masscan/1.0"]=1,["zgrab/0.x"]=1}}
--     ["update.hihonorcdn.com"]={ua = {["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/"]=1}}
-- }


-- local function func_limit_ua(host, ua)
--     -- local ua= ngx.req.get_headers()['User-Agent']
--     if ua == nil then
--         return
--     end
--     local domain_exist = false
--     local ua_domain_info = ua_domain_list[host]
--     if ua_domain_info ~= nil then 
--         domain_exist = true
--         if ua_domain_info.ua ~= nil then
--             for k,_ in pairs(ua_domain_info.ua) do
--                 if string.find(ua, k, 1, true) then
--                     -- ngx.header["err_msg"]="auth_error"
--                     -- ngx.exit(403)
--                     print("403")
--                     return
--                 end
--             end
--         end
--     end

--     if domain_exist == false then
--         -- ngx.log(ngx.INFO, "domain no exist")
--     end
-- end
-- func_limit_ua("update.hihonorcdn.com", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/")


-- local ua = "dasfsMozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/"
-- -- local res = string.find(ua, "Mozilla/5.0Windows NT 10.0Win64x64AppleWebKit537.36KHTMLlike GeckoChrome",1, true)
-- local res = string.match(ua, "^Mozilla/5.0 %(Windows NT 10.0; Win64; x64%) AppleWebKit/537.36 %(KHTML, like Gecko%) Chrome/")

-- if res ~= nil then
--     print("fsdfsfs")
    
--     return
-- end
-- print("2222")

-- local t= "A"
--  local r = string.lower(t1)
--  print(r)

-- local uriTmp = "/cdn/liveshow/migu/test.jpg"
-- local limitLocation = "/cdn/liveshow/"
-- if string.sub(uriTmp, 0, string.len(limitLocation)) == limitLocation then
--     print("111")
-- end

-- local geEscapeChar = {
--     ["^"] = true,
--     ["$"] = true,
--     ["("] = true,
--     [")"] = true,
--     ["%"] = true,
--     ["."] = true,
--     ["["] = true,
--     ["]"] = true,
--     ["*"] = true,
--     ["+"] = true,
--     ["-"] = true,
--     ["?"] = true
-- }
-- function split(szFullString, szSeparator)
--     local nFindStartIndex = 1
--     local nSplitIndex = 1
--     local nSplitArray = {}
--     if geEscapeChar[szSeparator] then
--         szSeparator = "%"..szSeparator
--     end
--     while true do
--        local nFindLastIndex = string.find(szFullString, szSeparator, nFindStartIndex)
--        if not nFindLastIndex then
--         nSplitArray[nSplitIndex] = string.sub(szFullString, nFindStartIndex, string.len(szFullString))
--         if nSplitArray[nSplitIndex] == "" then
--             nSplitArray[nSplitIndex] = "-"
--         end
--         break
--        end
--        nSplitArray[nSplitIndex] = string.sub(szFullString, nFindStartIndex, nFindLastIndex - 1)
--        if nSplitArray[nSplitIndex] == "" then
--           nSplitArray[nSplitIndex] = "-"
--        end
--        nFindStartIndex = nFindLastIndex + string.len(szSeparator)
--        nSplitIndex = nSplitIndex + 1
--     end
--     return nSplitArray
-- end

-- local function strSplit(str, reps)
--     -- local strlist = {}
--     -- string.gsub(str]=true,['[^' .. reps .. ']+', function(w) strlist[#strlist + 1] = w end)
--     -- return strlist
-- end

-- local function l_f_split(s,delimiter)
--     local result={}
--     local newdelimiter = delimiter
--     if geEscapeChar[delimiter] then
--         newdelimiter = "%"..delimiter
--     end
--     for match in (s..delimiter):gmatch("(.-)"..newdelimiter) do
--         if match == "" then
--             match = '-'
--         end
--             table.insert(result,match)
--     end
--     return result
-- end
-- local merge_info = "20241213T040301Z||2409:896d:5c:3255:c8f1:e1a8:7761:9fb2|2409:8087:4:12::12|GET|HTTP/2.0|p96-sign.douyinpic.com|/tos-cn-i-dy/_offtrans__864x540_b467fe94155d422e4870756395a38701_e6416bf8f19bedf12f5c807156f1ff0befabd9a9568d9c91a5eccdc9fbddbb22_vvic~tplv-dy-aweme-images-offline.image?lk3s=138a59ce&x-expires=1735272000&x-signature=ujX%2BtBWuW0vYWYnQpxqCK3%2BvQWg%3D&from=327834062&s=PackSourceEnum_FEED&se=false&sc=origin_cover&biz_tag=aweme_video&l=20241213120209C56E151EFF1B75E49E57|com.ss.android.ugc.aweme/320200 (Linux; U; Android 14; zh_CN; V2244A; Build/UP1A.231005.007; Cronet/TTNetVersion:3578d129 2024-11-05 QuicVersion:d9628e3d 2024-10-11)|NULL|image/vvic|200|TCP_HIT|443|5306|20241213T040301.195Z|20241213T040301.197Z|20241213T040301.197Z|2813084340|HIT|39.136.135.231|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|NULL|83525389|NULL|NULL|1|01|38724|bytedance$https$lk3s=138a59ce&x-expires=1735272000&x-signature=ujX%2BtBWuW0vYWYnQpxqCK3%2BvQWg%3D&se=false&sc=origin_cover&from=327834062&s=PackSourceEnum_FEED&l=20241213120209C56E151EFF1B75E49E57&biz_tag=aweme_video$356$n151-185-201, http/1.1 SD05-CCS-007-CMQD-CMG ( [cHs f ]), https/1.1 BJ02-SNS-002-CMBJ-CMG ( [cHs f ])$4316$NULL$00-be2f27a30d1b2637e50354dde1e90468-be2f27a30d1b263701$NULL$edge_hit$NULL$2409:896d:5c:3255:c8f1:e1a8:7761:9fb2$L1|356|4316|n151-185-201, http/1.1 SD05-CCS-007-CMQD-CMG ( [cHs f ]), https/1.1 BJ02-SNS-002-CMBJ-CMG ( [cHs f ])|NULL|NULL|L1|NULL|14400|9A9ACAAD0296524CD3E4AA8BE16AF2FC|7499|39.136.135.231:8686|0.002|0.003|0.000|0.003|20241213T040301.194Z,-,20241213T040301.194Z,20241213T040301.195Z,-,-,-,-,-,-,-,-,-,-,NULL|-%0.002%0.003%0.000%0.003%67667%7499%26%14400%TLSv1.3%.%5946%NULL%NULL|200|bj02-sns-004-cmbj-cmg"

-- print(merge_info)

-- -- local start_time = os.time()
-- for i = 1, 10000,1 do
--     local l_t_merge_1 = strSplit(merge_info or "","|")
--     local merge_info_new_1 = table.concat(l_t_merge_1, "|")
-- end


-- local startTime = os.clock()
-- for i = 1, 10000,1 do
--     local l_t_merge_1 = strSplit(merge_info or "","|")
--     local merge_info_new_1 = table.concat(l_t_merge_1, "|")
-- end
-- -- local end_time = os.time()
-- -- local elapsed_time = os.difftime(os.clock(), start_time)
-- -- os.clock(os.clock() -startTime)
-- print("strSplit---=", (os.clock() -startTime))

-- local startTime3 = os.clock()
-- for i = 1, 10000,1 do
--     local l_t_merge_1 = split(merge_info or "","|")
--     local merge_info_new_1 = table.concat(l_t_merge_1, "|")
-- end
-- -- local end_time = os.time()
-- -- local elapsed_time = os.difftime(os.clock(), start_time)
-- -- os.clock(os.clock() -startTime)
-- print("split----=", (os.clock() -startTime3))


-- local start_time2 = os.clock()
-- for j = 1, 10000,1 do
--     local l_t_merge_1 = l_f_split(merge_info or "","|")
--     local merge_info_new_1 = table.concat(l_t_merge_1, "|")
-- end
-- -- local end_time2 = os.time()
-- -- local elapsed_time2 = os.difftime(end_time2, start_time2)
-- print("l_f_split-----=",  os.clock()-start_time2)



-- local l_t_merge_1 = l_f_split(merge_info or "","|")
-- local merge_info_new_1 = table.concat(l_t_merge_1, "|")
-- print("1len===" .. #l_t_merge_1)
-- -- print(merge_info_new_1)
-- if merge_info ~= merge_info_new_1 then
--     print("merge_info ~= merge_info_new_1")
--     print("merge_info_new_1=".. merge_info_new_1)
-- end
--1$$3$$-
-- merge_info = "1$$3$$" -- "10$NULL$NULL$NULL$0$NULL$1$NULL$NULL$0$0" --"1|2||4|5|||"
-- local l_t_merge_2 = l_f_split(merge_info or "","$")
-- -- local merge_info_new_2 = table.concat(l_t_merge_2, "$")
-- print("2len===="..#l_t_merge_2)
-- -- print(merge_info_new_2)
-- print(l_t_merge_2[1])
-- print(l_t_merge_2[2])
-- print(l_t_merge_2[3])
-- print(l_t_merge_2[4])
-- print(l_t_merge_2[5])
-- print(merge_info_new_2)
-- if merge_info ~= merge_info_new_2 then
--     print("merge_info ~= merge_info_new_2")
--     print("merge_info_new_2=".. merge_info_new_2)
-- end


-- local request_uri = "bytes|zhang"
-- print( string.gsub(request_uri,"|","%%7C"))

-- s =  string.gsub(request_uri, "|", function(c) return string.format("%%%02X", string.byte(c)) end)

-- print(s)

-- local uri = "/biz-orange/DA/cdnCache/getCdnCacheData/test.jpg"
-- local res= string.match(uri,"^/biz(-)orange/DA/cdnCache/getCdnCacheData")
-- if res ~= nil then
--     print("testlog---res ~= nil")
-- end
-- local code = 200
-- if code == "200" then
--     print("111")
-- end

-- local conf_ua = "cms-agent;cmsrefresh"
-- local head_ua = "cms-agent" -- "cmsrefresh"
-- local conf_ua_list = {
--     ["cms-agent"] = true,
--     ["cmsrefresh"] = true,
-- }
-- -- if string.find(conf_ua, head_ua, 1, true) then
-- --     print("111")
-- -- end
-- -- if conf_ua_list[head_ua] then
-- --     print("2222")
-- -- end


-- -- local start_time = os.time()
-- for i = 1, 10000,1 do
--     if string.find(conf_ua, head_ua, 1, true) then
--     end
-- end


-- local startTime = os.clock()
-- for i = 1, 10000,1 do
--     if string.find(conf_ua, head_ua, 1, true) then
--         -- print("111")
--     end
-- end
-- print("strSplit---=", (os.clock() -startTime))

-- local startTime3 = os.clock()
-- for i = 1, 10000,1 do
--     if conf_ua_list[head_ua] then
--         -- print("2222")
--     end
-- end
-- print("split----=", (os.clock() -startTime3))

-- function contains(array, value)
--     for _, v in ipairs(array) do
--         if v == value then
--             return true
--         end
--     end
--     return false
-- end

-- function stbid(server_name, arg_stbid)
--     local denied_stbid = { '1', '2', '3', '4', '5', '6', '7', '8', '9', '10' }
--     local denied_server_name = {
--     'cache.ott.bestlive.itv.cmvideo.cn',
--     'cache.ott.ystenlive.itv.cmvideo.cn',
--     'cache.ott.wasulive.itv.cmvideo.cn',
--     'cache.ott.hnbblive.itv.cmvideo.cn',
--     'cache.ott.fifalive.itv.cmvideo.cn',
--     "cache.ott.mstgwlive.itv.cmvideo.cn",
--     'zqhswlive.itv.cmvideo.cn',
--     'studentlive.migucloud.com'
--     }
--     if contains(denied_server_name, server_name)  then
--         if contains(denied_stbid, arg_stbid) then
--             print("---403---")
--         end
--     end
-- end 
-- local mapdenied_stbid = {
--     ['1']= true,
--     ['2']=true,
--     ['3']=true,
--     ['4']=true,
--     ['5']=true,
--     ['6']=true,
--     ['7']=true,
--     ['8']=true,
--     ['9']=true,
--     ['10']=true,
-- }
-- local mapdenied_server_name = {
--     ['cache.ott.bestlive.itv.cmvideo.cn'] =  true,
--     ['cache.ott.ystenlive.itv.cmvideo.cn'] =  true,
--     ['cache.ott.wasulive.itv.cmvideo.cn'] =  true,
--     ['cache.ott.hnbblive.itv.cmvideo.cn'] =  true,
--     ['cache.ott.fifalive.itv.cmvideo.cn'] =  true,
--     ['cache.ott.mstgwlive.itv.cmvideo.cn'] =  true,
--     ['zqhswlive.itv.cmvideo.cn'] =  true,
--     ['studentlive.migucloud.com'] =  true,
-- }
-- function Stdidf(server_name, arg_stbid)
--     if mapdenied_server_name[server_name] and mapdenied_stbid[mapdenied_stbid] then
--         print("---403---")
--     end
--     if mapdenied_stbid[server_name] and mapdenied_stbid[arg_stbid]  then
--         ngx.exit(403)
--     else if mapdenied_stbid[server_name] and mapdenied_stbid[arg_stbid]  then
--         ngx.exit(403)
--         end
--     end
-- end

-- local server_name = ""
-- local arg_stbid = ""

-- stbid(server_name, arg_stbid)


-- local startTime1 = os.clock()
-- for i = 1, 10000,1 do
--     Stdidf(server_name, arg_stbid)
-- end
-- print("split--map--=", (os.clock() -startTime1))

-- local startTime2 = os.clock()
-- for i = 1, 10000,1 do
--     stbid(server_name, arg_stbid)
-- end
-- print("split----=", (os.clock() -startTime2))

-- local origin_url = "http://download.aifast.komect.com:443/portal/acc/dc/test.jpg?test=bbb"
-- local origin_url = "http://download.aifast.komect.com:80/portal/acc/dc/test.jpg?test=bbb"

-- local origin_url = "http://download.aifast.komect.com/portal/acc/dc/test.jpg?test=bbb"

-- if string.find(origin_url, ":443/", 1, true) or string.find(origin_url, ":80/", 1, true) then
--     cache_key = string.gsub(origin_url, ":443/", ":80/")
-- else 
--     cache_key = string.gsub(origin_url, "download.aifast.komect.com", "download.aifast.komect.com:80")

-- -- end   

-- cache_key = string.gsub(origin_url, ":443/", ":80/")
-- if cache_key == origin_url then
--     cache_key = string.gsub(origin_url, "download.aifast.komect.com/", "download.aifast.komect.com:80/")
-- end

  
-- print(cache_key)

-- function split(szFullString, szSeparator)
--     local nFindStartIndex = 1
--     local nSplitIndex = 1
--     local nSplitArray = {}
--     while true do
--        local nFindLastIndex = string.find(szFullString, szSeparator, nFindStartIndex)
--        if not nFindLastIndex then
--         nSplitArray[nSplitIndex] = string.sub(szFullString, nFindStartIndex, string.len(szFullString))
--         break
--        end
--        nSplitArray[nSplitIndex] = string.sub(szFullString, nFindStartIndex, nFindLastIndex - 1)
--        nFindStartIndex = nFindLastIndex + string.len(szSeparator)
--        nSplitIndex = nSplitIndex + 1
--     end
--     return nSplitArray
-- end

-- function parse_args(secure_query)
--     args_arr = split(secure_query,'&')
--     local res = {}
--     local args_num = 0
--     for k,v in pairs(args_arr) do
--         arg = split(v,'=')
--         -- if table.getn(arg) == 2 then
--           res[arg[1]]={}
--           res[arg[1]]=arg[2]
--           args_num = args_num + 1
--         -- end
--      end
--      return res,args_num
-- end

-- secure_query = "token=3&expireAt=2&session=1&policy=8&date=2014" -- "a=1&b=2&token=3"
-- local cache_url = "cdncs.static.smart.hainan.edu.cn" .. "/migu/sx/test.jpg"

-- local args_str = ""
--     if secure_query ~= nil then
--         local res, _ = parse_args(secure_query)
--         if res ~= nil then
--             res["token"] = nil
--             res["session"] = nil
--             res["expireAt"] = nil
--             res["policy"] = nil
--             res["date"] = nil
--             for k,v in pairs(res) do
--                 args_str = args_str .. k .. "=" .. v .. "&"
--             end
--             args_str = string.gsub(args_str, ".$", "")
--         end
--     end
--     if args_str ~= "" then
--         cache_url = cache_url .. "?" .. args_str
--     end

-- print("kkkkkk=" .. cache_url)


-- local geEscapeChar = {
--     ["^"] = true,
--     ["$"] = true,
--     ["("] = true,
--     [")"] = true,
--     ["%"] = true,
--     ["."] = true,
--     ["["] = true,
--     ["]"] = true,
--     ["*"] = true,
--     ["+"] = true,
--     ["-"] = true,
--     ["?"] = true
-- }
-- local function split(s,delimiter)
--     local result={}
--     local newdelimiter = delimiter
--     if geEscapeChar[delimiter] then
--         newdelimiter = "%"..delimiter
--     end
--     for match in (s..delimiter):gmatch("(.-)"..newdelimiter) do
--         if match == "" then
--             match = '-'
--         end
--             table.insert(result,match)
--     end
--     return result
-- end

-- local field_45 = "-"

-- local t_field_45 = split(field_45, "$")

-- if t_field_45[1] == nil then t_field_45[1] = "NULL" end
-- -- if t_field_45[2] == nil then t_field_45[2] = "NULL" end

-- t_field_45[2] = t_field_45[2] or "11"
-- print(t_field_45[2])


-- local url = "/wd_r2/jstv/jschengshi/600/index.m3u8?msisdn=1f261800d431ce86e8d6a97313bb6941&mdspid=&spid=699067&netType=4&sid=5500260508&pid=2028597139&timestamp=20250428142207&Channel_ID=0116_0132_10010001005&ProgramID=626064714&ParentNodeID=-99&assertID=5500260508&client_ip=240e:387:45f:e800::389&SecurityKey=20250428142207&promotionId=&mvid=5100186409&mcid=500020&playurlVersion=WX-A1-8.4.1-RELEASE&userid=1778733688&jmhm=1f261800d431ce86e8d6a97313bb6941&videocodec=h264&bean=mgspipad&tid=ipad&conFee=0&sv=10000&ct=www&encrypt=b3eb84936c7aa5e5f813d024e430e4f9"
-- local md5str = (string.match(url, "/(.*)&encrypt.*"))

-- local atslog = nil
-- local index1 = string.find(atslog, "|", 1)
-- print(index1)
-- local k = "X-Z-123"
-- if string.sub(string.lower(k), 1, 4) == "x-z-" then
--     print(k)
-- end
-- local max_age = ""
-- print("tonumber(max_age)=", tonumber(max_age))

-- local max_age = "abc"
-- print("tonumber(max_age=abc)=", tonumber(max_age))

-- local max_age = "1a"
-- print("tonumber(max_age=1a)=", tonumber(max_age))

-- local max_age = "0"
-- print("tonumber(max_age)=0", tonumber(max_age))

-- local max_age = "1"
-- print("tonumber(max_age)=1", tonumber(max_age))

-- local max_age = nil
-- print("tonumber(max_age)=nil", tonumber(max_age))

-- if tonumber(max_age) and tonumber(max_age) >= 0 then
--     print("tonumber(max_age) >= 0")
-- end



local function urlDecode(s)
    s = string.gsub(s, "+", " ")
    s = string.gsub(s, "%%(%x%x)", function(h)
        return string.char(tonumber(h, 16))
    end)
    return s
end

local function parse_args(url_args)
    local args = {}
    -- 匹配键值对（值可为空）
    for key, value in string.gmatch(url_args, "([^&=]*)=([^&=]*)&?") do
        args[key] = (value == "" and nil) or value
    end
    return args
end

local function parse_url_args(url_args)
    local args = {}
    for key, value in string.gmatch(url_args, "([^&=]+)=([^&=]+)") do
        args[key] = value
    end
    return args
end

local function parse_url_args2(url_args)
    local args = {}
    -- 匹配键值对（值可为空）
    for key, value in string.gmatch(url_args, "([^&=]*)=([^&=]*)&?") do
        args[key] = (value == "" and nil) or value
    end
    return args
end


-- local uri_args = "http_user-agent=android&http_referer=&a=1&b=2"
-- local uri_args = "http_user-agent=android%2C1.1&http_referer=&a=1&b=2"
-- local uri_args = "http_User-Agent=android%2C1.1&http_Referer=&a=1&b=2"
-- local uri_args = "http_User-Agent=android%2C1.1&http_Referer=&a=1&b=2&sent_http_access-control-allow-origin=*&sent_http_data="


-- print(string.format("testlg---uri_args       =%s",uri_args))
-- local args = {}
-- if uri_args then
--     uri_args = urlDecode(string.lower(uri_args))
--     print(string.format("testlg---uri_args decode=%s",uri_args))
--     -- args = parse_url_args2(uri_args)
--     args = parse_args(uri_args)
-- end

-- for k, v in pairs(args) do
--     print(string.format("testlg---%s=%s",k, v))
-- end
-- local uri = "/host2/depot/578081/migu/sx/test.jpg"

-- local isNull = {
--     [""] = true,
--     ["nil"] = true,
-- }
-- local uri = "/host2"

-- local host2, depot, depot_id, file = string.match(uri, "^/([^/]*)/([^/]*)/([^/]*)/(.*)")
-- print("host2--=", host2)
--     print("depot--=", depot)
--     print("depot_id--=", depot_id)
--     print("file--=", file)
-- if isNull[host2] or isNull[depot] or isNull[depot_id] or isNull[file] then
--     print("uri fail")
-- else
--     print("url ok")
-- end
-- t_reserved = {}
-- t_reserved[1] = "1"
-- t_reserved[2] = ""
-- t_reserved[3] = "3"
-- local reserved = table.concat(t_reserved, "|")
-- print(reserved)
-- local k = "x-z-via-rspip"
-- if string.sub(string.lower(k), 1, 4) == "x-z-" and k ~= "x-z-via-rspip" then
--     print(k)
-- end




-- local function getUpstreamList(str)
--     local pattern =  "[,:] "  -- "[,:]%s*"
--     local parts = {}
--     if not str then
--         return parts
--     end
--     local from = 1
--     local delim_from, delim_to = string.find(str, pattern, from)
--     while delim_from do
--         table.insert(parts, string.sub(str, from, delim_from - 1))
--         from = delim_to + 1
--         delim_from, delim_to = string.find(str, pattern, from)
--     end
--     table.insert(parts, string.sub(str, from))
--     return parts
-- end

-- local function getOneAddrStatus(addrlistStr, statuslistStr)
--     local addrParts = getUpstreamList(addrlistStr)
--     local statusParts = getUpstreamList(statuslistStr)

--     local addr2xx = ""
--     local status2xx = ""
--     local addr3xx = ""
--     local status3xx = ""
--     local addr4xx = ""
--     local status4xx = ""
--     local addr5xx = ""
--     local status5xx = ""
--     -- 4xx,5xx 取最后一个
--     -- 2xx,3xx 取第一个(3xx再3xx，取第一个3xx忽略后面的)
--     for i, part_tmp in ipairs(statusParts) do
--         local part = tonumber(part_tmp) or 0
--         if part >= 500 and part < 600 then
--             addr5xx = addrParts[i] or addrParts[#addrParts]
--             status5xx = part
--         elseif part >= 400 and part < 500  then
--             addr4xx = addrParts[i] or addrParts[#addrParts]
--             status4xx = part
--         elseif part >= 300 and part < 400 and (status3xx == nil or status3xx == "")   then
--             addr3xx = addrParts[i] or addrParts[#addrParts]
--             status3xx = part
--         elseif part >= 200 and part < 300 and (status2xx == nil or status2xx == "") then
--             addr2xx = addrParts[i] or addrParts[#addrParts]
--             status2xx = part
--         end
--     end
--     if status3xx ~= nil and status3xx ~= "" then
--         return addr3xx,statusParts[#statusParts]
--     end
--     if status2xx ~= nil and status2xx ~= "" then
--         return addr2xx,statusParts[#statusParts]
--     end
--     if status4xx ~= nil and status4xx ~= "" then
--         return addr4xx,statusParts[#statusParts]
--     end
--     if status5xx ~= nil and status5xx ~= "" then
--         return addr5xx,statusParts[#statusParts]
--     end
-- end

-- local function gethostFromUpstream(str)
--     if not str then
--         return "-"
--     end
--     if string.find(str, ".",1, true) then
--         -- ipv4
--         local index = string.find(str,":",1, true)
--         if index then
--             str = string.sub(str,0,index - 1)
--             return str
--         end
--     else
--         --ipv6
--         local index = string.find(str,"]",1, true)
--         if index then
--             str = string.sub(str,0,index)
--             return str
--         end
--     end
--     return str
-- end
-- -- ngx.header["cmcc-up-addr"]
-- -- ngx.header["cmcc-up-status"]
-- local cmcc_up_addr = "10.12.7.141:80"
-- local cmcc_up_status = "206"

-- local addr_tmp,status_tmp = getOneAddrStatus(cmcc_up_addr,cmcc_up_status)
-- print("addr_tmp--1--=",addr_tmp)

-- local addr_tmp_2 = gethostFromUpstream(addr_tmp)

-- local origin_port = string.gsub(addr_tmp, addr_tmp_2.. ":", "")
-- print("addr_tmp--2--=",addr_tmp_2)
-- print("status_tmp--=",status_tmp)
-- print("origin_port--=",origin_port)


-- local uri = "/remove/migu/sx/test.jpg?t=1&remove-=2"
-- -- local _,new_uri =  string.match(uri, "^/([remove%-]*)(.*)")

-- local new_uri = uri:gsub("^/(remove%-[^/]+)/", "/")
-- print(uri .. "||" .. new_uri)

-- local limit_rate_domain = {["gauss-otacostauto-cn.allawnfs.com"]="1"}

-- -- 分时段限速,19-1点限速300k/s,1-8点限速6Mb/s,8-19点限速500k/s;
-- local function limit_rate_by_time()
--     local host = "gauss-otacostauto-cn.allawnfs.com"
--     if not limit_rate_domain[host] or limit_rate_domain[host] ~= 1 then
--         print("000000")
--         return
--     end
    
--     print("1111")

-- end
-- limit_rate_by_time()

-- local uri = "/migu/sx/test.ZIP"
-- if uri:lower():match("%.zip$") then
--     local ua = ""
--     ua = ua:lower()
--     if ua:find("mozilla") or ua:find("go%-http%-client") then
--         print("1111")
--         return
--     end
--     print("00000")
--     return
-- end

-- print("-1")

-- local rate = 0
-- local type = 0
--     local hour = tonumber(os.date("%H"))
--     if hour >= 19 or hour < 1 then       -- 19:00-01:00时段
--         rate = 300 * 1024               -- 300KB/s
--         type = 1
--     elseif hour >= 1 and hour < 8 then   -- 01:00-08:00时段
--         rate = 6 * 1024 * 1024          -- 6MB/s
--         type = 2
--     else                                 -- 08:00-19:00时段
--         rate = 500 * 1024               -- 500KB/s
--         type = 3
--     end
-- print(hour .. "---" ..type)
-- print(os.date())


-- local limit_rate_domain = {
--     ["gauss-otacostauto-cn.allawnfs.com"] = {
--         { start_h = 19, end_h = 1, rate = "300k" },
--         { start_h = 1, end_h = 8, rate = "6M" },
--         { start_h = 8, end_h = 19, rate = "500k" }
--     },
--     ["gauss-sauapkcostauto-cn.allawnfs.com"] = {
--         { start_h = 19, end_h = 1, rate = "300k" },
--         { start_h = 1, end_h = 8, rate = "6M" },
--         { start_h = 8, end_h = 19, rate = "500k" }
--     },
--     ["gauss-compotacostauto-cn.allawnfs.com"] = {
--         { start_h = 21, end_h = 1, rate = "300k" },
--         { start_h = 1, end_h = 2, rate = "3M" },
--         { start_h = 2, end_h = 8, rate = "6M" },
--         { start_h = 8, end_h = 14, rate = "500k" },
--         { start_h = 14, end_h = 18, rate = "300k" },
--         { start_h = 18, end_h = 19, rate = "500k" },
--         { start_h = 19, end_h = 21, rate = "1M" }
--     }
-- }

-- local function limit_rate_by_time(host, current_hour)
--     if not limit_rate_domain[host] then
--         return nil
--     end
--     local rules = limit_rate_domain[host]

--     for _, rule in ipairs(rules) do
--         local start_h = rule.start_h
--         local end_h = rule.end_h
        
--         -- 处理跨天时间段（如23:00-1:00）
--         if start_h > end_h then
--             if current_hour >= start_h or current_hour < end_h then
--                 return  rule.rate
--             end
--         else
--             if current_hour >= start_h and current_hour < end_h then
--                 return rule.rate
--             end
--         end
--     end
-- end

-- local hour = 20
-- local rate = limit_rate_by_time("gauss-otacostauto-cn.allawnfs.com", hour)
-- print(string.format("time: %d rate: %s", hour, rate))
-- hour = 7
-- rate = limit_rate_by_time("gauss-otacostauto-cn.allawnfs.com", hour)
-- print(string.format("time: %d rate: %s", hour, rate))
-- hour = 7
-- rate = limit_rate_by_time("gauss-otacostauto-cn.allawnfs.com", hour)
-- print(string.format("time: %d rate: %s", hour, rate))


-- local check_ua_domain = {
--     ["gauss-otacostauto-cn.allawnfs.com"] = {"mozilla","go%-http%-client"}
-- }

-- local function check_ua(uri, host, ua)
--     if not check_ua_domain[host] then
--         return "1"
--     end

--     if uri:lower():match("%.zip$") then
--         ua = ua:lower()

--         local forbidden_uas = check_ua_domain[host]

--         for _, forbidden_ua in ipairs(forbidden_uas) do
--             if ua:find(forbidden_ua)  then
--                 return "UA_FORBIDDEN"
--             end
--         end
--     end
--     return "2"
-- end

-- local uri = "/migu/test.zip"
-- local ua = "sfsdmozilla"
-- local r = check_ua(uri, "gauss-otacostauto-cn.allawnfs.com", ua)
-- print(string.format("uri: %s ua: %s  result: %s", uri, ua, r))
-- uri = "/migu/test.jpg.zip"
-- ua = "Go-Http-client"
-- local r = check_ua(uri, "gauss-otacostauto-cn.allawnfs.com", ua)
-- print(string.format("uri: %s ua: %s  result: %s", uri, ua, r))
-- ua = "Mozi-http-client"
-- uri = "/migu/test.png"
-- local r = check_ua("", "gauss-otacostauto-cn.allawnfs.com", ua)
-- print(string.format("uri: %s ua: %s  result: %s", uri, ua, r))

-- local ua_domain = {
--     ["update.hihonorcdn.com"]={"moto","redmi","pixel", "hbp%-lx9"}
-- }
-- local host = "update.hihonorcdn.com"
-- local ua= "fsdfsMtofsdfREd111HBP-lx9"

-- local forbidden_uas = ua_domain[host]
-- if forbidden_uas then
--     ua = ua:lower()
--     for _, forbidden_ua in ipairs(forbidden_uas) do
--         if ua:find(forbidden_ua)  then
--             print("UA_FORBIDDEN")
--             return
--         end
--     end
-- end
-- print("OK")

-- local function replace_chars(str)
--     -- 使用gsub进行模式匹配替换
--     str = string.gsub(str, "|", "%%7C")  -- | → %7C
--     str = string.gsub(str, "%$", "%%24") -- $ → %24（$需转义）
--     return str
-- end

-- local t= "123|abc$def" -- "abc|def$ghe"
-- print(replace_chars(t))

-- local tmp_body_ts_relative_uri = "20240418185719sim121201/000/121201-000-1174.ts"
--             -- ngx.log(ngx.ERR, "testlog--tmp_body_ts_relative_uri=", tmp_body_ts_relative_uri)
--             local ts_path_pre = string.match(tmp_body_ts_relative_uri, "(.*)/.*/.*$")
-- print(ts_path_pre) 

--Accept请求头值不包含image/webp的增加改写规则：^/(.+).(png|jpg|jpeg|bmp|gif)\.(.*) -> /$1.$2,示例:/123.png.webp改写为123.png,/123.png.456.webp改写为123.png
-- local uri = "/123.png.456.webp"
-- local accept_value = "image/jpeg,image/png"
-- if accept_value and not accept_value:lower():find("image/webp") then
--     print("accept===", accept_value)
--     -- local base_path, ext = uri:match("^/(.+)%.(png|jpg|jpeg|bmp|gif)%..+$")
--     local base_path, ext,_ = string.match(uri, "^/(.+)%.(png|jpg)%.(.*)$")
--     print(base_path)
--     print(ext)

--     -- if base_path and ext then
--     --     local original_url = "/" .. base_path .. "." .. ext
--     --     print(original_url) -- 输出: /test.image.png
--     -- end

-- end

local cache_key_pre_list = {  
    ["fs.mvsso.ali.kugou.com"]="fsmvandroid.kugou.com",
    ["fs.mvsso.hw.kugou.com"]="fsmvandroid.kugou.com",
    ["fs.youthandroid2.kugou.com"]="fsandroid.kugou.com",
    ["fsandroid.kugou.com"]="fsandroid.kugou.com",
    ["fsdg360.hw.kugou.com"]="fsandroid.kugou.com",
    ["fsios.kugou.com"]="fsandroid.kugou.com",
    ["fsmvandroid.kugou.com"]="fsmvandroid.kugou.com",
    ["fspc.kugou.com"]="fsandroid.kugou.com",
    ["fsvippc.kugou.com"]="fsandroid.kugou.com",
    ["imgewebp.kugou.com"]="",
    ["webfile.yun.kugou.com"]=""}

local function get_path(uri, is_refresh)
    local result_path = uri

    local id3, p1, p2, _ = string.match(uri, "^/v3/([^/]+)/([^/]+)/([^/]+)$")
    if id3 and p1 and p2 then
        -- 刷新
        if is_refresh then
            result_path = string.format("/v3/%s/%s/%s", id3, p1, p2)
            print("/v3 刷新")
            return result_path
        end
    end

    local _, _, id3, p1, p2, _ = string.match(uri, "^/([^/]+)/([^/]+)/v3/([^/]+)/([^/]+)/([^/]+)(.*)")
    if id3 and p1 and p2 then
        result_path = string.format("/v3/%s/%s/%s", id3, p1, p2)
        print("/t/s/v3/ ")
        return result_path
    end

    local _, _, _, path2 = string.match(uri, "^/([^/]+)/([^/]+)/v2/([^/]+)(.*)")
    if path2 then
        result_path = string.format("%s", path2)
        print("/t/s/v2/ ")
        return result_path
    end

    local date_time, sign, path1 = string.match(uri, "^/([^/]+)/([^/]+)(.*)")
    if not date_time or string.len(date_time) ~= 12 or not sign or string.len(sign) ~= 32 then
        -- 刷新
        if is_refresh then
            result_path = uri
            print("/p 刷新 ")
            return result_path
        end    
    elseif path1 then
        result_path = string.format("%s", path1)
        print("/t/s/p ")
        return result_path
    end
    

    return result_path
end

local function get_cache_key(uri, is_refresh, host)
    local path = get_path(uri, is_refresh)
    local cache_key_pre = cache_key_pre_list[host]
    if not cache_key_pre or cache_key_pre == "" then
        cache_key_pre = host
    end
    local cache_key = string.format("/%s/Range:slice_range%s", cache_key_pre, path)
    return cache_key
end


local uri = "/202509090930/a67cf34d4de7d12f53c1d29d797b99f5/a1_b2_c3_s3877055665.mp3"

local uri2 = "/202509090930/a67cf34d4de7d12f53c1d29d797b99f5/v2/d18c7e16fd8b3cc70ecb59ff07821f69/a1_b2_c3_s3877055665.mp3"

local uri3 = "/202509090930/a67cf34d4de7d12f53c1d29d797b99f5/v3/d18c7e16fd8b3cc70ecb59ff07821f69/yp/full/a1_b2_c3_s3877055665.mp3"

local uri4 = "/v3/d18c7e16fd8b3cc70ecb59ff07821f69/yp/full"

local uri5 = "/a1_b2_c3_s3877055665.mp3"

local uri6 = "/202/a67cf34d4de7d12f53c1d29d797b99f5/a1_b2_c3_s3877055665.mp3"

-- local cache_key = get_cache_key(uri6, true, "fsvippc.kugou.com222")
-- print(cache_key)

-- local t= ""
-- if t  == "" then
--     print("1")
-- end


-- function getextension_uri(filename)
--     -- return filename:match(".+(/.+/.+%.%w+)$")
--     return filename:match("(/[^/]+/[^/]+%.[%w]+)$")
-- end

-- local cached_file_uri={["/4QbVtADbnLVIc/d.FxJzG50F.js"]=1,["/4QbVtADbnLVIc/c.FxJzG50F.js"]=1,["/4QbVtADbnLVIc/4rJFe6jNL52p.js"]=1,["/4QbVtADbnLVIc/jW39ezbWPr.js"]=1}


-- local file_uri = getextension_uri("/4QbVtLVIcADbn/MjUyNzc30k.png");
-- print(file_uri)

-- local file_uri = getextension_uri("/c/4QbVtADbnLVIc/d.FxJzG50F.js");
-- print(file_uri)



-- local file_uri = getextension_uri("/b/c/4QbVtADbnLVIc/c.FxJzG50F.js");
-- print(file_uri)


-- local file_uri = getextension_uri("a/bb/c/d/e/f/b/c/4QbVtLVIcADbn/MjUyNzc30k.png");
-- print(file_uri)


-- local file_uri = getextension_uri("/b/c/4QbVtLVIcADbn/MjUyNzc30k.png");
-- print(file_uri)

-- local cache_control_value = ""
-- if "/4QbVtLVIcADbn/MjUyNzc30k.png" == file_uri then
--     cache_control_value = "no-cache";
-- else 
--  cache_control_value = "max-age=86400";
-- end



local status_time = {}

local function prase(str)
    if nil == str or "" == str then
        return
    end

    local sp = string.find(str, ":")
    if nil == sp then
        return
    end
    
    local cache_time = string.sub(str, sp + 1, -1)
    if nil == cache_time or "" == cache_time then
        return
    end
    
    local status_str = string.sub(str, 1, sp - 1)
    
    local b_index = 1
    local e_index = 1
    while b_index < #status_str do
        e_index = string.find(status_str, "|", b_index)
        if nil ~= e_index then
            local temp_str = string.sub(status_str, b_index, e_index - 1)
            status_time[temp_str] = cache_time
            b_index = e_index + 1
        else
            e_index = #status_str
            local temp_str = string.sub(status_str, b_index, e_index)
            status_time[temp_str] = cache_time
            break
        end
    end

end

local function is_ex_status(initial_status, status)
    local initial_status_num = tonumber(initial_status)
    if initial_status_num and initial_status_num >= 3 and 304 ~= status then
        return true
    end

    return false
end

local function handler_exception_cache(str)
    -- "400|200|:5;404:5;500:5";
    if nil == str or "" == str then
        return
    end

    local b_index = 1
    local e_index = 1
    while b_index < #str + 1 do
        e_index = string.find(str, ";", b_index)
        if e_index ~= nil then
            local temp_str = string.sub(str, b_index, e_index - 1)
            prase(temp_str)
            b_index = e_index + 1
        else
            local temp_str = string.sub(str, b_index, #str)
            prase(temp_str)
            break
        end

    end

end

-- handler_exception_cache("2XX|304:31536000;4XX|5XX:10;3XX:0")
-- for k, v in pairs(status_time) do
--     print(k, v)
-- end
-- local status = "304"
-- local initial_status = string.sub(tostring(status), 1, 1)
-- local reg_st = initial_status .. "XX"

-- local status_cache_time = status_time[tostring(status)] or status_time[reg_st]
-- print("status_cache_time=" .. status_cache_time)

-- local function encode_uri(s)
--     s = string.gsub(s, "([^%w%.%-_~%/ ])", function(c) return string.format("%%%02X", string.byte(c)) end)
--     return string.gsub(s, " ", "+")
-- end

-- local function process_uriV1(uri)
--     local file = string.match(uri, "^.*/([^/]*)$")
--     print("file=" .. file)
--     local path = string.gsub(uri, file .. "$", "")
--     print("path=" .. path)
--     return path .. encode_uri(file)
-- end

-- local function escape_pattern(s)
--     -- 将模式特殊字符转义
--     return (string.gsub(s, "([%^%$%(%)%%%.%[%]%*%+%-%?])", "%%%1"))
-- end

-- local function process_uri(uri)
--     local file = string.match(uri, "^.*/([^/]*)$")
--     if not file then
--         return uri
--     end
--     print(file)
--     local escaped_file = escape_pattern(file)
--     print("escaped_file=" .. escaped_file)
--     local path = string.gsub(uri, escaped_file .. "$", "")  -- 确保只替换末尾的文件名
--     print("path=" .. path)
--     -- return path .. encode_uri(file)
--     if uri == nil or "" == uri then
--         return ""
--     end
    
--     return encode_uri(uri)
-- end
-- local  uri = "/resource/beta/apk/20231114094513/MiguPlay-V3.69.1.1_miguzsj中国.mp4"
-- local path = process_uriV1(uri)
-- -- local path = process_uri(uri)
-- print(path)


-- handler_exception_cache("400|200|:5;404:5;500:2")

-- local status = "300"
-- local reg_st = "3XX"
-- local status_cache_time = status_time[tostring(status)] or status_time[reg_st]
-- local origin_cache_precedence = nil or ""  -- 是否源站缓存优先
-- local origin_cache_control = "1s" or ""

-- print(status_cache_time)
-- if ( origin_cache_precedence ~= "true" and status_cache_time ~= nil )
--     or (origin_cache_precedence == "true" and origin_cache_control == "" and status_cache_time ~= nil)
-- then
--         print("ats set cache")
--         return
--     end

--     print("111")

--     "./oppo/store.heytapimage.com/CCS/ats/remap.config  /oppo修改回源/store.heytapimage.com/L2/etc/trafficserver/"
--     "./oppo/store01.heytapimage.com/CCS/ats/remap.config  /oppo修改回源/store01.heytapimage.com/L2/etc/trafficserver/"


-- local a = nil
-- if a ~= "" then
--     print("不鉴权")
--     return
-- end
-- print("鉴权")
-- local t =  string.format("TCP_MISS,TCP_MISS,%s", a or "TCP_MISS")
-- print(t)


-- local geEscapeChar = {
--     ["^"] = true,
--     ["$"] = true,
--     ["("] = true,
--     [")"] = true,
--     ["%"] = true,
--     ["."] = true,
--     ["["] = true,
--     ["]"] = true,
--     ["*"] = true,
--     ["+"] = true,
--     ["-"] = true,
--     ["?"] = true
-- }
-- local function split(s,delimiter)
--     local result={}
--     local newdelimiter = delimiter
--     if geEscapeChar[delimiter] then
--         newdelimiter = "%"..delimiter
--     end
--     for match in (s..delimiter):gmatch("(.-)"..newdelimiter) do
--         if match == "" then
--             match = '-'
--         end
--             table.insert(result,match)
--     end
--     return result
-- end
-- local via21 =""
-- local http_x_server_addr = "a,bbb"
-- local strs = split(http_x_server_addr, ",")
-- if #strs == 3 then
--     via21 = string.format("%s,%s", strs[1], strs[2])
-- end
-- print(via21)

-- local ip =  "10.1.5.205"
-- local str = "10.1.204,10.1.5.205,10.12.7.141"
-- local pattern = ip:gsub("%.", "%%.")
-- local res =str:match("(.*)"..pattern) or str
-- print(res)
-- local res = ""
-- local hit_type_tmp = "hit"
-- local http_x_server_addr = "10.1.5.205"
-- local server_addr = "10.1.5.205"
-- if hit_type_tmp == "hit" and (http_x_server_addr ~= nil) and http_x_server_addr ~= "" then

--     local pattern = server_addr:gsub("%.", "%%.")
--     print("111" ..pattern)
--     local ips =http_x_server_addr:match("(.*)"..pattern) or http_x_server_addr
--     if ips == "" then
--         res = server_addr
--     else
--         res = string.format("%s%s", ips, server_addr)
--     end
-- end
-- print(res)


-- local function check_file_type(uri)
--     --gif;png;bmp;jpeg;jpg
--     local file_suffix = {["gif"] = 1, ["png"] = 1, ["bmp"] = 1, ["jpeg"] = 1, ["jpg"] = 1}
--     local suffix = string.match(uri, "^[^%.]*%.(.*)$")
--     if file_suffix[suffix] == 1 then
--         return true
--     end
--     return false
-- end

-- print(check_file_type("/path/to/.gif.image.php"))
-- print(check_file_type("/path/to/image.ts"))

-- local dhfs_cache_key_and_time = {
--     ["^/userfiles/cms/video_inner_page"] = 300,      -- 5分钟短缓存
--     ["^/userfiles/cms/videolandingpage/browser"] = 3600, -- 1小时缓存
--     ["^/userfiles/pager_render"] = 600,              -- 10分钟缓存
--     ["^/magazine/api"] = 0,                          -- 禁用缓存
-- }

-- local function get_cache_key_and_time(uri, args_table)
--     local host = "dhfs.heytapimage.com"
--     -- local args = ngx.var.args
--     -- local args_table = ngx.req.get_uri_args()
--     -- 构建基础缓存键
--     local cache_key = string.format("/%s%s", host, uri)
--     local cache_time = 0
--     if host == "dhfs.heytapimage.com" then 
--         -- 缓存key有特殊参数
--         if args_table["x-oss-process"] then
--             cache_key = string.format("%s?x-oss-process=%s", cache_key, args_table["x-oss-process"])
--         end
--         -- 查找匹配的缓存时间
--         cache_time = 7776000 -- 默认90天缓存
--         for path, time in pairs(dhfs_cache_key_and_time) do
--             if string.match(uri, path) then
--                 cache_time = time
--                 break
--             end
--         end
--         -- 处理SN_choice特殊情况
--         if args_table["SN_choice"] and tonumber(args_table["SN_choice"]) == 9 then
--             cache_time = 0
--         end
--     end
--     return cache_key, cache_time
-- end


-- local cache_key, cache_time = get_cache_key_and_time("/userles/pager_render", {["x-oss-process"] = "image/resize,w_800", ["SN_choice"] = "0"})

-- print(cache_key, cache_time)

-- local x = "a"
-- local cache_key = "host?"
-- if a then
--     cache_key = string.format("%s&SN_choice=9", cache_key)
-- else
--     cache_key = string.format("%s?SN_choice=9", cache_key)
-- end
-- print("111"..cache_key)
-- local cache_key2= "host2"
-- cache_key2 = cache_key2 .. (a and "&" or "?") .. "SN_choice=9"
-- print("222"..cache_key2)

-- local function encode_uri(s)
--     s = string.gsub(s, "([^%w%.%-_~%/ ])", function(c) return string.format("%%%02X", string.byte(c)) end)
--     return string.gsub(s, " ", "+")
-- end

-- local function escape_pattern(s)
--     -- 将模式特殊字符转义
--     return (string.gsub(s, "([%^%$%(%)%%%.%[%]%*%+%-%?])", "%%%1"))
-- end

-- local function process_uri(uri)
--     local file = string.match(uri, "^.*/([^/]*)$")
--     if not file then
--         return uri
--     end
--     print("file:", file)
--     local escaped_file = escape_pattern(file)
--     print("escaped_file:", escaped_file)

--     local path = string.gsub(uri, escaped_file .. "$", "")  -- 确保只替换末尾的文件名
--     print("path:", path)
--     -- uri = path .. escaped_file
--     -- return path .. encode_uri(file)
--     -- print("uri:", uri)

--     return encode_uri(uri)
-- end
-- local res = process_uri("/zy/file/input/2025-07-23/11/f363123528851263488_unzip/贴纸-生活/标题_th.png")
-- print(res)  

-- local function encode_uri(s)
--     s = string.gsub(s, "([^%w%.%-_~%/ ])", function(c) return string.format("%%%02X", string.byte(c)) end)
--     return string.gsub(s, " ", "+")
-- end

-- local res = encode_uri("/zy/file/input/2025-07-23/11/f363123528851263488_unzip/贴纸-生活/底边3_th.png")
-- print(res)  

-- local function checkIdx(idx)
--     -- 检查是否为数字类型
--     -- print("type(idx)=",type(idx))
--     local idx_num = tonumber(idx)
--     if idx_num then
--         -- 进一步判断是否为整数
--         if idx_num == math.floor(idx_num) then
--             -- 整数
--             return true
--         else
--             --浮点数
--             print( "idx_num is floor")
--             return false
--         end
--     else
--         print( "idx_num is nil")
--         return false -- TODO
--     end
-- end


-- print("123==" , checkIdx("123"))
-- -- print("123.45==" , checkIdx("123.45"))
-- -- print("abc==" , checkIdx("abc"))
-- -- print("num 123==" , checkIdx(123))
-- -- print("float 123.45==" , checkIdx(123.45))

-- print("a123==" , checkIdx("a123"))
-- print("123bb==" , checkIdx("123bb"))
-- print("1c23==" , checkIdx("1c23"))
-- local function getUpstreamList(str)
--     local pattern =  "[,:] "  -- "[,:]%s*"
--     local parts = {}
--     if not str then
--         return parts
--     end
--     local from = 1
--     local delim_from, delim_to = string.find(str, pattern, from)
--     while delim_from do
--         table.insert(parts, string.sub(str, from, delim_from - 1))
--         from = delim_to + 1
--         delim_from, delim_to = string.find(str, pattern, from)
--     end
--     table.insert(parts, string.sub(str, from))
--     return parts
-- end
-- local res = getUpstreamList("302 : 200")
-- for k, v in ipairs(res) do
--     print(k, v)
-- end


-- local ngx_re_match = string.match -- ngx.re.match

-- -- 文件名特殊处理
-- local function handler_filename(filename)
--     -- 处理情况1: .ts.m3u8文件后缀只取到.ts
--     local m = ngx_re_match(filename, "^(.*)\\.ts\\.m3u8$", "ijo")
--     if m and m[1] then
--         return m[1] .. ".ts"
--     end
    
--     -- 处理情况2: .ts|.mp4去除文件后缀前的数字分片
--     m = ngx_re_match(filename, "^(.*)\\.(\\d+)\\.(ts|mp4)$", "ijo")
--     if m and m[1] and m[3] then
--         hls_mp4_flag = true
--         return m[1] .. "." .. m[3]
--     end
    
--     -- 处理情况3: mp4视频文件不做特殊处理，直接取文件名
--     -- 如果以上两种情况都不匹配，则直接返回原始文件名
--     return filename
-- end

-- local function escape_pattern(s)
--     -- 转义 Lua 模式中所有特殊字符：. * + ? - ^ $ ( ) [ ] { } %
--     return s:gsub("([%.%*%+%?%-%^%$%(%)%[%]%{%}%%])", "%%%1")
-- end

-- local function get_cache_key(_uri, _request_uri)
--     local cache_pre = string.format("%s/Range:slice_range", "finder.com")
--     local cache_path = ""

--     local vkey = nil
--     local filename = nil
--     local path_vkey = nil
--     -- local m = ngx_re_match(_uri, "^.*/([^/]+)/([^/?]+)(\\?.*)?$", "jo")
--     -- if m then
--     --     vkey = m[1]
--     --     filename = m[2]
--     -- end
--     local m = ngx_re_match(_uri, "^.*/([^/?]+)(\\?.*)?$", "jo")
--     if m then
--         filename = m[1]
--     end

--     m = ngx_re_match(_uri, "^.*/([^/]+)/([^/?]+)(\\?.*)?$", "jo")
--     if m then
--         path_vkey = m[1]
--     end

--     if ngx_re_match(filename, "\\.mp4$", "ijo") then
--         vkey = ngx_var.arg_vkey
--         cache_path = _request_uri -- mp4带参缓存，去除参数vkey
--     else
--         vkey = path_vkey
--         -- 做一致性hash的时候，去除path中的加密串
--         local escaped_path = escape_pattern(path_vkey)
--         local pattern = "/" .. escaped_path .. "/"
--         ngx.var.uri_args = string.gsub(_uri, pattern, "/")

--         if ngx_re_match(filename, "\\.xml$", "ijo") then
--             -- xml 去参数缓存，再去除path里的vkey
--             cache_path = ngx.var.uri_args
--         else
--             -- 其他带参数缓存
--             local escaped_path = escape_pattern(path_vkey)
--             local pattern = "/" .. escaped_path .. "/"
--             string.gsub(_request_uri, pattern, "/")
--         end
        
--     end

--     if hls_mp4_flag then
--         vkey = path_vkey
--         -- 做一致性hash的时候，去除path中的加密串
--         local escaped_path = escape_pattern(path_vkey)
--         local pattern = "/" .. escaped_path .. "/"
--         ngx.var.uri_args = string.gsub(_uri, pattern, "/")
--     end

--     -- ngx.var.cache_key = string.format("%s%s", cache_pre, cache_path)
--     local cache_key = string.format("%s%s", cache_pre, cache_path)

--     print("---"..cache_key)
-- end


-- -- B__PhtCuzrLhjVCzHFu09EgnZIU6izBg0tSkV-oAwkAPR3H88b18kGI_Uj0FhEYgcweR6Xhqy_phA4Igzg0K_xS8M0B60SoLV99xh4kxjagaQYrB3BT0cSMPXEf876oexKJX6eijIXEZ3dNXh5XnU6B-8ZlndjP1VrtzO6C6O2YxEhsyryvVOK3BbHRAKsUwXp/svp_50112/Z-oaicN1YGtmJkmvolwtfudoKs-ds4E6q8P-im_FwYQag42mWJnBfTLgqYUE4SrVIUXjmwAnEUrmM4YXGT5sfQ94b5Ki9S_8oSsiKtAPqccuhk563nbsBbadypsHdY4zwlb7ZdpDiX9hAdRzZXx9Len-hd7-Tk4VBKof0vua6uha-MUUi9QapilMzIn-n6Q56dgHA4L7KBtSGhXPkc-WGOYaat4oUYZaVb7CaK1HKfWYiL5YHrToVQ/084_gzc_1000102_0b53raaaaaaavqaivcaannu4bcgdaccaabca.f323024.3.ts?cdncode=%252F18907E7BE0798990%252F&index=84&start=844200&end=855760&brs=16959856&bre=17701139&ver=4&token=d3a38ed3284f3213fec881a91be3fd5a&cost=low&cdn_strategy=0
-- local uri = "/B__PhtCuzrLhjVCzHFu09EgnZIU6izBg0tSkV-oAwkAPR3H88b18kGI_Uj0FhEYgcweR6Xhqy_phA4Igzg0K_xS8M0B60SoLV99xh4kxjagaQYrB3BT0cSMPXEf876oexKJX6eijIXEZ3dNXh5XnU6B-8ZlndjP1VrtzO6C6O2YxEhsyryvVOK3BbHRAKsUwXp/svp_50112/Z-oaicN1YGtmJkmvolwtfudoKs-ds4E6q8P-im_FwYQag42mWJnBfTLgqYUE4SrVIUXjmwAnEUrmM4YXGT5sfQ94b5Ki9S_8oSsiKtAPqccuhk563nbsBbadypsHdY4zwlb7ZdpDiX9hAdRzZXx9Len-hd7-Tk4VBKof0vua6uha-MUUi9QapilMzIn-n6Q56dgHA4L7KBtSGhXPkc-WGOYaat4oUYZaVb7CaK1HKfWYiL5YHrToVQ/084_gzc_1000102_0b53raaaaaaavqaivcaannu4bcgdaccaabca.f323024.3.ts" 
-- -- ngx.var.uri
-- local request_uri = "B__PhtCuzrLhjVCzHFu09EgnZIU6izBg0tSkV-oAwkAPR3H88b18kGI_Uj0FhEYgcweR6Xhqy_phA4Igzg0K_xS8M0B60SoLV99xh4kxjagaQYrB3BT0cSMPXEf876oexKJX6eijIXEZ3dNXh5XnU6B-8ZlndjP1VrtzO6C6O2YxEhsyryvVOK3BbHRAKsUwXp/svp_50112/Z-oaicN1YGtmJkmvolwtfudoKs-ds4E6q8P-im_FwYQag42mWJnBfTLgqYUE4SrVIUXjmwAnEUrmM4YXGT5sfQ94b5Ki9S_8oSsiKtAPqccuhk563nbsBbadypsHdY4zwlb7ZdpDiX9hAdRzZXx9Len-hd7-Tk4VBKof0vua6uha-MUUi9QapilMzIn-n6Q56dgHA4L7KBtSGhXPkc-WGOYaat4oUYZaVb7CaK1HKfWYiL5YHrToVQ/084_gzc_1000102_0b53raaaaaaavqaivcaannu4bcgdaccaabca.f323024.3.ts?cdncode=%252F18907E7BE0798990%252F&index=84&start=844200&end=855760&brs=16959856&bre=17701139&ver=4&token=d3a38ed3284f3213fec881a91be3fd5a&cost=low&cdn_strategy=0"
-- --ngx.var.request_uri



-- set $cache_key "$host/Range:slice_range$uri";
-- 动态文件.php;.jsp;.asp;.aspx 不缓存，
-- .xml文件后缀  忽略参数缓存1800秒
-- 除动态、xml其余所有文件包括url参数缓存31536000秒
-- get_cache_key(uri, request_uri)


-- -- 获取原始请求URI
-- local request_uri = "/migu/sx/test.jpg?b=3&vkey=1&a=2"

-- -- 去除vkey参数并处理多余的分隔符
-- local cleaned_uri = request_uri:gsub("([?&])vkey=[^&]*", "")

-- -- 处理可能出现的?&情况
-- cleaned_uri = cleaned_uri:gsub("[?&]+", "?", 1)
-- cleaned_uri = cleaned_uri:gsub("?$", "")

-- -- 最终清理：如果以?结尾且后面没有参数，去除?
-- if cleaned_uri:sub(-1) == "?" then
--     cleaned_uri = cleaned_uri:sub(1, -2)
-- end

-- -- 输出处理结果（可根据实际需求使用）
-- print("原始URI: " .. request_uri)
-- print("处理后URI: " .. cleaned_uri)

-- -- 或者设置回ngx.var（如果需要修改请求）
-- -- ngx.var.request_uri = cleaned_uri
-- local str = 'https://findertt.video.qq.com.60.cdnhwcqir15.com/tx/file1_cdn_returned?a=ddd&token=Cvvj5Ix3eexZiajDdmtxmMEB25durgo0jjMtkfsY1wpLIiaqgRbBX2Xg4FUTQqyAKmNnuFvnUkIgVgWqJKkkaXeH6hiabgMInqD5pwFBNwnvx7JspGBEBzkvHkbcXyhjmfb&encfilekey=Cvvj5Ix3eez3Y79SxtvVL0L7CkPM6dFibFeI6caGYwFEtDXfWjlPoibde2kEoHPBGh8tDoSQeGx9pRrMaRZNohuy1ia4gSg680WBUF6lFSpZs30MOrYfwgjqjau1EsmdvksiaYwZ63HpHEm5elLO30icMyg&idx=1&basedata=CAASBnhXVDE1NhoGeFdUMTU2GgZ4V1QxNTc&sign=wFaiRDwQrHgk7xw6YWLi3Fxlhd2bmg-k37Tcz5U-s7LwmqTaNqZbTG9EkQTBmBt6YrNwyyhBaWDw7lR31xyv9Q&X-snsvideoflag=xWT149'
-- print('----'..#str)


-- local _request_uri = "/internel_tencent_mp4/migu/sx/test.jpg?t=100&start=100&end=500"

-- local function remove_param_vkey(_request_uri, param)
--     local cleaned_uri = _request_uri:gsub("([?&])".. param .. "=[^&]*", "")
--     -- 处理可能出现的?&情况
--     cleaned_uri = cleaned_uri:gsub("[?&]+", "?", 1)
--     cleaned_uri = cleaned_uri:gsub("?$", "")
--     -- 最终清理：如果以?结尾且后面没有参数，去除?
--     if cleaned_uri:sub(-1) == "?" then
--         cleaned_uri = cleaned_uri:sub(1, -2)
--     end

--     return cleaned_uri
-- end

-- local cache_path = remove_param_vkey(_request_uri, "start")
--     cache_path = remove_param_vkey(cache_path, "end")
-- print("处理后URI: " .. cache_path)

local addr_tmp_1 = " 10.12.7.141 :80 "
addr_tmp_1 = string.gsub(addr_tmp_1, "%s+", "")
local p = addr_tmp_1:match(":(%d+)$")
print(string.format("---%s", p))