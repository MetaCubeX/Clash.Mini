# 准备发布Release文件

cd $Env:BUILD_PATH
cp .\Clash.Mini_*.exe ${Env:PUBLISH_PATH}\
echo "Ready to upload the following file(s):"
ls ${Env:PUBLISH_PATH}