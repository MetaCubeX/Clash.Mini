# 配置构建信息

$BUILD_PATH = "$pwd\build"
$PUBLISH_PATH = "$BUILD_PATH\publish"
echo "BUILD_PATH=$BUILD_PATH" >> $Env:GITHUB_ENV
echo "PUBLISH_PATH=$PUBLISH_PATH" >> $Env:GITHUB_ENV

$BUILD_VERSION = (cat .\versioninfo.json | jq -r ".StringFileInfo.ProductVersion")
echo "BUILD_VERSION=$BUILD_VERSION" >> $Env:GITHUB_ENV
echo "Build Version: v$BUILD_VERSION"
$GIT_TAG = "v$BUILD_VERSION$( ${env:GIT_BRANCH} -ne 'release' ? '-pre' : '' )"
echo "GIT_TAG=$GIT_TAG" >> $Env:GITHUB_ENV
echo "Current Tag: $GIT_TAG"
echo "::set-output name=git-tag::$GIT_TAG"
