define(function (require) {
    function a(a) {
        location.href = a
    }

    function c(a, c, h) {
        var g = new Date;
        g.setDate(g.getDate() + h), document.cookie = a + "=" + escape(c) + (null == h ? "" : ";expires=" + g.toGMTString()) + ";path=/;domain=.imooc.com"
    }

    function h() {
        if ("number" != typeof continueTime) {
            continueTime = 0;
            var a = w.get("_vt");
            a && a[pageInfo.mid] && (continueTime = a[pageInfo.mid].st || 0)
        }
        g(continueTime)
    }

    function g(c) {
        var h = !1, g = window.location.protocol.split(":")[0],
            w = "http://www.imooc.com/course/playlist/" + pageInfo.mid + "?t=m3u8&_id=" + OP_CONFIG.mongo_id + "&ssl=0";
        "https" === g && (w = "https://www.imooc.com/course/playlist/" + pageInfo.mid + "?t=m3u8&_id=" + OP_CONFIG.mongo_id), window.thePlayer = mocoplayer($("#video-box"), {
            url: w, title: videoTitle, currentTime: c, events: {
                onReady: function () {
                    "html5" == this.getPlayType() && this.play()
                }, onComplete: function () {
                    window.clearInterval(y), h = !0, y = null;
                    var c = $(".J_next-box"), g = c.find(".J-next-course");
                    if (c.removeClass("hide").show(), b(), g.length) {
                        var w = $(".J-next-btn"), v = g.data("next-src");
                        w.removeClass("hide").on("click", function () {
                            a(v)
                        })
                    } else {
                        var _ = {
                            getData: function (a, c, h) {
                                var g = this;
                                $.ajax({
                                    type: "GET",
                                    url: "/course/" + a,
                                    data: {cid: course_id},
                                    dataType: "json",
                                    success: function (a) {
                                        if (0 == a.result) {
                                            var w = a.data;
                                            w.length && (g.bindData(w.slice(0, 2), c[0], h[0]), g.bindData(w.slice(2, 5), c[1], h[1]))
                                        }
                                    }
                                })
                            }, bindData: function (a, c, h) {
                                var g, w, v = "", y = "";
                                $(a).each(function (c) {
                                    2 == a[c].type ? (g = "//coding.imooc.com/class/" + a[c].id + ".html", w = a[c].pic) : (g = "/learn/" + a[c].id, w = "//img.mukewang.com/" + a[c].pic + "-240-135.jpg"), y += '<li class="l">                                                    <a href="' + g + '" target="_blank">                                                        <img src="' + w + '"/>                                                    </a>                                                    <a href="' + g + '" target="_blank" class="recom-course-title">                                                        ' + a[c].name + "                                                    </a>                                                </li>"
                                }), "first-row" == c ? (v = '<ul class="l">' + y + "</ul>", $(h).before(v)) : (v = '<div class="recom-course-row"><ul class="clearfix">' + y + "</ul></div>", $(h).after(v))
                            }
                        };
                        _.getData("ajaxlastmediarecom", ["first-row", "second-row"], [".first-row .other-oper", ".first-row"])
                    }
                }, onPlay: function () {
                    $(".J_next-box").hide()
                }, onError: function () {
                }
            }
        })
    }

    var w = require("store"), v = require("./common/animate-achievement");
    require("lib/juicer/juicer.min.js"), require("//static.mukewang.com/static/moco/player/3.0.5/mocoplayer"), require("/static/page/course/common/course_detail_common.js");
    var y = null;
    isLogin && h(), $(window).on("beforeunload", function () {
        function a() {
            var h, g, i = 0, w = (new Date).getTime();
            for (h in c) i++, c[h].t < w && (w = c[h].t, g = h);
            i >= 10 && (delete c[g], a())
        }

        var c = w.get("_vt") || {}, h = c[pageInfo.mid];
        if (thePlayer.getCurrentTime) return $(".J_next-box").hasClass("hide") ? void (h ? (h.t = (new Date).getTime(), h.st = thePlayer.getCurrentTime(), w.set("_vt", c)) : (h = {
            t: (new Date).getTime(),
            st: thePlayer.getCurrentTime()
        }, a(), c[pageInfo.mid] = h, w.set("_vt", c))) : (delete c[pageInfo.mid], void w.set("_vt", c))
    });
    var _ = function () {
        $.ajax({
            url: "//www.imooc.com/wechat/qrcode",
            type: "post",
            data: {method: -1},
            dataType: "json",
            success: function (a) {
                0 == a.result && ($(".js-wechat-box").hide(), c("showwechat", "1", 7))
            }
        })
    }, T = function () {
        var a = ($(".course-subnav").width() - $(".course-subnav ul").outerWidth()) / 2;
        a > 60 ? $("#js-ques-box").css("left", a) : $("#js-ques-box").css("left", 60)
    }, b = function () {
        if (OP_CONFIG.userInfo) {
            var a, c = {}, h = 0, g = (new Date).getTime();
            return c.mid = pageInfo.mid, y = window.setInterval(a = function () {
                var a, w;
                "object" == typeof thePlayer && thePlayer.getCurrentTime && (a = (new Date).getTime(), w = parseInt(a - g) / 1e3, c.time = w - h, c.learn_time = thePlayer.getCurrentTime(), $.ajax({
                    url: "/course/ajaxmediauser/",
                    data: c,
                    type: "POST",
                    dataType: "json",
                    success: function (a) {
                        if ("0" == a.result) {
                            h = w;
                            var c = a.data.media, g = a.data.course, a = [];
                            c && a.push({mp: c.mp.point, desc: c.mp.desc}), g && a.push({
                                mp: g.mp.point,
                                desc: g.mp.desc
                            }), v(a), c && ($("#J_ques_pop").show(), window.onresize = T, T())
                        } else "-105002" == a.result && $.alert("你的账号在另一地点登录，已被强迫下线。", {
                            info: "如果不是本人操作，建议你修改密码。",
                            callback: function () {
                                window.location.href = "//coding.imooc.com"
                            }
                        })
                    }
                }))
            }, 6e4), window.onbeforeunload = function () {
                var a, w;
                "object" == typeof thePlayer && thePlayer.getCurrentTime && (a = (new Date).getTime(), w = parseInt(a - g) / 1e3, c.time = w - h, c.learn_time = thePlayer.getCurrentTime(), $.ajax({
                    url: "/course/ajaxmediauser/",
                    data: c,
                    type: "POST",
                    async: !1,
                    dataType: "json",
                    success: function (a) {
                        "0" == a.result && (h = w)
                    }
                }))
            }, a
        }
    }();
    $("#js-ques-box").delegate(".Js-change-ques", "click", function () {
        var a, c = $("#ques-group li"), h = $("#ques-group").find("li.curr"), g = h.removeClass("curr").index(),
            w = c.length - 1;
        g === w ? c.eq(0).addClass("curr") : h.next().addClass("curr"), a = $("#ques-group .curr").find(".ques-title").attr("href"), $("#js-to-answer").attr("href", a)
    }), $("#js-ques-box").on("click", "#J_issue_Close,.ques-title,#js-to-answer", function () {
        $("#J_ques_pop").hide()
    }), $(document).on("click", ".js-wechat-close", function () {
        _()
    }), function () {
        var a = function () {
            var a = {};
            return a.resetSize = function () {
                var a = "100%", c = $("#header").outerHeight(!0), h = $(".js-course-subnav").outerHeight(!0),
                    g = $(window).height() - c - h;
                $(".js-box-wrap").css("width", a).css("height", g + "px"), $(".course-right-nano").css("height", g + "px"), g > 500 ? $(".question-tip-layer").css("marginTop", "6%") : g > 400 && 500 > g ? $(".question-tip-layer").css("marginTop", "3%") : 400 > g && $(".question-tip-layer").css("marginTop", "1%")
            }, a
        }(), c = function () {
            {
                var a = $(".js-course-menu"), c = $(".course-left"), h = c.offset().top, g = c.offset().left;
                $(".js-fixed-video")
            }
            return $(window).on("scroll", function () {
                var c = $(window).scrollTop();
                c >= h ? a.css("position", "fixed").css("left", g + "px") : a.css("position", "absolute").css("left", 0)
            }), {
                setLT: function () {
                    h = c.offset().top, g = c.offset().left
                }
            }
        }();
        a.resetSize(), c.setLT(), $(window).on("resize", function () {
            setTimeout(function () {
                a.resetSize(), c.setLT(), $(window).trigger("scroll")
            }, 200)
        })
    }(), function () {
        function a() {
            var n = 0, a = {};
            $.each(v, function (c, g) {
                $.getJSON(g[0], w[c], function (c) {
                    0 == c.result && (a[g[0]] = c.data), n++, n == v.length && h(a)
                })
            })
        }

        function c(a) {
            return $.each(a, function (a, c) {
                if (!isNaN(c.easy_type)) {
                    var h = "";
                    switch (Number(c.easy_type)) {
                        case 1:
                            h = "入门";
                            break;
                        case 2:
                            h = "初级";
                            break;
                        case 3:
                            h = "中级";
                            break;
                        case 4:
                            h = "高级";
                            break;
                        case 5:
                            h = "专业"
                    }
                    return c.easy_type = h, !0
                }
                return !1
            }), $.each(a, function (a, c) {
                "undefined" != typeof c.course_type ? "免费" == c.course_type ? c.course_type = "MF" : "实战" == c.course_type ? c.course_type = "SZ" : "路径" == c.course_type && (c.course_type = "LJ") : 1 == c.type ? c.course_type = "MF" : 2 == c.type && (c.course_type = "SZ")
            }), a
        }

        function h(a) {
            var h = "";
            $.each(v, function (w, v) {
                var y = a[v[0]];
                y && (h += juicer(g, {title: v[1], data: c(y)}))
            }), $(".video-panel .panel-container").append(h)
        }

        var g = ($("#v-teachers-tpl").html(), $("#v-course-tpl").html()),
            w = [{cid: course_id, tid: $("#teacher_id").val()}, {cid: course_id, type: "related"}],
            v = [["/lecturer/ajaxgettcourse", "讲师课程"], ["/course/ajaxcourserelate", "相关课程"]];
        a()
    }()
});