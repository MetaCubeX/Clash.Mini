# 获取Git信息

$GIT_BRANCH = $env:GITHUB_REF -replace "refs/heads/", ""
$VERSION_REGEXP = "^v?\d+\.\d+\.\d+(\.\d+)?(-pre)?$"
$VERSION_RELEASE_REGEXP = "^v?\d+\.\d+\.\d+(\.\d+)?$"
$VERSION_PRE_REGEXP = "^v?\d+\.\d+\.\d+(\.\d+)?-pre$"
echo "VERSION_REGEXP=$VERSION_REGEXP" >> $env:GITHUB_ENV
echo "VERSION_RELEASE_REGEXP=$VERSION_RELEASE_REGEXP" >> $env:GITHUB_ENV
echo "VERSION_PRE_REGEXP=$VERSION_PRE_REGEXP" >> $env:GITHUB_ENV
$GIT_BRANCH = $env:GITHUB_REF -replace "refs/heads/", ""
echo "GIT_BRANCH=$GIT_BRANCH" >> $env:GITHUB_ENV
echo "Current Branch: $GIT_BRANCH"

$SHORT_SHA = ${env:GITHUB_SHA}.Substring(0, 7)
echo "SHORT_SHA=$SHORT_SHA" >> $env:GITHUB_ENV
echo "Current commit hash id: ${env:GITHUB_SHA} ($SHORT_SHA)"

$GIT_TAG_RELEASE_VERSION = git tag -l | where { $_ -match $VERSION_RELEASE_REGEXP } | sort -Descending -Top 1
$preTagList = $( git tag -l | where { $_ -match $VERSION_PRE_REGEXP } )
for($i = 0; $i -lt $preTagList.Length; $i++) {
    $preTagList[$i] = $preTagList[$i].Replace("-pre", "").Replace("v", "")
}
$GIT_TAG_PRE_VERSION = $preTagList | sort -Property { [version]$_ } -Descending -Top 1

$GIT_TAG_PRE_LATEST = "v${GIT_TAG_PRE_VERSION}-pre"
$GIT_TAG_RELEASE_LATEST = $GIT_TAG_RELEASE_VERSION
$GIT_TAG_RELEASE_VERSION = $GIT_TAG_RELEASE_VERSION.Replace("v", "")
if ($GIT_TAG_PRE_VERSION -eq $GIT_TAG_RELEASE_VERSION)
{
    $GIT_TAG_LATEST = $GIT_TAG_RELEASE_VERSION
}
else
{
    $GIT_TAG_LATEST = @("v${GIT_TAG_RELEASE_VERSION}", "v${GIT_TAG_PRE_VERSION}-pre") | sort -Descending -Top 1
}
echo "GIT_TAG_PRE_LATEST=$GIT_TAG_PRE_LATEST" >> $env:GITHUB_ENV
echo "GIT_TAG_RELEASE_LATEST=$GIT_TAG_RELEASE_LATEST" >> $env:GITHUB_ENV
echo "GIT_TAG_LATEST=$GIT_TAG_LATEST" >> $env:GITHUB_ENV
echo "Latest Tag: $GIT_TAG_LATEST"

echo "::set-output name=git-branch::$GIT_BRANCH"
