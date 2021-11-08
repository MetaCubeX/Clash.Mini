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

sed -i "s/\{\{COMMIT_ID\}\}/${Env:COMMIT_ID}/g" "$pwd\app\app.go"
sed -i "s/\{\{BUGSNAG_KEY\}\}/${Env:BUGSNAG_KEY}/g" "$pwd\app\bugsnag.go"
sed -i "s/\{\{BRANCH\}\}/${Env:GIT_BRANCH}/g" "$pwd\app\bugsnag.go"
sed -i "s/\{\{MACHINE_ID_SECRET_VERSION\}\}/${Env:MACHINE_ID_SECRET_VERSION}/g" "$pwd\app\bugsnag.go"
sed -i "s/\{\{MACHINE_ID_SECRET\}\}/${Env:MACHINE_ID_SECRET}/g" "$pwd\app\bugsnag.go"

grep -r 'CommitId' "$pwd\app\app.go"
grep -r 'stage :=' "$pwd\app\bugsnag.go"
