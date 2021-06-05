# 获取Git信息

$GIT_BRANCH=$env:GITHUB_REF -replace "refs/heads/", ""
$VERSION_REGEXP="^v?\d+\.\d+\.\d+(\.\d+)?(-pre)?$"
$VERSION_RELEASE_REGEXP="^v?\d+\.\d+\.\d+(\.\d+)?$"
$VERSION_PRE_REGEXP="^v?\d+\.\d+\.\d+(\.\d+)?-pre$"
echo "VERSION_REGEXP=$VERSION_REGEXP" >> $env:GITHUB_ENV
echo "VERSION_RELEASE_REGEXP=$VERSION_RELEASE_REGEXP" >> $env:GITHUB_ENV
echo "VERSION_PRE_REGEXP=$VERSION_PRE_REGEXP" >> $env:GITHUB_ENV
$GIT_BRANCH=$env:GITHUB_REF -replace "refs/heads/", ""
echo "GIT_BRANCH=$GIT_BRANCH" >> $env:GITHUB_ENV
echo "Current Branch: $GIT_BRANCH"

$SHORT_SHA=${env:GITHUB_SHA}.Substring(0,7)
echo "SHORT_SHA=$SHORT_SHA" >> $env:GITHUB_ENV
echo "Current commit hash id: ${env:GITHUB_SHA} ($SHORT_SHA)"

$GIT_TAG_RELEASE_VERSION=git tag -l | where { $_ -match $VERSION_RELEASE_REGEXP } | sort -descending -top 1
$GIT_TAG_PRE_VERSION=git tag -l | where { $_ -match $VERSION_PRE_REGEXP } | sort -descending -top 1
if ($GIT_TAG_PRE_VERSION -eq $GIT_TAG_RELEASE_VERSION + "-pre") { $GIT_TAG_LATEST=$GIT_TAG_RELEASE_VERSION }
else { $GIT_TAG_LATEST=@($GIT_TAG_RELEASE_VERSION, $GIT_TAG_PRE_VERSION) | sort -descending -top 1 }
echo "GIT_TAG_LATEST=$GIT_TAG_LATEST" >> $env:GITHUB_ENV
echo "Latest Tag: $GIT_TAG_LATEST"
