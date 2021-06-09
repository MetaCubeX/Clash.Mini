# 推送到TG

if [ -f "./RELEASELOG_NOTIFY.md" ]; then
  RELEASE_LOG="./RELEASELOG_NOTIFY.md"
else
  RELEASE_LOG="./RELEASELOG.md"
fi
RLT=$(curl --location --request POST https://api.telegram.org/bot${TG_BOT_TOKEN}/sendMessage -s --form-string chat_id=${CHAT_ID} --form-string text="$(perl -lne 'print;' $RELEASE_LOG)" --form-string parse_mode="Markdown" --form-string disable_web_page_preview="true" --form-string allow_sending_without_reply="true" --form-string reply_markup="{\"inline_keyboard\":[[{\"text\":\"📦下载\",\"url\":\"https://github.com/Clash-Mini/Clash.Mini/releases/tag/${GIT_TAG}\"},{\"text\":\"⭐点个Star吧\",\"url\":\"https://github.com/Clash-Mini/Clash.Mini\"}]]}")
IS_OK=$(echo $RLT | jq ".ok")
echo $RLT | jq .
if $IS_OK; then
  MSG_ID=$(echo $RLT | jq ".result.message_id")
  echo "::set-output name=push-msg-id::$MSG_ID"
else
  echo "::error file=scripts/steps/notification/1_push_broadcast_tg.sh,line=15,col=1::push to channel failed."
  echo "Response: "
  echo "$RLT" | jq .
  exit 1
fi
