# 推送到TG

DT_STR=$(date "+%Y%m%d%H%M%S")
PART_X64=$(echo "${DT_STR}_Clash.Mini_X64_${GITHUB_SHA}" | base64 | tr -s "=" 2)
PART_X86=$(echo "${DT_STR}_Clash.Mini_X86_${GITHUB_SHA}" | base64 | tr -s "=" 2)

RUNNER_URL="https://github.com/JyCyunMe/Clash.Mini/actions/runs/${GITHUB_RUN_ID}"
echo "$RUNNER_URL"
ls -lahR "$ARTIFACTS_PATH"
ARTIFACT_X64="${ARTIFACTS_PATH}/${BUILD_X64_FILENAME}"
ARTIFACT_X86="${ARTIFACTS_PATH}/${BUILD_X86_FILENAME}"

ARTIFACT_X64_SHA256=$(cat "${ARTIFACT_X64}.sha256/${BUILD_X64_FILENAME}.sha256" | tr -d "\n")
ARTIFACT_X86_SHA256=$(cat "${ARTIFACT_X86}.sha256/${BUILD_X86_FILENAME}.sha256" | tr -d "\n")
RLT=$(curl --location --request POST https://api.telegram.org/bot${TG_BOT_TOKEN}/sendMediaGroup -s --form-string chat_id=${UPLOAD_CHAT_ID} --form "$PART_X64=@\"${ARTIFACT_X64}/${BUILD_X64_FILENAME}\"" --form "$PART_X86=@\"${ARTIFACT_X86}/${BUILD_X86_FILENAME}\"" --form-string media="[{\"type\": \"document\",\"media\": \"attach://$PART_X64\",\"caption\": \"SHA256: ${ARTIFACT_X64_SHA256}\",\"parse_mode\": \"MarkdownV2\"},{\"type\": \"document\",\"media\": \"attach://$PART_X86\",\"caption\": \"SHA256: ${ARTIFACT_X86_SHA256}\n\n_[${GITHUB_WORKFLOW} \\\\#${GITHUB_RUN_NUMBER}](${RUNNER_URL})_\",\"parse_mode\": \"MarkdownV2\"}]")
IS_OK=$(echo $RLT | jq ".ok")
echo $RLT | jq .
if ! $IS_OK; then
  echo "::error file=scripts/steps/notification/2_push_files_tg_dev.sh,line=19,col=1::pushing files to channel failed."
  echo "Response: "
  echo $RLT | jq .
  exit 1
fi
