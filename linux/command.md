### 查找进程
```linux
ps -aux | grep [进程名]
如：
ps -aux | grep goRenting
```
### 创建软连接
```linux
ln -s [目标文件地址] [软连接目录]
如：
ln -s /etc/sites/mailapi /etc/nginx/sites-enabled
```