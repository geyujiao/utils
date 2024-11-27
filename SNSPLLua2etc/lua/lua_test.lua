
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

function getlast(list)
    if not list then
        return "-"
    else
        local i = 0
        local index = 1
        while true do
            i = string.find(list,"[,:] ",i+1)
            if not i then
                break
            end
            index = i + 2
        end
        return string.sub(list,index,#list)
    end
end
local function getUpstreamList(str)
    local pattern =  "[,:] "  -- "[,:]%s*"
    local parts = {}
    if not str then
        return parts
    end
    local from = 1
    local delim_from, delim_to = string.find(str, pattern, from)
    while delim_from do
        table.insert(parts, string.sub(str, from, delim_from - 1))
        from = delim_to + 1
        delim_from, delim_to = string.find(str, pattern, from)
    end
    table.insert(parts, string.sub(str, from))
    return parts
end

local function gethostFromUpstream(str)
    if not str then
        return "-"
    end
    if string.find(str, ".",1, true) then
        -- ipv4
        local index = string.find(str,":",1, true)
        if index then
            str = string.sub(str,0,index - 1)
            return str
        end
    else
        --ipv6
        local index = string.find(str,"]",1, true)
        if index then
            str = string.sub(str,0,index)
            return str
        end
    end
    return str
end

local function getOneAddrStatus(addrlistStr, statuslistStr)
    local addrParts = getUpstreamList(addrlistStr)
    local statusParts = getUpstreamList(statuslistStr)

    local addr2xx = ""
    local status2xx = ""
    local addr3xx = ""
    local status3xx = ""
    local addr4xx = ""
    local status4xx = ""
    local addr5xx = ""
    local status5xx = ""
    -- 4xx,5xx 取最后一个
    -- 2xx,3xx 取第一个(3xx再3xx，取第一个3xx忽略后面的)
    for i, part_tmp in ipairs(statusParts) do
        local part = tonumber(part_tmp) or 0
        if part >= 500 and part < 600 then
            addr5xx = addrParts[i] or addrParts[#addrParts]
            status5xx = part
        elseif part >= 400 and part < 500  then
            addr4xx = addrParts[i] or addrParts[#addrParts]
            status4xx = part
        elseif part >= 300 and part < 400 and (status3xx == nil or status3xx == "")   then
            addr3xx = addrParts[i] or addrParts[#addrParts]
            status3xx = part
        elseif part >= 200 and part < 300 and (status2xx == nil or status2xx == "") then
            addr2xx = addrParts[i] or addrParts[#addrParts]
            status2xx = part
        end
    end
    if status3xx ~= nil and status3xx ~= "" then
        return addr3xx,status3xx
    end
    if status2xx ~= nil and status2xx ~= "" then
        return addr2xx,status2xx
    end
    if status4xx ~= nil and status4xx ~= "" then
        return addr4xx,status4xx
    end
    if status5xx ~= nil and status5xx ~= "" then
        return addr5xx,status5xx
    end
end

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
local need_cache = {"js", "css", "jpg", "jpeg", "png", "gif", "woff"}
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


local ua = "dasfsMozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/"
-- local res = string.find(ua, "Mozilla/5.0Windows NT 10.0Win64x64AppleWebKit537.36KHTMLlike GeckoChrome",1, true)
local res = string.match(ua, "^Mozilla/5.0 %(Windows NT 10.0; Win64; x64%) AppleWebKit/537.36 %(KHTML, like Gecko%) Chrome/")

if res ~= nil then
    print("fsdfsfs")
    
    return
end
print("2222")

