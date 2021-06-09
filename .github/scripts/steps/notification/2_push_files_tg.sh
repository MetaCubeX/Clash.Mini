# 推送到TG

DT_STR=$(date "+%Y%m%d%H%M%S")
PART_X64=$(echo "${DT_STR}_Clash.Mini_X64_${GITHUB_SHA}" | base64 | tr -s "=" 2)
PART_X86=$(echo "${DT_STR}_Clash.Mini_X86_${GITHUB_SHA}" | base64 | tr -s "=" 2)

RELEASE_URL="https://github.com/Clash-Mini/Clash.Mini/releases/download/${GIT_TAG}"
echo "$RELEASE_URL"
RELEASE_PATH="$(pwd)/releases"
mkdir -p "$RELEASE_PATH"
echo "${RELEASE_URL}/${RELEASE_PKG_X64}"
echo "${RELEASE_URL}/${RELEASE_PKG_X86}"
echo "${RELEASE_URL}/${RELEASE_PKG_X64}.sha256"
echo "${RELEASE_URL}/${RELEASE_PKG_X86}.sha256"
curl -L -o "${RELEASE_PATH}/${RELEASE_PKG_X64}" "${RELEASE_URL}/${RELEASE_PKG_X64}"
curl -L -o "${RELEASE_PATH}/${RELEASE_PKG_X86}" "${RELEASE_URL}/${RELEASE_PKG_X86}"
curl -L -o "${RELEASE_PATH}/${RELEASE_PKG_X64}.sha256" "${RELEASE_URL}/${RELEASE_PKG_X64}.sha256"
curl -L -o "${RELEASE_PATH}/${RELEASE_PKG_X86}.sha256" "${RELEASE_URL}/${RELEASE_PKG_X86}.sha256"
ls -lahR "$RELEASE_PATH"

RELEASE_X64_SHA256=$(cat "${RELEASE_PATH}/${RELEASE_PKG_X64}.sha256" | tr -d "\n")
RELEASE_X86_SHA256=$(cat "${RELEASE_PATH}/${RELEASE_PKG_X86}.sha256" | tr -d "\n")
RLT=$(curl --location --request POST https://api.telegram.org/bot${TG_BOT_TOKEN}/sendMediaGroup -s --form-string chat_id=${CHAT_ID} --form-string reply_to_message_id=$PUSH_MSG_ID --form $PART_X64=@"${RELEASE_PATH}/${RELEASE_PKG_X64}" --form $PART_X86=@"${RELEASE_PATH}/${RELEASE_PKG_X86}" --form-string media="[{\"type\": \"document\",\"media\": \"attach://$PART_X64\",\"caption\": \"SHA256: ${RELEASE_X64_SHA256}\",\"parse_mode\": \"Markdown\"},{\"type\": \"document\",\"media\": \"attach://$PART_X86\",\"caption\": \"SHA256: ${RELEASE_X86_SHA256}\",\"parse_mode\": \"Markdown\"}]")
IS_OK=$(echo "$RLT" | jq ".ok")
if ! $IS_OK; then
  echo "::error file=scripts/steps/notification/2_push_files_tg.sh,line=26,col=1::pushing files to channel failed."
  echo "Response: "
  echo "$RLT" | jq .
  exit 1
fi
