(self.webpackChunkyacd = self.webpackChunkyacd || []).push([
    [143], {
        50497: function (e, n, t) {
            "use strict";
            t.d(n, {
                T: function () {
                    return f
                }, w: function () {
                    return h
                }
            });
            var r = t(94949),
                o = t(80043),
                i = t(87757),
                c = t.n(i),
                s = (t(88674), t(41539), t(82526), t(57327), t(54747), t(49337), t(97943));

            function a(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function u(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? a(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : a(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var l = "/configs";

            function f(e) {
                return p.apply(this, arguments)
            }

            function p() {
                return (p = (0, o.Z)(c().mark((function e(n) {
                    var t, r, o;
                    return c().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return t = (0, s.g)(n), r = t.url, o = t.init, e.next = 3, fetch(r + l, o);
                            case 3:
                                return e.abrupt("return", e.sent);
                            case 4:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function d(e) {
                return "socks-port" in e && (e["socket-port"] = e["socks-port"]), e
            }

            function h(e, n) {
                return v.apply(this, arguments)
            }

            function v() {
                return (v = (0, o.Z)(c().mark((function e(n, t) {
                    var r, o, i, a;
                    return c().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return r = (0, s.g)(n), o = r.url, i = r.init, a = JSON.stringify(d(t)), e.next = 4, fetch(o + l, u(u({}, i), {}, {
                                    body: a,
                                    method: "PATCH"
                                }));
                            case 4:
                                return e.abrupt("return", e.sent);
                            case 5:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }
        }, 97750: function (e, n, t) {
            "use strict";
            t.d(n, {
                rQ: function () {
                    return d
                }, PI: function () {
                    return v
                }, $K: function () {
                    return g
                }, Sm: function () {
                    return y
                }
            });
            var r = t(94949),
                o = t(80043),
                i = t(87757),
                c = t.n(i),
                s = (t(82772), t(40561), t(88674), t(41539), t(82526), t(57327), t(54747), t(49337), t(97943));

            function a(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function u(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? a(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : a(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var l, f = "/connections",
                p = [];

            function d(e, n) {
                if (1 === l && n) return h(n);
                l = 1;
                var t = (0, s.P)(e, f),
                    r = new WebSocket(t);
                return r.addEventListener("error", (function () {
                    return l = 3
                })), r.addEventListener("message", (function (e) {
                    return function (e) {
                        var n;
                        try {
                            n = JSON.parse(e)
                        } catch (n) {
                            console.log("JSON.parse error", JSON.parse(e))
                        }
                        p.forEach((function (e) {
                            return e(n)
                        }))
                    }(e.data)
                })), n ? h(n) : void 0
            }

            function h(e) {
                return p.push(e),
                    function () {
                        var n = p.indexOf(e);
                        p.splice(n, 1)
                    }
            }

            function v(e) {
                return b.apply(this, arguments)
            }

            function b() {
                return (b = (0, o.Z)(c().mark((function e(n) {
                    var t, r, o;
                    return c().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return t = (0, s.g)(n), r = t.url, o = t.init, e.next = 3, fetch(r + f, u(u({}, o), {}, {
                                    method: "DELETE"
                                }));
                            case 3:
                                return e.abrupt("return", e.sent);
                            case 4:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function g(e) {
                return x.apply(this, arguments)
            }

            function x() {
                return (x = (0, o.Z)(c().mark((function e(n) {
                    var t, r, o;
                    return c().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return t = (0, s.g)(n), r = t.url, o = t.init, e.next = 3, fetch(r + f, u({}, o));
                            case 3:
                                return e.abrupt("return", e.sent);
                            case 4:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function y(e, n) {
                return j.apply(this, arguments)
            }

            function j() {
                return (j = (0, o.Z)(c().mark((function e(n, t) {
                    var r, o, i, a;
                    return c().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return r = (0, s.g)(n), o = r.url, i = r.init, a = `${o}${f}/${t}`, e.next = 4, fetch(a, u(u({}, i), {}, {
                                    method: "DELETE"
                                }));
                            case 4:
                                return e.abrupt("return", e.sent);
                            case 5:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }
        }, 41289: function (e, n, t) {
            "use strict";
            t.d(n, {
                r: function () {
                    return d
                }
            });
            t(54747), t(82772), t(40561), t(23123), t(88674), t(41539);
            var r, o = t(97943),
                i = "/traffic",
                c = new TextDecoder("utf-8"),
                s = 150,
                a = {
                    labels: Array(s),
                    up: Array(s),
                    down: Array(s),
                    size: s,
                    subscribers: [],
                    appendData(e) {
                        this.up.push(e.up), this.down.push(e.down);
                        var n = new Date,
                            t = "" + n.getMinutes() + n.getSeconds();
                        this.labels.push(t), this.up.length > this.size && this.up.shift(), this.down.length > this.size && this.down.shift(), this.labels.length > this.size && this.labels.shift(), this.subscribers.forEach((function (n) {
                            return n(e)
                        }))
                    },
                    subscribe(e) {
                        var n = this;
                        return this.subscribers.push(e),
                            function () {
                                var t = n.subscribers.indexOf(e);
                                n.subscribers.splice(t, 1)
                            }
                    }
                },
                u = !1,
                l = "";

            function f(e) {
                a.appendData(JSON.parse(e))
            }

            function p(e) {
                return e.read().then((function (n) {
                    for (var t = n.done, r = n.value, o = c.decode(r, {
                        stream: !t
                    }), i = (l += o).split("\n"), s = i[i.length - 1], a = 0; a < i.length - 1; a++) f(i[a]);
                    return t ? (f(s), l = "", console.log("GET /traffic streaming done"), void(u = !1)) : (l = s, p(e))
                }))
            }

            function d(e) {
                if (u || 1 === r) return a;
                r = 1;
                var n = (0, o.P)(e, i),
                    t = new WebSocket(n);
                return t.addEventListener("error", (function (e) {
                    r = 3
                })), t.addEventListener("close", (function (n) {
                    r = 3,
                        function (e) {
                            if (u) return a;
                            u = !0;
                            var n = (0, o.g)(e),
                                t = n.url,
                                r = n.init;
                            fetch(t + i, r).then((function (e) {
                                e.ok ? p(e.body.getReader()) : u = !1
                            }), (function (e) {
                                console.log("fetch /traffic error", e), u = !1
                            }))
                        }(e)
                })), t.addEventListener("message", (function (e) {
                    f(e.data)
                })), a
            }
        }, 58392: function (e, n, t) {
            "use strict";
            t(88674), t(41539), t(66992), t(33948);
            var r = t(14613),
                o = t(37110),
                i = t(83554),
                c = t(68718),
                s = {
                    zh: t.e(140).then(t.bind(t, 19965)),
                    en: t.e(36).then(t.bind(t, 66036))
                };
            r.Z.use(i.Z).use(c.Db).use(o.Z).init({
                debug: !1,
                backend: {
                    loadPath: "/__{{lng}}/{{ns}}.json",
                    request: function (e, n, t, r) {
                        var o;
                        switch (n) {
                            case "/__zh/translation.json":
                            case "/__zh-CN/translation.json":
                                o = s.zh;
                                break;
                            case "/__en/translation.json":
                            default:
                                o = s.en
                        }
                        o && o.then((function (e) {
                            r(null, {
                                status: 200,
                                data: e.data
                            })
                        }))
                    }
                },
                supportedLngs: ["en", "zh-CN"],
                fallbackLng: "en",
                interpolation: {
                    escapeValue: !1
                }
            });
            var a = t(67294),
                u = t(73935),
                l = t(83253),
                f = t.n(l),
                p = (t(57327), t(88767)),
                d = t(39711),
                h = t(96974),
                v = t(2804),
                b = t(46702),
                g = t(80043),
                x = t(87757),
                y = t.n(x),
                j = t(97943);

            function m() {
                return (m = (0, g.Z)(y().mark((function e(n, t) {
                    var r, o, i, c, s;
                    return y().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return r = {}, e.prev = 1, o = (0, j.g)(t), i = o.url, c = o.init, e.next = 5, fetch(i + n, c);
                            case 5:
                                if (!(s = e.sent).ok) {
                                    e.next = 10;
                                    break
                                }
                                return e.next = 9, s.json();
                            case 9:
                                r = e.sent;
                            case 10:
                                e.next = 15;
                                break;
                            case 12:
                                e.prev = 12, e.t0 = e.catch(1), console.log(`failed to fetch ${n}`, e.t0);
                            case 15:
                                return e.abrupt("return", r);
                            case 16:
                            case "end":
                                return e.stop()
                        }
                    }), e, null, [
                        [1, 12]
                    ])
                })))).apply(this, arguments)
            }
            var w = t(82569),
                O = t(85295),
                P = t(6055),
                C = "cHbZy_rAHf",
                k = "_2SNe_x81Ib",
                Z = "LUI6m76ply",
                N = t(85893);

            function S(e) {
                var n = e.name,
                    t = e.link,
                    r = e.version;
                return (0, N.jsxs)("div", {
                    className: C,
                    children: [(0, N.jsx)("h2", {
                        children: n
                    }), (0, N.jsxs)("p", {
                        children: [(0, N.jsx)("span", {
                            children: "Version "
                        }), (0, N.jsx)("span", {
                            className: k,
                            children: r
                        })]
                    }), (0, N.jsx)("p", {
                        children: (0, N.jsxs)("a", {
                            className: Z,
                            href: t,
                            target: "_blank",
                            rel: "noopener noreferrer",
                            children: [(0, N.jsx)(b.Z, {
                                size: 20
                            }), (0, N.jsx)("span", {
                                children: "Source"
                            })]
                        })
                    })]
                })
            }
            var E = (0, O.$j)((function (e) {
                    return {
                        apiConfig: (0, P.Y$)(e)
                    }
                }))((function (e) {
                    var n = (0, p.useQuery)(["/version", e.apiConfig], (function () {
                        return function (e, n) {
                            return m.apply(this, arguments)
                        }("/version", e.apiConfig)
                    })).data;
                    return (0, N.jsxs)(N.Fragment, {
                        children: [(0, N.jsx)(w.Z, {
                            title: "About"
                        }), n && n.version ? (0, N.jsx)(S, {
                            name: "Clash.Mini",
                            version: "0.1.1",
                            link: "https://github.com/Clash-Mini/Clash.Mini"

                        }): null,(0, N.jsx)(S, {
                            name: "Clash.Kernel",
                            version: "1.6.0 - Open source",
                            link: "https://github.com/Dreamacro/clash"
                        }),
                            (0, N.jsx)(S, {
                            name: "Yacd",
                            version: "0.2.15",
                            link: "https://github.com/haishanh/yacd"
                            })
                        ]
                    })
                })),
                D = (t(60285), t(64593));
            var R = (0, O.$j)((function (e) {
                    return {
                        apiConfig: (0, P.Y$)(e),
                        apiConfigs: (0, P.AJ)(e)
                    }
                }))((function (e) {
                    var n = e.apiConfig,
                        t = "Clash.Mini - Dashboard";
                    if (e.apiConfigs.length > 1) try {
                        t = `${new URL(n.baseURL).host} - yacd`
                    } catch (e) {}
                    return (0, N.jsx)(D.q, {
                        children: (0, N.jsx)("title", {
                            children: t
                        })
                    })
                })),
                L = t(82122),
                A = t(90433),
                I = new L.t,
                T = new A.S({
                    queryCache: I,
                    defaultOptions: {
                        queries: {
                            suspense: !0
                        }
                    }
                }),
                U = t(92669),
                _ = t(49522),
                $ = t(81125),
                B = t(71218),
                M = {
                    app: (0, P.E3)(),
                    modals: $.E3,
                    configs: U.E3,
                    proxies: B.E3,
                    logs: _.E3
                },
                F = {
                    selectChartStyleIndex: P.Pw,
                    updateAppConfig: P.N,
                    app: {
                        updateCollapsibleIsOpen: P.iB,
                        updateAppConfig: P.N,
                        removeClashAPIConfig: P.aj,
                        selectClashAPIConfig: P.O4
                    },
                    proxies: B.Nw
                },
                z = t(90924),
                W = t(50497),
                q = (t(21249), t(86010)),
                J = t(44309),
                G = t(12590),
                V = t(78268),
                H = a.useState,
                Q = a.useCallback;
            var Y = {
                    ul: "_1MP9tbO2C9",
                    li: "_39O4-s-qNw",
                    close: "_3U13UgV7Z1",
                    eye: "ipx6os2H89",
                    hasSecret: "_3GP8CDySTd",
                    url: "PK8GjRW5ZI",
                    secret: "_2-iwpHoCB6",
                    btn: "S1-PNvCcRP"
                },
                K = (0, O.$j)((function (e) {
                    return {
                        apiConfigs: (0, P.AJ)(e),
                        selectedClashAPIConfigIndex: (0, P.I4)(e)
                    }
                }))((function (e) {
                    var n = e.apiConfigs,
                        t = e.selectedClashAPIConfigIndex,
                        r = (0, O.WX)().app,
                        o = r.removeClashAPIConfig,
                        i = r.selectClashAPIConfig,
                        c = a.useCallback((function (e) {
                            o(e)
                        }), [o]),
                        s = a.useCallback((function (e) {
                            i(e)
                        }), [i]);
                    return (0, N.jsx)(N.Fragment, {
                        children: (0, N.jsx)("ul", {
                            className: Y.ul,
                            children: n.map((function (e, n) {
                                return (0, N.jsx)("li", {
                                    className: (0, q.Z)(Y.li, {
                                        [Y.hasSecret]: e.secret,
                                        [Y.isSelected]: n === t
                                    }),
                                    children: (0, N.jsx)(X, {
                                        disableRemove: n === t,
                                        baseURL: e.baseURL,
                                        secret: e.secret,
                                        onRemove: c,
                                        onSelect: s
                                    })
                                }, e.baseURL + e.secret)
                            }))
                        })
                    })
                }));

            function X(e) {
                var n = e.baseURL,
                    t = e.secret,
                    r = e.disableRemove,
                    o = e.onRemove,
                    i = e.onSelect,
                    c = function () {
                        var e = H(arguments.length > 0 && void 0 !== arguments[0] && arguments[0]),
                            n = (0, z.Z)(e, 2),
                            t = n[0],
                            r = n[1],
                            o = Q((function () {
                                return r((function (e) {
                                    return !e
                                }))
                            }), []);
                        return [t, o]
                    }(),
                    s = (0, z.Z)(c, 2),
                    u = s[0],
                    l = s[1],
                    f = u ? J.Z : G.Z,
                    p = a.useCallback((function (e) {
                        e.stopPropagation()
                    }), []);
                return (0, N.jsxs)(N.Fragment, {
                    children: [(0, N.jsx)(ee, {
                        disabled: r,
                        onClick: function () {
                            return o({
                                baseURL: n,
                                secret: t
                            })
                        }, className: Y.close,
                        children: (0, N.jsx)(V.Z, {
                            size: 20
                        })
                    }), (0, N.jsx)("span", {
                        className: Y.url,
                        tabIndex: 0,
                        role: "button",
                        onClick: function () {
                            return i({
                                baseURL: n,
                                secret: t
                            })
                        }, onKeyUp: p,
                        children: n
                    }), (0, N.jsx)("span", {}), t ? (0, N.jsxs)(N.Fragment, {
                        children: [(0, N.jsx)("span", {
                            className: Y.secret,
                            children: u ? t : "***"
                        }), (0, N.jsx)(ee, {
                            onClick: l,
                            className: Y.eye,
                            children: (0, N.jsx)(f, {
                                size: 20
                            })
                        })]
                    }) : null]
                })
            }

            function ee(e) {
                var n = e.children,
                    t = e.onClick,
                    r = e.className,
                    o = e.disabled;
                return (0, N.jsx)("button", {
                    disabled: o,
                    className: (0, q.Z)(r, Y.btn),
                    onClick: t,
                    children: n
                })
            }
            var ne = "_2-70itbuF1",
                te = "JKE-0c5hxF",
                re = "o2VhY_cs6Z",
                oe = "_3QDGxJ_pWs",
                ie = "tTZzzcEsTA",
                ce = "_2cCgtjpQZP",
                se = "_3OT00Mqmrw",
                ae = t(25904),
                ue = (t(82526), t(54747), t(49337), t(94949)),
                le = t(15116),
                fe = "_2uN43zExEi",
                pe = "_2gQ0j5OHC8";

            function de(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function he(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? de(Object(t), !0).forEach((function (n) {
                        (0, ue.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : de(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var ve = a.useCallback;

            function be(e) {
                var n = e.id,
                    t = e.label,
                    r = e.value,
                    o = e.onChange,
                    i = (0, le.Z)(e, ["id", "label", "value", "onChange"]),
                    c = ve((function (e) {
                        return o(e)
                    }), [o]),
                    s = (0, q.Z)({
                        [pe]: "string" == typeof r && "" !== r
                    });
                return (0, N.jsxs)("div", {
                    className: fe,
                    children: [(0, N.jsx)("input", he({
                        id: n,
                        value: r,
                        onChange: c
                    }, i)), (0, N.jsx)("label", {
                        htmlFor: n,
                        className: s,
                        children: t
                    })]
                })
            }
            var ge = t(4541),
                xe = a.useState,
                ye = a.useRef,
                je = a.useCallback;
            var me = (0, O.$j)((function (e) {
                return {
                    apiConfig: (0, P.Y$)(e)
                }
            }))((function (e) {
                var n = e.dispatch,
                    t = xe(""),
                    r = (0, z.Z)(t, 2),
                    o = r[0],
                    i = r[1],
                    c = xe(""),
                    s = (0, z.Z)(c, 2),
                    a = s[0],
                    u = s[1],
                    l = xe(""),
                    f = (0, z.Z)(l, 2),
                    p = f[0],
                    d = f[1],
                    h = ye(!1),
                    v = ye(null),
                    b = je((function (e) {
                        h.current = !0, d("");
                        var n = e.target,
                            t = n.name,
                            r = n.value;
                        switch (t) {
                            case "baseURL":
                                i(r);
                                break;
                            case "secret":
                                u(r);
                                break;
                            default:
                                throw new Error(`unknown input name ${t}`)
                        }
                    }), []),
                    g = je((function () {
                        (function (e) {
                            return we.apply(this, arguments)
                        })({
                            baseURL: o,
                            secret: a
                        }).then((function (e) {
                            0 !== e[0] ? d(e[1]) : n((0, P.sv)({
                                baseURL: o,
                                secret: a
                            }))
                        }))
                    }), [o, a, n]),
                    x = je((function (e) {
                        e.target instanceof Element && (!e.target.tagName || "INPUT" !== e.target.tagName.toUpperCase()) || "Enter" === e.key && g()
                    }), [g]);
                return (0, N.jsxs)("div", {
                    className: ne,
                    ref: v,
                    onKeyDown: x,
                    children: [(0, N.jsx)("div", {
                        className: te,
                        children: (0, N.jsx)("div", {
                            className: re,
                            children: (0, N.jsx)(ge.Z, {
                                width: 160,
                                height: 160
                            })
                        })
                    }), (0, N.jsx)("div", {
                        className: oe,
                        children: (0, N.jsxs)("div", {
                            className: ie,
                            children: [(0, N.jsx)(be, {
                                id: "baseURL",
                                name: "baseURL",
                                label: "API Base URL",
                                type: "text",
                                value: o,
                                onChange: b
                            }), (0, N.jsx)(be, {
                                id: "secret",
                                name: "secret",
                                label: "Secret(optional)",
                                value: a,
                                type: "text",
                                onChange: b
                            })]
                        })
                    }), (0, N.jsx)("div", {
                        className: ce,
                        children: p || null
                    }), (0, N.jsx)("div", {
                        className: se,
                        children: (0, N.jsx)(ae.Z, {
                            label: "Add",
                            onClick: g
                        })
                    }), (0, N.jsx)("div", {
                        style: {
                            height: 20
                        }
                    }), (0, N.jsx)(K, {})]
                })
            }));

            function we() {
                return (we = (0, g.Z)(y().mark((function e(n) {
                    var t, r;
                    return y().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                e.prev = 0, new URL(n.baseURL), e.next = 11;
                                break;
                            case 4:
                                if (e.prev = 4, e.t0 = e.catch(0), !n.baseURL) {
                                    e.next = 10;
                                    break
                                }
                                if ("http://" === (t = n.baseURL.substring(0, 7)) || "https:/" === t) {
                                    e.next = 10;
                                    break
                                }
                                return e.abrupt("return", [1, "Must starts with http:// or https://"]);
                            case 10:
                                return e.abrupt("return", [1, "Invalid URL"]);
                            case 11:
                                return e.prev = 11, e.next = 14, (0, W.T)(n);
                            case 14:
                                if (!((r = e.sent).status > 399)) {
                                    e.next = 17;
                                    break
                                }
                                return e.abrupt("return", [1, r.statusText]);
                            case 17:
                                return e.abrupt("return", [0]);
                            case 20:
                                return e.prev = 20, e.t1 = e.catch(11), e.abrupt("return", [1, "Failed to connect"]);
                            case 23:
                            case "end":
                                return e.stop()
                        }
                    }), e, null, [
                        [0, 4],
                        [11, 20]
                    ])
                })))).apply(this, arguments)
            }
            var Oe = {
                0: {
                    message: "Browser not supported!",
                    detail: 'This browser does not support "fetch", please choose another one.'
                },
                default: {
                    message: "Oops, something went wrong!"
                }
            };
            var Pe = "_2vs8ks4GvR",
                Ce = "Z8vSJz0PbL",
                ke = "EWfRQXOK8M",
                Ze = t(93621);

            function Ne(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function Se(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? Ne(Object(t), !0).forEach((function (n) {
                        (0, ue.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : Ne(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }

            function Ee(e) {
                var n = e.isOpen,
                    t = e.onRequestClose,
                    r = e.className,
                    o = e.overlayClassName,
                    i = e.children,
                    c = (0, le.Z)(e, ["isOpen", "onRequestClose", "className", "overlayClassName", "children"]),
                    s = (0, q.Z)(r, Ze.Z.content),
                    a = (0, q.Z)(o, Ze.Z.overlay);
                return (0, N.jsx)(f(), Se(Se({
                    isOpen: n,
                    onRequestClose: t,
                    className: s,
                    overlayClassName: a
                }, c), {}, {
                    children: i
                }))
            }
            var De = a.memo(Ee),
                Re = a.useCallback,
                Le = a.useEffect;
            var Ae = (0, O.$j)((function (e) {
                    return {
                        modals: e.modals,
                        apiConfig: (0, P.Y$)(e)
                    }
                }))((function (e) {
                    var n = e.dispatch,
                        t = e.apiConfig,
                        r = e.modals;
                    if (!window.fetch) {
                        var o = Oe[0].detail,
                            i = new Error(o);
                        throw i.code = 0, i
                    }
                    var c = Re((function () {
                        n((0, $.Mr)("apiConfig"))
                    }), [n]);
                    return Le((function () {
                        n((0, U.Tj)(t))
                    }), [n, t]), (0, N.jsx)(De, {
                        isOpen: r.apiConfig,
                        className: Pe,
                        overlayClassName: ke,
                        shouldCloseOnOverlayClick: !1,
                        shouldCloseOnEsc: !1,
                        onRequestClose: c,
                        children: (0, N.jsx)("div", {
                            className: Ce,
                            children: (0, N.jsx)(me, {})
                        })
                    })
                })),
                Ie = (t(12419), t(68670)),
                Te = t(83816),
                Ue = t(36678),
                _e = t(98766),
                $e = t(30929),
                Be = t(4656),
                Me = (t(92222), "_113JVByWGF"),
                Fe = "_1m2ZsnzFvt",
                ze = "_3TKFOM4Tgj";

            function We() {
                var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {},
                    n = e.width,
                    t = void 0 === n ? 24 : n,
                    r = e.height,
                    o = void 0 === r ? 24 : r;
                return (0, N.jsx)("svg", {
                    xmlns: "http://www.w3.org/2000/svg",
                    width: t,
                    height: o,
                    viewBox: "0 0 24 24",
                    fill: "none",
                    stroke: "currentColor",
                    strokeWidth: "2",
                    strokeLinecap: "round",
                    strokeLinejoin: "round",
                    children: (0, N.jsx)("path", {
                        d: "M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"
                    })
                })
            }
            var qe = function (e) {
                var n = e.message,
                    t = e.detail;
                return (0, N.jsxs)("div", {
                    className: Me,
                    children: [(0, N.jsx)("div", {
                        className: Fe,
                        children: (0, N.jsx)(ge.Z, {
                            width: 150,
                            height: 150
                        })
                    }), n ? (0, N.jsx)("h1", {
                        children: n
                    }) : null, t ? (0, N.jsx)("p", {
                        children: t
                    }) : null, (0, N.jsx)("p", {
                        children: (0, N.jsxs)("a", {
                            className: ze,
                            href: "https://github.com/haishanh/yacd/issues",
                            children: [(0, N.jsx)(We, {
                                width: 16,
                                height: 16
                            }), "haishanh/yacd"]
                        })
                    })]
                })
            };

            function Je(e) {
                var n = function () {
                    if ("undefined" == typeof Reflect || !Reflect.construct) return !1;
                    if (Reflect.construct.sham) return !1;
                    if ("function" == typeof Proxy) return !0;
                    try {
                        return Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], (function () {}))), !0
                    } catch (e) {
                        return !1
                    }
                }();
                return function () {
                    var t, r = (0, Be.Z)(e);
                    if (n) {
                        var o = (0, Be.Z)(this).constructor;
                        t = Reflect.construct(r, arguments, o)
                    } else t = r.apply(this, arguments);
                    return (0, $e.Z)(this, t)
                }
            }
            var Ge = function (e) {
                    (0, _e.Z)(t, e);
                    var n = Je(t);

                    function t() {
                        var e;
                        (0, Ie.Z)(this, t);
                        for (var r = arguments.length, o = new Array(r), i = 0; i < r; i++) o[i] = arguments[i];
                        return e = n.call.apply(n, [this].concat(o)), (0, ue.Z)((0, Ue.Z)(e), "state", {
                            error: null
                        }), e
                    }
                    return (0, Te.Z)(t, [{
                        key: "render",
                        value: function () {
                            if (this.state.error) {
                                var e = (r = this.state.error, "number" == typeof (o = r.code) ? Oe[o] : Oe.default),
                                    n = e.message,
                                    t = e.detail;
                                return (0, N.jsx)(qe, {
                                    message: n,
                                    detail: t
                                })
                            }
                            return this.props.children;
                            var r, o
                        }
                    }], [{
                        key: "getDerivedStateFromError",
                        value: function (e) {
                            return {
                                error: e
                            }
                        }
                    }]), t
                }(a.Component),
                Ve = t(64478),
                He = {
                    root: "_2kr0S-YLqE"
                },
                Qe = "_12V5kDiPEH",
                Ye = "_2FcudZSVil",
                Ke = function (e) {
                    var n = e.height,
                        t = n ? {
                            height: n
                        } : {};
                    return (0, N.jsx)("div", {
                        className: Qe,
                        style: t,
                        children: (0, N.jsx)("div", {
                            className: Ye
                        })
                    })
                },
                Xe = t(41289),
                en = t(66728),
                nn = t(35227);

            function tn(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function rn(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? tn(Object(t), !0).forEach((function (n) {
                        (0, ue.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : tn(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var on = a.useMemo,
                cn = {
                    position: "relative",
                    maxWidth: 1e3
                },
                sn = (0, O.$j)((function (e) {
                    return {
                        apiConfig: (0, P.Y$)(e),
                        selectedChartStyleIndex: (0, P.AM)(e)
                    }
                }))((function (e) {
                    var n = e.apiConfig,
                        t = e.selectedChartStyleIndex,
                        r = nn.A8.read(),
                        o = (0, Xe.r)(n),
                        i = (0, Ve.$)().t,
                        c = on((function () {
                            return {
                                labels: o.labels,
                                datasets: [rn(rn(rn({}, nn.IE), nn.Eu[t].up), {}, {
                                    label: i("Up"),
                                    data: o.up
                                }), rn(rn(rn({}, nn.IE), nn.Eu[t].down), {}, {
                                    label: i("Down"),
                                    data: o.down
                                })]
                            }
                        }), [o, t, i]);
                    return (0, en.Z)(r, "trafficChart", c, o), (0, N.jsx)("div", {
                        style: cn,
                        children: (0, N.jsx)("canvas", {
                            id: "trafficChart"
                        })
                    })
                }));
            var an = t(97750),
                un = t(11534),
                ln = "_2n4kL7wLDR",
                fn = a.useState,
                pn = a.useEffect,
                dn = a.useCallback,
                hn = (0, O.$j)((function (e) {
                    return {
                        apiConfig: (0, P.Y$)(e)
                    }
                }))((function (e) {
                    var n = e.apiConfig,
                        t = (0, Ve.$)().t,
                        r = function (e) {
                            var n = fn({
                                    upStr: "0 B/s",
                                    downStr: "0 B/s"
                                }),
                                t = (0, z.Z)(n, 2),
                                r = t[0],
                                o = t[1];
                            return pn((function () {
                                return (0, Xe.r)(e).subscribe((function (e) {
                                    return o({
                                        upStr: (0, un.Z)(e.up) + "/s",
                                        downStr: (0, un.Z)(e.down) + "/s"
                                    })
                                }))
                            }), [e]), r
                        }(n),
                        o = r.upStr,
                        i = r.downStr,
                        c = function (e) {
                            var n = fn({
                                    upTotal: "0 B",
                                    dlTotal: "0 B",
                                    connNumber: 0
                                }),
                                t = (0, z.Z)(n, 2),
                                r = t[0],
                                o = t[1],
                                i = dn((function (e) {
                                    var n = e.downloadTotal,
                                        t = e.uploadTotal,
                                        r = e.connections;
                                    o({
                                        upTotal: (0, un.Z)(t),
                                        dlTotal: (0, un.Z)(n),
                                        connNumber: r.length
                                    })
                                }), [o]);
                            return pn((function () {
                                return an.rQ(e, i)
                            }), [e, i]), r
                        }(n),
                        s = c.upTotal,
                        a = c.dlTotal,
                        u = c.connNumber;
                    return (0, N.jsxs)("div", {
                        className: ln,
                        children: [(0, N.jsxs)("div", {
                            className: "sec",
                            children: [(0, N.jsx)("div", {
                                children: t("Upload")
                            }), (0, N.jsx)("div", {
                                children: o
                            })]
                        }), (0, N.jsxs)("div", {
                            className: "sec",
                            children: [(0, N.jsx)("div", {
                                children: t("Download")
                            }), (0, N.jsx)("div", {
                                children: i
                            })]
                        }), (0, N.jsxs)("div", {
                            className: "sec",
                            children: [(0, N.jsx)("div", {
                                children: t("Upload Total")
                            }), (0, N.jsx)("div", {
                                children: s
                            })]
                        }), (0, N.jsxs)("div", {
                            className: "sec",
                            children: [(0, N.jsx)("div", {
                                children: t("Download Total")
                            }), (0, N.jsx)("div", {
                                children: a
                            })]
                        }), (0, N.jsxs)("div", {
                            className: "sec",
                            children: [(0, N.jsx)("div", {
                                children: t("Active Connections")
                            }), (0, N.jsx)("div", {
                                children: u
                            })]
                        })]
                    })
                }));

            function vn() {
                var e = (0, Ve.$)().t;
                return (0, N.jsxs)("div", {
                    children: [(0, N.jsx)(w.Z, {
                        title: e("Overview")
                    }), (0, N.jsxs)("div", {
                        className: He.root,
                        children: [(0, N.jsx)("div", {
                            children: (0, N.jsx)(hn, {})
                        }), (0, N.jsx)("div", {
                            className: He.chart,
                            children: (0, N.jsx)(a.Suspense, {
                                fallback: (0, N.jsx)(Ke, {
                                    height: "200px"
                                }),
                                children: (0, N.jsx)(sn, {})
                            })
                        })]
                    })]
                })
            }
            var bn = "_2fg1R7Zu62";
            var gn = function () {
                    return (0, N.jsx)("div", {
                        className: bn,
                        children: (0, N.jsx)(ge.Z, {
                            width: 280,
                            height: 280,
                            animate: !0,
                            c0: "transparent",
                            c1: "#646464"
                        })
                    })
                },
                xn = "_1X99PPGcn7",
                yn = "_2oV8uPP0dj",
                jn = t(98842),
                mn = t(73973),
                wn = t(59467),
                On = t(88757),
                Pn = "_3sTuXodYya",
                Cn = "_1WdrygzFVZ",
                kn = "q9nBJwAvlz",
                Zn = "_3yqSXpep4D",
                Nn = "_3wqPc_btNi",
                Sn = "_2sd59kze-t",
                En = "_2o6aUGu--S",
                Dn = "_11NITPe4W7",
                Rn = "_1kzEILbWHq",
                Ln = a.useCallback,
                An = {
                    activity: wn.o_x,
                    globe: wn.XUT,
                    command: wn.e71,
                    file: wn.NFo,
                    settings: wn.cKh,
                    link: wn.wWA
                },
                In = a.memo((function (e) {
                    var n = e.isActive,
                        t = e.to,
                        r = e.iconId,
                        o = e.labelText,
                        i = An[r],
                        c = (0, q.Z)(Zn, n ? Nn : null);
                    return (0, N.jsxs)(d.rU, {
                        to: t,
                        className: c,
                        children: [(0, N.jsx)(i, {}), (0, N.jsx)("div", {
                            className: Sn,
                            children: o
                        })]
                    })
                })),
                Tn = [{
                    to: "/",
                    iconId: "activity",
                    labelText: "Overview"
                }, {
                    to: "/proxies",
                    iconId: "globe",
                    labelText: "Proxies"
                }, {
                    to: "/rules",
                    iconId: "command",
                    labelText: "Rules"
                }, {
                    to: "/connections",
                    iconId: "link",
                    labelText: "Conns"
                }, {
                    to: "/configs",
                    iconId: "settings",
                    labelText: "Config"
                }, {
                    to: "/logs",
                    iconId: "file",
                    labelText: "Logs"
                }];

            function Un() {
                var e = On.U.read().motion;
                return (0, N.jsx)("svg", {
                    xmlns: "http://www.w3.org/2000/svg",
                    width: "20",
                    height: "20",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    stroke: "currentColor",
                    strokeWidth: "2",
                    strokeLinecap: "round",
                    strokeLinejoin: "round",
                    children: (0, N.jsx)(e.path, {
                        d: "M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z",
                        initial: {
                            rotate: -30
                        },
                        animate: {
                            rotate: 0
                        },
                        transition: {
                            duration: .7
                        }
                    })
                })
            }

            function _n() {
                var e = On.U.read().motion;
                return (0, N.jsxs)("svg", {
                    xmlns: "http://www.w3.org/2000/svg",
                    width: "20",
                    height: "20",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    stroke: "currentColor",
                    strokeWidth: "2",
                    strokeLinecap: "round",
                    strokeLinejoin: "round",
                    children: [(0, N.jsx)("circle", {
                        cx: "12",
                        cy: "12",
                        r: "5"
                    }), (0, N.jsxs)(e.g, {
                        initial: {
                            scale: .8
                        },
                        animate: {
                            scale: 1
                        },
                        transition: {
                            duration: .7
                        },
                        children: [(0, N.jsx)("line", {
                            x1: "12",
                            y1: "1",
                            x2: "12",
                            y2: "3"
                        }), (0, N.jsx)("line", {
                            x1: "12",
                            y1: "21",
                            x2: "12",
                            y2: "23"
                        }), (0, N.jsx)("line", {
                            x1: "4.22",
                            y1: "4.22",
                            x2: "5.64",
                            y2: "5.64"
                        }), (0, N.jsx)("line", {
                            x1: "18.36",
                            y1: "18.36",
                            x2: "19.78",
                            y2: "19.78"
                        }), (0, N.jsx)("line", {
                            x1: "1",
                            y1: "12",
                            x2: "3",
                            y2: "12"
                        }), (0, N.jsx)("line", {
                            x1: "21",
                            y1: "12",
                            x2: "23",
                            y2: "12"
                        }), (0, N.jsx)("line", {
                            x1: "4.22",
                            y1: "19.78",
                            x2: "5.64",
                            y2: "18.36"
                        }), (0, N.jsx)("line", {
                            x1: "18.36",
                            y1: "5.64",
                            x2: "19.78",
                            y2: "4.22"
                        })]
                    })]
                })
            }
            var $n = (0, O.$j)((function (e) {
                    return {
                        theme: (0, P.gh)(e)
                    }
                }))((function (e) {
                    var n = e.dispatch,
                        t = e.theme,
                        r = (0, Ve.$)().t,
                        o = (0, h.TH)(),
                        i = Ln((function () {
                            n((0, P.tj)())
                        }), [n]);
                    return (0, N.jsxs)("div", {
                        className: Pn,
                        children: [(0, N.jsx)("div", {
                            className: Cn
                        }), (0, N.jsx)("div", {
                            className: kn,
                            children: Tn.map((function (e) {
                                var n = e.to,
                                    t = e.iconId,
                                    i = e.labelText;
                                return (0, N.jsx)(In, {
                                    to: n,
                                    isActive: o.pathname === n,
                                    iconId: t,
                                    labelText: r(i)
                                }, n)
                            }))
                        }), (0, N.jsxs)("div", {
                            className: En,
                            children: [(0, N.jsx)(jn.ZP, {
                                label: r("theme"),
                                "aria-label": "switch to " + ("light" === t ? "dark" : "light") + " theme",
                                children: (0, N.jsx)("button", {
                                    className: (0, q.Z)(Dn, Rn),
                                    onClick: i,
                                    children: "light" === t ? (0, N.jsx)(Un, {}) : (0, N.jsx)(_n, {})
                                })
                            }), (0, N.jsx)(jn.ZP, {
                                label: r("about"),
                                children: (0, N.jsx)(d.rU, {
                                    to: "/about",
                                    className: Dn,
                                    children: (0, N.jsx)(mn.Z, {
                                        size: 20
                                    })
                                })
                            })]
                        })]
                    })
                })),
                Bn = t(34588),
                Mn = t(68970),
                Fn = t(17132),
                zn = t(26512),
                Wn = t(97148);

            function qn(e) {
                var n = function () {
                    if ("undefined" == typeof Reflect || !Reflect.construct) return !1;
                    if (Reflect.construct.sham) return !1;
                    if ("function" == typeof Proxy) return !0;
                    try {
                        return Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], (function () {}))), !0
                    } catch (e) {
                        return !1
                    }
                }();
                return function () {
                    var t, r = (0, Be.Z)(e);
                    if (n) {
                        var o = (0, Be.Z)(this).constructor;
                        t = Reflect.construct(r, arguments, o)
                    } else t = r.apply(this, arguments);
                    return (0, $e.Z)(this, t)
                }
            }

            function Jn(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function Gn(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? Jn(Object(t), !0).forEach((function (n) {
                        (0, ue.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : Jn(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var Vn = function () {},
                Hn = {
                    padding: "20px 0"
                },
                Qn = [{
                    label: "Global",
                    value: "Global"
                }, {
                    label: "Rule",
                    value: "Rule"
                }, {
                    label: "Direct",
                    value: "Direct"
                }],
                Yn = function (e) {
                    var n = e.children,
                        t = e.style;
                    return (0, N.jsx)("div", {
                        style: Gn(Gn({}, Hn), t),
                        children: n
                    })
                };

            function Kn() {
                var e = function () {
                        var e = arguments.length > 0 && void 0 !== arguments[0] && arguments[0],
                            n = a.useState(e),
                            t = (0, z.Z)(n, 2),
                            r = t[0],
                            o = t[1],
                            i = a.useCallback((function () {
                                o((function (e) {
                                    return !e
                                }))
                            }), []);
                        return [r, i]
                    }(!1),
                    n = (0, z.Z)(e, 2),
                    t = n[0],
                    r = n[1];
                return (0, N.jsx)(zn.Z, {
                    checked: t,
                    onChange: r
                })
            }
            a.PureComponent;
            var Xn = (0, a.lazy)((function () {
                    return Promise.all([t.e(776), t.e(88), t.e(170)]).then(t.bind(t, 64997))
                })),
                et = (0, a.lazy)((function () {
                    return t.e(497).then(t.bind(t, 9546))
                })),
                nt = (0, a.lazy)((function () {
                    return Promise.all([t.e(272), t.e(507)]).then(t.bind(t, 77098))
                })),
                tt = (0, a.lazy)((function () {
                    return Promise.all([t.e(776), t.e(857), t.e(641)]).then(t.bind(t, 22479))
                })),
                rt = (0, a.lazy)((function () {
                    return Promise.all([t.e(776), t.e(272), t.e(981)]).then(t.bind(t, 97193))
                })),
                ot = [{
                    path: "/",
                    element: (0, N.jsx)(vn, {})
                }, {
                    path: "/connections",
                    element: (0, N.jsx)(Xn, {})
                }, {
                    path: "/configs",
                    element: (0, N.jsx)(et, {})
                }, {
                    path: "/logs",
                    element: (0, N.jsx)(nt, {})
                }, {
                    path: "/proxies",
                    element: (0, N.jsx)(tt, {})
                }, {
                    path: "/rules",
                    element: (0, N.jsx)(rt, {})
                }, {
                    path: "/about",
                    element: (0, N.jsx)(E, {})
                }, !1].filter(Boolean);

            function it() {
                return (0, h.V$)(ot)
            }

            function ct() {
                return (0, N.jsxs)(N.Fragment, {
                    children: [(0, N.jsx)(Ae, {}), (0, N.jsx)($n, {}), (0, N.jsx)("div", {
                        className: yn,
                        children: (0, N.jsx)(a.Suspense, {
                            fallback: (0, N.jsx)(gn, {}),
                            children: (0, N.jsx)(it, {})
                        })
                    })]
                })
            }

            function st() {
                return (0, h.V$)([{
                    path: "/backend",
                    element: (0, N.jsx)(me, {})
                }, {
                    path: "*",
                    element: (0, N.jsx)(ct, {})
                }])
            }
            var at = function () {
                    return (0, N.jsx)(Ge, {
                        children: (0, N.jsx)(v.Wh, {
                            children: (0, N.jsx)(O.ZP, {
                                initialState: M,
                                actions: F,
                                children: (0, N.jsx)(p.QueryClientProvider, {
                                    client: T,
                                    children: (0, N.jsx)(d.UT, {
                                        children: (0, N.jsxs)("div", {
                                            className: xn,
                                            children: [(0, N.jsx)(R, {}), (0, N.jsx)(a.Suspense, {
                                                fallback: (0, N.jsx)(gn, {}),
                                                children: (0, N.jsx)(st, {})
                                            })]
                                        })
                                    })
                                })
                            })
                        })
                    })
                },
                ut = (t(4723), t(82772), Boolean("localhost" === window.location.hostname || "[::1]" === window.location.hostname || window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/)));

            function lt(e, n) {
                navigator.serviceWorker.register(e).then((function (e) {
                    e.onupdatefound = function () {
                        var t = e.installing;
                        null != t && (t.onstatechange = function () {
                            "installed" === t.state && (navigator.serviceWorker.controller ? (console.log("New content is available and will be used when all tabs for this page are closed. See https://cra.link/PWA."), n && n.onUpdate && n.onUpdate(e)) : (console.log("Content is cached for offline use."), n && n.onSuccess && n.onSuccess(e)))
                        })
                    }
                })).catch((function (e) {
                    console.error("Error during service worker registration:", e)
                }))
            }
            var ft = document.getElementById("app");
            f().setAppElement(ft), u.unstable_createRoot(ft).render((0, N.jsx)(at, {})),
                function (e) {
                    if ("serviceWorker" in navigator) {
                        if (new URL("", window.location.href).origin !== window.location.origin) return;
                        window.addEventListener("load", (function () {
                            var n = "/sw.js";
                            ut ? (! function (e, n) {
                                fetch(e, {
                                    headers: {
                                        "Service-Worker": "script"
                                    }
                                }).then((function (t) {
                                    var r = t.headers.get("content-type");
                                    404 === t.status || null != r && -1 === r.indexOf("javascript") ? navigator.serviceWorker.ready.then((function (e) {
                                        e.unregister().then((function () {
                                            window.location.reload()
                                        }))
                                    })) : lt(e, n)
                                })).catch((function () {
                                    console.log("No internet connection found. App is running in offline mode.")
                                }))
                            }(n, e), navigator.serviceWorker.ready.then((function () {
                                console.log("This web app is being served cache-first by a service worker")
                            }))) : lt(n, e)
                        }))
                    }
                }(), console.log("Checkout the repo: https://github.com/haishanh/yacd"), console.log("Version:", "0.2.15")
        }, 25904: function (e, n, t) {
            "use strict";
            t.d(n, {
                Z: function () {
                    return x
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(94949),
                o = t(15116),
                i = t(86010),
                c = t(67294),
                s = "_796AqwOFs_",
                a = "_1bLZvI40oA",
                u = "_1SrCTG7yDt",
                l = "_39VuJRXAmL",
                f = t(17132),
                p = t(85893);

            function d(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function h(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? d(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : d(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var v = c.forwardRef,
                b = c.useCallback;

            function g(e) {
                var n = e.children,
                    t = e.label,
                    r = e.text,
                    o = e.start;
                return (0, p.jsxs)(p.Fragment, {
                    children: [o ? (0, p.jsx)("span", {
                        className: u,
                        children: "function" == typeof o ? o() : o
                    }) : null, n || t || r]
                })
            }
            var x = v((function (e, n) {
                var t = e.onClick,
                    r = e.disabled,
                    c = void 0 !== r && r,
                    u = e.isLoading,
                    d = e.kind,
                    v = void 0 === d ? "primary" : d,
                    x = e.className,
                    y = e.children,
                    j = e.label,
                    m = e.text,
                    w = e.start,
                    O = (0, o.Z)(e, ["onClick", "disabled", "isLoading", "kind", "className", "children", "label", "text", "start"]),
                    P = {
                        children: y,
                        label: j,
                        text: m,
                        start: w
                    },
                    C = b((function (e) {
                        u || t && t(e)
                    }), [u, t]),
                    k = (0, i.Z)(s, {
                        [a]: "minimal" === v
                    }, x);
                return (0, p.jsx)("button", h(h({
                    className: k,
                    ref: n,
                    onClick: C,
                    disabled: c
                }, O), {}, {
                    children: u ? (0, p.jsxs)(p.Fragment, {
                        children: [(0, p.jsx)("span", {
                            style: {
                                display: "inline-flex",
                                opacity: 0
                            },
                            children: (0, p.jsx)(g, h({}, P))
                        }), (0, p.jsx)("span", {
                            className: l,
                            children: (0, p.jsx)(f.M, {})
                        })]
                    }) : (0, p.jsx)(g, h({}, P))
                }))
            }))
        }, 82569: function (e, n, t) {
            "use strict";
            t.d(n, {
                Z: function () {
                    return a
                }
            });
            var r = t(67294),
                o = "_24ddJm1Q5s",
                i = "B4QNkMu-0t",
                c = t(85893);

            function s(e) {
                var n = e.title;
                return (0, c.jsx)("div", {
                    className: o,
                    children: (0, c.jsx)("h1", {
                        className: i,
                        children: n
                    })
                })
            }
            var a = r.memo(s)
        }, 68970: function (e, n, t) {
            "use strict";
            t.d(n, {
                N: function () {
                    return b
                }, Z: function () {
                    return v
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(90924),
                o = t(15116),
                i = t(94949),
                c = t(67294),
                s = "_2DECxrOsTa",
                a = t(85893);

            function u(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function l(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? u(Object(t), !0).forEach((function (n) {
                        (0, i.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : u(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var f = c.useState,
                p = c.useRef,
                d = c.useEffect,
                h = c.useCallback;

            function v(e) {
                return (0, a.jsx)("input", l({
                    className: s
                }, e))
            }

            function b(e) {
                var n = e.value,
                    t = (0, o.Z)(e, ["value"]),
                    i = f(n),
                    c = (0, r.Z)(i, 2),
                    u = c[0],
                    v = c[1],
                    b = p(n);
                d((function () {
                    b.current !== n && v(n), b.current = n
                }), [n]);
                var g = h((function (e) {
                    return v(e.target.value)
                }), [v]);
                return (0, a.jsx)("input", l({
                    className: s,
                    value: u,
                    onChange: g
                }, t))
            }
        }, 85295: function (e, n, t) {
            "use strict";
            t.d(n, {
                WX: function () {
                    return m
                }, ZP: function () {
                    return w
                }, $j: function () {
                    return O
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(94949),
                o = t(90924),
                i = t(18172),
                c = t(67294),
                s = t(85893);

            function a(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function u(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? a(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : a(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            i.Fl(!1);
            var l = c.createContext,
                f = c.memo,
                p = c.useMemo,
                d = c.useRef,
                h = c.useEffect,
                v = c.useCallback,
                b = c.useContext,
                g = c.useState,
                x = l(null),
                y = l(null),
                j = l(null);

            function m() {
                return b(j)
            }

            function w(e) {
                var n = e.initialState,
                    t = e.actions,
                    r = void 0 === t ? {} : t,
                    c = e.children,
                    a = d(n),
                    u = g(n),
                    l = (0, o.Z)(u, 2),
                    f = l[0],
                    b = l[1],
                    m = v((function () {
                        return a.current
                    }), []);
                h((function () {
                    0
                }), [m]);
                var w = v((function (e, n) {
                        if ("function" == typeof e) return e(w, m);
                        var t = (0, i.ZP)(m(), n);
                        t !== a.current && (a.current = t, b(t))
                    }), [m]),
                    O = p((function () {
                        return C(r, w)
                    }), [r, w]);
                return (0, s.jsx)(x.Provider, {
                    value: f,
                    children: (0, s.jsx)(y.Provider, {
                        value: w,
                        children: (0, s.jsx)(j.Provider, {
                            value: O,
                            children: c
                        })
                    })
                })
            }

            function O(e) {
                return function (n) {
                    var t = f(n);
                    return function (n) {
                        var r = b(x),
                            o = b(y),
                            i = e(r, n),
                            c = u(u({
                                dispatch: o
                            }, n), i);
                        return (0, s.jsx)(t, u({}, c))
                    }
                }
            }

            function P(e, n) {
                return function () {
                    for (var t = arguments.length, r = new Array(t), o = 0; o < t; o++) r[o] = arguments[o];
                    return n(e.apply(this, r))
                }
            }

            function C(e, n) {
                var t = {};
                for (var r in e) {
                    var o = e[r];
                    "function" == typeof o ? t[r] = P(o, n) : "object" == typeof o && (t[r] = C(o, n))
                }
                return t
            }
        }, 4541: function (e, n, t) {
            "use strict";
            t.d(n, {
                Z: function () {
                    return c
                }
            });
            var r = t(86010),
                o = (t(67294), "_3MvgliBN_D"),
                i = t(85893);
            var c = function (e) {
                var n = e.width,
                    t = void 0 === n ? 320 : n,
                    c = e.height,
                    s = void 0 === c ? 320 : c,
                    a = e.animate,
                    u = void 0 !== a && a,
                    l = e.c0,
                    f = void 0 === l ? "currentColor" : l,
                    p = e.c1,
                    d = void 0 === p ? "#eee" : p,
                    h = (0, r.Z)({
                        [o]: u
                    });
                return (0, i.jsx)("svg", {
                    width: t,
                    height: s,
                    viewBox: "0 0 320 320",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, i.jsxs)("g", {
                        fill: "none",
                        fillRule: "evenodd",
                        children: [(0, i.jsx)("path", {
                            d: "M71.689 53.055c9.23-1.487 25.684 27.263 41.411 56.663 18.572-8.017 71.708-7.717 93.775 0 4.714-15.612 31.96-57.405 41.626-56.663 3.992.088 13.07 31.705 23.309 94.96 2.743 16.949 7.537 47.492 14.38 91.63-42.339 17.834-84.37 26.751-126.095 26.751-41.724 0-83.756-8.917-126.095-26.751C52.973 116.244 65.536 54.047 71.689 53.055z",
                            stroke: d,
                            strokeWidth: "4",
                            strokeLinecap: "round",
                            fill: f,
                            className: h
                        }), (0, i.jsx)("circle", {
                            fill: d,
                            cx: "216.5",
                            cy: "181.5",
                            r: "14.5"
                        }), (0, i.jsx)("circle", {
                            fill: d,
                            cx: "104.5",
                            cy: "181.5",
                            r: "14.5"
                        }), (0, i.jsx)("g", {
                            stroke: d,
                            strokeLinecap: "round",
                            strokeWidth: "4",
                            children: (0, i.jsx)("path", {
                                d: "M175.568 218.694c-2.494 1.582-5.534 2.207-8.563 1.508-3.029-.7-5.487-2.594-7.035-5.11M143.981 218.694c2.494 1.582 5.534 2.207 8.563 1.508 3.03-.7 5.488-2.594 7.036-5.11"
                            })
                        })]
                    })
                })
            }
        }, 26512: function (e, n, t) {
            "use strict";
            t(67294);
            var r = t(59936),
                o = t(6055),
                i = t(85295),
                c = t(85893);
            n.Z = (0, i.$j)((function (e) {
                return {
                    theme: (0, o.gh)(e)
                }
            }))((function (e) {
                var n = e.checked,
                    t = void 0 !== n && n,
                    o = e.onChange,
                    i = e.theme,
                    s = e.name,
                    a = "dark" === i ? "#393939" : "#e9e9e9";
                return (0, c.jsx)(r.default, {
                    onChange: o,
                    checked: t,
                    uncheckedIcon: !1,
                    checkedIcon: !1,
                    offColor: a,
                    onColor: "#792b4a",
                    offHandleColor: "#fff",
                    onHandleColor: "#fff",
                    handleDiameter: 24,
                    height: 28,
                    width: 44,
                    className: "rs",
                    name: s
                })
            }))
        }, 97148: function (e, n, t) {
            "use strict";
            t.d(n, {
                Z: function () {
                    return a
                }
            });
            t(82772), t(21249);
            var r = t(67294),
                o = "_2IgDTE__bQ",
                i = "_2IfOm9qQ_4",
                c = t(85893);

            function s(e) {
                var n = e.options,
                    t = e.value,
                    s = e.name,
                    a = e.onChange,
                    u = (0, r.useMemo)((function () {
                        return n.map((function (e) {
                            return e.value
                        })).indexOf(t)
                    }), [n, t]),
                    l = (0, r.useCallback)((function (e) {
                        var t = Math.floor(100 / n.length);
                        return e === n.length - 1 ? 100 - n.length * t + t : e > -1 ? t : void 0
                    }), [n]),
                    f = (0, r.useMemo)((function () {
                        return {
                            width: l(u) + "%",
                            left: u * l(0) + "%"
                        }
                    }), [u, l]);
                return (0, c.jsxs)("div", {
                    className: o,
                    children: [(0, c.jsx)("div", {
                        className: i,
                        style: f
                    }), n.map((function (e, n) {
                        var r = `${s}-${e.label}`,
                            o = 0 === n ? "" : "border-left";
                        return (0, c.jsxs)("label", {
                            htmlFor: r,
                            className: o,
                            style: {
                                width: l(n) + "%"
                            },
                            children: [(0, c.jsx)("input", {
                                id: r,
                                name: s,
                                type: "radio",
                                value: e.value,
                                checked: t === e.value,
                                onChange: a
                            }), (0, c.jsx)("div", {
                                children: e.label
                            })]
                        }, r)
                    }))]
                })
            }
            var a = r.memo(s)
        }, 17132: function (e, n, t) {
            "use strict";
            t.d(n, {
                M: function () {
                    return s
                }, $: function () {
                    return c
                }
            });
            t(67294);
            var r = "Q-CsP5Y3FT",
                o = "_3GL3LmFL_E",
                i = t(85893);

            function c(e) {
                var n = e.name,
                    t = e.type;
                return (0, i.jsxs)("h2", {
                    className: r,
                    children: [(0, i.jsx)("span", {
                        children: n
                    }), (0, i.jsx)("span", {
                        children: t
                    })]
                })
            }

            function s() {
                return (0, i.jsx)("span", {
                    className: o
                })
            }
        }, 66728: function (e, n, t) {
            "use strict";
            t.d(n, {
                Z: function () {
                    return l
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(94949),
                o = t(67294),
                i = t(35227);

            function c(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function s(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? c(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : c(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var a = o.useEffect,
                u = i.SB;

            function l(e, n, t, r) {
                var o = arguments.length > 4 && void 0 !== arguments[4] ? arguments[4] : {};
                a((function () {
                    var i = document.getElementById(n).getContext("2d"),
                        c = new e(i, {
                            type: "line",
                            data: t,
                            options: s(s({}, u), o)
                        }),
                        a = r && r.subscribe((function () {
                            return c.update()
                        }));
                    return function () {
                        a && a(), c.destroy()
                    }
                }), [e, n, t, r, o])
            }
        }, 35227: function (e, n, t) {
            "use strict";
            t.d(n, {
                A8: function () {
                    return i
                }, IE: function () {
                    return c
                }, SB: function () {
                    return s
                }, Eu: function () {
                    return a
                }
            });
            t(88674), t(41539), t(66992), t(33948);
            var r = t(4374),
                o = t(11534),
                i = (0, r.unstable_createResource)((function () {
                    return t.e(736).then(t.t.bind(t, 72037, 23)).then((function (e) {
                        return e.default
                    }))
                })),
                c = {
                    borderWidth: 1,
                    lineTension: 0,
                    pointRadius: 0
                },
                s = {
                    responsive: !0,
                    maintainAspectRatio: !0,
                    title: {
                        display: !1
                    },
                    legend: {
                        display: !0,
                        position: "top",
                        labels: {
                            fontColor: "#ccc",
                            boxWidth: 20
                        }
                    },
                    tooltips: {
                        enabled: !1,
                        mode: "index",
                        intersect: !1,
                        animationDuration: 100
                    },
                    hover: {
                        mode: "nearest",
                        intersect: !0
                    },
                    scales: {
                        xAxes: [{
                            display: !1,
                            gridLines: {
                                display: !1
                            }
                        }],
                        yAxes: [{
                            display: !0,
                            gridLines: {
                                display: !0,
                                color: "#555",
                                borderDash: [3, 6],
                                drawBorder: !1
                            },
                            ticks: {
                                callback: e => (0, o.Z)(e) + "/s "
                            }
                        }]
                    }
                },
                a = [{
                    down: {
                        backgroundColor: "rgba(176, 209, 132, 0.8)",
                        borderColor: "rgb(176, 209, 132)"
                    },
                    up: {
                        backgroundColor: "rgba(181, 220, 231, 0.8)",
                        borderColor: "rgb(181, 220, 231)"
                    }
                }, {
                    up: {
                        backgroundColor: "rgb(98, 190, 100)",
                        borderColor: "rgb(78,146,79)"
                    },
                    down: {
                        backgroundColor: "rgb(160, 230, 66)",
                        borderColor: "rgb(110, 156, 44)"
                    }
                }, {
                    up: {
                        backgroundColor: "rgba(94, 175, 223, 0.3)",
                        borderColor: "rgb(94, 175, 223)"
                    },
                    down: {
                        backgroundColor: "rgba(139, 227, 195, 0.3)",
                        borderColor: "rgb(139, 227, 195)"
                    }
                }, {
                    up: {
                        backgroundColor: "rgba(242, 174, 62, 0.3)",
                        borderColor: "rgb(242, 174, 62)"
                    },
                    down: {
                        backgroundColor: "rgba(69, 154, 248, 0.3)",
                        borderColor: "rgb(69, 154, 248)"
                    }
                }]
        }, 88757: function (e, n, t) {
            "use strict";
            t.d(n, {
                U: function () {
                    return r
                }
            });
            t(88674), t(41539), t(66992), t(33948);
            var r = function (e) {
                var n = {},
                    t = {},
                    r = {};

                function o() {
                    var o = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "default";
                    return t[o] = e(o).then((function (e) {
                        delete t[o], n[o] = e
                    })).catch((function (e) {
                        r[o] = e
                    })), t[o]
                }
                return {
                    preload: function () {
                        var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "default";
                        void 0 !== n[e] || t[e] || o(e)
                    }, read: function () {
                        var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "default";
                        if (void 0 !== n[e]) return n[e];
                        throw r[e] ? r[e] : t[e] ? t[e] : o(e)
                    }, clear: function (e) {
                        e ? delete n[e] : n = {}
                    }
                }
            }((function () {
                return t.e(354).then(t.bind(t, 47354))
            }))
        }, 11534: function (e, n, t) {
            "use strict";
            t.d(n, {
                Z: function () {
                    return o
                }
            });
            var r = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

            function o(e) {
                if (e < 1e3) return e + " B";
                var n = Math.min(Math.floor(Math.log10(e) / 3), r.length - 1);
                return (e = Number((e / Math.pow(1e3, n)).toPrecision(3))) + " " + r[n]
            }
        }, 97943: function (e, n, t) {
            "use strict";
            t.d(n, {
                g: function () {
                    return a
                }, P: function () {
                    return u
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(94949),
                o = (t(60285), t(41539), t(66992), t(33948), t(53062));

            function i(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }
            var c = {
                "Content-Type": "application/json"
            };

            function s(e) {
                var n = e.secret,
                    t = function (e) {
                        for (var n = 1; n < arguments.length; n++) {
                            var t = null != arguments[n] ? arguments[n] : {};
                            n % 2 ? i(Object(t), !0).forEach((function (n) {
                                (0, r.Z)(e, n, t[n])
                            })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : i(Object(t)).forEach((function (n) {
                                Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                            }))
                        }
                        return e
                    }({}, c);
                return n && (t.Authorization = `Bearer ${n}`), t
            }

            function a(e) {
                return {
                    url: e.baseURL,
                    init: {
                        headers: s({
                            secret: e.secret
                        })
                    }
                }
            }

            function u(e, n) {
                var t = e.baseURL,
                    r = e.secret,
                    i = "";
                "string" == typeof r && "" !== r && (i += "?token=" + encodeURIComponent(r));
                var c = new URL(t);
                return "https:" === c.protocol ? c.protocol = "wss:" : c.protocol = "ws:", `${(0,o.Os)(c.href)}${n}${i}`
            }
        }, 53062: function (e, n, t) {
            "use strict";
            t.d(n, {
                Ds: function () {
                    return r
                }, Os: function () {
                    return o
                }
            });
            t(15306);

            function r(e, n) {
                var t;
                return function () {
                    for (var r = arguments.length, o = new Array(r), i = 0; i < r; i++) o[i] = arguments[i];
                    t && clearTimeout(t), t = setTimeout((function () {
                        e.apply(void 0, o)
                    }), n)
                }
            }

            function o(e) {
                return e.replace(/\/$/, "")
            }
        }, 6055: function (e, n, t) {
            "use strict";
            t.d(n, {
                sv: function () {
                    return Z
                }, xE: function () {
                    return P
                }, Y$: function () {
                    return v
                }, AJ: function () {
                    return g
                }, VR: function () {
                    return m
                }, sU: function () {
                    return O
                }, Bg: function () {
                    return j
                }, S3: function () {
                    return w
                }, AM: function () {
                    return y
                }, I4: function () {
                    return b
                }, gh: function () {
                    return x
                }, E3: function () {
                    return U
                }, aj: function () {
                    return N
                }, Pw: function () {
                    return L
                }, O4: function () {
                    return S
                }, tj: function () {
                    return R
                }, N: function () {
                    return A
                }, iB: function () {
                    return I
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(94949),
                o = t(90924),
                i = t(80043),
                c = (t(40561), t(64765), t(23123), t(15306), t(60285), t(41539), t(66992), t(33948), t(87757)),
                s = t.n(c),
                a = "yacd.haishan.me";

            function u(e) {
                try {
                    var n = JSON.stringify(e);
                    localStorage.setItem(a, n)
                } catch (e) {}
            }
            var l, f, p = t(53062);
            t(92669), t(81125);

            function d(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function h(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? d(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : d(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var v = function (e) {
                    var n = e.app.selectedClashAPIConfigIndex;
                    return e.app.clashAPIConfigs[n]
                },
                b = function (e) {
                    return e.app.selectedClashAPIConfigIndex
                },
                g = function (e) {
                    return e.app.clashAPIConfigs
                },
                x = function (e) {
                    return e.app.theme
                },
                y = function (e) {
                    return e.app.selectedChartStyleIndex
                },
                j = function (e) {
                    return e.app.latencyTestUrl
                },
                m = function (e) {
                    return e.app.collapsibleIsOpen
                },
                w = function (e) {
                    return e.app.proxySortBy
                },
                O = function (e) {
                    return e.app.hideUnavailableProxies
                },
                P = function (e) {
                    return e.app.autoCloseOldConns
                },
                C = (0, p.Ds)(u, 600);

            function k(e, n) {
                for (var t = n.baseURL, r = n.secret, o = g(e()), i = 0; i < o.length; i++) {
                    var c = o[i];
                    if (c.baseURL === t && c.secret === r) return i
                }
            }

            function Z(e) {
                var n = e.baseURL,
                    t = e.secret;
                return function () {
                    var e = (0, i.Z)(s().mark((function e(r, o) {
                        var i;
                        return s().wrap((function (e) {
                            for (;;) switch (e.prev = e.next) {
                                case 0:
                                    if (!k(o, {
                                        baseURL: n,
                                        secret: t
                                    })) {
                                        e.next = 3;
                                        break
                                    }
                                    return e.abrupt("return");
                                case 3:
                                    i = {
                                        baseURL: n,
                                        secret: t,
                                        addedAt: Date.now()
                                    }, r("addClashAPIConfig", (function (e) {
                                        e.app.clashAPIConfigs.push(i)
                                    })), u(o().app);
                                case 6:
                                case "end":
                                    return e.stop()
                            }
                        }), e)
                    })));
                    return function (n, t) {
                        return e.apply(this, arguments)
                    }
                }()
            }

            function N(e) {
                var n = e.baseURL,
                    t = e.secret;
                return function () {
                    var e = (0, i.Z)(s().mark((function e(r, o) {
                        var i;
                        return s().wrap((function (e) {
                            for (;;) switch (e.prev = e.next) {
                                case 0:
                                    i = k(o, {
                                        baseURL: n,
                                        secret: t
                                    }), r("removeClashAPIConfig", (function (e) {
                                        e.app.clashAPIConfigs.splice(i, 1)
                                    })), u(o().app);
                                case 3:
                                case "end":
                                    return e.stop()
                            }
                        }), e)
                    })));
                    return function (n, t) {
                        return e.apply(this, arguments)
                    }
                }()
            }

            function S(e) {
                var n = e.baseURL,
                    t = e.secret;
                return function () {
                    var e = (0, i.Z)(s().mark((function e(r, o) {
                        var i;
                        return s().wrap((function (e) {
                            for (;;) switch (e.prev = e.next) {
                                case 0:
                                    i = k(o, {
                                        baseURL: n,
                                        secret: t
                                    }), b(o()) !== i && r("selectClashAPIConfig", (function (e) {
                                        e.app.selectedClashAPIConfigIndex = i
                                    })), u(o().app);
                                    try {
                                        window.location.reload()
                                    } catch (e) {}
                                case 5:
                                case "end":
                                    return e.stop()
                            }
                        }), e)
                    })));
                    return function (n, t) {
                        return e.apply(this, arguments)
                    }
                }()
            }
            var E = document.body;

            function D() {
                var e = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "dark";
                "dark" === e ? (E.classList.remove("light"), E.classList.add("dark")) : (E.classList.remove("dark"), E.classList.add("light"))
            }

            function R() {
                return function (e, n) {
                    var t = "light" === x(n()) ? "dark" : "light";
                    D(t), e("storeSwitchTheme", (function (e) {
                        e.app.theme = t
                    })), u(n().app)
                }
            }

            function L(e) {
                return function (n, t) {
                    n("appSelectChartStyleIndex", (function (n) {
                        n.app.selectedChartStyleIndex = Number(e)
                    })), u(t().app)
                }
            }

            function A(e, n) {
                return function (t, r) {
                    t("appUpdateAppConfig", (function (t) {
                        t.app[e] = n
                    })), u(r().app)
                }
            }

            function I(e, n, t) {
                return function (r, o) {
                    r("updateCollapsibleIsOpen", (function (r) {
                        r.app.collapsibleIsOpen[`${e}:${n}`] = t
                    })), C(o().app)
                }
            }
            var T = {
                selectedClashAPIConfigIndex: 0,
                clashAPIConfigs: [{
                    baseURL: null !== (l = null === (f = document.getElementById("app")) || void 0 === f ? void 0 : f.getAttribute("data-base-url")) && void 0 !== l ? l : "http://127.0.0.1:9090",
                    secret: "",
                    addedAt: 0
                }],
                latencyTestUrl: "http://www.gstatic.com/generate_204",
                selectedChartStyleIndex: 0,
                theme: "dark",
                collapsibleIsOpen: {},
                proxySortBy: "Natural",
                hideUnavailableProxies: !1,
                autoCloseOldConns: !1
            };

            function U() {
                var e = function () {
                    try {
                        var e = localStorage.getItem(a);
                        if (!e) return;
                        return JSON.parse(e)
                    } catch (e) {
                        return
                    }
                }();
                e = h(h({}, T), e);
                var n = function () {
                        var e = window.location.search,
                            n = {};
                        if ("string" != typeof e || "" === e) return n;
                        for (var t = e.replace(/^\?/, "").split("&"), r = 0; r < t.length; r++) {
                            var i = t[r].split("="),
                                c = (0, o.Z)(i, 2),
                                s = c[0],
                                a = c[1];
                            n[s] = encodeURIComponent(a)
                        }
                        return n
                    }(),
                    t = e.clashAPIConfigs[e.selectedClashAPIConfigIndex];
                if (t) {
                    var r = new URL(t.baseURL);
                    n.hostname && (r.hostname = n.hostname), n.port && (r.port = n.port), t.baseURL = (0, p.Os)(r.href), n.secret && (t.secret = n.secret)
                }
                return "dark" !== n.theme && "light" !== n.theme || (e.theme = n.theme), D(e.theme), e
            }
        }, 92669: function (e, n, t) {
            "use strict";
            t.d(n, {
                G_: function () {
                    return p
                }, ZO: function () {
                    return d
                }, Tj: function () {
                    return h
                }, wf: function () {
                    return v
                }, E3: function () {
                    return b
                }
            });
            t(82526), t(57327), t(54747), t(49337);
            var r = t(94949),
                o = t(80043),
                i = t(87757),
                c = t.n(i),
                s = t(50497),
                a = t(41289),
                u = t(81125);

            function l(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function f(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? l(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : l(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var p = function (e) {
                    return e.configs.configs
                },
                d = function (e) {
                    return e.configs.configs["log-level"]
                };

            function h(e) {
                return function () {
                    var n = (0, o.Z)(c().mark((function n(t, r) {
                        var o, i;
                        return c().wrap((function (n) {
                            for (;;) switch (n.prev = n.next) {
                                case 0:
                                    return n.prev = 0, n.next = 3, s.T(e);
                                case 3:
                                    o = n.sent, n.next = 10;
                                    break;
                                case 6:
                                    return n.prev = 6, n.t0 = n.catch(0), t((0, u.h7)("apiConfig")), n.abrupt("return");
                                case 10:
                                    if (o.ok) {
                                        n.next = 14;
                                        break
                                    }
                                    return console.log("Error fetch configs", o.statusText), t((0, u.h7)("apiConfig")), n.abrupt("return");
                                case 14:
                                    return n.next = 16, o.json();
                                case 16:
                                    i = n.sent, t("store/configs#fetchConfigs", (function (e) {
                                        e.configs.configs = i
                                    })), c = r(), c.configs.haveFetchedConfig ? a.r(e) : t((function (e) {
                                        e("store/configs#markHaveFetchedConfig", (function (e) {
                                            e.configs.haveFetchedConfig = !0
                                        }))
                                    }));
                                case 20:
                                case "end":
                                    return n.stop()
                            }
                            var c
                        }), n, null, [
                            [0, 6]
                        ])
                    })));
                    return function (e, t) {
                        return n.apply(this, arguments)
                    }
                }()
            }

            function v(e, n) {
                return function () {
                    var t = (0, o.Z)(c().mark((function t(r) {
                        return c().wrap((function (t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    s.w(e, n).then((function (e) {
                                        !1 === e.ok && console.log("Error update configs", e.statusText)
                                    }), (function (e) {
                                        throw console.log("Error update configs", e), e
                                    })).then((function () {
                                        r(h(e))
                                    })), r("storeConfigsOptimisticUpdateConfigs", (function (e) {
                                        e.configs.configs = f(f({}, e.configs.configs), n)
                                    }));
                                case 2:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function (e) {
                        return t.apply(this, arguments)
                    }
                }()
            }
            var b = {
                configs: {
                    port: 7890,
                    "socks-port": 7891,
                    "redir-port": 0,
                    "allow-lan": !1,
                    mode: "Rule",
                    "log-level": "info"
                },
                haveFetchedConfig: !1
            }
        }, 49522: function (e, n, t) {
            "use strict";
            t.d(n, {
                Rv: function () {
                    return c
                }, Xs: function () {
                    return s
                }, AR: function () {
                    return a
                }, TH: function () {
                    return u
                }, E3: function () {
                    return l
                }
            });
            t(57327), t(82772);
            var r = t(22222),
                o = function (e) {
                    return e.logs.logs
                },
                i = function (e) {
                    return e.logs.tail
                },
                c = function (e) {
                    return e.logs.searchText
                },
                s = (0, r.P1)(o, i, c, (function (e, n, t) {
                    for (var r = [], o = n; o >= 0; o--) r.push(e[o]);
                    if (300 === e.length)
                        for (var i = 299; i > n; i--) r.push(e[i]);
                    return "" === t ? r : r.filter((function (e) {
                        return e.payload.toLowerCase().indexOf(t) >= 0
                    }))
                }));

            function a(e) {
                return function (n) {
                    n("logsUpdateSearchText", (function (n) {
                        n.logs.searchText = e.toLowerCase()
                    }))
                }
            }

            function u(e) {
                return function (n, t) {
                    var r = t(),
                        c = o(r),
                        s = i(r),
                        a = s >= 299 ? 0 : s + 1;
                    c[a] = e, n("logsAppendLog", (function (e) {
                        e.logs.tail = a
                    }))
                }
            }
            var l = {
                searchText: "",
                logs: [],
                tail: -1
            }
        }, 81125: function (e, n, t) {
            "use strict";

            function r(e) {
                return function (n) {
                    n(`openModal:${e}`, (function (n) {
                        n.modals[e] = !0
                    }))
                }
            }

            function o(e) {
                return function (n) {
                    n(`closeModal:${e}`, (function (n) {
                        n.modals[e] = !1
                    }))
                }
            }
            t.d(n, {
                h7: function () {
                    return r
                }, Mr: function () {
                    return o
                }, E3: function () {
                    return i
                }
            });
            var i = {
                apiConfig: !1
            }
        }, 71218: function (e, n, t) {
            "use strict";
            t.d(n, {
                SJ: function () {
                    return L
                }, Nw: function () {
                    return re
                }, Ry: function () {
                    return B
                }, sj: function () {
                    return I
                }, yi: function () {
                    return A
                }, P_: function () {
                    return T
                }, a: function () {
                    return U
                }, DP: function () {
                    return $
                }, IA: function () {
                    return q
                }, E3: function () {
                    return D
                }, RE: function () {
                    return oe
                }, $3: function () {
                    return ee
                }, hU: function () {
                    return Y
                }, kL: function () {
                    return M
                }, AE: function () {
                    return F
                }
            });
            t(82526), t(54747), t(49337), t(47042), t(91038), t(41817);
            var r = t(94949),
                o = t(90924),
                i = t(80043),
                c = t(87757),
                s = t.n(c),
                a = (t(88674), t(41539), t(66992), t(33948), t(21249), t(57327), t(82772), t(2707), t(2804)),
                u = t(97750),
                l = t(97943);

            function f(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function p(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? f(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : f(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var d = "/proxies";

            function h(e) {
                return v.apply(this, arguments)
            }

            function v() {
                return (v = (0, i.Z)(s().mark((function e(n) {
                    var t, r, o, i;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return t = (0, l.g)(n), r = t.url, o = t.init, e.next = 3, fetch(r + d, o);
                            case 3:
                                return i = e.sent, e.next = 6, i.json();
                            case 6:
                                return e.abrupt("return", e.sent);
                            case 7:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function b(e, n, t) {
                return g.apply(this, arguments)
            }

            function g() {
                return (g = (0, i.Z)(s().mark((function e(n, t, r) {
                    var o, i, c, a, u;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return o = {
                                    name: r
                                }, i = (0, l.g)(n), c = i.url, a = i.init, u = `${c}/proxies/${t}`, e.next = 5, fetch(u, p(p({}, a), {}, {
                                    method: "PUT",
                                    body: JSON.stringify(o)
                                }));
                            case 5:
                                return e.abrupt("return", e.sent);
                            case 6:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function x(e, n) {
                return y.apply(this, arguments)
            }

            function y() {
                return (y = (0, i.Z)(s().mark((function e(n, t) {
                    var r, o, i, c, a, u, f = arguments;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return r = f.length > 2 && void 0 !== f[2] ? f[2] : "http://www.gstatic.com/generate_204", o = (0, l.g)(n), i = o.url, c = o.init, a = `timeout=5000&url=${r}`, u = `${i}/proxies/${encodeURIComponent(t)}/delay?${a}`, e.next = 6, fetch(u, c);
                            case 6:
                                return e.abrupt("return", e.sent);
                            case 7:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function j(e) {
                return m.apply(this, arguments)
            }

            function m() {
                return (m = (0, i.Z)(s().mark((function e(n) {
                    var t, r, o, i;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return t = (0, l.g)(n), r = t.url, o = t.init, e.next = 3, fetch(r + "/providers/proxies", o);
                            case 3:
                                if (404 !== (i = e.sent).status) {
                                    e.next = 6;
                                    break
                                }
                                return e.abrupt("return", {
                                    providers: {}
                                });
                            case 6:
                                return e.next = 8, i.json();
                            case 8:
                                return e.abrupt("return", e.sent);
                            case 9:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function w(e, n) {
                return O.apply(this, arguments)
            }

            function O() {
                return (O = (0, i.Z)(s().mark((function e(n, t) {
                    var r, o, i, c;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return r = (0, l.g)(n), o = r.url, i = r.init, c = p(p({}, i), {}, {
                                    method: "PUT"
                                }), e.next = 4, fetch(o + "/providers/proxies/" + t, c);
                            case 4:
                                return e.abrupt("return", e.sent);
                            case 5:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function P(e, n) {
                return C.apply(this, arguments)
            }

            function C() {
                return (C = (0, i.Z)(s().mark((function e(n, t) {
                    var r, o, i, c;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return r = (0, l.g)(n), o = r.url, i = r.init, c = p(p({}, i), {}, {
                                    method: "GET"
                                }), e.next = 4, fetch(o + "/providers/proxies/" + t + "/healthcheck", c);
                            case 4:
                                return e.abrupt("return", e.sent);
                            case 5:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }
            var k = t(6055);

            function Z(e, n) {
                var t;
                if ("undefined" == typeof Symbol || null == e[Symbol.iterator]) {
                    if (Array.isArray(e) || (t = function (e, n) {
                        if (!e) return;
                        if ("string" == typeof e) return N(e, n);
                        var t = Object.prototype.toString.call(e).slice(8, -1);
                        "Object" === t && e.constructor && (t = e.constructor.name);
                        if ("Map" === t || "Set" === t) return Array.from(e);
                        if ("Arguments" === t || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t)) return N(e, n)
                    }(e)) || n && e && "number" == typeof e.length) {
                        t && (e = t);
                        var r = 0,
                            o = function () {};
                        return {
                            s: o,
                            n: function () {
                                return r >= e.length ? {
                                    done: !0
                                } : {
                                    done: !1,
                                    value: e[r++]
                                }
                            }, e: function (e) {
                                throw e
                            }, f: o
                        }
                    }
                    throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }
                var i, c = !0,
                    s = !1;
                return {
                    s: function () {
                        t = e[Symbol.iterator]()
                    }, n: function () {
                        var e = t.next();
                        return c = e.done, e
                    }, e: function (e) {
                        s = !0, i = e
                    }, f: function () {
                        try {
                            c || null == t.return || t.return()
                        } finally {
                            if (s) throw i
                        }
                    }
                }
            }

            function N(e, n) {
                (null == n || n > e.length) && (n = e.length);
                for (var t = 0, r = new Array(n); t < n; t++) r[t] = e[t];
                return r
            }

            function S(e, n) {
                var t = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    n && (r = r.filter((function (n) {
                        return Object.getOwnPropertyDescriptor(e, n).enumerable
                    }))), t.push.apply(t, r)
                }
                return t
            }

            function E(e) {
                for (var n = 1; n < arguments.length; n++) {
                    var t = null != arguments[n] ? arguments[n] : {};
                    n % 2 ? S(Object(t), !0).forEach((function (n) {
                        (0, r.Z)(e, n, t[n])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : S(Object(t)).forEach((function (n) {
                        Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                    }))
                }
                return e
            }
            var D = {
                    proxies: {},
                    delay: {},
                    groupNames: [],
                    showModalClosePrevConns: !1
                },
                R = function () {
                    return null
                },
                L = ["Direct", "Fallback", "Reject", "Selector", "URLTest", "LoadBalance", "Unknown"],
                A = function (e) {
                    return e.proxies.proxies
                },
                I = function (e) {
                    return e.proxies.delay
                },
                T = function (e) {
                    return e.proxies.groupNames
                },
                U = function (e) {
                    return e.proxies.proxyProviders || []
                },
                _ = function (e) {
                    return e.proxies.dangleProxyNames
                },
                $ = function (e) {
                    return e.proxies.showModalClosePrevConns
                };

            function B(e) {
                return function () {
                    var n = (0, i.Z)(s().mark((function n(t, r) {
                        var i, c, a, u, l, f, p, d, v, b, g, x, y, m, w, O, P, C, k, N, S, D, R;
                        return s().wrap((function (n) {
                            for (;;) switch (n.prev = n.next) {
                                case 0:
                                    return n.next = 2, Promise.all([h(e), j(e)]);
                                case 2:
                                    for (i = n.sent, c = (0, o.Z)(i, 2), a = c[0], u = c[1], l = te(u.providers), f = l.providers, p = l.proxies, d = E(E({}, p), a.proxies), v = ne(d), b = (0, o.Z)(v, 2), g = b[0], x = b[1], y = I(r()), m = E({}, y), w = 0; w < x.length; w++) O = x[w], P = d[O] || {
                                        history: []
                                    }, C = P.history, (k = C[C.length - 1]) && "number" == typeof k.delay && (m[O] = {
                                        number: k.delay
                                    });
                                    N = [], S = Z(x);
                                    try {
                                        for (S.s(); !(D = S.n()).done;) R = D.value, p[R] || N.push(R)
                                    } catch (e) {
                                        S.e(e)
                                    } finally {
                                        S.f()
                                    }
                                    t("store/proxies#fetchProxies", (function (e) {
                                        e.proxies.proxies = d, e.proxies.groupNames = g, e.proxies.delay = m, e.proxies.proxyProviders = f, e.proxies.dangleProxyNames = N
                                    }));
                                case 16:
                                case "end":
                                    return n.stop()
                            }
                        }), n)
                    })));
                    return function (e, t) {
                        return n.apply(this, arguments)
                    }
                }()
            }

            function M(e, n) {
                return function () {
                    var t = (0, i.Z)(s().mark((function t(r) {
                        return s().wrap((function (t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return t.prev = 0, t.next = 3, w(e, n);
                                case 3:
                                    t.next = 7;
                                    break;
                                case 5:
                                    t.prev = 5, t.t0 = t.catch(0);
                                case 7:
                                    r(B(e));
                                case 8:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [0, 5]
                        ])
                    })));
                    return function (e) {
                        return t.apply(this, arguments)
                    }
                }()
            }

            function F(e, n) {
                return function () {
                    var t = (0, i.Z)(s().mark((function t(r) {
                        var o;
                        return s().wrap((function (t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    o = 0;
                                case 1:
                                    if (!(o < n.length)) {
                                        t.next = 12;
                                        break
                                    }
                                    return t.prev = 2, t.next = 5, w(e, n[o]);
                                case 5:
                                    t.next = 9;
                                    break;
                                case 7:
                                    t.prev = 7, t.t0 = t.catch(2);
                                case 9:
                                    o++, t.next = 1;
                                    break;
                                case 12:
                                    r(B(e));
                                case 13:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [2, 7]
                        ])
                    })));
                    return function (e) {
                        return t.apply(this, arguments)
                    }
                }()
            }

            function z(e, n) {
                return W.apply(this, arguments)
            }

            function W() {
                return (W = (0, i.Z)(s().mark((function e(n, t) {
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return e.prev = 0, e.next = 3, P(n, t);
                            case 3:
                                e.next = 7;
                                break;
                            case 5:
                                e.prev = 5, e.t0 = e.catch(0);
                            case 7:
                            case "end":
                                return e.stop()
                        }
                    }), e, null, [
                        [0, 5]
                    ])
                })))).apply(this, arguments)
            }

            function q(e, n) {
                return function () {
                    var t = (0, i.Z)(s().mark((function t(r) {
                        return s().wrap((function (t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return t.next = 2, z(e, n);
                                case 2:
                                    return t.next = 4, r(B(e));
                                case 4:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function (e) {
                        return t.apply(this, arguments)
                    }
                }()
            }

            function J() {
                return (J = (0, i.Z)(s().mark((function e(n, t, r) {
                    var o, i, c, a, l, f, p;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return e.next = 2, u.$K(n);
                            case 2:
                                return (o = e.sent).ok || console.log("unable to fetch all connections", o.statusText), e.next = 6, o.json();
                            case 6:
                                i = e.sent, c = i.connections, a = [], l = Z(c);
                                try {
                                    for (l.s(); !(f = l.n()).done;)(p = f.value).chains.indexOf(t) > -1 && p.chains.indexOf(r) < 0 && a.push(p.id)
                                } catch (e) {
                                    l.e(e)
                                } finally {
                                    l.f()
                                }
                                return e.next = 13, Promise.all(a.map((function (e) {
                                    return u.Sm(n, e).catch(R)
                                })));
                            case 13:
                            case "end":
                                return e.stop()
                        }
                    }), e)
                })))).apply(this, arguments)
            }

            function G(e, n, t, r, o) {
                return V.apply(this, arguments)
            }

            function V() {
                return (V = (0, i.Z)(s().mark((function e(n, t, r, o, i) {
                    var c;
                    return s().wrap((function (e) {
                        for (;;) switch (e.prev = e.next) {
                            case 0:
                                return e.prev = 0, e.next = 3, b(r, o, i);
                            case 3:
                                if (!1 !== e.sent.ok) {
                                    e.next = 6;
                                    break
                                }
                                throw new Error("failed to switch proxy: res.statusText");
                            case 6:
                                e.next = 12;
                                break;
                            case 8:
                                throw e.prev = 8, e.t0 = e.catch(0), console.log(e.t0, "failed to swith proxy"), e.t0;
                            case 12:
                                n(B(r)), (0, k.xE)(t()) && (c = A(t()), Q(r, c, {
                                    groupName: o,
                                    itemName: i
                                }));
                            case 15:
                            case "end":
                                return e.stop()
                        }
                    }), e, null, [
                        [0, 8]
                    ])
                })))).apply(this, arguments)
            }

            function H() {
                return function (e) {
                    e("closeModalClosePrevConns", (function (e) {
                        e.proxies.showModalClosePrevConns = !1
                    }))
                }
            }

            function Q(e, n, t) {
                var r = function (e, n, t) {
                    for (var r, o = [t, n], i = t;
                         (r = e[i]) && r.now;) o.unshift(r.now), i = r.now;
                    return o
                }(n, t.groupName, t.itemName);
                ! function (e, n, t) {
                    J.apply(this, arguments)
                }(e, t.groupName, r[0])
            }

            function Y(e, n, t) {
                return function () {
                    var r = (0, i.Z)(s().mark((function r(o, i) {
                        return s().wrap((function (r) {
                            for (;;) switch (r.prev = r.next) {
                                case 0:
                                    G(o, i, e, n, t).catch(R), o("store/proxies#switchProxy", (function (e) {
                                        var r = e.proxies.proxies;
                                        r[n] && r[n].now && (r[n].now = t)
                                    }));
                                case 2:
                                case "end":
                                    return r.stop()
                            }
                        }), r)
                    })));
                    return function (e, n) {
                        return r.apply(this, arguments)
                    }
                }()
            }

            function K(e, n) {
                return function () {
                    var t = (0, i.Z)(s().mark((function t(r, o) {
                        var i, c, a, u, l, f, p;
                        return s().wrap((function (t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return i = (0, k.Bg)(o()), t.next = 3, x(e, n, i);
                                case 3:
                                    return c = t.sent, a = "", !1 === c.ok && (a = c.statusText), t.next = 8, c.json();
                                case 8:
                                    u = t.sent, l = u.delay, f = I(o()), p = E(E({}, f), {}, {
                                        [n]: {
                                            error: a,
                                            number: l
                                        }
                                    }), r("requestDelayForProxyOnce", (function (e) {
                                        e.proxies.delay = p
                                    }));
                                case 13:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function (e, n) {
                        return t.apply(this, arguments)
                    }
                }()
            }

            function X(e, n) {
                return function () {
                    var t = (0, i.Z)(s().mark((function t(r) {
                        return s().wrap((function (t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return t.next = 2, r(K(e, n));
                                case 2:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function (e) {
                        return t.apply(this, arguments)
                    }
                }()
            }

            function ee(e) {
                return function () {
                    var n = (0, i.Z)(s().mark((function n(t, r) {
                        var o, i, c, a, u;
                        return s().wrap((function (n) {
                            for (;;) switch (n.prev = n.next) {
                                case 0:
                                    return o = _(r()), n.next = 3, Promise.all(o.map((function (n) {
                                        return t(X(e, n))
                                    })));
                                case 3:
                                    i = U(r()), c = Z(i), n.prev = 5, c.s();
                                case 7:
                                    if ((a = c.n()).done) {
                                        n.next = 13;
                                        break
                                    }
                                    return u = a.value, n.next = 11, z(e, u.name);
                                case 11:
                                    n.next = 7;
                                    break;
                                case 13:
                                    n.next = 18;
                                    break;
                                case 15:
                                    n.prev = 15, n.t0 = n.catch(5), c.e(n.t0);
                                case 18:
                                    return n.prev = 18, c.f(), n.finish(18);
                                case 21:
                                    return n.next = 23, t(B(e));
                                case 23:
                                case "end":
                                    return n.stop()
                            }
                        }), n, null, [
                            [5, 15, 18, 21]
                        ])
                    })));
                    return function (e, t) {
                        return n.apply(this, arguments)
                    }
                }()
            }

            function ne(e) {
                var n, t = [],
                    r = [];
                for (var o in e) {
                    var i = e[o];
                    i.all && Array.isArray(i.all) ? (t.push(o), "GLOBAL" === o && (n = i.all)) : L.indexOf(i.type) < 0 && r.push(o)
                }
                return n && (n.push("GLOBAL"), t = t.map((function (e) {
                    return [n.indexOf(e), e]
                })).sort((function (e, n) {
                    return e[0] - n[0]
                })).map((function (e) {
                    return e[1]
                }))), [t, r]
            }

            function te(e) {
                for (var n = Object.keys(e), t = [], r = {}, o = 0; o < n.length; o++) {
                    var i = e[n[o]];
                    if ("default" !== i.name && "Compatible" !== i.vehicleType) {
                        for (var c = i.proxies, s = [], a = 0; a < c.length; a++) {
                            var u = c[a];
                            r[u.name] = u, s.push(u.name)
                        }
                        i.proxies = s, t.push(i)
                    }
                }
                return {
                    providers: t,
                    proxies: r
                }
            }
            var re = {
                    requestDelayForProxies: function (e, n) {
                        return function () {
                            var t = (0, i.Z)(s().mark((function t(r, o) {
                                var i, c;
                                return s().wrap((function (t) {
                                    for (;;) switch (t.prev = t.next) {
                                        case 0:
                                            return i = _(o()), c = n.filter((function (e) {
                                                return i.indexOf(e) > -1
                                            })).map((function (n) {
                                                return r(X(e, n))
                                            })), t.next = 4, Promise.all(c);
                                        case 4:
                                            return t.next = 6, r(B(e));
                                        case 6:
                                        case "end":
                                            return t.stop()
                                    }
                                }), t)
                            })));
                            return function (e, n) {
                                return t.apply(this, arguments)
                            }
                        }()
                    }, closeModalClosePrevConns: H,
                    closePrevConnsAndTheModal: function (e) {
                        return function () {
                            var n = (0, i.Z)(s().mark((function n(t, r) {
                                var o, i, c, a;
                                return s().wrap((function (n) {
                                    for (;;) switch (n.prev = n.next) {
                                        case 0:
                                            if (i = r(), c = null === (o = i.proxies.switchProxyCtx) || void 0 === o ? void 0 : o.to) {
                                                n.next = 5;
                                                break
                                            }
                                            return t((function (e) {
                                                e("closeModalClosePrevConns", (function (e) {
                                                    e.proxies.showModalClosePrevConns = !1
                                                }))
                                            })), n.abrupt("return");
                                        case 5:
                                            a = i.proxies.proxies, Q(e, a, c), t("closePrevConnsAndTheModal", (function (e) {
                                                e.proxies.showModalClosePrevConns = !1, e.proxies.switchProxyCtx = void 0
                                            }));
                                        case 8:
                                        case "end":
                                            return n.stop()
                                    }
                                }), n)
                            })));
                            return function (e, t) {
                                return n.apply(this, arguments)
                            }
                        }()
                    }
                },
                oe = (0, a.cn)({
                    key: "proxyFilterText",
                    default: ""
                })
        }, 93621: function (e, n) {
            "use strict";
            n.Z = {
                overlay: "_2ueF0jmjym",
                content: "UZ5fqyDCOb"
            }
        }
    },
    function (e) {
        "use strict";
        var n;
        n = e.x, e.x = function () {
            var t = n();
            return [776, 88, 170, 857, 641, 272, 981, 507, 497, 736].map(e.E), t
        }
    },
    [
        [58392, 965, 977, 545, 623]
    ]
]);
//# sourceMappingURL=app.6706b8885424994ac6fe.js.map