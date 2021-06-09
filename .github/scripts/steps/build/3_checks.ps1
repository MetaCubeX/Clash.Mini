# 构建前检查

$NOT_PASSED = 0
echo "Build Version: v${Env:BUILD_VERSION}`nCurrent Tag: ${Env:GIT_TAG}"
echo "Latest Tag: ${Env:GIT_TAG_LATEST}`nLatest Release Tag: ${Env:GIT_TAG_RELEASE_LATEST}`nLatest Pre Tag: ${Env:GIT_TAG_PRE_LATEST}"

if (!(${Env:GIT_TAG} -match ${Env:VERSION_REGEXP}) -or ${Env:GIT_TAG} -eq "")
{
    $NOT_PASSED = 1
    echo "::error file=scripts/steps/build/3_checks.ps1,line=10,col=1::Cannot get the version information or it's incorrect."
}
if ($NOT_PASSED -eq 0)
{
    $INTERNAL_VERSION_REGEXP = "^(\d+\.\d+\.\d+)(\.\d+)?$"
    $fileVersion = (cat .\versioninfo.json | jq -r ".FixedFileInfo.FileVersion")
    $fileVersion = (echo $fileVersion | jq -r ".Major, .Minor, .Patch, .Build") -join "."
    $tmpVer = [regex]::Match($fileVersion, $INTERNAL_VERSION_REGEXP)
    if (!$tmpVer.success)
    {
        $NOT_PASSED = 1
    }
    else
    {
        $fileVersion = $tmpVer.Groups[1].Value
        if ($tmpVer.Groups[2].Value -ne ".0")
        {
            $fileVersion += $tmpVer.Groups[2].Value
        }
    }

    if ($NOT_PASSED -eq 0)
    {
        $productVersion = (cat .\versioninfo.json | jq -r ".FixedFileInfo.ProductVersion")
        $productVersion = (echo $productVersion | jq -r ".Major, .Minor, .Patch, .Build") -join "."
        $tmpVer = [regex]::Match($productVersion, $INTERNAL_VERSION_REGEXP)
        if (!$tmpVer.success)
        {
            $NOT_PASSED = 1
        }
        else
        {
            $productVersion = $tmpVer.Groups[1].Value
            if ($tmpVer.Groups[2].Value -ne ".0")
            {
                $productVersion += $tmpVer.Groups[2].Value
            }
        }
    }
}

if (($NOT_PASSED -eq 0) -and (($productVersion -ne $fileVersion) -or (${Env:BUILD_VERSION} -ne $productVersion)))
{
    $NOT_PASSED = 1
    echo "::error file=scripts/steps/build/3_checks.ps1,line=54,col=1::The version information has some differences.`nPlease check `"versioninfo.json`""
}
(git tag -l | where { $_ -eq "v0.1.3.21-pre" }).Count

if (($NOT_PASSED -eq 0) -and (((git tag -l | where { $_ -eq ${Env:GIT_TAG} }).Count) -gt 0 -or
        (${Env:GIT_TAG_RELEASE_LATEST} -ne "" -and ${Env:GIT_TAG} -eq ${Env:GIT_TAG_RELEASE_LATEST}) -or
        (${Env:GIT_TAG_PRE_LATEST} -ne "" -and ${Env:GIT_TAG} -eq ${Env:GIT_TAG_PRE_LATEST}) -or
        (${Env:GIT_TAG_LATEST} -ne "" -and (${Env:GIT_TAG}.replace("-pre", "") -lt ${Env:GIT_TAG_LATEST}.replace("-pre", "") -or
                (${Env:GIT_TAG}.replace("-pre", "") -eq ${Env:GIT_TAG_LATEST}.replace("-pre", "") -and
                        (${Env:GIT_TAG}.contains("-pre") -or !${Env:GIT_TAG_LATEST}.contains("-pre"))))
        )))
{
    $NOT_PASSED = 1
    echo "::error file=scripts/steps/build/3_checks.ps1,line=67,col=1::A newer or the current version already exists."
}
if ($NOT_PASSED -ne 0)
{
    echo "::error file=scripts/steps/build/3_checks.ps1,line=71,col=1::Check the version information is not passed."
    exit 1
}
