# lmdb
本地影视数据库

ffmpeg -i 2.rmvb -y -f image2 -ss 070.010 -t 0.010 -s  -a  b.jpg

ffmpeg -i 2.rmvb -y -vframes 2 -vf a thumb.jpg

ffmpeg -i 2.rmvb sample.jpg -ss 00:00:05  -r 1 -vframes 1 -an -vcodec mjpeg
