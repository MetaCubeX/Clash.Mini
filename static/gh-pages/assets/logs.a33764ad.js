var C=Object.defineProperty,E=Object.defineProperties;var b=Object.getOwnPropertyDescriptors;var h=Object.getOwnPropertySymbols;var j=Object.prototype.hasOwnProperty,A=Object.prototype.propertyIsEnumerable;var v=(e,t,n)=>t in e?C(e,t,{enumerable:!0,configurable:!0,writable:!0,value:n}):e[t]=n,p=(e,t)=>{for(var n in t||(t={}))j.call(t,n)&&v(e,n,t[n]);if(h)for(var n of h(t))A.call(t,n)&&v(e,n,t[n]);return e},w=(e,t)=>E(e,b(t));import{r as F,t as H,v as f}from"./index.decfcd35.js";var D;(function(e){e[e.Connecting=0]="Connecting",e[e.Open=1]="Open",e[e.Closing=2]="Closing",e[e.Closed=3]="Closed"})(D||(D={}));const L="/logs",J=new TextDecoder("utf-8"),N=()=>Math.floor((1+Math.random())*65536).toString(16);let M=!1,i=!1,a="",s,d;function m(e,t){let n;try{n=JSON.parse(e)}catch{console.log("JSON.parse error",JSON.parse(e))}const r=new Date,l=T(r);n.time=l,n.id=+r-0+N(),n.even=M=!M,t(n)}function T(e){const t=e.getFullYear()%100,n=f(e.getMonth()+1,2),r=f(e.getDate(),2),l=f(e.getHours(),2),o=f(e.getMinutes(),2),c=f(e.getSeconds(),2);return`${t}-${n}-${r} ${l}:${o}:${c}`}function O(e,t){return e.read().then(({done:n,value:r})=>{a+=J.decode(r,{stream:!n});const o=a.split(`
`),c=o[o.length-1];for(let g=0;g<o.length-1;g++)m(o[g],t);if(n){m(c,t),a="",console.log("GET /logs streaming done"),i=!1;return}else a=c;return O(e,t)})}function $(e){const t=Object.keys(e);return t.sort(),t.map(n=>e[n]).join("|")}let x,u;function Y(e,t){if(e.logLevel==="uninit"||i||s&&s.readyState===1)return;d=t;const n=F(e,L);s=new WebSocket(n),s.addEventListener("error",()=>{G(e,t)}),s.addEventListener("message",function(r){m(r.data,t)})}function q(){s.close(),u&&u.abort()}function z(e){!d||!s||(s.close(),i=!1,Y(e,d))}function G(e,t){if(u&&$(e)!==x)u.abort();else if(i)return;i=!0,x=$(e),u=new AbortController;const n=u.signal,{url:r,init:l}=H(e);fetch(r+L+"?level="+e.logLevel,w(p({},l),{signal:n})).then(o=>{const c=o.body.getReader();O(c,t)},o=>{i=!1,!n.aborted&&console.log("GET /logs error:",o.message)})}export{Y as f,z as r,q as s};
