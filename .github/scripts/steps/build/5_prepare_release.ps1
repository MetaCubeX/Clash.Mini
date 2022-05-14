# 准备发布Release文件

cd $env:BUILD_PATH
mkdir -Force $env:PUBLISH_PATH
$BUILD_X64_FILENAME = "Clash.Mini_${env:GIT_BRANCH}_v${env:BUILD_VERSION}_x64.exe"
$BUILD_X86_FILENAME = "Clash.Mini_${env:GIT_BRANCH}_v${env:BUILD_VERSION}_x86.exe"
echo "BUILD_X64_FILENAME=$BUILD_X64_FILENAME" >> $env:GITHUB_ENV
echo "BUILD_X86_FILENAME=$BUILD_X86_FILENAME" >> $env:GITHUB_ENV
cp .\Clash.Mini_*.exe ${Env:PUBLISH_PATH}\
echo "Ready to upload the following file(s):"
ls ${Env:PUBLISH_PATH}