local limit_rate_domain = {
    ["gauss-otacostauto-cn.allawnfs.com"] = {
        { start_h = 19, end_h = 1, rate = "300k" },
        { start_h = 1, end_h = 8, rate = "6M" },
        { start_h = 8, end_h = 19, rate = "500k" }
    },
    ["gauss-sauapkcostauto-cn.allawnfs.com"] = {
        { start_h = 19, end_h = 1, rate = "300k" },
        { start_h = 1, end_h = 8, rate = "6M" },
        { start_h = 8, end_h = 19, rate = "500k" }
    },
    ["gauss-compotacostauto-cn.allawnfs.com"] = {
        { start_h = 21, end_h = 1, rate = "300k" },
        { start_h = 1, end_h = 2, rate = "3M" },
        { start_h = 2, end_h = 8, rate = "6M" },
        { start_h = 8, end_h = 14, rate = "500k" },
        { start_h = 14, end_h = 18, rate = "300k" },
        { start_h = 18, end_h = 19, rate = "500k" },
        { start_h = 19, end_h = 21, rate = "1M" }
    },
    ["adps-cdn-cn.allawnfs.com"] = {
        { start_h = 18, end_h = 1, rate = "300k" },
        { start_h = 1, end_h = 11, rate = "5M" },
        { start_h = 11, end_h = 14, rate = "1M" },
        { start_h = 14, end_h = 18, rate = "5M" }
    }
}
