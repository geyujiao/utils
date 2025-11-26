local domain_mainos = require "luamap"

local req_server = "v5-e-mc.douyinvod.com"
local mainos, secondos = domain_mainos.find_origin_server(req_server)
if mainos then
    print( "mainos: ", mainos)
else
    print( "not found")
end
if secondos then
    print( "secondos: ", secondos)
else
    print( "not found")
end

-- 备注vcm不能为空
-- origin_type=默认 (原来逻辑) 表示 L3 > vcm
-- origin_type=1 (切量完成后) 表示 vcm > L3
-- origin_type=2 (切量过程中) 表示 回源映射表 > vcm
function get_mainos()
    local origin_type = ngx.var.origin_type
    if origin_type == "1" then
        return
    
    elseif origin_type == "2" then
        local mainos, _ = domain_mainos.find_origin_server(req_server)
        if mainos and "" ~= mainos then 
            local ret,scheme,domain,port = parse_urladdr(mainos)	
            if ret then
                ngx.var.mainsrcdomain = domain
                ngx.var.mainsrcport = port
                ngx.var.mainsrcscheme = scheme
                ngx.var.mainsrchost = domain
            end
        end
        return
    
    else
        local mainos_l3 = ngx.req.get_headers()["CM-os-mainaddr"]
        if mainos_l3 and "" ~= mainos_l3 then 
            local ret,scheme,domain,port = parse_urladdr(mainos_l3)	
            if ret then
                ngx.var.mainsrcdomain = domain
                ngx.var.mainsrcport = port
                ngx.var.mainsrcscheme = scheme
			    host = ngx.req.get_headers()["CM-os-mainhost"] or domain
                ngx.var.mainsrchost = host
            end
        end
        return

    end
end


local ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost=get_mainos("1")

print(string.format("1---%s,%s,%s,%s", ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost))


local ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost=get_mainos("2")

print(string.format("2---%s,%s,%s,%s", ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost))


local ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost=get_mainos("3")

print(string.format("3---%s,%s,%s,%s", ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost))


local ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost=get_mainos("")

print(string.format("---%s,%s,%s,%s", ngx.var.mainsrcdomain,ngx.var.mainsrcport,ngx.var.mainsrcscheme,ngx.var.mainsrchost))

