var Oe=Object.defineProperty,Te=Object.defineProperties;var Ce=Object.getOwnPropertyDescriptors;var G=Object.getOwnPropertySymbols;var we=Object.prototype.hasOwnProperty,xe=Object.prototype.propertyIsEnumerable;var J=(r,e,t)=>e in r?Oe(r,e,{enumerable:!0,configurable:!0,writable:!0,value:t}):r[e]=t,w=(r,e)=>{for(var t in e||(e={}))we.call(e,t)&&J(r,t,e[t]);if(G)for(var t of G(e))xe.call(e,t)&&J(r,t,e[t]);return r},U=(r,e)=>Te(r,Ce(e));import{ah as ae,ai as F,aj as ze,ak as oe,al as Ne,R as N,am as Pe,an as q,ao as Ee,ap as ke,aq as X,r as D,K as j,ar as Ae,A as le,X as Le,b as C,j as p,W as De,B as Fe,u as ue,d as We,g as Ue,C as $e}from"./index.76257afa.js";import{R as ce,T as Be}from"./TextFitler.c5efd418.js";import{f as qe}from"./index.cd34981b.js";import{F as je,p as He}from"./Fab.c287d48c.js";import{u as Ke}from"./useRemainingViewPortHeight.74fa6a72.js";import"./rotate-cw.5b6f0d7f.js";import"./debounce.d080d5e1.js";function Z(r,e){if(r==null)return{};var t={},i=Object.keys(r),n,s;for(s=0;s<i.length;s++)n=i[s],!(e.indexOf(n)>=0)&&(t[n]=r[n]);return t}var Qe=function(r){ae(e,r);function e(i,n){var s;return s=r.call(this)||this,s.client=i,s.setOptions(n),s.bindMethods(),s.updateResult(),s}var t=e.prototype;return t.bindMethods=function(){this.mutate=this.mutate.bind(this),this.reset=this.reset.bind(this)},t.setOptions=function(n){this.options=this.client.defaultMutationOptions(n)},t.onUnsubscribe=function(){if(!this.listeners.length){var n;(n=this.currentMutation)==null||n.removeObserver(this)}},t.onMutationUpdate=function(n){this.updateResult();var s={listeners:!0};n.type==="success"?s.onSuccess=!0:n.type==="error"&&(s.onError=!0),this.notify(s)},t.getCurrentResult=function(){return this.currentResult},t.reset=function(){this.currentMutation=void 0,this.updateResult(),this.notify({listeners:!0})},t.mutate=function(n,s){return this.mutateOptions=s,this.currentMutation&&this.currentMutation.removeObserver(this),this.currentMutation=this.client.getMutationCache().build(this.client,F({},this.options,{variables:typeof n!="undefined"?n:this.options.variables})),this.currentMutation.addObserver(this),this.currentMutation.execute()},t.updateResult=function(){var n=this.currentMutation?this.currentMutation.state:ze(),s=F({},n,{isLoading:n.status==="loading",isSuccess:n.status==="success",isError:n.status==="error",isIdle:n.status==="idle",mutate:this.mutate,reset:this.reset});this.currentResult=s},t.notify=function(n){var s=this;oe.batch(function(){s.mutateOptions&&(n.onSuccess?(s.mutateOptions.onSuccess==null||s.mutateOptions.onSuccess(s.currentResult.data,s.currentResult.variables,s.currentResult.context),s.mutateOptions.onSettled==null||s.mutateOptions.onSettled(s.currentResult.data,null,s.currentResult.variables,s.currentResult.context)):n.onError&&(s.mutateOptions.onError==null||s.mutateOptions.onError(s.currentResult.error,s.currentResult.variables,s.currentResult.context),s.mutateOptions.onSettled==null||s.mutateOptions.onSettled(void 0,s.currentResult.error,s.currentResult.variables,s.currentResult.context))),n.listeners&&s.listeners.forEach(function(o){o(s.currentResult)})})},e}(Ne);function de(r,e,t){var i=N.useRef(!1),n=N.useState(0),s=n[1],o=Pe(r,e,t),d=q(),c=N.useRef();c.current?c.current.setOptions(o):c.current=new Qe(d,o);var v=c.current.getCurrentResult();N.useEffect(function(){i.current=!0;var _=c.current.subscribe(oe.batchCalls(function(){i.current&&s(function(M){return M+1})}));return function(){i.current=!1,_()}},[]);var R=N.useCallback(function(_,M){c.current.mutate(_,M).catch(Ee)},[]);if(v.error&&ke(void 0,c.current.options.useErrorBoundary,[v.error]))throw v.error;return F({},v,{mutate:R,mutateAsync:v.mutate})}var Y=Number.isNaN||function(e){return typeof e=="number"&&e!==e};function Ve(r,e){return!!(r===e||Y(r)&&Y(e))}function Ge(r,e){if(r.length!==e.length)return!1;for(var t=0;t<r.length;t++)if(!Ve(r[t],e[t]))return!1;return!0}function $(r,e){e===void 0&&(e=Ge);var t,i=[],n,s=!1;function o(){for(var d=[],c=0;c<arguments.length;c++)d[c]=arguments[c];return s&&t===this&&e(d,i)||(n=r.apply(this,d),s=!0,t=this,i=d),n}return o}var Je=typeof performance=="object"&&typeof performance.now=="function",ee=Je?function(){return performance.now()}:function(){return Date.now()};function te(r){cancelAnimationFrame(r.id)}function Xe(r,e){var t=ee();function i(){ee()-t>=e?r.call(null):n.id=requestAnimationFrame(i)}var n={id:requestAnimationFrame(i)};return n}var x=null;function re(r){if(r===void 0&&(r=!1),x===null||r){var e=document.createElement("div"),t=e.style;t.width="50px",t.height="50px",t.overflow="scroll",t.direction="rtl";var i=document.createElement("div"),n=i.style;return n.width="100px",n.height="100px",e.appendChild(i),document.body.appendChild(e),e.scrollLeft>0?x="positive-descending":(e.scrollLeft=1,e.scrollLeft===0?x="negative":x="positive-ascending"),document.body.removeChild(e),x}return x}var Ze=150,Ye=function(e,t){return e};function et(r){var e,t=r.getItemOffset,i=r.getEstimatedTotalSize,n=r.getItemSize,s=r.getOffsetForIndexAndAlignment,o=r.getStartIndexForOffset,d=r.getStopIndexForStartIndex,c=r.initInstanceProps,v=r.shouldResetStyleCacheOnItemSizeChange,R=r.validateProps;return e=function(_){ae(M,_);function M(g){var a;return a=_.call(this,g)||this,a._instanceProps=c(a.props,X(a)),a._outerRef=void 0,a._resetIsScrollingTimeoutId=null,a.state={instance:X(a),isScrolling:!1,scrollDirection:"forward",scrollOffset:typeof a.props.initialScrollOffset=="number"?a.props.initialScrollOffset:0,scrollUpdateWasRequested:!1},a._callOnItemsRendered=void 0,a._callOnItemsRendered=$(function(l,u,f,m){return a.props.onItemsRendered({overscanStartIndex:l,overscanStopIndex:u,visibleStartIndex:f,visibleStopIndex:m})}),a._callOnScroll=void 0,a._callOnScroll=$(function(l,u,f){return a.props.onScroll({scrollDirection:l,scrollOffset:u,scrollUpdateWasRequested:f})}),a._getItemStyle=void 0,a._getItemStyle=function(l){var u=a.props,f=u.direction,m=u.itemSize,y=u.layout,h=a._getItemStyleCache(v&&m,v&&y,v&&f),I;if(h.hasOwnProperty(l))I=h[l];else{var S=t(a.props,l,a._instanceProps),O=n(a.props,l,a._instanceProps),T=f==="horizontal"||y==="horizontal",A=f==="rtl",L=T?S:0;h[l]=I={position:"absolute",left:A?void 0:L,right:A?L:void 0,top:T?0:S,height:T?"100%":O,width:T?O:"100%"}}return I},a._getItemStyleCache=void 0,a._getItemStyleCache=$(function(l,u,f){return{}}),a._onScrollHorizontal=function(l){var u=l.currentTarget,f=u.clientWidth,m=u.scrollLeft,y=u.scrollWidth;a.setState(function(h){if(h.scrollOffset===m)return null;var I=a.props.direction,S=m;if(I==="rtl")switch(re()){case"negative":S=-m;break;case"positive-descending":S=y-f-m;break}return S=Math.max(0,Math.min(S,y-f)),{isScrolling:!0,scrollDirection:h.scrollOffset<m?"forward":"backward",scrollOffset:S,scrollUpdateWasRequested:!1}},a._resetIsScrollingDebounced)},a._onScrollVertical=function(l){var u=l.currentTarget,f=u.clientHeight,m=u.scrollHeight,y=u.scrollTop;a.setState(function(h){if(h.scrollOffset===y)return null;var I=Math.max(0,Math.min(y,m-f));return{isScrolling:!0,scrollDirection:h.scrollOffset<I?"forward":"backward",scrollOffset:I,scrollUpdateWasRequested:!1}},a._resetIsScrollingDebounced)},a._outerRefSetter=function(l){var u=a.props.outerRef;a._outerRef=l,typeof u=="function"?u(l):u!=null&&typeof u=="object"&&u.hasOwnProperty("current")&&(u.current=l)},a._resetIsScrollingDebounced=function(){a._resetIsScrollingTimeoutId!==null&&te(a._resetIsScrollingTimeoutId),a._resetIsScrollingTimeoutId=Xe(a._resetIsScrolling,Ze)},a._resetIsScrolling=function(){a._resetIsScrollingTimeoutId=null,a.setState({isScrolling:!1},function(){a._getItemStyleCache(-1,null)})},a}M.getDerivedStateFromProps=function(a,l){return tt(a,l),R(a),null};var b=M.prototype;return b.scrollTo=function(a){a=Math.max(0,a),this.setState(function(l){return l.scrollOffset===a?null:{scrollDirection:l.scrollOffset<a?"forward":"backward",scrollOffset:a,scrollUpdateWasRequested:!0}},this._resetIsScrollingDebounced)},b.scrollToItem=function(a,l){l===void 0&&(l="auto");var u=this.props.itemCount,f=this.state.scrollOffset;a=Math.max(0,Math.min(a,u-1)),this.scrollTo(s(this.props,a,l,f,this._instanceProps))},b.componentDidMount=function(){var a=this.props,l=a.direction,u=a.initialScrollOffset,f=a.layout;if(typeof u=="number"&&this._outerRef!=null){var m=this._outerRef;l==="horizontal"||f==="horizontal"?m.scrollLeft=u:m.scrollTop=u}this._callPropsCallbacks()},b.componentDidUpdate=function(){var a=this.props,l=a.direction,u=a.layout,f=this.state,m=f.scrollOffset,y=f.scrollUpdateWasRequested;if(y&&this._outerRef!=null){var h=this._outerRef;if(l==="horizontal"||u==="horizontal")if(l==="rtl")switch(re()){case"negative":h.scrollLeft=-m;break;case"positive-ascending":h.scrollLeft=m;break;default:var I=h.clientWidth,S=h.scrollWidth;h.scrollLeft=S-I-m;break}else h.scrollLeft=m;else h.scrollTop=m}this._callPropsCallbacks()},b.componentWillUnmount=function(){this._resetIsScrollingTimeoutId!==null&&te(this._resetIsScrollingTimeoutId)},b.render=function(){var a=this.props,l=a.children,u=a.className,f=a.direction,m=a.height,y=a.innerRef,h=a.innerElementType,I=a.innerTagName,S=a.itemCount,O=a.itemData,T=a.itemKey,A=T===void 0?Ye:T,L=a.layout,ge=a.outerElementType,Ie=a.outerTagName,ye=a.style,Se=a.useIsScrolling,Re=a.width,H=this.state.isScrolling,W=f==="horizontal"||L==="horizontal",_e=W?this._onScrollHorizontal:this._onScrollVertical,K=this._getRangeToRender(),Me=K[0],be=K[1],Q=[];if(S>0)for(var E=Me;E<=be;E++)Q.push(D.exports.createElement(l,{data:O,key:A(E,O),index:E,isScrolling:Se?H:void 0,style:this._getItemStyle(E)}));var V=i(this.props,this._instanceProps);return D.exports.createElement(ge||Ie||"div",{className:u,onScroll:_e,ref:this._outerRefSetter,style:F({position:"relative",height:m,width:Re,overflow:"auto",WebkitOverflowScrolling:"touch",willChange:"transform",direction:f},ye)},D.exports.createElement(h||I||"div",{children:Q,ref:y,style:{height:W?"100%":V,pointerEvents:H?"none":void 0,width:W?V:"100%"}}))},b._callPropsCallbacks=function(){if(typeof this.props.onItemsRendered=="function"){var a=this.props.itemCount;if(a>0){var l=this._getRangeToRender(),u=l[0],f=l[1],m=l[2],y=l[3];this._callOnItemsRendered(u,f,m,y)}}if(typeof this.props.onScroll=="function"){var h=this.state,I=h.scrollDirection,S=h.scrollOffset,O=h.scrollUpdateWasRequested;this._callOnScroll(I,S,O)}},b._getRangeToRender=function(){var a=this.props,l=a.itemCount,u=a.overscanCount,f=this.state,m=f.isScrolling,y=f.scrollDirection,h=f.scrollOffset;if(l===0)return[0,0,0,0];var I=o(this.props,h,this._instanceProps),S=d(this.props,I,h,this._instanceProps),O=!m||y==="backward"?Math.max(1,u):1,T=!m||y==="forward"?Math.max(1,u):1;return[Math.max(0,I-O),Math.max(0,Math.min(l-1,S+T)),I,S]},M}(D.exports.PureComponent),e.defaultProps={direction:"ltr",itemData:void 0,layout:"vertical",overscanCount:2,useIsScrolling:!1},e}var tt=function(e,t){e.children,e.direction,e.height,e.layout,e.innerTagName,e.outerTagName,e.width,t.instance},rt=50,P=function(e,t,i){var n=e,s=n.itemSize,o=i.itemMetadataMap,d=i.lastMeasuredIndex;if(t>d){var c=0;if(d>=0){var v=o[d];c=v.offset+v.size}for(var R=d+1;R<=t;R++){var _=s(R);o[R]={offset:c,size:_},c+=_}i.lastMeasuredIndex=t}return o[t]},nt=function(e,t,i){var n=t.itemMetadataMap,s=t.lastMeasuredIndex,o=s>0?n[s].offset:0;return o>=i?fe(e,t,s,0,i):it(e,t,Math.max(0,s),i)},fe=function(e,t,i,n,s){for(;n<=i;){var o=n+Math.floor((i-n)/2),d=P(e,o,t).offset;if(d===s)return o;d<s?n=o+1:d>s&&(i=o-1)}return n>0?n-1:0},it=function(e,t,i,n){for(var s=e.itemCount,o=1;i<s&&P(e,i,t).offset<n;)i+=o,o*=2;return fe(e,t,Math.min(i,s-1),Math.floor(i/2),n)},ne=function(e,t){var i=e.itemCount,n=t.itemMetadataMap,s=t.estimatedItemSize,o=t.lastMeasuredIndex,d=0;if(o>=i&&(o=i-1),o>=0){var c=n[o];d=c.offset+c.size}var v=i-o-1,R=v*s;return d+R},st=et({getItemOffset:function(e,t,i){return P(e,t,i).offset},getItemSize:function(e,t,i){return i.itemMetadataMap[t].size},getEstimatedTotalSize:ne,getOffsetForIndexAndAlignment:function(e,t,i,n,s){var o=e.direction,d=e.height,c=e.layout,v=e.width,R=o==="horizontal"||c==="horizontal",_=R?v:d,M=P(e,t,s),b=ne(e,s),g=Math.max(0,Math.min(b-_,M.offset)),a=Math.max(0,M.offset-_+M.size);switch(i==="smart"&&(n>=a-_&&n<=g+_?i="auto":i="center"),i){case"start":return g;case"end":return a;case"center":return Math.round(a+(g-a)/2);case"auto":default:return n>=a&&n<=g?n:n<a?a:g}},getStartIndexForOffset:function(e,t,i){return nt(e,i,t)},getStopIndexForStartIndex:function(e,t,i,n){for(var s=e.direction,o=e.height,d=e.itemCount,c=e.layout,v=e.width,R=s==="horizontal"||c==="horizontal",_=R?v:o,M=P(e,t,n),b=i+_,g=M.offset+M.size,a=t;a<d-1&&g<b;)a++,g+=P(e,a,n).size;return a},initInstanceProps:function(e,t){var i=e,n=i.estimatedItemSize,s={itemMetadataMap:{},estimatedItemSize:n||rt,lastMeasuredIndex:-1};return t.resetAfterIndex=function(o,d){d===void 0&&(d=!0),s.lastMeasuredIndex=Math.min(s.lastMeasuredIndex,o-1),t._getItemStyleCache(-1),d&&t.forceUpdate()},s},shouldResetStyleCacheOnItemSizeChange:!1,validateProps:function(e){e.itemSize}});function ie(r,e){for(var t in r)if(!(t in e))return!0;for(var i in e)if(r[i]!==e[i])return!0;return!1}var at=["style"],ot=["style"];function lt(r,e){var t=r.style,i=Z(r,at),n=e.style,s=Z(e,ot);return!ie(t,n)&&!ie(i,s)}function ut(r){const e=r.providers,t=Object.keys(e),i={};for(let n=0;n<t.length;n++){const s=t[n];i[s]=U(w({},e[s]),{idx:n})}return{byName:i,names:t}}async function ct(r,e){const{url:t,init:i}=j(e);let n={providers:{}};try{const s=await fetch(t+r,i);s.ok&&(n=await s.json())}catch(s){console.log("failed to GET /providers/rules",s)}return ut(n)}async function me({name:r,apiConfig:e}){const{url:t,init:i}=j(e);try{return(await fetch(t+`/providers/rules/${r}`,w({method:"PUT"},i))).ok}catch(n){return console.log("failed to PUT /providers/rules/:name",n),!1}}async function dt({names:r,apiConfig:e}){for(let t=0;t<r.length;t++)await me({name:r[t],apiConfig:e})}var ft=function(r,e,t,i,n,s,o,d){if(!r){var c;if(e===void 0)c=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var v=[t,i,n,s,o,d],R=0;c=new Error(e.replace(/%s/g,function(){return v[R++]})),c.name="Invariant Violation"}throw c.framesToPop=1,c}},mt=ft;function ht(r){return mt(r.rules&&r.rules.length>=0,"there is no valid rules list in the rules API response"),r.rules.map((e,t)=>U(w({},e),{id:t}))}async function vt(r,e){let t={rules:[]};try{const{url:i,init:n}=j(e),s=await fetch(i+r,n);s.ok&&(t=await s.json())}catch(i){console.log("failed to fetch rules",i)}return ht(t)}const he=Ae({key:"ruleFilterText",default:""});function pt(r,e){const t=q(),{mutate:i,isLoading:n}=de(me,{onSuccess:()=>{t.invalidateQueries("/providers/rules")}});return[o=>{o.preventDefault(),i({name:r,apiConfig:e})},n]}function gt(r){const e=q(),{data:t}=ve(r),{mutate:i,isLoading:n}=de(dt,{onSuccess:()=>{e.invalidateQueries("/providers/rules")}});return[o=>{o.preventDefault(),i({names:t.names,apiConfig:r})},n]}function ve(r){return le(["/providers/rules",r],()=>ct("/providers/rules",r))}function It(r){const{data:e,isFetching:t}=le(["/rules",r],()=>vt("/rules",r)),{data:i}=ve(r),[n]=Le(he);if(n==="")return{rules:e,provider:i,isFetching:t};{const s=n.toLowerCase();return{rules:e.filter(o=>o.payload.toLowerCase().indexOf(s)>=0),isFetching:t,provider:{byName:i.byName,names:i.names.filter(o=>o.toLowerCase().indexOf(s)>=0)}}}}const yt="_RuleProviderItem_ly9yn_1",St="_left_ly9yn_7",Rt="_middle_ly9yn_14",_t="_gray_ly9yn_20",Mt="_refreshButtonWrapper_ly9yn_24";var z={RuleProviderItem:yt,left:St,middle:Rt,gray:_t,refreshButtonWrapper:Mt};function bt({idx:r,name:e,vehicleType:t,behavior:i,updatedAt:n,ruleCount:s,apiConfig:o}){const[d,c]=pt(e,o),v=qe(new Date(n),new Date);return C("div",{className:z.RuleProviderItem,children:[p("span",{className:z.left,children:r}),C("div",{className:z.middle,children:[p(De,{name:e,type:`${t} / ${i}`}),p("div",{className:z.gray,children:s<2?`${s} rule`:`${s} rules`}),C("small",{className:z.gray,children:["Updated ",v," ago"]})]}),p("span",{className:z.refreshButtonWrapper,children:p(Fe,{onClick:d,disabled:c,children:p(ce,{isRotating:c})})})]})}function Ot({apiConfig:r}){const[e,t]=gt(r),{t:i}=ue();return p(je,{icon:p(ce,{isRotating:t}),text:i("update_all_rule_provider"),style:He,onClick:e})}const Tt="_rule_14p9p_1",Ct="_left_14p9p_12",wt="_a_14p9p_19",xt="_b_14p9p_26",zt="_type_14p9p_37";var k={rule:Tt,left:Ct,a:wt,b:xt,type:zt};const B={_default:"#59caf9",DIRECT:"#f5bc41",REJECT:"#cb3166"};function Nt({proxy:r}){let e=B._default;return B[r]&&(e=B[r]),{color:e}}function Pt({type:r,payload:e,proxy:t,id:i}){const n=Nt({proxy:t});return C("div",{className:k.rule,children:[p("div",{className:k.left,children:i}),C("div",{children:[p("div",{className:k.b,children:e}),C("div",{className:k.a,children:[p("div",{className:k.type,children:r}),p("div",{style:n,children:t})]})]})]})}const Et="_header_1j1w3_1",kt="_RuleProviderItemWrapper_1j1w3_17";var pe={header:Et,RuleProviderItemWrapper:kt};const{memo:At}=N,se=30;function Lt(r,{rules:e,provider:t}){const i=t.names.length;return r<i?t.names[r]:e[r-i].id}function Dt({provider:r}){return function(t){const i=r.names.length;return t<i?90:60}}const Ft=At(({index:r,style:e,data:t})=>{const{rules:i,provider:n,apiConfig:s}=t,o=n.names.length;if(r<o){const c=n.names[r],v=n.byName[c];return p("div",{style:e,className:pe.RuleProviderItemWrapper,children:p(bt,w({apiConfig:s},v))})}const d=i[r-o];return p("div",{style:e,children:p(Pt,w({},d))})},lt),Wt=r=>({apiConfig:Ue(r)});var Gt=We(Wt)(Ut);function Ut({apiConfig:r}){const[e,t]=Ke(),{rules:i,provider:n}=It(r),s=Dt({provider:n}),{t:o}=ue();return C("div",{children:[C("div",{className:pe.header,children:[p($e,{title:o("Rules")}),p(Be,{placeholder:"Filter",textAtom:he})]}),p("div",{ref:e,style:{paddingBottom:se},children:p(st,{height:t-se,width:"100%",itemCount:i.length+n.names.length,itemSize:s,itemData:{rules:i,provider:n,apiConfig:r},itemKey:Lt,children:Ft})}),n&&n.names&&n.names.length>0?p(Ot,{apiConfig:r}):null]})}export{Gt as default};
