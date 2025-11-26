
#!/bin/bash

# 获取主机名并转为小写
host_name=$(hostname | tr '[:upper:]' '[:lower:]')

# 设置打包输出目录（支持不存在时自动创建）
output_dir="/home/sudoroot/temp"

# 自动创建输出目录（包括父目录）
if [ ! -d "$output_dir" ]; then
  echo "📁 目录 $output_dir 不存在，正在创建..."
  mkdir -p "$output_dir" 2>/dev/null
fi

# 再次检查是否创建成功
if [ ! -d "$output_dir" ]; then
  echo "❌ 错误：无法创建输出目录 $output_dir，权限不足或路径无效。"
  exit 1
fi

# 获取时间戳
timestamp=$(date +"%Y%m%d_%H%M%S")

# 构建备份文件路径
backup_file="$output_dir/config_backup-${host_name}-${timestamp}.tar.gz"

# 定义需要打包的目录和文件列表
files_to_backup=(
  "/etc/nginx"
  "/etc/trafficserver"
  "/etc/nginx/nginx.conf"
  "/etc/dnsmasq.conf"
  "/etc/dnsmasq.d/address.conf"
)

# 过滤掉不存在的路径
valid_files=()
for path in "${files_to_backup[@]}"; do
  if [ -e "$path" ]; then
    valid_files+=("$path")
  else
    echo "⚠️ 警告：文件或目录 $path 不存在，已跳过。"
  fi
done

# 判断是否有可打包的内容
if [ ${#valid_files[@]} -eq 0 ]; then
  echo "❌ 错误：没有找到任何有效的文件或目录可供打包。"
  exit 1
fi

# 开始打包压缩
echo "📦 正在打包配置文件到 $backup_file ..."
tar -czf "$backup_file" "${valid_files[@]}" --absolute-names

# 检查打包结果
if [ $? -eq 0 ]; then
  echo "✅ 成功打包完成！文件位于: $backup_file"
else
  echo "❌ 打包失败！请检查路径、权限或磁盘空间。"
  exit 1
fi
