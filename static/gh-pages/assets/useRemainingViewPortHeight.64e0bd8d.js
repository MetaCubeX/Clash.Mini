import{j as r}from"./vendor.287bb476.js";const{useState:s,useRef:u,useCallback:a,useLayoutEffect:c}=r;function g(){const t=u(null),[n,i]=s(200),e=a(()=>{const{top:o}=t.current.getBoundingClientRect();i(window.innerHeight-o)},[]);return c(()=>(e(),window.addEventListener("resize",e),()=>{window.removeEventListener("resize",e)}),[e]),[t,n]}export{g as u};
