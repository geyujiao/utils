
from Crypto.Cipher import AES
from Crypto.Random import get_random_bytes
import os

def encrypt_luac_file(input_file, output_file, key=None):
    """加密Lua字节码文件"""
    if key is None:
        key = get_random_bytes(16)  # AES-128
    
    # 读取原始字节码
    with open(input_file, 'rb') as f:
        plaintext = f.read()
    
    # 填充数据
    block_size = AES.block_size
    padding_length = block_size - len(plaintext) % block_size
    plaintext += bytes([padding_length]) * padding_length
    
    # 加密
    cipher = AES.new(key, AES.MODE_CBC)
    ciphertext = cipher.encrypt(plaintext)
    
    # 保存加密文件
    with open(output_file, 'wb') as f:
        f.write(cipher.iv)  # 保存初始化向量
        f.write(ciphertext)
    
    print(f"加密完成: {input_file} -> {output_file}")
    print(f"密钥(hex): {key.hex()}")
    return key

if __name__ == "__main__":
    # 示例用法
    input_luac = "example.luac"
    output_encrypted = "example_encrypted.luac"
    # input_luac = "example.lua"
    # output_encrypted = "example_encrypted.lua"
    
    # 使用固定密钥或随机生成
    fixed_key = b'abcdefghijklmnop'  # 16字节密钥
    encrypt_luac_file(input_luac, output_encrypted, fixed_key)
