# 准备上传Artifact文件

cd $env:BUILD_PATH
mkdir -Force $env:PUBLISH_PATH
$BUILD_X64_FILENAME = "Clash.Mini_${env:GIT_BRANCH}_v${env:BUILD_VERSION}_${env:SHORT_SHA}_x64.exe"
$BUILD_X86_FILENAME = "Clash.Mini_${env:GIT_BRANCH}_v${env:BUILD_VERSION}_${env:SHORT_SHA}_x86.exe"
echo "BUILD_X64_FILENAME=$BUILD_X64_FILENAME" >> $env:GITHUB_ENV
echo "BUILD_X86_FILENAME=$BUILD_X86_FILENAME" >> $env:GITHUB_ENV
cp .\Clash.Mini_dev_x64.exe $env:PUBLISH_PATH\$BUILD_X64_FILENAME
cp .\Clash.Mini_dev_x86.exe $env:PUBLISH_PATH\$BUILD_X86_FILENAME

echo "Ready to upload the following file(s):"
ls $env:PUBLISH_PATH

echo "::set-output name=build-x64-filename::$BUILD_X64_FILENAME"
echo "::set-output name=build-x86-filename::$BUILD_X86_FILENAME"
