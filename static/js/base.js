define(function (require, exports, module) {
    function a(a, c) {
        var a = a || "body", c = c || {}, h = c.curr || 0, w = c.menu || ".menu", g = c.event || "click",
            v = c.block || "block", b = c.currClass || "active", y = $(a).find(w), P = $(a).find(v);
        y.each(function (i) {
            $(this).attr("data-index", i)
        }), $(a).find(w).eq(h).show(), $(a).find(v).eq(h).show(), $(a).on(g, w, function () {
            k($(this)), c.callback && c.callback($(this).index())
        });
        var k = function (a) {
            var c = a.attr("data-index");
            y.removeClass(b), a.addClass(b), P.hide().eq(c).show()
        }
    }

    function c(a) {
        var c = 80, h = null, w = this;
        this.x = 0, this.leftPannel = $(a.left), this.rightPannel = $(a.right), this.dragger = $(a.dragger), this.container = $(a.container), this.dValue = a.dValue || 0, this.init = function () {
            this.dragger.on("mousedown", function (e) {
                w.x = e.clientX, w.lw = w.leftPannel.width(), w.rw = w.rightPannel.width(), w.mw = w.container.width(), w.mw = w.mw - w.dValue, w.minw = {
                    l: w.leftPannel.data("minw"),
                    r: w.rightPannel.data("minw")
                }, w.container.css({"user-select": "none"}).attr("unselectable", "on"), "function" == typeof a.start && a.start(), document.onmousemove = function (e) {
                    var g = e.clientX - w.x, v = w.lw;
                    g && (v = Math.max(w.minw.l, v + g), v = Math.min(v, w.mw - w.minw.r), "function" == typeof a.resize && a.resize(), w.rightPannel.css("width", w.mw - v), a.move && a.move(), h && clearTimeout(h), h = setTimeout(function () {
                        a.stop && a.stop()
                    }, c))
                }, document.onmouseup = function () {
                    w.x = 0, w.container.css({"user-select": "inherit"}).removeAttr("unselectable"), document.onmousemove = document.onmouseup = null, a.stop && a.stop()
                }
            })
        }, this.init()
    }

    var h = jQuery({});
    module.exports = {
        idrag: {
            create: function (a) {
                return new c(a)
            }
        }, tabSlider: {
            create: function (c, h) {
                return new a(c, h)
            }
        }, on: function () {
            h.on.apply(h, arguments)
        }, removeEvent: function () {
            h.off.apply(h, arguments)
        }, emit: function () {
            h.trigger.apply(h, arguments)
        }
    }
});