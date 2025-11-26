local _M = {
    _VERSION = "0.0.1"
}

--  字节所有视频域名回源映射表
local domain_mainos_map = {
    ["v5-e-mc.douyinvod.com"] = {
        mainos = "https://mainos.bytedance.com",
        secondos = "https://secondos.bytedance.com"
    },
    ["v5-f-mc.douyinvod.com"] = {
        mainos = "https://mainos.bytedance.com",
        secondos = "https://secondos.bytedance.com"
    }
}

-- 查询特定域名
function _M.find_origin_server(req_server)
    local origin_server_map = domain_mainos_map[req_server]
    if origin_server_map then
        return origin_server_map.mainos, origin_server_map.secondos
    else
        return nil,nil
    end
end

return _M