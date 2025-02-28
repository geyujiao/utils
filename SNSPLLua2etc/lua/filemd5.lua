local md5 = require 'md5'
local file = io.open('yourfile.txt', 'r') -- 替换为你的文件路径
local content = file:read('*a') -- 读取文件内容
file:close()
 
local digest = md5.sumhexa(content) -- 计算MD5并以十六进制字符串形式返回
print(digest)