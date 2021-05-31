package util

//import (
//	"os"
//	"strings"
//)
//
//const (
//	EmptyString		=	""
//)
//
//func joinPathFull(fileName string, suffix string, paths... string) string {
//	sb := strings.Builder{}
//	var i int
//	var needCut bool
//	for j, p := range paths {
//		needCut = false
//		for i = len(p) - 1; !needCut  && i > 0 && i > len(p) - 3; i-- {
//			if os.IsPathSeparator(p[i]) {
//				needCut = true
//			}
//		}
//		if needCut {
//			sb.WriteString(p[:i])
//		}
//		if j < len(paths) || fileName != EmptyString || suffix != EmptyString {
//			sb.WriteRune('/')
//		}
//	}
//	if fileName != EmptyString {
//		sb.WriteString(fileName)
//	}
//	if suffix != EmptyString {
//		sb.WriteString(suffix)
//	}
//	return sb.String()
//}
//
//// JoinPath 拼接
//func JoinPath(paths... string) string {
//	return joinPathFull(EmptyString, EmptyString, paths...)
//}
//
//func JoinPathWithSuffix(suffix string, paths... string) string {
//	//JoinPath([]string{ "13123", "13123" }...)
//	return joinPathFull(EmptyString, suffix, paths...)
//}
//
//func JoinPathWithFile(fileName string, suffix string, paths... string) string {
//	return joinPathFull(fileName, suffix, paths...)
//}
