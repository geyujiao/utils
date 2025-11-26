
local crypto = require("crypto")

-- 自定义解密加载器
local function encrypted_loader(name)
    local filename = string.gsub(name, "%.", "/") .. ".luac"
    
    -- 检查文件是否存在
    local file = io.open(filename, "rb")
    if not file then return nil end
    
    -- 读取加密数据
    local iv = file:read(16)  -- 读取初始化向量
    local encrypted_data = file:read("*a")
    file:close()
    
    -- 解密数据
    local key = "abcdefghijklmnop"  -- 必须与加密时使用的密钥一致
    local aes = crypto.cipher.new("aes-128-cbc", key, iv)
    local decrypted_data = aes:decrypt(encrypted_data)
    
    -- 移除填充
    local padding_length = decrypted_data:byte(-1)
    decrypted_data = decrypted_data:sub(1, -padding_length-1)
    
    -- 加载解密后的字节码
    local func, err = loadstring(decrypted_data)
    if not func then
        error("无法加载解密后的字节码: " .. (err or "未知错误"))
    end
    
    return func
end

-- 将自定义加载器插入到package.loaders中
table.insert(package.loaders, 2, encrypted_loader)

-- 测试函数
local function test_encrypted_module()
    -- 加载加密模块
    local success, module = pcall(require, "encrypted_module")
    if success then
        print("加密模块加载成功!")
        module.run()
    else
        print("加密模块加载失败: " .. module)
    end
end

-- 运行测试
test_encrypted_module()
