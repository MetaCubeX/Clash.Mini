import{g as i,r as a,O as l,N as u,j as p}from"./vendor.2c7996a3.js";import{d as x}from"./debounce.76599460.js";const _="_rotate_1dspl_1",g="_isRotating_1dspl_5",d="_rotating_1dspl_1";var r={rotate:_,isRotating:g,rotating:d};function E({isRotating:t}){const e=i(r.rotate,{[r.isRotating]:t});return a.exports.createElement("span",{className:e},a.exports.createElement(l,{width:16}))}const{useCallback:m,useState:R,useMemo:h}=p;function v(t){const[,e]=u(t),[n,c]=R(""),o=h(()=>x(e,300),[e]);return[m(s=>{c(s.target.value),o(s.target.value)},[o]),n]}const T="_input_7q916_1";var b={input:T};function N(t){const[e,n]=v(t.textAtom);return a.exports.createElement("input",{className:b.input,type:"text",value:n,onChange:e,placeholder:t.placeholder})}export{E as R,N as T};
