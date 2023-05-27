package wap

var technologies = `{
    "$schema": "../schema.json",
    "categories": {
      "1": {
        "name": "CMS",
        "priority": 1
      },
      "2": {
        "name": "Message boards",
        "priority": 1
      },
      "3": {
        "name": "Database managers",
        "priority": 2
      },
      "4": {
        "name": "Documentation",
        "priority": 2
      },
      "5": {
        "name": "Widgets",
        "priority": 9
      },
      "6": {
        "name": "Ecommerce",
        "priority": 1
      },
      "7": {
        "name": "Photo galleries",
        "priority": 1
      },
      "8": {
        "name": "Wikis",
        "priority": 1
      },
      "9": {
        "name": "Hosting panels",
        "priority": 1
      },
      "10": {
        "name": "Analytics",
        "priority": 9
      },
      "11": {
        "name": "Blogs",
        "priority": 1
      },
      "12": {
        "name": "JavaScript frameworks",
        "priority": 8
      },
      "13": {
        "name": "Issue trackers",
        "priority": 2
      },
      "14": {
        "name": "Video players",
        "priority": 7
      },
      "15": {
        "name": "Comment systems",
        "priority": 9
      },
      "16": {
        "name": "Security",
        "priority": 9
      },
      "17": {
        "name": "Font scripts",
        "priority": 9
      },
      "18": {
        "name": "Web frameworks",
        "priority": 7
      },
      "19": {
        "name": "Miscellaneous",
        "priority": 9
      },
      "20": {
        "name": "Editors",
        "priority": 4
      },
      "21": {
        "name": "LMS",
        "priority": 1
      },
      "22": {
        "name": "Web servers",
        "priority": 8
      },
      "23": {
        "name": "Caching",
        "priority": 7
      },
      "24": {
        "name": "Rich text editors",
        "priority": 5
      },
      "25": {
        "name": "JavaScript graphics",
        "priority": 6
      },
      "26": {
        "name": "Mobile frameworks",
        "priority": 8
      },
      "27": {
        "name": "Programming languages",
        "priority": 5
      },
      "28": {
        "name": "Operating systems",
        "priority": 6
      },
      "29": {
        "name": "Search engines",
        "priority": 4
      },
      "30": {
        "name": "Webmail",
        "priority": 2
      },
      "31": {
        "name": "CDN",
        "priority": 9
      },
      "32": {
        "name": "Marketing automation",
        "priority": 9
      },
      "33": {
        "name": "Web server extensions",
        "priority": 7
      },
      "34": {
        "name": "Databases",
        "priority": 5
      },
      "35": {
        "name": "Maps",
        "priority": 6
      },
      "36": {
        "name": "Advertising",
        "priority": 9
      },
      "37": {
        "name": "Network devices",
        "priority": 2
      },
      "38": {
        "name": "Media servers",
        "priority": 1
      },
      "39": {
        "name": "Webcams",
        "priority": 9
      },
      "41": {
        "name": "Payment processors",
        "priority": 8
      },
      "42": {
        "name": "Tag managers",
        "priority": 9
      },
      "44": {
        "name": "CI",
        "priority": 3
      },
      "45": {
        "name": "Control systems",
        "priority": 2
      },
      "46": {
        "name": "Remote access",
        "priority": 1
      },
      "47": {
        "name": "Development",
        "priority": 2
      },
      "48": {
        "name": "Network storage",
        "priority": 2
      },
      "49": {
        "name": "Feed readers",
        "priority": 1
      },
      "50": {
        "name": "DMS",
        "priority": 1
      },
      "51": {
        "name": "Page builders",
        "priority": 2
      },
      "52": {
        "name": "Live chat",
        "priority": 9
      },
      "53": {
        "name": "CRM",
        "priority": 5
      },
      "54": {
        "name": "SEO",
        "priority": 8
      },
      "55": {
        "name": "Accounting",
        "priority": 1
      },
      "56": {
        "name": "Cryptominers",
        "priority": 5
      },
      "57": {
        "name": "Static site generator",
        "priority": 1
      },
      "58": {
        "name": "User onboarding",
        "priority": 8
      },
      "59": {
        "name": "JavaScript libraries",
        "priority": 9
      },
      "60": {
        "name": "Containers",
        "priority": 8
      },
      "62": {
        "name": "PaaS",
        "priority": 8
      },
      "63": {
        "name": "IaaS",
        "priority": 8
      },
      "64": {
        "name": "Reverse proxies",
        "priority": 7
      },
      "65": {
        "name": "Load balancers",
        "priority": 7
      },
      "66": {
        "name": "UI frameworks",
        "priority": 7
      },
      "67": {
        "name": "Cookie compliance",
        "priority": 9
      },
      "68": {
        "name": "Accessibility",
        "priority": 9
      },
      "69": {
        "name": "Social logins",
        "priority": 6
      },
      "70": {
        "name": "SSL/TLS certificate authorities",
        "priority": 9
      },
      "71": {
        "name": "Affiliate programs",
        "priority": 9
      },
      "72": {
        "name": "Appointment scheduling",
        "priority": 9
      },
      "73": {
        "name": "Surveys",
        "priority": 9
      },
      "74": {
        "name": "A/B Testing",
        "priority": 9
      },
      "75": {
        "name": "Email",
        "priority": 9
      },
      "76": {
        "name": "Personalisation",
        "priority": 9
      },
      "77": {
        "name": "Retargeting",
        "priority": 9
      },
      "78": {
        "name": "RUM",
        "priority": 9
      },
      "79": {
        "name": "Geolocation",
        "priority": 9
      },
      "80": {
        "name": "WordPress themes",
        "priority": 9
      },
      "81": {
        "name": "Shopify themes",
        "priority": 9
      },
      "82": {
        "name": "Drupal themes",
        "priority": 9
      },
      "83": {
        "name": "Browser fingerprinting",
        "priority": 9
      }
    },
    "technologies": {
      "1C-Bitrix": {
        "cats": [
          1,
          6
        ],
        "description": "1C-Bitrix is a system of web project management, universal software for the creation, support and successful development of corporate websites and online stores.",
        "headers": {
          "Set-Cookie": "BITRIX_",
          "X-Powered-CMS": "Bitrix Site Manager"
        },
        "html": "(?:<link[^>]+components/bitrix|(?:src|href)=\"/bitrix/(?:js|templates))",
        "icon": "1C-Bitrix.svg",
        "implies": "PHP",
        "scripts": "1c-bitrix",
        "website": "http://www.1c-bitrix.ru"
      },
      "33Across": {
        "cats": [
          36
        ],
        "description": "33Across is a technology company focused on solving the challenge of consumer attention for automated advertising.",
        "dom": "iframe[src*='.33across.com'], link[href*='.33across.com']",
        "icon": "33Across.png",
        "saas": true,
        "website": "https://www.33across.com",
        "xhr": "\\.33across\\.com"
      },
      "3dCart": {
        "cats": [
          1,
          6
        ],
        "cookies": {
          "3dvisit": ""
        },
        "headers": {
          "X-Powered-By": "3DCART"
        },
        "icon": "3dCart.png",
        "scripts": "(?:twlh(?:track)?\\.asp|3d_upsell\\.js)",
        "website": "http://www.3dcart.com"
      },
      "4-Tell": {
        "cats": [
          76
        ],
        "cookies": {
          "4Tell": "",
          "4TellCart": "",
          "4TellSession": ""
        },
        "description": "4-Tell is an ecommerce software company for retailers with AI-powered personalisation and recommendations products.",
        "icon": "4-Tell.png",
        "js": {
          "_4TellBoost": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "4tellcdn\\.azureedge\\.net",
        "website": "https://4-tell.com"
      },
      "@sulu/web": {
        "cats": [
          59
        ],
        "icon": "Sulu.svg",
        "js": {
          "web.startComponents": ""
        },
        "website": "https://github.com/sulu/web-js"
      },
      "A-Frame": {
        "cats": [
          25
        ],
        "html": "<a-scene[^<>]*>",
        "icon": "A-Frame.svg",
        "implies": "three.js",
        "js": {
          "AFRAME.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "/?([\\d.]+)?/aframe(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://aframe.io"
      },
      "A8.net": {
        "cats": [
          71
        ],
        "description": " A8.net is an affiliate marketing network.",
        "dom": "img[src*='.a8.net']",
        "icon": "A8.net.png",
        "js": {
          "A8salesCookieRepository": "",
          "a8sales": "",
          "map_A8": ""
        },
        "scripts": "statics\\.a8\\.net",
        "website": "https://www.a8.net"
      },
      "AB Tasty": {
        "cats": [
          74
        ],
        "description": "AB Tasty is a customer experience optimisation company. AB Tasty offers AI-driven experimentation, personalisation, and product optimisation platforms for user testing.",
        "icon": "AB Tasty.svg",
        "js": {
          "ABTasty": "",
          "_abtasty": "",
          "loadABTasty": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "try\\.abtasty\\.com",
        "website": "https://www.abtasty.com"
      },
      "AD EBiS": {
        "cats": [
          10
        ],
        "html": [
          "<!-- EBiS contents tag",
          "<!--EBiS tag",
          "<!-- Tag EBiS",
          "<!-- EBiS common tag"
        ],
        "icon": "ebis.png",
        "website": "http://www.ebis.ne.jp"
      },
      "ADPLAN": {
        "cats": [
          10
        ],
        "icon": "ADPLAN.png",
        "scripts": [
          "^https?://[^.]+\\.adplan7\\.com/\\;version:7",
          "^https?://(?!o\\.)\\w+\\.advg\\.jp/"
        ],
        "website": "https://www.adplan7.com/"
      },
      "AMP": {
        "cats": [
          12
        ],
        "description": "AMP, originally created by Google, is an open-source HTML framework developed by the AMP open-source Project. AMP is designed to help webpages load faster.",
        "html": [
          "<html[^>]* (?:amp|⚡)[^-]",
          "<link rel=\"amphtml\""
        ],
        "icon": "Accelerated-Mobile-Pages.svg",
        "oss": true,
        "website": "https://www.amp.dev",
        "xhr": "cdn\\.ampproject\\.org"
      },
      "AMP Plugin": {
        "cats": [
          1,
          5
        ],
        "icon": "Accelerated-Mobile-Pages.svg",
        "implies": "WordPress",
        "meta": {
          "generator": "^AMP Plugin v(\\d+\\.\\d+.*)$\\;version:\\1"
        },
        "website": "https://amp-wp.org"
      },
      "AOLserver": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:aol:aolserver",
        "headers": {
          "Server": "AOLserver/?([\\d.]+)?\\;version:\\1"
        },
        "icon": "AOLserver.png",
        "website": "http://aolserver.com"
      },
      "AOS": {
        "cats": [
          59
        ],
        "description": "JavaScript library to animate elements on your page as you scroll.",
        "icon": "AOS.svg",
        "js": {
          "AOS.init": "\\;confidence:25",
          "AOS.refresh": "\\;confidence:50",
          "AOS.refreshHard": "\\;confidence:50"
        },
        "scripts": "unpkg\\.com/aos@(next)/dist/aos\\.js\\;version:\\1",
        "website": "http://michalsnik.github.io/aos/"
      },
      "AT Internet Analyzer": {
        "cats": [
          10
        ],
        "icon": "AT Internet.png",
        "js": {
          "ATInternet": "",
          "xtsite": ""
        },
        "website": "http://atinternet.com/en"
      },
      "AT Internet XiTi": {
        "cats": [
          10
        ],
        "icon": "AT Internet.png",
        "js": {
          "xt_click": ""
        },
        "scripts": "xiti\\.com/hit\\.xiti",
        "website": "http://atinternet.com/en"
      },
      "ATSHOP": {
        "cats": [
          6
        ],
        "description": "ATSHOP is an all-in-one ecommerce platform.",
        "dom": "link[href*='cdn.atshop.io']",
        "icon": "ATSHOP.png",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.atshop\\.io",
        "website": "https://atshop.io"
      },
      "AWIN": {
        "cats": [
          71
        ],
        "cookies": {
          "BAGawin": "",
          "_aw_xid": ""
        },
        "description": "AWIN is a global affiliate marketing network.",
        "icon": "AWIN.svg",
        "js": {
          "AWIN.Tracking": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "dwin1\\.com",
        "website": "https://www.awin.com"
      },
      "AWS Certificate Manager": {
        "cats": [
          70
        ],
        "certIssuer": "Amazon",
        "description": "AWS Certificate Manager is a service that lets you easily provision, manage, and deploy public and private Secure Sockets Layer/Transport Layer Security (SSL/TLS) certificates for use with AWS services and your internal connected resources.",
        "icon": "aws.svg",
        "implies": "Amazon Web Services",
        "saas": true,
        "website": "https://aws.amazon.com/certificate-manager/"
      },
      "AWStats": {
        "cats": [
          10
        ],
        "cpe": "cpe:/a:laurent_destailleur:awstats",
        "icon": "AWStats.png",
        "implies": "Perl",
        "meta": {
          "generator": "AWStats ([\\d.]+(?: \\(build [\\d.]+\\))?)\\;version:\\1"
        },
        "website": "http://awstats.sourceforge.net"
      },
      "Abicart": {
        "cats": [
          6
        ],
        "description": "Abicart is an ecommerce platform developed by the Swedish company Abicart AB.",
        "icon": "abicart.png",
        "meta": {
          "generator": [
            "Abicart",
            "Textalk Webshop"
          ]
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://abicart.com/"
      },
      "AccessTrade": {
        "cats": [
          71
        ],
        "description": "AccessTrade is an affiliate marketing platform based on the CPA model developed by Interspace Co.",
        "dom": "img[src*='.accesstrade.net'],img[data-src*='.accesstrade.net']",
        "icon": "AccessTrade.png",
        "website": "https://www.accesstrade.ne.jp"
      },
      "AccessiBe": {
        "cats": [
          68
        ],
        "description": "AccessiBe is an automated web accessibility solution for ADA and WCAG compliance. The system scans and analyzes a website using AI technology, and applies all the required adjustments to become ADA and WCAG 2.1 compliant.",
        "icon": "AccessiBe.svg",
        "js": {
          "acsb": "\\;confidence:50",
          "acsbJS": "\\;confidence:50"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "acsbapp?\\.com/.*/acsb\\.js",
        "website": "https://accessibe.com"
      },
      "Accesso": {
        "cats": [
          6,
          72
        ],
        "description": "Accesso provides ticketing, ecommerce and Point-of-Sale (PoS) solutions.",
        "icon": "Accesso.svg",
        "js": {
          "accesso": ""
        },
        "scripts": "/embed/accesso\\.js",
        "website": "https://accesso.com/"
      },
      "Ackee": {
        "cats": [
          10
        ],
        "description": "Ackee is a self-hosted, Node.js based analytics tool with a focus on privacy.",
        "dom": "[data-ackee-domain-id], [data-ackee-server]",
        "icon": "Ackee.png",
        "js": {
          "ackeeTracker": ""
        },
        "oss": true,
        "website": "https://ackee.electerious.com"
      },
      "Acquia Cloud Platform": {
        "cats": [
          62
        ],
        "description": "Acquia Cloud Platform is a Drupal-tuned application lifecycle management suite with an infrastructure to support Drupal deployment workflow processes.",
        "headers": {
          "X-AH-Environment": "^\\w+$"
        },
        "icon": "acquia-cloud.png",
        "implies": [
          "Drupal\\;confidence:95",
          "Apache",
          "Percona",
          "Amazon EC2"
        ],
        "saas": true,
        "website": "https://www.acquia.com/"
      },
      "Act-On": {
        "cats": [
          32
        ],
        "description": "Act-On is a cloud-based SaaS product for marketing automation.",
        "icon": "Act-On.svg",
        "js": {
          "ActOn": ""
        },
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "/cdnr/\\d+/acton/bn/tracker/\\d+",
        "website": "http://act-on.com"
      },
      "Actito": {
        "cats": [
          32,
          76
        ],
        "cookies": {
          "SmartFocus": ""
        },
        "description": "Actito is an agile SaaS marketing automation platform.",
        "icon": "Actito.png",
        "js": {
          "_actGoal": "",
          "smartFocus": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.actito\\.be",
          "\\.advisor\\.smartfocus\\.com"
        ],
        "website": "https://www.actito.com"
      },
      "Acuity Scheduling": {
        "cats": [
          72
        ],
        "description": "Acuity Scheduling is a cloud-based appointment scheduling software solution.",
        "dom": "a[href*='app.acuityscheduling.com']",
        "icon": "Acuity Scheduling.png",
        "js": {
          "ACUITY_MODAL_INIT": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.acuityscheduling\\.com",
        "website": "https://acuityscheduling.com"
      },
      "Ad Lightning": {
        "cats": [
          36
        ],
        "description": "Ad Lightning is an programmatic ads monitoring and audit service.",
        "icon": "Ad Lightning.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.adlightning\\.com",
        "website": "https://www.adlightning.com"
      },
      "AdInfinity": {
        "cats": [
          36
        ],
        "icon": "AdInfinity.png",
        "scripts": "adinfinity\\.com\\.au",
        "website": "http://adinfinity.com.au"
      },
      "AdOcean": {
        "cats": [
          36
        ],
        "icon": "AdOcean.png",
        "implies": "Gemius",
        "js": {
          "ado.master": "",
          "ado.placement": "",
          "ado.slave": ""
        },
        "scripts": [
          "adocean\\.pl/files/js/ado\\.js",
          "adocean\\.pl\\;confidence:80"
        ],
        "website": "https://adocean-global.com"
      },
      "AdRiver": {
        "cats": [
          36
        ],
        "html": "(?:<embed[^>]+(?:src=\"https?://mh\\d?\\.adriver\\.ru/|flashvars=\"[^\"]*(?:http:%3A//(?:ad|mh\\d?)\\.adriver\\.ru/|adriver_banner))|<(?:(?:iframe|img)[^>]+src|a[^>]+href)=\"https?://ad\\.adriver\\.ru/)",
        "icon": "AdRiver.png",
        "js": {
          "adriver": ""
        },
        "scripts": "(?:adriver\\.core\\.\\d\\.js|https?://(?:content|ad|masterh\\d)\\.adriver\\.ru/)",
        "website": "http://adriver.ru"
      },
      "AdRoll": {
        "cats": [
          36,
          77
        ],
        "description": "AdRoll is a digital marketing technology platform that specializes in retargeting.",
        "icon": "AdRoll.svg",
        "js": {
          "adroll_adv_id": "",
          "adroll_pix_id": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "(?:a|s)\\.adroll\\.com",
        "website": "http://adroll.com"
      },
      "AdThrive": {
        "cats": [
          36
        ],
        "description": "AdThrive is an online advertising network aka ad provider for bloggers for blog monetisation.",
        "icon": "AdThrive.png",
        "js": {
          "adthrive": "",
          "adthriveVideosInjected": ""
        },
        "saas": true,
        "scripts": "ads\\.adthrive\\.com",
        "website": "https://www.adthrive.com"
      },
      "Ada": {
        "cats": [
          52
        ],
        "description": "Ada is an automated customer experience company that provides chat bots used in customer support.",
        "icon": "Ada.svg",
        "js": {
          "__AdaEmbedConstructor": "",
          "adaEmbed": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.ada\\.support",
        "website": "https://www.ada.cx"
      },
      "Adabra": {
        "cats": [
          32
        ],
        "description": "Adabra is a SaaS omnichannel marketing automation platform to help boost sales. Adabra allows you to manage user segmentation, create workflow and campaigns through email, social, SMS and more.",
        "icon": "Adabra.svg",
        "js": {
          "adabraPreview": "",
          "adabra_version_panel": "(^.+$)\\;version:\\1",
          "adabra_version_track": "(^.+$)\\;version:\\1"
        },
        "pricing": [
          "poa",
          "recurring"
        ],
        "saas": true,
        "scripts": "track\\.adabra\\.com",
        "website": "https://www.adabra.com",
        "xhr": "my\\.adabra\\.com"
      },
      "Adally": {
        "cats": [
          68
        ],
        "icon": "Adally.png",
        "scripts": "cloudfront\\.net/.*/adally\\.js",
        "website": "https://adally.com/"
      },
      "Adcash": {
        "cats": [
          36
        ],
        "icon": "Adcash.svg",
        "js": {
          "SuLoaded": "",
          "SuUrl": "",
          "ac_bgclick_URL": "",
          "ct_nOpp": "",
          "ct_nSuUrl": "",
          "ct_siteunder": "",
          "ct_tag": ""
        },
        "scripts": "^[^\\/]*//(?:[^\\/]+\\.)?adcash\\.com/(?:script|ad)/",
        "url": "^https?://(?:[^\\/]+\\.)?adcash\\.com/script/pop_",
        "website": "http://adcash.com"
      },
      "AddShoppers": {
        "cats": [
          5,
          10
        ],
        "description": "AddShoppers is the social media marketing command center for small-medium online retailers.",
        "icon": "AddShoppers.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "(?:cdn\\.)?shop\\.pe/widget/",
        "website": "http://www.addshoppers.com"
      },
      "AddThis": {
        "cats": [
          5
        ],
        "description": "AddThis is a social bookmarking service that can be integrated into a website with the use of a web widget.",
        "icon": "AddThis.svg",
        "js": {
          "addthis": ""
        },
        "scripts": "addthis\\.com/js/",
        "website": "http://www.addthis.com"
      },
      "AddToAny": {
        "cats": [
          5
        ],
        "description": "AddToAny is a universal sharing platform that can be integrated into a website by use of a web widget or plugin.",
        "icon": "AddToAny.png",
        "js": {
          "a2apage_init": ""
        },
        "scripts": "addtoany\\.com/menu/page\\.js",
        "website": "http://www.addtoany.com"
      },
      "Adminer": {
        "cats": [
          3
        ],
        "html": [
          "Adminer</a> <span class=\"version\">([\\d.]+)</span>\\;version:\\1",
          "onclick=\"bodyClick\\(event\\);\" onload=\"verifyVersion\\('([\\d.]+)'\\);\">\\;version:\\1"
        ],
        "icon": "adminer.png",
        "implies": "PHP",
        "website": "http://www.adminer.org"
      },
      "Admitad": {
        "cats": [
          71
        ],
        "description": "Admitad is an affiliate network that acts as an intermediary between advertisers and publishers.",
        "icon": "Admitad.svg",
        "js": {
          "ADMITAD": "",
          "admitad": ""
        },
        "pricing": [
          "payg"
        ],
        "scripts": [
          "artfut\\.com/static/(?:tracking|crossdevice)\\.min\\.js",
          "cdn\\.admitad\\.com"
        ],
        "website": "https://www.admitad.com"
      },
      "Adnegah": {
        "cats": [
          36
        ],
        "headers": {
          "X-Advertising-By": "adnegah\\.net"
        },
        "html": "<iframe [^>]*src=\"[^\"]+adnegah\\.net",
        "icon": "adnegah.png",
        "scripts": "[^a-z]adnegah.*\\.js$",
        "website": "https://Adnegah.net"
      },
      "Adobe Analytics": {
        "cats": [
          10
        ],
        "description": "Adobe Analytics is a web analytics, marketing and cross-channel analytics application.",
        "icon": "Adobe Analytics.svg",
        "js": {
          "s_c_il.0._c": "s_c",
          "s_c_il.0.constructor.name": "AppMeasurement",
          "s_c_il.1._c": "s_c",
          "s_c_il.1.constructor.name": "AppMeasurement",
          "s_c_il.2._c": "s_c",
          "s_c_il.2.constructor.name": "AppMeasurement",
          "s_c_il.3._c": "s_c",
          "s_c_il.3.constructor.name": "AppMeasurement",
          "s_c_il.4._c": "s_c",
          "s_c_il.4.constructor.name": "AppMeasurement",
          "s_c_il.5._c": "s_c",
          "s_c_il.5.constructor.name": "AppMeasurement"
        },
        "saas": true,
        "website": "https://www.adobe.com/analytics/adobe-analytics.html"
      },
      "Adobe ColdFusion": {
        "cats": [
          18
        ],
        "cpe": "cpe:/a:adobe:coldfusion",
        "headers": {
          "Cookie": "CFTOKEN="
        },
        "html": "<!-- START headerTags\\.cfm",
        "icon": "Adobe ColdFusion.svg",
        "implies": "CFML",
        "js": {
          "_cfEmails": ""
        },
        "scripts": "/cfajax/",
        "url": "\\.cfm(?:$|\\?)",
        "website": "http://adobe.com/products/coldfusion-family.html"
      },
      "Adobe DTM": {
        "cats": [
          42
        ],
        "description": "Dynamic Tag Management (DTM) is a tag management solution for Adobe Experience Cloud applications and others.",
        "icon": "adobedtm.png",
        "js": {
          "_satellite.buildDate": ""
        },
        "saas": true,
        "website": "https://marketing.adobe.com/resources/help/en_US/dtm/c_overview.html"
      },
      "Adobe Experience Manager": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:adobe:experience_manager",
        "description": "Adobe Experience Manager (AEM) is a content management solution for building websites, mobile apps and forms.",
        "html": [
          "<div class=\"[^\"]*parbase",
          "<div[^>]+data-component-path=\"[^\"+]jcr:",
          "<div class=\"[^\"]*aem-Grid"
        ],
        "implies": "Java",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "/etc/designs/",
          "/etc/clientlibs/",
          "/etc\\.clientlibs/"
        ],
        "website": "https://www.adobe.com/marketing/experience-manager.html"
      },
      "Adobe Experience Platform Identity Service": {
        "cats": [
          10
        ],
        "description": "Adobe Experience Platform Identity Service creates identity graphs that hold customer profiles and the known identifiers that belong to individual consumers.",
        "icon": "Adobe.svg",
        "js": {
          "s_c_il.0._c": "Visitor",
          "s_c_il.1._c": "Visitor",
          "s_c_il.2._c": "Visitor",
          "s_c_il.3._c": "Visitor",
          "s_c_il.4._c": "Visitor",
          "s_c_il.5._c": "Visitor"
        },
        "website": "https://docs.adobe.com/content/help/en/id-service/using/home.html"
      },
      "Adobe Experience Platform Launch": {
        "cats": [
          42
        ],
        "description": "Adobe Experience Cloud Launch is an extendable tag management solution for Adobe Experience Cloud, Adobe Experience Platform, and other applications.",
        "icon": "adobe_experience_platform.png",
        "js": {
          "_satellite.buildInfo": ""
        },
        "website": "https://docs.adobelaunch.com/getting-started"
      },
      "Adobe Flash": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:adobe:flash",
        "description": "Adobe Flash is a multimedia software platform used for production of animations, rich web applications and embedded web browser video players.",
        "dom": [
          "object[type='application/x-shockwave-flash']",
          "param[value*='.swf']"
        ],
        "icon": "Adobe Flash.svg",
        "website": "https://www.adobe.com/products/flashplayer"
      },
      "Adobe GoLive": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:adobe:golive",
        "description": "Adobe GoLive is a WYSIWYG HTML editor and web site management application.",
        "icon": "Adobe GoLive.png",
        "meta": {
          "generator": "Adobe GoLive(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "website": "http://www.adobe.com/products/golive"
      },
      "Adobe RoboHelp": {
        "cats": [
          4
        ],
        "cpe": "cpe:/a:adobe:robohelp",
        "description": "Adobe RoboHelp is a Help Authoring Tool (HAT) that allows you to create help systems, e-learning content and knowledge bases.",
        "icon": "Adobe RoboHelp.svg",
        "js": {
          "gbWhLang": "",
          "gbWhMsg": "",
          "gbWhProxy": "",
          "gbWhUtil": "",
          "gbWhVer": ""
        },
        "meta": {
          "generator": "^Adobe RoboHelp(?: ([\\d]+))?\\;version:\\1"
        },
        "scripts": "(?:wh(?:utils|ver|proxy|lang|topic|msg)|ehlpdhtm)\\.js",
        "website": "http://adobe.com/products/robohelp.html"
      },
      "Adobe Target": {
        "cats": [
          32,
          74
        ],
        "description": "Adobe Target is an A/B testing, multi-variate testing, personalisation, and optimisation application",
        "icon": "Adobe.svg",
        "js": {
          "adobe.target": "",
          "adobe.target.VERSION": "^(.+)$\\;version:\\1"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.adobe.com/marketing/target.html"
      },
      "AdonisJS": {
        "cats": [
          18
        ],
        "cookies": {
          "adonis-session": "",
          "adonis-session-values": ""
        },
        "icon": "AdonisJS.png",
        "implies": "Node.js",
        "website": "https://adonisjs.com"
      },
      "Advally": {
        "cats": [
          36
        ],
        "description": "Advally is an advertising platform for publishers.",
        "icon": "Advally.png",
        "js": {
          "advally": ""
        },
        "pricing": [
          "payg",
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.adligature\\.com/.+/advally-([\\d.]+)\\.js\\;version:\\1",
        "website": "https://www.advally.com"
      },
      "Advert Stream": {
        "cats": [
          36
        ],
        "icon": "Advert Stream.png",
        "js": {
          "advst_is_above_the_fold": ""
        },
        "scripts": "(?:ad\\.advertstream\\.com|adxcore\\.com)",
        "website": "http://www.advertstream.com"
      },
      "Adverticum": {
        "cats": [
          36
        ],
        "html": "<div (?:id=\"[a-zA-Z0-9_]*\" )?class=\"goAdverticum\"",
        "icon": "Adverticum.svg",
        "scripts": "(?:ad\\.)?adverticum\\.net/g3\\.js",
        "website": "http://adverticum.net"
      },
      "Adyen": {
        "cats": [
          41
        ],
        "description": "Adyen allows businesses to accept ecommerce, mobile, and point-of-sale payments.",
        "icon": "Adyen.svg",
        "js": {
          "adyen.encrypt.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://www.adyen.com"
      },
      "Adzerk": {
        "cats": [
          36
        ],
        "html": "<iframe [^>]*src=\"[^\"]+adzerk\\.net",
        "icon": "Adzerk.png",
        "js": {
          "ados": "",
          "adosResults": ""
        },
        "scripts": "adzerk\\.net/ados\\.js",
        "website": "http://adzerk.com"
      },
      "Aegea": {
        "cats": [
          11
        ],
        "headers": {
          "X-Powered-By": "^E2 Aegea v(\\d+)$\\;version:\\1"
        },
        "icon": "Aegea.png",
        "implies": [
          "PHP",
          "jQuery"
        ],
        "website": "http://blogengine.ru"
      },
      "Affiliate B": {
        "cats": [
          71,
          36
        ],
        "description": "Affiliate B is an advertising system that allows site operators (HP, blogs, e-mail newsletters, etc.) to place advertiser advertisements on their own sites.",
        "dom": "img[src*='www.afi-b.com']",
        "icon": "Affiliate B.svg",
        "scripts": "t\\.afi-b\\.com",
        "website": "https://affiliate-b.com"
      },
      "Affiliate Future": {
        "cats": [
          71
        ],
        "description": "Affiliate Future is a provider of advertisers with marketing solution through its affiliate network and tools.",
        "dom": "img[src*='banners.affiliatefuture.com']",
        "icon": "Affiliate Future.png",
        "scripts": "tags\\.affiliatefuture\\.com",
        "website": "http://affiliatefuture.com"
      },
      "Afosto": {
        "cats": [
          6
        ],
        "headers": {
          "X-Powered-By": "Afosto SaaS BV"
        },
        "icon": "Afosto.svg",
        "website": "http://afosto.com"
      },
      "AfterBuy": {
        "cats": [
          6
        ],
        "html": [
          "<dd>This OnlineStore is brought to you by ViA-Online GmbH Afterbuy\\. Information and contribution at https://www\\.afterbuy\\.de</dd>"
        ],
        "icon": "after-buy.png",
        "scripts": "shop-static\\.afterbuy\\.de",
        "website": "http://www.afterbuy.de"
      },
      "Afterpay": {
        "cats": [
          41
        ],
        "cpe": "cpe:/a:afterpay:afterpay",
        "description": "Afterpay is a 'buy now, pay later' platform that makes it possible to pay off purchased goods in fortnightly instalments.",
        "dom": "#afterpay, .afterpay, [aria-label='Afterpay']",
        "icon": "afterpay.png",
        "js": {
          "Afterpay": "",
          "afterpay_product": ""
        },
        "saas": true,
        "scripts": [
          "portal\\.afterpay\\.com",
          "static\\.afterpay\\.com"
        ],
        "website": "https://www.afterpay.com/"
      },
      "Ahoy": {
        "cats": [
          10
        ],
        "cookies": {
          "ahoy_track": "",
          "ahoy_visit": "",
          "ahoy_visitor": ""
        },
        "js": {
          "ahoy": ""
        },
        "website": "https://github.com/ankane/ahoy"
      },
      "Aimtell": {
        "cats": [
          32
        ],
        "description": "Aimtell is a cloud-hosted marketing platform that allows digital marketers and businesses to deliver web-based push notifications.",
        "icon": "Aimtell.png",
        "js": {
          "_aimtellLoad": "",
          "_aimtellPushToken": "",
          "_aimtellWebhook": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.aimtell\\.\\w+/",
        "website": "https://aimtell.com"
      },
      "Aircall": {
        "cats": [
          52
        ],
        "description": "Aircall is a cloud-based phone system for customer support and sales teams.",
        "icon": "aircall.png",
        "saas": true,
        "scripts": "^https?://cdn\\.aircall\\.io/",
        "website": "http://aircall.io"
      },
      "Airee": {
        "cats": [
          31
        ],
        "headers": {
          "Server": "^Airee"
        },
        "icon": "Airee.png",
        "website": "http://xn--80aqc2a.xn--p1ai"
      },
      "Airform": {
        "cats": [
          19
        ],
        "html": [
          "<form[^>]+?action=\"[^\"]*airform\\.io[^>]+?>"
        ],
        "icon": "Airform.svg",
        "website": "https://airform.io"
      },
      "Airship": {
        "cats": [
          32,
          10
        ],
        "description": "Airship is an American company that provides marketing and branding services. Airship allows companies to generate custom messages to consumers via push notifications, SMS messaging, and similar, and provides customer analytics services.",
        "icon": "Airship.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "urbanairship\\.\\w+/notify/v([\\d.]+)\\;version:\\1",
        "website": "https://www.airship.com"
      },
      "Akamai": {
        "cats": [
          31
        ],
        "description": "Akamai is global content delivery network (CDN) services provider for media and software delivery, and cloud security solutions.",
        "headers": {
          "X-Akamai-Transformed": "",
          "X-EdgeConnect-MidMile-RTT": "",
          "X-EdgeConnect-Origin-MEX-Latency": ""
        },
        "icon": "akamai.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "http://akamai.com"
      },
      "Akamai Bot Manager": {
        "cats": [
          16
        ],
        "cookies": {
          "ak_bmsc": "",
          "bm_sv": "",
          "bm_sz": ""
        },
        "description": "Akamai Bot Manager detect bots using device fingerprinting bot signatures.",
        "icon": "akamai.svg",
        "implies": "Akamai",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.akamai.com/us/en/products/security/bot-manager.jsp"
      },
      "Akamai Web Application Protector": {
        "cats": [
          16
        ],
        "description": "Akamai Web Application Protector is designed for companies looking for a more automated approach to web application firewall (WAF) and distributed denial-of-service (DDoS) security.",
        "icon": "akamai.svg",
        "implies": "Akamai",
        "js": {
          "AKSB": "\\;confidence:50"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "ds-aksb-a\\.akamaihd\\.net/aksb.min.js",
          "aksb\\.min\\.js\\;confidence:50"
        ],
        "website": "https://www.akamai.com/us/en/products/security/web-application-protector-enterprise-waf-firewall-ddos-protection.jsp"
      },
      "Akamai mPulse": {
        "cats": [
          10,
          78
        ],
        "cookies": {
          "akaas_AB-Testing": ""
        },
        "description": "Akamai mPulse is a real user monitoring (RUM) solution that enables companies to monitor, find, and fix website and application performance issues.",
        "html": [
          "<script>[\\s\\S]*?go-mpulse\\.net\\/boomerang[\\s\\S]*?</script>"
        ],
        "icon": "akamai.svg",
        "implies": [
          "Boomerang"
        ],
        "js": {
          "BOOMR_API_key": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://developer.akamai.com/akamai-mpulse-real-user-monitoring-solution"
      },
      "Akaunting": {
        "cats": [
          55
        ],
        "description": "Akaunting is a free and online accounting software.",
        "headers": {
          "X-Akaunting": "^Free Accounting Software$"
        },
        "html": [
          "<link[^>]+akaunting-green\\.css",
          "Powered By Akaunting: <a [^>]*href=\"https?://(?:www\\.)?akaunting\\.com[^>]+>"
        ],
        "icon": "akaunting.svg",
        "implies": "Laravel",
        "oss": true,
        "website": "https://akaunting.com"
      },
      "Akka HTTP": {
        "cats": [
          18,
          22
        ],
        "cpe": "cpe:/a:lightbend:akka_http",
        "headers": {
          "Server": "akka-http(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "akka-http.png",
        "website": "http://akka.io"
      },
      "Albacross": {
        "cats": [
          10,
          77
        ],
        "description": "Albacross is a lead generation and account intelligence platform. It helps marketing and sales teams identify their ideal customers visiting their website and gives them the insights they need to generate more qualified leads, make prospecting more efficient and close more deals.",
        "icon": "Albacross.svg",
        "js": {
          "_nQsv": "^([\\d.])$\\;version:\\1"
        },
        "pricing": [
          "low",
          "recurring",
          "poa"
        ],
        "saas": true,
        "scripts": "\\.albacross\\.com",
        "website": "https://albacross.com"
      },
      "AlertifyJS": {
        "cats": [
          12
        ],
        "description": "AlertifyJS is a javascript framework for developing browser dialogs and notifications.",
        "icon": "AlertifyJS.png",
        "js": {
          "alertify.defaults.autoReset": ""
        },
        "oss": true,
        "scripts": "/alertify/alertify\\.min\\.js",
        "website": "https://alertifyjs.com"
      },
      "Algolia": {
        "cats": [
          29
        ],
        "description": "Algolia offers a hosted web search product delivering real-time results.",
        "icon": "Algolia.svg",
        "js": {
          "AlgoliaSearch": "",
          "__algolia": "",
          "algoliasearch.version": "^(.+)$\\;version:\\1"
        },
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "website": "http://www.algolia.com"
      },
      "All in One SEO Pack": {
        "cats": [
          54
        ],
        "cpe": "cpe:/a:semperfiwebdesign:all_in_one_seo_pack",
        "description": "All in One SEO plugin optimizes WordPress website and its content for search engines.",
        "html": "<!-- All in One SEO Pack ([\\d.]+) \\;version:\\1",
        "icon": "all-in-One-SEO-Pack.png",
        "implies": "WordPress",
        "website": "https://wordpress.org/plugins/all-in-one-seo-pack/"
      },
      "AlloyUI": {
        "cats": [
          12
        ],
        "icon": "AlloyUI.png",
        "implies": [
          "Bootstrap",
          "YUI"
        ],
        "js": {
          "AUI": ""
        },
        "scripts": "^https?://cdn\\.alloyui\\.com/",
        "website": "http://www.alloyui.com"
      },
      "Alpine.js": {
        "cats": [
          12
        ],
        "html": "<[^>]+[^\\w-]x-data[^\\w-][^<]+\\;confidence:75",
        "icon": "Alpine.js.png",
        "js": {
          "Alpine.version": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "/alpine(?:\\.min)?\\.js"
        ],
        "website": "https://github.com/alpinejs/alpine"
      },
      "Amaya": {
        "cats": [
          20
        ],
        "description": "Amaya is an open-source web browser editor to create and update documents on the web.",
        "icon": "Amaya.png",
        "meta": {
          "generator": "Amaya(?: V?([\\d.]+[a-z]))?\\;version:\\1"
        },
        "website": "http://www.w3.org/Amaya"
      },
      "Amazon ALB": {
        "cats": [
          65
        ],
        "cookies": {
          "AWSALB": "",
          "AWSALBCORS": ""
        },
        "description": "Amazon Application Load Balancer (ALB) distributes incoming application traffic to increase availability and support content-based routing.",
        "icon": "aws-elb.png",
        "implies": "Amazon Web Services",
        "website": "https://aws.amazon.com/elasticloadbalancing/"
      },
      "Amazon Advertising": {
        "cats": [
          36
        ],
        "description": "Amazon Advertising (formerly AMS or Amazon Marketing Services) is a service that works in a similar way to pay-per-click ads on Google.",
        "icon": "Amazon.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.amazon-adsystem\\.com",
        "website": "https://advertising.amazon.com",
        "xhr": "\\.amazon-adsystem\\.com"
      },
      "Amazon Associates": {
        "cats": [
          71
        ],
        "description": "Amazon Associates is an affiliate marketing program that allows website owners and bloggers to create links and earn referral fees when customers click through and buy products from Amazon.",
        "dom": {
          "a[href*='amazon.com']": {
            "attributes": {
              "href": "^https?://amazon.com.+&tag="
            }
          },
          "a[href*='amzn.to']": {
            "attributes": {
              "href": ""
            }
          }
        },
        "icon": "Amazon.svg",
        "saas": true,
        "scripts": "\\.associates-amazon\\.com",
        "website": "https://affiliate-program.amazon.com"
      },
      "Amazon Cloudfront": {
        "cats": [
          31
        ],
        "description": "Amazon CloudFront is a fast content delivery network (CDN) service that securely delivers data, videos, applications, and APIs to customers globally with low latency, high transfer speeds.",
        "dns": {
          "CNAME": "^[a-z0-9]+\\.cloudfront.net\\.?$"
        },
        "headers": {
          "Via": "\\(CloudFront\\)$",
          "X-Amz-Cf-Id": ""
        },
        "icon": "Amazon-Cloudfront.svg",
        "implies": "Amazon Web Services",
        "saas": true,
        "website": "http://aws.amazon.com/cloudfront/"
      },
      "Amazon EC2": {
        "cats": [
          22
        ],
        "description": "Amazon Elastic Compute Cloud is a part of Amazon.com's cloud-computing platform, Amazon Web Services, that allows users to rent virtual computers on which to run their own computer applications.",
        "headers": {
          "Server": "\\(Amazon\\)"
        },
        "icon": "aws-ec2.svg",
        "implies": "Amazon Web Services",
        "saas": true,
        "website": "http://aws.amazon.com/ec2/"
      },
      "Amazon ECS": {
        "cats": [
          63
        ],
        "headers": {
          "Server": "^ECS"
        },
        "icon": "aws.svg",
        "implies": [
          "Amazon Web Services",
          "Docker"
        ],
        "saas": true,
        "website": "https://aws.amazon.com/elasticloadbalancing/"
      },
      "Amazon ELB": {
        "cats": [
          65
        ],
        "cookies": {
          "AWSELB": ""
        },
        "icon": "aws-elb.png",
        "implies": "Amazon Web Services",
        "saas": true,
        "website": "https://aws.amazon.com/elasticloadbalancing/"
      },
      "Amazon Pay": {
        "cats": [
          41
        ],
        "description": "Amazon Pay is an online payments processing service that is owned by Amazon. It lets you use the payment methods associated with your Amazon account to make payments for goods and services.",
        "dom": "img[src*='amazonpay']",
        "icon": "Amazon Pay.svg",
        "js": {
          "AmazonPayments": "",
          "OffAmazonPayments": "",
          "enableAmazonPay": "",
          "onAmazonPaymentsReady": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": [
          "/amazonpayments(?:\\.min)?\\.js",
          "\\.payments-amazon\\.com/OffAmazonPayments"
        ],
        "website": "https://pay.amazon.com",
        "xhr": "payments\\.amazon\\.com"
      },
      "Amazon S3": {
        "cats": [
          19
        ],
        "description": "Amazon S3 or Amazon Simple Storage Service is a service offered by Amazon Web Services (AWS) that provides object storage through a web service interface.",
        "headers": {
          "server": "^AmazonS3$"
        },
        "icon": "aws-s3.svg",
        "implies": "Amazon Web Services",
        "saas": true,
        "website": "http://aws.amazon.com/s3/"
      },
      "Amazon SES": {
        "cats": [
          75
        ],
        "description": "Amazon Simple Email Service (SES) is an email service that enables developers to send mail from within any application.",
        "dns": {
          "TXT": [
            "amazonses\\.com"
          ]
        },
        "icon": "aws.svg",
        "implies": "Amazon Web Services",
        "saas": true,
        "website": "https://aws.amazon.com/ses/"
      },
      "Amazon Web Services": {
        "cats": [
          62
        ],
        "description": "Amazon Web Services (AWS) is a comprehensive cloud services platform offering compute power, database storage, content delivery and other functionality.",
        "dns": {
          "NS": "\\.awsdns-"
        },
        "headers": {
          "x-amz-delete-marker": "",
          "x-amz-err-code": "",
          "x-amz-err-message": "",
          "x-amz-id-2": "",
          "x-amz-req-time-micros": "",
          "x-amz-request-id": "",
          "x-amz-rid": "",
          "x-amz-version-id": ""
        },
        "icon": "aws.svg",
        "saas": true,
        "website": "https://aws.amazon.com/"
      },
      "Amazon Webstore": {
        "cats": [
          6
        ],
        "description": "Amazon Webstore is an all-in-one hosted ecommerce website solution.",
        "icon": "Amazon Webstore.svg",
        "js": {
          "amzn": ""
        },
        "saas": true,
        "website": "https://aws.amazon.com/marketplace/pp/Amazon-Web-Services-Amazon-Webstore/B007NLVI2S"
      },
      "Amber": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "X-Powered-By": "^Amber$"
        },
        "icon": "amber.png",
        "website": "https://amberframework.org"
      },
      "American Express": {
        "cats": [
          41
        ],
        "description": "American Express, also known as Amex, facilitates electronic funds transfers throughout the world, most commonly through branded credit cards, debit cards and prepaid cards.",
        "html": "<[^>]+aria-labelledby=\"pi-american_express",
        "icon": "Amex.svg",
        "website": "https://www.americanexpress.com"
      },
      "Ametys": {
        "cats": [
          1
        ],
        "icon": "Ametys.png",
        "implies": "Java",
        "meta": {
          "generator": "(?:Ametys|Anyware Technologies)"
        },
        "scripts": "ametys\\.js",
        "website": "http://ametys.org"
      },
      "Amex Express Checkout": {
        "cats": [
          41
        ],
        "description": "Amex Express Checkout is a service that simplifies the checkout experience by auto-filling necessary cardholder payment data into merchant checkout fields.",
        "icon": "amex.png",
        "scripts": "aexp-static\\.com",
        "website": "https://www.americanexpress.com/us/express-checkout/"
      },
      "Amiro.CMS": {
        "cats": [
          1
        ],
        "icon": "Amiro.CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "Amiro"
        },
        "website": "http://amirocms.com"
      },
      "Amplitude": {
        "cats": [
          10
        ],
        "description": "Amplitude is a web and mobile analytics solution with cross-platform user journey tracking, user behavior analysis and segmentation capabilities.",
        "icon": "Amplitude.svg",
        "js": {
          "AMPLITUDE_KEY": ""
        },
        "pricing": [
          "freemium",
          "poa",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.amplitude\\.com",
        "website": "https://amplitude.com",
        "xhr": "\\.amplitude\\.com"
      },
      "Analysys Ark": {
        "cats": [
          10
        ],
        "cookies": {
          "ARK_ID": ""
        },
        "icon": "Analysys Ark.svg",
        "js": {
          "AnalysysAgent": ""
        },
        "scripts": "AnalysysFangzhou_JS_SDK\\.min\\.js\\?v=([\\d.]+)\\;version:\\1",
        "website": "https://ark.analysys.cn"
      },
      "Anetwork": {
        "cats": [
          36
        ],
        "icon": "Anetwork.png",
        "scripts": "static-cdn\\.anetwork\\.ir/",
        "website": "https://www.anetwork.ir"
      },
      "Angular": {
        "cats": [
          12
        ],
        "cpe": "cpe:/a:angularjs:angular.js",
        "description": "Angular is a TypeScript-based open-source web application framework led by the Angular Team at Google.",
        "dom": {
          "[ng-version]": {
            "attributes": {
              "ng-version": "^([\\d.]+)\\;version:\\1"
            }
          }
        },
        "excludes": [
          "AngularDart",
          "AngularJS"
        ],
        "icon": "Angular.svg",
        "implies": "TypeScript",
        "js": {
          "ng.coreTokens": "",
          "ng.probe": ""
        },
        "oss": true,
        "website": "https://angular.io"
      },
      "Angular Material": {
        "cats": [
          66
        ],
        "description": "Angular Material is a UI component library for Angular JS developers. Angular Material components assist in constructing attractive, consistent, and functional web pages and web applications.",
        "icon": "AngularJS.svg",
        "implies": "AngularJS",
        "js": {
          "ngMaterial": ""
        },
        "scripts": "/([\\d.rc-]+)?/angular-material(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://material.angularjs.org"
      },
      "AngularDart": {
        "cats": [
          18
        ],
        "excludes": [
          "Angular",
          "AngularJS"
        ],
        "icon": "AngularDart.svg",
        "implies": "Dart",
        "js": {
          "ngTestabilityRegistries": ""
        },
        "website": "https://webdev.dartlang.org/angular/"
      },
      "AngularJS": {
        "cats": [
          12
        ],
        "description": "AngularJS is a JavaScript-based open-source web application framework led by the Angular Team at Google.",
        "excludes": [
          "Angular",
          "AngularDart"
        ],
        "html": [
          "<(?:div|html)[^>]+ng-app=",
          "<ng-app"
        ],
        "icon": "AngularJS.svg",
        "js": {
          "angular": "",
          "angular.version.full": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "angular[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "/([\\d.]+(?:-?rc[.\\d]*)*)/angular(?:\\.min)?\\.js\\;version:\\1",
          "\\bangular.{0,32}\\.js"
        ],
        "website": "https://angularjs.org"
      },
      "Ant Design": {
        "cats": [
          12
        ],
        "html": [
          "<[^>]*class=\"ant-(?:btn|col|row|layout|breadcrumb|menu|pagination|steps|select|cascader|checkbox|calendar|form|input-number|input|mention|rate|radio|slider|switch|tree-select|time-picker|transfer|upload|avatar|badge|card|carousel|collapse|list|popover|tooltip|table|tabs|tag|timeline|tree|alert|modal|message|notification|progress|popconfirm|spin|anchor|back-top|divider|drawer)",
          "<i class=\"anticon anticon-"
        ],
        "icon": "Ant Design.svg",
        "js": {
          "antd": ""
        },
        "website": "https://ant.design"
      },
      "AnyClip": {
        "cats": [
          36
        ],
        "description": "AnyClip is a video engagement platform that uses an AI-driven content analysis engine to analyze and categorize video content in real-time to create personalised video feeds.",
        "dom": "img[src*='.anyclip.com'], video[poster*='.anyclip.com']",
        "icon": "AnyClip.svg",
        "js": {
          "anyclip": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.anyclip\\.com",
        "website": "https://www.anyclip.com"
      },
      "Apache": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:apache:http_server",
        "description": "Apache is a free and open-source cross-platform web server software.",
        "headers": {
          "Server": "(?:Apache(?:$|/([\\d.]+)|[^/-])|(?:^|\\b)HTTPD)\\;version:\\1"
        },
        "icon": "Apache.svg",
        "website": "http://apache.org"
      },
      "Apache JSPWiki": {
        "cats": [
          8
        ],
        "cpe": "cpe:/a:apache:jspwiki",
        "description": "Apache JSPWiki is an open-source Wiki engine, built around standard JEE components (Java, servlets, JSP).",
        "html": "<html[^>]* xmlns:jspwiki=",
        "icon": "Apache JSPWiki.png",
        "implies": "Apache Tomcat",
        "scripts": "jspwiki",
        "url": "wiki\\.jsp",
        "website": "http://jspwiki.org"
      },
      "Apache Tomcat": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:apache:tomcat",
        "description": "Apache Tomcat is an open-source implementation of the Java Servlet, JavaServer Pages, Java Expression Language and WebSocket technologies.",
        "headers": {
          "Server": "^Apache-Coyote",
          "X-Powered-By": "\\bTomcat\\b(?:-([\\d.]+))?\\;version:\\1"
        },
        "icon": "Apache Tomcat.svg",
        "implies": "Java",
        "website": "http://tomcat.apache.org"
      },
      "Apache Traffic Server": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:apache:traffic_server",
        "headers": {
          "Server": "ATS/?([\\d.]+)?\\;version:\\1"
        },
        "icon": "Apache Traffic Server.png",
        "website": "http://trafficserver.apache.org/"
      },
      "Apache Wicket": {
        "cats": [
          18
        ],
        "cpe": "cpe:/a:apache:wicket",
        "icon": "Apache Wicket.svg",
        "implies": "Java",
        "js": {
          "Wicket": ""
        },
        "website": "http://wicket.apache.org"
      },
      "ApexPages": {
        "cats": [
          51
        ],
        "headers": {
          "X-Powered-By": "Salesforce\\.com ApexPages"
        },
        "icon": "ApexPages.png",
        "implies": "Salesforce",
        "website": "https://developer.salesforce.com/docs/atlas.en-us.apexcode.meta/apexcode/apex_intro.htm"
      },
      "Apigee": {
        "cats": [
          4
        ],
        "description": "Apigee is an API gateway management tool to exchange data across cloud services and applications",
        "html": "<script>[^>]{0,50}script src=[^>]/profiles/apigee",
        "icon": "apigee.svg",
        "scripts": "/profiles/apigee",
        "website": "https://cloud.google.com/apigee/"
      },
      "Apisearch": {
        "cats": [
          29
        ],
        "description": "Apisearch is a real-time search platform for ecommerce.",
        "icon": "Apisearch.png",
        "oss": true,
        "pricing": [
          "freemium",
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "static\\.apisearch\\.cloud",
        "website": "https://apisearch.io"
      },
      "Aplazame": {
        "cats": [
          41
        ],
        "description": "Aplazame is a consumer credit company that provides instant financing service for online purchases. It combines an overtime payment method integrated at the ecommerce checkout with marketing tools to enable ecommerce to use financing as a promotional lever to boost sales.",
        "icon": "Aplazame.svg",
        "js": {
          "aplazame": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.aplazame\\.com/aplazame\\.js",
          "aplazame\\.com/static/aplazame\\.js"
        ],
        "website": "https://aplazame.com"
      },
      "Apollo": {
        "cats": [
          59
        ],
        "icon": "Apollo.svg",
        "implies": [
          "GraphQL",
          "TypeScript\\;confidence:50"
        ],
        "js": {
          "__APOLLO_CLIENT__": "",
          "__APOLLO_CLIENT__.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://www.apollographql.com"
      },
      "ApostropheCMS": {
        "cats": [
          1
        ],
        "html": "<[^>]+data-apos-refreshable[^>]",
        "icon": "ApostropheCMS.svg",
        "implies": "Node.js",
        "website": "http://apostrophecms.org"
      },
      "AppDynamics": {
        "cats": [
          10,
          78
        ],
        "description": "AppDynamics is an application performance management (APM) and IT operations analytics (ITOA) company based in San Francisco.",
        "icon": "AppDynamics.png",
        "js": {
          "ADRUM.conf.agentVer": "^(.+)$\\;version:\\1"
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "adrum",
        "website": "https://appdynamics.com"
      },
      "AppNexus": {
        "cats": [
          36
        ],
        "description": "AppNexus is a cloud-based software platform that enables and optimizes programmatic online advertising.",
        "dom": "iframe[src*='.adnxs.com'], img[src*='.adnxs.com']",
        "icon": "AppNexus.svg",
        "js": {
          "appnexus": "",
          "appnexusVideo": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "adnxs\\.(?:net|com)",
        "website": "http://appnexus.com",
        "xhr": "prebid\\.adnxs\\.com"
      },
      "Appian": {
        "cats": [
          62
        ],
        "description": "Appian is an enterprise low-code application platform.",
        "icon": "Appian.png",
        "saas": true,
        "scripts": "/suite/tempo/ui/sail-client/embeddedBootstrap.nocache.js",
        "website": "https://www.appian.com"
      },
      "Apple Pay": {
        "cats": [
          41
        ],
        "description": "Apple Pay is a mobile payment and digital wallet service by Apple that allows users to make payments in person, in iOS apps, and on the web.",
        "html": [
          "<[^>]+aria-labelledby=\"pi-apple_pay",
          "<script id=\"apple-pay"
        ],
        "icon": "Apple.svg",
        "js": {
          "ApplePay": "",
          "enableApplePay": ""
        },
        "website": "https://www.apple.com/apple-pay"
      },
      "Apple Sign-in": {
        "cats": [
          69
        ],
        "description": "Apple Sign-in is based on OAuth 2.0 and OpenID Connect, and provides a privacy-friendly way for users to sign in to websites and apps.",
        "dom": {
          "a[href*='appleid.apple.com/auth/authorize']": {
            "exists": ""
          },
          "button": {
            "text": "(Sign (in|up)|Log in|Continue) with Apple"
          }
        },
        "html": "<meta[ˆ>]*appleid-signin-client-id",
        "icon": "Apple.svg",
        "js": {
          "AppleID": ""
        },
        "scripts": "appleid\\.auth\\.js",
        "website": "https://developer.apple.com/sign-in-with-apple/"
      },
      "Appointy": {
        "cats": [
          72
        ],
        "description": "Appointy’s online scheduling software.",
        "html": "<iframe[^>]+src=\"?https://[\\w\\d\\-]+\\.appointy\\.com",
        "icon": "Appointy.png",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.appointy.com/"
      },
      "Apptus": {
        "cats": [
          76
        ],
        "cookies": {
          "apptus.customerKey": "",
          "apptus.sessionKey": ""
        },
        "description": "Apptus is an AI-powered ecommerce optimisation software provider.",
        "icon": "Apptus.svg",
        "js": {
          "ApptusEsales": "",
          "apptusConfig": "",
          "apptusDebug": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.esales\\.apptus\\.com.+(?:apptus-esales-api-([\\d.]+))\\.min\\.js\\;version:\\1",
        "website": "https://www.apptus.com"
      },
      "AquilaCMS": {
        "cats": [
          1
        ],
        "description": "AquilaCMS is a fullstack, headless CMS written in JavaScript.",
        "icon": "AquilaCMS.svg",
        "implies": [
          "Next.js",
          "Node.js",
          "React",
          "MongoDB",
          "Amazon Web Services"
        ],
        "meta": {
          "powered-by": "AquilaCMS"
        },
        "oss": true,
        "website": "https://www.aquila-cms.com/"
      },
      "Arastta": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:arastta:ecommerce",
        "excludes": "OpenCart",
        "headers": {
          "Arastta": "^(.+)$\\;version:\\1",
          "X-Arastta": ""
        },
        "html": "Powered by <a [^>]*href=\"https?://(?:www\\.)?arastta\\.org[^>]+>Arastta",
        "icon": "Arastta.svg",
        "implies": "PHP",
        "scripts": "arastta\\.js",
        "website": "http://arastta.org"
      },
      "Arc": {
        "cats": [
          31
        ],
        "description": "Arc is a peer-to-peer CDN that pays site owners for using it. Instead of expensive servers in distant datacenters, Arc's network is comprised of browsers.",
        "dom": "#arc-widget",
        "icon": "Arc.svg",
        "js": {
          "arc.p2pClient": "",
          "arcWidgetJsonp": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "arc\\.io/widget\\.js",
        "website": "https://arc.io",
        "xhr": "\\.arc\\.io"
      },
      "Arc Publishing": {
        "cats": [
          1
        ],
        "html": "<div [^>]*id=\"pb-root\"",
        "icon": "Arc-Publishing.svg",
        "js": {
          "Fusion.arcSite": ""
        },
        "website": "https://www.arcpublishing.com/"
      },
      "ArcGIS API for JavaScript": {
        "cats": [
          35
        ],
        "description": "ArcGIS API for JavaScript is a tool used to embed maps and tasks in web applications.",
        "icon": "arcgis_icon.png",
        "scripts": [
          "js\\.arcgis\\.com",
          "basemaps\\.arcgis\\.com"
        ],
        "website": "https://developers.arcgis.com/javascript/"
      },
      "Artifactory": {
        "cats": [
          47
        ],
        "cpe": "cpe:/a:jfrog:artifactory",
        "html": [
          "<span class=\"version\">Artifactory(?: Pro)?(?: Power Pack)?(?: ([\\d.]+))?\\;version:\\1"
        ],
        "icon": "Artifactory.svg",
        "js": {
          "ArtifactoryUpdates": ""
        },
        "scripts": [
          "wicket/resource/org\\.artifactory\\."
        ],
        "website": "http://jfrog.com/open-source/#os-arti"
      },
      "Artifactory Web Server": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:jfrog:artifactory",
        "headers": {
          "Server": "Artifactory(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Artifactory.svg",
        "implies": "Artifactory",
        "website": "http://jfrog.com/open-source/#os-arti"
      },
      "ArvanCloud": {
        "cats": [
          31
        ],
        "description": "ArvanCloud is a Cloud Services Provider, offering a wide range of incorporated cloud services including CDN, DDoS mitigation, Cloud Managed DNS, Cloud Security, VoD/AoD Streaming, Live Streaming, Cloud Compute, Cloud Object Storage, and PaaS.",
        "headers": {
          "AR-PoweredBy": "Arvan Cloud \\(arvancloud\\.com\\)"
        },
        "icon": "ArvanCloud.png",
        "js": {
          "ArvanCloud": ""
        },
        "website": "http://www.ArvanCloud.com"
      },
      "AsciiDoc": {
        "cats": [
          1,
          20,
          27
        ],
        "description": "AsciiDoc is a text document format for writing documentation, slideshows, web pages, man pages and blogs. AsciiDoc files can be translated to many formats including HTML, PDF, EPUB, man page.",
        "icon": "AsciiDoc.png",
        "js": {
          "asciidoc": ""
        },
        "meta": {
          "generator": "^AsciiDoc ([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.methods.co.nz/asciidoc"
      },
      "Asciinema": {
        "cats": [
          14
        ],
        "description": "Asciinema is a free and open-source solution for recording terminal sessions and sharing them on the web.",
        "html": "<asciinema-player",
        "icon": "asciinema.png",
        "js": {
          "asciinema": ""
        },
        "scripts": "asciinema\\.org/",
        "website": "https://asciinema.org/"
      },
      "Astute Solutions": {
        "cats": [
          53
        ],
        "description": "Astute Solutions is a customer engagement software.",
        "dom": "iframe[src*='.iperceptions.com'], link[href*='.iperceptions.com']",
        "icon": "Astute Solutions.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.iperceptions\\.com",
        "website": "https://astutesolutions.com"
      },
      "Athena Search": {
        "cats": [
          29
        ],
        "description": "Athena Search is a customizable autocomplete, feature-rich dashboard, smart predictions, real-time reports search engine developed from scratch by Syncit Group’s.",
        "dom": "link[href*='athena/autocomplete/css/autocomplete.css']",
        "icon": "Athena Search.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "wp-content/plugins/athena-search",
        "website": "https://www.athenasearch.io"
      },
      "Atlassian Bitbucket": {
        "cats": [
          47
        ],
        "cpe": "cpe:/a:atlassian:bitbucket",
        "description": "Bitbucket is a web-based version control repository hosting service for source code and development projects that use either Mercurial or Git revision control systems.",
        "html": "<li>Atlassian Bitbucket <span title=\"[a-z0-9]+\" id=\"product-version\" data-commitid=\"[a-z0-9]+\" data-system-build-number=\"[a-z0-9]+\"> v([\\d.]+)<\\;version:\\1",
        "icon": "Atlassian Bitbucket.svg",
        "implies": "Python",
        "js": {
          "bitbucket": ""
        },
        "meta": {
          "application-name": "Bitbucket"
        },
        "website": "http://www.atlassian.com/software/bitbucket/overview/"
      },
      "Atlassian Confluence": {
        "cats": [
          8
        ],
        "cpe": "cpe:/a:atlassian:confluence",
        "description": "Atlassian Confluence is a web-based collaboration wiki tool.",
        "dom": {
          "li.print-only": {
            "text": "Atlassian Confluence ([\\d.]+)\\;version:\\1"
          }
        },
        "headers": {
          "X-Confluence-Request-Time": ""
        },
        "icon": "Atlassian Confluence.svg",
        "implies": "Java",
        "meta": {
          "confluence-request-time": ""
        },
        "website": "http://www.atlassian.com/software/confluence/overview/team-collaboration-software"
      },
      "Atlassian FishEye": {
        "cats": [
          47
        ],
        "cookies": {
          "FESESSIONID": ""
        },
        "cpe": "cpe:/a:atlassian:fisheye",
        "html": "<title>(?:Log in to )?FishEye (?:and Crucible )?([\\d.]+)?</title>\\;version:\\1",
        "icon": "Atlassian FishEye.svg",
        "website": "http://www.atlassian.com/software/fisheye/overview/"
      },
      "Atlassian Jira": {
        "cats": [
          13
        ],
        "cpe": "cpe:/a:atlassian:jira",
        "dom": "#jira",
        "icon": "Atlassian Jira.svg",
        "implies": "Java",
        "js": {
          "jira.id": ""
        },
        "meta": {
          "application-name": "JIRA",
          "data-version": "([\\d.]+)\\;version:\\1\\;confidence:0"
        },
        "website": "http://www.atlassian.com/software/jira/overview/"
      },
      "Atlassian Jira Issue Collector": {
        "cats": [
          13,
          47
        ],
        "description": "Atlassian Jira Issue Collector is a tool used to download a list of websites using with email addresses, phone numbers and LinkedIn profiles.",
        "icon": "Atlassian Jira.svg",
        "scripts": [
          "jira-issue-collector-plugin",
          "atlassian\\.jira\\.collector\\.plugin"
        ],
        "website": "http://www.atlassian.com/software/jira/overview/"
      },
      "Atlassian Statuspage": {
        "cats": [
          13,
          62
        ],
        "description": "Statuspage is a status and incident communication tool.",
        "headers": {
          "X-StatusPage-Skip-Logging": "",
          "X-StatusPage-Version": ""
        },
        "html": "<a[^>]*href=\"https?://(?:www\\.)?statuspage\\.io/powered-by[^>]+>",
        "icon": "Atlassian Statuspage.svg",
        "website": "https://www.statuspage.io/"
      },
      "Attentive": {
        "cats": [
          76,
          5
        ],
        "description": "Attentive is a personalised mobile messaging platform that helps retail & ecommerce brands acquire, retain, and interact with mobile shoppers.",
        "icon": "Attentive.svg",
        "js": {
          "__attentive": "",
          "__attentive_domain": "",
          "attn_email_save": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.attn\\.tv",
        "website": "https://www.attentivemobile.com"
      },
      "AudioEye": {
        "cats": [
          68
        ],
        "description": "AudioEye automatically evaluates and adjusts website content to be accessible.",
        "html": "<iframe[^>]*audioeye\\.com/frame/cookieStorage",
        "icon": "AudioEye.svg",
        "scripts": "audioeye\\.com/ae\\.js",
        "website": "https://www.audioeye.com/"
      },
      "Aurelia": {
        "cats": [
          12
        ],
        "html": [
          "<[^>]+aurelia-app=[^>]",
          "<[^>]+data-main=[^>]aurelia-bootstrapper",
          "<[^>]+au-target-id=[^>]\\d"
        ],
        "icon": "Aurelia.svg",
        "scripts": [
          "aurelia(?:\\.min)?\\.js"
        ],
        "website": "http://aurelia.io"
      },
      "Auth0": {
        "cats": [
          19
        ],
        "description": "Auth0 provides authentication and authorisation as a service.",
        "icon": "Auth0.png",
        "scripts": [
          "/auth0(?:-js)?/([\\d.]+)/auth0(?:.min)?\\.js\\;version:\\1",
          "/auth0-js@([\\d.]+)/([a-z]+)/auth0\\.min\\.js\\;version:\\1"
        ],
        "website": "https://auth0.github.io/auth0.js/index.html"
      },
      "Auth0 Lock": {
        "cats": [
          19
        ],
        "description": "Auth0's signin solution",
        "icon": "Auth0.png",
        "scripts": "/lock/([\\d.]+)/lock(?:.min)?\\.js\\;version:\\1",
        "website": "https://auth0.com/docs/libraries/lock"
      },
      "Automattic": {
        "cats": [
          62
        ],
        "description": "Automattic is a freemium blogging service and an open-source blogging software.",
        "headers": {
          "X-Hacker": "(?:automattic\\.com/jobs|wpvip\\.com/careers)"
        },
        "icon": "automattic.png",
        "implies": "WordPress",
        "website": "https://automattic.com/"
      },
      "Autopilot": {
        "cats": [
          32,
          74,
          75
        ],
        "description": "Autopilot is a visual marketing software that enables users to create marketing campaigns and manage lead conversions. ",
        "icon": "Autopilot.png",
        "js": {
          "Autopilot": "\\;confidence:50",
          "AutopilotAnywhere": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.autopilothq.com"
      },
      "Avangate": {
        "cats": [
          6
        ],
        "html": "<link[^>]* href=\"https?://edge\\.avangate\\.net/",
        "icon": "Avangate.svg",
        "js": {
          "__avng8_": "",
          "avng8_": ""
        },
        "scripts": "^https?://edge\\.avangate\\.net/",
        "website": "http://avangate.com"
      },
      "Avasize": {
        "cats": [
          5
        ],
        "icon": "Avasize.png",
        "scripts": "^https?://cdn\\.avasize\\.com/",
        "website": "https://www.avasize.com"
      },
      "Aweber": {
        "cats": [
          32,
          75
        ],
        "description": "AWeber is an email marketing service.",
        "dom": "form[action*='aweber.com']",
        "icon": "Aweber.svg",
        "js": {
          "awt_analytics": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.aweber\\.com/",
        "website": "https://www.aweber.com"
      },
      "Awesomplete": {
        "cats": [
          29
        ],
        "description": "Awesomplete is a tool in the Javascript UI Libraries category of a tech stack.",
        "html": "<link[^>]+href=\"[^>]*awesomplete(?:\\.min)?\\.css",
        "js": {
          "awesomplete": ""
        },
        "scripts": "/awesomplete\\.js(?:$|\\?)",
        "website": "https://leaverou.github.io/awesomplete/"
      },
      "Axios": {
        "cats": [
          59
        ],
        "description": "Promise based HTTP client for the browser and node.js",
        "js": {
          "axios.get": ""
        },
        "scripts": [
          "/axios(@|/)([\\d.]+)(?:/[a-z]+)?/axios(?:.min)?\\.js\\;version:\\2"
        ],
        "website": "https://github.com/axios/axios"
      },
      "Azion": {
        "cats": [
          31
        ],
        "description": "",
        "dns": {
          "CNAME": "azioncdn\\.net\\.?$"
        },
        "headers": {
          "Server": "^Azion "
        },
        "icon": "Azion.svg",
        "pricing": [
          "poa",
          "payg"
        ],
        "saas": true,
        "website": "https://www.azion.com/"
      },
      "Azure": {
        "cats": [
          62
        ],
        "cookies": {
          "ARRAffinity": "",
          "TiPMix": ""
        },
        "description": "Azure is a cloud computing service for building, testing, deploying, and managing applications and services through Microsoft-managed data centers.",
        "dns": {
          "NS": [
            "\\.azure-dns\\."
          ],
          "SOA": [
            "azuredns-cloud\\.net"
          ]
        },
        "headers": {
          "azure-regionname": "",
          "azure-sitename": "",
          "azure-slotname": "",
          "azure-version": ""
        },
        "icon": "azure.svg",
        "website": "https://azure.microsoft.com"
      },
      "Azure CDN": {
        "cats": [
          31
        ],
        "description": "Azure Content Delivery Network (CDN) reduces load times, save bandwidth and speed responsiveness.",
        "headers": {
          "X-EC-Debug": "",
          "server": "^(?:ECAcc|ECS|ECD)"
        },
        "icon": "azure.svg",
        "website": "https://azure.microsoft.com/en-us/services/cdn/"
      },
      "BEM": {
        "cats": [
          12
        ],
        "html": "<[^>]+data-bem",
        "icon": "BEM.png",
        "website": "http://en.bem.info"
      },
      "BIGACE": {
        "cats": [
          1
        ],
        "html": "(?:Powered by <a href=\"[^>]+BIGACE|<!--\\s+Site is running BIGACE)",
        "icon": "BIGACE.png",
        "implies": "PHP",
        "meta": {
          "generator": "BIGACE ([\\d.]+)\\;version:\\1"
        },
        "website": "http://bigace.de"
      },
      "BOOM": {
        "cats": [
          1
        ],
        "headers": {
          "X-Supplied-By": "MANA"
        },
        "icon": "boom.svg",
        "implies": "WordPress",
        "meta": {
          "generator": "^boom site builder$"
        },
        "website": "http://manaandisheh.com"
      },
      "Babel": {
        "cats": [
          19
        ],
        "description": "Babel is a free and open-source transcompiler for writing next generation JavaScript.",
        "icon": "Babel.svg",
        "js": {
          "_babelPolyfill": ""
        },
        "website": "https://babeljs.io"
      },
      "Bablic": {
        "cats": [
          3,
          9
        ],
        "description": "Bablic is a localisation solution to translate your website.",
        "icon": "bablic.png",
        "js": {
          "bablic": ""
        },
        "website": "https://www.bablic.com/"
      },
      "Backbone.js": {
        "cats": [
          12
        ],
        "cpe": "cpe:/a:backbone_project:backbone",
        "description": "BackboneJS is a JavaScript library that allows to develop and structure the client side applications that run in a web browser.",
        "icon": "Backbone.js.png",
        "implies": "Underscore.js",
        "js": {
          "Backbone": "",
          "Backbone.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "backbone.*\\.js",
        "website": "http://backbonejs.org"
      },
      "Backdrop": {
        "cats": [
          1
        ],
        "excludes": "Drupal",
        "headers": {
          "X-Backdrop-Cache": "",
          "X-Generator": "^Backdrop CMS(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "icon": "Backdrop.png",
        "implies": "PHP",
        "js": {
          "Backdrop": ""
        },
        "meta": {
          "generator": "^Backdrop CMS(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "website": "https://backdropcms.org"
      },
      "Baidu Analytics (百度统计)": {
        "cats": [
          10
        ],
        "icon": "Baidu Tongji.png",
        "scripts": "hm\\.baidu\\.com/hm\\.js",
        "website": "https://tongji.baidu.com/"
      },
      "Banshee": {
        "cats": [
          1,
          18
        ],
        "html": "Built upon the <a href=\"[^>]+banshee-php\\.org/\">[a-z]+</a>(?:v([\\d.]+))?\\;version:\\1",
        "icon": "Banshee.png",
        "implies": "PHP",
        "meta": {
          "generator": "Banshee PHP"
        },
        "website": "http://www.banshee-php.org"
      },
      "Base": {
        "cats": [
          6
        ],
        "description": "Base is a hosted ecommerce platform that allows business owners to set up an online store and sell their products online.",
        "icon": "Base.svg",
        "js": {
          "Base.App.open_nav": ""
        },
        "meta": {
          "base-theme-name": "\\d+"
        },
        "scripts": "thebase\\.in/js",
        "website": "https://thebase.in"
      },
      "Basic": {
        "cats": [
          16
        ],
        "description": "Basic is an authetication method used by some web servers.",
        "headers": {
          "WWW-Authenticate": "^Basic"
        },
        "website": "https://tools.ietf.org/html/rfc7617"
      },
      "Bentobox": {
        "cats": [
          1
        ],
        "description": "Bentobox is a restaurant website platform that handles menus, reservations, gift cards and more.",
        "html": "<!-- Powered by BentoBox",
        "icon": "Bentobox.svg",
        "js": {
          "BentoAnalytics": ""
        },
        "website": "https://getbento.com"
      },
      "Beyable": {
        "cats": [
          76
        ],
        "cookies": {
          "beyable-cart": "",
          "beyable-cartd": ""
        },
        "description": "Beyable is a suite of tools that analyze website traffic to understand visitors' behaviors in real-time, through multi-channel in order to optimise conversion rate.",
        "icon": "Beyable.png",
        "js": {
          "BEYABLE": "",
          "beYableDomain": "",
          "beYableKey": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "msecnd\\.net/api/beYableJSv(\\d+)\\;version:\\1",
        "website": "https://beyable.com"
      },
      "Big Cartel": {
        "cats": [
          6
        ],
        "icon": "bigcartel.png",
        "meta": {
          "generator": "Big Cartel"
        },
        "website": "https://www.bigcartel.com"
      },
      "BigCommerce": {
        "cats": [
          6
        ],
        "description": "BigCommerce is a hosted ecommerce platform that allows business owners to set up an online store and sell their products online.",
        "dom": "img[data-src*='.bigcommerce.com'], img[src*='.bigcommerce.com'], link[href*='.bigcommerce.com']",
        "icon": "BigCommerce.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": ".bigcommerce\\.com",
        "url": "mybigcommerce\\.com",
        "website": "http://www.bigcommerce.com"
      },
      "BigDump": {
        "cats": [
          3
        ],
        "html": "<!-- <h1>BigDump: Staggered MySQL Dump Importer ver\\. ([\\d.b]+)\\;version:\\1",
        "implies": [
          "MySQL",
          "PHP"
        ],
        "website": "http://www.ozerov.de/bigdump.php"
      },
      "Bigware": {
        "cats": [
          6
        ],
        "cookies": {
          "bigWAdminID": "",
          "bigwareCsid": ""
        },
        "cpe": "cpe:/a:bigware:bigware_shop",
        "html": "(?:Diese <a href=[^>]+bigware\\.de|<a href=[^>]+/main_bigware_\\d+\\.php)",
        "icon": "Bigware.png",
        "implies": "PHP",
        "url": "(?:\\?|&)bigWAdminID=",
        "website": "http://bigware.de"
      },
      "Birdeye": {
        "cats": [
          32,
          5
        ],
        "description": "Birdeye is an all-in-one customer experience platform.",
        "icon": "Birdeye.svg",
        "js": {
          "bfiframe": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "birdeye\\.com/embed",
          "birdeye\\.com"
        ],
        "website": "https://birdeye.com"
      },
      "BittAds": {
        "cats": [
          36
        ],
        "icon": "BittAds.png",
        "js": {
          "bitt": ""
        },
        "scripts": "bittads\\.com/js/bitt\\.js$",
        "website": "http://bittads.com"
      },
      "Bizweb": {
        "cats": [
          6
        ],
        "icon": "bizweb.png",
        "js": {
          "Bizweb": ""
        },
        "website": "https://www.bizweb.vn"
      },
      "Blackbaud Luminate Online": {
        "cats": [
          41,
          32
        ],
        "description": "Blackbaud Luminate Online provides online fundraising and marketing automation for nonprofits.",
        "icon": "Blackbaud-Luminate-Online.png",
        "js": {
          "don_premium_map": ""
        },
        "scripts": "js/convio/modules\\.js",
        "url": "/site/Donation2?.*df_id=",
        "website": "https://www.blackbaud.com/products/blackbaud-luminate-online"
      },
      "Blade": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "X-Powered-By": "blade-([\\w.]+)?\\;version:\\1"
        },
        "icon": "Blade.png",
        "implies": "Java",
        "website": "https://lets-blade.com"
      },
      "Blazor": {
        "cats": [
          18
        ],
        "icon": "Blazor.png",
        "implies": "Microsoft ASP.NET",
        "scripts": [
          "blazor\\.server\\.js",
          "blazor\\.host\\.min\\.js",
          "blazor\\.webassembly\\.js"
        ],
        "website": "https://dotnet.microsoft.com/apps/aspnet/web-apps/blazor"
      },
      "Blessing Skin": {
        "cats": [
          7
        ],
        "description": "Blessing Skin is a plubin that brings your custom skins back in offline Minecraft servers.",
        "icon": "Blessing Skin.png",
        "implies": "Laravel",
        "js": {
          "blessing.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://github.com/bs-community/blessing-skin-server"
      },
      "Blesta": {
        "cats": [
          6
        ],
        "cookies": {
          "blesta_sid": ""
        },
        "icon": "Blesta.png",
        "website": "http://www.blesta.com"
      },
      "Blip.tv": {
        "cats": [
          14
        ],
        "description": "Blip.tv is a media platform for web series content and also a dashboard for producers of original web series to distribute and monetise their productions.",
        "html": "<(?:param|embed|iframe)[^>]+blip\\.tv/play",
        "icon": "Blip.tv.png",
        "website": "http://blip.tv"
      },
      "Blogger": {
        "cats": [
          11
        ],
        "description": "Blogger is a blog-publishing service that allows multi-user blogs with time-stamped entries.",
        "icon": "Blogger.png",
        "implies": "Python",
        "meta": {
          "generator": "^Blogger$"
        },
        "url": "^https?://[^/]+\\.(?:blogspot|blogger)\\.com",
        "website": "http://www.blogger.com"
      },
      "Bloomreach": {
        "cats": [
          1
        ],
        "html": "<[^>]+/binaries/(?:[^/]+/)*content/gallery/",
        "icon": "Bloomreach.png",
        "website": "https://developers.bloomreach.com"
      },
      "Bloomreach Search & Merchandising": {
        "cats": [
          29,
          10,
          32,
          74
        ],
        "description": "Bloomreach Search & Merchandising (brSM) is an AI-driven platform for ecommerce.",
        "icon": "Bloomreach.png",
        "js": {
          "BrTrk": "",
          "br_data": ""
        },
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "cdns\\.brsrvr\\.com/v([\\d.]+)\\;version:\\1",
          "cdn\\.brcdn\\.com/v([\\d.]+)\\;version:\\1"
        ],
        "website": "https://developers.bloomreach.com/products/personalization/dev-studio"
      },
      "Blue": {
        "cats": [
          77
        ],
        "description": "Blue is a ecommerce data marketing, lead generation, real time bidding and recommendation solutions.",
        "dom": "iframe[src*='.getblue.io']",
        "icon": "Blue.svg",
        "js": {
          "blueProductId": "",
          "bluecpy_id": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.getblue\\.io",
        "website": "https://web.getblue.io/en/"
      },
      "Blue Triangle": {
        "cats": [
          16,
          78
        ],
        "description": "Blue Triangle is a connected view of marketing, web performance, and third-party tag analytics while constantly monitoring website code for security vulnerabilities.",
        "icon": "Blue Triangle.png",
        "js": {
          "_bttUtil.version": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "poa",
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.btttag\\.com/btt\\.js",
        "website": "https://bluetriangle.com",
        "xhr": "\\.btttag\\.com"
      },
      "Bluecore": {
        "cats": [
          32,
          75
        ],
        "description": "Bluecore is a retail marketing technology that uses data gained from direct marketing like email, social media, site activity.",
        "icon": "Bluecore.svg",
        "js": {
          "_bluecoreTrack": "",
          "bluecore_action_trigger": "",
          "triggermail": "",
          "triggermail_email_address": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.bluecore\\.com",
        "website": "https://www.bluecore.com"
      },
      "Bluefish": {
        "cats": [
          20
        ],
        "description": "Bluefish is a free software text editor with a variety of tools for programming in general and the development of websites.",
        "icon": "Bluefish.png",
        "meta": {
          "generator": "Bluefish(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "website": "http://sourceforge.net/projects/bluefish"
      },
      "Blueknow": {
        "cats": [
          76
        ],
        "description": "Blueknow is a ecommerce personalisation software designed to serve enterprises, SMEs.",
        "icon": "Blueknow.png",
        "js": {
          "Blueknow": "",
          "BlueknowTracker": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.blueknow\\.com",
        "website": "https://www.blueknow.com"
      },
      "Boba.js": {
        "cats": [
          59
        ],
        "implies": "Google Analytics",
        "scripts": "boba(?:\\.min)?\\.js",
        "website": "http://boba.space150.com"
      },
      "Bokeh": {
        "cats": [
          25
        ],
        "icon": "bokeh.png",
        "implies": "Python",
        "js": {
          "Bokeh": "",
          "Bokeh.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "bokeh.*\\.js",
        "website": "https://bokeh.org"
      },
      "Bold Chat": {
        "cats": [
          52
        ],
        "description": "BoldChat is a live chat platform.",
        "icon": "BoldChat.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "^https?://vmss\\.boldchat\\.com/aid/\\d{18}/bc\\.vms4/vms\\.js",
        "website": "https://www.boldchat.com/"
      },
      "BoldGrid": {
        "cats": [
          1,
          11
        ],
        "description": "BoldGrid is a free website builder for WordPress websites.",
        "html": [
          "<link rel=[\"']stylesheet[\"'] [^>]+boldgrid",
          "<link rel=[\"']stylesheet[\"'] [^>]+post-and-page-builder",
          "<link[^>]+s\\d+\\.boldgrid\\.com"
        ],
        "implies": "WordPress",
        "scripts": "/wp-content/plugins/post-and-page-builder",
        "website": "https://boldgrid.com"
      },
      "Bolt CMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:bolt:bolt",
        "icon": "Bolt CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "Bolt"
        },
        "website": "http://bolt.cm"
      },
      "Bolt Payments": {
        "cats": [
          41
        ],
        "description": "Bolt powers a checkout experience designed to convert shoppers.",
        "dom": "bolt-checkout-button",
        "icon": "Bolt.svg",
        "js": {
          "BoltCheckout": "",
          "bolt_callbacks": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.bolt.com/",
        "xhr": "connect\\.bolt\\.com"
      },
      "Bonfire": {
        "cats": [
          18
        ],
        "cookies": {
          "bf_session": ""
        },
        "html": "Powered by <a[^>]+href=\"https?://(?:www\\.)?cibonfire\\.com[^>]*>Bonfire v([^<]+)\\;version:\\1",
        "icon": "Bonfire.png",
        "implies": "CodeIgniter",
        "website": "http://cibonfire.com"
      },
      "BookDinners": {
        "cats": [
          5,
          72
        ],
        "description": "BookDinners is a restaurant table booking widget.",
        "icon": "BookDinners.svg",
        "scripts": "bookdinners\\.nl/widget\\.js",
        "website": "https://www.bookdinners.nl"
      },
      "Bookatable": {
        "cats": [
          5,
          72
        ],
        "description": "Bookatable is a restaurant table booking widget.",
        "icon": "Bookatable.svg",
        "scripts": "bda\\.bookatable\\.com/deploy/lbui\\.direct\\.min\\.js",
        "website": "https://www.bookatable.co.uk"
      },
      "Bookero": {
        "cats": [
          72
        ],
        "description": "Bookero is online booking system for you website or Facebook page.",
        "icon": "Bookero.svg",
        "js": {
          "bookero_config": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "scripts": "cdn\\.bookero\\.pl",
        "url": "\\.bookero\\.(?:org|pl)",
        "website": "https://www.bookero.org"
      },
      "Bookingkit": {
        "cats": [
          5,
          72
        ],
        "description": "Bookingkit is an online booking management solution. Bookingkit helps its users generate PDF invoices, manage day-to-day scheduling operations, and automatically sync availabilities in real time.",
        "icon": "Bookingkit.svg",
        "js": {
          "BookingKitApp": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://bookingkit.net/"
      },
      "Booksy": {
        "cats": [
          5,
          72
        ],
        "description": "Booksy is a booking system for people looking to schedule appointments for health and beauty services.",
        "icon": "Booksy.svg",
        "js": {
          "booksy": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "scripts": "booksy\\.com/widget/code\\.js",
        "website": "https://booksy.com/"
      },
      "Boomerang": {
        "cats": [
          59,
          78
        ],
        "description": "boomerang is a JavaScript library that measures the page load time experienced by real users, commonly called RUM (Real User Measurement).",
        "icon": "boomerang.svg",
        "js": {
          "BOOMR": "",
          "BOOMR_lstart": "",
          "BOOMR_mq": ""
        },
        "oss": true,
        "website": "https://akamai.github.io/boomerang"
      },
      "Bootstrap": {
        "cats": [
          66
        ],
        "cpe": "cpe:/a:getbootstrap:bootstrap",
        "description": "Bootstrap is a free and open-source CSS framework directed at responsive, mobile-first front-end web development. It contains CSS and JavaScript-based design templates for typography, forms, buttons, navigation, and other interface components.",
        "html": [
          "<style>\\s+/\\*!\\s+\\* Bootstrap v(\\d\\.\\d\\.\\d)\\;version:\\1",
          "<link[^>]* href=[^>]*?bootstrap(?:[^>]*?([0-9a-fA-F]{7,40}|[\\d]+(?:.[\\d]+(?:.[\\d]+)?)?)|)[^>]*?(?:\\.min)?\\.css\\;version:\\1"
        ],
        "icon": "Bootstrap.svg",
        "implies": [
          "jQuery\\;confidence:50"
        ],
        "js": {
          "bootstrap.Alert.VERSION": "^(.+)$\\;version:\\1",
          "jQuery.fn.tooltip.Constructor.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "bootstrap(?:[^>]*?([0-9a-fA-F]{7,40}|[\\d]+(?:.[\\d]+(?:.[\\d]+)?)?)|)[^>]*?(?:\\.min)?\\.js\\;version:\\1"
        ],
        "website": "https://getbootstrap.com"
      },
      "Bootstrap Table": {
        "cats": [
          59
        ],
        "html": "<link[^>]+href=\"[^>]*bootstrap-table(?:\\.min)?\\.css",
        "icon": "Bootstrap Table.svg",
        "implies": [
          "Bootstrap",
          "jQuery"
        ],
        "scripts": "bootstrap-table(?:\\.min)?\\.js",
        "website": "http://bootstrap-table.wenzhixin.net.cn/"
      },
      "Booxi": {
        "cats": [
          72
        ],
        "description": "Booxi is a cloud-based appointment management platform for small to midsize businesses.",
        "icon": "Booxi.svg",
        "js": {
          "booxi": "",
          "booxiController": "",
          "bxe_core": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "/bxe_core\\.js",
        "website": "https://www.booxi.com"
      },
      "Borderfree": {
        "cats": [
          6
        ],
        "cookies": {
          "bfx.apiKey:": "^[\\w\\d-]+$\\;confidence:25",
          "bfx.country:": "^\\w+$\\;confidence:25",
          "bfx.language": "^\\w+$\\;confidence:25",
          "bfx.logLevel": "^\\w+$\\;confidence:25"
        },
        "description": "Borderfree is an cross-border ecommerce solutions provider.",
        "icon": "Borderfree.png",
        "js": {
          "bfx._apiKey": "\\;confidence:50",
          "bfx._brand": "\\;confidence:50"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "global\\.prd\\.borderfree\\.com",
          "wm\\.prd\\.borderfree\\.com",
          "bfx\\.js"
        ],
        "website": "https://www.borderfree.com"
      },
      "Borlabs Cookie": {
        "cats": [
          67
        ],
        "description": "Borlabs Cookie is a GDPR cookie consent plugin for WordPress.",
        "dom": "#BorlabsCookieBox",
        "icon": "Borlabs Cookie.svg",
        "implies": "WordPress",
        "js": {
          "borlabsCookieConfig": ""
        },
        "pricing": [
          "low",
          "onetime"
        ],
        "website": "https://borlabs.io/borlabs-cookie/"
      },
      "Botble CMS": {
        "cats": [
          1,
          6
        ],
        "cookies": {
          "botble_session": ""
        },
        "headers": {
          "CMS-Version": "^(.+)$\\;version:\\1\\;confidence:0"
        },
        "icon": "Botble-CMS.png",
        "implies": "Laravel",
        "website": "https://botble.com"
      },
      "BrainSINS": {
        "cats": [
          76
        ],
        "description": "BrainSINS is a personalisation technology and ecommerce analytics services to online retailers.",
        "icon": "BrainSINS.png",
        "js": {
          "BrainSINS": "",
          "BrainSINSRecommender": "",
          "brainsins_token": "",
          "launchBrainSINS": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "mw\\.brainsins\\.com",
          "cloudfront\\.net/brainsins(?:_v)?(\\d+)\\.js\\;version:\\1"
        ],
        "website": "http://brainsins.com"
      },
      "Braintree": {
        "cats": [
          41
        ],
        "description": "Braintree, a division of PayPal, specializes in mobile and web payment systems for ecommerce companies. Braintree provides clients with a merchant account and a payment gateway.",
        "icon": "Braintree.svg",
        "js": {
          "Braintree": "",
          "Braintree.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "js\\.braintreegateway\\.com",
        "website": "https://www.braintreepayments.com"
      },
      "Branch": {
        "cats": [
          32,
          10
        ],
        "description": "Branch is a mobile deep linking system to increase engagement and retention.",
        "icon": "Branch.png",
        "js": {
          "branch.setBranchViewData": "",
          "branch_callback__0": ""
        },
        "pricing": [
          "poa",
          "freemium"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.branch\\.io",
          "app\\.link/_r\\?sdk=web([\\d.]+)\\;version:\\1"
        ],
        "website": "https://branch.io"
      },
      "Braze": {
        "cats": [
          32,
          10
        ],
        "description": "Braze is a customer engagement platform that delivers messaging experiences across push, email, in-product, and more.",
        "icon": "Braze.svg",
        "js": {
          "appboy": "",
          "appboyQueue": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "js\\.appboycdn\\.com/web-sdk/([\\d.]+)\\;version:\\1",
        "website": "https://www.braze.com"
      },
      "Bread": {
        "cats": [
          41
        ],
        "description": "Bread is a buy now, pay later platform for ecommerce websites.",
        "dom": "#bread-mini-cart-btn",
        "icon": "Bread.svg",
        "js": {
          "BreadCalc": "",
          "BreadError": "",
          "BreadLoaded": "",
          "BreadShopify": "",
          "bread.appHost": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.getbread\\.com",
        "website": "https://www.breadpayments.com"
      },
      "BrightInfo": {
        "cats": [
          32,
          74
        ],
        "description": "BrightInfo is an automated content personalisation solution.",
        "icon": "BrightInfo.png",
        "js": {
          "_BI_": "\\;confidence:50",
          "_biq": "\\;confidence:50",
          "biJsUrl": "//app\\.brightinfo\\.com"
        },
        "pricing": [
          "poa",
          "recurring"
        ],
        "saas": true,
        "scripts": "app\\.brightinfo\\.com",
        "website": "https://www.brightinfo.com"
      },
      "Brightspot": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "^Brightspot$"
        },
        "icon": "Brightspot.svg",
        "implies": "Java",
        "website": "https://www.brightspot.com"
      },
      "Broadstreet": {
        "cats": [
          36
        ],
        "description": "Broadstreet is an ad manager that caters specifically to the needs of direct, digital ad sales.",
        "icon": "Broadstreet.png",
        "js": {
          "broadstreet": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.broadstreetads\\.com",
        "website": "https://broadstreetads.com"
      },
      "Bronto": {
        "cats": [
          32,
          75
        ],
        "description": "Bronto is a cloud-based email marketing automation software.",
        "icon": "Bronto.svg",
        "js": {
          "BrontoShopify": "",
          "bronto.versions.sca": "(.+)\\;version:\\1",
          "brontoCookieConsent": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "(?:snip|cdn)\\.bronto\\.com",
        "website": "https://bronto.com"
      },
      "Brownie": {
        "cats": [
          1,
          6
        ],
        "description": "Brownie is a framework, CMS, ecommerce and ERP omni-channel platform to manage your entire business in one cloud solution.",
        "dom": "a[href*='browniesuite.com'][target='_blank'] img[src*='brownie']",
        "headers": {
          "X-Powered-By": "Brownie"
        },
        "icon": "Brownie.svg",
        "implies": [
          "PHP",
          "MySQL",
          "Amazon Web Services",
          "Bootstrap",
          "jQuery"
        ],
        "pricing": [
          "mid",
          "recurring",
          "poa",
          "payg"
        ],
        "saas": true,
        "scripts": "assets\\.youthsrl\\.com/brownie",
        "website": "https://www.browniesuite.com"
      },
      "BrowserCMS": {
        "cats": [
          1
        ],
        "icon": "BrowserCMS.png",
        "implies": "Ruby",
        "meta": {
          "generator": "BrowserCMS ([\\d.]+)\\;version:\\1"
        },
        "website": "http://browsercms.org"
      },
      "Bsale": {
        "cats": [
          6
        ],
        "cookies": {
          "_bsalemarket_session": ""
        },
        "description": "Bsale is an store management solution for retail businesses that sell both in store and online.",
        "icon": "Bsale.png",
        "implies": "Nginx",
        "js": {
          "Bsale.version": "([\\d.]+)\\;version:\\1"
        },
        "meta": {
          "autor": "Bsale",
          "generator": "Bsale"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.bsale.cl"
      },
      "Bubble": {
        "cats": [
          51,
          18
        ],
        "description": "Bubble is a no-code platform that lets anyone build web apps without writing any code.",
        "headers": {
          "x-bubble-capacity-limit": "",
          "x-bubble-capacity-used": "",
          "x-bubble-perf": ""
        },
        "icon": "bubble.png",
        "implies": "Node.js",
        "js": {
          "_bubble_page_load_data": "",
          "bubble_environment": "",
          "bubble_hostname_modifier": "",
          "bubble_version": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "website": "http://bubble.io"
      },
      "BugSnag": {
        "cats": [
          10
        ],
        "description": "Bugsnag is a cross-platform error monitoring, reporting, and resolution software.",
        "icon": "BugSnag.png",
        "js": {
          "Bugsnag": "",
          "bugsnag": "",
          "bugsnagClient": ""
        },
        "scripts": "/bugsnag.*\\.js",
        "website": "http://bugsnag.com"
      },
      "Bugcrowd": {
        "cats": [
          16
        ],
        "description": "Bugcrowd is a crowdsourced cybersecurity platform.",
        "dns": {
          "TXT": [
            "bugcrowd-verification"
          ]
        },
        "icon": "Bugcrowd.svg",
        "website": "https://www.bugcrowd.com"
      },
      "Bugzilla": {
        "cats": [
          13
        ],
        "cookies": {
          "Bugzilla_login_request_cookie": ""
        },
        "cpe": "cpe:/a:mozilla:bugzilla",
        "html": [
          "href=\"enter_bug\\.cgi\">",
          "<main id=\"bugzilla-body\"",
          "<a href=\"https?://www\\.bugzilla\\.org/docs/([0-9.]+)/[^>]+>Help<\\;version:\\1",
          "<span id=\"information\" class=\"header_addl_info\">version ([\\d.]+)<\\;version:\\1"
        ],
        "icon": "Bugzilla.png",
        "implies": "Perl",
        "js": {
          "BUGZILLA": ""
        },
        "meta": {
          "generator": "Bugzilla ?([\\d.]+)?\\;version:\\1"
        },
        "website": "http://www.bugzilla.org"
      },
      "Builder": {
        "cats": [
          1
        ],
        "description": "Builder.io is a headless CMS with a powerful drag-and-drop visual editor that lets you build and optimize digital experiences with speed and flexibility. ",
        "dom": "[data-builder-content-id], img[src*='cdn.builder.io']",
        "icon": "Builder.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "website": "https://builder.io",
        "xhr": "cdn\\.builder\\.io"
      },
      "Buildertrend": {
        "cats": [
          19
        ],
        "dom": "iframe[src*='buildertrend.net'], script[src*='buildertrend.net']",
        "icon": "Buildertrend.svg",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://buildertrend.com"
      },
      "Bulma": {
        "cats": [
          66
        ],
        "html": "<link[^>]+?href=\"[^\"]+bulma(?:\\.min)?\\.css",
        "icon": "Bulma.png",
        "website": "http://bulma.io"
      },
      "Bump": {
        "cats": [
          4
        ],
        "description": "Bump is an API contract management platform that helps document and track APIs by identifying changes in API structure, and keeping developers informed through an elegant documentation.",
        "dom": {
          ".doc-navigation footer, footer.catalog-footer": {
            "text": "Powered by Bump"
          }
        },
        "icon": "Bump.svg",
        "pricing": [
          "recurring",
          "freemium",
          "low"
        ],
        "saas": true,
        "website": "https://bump.sh"
      },
      "Business Catalyst": {
        "cats": [
          1
        ],
        "html": "<!-- BC_OBNW -->",
        "icon": "Business Catalyst.svg",
        "scripts": "CatalystScripts",
        "website": "http://businesscatalyst.com"
      },
      "Buy me a coffee": {
        "cats": [
          5,
          41
        ],
        "description": "Buy me a coffee is a service for online content creators that they may use to receive tips and donations to support their work.",
        "dom": "a[href*='www.buymeacoffee.com/'][target='_blank']",
        "icon": "Buy me a coffee.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "cdnjs\\.buymeacoffee\\.com/([\\d.]+)\\;version:\\1",
        "website": "https://www.buymeacoffee.com"
      },
      "BuySellAds": {
        "cats": [
          36
        ],
        "icon": "BuySellAds.svg",
        "js": {
          "_bsa": "",
          "_bsaPRO": "",
          "_bsap": "",
          "_bsap_serving_callback": ""
        },
        "scripts": [
          "^https?://s\\d\\.buysellads\\.com/",
          "servedby-buysellads\\.com/monetization(?:\\.[\\w\\d]+)?\\.js"
        ],
        "website": "http://buysellads.com"
      },
      "CCV Shop": {
        "cats": [
          6
        ],
        "icon": "ccvshop.png",
        "scripts": "/website/JavaScript/Vertoshop\\.js",
        "website": "https://ccvshop.be"
      },
      "CDN77": {
        "cats": [
          31
        ],
        "description": "CDN77 is a content delivery network (CDN).",
        "headers": {
          "Server": "^CDN77-Turbo$"
        },
        "icon": "CDN77.png",
        "website": "https://www.cdn77.com"
      },
      "CFML": {
        "cats": [
          27
        ],
        "description": "ColdFusion Markup Language (CFML), is a scripting language for web development that runs on the JVM, the .NET framework, and Google App Engine.",
        "icon": "CFML.png",
        "website": "http://adobe.com/products/coldfusion-family.html"
      },
      "CIVIC": {
        "cats": [
          67
        ],
        "description": "Civic provides cookie control for user consent and the use of cookies.",
        "icon": "civic.png",
        "scripts": "cc\\.cdn\\.civiccomputing\\.com",
        "website": "https://www.civicuk.com/cookie-control"
      },
      "CKEditor": {
        "cats": [
          24
        ],
        "cpe": "cpe:/a:ckeditor:ckeditor",
        "description": "CKEditor is a WYSIWYG rich text editor which enables writing content directly inside of web pages or online applications. Its core code is written in JavaScript and it is developed by CKSource. CKEditor is available under open-source and commercial licenses.",
        "icon": "CKEditor.png",
        "js": {
          "CKEDITOR": "",
          "CKEDITOR.version": "^(.+)$\\;version:\\1",
          "CKEDITOR_BASEPATH": ""
        },
        "website": "http://ckeditor.com"
      },
      "CMS Made Simple": {
        "cats": [
          1
        ],
        "cookies": {
          "CMSSESSID": ""
        },
        "cpe": "cpe:/a:cmsmadesimple:cms_made_simple",
        "icon": "CMS Made Simple.png",
        "implies": "PHP",
        "meta": {
          "generator": "CMS Made Simple"
        },
        "website": "http://cmsmadesimple.org"
      },
      "CMSimple": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:cmsimple:cmsimple",
        "implies": "PHP",
        "meta": {
          "generator": "CMSimple( [\\d.]+)?\\;version:\\1"
        },
        "website": "http://www.cmsimple.org/en"
      },
      "CNZZ": {
        "cats": [
          10
        ],
        "icon": "cnzz.png",
        "js": {
          "cnzz_protocol": ""
        },
        "scripts": "//[^./]+\\.cnzz\\.com/(?:z_stat.php|core)\\?",
        "website": "https://web.umeng.com/"
      },
      "CPG Dragonfly": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "^Dragonfly CMS"
        },
        "icon": "CPG Dragonfly.png",
        "implies": "PHP",
        "meta": {
          "generator": "CPG Dragonfly"
        },
        "website": "http://dragonflycms.org"
      },
      "CS Cart": {
        "cats": [
          6
        ],
        "html": [
          "&nbsp;Powered by (?:<a href=[^>]+cs-cart\\.com|CS-Cart)",
          "\\.cm-noscript[^>]+</style>"
        ],
        "icon": "CS Cart.png",
        "implies": "PHP",
        "js": {
          "fn_compare_strings": ""
        },
        "pricing": [
          "mid",
          "onetime"
        ],
        "website": "http://www.cs-cart.com"
      },
      "CacheFly": {
        "cats": [
          31
        ],
        "description": "CacheFly is a content delivery network (CDN) which offers CDN service that relies solely on IP anycast for routing, rather than DNS based global load balancing.",
        "headers": {
          "Server": "^CFS ",
          "X-CF1": "",
          "X-CF2": ""
        },
        "icon": "CacheFly.svg",
        "website": "http://www.cachefly.com"
      },
      "Caddy": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^Caddy$"
        },
        "icon": "caddy.svg",
        "implies": "Go",
        "website": "http://caddyserver.com"
      },
      "Cafe24": {
        "cats": [
          6
        ],
        "icon": "Cafe24.png",
        "js": {
          "EC_GLOBAL_DATETIME": "",
          "EC_GLOBAL_INFO": "",
          "EC_ROOT_DOMAIN": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "https://ec.cafe24.com/"
      },
      "CakePHP": {
        "cats": [
          18
        ],
        "cookies": {
          "cakephp": ""
        },
        "cpe": "cpe:/a:cakephp:cakephp",
        "icon": "CakePHP.png",
        "implies": "PHP",
        "meta": {
          "application-name": "CakePHP"
        },
        "website": "http://cakephp.org"
      },
      "Calendly": {
        "cats": [
          72
        ],
        "description": "Calendly is an app for scheduling appointments, meetings, and events.",
        "icon": "Calendly.svg",
        "js": {
          "Calendly": ""
        },
        "pricing": [
          "low",
          "freemium"
        ],
        "saas": true,
        "scripts": "https://assets\\.calendly\\.com/assets/external/widget\\.js",
        "website": "https://calendly.com/"
      },
      "Captch Me": {
        "cats": [
          16,
          36
        ],
        "icon": "Captch Me.svg",
        "js": {
          "Captchme": ""
        },
        "scripts": "^https?://api\\.captchme\\.net/",
        "website": "http://captchme.com"
      },
      "Carbon Ads": {
        "cats": [
          36
        ],
        "html": "<[a-z]+ [^>]*id=\"carbonads-container\"",
        "icon": "Carbon Ads.png",
        "js": {
          "_carbonads": ""
        },
        "scripts": "carbonads\\.com",
        "website": "http://carbonads.net"
      },
      "Cargo": {
        "cats": [
          1
        ],
        "html": "<link [^>]+Cargo feed",
        "icon": "Cargo.svg",
        "implies": "PHP",
        "meta": {
          "cargo_title": ""
        },
        "scripts": "/cargo\\.",
        "website": "http://cargocollective.com"
      },
      "Cart Functionality": {
        "cats": [
          6
        ],
        "description": "Websites that have a shopping cart or checkout page, either using a known ecommerce platform or a custom solution.",
        "dom": "a[href*='/cart'], a[href*='/basket'], a[href*='/trolley'], a[href*='/bag'], a[href*='/shoppingbag'], a[href*='/checkout'], [aria-controls='cart']",
        "icon": "Cart-generic.svg",
        "website": "https://www.wappalyzer.com/technologies/ecommerce/cart-functionality"
      },
      "CartStack": {
        "cats": [
          6,
          10
        ],
        "description": "CartStack is a SaaS solution that allows any company with an ecommerce site or reservation system to increase revenue through reminding/encouraging consumers to return to their abandoned cart and complete their purchase.",
        "icon": "CartStack.svg",
        "js": {
          "_cartstack": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "api\\.cartstack\\.\\w+",
        "website": "https://www.cartstack.com"
      },
      "Catberry.js": {
        "cats": [
          12,
          18
        ],
        "headers": {
          "X-Powered-By": "Catberry"
        },
        "icon": "Catberry.js.png",
        "implies": "Node.js",
        "js": {
          "catberry": "",
          "catberry.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://catberry.github.io/"
      },
      "Cecil": {
        "cats": [
          57
        ],
        "description": "Cecil is a CLI application, powered by PHP, that merge plain text files (written in Markdown), images and Twig templates to generate a static website.",
        "icon": "Cecil.svg",
        "meta": {
          "generator": "^Cecil(?: ([0-9.]+))?$\\;version:\\1"
        },
        "website": "https://cecil.app"
      },
      "CentOS": {
        "cats": [
          28
        ],
        "cpe": "cpe:/o:centos:centos",
        "description": "CentOS is a Linux distribution that provides a free, community-supported computing platform functionally compatible with its upstream source, Red Hat Enterprise Linux (RHEL).",
        "headers": {
          "Server": "CentOS",
          "X-Powered-By": "CentOS"
        },
        "icon": "CentOS.png",
        "website": "http://centos.org"
      },
      "Centminmod": {
        "cats": [
          22
        ],
        "headers": {
          "X-Powered-By": "centminmod"
        },
        "icon": "centminmod.png",
        "implies": [
          "CentOS",
          "Nginx",
          "PHP"
        ],
        "website": "https://centminmod.com"
      },
      "Ceres": {
        "cats": [
          6
        ],
        "headers": {
          "X-Plenty-Shop": "Ceres"
        },
        "icon": "Ceres.svg",
        "website": "https://www.plentymarkets.com/"
      },
      "Chamilo": {
        "cats": [
          21
        ],
        "cpe": "cpe:/a:chamilo:chamilo_lms",
        "description": "Chamilo is an open-source learning management and collaboration system.",
        "headers": {
          "X-Powered-By": "Chamilo ([\\d.]+)\\;version:\\1"
        },
        "html": "\">Chamilo ([\\d.]+)</a>\\;version:\\1",
        "icon": "Chamilo.png",
        "implies": "PHP",
        "meta": {
          "generator": "Chamilo ([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.chamilo.org"
      },
      "Chaport": {
        "cats": [
          52
        ],
        "description": "Chaport is a multi-channel live chat and chatbot software for business.",
        "icon": "Chaport.png",
        "js": {
          "chaport": "",
          "chaportConfig": ""
        },
        "pricing": [
          "freemium",
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.chaport\\.com",
        "website": "https://www.chaport.com"
      },
      "Chargebee": {
        "cats": [
          41
        ],
        "description": "Chargebee is a PCI Level 1 certified recurring billing platform for SaaS and subscription-based businesses.",
        "icon": "Chargebee.svg",
        "js": {
          "Chargebee": "",
          "chargebeeTrackFunc": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "js\\.chargebee\\.com/v([\\d.]+)\\;version:\\1",
        "website": "https://www.chargebee.com"
      },
      "Chart.js": {
        "cats": [
          25
        ],
        "description": "Chart.js is an open-source JavaScript library that allows you to draw different types of charts by using the HTML5 canvas element.",
        "icon": "Chart.js.svg",
        "js": {
          "Chart": "\\;confidence:50",
          "Chart.defaults.doughnut": "",
          "chart.ctx.bezierCurveTo": ""
        },
        "scripts": [
          "/Chart(?:\\.bundle)?(?:\\.min)?\\.js\\;confidence:75",
          "chartjs\\.org/dist/([\\d.]+(?:-[^/]+)?|master|latest)/Chart.*\\.js\\;version:\\1",
          "cdnjs\\.cloudflare\\.com/ajax/libs/Chart\\.js/([\\d.]+(?:-[^/]+)?)/Chart.*\\.js\\;version:\\1",
          "cdn\\.jsdelivr\\.net/(?:npm|gh/chartjs)/chart\\.js@([\\d.]+(?:-[^/]+)?|latest)/dist/Chart.*\\.js\\;version:\\1"
        ],
        "website": "https://www.chartjs.org"
      },
      "Chartbeat": {
        "cats": [
          10
        ],
        "icon": "Chartbeat.png",
        "js": {
          "_sf_async_config": "",
          "_sf_endpt": ""
        },
        "scripts": "chartbeat\\.js",
        "website": "http://chartbeat.com"
      },
      "Chatango": {
        "cats": [
          5
        ],
        "description": "Chatango is a website used for connecting to a large selection of users.",
        "dom": "iframe[src*='st.chatango.com']",
        "icon": "Chatango.png",
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "st\\.chatango\\.com",
        "website": "https://chatango.com"
      },
      "Chatra": {
        "cats": [
          52
        ],
        "description": "Chatra is a cloud-based live chat platform aimed at small businesses and ecommerce retailers.",
        "icon": "Chatra.svg",
        "js": {
          "ChatraID": "",
          "ChatraSetup": ""
        },
        "pricing": [
          "freemium",
          "payg",
          "recurring"
        ],
        "saas": true,
        "scripts": "call\\.chatra\\.io/chatra\\.js",
        "website": "https://chatra.com",
        "xhr": "chat\\.chatra\\.io/"
      },
      "Checkfront": {
        "cats": [
          5,
          6,
          72
        ],
        "description": "Checkfront is a cloud-based booking management application and ecommerce platform.",
        "icon": "Checkfront.svg",
        "pricing": [
          "low",
          "recurring",
          "poa"
        ],
        "saas": true,
        "scripts": "\\.checkfront\\.com/",
        "website": "https://www.checkfront.com"
      },
      "Cherokee": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:cherokee-project:cherokee",
        "headers": {
          "Server": "^Cherokee(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Cherokee.png",
        "website": "http://www.cherokee-project.com"
      },
      "CherryPy": {
        "cats": [
          22
        ],
        "description": "CherryPy is an object-oriented web application framework using the Python programming language.",
        "headers": {
          "Server": "CherryPy(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "CherryPy.svg",
        "website": "https://cherrypy.org/"
      },
      "Chevereto": {
        "cats": [
          7
        ],
        "description": "Chevereto is an image hosting software that allows you to create a full-featured image hosting website on your own server.",
        "html": "Powered by <a href=\"https?://chevereto\\.com\">",
        "icon": "chevereto.png",
        "implies": "PHP",
        "meta": {
          "generator": "^Chevereto ?([0-9.]+)?$\\;version:\\1"
        },
        "scripts": "/chevereto\\.js",
        "website": "https://chevereto.com/"
      },
      "Chili Piper": {
        "cats": [
          72
        ],
        "description": "Chili Piper is a suite of automated scheduling tools that help revenue teams convert leads.",
        "icon": "Chili Piper.svg",
        "js": {
          "ChiliPiper": ""
        },
        "pricing": [
          "low",
          "freemium"
        ],
        "saas": true,
        "scripts": "js\\.chilipiper\\.com/marketing\\.js",
        "website": "https://www.chilipiper.com/"
      },
      "Chitika": {
        "cats": [
          36
        ],
        "icon": "Chitika.png",
        "js": {
          "ch_client": "",
          "ch_color_site_link": ""
        },
        "scripts": "scripts\\.chitika\\.net/",
        "website": "http://chitika.com"
      },
      "Choices": {
        "cats": [
          19
        ],
        "description": "Choices.js is a lightweight, configurable select box/text input plugin.",
        "icon": "Choices.png",
        "scripts": [
          "choices\\.js/|@([\\d.]+)(?:/assets)?(?:/scripts)?(?:/styles)?(?:/dist)?/choices(?:\\.min)?\\.js|css\\;version:\\1"
        ],
        "website": "https://joshuajohnson.co.uk/Choices/"
      },
      "Chorus": {
        "cats": [
          1
        ],
        "html": "<meta data-chorus-version=",
        "icon": "Chorus.png",
        "website": "https://getchorus.voxmedia.com/"
      },
      "Ckan": {
        "cats": [
          1
        ],
        "headers": {
          "Access-Control-Allow-Headers": "X-CKAN-API-KEY",
          "Link": "<http://ckan\\.org/>; rel=shortlink"
        },
        "icon": "Ckan.png",
        "implies": [
          "Python",
          "Solr",
          "Java",
          "PostgreSQL"
        ],
        "meta": {
          "generator": "^ckan ?([0-9.]+)$\\;version:\\1"
        },
        "website": "http://ckan.org/"
      },
      "Clarity": {
        "cats": [
          66
        ],
        "html": [
          "<clr-main-container",
          "<link [^>]*href=\"[^\"]*clr-ui(?:\\.min)?\\.css"
        ],
        "icon": "clarity.svg",
        "implies": "Angular",
        "js": {
          "ClarityIcons": ""
        },
        "scripts": "clr-angular(?:\\.umd)?(?:\\.min)?\\.js",
        "website": "https://clarity.design/"
      },
      "Classy": {
        "cats": [
          51
        ],
        "description": "Classy is a class library for JavaScript applications.",
        "icon": "classy.png",
        "js": {
          "Classy": ""
        },
        "website": "https://www.classy.org/"
      },
      "ClearSale": {
        "cats": [
          10
        ],
        "description": "ClearSale offers fraud management and chargeback protection services.",
        "icon": "ClearSale.svg",
        "js": {
          "csdm": "\\;confidence:50"
        },
        "scripts": [
          "device\\.clearsale\\.com\\.br"
        ],
        "website": "https://www.clear.sale/"
      },
      "Clearbit Reveal": {
        "cats": [
          10
        ],
        "description": "Clearbit Reveal identifies anonymous visitors to websites.",
        "icon": "Clearbit.png",
        "scripts": [
          "reveal\\.clearbit\\.com/v[(0-9)]/"
        ],
        "website": "https://clearbit.com/reveal"
      },
      "Clerk.io": {
        "cats": [
          76
        ],
        "description": "Clerk.io is an all-in-one ecommerce personalisation platform.",
        "icon": "Clerk.io.svg",
        "js": {
          "__clerk_cb_0": "",
          "__clerk_q": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "\\.clerk\\.io/"
        ],
        "website": "https://clerk.io"
      },
      "CleverTap": {
        "cats": [
          32,
          10
        ],
        "description": "CleverTap is a SaaS based customer lifecycle management and mobile marketing company headquartered in Mountain View, California.",
        "icon": "CleverTap.png",
        "js": {
          "clevertap": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://clevertap.com"
      },
      "Cleverbridge": {
        "cats": [
          6
        ],
        "description": "Cleverbridge is a all-in-one ecommerce and subscription billing solution for software, (SaaS) and digital goods.",
        "icon": "Cleverbridge.svg",
        "js": {
          "cbCartProductSelection": ""
        },
        "scripts": [
          "static-cf\\.cleverbridge\\.\\w+/js/Shop\\.js"
        ],
        "website": "https://www.cleverbridge.com"
      },
      "ClickFunnels": {
        "cats": [
          32
        ],
        "description": "Clickfunnels is an online sales funnel builder that helps businesses market, sell, and deliver their products online.",
        "html": "<meta property=\"cf:app_domain\" content=\"app\\.clickfunnels\\.com\"",
        "icon": "ClickFunnels.png",
        "website": "https://www.clickfunnels.com"
      },
      "ClickHeat": {
        "cats": [
          10
        ],
        "icon": "ClickHeat.png",
        "implies": "PHP",
        "js": {
          "clickHeatServer": ""
        },
        "scripts": "clickheat.*\\.js",
        "website": "http://www.labsmedia.com/clickheat/index.html"
      },
      "ClickTale": {
        "cats": [
          10
        ],
        "description": "ClickTale is a SaaS solution enabling organisations to gain visual in-page analytics.",
        "icon": "ClickTale.png",
        "js": {
          "ClickTale": "",
          "ClickTaleEvent": "",
          "ClickTaleGlobal": "",
          "clickTaleStartEventSignal": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.clicktale\\.net",
        "website": "http://www.clicktale.com"
      },
      "Clickbank": {
        "cats": [
          71
        ],
        "dom": "a[href*='pay.clickbank.net?cbfid'], img[width='1'][src*='hop.clickbank.net?affiliate']",
        "icon": "Clickbank.svg",
        "js": {
          "cbtb": ""
        },
        "scripts": "r\\.wdfl\\.co",
        "website": "https://www.clickbank.com/"
      },
      "Clicky": {
        "cats": [
          10
        ],
        "icon": "Clicky.png",
        "js": {
          "clicky": ""
        },
        "scripts": "static\\.getclicky\\.com",
        "website": "http://getclicky.com"
      },
      "ClientJS": {
        "cats": [
          59,
          83
        ],
        "description": "ClientJS is a JavaScript library for generating browser fingerprints, exposing all the browser data-points.",
        "icon": "ClientJS.png",
        "js": {
          "ClientJS": ""
        },
        "oss": true,
        "scripts": [
          "/clientjs@(\\d.*?)/\\;version:\\1",
          "/ClientJS/(?:(\\d.*?)/)?\\;version:\\1"
        ],
        "website": "http://clientjs.org"
      },
      "Clipboard.js": {
        "cats": [
          19
        ],
        "icon": "Clipboard.js.svg",
        "scripts": "clipboard(?:-([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://clipboardjs.com/"
      },
      "CloudCart": {
        "cats": [
          6
        ],
        "icon": "cloudcart.svg",
        "meta": {
          "author": "^CloudCart LLC$"
        },
        "scripts": "/cloudcart-(?:assets|storage)/",
        "website": "http://cloudcart.com"
      },
      "CloudSuite": {
        "cats": [
          6
        ],
        "cookies": {
          "cs_secure_session": ""
        },
        "icon": "CloudSuite.svg",
        "website": "https://cloudsuite.com"
      },
      "Cloudera": {
        "cats": [
          34
        ],
        "description": "Cloudera is a software platform for data engineering, data warehousing, machine learning and analytics that runs in the cloud or on-premises.",
        "headers": {
          "Server": "cloudera"
        },
        "icon": "Cloudera.png",
        "website": "http://www.cloudera.com"
      },
      "Cloudflare": {
        "cats": [
          31
        ],
        "cookies": {
          "__cfduid": ""
        },
        "description": "Cloudflare is a web-infrastructure and website-security company, providing content-delivery-network services, DDoS mitigation, Internet security, and distributed domain-name-server services.",
        "headers": {
          "Server": "^cloudflare$",
          "cf-cache-status": "",
          "cf-ray": ""
        },
        "icon": "CloudFlare.svg",
        "js": {
          "CloudFlare": ""
        },
        "website": "http://www.cloudflare.com"
      },
      "Cloudflare Browser Insights": {
        "cats": [
          10,
          78
        ],
        "description": "Cloudflare Browser Insights is a tool tool that measures the performance of websites from the perspective of users.",
        "icon": "CloudFlare.svg",
        "js": {
          "__cfBeaconCustomTag": ""
        },
        "scripts": "static\\.cloudflareinsights\\.com/beacon(?:\\.min)?\\.js",
        "website": "http://www.cloudflare.com"
      },
      "Cloudinary": {
        "cats": [
          31
        ],
        "html": [
          "<img[^>]+\\.cloudinary\\.com\\;confidence:80"
        ],
        "icon": "Cloudinary.svg",
        "website": "https://cloudinary.com"
      },
      "ClustrMaps Widget": {
        "cats": [
          35
        ],
        "description": "ClustrMaps widget is a visitor tracker, designed for general web and blog use.",
        "dom": "img[src*='clustrmaps.com']",
        "icon": "ClustrMaps.svg",
        "scripts": "clustrmaps\\.com",
        "website": "https://clustrmaps.com/"
      },
      "CoConstruct": {
        "cats": [
          19
        ],
        "dom": "a[href*='co-construct.com/skins'], iframe[src*='co-construct.com']",
        "icon": "CoConstruct.png",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.coconstruct.com"
      },
      "Coaster CMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:web-feet:coaster_cms",
        "icon": "coaster-cms.png",
        "implies": "Laravel",
        "meta": {
          "generator": "^Coaster CMS v([\\d.]+)$\\;version:\\1"
        },
        "website": "https://www.coastercms.org"
      },
      "Cococart": {
        "cats": [
          6
        ],
        "description": "Cococart is an ecommerce platform.",
        "dom": [
          "meta[property='og:image'][content*='static.cococart.co']",
          "div[style*='static.cococart.co']"
        ],
        "icon": "Cococart.png",
        "pricing": [
          "freemium",
          "recurring",
          "payg"
        ],
        "saas": true,
        "website": "https://www.cococart.co"
      },
      "CoconutSoftware": {
        "cats": [
          5,
          72
        ],
        "cookies": {
          "coconut_calendar": ""
        },
        "description": "Coconut is a cloud-based appointment scheduling solution designed for enterprise financial services organisations such as credit unions, retail banks and more.",
        "icon": "CoconutSoftware.svg",
        "website": "https://www.coconutsoftware.com/"
      },
      "CodeIgniter": {
        "cats": [
          18
        ],
        "cookies": {
          "ci_csrf_token": "^(.+)$\\;version:\\1?2+:",
          "ci_session": "",
          "exp_last_activity": "",
          "exp_tracker": ""
        },
        "cpe": "cpe:/a:codeigniter:codeigniter",
        "html": "<input[^>]+name=\"ci_csrf_token\"\\;version:2+",
        "icon": "CodeIgniter.png",
        "implies": "PHP",
        "website": "http://codeigniter.com"
      },
      "CodeMirror": {
        "cats": [
          19
        ],
        "description": "CodeMirror is a JavaScript component that provides a code editor in the browser.",
        "icon": "CodeMirror.png",
        "js": {
          "CodeMirror": "",
          "CodeMirror.version": "^(.+)$\\;version:\\1"
        },
        "website": "http://codemirror.net"
      },
      "CoinHive": {
        "cats": [
          56
        ],
        "description": "Coinhive is a cryptocurrency mining service.",
        "icon": "CoinHive.svg",
        "js": {
          "CoinHive": ""
        },
        "scripts": [
          "\\/(?:coinhive|(authedmine))(?:\\.min)?\\.js\\;version:\\1?opt-in:",
          "coinhive\\.com/lib"
        ],
        "url": "https?://cnhv\\.co/",
        "website": "https://coinhive.com"
      },
      "CoinHive Captcha": {
        "cats": [
          16,
          56
        ],
        "html": "(?:<div[^>]+class=\"coinhive-captcha[^>]+data-key|<div[^>]+data-key[^>]+class=\"coinhive-captcha)",
        "icon": "CoinHive.svg",
        "scripts": "https?://authedmine\\.com/(?:lib/captcha|captcha)",
        "website": "https://coinhive.com"
      },
      "Coinbase Commerce": {
        "cats": [
          41
        ],
        "description": "Coinbase Commerce is a platform that enables merchants to accept cryptocurrency payments.",
        "dom": "a[href^='https://commerce.coinbase.com/checkout/']",
        "icon": "Coinbase.svg",
        "website": "https://commerce.coinbase.com/"
      },
      "Coinhave": {
        "cats": [
          56
        ],
        "description": "CoinHave is a cryptocurrency mining service.",
        "icon": "coinhave.png",
        "scripts": "https?://coin-have\\.com/c/[0-9a-zA-Z]{4}\\.js",
        "website": "https://coin-have.com/"
      },
      "Coinimp": {
        "cats": [
          56
        ],
        "description": "CoinImp is a cryptocurrency mining service.",
        "icon": "coinimp.png",
        "js": {
          "Client.Anonymous": "\\;confidence:50"
        },
        "scripts": "https?://www\\.hashing\\.win/scripts/min\\.js",
        "website": "https://www.coinimp.com"
      },
      "ColorMeShop": {
        "cats": [
          6
        ],
        "icon": "colormeshop.png",
        "js": {
          "Colorme": ""
        },
        "website": "https://shop-pro.jp"
      },
      "Comandia": {
        "cats": [
          6
        ],
        "html": "<link[^>]+=['\"]//cdn\\.mycomandia\\.com",
        "icon": "Comandia.svg",
        "js": {
          "Comandia": ""
        },
        "website": "http://comandia.com"
      },
      "Combeenation": {
        "cats": [
          6
        ],
        "html": "<iframe[^>]+src=\"[^>]+portal\\.combeenation\\.com",
        "icon": "Combeenation.png",
        "website": "https://www.combeenation.com"
      },
      "Commerce Server": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:microsoft:commerce_server",
        "headers": {
          "COMMERCE-SERVER-SOFTWARE": ""
        },
        "icon": "Commerce Server.png",
        "implies": "Microsoft ASP.NET",
        "website": "http://commerceserver.net"
      },
      "Commerce.js": {
        "cats": [
          6
        ],
        "description": "Commerce.js is an API-first ecommerce platform for developers and businesses.",
        "headers": {
          "Chec-Version": ".*",
          "X-Powered-By": "Commerce.js"
        },
        "icon": "commercejs.png",
        "js": {
          "CommercejsSpace": ""
        },
        "scripts": [
          "cdn\\.chec\\.io/v(\\d+)/commerce\\.js\\;version:\\1",
          "chec/commerce\\.js"
        ],
        "url": "\\.spaces.chec\\.io",
        "website": "https://www.commercejs.com"
      },
      "Commerce7": {
        "cats": [
          6
        ],
        "description": "Commerce7 is an ecommerce platform for wineries.",
        "icon": "Commerce7.svg",
        "pricing": [
          "mid",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": "cdn\\.commerce7\\.com",
        "website": "https://commerce7.com",
        "xhr": "api\\.commerce7\\.com"
      },
      "Concrete5": {
        "cats": [
          1
        ],
        "cookies": {
          "CONCRETE5": ""
        },
        "cpe": "cpe:/a:concrete5:concrete5",
        "icon": "Concrete5.png",
        "implies": "PHP",
        "js": {
          "CCM_IMAGE_PATH": ""
        },
        "meta": {
          "generator": "^concrete5 - ([\\d.]+)$\\;version:\\1"
        },
        "scripts": "/concrete/js/",
        "website": "https://concrete5.org"
      },
      "Conekta": {
        "cats": [
          41
        ],
        "description": "Conekta is a Mexican payment platform.",
        "icon": "Conekta.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": [
          "conektaapi/v([\\d.]+)\\;version:\\1",
          "cdn\\.conekta\\.\\w+/js/(?:v([\\d.]+)|)\\;version:\\1"
        ],
        "website": "https://conekta.com"
      },
      "Contao": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:contao:contao_cms",
        "description": "Contao is an open source CMS that allows you to create websites and scalable web applications.",
        "html": [
          "<!--[^>]+powered by (?:TYPOlight|Contao)[^>]*-->",
          "<link[^>]+(?:typolight|contao)\\.css"
        ],
        "icon": "Contao.svg",
        "implies": "PHP",
        "meta": {
          "generator": "^Contao Open Source CMS$"
        },
        "oss": true,
        "website": "http://contao.org"
      },
      "Contenido": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:contenido:contendio",
        "icon": "Contenido.png",
        "implies": "PHP",
        "meta": {
          "generator": "Contenido ([\\d.]+)\\;version:\\1"
        },
        "website": "http://contenido.org/en"
      },
      "Contensis": {
        "cats": [
          1
        ],
        "icon": "Contensis.png",
        "implies": [
          "Java",
          "CFML"
        ],
        "meta": {
          "generator": "Contensis CMS Version ([\\d.]+)\\;version:\\1"
        },
        "website": "https://zengenti.com/en-gb/products/contensis"
      },
      "ContentBox": {
        "cats": [
          1,
          11
        ],
        "icon": "ContentBox.png",
        "implies": "Adobe ColdFusion",
        "meta": {
          "generator": "ContentBox powered by ColdBox"
        },
        "website": "http://www.gocontentbox.org"
      },
      "ContentSquare": {
        "cats": [
          58
        ],
        "description": "ContentSquare is an enterprise-level UX optimisation platform.",
        "icon": "ContentSquare.png",
        "js": {
          "CS_CONF.trackerDomain": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.contentsquare\\.net/",
        "website": "https://contentsquare.com"
      },
      "Contentful": {
        "cats": [
          1
        ],
        "description": "Contentful is an API-first content management platform to create, manage and publish content on any digital channel.",
        "headers": {
          "x-contentful-request-id": ""
        },
        "html": "<[^>]+(?:https?:)?//(?:assets|downloads|images|videos)\\.(?:ct?fassets\\.net|contentful\\.com)",
        "icon": "Contentful.svg",
        "pricing": [
          "mid",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "website": "http://www.contentful.com",
        "xhr": "cdn\\.contentful\\.com"
      },
      "Contentstack": {
        "cats": [
          1
        ],
        "description": "Contentstack is a headless CMS software designed to help businesses deliver personalised content experiences to audiences via multiple channels.",
        "dom": "img[src*='images.contentstack.io']",
        "icon": "Contentstack.png",
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.contentstack.com"
      },
      "Convert": {
        "cats": [
          74
        ],
        "description": "Convert Experiences is an enterprise A/B testing and personalisation solution for conversion optimisation and data-driven decisions in high-traffic websites.",
        "icon": "Convert.png",
        "js": {
          "convert": "\\;confidence:34",
          "convertData": "\\;confidence:33",
          "convert_temp": "\\;confidence:33"
        },
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.convertexperiments\\.com/js",
        "website": "https://www.convert.com"
      },
      "ConvertFlow": {
        "cats": [
          10,
          74
        ],
        "description": "ConvertFlow is the all-in-one conversion marketing platform.",
        "icon": "ConvertFlow.svg",
        "js": {
          "convertflow": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "(?:app|js)\\.convertflow\\.co",
        "website": "https://www.convertflow.com"
      },
      "ConvertKit": {
        "cats": [
          32,
          75
        ],
        "description": "ConvertKit is an email marketing tool built for content creators.",
        "dom": "form[action*='.convertkit.com'], link[href*='.convertkit.com']",
        "icon": "ConvertKit.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.convertkit\\.com",
        "website": "https://convertkit.com"
      },
      "Cookie Script": {
        "cats": [
          67
        ],
        "description": "Cookie-Script automatically scans, categorizes and adds description to all cookies found on your website.",
        "icon": "CookieScript.png",
        "scripts": "//cookie-script\\.com/s/",
        "website": "https://cookie-script.com/"
      },
      "CookieHub": {
        "cats": [
          67
        ],
        "icon": "CookieHub.png",
        "scripts": [
          "cookiehub\\.net/.*\\.js"
        ],
        "website": "https://www.cookiehub.com"
      },
      "CookieYes": {
        "cats": [
          67
        ],
        "dom": {
          "#cookie-law-info-bar": {
            "text": ""
          }
        },
        "icon": "cookieyes.svg",
        "scripts": [
          "app\\.cookieyes\\.com/client_data/",
          "cdn-cookieyes\\.com/client_data/"
        ],
        "website": "https://www.cookieyes.com/"
      },
      "Cookiebot": {
        "cats": [
          67
        ],
        "description": "Cookiebot is a cloud-driven solution that automatically controls cookies and trackers, enabling full GDPR/ePrivacy and CCPA compliance for websites.",
        "icon": "Cookiebot.svg",
        "scripts": "consent\\.cookiebot\\.com",
        "website": "http://www.cookiebot.com"
      },
      "Cooladata": {
        "cats": [
          76,
          10
        ],
        "description": "Cooladata is a data warehouse and behavioral analytics platform designed for gaming, elearning, ecommerce, SaaS, and media companies.",
        "icon": "Cooladata.png",
        "js": {
          "cooladata": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.cooladata\\.com/",
        "website": "https://www.cooladata.com"
      },
      "Coppermine": {
        "cats": [
          7
        ],
        "cpe": "cpe:/a:coppermine-gallery:coppermine_photo_gallery",
        "description": "Coppermine is an open-source image gallery application.",
        "html": "<!--Coppermine Photo Gallery ([\\d.]+)\\;version:\\1",
        "icon": "Coppermine.png",
        "implies": "PHP",
        "website": "http://coppermine-gallery.net"
      },
      "Corebine": {
        "cats": [
          1
        ],
        "description": "Corebine is a content management system designed for Sports",
        "dom": "#corebine-app",
        "icon": "Corebine.png",
        "js": {
          "corebine": ""
        },
        "pricing": [
          "poa"
        ],
        "website": "https://corebine.com"
      },
      "Cosmoshop": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:cosmoshop:cosmoshop",
        "icon": "Cosmoshop.png",
        "scripts": "cosmoshop_functions\\.js",
        "website": "http://cosmoshop.de"
      },
      "Cotonti": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:cotonti:cotonti_siena",
        "icon": "Cotonti.png",
        "implies": "PHP",
        "meta": {
          "generator": "Cotonti"
        },
        "website": "http://www.cotonti.com"
      },
      "CouchDB": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:apache:couchdb",
        "headers": {
          "Server": "CouchDB/([\\d.]+)\\;version:\\1"
        },
        "icon": "CouchDB.png",
        "website": "http://couchdb.apache.org"
      },
      "Countly": {
        "cats": [
          10
        ],
        "icon": "Countly.png",
        "js": {
          "Countly": ""
        },
        "website": "https://count.ly"
      },
      "CoverManager": {
        "cats": [
          5,
          72
        ],
        "description": "CoverManager is a restaurant table booking widget.",
        "html": "<iframe[^>]*covermanager\\.com/reservation",
        "icon": "CoverManager.svg",
        "website": "https://www.covermanager.com"
      },
      "Cowboy": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^Cowboy$"
        },
        "icon": "Cowboy.png",
        "website": "http://ninenines.eu"
      },
      "CppCMS": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "^CppCMS/([\\d.]+)$\\;version:\\1"
        },
        "icon": "CppCMS.png",
        "website": "http://cppcms.com"
      },
      "Craft CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "CraftSessionId": ""
        },
        "cpe": "cpe:/a:craftcms:craft_cms",
        "description": "Craft CMS is a content management system for building bespoke websites.",
        "headers": {
          "X-Powered-By": "\\bCraft CMS\\b"
        },
        "icon": "Craft CMS.svg",
        "implies": "Yii",
        "oss": true,
        "pricing": [
          "low",
          "freemium",
          "recurring",
          "onetime"
        ],
        "website": "https://craftcms.com"
      },
      "Craft Commerce": {
        "cats": [
          6
        ],
        "headers": {
          "X-Powered-By": "\\bCraft Commerce\\b"
        },
        "icon": "Craft CMS.svg",
        "implies": "Craft CMS",
        "website": "https://craftcommerce.com"
      },
      "Crazy Egg": {
        "cats": [
          10
        ],
        "icon": "Crazy Egg.png",
        "js": {
          "CE2": ""
        },
        "scripts": "script\\.crazyegg\\.com/pages/scripts/\\d+/\\d+\\.js",
        "website": "http://crazyegg.com"
      },
      "Crisp Live Chat": {
        "cats": [
          52
        ],
        "description": "Crisp Live Chat is a live chat solution with free and paid options.",
        "icon": "Crisp Live Chat.svg",
        "js": {
          "$crisp": "",
          "CRISP_WEBSITE_ID": ""
        },
        "website": "https://crisp.chat/"
      },
      "Criteo": {
        "cats": [
          36,
          77
        ],
        "description": "Criteo provides personalised retargeting that works with Internet retailers to serve personalised online display advertisements to consumers who have previously visited the advertiser's website.",
        "icon": "Criteo.svg",
        "js": {
          "Criteo": "",
          "criteo_pubtag": "",
          "criteo_q": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "//(?:cas\\.criteo\\.com|(?:[^/]\\.)?criteo\\.net)/",
          "//static\\.criteo\\.net/js/ld/ld\\.js"
        ],
        "website": "http://criteo.com"
      },
      "Cross Pixel": {
        "cats": [
          77
        ],
        "description": "Cross Pixel is an advertising platform through which advertisers can leverage the marriage of partner audience synergies with the power of retargeting.",
        "icon": "Cross Pixel.png",
        "js": {
          "cp_C4w1ldN2d9PmVrkN": ""
        },
        "scripts": "tag\\.crsspxl\\.com",
        "website": "http://crosspixel.net"
      },
      "CrossBox": {
        "cats": [
          30
        ],
        "description": "CrossBox is a webmail client.",
        "headers": {
          "server": "CBX-WS"
        },
        "icon": "CrossBox.png",
        "website": "https://crossbox.io"
      },
      "Crownpeak": {
        "cats": [
          67
        ],
        "description": "CrownPeak is a cloud-based Digital Experience Management (DEM) platform that is designed to in the management of digital experiences across multiple touch-points, especially for marketing and a freer IT architecture.",
        "icon": "Crownpeak.png",
        "js": {
          "CrownPeakAutocomplete": "",
          "CrownPeakSearch": ""
        },
        "scripts": [
          "c\\.evidon\\.com",
          "js/crownpeak\\."
        ],
        "website": "http://www.crownpeak.com"
      },
      "Crypto-Loot": {
        "cats": [
          56
        ],
        "description": "Crypto-Loot is a browser based web miner for the uPlexa Blockchain.",
        "icon": "Crypto-Loot.png",
        "js": {
          "CRLT.CONFIG.ASMJS_NAME": "",
          "CryptoLoot": ""
        },
        "scripts": [
          "^/crypto-loot\\.com/lib/",
          "^/webmine\\.pro/",
          "^/cryptoloot\\.pro/",
          "/crlt\\.js\\;confidence:75"
        ],
        "website": "https://crypto-loot.com/"
      },
      "CubeCart": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:cubecart:cubecart",
        "html": "(?:Powered by <a href=[^>]+cubecart\\.com|<p[^>]+>Powered by CubeCart)",
        "icon": "CubeCart.png",
        "implies": "PHP",
        "meta": {
          "generator": "cubecart"
        },
        "website": "http://www.cubecart.com"
      },
      "Cufon": {
        "cats": [
          17
        ],
        "description": "Cufon is a tool used to overlap real text with an image.",
        "icon": "Cufon.png",
        "js": {
          "Cufon": ""
        },
        "scripts": "cufon-yui\\.js",
        "website": "http://cufon.shoqolate.com"
      },
      "Cybersource": {
        "cats": [
          41
        ],
        "description": "Cybersource is an ecommerce credit card payment system solution.",
        "icon": "cybersource.png",
        "scripts": "cybersource\\..+\\.js",
        "website": "https://www.cybersource.com/"
      },
      "Czater": {
        "cats": [
          52
        ],
        "description": "Czater is an live chat solution with extended CRM and videochat features.",
        "icon": "Czater.svg",
        "js": {
          "$czater": "",
          "$czaterMethods": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.czater\\.pl",
        "website": "https://www.czater.pl"
      },
      "D3": {
        "cats": [
          25
        ],
        "cpe": "cpe:/a:d3.js_project:d3.js",
        "description": "D3.js is a JavaScript library for producing dynamic, interactive data visualisations in web browsers.",
        "icon": "D3.png",
        "js": {
          "d3.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "/d3(?:\\. v\\d+)?(?:\\.min)?\\.js",
        "website": "http://d3js.org"
      },
      "DERAK.CLOUD": {
        "cats": [
          31
        ],
        "cookies": {
          "__derak_auth": "",
          "__derak_user": ""
        },
        "headers": {
          "Derak-Umbrage": "",
          "Server": "^DERAK\\.CLOUD$"
        },
        "icon": "DerakCloud.png",
        "js": {
          "derakCloud.init": ""
        },
        "website": "https://derak.cloud"
      },
      "DHTMLX": {
        "cats": [
          59
        ],
        "icon": "DHTMLX.png",
        "scripts": "dhtmlxcommon\\.js",
        "website": "http://dhtmlx.com"
      },
      "DM Polopoly": {
        "cats": [
          1
        ],
        "html": "<(?:link [^>]*href|img [^>]*src)=\"/polopoly_fs/",
        "icon": "DM Polopoly.png",
        "implies": "Java",
        "website": "http://www.atex.com/products/dm-polopoly"
      },
      "DNN": {
        "cats": [
          1
        ],
        "cookies": {
          "DotNetNukeAnonymous": ""
        },
        "cpe": "cpe:/a:dnnsoftware:dotnetnuke",
        "headers": {
          "Cookie": "dnn_IsMobile=",
          "DNNOutputCache": "",
          "X-Compressed-By": "DotNetNuke"
        },
        "html": [
          "<!-- by DotNetNuke Corporation",
          "<!-- DNN Platform"
        ],
        "icon": "DNN.png",
        "implies": "Microsoft ASP.NET",
        "js": {
          "DotNetNuke": "",
          "dnn.apiversion": "^(.+)$\\;version:\\1"
        },
        "meta": {
          "generator": "DotNetNuke"
        },
        "oss": true,
        "scripts": [
          "/js/dnncore\\.js",
          "/js/dnn\\.js"
        ],
        "website": "http://dnnsoftware.com"
      },
      "Dancer": {
        "cats": [
          18
        ],
        "description": "Mono.net delivers the a Software-as-a-Service (SaaS) platform to build and sell websites and other digital products.",
        "headers": {
          "Server": "Perl Dancer ([\\d.]+)\\;version:\\1",
          "X-Powered-By": "Perl Dancer ([\\d.]+)\\;version:\\1"
        },
        "icon": "Dancer.png",
        "implies": "Perl",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "http://perldancer.org"
      },
      "Danneo CMS": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "CMS Danneo ([\\d.]+)\\;version:\\1"
        },
        "icon": "Danneo CMS.png",
        "implies": [
          "Apache",
          "PHP"
        ],
        "meta": {
          "generator": "Danneo CMS ([\\d.]+)\\;version:\\1"
        },
        "website": "http://danneo.com"
      },
      "Dart": {
        "cats": [
          27
        ],
        "excludes": [
          "Angular",
          "AngularJS"
        ],
        "html": "/(?:<script)[^>]+(?:type=\"application/dart\")/",
        "icon": "Dart.svg",
        "implies": "AngularDart",
        "js": {
          "___dart__$dart_dartObject_ZxYxX_0_": "",
          "___dart_dispatch_record_ZxYxX_0_": ""
        },
        "scripts": [
          "/(?:\\.)?(?:dart)(?:\\.js)?/",
          "packages/browser/dart\\.js"
        ],
        "website": "https://www.dartlang.org"
      },
      "Darwin": {
        "cats": [
          28
        ],
        "description": "Darwin is the open-source operating system from Apple that forms the basis for macOS.",
        "headers": {
          "Server": "Darwin",
          "X-Powered-By": "Darwin"
        },
        "icon": "Apple.svg",
        "website": "https://opensource.apple.com"
      },
      "DataLife Engine": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:dleviet:datalife_engine",
        "icon": "DataLife Engine.svg",
        "implies": [
          "PHP",
          "Apache"
        ],
        "js": {
          "dle_root": ""
        },
        "meta": {
          "generator": "DataLife Engine"
        },
        "website": "https://dle-news.ru"
      },
      "DataTables": {
        "cats": [
          59
        ],
        "icon": "DataTables.png",
        "implies": "jQuery",
        "scripts": "dataTables.*\\.js",
        "website": "http://datatables.net"
      },
      "Datadog": {
        "cats": [
          78,
          10
        ],
        "description": "Datadog is a SaaS-based monitoring and analytics platform for large-scale applications and infrastructure.",
        "icon": "Datadog.svg",
        "js": {
          "DD_LOGS": "\\;confidence:1",
          "DD_RUM": "\\;confidence:99"
        },
        "pricing": [
          "low",
          "payg",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "scripts": "www\\.datadoghq-browser-agent\\.com",
        "website": "https://www.datadoghq.com"
      },
      "Datadome": {
        "cats": [
          16
        ],
        "cookies": {
          "datadome": "",
          "datadome-_zldp": "",
          "datadome-_zldt": ""
        },
        "headers": {
          "Server": "^DataDome$",
          "X-DataDome": "",
          "X-DataDome-CID": ""
        },
        "icon": "datadome.png",
        "scripts": "^https://ct\\.datadome\\.co/[a-z]\\.js$",
        "website": "https://datadome.co/"
      },
      "DatoCMS": {
        "cats": [
          1
        ],
        "html": "<[^>]+https://www\\.datocms-assets\\.com",
        "icon": "datocms.svg",
        "website": "https://www.datocms.com"
      },
      "Day.js": {
        "cats": [
          59
        ],
        "icon": "Day.js.svg",
        "js": {
          "dayjs": ""
        },
        "website": "https://github.com/iamkun/dayjs"
      },
      "Dealer Spike": {
        "cats": [
          6,
          32
        ],
        "description": "Dealer Spike is a digital marketing and advertising company focused that helps dealers grow their business.",
        "dom": "meta[name='author'][content*='Dealer Spike']",
        "icon": "Dealer Spike.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.dealerspike\\.com",
        "website": "https://www.dealerspike.com",
        "xhr": "\\.dealerspike\\.com"
      },
      "Debian": {
        "cats": [
          28
        ],
        "cpe": "cpe:/o:debian:debian_linux",
        "description": "Debian is a Linux software which is a free open-source software.",
        "headers": {
          "Server": "Debian",
          "X-Powered-By": "(?:Debian|dotdeb|(potato|woody|sarge|etch|lenny|squeeze|wheezy|jessie|stretch|buster|sid))\\;version:\\1"
        },
        "icon": "Debian.png",
        "website": "https://debian.org"
      },
      "Decibel": {
        "cats": [
          10,
          74
        ],
        "description": "Decibel is a behavioral analysis solution that helps users gain actionable insights about their digital audience.",
        "icon": "Decibel.svg",
        "js": {
          "decibelInsight": "",
          "decibelInsightLayer": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.decibelinsight\\.net",
        "website": "https://decibel.com"
      },
      "DedeCMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:dedecms:dedecms",
        "icon": "DedeCMS.png",
        "implies": "PHP",
        "js": {
          "DedeContainer": ""
        },
        "scripts": "dedeajax",
        "website": "http://dedecms.com"
      },
      "Detectify": {
        "cats": [
          16
        ],
        "description": "Detectify is an automated scanner that checks your web application for vulnerabilities.",
        "dns": {
          "TXT": [
            "detectify-verification"
          ]
        },
        "icon": "Detectify.svg",
        "website": "https://detectify.com/"
      },
      "Didomi": {
        "cats": [
          67
        ],
        "description": "Didomi is a consent management platform helping brands and businesses collect, store and leverage their customer consents.",
        "icon": "didomi.png",
        "scripts": "sdk\\.privacy-center\\.org/.*/loader\\.js",
        "website": "https://www.didomi.io/en/consent-preference-management"
      },
      "Digest": {
        "cats": [
          16
        ],
        "description": "Digest is an authentication method based on a MD5 hash used by web servers.",
        "headers": {
          "WWW-Authenticate": "^Digest"
        },
        "website": "https://tools.ietf.org/html/rfc7616"
      },
      "DigiCert": {
        "cats": [
          70
        ],
        "certIssuer": "DigiCert",
        "icon": "DigiCert.svg",
        "website": "https://www.digicert.com/"
      },
      "Digistore24": {
        "cats": [
          71
        ],
        "description": "Digistore24 is a German digital reselling and affiliate marketing platform.",
        "dom": "a[href*='www.digistore24.com'][target='_blank']",
        "icon": "Digistore24.svg",
        "js": {
          "DIGISTORE_LINK_ID_KEY": "",
          "DIGISTORE_VENDORKEY": "",
          "getTheSourceForDigistoreLinks": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": [
          "digistore/digistore\\.js",
          "www\\.digistore24\\.com"
        ],
        "website": "https://www.digistore24.com"
      },
      "DigitalRiver": {
        "cats": [
          6
        ],
        "cookies": {
          "X-DR-SHOPPER-ets": "",
          "X-DR-THEME": "^\\d+$"
        },
        "description": "Digital River provides global ecommerce, payments and marketing services.",
        "icon": "DigitalRiver.svg",
        "scripts": "/drh\\.img\\.digitalriver\\.\\w+/DRHM/",
        "website": "https://www.digitalriver.com"
      },
      "DirectAdmin": {
        "cats": [
          9
        ],
        "cpe": "cpe:/a:directadmin:directadmin",
        "description": "DirectAdmin is a graphical web-based web hosting control panel designed to make administration of websites easier.",
        "headers": {
          "Server": "DirectAdmin Daemon v([\\d.]+)\\;version:\\1"
        },
        "html": "<a[^>]+>DirectAdmin</a> Web Control Panel",
        "icon": "DirectAdmin.png",
        "implies": [
          "PHP",
          "Apache"
        ],
        "website": "https://www.directadmin.com"
      },
      "Discourse": {
        "cats": [
          2
        ],
        "description": "Discourse is an open-source internet forum and mailing list management software application.",
        "icon": "Discourse.png",
        "implies": "Ruby on Rails",
        "js": {
          "Discourse": ""
        },
        "meta": {
          "generator": "Discourse(?: ?/?([\\d.]+\\d))?\\;version:\\1"
        },
        "website": "https://discourse.org"
      },
      "Discuz! X": {
        "cats": [
          2
        ],
        "description": "Discuz! X is an internet forum software written in PHP and supports MySQL and PostgreSQL databases.",
        "icon": "Discuz X.png",
        "implies": "PHP",
        "js": {
          "DISCUZCODE": "",
          "discuzVersion": "^(.+)$\\;version:\\1",
          "discuz_uid": ""
        },
        "meta": {
          "generator": "Discuz! X([\\d\\.]+)?\\;version:\\1"
        },
        "website": "http://www.discuz.net"
      },
      "Disqus": {
        "cats": [
          15
        ],
        "cpe": "cpe:/a:disqus:disqus_comment_system",
        "description": "Disqus is a worldwide blog comment hosting service for web sites and online communities that use a networked platform.",
        "html": "<div[^>]+id=\"disqus_thread\"",
        "icon": "Disqus.svg",
        "js": {
          "DISQUS": "",
          "disqus_shortname": "",
          "disqus_url": ""
        },
        "scripts": "disqus_url",
        "website": "https://disqus.com"
      },
      "District M": {
        "cats": [
          36
        ],
        "description": "District M is a programmatic advertising exchange.",
        "dom": "iframe[src*='.districtm.io']",
        "icon": "District M.svg",
        "saas": true,
        "website": "https://districtm.net",
        "xhr": "\\.districtm\\.io"
      },
      "Divi": {
        "cats": [
          51
        ],
        "description": "Divi is a WordPress Theme and standalone WordPress plugin from Elegant themes that allows users to build websites using the visual drag-and-drop Divi page builder.",
        "icon": "Divi.png",
        "js": {
          "DIVI": ""
        },
        "meta": {
          "generator": "Divi"
        },
        "pricing": [
          "low",
          "recurring",
          "onetime"
        ],
        "saas": true,
        "scripts": "Divi/js/custom\\.(?:min|unified)\\.js\\?ver=([\\d.]+)\\;version:\\1",
        "website": "https://www.elegantthemes.com/gallery/divi"
      },
      "Django": {
        "cats": [
          18
        ],
        "cpe": "cpe:/a:djangoproject:django",
        "description": "Django is a Python-based free and open-source web application framework.",
        "html": "(?:powered by <a[^>]+>Django ?([\\d.]+)?<\\/a>|<input[^>]*name=[\"']csrfmiddlewaretoken[\"'][^>]*>)\\;version:\\1",
        "icon": "Django.png",
        "implies": "Python",
        "js": {
          "__admin_media_prefix__": "",
          "django": ""
        },
        "website": "https://djangoproject.com"
      },
      "DocFX": {
        "cats": [
          4
        ],
        "description": "DocFX is a tool for building and publishing API documentation for .NET projects.",
        "icon": "DocFX.svg",
        "meta": {
          "docfx:navrel": "toc.html",
          "docfx:tocrel": "toc.html",
          "generator": "docfx\\s([\\d\\.]+)\\;version:\\1"
        },
        "oss": true,
        "website": "https://github.com/dotnet/docfx"
      },
      "Docker": {
        "cats": [
          60
        ],
        "cpe": "cpe:/a:docker:engine",
        "description": "Docker is a tool designed to make it easier to create, deploy, and run applications by using containers.",
        "html": "<!-- This comment is expected by the docker HEALTHCHECK  -->",
        "icon": "Docker.svg",
        "website": "https://www.docker.com/"
      },
      "DocuSign": {
        "cats": [
          19
        ],
        "description": "DocuSign allows organisations to manage electronic agreements.",
        "dns": {
          "TXT": [
            "docusign"
          ]
        },
        "icon": "DocuSign.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.docusign.com",
        "xhr": "docusign\\.net"
      },
      "Docusaurus": {
        "cats": [
          4,
          57
        ],
        "description": "Docusaurus is a tool for teams to publish documentation websites.",
        "icon": "docusaurus.svg",
        "implies": [
          "React",
          "webpack"
        ],
        "js": {
          "search.indexName": ""
        },
        "meta": {
          "generator": "^Docusaurus(?: v(.+))?$\\;version:\\1"
        },
        "website": "https://docusaurus.io/"
      },
      "Dojo": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:dojotoolkit:dojo",
        "icon": "Dojo.png",
        "js": {
          "dojo": "",
          "dojo.version.major": "^(.+)$\\;version:\\1"
        },
        "scripts": "([\\d.]+)/dojo/dojo(?:\\.xd)?\\.js\\;version:\\1",
        "website": "https://dojotoolkit.org"
      },
      "Dokeos": {
        "cats": [
          21
        ],
        "description": "Dokeos is an e-learning and course management web application.",
        "headers": {
          "X-Powered-By": "Dokeos"
        },
        "html": "(?:Portal <a[^>]+>Dokeos|@import \"[^\"]+dokeos_blue)",
        "icon": "Dokeos.png",
        "implies": [
          "PHP",
          "Xajax",
          "jQuery",
          "CKEditor"
        ],
        "meta": {
          "generator": "Dokeos"
        },
        "website": "https://dokeos.com"
      },
      "DokuWiki": {
        "cats": [
          8
        ],
        "cookies": {
          "DokuWiki": ""
        },
        "cpe": "cpe:/a:dokuwiki:dokuwiki",
        "description": "DokuWiki is a Free open-source wiki software.",
        "html": [
          "<div[^>]+id=\"dokuwiki__>",
          "<a[^>]+href=\"#dokuwiki__"
        ],
        "icon": "DokuWiki.png",
        "implies": "PHP",
        "meta": {
          "generator": "^DokuWiki( Release [\\d-]+)?\\;version:\\1"
        },
        "website": "https://www.dokuwiki.org"
      },
      "Dotclear": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:dotclear:dotclear",
        "headers": {
          "X-Dotclear-Static-Cache": ""
        },
        "icon": "Dotclear.png",
        "implies": "PHP",
        "website": "http://dotclear.org"
      },
      "Dotdigital": {
        "cats": [
          32,
          10
        ],
        "description": "Dotdigital is an all-in-one cloud-based customer engagement multichannel marketing platform.",
        "icon": "Dotdigital.svg",
        "js": {
          "dmPt": "\\;confidence:25",
          "dm_insight_id": "",
          "dmtrackingobjectname": "\\;confidence:75"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "js/_dmptv([\\d.]+)\\.js\\;version:\\1",
        "website": "https://dotdigital.com"
      },
      "DoubleClick Ad Exchange (AdX)": {
        "cats": [
          36
        ],
        "description": "DoubleClick Ad Exchange is a real-time marketplace to buy and sell display advertising space.",
        "icon": "DoubleClick.svg",
        "saas": true,
        "scripts": [
          "googlesyndication\\.com/pagead/show_ads\\.js",
          "tpc\\.googlesyndication\\.com/safeframe",
          "googlesyndication\\.com.*abg\\.js"
        ],
        "website": "http://www.doubleclickbygoogle.com/solutions/digital-marketing/ad-exchange/"
      },
      "DoubleClick Campaign Manager (DCM)": {
        "cats": [
          36
        ],
        "icon": "DoubleClick.svg",
        "scripts": "2mdn\\.net",
        "website": "http://www.doubleclickbygoogle.com/solutions/digital-marketing/campaign-manager/"
      },
      "DoubleClick Floodlight": {
        "cats": [
          36
        ],
        "icon": "DoubleClick.svg",
        "scripts": "https?://fls\\.doubleclick\\.net",
        "website": "http://support.google.com/ds/answer/6029713?hl=en"
      },
      "DoubleClick for Publishers (DFP)": {
        "cats": [
          36
        ],
        "description": "DoubleClick for Publishers (DFP) is a hosted ad serving platform that streamlines your ad management.",
        "icon": "DoubleClick.svg",
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "googletagservices\\.com/tag/js/gpt(?:_mobile)?\\.js",
        "website": "http://www.google.com/dfp"
      },
      "Doxygen": {
        "cats": [
          4
        ],
        "cpe": "cpe:/a:doxygen:doxygen",
        "description": "Doxygen is a documentation generator, a tool for writing software reference documentation.",
        "html": "(?:<!-- Generated by Doxygen ([\\d.]+)|<link[^>]+doxygen\\.css)\\;version:\\1",
        "icon": "Doxygen.png",
        "meta": {
          "generator": "Doxygen ([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.doxygen.nl/"
      },
      "Draft.js": {
        "cats": [
          20
        ],
        "description": "Draft.js is a JavaScript rich text editor framework, built for React.",
        "icon": "draftjs.png",
        "scripts": "draft-js(@|/)([\\d.]+)\\;version:\\2",
        "website": "https://draftjs.org/"
      },
      "DreamWeaver": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:adobe:dreamweaver",
        "description": "Dreamweaver is a development tool for creating, publishing, and managing websites and mobile content.",
        "html": "<!--[^>]*(?:InstanceBeginEditable|Dreamweaver([^>]+)target|DWLayoutDefaultTable)\\;version:\\1",
        "icon": "DreamWeaver.png",
        "js": {
          "MM_preloadImages": "",
          "MM_showHideLayers": "",
          "MM_showMenu": ""
        },
        "website": "https://www.adobe.com/products/dreamweaver.html"
      },
      "Drift": {
        "cats": [
          52
        ],
        "description": "Drift is a conversational marketing platform.",
        "icon": "Drift.svg",
        "js": {
          "drift": "",
          "driftt": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.drift.com/"
      },
      "Drupal": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:drupal:drupal",
        "description": "Drupal is a free and open-source web content management framework.",
        "headers": {
          "Expires": "19 Nov 1978",
          "X-Drupal-Cache": "",
          "X-Generator": "^Drupal(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "html": "<(?:link|style)[^>]+\"/sites/(?:default|all)/(?:themes|modules)/",
        "icon": "Drupal.svg",
        "implies": "PHP",
        "js": {
          "Drupal": ""
        },
        "meta": {
          "generator": "^Drupal(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "oss": true,
        "scripts": "drupal\\.js",
        "website": "https://drupal.org"
      },
      "Drupal Commerce": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:commerceguys:commerce",
        "description": "Drupal Commerce is open-source ecommerce software that augments the content management system Drupal.",
        "dom": "aside#cart-offcanvas, form.commerce-order-item-add-to-cart-form,form.commerce-add-to-cart",
        "html": "<[^>]+(?:id=\"block[_-]commerce[_-]cart[_-]cart|class=\"commerce[_-]product[_-]field)",
        "icon": "Drupal Commerce.png",
        "implies": "Drupal",
        "oss": true,
        "scripts": [
          "/modules/(?:contrib/)?commerce/js/conditions\\.js\\;confidence:50",
          "/profiles/commerce_kickstart/modules/contrib/commerce/modules/checkout/commerce_checkout\\.js\\;confidence:50",
          "/sites/(?:all|default)/modules/(?:contrib/)?commerce/modules/checkout/commerce_checkout\\.js\\;confidence:50"
        ],
        "website": "http://drupalcommerce.org"
      },
      "Duda": {
        "cats": [
          1
        ],
        "html": "<div[^>]*id=\"P6iryBW0Wu\"",
        "icon": "duda.png",
        "js": {
          "SystemID": "^.*DIRECT.*$",
          "version": "^(.*)$\\;version:\\1\\;confidence:0"
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "https://www.duda.co/website-builder"
      },
      "Duopana": {
        "cats": [
          1,
          11,
          51,
          53
        ],
        "description": "Duopana is a platform for creating online communities, blogs and managing collaborative content.",
        "html": "(?:<!-- /*BeraCode script)",
        "icon": "Duopana.svg",
        "scripts": "\\*berajs.beracode.com\\*",
        "website": "https://duopana.com/"
      },
      "Dynamic Yield": {
        "cats": [
          74,
          76
        ],
        "cookies": {
          "_dy_geo": "",
          "_dy_ses_load_seq": ""
        },
        "description": "Dynamic Yield is a provider of automated conversion optimisation tools for marketers and retailers.",
        "icon": "DynamicYield.svg",
        "js": {
          "recommendationContext": ""
        },
        "scripts": "cdn\\.dynamicyield\\.\\w+/",
        "website": "https://www.dynamicyield.com"
      },
      "Dynamicweb": {
        "cats": [
          1,
          6,
          10
        ],
        "cookies": {
          "Dynamicweb": ""
        },
        "icon": "Dynamicweb.svg",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "Dynamicweb ([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.dynamicweb.dk"
      },
      "Dynatrace": {
        "cats": [
          10
        ],
        "cookies": {
          "dtCookie1": ""
        },
        "icon": "Dynatrace.png",
        "js": {
          "dtrum": ""
        },
        "website": "https://www.dynatrace.com"
      },
      "EC-CUBE": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:lockon:ec-cube",
        "icon": "ec-cube.png",
        "implies": "PHP",
        "oss": true,
        "scripts": [
          "eccube\\.js",
          "win_op\\.js"
        ],
        "website": "http://www.ec-cube.net"
      },
      "EKM": {
        "cats": [
          6
        ],
        "cookies": {
          "ekmpowershop": ""
        },
        "icon": "EKM.png",
        "js": {
          "_ekmpinpoint": ""
        },
        "website": "https://www.ekm.com"
      },
      "ELOG": {
        "cats": [
          19
        ],
        "html": "<title>ELOG Logbook Selection</title>",
        "icon": "ELOG.png",
        "website": "http://midas.psi.ch/elog"
      },
      "ELOG HTTP": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "ELOG HTTP ?([\\d.-]+)?\\;version:\\1"
        },
        "icon": "ELOG.png",
        "implies": "ELOG",
        "website": "http://midas.psi.ch/elog"
      },
      "EPages": {
        "cats": [
          6
        ],
        "headers": {
          "X-epages-RequestId": ""
        },
        "icon": "epages.png",
        "js": {
          "epages": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "http://www.epages.com/"
      },
      "EPiServer": {
        "cats": [
          1
        ],
        "cookies": {
          "EPiServer": "",
          "EPiTrace": ""
        },
        "cpe": "cpe:/a:episerver:episerver",
        "icon": "EPiServer.png",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "EPiServer"
        },
        "website": "http://episerver.com"
      },
      "EPrints": {
        "cats": [
          19
        ],
        "icon": "EPrints.png",
        "implies": "Perl",
        "js": {
          "EPJS_menu_template": "",
          "EPrints": ""
        },
        "meta": {
          "generator": "EPrints ([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.eprints.org"
      },
      "EX.CO": {
        "cats": [
          5
        ],
        "description": "EX.CO (formerly Playbuzz) is an online publishing platform for publishers, brand agencies, and individual content creators to create content in interactive formats such as polls, quizzes, lists, video snippets, slideshows, and countdowns.",
        "icon": "EX.CO.svg",
        "js": {
          "__EXCO": "",
          "__EXCO_INTEGRATION_TYPE": "",
          "excoPixelUrl": ""
        },
        "pricing": [
          "poa"
        ],
        "scripts": [
          "\\.ex\\.co",
          "\\.playbuzz\\.com"
        ],
        "website": "https://ex.co",
        "xhr": "prd-collector-anon\\.ex\\.co"
      },
      "EasyDigitalDownloads": {
        "cats": [
          6
        ],
        "description": "Easy Digital Downloads is a WordPress ecommerce plugin that focuses purely on digital products.",
        "icon": "EasyDigitalDownloads.svg",
        "meta": {
          "generator": "^Easy Digital Downloads v(.*)$\\;version:\\1"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://easydigitaldownloads.com"
      },
      "EasyEngine": {
        "cats": [
          47,
          9
        ],
        "description": "EasyEngine is a command-line tool for the Nginx web servers to manage WordPress sites that are running on the LEMP Stack.",
        "headers": {
          "x-powered-by": "^EasyEngine (.*)$\\;version:\\1"
        },
        "icon": "EasyEngine.png",
        "implies": "Docker",
        "website": "https://easyengine.io"
      },
      "Ecwid": {
        "cats": [
          6
        ],
        "icon": "ecwid.svg",
        "js": {
          "Ecwid": "",
          "EcwidCart": ""
        },
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": [
          "https://app\\.multiscreenstore\\.com/script\\.js",
          "https://app\\.ecwid\\.com/script\\.js"
        ],
        "website": "https://www.ecwid.com/"
      },
      "EdgeCast": {
        "cats": [
          31
        ],
        "description": "EdgeCast is a content delivery network (CDN) that accelerated and delivers static content to users around the world.",
        "headers": {
          "Server": "^ECD\\s\\(\\S+\\)"
        },
        "icon": "EdgeCast.png",
        "url": "https?://(?:[^/]+\\.)?edgecastcdn\\.net/",
        "website": "http://www.edgecast.com"
      },
      "Elasticsearch": {
        "cats": [
          29
        ],
        "description": "Elasticsearch is a search engine based on the Lucene library. It provides a distributed, multitenant-capable full-text search engine with an HTTP web interface and schema-free JSON documents.",
        "icon": "Elasticsearch.svg",
        "website": "https://www.elastic.co"
      },
      "Elcodi": {
        "cats": [
          6
        ],
        "headers": {
          "X-Elcodi": ""
        },
        "icon": "Elcodi.png",
        "implies": [
          "PHP",
          "Symfony"
        ],
        "website": "http://elcodi.io"
      },
      "Eleanor CMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:eleanor-cms:eleanor_cms",
        "icon": "Eleanor CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "Eleanor"
        },
        "website": "http://eleanor-cms.ru"
      },
      "Element UI": {
        "cats": [
          12
        ],
        "html": [
          "<(?:div|button) class=\"el-(?:table-column|table-filter|popper|pagination|pager|select-group|form|form-item|color-predefine|color-hue-slider|color-svpanel|color-alpha-slider|color-dropdown|color-picker|badge|tree|tree-node|select|message|dialog|checkbox|checkbox-button|checkbox-group|container|steps|carousel|menu|menu-item|submenu|menu-item-group|button|button-group|card|table|select-dropdown|row|tabs|notification|radio|progress|progress-bar|tag|popover|tooltip|cascader|cascader-menus|cascader-menu|time-spinner|spinner|spinner-inner|transfer|transfer-panel|rate|slider|dropdown|dropdown-menu|textarea|input|input-group|popup-parent|radio-group|main|breadcrumb|time-range-picker|date-range-picker|year-table|date-editor|range-editor|time-spinner|date-picker|time-panel|date-table|month-table|picker-panel|collapse|collapse-item|alert|select-dropdown|select-dropdown__empty|select-dropdown__wrap|select-dropdown__list|scrollbar|switch|carousel|upload|upload-dragger|upload-list|upload-cover|aside|input-number|header|message-box|footer|radio-button|step|autocomplete|autocomplete-suggestion|loading-parent|loading-mask|loading-spinner|)"
        ],
        "icon": "ElementUI.svg",
        "implies": "Vue.js",
        "website": "https://element.eleme.io/"
      },
      "Elementor": {
        "cats": [
          51
        ],
        "description": "Elementor is a website builder platform for professionals on WordPress.",
        "html": [
          "<div class=(?:\"|')[^\"']*elementor",
          "<section class=(?:\"|')[^\"']*elementor",
          "<link [^>]*href=(?:\"|')[^\"']*elementor/assets",
          "<link [^>]*href=(?:\"|')[^\"']*uploads/elementor/css"
        ],
        "icon": "Elementor.png",
        "js": {
          "elementorFrontend.getElements": ""
        },
        "scripts": "elementor/assets/js/[^/]+\\.js\\?ver=([\\d.]+)$\\;version:\\1",
        "website": "https://elementor.com"
      },
      "Elm": {
        "cats": [
          27,
          12
        ],
        "icon": "elm.svg",
        "js": {
          "Elm.Main.embed": "\\;version:0.18",
          "Elm.Main.init": "\\;version:0.19"
        },
        "website": "https://elm-lang.org/"
      },
      "Elm-ui": {
        "cats": [
          66
        ],
        "html": "<style>[\\s\\S]*\\.explain > \\.s[\\s\\S]*\\.explain > \\.ctr > \\.s",
        "icon": "elm.svg",
        "implies": "Elm",
        "website": "https://github.com/mdgriffith/elm-ui"
      },
      "Eloqua": {
        "cats": [
          32
        ],
        "description": "Eloqua is a Software-as-a-Service (SaaS) platform for marketing automation offered that aims to help B2B marketers and organisations manage marketing campaigns and sales lead generation.",
        "icon": "Oracle.png",
        "js": {
          "elqCurESite": "",
          "elqLoad": "",
          "elqSiteID": "",
          "elq_global": ""
        },
        "scripts": "elqCfg\\.js",
        "website": "http://eloqua.com"
      },
      "Emarsys": {
        "cats": [
          32,
          76
        ],
        "description": "Emarsys is a cloud-based B2C marketing platform.",
        "icon": "Emarsys.svg",
        "js": {
          "Scarab": "",
          "ScarabQueue": ""
        },
        "scripts": "(?:static|cdn)\\.scarabresearch\\.com",
        "website": "https://emarsys.com/"
      },
      "EmbedThis Appweb": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:embedthis:appweb",
        "headers": {
          "Server": "Mbedthis-Appweb(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Embedthis.png",
        "website": "http://embedthis.com/appweb"
      },
      "Ember.js": {
        "cats": [
          12
        ],
        "cpe": "cpe:/a:emberjs:ember.js",
        "icon": "Ember.js.png",
        "implies": "Handlebars",
        "js": {
          "Ember": "",
          "Ember.VERSION": "^(.+)$\\;version:\\1"
        },
        "website": "http://emberjs.com"
      },
      "Emotion": {
        "cats": [
          12,
          47
        ],
        "description": "Emotion is a library designed for writing CSS styles with JavaScript.",
        "dom": "style[data-emotion], style[data-emotion-css]",
        "icon": "Emotion.png",
        "website": "http://emotion.sh"
      },
      "Engagio": {
        "cats": [
          10
        ],
        "icon": "Engagio.png",
        "scripts": [
          "web-analytics\\.engagio\\.com/js/ei\\.js",
          "web-analytics\\.engagio\\.com/api/"
        ],
        "website": "https://www.engagio.com/"
      },
      "Ensighten": {
        "cats": [
          42
        ],
        "description": "Ensighten is a solution that enables secure management, implementation and control of website technologies.",
        "icon": "ensighten.png",
        "scripts": "//nexus\\.ensighten\\.com/",
        "website": "https://success.ensighten.com/hc/en-us"
      },
      "Envoy": {
        "cats": [
          64
        ],
        "cpe": "cpe:/a:envoyproxy:envoy",
        "description": "Envoy is an open-source edge and service proxy, designed for cloud-native applications.",
        "headers": {
          "Server": "^envoy$",
          "x-envoy-upstream-service-time": ""
        },
        "icon": "Envoy.png",
        "website": "https://www.envoyproxy.io/"
      },
      "Enyo": {
        "cats": [
          12,
          26
        ],
        "description": "Enyo is an open-source JavaScript framework for cross-platform for mobile, desktop, TV and web applications.",
        "icon": "Enyo.png",
        "js": {
          "enyo": ""
        },
        "scripts": "enyo\\.js",
        "website": "http://enyojs.com"
      },
      "Epoch": {
        "cats": [
          25
        ],
        "html": "<link[^>]+?href=\"[^\"]+epoch(?:\\.min)?\\.css",
        "implies": "D3",
        "scripts": "epoch(?:\\.min)?\\.js",
        "website": "https://fastly.github.io/epoch"
      },
      "Epom": {
        "cats": [
          36
        ],
        "icon": "Epom.png",
        "js": {
          "epomCustomParams": ""
        },
        "url": "^https?://(?:[^/]+\\.)?ad(?:op)?shost1\\.com/",
        "website": "http://epom.com"
      },
      "EqualWeb": {
        "cats": [
          68
        ],
        "description": "EqualWeb provides web accessibility, digital accessibility solutions and helps people with disabilities access digital information.",
        "icon": "EqualWeb.png",
        "scripts": "cdn\\.equalweb\\.com.*\\.js",
        "website": "https://www.equalweb.com/"
      },
      "Erlang": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:erlang:erlang%2fotp",
        "description": "Erlang is a general-purpose, concurrent, functional programming language, and a garbage-collected runtime system.",
        "headers": {
          "Server": "Erlang( OTP/(?:[\\d.ABR-]+))?\\;version:\\1"
        },
        "icon": "Erlang.png",
        "website": "http://www.erlang.org"
      },
      "Essential JS 2": {
        "cats": [
          12,
          59
        ],
        "html": "<[^>]+ class ?= ?\"(?:e-control|[^\"]+ e-control)(?: )[^\"]* e-lib\\b",
        "icon": "syncfusion.svg",
        "website": "https://www.syncfusion.com/javascript-ui-controls"
      },
      "Estore Compare": {
        "cats": [
          74
        ],
        "description": "Estore Compare is a website optimisation software that offers A/B testing, CVR and LTV measuring tools.",
        "icon": "EstoreCompare.svg",
        "scripts": "cdn\\d+\\.estore\\.jp/",
        "website": "https://estore.co.jp/estorecompare/"
      },
      "Estore Shopserve": {
        "cats": [
          6
        ],
        "description": "Estore Shopserve is an all-in-one payment processing and ecommerce solution.",
        "icon": "EstoreShopserve.svg",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cart\\d+\\.shopserve\\.jp/",
        "website": "https://estore.co.jp/shopserve"
      },
      "Etherpad": {
        "cats": [
          24
        ],
        "cpe": "cpe:/a:etherpad:etherpad",
        "description": "Etherpad is an open-source, web-based collaborative real-time editor, allowing authors to simultaneously edit a text document, and see all of the participants' edits in real-time, with the ability to display each author's text in their own colour.",
        "headers": {
          "Server": "^Etherpad"
        },
        "icon": "etherpad.png",
        "implies": "Node.js",
        "js": {
          "padeditbar": "",
          "padimpexp": ""
        },
        "scripts": [
          "/ep_etherpad-lite/"
        ],
        "website": "https://etherpad.org"
      },
      "Etracker": {
        "cats": [
          10,
          74
        ],
        "description": "Etracker is a web optimisation solution.",
        "icon": "Etracker.png",
        "js": {
          "_etracker": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.etracker\\.com",
        "website": "https://www.etracker.com"
      },
      "Eveve": {
        "cats": [
          5,
          72
        ],
        "description": "Eveve is a restaurant table booking widget.",
        "html": "<iframe[^>]*[\\w]+\\.eveve\\.com",
        "icon": "Eveve.svg",
        "implies": "PHP",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.eveve.com"
      },
      "Exhibit": {
        "cats": [
          25
        ],
        "icon": "Exhibit.png",
        "js": {
          "Exhibit": "",
          "Exhibit.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "exhibit.*\\.js",
        "website": "http://simile-widgets.org/exhibit/"
      },
      "ExitIntel": {
        "cats": [
          32
        ],
        "description": "ExitIntel is a full service conversion optimisation agency that focuses on ecommerce companies.",
        "icon": "ExitIntel.png",
        "js": {
          "exitintel.VERSION": "(.+)$\\;version:\\1",
          "exitintelAccount": "",
          "exitintelConfig": ""
        },
        "pricing": [
          "high",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": "(?:get.)?exitintel\\.com",
        "website": "https://exitintelligence.com"
      },
      "ExpertRec": {
        "cats": [
          29
        ],
        "description": "ExpertRec is a collaborative Web search engine, which allows users share search histories through a web browser.",
        "icon": "ExpertRec.png",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "expertrec\\.com/api/js/ci_common\\.js\\?id=.*",
        "website": "https://www.expertrec.com"
      },
      "Express": {
        "cats": [
          18,
          22
        ],
        "cpe": "cpe:/a:expressjs:express",
        "description": "Express is a web application framework for Node.js, released as free and open-source software under the MIT License. It is designed for building web applications and APIs.",
        "headers": {
          "X-Powered-By": "^Express$"
        },
        "icon": "Express.png",
        "implies": "Node.js",
        "website": "http://expressjs.com"
      },
      "ExpressionEngine": {
        "cats": [
          1
        ],
        "cookies": {
          "exp_csrf_token": "",
          "exp_last_activity": "",
          "exp_tracker": ""
        },
        "cpe": "cpe:/a:ellislab:expressionengine",
        "description": "ExpressionEngine is a free and open-source CMS.",
        "icon": "ExpressionEngine.svg",
        "implies": "PHP",
        "oss": true,
        "website": "http://expressionengine.com"
      },
      "ExtJS": {
        "cats": [
          12
        ],
        "cpe": "cpe:/a:sencha:ext_js",
        "icon": "ExtJS.png",
        "js": {
          "Ext": "",
          "Ext.version": "^(.+)$\\;version:\\1",
          "Ext.versions.extjs.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "ext-base\\.js",
        "website": "https://www.sencha.com"
      },
      "Ezoic": {
        "cats": [
          10,
          36
        ],
        "description": "Ezoic is a website optimisation platform for digital publishers and website owners powered by machine learning.",
        "icon": "Ezoic.svg",
        "js": {
          "EzoicA": "",
          "EzoicBanger": "",
          "ezoicTestActive": ""
        },
        "pricing": [
          "freemium",
          "recurring",
          "poa",
          "payg"
        ],
        "saas": true,
        "scripts": "\\.ezo(?:js|ic|dn)\\.(?:com|net)",
        "website": "https://www.ezoic.com"
      },
      "F5 BigIP": {
        "cats": [
          64
        ],
        "cookies": {
          "F5_HT_shrinked": "",
          "F5_ST": "",
          "F5_fullWT": "",
          "LastMRH_Session": "",
          "MRHSHint": "",
          "MRHSequence": "",
          "MRHSession": "",
          "TIN": ""
        },
        "description": "F5's BIG-IP is a family of products covering software and hardware designed around application availability, access control, and security solutions.",
        "headers": {
          "server": "^big-?ip$"
        },
        "icon": "F5.png",
        "website": "https://www.f5.com/products/big-ip-services"
      },
      "FAST ESP": {
        "cats": [
          29
        ],
        "description": "FAST ESP is a service-oriented architecture development platform which is geared towards production searchable indexes. It provided a flexible framework for creating ETL applications for efficient indexing of searchable content.",
        "html": "<form[^>]+id=\"fastsearch\"",
        "icon": "FAST ESP.png",
        "website": "http://microsoft.com/enterprisesearch"
      },
      "FAST Search for SharePoint": {
        "cats": [
          29
        ],
        "description": "FAST Search for SharePoint is the search engine for Microsoft's SharePoint collaboration platform. Its purpose is helping SharePoint users locate and retrieve stored enterprise content.",
        "html": "<input[^>]+ name=\"ParametricSearch",
        "icon": "FAST Search for SharePoint.png",
        "implies": [
          "Microsoft SharePoint",
          "Microsoft ASP.NET"
        ],
        "url": "Pages/SearchResults\\.aspx\\?k=",
        "website": "http://sharepoint.microsoft.com/en-us/product/capabilities/search/Pages/Fast-Search.aspx"
      },
      "Facebook": {
        "cats": [
          5
        ],
        "description": "Facebook is an online social media and social networking service.",
        "icon": "Facebook.svg",
        "scripts": "//connect\\.facebook\\.([a-z]+)/[^/]*/[a-z]*\\.js",
        "website": "http://facebook.com"
      },
      "Facebook Login": {
        "cats": [
          69
        ],
        "description": "Facebook Login is a way for people to create accounts and log into your app across multiple platforms.",
        "icon": "Facebook.svg",
        "js": {
          "FB.getLoginStatus": ""
        },
        "website": "https://developers.facebook.com/docs/facebook-login/"
      },
      "Facebook Pixel": {
        "cats": [
          10
        ],
        "description": "Facebook pixel is an analytics tool that allows you to measure the effectiveness of your advertising.",
        "dom": "img[src*='facebook.com/tr']",
        "icon": "Facebook.svg",
        "js": {
          "_fbq": ""
        },
        "website": "http://facebook.com"
      },
      "Facil-iti": {
        "cats": [
          68
        ],
        "description": "Facil-iti is solution for website owners, providing accessibility to disability users and seniors.",
        "icon": "Facil-iti.svg",
        "scripts": "ws\\.facil-iti\\.com/tag/faciliti-tag\\.min\\.js",
        "website": "https://www.facil-iti.com/"
      },
      "Fact Finder": {
        "cats": [
          29
        ],
        "description": "Fact Finder is a platform which, combines search, navigation, and merchandising solutions to streamline online search and power sales.",
        "html": "<!-- Factfinder",
        "icon": "Fact Finder.png",
        "scripts": "Suggest\\.ff",
        "url": "(?:/ViewParametricSearch|ffsuggest\\.[a-z]htm)",
        "website": "http://fact-finder.com"
      },
      "FancyBox": {
        "cats": [
          59
        ],
        "description": "FancyBox is a tool for displaying images, html content and multi-media in a Mac-style 'lightbox' that floats overtop of web page.",
        "icon": "FancyBox.png",
        "implies": "jQuery",
        "js": {
          "$.fancybox.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "jquery\\.fancybox(?:\\.pack|\\.min)?\\.js(?:\\?v=([\\d.]+))?$\\;version:\\1",
        "website": "http://fancyapps.com/fancybox"
      },
      "FaraPy": {
        "cats": [
          1
        ],
        "html": "<!-- Powered by FaraPy.",
        "icon": "FaraPy.png",
        "implies": "Python",
        "website": "https://faral.tech"
      },
      "FareHarbor": {
        "cats": [
          5,
          72
        ],
        "description": "FareHarbor is an booking and schedule management solution intended for tour and activity companies.",
        "html": "<iframe[^>]+fareharbor",
        "icon": "FareHarbor.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "fareharbor\\.com/embeds/api/",
        "website": "https://fareharbor.com"
      },
      "Fastcommerce": {
        "cats": [
          6
        ],
        "icon": "Fastcommerce.png",
        "meta": {
          "generator": "^Fastcommerce"
        },
        "website": "https://www.fastcommerce.com.br"
      },
      "Fastly": {
        "cats": [
          31
        ],
        "description": "Fastly is a cloud computing services provider. Fastly's cloud platform provides a content delivery network, Internet security services, load balancing, and video & streaming services.",
        "headers": {
          "Fastly-Debug-Digest": "",
          "Vary": "Fastly-SSL",
          "X-Fastly-Request-ID": "",
          "x-fastly-origin": "",
          "x-via-fastly:": ""
        },
        "icon": "Fastly.svg",
        "pricing": [
          "payg"
        ],
        "website": "https://www.fastly.com",
        "xhr": "\\.fastly\\.net"
      },
      "Fastspring": {
        "cats": [
          6
        ],
        "html": [
          "<a [^>]*href=\"https?://sites\\.fastspring\\.com",
          "<form [^>]*action=\"https?://sites\\.fastspring\\.com"
        ],
        "icon": "fastspring.png",
        "website": "https://fastspring.com"
      },
      "Fat Zebra": {
        "cats": [
          41
        ],
        "description": "Fat Zebra provides a process of accepting credit card payments online.",
        "html": [
          "<(?:iframe|img|form)[^>]+paynow\\.pmnts\\.io",
          "<(?:iframe)[^>]+FatZebraFrame"
        ],
        "icon": "fatzebra.png",
        "scripts": "paynow\\.pmnts\\.io",
        "website": "https://www.fatzebra.com/"
      },
      "Fat-Free Framework": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "^Fat-Free Framework$"
        },
        "icon": "Fat-Free Framework.png",
        "implies": "PHP",
        "website": "http://fatfreeframework.com"
      },
      "Fbits": {
        "cats": [
          6
        ],
        "icon": "Fbits.png",
        "js": {
          "fbits": ""
        },
        "website": "https://www.traycorp.com.br"
      },
      "Fedora": {
        "cats": [
          28
        ],
        "cpe": "cpe:/o:fedoraproject:fedora",
        "description": "Fedora is a free open-source Linux-based operating system.",
        "headers": {
          "Server": "Fedora"
        },
        "icon": "Fedora.png",
        "website": "http://fedoraproject.org"
      },
      "Feedback Fish": {
        "cats": [
          13
        ],
        "description": "Feedback Fish is a widget for collecting website feedback from users.",
        "icon": "feedback-fish.svg",
        "scripts": "^https://feedback\\.fish/ff\\.js",
        "website": "https://feedback.fish"
      },
      "FilePond": {
        "cats": [
          59
        ],
        "description": "FilePond is a JavaScript library for file uploads.",
        "icon": "filepond.svg",
        "js": {
          "FilePond": "",
          "FilePond.create": "",
          "FilePond.setOptions": ""
        },
        "scripts": "filepond.js",
        "website": "https://pqina.nl/filepond/"
      },
      "FinanceAds": {
        "cats": [
          71
        ],
        "description": "FinanceAds is a performance marketing agency that has grown affiliate network for the financial sector.",
        "dom": "link[href*='js.financeads.net'], link[href*='js.financeads.com'], a[href*='www.financeads.net/tc.php']",
        "icon": "Financeads.png",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://www.financeads.com"
      },
      "FingerprintJS": {
        "cats": [
          59,
          83
        ],
        "description": "FingerprintJS is a browser fingerprinting library that queries browser attributes and computes a hashed visitor identifier from them.",
        "icon": "FingerprintJS.svg",
        "js": {
          "Fingerprint": "(\\d)?$\\;version:\\1",
          "Fingerprint2": "",
          "Fingerprint2.VERSION": "^(.+)$\\;version:\\1",
          "FingerprintJS": ""
        },
        "oss": true,
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": [
          "fingerprint(?:/fp)?(\\d)?(?:\\.min)?\\.js\\;version:\\1",
          "fingerprintjs(?:\\-pro|2)?(?:@|/)(\\d.*?)?/\\;version:\\1"
        ],
        "website": "https://fingerprintjs.com"
      },
      "Firebase": {
        "cats": [
          34
        ],
        "description": "Firebase is a Backend-as-a-Service (Baas) for creating mobile and web applications.",
        "icon": "Firebase.png",
        "js": {
          "firebase.SDK_VERSION": "([\\d.]+)$\\;version:\\1"
        },
        "scripts": [
          "/(?:([\\d.]+)/)?firebase(?:\\.min)?\\.js\\;version:\\1",
          "/firebasejs/([\\d.]+)/firebase\\;version:\\1"
        ],
        "website": "https://firebase.com"
      },
      "Fireblade": {
        "cats": [
          31
        ],
        "headers": {
          "Server": "fbs"
        },
        "icon": "Fireblade.png",
        "website": "http://fireblade.com"
      },
      "Firepush": {
        "cats": [
          32
        ],
        "description": "Firepush is an omnichannel marketing app that helps Shopify stores to drive sales with automated web push, email and SMS campaigns.",
        "dom": "link[href*='cdn.firepush.net']",
        "icon": "Firepush.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.firepush\\.\\w+",
        "website": "https://getfirepush.com"
      },
      "FirstImpression.io": {
        "cats": [
          36
        ],
        "description": "FirstImpression.io is a customized ad placements for publisher websites on their platform, with zero technical work.",
        "icon": "FirstImpression.io.png",
        "js": {
          "FI.options": "",
          "fiPrebidAnalyticsHandler": ""
        },
        "scripts": "\\.firstimpression\\.io",
        "website": "https://www.firstimpression.io",
        "xhr": "\\.firstimpression\\.io"
      },
      "Fit Analytics": {
        "cats": [
          5,
          76
        ],
        "description": "Fit Analytics is a platform that provides clothing size recommendations for online customers by measuring individual dimensions via webcams.",
        "icon": "Fit Analytics.png",
        "js": {
          "FitAnalyticsWidget": "",
          "_fitAnalytics": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.fitanalytics\\.com",
        "website": "https://www.fitanalytics.com"
      },
      "Flarum": {
        "cats": [
          2
        ],
        "cpe": "cpe:/a:flarum:flarum",
        "description": "Flarum is a discussion platform.",
        "html": "<div id=\"flarum-loading\"",
        "icon": "flarum.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "app.cache.discussionList": "",
          "app.forum.freshness": ""
        },
        "website": "http://flarum.org/"
      },
      "Flask": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "Server": "Werkzeug/?([\\d\\.]+)?\\;version:\\1"
        },
        "icon": "Flask.png",
        "implies": "Python",
        "website": "http://flask.pocoo.org"
      },
      "Flat UI": {
        "cats": [
          66
        ],
        "html": "<link[^>]* href=[^>]+flat-ui(?:\\.min)?\\.css",
        "icon": "Flat UI.png",
        "implies": "Bootstrap",
        "website": "https://designmodo.github.io/Flat-UI/"
      },
      "FlexCMP": {
        "cats": [
          1
        ],
        "headers": {
          "X-Flex-Lang": "",
          "X-Powered-By": "FlexCMP.+\\[v\\. ([\\d.]+)\\;version:\\1"
        },
        "html": "<!--[^>]+FlexCMP[^>v]+v\\. ([\\d.]+)\\;version:\\1",
        "icon": "FlexCMP.png",
        "meta": {
          "generator": "^FlexCMP"
        },
        "website": "http://www.flexcmp.com/cms/home"
      },
      "FlexSlider": {
        "cats": [
          5
        ],
        "description": "FlexSlider is a free jQuery slider plugin.",
        "icon": "FlexSlider.png",
        "implies": "jQuery",
        "scripts": [
          "jquery\\.flexslider(?:\\.min)?\\.js$"
        ],
        "website": "https://woocommerce.com/flexslider/"
      },
      "Flickity": {
        "cats": [
          59
        ],
        "js": {
          "Flickity": ""
        },
        "scripts": "/flickity(?:\\.pkgd)?(?:\\.min)?\\.js",
        "website": "https://flickity.metafizzy.co/"
      },
      "Flow": {
        "cats": [
          6
        ],
        "icon": "flow.png",
        "scripts": [
          "shopify-cdn\\.flow\\.io",
          "flow\\.min\\.js"
        ],
        "website": "https://www.flow.io/"
      },
      "FluxBB": {
        "cats": [
          2
        ],
        "cpe": "cpe:/a:fluxbb:fluxbb",
        "html": "<p id=\"poweredby\">[^<]+<a href=\"https?://fluxbb\\.org/\">",
        "icon": "FluxBB.png",
        "implies": "PHP",
        "website": "https://fluxbb.org"
      },
      "Flyspray": {
        "cats": [
          13
        ],
        "cookies": {
          "flyspray_project": ""
        },
        "html": "(?:<a[^>]+>Powered by Flyspray|<map id=\"projectsearchform)",
        "icon": "Flyspray.png",
        "implies": "PHP",
        "website": "http://flyspray.org"
      },
      "Flywheel": {
        "cats": [
          62
        ],
        "headers": {
          "x-fw-hash": "",
          "x-fw-serve": "",
          "x-fw-server": "^Flywheel(?:/([\\d.]+))?\\;version:\\1",
          "x-fw-static": "",
          "x-fw-type": ""
        },
        "icon": "flywheel.svg",
        "implies": "WordPress",
        "website": "https://getflywheel.com"
      },
      "Fomo": {
        "cats": [
          32
        ],
        "description": "Fomo is a social proof marketing platform.",
        "icon": "Fomo.png",
        "js": {
          "fomo.version": "(.+)\\;version:\\1"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "fomo\\.com/api/v",
        "website": "https://fomo.com"
      },
      "Font Awesome": {
        "cats": [
          17
        ],
        "description": "Font Awesome is a font and icon toolkit based on CSS and Less.",
        "html": [
          "<link[^>]* href=[^>]+(?:([\\d.]+)/)?(?:css/)?font-awesome(?:\\.min)?\\.css\\;version:\\1",
          "<link[^>]* href=[^>]*kit\\-pro\\.fontawesome\\.com/releases/v([0-9.]+)/\\;version:\\1"
        ],
        "icon": "font-awesome.svg",
        "js": {
          "FontAwesomeCdnConfig": "",
          "___FONT_AWESOME___": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "scripts": [
          "(?:F|f)o(?:n|r)t-?(?:A|a)wesome(?:.*?([0-9a-fA-F]{7,40}|[\\d]+(?:.[\\d]+(?:.[\\d]+)?)?)|)",
          "\\.fontawesome\\.com/([0-9a-z]+).js"
        ],
        "website": "https://fontawesome.com/"
      },
      "Fork CMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:fork-cms:fork_cms",
        "icon": "ForkCMS.png",
        "implies": "Symfony",
        "meta": {
          "generator": "^Fork CMS$"
        },
        "website": "http://www.fork-cms.com/"
      },
      "Formitable": {
        "cats": [
          5,
          6,
          72
        ],
        "description": "Formitable is an reservation management system for restaurants.",
        "icon": "Formitable.svg",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "formitable\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
          "cdn\\.formitable\\.com"
        ],
        "website": "https://formitable.com"
      },
      "ForoshGostar": {
        "cats": [
          6
        ],
        "cookies": {
          "Aws.customer": ""
        },
        "icon": "ForoshGostar.svg",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "^Forosh\\s?Gostar.*|Arsina Webshop.*$"
        },
        "website": "https://www.foroshgostar.com"
      },
      "Forte": {
        "cats": [
          41
        ],
        "description": "Forte, a CSG Company offers merchants and partners a broad range of payment solutions.",
        "icon": "Forte.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "checkout\\.forte\\.net",
        "website": "https://www.forte.net"
      },
      "Forter": {
        "cats": [
          6,
          16
        ],
        "cookies": {
          "forterToken": ""
        },
        "description": "Forter is a SaaS company that provides fraud prevention technology for online retailers and marketplaces.",
        "icon": "Forter.svg",
        "js": {
          "ftr__startScriptLoad": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "forter\\.com",
        "website": "https://www.forter.com"
      },
      "Fortune3": {
        "cats": [
          6
        ],
        "html": "(?:<link [^>]*href=\"[^\\/]*\\/\\/www\\.fortune3\\.com\\/[^\"]*siterate\\/rate\\.css|Powered by <a [^>]*href=\"[^\"]+fortune3\\.com)",
        "icon": "Fortune3.png",
        "scripts": "cartjs\\.php\\?(?:.*&)?s=[^&]*myfortune3cart\\.com",
        "website": "http://fortune3.com"
      },
      "Foswiki": {
        "cats": [
          8
        ],
        "cookies": {
          "FOSWIKISTRIKEONE": "",
          "SFOSWIKISID": ""
        },
        "cpe": "cpe:/a:foswiki:foswiki",
        "description": "Foswiki is a free open-source collaboration platform.",
        "headers": {
          "X-Foswikiaction": "",
          "X-Foswikiuri": ""
        },
        "html": [
          "<div class=\"foswiki(?:Copyright|Page|Main)\">"
        ],
        "icon": "foswiki.png",
        "implies": "Perl",
        "js": {
          "foswiki": ""
        },
        "meta": {
          "foswiki.SERVERTIME": "",
          "foswiki.WIKINAME": ""
        },
        "website": "http://foswiki.org"
      },
      "FreeBSD": {
        "cats": [
          28
        ],
        "cpe": "cpe:/o:freebsd:freebsd",
        "description": "FreeBSD is a free and open-source Unix-like operating system.",
        "headers": {
          "Server": "FreeBSD(?: ([\\d.]+))?\\;version:\\1"
        },
        "icon": "FreeBSD.png",
        "website": "http://freebsd.org"
      },
      "FreeTextBox": {
        "cats": [
          24
        ],
        "description": "FreeTextBox is a free open-source HTML Editor.",
        "html": "<!-- \\* FreeTextBox v\\d \\((\\d+\\.\\d+\\.\\d+)\\;version:\\1",
        "icon": "FreeTextBox.png",
        "implies": "Microsoft ASP.NET",
        "js": {
          "FTB_API": "",
          "FTB_AddEvent": ""
        },
        "website": "http://freetextbox.com"
      },
      "Freespee": {
        "cats": [
          10
        ],
        "icon": "Freespee.svg",
        "scripts": "analytics\\.freespee\\.com/js/external/fs\\.(?:min\\.)?js",
        "website": "https://www.freespee.com"
      },
      "Freshchat": {
        "cats": [
          52
        ],
        "description": "Freshchat is a cloud-hosted live messaging and engagement application.",
        "icon": "Freshchat.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "wchat\\.freshchat\\.com/js/widget\\.js",
        "website": "https://www.freshworks.com/live-chat-software/"
      },
      "Freshworks CRM": {
        "cats": [
          53,
          32,
          74
        ],
        "description": "Freshworks CRM is a cloud-based customer relationship management (CRM) solution. Key features include one-click phone, sales lead tracking, sales management, event tracking and more.",
        "icon": "Freshworks CRM.svg",
        "js": {
          "ZargetForm": "",
          "zarget": "",
          "zargetAPI": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.freshmarketer\\.com",
          "cdn\\.zarget\\.com"
        ],
        "website": "https://www.freshworks.com/crm"
      },
      "Froala Editor": {
        "cats": [
          24
        ],
        "description": "Froala Editor is a WYSIWYG HTML Editor written in Javascript that enables rich text editing capabilities for applications.",
        "html": "<[^>]+class=\"[^\"]*(?:fr-view|fr-box)",
        "icon": "Froala.svg",
        "implies": [
          "jQuery",
          "Font Awesome"
        ],
        "website": "http://froala.com/wysiwyg-editor"
      },
      "FrontPage": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:microsoft:frontpage",
        "description": "FrontPage is a program for developing and maintaining websites.",
        "icon": "FrontPage.png",
        "meta": {
          "ProgId": "^FrontPage\\.",
          "generator": "Microsoft FrontPage(?:\\s((?:Express )?[\\d.]+))?\\;version:\\1"
        },
        "website": "http://office.microsoft.com/frontpage"
      },
      "Frontity": {
        "cats": [
          12,
          18
        ],
        "description": "Frontity is a React open-source framework focused on WordPress.",
        "icon": "frontity.png",
        "implies": [
          "React",
          "webpack",
          "WordPress"
        ],
        "meta": {
          "generator": "^Frontity"
        },
        "website": "https://frontity.org"
      },
      "Frosmo": {
        "cats": [
          32,
          74
        ],
        "description": "Frosmo is a SaaS company which provides AI-driven personalisation products.",
        "icon": "Frosmo.png",
        "js": {
          "_frosmo": "",
          "frosmo": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "frosmo\\.easy\\.js",
        "website": "https://frosmo.com"
      },
      "Fusion Ads": {
        "cats": [
          36
        ],
        "icon": "Fusion Ads.png",
        "js": {
          "_fusion": ""
        },
        "scripts": "^[^\\/]*//[ac]dn\\.fusionads\\.net/(?:api/([\\d.]+)/)?\\;version:\\1",
        "website": "http://fusionads.net"
      },
      "Future Shop": {
        "cats": [
          6
        ],
        "icon": "futureshop.png",
        "scripts": "future-shop.*\\.js",
        "website": "https://www.future-shop.jp"
      },
      "GOV.UK Elements": {
        "cats": [
          66
        ],
        "html": [
          "<link[^>]+elements-page[^>\"]+css\\;confidence:25",
          "<div[^>]+phase-banner-alpha\\;confidence:25",
          "<div[^>]+phase-banner-beta\\;confidence:25",
          "<div[^>]+govuk-box-highlight\\;confidence:25"
        ],
        "icon": "govuk.png",
        "implies": "GOV.UK Toolkit",
        "website": "https://github.com/alphagov/govuk_elements/"
      },
      "GOV.UK Frontend": {
        "cats": [
          66
        ],
        "html": [
          "<link[^>]* href=[^>]*?govuk-frontend(?:[^>]*?([0-9a-fA-F]{7,40}|[\\d]+(?:.[\\d]+(?:.[\\d]+)?)?)|)[^>]*?(?:\\.min)?\\.css\\;version:\\1",
          "<body[^>]+govuk-template__body\\;confidence:80",
          "<a[^>]+govuk-link\\;confidence:10"
        ],
        "icon": "govuk.png",
        "js": {
          "GOVUKFrontend": ""
        },
        "scripts": [
          "govuk-frontend(?:[^>]*?([0-9a-fA-F]{7,40}|[\\d]+(?:.[\\d]+(?:.[\\d]+)?)?)|)[^>]*?(?:\\.min)?\\.js\\;version:\\1"
        ],
        "website": "https://design-system.service.gov.uk/"
      },
      "GOV.UK Template": {
        "cats": [
          66
        ],
        "html": [
          "<link[^>]+govuk-template[^>\"]+css",
          "<link[^>]+govuk-template-print[^>\"]+css",
          "<link[^>]+govuk-template-ie6[^>\"]+css",
          "<link[^>]+govuk-template-ie7[^>\"]+css",
          "<link[^>]+govuk-template-ie8[^>\"]+css"
        ],
        "icon": "govuk.png",
        "js": {
          "GOVUK": ""
        },
        "scripts": [
          "govuk-template\\.js"
        ],
        "website": "https://github.com/alphagov/govuk_template/"
      },
      "GOV.UK Toolkit": {
        "cats": [
          66
        ],
        "icon": "govuk.png",
        "js": {
          "GOVUK.details": "",
          "GOVUK.modules": "",
          "GOVUK.primaryLinks": ""
        },
        "website": "https://github.com/alphagov/govuk_frontend_toolkit"
      },
      "GSAP": {
        "cats": [
          12
        ],
        "description": "GSAP is an animation library that allows you to create animations with JavaScript.",
        "icon": "TweenMax.png",
        "js": {
          "TweenLite.version": "([\\d.]+)\\;version:\\1",
          "TweenMax.version": "([\\d.]+)\\;version:\\1",
          "gsap.version": "([\\d.]+)\\;version:\\1",
          "gsapVersions": ""
        },
        "scripts": "TweenMax(?:\\.min)?\\.js",
        "website": "https://greensock.com/gsap"
      },
      "GX WebManager": {
        "cats": [
          1
        ],
        "html": "<!--\\s+Powered by GX",
        "icon": "GX WebManager.png",
        "meta": {
          "generator": "GX WebManager(?: ([\\d.]+))?\\;version:\\1"
        },
        "website": "http://www.gxsoftware.com/en/products/web-content-management.htm"
      },
      "Gallery": {
        "cats": [
          7
        ],
        "description": "Gallery is an open-source web based photo album organiser.",
        "html": [
          "<div id=\"gsNavBar\" class=\"gcBorder1\">",
          "<a href=\"http://gallery\\.sourceforge\\.net\"><img[^>]+Powered by Gallery\\s*(?:(?:v|Version)\\s*([0-9.]+))?\\;version:\\1"
        ],
        "icon": "Gallery.png",
        "js": {
          "$.fn.gallery_valign": "",
          "galleryAuthToken": ""
        },
        "website": "http://galleryproject.org/"
      },
      "Gambio": {
        "cats": [
          6
        ],
        "html": "(?:<link[^>]* href=\"templates/gambio/|<a[^>]content\\.php\\?coID=\\d|<!-- gambio eof -->|<!--[\\s=]+Shopsoftware by Gambio GmbH \\(c\\))",
        "icon": "Gambio.png",
        "implies": "PHP",
        "js": {
          "gambio": ""
        },
        "scripts": "gm_javascript\\.js\\.php",
        "website": "http://gambio.de"
      },
      "Gatsby": {
        "cats": [
          57,
          12
        ],
        "description": "Gatsby is a React-based open-source framework with performance, scalability and security built-in.",
        "html": [
          "<div id=\"___gatsby\">",
          "<style id=\"gatsby-inlined-css\">"
        ],
        "icon": "Gatsby.svg",
        "implies": [
          "React",
          "webpack"
        ],
        "meta": {
          "generator": "^Gatsby(?: ([0-9.]+))?$\\;version:\\1"
        },
        "website": "https://www.gatsbyjs.org/"
      },
      "Gauges": {
        "cats": [
          10
        ],
        "cookies": {
          "_gauges_": ""
        },
        "icon": "Gauges.png",
        "js": {
          "_gauges": ""
        },
        "website": "https://get.gaug.es"
      },
      "Gemius": {
        "cats": [
          10
        ],
        "html": "<a [^>]*onclick=\"gemius_hit",
        "icon": "Gemius.png",
        "js": {
          "gemius_hit": "",
          "gemius_init": "",
          "gemius_pending": "",
          "pp_gemius_hit": ""
        },
        "scripts": [
          "hit\\.gemius\\.pl/xgemius\\.js",
          "hit\\.gemius\\.pl\\;confidence:80",
          "xgemius\\.js\\;confidence:30"
        ],
        "website": "https://www.gemius.com"
      },
      "GeneXus": {
        "cats": [
          27
        ],
        "html": [
          "<link[^>]+?id=\"gxtheme_css_reference\""
        ],
        "icon": "GeneXus.png",
        "js": {
          "gx.gxVersion": "^(.+)-.*$\\;version:\\1"
        },
        "scripts": [
          "/static/gxgral\\.js",
          "/static/gxtimezone\\.js"
        ],
        "website": "https://www.genexus.com/"
      },
      "Genesis theme": {
        "cats": [
          80
        ],
        "description": "Genesis theme is a WordPress theme that has been developed using the Genesis Framework from Studiopress.",
        "icon": "Genesis theme.svg",
        "js": {
          "genesisBlocksShare": "",
          "genesis_responsive_menu": ""
        },
        "requires": "WordPress",
        "scripts": "/wp-content/themes/genesis/lib/js/",
        "website": "https://www.studiopress.com/themes/genesis"
      },
      "Genesys Cloud": {
        "cats": [
          32,
          5,
          75
        ],
        "description": "Genesys Cloud is an all-in-one cloud-based contact center software built to improve the customer experience.",
        "icon": "Genesys Cloud.svg",
        "js": {
          "PURECLOUD_WEBCHAT_FRAME_CONFIG": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "apps\\.mypurecloud\\.\\w+/widgets/([\\d.]+)\\;version:\\1",
          "apps\\.mypurecloud\\.\\w+"
        ],
        "website": "https://www.genesys.com"
      },
      "Gentoo": {
        "cats": [
          28
        ],
        "cpe": "cpe:/o:gentoo:linux",
        "description": "Gentoo is a free operating system based on Linux.",
        "headers": {
          "X-Powered-By": "gentoo"
        },
        "icon": "Gentoo.png",
        "website": "http://www.gentoo.org"
      },
      "Gerrit": {
        "cats": [
          47
        ],
        "html": [
          ">Gerrit Code Review</a>\\s*\"\\s*\\(([0-9.]+)\\)\\;version:\\1",
          "<(?:div|style) id=\"gerrit_"
        ],
        "icon": "gerrit.svg",
        "implies": [
          "Java",
          "git"
        ],
        "js": {
          "Gerrit": "",
          "gerrit_ui": ""
        },
        "meta": {
          "title": "^Gerrit Code Review$"
        },
        "scripts": "^gerrit_ui/gerrit_ui",
        "website": "http://www.gerritcodereview.com"
      },
      "Get Satisfaction": {
        "cats": [
          13
        ],
        "icon": "Get Satisfaction.png",
        "js": {
          "GSFN": ""
        },
        "website": "https://getsatisfaction.com/corp/"
      },
      "GetSimple CMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:get-simple:getsimple_cms",
        "icon": "GetSimple CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "GetSimple"
        },
        "website": "http://get-simple.info"
      },
      "GetSocial": {
        "cats": [
          69,
          10
        ],
        "description": "GetSocial is a social analytics and publishing platform.",
        "icon": "GetSocial.png",
        "js": {
          "GETSOCIAL_VERSION": "(.+)\\;version:\\1"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "api\\.at\\.getsocial\\.io",
        "website": "https://getsocial.io"
      },
      "Getintent": {
        "cats": [
          36
        ],
        "description": "Getintent is an adtech company that offers AI-powered programmatic suite for agencies, publishers, broadcasters and content owners.",
        "icon": "Getintent.png",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://getintent.com",
        "xhr": "\\.adhigh\\.net"
      },
      "Ghost": {
        "cats": [
          1,
          11
        ],
        "description": "Ghost is a free and open-source blogging platform written in JavaScript, designed to simplify the process of online publishing for individual bloggers as well as online publications.",
        "headers": {
          "X-Ghost-Cache-Status": ""
        },
        "icon": "Ghost.png",
        "implies": "Node.js",
        "meta": {
          "generator": "Ghost(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "website": "http://ghost.org"
      },
      "GitBook": {
        "cats": [
          4
        ],
        "description": "GitBook is a command-line tool for creating documentation using Git and Markdown.",
        "icon": "GitBook.png",
        "meta": {
          "generator": "GitBook ([\\d.]+)?\\;version:\\1"
        },
        "url": "^https?://[^/]+\\.gitbook\\.com/",
        "website": "https://www.gitbook.com"
      },
      "GitHub Pages": {
        "cats": [
          62
        ],
        "description": "GitHub Pages is a static site hosting service.",
        "headers": {
          "Server": "^GitHub\\.com$",
          "X-GitHub-Request-Id": ""
        },
        "icon": "GitHub.svg",
        "implies": "Ruby on Rails",
        "url": "^https?://[^/]+\\.github\\.io",
        "website": "https://pages.github.com/"
      },
      "GitLab": {
        "cats": [
          13,
          47
        ],
        "cookies": {
          "_gitlab_session": ""
        },
        "description": "GitLab is a web-based DevOps lifecycle tool that provides a Git-repository manager providing wiki, issue-tracking and continuous integration and deployment pipeline features, using an open-source license.",
        "html": [
          "<meta content=\"https?://[^/]+/assets/gitlab_logo-",
          "<header class=\"navbar navbar-fixed-top navbar-gitlab with-horizontal-nav\">"
        ],
        "icon": "GitLab.svg",
        "implies": [
          "Ruby on Rails",
          "Vue.js"
        ],
        "js": {
          "GitLab": "",
          "gl.dashboardOptions": ""
        },
        "meta": {
          "og:site_name": "^GitLab$"
        },
        "website": "https://about.gitlab.com"
      },
      "GitLab CI/CD": {
        "cats": [
          44,
          47
        ],
        "icon": "GitLab CI.png",
        "implies": "Ruby on Rails",
        "meta": {
          "description": "GitLab CI/CD is a tool built into GitLab for software development through continuous methodologies."
        },
        "website": "http://about.gitlab.com/gitlab-ci"
      },
      "Gitea": {
        "cats": [
          47
        ],
        "cookies": {
          "i_like_gitea": ""
        },
        "cpe": "cpe:/a:gitea:gitea",
        "description": "Gitea is an open-source forge software package for hosting software development version control using Git as well as other collaborative features like bug tracking, wikis and code review. It supports self-hosting but also provides a free public first-party instance hosted on DiDi's cloud.",
        "html": [
          "<div class=\"ui left\">\\n\\s+© Gitea Version: ([\\d.]+)\\;version:\\1"
        ],
        "icon": "gitea.svg",
        "meta": {
          "keywords": "^go,git,self-hosted,gitea$"
        },
        "website": "https://gitea.io"
      },
      "Gitiles": {
        "cats": [
          47
        ],
        "html": "Powered by <a href=\"https://gerrit\\.googlesource\\.com/gitiles/\">Gitiles<",
        "implies": [
          "Java",
          "git"
        ],
        "website": "http://gerrit.googlesource.com/gitiles/"
      },
      "GlassFish": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:oracle:glassfish_server",
        "headers": {
          "Server": "GlassFish(?: Server)?(?: Open Source Edition)?(?: ?/?([\\d.]+))?\\;version:\\1"
        },
        "icon": "GlassFish.png",
        "implies": "Java",
        "website": "http://glassfish.java.net"
      },
      "Global-e": {
        "cats": [
          6
        ],
        "cookies": {
          "GlobalE_CT_Data": "",
          "GlobalE_Data": "",
          "GlobalE_SupportThirdPartCookies": ""
        },
        "description": "Global-e is a provider of cross-border ecommerce solutions.",
        "icon": "Global-e.svg",
        "js": {
          "GLOBALE_ENGINE_CONFIG": "",
          "GlobalE": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.global-e\\.com",
        "website": "https://www.global-e.com"
      },
      "Glyphicons": {
        "cats": [
          17
        ],
        "description": "Glyphicons are icon fonts which you can use in your web projects.",
        "html": "(?:<link[^>]* href=[^>]+glyphicons(?:\\.min)?\\.css|<img[^>]* src=[^>]+glyphicons)",
        "icon": "Glyphicons.png",
        "website": "http://glyphicons.com"
      },
      "Go": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:golang:go",
        "icon": "Go.svg",
        "website": "https://golang.org"
      },
      "GoAhead": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:embedthis:goahead",
        "headers": {
          "Server": "GoAhead"
        },
        "icon": "GoAhead.png",
        "website": "http://embedthis.com/products/goahead/index.html"
      },
      "GoCache": {
        "cats": [
          31
        ],
        "description": "GoCache is an in-memory key:value store/cache similar to memcached that is suitable for applications running on a single machine.",
        "headers": {
          "Server": "^gocache$",
          "X-GoCache-CacheStatus": ""
        },
        "icon": "GoCache.png",
        "website": "https://www.gocache.com.br/"
      },
      "GoDaddy Website Builder": {
        "cats": [
          1
        ],
        "cookies": {
          "dps_site_id": ""
        },
        "icon": "godaddy.svg",
        "meta": {
          "generator": "Go Daddy Website Builder (.+)\\;version:\\1"
        },
        "pricing": [
          "mid"
        ],
        "saas": true,
        "website": "https://www.godaddy.com/websites/website-builder"
      },
      "GoJS": {
        "cats": [
          25
        ],
        "icon": "GoJS.png",
        "js": {
          "go.GraphObject": "",
          "go.version": "(.*)\\;version:\\1"
        },
        "website": "https://gojs.net/"
      },
      "GoStats": {
        "cats": [
          10
        ],
        "icon": "GoStats.png",
        "js": {
          "_goStatsRun": "",
          "_go_track_src": "",
          "go_msie": ""
        },
        "website": "http://gostats.com/"
      },
      "Gogs": {
        "cats": [
          47
        ],
        "cookies": {
          "i_like_gogits": ""
        },
        "cpe": "cpe:/a:gogs:gogs",
        "description": "Gogs is a self-hosted Git service written in Go.",
        "html": [
          "<div class=\"ui left\">\\n\\s+© \\d{4} Gogs Version: ([\\d.]+) Page:\\;version:\\1",
          "<button class=\"ui basic clone button\" id=\"repo-clone-ssh\" data-link=\"gogs@"
        ],
        "icon": "gogs.png",
        "meta": {
          "keywords": "go, git, self-hosted, gogs"
        },
        "scripts": "js/gogs\\.js",
        "website": "http://gogs.io"
      },
      "Google AdSense": {
        "cats": [
          36
        ],
        "description": "Google AdSense is a program run by Google through which website publishers serve advertisements that are targeted to the site content and audience.",
        "icon": "Google AdSense.svg",
        "js": {
          "Goog_AdSense_": "",
          "Goog_AdSense_OsdAdapter": "",
          "__google_ad_urls": "",
          "google_ad_": ""
        },
        "saas": true,
        "scripts": [
          "googlesyndication\\.com/",
          "ad\\.ca\\.doubleclick\\.net",
          "2mdn\\.net",
          "ad\\.ca\\.doubleclick\\.net"
        ],
        "website": "https://www.google.fr/adsense/start/"
      },
      "Google Ads": {
        "cats": [
          36
        ],
        "description": "Google Ads is an online advertising platform developed by Google.",
        "icon": "Google Ads.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://ads.google.com"
      },
      "Google Ads Conversion Tracking": {
        "cats": [
          10
        ],
        "description": "Google Ads Conversion Tracking is a free tool that shows you what happens after a customer interacts with your ads.",
        "icon": "Google.svg",
        "js": {
          "google_trackConversion": ""
        },
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "\\.googleadservices\\.com/pagead/conversion_async\\.js",
        "website": "https://support.google.com/google-ads/answer/1722022"
      },
      "Google Analytics": {
        "cats": [
          10
        ],
        "cookies": {
          "__utma": "",
          "_ga": "",
          "_gat": ""
        },
        "description": "Google Analytics is a free web analytics service that tracks and reports website traffic.",
        "html": "<amp-analytics [^>]*type=[\"']googleanalytics[\"']",
        "icon": "Google Analytics.svg",
        "js": {
          "GoogleAnalyticsObject": "",
          "gaGlobal": ""
        },
        "scripts": [
          "google-analytics\\.com/(?:ga|urchin|analytics)\\.js",
          "googletagmanager\\.com/gtag/js"
        ],
        "website": "http://google.com/analytics"
      },
      "Google Analytics Enhanced eCommerce": {
        "cats": [
          10
        ],
        "description": "Google analytics enhanced ecommerce is a plug-in which enables the measurement of user interactions with products on ecommerce websites.",
        "icon": "Google Analytics.svg",
        "implies": "Google Analytics",
        "js": {
          "gaplugins.EC": ""
        },
        "scripts": "google-analytics\\.com\\/plugins\\/ua\\/(?:ec|ecommerce)\\.js",
        "website": "https://developers.google.com/analytics/devguides/collection/analyticsjs/enhanced-ecommerce"
      },
      "Google App Engine": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Google Frontend"
        },
        "icon": "Google App Engine.svg",
        "website": "http://code.google.com/appengine"
      },
      "Google Call Conversion Tracking": {
        "cats": [
          10
        ],
        "description": "Google Call Conversion Tracking is conversion tracking that shows which search keywords are driving the most calls.",
        "icon": "Google.svg",
        "js": {
          "_googCallTrackingImpl": "",
          "google_wcc_status": ""
        },
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "gstatic\\.com/call-tracking/.+\\.js",
        "website": "https://support.google.com/google-ads/answer/6100664"
      },
      "Google Charts": {
        "cats": [
          25
        ],
        "description": "Google Charts is an interactive web service that creates graphical charts from user-supplied information.",
        "icon": "Google Charts.png",
        "js": {
          "__googleVisualizationAbstractRendererElementsCount__": "",
          "__gvizguard__": ""
        },
        "website": "http://developers.google.com/chart/"
      },
      "Google Cloud": {
        "cats": [
          31
        ],
        "cpe": "cpe:/a:google:cloud_platform",
        "description": "Google Cloud is a suite of cloud computing services.",
        "headers": {
          "Via": "^1\\.1 google$"
        },
        "icon": "google_cloud.svg",
        "website": "https://cloud.google.com"
      },
      "Google Cloud Storage": {
        "cats": [
          19
        ],
        "description": "Google Cloud Storage allows world-wide storage and retrieval of any amount of data at any time.",
        "headers": {
          "x-goog-storage-class": "^\\w+$"
        },
        "icon": "google-cloud-storage.svg",
        "implies": "Google Cloud",
        "website": "https://cloud.google.com/storage"
      },
      "Google Code Prettify": {
        "cats": [
          19
        ],
        "icon": "Google.svg",
        "js": {
          "prettyPrint": ""
        },
        "website": "http://code.google.com/p/google-code-prettify"
      },
      "Google Font API": {
        "cats": [
          17
        ],
        "description": "Google Font API is a web service that supports open-source font files that can be used on your web designs.",
        "html": "<link[^>]* href=[^>]+fonts\\.(?:googleapis|google)\\.com",
        "icon": "Google Font API.png",
        "js": {
          "WebFonts": ""
        },
        "scripts": "googleapis\\.com/.+webfont",
        "website": "http://google.com/fonts"
      },
      "Google Maps": {
        "cats": [
          35
        ],
        "description": "Google Maps is a web mapping service. It offers satellite imagery, aerial photography, street maps, 360° interactive panoramic views of streets, real-time traffic conditions, and route planning for traveling by foot, car, bicycle and air, or public transportation.",
        "icon": "Google Maps.svg",
        "scripts": [
          "(?:maps\\.google\\.com/maps\\?file=api(?:&v=([\\d.]+))?|maps\\.google\\.com/maps/api/staticmap)\\;version:API v\\1",
          "//maps\\.google(?:apis)?\\.com/maps/api/js"
        ],
        "website": "http://maps.google.com"
      },
      "Google My Business": {
        "cats": [
          1
        ],
        "icon": "Google.svg",
        "url": "https?://[^.]+\\.business\\.site",
        "website": "https://www.google.com/business/website-builder"
      },
      "Google Optimize": {
        "cats": [
          74
        ],
        "description": "Google Optimize allows you to test variants of web pages and see how they perform.",
        "icon": "Google Optimize.svg",
        "js": {
          "google_optimize": ""
        },
        "scripts": "googleoptimize\\.com/optimize\\.js",
        "website": "https://optimize.google.com"
      },
      "Google PageSpeed": {
        "cats": [
          23,
          33
        ],
        "description": "Google PageSpeed is a family of tools designed to help websites performance optimisations.",
        "headers": {
          "X-Mod-Pagespeed": "([\\d.]+)\\;version:\\1",
          "X-Page-Speed": "(.+)\\;version:\\1"
        },
        "icon": "Google PageSpeed.svg",
        "website": "http://developers.google.com/speed/pagespeed/mod"
      },
      "Google Pay": {
        "cats": [
          41
        ],
        "description": "Google Pay is a digital wallet platform and online payment system developed by Google to power in-app and tap-to-pay purchases on mobile devices, enabling users to make payments with Android phones, tablets or watches.",
        "dom": "[aria-labelledby='pi-google_pay'], ul[data-shopify-buttoncontainer] li",
        "icon": "Google.svg",
        "scripts": [
          "pay\\.google\\.com/([a-z/]+)/pay\\.js"
        ],
        "website": "https://pay.google.com"
      },
      "Google Plus": {
        "cats": [
          5
        ],
        "description": "Google Plus is a social network.",
        "icon": "Google Plus.svg",
        "scripts": "apis\\.google\\.com/js/[a-z]*\\.js",
        "website": "http://plus.google.com"
      },
      "Google Publisher Tag": {
        "cats": [
          36
        ],
        "description": "Google Publisher Tag (GPT) is an ad tagging library for Google Ad Manager which is used to dynamically build ad requests.",
        "icon": "Google Publisher Tag.svg",
        "scripts": [
          "googletagservices\\.com/tag/js/gpt\\.js",
          "securepubads\\.g\\.doubleclick\\.net/gpt"
        ],
        "website": "https://developers.google.com/publisher-tag/guides/get-started"
      },
      "Google Remarketing Tag": {
        "cats": [
          77
        ],
        "description": "Google Remarketing Tag is a web tagging library for Google's site measurement, conversion tracking and remarketing products.",
        "icon": "Google Tag Manager.svg",
        "js": {
          "gtag": ""
        },
        "saas": true,
        "website": "https://support.google.com/google-ads/answer/2476688"
      },
      "Google Sign-in": {
        "cats": [
          69
        ],
        "description": "Google Sign-In is a secure authentication system that reduces the burden of login for users, by enabling them to sign in with their Google account.",
        "dom": {
          "button": {
            "text": "(Sign|Log) in with Google"
          }
        },
        "html": [
          "<meta[^>]*google-signin-client_id",
          "<meta[^>]*google-signin-scope",
          "<iframe[^>]*accounts\\.google\\.com/o/oauth2",
          "<a[^>]*accounts\\.google\\.com/o/oauth2"
        ],
        "icon": "Google.svg",
        "scripts": [
          "apis\\.google\\.com/js/platform\\.js",
          "accounts\\.google\\.com/gsi/client"
        ],
        "website": "https://developers.google.com/identity/sign-in/web"
      },
      "Google Sites": {
        "cats": [
          1
        ],
        "icon": "Google Sites.png",
        "url": "^https?://sites\\.google\\.com",
        "website": "http://sites.google.com"
      },
      "Google Tag Manager": {
        "cats": [
          42
        ],
        "description": "Google Tag Manager is a tag management system (TMS) that allows you to quickly and easily update measurement codes and related code fragments collectively known as tags on your website or mobile app.",
        "html": [
          "googletagmanager\\.com/ns\\.html[^>]+></iframe>",
          "<!-- (?:End )?Google Tag Manager -->"
        ],
        "icon": "Google Tag Manager.svg",
        "js": {
          "google_tag_manager": "",
          "googletag": ""
        },
        "saas": true,
        "scripts": "googletagmanager\\.com/gtm\\.js",
        "website": "http://www.google.com/tagmanager"
      },
      "Google Wallet": {
        "cats": [
          41
        ],
        "icon": "Google Wallet.png",
        "saas": true,
        "scripts": [
          "checkout\\.google\\.com",
          "wallet\\.google\\.com"
        ],
        "website": "http://wallet.google.com"
      },
      "Google Web Server": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:google:web_server",
        "headers": {
          "Server": "gws"
        },
        "icon": "Google.svg",
        "website": "http://en.wikipedia.org/wiki/Google_Web_Server"
      },
      "Google Web Toolkit": {
        "cats": [
          18
        ],
        "cpe": "cpe:/a:google:web_toolkit",
        "description": "Google Web Toolkit (GWT) is an open-source Java software development framework that makes writing AJAX applications.",
        "icon": "Google Web Toolkit.png",
        "implies": "Java",
        "js": {
          "__gwt_": "",
          "__gwt_activeModules": "",
          "__gwt_getMetaProperty": "",
          "__gwt_isKnownPropertyValue": "",
          "__gwt_stylesLoaded": "",
          "__gwtlistener": ""
        },
        "meta": {
          "gwt:property": ""
        },
        "website": "http://developers.google.com/web-toolkit"
      },
      "Google Workspace": {
        "cats": [
          30,
          75
        ],
        "description": "Google Workspace, formerly G Suite, is a collection of cloud computing, productivity and collaboration tools.",
        "dns": {
          "MX": [
            "aspmx\\.l\\.google\\.com",
            "googlemail\\.com"
          ]
        },
        "icon": "Google Workspace.svg",
        "website": "https://workspace.google.com/"
      },
      "Graffiti CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "graffitibot": ""
        },
        "icon": "Graffiti CMS.png",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "Graffiti CMS ([^\"]+)\\;version:\\1"
        },
        "scripts": "/graffiti\\.js",
        "website": "http://graffiticms.codeplex.com"
      },
      "GrandNode": {
        "cats": [
          6
        ],
        "cookies": {
          "Grand.customer": ""
        },
        "html": "(?:<!--GrandNode |<a[^>]+grandnode - Powered by |Powered by: <a[^>]+nopcommerce)",
        "icon": "GrandNode.svg",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "grandnode"
        },
        "website": "https://grandnode.com"
      },
      "GraphCMS": {
        "cats": [
          1
        ],
        "description": "GraphCMS is a GraphQL headless CMS for content federation and omnichannel headless content management.",
        "dom": {
          ".graphcms-image-wrapper": {
            "text": ""
          },
          "img[src*='media.graphcms.com']": {
            "text": ""
          }
        },
        "icon": "GraphCMS.svg",
        "implies": [
          "GraphQL",
          "PostgreSQL",
          "Go",
          "TypeScript"
        ],
        "pricing": [
          "mid",
          "recurring",
          "freemium"
        ],
        "website": "https://graphcms.com",
        "xhr": "\\.graphcms\\.com"
      },
      "GraphQL": {
        "cats": [
          27
        ],
        "description": "GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data.",
        "icon": "GraphQL.svg",
        "oss": true,
        "website": "https://graphql.org"
      },
      "Grav": {
        "cats": [
          1
        ],
        "icon": "Grav.png",
        "implies": "PHP",
        "meta": {
          "generator": "GravCMS(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "website": "http://getgrav.org"
      },
      "Gravatar": {
        "cats": [
          19
        ],
        "description": "Gravatar is a service for providing globally unique avatars.",
        "html": "<[^>]+gravatar\\.com/avatar/",
        "icon": "Gravatar.png",
        "js": {
          "Gravatar": ""
        },
        "website": "http://gravatar.com"
      },
      "Gravitec": {
        "cats": [
          5
        ],
        "description": "Gravitec is a push notification tool.",
        "icon": "Gravitec.png",
        "js": {
          "Gravitec": "",
          "gravitecWebpackJsonp": ""
        },
        "pricing": [
          "freemium",
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "(?:cdn|id)\\.gravitec\\.(?:media|net)",
        "website": "https://gravitec.net"
      },
      "Gravity Forms": {
        "cats": [
          19
        ],
        "html": [
          "<div class=(?:\"|')[^>]*gform_wrapper",
          "<div class=(?:\"|')[^>]*gform_body",
          "<ul [^>]*class=(?:\"|')[^>]*gform_fields",
          "<link [^>]*href=(?:\"|')[^>]*wp-content/plugins/gravityforms/css/"
        ],
        "icon": "gravityforms.svg",
        "implies": "WordPress",
        "scripts": "/wp-content/plugins/gravityforms/js/[^/]+\\.js\\?ver=([\\d.]+)$\\;version:\\1",
        "website": "http://gravityforms.com"
      },
      "Green Valley CMS": {
        "cats": [
          1
        ],
        "html": "<img[^>]+/dsresource\\?objectid=",
        "icon": "Green Valley CMS.png",
        "implies": "Apache Tomcat",
        "meta": {
          "DC.identifier": "/content\\.jsp\\?objectid="
        },
        "website": "http://www.greenvalley.nl/Public/Producten/Content_Management/CMS"
      },
      "Griddo": {
        "cats": [
          1
        ],
        "description": "Griddo is an Martech Experience Platform for creating custom digital experiences.",
        "icon": "Griddo.svg",
        "implies": [
          "React",
          "Gatsby"
        ],
        "meta": {
          "generator": "^Griddo$"
        },
        "pricing": [
          "high",
          "poa"
        ],
        "saas": true,
        "website": "https://griddo.io"
      },
      "Gridsome": {
        "cats": [
          57
        ],
        "description": "Gridsome is a free and open-source Vue-powered static site generator for building static websites.",
        "icon": "Gridsome.svg",
        "implies": "Vue.js",
        "meta": {
          "generator": "^Gridsome v([\\d.]+)$\\;version:\\1"
        },
        "website": "https://gridsome.org"
      },
      "GrowingIO": {
        "cats": [
          10
        ],
        "cookies": {
          "gr_user_id": "",
          "grwng_uid": ""
        },
        "icon": "GrowingIO.png",
        "scripts": "assets\\.growingio\\.com/([\\d.]+)/gio\\.js\\;version:\\1",
        "website": "https://www.growingio.com/"
      },
      "Guestonline": {
        "cats": [
          5,
          72
        ],
        "description": "Guestonline is a restaurant table booking widget.",
        "dom": "iframe[src*='ib.guestonline.']",
        "icon": "Guestonline.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "ib\\.guestonline\\.\\w+",
        "website": "https://www.guestonline.io"
      },
      "GumGum": {
        "cats": [
          36
        ],
        "description": "GumGum is a technology and media company specializing in contextual intelligence.",
        "dom": "iframe[src*='gumgum.com'], img[src*='gumgum.com']",
        "icon": "GumGum.png",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://gumgum.com",
        "xhr": "\\.gumgum\\.com"
      },
      "HCL Commerce": {
        "cats": [
          6
        ],
        "html": "<(?:a|link|script)[^>]*(?:href|src)=\".*(?:/wcsstore/|webapp\\/wcs)",
        "icon": "HCL Commerce.svg",
        "implies": "Java",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "url": "/wcs/",
        "website": "https://www.hcltechsw.com/products/commerce"
      },
      "HCL Digital Experience": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:ibm:websphere_portal",
        "headers": {
          "IBM-Web2-Location": "",
          "Itx-Generated-Timestamp": ""
        },
        "icon": "IBM.svg",
        "implies": "Java",
        "url": "/wps/",
        "website": "https://www.hcltechsw.com/products/dx"
      },
      "HHVM": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:facebook:hhvm",
        "headers": {
          "X-Powered-By": "HHVM/?([\\d.]+)?\\;version:\\1"
        },
        "icon": "HHVM.png",
        "implies": "PHP\\;confidence:75",
        "website": "http://hhvm.com"
      },
      "HP Compact Server": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "HP_Compact_Server(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "HP.svg",
        "website": "http://hp.com"
      },
      "HP iLO": {
        "cats": [
          22,
          46
        ],
        "description": "HP iLO is a tool that provides multiple ways to configure, update, monitor, and run servers remotely.",
        "headers": {
          "Server": "HP-iLO-Server(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "HP.svg",
        "website": "http://hp.com"
      },
      "HTTP/2": {
        "cats": [
          19
        ],
        "description": "HTTP/2 (originally named HTTP/2.0) is a major revision of the HTTP network protocol used by the World Wide Web.",
        "excludes": "SPDY",
        "headers": {
          "X-Firefox-Spdy": "h2"
        },
        "icon": "http2.png",
        "website": "https://http2.github.io"
      },
      "Haddock": {
        "cats": [
          4
        ],
        "description": "Haddock is a tool for automatically generating documentation from annotated Haskell source code.",
        "html": "<p>Produced by <a href=\"http://www\\.haskell\\.org/haddock/\">Haddock</a> version ([0-9.]+)</p>\\;version:\\1",
        "scripts": "haddock-util\\.js",
        "website": "http://www.haskell.org/haddock/"
      },
      "Halo": {
        "cats": [
          1,
          11
        ],
        "icon": "Halo.png",
        "implies": "Java",
        "meta": {
          "generator": "Halo ([\\d.]+)?\\;version:\\1"
        },
        "website": "https://halo.run"
      },
      "Hammer.js": {
        "cats": [
          59
        ],
        "icon": "Hammer.js.png",
        "js": {
          "Ha.VERSION": "^(.+)$\\;version:\\1",
          "Hammer": "",
          "Hammer.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "hammer(?:\\.min)?\\.js",
        "website": "https://hammerjs.github.io"
      },
      "Handlebars": {
        "cats": [
          12
        ],
        "cpe": "cpe:/a:handlebars.js_project:handlebars.js",
        "description": "Handlebars is a JavaScript library used to create reusable webpage templates.",
        "html": "<[^>]*type=[^>]text\\/x-handlebars-template",
        "icon": "Handlebars.png",
        "js": {
          "Handlebars": "",
          "Handlebars.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "handlebars(?:\\.runtime)?(?:-v([\\d.]+?))?(?:\\.min)?\\.js\\;version:\\1",
        "website": "http://handlebarsjs.com"
      },
      "Haravan": {
        "cats": [
          6
        ],
        "icon": "Haravan.png",
        "js": {
          "Haravan": ""
        },
        "scripts": "haravan.*\\.js",
        "website": "https://www.haravan.com"
      },
      "Haskell": {
        "cats": [
          27
        ],
        "icon": "Haskell.png",
        "website": "http://wiki.haskell.org/Haskell"
      },
      "HeadJS": {
        "cats": [
          59
        ],
        "html": "<[^>]*data-headjs-load",
        "icon": "HeadJS.png",
        "js": {
          "head.browser.name": ""
        },
        "scripts": "head\\.(?:core|load)(?:\\.min)?\\.js",
        "website": "http://headjs.com"
      },
      "Heap": {
        "cats": [
          10
        ],
        "description": "Heap is an analytics platform.",
        "icon": "Heap.svg",
        "js": {
          "heap.version.heapJsVersion": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "freemium",
          "high"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.heapanalytics\\.com",
          "heap-\\d+\\.js"
        ],
        "website": "https://heap.io"
      },
      "Heartland Payment Systems": {
        "cats": [
          41
        ],
        "description": "Heartland Payment Systems is a US-based payment processing and technology provider.",
        "icon": "Heartland Payment Systems.svg",
        "pricing": [
          "payg",
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.heartlandportico\\.com",
        "website": "https://www.heartlandpaymentsystems.com"
      },
      "Helix Ultimate": {
        "cats": [
          18
        ],
        "description": "Helix Ultimate a free template framework for Joomla.",
        "dom": "header#sp-header, body.helix-ultimate",
        "icon": "HelixUltimate.svg",
        "implies": "Joomla",
        "oss": true,
        "website": "https://www.joomshaper.com/joomla-templates/helixultimate"
      },
      "Hello Bar": {
        "cats": [
          5
        ],
        "icon": "Hello Bar.png",
        "js": {
          "HelloBar": ""
        },
        "scripts": "hellobar\\.js",
        "website": "http://hellobar.com"
      },
      "HelpDocs": {
        "cats": [
          4,
          19
        ],
        "description": "HelpDocs is an knowledge management system.",
        "icon": "HelpDocs.svg",
        "js": {
          "HDAnalytics": "\\;confidence:25",
          "HDUtils": "\\;confidence:25",
          "hd_instant_search": "\\;confidence:50"
        },
        "pricing": [
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.helpdocs\\.io",
        "website": "https://www.helpdocs.io"
      },
      "Helpscout": {
        "cats": [
          52
        ],
        "description": "Helpscout is a customer service platform including email, a knowledge base tool and live chat.",
        "icon": "Helpscout.svg",
        "js": {
          "Beacon": "\\;confidence:25"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "beacon-v2\\.helpscout\\.net",
        "website": "https://www.helpscout.com/"
      },
      "Heroku": {
        "cats": [
          62
        ],
        "description": "Heroku is a cloud platform as a service (PaaS) supporting several programming languages.",
        "dns": {
          "TXT": [
            "heroku-domain-verification"
          ]
        },
        "headers": {
          "Via": "[\\d.-]+ vegur$"
        },
        "icon": "heroku.svg",
        "url": "\\.herokuapp\\.com",
        "website": "https://www.heroku.com/"
      },
      "Hexo": {
        "cats": [
          57
        ],
        "description": "Hexo is a blog framework powered by Node.js.",
        "html": [
          "Powered by <a href=\"https?://hexo\\.io/?\"[^>]*>Hexo</"
        ],
        "icon": "Hexo.png",
        "implies": "Node.js",
        "meta": {
          "generator": "Hexo(?: v?([\\d.]+))?\\;version:\\1"
        },
        "website": "https://hexo.io"
      },
      "HiConversion": {
        "cats": [
          74
        ],
        "description": "HiConversion is a SaaS solution that caters to ecommerce brands seeking actionable insights through their adaptive customer experience optimisation.",
        "icon": "HiConversion.png",
        "js": {
          "__hic.version": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "poa",
          "recurring"
        ],
        "saas": true,
        "scripts": "deploy\\.hiconversion\\.com",
        "website": "https://www.hiconversion.com"
      },
      "Hiawatha": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Hiawatha v([\\d.]+)\\;version:\\1"
        },
        "icon": "Hiawatha.png",
        "website": "http://hiawatha-webserver.org"
      },
      "Highcharts": {
        "cats": [
          25
        ],
        "description": "Highcharts is a charting library written in pure JavaScript, for adding interactive charts to a website or web application.",
        "html": "<svg[^>]*><desc>Created with Highcharts ([\\d.]*)\\;version:\\1",
        "icon": "Highcharts.png",
        "js": {
          "Highcharts": "",
          "Highcharts.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "highcharts.*\\.js",
        "website": "https://www.highcharts.com"
      },
      "Highlight.js": {
        "cats": [
          19
        ],
        "icon": "Highlight.js.png",
        "js": {
          "hljs.highlightBlock": "",
          "hljs.listLanguages": ""
        },
        "scripts": "/(?:([\\d.])+/)?highlight(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://highlightjs.org/"
      },
      "Highstock": {
        "cats": [
          25
        ],
        "html": "<svg[^>]*><desc>Created with Highstock ([\\d.]*)\\;version:\\1",
        "icon": "Highcharts.png",
        "scripts": "highstock[.-]?([\\d\\.]*\\d).*\\.js\\;version:\\1",
        "website": "http://highcharts.com/products/highstock"
      },
      "HikeOrders": {
        "cats": [
          68
        ],
        "description": "HikeOrders web accessibility automated plugin.",
        "icon": "HikeOrders.png",
        "scripts": "hikeorders\\.com/main/assets/js/hko-accessibility\\.min\\.js",
        "website": "https://hikeorders.com/"
      },
      "Hinza Advanced CMS": {
        "cats": [
          1,
          6
        ],
        "icon": "hinza_advanced_cms.svg",
        "implies": "Laravel",
        "meta": {
          "generator": "hinzacms"
        },
        "website": "http://hinzaco.com"
      },
      "History": {
        "cats": [
          19
        ],
        "description": "Manage session history with JavaScript",
        "scripts": [
          "/history(@|/)([\\d.]+)(?:/[a-z]+)?/history(?:(.production|.development))?(?:.min)?\\.js\\;version:\\2"
        ],
        "website": "https://github.com/ReactTraining/history"
      },
      "Hogan.js": {
        "cats": [
          12
        ],
        "icon": "Hogan.js.png",
        "js": {
          "Hogan": ""
        },
        "scripts": [
          "hogan-[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "([\\d.]+)/hogan(?:\\.min)?\\.js\\;version:\\1"
        ],
        "website": "https://twitter.github.io/hogan.js/"
      },
      "Hostmeapp": {
        "cats": [
          5,
          72
        ],
        "description": "Hostmeapp is an restaurant software. Includes reservation, waitlist, guestbook and marketing tools.",
        "icon": "Hostmeapp.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "tables\\.hostmeapp\\.com",
        "website": "https://www.hostmeapp.com"
      },
      "Hotaru CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "hotaru_mobile": ""
        },
        "icon": "Hotaru CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "Hotaru CMS"
        },
        "website": "http://hotarucms.org"
      },
      "Hotjar": {
        "cats": [
          10
        ],
        "description": "Hotjar is a suite of analytic tools to assist in the gathering of qualitative data, providing feedback through tools such as heatmaps, session recordings, and surveys.",
        "icon": "Hotjar.svg",
        "js": {
          "HotLeadfactory": "",
          "HotleadController": "",
          "hj.apiUrlBase": ""
        },
        "scripts": "//static\\.hotjar\\.com/",
        "website": "https://www.hotjar.com"
      },
      "Howler.js": {
        "cats": [
          59
        ],
        "description": "Howler.js is an audio library with support for the Web Audio API and a fallback mechanism for HTML5 Audio.",
        "icon": "Howler.js.svg",
        "js": {
          "Howler": "",
          "HowlerGlobal": ""
        },
        "oss": true,
        "saas": false,
        "scripts": [
          "howler@([\\d.]+)/dist/howler\\.min\\.js\\;version:\\1",
          "howler/([\\d.]+)/howler(?:\\.core)?\\.min\\.js\\;version:\\1"
        ],
        "website": "https://howlerjs.com"
      },
      "HubSpot": {
        "cats": [
          32
        ],
        "description": "HubSpot is a marketing and sales software that helps companies attract visitors, convert leads, and close customers.",
        "dns": {
          "TXT": [
            "hubspotemail\\.net"
          ]
        },
        "html": "<!-- Start of Async HubSpot",
        "icon": "HubSpot.png",
        "js": {
          "_hsq": "",
          "hubspot": ""
        },
        "website": "https://www.hubspot.com"
      },
      "HubSpot Analytics": {
        "cats": [
          10
        ],
        "description": "HubSpot is a marketing and sales software that helps companies attract visitors, convert leads, and close customers.",
        "icon": "HubSpot.png",
        "scripts": "js\\.hs-analytics\\.net/analytics",
        "website": "https://www.hubspot.com/products/marketing/analytics"
      },
      "HubSpot CMS Hub": {
        "cats": [
          1
        ],
        "description": "CMS Hub is a content management platform by HubSpot for marketers to manage, optimize, and track content performance on websites, blogs, and landing pages.",
        "headers": {
          "x-hs-hub-id": "",
          "x-powered-by": "HubSpot"
        },
        "icon": "HubSpot.png",
        "implies": "HubSpot",
        "meta": {
          "generator": "HubSpot"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.hubspot.com/products/cms"
      },
      "HubSpot Chat": {
        "cats": [
          52
        ],
        "description": "HubSpot Chat is a tool where you can view, manage, and reply to incoming messages from multiple channels.",
        "icon": "HubSpot.png",
        "js": {
          "HubSpotConversations": ""
        },
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "js\\.usemessages\\.com",
        "website": "https://www.hubspot.com/products/crm/live-chat"
      },
      "Hugo": {
        "cats": [
          57
        ],
        "description": "Hugo is an open-source static site generator written in Go.",
        "html": "powered by <a [^>]*href=\"http://hugo\\.spf13\\.com",
        "icon": "Hugo.png",
        "meta": {
          "generator": "Hugo ([\\d.]+)?\\;version:\\1"
        },
        "website": "http://gohugo.io"
      },
      "IBM Coremetrics": {
        "cats": [
          10
        ],
        "icon": "IBM.svg",
        "scripts": "cmdatatagutils\\.js",
        "website": "http://ibm.com/software/marketing-solutions/coremetrics"
      },
      "IBM DataPower": {
        "cats": [
          64
        ],
        "cpe": "cpe:/a:ibm:datapower_gateway",
        "description": "IBM DataPower Gateway is a single multi-channel gateway designed to help provide security, control, integration and optimized access to a full range of mobile, web, application programming interface (API), service-oriented architecture (SOA), B2B and cloud workloads.",
        "headers": {
          "X-Backside-Transport": "",
          "X-Global-Transaction-ID": ""
        },
        "icon": "DataPower.png",
        "website": "https://www.ibm.com/products/datapower-gateway"
      },
      "IBM HTTP Server": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:ibm:http_server",
        "headers": {
          "Server": "IBM_HTTP_Server(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "IBM.svg",
        "website": "http://ibm.com/software/webservers/httpservers"
      },
      "IIS": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:microsoft:internet_information_server",
        "description": "Internet Information Services (IIS) is an extensible web server software created by Microsoft for use with the Windows NT family.",
        "headers": {
          "Server": "^(?:Microsoft-)?IIS(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "IIS.png",
        "implies": "Windows Server",
        "website": "http://www.iis.net"
      },
      "INFOnline": {
        "cats": [
          10
        ],
        "icon": "INFOnline.png",
        "js": {
          "iam_data": "",
          "szmvars": ""
        },
        "scripts": "^https?://(?:[^/]+\\.)?i(?:oam|v)wbox\\.de/",
        "website": "https://www.infonline.de"
      },
      "IPB": {
        "cats": [
          2
        ],
        "cookies": {
          "ipbWWLmodpids": "",
          "ipbWWLsession_id": ""
        },
        "html": "<link[^>]+ipb_[^>]+\\.css",
        "icon": "IPB.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "IPBoard": "",
          "ipb_var": "",
          "ipsSettings": ""
        },
        "scripts": "jscripts/ips_",
        "website": "https://invisioncommunity.com/"
      },
      "IPInfoDB": {
        "cats": [
          79
        ],
        "description": "IPInfoDB is the API that returns the location of an IP address.",
        "icon": "IPInfoDB.png",
        "saas": true,
        "website": "https://www.ipinfodb.com/",
        "xhr": "api\\.ipinfodb\\.com"
      },
      "Ideasoft": {
        "cats": [
          6
        ],
        "description": "Ideasoft is a Turkish software company providing web-based software solutions, software design, ecommerce solutions, and other services.",
        "icon": "Ideasoft.png",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "\\.myideasoft\\.com/([\\d.]+)\\;version:\\1"
        ],
        "website": "https://www.ideasoft.com.tr"
      },
      "Identrust": {
        "cats": [
          70
        ],
        "certIssuer": "TrustID",
        "description": "denTrust is a digital identity authentication solution.",
        "icon": "Identrust.svg",
        "website": "https://www.identrust.com/"
      },
      "IdoSell Shop": {
        "cats": [
          6
        ],
        "icon": "idosellshop.png",
        "js": {
          "IAI_Ajax": ""
        },
        "website": "https://www.idosell.com"
      },
      "Immutable.js": {
        "cats": [
          59
        ],
        "icon": "Immutable.js.png",
        "js": {
          "Immutable": "",
          "Immutable.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "^immutable\\.(?:min\\.)?js$",
        "website": "https://facebook.github.io/immutable-js/"
      },
      "Impact": {
        "cats": [
          71
        ],
        "description": "Impact helps businesses contract and pay partners.",
        "icon": "Impact.svg",
        "js": {
          "ImpactRadiusEvent": "",
          "irEvent": ""
        },
        "scripts": "d\\.impactradius-event\\.com",
        "website": "https://impact.com/"
      },
      "Imperva": {
        "cats": [
          16
        ],
        "icon": "Imperva.svg",
        "scripts": [
          "/_Incapsula_Resource"
        ],
        "website": "https://www.imperva.com/"
      },
      "ImpressCMS": {
        "cats": [
          1
        ],
        "cookies": {
          "ICMSSession": "",
          "ImpressCMS": ""
        },
        "cpe": "cpe:/a:impresscms:impresscms",
        "icon": "ImpressCMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "ImpressCMS"
        },
        "scripts": "include/linkexternal\\.js",
        "website": "http://www.impresscms.org"
      },
      "ImpressPages": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:impresspages:impresspages_cms",
        "icon": "ImpressPages.png",
        "implies": "PHP",
        "meta": {
          "generator": "ImpressPages(?: CMS)?( [\\d.]*)?\\;version:\\1"
        },
        "website": "http://impresspages.org"
      },
      "Incapsula": {
        "cats": [
          31
        ],
        "description": "Incapsula is a cloud-based application delivery platform. It uses a global content delivery network to provide web application security, DDoS mitigation, content caching, application delivery, load balancing and failover services.",
        "headers": {
          "X-CDN": "Incapsula"
        },
        "icon": "Incapsula.png",
        "website": "http://www.incapsula.com"
      },
      "Includable": {
        "cats": [
          18
        ],
        "headers": {
          "X-Includable-Version": ""
        },
        "icon": "Includable.svg",
        "website": "http://includable.com"
      },
      "Index Exchange": {
        "cats": [
          36
        ],
        "description": "Index Exchange is a customizable exchange technology that enables sell side media firms to monetize ad inventories programmatically and in real time.",
        "icon": "Index Exchange.svg",
        "website": "https://www.indexexchange.com",
        "xhr": "\\.casalemedia\\.com"
      },
      "Indexhibit": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:indexhibit:indexhibit",
        "html": "<(?:link|a href) [^>]+ndxz-studio",
        "implies": [
          "PHP",
          "Apache",
          "Exhibit"
        ],
        "meta": {
          "generator": "Indexhibit"
        },
        "website": "http://www.indexhibit.org"
      },
      "Indico": {
        "cats": [
          1
        ],
        "cookies": {
          "MAKACSESSION": ""
        },
        "html": "Powered by\\s+(?:CERN )?<a href=\"http://(?:cdsware\\.cern\\.ch/indico/|indico-software\\.org|cern\\.ch/indico)\">(?:CDS )?Indico( [\\d\\.]+)?\\;version:\\1",
        "icon": "Indico.png",
        "website": "http://indico-software.org"
      },
      "Indy": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Indy(?:/([\\d.]+))?\\;version:\\1"
        },
        "website": "http://indyproject.org"
      },
      "Inertia": {
        "cats": [
          12
        ],
        "description": "Inertia is a protocol for creating monolithic single-page applications.",
        "dom": {
          "div[data-page*='component']": {
            "attributes": {
              "data-page": "component.+props.+url"
            }
          }
        },
        "headers": {
          "X-Inertia": ""
        },
        "icon": "Inertia.svg",
        "oss": true,
        "saas": false,
        "website": "https://inertiajs.com"
      },
      "InfernoJS": {
        "cats": [
          12
        ],
        "icon": "InfernoJS.png",
        "js": {
          "Inferno": "",
          "Inferno.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://infernojs.org"
      },
      "Insider": {
        "cats": [
          32
        ],
        "description": "Insider is the first integrated Growth Management Platform helping digital marketers drive growth across the funnel, from Acquisition to Activation, Retention, and Revenue from a unified platform powered by Artificial Intelligence and Machine Learning.",
        "icon": "Insider.svg",
        "js": {
          "Insider": "\\;confidence:20"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "api\\.useinsider\\.\\w+/",
        "website": "https://useinsider.com"
      },
      "Inspectlet": {
        "cats": [
          10
        ],
        "html": [
          "<!-- (?:Begin|End) Inspectlet Embed Code -->"
        ],
        "icon": "inspectlet.png",
        "js": {
          "__insp": "",
          "__inspld": ""
        },
        "scripts": [
          "cdn\\.inspectlet\\.com"
        ],
        "website": "https://www.inspectlet.com/"
      },
      "Instabot": {
        "cats": [
          5,
          10,
          32,
          52,
          58
        ],
        "description": "Instabot is a conversion chatbot that understands your users, and curates information, answers questions, captures contacts, and books meetings instantly.",
        "icon": "Instabot.png",
        "js": {
          "Instabot": ""
        },
        "scripts": "/rokoInstabot\\.js",
        "website": "https://instabot.io/"
      },
      "Instana": {
        "cats": [
          10,
          13,
          78
        ],
        "description": "Instana is a Kubernetes-native APM tool which is built for new-stack including Microservices and lately Serverless but also supports the existing VM based stacks  including several supported technologies.",
        "icon": "Instana.svg",
        "js": {
          "ineum": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "eum\\.instana\\.io",
        "website": "https://www.instana.com"
      },
      "InstantCMS": {
        "cats": [
          1
        ],
        "cookies": {
          "InstantCMS[logdate]": ""
        },
        "cpe": "cpe:/a:instantcms:instantcms",
        "icon": "InstantCMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "InstantCMS"
        },
        "website": "http://www.instantcms.ru"
      },
      "Instapage": {
        "cats": [
          51,
          74,
          10
        ],
        "description": "Instapage is a cloud-based landing page platform designed for marketing teams and agencies.",
        "icon": "Instapage.svg",
        "implies": [
          "Lua",
          "Node.js"
        ],
        "js": {
          "_instapageSnowplow": "",
          "instapageSp": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.instapagemetrics\\.com",
          "heatmap-events-collector\\.instapage\\.com"
        ],
        "website": "https://instapage.com"
      },
      "Intel Active Management Technology": {
        "cats": [
          22,
          46
        ],
        "cpe": "cpe:/a:intel:active_management_technology",
        "description": "Intel Active Management Technology (AMT) is a proprietary remote management and control system for personal computers with Intel CPUs.",
        "headers": {
          "Server": "Intel\\(R\\) Active Management Technology(?: ([\\d.]+))?\\;version:\\1"
        },
        "icon": "Intel Active Management Technology.png",
        "website": "http://intel.com"
      },
      "IntenseDebate": {
        "cats": [
          15
        ],
        "description": "IntenseDebate is a blog commenting system that supports Typepad, Blogger and Wordpress blogs. The system allows blog owners to track and moderate comments from one place with features like threading, comment analytics, user reputation, and comment aggregation.",
        "icon": "IntenseDebate.png",
        "scripts": "intensedebate\\.com",
        "website": "http://intensedebate.com"
      },
      "Intercom": {
        "cats": [
          52,
          53
        ],
        "description": "Intercom is an American software company that produces a messaging platform which allows businesses to communicate with prospective and existing customers within their app, on their website, through social media, or via email.",
        "icon": "Intercom.svg",
        "js": {
          "Intercom": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "(?:api\\.intercom\\.io/api|static\\.intercomcdn\\.com/intercom\\.v1)",
        "website": "https://www.intercom.com"
      },
      "Intercom Articles": {
        "cats": [
          4
        ],
        "description": "Intercom Articles is a tool to create, organise and publish help articles.",
        "html": "<a href=\"https://www.intercom.com/intercom-link[^\"]+solution=customer-support[^>]+>We run on Intercom",
        "icon": "Intercom.svg",
        "website": "https://www.intercom.com/articles"
      },
      "Intershop": {
        "cats": [
          6
        ],
        "html": "<ish-root",
        "icon": "Intershop.png",
        "scripts": "(?:is-bin|INTERSHOP)",
        "website": "http://intershop.com"
      },
      "Invenio": {
        "cats": [
          50
        ],
        "cookies": {
          "INVENIOSESSION": ""
        },
        "description": "Invenio is an open-source software framework for large-scale digital repositories that provides the tools for management of digital assets in an institutional repository and research data management systems.",
        "html": "(?:Powered by|System)\\s+(?:CERN )?<a (?:class=\"footer\" )?href=\"http://(?:cdsware\\.cern\\.ch(?:/invenio)?|invenio-software\\.org|cern\\.ch/invenio)(?:/)?\">(?:CDS )?Invenio</a>\\s*v?([\\d\\.]+)?\\;version:\\1",
        "icon": "Invenio.png",
        "website": "http://invenio-software.org"
      },
      "Inveon": {
        "cats": [
          6
        ],
        "cookies": {
          "INV.Customer": "\\;confidence:50",
          "inveonSessionId": ""
        },
        "description": "Inveon is a technology company that has been delivering ecommerce infrastructure software and mcommerce applications.",
        "icon": "Inveon.svg",
        "js": {
          "InvApp": "\\;confidence:50",
          "invTagManagerParams": ""
        },
        "scripts": "Scripts/_app/Inv(?:\\w+)\\.js\\?v=(.+)$\\;version:\\1",
        "website": "https://www.inveon.com"
      },
      "Ionic": {
        "cats": [
          18
        ],
        "icon": "ionic.png",
        "js": {
          "Ionic.config": "",
          "Ionic.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://ionicframework.com"
      },
      "Ionicons": {
        "cats": [
          17
        ],
        "description": "Ionicons is an open-source icon set crafted for web, iOS, Android, and desktop apps.",
        "html": "<link[^>]* href=[^>]+ionicons(?:\\.min)?\\.css",
        "icon": "Ionicons.png",
        "website": "http://ionicons.com"
      },
      "Irroba": {
        "cats": [
          6
        ],
        "html": "<a[^>]*href=\"https://www\\.irroba\\.com\\.br",
        "icon": "irroba.svg",
        "website": "https://www.irroba.com.br/"
      },
      "Isotope": {
        "cats": [
          59
        ],
        "description": "Isotope.js is a JavaScript library that makes it easy to sort, filter, and add Masonry layouts to items on a webpage.",
        "icon": "Isotope.svg",
        "js": {
          "Isotope": "",
          "init_isotope": ""
        },
        "oss": true,
        "pricing": [
          "low",
          "freemium",
          "onetime"
        ],
        "website": "https://isotope.metafizzy.co"
      },
      "Izooto": {
        "cats": [
          32,
          5
        ],
        "description": "iZooto is a user engagement and retention tool that leverages web push notifications to help business to drive repeat traffic, leads and sales.",
        "icon": "Izooto.png",
        "js": {
          "Izooto": "",
          "_izooto": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.izooto\\.\\w+",
        "website": "https://www.izooto.com"
      },
      "J2Store": {
        "cats": [
          6
        ],
        "icon": "j2store.png",
        "implies": "Joomla",
        "js": {
          "j2storeURL": ""
        },
        "website": "https://www.j2store.org/"
      },
      "JANet": {
        "cats": [
          71
        ],
        "description": "JANet is an affiliate marketing network.",
        "dom": "img[src*='.j-a-net.jp'],img[data-src*='.j-a-net.jp']",
        "icon": "JANet.png",
        "website": "https://j-a-net.jp"
      },
      "JAlbum": {
        "cats": [
          7
        ],
        "description": "jAlbum is across-platform photo website software for creating and uploading galleries from images and videos.",
        "icon": "JAlbum.png",
        "implies": "Java",
        "meta": {
          "generator": "JAlbum( [\\d.]+)?\\;version:\\1"
        },
        "website": "http://jalbum.net/en"
      },
      "JBoss Application Server": {
        "cats": [
          22
        ],
        "headers": {
          "X-Powered-By": "JBoss(?:-([\\d.]+))?\\;version:\\1"
        },
        "icon": "JBoss Application Server.png",
        "website": "http://jboss.org/jbossas.html"
      },
      "JBoss Web": {
        "cats": [
          22
        ],
        "excludes": "Apache Tomcat",
        "headers": {
          "X-Powered-By": "JBossWeb(?:-([\\d.]+))?\\;version:\\1"
        },
        "icon": "JBoss Web.png",
        "implies": "JBoss Application Server",
        "website": "http://jboss.org/jbossweb"
      },
      "JET Enterprise": {
        "cats": [
          6
        ],
        "headers": {
          "powered": "jet-enterprise"
        },
        "icon": "JET Enterprise.svg",
        "website": "http://www.jetecommerce.com.br/"
      },
      "JS Charts": {
        "cats": [
          25
        ],
        "icon": "JS Charts.png",
        "js": {
          "JSChart": ""
        },
        "scripts": "jscharts.{0,32}\\.js",
        "website": "http://www.jscharts.com"
      },
      "JSEcoin": {
        "cats": [
          56
        ],
        "description": "JSEcoin is a way to mine, receive payments for your goods or services and transfer cryptocurrency",
        "icon": "JSEcoin.png",
        "js": {
          "jseMine": ""
        },
        "scripts": "^(?:https):?//load\\.jsecoin\\.com/load/",
        "website": "https://jsecoin.com/"
      },
      "JShop": {
        "cats": [
          6
        ],
        "description": "JShop is the ecommerce database solution marketed by Whorl Ltd. worldwide.",
        "icon": "JShop.svg",
        "js": {
          "jss_1stepDeliveryType": "",
          "jss_1stepFillShipping": ""
        },
        "website": "http://www.whorl.co.uk"
      },
      "JTL Shop": {
        "cats": [
          6
        ],
        "cookies": {
          "JTLSHOP": ""
        },
        "html": "(?:<input[^>]+name=\"JTLSHOP|<a href=\"jtl\\.php)",
        "icon": "JTL Shop.png",
        "website": "http://www.jtl-software.de/produkte/jtl-shop3"
      },
      "JW Player": {
        "cats": [
          14
        ],
        "description": "JW Player is a online video player with video engagement analytics, custom video player skins, and live video streaming capability.",
        "dom": "div[data-video-provider*=jwplayer]",
        "icon": "JW Player.svg",
        "js": {
          "jwDefaults": "",
          "jwplayer": "",
          "jwplayerApiUrl": "",
          "webpackJsonpjwplayer": ""
        },
        "oss": true,
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "scripts": [
          "\\.jwplayer\\.com",
          "\\.jwpcdn\\.com"
        ],
        "website": "https://www.jwplayer.com",
        "xhr": "\\.jwpsrv\\.com"
      },
      "Jahia DX": {
        "cats": [
          1,
          47
        ],
        "html": "<script id=\"staticAssetAggregatedJavascrip",
        "icon": "JahiaDX.svg",
        "website": "http://www.jahia.com/dx"
      },
      "Jalios": {
        "cats": [
          1
        ],
        "icon": "Jalios.png",
        "meta": {
          "generator": "Jalios"
        },
        "website": "http://www.jalios.com"
      },
      "Java": {
        "cats": [
          27
        ],
        "cookies": {
          "JSESSIONID": ""
        },
        "description": "Java is a class-based, object-oriented programming language that is designed to have as few implementation dependencies as possible.",
        "icon": "Java.png",
        "website": "http://java.com"
      },
      "Java Servlet": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "Servlet(?:\\/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Java.png",
        "implies": "Java",
        "website": "http://www.oracle.com/technetwork/java/index-jsp-135475.html"
      },
      "JavaScript Infovis Toolkit": {
        "cats": [
          25
        ],
        "icon": "JavaScript Infovis Toolkit.png",
        "js": {
          "$jit": "",
          "$jit.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "jit(?:-yc)?\\.js",
        "website": "https://philogb.github.io/jit/"
      },
      "JavaServer Faces": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "JSF(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "JavaServer Faces.png",
        "implies": "Java",
        "website": "http://javaserverfaces.java.net"
      },
      "JavaServer Pages": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "JSP(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Java.png",
        "implies": "Java",
        "website": "http://www.oracle.com/technetwork/java/javaee/jsp/index.html"
      },
      "Javadoc": {
        "cats": [
          4
        ],
        "description": "Javadoc is a tool used for generating Java code documentation in HTML format from Java source code.",
        "html": "<!-- Generated by javadoc -->",
        "icon": "Java.png",
        "website": "https://docs.oracle.com/javase/8/docs/technotes/tools/windows/javadoc.html"
      },
      "Jekyll": {
        "cats": [
          57
        ],
        "cpe": "cpe:/a:jekyllrb:jekyll",
        "description": "Jekyll is a blog-aware, static site generator for personal, project, or organisation sites.",
        "html": [
          "Powered by <a href=\"https?://jekyllrb\\.com\"[^>]*>Jekyll</",
          "<!-- Created with Jekyll Now -",
          "<!-- Begin Jekyll SEO tag"
        ],
        "icon": "Jekyll.png",
        "meta": {
          "generator": "Jekyll (v[\\d.]+)?\\;version:\\1"
        },
        "website": "http://jekyllrb.com"
      },
      "Jenkins": {
        "cats": [
          44
        ],
        "description": "Jenkins is an open-source automation tool written in Java with plugins built for Continuous Integration (CI) purposes.",
        "headers": {
          "X-Jenkins": "([\\d.]+)\\;version:\\1"
        },
        "html": "<span class=\"jenkins_ver\"><a href=\"https://jenkins\\.io/\">Jenkins ver\\. ([\\d.]+)\\;version:\\1",
        "icon": "Jenkins.png",
        "implies": "Java",
        "js": {
          "jenkinsCIGlobal": "",
          "jenkinsRules": ""
        },
        "website": "https://jenkins.io/"
      },
      "Jetshop": {
        "cats": [
          6
        ],
        "html": "<(?:div|aside) id=\"jetshop-branding\">",
        "icon": "Jetshop.png",
        "js": {
          "JetshopData": ""
        },
        "website": "http://jetshop.se"
      },
      "Jetty": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Jetty(?:\\(([\\d\\.]*\\d+))?\\;version:\\1"
        },
        "icon": "Jetty.png",
        "implies": "Java",
        "website": "http://www.eclipse.org/jetty"
      },
      "Jibres": {
        "cats": [
          6,
          55
        ],
        "cookies": {
          "jibres": ""
        },
        "description": "Jibres is an ecommerce solution with an online store builder and Point-of-Sale (PoS) software.",
        "headers": {
          "X-Powered-By": "Jibres"
        },
        "icon": "Jibres.svg",
        "js": {
          "jibres": ""
        },
        "meta": {
          "generator": "Jibres"
        },
        "scripts": "/jibres\\.js",
        "website": "https://jibres.com"
      },
      "Jimdo": {
        "cats": [
          1
        ],
        "description": "Jimdo is a website-builder and all-in-one hosting solution, designed to enable users to build their own websites.",
        "headers": {
          "X-Jimdo-Instance": "",
          "X-Jimdo-Wid": ""
        },
        "icon": "jimdo.png",
        "pricing": [
          "low",
          "freemium"
        ],
        "saas": true,
        "url": "\\.jimdo\\.com/",
        "website": "https://www.jimdo.com"
      },
      "Jirafe": {
        "cats": [
          10,
          32
        ],
        "icon": "Jirafe.png",
        "js": {
          "jirafe": ""
        },
        "scripts": "/jirafe\\.js",
        "website": "https://docs.jirafe.com"
      },
      "Jitsi": {
        "cats": [
          52
        ],
        "description": "Jitsi is a free and open-source multiplatform voice (VoIP), videoconferencing and instant messaging applications for the web platform.",
        "icon": "Jitsi.png",
        "scripts": "lib-jitsi-meet.*\\.js",
        "website": "https://jitsi.org"
      },
      "Jive": {
        "cats": [
          19
        ],
        "headers": {
          "X-JIVE-USER-ID": "",
          "X-JSL": "",
          "X-Jive-Flow-Id": "",
          "X-Jive-Request-Id": "",
          "x-jive-chrome-wrapped": ""
        },
        "icon": "Jive.png",
        "website": "http://www.jivesoftware.com"
      },
      "JivoChat": {
        "cats": [
          52
        ],
        "description": "JivoChat is a live chat solution for websites offering customizable web and mobile chat widgets.",
        "icon": "JivoChat.png",
        "js": {
          "jivo_api": "",
          "jivo_version": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "freemium",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": "\\.jivosite\\.com",
        "website": "https://www.jivosite.com"
      },
      "JobberBase": {
        "cats": [
          19
        ],
        "icon": "JobberBase.png",
        "implies": "PHP",
        "js": {
          "Jobber": ""
        },
        "meta": {
          "generator": "Jobberbase"
        },
        "website": "http://www.jobberbase.com"
      },
      "JoomShopping": {
        "cats": [
          6
        ],
        "description": "JoomShopping is an open-source ecommerce plugin for Joomla.",
        "icon": "JoomShopping.png",
        "implies": "Joomla",
        "js": {
          "joomshoppingVideoHtml5": ""
        },
        "oss": true,
        "pricing": [
          "freemium",
          "onetime"
        ],
        "scripts": "/components/com_jshopping/",
        "website": "https://www.webdesigner-profi.de/joomla-webdesign/joomla-shop"
      },
      "Joomla": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:joomla:joomla",
        "description": "Joomla is a free and open-source content management system for publishing web content.",
        "headers": {
          "X-Content-Encoded-By": "Joomla! ([\\d.]+)\\;version:\\1"
        },
        "html": "(?:<div[^>]+id=\"wrapper_r\"|<(?:link|script)[^>]+(?:feed|components)/com_|<table[^>]+class=\"pill)\\;confidence:50",
        "icon": "Joomla.svg",
        "implies": "PHP",
        "js": {
          "Joomla": "",
          "jcomments": ""
        },
        "meta": {
          "generator": "Joomla!(?: ([\\d.]+))?\\;version:\\1"
        },
        "oss": true,
        "url": "option=com_",
        "website": "https://www.joomla.org"
      },
      "Jumpseller": {
        "cats": [
          6
        ],
        "description": "Jumpseller is a cloud ecommerce solution for small businesses.",
        "icon": "Jumpseller.svg",
        "js": {
          "Jumpseller": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "assets\\.jumpseller\\.\\w+/",
          "jumpseller-apps\\.herokuapp\\.\\w+/"
        ],
        "website": "https://jumpseller.com"
      },
      "K2": {
        "cats": [
          19
        ],
        "html": "<!--(?: JoomlaWorks \"K2\"| Start K2)",
        "icon": "K2.png",
        "implies": "Joomla",
        "js": {
          "K2RatingURL": ""
        },
        "website": "https://getk2.org"
      },
      "KISSmetrics": {
        "cats": [
          10
        ],
        "icon": "KISSmetrics.png",
        "js": {
          "KM_COOKIE_DOMAIN": ""
        },
        "website": "https://www.kissmetrics.com"
      },
      "KQS.store": {
        "cats": [
          6
        ],
        "description": "KQS.store is an ecommerce software.",
        "dom": [
          "a[href*='kqsdesign.pl'][target='_blank']",
          "a[href*='kqs.pl'][target='_blank']",
          "#kqs-box,kqs-cookie"
        ],
        "icon": "KQS.store.png",
        "js": {
          "kqs_box": "\\;confidence:50",
          "kqs_off": "\\;confidence:50"
        },
        "pricing": [
          "onetime"
        ],
        "saas": false,
        "website": "https://www.kqs.pl"
      },
      "KaTeX": {
        "cats": [
          25
        ],
        "description": "KaTeX is a cross-browser JavaScript library that displays mathematical notation in web browsers.",
        "dom": {
          "link[href*=katex]": {
            "attributes": {
              "href": "katex(?:\\.min)?\\.css"
            }
          }
        },
        "icon": "KaTeX.svg",
        "js": {
          "katex": "",
          "katex.version": "^(.+)$\\;version:\\1"
        },
        "oss": true,
        "scripts": "katex(@|/)[0-9.]+(?:/dist)?/katex(?:\\.min)?\\.(mjs|js|css)\\;version:\\1",
        "website": "https://katex.org/"
      },
      "Kajabi": {
        "cats": [
          6
        ],
        "cookies": {
          "_kjb_session": ""
        },
        "icon": "Kajabi.svg",
        "js": {
          "Kajabi": ""
        },
        "pricing": [
          "mid"
        ],
        "saas": true,
        "website": "https://newkajabi.com"
      },
      "Kameleoon": {
        "cats": [
          74
        ],
        "cookies": {
          "kameleoonVisitorCode": ""
        },
        "description": "Kameleoon is a personalisation technology platform for real-time omnichannel optimisation and conversion.",
        "dom": "link[href*='.kameleoon.eu/kameleoon.js']",
        "icon": "Kameleoon.svg",
        "js": {
          "Kameleoon.Gatherer.SCRIPT_VERSION": "(.+)\\;version:\\1",
          "kameleoonEndLoadTime": "",
          "kameleoonS": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.kameleoon\\.\\w+/kameleoon\\.js",
        "website": "https://kameleoon.com/"
      },
      "Kampyle": {
        "cats": [
          10,
          13
        ],
        "cookies": {
          "k_visit": ""
        },
        "description": "Kampyle is a complete customer feedback platform that helps digital enterprises listen, understand, and act across all digital touch-points.",
        "icon": "Kampyle.png",
        "js": {
          "KAMPYLE_COMMON": "",
          "k_track": "",
          "kampyle": ""
        },
        "scripts": "cf\\.kampyle\\.com/k_button\\.js",
        "website": "http://www.kampyle.com"
      },
      "Kamva": {
        "cats": [
          6
        ],
        "icon": "Kamva.svg",
        "js": {
          "Kamva": ""
        },
        "meta": {
          "generator": "[CK]amva"
        },
        "scripts": "cdn\\.mykamva\\.ir",
        "website": "https://kamva.ir"
      },
      "Karma": {
        "cats": [
          59
        ],
        "description": "Karma is a test runner for JavaScript that runs on Node.js.",
        "icon": "Karma.svg",
        "implies": "Node.js",
        "js": {
          "karma.vars.version": "(.+)\\;version:\\1"
        },
        "oss": true,
        "scripts": "karma\\.mdpcdn\\.com",
        "website": "https://karma-runner.github.io"
      },
      "Kartra": {
        "cats": [
          32
        ],
        "description": "Kartra is an online business platform that offers marketing and sales tools.",
        "dom": "form[action*='app.kartra.com']",
        "icon": "Kartra.png",
        "js": {
          "init_kartra_tracking": "",
          "kartra_tracking_loaded": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "app\\.kartra\\.com",
        "website": "https://home.kartra.com"
      },
      "Keap": {
        "cats": [
          53
        ],
        "description": "Keap offers an email marketing and sales platform for small businesses, including products to manage customers, customer relationship management, marketing, and ecommerce.",
        "dom": "form[action*='property.infusionsoft.com']",
        "icon": "Keap.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "property\\.infusionsoft\\.com",
        "website": "https://keap.com",
        "xhr": "property\\.infusionsoft\\.com"
      },
      "Kemal": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "X-Powered-By": "Kemal"
        },
        "icon": "kemalcr.png",
        "website": "http://kemalcr.com"
      },
      "Kendo UI": {
        "cats": [
          66
        ],
        "description": "Kendo UI is a HTML5 user interface framework for building interactive and high-performance websites and applications.",
        "html": "<link[^>]*\\s+href=[^>]*styles/kendo\\.common(?:\\.min)?\\.css[^>]*/>",
        "icon": "Kendo UI.png",
        "implies": "jQuery",
        "js": {
          "kendo": "",
          "kendo.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://www.telerik.com/kendo-ui"
      },
      "Kentico CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "CMSPreferredCulture": ""
        },
        "cpe": "cpe:/a:kentico:kentico_cms",
        "description": "Kentico CMS is a web content management system for building websites, online stores, intranets, and Web 2.0 community sites.",
        "icon": "Kentico CMS.png",
        "js": {
          "CMS.Application": ""
        },
        "meta": {
          "generator": "Kentico CMS ([\\d.R]+ \\(build [\\d.]+\\))\\;version:\\1"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "/CMSPages/GetResource\\.ashx",
        "website": "http://www.kentico.com"
      },
      "Kerberos": {
        "cats": [
          16
        ],
        "description": "Kerberos is an authentication method commonly used by Windows servers",
        "headers": {
          "WWW-Authenticate": "^Kerberos"
        },
        "website": "https://tools.ietf.org/html/rfc4559"
      },
      "Kestrel": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^Kestrel"
        },
        "icon": "kestrel.svg",
        "implies": "Microsoft ASP.NET",
        "website": "https://docs.microsoft.com/en-us/aspnet/core/fundamentals/servers/kestrel"
      },
      "KeyCDN": {
        "cats": [
          31
        ],
        "description": "KeyCDN is a content delivery network (CDN).",
        "headers": {
          "Server": "^keycdn-engine$"
        },
        "icon": "KeyCDN.png",
        "website": "http://www.keycdn.com"
      },
      "Kibana": {
        "cats": [
          29,
          25
        ],
        "cpe": "cpe:/a:elasticsearch:kibana",
        "description": "Kibana is an open-source data visualisation dashboard for Elasticsearch. It provides visualisation capabilities on top of the content indexed on an Elasticsearch cluster. Users can create bar, line and scatter plots, or pie charts and maps on top of large volumes of data.",
        "headers": {
          "kbn-name": "kibana",
          "kbn-version": "^([\\d.]+)$\\;version:\\1"
        },
        "html": "<title>Kibana</title>",
        "icon": "kibana.svg",
        "implies": [
          "Node.js",
          "Elasticsearch"
        ],
        "url": "kibana#/dashboard/",
        "website": "http://www.elastic.co/products/kibana"
      },
      "Kibo Commerce": {
        "cats": [
          6
        ],
        "description": "Kibo Commerce is a enterprise ecommerce platform  that offers a cloud-based, end-to-end commerce solution for retailers and branded manufacturers.",
        "icon": "Kibo.png",
        "pricing": [
          "poa"
        ],
        "scripts": "cdn-tp\\d+\\.mozu\\.com",
        "website": "https://kibocommerce.com"
      },
      "Kibo Personalization": {
        "cats": [
          76,
          74
        ],
        "description": "Kibo Personalization is a omnichannel personalisation software powered by machine learning to deliver individualized customer experiences and powered by Monetate and Certona.",
        "icon": "Kibo.png",
        "js": {
          "BaynoteAPI": "",
          "BaynoteJSVersion": "",
          "monetate": "",
          "monetateQ": "",
          "monetateT": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "\\.monetate\\.net",
          "\\.baynote\\.net"
        ],
        "website": "https://kibocommerce.com/personalization-software"
      },
      "KineticJS": {
        "cats": [
          25
        ],
        "icon": "KineticJS.png",
        "js": {
          "Kinetic": "",
          "Kinetic.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "kinetic(?:-v?([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://github.com/ericdrowell/KineticJS/"
      },
      "Kinsta": {
        "cats": [
          62
        ],
        "headers": {
          "x-kinsta-cache": ""
        },
        "icon": "kinsta.svg",
        "implies": "WordPress",
        "website": "https://kinsta.com"
      },
      "Klarna Checkout": {
        "cats": [
          41,
          6
        ],
        "description": "Klarna Checkout is a complete payment solution where Klarna handles a store's entire checkout.",
        "icon": "Klarna.svg",
        "js": {
          "_klarnaCheckout": ""
        },
        "website": "https://www.klarna.com/international/"
      },
      "Klaviyo": {
        "cats": [
          32
        ],
        "description": "Klaviyo is an email marketing platform for online businesses.",
        "icon": "Klaviyo.svg",
        "js": {
          "KlaviyoSubscribe": "",
          "klaviyo": ""
        },
        "scripts": "klaviyo\\.com",
        "website": "https://www.klaviyo.com/"
      },
      "Klevu": {
        "cats": [
          29,
          6
        ],
        "description": "Klevu is a highly advanced AI-Powered search solution for ecommerce platforms.",
        "icon": "Klevu.png",
        "js": {
          "klevu_apiKey": "",
          "klevu_layout": "",
          "klevu_sessionId": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "js\\.klevu\\.\\w+/klevu-js-v([\\d.]+)\\;version:\\1",
        "website": "https://www.klevu.com"
      },
      "Knockout.js": {
        "cats": [
          12
        ],
        "icon": "Knockout.js.png",
        "js": {
          "ko.version": "^(.+)$\\;version:\\1"
        },
        "website": "http://knockoutjs.com"
      },
      "Ko-fi": {
        "cats": [
          5,
          41
        ],
        "description": "Ko-fi is an online platform that helps content creators get the financial support.",
        "dom": "a[href*='ko-fi.com/'][target='_blank']",
        "icon": "Ko-fi.svg",
        "js": {
          "kofiwidget2": ""
        },
        "pricing": [
          "freemium",
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "ko-fi\\.com/widgets",
        "website": "https://ko-fi.com"
      },
      "Koa": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "X-Powered-By": "^koa$"
        },
        "icon": "Koa.png",
        "implies": "Node.js",
        "website": "http://koajs.com"
      },
      "Koala Framework": {
        "cats": [
          1,
          18
        ],
        "html": "<!--[^>]+This website is powered by Koala Web Framework CMS",
        "icon": "Koala Framework.png",
        "implies": "PHP",
        "meta": {
          "generator": "^Koala Web Framework CMS"
        },
        "website": "http://koala-framework.org"
      },
      "KobiMaster": {
        "cats": [
          6
        ],
        "icon": "Kobimaster.png",
        "implies": "Microsoft ASP.NET",
        "js": {
          "kmGetSession": "",
          "kmPageInfo": ""
        },
        "website": "https://www.kobimaster.com.tr"
      },
      "Koha": {
        "cats": [
          50
        ],
        "cpe": "cpe:/a:koha:koha",
        "description": "Koha is an open-source Integrated Library System (ILS).",
        "html": [
          "<input name=\"koha_login_context\" value=\"intranet\" type=\"hidden\">",
          "<a href=\"/cgi-bin/koha/"
        ],
        "icon": "koha.png",
        "implies": "Perl",
        "js": {
          "KOHA": ""
        },
        "meta": {
          "generator": "^Koha ([\\d.]+)$\\;version:\\1"
        },
        "website": "https://koha-community.org/"
      },
      "Kohana": {
        "cats": [
          18
        ],
        "cookies": {
          "kohanasession": ""
        },
        "cpe": "cpe:/a:kohanaframework:kohana",
        "headers": {
          "X-Powered-By": "Kohana Framework ([\\d.]+)\\;version:\\1"
        },
        "icon": "Kohana.png",
        "implies": "PHP",
        "website": "http://kohanaframework.org"
      },
      "Koken": {
        "cats": [
          1
        ],
        "cookies": {
          "koken_referrer": ""
        },
        "html": [
          "<html lang=\"en\" class=\"k-source-essays k-lens-essays\">",
          "<!--\\s+KOKEN DEBUGGING"
        ],
        "icon": "Koken.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "meta": {
          "generator": "Koken ([\\d.]+)\\;version:\\1"
        },
        "scripts": "koken(?:\\.js\\?([\\d.]+)|/storage)\\;version:\\1",
        "website": "http://koken.me"
      },
      "Komodo CMS": {
        "cats": [
          1
        ],
        "icon": "Komodo CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "^Komodo CMS"
        },
        "website": "http://www.komodocms.com"
      },
      "Konduto": {
        "cats": [
          16
        ],
        "description": "Konduto is a fraud detection service for ecommerce.",
        "dom": "link[href*='.konduto.com']",
        "icon": "Konduto.png",
        "js": {
          "Konduto": "",
          "getKondutoID": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.k-analytix\\.com",
        "website": "https://www.konduto.com"
      },
      "Koobi": {
        "cats": [
          1
        ],
        "html": "<!--[^K>-]+Koobi ([a-z\\d.]+)\\;version:\\1",
        "icon": "Koobi.png",
        "meta": {
          "generator": "Koobi"
        },
        "website": "http://dream4.de/cms"
      },
      "Kooboo CMS": {
        "cats": [
          1
        ],
        "headers": {
          "X-KoobooCMS-Version": "^(.+)$\\;version:\\1"
        },
        "icon": "Kooboo CMS.png",
        "implies": "Microsoft ASP.NET",
        "scripts": "/Kooboo",
        "website": "http://kooboo.com"
      },
      "Kooomo": {
        "cats": [
          6
        ],
        "description": "Kooomo is a SaaS ecommerce platform with a focus on localisation and internationalisation.",
        "dom": "img[src*='.kooomo-cloud.com']",
        "icon": "Kooomo.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "meta": {
          "generator": "Kooomo(?: v([\\d.]+))?\\;version:\\1"
        },
        "pricing": [
          "mid",
          "recurring",
          "payg"
        ],
        "saas": true,
        "website": "https://www.kooomo.com"
      },
      "Kotisivukone": {
        "cats": [
          1
        ],
        "icon": "Kotisivukone.png",
        "scripts": "kotisivukone(?:\\.min)?\\.js",
        "website": "http://www.kotisivukone.fi"
      },
      "Kount": {
        "cats": [
          6,
          16
        ],
        "description": "Kount is a suite of fraud detection and prevention solutions for ecommerce businesses.",
        "icon": "Kount.svg",
        "js": {
          "ka.ClientSDK": "",
          "ka.collectData": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "shopify\\.kount\\.net/js"
        ],
        "website": "https://kount.com"
      },
      "Kubernetes Dashboard": {
        "cats": [
          47
        ],
        "cpe": "cpe:/a:kubernetes:kubernetes",
        "html": "<html ng-app=\"kubernetesDashboard\">",
        "icon": "Kubernetes.svg",
        "website": "https://kubernetes.io/"
      },
      "LEPTON": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:lepton-cms:lepton",
        "icon": "LEPTON.png",
        "implies": "PHP",
        "meta": {
          "generator": "LEPTON"
        },
        "website": "http://www.lepton-cms.org"
      },
      "LKQD": {
        "cats": [
          36
        ],
        "description": "LKQD is a video advertising platform that enables publishers to serve video ads across multiple devices and formats.",
        "dom": "iframe[src*='.lkqd.net']",
        "icon": "LKQD.svg",
        "js": {
          "lkqdCall": "",
          "lkqdErrorCount": "",
          "lkqdSettings": "",
          "lkqd_http_response": ""
        },
        "scripts": "\\.lkqd\\.net",
        "website": "https://wiki.lkqd.com",
        "xhr": "\\.lkqd\\.net"
      },
      "LOU": {
        "cats": [
          58
        ],
        "description": "LOU is a Digital Adoption Platform that streamlines user onboarding and training.",
        "icon": "LOU.png",
        "scripts": "cdn\\.louassist\\.com*",
        "website": "https://www.louassist.com"
      },
      "Lagoon": {
        "cats": [
          62
        ],
        "headers": {
          "X-LAGOON": "",
          "x-lagoon": ""
        },
        "icon": "Lagoon.png",
        "website": "https://www.amazee.io/hosting"
      },
      "Laravel": {
        "cats": [
          18
        ],
        "cookies": {
          "laravel_session": ""
        },
        "cpe": "cpe:/a:laravel:laravel",
        "description": "Laravel is a free, open-source PHP web framework.",
        "icon": "Laravel.svg",
        "implies": "PHP",
        "js": {
          "Laravel": ""
        },
        "website": "https://laravel.com"
      },
      "Laterpay": {
        "cats": [
          41
        ],
        "description": "Laterpay is a service that simplifies payments on the Internet. In addition to the classic immediate purchase option, Laterpay also allows you to buy digital content such as articles or videos now and pay later.",
        "icon": "laterpay.png",
        "meta": {
          "laterpay:connector:callbacks:on_user_has_access": "deobfuscateText"
        },
        "scripts": "https?://connectormwi\\.laterpay\\.net/([0-9.]+)[a-zA-z-]*/live/[\\w-]+\\.js\\;version:\\1",
        "website": "https://www.laterpay.net/"
      },
      "Launchrock": {
        "cats": [
          74,
          51
        ],
        "description": "Launchrock is an online tool designed to help capture email addresses and create online product launching events.",
        "icon": "Launchrock.svg",
        "js": {
          "lrIgnition": "\\;confidence:25",
          "lrLoadedJs": "\\;confidence:25",
          "lrSiteSettingAsBoolean": "\\;confidence:25",
          "lrignition": "\\;confidence:25"
        },
        "pricing": [
          "freemium",
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "js/ignition-current\\.min\\.js\\;confidence:50",
        "website": "https://www.launchrock.com"
      },
      "Leadinfo": {
        "cats": [
          10
        ],
        "description": "Leadinfo is a lead generation software that enables you to recognise B2B website visitors.",
        "icon": "Leadinfo.svg",
        "js": {
          "GlobalLeadinfoNamespace": "\\;confidence:50",
          "leadinfo": "\\;confidence:50"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.leadinfo\\.net",
        "website": "https://www.leadinfo.com"
      },
      "Leaflet": {
        "cats": [
          35
        ],
        "description": "Leaflet is the open-source JavaScript library for mobile-friendly interactive maps.",
        "icon": "Leaflet.png",
        "js": {
          "L.DistanceGrid": "",
          "L.PosAnimation": "",
          "L.version": "^(.+)$\\;version:\\1\\;confidence:0"
        },
        "scripts": "leaflet.{0,32}\\.js",
        "website": "http://leafletjs.com"
      },
      "Leanplum": {
        "cats": [
          32,
          74
        ],
        "description": "Leanplum is a multi-channel messaging and campaign orchestration platform.",
        "icon": "Leanplum.svg",
        "js": {
          "Leanplum": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "npm/leanplum-sdk\\@([\\d.]+)\\;version:\\1",
        "website": "https://www.leanplum.com"
      },
      "Lede": {
        "cats": [
          1
        ],
        "description": "Lede is a publishing platform and growth program designed to support journalism startups and news media.",
        "html": [
          "<a [^>]*href=\"[^\"]+joinlede.com"
        ],
        "icon": "lede.png",
        "implies": [
          "WordPress",
          "WordPress VIP"
        ],
        "js": {
          "ledeChartbeatViews": "",
          "ledeEngagement": "",
          "ledeEngagementReset": ""
        },
        "meta": {
          "og:image": "https?\\:\\/\\/lede-admin"
        },
        "saas": true,
        "website": "https://joinlede.com/"
      },
      "Less": {
        "cats": [
          19
        ],
        "html": "<link[^>]+ rel=\"stylesheet/less\"",
        "icon": "Less.png",
        "website": "http://lesscss.org"
      },
      "Let's Encrypt": {
        "cats": [
          70
        ],
        "certIssuer": "Let's Encrypt",
        "description": "Let's Encrypt is a free, automated, and open certificate authority.",
        "icon": "Lets Encrypt.svg",
        "website": "https://letsencrypt.org/"
      },
      "Liferay": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:liferay:liferay_portal",
        "description": "Liferay is an open-source company that provides free documentation and paid professional service to users of its software.",
        "headers": {
          "Liferay-Portal": "[a-z\\s]+([\\d.]+)\\;version:\\1"
        },
        "icon": "Liferay.svg",
        "js": {
          "Liferay": ""
        },
        "oss": true,
        "pricing": [
          "poa"
        ],
        "website": "https://www.liferay.com"
      },
      "Lift": {
        "cats": [
          18
        ],
        "cpe": "cpe:/a:liftweb:lift",
        "headers": {
          "X-Lift-Version": "(.+)\\;version:\\1"
        },
        "icon": "Lift.png",
        "implies": "Scala",
        "website": "http://liftweb.net"
      },
      "LightMon Engine": {
        "cats": [
          1
        ],
        "cookies": {
          "lm_online": ""
        },
        "html": "<!-- Lightmon Engine Copyright Lightmon",
        "icon": "LightMon Engine.png",
        "implies": "PHP",
        "meta": {
          "generator": "LightMon Engine"
        },
        "website": "http://lightmon.ru"
      },
      "Lightbox": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:lightbox_photo_gallery_project:lightbox_photo_gallery",
        "html": "<link [^>]*href=\"[^\"]+lightbox(?:\\.min)?\\.css",
        "icon": "Lightbox.png",
        "scripts": "lightbox(?:-plus-jquery)?.{0,32}\\.js",
        "website": "http://lokeshdhakar.com/projects/lightbox2/"
      },
      "Lightspeed eCom": {
        "cats": [
          6
        ],
        "html": "<!-- \\[START\\] 'blocks/head\\.rain' -->",
        "icon": "Lightspeed.svg",
        "pricing": [
          "low"
        ],
        "saas": true,
        "scripts": "http://assets\\.webshopapp\\.com",
        "url": "seoshop.webshopapp.com",
        "website": "http://www.lightspeedhq.com/products/ecommerce/"
      },
      "LinkSmart": {
        "cats": [
          36
        ],
        "icon": "LinkSmart.png",
        "js": {
          "LS_JSON": "",
          "LinkSmart": "",
          "_mb_site_guid": ""
        },
        "scripts": "^https?://cdn\\.linksmart\\.com/linksmart_([\\d.]+?)(?:\\.min)?\\.js\\;version:\\1",
        "website": "http://linksmart.com"
      },
      "Linkedin Insight Tag": {
        "cats": [
          10
        ],
        "dom": "noscript > img[src*='dc.ads.linkedin.com']",
        "icon": "Linkedin.svg",
        "js": {
          "_linkedin_data_partner_id": ""
        },
        "scripts": "snap\\.licdn\\.com/li\\.lms-analytics/insight\\.min\\.js",
        "website": "https://business.linkedin.com/marketing-solutions/insight-tag"
      },
      "Linkedin Sign-in": {
        "cats": [
          69
        ],
        "description": "Linkedin Sign-In is an authentication system that reduces the burden of login for users, by enabling them to sign in with their Linkedin account.",
        "icon": "Linkedin.svg",
        "js": {
          "OnLinkedInAuth": "",
          "onLinkedInLoad": ""
        },
        "scripts": "platform\\.linkedin\\.com/(?:.*)?in\\.js(?:\\?version)?([\\d.]+)?\\;version:\\1",
        "website": "https://www.linkedin.com/developers"
      },
      "Liquid Web": {
        "cats": [
          62
        ],
        "headers": {
          "x-lw-cache": ""
        },
        "icon": "liquidweb.svg",
        "website": "https://www.liquidweb.com"
      },
      "List.js": {
        "cats": [
          59
        ],
        "icon": "List.js.png",
        "js": {
          "List": "\\;confidence:50"
        },
        "scripts": [
          "list\\.js/\\;confidence:50",
          "@([\\d.]+)/(?:/dist)?list\\.(?:min\\.)?js\\;version:\\1"
        ],
        "website": "http://listjs.com"
      },
      "Listrak": {
        "cats": [
          32
        ],
        "description": "Listrak is a AI-based marketing automation and CRM solutions that unify, interpret and personalise data to engage customer across channels and devices.",
        "icon": "Listrak.png",
        "js": {
          "_LTKSignup": "",
          "_LTKSubscriber": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "(?:cdn|s1)\\.listrakbi\\.com",
          "services\\.listrak\\.com"
        ],
        "website": "https://www.listrak.com"
      },
      "LiteSpeed": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:litespeedtech:litespeed_web_server",
        "description": "LiteSpeed is a high-scalability web server.",
        "headers": {
          "Server": "^LiteSpeed$"
        },
        "icon": "LiteSpeed.svg",
        "website": "http://litespeedtech.com"
      },
      "Litespeed Cache": {
        "cats": [
          23
        ],
        "description": "LiteSpeed Cache is an all-in-one site acceleration plugin for WordPress.",
        "headers": {
          "x-litespeed-cache": ""
        },
        "icon": "litespeed-cache.png",
        "implies": "LiteSpeed",
        "website": "https://wordpress.org/plugins/litespeed-cache/"
      },
      "Lithium": {
        "cats": [
          1
        ],
        "cookies": {
          "LithiumVisitor": ""
        },
        "html": " <a [^>]+Powered by Lithium",
        "icon": "Lithium.png",
        "implies": "PHP",
        "js": {
          "LITHIUM": ""
        },
        "website": "https://www.lithium.com"
      },
      "Live Story": {
        "cats": [
          1
        ],
        "icon": "LiveStory.png",
        "js": {
          "LSHelpers": "",
          "LiveStory": ""
        },
        "website": "https://www.livestory.nyc/"
      },
      "LiveAgent": {
        "cats": [
          52
        ],
        "description": "LiveAgent is an online live chat platform. The software provides a ticket management system.",
        "icon": "LiveAgent.png",
        "js": {
          "LiveAgent": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://www.liveagent.com"
      },
      "LiveChat": {
        "cats": [
          52
        ],
        "description": "LiveChat is an online customer service software with online chat, help desk software, and web analytics capabilities.",
        "icon": "LiveChat.png",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "cdn\\.livechatinc\\.com/.*tracking\\.js",
        "website": "http://livechatinc.com"
      },
      "LiveHelp": {
        "cats": [
          52,
          53
        ],
        "description": "LiveHelp is an online chat tool.",
        "icon": "LiveHelp.png",
        "js": {
          "LHready": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "http://www.livehelp.it"
      },
      "LiveIntent": {
        "cats": [
          75,
          36
        ],
        "description": "LiveIntent is an email ad monetization platform.",
        "icon": "LiveIntent.svg",
        "js": {
          "LI.advertiserId": "\\d+"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.liadm\\.com",
        "website": "https://www.liveintent.com",
        "xhr": "\\.liadm\\.com"
      },
      "LiveJournal": {
        "cats": [
          11
        ],
        "description": "LiveJournal is a social networking service where users can keep a blog, journal or diary.",
        "icon": "LiveJournal.png",
        "url": "\\.livejournal\\.com",
        "website": "http://www.livejournal.com"
      },
      "LivePerson": {
        "cats": [
          52
        ],
        "description": "LivePerson is a tool for conversational chatbots and messaging.",
        "icon": "LivePerson.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "^https?://lptag\\.liveperson\\.net/tag/tag\\.js",
        "website": "https://www.liveperson.com/"
      },
      "LiveRamp PCM": {
        "cats": [
          67
        ],
        "description": "LiveRamp PCM is a preference and consent management platform that enables comply with the ePrivacy Directive, GDPR, CCPA, and other data protection and privacy laws and regulations.",
        "dom": "iframe[src*='gdpr-consent-tool\\.privacymanager\\.io']",
        "icon": "LiveRamp.svg",
        "js": {
          "wpJsonpLiverampGdprCmp": ""
        },
        "scripts": "gdpr\\.privacymanager\\.io",
        "website": "https://liveramp.com/our-platform/preference-consent-management"
      },
      "LiveStreet CMS": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "LiveStreet CMS"
        },
        "icon": "LiveStreet CMS.png",
        "implies": "PHP",
        "js": {
          "LIVESTREET_SECURITY_KEY": ""
        },
        "website": "http://livestreetcms.com"
      },
      "LiveZilla": {
        "cats": [
          52
        ],
        "description": "LiveZilla is a web-based live support platform.",
        "dom": "#lz_overlay_chat",
        "icon": "LiveZilla.png",
        "js": {
          "lz_chat_execute": "",
          "lz_code_id": "(?:[\\w\\d]+)",
          "lz_tracking_set_widget_visibility": ""
        },
        "pricing": [
          "onetime",
          "mid"
        ],
        "saas": false,
        "website": "https://www.livezilla.net"
      },
      "Livefyre": {
        "cats": [
          15
        ],
        "description": "Livefyre is a platform that integrates with the social web to boost social interaction.",
        "html": "<[^>]+(?:id|class)=\"livefyre",
        "icon": "Livefyre.png",
        "js": {
          "FyreLoader": "",
          "L.version": "^(.+)$\\;confidence:0\\;version:\\1",
          "LF.CommentCount": "",
          "fyre": ""
        },
        "scripts": "livefyre_init\\.js",
        "website": "http://livefyre.com"
      },
      "Liveinternet": {
        "cats": [
          10
        ],
        "html": [
          "<script [^>]*>[\\s\\S]*//counter\\.yadro\\.ru/hit",
          "<!--LiveInternet counter-->",
          "<!--/LiveInternet-->",
          "<a href=\"http://www\\.liveinternet\\.ru/click\""
        ],
        "icon": "Liveinternet.png",
        "website": "http://liveinternet.ru/rating/"
      },
      "Livewire": {
        "cats": [
          18,
          19
        ],
        "html": "<[^>]{1,512}\\bwire:",
        "icon": "Livewire.png",
        "implies": "Laravel",
        "js": {
          "livewire": ""
        },
        "scripts": "livewire(?:\\.min)?\\.js",
        "website": "https://laravel-livewire.com"
      },
      "LocalFocus": {
        "cats": [
          25
        ],
        "html": "<iframe[^>]+\\blocalfocus\\b",
        "icon": "LocalFocus.png",
        "implies": [
          "Angular",
          "D3"
        ],
        "website": "https://www.localfocus.nl/en/"
      },
      "Localised": {
        "cats": [
          6
        ],
        "description": "Localised is local-first ecommerce platform.",
        "dom": {
          "img[src='.localised.com']": {
            "attributes": {
              "src": ""
            }
          },
          "span": {
            "attributes": {
              "text": ".+Localised Inc\\..+"
            }
          }
        },
        "icon": "Localised.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.localised.com",
        "xhr": "api\\.localised\\.com"
      },
      "LocomotiveCMS": {
        "cats": [
          1
        ],
        "html": "<link[^>]*/sites/[a-z\\d]{24}/theme/stylesheets",
        "icon": "LocomotiveCMS.png",
        "implies": [
          "Ruby on Rails",
          "MongoDB"
        ],
        "website": "https://www.locomotivecms.com"
      },
      "Lodash": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:lodash:lodash",
        "description": "Lodash is a JavaScript library which provides utility functions for common programming tasks using the functional programming paradigm.",
        "excludes": "Underscore.js",
        "icon": "Lo-dash.png",
        "js": {
          "_.VERSION": "^(.+)$\\;confidence:0\\;version:\\1",
          "_.differenceBy": "",
          "_.templateSettings.imports._.templateSettings.imports._.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "lodash.*\\.js",
        "website": "http://www.lodash.com"
      },
      "LogRocket": {
        "cats": [
          10
        ],
        "description": "LogRocket records videos of user sessions with logs and network data.",
        "icon": "LogRocket.svg",
        "scripts": [
          "cdn\\.logrocket\\.(com|io)",
          "cdn\\.lr-ingest\\.io"
        ],
        "website": "https://logrocket.com/"
      },
      "Login with Amazon": {
        "cats": [
          69
        ],
        "description": "Login with Amazon allows you use your Amazon user name and password to sign into and share information with participating third-party websites or apps.",
        "icon": "Amazon.svg",
        "js": {
          "onAmazonLoginReady": ""
        },
        "website": "https://developer.amazon.com/apps-and-games/login-with-amazon"
      },
      "LoginRadius": {
        "cats": [
          5,
          69
        ],
        "description": "LoginRadius is a cloud based SaaS Customer Identity Access Management platform based in Canada.",
        "icon": "LoginRadius.svg",
        "js": {
          "LoginRadius": "",
          "LoginRadiusDefaults": "",
          "LoginRadiusSDK": "",
          "LoginRadiusUtility": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "\\.loginradius\\.com",
          "\\.lrcontent\\.com"
        ],
        "website": "https://www.loginradius.com"
      },
      "Loja Integrada": {
        "cats": [
          6
        ],
        "headers": {
          "X-Powered-By": "vtex-integrated-store"
        },
        "icon": "Loja Integrada.png",
        "js": {
          "LOJA_ID": ""
        },
        "website": "https://lojaintegrada.com.br/"
      },
      "Loja Mestre": {
        "cats": [
          6
        ],
        "description": "Loja Mestre is an all-in-one ecommerce platform from Brazil.",
        "icon": "Loja Mestre.svg",
        "meta": {
          "webmaster": "www\\.lojamestre\\.\\w+\\.br"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "lojamestre\\.\\w+\\.br",
        "website": "https://www.lojamestre.com.br/"
      },
      "Loja Virtual": {
        "cats": [
          6
        ],
        "description": "Loja Virtual is an all-in-one ecommerce platform from Brazil.",
        "icon": "Loja Virtual.svg",
        "js": {
          "id_loja_virtual": "",
          "link_loja_virtual": "",
          "loja_sem_dominio": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "/js/ljvt_v(\\d+)/\\;version:\\1\\;confidence:20",
          "cdn1\\.solojavirtual\\.com"
        ],
        "website": "https://www.lojavirtual.com.br"
      },
      "Loja2": {
        "cats": [
          6
        ],
        "description": "Loja2 is an all-in-one ecommerce platform from Brazil.",
        "icon": "Loja2.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "loja2\\.com\\.br",
        "website": "https://www.loja2.com.br"
      },
      "Loox": {
        "cats": [
          5,
          6
        ],
        "description": "Loox is a reviews app for Shopify that helps you gather reviews and user-generated photos from your customers.",
        "icon": "Loox.svg",
        "js": {
          "loox_global_hash": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "loox\\.io/widget",
        "website": "https://loox.app"
      },
      "Lotus Domino": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:ibm:lotus_domino",
        "headers": {
          "Server": "Lotus-Domino"
        },
        "icon": "Lotus Domino.png",
        "implies": "Java",
        "website": "http://www-01.ibm.com/software/lotus/products/domino"
      },
      "Lua": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:lua:lua",
        "description": "Lua is a multi-paradigm programming language designed primarily for embedded use in applications.",
        "headers": {
          "X-Powered-By": "\\bLua(?: ([\\d.]+))?\\;version:\\1"
        },
        "icon": "Lua.png",
        "website": "http://www.lua.org"
      },
      "Lucene": {
        "cats": [
          34
        ],
        "cpe": "cpe:/a:apache:lucene",
        "description": "Lucene is a free and open-source search engine software library.",
        "icon": "Lucene.png",
        "implies": "Java",
        "website": "http://lucene.apache.org/core/"
      },
      "Luigi’s Box": {
        "cats": [
          10,
          29
        ],
        "icon": "Luigisbox.svg",
        "js": {
          "Luigis": ""
        },
        "website": "https://www.luigisbox.com"
      },
      "MODX": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:modx:modx_revolution",
        "description": "MODX is an open source content management system and web application framework for publishing content on the world wide web and intranets.",
        "headers": {
          "X-Powered-By": "^MODX"
        },
        "html": [
          "<a[^>]+>Powered by MODX</a>",
          "<(?:link|script)[^>]+assets/snippets/\\;confidence:20",
          "<form[^>]+id=\"ajaxSearch_form\\;confidence:20",
          "<input[^>]+id=\"ajaxSearch_input\\;confidence:20"
        ],
        "icon": "MODX.png",
        "implies": "PHP",
        "js": {
          "MODX": "",
          "MODX_MEDIA_PATH": ""
        },
        "meta": {
          "generator": "MODX[^\\d.]*([\\d.]+)?\\;version:\\1"
        },
        "oss": true,
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "http://modx.com"
      },
      "MadAdsMedia": {
        "cats": [
          36
        ],
        "icon": "MadAdsMedia.png",
        "js": {
          "setMIframe": "",
          "setMRefURL": ""
        },
        "scripts": "^https?://(?:ads-by|pixel)\\.madadsmedia\\.com/",
        "website": "http://madadsmedia.com"
      },
      "Magento": {
        "cats": [
          6
        ],
        "cookies": {
          "frontend": "\\;confidence:50"
        },
        "cpe": "cpe:/a:magento:magento",
        "description": "Magento is an open-source ecommerce platform written in PHP.",
        "html": [
          "<script [^>]+data-requiremodule=\"mage/\\;version:2",
          "<script [^>]+data-requiremodule=\"Magento_\\;version:2",
          "<script type=\"text/x-magento-init\">"
        ],
        "icon": "Magento.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "Mage": "",
          "VarienForm": ""
        },
        "magento": "Magento/([0-9.]+)\\;version:\\1",
        "oss": true,
        "scripts": [
          "js/mage",
          "skin/frontend/(?:default|(enterprise))\\;version:\\1?Enterprise:Community",
          "static/_requirejs\\;confidence:50\\;version:2"
        ],
        "website": "https://magento.com"
      },
      "MailChimp": {
        "cats": [
          32,
          75
        ],
        "cpe": "cpe:/a:thinkshout:mailchimp",
        "description": "Mailchimp is a marketing automation platform and email marketing service.",
        "dns": {
          "TXT": [
            "spf\\.mandrillapp\\.com"
          ]
        },
        "html": [
          "<form [^>]*data-mailchimp-url",
          "<form [^>]*id=\"mc-embedded-subscribe-form\"",
          "<form [^>]*name=\"mc-embedded-subscribe-form\"",
          "<input [^>]*id=\"mc-email\"\\;confidence:20",
          "<!-- Begin MailChimp Signup Form -->"
        ],
        "icon": "mailchimp.svg",
        "js": {
          "mc4wp": "\\;confidence:50"
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "s3\\.amazonaws\\.com/downloads\\.mailchimp\\.com/js/mc-validate\\.js",
          "cdn-images\\.mailchimp\\.com/[^>]*\\.css",
          "mailchimp-woocommerce-public\\.min\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
          "mailchimp-for-wp/assets/js/forms\\.min\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
          "chimpstatic\\.com/mcjs-connected"
        ],
        "website": "http://mailchimp.com"
      },
      "Mailgun": {
        "cats": [
          75
        ],
        "description": "Mailgun is a transactional email API service for developers.",
        "dns": {
          "TXT": [
            "mailgun\\.org"
          ]
        },
        "icon": "Mailgun.svg",
        "website": "https://www.mailgun.com/"
      },
      "Mailjet": {
        "cats": [
          75
        ],
        "description": "Mailjet is an email delivery service for marketing and developer teams.",
        "dns": {
          "TXT": [
            "mailjet\\.com"
          ]
        },
        "icon": "Mailjet.svg",
        "website": "https://www.mailjet.com/"
      },
      "Make-Sense": {
        "cats": [
          68
        ],
        "icon": "Make-Sense.png",
        "scripts": "mk-sense\\.com/aweb\\?license",
        "website": "https://mk-sense.com/"
      },
      "MakeShopKorea": {
        "cats": [
          6
        ],
        "icon": "MakeShopKorea.png",
        "js": {
          "Makeshop": "",
          "MakeshopLogUniqueId": ""
        },
        "website": "https://www.makeshop.co.kr"
      },
      "Mambo": {
        "cats": [
          1
        ],
        "excludes": "Joomla",
        "icon": "Mambo.png",
        "meta": {
          "generator": "Mambo"
        },
        "website": "http://mambo-foundation.org"
      },
      "Mangeznotez": {
        "cats": [
          5,
          72
        ],
        "description": "Mangeznotez is a restaurant table booking widget.",
        "icon": "Mangeznotez.svg",
        "scripts": [
          "www\\.mangeznotez\\.\\w+",
          "\\w+.mangeznotez\\.\\w+(?:.*\\?ver=([\\d.]+))?\\;version:\\1"
        ],
        "website": "https://www.mangeznotez.com"
      },
      "MantisBT": {
        "cats": [
          13
        ],
        "cpe": "cpe:/a:mantisbt:mantisbt",
        "html": "<img[^>]+ alt=\"Powered by Mantis Bugtracker",
        "icon": "MantisBT.png",
        "implies": "PHP",
        "website": "http://www.mantisbt.org"
      },
      "ManyChat": {
        "cats": [
          32,
          52
        ],
        "description": "ManyChat is a service that allows you to create chatbots for Facebook Messenger.",
        "icon": "ManyChat.svg",
        "js": {
          "mcwidget": ""
        },
        "scripts": "widget\\.manychat\\.com",
        "website": "https://manychat.com/"
      },
      "ManyContacts": {
        "cats": [
          32,
          5
        ],
        "description": "ManyContacts is an attention-grabbing contact form sitting on top of your website that helps to increase conversion by converting visitors into leads.",
        "icon": "ManyContacts.png",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.manycontacts\\.com",
        "website": "https://www.manycontacts.com"
      },
      "Mapbox GL JS": {
        "cats": [
          35
        ],
        "description": "Mapbox GL JS is a JavaScript library that uses WebGL to render interactive maps from vector tiles and Mapbox styles.",
        "dom": "link[href*='mapbox-gl.css']",
        "icon": "Mapbogljs.png",
        "js": {
          "mapboxgl.version": "^(.+)$\\;version:\\1\\;confidence:0"
        },
        "scripts": "mapbox-gl.js",
        "website": "https://github.com/mapbox/mapbox-gl-js"
      },
      "MariaDB": {
        "cats": [
          34
        ],
        "cpe": "cpe:/a:mariadb_project:mariadb",
        "description": "MariaDB is an open-source relational database management system compatible with MySQL.",
        "icon": "mariadb.svg",
        "website": "https://mariadb.org"
      },
      "Marionette.js": {
        "cats": [
          12
        ],
        "description": "Marionette is a composite application library for Backbone.js that aims to simplify the construction of large scale JavaScript applications. It is a collection of common design and implementation patterns found in applications.",
        "icon": "Marionette.js.svg",
        "implies": [
          "Underscore.js",
          "Backbone.js"
        ],
        "js": {
          "Marionette": "",
          "Marionette.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "backbone\\.marionette.*\\.js",
        "website": "https://marionettejs.com"
      },
      "Marked": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:marked_project:marked",
        "icon": "marked.svg",
        "js": {
          "marked": ""
        },
        "scripts": "/marked(?:\\.min)?\\.js",
        "website": "https://marked.js.org"
      },
      "Marketo": {
        "cats": [
          32
        ],
        "description": "Marketo develops and sells marketing automation software for account-based marketing and other marketing services and products including SEO and content creation.",
        "icon": "Marketo.png",
        "js": {
          "Munchkin": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "munchkin\\.marketo\\.\\w+/(?:([\\d.]+)/)?munchkin\\.js\\;version:\\1",
        "website": "https://www.marketo.com"
      },
      "Marketo Forms": {
        "cats": [
          5
        ],
        "description": "Marketo Forms help create web forms without programming knowledge. Forms can reside on Marketo landing pages and also be embedded on any page of website.",
        "icon": "Marketo.png",
        "js": {
          "formatMarketoForm": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "marketo\\.\\w+/js/forms(?:[\\d.]+)/js/forms([\\d.]+)\\.min\\.js\\;version:\\1",
          "v([\\d.]+)/js/marketo-alt-form\\.min\\.js\\;version:\\1"
        ],
        "website": "https://www.marketo.com"
      },
      "Mastercard": {
        "cats": [
          41
        ],
        "description": "MasterCard facilitates electronic funds transfers throughout the world, most commonly through branded credit cards, debit cards and prepaid cards.",
        "html": "<[^>]+aria-labelledby=\"pi-mastercard",
        "icon": "Mastercard.svg",
        "website": "https://www.mastercard.com"
      },
      "MasterkinG32 Framework": {
        "cats": [
          18
        ],
        "description": "MasterkinG32 framework.",
        "headers": {
          "X-Powered-Framework": "MasterkinG(?:)"
        },
        "icon": "Masterking32.png",
        "meta": {
          "generator": "^MasterkinG(?:)"
        },
        "website": "https://masterking32.com"
      },
      "Mastodon": {
        "cats": [
          2
        ],
        "cookies": {
          "_mastodon_session": ""
        },
        "description": "Mastodon is a free and open-source self-hosted social networking service.",
        "headers": {
          "Server": "^Mastodon$"
        },
        "icon": "Mastodon.svg",
        "website": "https://joinmastodon.org"
      },
      "Material Design Lite": {
        "cats": [
          66
        ],
        "description": "Material Design Lite is a library of components for web developers.",
        "html": "<link[^>]* href=\"[^\"]*material(?:\\.[\\w]+-[\\w]+)?(?:\\.min)?\\.css",
        "icon": "Material Design Lite.png",
        "js": {
          "MaterialIconToggle": ""
        },
        "scripts": "(?:/([\\d.]+))?/material(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://getmdl.io"
      },
      "Materialize CSS": {
        "cats": [
          66
        ],
        "description": "Materialize CSS is a css framework which is used to create responsive websites.",
        "html": "<link[^>]* href=\"[^\"]*materialize(?:\\.min)?\\.css",
        "icon": "Materialize CSS.png",
        "scripts": "materialize(?:\\.min)?\\.js",
        "website": "http://materializecss.com"
      },
      "MathJax": {
        "cats": [
          25
        ],
        "description": "MathJax is a cross-browser JavaScript library that displays mathematical notation in web browsers, using MathML, LaTeX and ASCIIMathML markup.",
        "icon": "MathJax.png",
        "js": {
          "MathJax": "",
          "MathJax.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "([\\d.]+)?/mathjax\\.js\\;version:\\1",
        "website": "https://www.mathjax.org"
      },
      "Matomo Analytics": {
        "cats": [
          10
        ],
        "cookies": {
          "PIWIK_SESSID": ""
        },
        "cpe": "cpe:/a:matomo:matomo",
        "description": "Matomo Analytics is a free and open-source web analytics applicatio, that runs on a PHP/MySQL web-server.",
        "icon": "Matomo.png",
        "js": {
          "Matomo": "",
          "Piwik": "",
          "_paq": ""
        },
        "meta": {
          "apple-itunes-app": "app-id=737216887",
          "generator": "(?:Matomo|Piwik) - Open Source Web Analytics",
          "google-play-app": "app-id=org\\.piwik\\.mobile2"
        },
        "scripts": "piwik\\.js|piwik\\.php",
        "website": "https://matomo.org"
      },
      "Matomo Tag Manager": {
        "cats": [
          42
        ],
        "description": "Matomo Tag Manager manages tracking and marketing tags.",
        "icon": "Matomo.png",
        "js": {
          "MatomoTagManager": ""
        },
        "website": "https://developer.matomo.org/guides/tagmanager/introduction"
      },
      "Mattermost": {
        "cats": [
          2
        ],
        "cpe": "cpe:/a:jenkins:mattermost",
        "html": "<noscript> To use Mattermost, please enable JavaScript\\. </noscript>",
        "icon": "mattermost.png",
        "implies": [
          "Go",
          "React"
        ],
        "js": {
          "mm_config": "",
          "mm_current_user_id": "",
          "mm_license": "",
          "mm_user": ""
        },
        "website": "https://about.mattermost.com"
      },
      "Mautic": {
        "cats": [
          32
        ],
        "cpe": "cpe:/a:mautic:mautic",
        "description": "Mautic is a free and open-source marketing automation tool for Content Management, Social Media, Email Marketing, and can be used for the integration of social networks, campaign management, forms, questionnaires, reports, etc.",
        "icon": "mautic.svg",
        "js": {
          "MauticTrackingObject": ""
        },
        "scripts": "[^a-z]mtc.*\\.js",
        "website": "https://www.mautic.org/"
      },
      "MaxCDN": {
        "cats": [
          31
        ],
        "description": "MaxCDN is a content delivery network, which accelerates web-sites and decreases the server load.",
        "headers": {
          "Server": "^NetDNA",
          "X-CDN-Forward": "^maxcdn$"
        },
        "icon": "MaxCDN.png",
        "website": "http://www.maxcdn.com"
      },
      "MaxMind": {
        "cats": [
          79,
          83
        ],
        "description": "MaxMind is a provider of geolocation and online fraud detection tools.",
        "icon": "MaxMind.png",
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "scripts": [
          "[device|js]\\.maxmind\\.com/",
          "geoip\\.maxmind\\.min\\.js"
        ],
        "website": "https://www.maxmind.com",
        "xhr": "\\.maxmind\\.com"
      },
      "MaxSite CMS": {
        "cats": [
          1
        ],
        "icon": "MaxSite CMS.png",
        "implies": "PHP",
        "meta": {
          "generator": "MaxSite CMS"
        },
        "website": "http://max-3000.com"
      },
      "Maxemail": {
        "cats": [
          32
        ],
        "icon": "Maxemail.svg",
        "js": {
          "Mxm.Basket": "",
          "Mxm.FormHandler": "",
          "Mxm.Tracker": ""
        },
        "website": "https://maxemail.xtremepush.com"
      },
      "MediaElement.js": {
        "cats": [
          14
        ],
        "description": "MediaElement.js is a set of custom Flash plugins that mimic the HTML5 MediaElement API for browsers that don't support HTML5 or don't support the media codecs.",
        "icon": "MediaElement.js.png",
        "js": {
          "mejs": "",
          "mejs.version": "^(.+)$\\;version:\\1"
        },
        "website": "http://www.mediaelementjs.com"
      },
      "MediaWiki": {
        "cats": [
          8
        ],
        "cpe": "cpe:/a:mediawiki:mediawiki",
        "description": "MediaWiki is a free and open-source wiki engine.",
        "html": [
          "<body[^>]+class=\"mediawiki\"",
          "<(?:a|img)[^>]+>Powered by MediaWiki</a>",
          "<a[^>]+/Special:WhatLinksHere/"
        ],
        "icon": "MediaWiki.svg",
        "implies": "PHP",
        "js": {
          "mw.util.toggleToc": "",
          "wgTitle": "",
          "wgVersion": "^(.+)$\\;version:\\1"
        },
        "meta": {
          "generator": "^MediaWiki ?(.+)$\\;version:\\1"
        },
        "website": "https://www.mediawiki.org"
      },
      "Medium": {
        "cats": [
          11
        ],
        "description": "Medium is an online publishing platform.",
        "headers": {
          "X-Powered-By": "^Medium$"
        },
        "icon": "Medium.svg",
        "implies": "Node.js",
        "scripts": "medium\\.com",
        "url": "^https?://(?:www\\.)?medium\\.com",
        "website": "https://medium.com"
      },
      "Meebo": {
        "cats": [
          5
        ],
        "html": "(?:<iframe id=\"meebo-iframe\"|Meebo\\('domReady'\\))",
        "icon": "Meebo.png",
        "website": "http://www.meebo.com"
      },
      "Meeting Scheduler": {
        "cats": [
          5,
          72
        ],
        "description": "Meeting Scheduler is a schedule appointments widget.",
        "dom": "a[href*='bookmenow.info/book']",
        "icon": "Meeting Scheduler.png",
        "scripts": "bookmenow\\.info/(?:runtime|main).+\\.js",
        "website": "https://bookmenow.info"
      },
      "Melis Platform": {
        "cats": [
          1,
          6,
          11,
          32
        ],
        "cpe": "cpe:/a:melisplatform:melisplatform",
        "html": [
          "<!-- Rendered with Melis CMS V2",
          "<!-- Rendered with Melis Platform"
        ],
        "icon": "melis-platform.svg",
        "implies": [
          "Apache",
          "PHP",
          "MySQL",
          "Symfony",
          "Laravel",
          "Zend"
        ],
        "meta": {
          "generator": "^Melis Platform\\.",
          "powered-by": "^Melis CMS\\."
        },
        "website": "https://www.melistechnology.com/"
      },
      "MemberStack": {
        "cats": [
          6,
          47
        ],
        "cookies": {
          "memberstack": ""
        },
        "description": "MemberStack is a no-code membership platform for Webflow.",
        "icon": "MemberStack.png",
        "js": {
          "MemberStack": ""
        },
        "scripts": "memberstack\\.js",
        "url": "^https?//.+\\.memberstack\\.io",
        "website": "https://www.memberstack.io"
      },
      "Mercado Shops": {
        "cats": [
          6
        ],
        "cookies": {
          "_mshops_ga_gid": ""
        },
        "description": "Mercado Shops is an all-in-one ecommerce platform.",
        "icon": "Mercado Shops.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "frontend-assets/mshops-web-home/vendor",
        "website": "https://www.mercadoshops.com"
      },
      "Mermaid": {
        "cats": [
          25
        ],
        "html": "<div [^>]*class=[\"']mermaid[\"']>\\;confidence:90",
        "js": {
          "mermaid": ""
        },
        "scripts": "/mermaid(?:\\.min)?\\.js",
        "website": "https://mermaidjs.github.io/"
      },
      "Meteor": {
        "cats": [
          12,
          18
        ],
        "html": "<link[^>]+__meteor-css__",
        "icon": "Meteor.png",
        "implies": [
          "MongoDB",
          "Node.js"
        ],
        "js": {
          "Meteor": "",
          "Meteor.release": "^METEOR@([\\d.]+)\\;version:\\1"
        },
        "website": "https://www.meteor.com"
      },
      "Methode": {
        "cats": [
          1
        ],
        "html": "<!-- Methode uuid: \"[a-f\\d]+\" ?-->",
        "icon": "Methode.png",
        "meta": {
          "eomportal-id": "\\d+",
          "eomportal-instanceid": "\\d+",
          "eomportal-lastUpdate": "",
          "eomportal-loid": "[\\d.]+",
          "eomportal-uuid": "[a-f\\d]+"
        },
        "website": "https://www.eidosmedia.com/"
      },
      "Metomic": {
        "cats": [
          67
        ],
        "icon": "metomic.png",
        "scripts": [
          "metomic\\.js"
        ],
        "website": "https://metomic.io"
      },
      "Microsoft 365": {
        "cats": [
          30,
          75
        ],
        "description": "Microsoft 365 is a line of subscription services offered by Microsoft as part of the Microsoft Office product line.",
        "dns": {
          "MX": [
            "outlook\\.com"
          ]
        },
        "icon": "Microsoft 365.svg",
        "website": "https://www.microsoft.com/microsoft-365"
      },
      "Microsoft ASP.NET": {
        "cats": [
          18
        ],
        "cookies": {
          "ASP.NET_SessionId": "",
          "ASPSESSION": ""
        },
        "cpe": "cpe:/a:microsoft:asp.net",
        "description": "ASP.NET is an open-source, server-side web-application framework designed for web development to produce dynamic web pages.",
        "headers": {
          "X-AspNet-Version": "(.+)\\;version:\\1",
          "X-Powered-By": "^ASP\\.NET"
        },
        "html": "<input[^>]+name=\"__VIEWSTATE",
        "icon": "Microsoft ASP.NET.png",
        "implies": "IIS\\;confidence:50",
        "url": "\\.aspx?(?:$|\\?)",
        "website": "https://www.asp.net"
      },
      "Microsoft Advertising": {
        "cats": [
          36
        ],
        "cookies": {
          "_uetsid": "\\w+",
          "_uetvid": "\\w+"
        },
        "description": "Microsoft Advertising is an online advertising platform developed by Microsoft.",
        "icon": "Microsoft.png",
        "js": {
          "UET": "",
          "uetq": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "bat\\.bing\\.com/bat\\.js",
        "website": "https://ads.microsoft.com"
      },
      "Microsoft Clarity": {
        "cats": [
          10
        ],
        "description": "Microsoft's Clarity is a analytics tool which provides website usage statistics, session recording, and heatmaps.",
        "icon": "Microsoft.png",
        "js": {
          "clarity": ""
        },
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "www\\.clarity\\.ms/.+/([\\d.]+)/clarity\\.js\\;version:\\1",
        "website": "https://clarity.microsoft.com"
      },
      "Microsoft Excel": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:microsoft:excel",
        "description": "Microsoft Excel is an electronic spreadsheet program used for storing, organizing, and manipulating data.",
        "html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:excel\"|<!--\\s*(?:START|END) OF OUTPUT FROM EXCEL PUBLISH AS WEB PAGE WIZARD\\s*-->|<div [^>]*x:publishsource=\"?Excel\"?)",
        "icon": "Microsoft Excel.svg",
        "meta": {
          "ProgId": "^Excel\\.",
          "generator": "Microsoft Excel( [\\d.]+)?\\;version:\\1"
        },
        "website": "https://office.microsoft.com/excel"
      },
      "Microsoft HTTPAPI": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Microsoft-HTTPAPI(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Microsoft.png",
        "website": "http://microsoft.com"
      },
      "Microsoft PowerPoint": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:microsoft:powerpoint",
        "description": "Microsoft PowerPoint is a tool to create presentations using simple drag and drop tools.",
        "html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:powerpoint\"|<link rel=\"?Presentation-XML\"? href=\"?[^\"]+\\.xml\"?>|<o:PresentationFormat>[^<]+</o:PresentationFormat>[^!]+<o:Slides>\\d+</o:Slides>(?:[^!]+<o:Version>([\\d.]+)</o:Version>)?)\\;version:\\1",
        "icon": "Microsoft PowerPoint.svg",
        "meta": {
          "ProgId": "^PowerPoint\\.",
          "generator": "Microsoft PowerPoint ( [\\d.]+)?\\;version:\\1"
        },
        "website": "https://office.microsoft.com/powerpoint"
      },
      "Microsoft Publisher": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:microsoft:publisher",
        "description": "Microsoft Publisher is an application that allows you to create professional documents such as newsletters, postcards, flyers, invitations, brochures.",
        "html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:publisher\"|<!--[if pub]><xml>)",
        "icon": "Microsoft Publisher.svg",
        "meta": {
          "ProgId": "^Publisher\\.",
          "generator": "Microsoft Publisher( [\\d.]+)?\\;version:\\1"
        },
        "website": "https://office.microsoft.com/publisher"
      },
      "Microsoft SharePoint": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:microsoft:sharepoint_server",
        "description": "SharePoint is a cloud-based collaborative platform to manage and share content.",
        "headers": {
          "MicrosoftSharePointTeamServices": "^(.+)$\\;version:\\1",
          "SPRequestGuid": "",
          "SharePointHealthScore": "",
          "X-SharePointHealthScore": ""
        },
        "icon": "Microsoft SharePoint.png",
        "js": {
          "SPDesignerProgID": "",
          "_spBodyOnLoadCalled": ""
        },
        "meta": {
          "generator": "Microsoft SharePoint"
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "https://sharepoint.microsoft.com"
      },
      "Microsoft Word": {
        "cats": [
          20
        ],
        "cpe": "cpe:/a:microsoft:word",
        "description": "MS Word is a word-processing program used primarily for creating documents.",
        "html": "(?:<html [^>]*xmlns:w=\"urn:schemas-microsoft-com:office:word\"|<w:WordDocument>|<div [^>]*class=\"?WordSection1[\" >]|<style[^>]*>[^>]*@page WordSection1)",
        "icon": "Microsoft Word.svg",
        "meta": {
          "ProgId": "^Word\\.",
          "generator": "Microsoft Word( [\\d.]+)?\\;version:\\1"
        },
        "website": "https://office.microsoft.com/word"
      },
      "Miestro": {
        "cats": [
          6
        ],
        "description": "Miestro is an all-in-one ecommerce platform wich allow create online course and membership sites.",
        "icon": "Miestro.svg",
        "meta": {
          "base_url": ".+\\.miestro\\.com"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.miestro\\.com",
        "website": "https://miestro.com"
      },
      "Milligram": {
        "cats": [
          66
        ],
        "html": [
          "<link[^>]+?href=\"[^\"]+milligram(?:\\.min)?\\.css"
        ],
        "icon": "Milligram.png",
        "website": "https://milligram.io"
      },
      "MindBody": {
        "cats": [
          5,
          72
        ],
        "description": "Mindbody is a (SaaS) company that provides cloud-based online scheduling and other business management software for the wellness services industry.",
        "icon": "MindBody.svg",
        "js": {
          "HealcodeWidget": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\w+\\.healcode\\.com",
        "website": "https://www.mindbodyonline.com"
      },
      "Minero.cc": {
        "cats": [
          56
        ],
        "description": "Minero.cc is a bot that helps run crypto mining application.",
        "scripts": [
          "//minero\\.cc/lib/minero(?:-miner|-hidden)?\\.min\\.js"
        ],
        "website": "http://minero.cc/"
      },
      "Mint": {
        "cats": [
          10
        ],
        "icon": "Mint.png",
        "js": {
          "Mint": ""
        },
        "scripts": "mint/\\?js",
        "website": "https://haveamint.com"
      },
      "Misskey": {
        "cats": [
          2
        ],
        "description": "Misskey is a distributed microblogging platform.",
        "html": "<!-- Thank you for using Misskey! @syuilo -->",
        "icon": "Misskey.svg",
        "meta": {
          "application-name": "Misskey"
        },
        "website": "https://join.misskey.page/"
      },
      "Miva": {
        "cats": [
          6
        ],
        "headers": {
          "content-disposition": "filename=(?:mvga\\.js|MivaEvents\\.js)"
        },
        "icon": "miva.png",
        "js": {
          "MivaVM_API": "",
          "MivaVM_Version": "^(.+)$\\;version:\\1",
          "mivaJS": "",
          "mivaJS.Page": "",
          "mivaJS.Product_Code": "",
          "mivaJS.Product_ID": "",
          "mivaJS.Screen": "",
          "mivaJS.Store_Code": ""
        },
        "scripts": "mvga\\.js",
        "website": "http://www.miva.com"
      },
      "Mixpanel": {
        "cats": [
          10
        ],
        "description": "Mixpanel provides a business analytics service. It tracks user interactions with web and mobile applications and provides tools for targeted communication with them. Its toolset contains in-app A/B tests and user survey forms.",
        "dns": {
          "TXT": [
            "mixpanel-domain-verify"
          ]
        },
        "icon": "Mixpanel.png",
        "js": {
          "mixpanel": ""
        },
        "scripts": [
          "cdn\\.mxpnl\\.com/libs/mixpanel\\-([0-9.]+)\\.min\\.js\\;version:\\1",
          "api\\.mixpanel\\.com/track"
        ],
        "website": "https://mixpanel.com"
      },
      "MkDocs": {
        "cats": [
          4
        ],
        "description": "MkDocs is a static site generator, used for building project documentation.",
        "icon": "mkdocs.png",
        "meta": {
          "generator": "^mkdocs-([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.mkdocs.org/"
      },
      "MoEngage": {
        "cats": [
          32,
          10
        ],
        "description": "MoEngage is an intelligent customer engagement platform for the customer-obsessed marketer.",
        "icon": "MoEngage.png",
        "js": {
          "MOENGAGE_API_KEY": "",
          "Moengage": "",
          "downloadMoengage": "",
          "moengage_object": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.moengage\\.\\w+",
        "website": "https://www.moengage.com"
      },
      "Moat": {
        "cats": [
          10
        ],
        "description": "Moat is a digital ad analytics tool.",
        "icon": "Moat.svg",
        "scripts": [
          "moatads\\.com"
        ],
        "website": "https://moat.com/"
      },
      "MobX": {
        "cats": [
          59
        ],
        "icon": "MobX.svg",
        "js": {
          "__mobxGlobal": "",
          "__mobxGlobals": "",
          "__mobxInstanceCount": ""
        },
        "scripts": "(?:/([\\d\\.]+))?/mobx(?:\\.[a-z]+){0,2}\\.js(?:$|\\?)\\;version:\\1",
        "website": "https://mobx.js.org"
      },
      "Mobify": {
        "cats": [
          6,
          26
        ],
        "description": "Mobify is a web storefront platform for headless commerce.",
        "headers": {
          "X-Powered-By": "Mobify"
        },
        "icon": "Mobify.png",
        "js": {
          "Mobify": ""
        },
        "scripts": [
          "//cdn\\.mobify\\.com/",
          "//a\\.mobify\\.com/"
        ],
        "website": "https://www.mobify.com"
      },
      "Mobirise": {
        "cats": [
          51
        ],
        "description": "Mobirise is a free offline app for Windows and Mac to easily create small/medium websites, landing pages, online resumes and portfolios.",
        "html": [
          "<!-- Site made with Mobirise Website Builder v([\\d.]+)\\;version:\\1"
        ],
        "icon": "mobirise.png",
        "meta": {
          "generator": "^Mobirise v([\\d.]+)\\;version:\\1"
        },
        "website": "https://mobirise.com"
      },
      "MochiKit": {
        "cats": [
          59
        ],
        "icon": "MochiKit.png",
        "js": {
          "MochiKit": "",
          "MochiKit.MochiKit.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "MochiKit(?:\\.min)?\\.js",
        "website": "https://mochi.github.io/mochikit/"
      },
      "MochiWeb": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:mochiweb_project:mochiweb",
        "headers": {
          "Server": "MochiWeb(?:/([\\d.]+))?\\;version:\\1"
        },
        "website": "https://github.com/mochi/mochiweb"
      },
      "Modernizr": {
        "cats": [
          59
        ],
        "description": "Modernizr is a JavaScript library that detects the features available in a user's browser.",
        "icon": "Modernizr.svg",
        "js": {
          "Modernizr._version": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "([\\d.]+)?/modernizr(?:\\.([\\d.]+))?.*\\.js\\;version:\\1?\\1:\\2"
        ],
        "website": "https://modernizr.com"
      },
      "Modified": {
        "cats": [
          6
        ],
        "icon": "modified.png",
        "meta": {
          "generator": "\\(c\\) by modified eCommerce Shopsoftware ------ http://www\\.modified-shop\\.org"
        },
        "website": "http://www.modified-shop.org/"
      },
      "Moguta.CMS": {
        "cats": [
          1,
          6
        ],
        "html": "<link[^>]+href=[\"'][^\"]+mg-(?:core|plugins|templates)/",
        "icon": "Moguta.CMS.png",
        "implies": "PHP",
        "scripts": "mg-(?:core|plugins|templates)/",
        "website": "https://moguta.ru"
      },
      "MoinMoin": {
        "cats": [
          8
        ],
        "cookies": {
          "MOIN_SESSION": ""
        },
        "cpe": "cpe:/a:moinmo:moinmoin",
        "description": "MoinMoin is a wiki engine implemented in Python.",
        "icon": "MoinMoin.png",
        "implies": "Python",
        "js": {
          "show_switch2gui": ""
        },
        "scripts": "moin(?:_static(\\d)(\\d)(\\d)|.+)/common/js/common\\.js\\;version:\\1.\\2.\\3",
        "website": "https://moinmo.in"
      },
      "Mojolicious": {
        "cats": [
          18
        ],
        "headers": {
          "server": "^mojolicious",
          "x-powered-by": "mojolicious"
        },
        "icon": "Mojolicious.png",
        "implies": "Perl",
        "website": "http://mojolicio.us"
      },
      "Mollom": {
        "cats": [
          16
        ],
        "cpe": "cpe:/a:acquia:mollom",
        "html": "<img[^>]+\\.mollom\\.com",
        "icon": "Mollom.png",
        "scripts": "mollom(?:\\.min)?\\.js",
        "website": "http://mollom.com"
      },
      "Moment Timezone": {
        "cats": [
          59
        ],
        "icon": "Moment.js.svg",
        "implies": "Moment.js",
        "scripts": "moment-timezone(?:-data)?(?:\\.min)?\\.js",
        "website": "http://momentjs.com/timezone/"
      },
      "Moment.js": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:momentjs:moment",
        "description": "Moment.js is a free and open-source JavaScript library that removes the need to use the native JavaScript Date object directly.",
        "icon": "Moment.js.svg",
        "js": {
          "moment": "",
          "moment.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "moment(?:\\.min)?\\.js",
        "website": "https://momentjs.com"
      },
      "Mondo Media": {
        "cats": [
          6
        ],
        "icon": "Mondo Media.png",
        "meta": {
          "generator": "Mondo Shop"
        },
        "website": "http://mondo-media.de"
      },
      "MongoDB": {
        "cats": [
          34
        ],
        "cpe": "cpe:/a:mongodb:mongodb",
        "description": "MongoDB is a document-oriented NoSQL database used for high volume data storage.",
        "icon": "MongoDB.png",
        "website": "http://www.mongodb.org"
      },
      "Mongrel": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:zed_shaw:mongrel",
        "headers": {
          "Server": "Mongrel"
        },
        "icon": "Mongrel.png",
        "implies": "Ruby",
        "website": "http://mongrel2.org"
      },
      "Monkey HTTP Server": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Monkey/?([\\d.]+)?\\;version:\\1"
        },
        "icon": "Monkey HTTP Server.png",
        "website": "http://monkey-project.com"
      },
      "Mono": {
        "cats": [
          18
        ],
        "cpe": "cpe:/a:mono:mono",
        "headers": {
          "X-Powered-By": "Mono"
        },
        "icon": "Mono.png",
        "website": "http://mono-project.com"
      },
      "Mono.net": {
        "cats": [
          1
        ],
        "icon": "Mono.net.png",
        "implies": "Matomo Analytics",
        "js": {
          "_monoTracker": ""
        },
        "scripts": "monotracker(?:\\.min)?\\.js",
        "website": "https://www.mono.net/en"
      },
      "MooTools": {
        "cats": [
          12
        ],
        "icon": "MooTools.png",
        "js": {
          "MooTools": "",
          "MooTools.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "mootools.*\\.js",
        "website": "https://mootools.net"
      },
      "Moodle": {
        "cats": [
          21
        ],
        "cookies": {
          "MOODLEID_": "",
          "MoodleSession": ""
        },
        "cpe": "cpe:/a:moodle:moodle",
        "description": "Moodle is a free and open-source Learning Management System (LMS) written in PHP and distributed under the GNU General Public License.",
        "html": "<img[^>]+moodlelogo",
        "icon": "Moodle.png",
        "implies": "PHP",
        "js": {
          "M.core": "",
          "Y.Moodle": ""
        },
        "meta": {
          "keywords": "^moodle"
        },
        "website": "http://moodle.org"
      },
      "Moon": {
        "cats": [
          12
        ],
        "icon": "moon.svg",
        "scripts": "/moon(?:\\.min)?\\.js$",
        "website": "https://kbrsh.github.io/moon/"
      },
      "MotoCMS": {
        "cats": [
          1
        ],
        "html": "<link [^>]*href=\"[^>]*\\/mt-content\\/[^>]*\\.css",
        "icon": "MotoCMS.svg",
        "implies": [
          "PHP",
          "AngularJS",
          "jQuery"
        ],
        "scripts": "/mt-includes/js/website(?:assets)?\\.(?:min)?\\.js",
        "website": "http://motocms.com"
      },
      "Mouse Flow": {
        "cats": [
          10
        ],
        "icon": "mouseflow.png",
        "js": {
          "_mfq": ""
        },
        "scripts": [
          "cdn\\.mouseflow\\.com"
        ],
        "website": "https://mouseflow.com/"
      },
      "Movable Ink": {
        "cats": [
          76
        ],
        "description": "Movable Ink is a technology company that allows marketers to change emails after they have been sent out.",
        "icon": "Movable Ink.png",
        "js": {
          "MovableInkTrack": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://movableink.com"
      },
      "Movable Type": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:sixapart:movable_type",
        "icon": "Movable Type.svg",
        "meta": {
          "generator": "Movable Type"
        },
        "website": "http://movabletype.org"
      },
      "Mozard Suite": {
        "cats": [
          1
        ],
        "icon": "Mozard Suite.png",
        "meta": {
          "author": "Mozard"
        },
        "url": "/mozard/!suite",
        "website": "http://mozard.nl"
      },
      "Mura CMS": {
        "cats": [
          1,
          11
        ],
        "icon": "Mura CMS.png",
        "implies": "Adobe ColdFusion",
        "meta": {
          "generator": "Mura CMS ([\\d]+)\\;version:\\1"
        },
        "website": "http://www.getmura.com"
      },
      "Mustache": {
        "cats": [
          12
        ],
        "description": "Mustache is a web template system.",
        "icon": "Mustache.png",
        "js": {
          "Mustache.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "mustache(?:\\.min)?\\.js",
        "website": "https://mustache.github.io"
      },
      "My Food Link": {
        "cats": [
          6
        ],
        "html": [
          "<div class='mfl-made-by-myfoodlink'>",
          "<a href=\"https://www.myfoodlink.com.au"
        ],
        "icon": "myfoodlink.png",
        "website": "https://www.myfoodlink.com.au/"
      },
      "MyBB": {
        "cats": [
          2
        ],
        "cpe": "cpe:/a:mybb:mybb",
        "description": "MyBB is a free and open-source forum software written in PHP.",
        "html": "(?:<script [^>]+\\s+<!--\\s+lang\\.no_new_posts|<a[^>]* title=\"Powered By MyBB)",
        "icon": "MyBB.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "MyBB": ""
        },
        "website": "https://mybb.com"
      },
      "MyBlogLog": {
        "cats": [
          5
        ],
        "icon": "MyBlogLog.png",
        "scripts": "pub\\.mybloglog\\.com",
        "website": "http://www.mybloglog.com"
      },
      "MyCashFlow": {
        "cats": [
          6
        ],
        "headers": {
          "X-MCF-ID": ""
        },
        "icon": "mycashflow.png",
        "website": "https://www.mycashflow.fi/"
      },
      "MyLiveChat": {
        "cats": [
          52
        ],
        "description": "MyLiveChat is a live chat developed by CuteSoft.",
        "icon": "MyLiveChat.png",
        "js": {
          "MyLiveChat.Version": "(.+)$\\;version:\\1"
        },
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "scripts": "mylivechat\\.com/",
        "website": "https://mylivechat.com"
      },
      "MySQL": {
        "cats": [
          34
        ],
        "cpe": "cpe:/a:mysql:mysql",
        "description": "MySQL is an open-source relational database management system.",
        "icon": "MySQL.svg",
        "website": "http://mysql.com"
      },
      "MyWebsite": {
        "cats": [
          1
        ],
        "description": "MyWebsite is website builder designed for easy editing and fast results.",
        "icon": "Ionos-by-1and1-logo.svg",
        "js": {
          "SystemID": "^.*1AND1.*$",
          "version": "^(.*)$\\;version:\\1\\;confidence:0"
        },
        "meta": {
          "generator": "^.*MyWebsite.*$\\;version:8"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\/\\/integration\\.mywebsite-editor\\.com.*\\.js\\;version:9",
        "website": "https://www.ionos.com"
      },
      "Mynetcap": {
        "cats": [
          1
        ],
        "icon": "Mynetcap.png",
        "meta": {
          "generator": "Mynetcap"
        },
        "website": "http://www.netcap-creation.fr"
      },
      "NEO - Omnichannel Commerce Platform": {
        "cats": [
          6
        ],
        "description": "NEO is an ecommerce software that manages multiple online stores.",
        "dom": "#svr[value^=\"NEOWEBV\"]",
        "headers": {
          "powered": "jet-neo"
        },
        "icon": "Plataforma NEO.svg",
        "url": "\\.plataformaneo\\.com",
        "website": "https://www.jetecommerce.com.br"
      },
      "NSW Design System": {
        "cats": [
          66
        ],
        "dom": ".nsw-container, .nsw-header, .nsw-icon, link[href*='nsw-design-system']",
        "icon": "NSW Design System.svg",
        "js": {
          "NSW.initSite": ""
        },
        "website": "https://www.digital.nsw.gov.au/digital-design-system"
      },
      "NTLM": {
        "cats": [
          16
        ],
        "description": "NTLM is an authentication method commonly used by Windows servers",
        "headers": {
          "WWW-Authenticate": "^NTLM"
        },
        "website": "https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-ntht/"
      },
      "NVD3": {
        "cats": [
          25
        ],
        "description": "NVD3 is a JavaScript visualisation library that is a free open-source tool.",
        "html": "<link[^>]* href=[^>]+nv\\.d3(?:\\.min)?\\.css",
        "icon": "NVD3.png",
        "implies": "D3",
        "js": {
          "nv.addGraph": "",
          "nv.version": "^(.+)$\\;confidence:0\\;version:\\1"
        },
        "scripts": "nv\\.d3(?:\\.min)?\\.js",
        "website": "http://nvd3.org"
      },
      "Najva": {
        "cats": [
          32
        ],
        "description": "Najva is a retention marketing solution that offers push notification and email marketing.",
        "icon": "Najva.png",
        "js": {
          "Najva.identifyEmailSubscriber": ""
        },
        "pricing": [
          "low",
          "freemium"
        ],
        "saas": true,
        "scripts": "app\\.najva\\.com/",
        "website": "https://www.najva.com"
      },
      "Navegg": {
        "cats": [
          10
        ],
        "icon": "Navegg.png",
        "scripts": "tag\\.navdmp\\.com",
        "website": "https://www.navegg.com/"
      },
      "Neos CMS": {
        "cats": [
          1
        ],
        "excludes": "TYPO3 CMS",
        "headers": {
          "X-Flow-Powered": "Neos/?(.+)?$\\;version:\\1"
        },
        "icon": "Neos.svg",
        "implies": "Neos Flow",
        "url": "/neos/",
        "website": "https://neos.io"
      },
      "Neos Flow": {
        "cats": [
          18
        ],
        "excludes": "TYPO3 CMS",
        "headers": {
          "X-Flow-Powered": "Flow/?(.+)?$\\;version:\\1"
        },
        "icon": "Neos.svg",
        "implies": "PHP",
        "website": "https://flow.neos.io"
      },
      "Nepso": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-CMS": "Nepso"
        },
        "icon": "nepso.svg",
        "website": "https://www.nepso.com"
      },
      "NetSuite": {
        "cats": [
          6
        ],
        "cookies": {
          "NS_VER": ""
        },
        "icon": "NetSuite.png",
        "website": "http://netsuite.com"
      },
      "Netlify": {
        "cats": [
          62,
          31
        ],
        "description": "Netlify providers hosting and server-less backend services for web applications and static websites.",
        "headers": {
          "Server": "^Netlify",
          "X-NF-Request-ID": ""
        },
        "icon": "Netlify.svg",
        "url": "^https?://[^/]+\\.netlify\\.(?:com|app)/",
        "website": "https://www.netlify.com/",
        "xhr": "cdn\\.netlify\\.com"
      },
      "Neto": {
        "cats": [
          6
        ],
        "icon": "Neto.svg",
        "js": {
          "NETO": ""
        },
        "scripts": "jquery\\.neto.*\\.js",
        "website": "https://www.neto.com.au"
      },
      "Nette Framework": {
        "cats": [
          18
        ],
        "cookies": {
          "nette-browser": ""
        },
        "headers": {
          "X-Powered-By": "^Nette Framework"
        },
        "html": [
          "<input[^>]+data-nette-rules",
          "<div[^>]+id=\"snippet-",
          "<input[^>]+id=\"frm-"
        ],
        "icon": "Nette Framework.png",
        "implies": "PHP",
        "js": {
          "Nette": "",
          "Nette.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://nette.org"
      },
      "New Relic": {
        "cats": [
          10
        ],
        "description": "New Relic is a SaaS offering that focuses on performance and availability monitoring.",
        "dom": "link[href*='.newrelic.com']",
        "icon": "New Relic.svg",
        "js": {
          "NREUM": "",
          "newrelic": ""
        },
        "pricing": [
          "freemium",
          "payg",
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://newrelic.com"
      },
      "Next.js": {
        "cats": [
          18,
          22
        ],
        "cpe": "cpe:/a:zeit:next.js",
        "description": "Next.js is a React framework for developing single page Javascript applications.",
        "headers": {
          "x-powered-by": "^Next\\.js ?([0-9.]+)?\\;version:\\1"
        },
        "icon": "vercel.svg",
        "implies": [
          "React",
          "webpack",
          "Node.js"
        ],
        "js": {
          "__NEXT_DATA__": ""
        },
        "website": "https://nextjs.org"
      },
      "NextGEN Gallery": {
        "cats": [
          7
        ],
        "cpe": "cpe:/a:imagely:nextgen_gallery",
        "description": "NextGEN Gallery is a free open-source image management plugin for the WordPress content management system.",
        "html": [
          "<!-- <meta name=\"NextGEN\" version=\"([\\d.]+)\" /> -->\\;version:\\1"
        ],
        "icon": "NextGEN Gallery.png",
        "implies": "WordPress",
        "scripts": "/nextgen-gallery/js/",
        "website": "https://www.imagely.com/wordpress-gallery-plugin"
      },
      "Nextsale": {
        "cats": [
          5,
          32
        ],
        "description": "Nextsale is a conversion optimisation platform that provides social proof and urgency tools for ecommerce websites.",
        "icon": "Nextsale.svg",
        "js": {
          "NextsaleObject": ""
        },
        "pricing": [
          "freemium",
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "(?:api|sdk)\\.nextsale\\.io/",
        "website": "https://nextsale.io"
      },
      "Nginx": {
        "cats": [
          22,
          64
        ],
        "cpe": "cpe:/a:nginx:nginx",
        "description": "Nginx is a web server that can also be used as a reverse proxy, load balancer, mail proxy and HTTP cache.",
        "headers": {
          "Server": "nginx(?:/([\\d.]+))?\\;version:\\1",
          "X-Fastcgi-Cache": ""
        },
        "icon": "Nginx.svg",
        "website": "http://nginx.org/en"
      },
      "NitroPack": {
        "cats": [
          23
        ],
        "description": "NitroPack creates optimised HTML cache for fast page loading experience.",
        "icon": "NitroPack.svg",
        "meta": {
          "generator": "NitroPack"
        },
        "website": "https://nitropack.io/"
      },
      "Node.js": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:nodejs:node.js",
        "description": "Node.js is an open-source, cross-platform, JavaScript runtime environment that executes JavaScript code outside a web browser.",
        "icon": "node.js.png",
        "website": "http://nodejs.org"
      },
      "NodeBB": {
        "cats": [
          2
        ],
        "cpe": "cpe:/a:nodebb:nodebb",
        "headers": {
          "X-Powered-By": "^NodeBB$"
        },
        "icon": "NodeBB.png",
        "implies": "Node.js",
        "scripts": "^/nodebb\\.min\\.js\\?",
        "website": "https://nodebb.org"
      },
      "Nosto": {
        "cats": [
          76,
          74
        ],
        "description": "Nosto is an ecommerce platform providing product recommendations based on individual behavioral data.",
        "icon": "Nosto.svg",
        "js": {
          "nosto": "\\;confidence:50",
          "nostojs": "\\;confidence:50"
        },
        "meta": {
          "nosto-version": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "connect\\.nosto\\.\\w+/",
        "website": "https://www.nosto.com"
      },
      "Nukeviet CMS": {
        "cats": [
          1
        ],
        "description": "NukeViet CMS is a Vietnamese content management system.",
        "icon": "Nukeviet CMS.png",
        "js": {
          "nv_DigitalClock": "\\;confidence:50",
          "nv_is_change_act_confirm": "\\;confidence:50"
        },
        "meta": {
          "generator": "NukeViet v([\\d.]+)\\;version:\\1"
        },
        "oss": true,
        "pricing": [
          "freemium"
        ],
        "website": "https://nukeviet.vn/en/"
      },
      "Nuqlium": {
        "cats": [
          76,
          29
        ],
        "description": "Nuqlium is an integrated cloud-based online merchandising platform.",
        "icon": "Nuqlium.png",
        "js": {
          "nuqliumObject": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.nuqlium\\.com/api",
        "website": "https://www.nuqlium.com"
      },
      "Nuvemshop": {
        "cats": [
          6
        ],
        "html": "<a target=\"_blank\" title=\"Nuvemshop\"",
        "icon": "nuvem.png",
        "scripts": "Nuvem",
        "website": "https://www.nuvemshop.com.br/"
      },
      "Nuxt.js": {
        "cats": [
          12,
          18,
          22,
          57
        ],
        "description": "Nuxt is a Vue framework for developing modern web applications.",
        "html": [
          "<div [^>]*id=\"__nuxt\"",
          "<script [^>]*>window\\.__NUXT__"
        ],
        "icon": "Nuxt.js.svg",
        "implies": [
          "Vue.js",
          "Node.js"
        ],
        "js": {
          "$nuxt": ""
        },
        "scripts": [
          "/_nuxt/"
        ],
        "website": "https://nuxtjs.org"
      },
      "OWL Carousel": {
        "cats": [
          5
        ],
        "description": "OWL Carousel is an enabled jQuery plugin that lets you create responsive carousel sliders.",
        "html": "<link [^>]*href=\"[^\"]+owl\\.carousel(?:\\.min)?\\.css",
        "icon": "OWL Carousel.png",
        "implies": "jQuery",
        "scripts": "owl\\.carousel.*\\.js",
        "website": "https://owlcarousel2.github.io/OwlCarousel2/"
      },
      "OXID eShop": {
        "cats": [
          6
        ],
        "cookies": {
          "sid_key": "oxid"
        },
        "description": "OXID eShop is a free, open source ecommerce solution built using object oriented programming and PHP.",
        "icon": "OXID eShop.svg",
        "implies": "PHP",
        "js": {
          "oxCookieNote": "",
          "oxInputValidator": "",
          "oxLoginBox": "",
          "oxMiniBasket": "",
          "oxModalPopup": "",
          "oxTopMenu": ""
        },
        "website": "https://www.oxid-esales.com"
      },
      "OXID eShop Community Edition": {
        "cats": [
          6
        ],
        "description": "OXID eShop Community Edition is a free, open source ecommerce solution built using object oriented programming and PHP.",
        "excludes": "OXID eShop",
        "html": "<!--[^-]*OXID eShop Community Edition, Version (\\d+)\\;version:\\1",
        "icon": "OXID eShop.svg",
        "implies": "PHP",
        "website": "https://www.oxid-esales.com"
      },
      "OXID eShop Enterprise Edition": {
        "cats": [
          6,
          62
        ],
        "description": "OXID eShop Enterprise Edition is a B2B or B2C ecommerce solution built using object oriented programming and PHP.",
        "excludes": "OXID eShop",
        "html": "<!--[^-]*OXID eShop Enterprise Edition, Version (\\d+)\\;version:\\1",
        "icon": "OXID eShop.svg",
        "implies": "PHP",
        "website": "https://www.oxid-esales.com"
      },
      "Occasion": {
        "cats": [
          72
        ],
        "description": "Occasion is an online booking system.",
        "dom": [
          "iframe[src*='app.getoccasion.com']",
          "a[href*='app.getoccasion.com']"
        ],
        "icon": "Occasion.png",
        "js": {
          "OCCSN.stack": "",
          "occsnMerchantToken": ""
        },
        "pricing": [
          "low",
          "payg",
          "recurring"
        ],
        "saas": true,
        "scripts": "app\\.getoccasion\\.com",
        "website": "https://www.getoccasion.com"
      },
      "Ochanoko": {
        "cats": [
          6
        ],
        "description": "Ochanoko is a ecommerce online shopping cart solutions, ecommerce web site hosting.",
        "icon": "Ochanoko.svg",
        "js": {
          "ocnkProducts": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "ocnk-min\\.js",
        "website": "https://www.ocnk.com"
      },
      "October CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "october_session": ""
        },
        "description": "October is a free, open-source, self-hosted CMS platform based on the Laravel PHP Framework.",
        "icon": "October CMS.png",
        "implies": "Laravel",
        "meta": {
          "generator": "OctoberCMS"
        },
        "oss": true,
        "website": "http://octobercms.com"
      },
      "Octopress": {
        "cats": [
          57
        ],
        "description": "Octopress is a static blogging framework.",
        "html": "Powered by <a href=\"http://octopress\\.org\">",
        "icon": "octopress.png",
        "implies": "Jekyll",
        "meta": {
          "generator": "Octopress"
        },
        "scripts": "/octopress\\.js",
        "website": "http://octopress.org"
      },
      "Odoo": {
        "cats": [
          1,
          6
        ],
        "cpe": "cpe:/a:odoo:odoo",
        "description": "Odoo is business management software which includes CRM, ecommerce, billing, accounting, manufacturing, warehouse, project management, and inventory management.",
        "html": "<link[^>]* href=[^>]+/web/css/(?:web\\.assets_common/|website\\.assets_frontend/)\\;confidence:25",
        "icon": "Odoo.png",
        "implies": [
          "Python",
          "PostgreSQL",
          "Node.js",
          "Less"
        ],
        "meta": {
          "generator": "Odoo"
        },
        "oss": true,
        "scripts": "/web/js/(?:web\\.assets_common/|website\\.assets_frontend/)\\;confidence:25",
        "website": "http://odoo.com"
      },
      "Okendo": {
        "cats": [
          5,
          76
        ],
        "description": "Okendo is a customer marketing platform with product ratings and reviews, customer photos and videos to help personalise experiences.",
        "dom": {
          "div.okeReviews": {
            "attributes": {
              "data-oke-reviews-version": "^([\\d.]+)$\\;version:\\1"
            }
          }
        },
        "icon": "Okendo.svg",
        "js": {
          "okeReviewsWidgetOnInit": "",
          "okeWidgetControlInit": "",
          "okendoReviews": ""
        },
        "pricing": [
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.okendo.io"
      },
      "Okta": {
        "cats": [
          19
        ],
        "description": "Okta is a platform in the Identity-as-a-Service (IDaaS) category. Okta features include Provisioning, Single Sign-On (SSO), Active Directory (AD) and LDAP integration, the centralized de-provisioning of users, multi-factor authentication (MFA), mobile identity management.",
        "icon": "Okta.svg",
        "js": {
          "OktaAuth": "",
          "isOktaEnabled": "",
          "oktaCurrentSessionUrl": ""
        },
        "pricing": [
          "poa",
          "freemium",
          "recurring"
        ],
        "scripts": "oktacdn\\.com/.+/([\\d.]+)/\\;version:\\1",
        "website": "https://developer.okta.com"
      },
      "Olark": {
        "cats": [
          52
        ],
        "description": "Olark is a cloud-based live chat solution.",
        "icon": "Olark.png",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.olark\\.com",
        "website": "https://www.olark.com/"
      },
      "Ometria": {
        "cats": [
          32
        ],
        "cookies": {
          "ometria": ""
        },
        "description": "Ometria is a customer insight and marketing automation platform.",
        "dom": "form[action*='api.ometria.com']",
        "icon": "Ometria.svg",
        "js": {
          "AddOmetriaBasket": "",
          "AddOmetriaIdentify": "",
          "ometria": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.ometria\\.com",
        "website": "https://ometria.com"
      },
      "Omise": {
        "cats": [
          41
        ],
        "description": "Omise is a payment gateway for Thailand, Japan and Singapore. Providing both online and offline payment solutions to merchants.",
        "icon": "Omise.svg",
        "js": {
          "Omise": "",
          "OmiseCard": ""
        },
        "pricing": [
          "payg"
        ],
        "scripts": "cdn\\.omise\\.co",
        "website": "https://www.omise.co"
      },
      "Omniconvert": {
        "cats": [
          74
        ],
        "description": "Omniconvert is an award-winning conversion rate optimisation (CRO) software that can be used for A/B testing, online surveys, traffic segmentation.",
        "dom": "link[href*='app.omniconvert.com']",
        "icon": "Omniconvert.png",
        "js": {
          "_omni": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "scripts": "cdn\\.omniconvert\\.com",
        "website": "https://www.omniconvert.com"
      },
      "Omnisend": {
        "cats": [
          32,
          6
        ],
        "cookies": {
          "omnisendSessionID": ""
        },
        "description": "Omnisend is an ecommerce marketing automation platform that provides an omnichannel marketing strategy for businesses.",
        "icon": "Omnisend.svg",
        "js": {
          "_omnisend": ""
        },
        "meta": {
          "omnisend-site-verification": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "omnisrc\\.com",
        "website": "https://www.omnisend.com"
      },
      "OneAPM": {
        "cats": [
          10
        ],
        "icon": "OneAPM.png",
        "js": {
          "BWEUM": ""
        },
        "website": "http://www.oneapm.com"
      },
      "OneAll": {
        "cats": [
          69
        ],
        "description": "OneAll is a social login solution enables your users to sign into their accounts on your website or mobile app using their login details from networking sites.",
        "icon": "OneAll.png",
        "js": {
          "oa_social_login": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "api\\.oneall\\.com/socialize",
        "website": "https://www.oneall.com"
      },
      "OneSignal": {
        "cats": [
          32,
          74
        ],
        "description": "OneSignal is a customer engagement messaging solution.",
        "icon": "OneSignal.svg",
        "js": {
          "OneSignal": "",
          "__oneSignalSdkLoadCount": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.onesignal\\.com",
        "website": "https://onesignal.com"
      },
      "OneStat": {
        "cats": [
          10
        ],
        "icon": "OneStat.png",
        "js": {
          "OneStat_Pageview": ""
        },
        "website": "http://www.onestat.com"
      },
      "OneTrust": {
        "cats": [
          67
        ],
        "cookies": {
          "OptanonConsent": ""
        },
        "description": "OneTrust is a cloud-based data privacy management compliance platform.",
        "icon": "OneTrust.png",
        "scripts": [
          "cdn\\.cookielaw\\.org",
          "optanon\\.blob\\.core\\.windows\\.net",
          "otSDKStub\\.js",
          "cdn\\.cookielaw\\.org",
          "optanon\\.blob\\.core\\.windows\\.net"
        ],
        "website": "http://www.onetrust.com"
      },
      "Onshop": {
        "cats": [
          6
        ],
        "description": "OpenCartfree open-source ecommerce platform for online merchants.",
        "excludes": "OpenCart",
        "icon": "Onshop.svg",
        "implies": "PHP",
        "meta": {
          "generator": "Onshop Ecommerce"
        },
        "scripts": "/opencart_custom\\.js",
        "website": "https://onshop.asia"
      },
      "Open AdStream": {
        "cats": [
          36
        ],
        "icon": "Open AdStream.svg",
        "js": {
          "OAS_AD": ""
        },
        "website": "https://www.xaxis.com"
      },
      "Open Classifieds": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:open-classifieds:open_classifieds",
        "icon": "Open Classifieds.png",
        "meta": {
          "author": "open-classifieds\\.com",
          "copyright": "Open Classifieds ?([0-9.]+)?\\;version:\\1"
        },
        "website": "http://open-classifieds.com"
      },
      "Open Journal Systems": {
        "cats": [
          50
        ],
        "cookies": {
          "OJSSID": ""
        },
        "cpe": "cpe:/a:public_knowledge_project:open_journal_systems",
        "description": "Open Journal Systems (OJS) is an open-source software application for managing and publishing scholarly journals.",
        "icon": "Open Journal Systems.png",
        "implies": "PHP",
        "meta": {
          "generator": "Open Journal Systems(?: ([\\d.]+))?\\;version:\\1"
        },
        "website": "http://pkp.sfu.ca/ojs"
      },
      "Open Web Analytics": {
        "cats": [
          10
        ],
        "cpe": "cpe:/a:openwebanalytics:open_web_analytics",
        "html": "<!-- (?:Start|End) Open Web Analytics Tracker -->",
        "icon": "Open Web Analytics.png",
        "js": {
          "OWA.config.baseUrl": "",
          "owa_baseUrl": "",
          "owa_cmds": ""
        },
        "website": "http://www.openwebanalytics.com"
      },
      "Open eShop": {
        "cats": [
          6
        ],
        "icon": "Open eShop.png",
        "implies": "PHP",
        "meta": {
          "author": "open-eshop\\.com",
          "copyright": "Open eShop ?([0-9.]+)?\\;version:\\1"
        },
        "website": "http://open-eshop.com/"
      },
      "OpenBSD httpd": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^OpenBSD httpd"
        },
        "website": "https://man.openbsd.org/httpd.8"
      },
      "OpenCart": {
        "cats": [
          6
        ],
        "cookies": {
          "OCSESSID": ""
        },
        "cpe": "cpe:/a:opencart:opencart",
        "icon": "OpenCart.png",
        "implies": "PHP",
        "website": "http://www.opencart.com"
      },
      "OpenCms": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:alkacon:opencms",
        "headers": {
          "Server": "OpenCms"
        },
        "html": "<link href=\"/opencms/",
        "icon": "OpenCms.png",
        "implies": "Java",
        "scripts": "opencms",
        "website": "http://www.opencms.org"
      },
      "OpenGSE": {
        "cats": [
          22
        ],
        "description": "OpenGSE is a test suite used for testing servlet compliance. It is deployed by using WAR files that are deployed on the server engine.",
        "headers": {
          "Server": "GSE"
        },
        "icon": "Google.svg",
        "implies": "Java",
        "website": "http://code.google.com/p/opengse"
      },
      "OpenGrok": {
        "cats": [
          19
        ],
        "cookies": {
          "OpenGrok": ""
        },
        "icon": "OpenGrok.png",
        "implies": "Java",
        "meta": {
          "generator": "OpenGrok(?: v?([\\d.]+))?\\;version:\\1"
        },
        "website": "http://hub.opensolaris.org/bin/view/Project+opengrok/WebHome"
      },
      "OpenLayers": {
        "cats": [
          35
        ],
        "description": "OpenLayers is an open-source JavaScript library for displaying map data in web browser.",
        "icon": "OpenLayers.png",
        "js": {
          "OpenLayers.VERSION_NUMBER": "([\\d.]+)\\;version:\\1",
          "ol.CanvasMap": ""
        },
        "scripts": "openlayers",
        "website": "https://openlayers.org"
      },
      "OpenNemas": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "OpenNemas"
        },
        "icon": "OpenNemas.png",
        "meta": {
          "generator": "OpenNemas"
        },
        "website": "http://www.opennemas.com"
      },
      "OpenPay": {
        "cats": [
          41
        ],
        "description": "Openpay is an innovative online and in-store payment solution enabling you to purchase now and pay later, with no interest.",
        "icon": "openpay.png",
        "scripts": "openpay\\.com.\\au",
        "website": "https://www.openpay.com.au/"
      },
      "OpenResty": {
        "cats": [
          22
        ],
        "description": "OpenResty is a web platform based on nginx which can run Lua scripts using its LuaJIT engine.",
        "headers": {
          "Server": "openresty(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "OpenResty.png",
        "implies": [
          "Nginx"
        ],
        "website": "http://openresty.org"
      },
      "OpenSSL": {
        "cats": [
          33
        ],
        "cpe": "cpe:/a:openssl:openssl",
        "description": "OpenSSL is a software library for applications that secure communications over computer networks against eavesdropping or need to identify the party at the other end.",
        "headers": {
          "Server": "OpenSSL(?:/([\\d.]+[a-z]?))?\\;version:\\1"
        },
        "icon": "OpenSSL.png",
        "website": "http://openssl.org"
      },
      "OpenStreetMap": {
        "cats": [
          35
        ],
        "description": "OpenStreetMap is a free, editable map of the whole world that is being built by volunteers largely from scratch and released with an open-content license.",
        "dom": "iframe[src*='openstreetmap.org'],iframe[data-lazy-src*='openstreetmap.org']",
        "icon": "OpenStreetMap.svg",
        "oss": true,
        "website": "https://www.openstreetmap.org"
      },
      "OpenText Web Solutions": {
        "cats": [
          1
        ],
        "html": "<!--[^>]+published by Open Text Web Solutions",
        "icon": "OpenText Web Solutions.png",
        "implies": "Microsoft ASP.NET",
        "website": "http://websolutions.opentext.com"
      },
      "OpenUI5": {
        "cats": [
          12
        ],
        "icon": "OpenUI5.png",
        "js": {
          "sap.ui.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "sap-ui-core\\.js",
        "website": "http://openui5.org/"
      },
      "OpenX": {
        "cats": [
          36
        ],
        "cpe": "cpe:/a:openx:openx",
        "description": "OpenX is a programmatic advertising technology company.",
        "dom": "iframe[src*='.openx.net'], img[src*='.openx.net']",
        "icon": "OpenX.png",
        "js": {
          "openx.name": "openx"
        },
        "saas": true,
        "scripts": [
          "https?://[^/]*\\.openx\\.net",
          "https?://[^/]*\\.servedbyopenx\\.com"
        ],
        "website": "http://openx.com"
      },
      "OpinionLab": {
        "cats": [
          5,
          73
        ],
        "description": "OpinionLab is a omnichannel customer feedback solution provider.",
        "icon": "OpinionLab.png",
        "js": {
          "OOo.Browser": "",
          "OOo.version": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.foresee\\.com/code/([\\d.]+)-oo/oo_engine\\.min\\.js\\;version:\\1",
        "website": "https://www.opinionlab.com"
      },
      "Optimise": {
        "cats": [
          71,
          36
        ],
        "description": "Optimise Media Group is a UK-based performance advertising network.",
        "icon": "Optimise.svg",
        "js": {
          "OMID": "^[0-9]+$"
        },
        "scripts": "track\\.omguk\\.com",
        "website": "https://www.optimisemedia.com"
      },
      "Optimizely": {
        "cats": [
          74
        ],
        "icon": "Optimizely.svg",
        "js": {
          "optimizely": ""
        },
        "pricing": [
          "poa",
          "high"
        ],
        "saas": true,
        "scripts": "optimizely\\.com.*\\.js",
        "website": "https://www.optimizely.com"
      },
      "Oracle Application Server": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:oracle:application_server",
        "headers": {
          "Server": "Oracle[- ]Application[- ]Server(?: Containers for J2EE)?(?:[- ](\\d[\\da-z./]+))?\\;version:\\1"
        },
        "icon": "Oracle.png",
        "website": "http://www.oracle.com/technetwork/middleware/ias/overview/index.html"
      },
      "Oracle Commerce": {
        "cats": [
          6
        ],
        "cpe": "cpe:/a:oracle:commerce_platform",
        "headers": {
          "X-ATG-Version": "(?:ATGPlatform/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Oracle.png",
        "website": "http://www.oracle.com/applications/customer-experience/commerce/products/commerce-platform/index.html"
      },
      "Oracle Commerce Cloud": {
        "cats": [
          6
        ],
        "headers": {
          "OracleCommerceCloud-Version": "^(.+)$\\;version:\\1"
        },
        "html": "<[^>]+id=\"oracle-cc\"",
        "icon": "Oracle.png",
        "website": "http://cloud.oracle.com/commerce-cloud"
      },
      "Oracle Dynamic Monitoring Service": {
        "cats": [
          19
        ],
        "headers": {
          "x-oracle-dms-ecid": ""
        },
        "icon": "Oracle.png",
        "website": "http://oracle.com"
      },
      "Oracle HTTP Server": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:oracle:http_server",
        "headers": {
          "Server": "Oracle-HTTP-Server(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Oracle.png",
        "website": "http://oracle.com"
      },
      "Oracle Infinity": {
        "cats": [
          10
        ],
        "description": "Oracle Infinity is a digital analytics platform for tracking, measuring, and optimizing the performance and visitor behavior of enterprise websites and mobile apps.",
        "icon": "Oracle.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "c\\.oracleinfinity\\.io",
        "website": "https://www.oracle.com/cx/marketing/digital-intelligence/"
      },
      "Oracle Maxymiser": {
        "cats": [
          74,
          76
        ],
        "description": "Oracle Maxymiser is a real-time behavioral targeting, in-session personalisation, and product recommendations platform.",
        "icon": "Oracle.png",
        "js": {
          "maxy": "\\;confidence:20",
          "mmsystem.getConfig": "\\;confidence:80"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "service\\.maxymiser\\.net",
        "website": "https://www.oracle.com/uk/cx/marketing/personalization-testing"
      },
      "Oracle Recommendations On Demand": {
        "cats": [
          10
        ],
        "icon": "Oracle.png",
        "scripts": "atgsvcs.+atgsvcs\\.js",
        "website": "http://www.oracle.com/us/products/applications/commerce/recommendations-on-demand/index.html"
      },
      "Oracle Web Cache": {
        "cats": [
          23
        ],
        "cpe": "cpe:/a:oracle:web_cache",
        "description": "Oracle Web Cache is a browser and content management server, which improves the performance of web sites.",
        "headers": {
          "Server": "Oracle(?:AS)?[- ]Web[- ]Cache(?:[- /]([\\da-z./]+))?\\;version:\\1"
        },
        "icon": "Oracle.png",
        "website": "http://oracle.com"
      },
      "Orchard CMS": {
        "cats": [
          1
        ],
        "icon": "Orchard CMS.png",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "Orchard"
        },
        "website": "http://orchardproject.net"
      },
      "Oribi": {
        "cats": [
          10
        ],
        "description": "Oribi is an all-in-one marketing analytics tool.",
        "icon": "Oribi.svg",
        "js": {
          "ORIBI": ""
        },
        "scripts": "cdn\\.oribi\\.io",
        "website": "https://oribi.io/"
      },
      "OroCommerce": {
        "cats": [
          6
        ],
        "html": [
          "<script [^>]+data-requiremodule=\"oro/",
          "<script [^>]+data-requiremodule=\"oroui/"
        ],
        "icon": "orocommerce.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "scripts": [
          "oro\\.min\\.js\\?version=([\\d.]+)\\;version:\\1"
        ],
        "website": "https://oroinc.com"
      },
      "Osano": {
        "cats": [
          67
        ],
        "description": "Osano is a data privacy platform that helps your website become compliant with laws such as GDPR and CCPA.",
        "icon": "Osano.svg",
        "js": {
          "Osano": ""
        },
        "pricing": [
          "freemium",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cookieconsent\\.min\\.js",
        "website": "https://www.osano.com"
      },
      "Outbrain": {
        "cats": [
          5,
          36
        ],
        "description": "Outbrain is a web advertising platform that displays boxes of links, known as chumboxes, to pages within websites.",
        "icon": "Outbrain.svg",
        "js": {
          "OB_ADV_ID": "",
          "OB_releaseVer": "^(.+)$\\;version:\\1",
          "OutbrainPermaLink": "",
          "obApi.version": "^[\\d.]+$"
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "widgets\\.outbrain\\.com/outbrain\\.js",
        "website": "https://www.outbrain.com"
      },
      "Outlook Web App": {
        "cats": [
          30
        ],
        "description": "Outlook on the web is an information manager web app. It includes a web-based email client, a calendar tool, a contact manager, and a task manager.",
        "html": "<link\\s[^>]*href=\"[^\"]*?([\\d.]+)/themes/resources/owafont\\.css\\;version:\\1",
        "icon": "Outlook.svg",
        "implies": "Microsoft ASP.NET",
        "js": {
          "IsOwaPremiumBrowser": ""
        },
        "url": "/owa/auth/log(?:on|off)\\.aspx",
        "website": "http://help.outlook.com"
      },
      "Oxatis": {
        "cats": [
          6
        ],
        "description": "Oxatis is a cloud-based ecommerce solution which enables users to create and manage their own online store websites.",
        "icon": "Oxatis.svg",
        "meta": {
          "generator": "^Oxatis\\s\\(www\\.oxatis\\.com\\)$"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.oxatis.com/"
      },
      "Oxygen": {
        "cats": [
          51
        ],
        "description": "Oxygen Builder is a tool to build a WordPress website.",
        "html": [
          "<body class=(?:\"|')[^\"']*oxygen-body",
          "<link [^>]*href=(?:\"|')[^>]*wp-content/plugins/oxygen/"
        ],
        "icon": "Oxygen.png",
        "implies": "WordPress",
        "scripts": [
          "wp-content/plugins/oxygen"
        ],
        "website": "https://oxygenbuilder.com"
      },
      "PDF.js": {
        "cats": [
          19
        ],
        "html": "<\\/div>\\s*<!-- outerContainer -->\\s*<div\\s*id=\"printContainer\"><\\/div>",
        "icon": "PDF.js.svg",
        "js": {
          "PDFJS": "",
          "PDFJS.version": "^(.+)$\\;version:\\1",
          "_pdfjsCompatibilityChecked": "",
          "pdfjs-dist/build/pdf.version": "^(.+)$\\;version:\\1",
          "pdfjsDistBuildPdf.version": "^(.+)$\\;version:\\1",
          "pdfjsLib.version": "^(.+)$\\;version:\\1"
        },
        "url": "/web/viewer\\.html?file=[^&]\\.pdf",
        "website": "https://mozilla.github.io/pdf.js/"
      },
      "PHP": {
        "cats": [
          27
        ],
        "cookies": {
          "PHPSESSID": ""
        },
        "cpe": "cpe:/a:php:php",
        "description": "PHP is a general-purpose scripting language used for web development.",
        "headers": {
          "Server": "php/?([\\d.]+)?\\;version:\\1",
          "X-Powered-By": "^php/?([\\d.]+)?\\;version:\\1"
        },
        "icon": "PHP.svg",
        "url": "\\.php(?:$|\\?)",
        "website": "http://php.net"
      },
      "PHP-Nuke": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:phpnuke:php-nuke",
        "html": "<[^>]+Powered by PHP-Nuke",
        "icon": "PHP-Nuke.png",
        "implies": "PHP",
        "meta": {
          "generator": "PHP-Nuke"
        },
        "website": "http://phpnuke.org"
      },
      "PHPDebugBar": {
        "cats": [
          47
        ],
        "icon": "phpdebugbar.png",
        "js": {
          "PhpDebugBar": "",
          "phpdebugbar": ""
        },
        "scripts": [
          "debugbar.*\\.js"
        ],
        "website": "http://phpdebugbar.com/"
      },
      "PHPFusion": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:phpfusion:phpfusion",
        "headers": {
          "X-PHPFusion": "(.+)$\\;version:\\1",
          "X-Powered-By": "PHPFusion (.+)$\\;version:\\1"
        },
        "html": [
          "Powered by <a href=\"[^>]+phpfusion",
          "Powered by <a href=\"[^>]+php-fusion"
        ],
        "icon": "PHPFusion.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "website": "https://phpfusion.com"
      },
      "PIXIjs": {
        "cats": [
          25
        ],
        "description": "PixiJS - The HTML5 Creation Engine. 2D WebGL renderer.",
        "icon": "pixijs.png",
        "js": {
          "PIXI": "",
          "PIXI.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "pixi\\.(min\\.)?js$",
        "url": ".+\\.pixijs\\.com",
        "website": "https://www.pixijs.com/"
      },
      "POWR": {
        "cats": [
          5
        ],
        "description": "POWR is a cloud-based system of plugins that work on almost any website.",
        "icon": "POWR.svg",
        "js": {
          "POWR_RECEIVERS": "",
          "loadPowr": ""
        },
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "scripts": "www\\.powr\\.io/powr\\.js",
        "website": "https://www.powr.io"
      },
      "Paddle": {
        "cats": [
          41
        ],
        "description": "Paddle is a billing and payment gateway for B2B SaaS companies.",
        "icon": "Paddle.svg",
        "js": {
          "Paddle.Checkout": "",
          "PaddleScriptLocation": ""
        },
        "pricing": [
          "recurring",
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.paddle\\.com/paddle/paddle\\.js",
        "website": "https://paddle.com/"
      },
      "PageFly": {
        "cats": [
          51
        ],
        "description": "PageFly is an app for Shopify that allows you to build landing pages, product pages, blogs, and FAQs.",
        "icon": "pagefly.png",
        "scripts": "pagefly\\.io",
        "website": "https://pagefly.io"
      },
      "Pagekit": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:pagekit:pagekit",
        "icon": "Pagekit.png",
        "meta": {
          "generator": "Pagekit"
        },
        "website": "http://pagekit.com"
      },
      "Pagely": {
        "cats": [
          62
        ],
        "headers": {
          "Server": "^Pagely"
        },
        "icon": "pagely.svg",
        "implies": [
          "WordPress",
          "Amazon Web Services"
        ],
        "website": "https://pagely.com/"
      },
      "Pagevamp": {
        "cats": [
          1
        ],
        "headers": {
          "X-ServedBy": "pagevamp"
        },
        "icon": "Pagevamp.png",
        "js": {
          "Pagevamp": ""
        },
        "website": "https://www.pagevamp.com"
      },
      "Pantheon": {
        "cats": [
          62
        ],
        "description": "Pantheon is a WebOps (Website Operations) and Management Platform for WordPress and Drupal.",
        "headers": {
          "Server": "^Pantheon",
          "x-pantheon-styx-hostname": "",
          "x-styx-req-id": ""
        },
        "icon": "pantheon.svg",
        "implies": [
          "PHP",
          "Nginx",
          "MariaDB",
          "Fastly"
        ],
        "website": "https://pantheon.io/"
      },
      "Pardot": {
        "cats": [
          32
        ],
        "description": "Pardot is a Software-as-a-Service (SaaS) marketing automation platform.",
        "dns": {
          "TXT": [
            "pardot"
          ]
        },
        "headers": {
          "X-Pardot-LB": "",
          "X-Pardot-Route": "",
          "X-Pardot-Rsp": ""
        },
        "icon": "Pardot.png",
        "js": {
          "piAId": "",
          "piCId": "",
          "piHostname": "",
          "piProtocol": "",
          "piTracker": ""
        },
        "website": "https://www.pardot.com"
      },
      "Pars Elecom Portal": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "Pars Elecom Portal"
        },
        "icon": "parselecom.png",
        "implies": [
          "Microsoft ASP.NET",
          "IIS",
          "Windows Server"
        ],
        "meta": {
          "copyright": "Pars Elecom Portal"
        },
        "website": "http://parselecom.net"
      },
      "Parse.ly": {
        "cats": [
          10
        ],
        "icon": "Parse.ly.png",
        "js": {
          "PARSELY": ""
        },
        "website": "https://www.parse.ly"
      },
      "Paths.js": {
        "cats": [
          25
        ],
        "scripts": "paths(?:\\.min)?\\.js",
        "website": "https://github.com/andreaferretti/paths-js"
      },
      "Patreon": {
        "cats": [
          5,
          41
        ],
        "description": "Patreon is an American membership platform that provides business tools for content creators to run a subscription service.",
        "dom": {
          "a[href*='www.patreon.com/']": {
            "attributes": {
              "href": "patreon\\.com/.+"
            }
          }
        },
        "icon": "Patreon.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "patreon-connect/assets/.+ver=([\\d.]+)\\;version:\\1",
        "website": "https://www.patreon.com"
      },
      "PayBright": {
        "cats": [
          41
        ],
        "description": "PayBright is a Canadian fintech company that offers short-term interest-free installment loans for online shopping to consumers at checkout.",
        "dom": "link[href*='app.paybright.com']",
        "icon": "PayBright.png",
        "js": {
          "_paybright_config": ""
        },
        "scripts": "app\\.paybright\\.com",
        "website": "https://paybright.com"
      },
      "PayKickStart": {
        "cats": [
          5,
          71
        ],
        "description": "PayKickstart is an online shopping cart and affiliate management platform with built-in conversion enhancing features like one-click upsells for credit card/paypal, order bumps, custom checkout pages/widgets/embed forms, coupon management, auto-complete shipping fields, subscription saver sequences, and more.",
        "icon": "PayKickStart.png",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "app\\.paykickstart\\.com",
        "website": "https://paykickstart.com"
      },
      "PayPal": {
        "cats": [
          41
        ],
        "cpe": "cpe:/a:paypal:paypal",
        "description": "PayPal is an online payments system that supports online money transfers and serves as an electronic alternative to traditional paper methods like checks and money orders.",
        "dom": {
          "button": {
            "text": "PayPal"
          },
          "iframe[src*='paypal.com'], img[src*='paypal.com'], img[src*='paypalobjects.com'], link[href*='paypal-express-checkout']": {
            "exists": ""
          }
        },
        "html": "<input[^>]+_s-xclick",
        "icon": "PayPal.svg",
        "js": {
          "PAYPAL": "",
          "enablePaypal": "",
          "paypalClientId": "",
          "paypalJs": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "paypalobjects\\.com",
        "website": "https://paypal.com"
      },
      "Peek": {
        "cats": [
          5,
          72
        ],
        "description": "Peek is a online booking system for tour and activity providers.",
        "icon": "Peek.svg",
        "js": {
          "Peek": "",
          "PeekJsApi": "",
          "_peekConfig": ""
        },
        "scripts": "js\\.peek\\.\\w+",
        "website": "https://www.peek.com/"
      },
      "PeerTube": {
        "cats": [
          2
        ],
        "description": "PeerTube is a free and open-source, decentralized, federated video platform powered by ActivityPub and WebTorrent.",
        "dom": {
          "#incompatible-browser p": {
            "text": "^We are sorry but it seems that PeerTube is not compatible with your web browser\\.$"
          }
        },
        "icon": "PeerTube.svg",
        "meta": {
          "og:platform": "^PeerTube$"
        },
        "website": "https://joinpeertube.org/"
      },
      "Pelican": {
        "cats": [
          57
        ],
        "html": [
          "powered by <a href=\"[^>]+getpelican\\.com",
          "powered by <a href=\"https?://pelican\\.notmyidea\\.org"
        ],
        "icon": "pelican.png",
        "implies": "Python",
        "website": "https://blog.getpelican.com/"
      },
      "PencilBlue": {
        "cats": [
          1,
          11
        ],
        "headers": {
          "X-Powered-By": "PencilBlue"
        },
        "icon": "PencilBlue.png",
        "implies": "Node.js",
        "website": "http://pencilblue.org"
      },
      "Pendo": {
        "cats": [
          58
        ],
        "description": "Pendo is a product experience platform for teams to capture product usage data and behavior.",
        "icon": "Pendo.svg",
        "js": {
          "pendo.VERSION": "(.+)\\;version:\\1"
        },
        "pricing": [
          "freemium",
          "poa"
        ],
        "saas": true,
        "scripts": "\\.pendo\\.io/",
        "website": "https://www.pendo.io"
      },
      "Percona": {
        "cats": [
          34
        ],
        "description": "Percona server is an opensource, fully compatible, enhanced drop-in replacement for MySQL, providing superior performance, scalability, and instrumentation.",
        "icon": "percona.svg",
        "website": "https://www.percona.com"
      },
      "Percussion": {
        "cats": [
          1
        ],
        "html": "<[^>]+class=\"perc-region\"",
        "icon": "Percussion.png",
        "meta": {
          "generator": "(?:Percussion|Rhythmyx)"
        },
        "website": "http://percussion.com"
      },
      "PerimeterX": {
        "cats": [
          16
        ],
        "cookies": {
          "_px3": "",
          "_pxff_cc": "",
          "_pxhd": "",
          "_pxvid": ""
        },
        "description": "PerimeterX is a provider of scalable, behavior-based threat protection technology for the web, cloud, and mobile.",
        "icon": "perimeterx.svg",
        "js": {
          "_pxAppId": ""
        },
        "pricing": [
          "payg",
          "recurring",
          "poa"
        ],
        "saas": true,
        "scripts": "client\\.a\\.pxi\\.pub/",
        "website": "https://www.perimeterx.com"
      },
      "Perl": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:perl:perl",
        "description": "Perl is a family of two high-level, general-purpose, interpreted, dynamic programming languages.",
        "headers": {
          "Server": "\\bPerl\\b(?: ?/?v?([\\d.]+))?\\;version:\\1"
        },
        "icon": "Perl.png",
        "website": "http://perl.org"
      },
      "Permutive": {
        "cats": [
          19
        ],
        "description": "Permutive is a publisher-focused data management platform.",
        "icon": "Permutive.png",
        "js": {
          "permutive": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.permutive\\.com",
        "website": "https://permutive.com",
        "xhr": "api\\.permutive\\.com"
      },
      "PersonaClick": {
        "cats": [
          76
        ],
        "description": "PersonaClick is a provide personalisation, recommandation and multi channel services.",
        "icon": "PersonaClick.png",
        "js": {
          "personaclick": "",
          "personaclick_callback": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.personaclick\\.com/v([\\d.]+)\\.js\\;version:\\1",
        "website": "https://www.personaclick.com"
      },
      "Phabricator": {
        "cats": [
          13,
          47
        ],
        "cookies": {
          "phsid": ""
        },
        "description": "Phabricator is a suite of web-based software development collaboration tools, including the Differential code review tool, the Diffusion repository browser, the Herald change monitoring tool, the Maniphest bug tracker and the Phriction wiki. Phabricator integrates with Git, Mercurial, and Subversion.",
        "html": "<[^>]+(?:class|id)=\"phabricator-",
        "icon": "Phabricator.png",
        "implies": "PHP",
        "scripts": "/phabricator/[a-f0-9]{8}/rsrc/js/phui/[a-z-]+\\.js$",
        "website": "http://phacility.com"
      },
      "Phaser": {
        "cats": [
          12
        ],
        "icon": "Phaser.png",
        "js": {
          "Phaser": "",
          "Phaser.VERSION": "^(.+)$\\;version:\\1"
        },
        "website": "https://phaser.io"
      },
      "Phenomic": {
        "cats": [
          57
        ],
        "description": "Phenomic is a modular website compiler.",
        "html": [
          "<[^>]+id=\"phenomic(?:root)?\""
        ],
        "icon": "Phenomic.svg",
        "implies": "React",
        "scripts": "/phenomic\\.browser\\.[a-f0-9]+\\.js",
        "website": "https://phenomic.io/"
      },
      "Phoenix": {
        "cats": [
          18
        ],
        "icon": "sazito-phoenix.png",
        "implies": [
          "React",
          "webpack",
          "Node.js"
        ],
        "js": {
          "Phoenix": ""
        },
        "meta": {
          "generator": "^phoenix"
        },
        "website": "https://github.com/Sazito/phoenix/"
      },
      "PhotoShelter": {
        "cats": [
          1
        ],
        "html": [
          "<!--\\s+#+ Powered by the PhotoShelter Beam platform",
          "<link[^>]+c\\.photoshelter\\.com"
        ],
        "icon": "PhotoShelter.png",
        "implies": [
          "PHP",
          "MySQL",
          "jQuery"
        ],
        "url": "photoshelter\\.com",
        "website": "https://www.photoshelter.com"
      },
      "Phusion Passenger": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:phusionpassenger:phusion_passenger",
        "description": "Phusion Passenger is a free web server and application server with support for Ruby, Python and Node.js.",
        "headers": {
          "Server": "Phusion Passenger ([\\d.]+)\\;version:\\1",
          "X-Powered-By": "Phusion Passenger ?([\\d.]+)?\\;version:\\1"
        },
        "icon": "Phusion Passenger.png",
        "website": "https://phusionpassenger.com"
      },
      "Piano": {
        "cats": [
          76
        ],
        "description": "Piano is a enterprise SaaS company which specializing in advanced media business processes and ecommerce optimisation.",
        "icon": "Piano.png",
        "js": {
          "PianoESPConfig": "",
          "gciDataPiano": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "\\.tinypass\\.com",
          "\\.piano\\.io"
        ],
        "website": "https://piano.io"
      },
      "Pico": {
        "cats": [
          53
        ],
        "icon": "pico.svg",
        "js": {
          "Pico": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "widget\\.pico\\.tools"
        ],
        "website": "https://trypico.com"
      },
      "Picreel": {
        "cats": [
          77,
          5
        ],
        "description": "Picreel is a conversion optimisation software.",
        "dom": "iframe[src*='app.picreel.com']",
        "icon": "Picreel.svg",
        "js": {
          "picreel": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.picreel\\.com",
        "website": "https://www.picreel.com"
      },
      "Pimcore": {
        "cats": [
          1,
          6
        ],
        "cpe": "cpe:/a:pimcore:pimcore",
        "dom": ".pimcore_area_content",
        "headers": {
          "X-Powered-By": "^pimcore$"
        },
        "icon": "pimcore.svg",
        "implies": "PHP",
        "website": "http://pimcore.org"
      },
      "Pin Payments": {
        "cats": [
          41
        ],
        "description": "Pin Payments is an all-in-one online payment system. It offers merchants a simple JSON API, secure credit card storage, multi-currency capabilities, bank account compatibility, onsite payment processing and automatic fund transfer to specified bank accounts.",
        "icon": "pinpayments.png",
        "scripts": "api\\.pinpayments\\.com",
        "website": "https://www.pinpayments.com/"
      },
      "Pingdom": {
        "cats": [
          78
        ],
        "description": "Pingdom is a Swedish website monitoring software as a service company.",
        "icon": "Pingdom.svg",
        "pricing": [
          "payg",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.pingdom\\.net",
        "website": "https://www.pingdom.com"
      },
      "Pingoteam": {
        "cats": [
          1
        ],
        "icon": "Pingoteam.svg",
        "implies": "PHP",
        "meta": {
          "designer": "Pingoteam"
        },
        "website": "https://www.pingoteam.ir/"
      },
      "Pinterest": {
        "cats": [
          5
        ],
        "description": "Pinterest is an image sharing and social media service designed to enable saving and discovery of information.",
        "icon": "Pinterest.svg",
        "scripts": "//assets\\.pinterest\\.com/js/pinit\\.js",
        "website": "http://pinterest.com"
      },
      "Pinterest Ads": {
        "cats": [
          36
        ],
        "description": "Pinterest Ads is an online advertising platform developed by Pinterest.",
        "icon": "Pinterest.svg",
        "website": "https://ads.pinterest.com/",
        "xhr": "ct\\.pinterest\\.com"
      },
      "Pinterest Conversion Tag": {
        "cats": [
          10
        ],
        "description": "Pinterest Conversion Tag allows you to track actions people take on your website after viewing your Promoted Pin.",
        "dom": "img[src*='ct.pinterest.com/v3/?tid']",
        "icon": "Pinterest.svg",
        "js": {
          "pintrk": ""
        },
        "website": "https://www.pinterest.com.au/business/"
      },
      "Pipedrive": {
        "cats": [
          52,
          53
        ],
        "description": "Pipedrive is a cloud-based sales CRM.",
        "icon": "Pipedrive.svg",
        "js": {
          "LeadBooster": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "https://www.pipedrive.com/"
      },
      "PixelFed": {
        "cats": [
          2
        ],
        "description": "PixelFed is an activitypub based image sharing platform.",
        "dom": {
          "a[href='pixelfed.org'][title*='version']": {
            "attributes": {
              "title": "^version ([0-9.]+)$\\;version:\\1"
            },
            "text": "^Powered by Pixelfed$"
          }
        },
        "icon": "PixelFed.svg",
        "implies": [
          "Laravel"
        ],
        "website": "https://pixelfed.org"
      },
      "Pixlee": {
        "cats": [
          32
        ],
        "description": "Pixlee is a visual marketing platform that helps brands leverage user-generated content to improve their marketing performance.",
        "icon": "Pixlee.png",
        "js": {
          "Pixlee": "",
          "Pixlee_Analytics": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "assets\\.pixlee\\.com",
        "website": "https://pixlee.com"
      },
      "Planet": {
        "cats": [
          49
        ],
        "description": "Planet is a feed aggregator, which creates pages with entries from the original feeds in chronological order, most recent entries first.",
        "icon": "Planet.png",
        "meta": {
          "generator": "^Planet(?:/([\\d.]+))?\\;version:\\1"
        },
        "website": "http://planetplanet.org"
      },
      "Platform.sh": {
        "cats": [
          62
        ],
        "headers": {
          "x-platform-cluster": "",
          "x-platform-processor": "",
          "x-platform-router": "",
          "x-platform-server": ""
        },
        "icon": "platformsh.svg",
        "website": "https://platform.sh"
      },
      "PlatformOS": {
        "cats": [
          1,
          62
        ],
        "headers": {
          "x-powered-by": "^platformOS$"
        },
        "icon": "PlatformOS.svg",
        "website": "https://www.platform-os.com"
      },
      "Plausible": {
        "cats": [
          10
        ],
        "cookies": {},
        "description": "Plausible is an open-source alternative to Google Analytics.",
        "icon": "Plausible.svg",
        "js": {
          "plausible": ""
        },
        "scripts": "plausible\\.io/js/plausible\\.js",
        "website": "https://plausible.io/"
      },
      "Play": {
        "cats": [
          18
        ],
        "cookies": {
          "PLAY_SESSION": ""
        },
        "cpe": "cpe:/a:playframework:play_framework",
        "icon": "Play.svg",
        "implies": "Scala",
        "website": "https://www.playframework.com"
      },
      "Pleroma": {
        "cats": [
          2
        ],
        "description": "Pleroma is a free, federated social networking server built on open protocols.",
        "dom": {
          "noscript": {
            "text": "^To use Pleroma, please enable JavaScript.$"
          },
          "title": {
            "text": "^Pleroma$"
          }
        },
        "icon": "Pleroma.svg",
        "website": "https://pleroma.social/"
      },
      "Plesk": {
        "cats": [
          9
        ],
        "description": "Plesk is a web hosting and server data centre automation software with a control panel developed for Linux and Windows-based retail hosting service providers.",
        "headers": {
          "X-Powered-By": "^Plesk(?:L|W)in",
          "X-Powered-By-Plesk": "^Plesk"
        },
        "icon": "Plesk.png",
        "scripts": "common\\.js\\?plesk",
        "website": "https://www.plesk.com/"
      },
      "Pligg": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:pligg:pligg_cms",
        "html": "<span[^>]+id=\"xvotes-0",
        "icon": "Pligg.svg",
        "js": {
          "pligg_": ""
        },
        "meta": {
          "generator": "Pligg"
        },
        "website": "http://pligg.com"
      },
      "Plone": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:plone:plone",
        "icon": "Plone.svg",
        "implies": "Python",
        "meta": {
          "generator": "Plone"
        },
        "website": "http://plone.org"
      },
      "Plotly": {
        "cats": [
          25
        ],
        "icon": "Plotly.png",
        "implies": "D3",
        "js": {
          "Plotly.version": "([\\d.])\\;version:\\1"
        },
        "scripts": "https?://cdn\\.plot\\.ly/plotly",
        "website": "https://plot.ly/javascript/"
      },
      "Plyr": {
        "cats": [
          14
        ],
        "description": "Rlyr is a set of tools for splitting, applying and combining Data. ",
        "icon": "Plyr.png",
        "js": {
          "Plyr": ""
        },
        "scripts": "https://cdn\\.plyr\\.io/([0-9.]+)/.+\\.js\\;version:\\1",
        "website": "https://plyr.io/"
      },
      "Po.st": {
        "cats": [
          5
        ],
        "icon": "Po.st.png",
        "js": {
          "pwidget_config": ""
        },
        "website": "http://www.po.st/"
      },
      "Podia": {
        "cats": [
          21,
          6
        ],
        "cookies": {
          "_podia_storefront_visitor_id": ""
        },
        "description": "Podia is a platform to host and sell online courses, memberships, and digital downloads.",
        "icon": "Podia.svg",
        "js": {
          "Podia.Checkout": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.podia\\.com",
        "website": "https://www.podia.com"
      },
      "Podium": {
        "cats": [
          5,
          52
        ],
        "description": "Podium is a customer communication platform for businesses who interact with customers on a local level.",
        "icon": "Podium.png",
        "js": {
          "PodiumWebChat": "",
          "podiumWebsiteWidgetLoaded": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.podium\\.com/",
        "website": "https://www.podium.com"
      },
      "Polyfill": {
        "cats": [
          59
        ],
        "icon": "polyfill.svg",
        "scripts": [
          "^https?://cdn\\.polyfill\\.io/",
          "/polyfill\\.min\\.js"
        ],
        "website": "https://polyfill.io"
      },
      "Polymer": {
        "cats": [
          12
        ],
        "html": "(?:<polymer-[^>]+|<link[^>]+rel=\"import\"[^>]+/polymer\\.html\")",
        "icon": "Polymer.png",
        "js": {
          "Polymer.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "polymer\\.js",
        "website": "http://polymer-project.org"
      },
      "Popper": {
        "cats": [
          19
        ],
        "description": "Popper is an tolltip & popover positioning engine",
        "html": [
          "<script [^>]*src=\"[^\"]*/popper\\.js/([0-9.]+)\\;version:\\1"
        ],
        "icon": "Popper.png",
        "js": {
          "createPopper": ""
        },
        "scripts": [
          "/popper\\.js/([0-9.]+)\\;version:\\1"
        ],
        "website": "https://popper.js.org"
      },
      "Post Affiliate Pro": {
        "cats": [
          71
        ],
        "description": "Post Affiliate Pro is a software built for online stores and ecommerce websites that need to track and monitor their affiliate network.",
        "icon": "Post Affiliate Pro.svg",
        "js": {
          "PostAffAction": "",
          "PostAffCookie": "",
          "PostAffInfo": "",
          "PostAffTracker": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "postaffiliatepro\\.com/scripts/trackjs\\.js",
          "(?:affiliate|associate)\\..+/scripts/trackjs\\.js"
        ],
        "website": "https://www.postaffiliatepro.com"
      },
      "Posterous": {
        "cats": [
          1,
          11
        ],
        "html": "<div class=\"posterous",
        "icon": "Posterous.png",
        "js": {
          "Posterous": ""
        },
        "website": "http://posterous.com"
      },
      "PostgreSQL": {
        "cats": [
          34
        ],
        "cpe": "cpe:/a:postgresql:postgresql",
        "description": "PostgreSQL, also known as Postgres, is a free and open-source relational database management system emphasizing extensibility and SQL compliance.",
        "icon": "PostgreSQL.png",
        "website": "http://www.postgresql.org/"
      },
      "Powerboutique": {
        "cats": [
          6
        ],
        "icon": "powerboutique.png",
        "scripts": "powerboutique",
        "website": "https://www.powerboutique.com/"
      },
      "Powergap": {
        "cats": [
          6
        ],
        "html": [
          "<a[^>]+title=\"POWERGAP",
          "<input type=\"hidden\" name=\"shopid\""
        ],
        "icon": "Powergap.png",
        "saas": true,
        "website": "http://powergap.de"
      },
      "Preact": {
        "cats": [
          59
        ],
        "description": "Preact is a JavaScript library that describes itself as a fast 3kB alternative to React with the same ES6 API.",
        "dom": "meta[data-preact-helmet='true']",
        "icon": "Preact.svg",
        "oss": true,
        "website": "https://preactjs.com"
      },
      "Prebid": {
        "cats": [
          36
        ],
        "description": "Prebid is an open-source header bidding wrapper. It forms the core of our Nucleus ad platform, helping maximize revenue and performance for publishers.",
        "icon": "Prebid.png",
        "js": {
          "PREBID_TIMEOUT": "",
          "pbjs": ""
        },
        "oss": true,
        "pricing": [
          "freemium"
        ],
        "scripts": [
          "/prebid\\.js",
          "adnxs\\.com/[^\"]*(?:prebid|/pb\\.js)"
        ],
        "website": "http://prebid.org"
      },
      "Prediggo": {
        "cats": [
          76,
          32
        ],
        "description": "Prediggo is an ecommerce personalisation and marketing automation software provider.",
        "icon": "Prediggo.png",
        "js": {
          "Prediggo": "",
          "PrediggoSearchFormExternalAc": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "js/prediggo/(?:[\\w]+)\\.js",
        "website": "https://prediggo.com"
      },
      "Prefix-Free": {
        "cats": [
          19
        ],
        "icon": "Prefix-Free.png",
        "js": {
          "PrefixFree": ""
        },
        "scripts": "prefixfree\\.js",
        "website": "https://leaverou.github.io/prefixfree/"
      },
      "PrestaShop": {
        "cats": [
          6
        ],
        "cookies": {
          "PrestaShop": ""
        },
        "cpe": "cpe:/a:prestashop:prestashop",
        "description": "PrestaShop is a freemium, open-source ecommerce solution, written in the PHP programming language with support for the MySQL database management system.",
        "headers": {
          "Powered-By": "^Prestashop$"
        },
        "html": [
          "Powered by <a\\s+[^>]+>PrestaShop",
          "<!-- /Block [a-z ]+ module (?:HEADER|TOP)?\\s?-->",
          "<!-- /Module Block [a-z ]+ -->"
        ],
        "icon": "PrestaShop.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "freeProductTranslation": "\\;confidence:25",
          "prestashop": "",
          "priceDisplayMethod": "\\;confidence:25",
          "priceDisplayPrecision": "\\;confidence:25",
          "rcAnalyticsEvents.eventPrestashopCheckout": ""
        },
        "meta": {
          "generator": "PrestaShop"
        },
        "oss": true,
        "pricing": [
          "freemium"
        ],
        "website": "http://www.prestashop.com"
      },
      "Primis": {
        "cats": [
          36
        ],
        "description": "Primis is a video discovery platform for publishers.",
        "dom": {
          "iframe[src*='.sekindo.com']": {
            "attributes": {
              "src": ""
            }
          },
          "img[src*='.sekindo.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "Primis.svg",
        "js": {
          "SekindoNativeSkinApi": "",
          "sekindoDisplayedPlacement": "",
          "sekindoFlowingPlayerOn": ""
        },
        "saas": true,
        "scripts": "\\.sekindo\\.com",
        "website": "https://www.primis.tech",
        "xhr": "\\.sekindo\\.com"
      },
      "Prism": {
        "cats": [
          19
        ],
        "description": "Prism is an extensible syntax highlighter.",
        "icon": "Prism.svg",
        "js": {
          "Prism": ""
        },
        "scripts": "prism\\.js",
        "website": "http://prismjs.com"
      },
      "Prismic": {
        "cats": [
          1
        ],
        "description": "Prismic is a headless CMS for Jamstack.",
        "icon": "Prismic.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "website": "https://prismic.io",
        "xhr": "\\.cdn\\.prismic\\.io"
      },
      "Profitwell": {
        "cats": [
          10
        ],
        "icon": "Profitwell.svg",
        "scripts": [
          "public\\.profitwell\\.com/js/profitwell\\.js"
        ],
        "website": "https://www.profitwell.com/"
      },
      "Project Wonderful": {
        "cats": [
          36
        ],
        "html": "<div[^>]+id=\"pw_adbox_",
        "icon": "Project Wonderful.png",
        "js": {
          "pw_adloader": ""
        },
        "scripts": "^https?://(?:www\\.)?projectwonderful\\.com/(?:pwa\\.js|gen\\.php)",
        "website": "http://projectwonderful.com"
      },
      "Projesoft": {
        "cats": [
          6
        ],
        "icon": "projesoft.png",
        "scripts": [
          "projesoft\\.js"
        ],
        "website": "https://www.projesoft.com.tr"
      },
      "Prototype": {
        "cats": [
          12
        ],
        "cpe": "cpe:/a:prototypejs:prototype",
        "description": "Prototype is a JavaScript Framework that aims to ease development of web applications.",
        "icon": "Prototype.png",
        "js": {
          "Prototype.Version": "^(.+)$\\;version:\\1"
        },
        "scripts": "(?:prototype|protoaculous)(?:-([\\d.]*[\\d]))?.*\\.js\\;version:\\1",
        "website": "http://www.prototypejs.org"
      },
      "Protovis": {
        "cats": [
          25
        ],
        "js": {
          "protovis": ""
        },
        "scripts": "protovis.*\\.js",
        "website": "http://mbostock.github.io/protovis"
      },
      "ProvenExpert": {
        "cats": [
          5
        ],
        "description": "ProvenExpert is a review based marketing platform that allows users to create customer surveys, provides aggregate reviews and ratings.",
        "dom": {
          "img[src*='provenexpert']": {
            "attributes": {
              "src": "images\\.provenexpert\\.\\w+"
            }
          }
        },
        "icon": "ProvenExpert.svg",
        "scripts": "provenexpert\\.\\w+/widget",
        "website": "https://www.provenexpert.com"
      },
      "Provide Support": {
        "cats": [
          52
        ],
        "description": "Provide Support is a SaaS for customer service that includes live chat, real-time website monitoring, chat statistics.",
        "icon": "Provide Support.svg",
        "pricing": [
          "low",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": "\\.providesupport\\.com",
        "website": "https://www.providesupport.com"
      },
      "Proximis": {
        "cats": [
          5,
          6
        ],
        "icon": "Proximis Omnichannel.png",
        "scripts": "widget-commerce(?:\\.min)?\\.js",
        "website": "https://www.proximis.com"
      },
      "Proximis Unified Commerce": {
        "cats": [
          6,
          1
        ],
        "html": "<html[^>]+data-ng-app=\"RbsChangeApp\"",
        "icon": "Proximis Omnichannel.png",
        "implies": [
          "PHP",
          "AngularJS"
        ],
        "js": {
          "__change": ""
        },
        "meta": {
          "generator": "Proximis Unified Commerce"
        },
        "website": "https://www.proximis.com"
      },
      "PubMatic": {
        "cats": [
          36
        ],
        "description": "PubMatic is a company that develops and implements online advertising software and strategies for the digital publishing and advertising industry.",
        "dom": {
          "iframe[src*='.pubmatic.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "PubMatic.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "https?://[^/]*\\.pubmatic\\.com",
        "website": "http://www.pubmatic.com/",
        "xhr": "\\.pubmatic\\.com"
      },
      "Public CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "PUBLICCMS_USER": ""
        },
        "headers": {
          "X-Powered-PublicCMS": "^(.+)$\\;version:\\1"
        },
        "icon": "Public CMS.png",
        "implies": "Java",
        "website": "http://www.publiccms.com"
      },
      "Pulse Secure": {
        "cats": [
          46
        ],
        "cookies": {
          "DSSIGNIN": ""
        },
        "description": "Pulse Secure allows to deploy VPNs to securely to your internal resources.",
        "icon": "PulseSecure.png",
        "url": "/dana-na/auth/",
        "website": "https://www.pulsesecure.net/products/remote-access-overview/"
      },
      "Pure CSS": {
        "cats": [
          66
        ],
        "description": "Pure CSS is a set of small, responsive CSS modules that can be used in web projects.",
        "html": [
          "<link[^>]+(?:([\\d.])+/)?pure(?:-min)?\\.css\\;version:\\1",
          "<div[^>]+class=\"[^\"]*pure-u-(?:sm-|md-|lg-|xl-)?\\d-\\d"
        ],
        "icon": "Pure CSS.png",
        "website": "http://purecss.io"
      },
      "Pure Chat": {
        "cats": [
          52
        ],
        "description": "Pure Chat is a live chat solution for small to mid-sized teams.",
        "icon": "Pure Chat.png",
        "js": {
          "PCWidget": "",
          "purechatApi": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "app\\.purechat\\.com",
        "website": "https://www.purechat.com"
      },
      "PushEngage": {
        "cats": [
          32
        ],
        "description": "PushEngage is a browser push notification platform that helps content website managers engage visitors by automatically segmenting and sending web push messages.",
        "icon": "PushEngage.png",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "clientcdn\\.pushengage\\.\\w+/core",
        "website": "https://www.pushengage.com"
      },
      "PushOwl": {
        "cats": [
          32
        ],
        "description": " PushOwl is a push notification platform for ecommerce stores.",
        "icon": "PushOwl.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.pushowl\\.com",
        "website": "https://pushowl.com"
      },
      "Pushnami": {
        "cats": [
          32
        ],
        "description": "Pushnami is an AI-powered messaging platform that uses intelligent analytics to deliver superior push, social, and email performance.",
        "icon": "Pushnami.svg",
        "scripts": "api\\.pushnami\\.com",
        "website": "https://pushnami.com"
      },
      "Pygments": {
        "cats": [
          19
        ],
        "cpe": "cpe:/a:pygments:pygments",
        "html": "<link[^>]+pygments\\.css[\"']",
        "icon": "pygments.png",
        "website": "http://pygments.org"
      },
      "PyroCMS": {
        "cats": [
          1
        ],
        "cookies": {
          "pyrocms": ""
        },
        "headers": {
          "X-Streams-Distribution": "PyroCMS"
        },
        "icon": "PyroCMS.png",
        "implies": "Laravel",
        "website": "http://pyrocms.com"
      },
      "Python": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:python:python",
        "description": "Python is an interpreted and general-purpose programming language.",
        "headers": {
          "Server": "(?:^|\\s)Python(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Python.png",
        "website": "http://python.org"
      },
      "Quadpay": {
        "cats": [
          41
        ],
        "description": "Quadpay is a payment platform.",
        "icon": "Quadpay.svg",
        "js": {
          "QuadPayShopify": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.quadpay\\.com/quadpay-widget-([\\d.]+)\\.js\\;version:\\1",
        "website": "https://www.quadpay.com"
      },
      "Quantcast Choice": {
        "cats": [
          67
        ],
        "description": "Quantcast Choice is a free consent management platform to meet key privacy requirements stemming from ePrivacy Directive, GDPR, and CCPA.",
        "icon": "Quantcast.png",
        "js": {
          "quantserve": ""
        },
        "scripts": "quantcast\\.mgr\\.consensu\\.org",
        "website": "https://www.quantcast.com/gdpr/consent-management-solution/"
      },
      "Quantcast Measure": {
        "cats": [
          10
        ],
        "icon": "Quantcast.png",
        "js": {
          "quantserve": ""
        },
        "scripts": "\\.quantserve\\.com/quant\\.js",
        "website": "https://www.quantcast.com/products/measure-audience-insights/"
      },
      "Qubit": {
        "cats": [
          74,
          76
        ],
        "description": "Qubit is a SaaS based persuasive personalisation at scale services.",
        "icon": "Qubit.png",
        "js": {
          "__qubit": "",
          "onQubitReady": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "static\\.goqubit\\.com",
        "website": "https://www.qubit.com"
      },
      "Question2Answer": {
        "cats": [
          15
        ],
        "description": "Question2Answer (Q2A) is a popular open-source Q&A platform for PHP/MySQL.",
        "html": "<!-- Powered by Question2Answer",
        "icon": "question2answer.png",
        "implies": "PHP",
        "scripts": "\\./qa-content/qa-page\\.js\\?([0-9.]+)\\;version:\\1",
        "website": "http://www.question2answer.org"
      },
      "Quick.CMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:opensolution:quick.cms",
        "html": "<a href=\"[^>]+opensolution\\.org/\">CMS by",
        "icon": "Quick.CMS.png",
        "meta": {
          "generator": "Quick\\.CMS(?: v([\\d.]+))?\\;version:\\1"
        },
        "website": "http://opensolution.org"
      },
      "Quick.Cart": {
        "cats": [
          6
        ],
        "html": "<a href=\"[^>]+opensolution\\.org/\">(?:Shopping cart by|Sklep internetowy)",
        "icon": "Quick.Cart.png",
        "meta": {
          "generator": "Quick\\.Cart(?: v([\\d.]+))?\\;version:\\1"
        },
        "website": "http://opensolution.org"
      },
      "Quill": {
        "cats": [
          24
        ],
        "description": "Quill is a free open-source WYSIWYG editor.",
        "icon": "Quill.png",
        "js": {
          "Quill": ""
        },
        "website": "http://quilljs.com"
      },
      "RBS Change": {
        "cats": [
          1,
          6
        ],
        "html": "<html[^>]+xmlns:change=",
        "icon": "RBS Change.png",
        "implies": "PHP",
        "meta": {
          "generator": "RBS Change"
        },
        "website": "http://www.rbschange.fr"
      },
      "RCMS": {
        "cats": [
          1
        ],
        "icon": "RCMS.png",
        "meta": {
          "generator": "^(?:RCMS|ReallyCMS)"
        },
        "website": "http://www.rcms.fi"
      },
      "RD Station": {
        "cats": [
          32
        ],
        "description": "RD Station is a platform that helps medium and small businesses manage and automate their Digital Marketing strategy.",
        "icon": "RD Station.png",
        "js": {
          "RDStation": ""
        },
        "scripts": "d335luupugsy2\\.cloudfront\\.net/js/loader-scripts/.*-loader\\.js",
        "website": "http://rdstation.com.br"
      },
      "RDoc": {
        "cats": [
          4
        ],
        "cpe": "cpe:/a:dave_thomas:rdoc",
        "description": "RDoc produces HTML and command-line documentation for Ruby projects.",
        "html": [
          "<link[^>]+href=\"[^\"]*rdoc-style\\.css",
          "Generated by <a[^>]+href=\"https?://rdoc\\.rubyforge\\.org[^>]+>RDoc</a> ([\\d.]*\\d)\\;version:\\1",
          "Generated by <a href=\"https:\\/\\/ruby\\.github\\.io\\/rdoc\\/\">RDoc<\\/a> ([\\d.]*\\d)\\;version:\\1"
        ],
        "icon": "RDoc.png",
        "implies": "Ruby",
        "js": {
          "rdoc_rel_prefix": ""
        },
        "website": "https://github.com/ruby/rdoc"
      },
      "RTB House": {
        "cats": [
          77
        ],
        "description": "RTB House is a company that provides learning-powered retargeting solutions for brands and agencies.",
        "icon": "RTB House.png",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://www.rtbhouse.com",
        "xhr": "\\.creativecdn\\.com"
      },
      "RX Web Server": {
        "cats": [
          22
        ],
        "headers": {
          "X-Powered-By": "RX-WEB"
        },
        "icon": "RXWeb.svg",
        "website": "http://developers.rokitax.co.uk/projects/rxweb"
      },
      "RackCache": {
        "cats": [
          23
        ],
        "description": "RackCache is a quick drop-in component to enable HTTP caching for Rack-based applications.",
        "headers": {
          "X-Rack-Cache": ""
        },
        "icon": "RackCache.png",
        "implies": "Ruby",
        "website": "https://github.com/rtomayko/rack-cache"
      },
      "RainLoop": {
        "cats": [
          30
        ],
        "description": "RainLoop is a web-based email client.",
        "headers": {
          "Server": "^RainLoop"
        },
        "html": [
          "<link[^>]href=\"rainloop/v/([0-9.]+)/static/apple-touch-icon\\.png/>\\;version:\\1"
        ],
        "icon": "RainLoop.png",
        "implies": "PHP",
        "js": {
          "rainloop": "",
          "rainloopI18N": ""
        },
        "meta": {
          "rlAppVersion": "^([0-9.]+)$\\;version:\\1"
        },
        "scripts": "^rainloop/v/([0-9.]+)/\\;version:\\1",
        "website": "https://www.rainloop.net/"
      },
      "Rakuten": {
        "cats": [
          71
        ],
        "cookies": {
          "rakuten-source": ""
        },
        "description": "Rakuten (formerly Ebates) allows you to earn cash-back rewards.",
        "icon": "Rakuten.svg",
        "js": {
          "rakutenRanMID": "",
          "rakutenSource": ""
        },
        "scripts": "tag\\.rmp\\.rakuten\\.com",
        "website": "https://www.rakuten.com/"
      },
      "Rakuten Digital Commerce": {
        "cats": [
          6
        ],
        "icon": "RakutenDigitalCommerce.png",
        "js": {
          "RakutenApplication": ""
        },
        "website": "https://digitalcommerce.rakuten.com.br"
      },
      "Ramda": {
        "cats": [
          59
        ],
        "icon": "Ramda.png",
        "scripts": "ramda.*\\.js",
        "website": "http://ramdajs.com"
      },
      "Raphael": {
        "cats": [
          25
        ],
        "description": "Raphael is a cross-browser JavaScript library that draws Vector graphics for websites.",
        "icon": "Raphael.png",
        "js": {
          "Raphael.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "raphael(?:-([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://dmitrybaranovskiy.github.io/raphael/"
      },
      "RapidSpike": {
        "cats": [
          13,
          16
        ],
        "description": "RapidSpike is an uptime and performance monitoring service for web sites and applications.",
        "icon": "RapidSpike.svg",
        "pricing": [
          "poa",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.rapidspike\\.com/",
        "website": "https://www.rapidspike.com"
      },
      "Raptor": {
        "cats": [
          76
        ],
        "description": "Raptor is a personalisation engine based on machine learning that analyses and learns about the user's behavior and unique browser history.",
        "icon": "Raptor.png",
        "js": {
          "Raptor": "",
          "onRaptorLoaded": "",
          "raptorBase64": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "api\\.raptorsmartadvisor\\.com",
          "msecnd\\.net/script/raptor-([\\d.]+)\\.js\\;version:\\1"
        ],
        "website": "https://raptorsmartadvisor.com"
      },
      "Raspbian": {
        "cats": [
          28
        ],
        "description": "Raspbian is a free operating system for the Raspberry Pi hardware.",
        "headers": {
          "Server": "Raspbian",
          "X-Powered-By": "Raspbian"
        },
        "icon": "Raspbian.svg",
        "website": "https://www.raspbian.org/"
      },
      "Raychat": {
        "cats": [
          52
        ],
        "description": "Raychat is a free customer messaging platform.",
        "icon": "raychat.png",
        "js": {
          "Raychat": ""
        },
        "scripts": "app\\.raychat\\.io/scripts/js",
        "website": "https://raychat.io"
      },
      "Raygun": {
        "cats": [
          78,
          13
        ],
        "description": "Raygun is a cloud-based networking monitoring and bug tracking application.",
        "icon": "Raygun.svg",
        "js": {
          "Raygun": "",
          "raygunEnabled": "",
          "raygunFactory": ""
        },
        "pricing": [
          "payg",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.raygun\\.io",
        "website": "https://raygun.com"
      },
      "Rayo": {
        "cats": [
          1
        ],
        "icon": "Rayo.png",
        "implies": [
          "AngularJS",
          "Microsoft ASP.NET"
        ],
        "js": {
          "Rayo": ""
        },
        "meta": {
          "generator": "^Rayo"
        },
        "website": "http://www.rayo.ir"
      },
      "ReDoc": {
        "cats": [
          4
        ],
        "description": "Redoc is an open-source tool that generates API documentation from OpenAPI specifications.",
        "html": "<redoc ",
        "icon": "redoc.png",
        "implies": "React",
        "js": {
          "Redoc.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "/redoc\\.(?:min\\.)?js",
        "website": "https://github.com/Rebilly/ReDoc"
      },
      "React": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:facebook:react",
        "description": "React is an open-source JavaScript library for building user interfaces or UI components.",
        "dom": {
          "body > div": {
            "properties": {
              "_reactRootContainer": ""
            }
          }
        },
        "html": "<[^>]+data-react",
        "icon": "React.png",
        "js": {
          "React.version": "^(.+)$\\;version:\\1",
          "react.version": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "react(?:-with-addons)?[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "/([\\d.]+)/react(?:\\.min)?\\.js\\;version:\\1",
          "^react\\b.*\\.js"
        ],
        "website": "https://reactjs.org"
      },
      "React Redux": {
        "cats": [
          12
        ],
        "description": "Official React bindings for Redux",
        "icon": "Redux.png",
        "implies": [
          "React",
          "Redux"
        ],
        "scripts": [
          "/react-redux(@|/)([\\d.]+)(?:/[a-z]+)?/react-redux(?:.min)?\\.js\\;version:\\2"
        ],
        "website": "https://react-redux.js.org/"
      },
      "React Router": {
        "cats": [
          12
        ],
        "description": "React Router provides declarative routing for React.",
        "icon": "React Router.png",
        "implies": "React",
        "scripts": [
          "/react-router(@|/)([\\d.]+)(?:/[a-z]+)?/react-router(?:.min)?\\.js\\;version:\\2"
        ],
        "website": "https://reactrouter.com/"
      },
      "RebelMouse": {
        "cats": [
          1
        ],
        "headers": {
          "x-rebelmouse-cache-control": "",
          "x-rebelmouse-surrogate-control": ""
        },
        "html": "<!-- Powered by RebelMouse\\.",
        "icon": "RebelMouse.svg",
        "website": "https://www.rebelmouse.com/"
      },
      "Recart": {
        "cats": [
          32
        ],
        "description": "Recart is a tool to engage users who abandoned their shopping cart via Facebook Messenger.",
        "icon": "Recart.svg",
        "js": {
          "__recart": "",
          "recart": ""
        },
        "scripts": "api\\.recart\\.com",
        "website": "https://recart.com/"
      },
      "Recite Me": {
        "cats": [
          68
        ],
        "description": "Recite Me is a cloud-based web accessibility assistive toolbar that allows website visitors to customize a site in a way that works for them.",
        "icon": "Recite Me.png",
        "scripts": "api\\.reciteme\\.com/asset/js",
        "website": "https://reciteme.com/"
      },
      "Recurly": {
        "cats": [
          41
        ],
        "description": "Recurly provides enterprise-class subscription billing and recurring payment management for thousands of businesses worldwide.",
        "html": "<input[^>]+data-recurly",
        "icon": "Recurly.png",
        "js": {
          "recurly.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "js\\.recurly\\.com",
        "website": "https://recurly.com"
      },
      "Red Hat": {
        "cats": [
          28
        ],
        "cpe": "cpe:/o:redhat:linux",
        "description": "Red Hat is an open-source Linux operating system.",
        "headers": {
          "Server": "Red Hat",
          "X-Powered-By": "Red Hat"
        },
        "icon": "Red Hat.svg",
        "website": "https://www.redhat.com"
      },
      "RedCart": {
        "cats": [
          6
        ],
        "cookies": {
          "rc2c-erotica": "\\d+"
        },
        "description": "RedCart is an all-in-one ecommerce platform from Poland.",
        "icon": "RedCart.svg",
        "js": {
          "RC_SHOP_ID": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://redcart.pl"
      },
      "Reddit": {
        "cats": [
          2
        ],
        "html": "(?:<a[^>]+Powered by Reddit|powered by <a[^>]+>reddit<)",
        "icon": "Reddit.png",
        "implies": "Python",
        "js": {
          "reddit": ""
        },
        "url": "^https?://(?:www\\.)?reddit\\.com",
        "website": "http://code.reddit.com"
      },
      "Redis": {
        "cats": [
          34
        ],
        "cpe": "cpe:/a:redislabs:redis",
        "description": "Redis is an in-memory data structure project implementing a distributed, in-memory key–value database with optional durability. Redis supports different kinds of abstract data structures, such as strings, lists, maps, sets, sorted sets, HyperLogLogs, bitmaps, streams, and spatial indexes.",
        "icon": "Redis.svg",
        "website": "https://redis.io"
      },
      "Redis Object Cache": {
        "cats": [
          23
        ],
        "html": "<!--\\s+Performance optimized by Redis Object Cache",
        "icon": "RedisObjectCache.svg",
        "implies": [
          "Redis",
          "WordPress"
        ],
        "website": "https://wprediscache.com"
      },
      "Redmine": {
        "cats": [
          13
        ],
        "cookies": {
          "_redmine_session": ""
        },
        "cpe": "cpe:/a:redmine:redmine",
        "description": "Redmine is a free and open-source, web-based project management and issue tracking tool.",
        "html": "Powered by <a href=\"[^>]+Redmine",
        "icon": "Redmine.png",
        "implies": "Ruby on Rails",
        "meta": {
          "description": "Redmine"
        },
        "website": "http://www.redmine.org"
      },
      "Redux": {
        "cats": [
          12
        ],
        "description": "Redux is a predictable state container for JavaScript applications.",
        "icon": "Redux.png",
        "scripts": [
          "/redux(@|/)([\\d.]+)(?:/[a-z]+)?/redux(?:.min)?\\.js\\;version:\\2"
        ],
        "website": "https://redux.js.org/"
      },
      "Refersion": {
        "cats": [
          71,
          36
        ],
        "description": "Refersion is an affiliate management app.",
        "dom": {
          "a[href*='.refersion.com']": {
            "attributes": {
              "href": ""
            }
          },
          "img[src*='cdn.refersion.com'],img[src*='s3.amazonaws.com/refersion_client/']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "Refersion.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.refersion\\.com",
        "website": "http://refersion.com"
      },
      "Regiondo": {
        "cats": [
          5,
          72
        ],
        "description": "Regiondo is a online booking system for tour and activity providers.",
        "icon": "Regiondo.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.regiondo\\.net",
        "website": "https://www.regiondo.com"
      },
      "Reinvigorate": {
        "cats": [
          10
        ],
        "icon": "Reinvigorate.png",
        "js": {
          "reinvigorate": ""
        },
        "website": "http://www.reinvigorate.net"
      },
      "RequireJS": {
        "cats": [
          12
        ],
        "description": "RequireJS is a JavaScript library and file loader which manages the dependencies between JavaScript files and in modular programming.",
        "icon": "RequireJS.png",
        "js": {
          "requirejs.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "require.*\\.js",
        "website": "http://requirejs.org"
      },
      "ResDiary": {
        "cats": [
          5,
          72
        ],
        "description": "ResDiary, is a online reservation system for hospitality operators.",
        "dom": {
          "iframe[src*='resdiary']": {
            "attributes": {
              "src": "\\.resdiary\\.\\w+/"
            }
          }
        },
        "icon": "ResDiary.svg",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.resdiary\\.\\w+/",
        "website": "https://www.resdiary.com"
      },
      "Resengo": {
        "cats": [
          5,
          72
        ],
        "description": "Resengo is a restaurant table booking widget.",
        "dom": {
          "iframe[src*='resengo']": {
            "attributes": {
              "src": "www\\.resengo\\.\\w+"
            }
          }
        },
        "icon": "Resengo.svg",
        "js": {
          "wpJsonpResengoReservationWidget": ""
        },
        "scripts": "www\\.resengo\\.\\w+",
        "website": "https://wwc.resengo.com"
      },
      "Reservio": {
        "cats": [
          72
        ],
        "description": "Reservio is a cloud-based appointment scheduling and online booking solution.",
        "dom": [
          ".reservio-booking-button",
          "a[href*='.reservio.com'][target='_blank']",
          "a[href*='bookings.reservio.com']"
        ],
        "icon": "Reservio.svg",
        "pricing": [
          "freemium",
          "recurring",
          "low"
        ],
        "saas": true,
        "scripts": "static\\.reservio\\.com",
        "website": "https://www.reservio.com"
      },
      "Resin": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:caucho:resin",
        "headers": {
          "Server": "^Resin(?:/(\\S*))?\\;version:\\1"
        },
        "icon": "Resin.png",
        "implies": "Java",
        "website": "http://caucho.com"
      },
      "Resmio": {
        "cats": [
          5,
          72
        ],
        "description": "Resmio is a restaurant table booking widget.",
        "icon": "Resmio.svg",
        "js": {
          "ResmioButton": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": "static\\.resmio\\.\\w+/static/",
        "website": "https://www.resmio.com"
      },
      "Resy": {
        "cats": [
          5,
          72
        ],
        "description": "Resy is an technology and media company that provides an app and back-end management software for restaurant reservations.",
        "icon": "Resy.svg",
        "js": {
          "resyWidget": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "widgets\\.resy\\.\\w+",
        "website": "https://resy.com"
      },
      "Retail Rocket": {
        "cats": [
          76
        ],
        "cookies": {
          "rr-testCookie": "testvalue",
          "rrpvid": "^\\d+$"
        },
        "description": "Retail Rocket is a big data-based personalisation platform for ecommerce websites.",
        "icon": "Retail Rocket.png",
        "js": {
          "retailrocket": "",
          "rrAddToBasket": "\\;confidence:25",
          "rrApiOnReady": "\\;confidence:25",
          "rrLibrary": "\\;confidence:25",
          "rrPartnerId": "\\;confidence:25"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.retailrocket\\.net",
        "website": "https://retailrocket.net"
      },
      "RevLifter": {
        "cats": [
          76
        ],
        "cookies": {
          "REVLIFTER": ""
        },
        "description": "RevLifter is an AI-powered coupon technology which allows brands to offer personalised incentives to their customers based on real-time basket data.",
        "icon": "RevLifter.png",
        "js": {
          "RevLifterObject": "",
          "revlifter": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "assets\\.revlifter\\.io",
        "website": "https://www.revlifter.com"
      },
      "Reveal.js": {
        "cats": [
          12
        ],
        "icon": "Reveal.js.png",
        "implies": "Highlight.js",
        "js": {
          "Reveal.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": "(?:^|/)reveal(?:\\.min)?\\.js",
        "website": "http://lab.hakim.se/reveal-js"
      },
      "Revel": {
        "cats": [
          18
        ],
        "cookies": {
          "REVEL_FLASH": "",
          "REVEL_SESSION": ""
        },
        "icon": "Revel.png",
        "implies": "Go",
        "website": "https://revel.github.io"
      },
      "RevolverMaps": {
        "cats": [
          35
        ],
        "description": "RevolverMaps is a collection of real-time visitor statistics widgets for website or blog. Interactive visitor mappings to a globe rendered by the Revolver Engine.",
        "icon": "RevolverMaps.svg",
        "scripts": "\\.revolvermaps\\.com",
        "website": "https://www.revolvermaps.com"
      },
      "Revslider": {
        "cats": [
          19
        ],
        "description": "Slider Revolution is a WordPress plugin that allows you to create responsive sliders with many animation effects, text, image and video layers, and many other features.",
        "html": [
          "<link[^>]* href=[\\'\"][^']+revslider[/\\w-]+\\.css\\?ver=([0-9.]+)[\\'\"]\\;version:\\1"
        ],
        "icon": "revslider.png",
        "scripts": "/revslider/[/\\w-]+/js",
        "website": "https://revolution.themepunch.com/"
      },
      "Rewardful": {
        "cats": [
          71
        ],
        "description": "Rewardful is a way for SaaS companies to setup affiliate and referral programs with Stripe.",
        "icon": "Rewardful.svg",
        "js": {
          "Rewardful": ""
        },
        "scripts": "r\\.wdfl\\.co",
        "website": "https://www.getrewardful.com/"
      },
      "Rezdy": {
        "cats": [
          5,
          72
        ],
        "description": "Rezdy is an online booking software for tours and attractions.",
        "icon": "Rezdy.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "rezdy\\.\\w+/pluginJs",
        "website": "https://www.rezdy.com"
      },
      "Rezgo": {
        "cats": [
          5,
          72
        ],
        "description": "Rezgo is a tour operator software that provides online booking system.",
        "dom": {
          "iframe": {
            "attributes": {
              "id": "rezgo_content_frame"
            }
          },
          "link": {
            "attributes": {
              "href": "wp-content/plugins/rezgo/rezgo/templates"
            }
          }
        },
        "icon": "Rezgo.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://www.rezgo.com"
      },
      "RichRelevance": {
        "cats": [
          76
        ],
        "description": "RichRelevance is a cloud-based omnichannel personalisation platform built to help Retailers, B2B, financial services, travel and hospitality, and branded manufacturers personalise their customer experiences.",
        "icon": "RichRelevance.png",
        "js": {
          "RR.U": "\\;confidence:50",
          "rr_v": "([\\d.]+)\\;version:\\1\\;confidence:50"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.richrelevance\\.com/",
        "website": "https://richrelevance.com"
      },
      "Rickshaw": {
        "cats": [
          25
        ],
        "implies": "D3",
        "js": {
          "Rickshaw": ""
        },
        "scripts": "rickshaw(?:\\.min)?\\.js",
        "website": "http://code.shutterstock.com/rickshaw/"
      },
      "RightJS": {
        "cats": [
          12
        ],
        "icon": "RightJS.png",
        "js": {
          "RightJS": ""
        },
        "scripts": "right\\.js",
        "website": "http://rightjs.org"
      },
      "Riot": {
        "cats": [
          12
        ],
        "icon": "Riot.png",
        "js": {
          "riot": ""
        },
        "scripts": "riot(?:\\+compiler)?(?:\\.min)?\\.js",
        "website": "https://riot.js.org/"
      },
      "Riskified": {
        "cats": [
          6
        ],
        "headers": {
          "server": "Riskified Server"
        },
        "html": [
          "<[^>]*beacon\\.riskified\\.com",
          "<[^>]*c\\.riskified\\.com"
        ],
        "icon": "riskified.png",
        "js": {
          "RISKX": "",
          "riskifiedBeaconLoad": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.riskified.com/"
      },
      "RiteCMS": {
        "cats": [
          1
        ],
        "icon": "RiteCMS.png",
        "implies": [
          "PHP",
          "SQLite\\;confidence:80"
        ],
        "meta": {
          "generator": "^RiteCMS(?: (.+))?\\;version:\\1"
        },
        "website": "http://ritecms.com"
      },
      "Roadiz CMS": {
        "cats": [
          1,
          11
        ],
        "headers": {
          "X-Powered-By": "Roadiz CMS"
        },
        "icon": "Roadiz CMS.png",
        "implies": [
          "PHP",
          "Symfony"
        ],
        "meta": {
          "generator": "^Roadiz ?(?:master|develop)? v?([0-9\\.]+)\\;version:\\1"
        },
        "website": "https://www.roadiz.io"
      },
      "Robin": {
        "cats": [
          6
        ],
        "icon": "Robin.png",
        "js": {
          "_robin_getRobinJs": "",
          "robin_settings": "",
          "robin_storage_settings": ""
        },
        "website": "http://www.robinhq.com"
      },
      "RockRMS": {
        "cats": [
          1,
          11,
          32
        ],
        "description": "Rock RMS is a free, open-source Relationship Management System (RMS) built for churches and businesses.",
        "icon": "RockRMS.svg",
        "implies": [
          "Windows Server",
          "IIS",
          "Microsoft ASP.NET"
        ],
        "meta": {
          "generator": "^Rock v([0-9.]+)\\;version:\\1"
        },
        "website": "http://www.rockrms.com"
      },
      "Rollbar": {
        "cats": [
          13
        ],
        "icon": "Rollbar.svg",
        "scripts": [
          "rollbar\\.js/([0-9.]+)\\;version:\\1"
        ],
        "website": "https://rollbar.com/"
      },
      "RoundCube": {
        "cats": [
          30
        ],
        "description": "RoundCube is free and open-source web-based IMAP email client.",
        "html": "<title>RoundCube",
        "icon": "RoundCube.png",
        "implies": "PHP",
        "js": {
          "rcmail": "",
          "roundcube": ""
        },
        "website": "http://roundcube.net"
      },
      "Rubicon Project": {
        "cats": [
          36
        ],
        "description": "Rubicon Project is an advertising automation platform enabling publishers to transact advertising brands.",
        "dom": {
          "iframe[src*='.rubiconproject.com']": {
            "attributes": {
              "src": ""
            }
          },
          "img[src*='.rubiconproject.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "Rubicon Project.svg",
        "saas": true,
        "scripts": "https?://[^/]*\\.rubiconproject\\.com",
        "website": "http://rubiconproject.com/",
        "xhr": "\\.rubiconproject\\.com"
      },
      "Ruby": {
        "cats": [
          27
        ],
        "cpe": "cpe:/a:ruby-lang:ruby",
        "description": "Ruby is an open-source object-oriented programming language.",
        "headers": {
          "Server": "(?:Mongrel|WEBrick|Ruby)"
        },
        "icon": "Ruby.png",
        "website": "http://ruby-lang.org"
      },
      "Ruby Receptionists": {
        "cats": [
          5
        ],
        "description": "Ruby Receptionists is a Portland, Oregon based virtual answering service for small businesses.",
        "icon": "Ruby Receptionists.svg",
        "js": {
          "rubyApi": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "chatwidget\\.ruby\\.com",
        "website": "https://www.ruby.com"
      },
      "Ruby on Rails": {
        "cats": [
          18
        ],
        "cookies": {
          "_session_id": "\\;confidence:75"
        },
        "cpe": "cpe:/a:rubyonrails:rails",
        "description": "Ruby on Rails is a server-side web application framework written in Ruby under the MIT License.",
        "headers": {
          "Server": "mod_(?:rails|rack)",
          "X-Powered-By": "mod_(?:rails|rack)"
        },
        "icon": "Ruby on Rails.png",
        "implies": "Ruby",
        "meta": {
          "csrf-param": "^authenticity_token$\\;confidence:50"
        },
        "scripts": "/assets/application-[a-z\\d]{32}/\\.js\\;confidence:50",
        "website": "https://rubyonrails.org"
      },
      "RxJS": {
        "cats": [
          12
        ],
        "icon": "RxJS.png",
        "js": {
          "Rx.CompositeDisposable": "",
          "Rx.Symbol": ""
        },
        "scripts": "rx(?:\\.\\w+)?(?:\\.compat|\\.global)?(?:\\.min)?\\.js",
        "website": "http://reactivex.io"
      },
      "SAP": {
        "cats": [
          53
        ],
        "headers": {
          "Server": "SAP NetWeaver Application Server"
        },
        "icon": "SAP.svg",
        "website": "http://sap.com"
      },
      "SAP Commerce Cloud": {
        "cats": [
          6
        ],
        "cookies": {
          "_hybris": ""
        },
        "cpe": "cpe:/a:sap:commerce_cloud",
        "description": "SAP Commerce Cloud is a cloud-native omnichannel commerce solution for B2B, B2C, and B2B2C companies.",
        "html": [
          "<[^>]+/(?:sys_master|hybr|_ui/(?:.*responsive/)?(?:desktop|common(?:/images|/img|/css|ico)?))/",
          "<script[^>].*hybris.*.js"
        ],
        "icon": "SAP.svg",
        "implies": "Java",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.sap.com/products/crm/e-commerce-platforms.html"
      },
      "SAP Customer Data Cloud Sign-in": {
        "cats": [
          69
        ],
        "icon": "SAP.svg",
        "scripts": "\\.gigya\\.com/JS/gigya\\.js",
        "website": "https://www.sap.com/uk/acquired-brands/what-is-gigya.html"
      },
      "SDL Tridion": {
        "cats": [
          1
        ],
        "html": "<img[^>]+_tcm\\d{2,3}-\\d{6}\\.",
        "icon": "SDL Tridion.svg",
        "website": "http://www.sdl.com/products/tridion"
      },
      "SEMrush": {
        "cats": [
          32
        ],
        "description": "SEMrush is an all-in-one tool suite for improving online visibility and discovering marketing insights.",
        "icon": "SEMrush.svg",
        "js": {
          "SEMRUSH": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "www\\.semrush\\.com",
        "website": "https://www.semrush.com"
      },
      "SIMsite": {
        "cats": [
          1
        ],
        "icon": "SIMsite.png",
        "meta": {
          "SIM.medium": ""
        },
        "scripts": "/sim(?:site|core)/js",
        "website": "http://simgroep.nl/internet/portfolio-contentbeheer_41623/"
      },
      "SMF": {
        "cats": [
          2
        ],
        "description": "SMF is a free open-source software, used for community forums and is written in PHP.",
        "html": "credits/?\" title=\"Simple Machines Forum\" target=\"_blank\" class=\"new_win\">SMF ([0-9.]+)</a>\\;version:\\1",
        "icon": "SMF.png",
        "implies": "PHP",
        "js": {
          "smf_": ""
        },
        "website": "http://www.simplemachines.org"
      },
      "SOBI 2": {
        "cats": [
          19
        ],
        "html": "(?:<!-- start of Sigsiu Online Business Index|<div[^>]* class=\"sobi2)",
        "icon": "SOBI 2.png",
        "implies": "Joomla",
        "website": "http://www.sigsiu.net/sobi2.html"
      },
      "SPDY": {
        "cats": [
          19
        ],
        "excludes": "HTTP/2",
        "headers": {
          "X-Firefox-Spdy": "\\d\\.\\d"
        },
        "icon": "SPDY.png",
        "website": "http://chromium.org/spdy"
      },
      "SPIP": {
        "cats": [
          1
        ],
        "headers": {
          "Composed-By": "SPIP ([\\d.]+) @\\;version:\\1",
          "X-Spip-Cache": ""
        },
        "icon": "spip.svg",
        "implies": "PHP",
        "meta": {
          "generator": "(?:^|\\s)SPIP(?:\\s([\\d.]+(?:\\s\\[\\d+\\])?))?\\;version:\\1"
        },
        "website": "http://www.spip.net"
      },
      "SPNEGO": {
        "cats": [
          16
        ],
        "description": "SPNEGO is an authentication method commonly used in Windows servers to allow NTLM or Kerberos authentication",
        "headers": {
          "WWW-Authenticate": "^Negotiate"
        },
        "website": "https://tools.ietf.org/html/rfc4559"
      },
      "SQL Buddy": {
        "cats": [
          3
        ],
        "description": "SQL Buddy is an open-source web-based application written in PHP to handle the administration of MySQL and SQLite with the use of a Web browser.",
        "html": "(?:<title>SQL Buddy</title>|<[^>]+onclick=\"sideMainClick\\(\"home\\.php)",
        "icon": "SQL Buddy.png",
        "implies": "PHP",
        "website": "http://www.sqlbuddy.com"
      },
      "SQLite": {
        "cats": [
          34
        ],
        "icon": "SQLite.png",
        "website": "http://www.sqlite.org"
      },
      "STUDIO": {
        "cats": [
          51
        ],
        "description": "STUDIO is a Japan-based company and SaaS application for designing and hosting websites. The service includes a visual editor with built-in CMS and analytics.",
        "dom": ".StudioCanvas, .publish-studio-style",
        "icon": "STUDIO.svg",
        "implies": [
          "Vue.js",
          "Nuxt.js",
          "Firebase",
          "Google Cloud",
          "Google Tag Manager"
        ],
        "meta": {
          "generator": "^STUDIO$"
        },
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "website": "https://studio.design"
      },
      "SUSE": {
        "cats": [
          28
        ],
        "description": "SUSE is a Linux-based server operating system.",
        "headers": {
          "Server": "SUSE(?:/?\\s?-?([\\d.]+))?\\;version:\\1",
          "X-Powered-By": "SUSE(?:/?\\s?-?([\\d.]+))?\\;version:\\1"
        },
        "icon": "SUSE.png",
        "website": "http://suse.com"
      },
      "SWFObject": {
        "cats": [
          19
        ],
        "description": "SWFObject is an open-source JavaScript library used to embed Adobe Flash content onto web pages.",
        "icon": "SWFObject.png",
        "js": {
          "SWFObject": ""
        },
        "oss": true,
        "scripts": "swfobject.*\\.js",
        "website": "https://github.com/swfobject/swfobject"
      },
      "Saber": {
        "cats": [
          57
        ],
        "description": "Saber is a framework for building static websites.",
        "html": [
          "<div [^>]*id=\"_saber\""
        ],
        "icon": "Saber.svg",
        "implies": "Vue.js",
        "meta": {
          "generator": "^Saber v([\\d.]+)$\\;version:\\1"
        },
        "website": "https://saber.land/"
      },
      "Sails.js": {
        "cats": [
          18
        ],
        "cookies": {
          "sails.sid": ""
        },
        "headers": {
          "X-Powered-By": "^Sails(?:$|[^a-z0-9])"
        },
        "icon": "Sails.js.svg",
        "implies": "Express",
        "website": "http://sailsjs.org"
      },
      "Sailthru": {
        "cats": [
          32,
          75,
          76
        ],
        "cookies": {
          "sailthru_pageviews": ""
        },
        "description": "Sailthru is a marketing automation software and multi-channel personalisation tool that serves ecommerce and media brands.",
        "dom": {
          "link[href*='ak.sail-horizon.com'],link[href*='api.sail-personalize.com'],link[href*='api.sail-track.com']": {
            "attributes": {
              "href": ""
            }
          }
        },
        "icon": "Sailthru.svg",
        "js": {
          "Sailthru": "",
          "sailthruIdentify": "",
          "sailthruNewsletterRegistration": "",
          "trackSailthruUser": ""
        },
        "meta": {
          "sailthru.image.full": "",
          "sailthru.title": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "ak\\.sail-horizon\\.com",
        "website": "https://www.sailthru.com"
      },
      "SaleCycle": {
        "cats": [
          6,
          76
        ],
        "description": "SaleCycle is a UK based global behavioral marketing firm.",
        "dom": "iframe[src*='.salecycle.com'][target='_self']",
        "icon": "salecycle.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.salecycle.com"
      },
      "Salesfloor": {
        "cats": [
          6
        ],
        "icon": "salesfloor.png",
        "scripts": [
          "salesFloor\\.js"
        ],
        "website": "https://salesfloor.net/"
      },
      "Salesforce": {
        "cats": [
          53
        ],
        "cookies": {
          "com.salesforce": ""
        },
        "description": "Salesforce is a cloud computing service software (SaaS) that specializes in customer relationship management (CRM).",
        "dns": {
          "TXT": [
            "salesforce\\.com"
          ]
        },
        "html": "<[^>]+=\"brandQuaternaryFgrs\"",
        "icon": "Salesforce.svg",
        "js": {
          "SFDCApp": "",
          "SFDCCmp": "",
          "SFDCPage": "",
          "SFDCSessionVars": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "https://www.salesforce.com"
      },
      "Salesforce Commerce Cloud": {
        "cats": [
          6
        ],
        "description": "Salesforce Commerce Cloud is a cloud-based software-as-a-service (SaaS) ecommerce solution.",
        "headers": {
          "Server": "Demandware eCommerce Server"
        },
        "html": [
          "<[^>]+demandware\\.edgesuite"
        ],
        "icon": "Salesforce.svg",
        "implies": "Salesforce",
        "js": {
          "dwAnalytics": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "/demandware\\.static/",
        "website": "http://demandware.com"
      },
      "Salesforce Interaction Studio": {
        "cats": [
          76
        ],
        "description": "Salesforce Interaction Studio (formerly Evergage) is a cloud-based software that allows users to collect, analyze, and respond to user behavior on their websites and web applications in real-time.",
        "icon": "Salesforce.svg",
        "js": {
          "Evergage": "",
          "evergageHideSections": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.evgnet\\.com",
        "website": "https://www.salesforce.com/products/marketing-cloud/customer-interaction"
      },
      "Salesforce Service Cloud": {
        "cats": [
          52,
          53
        ],
        "description": "Salesforce Service Cloud is a customer relationship management (CRM) platform for customer service and support.",
        "icon": "Salesforce.svg",
        "implies": "Salesforce",
        "js": {
          "embedded_svc": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "scripts": "service\\.force\\.com",
        "website": "https://www.salesforce.com/au/products/service-cloud/"
      },
      "Sana Commerce": {
        "cats": [
          6
        ],
        "description": "Sana Commerce is an ecommerce platform for SAP and Microsoft Dynamics.",
        "icon": "Sana Commerce.svg",
        "js": {
          "Sana.UI": ""
        },
        "pricing": [
          "recurring",
          "onetime"
        ],
        "saas": true,
        "website": "https://www.sana-commerce.com"
      },
      "Sanity": {
        "cats": [
          1
        ],
        "description": "Sanity is a platform for structured content. It comes with an open-source, headless CMS that can be customized with Javascript, a real-time hosted data store and an asset delivery pipeline.",
        "dom": {
          "img[src*='cdn.sanity.io'],img[srcset*='cdn.sanity.io'],video[src*='cdn.sanity.io'],source[src*='cdn.sanity.io'],source[srcset*='cdn.sanity.io'],track[src*='cdn.sanity.io']": {
            "attributes": {
              "src": "",
              "srcset": ""
            }
          }
        },
        "headers": {
          "x-sanity-shard": ""
        },
        "icon": "Sanity.svg",
        "pricing": [
          "freemium",
          "recurring",
          "payg"
        ],
        "saas": true,
        "website": "https://www.sanity.io",
        "xhr": "api(?:cdn)?\\.sanity\\.io"
      },
      "Sapper": {
        "cats": [
          18
        ],
        "html": [
          "<script[^>]*>__SAPPER__"
        ],
        "icon": "Sapper.svg",
        "implies": [
          "Svelte",
          "Node.js"
        ],
        "js": {
          "__SAPPER__": ""
        },
        "website": "https://sapper.svelte.dev"
      },
      "Sarka-SPIP": {
        "cats": [
          1
        ],
        "icon": "Sarka-SPIP.png",
        "implies": "SPIP",
        "meta": {
          "generator": "Sarka-SPIP(?:\\s([\\d.]+))?\\;version:\\1"
        },
        "website": "http://sarka-spip.net"
      },
      "Sazito": {
        "cats": [
          6
        ],
        "icon": "Sazito.svg",
        "js": {
          "Sazito": ""
        },
        "meta": {
          "generator": "^Sazito"
        },
        "website": "http://sazito.com"
      },
      "Scala": {
        "cats": [
          27
        ],
        "description": "Scala is a general-purpose programming language providing support for both object-oriented programming and functional programming.",
        "icon": "Scala.png",
        "website": "http://www.scala-lang.org"
      },
      "Schedule Engine": {
        "cats": [
          52
        ],
        "description": "Schedule Engine is a customer support solution built for contractors.",
        "icon": "Schedule Engine.svg",
        "scripts": "webchat.scheduleengine.net",
        "website": "https://www.scheduleengine.com/"
      },
      "Scientific Linux": {
        "cats": [
          28
        ],
        "description": "Scientific Linux (SL) is a free open-source operating system based on Red Hat Enterprise Linux.",
        "headers": {
          "Server": "Scientific Linux",
          "X-Powered-By": "Scientific Linux"
        },
        "icon": "Scientific Linux.png",
        "website": "http://scientificlinux.org"
      },
      "Scorpion": {
        "cats": [
          1
        ],
        "description": "Scorpion is a marketing and technology provider.",
        "html": "<[^>]+id=\"HSScorpion",
        "icon": "Scorpion.svg",
        "js": {
          "Process.UserData": ""
        },
        "scripts": "cdn.cxc.scorpion.direct",
        "website": "https://www.scorpion.co/"
      },
      "SeamlessCMS": {
        "cats": [
          1
        ],
        "icon": "SeamlessCMS.png",
        "meta": {
          "generator": "^Seamless\\.?CMS"
        },
        "website": "http://www.seamlesscms.com"
      },
      "Searchspring": {
        "cats": [
          6,
          29
        ],
        "description": "Searchspring is a highly advanced search solution for ecommerce platforms.",
        "icon": "Searchspring.svg",
        "js": {
          "SearchSpring": "",
          "SearchSpringConf": "",
          "SearchSpringInit": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.searchspring\\.net",
        "website": "https://searchspring.com"
      },
      "Sectigo": {
        "cats": [
          70
        ],
        "certIssuer": "Sectigo",
        "description": "Sectigo provides SSL certificate and computer security products.",
        "icon": "Sectigo.svg",
        "website": "https://sectigo.com/"
      },
      "Section.io": {
        "cats": [
          31
        ],
        "description": "Section.io is a Content Delivery Network (CDN).",
        "headers": {
          "section-io-id": "",
          "section-io-origin-status": "",
          "section-io-origin-time-seconds": ""
        },
        "icon": "sectionio.svg",
        "website": "https://www.section.io"
      },
      "Seers": {
        "cats": [
          67
        ],
        "icon": "seersco.png",
        "scripts": [
          "seersco.com/script/cb\\.js"
        ],
        "website": "http://www.seersco.com"
      },
      "Segmanta": {
        "cats": [
          73
        ],
        "description": "Segmanta is a mobile-first survey platform designed for product feedback, brand awareness and concept testing research.",
        "icon": "Segmanta.png",
        "js": {
          "SEGMANTA__DYNAMIC_EMBED_CONFIG": "",
          "SEGMANTA__USER_METADATA": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "pge\\.segmanta\\.com/widget_embed_js(?:/widgetEmbed-v([\\d.]+)\\.min\\.js)?\\;version:\\1",
        "website": "https://segmanta.com"
      },
      "Segment": {
        "cats": [
          10
        ],
        "description": "Segment is a customer data platform (CDP) that helps you collect, clean, and control your customer data.",
        "dns": {
          "TXT": [
            "segment-site-verification"
          ]
        },
        "icon": "Segment.svg",
        "js": {
          "analytics.VERSION": "(.+)\\;version:\\1"
        },
        "pricing": [
          "mid",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.segment\\.com/analytics\\.js",
          "/segment-wrapper\\.min\\.js"
        ],
        "website": "https://segment.com"
      },
      "SegmentStream": {
        "cats": [
          10
        ],
        "description": "SegmentStream is a AI-powered marketing analytics platform built for data-driven CMOs, web analysts and performance marketing teams.",
        "icon": "SegmentStream.svg",
        "js": {
          "segmentstream.VERSION": "(.+)\\;version:\\1"
        },
        "pricing": [
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.segmentstream\\.com",
        "website": "https://segmentstream.com"
      },
      "Select2": {
        "cats": [
          59
        ],
        "description": "Select2 is a jQuery based replacement for select boxes. It supports searching, remote data sets, and infinite scrolling of results.",
        "icon": "Select2.png",
        "implies": "jQuery",
        "js": {
          "jQuery.fn.select2": ""
        },
        "scripts": "select2(?:\\.min|\\.full)?\\.js",
        "website": "https://select2.org/"
      },
      "Sellacious": {
        "cats": [
          6
        ],
        "description": "Sellacious is an open-source ecommerce and marketplace platform for integrated POS and online stores.",
        "dom": {
          ".mod-sellacious-cart": {
            "text": ""
          }
        },
        "icon": "Sellacious.png",
        "implies": "Joomla",
        "js": {
          "SellaciousViewCartAIO": ""
        },
        "oss": true,
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "website": "https://www.sellacious.com"
      },
      "Sellingo": {
        "cats": [
          6
        ],
        "description": "Sellingo is a Polish ecommerce platform.",
        "dom": {
          "a[href*='sellingo.pl'][target='_blank']": {
            "text": "Oprogramowanie sklepu internetowego Sellingo.pl"
          }
        },
        "icon": "Sellingo.svg",
        "js": {
          "sellingoQuantityCalc": ""
        },
        "pricing": [
          "recurring",
          "low",
          "freemium"
        ],
        "saas": true,
        "website": "https://sellingo.pl"
      },
      "Sellix": {
        "cats": [
          6
        ],
        "description": "Sellix is an ecommerce payment processor. It accepts PayPal, PerfectMoney and popular cryptocurrencies.",
        "icon": "Sellix.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.sellix\\.io/static/js/embed\\.js",
        "website": "https://sellix.io/"
      },
      "Selly": {
        "cats": [
          6
        ],
        "description": "Selly is an ecommerce payment processor. It accepts PayPal, Bitcoin, Ethereum, Litecoin, Stripe and more.",
        "icon": "Selly.svg",
        "pricing": [
          "low",
          "poa",
          "recurring"
        ],
        "saas": true,
        "scripts": "embed\\.selly\\.(?:gg|io)",
        "website": "https://selly.io/"
      },
      "Semantic UI": {
        "cats": [
          66
        ],
        "description": "Semantic UI is a front-end development framework, powered by LESS and jQuery.",
        "html": [
          "<link[^>]+semantic(?:\\.min)\\.css\""
        ],
        "icon": "Semantic-ui.png",
        "scripts": "/semantic(?:-([\\d.]+))?(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://semantic-ui.com/"
      },
      "Sencha Touch": {
        "cats": [
          12,
          26
        ],
        "description": "Sencha Touch is a user interface (UI) JavaScript library, or web framework, specifically built for the Mobile Web.",
        "icon": "Sencha Touch.png",
        "scripts": "sencha-touch.*\\.js",
        "website": "http://www.sencha.com/products/touch"
      },
      "Sendgrid": {
        "cats": [
          75
        ],
        "description": "SendGrid is a cloud-based email delivery platform for transactional and marketing emails.",
        "dns": {
          "TXT": [
            "sendgrid\\.net"
          ]
        },
        "icon": "SendGrid.svg",
        "website": "https://sendgrid.com/"
      },
      "Sensors Data": {
        "cats": [
          10
        ],
        "cookies": {
          "sensorsdata2015jssdkcross": "",
          "sensorsdata2015session": ""
        },
        "icon": "Sensors Data.svg",
        "js": {
          "sa.lib_version": "([\\d.]+)\\;version:\\1",
          "sensorsdata_app_js_bridge_call_js": ""
        },
        "scripts": "sensorsdata",
        "website": "https://www.sensorsdata.cn"
      },
      "Sentry": {
        "cats": [
          13
        ],
        "description": "Sentry is an open-source platform for workflow productivity, aggregating errors from across the stack in real time.",
        "html": [
          "<script[^>]*>\\s*Raven\\.config\\('[^']*', \\{\\s+release: '([0-9\\.]+)'\\;version:\\1",
          "<script[^>]*src=\"[^\"]*browser\\.sentry\\-cdn\\.com/([0-9.]+)/bundle(?:\\.tracing)?(?:\\.min)?\\.js\\;version:\\1"
        ],
        "icon": "Sentry.svg",
        "js": {
          "Raven.config": "",
          "Sentry": "",
          "Sentry.SDK_VERSION": "(.+)\\;version:\\1",
          "__SENTRY__": "",
          "ravenOptions.whitelistUrls": ""
        },
        "scripts": [
          "browser\\.sentry\\-cdn\\.com/([0-9.]+)/bundle(?:\\.tracing)?(?:\\.min)?\\.js\\;version:\\1"
        ],
        "website": "https://sentry.io/"
      },
      "Seravo": {
        "cats": [
          62
        ],
        "headers": {
          "x-powered-by": "^Seravo"
        },
        "icon": "seravo.svg",
        "implies": "WordPress",
        "website": "https://seravo.com"
      },
      "Serendipity": {
        "cats": [
          1,
          11
        ],
        "icon": "Serendipity.png",
        "implies": "PHP",
        "meta": {
          "Powered-By": "Serendipity v\\.([\\d.]+)\\;version:\\1",
          "generator": "Serendipity(?: v\\.([\\d.]+))?\\;version:\\1"
        },
        "website": "http://s9y.org"
      },
      "Service Provider Pro": {
        "cats": [
          41,
          53
        ],
        "cookies": {
          "spp_csrf": "\\;confidence:25",
          "spp_orderform": "",
          "spp_session": "\\;confidence:25"
        },
        "description": "Service Provider Pro is a client management & billing software for productized service agencies.",
        "icon": "Service Provider Pro.svg",
        "js": {
          "sppOrderform": ""
        },
        "meta": {
          "server": "app.spp.co"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "js/spp_clients\\.js\\;confidence:50",
          "\\.spp\\.io/js/"
        ],
        "website": "https://spp.co"
      },
      "ServiceNow": {
        "cats": [
          19
        ],
        "description": "ServiceNow is a cloud computing platform to help companies manage digital workflows for enterprise operations.",
        "dom": {
          "a[href*='.service-now.com']": {
            "attributes": {
              "href": "https://[^.]+\\.service-now\\.com/.+"
            }
          }
        },
        "icon": "ServiceNow.svg",
        "pricing": [
          "poa",
          "high"
        ],
        "website": "https://www.servicenow.com"
      },
      "Setmore": {
        "cats": [
          5,
          72
        ],
        "description": "Setmore is a cloud-based appointment scheduling solution.",
        "icon": "Setmore.svg",
        "js": {
          "setmorePopup": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "my\\.setmore\\.\\w+/",
          "/setmore-appointments/script/"
        ],
        "website": "https://www.setmore.com"
      },
      "SevenRooms": {
        "cats": [
          5,
          72
        ],
        "description": "SevenRooms is an fully-integrated reservation, seating and restaurant management system.",
        "dom": {
          "iframe[src*='sevenrooms']": {
            "attributes": {
              "src": "\\.sevenrooms\\.\\w+/"
            }
          }
        },
        "icon": "SevenRooms.svg",
        "js": {
          "SevenroomsWidget": ""
        },
        "scripts": "sevenrooms\\.\\w+/widget/embed\\.js",
        "website": "https://sevenrooms.com"
      },
      "Shapecss": {
        "cats": [
          66
        ],
        "html": "<link[^>]* href=\"[^\"]*shapecss(?:\\.min)?\\.css",
        "icon": "Shapecss.svg",
        "js": {
          "Shapecss": ""
        },
        "scripts": [
          "shapecss[-.]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "/([\\d.]+)/shapecss(?:\\.min)?\\.js\\;version:\\1",
          "shapecss.*\\.js"
        ],
        "website": "https://shapecss.com"
      },
      "ShareThis": {
        "cats": [
          5
        ],
        "description": "ShareThis provides free engagement and growth tools (e.g., share buttons, follow buttons, and reaction buttons) for site owners.",
        "icon": "ShareThis.png",
        "js": {
          "SHARETHIS": ""
        },
        "scripts": "w\\.sharethis\\.com/",
        "website": "http://sharethis.com"
      },
      "SharpSpring": {
        "cats": [
          32
        ],
        "description": "SharpSpring is a cloud-based marketing tool that offers customer relationship management, marketing automation, mobile and social marketing, sales team automation, customer service and more, all within one solution.",
        "icon": "SharpSpring.png",
        "js": {
          "sharpspring_tracking_installed": ""
        },
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.marketingautomation\\.services.+(?:ver=)([\\d.]+)\\;version:\\1",
        "website": "https://sharpspring.com"
      },
      "SharpSpring Ads": {
        "cats": [
          77
        ],
        "description": "SharpSpring Ads is an all-in-one retargeting platform.",
        "icon": "SharpSpring.png",
        "js": {
          "_pa": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.perfectaudience\\.com",
        "website": "https://sharpspring.com/ads"
      },
      "ShellInABox": {
        "cats": [
          46
        ],
        "description": "Shell In A Box implements a web server that can export arbitrary command line tools to a web based terminal emulator.",
        "html": [
          "<title>Shell In A Box</title>",
          "must be enabled for ShellInABox</noscript>"
        ],
        "icon": "ShellInABox.png",
        "js": {
          "ShellInABox": ""
        },
        "website": "http://shellinabox.com"
      },
      "Shiny": {
        "cats": [
          18
        ],
        "icon": "Shiny.png",
        "js": {
          "Shiny.addCustomMessageHandler": ""
        },
        "website": "https://shiny.rstudio.com"
      },
      "ShinyStat": {
        "cats": [
          10
        ],
        "html": "<img[^>]*\\s+src=['\"]?https?://www\\.shinystat\\.com/cgi-bin/shinystat\\.cgi\\?[^'\"\\s>]*['\"\\s/>]",
        "icon": "ShinyStat.png",
        "js": {
          "SSsdk": ""
        },
        "scripts": "^https?://codice(?:business|ssl|pro|isp)?\\.shinystat\\.com/cgi-bin/getcod\\.cgi",
        "website": "http://shinystat.com"
      },
      "Shop Pay": {
        "cats": [
          41
        ],
        "description": "Shop Pay is an accelerated checkout that lets customers save their email address, credit card, and shipping and billing information so they can complete their transaction faster the next time they are directed to the Shopify checkout.",
        "dom": {
          "[aria-labelledby='pi-shopify_pay']": {
            "text": ""
          },
          "ul[data-shopify-buttoncontainer] li": {
            "text": "ShopPay"
          }
        },
        "icon": "Shopify.svg",
        "scripts": [
          "cdn\\.shopify\\.com/shopifycloud/shopify_pay/"
        ],
        "website": "https://shop.app"
      },
      "ShopGold": {
        "cats": [
          6
        ],
        "cookies": {
          "eGold": "^\\w+$",
          "popup_shopGold": "",
          "popup_shopGold_time": ""
        },
        "description": "ShopGold is an all-in-one payment processing and ecommerce solution.",
        "icon": "ShopGold.svg",
        "website": "https://www.shopgold.pl"
      },
      "Shopatron": {
        "cats": [
          6
        ],
        "html": [
          "<body class=\"shopatron",
          "<img[^>]+mediacdn\\.shopatron\\.com\\;confidence:50"
        ],
        "icon": "Shopatron.png",
        "js": {
          "shptUrl": ""
        },
        "meta": {
          "keywords": "Shopatron"
        },
        "scripts": "mediacdn\\.shopatron\\.com",
        "website": "http://ecommerce.shopatron.com"
      },
      "Shopcada": {
        "cats": [
          6
        ],
        "icon": "Shopcada.png",
        "js": {
          "Shopcada": ""
        },
        "website": "http://shopcada.com"
      },
      "Shoper": {
        "cats": [
          6
        ],
        "icon": "Shoper.svg",
        "js": {
          "shoper": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "website": "https://www.shoper.pl"
      },
      "Shopery": {
        "cats": [
          6
        ],
        "headers": {
          "X-Shopery": ""
        },
        "icon": "Shopery.svg",
        "implies": [
          "PHP",
          "Symfony",
          "Elcodi"
        ],
        "website": "http://shopery.com"
      },
      "Shopfa": {
        "cats": [
          6
        ],
        "headers": {
          "X-Powered-By": "^ShopFA ([\\d.]+)$\\;version:\\1"
        },
        "icon": "Shopfa.svg",
        "js": {
          "shopfa": ""
        },
        "meta": {
          "generator": "^ShopFA ([\\d.]+)$\\;version:\\1"
        },
        "website": "https://shopfa.com"
      },
      "Shopify": {
        "cats": [
          6
        ],
        "cookies": {
          "_shopify_s": "",
          "_shopify_y": ""
        },
        "description": "Shopify is a subscription-based software that allows anyone to set up an online store and sell their products. Shopify store owners can also sell in physical locations using Shopify POS, a point-of-sale app and accompanying hardware.",
        "dom": {
          "link[href*='cdn.shopify.com']": {
            "attributes": {
              "href": "cdn\\.shopify\\.com\\;confidence:50"
            }
          }
        },
        "headers": {
          "x-shopid": "\\;confidence:50",
          "x-shopify-stage": "\\;confidence:50"
        },
        "icon": "Shopify.svg",
        "js": {
          "Shopify": "\\;confidence:25",
          "ShopifyAPI": ""
        },
        "meta": {
          "shopify-checkout-api-token": "",
          "shopify-digital-wallet": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "sdks\\.shopifycdn\\.com",
          "cdn\\.shopify\\.com"
        ],
        "url": "^https?//.+\\.myshopify\\.com",
        "website": "http://shopify.com"
      },
      "Shoplo": {
        "cats": [
          6
        ],
        "description": "Shoplo is an all-in-one ecommerce platform.",
        "icon": "Shoplo.svg",
        "js": {
          "SHOPLOAJAX": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.shoplo\\.\\w+/",
        "website": "https://www.shoplo.com"
      },
      "Shoporama": {
        "cats": [
          6
        ],
        "icon": "Shoporama.svg",
        "meta": {
          "generator": "Shoporama"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "robots": "# Powered by Shoporama",
        "website": "https://www.shoporama.dk"
      },
      "Shoppy": {
        "cats": [
          6
        ],
        "description": "Shoppy is an all-in-one payment processing and ecommerce solution.",
        "icon": "Shoppy.svg",
        "js": {
          "Shoppy": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.shoppy\\.gg",
        "website": "https://shoppy.gg"
      },
      "Shoptet": {
        "cats": [
          6
        ],
        "html": "<link [^>]*href=\"https?://cdn\\.myshoptet\\.com/",
        "icon": "Shoptet.svg",
        "implies": "PHP",
        "js": {
          "shoptet": ""
        },
        "meta": {
          "web_author": "^Shoptet"
        },
        "scripts": [
          "^https?://cdn\\.myshoptet\\.com/"
        ],
        "website": "http://www.shoptet.cz"
      },
      "Shopware": {
        "cats": [
          6
        ],
        "description": "Shopware is an enterprise-level ecommerce platform.",
        "headers": {
          "sw-context-token": "^[\\w]{32}$\\;version:6",
          "sw-invalidation-states": "\\;version:6",
          "sw-language-id": "^[a-fA-F0-9]{32}$\\;version:6",
          "sw-version-id": "\\;version:6"
        },
        "html": "<title>Shopware ([\\d\\.]+) [^<]+\\;version:\\1",
        "icon": "Shopware.svg",
        "implies": [
          "PHP",
          "MySQL",
          "jQuery",
          "Symfony"
        ],
        "meta": {
          "application-name": "Shopware"
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "(?:(shopware)|/web/cache/[0-9]{10}_.+)\\.js\\;version:\\1?4:5",
          "/jquery\\.shopware\\.min\\.js",
          "/engine/Shopware/"
        ],
        "website": "https://www.shopware.com"
      },
      "Shuttle": {
        "cats": [
          1
        ],
        "description": "Shuttle is a website development platform.",
        "icon": "Devisto.svg",
        "implies": [
          "Laravel",
          "PHP",
          "Amazon Web Services"
        ],
        "js": {
          "Shuttle.FrontApp": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "shuttle(?:-assets-new|-storage)\\.s3\\.amazonaws\\.com",
        "website": "https://www.devisto.com"
      },
      "Sift": {
        "cats": [
          10,
          16
        ],
        "description": "Sift is a CA-based fraud prevention company.",
        "icon": "Sift.png",
        "js": {
          "__siftFlashCB": "",
          "_sift": ""
        },
        "pricing": [
          "mid",
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.sift(?:science)?\\.com/s\\.js",
        "website": "https://sift.com/"
      },
      "Signal": {
        "cats": [
          32,
          42
        ],
        "description": "Signal is a cross-platform encrypted messaging service.",
        "icon": "signal.png",
        "js": {
          "signalData": ""
        },
        "scripts": [
          "//s\\.btstatic\\.com/tag\\.js",
          "//s\\.thebrighttag\\.com/iframe\\?"
        ],
        "website": "https://www.signal.co/"
      },
      "Signifyd": {
        "cats": [
          10,
          16
        ],
        "description": "Signifyd is a provider of an enterprise-grade fraud technology solution for ecommerce stores.",
        "icon": "Signifyd.svg",
        "js": {
          "SIGNIFYD_GLOBAL": ""
        },
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.signifyd\\.com",
        "website": "https://www.signifyd.com"
      },
      "SilverStripe": {
        "cats": [
          1
        ],
        "description": "Silverstripe CMS is a free and open source Content Management System and Framework for creating and maintaining websites and web applications.",
        "html": "Powered by <a href=\"[^>]+SilverStripe",
        "icon": "SilverStripe.svg",
        "implies": "PHP",
        "meta": {
          "generator": "^SilverStripe"
        },
        "oss": true,
        "website": "https://www.silverstripe.org"
      },
      "Simbel": {
        "cats": [
          6
        ],
        "headers": {
          "powered": "simbel"
        },
        "icon": "simbel.svg",
        "website": "http://simbel.com.ar/"
      },
      "SimpleHTTP": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "SimpleHTTP(?:/([\\d.]+))?\\;version:\\1"
        },
        "website": "http://example.com"
      },
      "Simplo7": {
        "cats": [
          6
        ],
        "description": "Simplo7 is an all-in-one ecommerce product.",
        "icon": "Simplo7.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.simplo7\\.\\w+/",
        "website": "https://www.simplo7.com.br"
      },
      "Simplébo": {
        "cats": [
          1
        ],
        "headers": {
          "X-ServedBy": "simplebo"
        },
        "icon": "Simplebo.png",
        "website": "https://www.simplebo.fr"
      },
      "Site Kit": {
        "cats": [
          10
        ],
        "description": "Site Kit is a one-stop solution for WordPress users to use everything Google has to offer to make them successful on the web.",
        "icon": "Google.svg",
        "implies": "WordPress",
        "meta": {
          "generator": "^Site Kit by Google ?([\\d.]+)?\\;version:\\1"
        },
        "website": "https://sitekit.withgoogle.com/"
      },
      "Site Meter": {
        "cats": [
          10
        ],
        "icon": "Site Meter.png",
        "scripts": "sitemeter\\.com/js/counter\\.js\\?site=",
        "website": "http://www.sitemeter.com"
      },
      "Site24x7": {
        "cats": [
          78
        ],
        "description": "Site24x7 is a cloud-based website and server monitoring platform.",
        "icon": "Site24x7.png",
        "js": {
          "S247RumQueueImpl": "",
          "s247RUM": "",
          "site24x7RumError": "",
          "site24x7rum": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.site24x7rum\\.com",
        "website": "https://www.site24x7.com"
      },
      "SiteEdit": {
        "cats": [
          1
        ],
        "icon": "SiteEdit.png",
        "meta": {
          "generator": "SiteEdit"
        },
        "website": "http://www.siteedit.ru"
      },
      "SiteGround": {
        "cats": [
          62
        ],
        "description": "SiteGround is a web hosting service.",
        "headers": {
          "host-header": "192fc2e7e50945beb8231a492d6a8024|b7440e60b07ee7b8044761568fab26e8|624d5be7be38418a3e2a818cc8b7029b|6b7412fb82ca5edfd0917e3957f05d89"
        },
        "icon": "siteground.svg",
        "website": "https://www.siteground.com"
      },
      "SiteMinder": {
        "cats": [
          5,
          72
        ],
        "description": "SiteMinder is a appointment booking solution designed for hotels.",
        "icon": "SiteMinder.svg",
        "pricing": [
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "widget\\.siteminder\\.com",
        "website": "https://www.siteminder.com"
      },
      "SiteSpect": {
        "cats": [
          74
        ],
        "description": "SiteSpect is the A/B testing and optimisation solution.",
        "icon": "SiteSpect.png",
        "js": {
          "ss_dom_var": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "/__ssobj/core\\.js",
        "website": "https://www.sitespect.com"
      },
      "Sitecore": {
        "cats": [
          1
        ],
        "cookies": {
          "SC_ANALYTICS_GLOBAL_COOKIE": ""
        },
        "description": "Sitecore provides web content management, and multichannel marketing automation software.",
        "html": "<img[^>]+src=\"[^>]*/~/media/[^>]+\\.ashx",
        "icon": "Sitecore.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "http://sitecore.net"
      },
      "Sitefinity": {
        "cats": [
          1
        ],
        "description": "Sitefinity is a content management system (CMS) that you use to create, store, manage, and present content on your website.",
        "icon": "Sitefinity.svg",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "^Sitefinity (.+)$\\;version:\\1"
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "http://www.sitefinity.com"
      },
      "Siteglide": {
        "cats": [
          1,
          53,
          6
        ],
        "description": "SiteGlide is a Digital Experience Platform (DEP) for ecommerce stores, membership sites and customer portals.",
        "icon": "Siteglide.svg",
        "implies": "PlatformOS",
        "pricing": [
          "low"
        ],
        "saas": true,
        "scripts": "siteglide\\.js",
        "website": "https://www.siteglide.com"
      },
      "Sitevision CMS": {
        "cats": [
          1
        ],
        "cookies": {
          "SiteVisionLTM": ""
        },
        "description": "Sitevision CMS is a platform for web publishing that consists of flexible and pre-made modules. Available as self-hosed software and Cloud SaaS.",
        "icon": "Sitevision CMS.png",
        "pricing": [
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": [
          "sitevision/system-resource/(?:[\\w\\d]+)/js/docready-min\\.js\\;confidence:25",
          "sitevision/system-resource/(?:[\\w\\d]+)/js/AppRegistry\\.js\\;confidence:25",
          "sitevision/system-resource/(?:[\\w\\d]+)/webapps/webapp_sdk-min\\.js\\;confidence:50",
          "sitevision/system-resource/(?:[\\w\\d]+)/envision/envision\\.js\\;confidence:50"
        ],
        "website": "https://www.sitevision.se"
      },
      "Sivuviidakko": {
        "cats": [
          1
        ],
        "icon": "Sivuviidakko.png",
        "meta": {
          "generator": "Sivuviidakko"
        },
        "website": "http://sivuviidakko.fi"
      },
      "Sizmek": {
        "cats": [
          36
        ],
        "html": "(?:<a [^>]*href=\"[^/]*//[^/]*serving-sys\\.com/|<img [^>]*src=\"[^/]*//[^/]*serving-sys\\.com/)",
        "icon": "Sizmek.png",
        "scripts": "serving-sys\\.com/",
        "website": "http://sizmek.com"
      },
      "Skedify": {
        "cats": [
          72
        ],
        "description": "Skedify is an appointment booking solution created for enterprises.",
        "icon": "Skedify.svg",
        "js": {
          "Skedify.Plugin.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "plugin\\.skedify\\.io",
        "website": "https://calendly.com/"
      },
      "Skimlinks": {
        "cats": [
          71
        ],
        "description": "Skimlinks is a content monetization platform for online publishers.",
        "icon": "Skimlinks.svg",
        "js": {
          "__SKIM_JS_GLOBAL__": "",
          "addSkimlinks": "",
          "skimlinksAPI": ""
        },
        "pricing": [
          "payg"
        ],
        "scripts": "\\.skimresources\\.com",
        "website": "https://skimlinks.com",
        "xhr": "\\.skimresources\\.com"
      },
      "Sky-Shop": {
        "cats": [
          6
        ],
        "description": "Sky-Shop.pl is a platform for dropshipping an online sales on Allegro, eBay and Amazon.",
        "dom": ".skyshop-container",
        "icon": "Sky-Shop.svg",
        "implies": [
          "PHP",
          "Bootstrap",
          "jQuery"
        ],
        "js": {
          "L.CONTINUE_SHOPPING": ""
        },
        "meta": {
          "generator": "Sky-Shop"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://sky-shop.pl"
      },
      "SkyVerge": {
        "cats": [
          41
        ],
        "description": "SkyVerge  is a company which develop extension tools for WooCommerce stores.",
        "icon": "SkyVerge.svg",
        "implies": "WooCommerce",
        "js": {
          "sv_wc_payment_gateway_payment_form_param": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "sv-wc-payment-gateway-payment-form\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
        "website": "https://www.skyverge.com"
      },
      "Slick": {
        "cats": [
          59
        ],
        "html": "<link [^>]+(?:/([\\d.]+)/)?slick-theme\\.css\\;version:\\1",
        "implies": "jQuery",
        "scripts": "(?:/([\\d.]+))?/slick(?:\\.min)?\\.js\\;version:\\1",
        "website": "https://kenwheeler.github.io/slick"
      },
      "SlickStack": {
        "cats": [
          47,
          9
        ],
        "description": "SlickStack is a free LEMP stack automation script written in Bash designed to enhance and simplify WordPress provisioning, performance, and security.",
        "headers": {
          "x-powered-by": "SlickStack"
        },
        "icon": "SlickStack.png",
        "implies": "WordPress",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://slickstack.io"
      },
      "Slimbox": {
        "cats": [
          59
        ],
        "html": "<link [^>]*href=\"[^/]*slimbox(?:-rtl)?\\.css",
        "icon": "Slimbox.png",
        "implies": "MooTools",
        "scripts": "slimbox\\.js",
        "website": "http://www.digitalia.be/software/slimbox"
      },
      "Slimbox 2": {
        "cats": [
          59
        ],
        "html": "<link [^>]*href=\"[^/]*slimbox2(?:-rtl)?\\.css",
        "icon": "Slimbox 2.png",
        "implies": "jQuery",
        "scripts": "slimbox2\\.js",
        "website": "http://www.digitalia.be/software/slimbox2"
      },
      "Smart Ad Server": {
        "cats": [
          36
        ],
        "html": "<img[^>]+smartadserver\\.com\\/call",
        "icon": "Smart Ad Server.png",
        "js": {
          "SmartAdServer": ""
        },
        "website": "http://smartadserver.com"
      },
      "SmartSite": {
        "cats": [
          1
        ],
        "html": "<[^>]+/smartsite\\.(?:dws|shtml)\\?id=",
        "icon": "SmartSite.png",
        "meta": {
          "author": "Redacteur SmartInstant"
        },
        "website": "http://www.seneca.nl/pub/Smartsite/Smartsite-Smartsite-iXperion"
      },
      "Smarter Click": {
        "cats": [
          77
        ],
        "description": "Smarter Click is a marketing technology company.",
        "icon": "Smarter Click.png",
        "js": {
          "$smcInstall": "",
          "$smcT5": "",
          "$smctData": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.?smct\\.(?:io|co)/",
        "website": "https://smarterclick.com"
      },
      "Smartstore": {
        "cats": [
          6
        ],
        "cookies": {
          "SMARTSTORE.CUSTOMER": "",
          "SMARTSTORE.VISITOR": ""
        },
        "description": "Smartstore is an open-source ecommerce system with CMS capabilities.",
        "html": [
          "<!--Powered by Smart[sS]tore",
          "<meta property=\"sm:pagedata\""
        ],
        "icon": "Smartstore.png",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "^Smart[sS]tore(.NET)? (.+)$\\;version:\\2"
        },
        "oss": true,
        "website": "https://www.smartstore.com"
      },
      "Smartstore Page Builder": {
        "cats": [
          1
        ],
        "css": "\\.g-stage \\.g-stage-root",
        "html": "<section[^>]+class=\"g-stage",
        "icon": "Smartstore.png",
        "implies": "Microsoft ASP.NET",
        "website": "https://www.smartstore.com"
      },
      "Smartstore biz": {
        "cats": [
          6
        ],
        "icon": "Smartstore.biz.png",
        "scripts": "smjslib\\.js",
        "website": "http://smartstore.com"
      },
      "Smartsupp": {
        "cats": [
          52
        ],
        "description": "Smartsupp is a live chat tool that offers visitor recording feature.",
        "icon": "Smartsupp.svg",
        "js": {
          "$smartsupp.options.widgetVersion": "([\\d.]+)\\;version:\\1",
          "smartsupp": ""
        },
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "scripts": "\\.smartsuppchat\\.com",
        "website": "https://www.smartsupp.com"
      },
      "SmugMug": {
        "cats": [
          7
        ],
        "description": "SmugMug is a paid image sharing, image hosting service, and online video platform on which users can upload photos and videos.",
        "dom": {
          ".sm-page-footer-copyright": {
            "text": "SmugMug"
          }
        },
        "headers": {
          "Smug-CDN": ""
        },
        "icon": "SmugMug.svg",
        "js": {
          "_smugsp": ""
        },
        "pricing": [
          "recurring",
          "low"
        ],
        "saas": true,
        "website": "https://www.smugmug.com"
      },
      "Snap": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "Server": "Snap/([.\\d]+)\\;version:\\1"
        },
        "icon": "Snap.png",
        "implies": "Haskell",
        "website": "http://snapframework.com"
      },
      "Snap Pixel": {
        "cats": [
          36
        ],
        "description": "Snap Pixel is a piece of JavaScript code that helps advertisers measure the cross-device impact of campaigns.",
        "icon": "Snap Pixel.svg",
        "js": {
          "__SnapPixel": "",
          "snaptr.pixelIdList": ""
        },
        "saas": true,
        "scripts": "intg\\.snapchat\\.com/",
        "website": "https://businesshelp.snapchat.com/s/article/snap-pixel-about"
      },
      "Snap.svg": {
        "cats": [
          59
        ],
        "icon": "Snap.svg.svg",
        "js": {
          "Snap.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "snap\\.svg(?:-min)?\\.js",
        "website": "http://snapsvg.io"
      },
      "SnapEngage": {
        "cats": [
          52
        ],
        "description": "SnapEngage is a live chat solution that caters to businesses across various industries.",
        "html": "<!-- begin SnapEngage",
        "icon": "SnapEngage.svg",
        "js": {
          "SnapEngage": "",
          "SnapEngageChat": "",
          "snapengage_mobile": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://snapengage.com/"
      },
      "Snipcart": {
        "cats": [
          6
        ],
        "cookies": {
          "snipcart-cart": ""
        },
        "dom": {
          "div#snipcart": {
            "text": ""
          },
          "link[href*='snipcart.css']": {
            "attributes": {
              "href": ""
            }
          }
        },
        "icon": "Snipcart.png",
        "scripts": "https://cdn\\.snipcart\\.com/themes/v([\\w.]+)/default/snipcart\\.js\\;version:\\1",
        "website": "https://snipcart.com"
      },
      "Snoobi": {
        "cats": [
          10
        ],
        "icon": "Snoobi.png",
        "js": {
          "snoobi": ""
        },
        "scripts": "snoobi\\.com/snoop\\.php",
        "website": "http://www.snoobi.com"
      },
      "Snowplow Analytics": {
        "cats": [
          10,
          63
        ],
        "cookies": {
          "_sp_id": "",
          "sp": "\\;confidence:50"
        },
        "description": "Snowplow is an open-source behavioral data management platform for businesses.",
        "icon": "Snowplow.svg",
        "js": {
          "GlobalSnowplowNamespace": "",
          "Snowplow": ""
        },
        "oss": true,
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "sp\\.js\\;confidence:50",
          "d1fc8wv8zag5ca\\.cloudfront\\.net/.+/sp\\.js",
          "cdn\\.jsdelivr\\.net/gh/snowplow/sp-js-assets@(?:.+)/sp\\.js",
          "cdnjs\\.cloudflare\\.com/ajax/libs/snowplow/(?:.+)/sp.*\\.js",
          "unpkg.com/@snowplow/javascript-tracker@(?:.+)/dist/sp.*\\.js",
          "cdn\\.jsdelivr\\.net/npm/@snowplow/javascript-tracker@(?:.+)/dist/sp\\.js"
        ],
        "website": "https://snowplowanalytics.com"
      },
      "SobiPro": {
        "cats": [
          19
        ],
        "icon": "SobiPro.png",
        "implies": "Joomla",
        "js": {
          "SobiProUrl": ""
        },
        "website": "http://sigsiu.net/sobipro.html"
      },
      "Social9": {
        "cats": [
          5
        ],
        "description": "Social9 is a social sharing widgets and plugins.",
        "icon": "Social9.svg",
        "pricing": [
          "poa",
          "freemium"
        ],
        "saas": true,
        "scripts": "social9\\.com/.+\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
        "website": "https://social9.com"
      },
      "Socket.io": {
        "cats": [
          12
        ],
        "icon": "Socket.IO.svg",
        "implies": "Node.js",
        "js": {
          "io.Socket": "",
          "io.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "socket\\.io.*\\.js",
        "website": "https://socket.io"
      },
      "SoftTr": {
        "cats": [
          6
        ],
        "icon": "softtr.png",
        "meta": {
          "author": "SoftTr E-Ticaret Sitesi Yazılımı"
        },
        "website": "http://www.softtr.com"
      },
      "SolidPixels": {
        "cats": [
          1,
          6,
          4
        ],
        "description": "Solidpixels is platform to build websites.",
        "icon": "SolidPixels.png",
        "implies": "React",
        "meta": {
          "web_author": "^solidpixels"
        },
        "scripts": [
          "^https?://cdn\\.solidpixels\\.net/"
        ],
        "website": "https://www.solidpixels.net"
      },
      "Solodev": {
        "cats": [
          1
        ],
        "headers": {
          "solodev_session": ""
        },
        "html": "<div class=[\"']dynamicDiv[\"'] id=[\"']dd\\.\\d\\.\\d(?:\\.\\d)?[\"']>",
        "icon": "Solodev.png",
        "implies": "PHP",
        "website": "http://www.solodev.com"
      },
      "Solr": {
        "cats": [
          34
        ],
        "description": "Solr is an open-source enterprise-search platform, written in Java.",
        "icon": "Solr.png",
        "implies": "Lucene",
        "website": "http://lucene.apache.org/solr/"
      },
      "Solusquare OmniCommerce Cloud": {
        "cats": [
          6
        ],
        "cookies": {
          "_solusquare": ""
        },
        "icon": "Solusquare.png",
        "implies": "Adobe ColdFusion",
        "meta": {
          "generator": "^Solusquare$"
        },
        "website": "https://www.solusquare.com"
      },
      "Solve Media": {
        "cats": [
          16,
          36
        ],
        "icon": "Solve Media.png",
        "js": {
          "ACPuzzle": "",
          "_ACPuzzle": "",
          "_adcopy-puzzle-image-image": "",
          "adcopy-puzzle-image-image": ""
        },
        "scripts": "^https?://api\\.solvemedia\\.com/",
        "website": "http://solvemedia.com"
      },
      "SonarQubes": {
        "cats": [
          47
        ],
        "description": "SonarQube is an open-source platform for the continuous inspection of code quality to perform automatic reviews with static analysis of code to detect bugs, code smells, and security vulnerabilities on 20+ programming languages.",
        "html": [
          "<link href=\"/css/sonar\\.css\\?v=([\\d.]+)\\;version:\\1",
          "<title>SonarQube</title>"
        ],
        "icon": "sonar.png",
        "implies": "Java",
        "js": {
          "SonarMeasures": "",
          "SonarRequest": ""
        },
        "meta": {
          "application-name": "^SonarQubes$"
        },
        "scripts": "^/js/bundles/sonar\\.js?v=([\\d.]+)$\\;version:\\1",
        "website": "https://www.sonarqube.org/"
      },
      "Sonobi": {
        "cats": [
          36
        ],
        "description": "Sonobi is an ad technology developer that designs advertising tools and solutions for the media publishers, brand advertisers and media agencies.",
        "icon": "Sonobi.png",
        "saas": true,
        "website": "https://sonobi.com",
        "xhr": "apex\\.go\\.sonobi\\.com"
      },
      "SoteShop": {
        "cats": [
          6
        ],
        "cookies": {
          "soteshop": "^\\w+$"
        },
        "description": "SoteShop is an e-shop management software.",
        "icon": "SoteShop.svg",
        "implies": "PHP",
        "website": "https://www.soteshop.com/"
      },
      "Sotel": {
        "cats": [
          1
        ],
        "icon": "Sotel.png",
        "meta": {
          "generator": "sotel"
        },
        "website": "https://www.soteledu.com/en/"
      },
      "SoundManager": {
        "cats": [
          59
        ],
        "icon": "SoundManager.png",
        "js": {
          "BaconPlayer": "",
          "SoundManager": "",
          "soundManager.version": "V(.+) \\;version:\\1"
        },
        "website": "http://www.schillmania.com/projects/soundmanager2"
      },
      "Sovrn": {
        "cats": [
          36
        ],
        "description": "Sovrn is a advertising products and services provider for publishers.",
        "icon": "Sovrn.png",
        "js": {
          "sovrn": "",
          "sovrn_render": ""
        },
        "saas": true,
        "scripts": "\\.lijit\\.com",
        "website": "https://www.sovrn.com",
        "xhr": "\\.lijit\\.com"
      },
      "Sovrn//Commerce": {
        "cats": [
          71
        ],
        "description": "Sovrn//Commerce is  a content monetization tool for publishers.",
        "icon": "Sovrn.png",
        "js": {
          "vglnk": "",
          "vl_cB": "",
          "vl_disable": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "(?:^[^/]*//[^/]*viglink\\.com/api/|vglnk\\.js)",
        "website": "https://www.sovrn.com/publishers/commerce/"
      },
      "SparkPost": {
        "cats": [
          75
        ],
        "description": "SparkPost is an email infrastructure provider.",
        "dns": {
          "TXT": [
            "sparkpostmail\\.com"
          ]
        },
        "icon": "SparkPost.svg",
        "website": "https://www.sparkpost.com/"
      },
      "SpeedCurve": {
        "cats": [
          78
        ],
        "description": "SpeedCurve is a front-end performance monitoring service.",
        "icon": "SpeedCurve.svg",
        "js": {
          "LUX.version": "([\\d.]+)\\;version:\\1",
          "LUX_t_end": "\\d+\\;confidence:50",
          "LUX_t_start": "\\d+\\;confidence:50"
        },
        "pricing": [
          "mid",
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.speedcurve\\.com",
        "website": "https://www.speedcurve.com"
      },
      "Sphinx": {
        "cats": [
          4
        ],
        "description": "Sphinx is a tool that makes it easy to create documentation.",
        "html": "Created using <a href=\"https?://(?:www\\.)?sphinx-doc\\.org/\">Sphinx</a> ([0-9.]+)\\.\\;version:\\1",
        "icon": "Sphinx.png",
        "js": {
          "DOCUMENTATION_OPTIONS": ""
        },
        "website": "https://www.sphinx-doc.org/"
      },
      "Spinnakr": {
        "cats": [
          10
        ],
        "description": "Spinnakr is a startup with a platform designed to personalise messages on blogs and websites.",
        "icon": "Spinnakr.png",
        "js": {
          "_spinnakr_site_id": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.spinnakr.com"
      },
      "Splitbee": {
        "cats": [
          10
        ],
        "icon": "splitbee.svg",
        "js": {
          "splitbee": ""
        },
        "scripts": "^https://cdn\\.splitbee\\.io/sb\\.js",
        "website": "https://splitbee.io"
      },
      "Splunk": {
        "cats": [
          19
        ],
        "html": "<p class=\"footer\">&copy; [-\\d]+ Splunk Inc\\.(?: Splunk ([\\d\\.]+(?: build [\\d\\.]*\\d)?))?[^<]*</p>\\;version:\\1",
        "icon": "Splunk.png",
        "meta": {
          "author": "Splunk Inc\\;confidence:50"
        },
        "website": "http://splunk.com"
      },
      "Splunkd": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Splunkd"
        },
        "icon": "Splunk.png",
        "website": "http://splunk.com"
      },
      "SpotX": {
        "cats": [
          36
        ],
        "description": "SpotX is a video advertising platform.",
        "dom": {
          "link[href*='.spotxchange.com']": {
            "attributes": {
              "href": ""
            }
          }
        },
        "icon": "SpotX.png",
        "js": {
          "SpotX.VERSION": "([\\d.]+)\\;version:\\1"
        },
        "saas": true,
        "scripts": "js\\.spotx\\.tv",
        "website": "https://www.spotx.tv",
        "xhr": "\\.spotx(?:change|cdn)\\.com"
      },
      "Spree": {
        "cats": [
          6
        ],
        "html": "(?:<link[^>]*/assets/store/all-[a-z\\d]{32}\\.css[^>]+>|<script>\\s*Spree\\.(?:routes|translations|api_key))",
        "icon": "Spree.png",
        "implies": "Ruby on Rails",
        "website": "https://spreecommerce.org"
      },
      "Spring": {
        "cats": [
          18
        ],
        "headers": {
          "X-Application-Context": ""
        },
        "icon": "Spring.png",
        "implies": "Java",
        "website": "https://spring.io/"
      },
      "Sqreen": {
        "cats": [
          19
        ],
        "headers": {
          "X-Protected-By": "^Sqreen$"
        },
        "icon": "Sqreen.png",
        "website": "https://sqreen.io"
      },
      "Square": {
        "cats": [
          41
        ],
        "description": "Square is a mobile payment company that offers business software, payment hardware products and small business services.",
        "icon": "Square.svg",
        "js": {
          "SqPaymentForm": "",
          "Square.Analytics": "",
          "__BOOTSTRAP_STATE__.storeInfo.square_application_id": ""
        },
        "xhr": "\\.squareup\\.com",
        "scripts": "js\\.squareup\\.com",
        "website": "https://squareup.com/"
      },
      "Squarespace": {
        "cats": [
          1
        ],
        "description": "Squarespace provides Software-as-a-Service (SaaS) for website building and hosting, and allows users to use pre-built website templates.",
        "headers": {
          "X-ServedBy": "squarespace"
        },
        "html": "<!-- This is Squarespace\\. -->",
        "icon": "Squarespace.png",
        "js": {
          "Squarespace": "",
          "Static.SQUARESPACE_CONTEXT": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "http://www.squarespace.com"
      },
      "Squarespace Commerce": {
        "cats": [
          6
        ],
        "description": "Squarespace Commerce is an ecommerce platform designed to facilitate the creation of websites and online stores, with domain registration and web hosting included.",
        "icon": "Squarespace.png",
        "implies": "Squarespace",
        "js": {
          "SQUARESPACE_ROLLUPS.squarespace-commerce": "",
          "SquarespaceCommerceCartBundle": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "assets\\.squarespace\\.\\w+/universal/scripts-compressed/commerce-\\w+-min\\.[\\w+\\-]+\\.js",
        "website": "https://www.squarespace.com/ecommerce-website"
      },
      "SquirrelMail": {
        "cats": [
          30
        ],
        "description": "SquirrelMail is an open-source web-mail package that is based on PHP language. It provides a web-based-email client and a proxy server for IMAP protocol.",
        "html": "<small>SquirrelMail version ([.\\d]+)[^<]*<br \\;version:\\1",
        "icon": "SquirrelMail.png",
        "implies": "PHP",
        "js": {
          "squirrelmail_loginpage_onload": ""
        },
        "url": "/src/webmail\\.php(?:$|\\?)",
        "website": "http://squirrelmail.org"
      },
      "Squiz Matrix": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "Squiz Matrix"
        },
        "html": "<!--\\s+Running (?:MySource|Squiz) Matrix",
        "icon": "Squiz Matrix.png",
        "implies": "PHP",
        "meta": {
          "generator": "Squiz Matrix"
        },
        "website": "http://squiz.net"
      },
      "Stack Analytix": {
        "cats": [
          10
        ],
        "description": "Stack Analytix offers heatmaps, session recording, conversion analysis and user engagement tools.",
        "icon": "Stack Analytix.svg",
        "js": {
          "stackAnalysis": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "api\\.stackanalytix\\.com",
        "website": "https://www.stackanalytix.com"
      },
      "StackPath": {
        "cats": [
          31
        ],
        "description": "StackPath is a cloud computing and services provider.",
        "headers": {
          "x-backend-server": "hosting\\.stackcp\\.net$",
          "x-provided-by": "^StackCDN(?: ([\\d.]+))?\\;version:\\1"
        },
        "icon": "StackPath.svg",
        "website": "https://www.stackpath.com"
      },
      "Stackla": {
        "cats": [
          5
        ],
        "icon": "Stackla.png",
        "js": {
          "Stackla": ""
        },
        "scripts": "assetscdn\\.stackla\\.com\\/media\\/js\\/widget\\/(?:[a-zA-Z0-9.]+)?\\.js",
        "website": "http://stackla.com/"
      },
      "Starhost": {
        "cats": [
          51,
          62
        ],
        "description": "Starhost provides a Platform-as-a-Service (PaaS) for website building, web hosting, and domain registration.",
        "headers": {
          "Cache-Control": "Starhost",
          "X-Starhost": ""
        },
        "icon": "Starhost.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://starhost.verbosec.com"
      },
      "Starlet": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^Plack::Handler::Starlet"
        },
        "icon": "Starlet.png",
        "implies": "Perl",
        "website": "http://metacpan.org/pod/Starlet"
      },
      "Statamic": {
        "cats": [
          1
        ],
        "description": "Statamic is an open-source and self-hosted content management system based on the PHP programming language.",
        "headers": {
          "x-powered-by": "^Statamic$"
        },
        "icon": "Statamic.svg",
        "implies": [
          "PHP",
          "Laravel"
        ],
        "oss": true,
        "pricing": [
          "freemium",
          "mid",
          "payg"
        ],
        "saas": false,
        "website": "https://statamic.com"
      },
      "Statcounter": {
        "cats": [
          10
        ],
        "icon": "Statcounter.svg",
        "js": {
          "_statcounter": "",
          "sc_project": "\\;confidence:50",
          "sc_security": "\\;confidence:50"
        },
        "scripts": "statcounter\\.com/counter/counter",
        "website": "http://www.statcounter.com"
      },
      "Statically": {
        "cats": [
          31
        ],
        "headers": {
          "Server": "^statically$"
        },
        "html": "<link [^>]*?href=\"?[a-z]*?:?//cdn\\.statically\\.io/",
        "icon": "Statically.svg",
        "website": "https://statically.io"
      },
      "SteelHouse": {
        "cats": [
          77
        ],
        "description": "SteelHouse is an advertising software company which provides services for brands, agencies, and direct marketers.",
        "icon": "SteelHouse.png",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.steelhousemedia\\.com/(?:spx\\?dxver=([\\d.]+))?\\;version:\\1",
        "website": "https://steelhouse.com"
      },
      "Stimulus": {
        "cats": [
          12
        ],
        "description": "A modest JavaScript framework for the HTML you already have.",
        "html": "<[^>]+data-controller",
        "icon": "Stimulus.png",
        "website": "https://stimulusjs.org/"
      },
      "Store Systems": {
        "cats": [
          6
        ],
        "html": "Shopsystem von <a href=[^>]+store-systems\\.de\"",
        "icon": "Store Systems.png",
        "website": "http://store-systems.de"
      },
      "Storeden": {
        "cats": [
          6
        ],
        "headers": {
          "X-Powered-By": "Storeden"
        },
        "icon": "storeden.svg",
        "website": "https://www.storeden.com"
      },
      "Storyblok": {
        "cats": [
          1
        ],
        "icon": "storyblok.png",
        "meta": {
          "generator": "storyblok"
        },
        "website": "https://www.storyblok.com"
      },
      "Strapdown.js": {
        "cats": [
          12
        ],
        "icon": "strapdown.js.png",
        "implies": [
          "Bootstrap",
          "Google Code Prettify"
        ],
        "scripts": "strapdown\\.js",
        "website": "http://strapdownjs.com"
      },
      "Strapi": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "^Strapi"
        },
        "icon": "Strapi.png",
        "website": "https://strapi.io"
      },
      "Strato": {
        "cats": [
          6
        ],
        "html": "<a href=\"http://www\\.strato\\.de/\" target=\"_blank\">",
        "icon": "strato.png",
        "website": "http://shop.strato.com"
      },
      "Strikingly": {
        "cats": [
          1
        ],
        "html": "<!-- Powered by Strikingly\\.com",
        "icon": "Strikingly.png",
        "website": "https://strikingly.com"
      },
      "Stripe": {
        "cats": [
          41
        ],
        "cookies": {
          "__stripe_mid": "",
          "__stripe_sid": ""
        },
        "description": "Stripe offers online payment processing for internet businesses as well as fraud prevention, invoicing and subscription management.",
        "dns": {
          "TXT": "stripe-verification="
        },
        "html": "<input[^>]+data-stripe",
        "icon": "Stripe.svg",
        "js": {
          "Stripe.version": "^(.+)$\\;version:\\1"
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "js\\.stripe\\.com",
        "website": "http://stripe.com"
      },
      "Stylitics": {
        "cats": [
          5,
          32
        ],
        "description": "Stylitics is a cloud-based SaaS platform for retailers to automate and distribute visual content at scale.",
        "dom": "link[href*='.stylitics.com']",
        "icon": "Stylitics.svg",
        "js": {
          "Stylitics": "",
          "stylitics": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "\\.stylitics\\.com/v([\\d.]+)\\;version:\\1",
          "/stylitics/js/stylitics\\.js\\?ver=v([\\d.]+)\\;version:\\1"
        ],
        "website": "https://stylitics.com"
      },
      "SublimeVideo": {
        "cats": [
          14
        ],
        "description": "SublimeVideo is a framework for HTML5 Video Player.",
        "icon": "SublimeVideo.png",
        "js": {
          "sublimevideo": ""
        },
        "scripts": "cdn\\.sublimevideo\\.net/js/[a-z\\d]+\\.js",
        "website": "http://sublimevideo.net"
      },
      "Subrion": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-CMS": "Subrion CMS"
        },
        "icon": "Subrion.png",
        "implies": "PHP",
        "meta": {
          "generator": "^Subrion "
        },
        "website": "http://subrion.com"
      },
      "Sucuri": {
        "cats": [
          31
        ],
        "description": "Sucuri is a WordPress security plugin used to protect websites from malware and hacks.",
        "headers": {
          "x-sucuri-cache:": "",
          "x-sucuri-id": ""
        },
        "icon": "sucuri.png",
        "website": "https://sucuri.net/"
      },
      "Sulu": {
        "cats": [
          1
        ],
        "headers": {
          "X-Generator": "Sulu/?(.+)?$\\;version:\\1"
        },
        "icon": "Sulu.svg",
        "implies": "Symfony",
        "website": "http://sulu.io"
      },
      "SummerCart": {
        "cats": [
          6
        ],
        "description": "SummerCart is an ecommerce platform written in PHP.",
        "dom": {
          "link": {
            "attributes": {
              "href": "com\\.summercart\\.\\;confidence:100"
            }
          }
        },
        "icon": "SummerCart.svg",
        "implies": "PHP",
        "js": {
          "SC": "\\;confidence:10",
          "SCEvents": "\\;confidence:40"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "http://www.summercart.com"
      },
      "SumoMe": {
        "cats": [
          5,
          32
        ],
        "description": "Sumo is a plugin offering set of marketing tools to automate a website's growth. These tools help drive extra traffic, engage visitors, increase email subscribers and build community.",
        "icon": "SumoMe.png",
        "scripts": "load\\.sumome\\.com",
        "website": "http://sumome.com"
      },
      "SunOS": {
        "cats": [
          28
        ],
        "description": "SunOS is a Unix-branded operating system developed by Sun Microsystems for their workstation and server computer systems.",
        "headers": {
          "Server": "SunOS( [\\d\\.]+)?\\;version:\\1",
          "Servlet-engine": "SunOS( [\\d\\.]+)?\\;version:\\1"
        },
        "icon": "Oracle.png",
        "website": "http://oracle.com/solaris"
      },
      "Super Socializer": {
        "cats": [
          69
        ],
        "description": "Super Socializer is a multipurpose social media plugin for WordPress.",
        "icon": "Super Socializer.png",
        "implies": "WordPress",
        "oss": true,
        "scripts": "plugins/super-socializer/.+?ver=([\\d.]+)\\;version:\\1",
        "website": "https://super-socializer-wordpress.heateor.com"
      },
      "Supersized": {
        "cats": [
          25
        ],
        "icon": "Supersized.png",
        "scripts": "supersized(?:\\.([\\d.]*[\\d]))?.*\\.js\\;version:\\1",
        "website": "http://buildinternet.com/project/supersized"
      },
      "Svbtle": {
        "cats": [
          11
        ],
        "icon": "svbtle.png",
        "meta": {
          "generator": "^Svbtle\\.com$"
        },
        "url": "^https?://[^/]+\\.svbtle\\.com",
        "website": "https://www.svbtle.com"
      },
      "Svelte": {
        "cats": [
          12
        ],
        "description": "Svelte is a free and open-source front end compiler created by Rich Harris and maintained by the Svelte core team members.",
        "html": "<[^>]+class=\\\"[^\\\"]+\\ssvelte-[\\w]*\\\"",
        "icon": "Svelte.svg",
        "oss": true,
        "website": "https://svelte.dev"
      },
      "SvelteKit": {
        "cats": [
          66
        ],
        "description": "SvelteKit is the official Svelte framework for building web applications with a flexible filesystem-based routing.",
        "dom": {
          "a": {
            "attributes": {
              "sveltekit:prefetch": ""
            }
          }
        },
        "icon": "Svelte.svg",
        "oss": true,
        "website": "https://kit.svelte.dev"
      },
      "SweetAlert": {
        "cats": [
          59
        ],
        "description": "SweetAlert is a beautiful replacement for JavaScript's alert.",
        "html": "<link[^>]+?href=\"[^\"]+sweet-alert(?:\\.min)?\\.css",
        "icon": "SweetAlert.png",
        "oss": true,
        "scripts": "sweet(?:-)?alert(?:\\.min)?\\.js",
        "website": "https://sweetalert.js.org"
      },
      "SweetAlert2": {
        "cats": [
          59
        ],
        "description": "SweetAlert2 is a beautiful, responsive, customizable, accessible replacement for JavaScript's popup boxes.",
        "excludes": "SweetAlert",
        "html": "<link[^>]+?href=\"[^\"]+sweetalert2(?:\\.min)?\\.css",
        "icon": "SweetAlert2.png",
        "js": {
          "Sweetalert2": ""
        },
        "oss": true,
        "scripts": [
          "sweetalert2(?:\\.all)?(?:\\.min)?\\.js",
          "/npm/sweetalert2@([\\d.]+)\\;version:\\1",
          "sweetalert2@([\\d.]+)/dist/sweetalert2(?:\\.all)(?:\\.min)\\.js",
          "limonte-sweetalert2/([\\d.]+)/sweetalert2(?:\\.all)(?:\\.min)\\.js"
        ],
        "website": "https://sweetalert2.github.io/"
      },
      "Swell": {
        "cats": [
          6
        ],
        "cookies": {
          "swell-session": ""
        },
        "html": [
          "<[^>]*swell\\.is",
          "<[^>]*swell\\.store",
          "<[^>]*schema\\.io"
        ],
        "icon": "Swell.svg",
        "js": {
          "swell.version": "^(.+)$\\;version:\\1"
        },
        "website": "https://www.swell.is/"
      },
      "Swiftype": {
        "cats": [
          29
        ],
        "description": "Swiftype provides search software for organisations, websites, and computer programs.",
        "icon": "swiftype.png",
        "js": {
          "Swiftype": ""
        },
        "scripts": "swiftype\\.com/embed\\.js$",
        "website": "http://swiftype.com"
      },
      "Swiper Slider": {
        "cats": [
          19
        ],
        "html": "<[^>]+=swiper-container",
        "icon": "swiper.svg",
        "js": {
          "Swiper": ""
        },
        "scripts": "swiper(?:\\.min)?\\.js",
        "website": "https://swiperjs.com"
      },
      "Sylius": {
        "cats": [
          6
        ],
        "description": "Sylius is an open source ecommerce framework based on Symfony full stack.",
        "dom": {
          "body.sylius_homepage": {
            "attributes": {
              "class": ""
            }
          },
          "div#sylius-cart-button": {
            "attributes": {
              "id": ""
            }
          }
        },
        "icon": "Sylius.svg",
        "implies": "Symfony",
        "oss": true,
        "scripts": [
          "syliusshop/script\\.js",
          "syliusgtmenhancedecommerceplugin"
        ],
        "website": "https://sylius.com"
      },
      "Symfony": {
        "cats": [
          18
        ],
        "cookies": {
          "sf_redirect": ""
        },
        "description": "Symfony is a PHP web application framework and a set of reusable PHP components/libraries.",
        "html": [
          "<div id=\"sfwdt[^\"]+\" class=\"[^\"]*sf-toolbar"
        ],
        "icon": "Symfony.svg",
        "implies": "PHP",
        "js": {
          "Sfjs": ""
        },
        "website": "https://symfony.com"
      },
      "Sympa": {
        "cats": [
          30
        ],
        "description": "Sympa is an open-source mailing list management (MLM) software.",
        "html": "<a href=\"https?://www\\.sympa\\.org\">\\s*Powered by Sympa\\s*</a>",
        "icon": "sympa.png",
        "implies": "Perl",
        "meta": {
          "generator": "^Sympa$"
        },
        "website": "https://www.sympa.org/"
      },
      "Synerise": {
        "cats": [
          10
        ],
        "icon": "Synerise.svg",
        "scripts": "snrcdn\\.net/sdk/(3\\.0)/synerise-javascript-sdk\\.min\\.js\\;version:\\1",
        "website": "https://synerise.com/"
      },
      "Synology DiskStation": {
        "cats": [
          48
        ],
        "description": "DiskStation provides a full-featured network attached storage.",
        "html": "<noscript><div class='syno-no-script'",
        "icon": "Synology DiskStation.svg",
        "meta": {
          "application-name": "Synology DiskStation",
          "description": "^DiskStation provides a full-featured network attached storage"
        },
        "scripts": "webapi/entry\\.cgi\\?api=SYNO\\.(?:Core|Filestation)\\.Desktop\\.",
        "website": "http://synology.com"
      },
      "SyntaxHighlighter": {
        "cats": [
          19
        ],
        "html": "<(?:script|link)[^>]*sh(?:Core|Brush|ThemeDefault)",
        "icon": "SyntaxHighlighter.png",
        "js": {
          "SyntaxHighlighter": ""
        },
        "website": "https://github.com/syntaxhighlighter"
      },
      "Syte": {
        "cats": [
          76,
          29
        ],
        "description": "Syte is a provider of visual AI technology that aims to improve retailers' site navigation, product discovery, and user experience by powering solutions that engage and convert shoppers.",
        "dom": "img[src*='cdn.syteapi.com']",
        "icon": "Syte.svg",
        "js": {
          "SyteApi.getBinImageBB": "",
          "SyteApp.Analytics": "",
          "SytePixel": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "cdn\\.syteapi\\.com/",
        "website": "https://www.syte.ai"
      },
      "T-Soft": {
        "cats": [
          6
        ],
        "html": "<a href=\"http://www\\.tsoft\\.com\\.tr\" target=\"_blank\" title=\"T-Soft E-ticaret Sistemleri\">",
        "icon": "Tsoft.png",
        "website": "https://www.tsoft.com.tr/"
      },
      "TN Express Web": {
        "cats": [
          1
        ],
        "cookies": {
          "TNEW": ""
        },
        "description": "Tessitura is an enterprise application to manage activities in ticketing, fundraising, customer relationship management, and marketing.",
        "icon": "tessitura.svg",
        "implies": "Tessitura",
        "website": "https://www.tessituranetwork.com"
      },
      "TNS Payments": {
        "cats": [
          41
        ],
        "description": "TNS Payments, is designed to deliver payment transaction information to banks, merchants, processors and other payment institutions.",
        "icon": "tnsi.png",
        "scripts": "secure\\.ap\\.tnspayments\\.com",
        "website": "https://tnsi.com/products/payments/"
      },
      "TWiki": {
        "cats": [
          8
        ],
        "cookies": {
          "TWIKISID": ""
        },
        "description": "TWiki is an open-source wiki and application platform.",
        "html": "<img [^>]*(?:title|alt)=\"This site is powered by the TWiki collaboration platform",
        "icon": "TWiki.png",
        "implies": "Perl",
        "scripts": "(?:TWikiJavascripts|twikilib(?:\\.min)?\\.js)",
        "website": "http://twiki.org"
      },
      "TYPO3 CMS": {
        "cats": [
          1
        ],
        "description": "TYPO3 is a free and open-source Web content management system written in PHP.",
        "html": [
          "<link[^>]+ href=\"/?typo3(?:conf|temp)/",
          "<img[^>]+ src=\"/?typo3(?:conf|temp)/",
          "<!--\n\tThis website is powered by TYPO3"
        ],
        "icon": "TYPO3.svg",
        "implies": "PHP",
        "meta": {
          "generator": "TYPO3\\s+(?:CMS\\s+)?(?:[\\d.]+)?(?:\\s+CMS)?"
        },
        "oss": true,
        "scripts": "^/?typo3(?:conf|temp)/",
        "url": "/typo3/",
        "website": "https://typo3.org/"
      },
      "TableBooker": {
        "cats": [
          5,
          72
        ],
        "description": "Tablebooker is an online reservation system for restaurants.",
        "dom": {
          "iframe": {
            "attributes": {
              "id": "tablebookerResFrame_"
            }
          }
        },
        "icon": "TableBooker.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "reservations\\.tablebooker\\.\\w+/",
        "website": "https://www.tablebooker.com"
      },
      "TableCheck": {
        "cats": [
          5,
          72
        ],
        "description": "TableCheck is a restaurant table booking widget.",
        "dom": {
          "form[action*='tablecheck']": {
            "attributes": {
              "action": "\\.tablecheck\\.\\w+/"
            }
          }
        },
        "icon": "TableCheck.svg",
        "scripts": "tc_widget\\.js",
        "website": "https://www.tablecheck.com"
      },
      "Taboola": {
        "cats": [
          36
        ],
        "description": "Taboola is a content discovery & native advertising platform for publishers and advertisers.",
        "dom": {
          "link[href*='.taboola.com']": {
            "attributes": {
              "href": ""
            }
          }
        },
        "icon": "Taboola.svg",
        "js": {
          "_taboola": "",
          "_taboolaNetworkMode": "",
          "taboola_view_id": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.taboola\\.com",
        "website": "https://www.taboola.com",
        "xhr": "\\.taboola\\.com"
      },
      "TagCommander": {
        "cats": [
          42
        ],
        "description": "TagCommander is an enterprise tag management system (TMS) that makes it easy to add and remove vendor tags from your Drupal-based websites.",
        "icon": "tagcommander.png",
        "js": {
          "tc_vars": ""
        },
        "scripts": "\\.tagcommander\\.com",
        "website": "https://www.commandersact.com/en/solutions/tagcommander/"
      },
      "Taggbox": {
        "cats": [
          5
        ],
        "description": "Taggbox is an UGC platform to collect, curate, manage rights, and publish user-generated content marketing campaigns across multiple channels.",
        "icon": "Taggbox.svg",
        "js": {
          "taggboxAjaxurl": ""
        },
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "scripts": [
          "\\.taggbox\\.com",
          "taggbox\\.com/app/js/embed\\.min\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1"
        ],
        "url": "\\.taggbox\\.com",
        "website": "https://taggbox.com/"
      },
      "Taiga": {
        "cats": [
          13
        ],
        "icon": "Taiga.png",
        "implies": [
          "Django",
          "AngularJS"
        ],
        "js": {
          "taigaConfig": ""
        },
        "website": "http://taiga.io"
      },
      "TakeDrop": {
        "cats": [
          6
        ],
        "description": "TakeDrop is an ecommerce platform.",
        "dom": "img[src*='main.takedropstorage.com']",
        "icon": "TakeDrop.png",
        "js": {
          "webpackJsonptakedrop-react": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://takedrop.pl"
      },
      "Tamago": {
        "cats": [
          5
        ],
        "html": "<link [^>]*href=\"http://tamago\\.temonalab\\.com",
        "icon": "Tamago.png",
        "website": "http://tamago.temonalab.com"
      },
      "Tawk.to": {
        "cats": [
          52
        ],
        "cookies": {
          "TawkConnectionTime": ""
        },
        "description": "Tawk.to is a free messaging app to monitor and chat with the visitors to a website, mobile app.",
        "icon": "TawkTo.png",
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "scripts": "//embed\\.tawk\\.to",
        "website": "http://tawk.to"
      },
      "Teachable": {
        "cats": [
          21
        ],
        "cookies": {
          "_gat_teachableTracker": "\\d+"
        },
        "description": "Teachable is a learning management system or LMS platform.",
        "icon": "Teachable.png",
        "js": {
          "isTeachable": "\\;confidence:50",
          "teachableIcons": "\\;confidence:50",
          "trackTeachableGAEvent": ""
        },
        "meta": {
          "asset_host": "\\.teachablecdn\\.com"
        },
        "pricing": [
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.teachablecdn\\.com",
        "website": "https://teachable.com"
      },
      "Teads": {
        "cats": [
          36
        ],
        "description": "Teads is a technology provider that sells ads on publisher websites.",
        "icon": "Teads.svg",
        "website": "https://www.teads.com",
        "xhr": "\\.teads\\.tv"
      },
      "Tealeaf": {
        "cats": [
          10
        ],
        "icon": "Tealeaf.png",
        "js": {
          "TeaLeaf": ""
        },
        "website": "http://www.tealeaf.com"
      },
      "Tealium": {
        "cats": [
          42
        ],
        "description": "Tealium provides a sales enterprise tag management system and marketing software.",
        "icon": "Tealium.png",
        "js": {
          "TEALIUMENABLED": ""
        },
        "scripts": [
          "^(?:https?:)?//tags\\.tiqcdn\\.com/",
          "/tealium/utag\\.js$"
        ],
        "website": "http://tealium.com"
      },
      "TeamCity": {
        "cats": [
          44
        ],
        "description": "TeamCity is a build management and continuous integration server from JetBrains.",
        "html": "<span class=\"versionTag\"><span class=\"vWord\">Version</span> ([\\d\\.]+)\\;version:\\1",
        "icon": "TeamCity.svg",
        "implies": [
          "Apache Tomcat",
          "Java",
          "jQuery",
          "Moment.js",
          "Prototype",
          "React",
          "Underscore.js"
        ],
        "meta": {
          "application-name": "TeamCity"
        },
        "website": "https://www.jetbrains.com/teamcity/"
      },
      "Telescope": {
        "cats": [
          1
        ],
        "icon": "Telescope.png",
        "implies": [
          "Meteor",
          "React"
        ],
        "js": {
          "Telescope": ""
        },
        "website": "http://telescopeapp.org"
      },
      "Tencent Analytics (腾讯分析)": {
        "cats": [
          10
        ],
        "icon": "tajs.png",
        "scripts": "tajs\\.qq\\.com/stats",
        "website": "https://ta.qq.com/"
      },
      "Tencent Waterproof Wall": {
        "cats": [
          9,
          16
        ],
        "icon": "TencentWaterproofWall.png",
        "scripts": [
          "/TCaptcha\\.js",
          "captcha\\.qq\\.com/.*"
        ],
        "website": "https://007.qq.com/"
      },
      "Tengine": {
        "cats": [
          22
        ],
        "description": "Tengine is a web server which is based on the Nginx HTTP server.",
        "headers": {
          "Server": "Tengine"
        },
        "icon": "Tengine.png",
        "website": "http://tengine.taobao.org"
      },
      "Termly": {
        "cats": [
          67
        ],
        "description": "Termly provides free website policy resources and web-based policy creation software.",
        "icon": "termly.svg",
        "scripts": "app\\.termly\\.io/embed\\.min\\.js",
        "website": "https://termly.io/"
      },
      "Tessitura": {
        "cats": [
          53
        ],
        "html": "<!--[^>]+Tessitura Version: (\\d*\\.\\d*\\.\\d*)?\\;version:\\1",
        "icon": "tessitura.svg",
        "implies": [
          "Microsoft ASP.NET",
          "IIS",
          "Windows Server"
        ],
        "website": "https://www.tessituranetwork.com"
      },
      "Texthelp": {
        "cats": [
          68
        ],
        "description": "TextHelp is a literacy, accessibility and dyslexia software developer for people with reading and writing difficulties.",
        "icon": "Texthelp.svg",
        "scripts": "browsealoud\\.com/.*/browsealoud\\.js",
        "website": "https://www.texthelp.com/en-gb/products/browsealoud/"
      },
      "Textpattern CMS": {
        "cats": [
          1
        ],
        "icon": "Textpattern CMS.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "meta": {
          "generator": "Textpattern"
        },
        "website": "http://textpattern.com"
      },
      "The Hut Group": {
        "cats": [
          6
        ],
        "icon": "TheHutGroup.png",
        "scripts": [
          "THEHUT-.*\\.js"
        ],
        "website": "https://www.thg.com/"
      },
      "Thefork": {
        "cats": [
          5,
          72
        ],
        "description": "Thefork is a restaurant booking, managing system.",
        "dom": {
          "iframe[data-src*='lafourchette']": {
            "attributes": {
              "data-src": "module\\.lafourchette.\\w+"
            }
          },
          "iframe[src*='lafourchette']": {
            "attributes": {
              "src": "module\\.lafourchette.\\w+"
            }
          }
        },
        "icon": "Thefork.svg",
        "pricing": [
          "mid",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.thefork.com"
      },
      "Thelia": {
        "cats": [
          1,
          6
        ],
        "html": "<(?:link|style|script)[^>]+/assets/frontOffice/",
        "icon": "Thelia.png",
        "implies": [
          "PHP",
          "Symfony"
        ],
        "website": "http://thelia.net"
      },
      "ThinkPHP": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "ThinkPHP"
        },
        "icon": "ThinkPHP.png",
        "implies": "PHP",
        "website": "http://www.thinkphp.cn"
      },
      "Thinkific": {
        "cats": [
          21
        ],
        "cookies": {
          "_thinkific_session": ""
        },
        "description": "Thinkific is a software platform that enables entrepreneurs to create, market, sell, and deliver their own online courses.",
        "icon": "Thinkific.svg",
        "js": {
          "Thinkific": "",
          "ThinkificAnalytics": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn(?:-themes)?\\.thinkific\\.com",
        "website": "https://www.thinkific.com"
      },
      "ThreatMetrix": {
        "cats": [
          16,
          83
        ],
        "description": "LexisNexis ThreatMetrix is an enterprise solution for online risk and fraud protection ('digital identity intelligence and digital authentication').",
        "icon": "ThreatMetrix.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.online-metrix\\.net",
        "website": "https://risk.lexisnexis.com/products/threatmetrix"
      },
      "ThriveCart": {
        "cats": [
          6
        ],
        "description": "ThriveCart is a sales cart solution that lets you create checkout pages for your online products and services.",
        "icon": "ThriveCart.svg",
        "js": {
          "ThriveCart": ""
        },
        "pricing": [
          "mid",
          "onetime"
        ],
        "saas": true,
        "scripts": "thrivecart\\.js",
        "website": "https://thrivecart.com"
      },
      "Ticimax": {
        "cats": [
          6
        ],
        "icon": "Ticimax.png",
        "scripts": [
          "cdn\\.ticimax\\.com/"
        ],
        "website": "https://www.ticimax.com"
      },
      "Tictail": {
        "cats": [
          6
        ],
        "html": "<link[^>]*tictail\\.com",
        "icon": "tictail.png",
        "website": "https://tictail.com"
      },
      "TiddlyWiki": {
        "cats": [
          1,
          2,
          4,
          8
        ],
        "description": "TiddlyWiki is an open-source notebook for capturing, organising and sharing complex information.",
        "html": "<[^>]*type=[^>]text\\/vnd\\.tiddlywiki",
        "icon": "TiddlyWiki.png",
        "js": {
          "tiddler": ""
        },
        "meta": {
          "application-name": "^TiddlyWiki$",
          "copyright": "^TiddlyWiki created by Jeremy Ruston",
          "generator": "^TiddlyWiki$",
          "tiddlywiki-version": "^(.+)$\\;version:\\1"
        },
        "website": "http://tiddlywiki.com"
      },
      "Tidio": {
        "cats": [
          32,
          5
        ],
        "description": "Tidio is a customer communication product. It provides multi-channel support so users can communicate with customers on the go. Live chat, messenger, or email are all supported.",
        "icon": "Tidio.svg",
        "js": {
          "tidioChatApi": ""
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "code\\.tidio\\.co",
        "website": "https://www.tidio.com"
      },
      "TikTok Pixel": {
        "cats": [
          10
        ],
        "description": "",
        "dom": {
          "script[data-hid='tiktok']": {
            "attributes": {
              "data-hid": ""
            }
          }
        },
        "icon": "TikTok.svg",
        "js": {
          "TiktokAnalyticsObject": ""
        },
        "website": "https://ads.tiktok.com",
        "xhr": "analytics\\.tiktok\\.com"
      },
      "Tiki Wiki CMS Groupware": {
        "cats": [
          1,
          2,
          8,
          11,
          13
        ],
        "description": "Tiki Wiki is a free and open-source wiki-based content management system and online office suite written primarily in PHP.",
        "icon": "Tiki Wiki CMS Groupware.png",
        "meta": {
          "generator": "^Tiki"
        },
        "scripts": "(?:/|_)tiki",
        "website": "http://tiki.org"
      },
      "Tilda": {
        "cats": [
          1
        ],
        "description": "Tilda is a web design tool.",
        "html": "<link[^>]* href=[^>]+tilda(?:cdn|\\.ws|-blocks)",
        "icon": "Tilda.svg",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "tilda(?:cdn|\\.ws|-blocks)",
        "website": "https://tilda.cc"
      },
      "Timeplot": {
        "cats": [
          25
        ],
        "icon": "Timeplot.png",
        "js": {
          "Timeplot": ""
        },
        "scripts": "timeplot.*\\.js",
        "website": "http://www.simile-widgets.org/timeplot/"
      },
      "Timify": {
        "cats": [
          72
        ],
        "description": "The Online Scheduling and Resource Management Software.",
        "icon": "Timify.svg",
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "scripts": "https://widget\\.timify\\.com/js/widget\\.js",
        "website": "https://www.timify.com/"
      },
      "TinyMCE": {
        "cats": [
          24
        ],
        "cpe": "cpe:/a:tiny:tinymce",
        "description": "TinyMCE is an online rich-text editor released as open-source software. TinyMCE is designed to integrate with JavaScript libraries, Vue.js, and AngularJS as well as content management systems such as Joomla!, and WordPress.",
        "icon": "TinyMCE.png",
        "js": {
          "tinyMCE.majorVersion": "([\\d.]+)\\;version:\\1"
        },
        "scripts": "/tiny_?mce(?:\\.min)?\\.js",
        "website": "http://tinymce.com"
      },
      "Titan": {
        "cats": [
          36
        ],
        "icon": "Titan.png",
        "js": {
          "titan": "",
          "titanEnabled": ""
        },
        "website": "http://titan360.com"
      },
      "TomatoCart": {
        "cats": [
          6
        ],
        "icon": "TomatoCart.png",
        "js": {
          "AjaxShoppingCart": ""
        },
        "meta": {
          "generator": "TomatoCart"
        },
        "website": "http://tomatocart.com"
      },
      "TornadoServer": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "TornadoServer(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "TornadoServer.png",
        "website": "http://tornadoweb.org"
      },
      "TotalCode": {
        "cats": [
          6
        ],
        "headers": {
          "X-Powered-By": "^TotalCode$"
        },
        "icon": "TotalCode.png",
        "website": "http://www.totalcode.com"
      },
      "Trac": {
        "cats": [
          13
        ],
        "html": [
          "<a id=\"tracpowered",
          "Powered by <a href=\"[^\"]*\"><strong>Trac(?:[ /]([\\d.]+))?\\;version:\\1"
        ],
        "icon": "Trac.png",
        "implies": "Python",
        "website": "http://trac.edgewall.org"
      },
      "TrackJs": {
        "cats": [
          10
        ],
        "description": "TrackJS is an error monitoring agent for production web sites and applications.",
        "icon": "TrackJs.svg",
        "js": {
          "TrackJs": ""
        },
        "website": "http://trackjs.com"
      },
      "Tradedoubler": {
        "cats": [
          71
        ],
        "description": "Tradedoubler is a global affiliate marketing network.",
        "dom": {
          "a[href*='clk.tradedoubler.com/click']": {
            "text": ""
          },
          "img[src*='impes.tradedoubler.com/imp']": {
            "text": ""
          }
        },
        "icon": "Tradedoubler.svg",
        "website": "https://www.tradedoubler.com/"
      },
      "Transifex": {
        "cats": [
          12
        ],
        "icon": "transifex.png",
        "js": {
          "Transifex.live.lib_version": "(.+)\\;version:\\1"
        },
        "website": "https://www.transifex.com"
      },
      "Translucide": {
        "cats": [
          1
        ],
        "icon": "translucide.svg",
        "implies": [
          "PHP",
          "jQuery"
        ],
        "scripts": "lucide\\.init(?:\\.min)?\\.js",
        "website": "http://www.translucide.net"
      },
      "Tray": {
        "cats": [
          6
        ],
        "icon": "tray.png",
        "scripts": "tcdn\\.com\\.br",
        "website": "https://www.tray.com.br"
      },
      "Trbo": {
        "cats": [
          74,
          76
        ],
        "cookies": {
          "trbo_session": "^(?:[\\d]+)$",
          "trbo_usr": "^(?:[\\d\\w]+)$"
        },
        "description": "Trbo is a personalisation, recommendations, A/B testing platform from Germany.",
        "icon": "Trbo.svg",
        "js": {
          "_trbo": "",
          "_trbo_start": "",
          "_trboq": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.trbo\\.com",
        "website": "https://www.trbo.com"
      },
      "Tripadviser.Widget": {
        "cats": [
          5
        ],
        "description": "Tripadvisor embed reviews widget.",
        "icon": "Tripadviser.Widget.svg",
        "scripts": "tripadvisor\\.[\\w]+/WidgetEmbed",
        "website": "https://www.tripadvisor.com/Widgets"
      },
      "TripleLift": {
        "cats": [
          36
        ],
        "description": "TripleLift is an advertising technology company providing native programmatic to buyers and sellers.",
        "dom": {
          "iframe[src*='.3lift.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "TripleLift.png",
        "pricing": [
          "low",
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://triplelift.com",
        "xhr": "\\.3lift\\.com"
      },
      "TruValidate": {
        "cats": [
          16,
          83
        ],
        "description": "TransUnion TruValidate (previously ReputationShield/IDVision from iovation) is an online risk and fraud detection platform.",
        "icon": "TruValidate.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "mpsnare\\.iesnare\\.com",
          "ci-mpsnare\\.iovation\\.com"
        ],
        "website": "https://www.transunion.com/solution/truvalidate"
      },
      "TrustArc": {
        "cats": [
          67
        ],
        "description": "TrustArc provides software and services to help corporations update their privacy management processes so they comply with government laws and best practices.",
        "icon": "TrustArc.svg",
        "scripts": "consent\\.trustarc\\.com",
        "website": "http://trustarc.com"
      },
      "TrustYou": {
        "cats": [
          5
        ],
        "description": "TrustYou is guest feedback and hotel reputation software company located in Germany.",
        "dom": {
          "iframe[src*='api.trustyou.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "TrustYou.svg",
        "saas": true,
        "website": "https://www.trustyou.com"
      },
      "Trustpilot": {
        "cats": [
          5
        ],
        "description": "Trustpilot is a Danish consumer review website which provide embed stand-alone applications in your website to show your most recent reviews, TrustScore, and star ratings.",
        "icon": "Trustpilot.svg",
        "js": {
          "Trustpilot": ""
        },
        "pricing": [
          "mid",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.trustpilot\\.com",
        "website": "https://business.trustpilot.com"
      },
      "Tumblr": {
        "cats": [
          11
        ],
        "description": "Tumblr is a microblogging and social networking website. The service allows users to post multimedia and other content to a short-form blog.",
        "headers": {
          "X-Tumblr-User": ""
        },
        "html": "<iframe src=\"[^>]+tumblr\\.com",
        "icon": "Tumblr.png",
        "url": "^https?://(?:www\\.)?[^/]+\\.tumblr\\.com/",
        "website": "http://www.tumblr.com"
      },
      "TurfJS": {
        "cats": [
          59
        ],
        "description": "Turf is a modular geospatial engine written in JavaScript.",
        "icon": "TurfJS.svg",
        "js": {
          "turf.feature": "",
          "turf.point": "",
          "turf.random": ""
        },
        "scripts": "(turf@[\\d.]+)?/?turf\\.min\\.js",
        "website": "https://turfjs.org/"
      },
      "TwicPics": {
        "cats": [
          31,
          59
        ],
        "description": "TwicPics offers on-demand responsive image generation.",
        "dom": {
          ".twic": {
            "attributes": {
              "data-src": ""
            }
          },
          "[data-twic-src]": {
            "attributes": {
              "data-twic-src": ""
            }
          },
          "data-twic-background": {
            "attributes": {
              "data-twic-background": ""
            }
          }
        },
        "headers": {
          "server": "^TwicPics/\\d+\\.\\d+\\.\\d+$"
        },
        "icon": "TwicPics.svg",
        "pricing": [
          "freemium",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": ".+\\.twic\\.pics",
        "website": "https://www.twicpics.com"
      },
      "Twilight CMS": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-CMS": "Twilight CMS"
        },
        "icon": "Twilight CMS.png",
        "website": "http://www.twilightcms.com"
      },
      "TwistPHP": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "TwistPHP"
        },
        "icon": "TwistPHP.png",
        "implies": "PHP",
        "website": "http://twistphp.com"
      },
      "TwistedWeb": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "TwistedWeb(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "TwistedWeb.png",
        "website": "http://twistedmatrix.com/trac/wiki/TwistedWeb"
      },
      "Twitter": {
        "cats": [
          5
        ],
        "description": "Twitter is a 'microblogging' system that allows you to send and receive short posts called tweets.",
        "icon": "Twitter.svg",
        "scripts": "//platform\\.twitter\\.com/widgets\\.js",
        "website": "http://twitter.com"
      },
      "Twitter Ads": {
        "cats": [
          36
        ],
        "description": "Twitter Ads is an advertising platform for Twitter 'microblogging' system.",
        "icon": "Twitter.svg",
        "js": {
          "twttr": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "static\\.ads-twitter\\.com/uwt\\.js",
        "website": "https://ads.twitter.com"
      },
      "Twitter Analytics": {
        "cats": [
          10
        ],
        "description": "Twitter Analytics is a built-in data-tracking platform that shows you insights specific to your Twitter account and activity.",
        "icon": "Twitter.svg",
        "pricing": [
          "freemium"
        ],
        "saas": true,
        "scripts": "analytics\\.twitter\\.com",
        "website": "https://analytics.twitter.com"
      },
      "Twitter Emoji (Twemoji)": {
        "cats": [
          17
        ],
        "description": "Twitter Emoji is a set of open-source emoticons and emojis for Twitter, TweetDeck, and also for Android and iOS versions of the application.",
        "js": {
          "twemoji": ""
        },
        "scripts": "twemoji(?:\\.min)?\\.js",
        "website": "https://twitter.github.io/twemoji/"
      },
      "Twitter Flight": {
        "cats": [
          12
        ],
        "icon": "Twitter Flight.png",
        "implies": "jQuery",
        "js": {
          "flight": ""
        },
        "website": "https://flightjs.github.io/"
      },
      "Twitter typeahead.js": {
        "cats": [
          59
        ],
        "icon": "Twitter typeahead.js.png",
        "implies": "jQuery",
        "js": {
          "typeahead": ""
        },
        "scripts": "(?:typeahead|bloodhound)\\.(?:jquery|bundle)?(?:\\.min)?\\.js",
        "website": "https://twitter.github.io/typeahead.js"
      },
      "TypeDoc": {
        "cats": [
          4
        ],
        "description": "TypeDoc is an API documentation generator for TypeScript projects.",
        "icon": "TypeDoc.png",
        "implies": "TypeScript",
        "oss": true,
        "website": "https://typedoc.org"
      },
      "TypePad": {
        "cats": [
          11
        ],
        "description": "Typepad is a blog hosting service.",
        "icon": "TypePad.png",
        "meta": {
          "generator": "typepad"
        },
        "url": "typepad\\.com",
        "website": "http://www.typepad.com"
      },
      "TypeScript": {
        "cats": [
          27
        ],
        "description": "TypeScript is an open-source language which builds on JavaScript  by adding static type definitions.",
        "icon": "TypeScript.svg",
        "oss": true,
        "website": "https://www.typescriptlang.org"
      },
      "Typecho": {
        "cats": [
          11
        ],
        "description": "Typecho is a PHP Blogging Platform.",
        "icon": "typecho.svg",
        "implies": "PHP",
        "js": {
          "TypechoComment": ""
        },
        "meta": {
          "generator": "Typecho( [\\d.]+)?\\;version:\\1"
        },
        "url": "/admin/login\\.php?referer=http%3A%2F%2F",
        "website": "http://typecho.org/"
      },
      "Typekit": {
        "cats": [
          17
        ],
        "description": "Typekit is an online service which offers a subscription library of fonts.",
        "html": "<link [^>]*href=\"[^\"]+use\\.typekit\\.(?:net|com)",
        "icon": "Typekit.png",
        "js": {
          "Typekit.config.js": "^(.+)$\\;version:\\1"
        },
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "scripts": "use\\.typekit\\.com",
        "website": "http://typekit.com"
      },
      "UIKit": {
        "cats": [
          66
        ],
        "description": "UIKit is the framework used for developing iOS applications.",
        "html": "<[^>]+class=\"[^\"]*(?:uk-container|uk-section)",
        "icon": "UIKit.png",
        "scripts": "uikit.*\\.js",
        "website": "https://getuikit.com"
      },
      "UMI.CMS": {
        "cats": [
          1
        ],
        "headers": {
          "X-Generated-By": "UMI\\.CMS"
        },
        "icon": "UMI.CMS.png",
        "implies": "PHP",
        "website": "https://www.umi-cms.ru"
      },
      "UNIX": {
        "cats": [
          28
        ],
        "description": "Unix is a family of multitasking, multiuser computer operating systems.",
        "headers": {
          "Server": "Unix"
        },
        "icon": "UNIX.png",
        "website": "http://unix.org"
      },
      "Ubercart": {
        "cats": [
          6
        ],
        "icon": "Ubercart.png",
        "implies": "Drupal",
        "scripts": "uc_cart/uc_cart_block\\.js",
        "website": "http://www.ubercart.org"
      },
      "Ubuntu": {
        "cats": [
          28
        ],
        "description": "Ubuntu is a free and open-source operating system on Linux for the enterprise server, desktop, cloud, and IoT.",
        "headers": {
          "Server": "Ubuntu",
          "X-Powered-By": "Ubuntu"
        },
        "icon": "Ubuntu.png",
        "website": "http://www.ubuntu.com/server"
      },
      "UltraCart": {
        "cats": [
          6
        ],
        "html": "<form [^>]*action=\"[^\"]*\\/cgi-bin\\/UCEditor\\?(?:[^\"]*&)?merchantId=[^\"]",
        "icon": "UltraCart.png",
        "js": {
          "ucCatalog": ""
        },
        "scripts": "cgi-bin\\/UCJavaScript\\?",
        "url": "/cgi-bin/UCEditor\\?",
        "website": "http://ultracart.com"
      },
      "Umbraco": {
        "cats": [
          1
        ],
        "headers": {
          "X-Umbraco-Version": "^(.+)$\\;version:\\1"
        },
        "html": "powered by <a href=[^>]+umbraco",
        "icon": "Umbraco.png",
        "implies": "Microsoft ASP.NET",
        "js": {
          "UC_IMAGE_SERVICE|ITEM_INFO_SERVICE": "",
          "UC_ITEM_INFO_SERVICE": "",
          "UC_SETTINGS": "",
          "Umbraco": ""
        },
        "meta": {
          "generator": "umbraco"
        },
        "url": "/umbraco/login\\.aspx(?:$|\\?)",
        "website": "http://umbraco.com"
      },
      "Unbounce": {
        "cats": [
          20,
          51
        ],
        "description": "Unbounce is a tool to build landing pages.",
        "headers": {
          "X-Unbounce-PageId": ""
        },
        "icon": "Unbounce.png",
        "scripts": "ubembed\\.com",
        "website": "http://unbounce.com"
      },
      "Underscore.js": {
        "cats": [
          59
        ],
        "description": "Underscore.js is a JavaScript library which provides utility functions for common programming tasks. It is comparable to features provided by Prototype.js and the Ruby language, but opts for a functional programming design instead of extending object prototypes.",
        "excludes": "Lodash",
        "icon": "Underscore.js.png",
        "js": {
          "_.VERSION": "^(.+)$\\;confidence:0\\;version:\\1",
          "_.restArguments": ""
        },
        "scripts": "underscore.*\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1",
        "website": "http://underscorejs.org"
      },
      "Uniconsent": {
        "cats": [
          67
        ],
        "icon": "Uniconsent.png",
        "scripts": "cmp\\.uniconsent\\.mgr\\.consensu\\.org/dfp\\.js",
        "website": "https://www.uniconsent.com/"
      },
      "Unpkg": {
        "cats": [
          31
        ],
        "description": "Unpkg is a content delivery network for everything on npm.",
        "dom": "link[href*='unpkg.com']",
        "icon": "Unpkg.png",
        "oss": true,
        "scripts": "unpkg\\.com/",
        "website": "https://unpkg.com"
      },
      "UpSellit": {
        "cats": [
          76
        ],
        "description": "UpSellit is a performance based lead and cart abandonment recovery solutions.",
        "icon": "UpSellit.png",
        "js": {
          "usi_analytics": "\\;confidence:25",
          "usi_app": "\\;confidence:25",
          "usi_commons": "\\;confidence:25",
          "usi_cookies": "\\;confidence:25"
        },
        "pricing": [
          "poa",
          "mid"
        ],
        "saas": true,
        "scripts": "www\\.upsellit\\.com/active",
        "website": "https://us.upsellit.com"
      },
      "Upvoty": {
        "cats": [
          47
        ],
        "dom": {
          "div.upvotyContainer": {
            "text": ""
          }
        },
        "icon": "upvoty.png",
        "implies": "PHP",
        "js": {
          "upvoty": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://upvoty.com"
      },
      "Usabilla": {
        "cats": [
          13
        ],
        "description": "Usabilla is a user feedback solution for collecting qualitative and quantitative user feedback across digital channels including websites, apps and emails.",
        "icon": "Usabilla.svg",
        "js": {
          "usabilla_live": ""
        },
        "website": "http://usabilla.com"
      },
      "UsableNet": {
        "cats": [
          68
        ],
        "description": "UsableNet is a technology for web accessibility and digital transformation, providing end-to-end services.",
        "html": "<iframe[ˆ>]*\\.usablenet\\.com/pt/",
        "icon": "UsableNet.png",
        "scripts": "\\.usablenet\\.com/pt/",
        "website": "https://usablenet.com/"
      },
      "Uscreen": {
        "cats": [
          1
        ],
        "description": "Uscreen is a CMS to monetize VOD and live content. They provide site hosting, video hosting, and a payment gateway for selling video based content.",
        "dom": {
          ".powered-by-uscreen": {
            "text": ""
          }
        },
        "icon": "Uscreen.svg",
        "js": {
          "analyticsHost": "stats\\.uscreen\\.io"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://uscreen.tv/"
      },
      "UserLike": {
        "cats": [
          52
        ],
        "description": "Userlike is a cloud-based live chat solution that can be integrated with existing websites.",
        "icon": "UserLike.svg",
        "js": {
          "userlike": "",
          "userlikeInit": ""
        },
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "scripts": [
          "userlike\\.min\\.js",
          "userlikelib\\.min\\.js",
          "//userlike-cdn-widgets\\.",
          "api\\.userlike\\.com"
        ],
        "website": "http://userlike.com"
      },
      "UserRules": {
        "cats": [
          13
        ],
        "icon": "UserRules.png",
        "js": {
          "_usrp": ""
        },
        "website": "http://www.userrules.com"
      },
      "UserVoice": {
        "cats": [
          13
        ],
        "description": "UserVoice is a management to collect and analyse feedback from customers.",
        "icon": "UserVoice.png",
        "js": {
          "UserVoice": ""
        },
        "website": "http://uservoice.com"
      },
      "UserWay": {
        "cats": [
          68
        ],
        "description": "UserWay is a web accessibility widget add-on for websites to help improve their compliance with accessibility standards.",
        "icon": "UserWay.png",
        "scripts": "cdn\\.userway\\.org/widget.*\\.js",
        "website": "https://userway.org/"
      },
      "Ushahidi": {
        "cats": [
          1,
          35
        ],
        "cookies": {
          "ushahidi": ""
        },
        "description": "Ushahidi is a tool that collects crowdsourced data and targeted survey responses from multiple data sources.",
        "icon": "Ushahidi.png",
        "implies": [
          "PHP",
          "MySQL",
          "OpenLayers"
        ],
        "js": {
          "Ushahidi": ""
        },
        "scripts": "/js/ushahidi\\.js$",
        "website": "http://www.ushahidi.com"
      },
      "VIVVO": {
        "cats": [
          1
        ],
        "cookies": {
          "VivvoSessionId": ""
        },
        "icon": "VIVVO.png",
        "js": {
          "vivvo": ""
        },
        "website": "http://vivvo.net"
      },
      "VP-ASP": {
        "cats": [
          6
        ],
        "html": "<a[^>]+>Powered By VP-ASP Shopping Cart</a>",
        "icon": "VP-ASP.png",
        "implies": "Microsoft ASP.NET",
        "scripts": "vs350\\.js",
        "website": "http://www.vpasp.com"
      },
      "VTEX": {
        "cats": [
          6
        ],
        "cookies": {
          "VtexFingerPrint": "",
          "VtexWorkspace": "",
          "vtex_session": ""
        },
        "description": "VTEX is an ecommerce software that manages multiple online stores.",
        "headers": {
          "Server": "^VTEX IO$",
          "powered": "vtex"
        },
        "icon": "VTEX.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "website": "https://vtex.com/"
      },
      "VWO": {
        "cats": [
          10,
          74
        ],
        "description": "VWO is a website testing and conversion optimisation platform.",
        "icon": "VWO.svg",
        "js": {
          "VWO": "",
          "__vwo": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "dev\\.visualwebsiteoptimizer\\.com/?([\\d.]+)\\;version:\\1"
        ],
        "website": "https://vwo.com/"
      },
      "VWO Engage": {
        "cats": [
          32
        ],
        "description": "VWO Engage is a part of the VWO Platform, which is a web-based push notification platform used by SaaS and B2B marketers, online content publishers, and ecommerce store owners.",
        "icon": "VWO.svg",
        "js": {
          "_pushcrewDebuggingQueue": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.pushcrew\\.\\w+",
        "website": "https://vwo.com/engage"
      },
      "Vaadin": {
        "cats": [
          18
        ],
        "icon": "Vaadin.svg",
        "implies": "Java",
        "js": {
          "vaadin": ""
        },
        "scripts": "vaadinBootstrap\\.js(?:\\?v=([\\d.]+))?\\;version:\\1",
        "website": "https://vaadin.com"
      },
      "ValueCommerce": {
        "cats": [
          71
        ],
        "description": "ValueCommerce is a pay-per-performance advertising affiliate marketing network.",
        "dom": {
          "a[href*='ap.valuecommerce.com']": {
            "attributes": {
              "href": ""
            }
          },
          "img[src*='ap.valuecommerce.com'],img[data-src*='ap.valuecommerce.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "ValueCommerce.png",
        "scripts": "\\.valuecommerce\\.com",
        "website": "https://www.valuecommerce.co.jp"
      },
      "Vanilla": {
        "cats": [
          2
        ],
        "description": "Vanilla is a both a cloud-based (SaaS) open-source community forum software.",
        "headers": {
          "X-Powered-By": "Vanilla"
        },
        "html": "<body id=\"(?:DiscussionsPage|vanilla)",
        "icon": "Vanilla.png",
        "implies": "PHP",
        "website": "http://vanillaforums.org"
      },
      "Varbase": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:vardot:varbase",
        "icon": "varbase.png",
        "implies": "Drupal",
        "meta": {
          "generator": "Varbase"
        },
        "website": "https://drupal.org/project/varbase"
      },
      "Varnish": {
        "cats": [
          23
        ],
        "description": "Varnish is a reverse caching proxy.",
        "headers": {
          "Via": "varnish(?: \\(Varnish/([\\d.]+)\\))?\\;version:\\1",
          "X-Varnish": "",
          "X-Varnish-Action": "",
          "X-Varnish-Age": "",
          "X-Varnish-Cache": "",
          "X-Varnish-Hostname": ""
        },
        "icon": "Varnish.svg",
        "website": "http://www.varnish-cache.org"
      },
      "Venmo": {
        "cats": [
          41
        ],
        "description": "Venmo is a mobile payment service owned by PayPal. Venmo account holders can transfer funds to others via a mobile phone app.",
        "html": "<[^>]+aria-labelledby=\"pi-venmo",
        "icon": "Venmo.svg",
        "website": "https://venmo.com"
      },
      "Veoxa": {
        "cats": [
          36
        ],
        "html": "<img [^>]*src=\"[^\"]+tracking\\.veoxa\\.com",
        "icon": "Veoxa.png",
        "js": {
          "VuVeoxaContent": ""
        },
        "scripts": "tracking\\.veoxa\\.com",
        "website": "http://veoxa.com"
      },
      "Vercel": {
        "cats": [
          22
        ],
        "headers": {
          "server": "^now$",
          "x-now-trace": "",
          "x-vercel-cache": "",
          "x-vercel-id": ""
        },
        "icon": "vercel.svg",
        "website": "https://vercel.com"
      },
      "VerifyPass": {
        "cats": [
          5,
          76
        ],
        "description": "VerifyPass is a company which provide secure identity proofing, authentication, and group affiliation verification for teachers and students.",
        "icon": "VerifyPass.png",
        "js": {
          "verifypass_api_instantiator": "",
          "verifypass_is_loaded": "",
          "verifypass_popup": ""
        },
        "pricing": [
          "payg",
          "recurring"
        ],
        "saas": true,
        "scripts": [
          "cdn\\.id\\.services",
          "cdn\\.verifypass\\.com"
        ],
        "website": "https://verifypass.com"
      },
      "Verizon Media": {
        "cats": [
          36
        ],
        "description": "Verizon Media is a tech and media company with global assets for advertisers, consumers and media companies.",
        "dom": {
          "img[src*='pixel.advertising.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "Verizon Media.svg",
        "scripts": "\\.advertising\\.com",
        "website": "https://www.verizonmedia.com"
      },
      "Vidazoo": {
        "cats": [
          36
        ],
        "description": "Vidazoo is a video content and yield management platform.",
        "icon": "Vidazoo.svg",
        "js": {
          "__vidazooPlayer__": "",
          "vidazoo": "",
          "vidazoo.version": "(.+)\\;version:\\1"
        },
        "pricing": [
          "high",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.vidazoo\\.com",
        "website": "https://www.vidazoo.com"
      },
      "VideoJS": {
        "cats": [
          14
        ],
        "description": "Video.js is a JavaScript and CSS library that makes it easier to work with and build on HTML5 video.",
        "html": "<div[^>]+class=\"video-js+\">",
        "icon": "VideoJS.svg",
        "js": {
          "VideoJS": "",
          "videojs": "",
          "videojs.VERSION": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "zencdn\\.net/c/video\\.js",
          "cdnjs\\.cloudflare\\.com\\/ajax\\/libs\\/video\\.js\\/([\\d\\.]+)\\/\\;version:\\1"
        ],
        "website": "http://videojs.com"
      },
      "Vigbo": {
        "cats": [
          1
        ],
        "cookies": {
          "_gphw_mode": ""
        },
        "html": "<link[^>]* href=[^>]+(?:\\.vigbo\\.com|\\.gophotoweb\\.com)",
        "icon": "vigbo.png",
        "scripts": "(?:\\.vigbo\\.com|\\.gophotoweb\\.com)",
        "website": "https://vigbo.com"
      },
      "Vignette": {
        "cats": [
          1
        ],
        "html": "<[^>]+=\"vgn-?ext",
        "icon": "Vignette.png",
        "website": "http://www.vignette.com"
      },
      "Vimeo": {
        "cats": [
          14
        ],
        "description": "Vimeo is a video hosting, sharing and services platform. Vimeo operation an ad-free basis by providing subscription plans.",
        "html": "(?:<(?:param|embed)[^>]+vimeo\\.com/moogaloop|<iframe[^>]player\\.vimeo\\.com)",
        "icon": "Vimeo.svg",
        "js": {
          "Vimeo.Player": "",
          "VimeoPlayer": ""
        },
        "website": "http://vimeo.com"
      },
      "Virgool": {
        "cats": [
          11
        ],
        "headers": {
          "X-Powered-By": "^Virgool$"
        },
        "icon": "Virgool.svg",
        "url": "^https?://(?:www\\.)?virgool\\.io",
        "website": "https://virgool.io"
      },
      "VirtueMart": {
        "cats": [
          6
        ],
        "html": "<div id=\"vmMainPage",
        "icon": "VirtueMart.png",
        "implies": "Joomla",
        "website": "http://virtuemart.net"
      },
      "Virtuoso": {
        "cats": [
          34
        ],
        "headers": {
          "Server": "Virtuoso/?([0-9.]+)?\\;version:\\1"
        },
        "meta": {
          "Copyright": "^Copyright &copy; \\d{4} OpenLink Software",
          "Keywords": "^OpenLink Virtuoso Sparql"
        },
        "url": "/sparql",
        "website": "https://virtuoso.openlinksw.com/"
      },
      "Visa": {
        "cats": [
          41
        ],
        "html": "<[^>]+aria-labelledby=\"pi-visa",
        "icon": "Visa.svg",
        "js": {
          "visaApi": "",
          "visaImage": "",
          "visaSrc": ""
        },
        "website": "https://www.visa.com"
      },
      "Visa Checkout": {
        "cats": [
          41
        ],
        "description": "Visa facilitates electronic funds transfers throughout the world, most commonly through Visa-branded credit cards, debit cards and prepaid cards.",
        "icon": "visa.png",
        "scripts": "secure\\.checkout\\.visa\\.com",
        "website": "https://checkout.visa.com"
      },
      "Visual Composer": {
        "cats": [
          51
        ],
        "description": "Visual Composer is an all-in-one web design tool for anyone who uses WordPress.",
        "icon": "visualcomposer.svg",
        "implies": [
          "WordPress",
          "PHP"
        ],
        "meta": {
          "generator": "Powered by Visual Composer Website Builder"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "website": "https://visualcomposer.com"
      },
      "Visualsoft": {
        "cats": [
          6
        ],
        "cookies": {
          "vscommerce": ""
        },
        "description": "Visualsoft is an ecommerce agency that delivers web design, development and marketing services to online retailers.",
        "icon": "Visualsoft.svg",
        "meta": {
          "vs_status_checker_version": "\\d+",
          "vsvatprices": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.visualsoft.co.uk/"
      },
      "Vizury": {
        "cats": [
          32
        ],
        "description": "Vizury is a commerce marketing platform.",
        "icon": "Vizury.png",
        "js": {
          "safariVizury": "",
          "vizury_data": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.vizury\\.com",
        "website": "https://www.vizury.com"
      },
      "Volusion": {
        "cats": [
          6
        ],
        "html": [
          "<link [^>]*href=\"[^\"]*/vspfiles/\\;version:1",
          "<body [^>]*data-vn-page-name\\;version:2"
        ],
        "icon": "Volusion.svg",
        "js": {
          "volusion": ""
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "/volusion\\.js(?:\\?([\\d.]*))?\\;version:\\1",
        "website": "https://www.volusion.com"
      },
      "Voog.com Website Builder": {
        "cats": [
          1,
          6
        ],
        "html": "<script [^>]*src=\"[^\"]*voog\\.com/tracker\\.js",
        "icon": "Voog.png",
        "scripts": "voog\\.com/tracker\\.js",
        "website": "https://www.voog.com/"
      },
      "Voracio": {
        "cats": [
          6
        ],
        "cookies": {
          "voracio_csrf_token": "",
          "voracio_sessionid": ""
        },
        "description": "Voracio is a cloud SaaS ecommerce platform powered by Microsoft .NET and built on the Microsoft Azure cloud framework.",
        "icon": "Voracio.svg",
        "js": {
          "voracio": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "website": "https://www.voracio.co.uk"
      },
      "Vue.js": {
        "cats": [
          12
        ],
        "description": "Vue.js is an open-source model–view–viewmodel JavaScript framework for building user interfaces and single-page applications.",
        "html": "<[^>]+\\sdata-v(?:ue)?-",
        "icon": "vue.svg",
        "js": {
          "Vue.version": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "vue[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "(?:/([\\d.]+))?/vue(?:\\.min)?\\.js\\;version:\\1"
        ],
        "website": "https://vuejs.org"
      },
      "VuePress": {
        "cats": [
          57
        ],
        "description": "VuePress is a static site generator with a Vue-powered theming system, and a default theme for writing technical documentation.",
        "icon": "VuePress.svg",
        "implies": "Vue.js",
        "meta": {
          "generator": "^VuePress(?: ([0-9.]+))?$\\;version:\\1"
        },
        "website": "https://vuepress.vuejs.org/"
      },
      "Vuetify": {
        "cats": [
          66
        ],
        "css": "\\.v-application \\.d-block ",
        "html": "<div data-app[^>]+class=\"v-application",
        "icon": "Vuetify.svg",
        "implies": "Vue.js",
        "website": "https://vuetifyjs.com"
      },
      "W3 Total Cache": {
        "cats": [
          23
        ],
        "description": "W3 Total Cache (W3TC) improves the SEO and increases website performance and reducing load times by leveraging features like content delivery network (CDN) integration and the latest best practices.",
        "headers": {
          "X-Powered-By": "W3 Total Cache(?:/([\\d.]+))?\\;version:\\1"
        },
        "html": "<!--[^>]+W3 Total Cache",
        "icon": "W3 Total Cache.png",
        "implies": "WordPress",
        "website": "http://www.w3-edge.com/wordpress-plugins/w3-total-cache"
      },
      "W3Counter": {
        "cats": [
          10
        ],
        "icon": "W3Counter.png",
        "scripts": "w3counter\\.com/tracker\\.js",
        "website": "http://www.w3counter.com"
      },
      "WEBDEV": {
        "cats": [
          20
        ],
        "description": "WEBDEV is a tool to develop internet and intranet sites and applications that support data and processes",
        "headers": {
          "WebDevSrc": ""
        },
        "html": "<!-- [a-zA-Z0-9_]+ [\\d/]+ [\\d:]+ WebDev \\d\\d ([\\d.]+) -->\\;version:\\1",
        "icon": "webdev.png",
        "meta": {
          "generator": "^WEBDEV$"
        },
        "website": "https://www.windev.com/webdev/index.html"
      },
      "WEBXPAY": {
        "cats": [
          6
        ],
        "html": "Powered by <a href=\"https://www\\.webxpay\\.com\">WEBXPAY<",
        "icon": "WEBXPAY.png",
        "js": {
          "WEBXPAY": ""
        },
        "website": "https://webxpay.com"
      },
      "WHMCS": {
        "cats": [
          6
        ],
        "description": "WHMCS is an automation platform that simplifies and automates all aspects of operating an online web hosting and domain registrar business.",
        "icon": "WHMCS.png",
        "js": {
          "WHMCS": ""
        },
        "oss": true,
        "pricing": [
          "low",
          "recurring"
        ],
        "website": "http://www.whmcs.com"
      },
      "WP Engine": {
        "cats": [
          62
        ],
        "description": "WP Engine is a website hosting provider.",
        "headers": {
          "X-Pass-Why": "",
          "X-Powered-By": "WP Engine",
          "X-WPE-Loopback-Upstream-Addr": "",
          "wpe-backend": ""
        },
        "icon": "wpengine.svg",
        "implies": "WordPress",
        "website": "https://wpengine.com"
      },
      "WP Rocket": {
        "cats": [
          23
        ],
        "description": "WP Rocket is a caching and performance optimisation plugin to improve the loading speed of WordPress websites.",
        "headers": {
          "X-Powered-By": "WP Rocket(?:/([\\d.]+))?\\;version:\\1",
          "X-Rocket-Nginx-Bypass": ""
        },
        "html": "<!--[^>]+WP Rocket",
        "icon": "WP Rocket.png",
        "implies": "WordPress",
        "website": "http://wp-rocket.me"
      },
      "WP-Statistics": {
        "cats": [
          10
        ],
        "html": [
          "<!-- Analytics by WP-Statistics v([\\d.]+) -\\;version:\\1"
        ],
        "icon": "WP-Statistics.png",
        "implies": "WordPress",
        "website": "https://wp-statistics.com"
      },
      "WPCacheOn": {
        "cats": [
          23
        ],
        "description": "WPCacheOn is a caching and performance optimisation plugin, which improves the loading speed of WordPress websites.",
        "headers": {
          "x-powered-by": "^Optimized by WPCacheOn"
        },
        "icon": "WPCacheOn.png",
        "implies": [
          "WordPress"
        ],
        "website": "https://wpcacheon.io"
      },
      "Wagtail": {
        "cats": [
          1
        ],
        "description": "Wagtail is a Django content management system (CMS) focused on flexibility and user experience.",
        "dom": {
          "[style*='images/']": {
            "attributes": {
              "style": "(?:\\.[a-z]+|/media)(?:/[\\w-]+)?/(?:original_images/[\\w-]+|images/[\\w-.]+\\.(?:(?:fill|max|min)-\\d+x\\d+(?:-c\\d+)?|(?:width|height|scale)-\\d+|original))\\."
            }
          },
          "img[src*='images/']": {
            "attributes": {
              "src": "(?:\\.[a-z]+|/media)(?:/[\\w-]+)?/(?:original_images/[\\w-]+|images/[\\w-.]+\\.(?:(?:fill|max|min)-\\d+x\\d+(?:-c\\d+)?|(?:width|height|scale)-\\d+|original))\\."
            }
          },
          "img[srcset*='images/'], source[srcset*='images/']": {
            "attributes": {
              "srcset": "(?:\\.[a-z]+|/media)(?:/[\\w-]+)?/(?:original_images/[\\w-]+|images/[\\w-.]+\\.(?:(?:fill|max|min)-\\d+x\\d+(?:-c\\d+)?|(?:width|height|scale)-\\d+|original))\\."
            }
          },
          "meta[content*='images/']": {
            "attributes": {
              "content": "(?:\\.[a-z]+|/media)(?:/[\\w-]+)?/(?:original_images/[\\w-]+|images/[\\w-.]+\\.(?:(?:fill|max|min)-\\d+x\\d+(?:-c\\d+)?|(?:width|height|scale)-\\d+|original))\\."
            }
          },
          "style, script": {
            "text": "(?:\\.[a-z]+|/media)(?:/[\\w-]+)?/(?:original_images/[\\w-]+|images/[\\w-.]+\\.(?:(?:fill|max|min)-\\d+x\\d+(?:-c\\d+)?|(?:width|height|scale)-\\d+|original))\\."
          },
          "video[poster*='images/']": {
            "attributes": {
              "poster": "(?:\\.[a-z]+|/media)(?:/[\\w-]+)?/(?:original_images/[\\w-]+|images/[\\w-.]+\\.(?:(?:fill|max|min)-\\d+x\\d+(?:-c\\d+)?|(?:width|height|scale)-\\d+|original))\\."
            }
          }
        },
        "icon": "Wagtail.svg",
        "implies": "Django",
        "website": "https://wagtail.io/"
      },
      "Wair": {
        "cats": [
          76,
          5
        ],
        "description": "Wair is the widget to personalised fit.",
        "icon": "Wair.png",
        "js": {
          "PredictV3.default.version": "([\\d.]+)\\;version:\\1",
          "predictWidget": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "getwair\\.com",
        "website": "https://getwair.com"
      },
      "Warp": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^Warp/(\\d+(?:\\.\\d+)+)?$\\;version:\\1"
        },
        "icon": "Warp.png",
        "implies": "Haskell",
        "website": "http://www.stackage.org/package/warp"
      },
      "Web2py": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "web2py"
        },
        "icon": "Web2py.png",
        "implies": [
          "Python",
          "jQuery"
        ],
        "meta": {
          "generator": "^Web2py"
        },
        "scripts": "web2py\\.js",
        "website": "http://web2py.com"
      },
      "WebAR": {
        "cats": [
          19
        ],
        "html": [
          "<model-viewer"
        ],
        "icon": "webAR.svg",
        "website": "https://modelviewer.dev/"
      },
      "WebAssembly": {
        "cats": [
          27
        ],
        "description": "WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable compilation target for programming languages, enabling deployment on the web for client and server applications.",
        "headers": {
          "Content-Type": "application/wasm"
        },
        "icon": "WebAssembly.svg",
        "oss": true,
        "website": "https://webassembly.org/"
      },
      "WebEngage": {
        "cats": [
          32,
          76
        ],
        "description": "WebEngage is a customer data platform and marketing automation suite.",
        "icon": "WebEngage.png",
        "js": {
          "webengage.__v": "([\\d.]+)\\;version:\\1"
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.webengage\\.co(?:m)?/",
        "website": "https://webengage.com"
      },
      "WebGUI": {
        "cats": [
          1
        ],
        "cookies": {
          "wgSession": ""
        },
        "icon": "WebGUI.png",
        "implies": "Perl",
        "meta": {
          "generator": "^WebGUI ([\\d.]+)\\;version:\\1"
        },
        "website": "http://www.webgui.org"
      },
      "WebMetric": {
        "cats": [
          10
        ],
        "cookies": {
          "_wmuid": ""
        },
        "icon": "WebMetric.svg",
        "js": {
          "_wmid": ""
        },
        "scripts": "cdn\\.webmetric\\.ir",
        "website": "https://webmetric.ir/"
      },
      "WebSite X5": {
        "cats": [
          20
        ],
        "description": "WebSite X5 is a tools to create and publish websites.",
        "icon": "WebSite X5.png",
        "meta": {
          "generator": "Incomedia WebSite X5 (\\w+ [\\d.]+)\\;version:\\1"
        },
        "website": "http://websitex5.com"
      },
      "Webflow": {
        "cats": [
          51
        ],
        "description": "Webflow is Software-as-a-Service (Saas) for website building and hosting.",
        "html": "<html[^>]+data-wf-site",
        "icon": "webflow.svg",
        "js": {
          "Webflow": ""
        },
        "meta": {
          "generator": "Webflow"
        },
        "website": "https://webflow.com"
      },
      "Webgains": {
        "cats": [
          71
        ],
        "description": "Webgains is an affiliate marketing network.",
        "icon": "Webgains.svg",
        "js": {
          "ITCLKQ": ""
        },
        "scripts": "analytics\\.webgains\\.io",
        "website": "https://www.webgains.com/"
      },
      "Webix": {
        "cats": [
          12
        ],
        "icon": "Webix.png",
        "js": {
          "webix": ""
        },
        "scripts": "\\bwebix\\.js",
        "website": "http://webix.com"
      },
      "Weblium": {
        "cats": [
          1,
          6
        ],
        "description": "Weblium let's you create a web site or online store without the need for a web developer or designer.",
        "dom": "link[href*='res2.weblium.site']",
        "icon": "Weblium.svg",
        "implies": [
          "Node.js",
          "OpenResty",
          "React"
        ],
        "pricing": [
          "low",
          "freemium",
          "recurring"
        ],
        "saas": true,
        "scripts": "res2\\.weblium\\.site/common/core\\.min\\.js",
        "url": "\\.weblium\\.site",
        "website": "https://weblium.com"
      },
      "WebsPlanet": {
        "cats": [
          1
        ],
        "icon": "WebsPlanet.png",
        "meta": {
          "generator": "WebsPlanet"
        },
        "website": "http://websplanet.com"
      },
      "Websale": {
        "cats": [
          6
        ],
        "cookies": {
          "websale_ac": ""
        },
        "icon": "Websale.png",
        "website": "http://websale.de"
      },
      "Website Creator": {
        "cats": [
          1
        ],
        "icon": "WebsiteCreator.png",
        "implies": [
          "PHP",
          "MySQL",
          "Vue.js"
        ],
        "meta": {
          "generator": "Website Creator by hosttech",
          "wsc_rendermode": ""
        },
        "website": "https://www.hosttech.ch/websitecreator"
      },
      "WebsiteBaker": {
        "cats": [
          1
        ],
        "icon": "WebsiteBaker.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "meta": {
          "generator": "WebsiteBaker"
        },
        "website": "http://websitebaker2.org/en/home.php"
      },
      "Websocket": {
        "cats": [
          19
        ],
        "html": [
          "<link[^>]+rel=[\"']web-socket[\"']",
          "<(?:link|a)[^>]+href=[\"']wss?://"
        ],
        "icon": "websocket.png",
        "website": "https://en.wikipedia.org/wiki/WebSocket"
      },
      "Webtrekk": {
        "cats": [
          10
        ],
        "icon": "Webtrekk.png",
        "js": {
          "WebtrekkV3": "",
          "webtrekk": "",
          "webtrekkConfig": "",
          "webtrekkHeatmapObjects": "",
          "webtrekkLinktrackObjects": "",
          "webtrekkUnloadObjects": "",
          "webtrekkV3": "",
          "wt_tt": "",
          "wt_ttv2": ""
        },
        "website": "http://www.webtrekk.com"
      },
      "Webtrends": {
        "cats": [
          10
        ],
        "html": "<img[^>]+id=\"DCSIMG\"[^>]+webtrends",
        "icon": "Webtrends.png",
        "js": {
          "WTOptimize": "",
          "WebTrends": ""
        },
        "website": "http://worldwide.webtrends.com"
      },
      "Webzi": {
        "cats": [
          1
        ],
        "icon": "Webzi.svg",
        "js": {
          "Webzi": ""
        },
        "meta": {
          "generator": "^Webzi"
        },
        "scripts": "cdn\\.6th\\.ir",
        "website": "https://webzi.ir"
      },
      "Weebly": {
        "cats": [
          1
        ],
        "description": "Weebly is a website and eCommerce service.",
        "icon": "Weebly.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "_W.configDomain": ""
        },
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "scripts": "cdn\\d+\\.editmysite\\.com",
        "website": "https://www.weebly.com"
      },
      "Weglot": {
        "cats": [
          19
        ],
        "headers": {
          "Weglot-Translated": ""
        },
        "icon": "Weglot.png",
        "scripts": [
          "cdn\\.weglot\\.com",
          "wp-content/plugins/weglot"
        ],
        "website": "https://www.weglot.com"
      },
      "Welcart": {
        "cats": [
          6
        ],
        "cookies": {
          "usces_cookie": ""
        },
        "cpe": "cpe:/a:welcart:welcart",
        "html": [
          "<link[^>]+?href=\"[^\"]+usces_default(?:\\.min)?\\.css",
          "<!-- Welcart version : v([\\d.]+)\\;version:\\1"
        ],
        "icon": "welcart.png",
        "implies": [
          "PHP",
          "WordPress"
        ],
        "scripts": "uscesL10n",
        "website": "https://www.welcart.com"
      },
      "Whatfix": {
        "cats": [
          19
        ],
        "description": "Whatfix is a SaaS based platform which provides in-app guidance and performance support for web applications and software products.",
        "icon": "Whatfix.png",
        "js": {
          "_wfx_add_logger": "",
          "_wfx_settings": "",
          "wfx_is_playing__": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "whatfix\\.com",
        "website": "https://whatfix.com"
      },
      "Whooshkaa": {
        "cats": [
          5
        ],
        "html": "<iframe src=\"[^>]+whooshkaa\\.com",
        "icon": "Whooshkaa.svg",
        "website": "https://www.whooshkaa.com"
      },
      "Wiki.js": {
        "cats": [
          4
        ],
        "description": "Wiki.js is a wiki engine running on Node.js and written in JavaScript.",
        "icon": "Wiki.js.png",
        "implies": "Node.js",
        "js": {
          "WIKI.$_apolloInitData": "",
          "WIKI.$apolloProvider": ""
        },
        "oss": true,
        "website": "https://js.wiki"
      },
      "Wikinggruppen": {
        "cats": [
          6
        ],
        "html": [
          "<!-- WIKINGGRUPPEN"
        ],
        "icon": "wikinggruppen.png",
        "website": "https://wikinggruppen.se/"
      },
      "WikkaWiki": {
        "cats": [
          8
        ],
        "description": "WikkaWiki is an open-source wiki application written in PHP.",
        "html": "Powered by <a href=\"[^>]+WikkaWiki",
        "icon": "WikkaWiki.png",
        "meta": {
          "generator": "WikkaWiki"
        },
        "website": "http://wikkawiki.org"
      },
      "Windows CE": {
        "cats": [
          28
        ],
        "description": "Windows CE is an operating system designed for small footprint devices or embedded systems.",
        "headers": {
          "Server": "\\bWinCE\\b"
        },
        "icon": "Microsoft.png",
        "website": "http://microsoft.com"
      },
      "Windows Server": {
        "cats": [
          28
        ],
        "description": "Windows Server is a brand name for a group of server operating systems.",
        "headers": {
          "Server": "Win32|Win64"
        },
        "icon": "WindowsServer.png",
        "website": "http://microsoft.com/windowsserver"
      },
      "Wink": {
        "cats": [
          26,
          12
        ],
        "description": "Wink Toolkit is a JavaScript toolkit used to build mobile web apps.",
        "icon": "Wink.png",
        "js": {
          "wink.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "(?:_base/js/base|wink).*\\.js",
        "website": "http://winktoolkit.org"
      },
      "Winstone Servlet Container": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Winstone Servlet (?:Container|Engine) v?([\\d.]+)?\\;version:\\1",
          "X-Powered-By": "Winstone(?:\\/([\\d.]+))?\\;version:\\1"
        },
        "website": "http://winstone.sourceforge.net"
      },
      "Wistia": {
        "cats": [
          14
        ],
        "description": "Wistia is designed exclusively to serve companies using video on their websites for marketing, support, and sales.",
        "icon": "Wistia.svg",
        "js": {
          "Wistia": "",
          "wistiaEmbeds": "",
          "wistiaUtils": ""
        },
        "pricing": [
          "freemium",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.wistia\\.com",
        "website": "https://wistia.com"
      },
      "Wix": {
        "cats": [
          1,
          11
        ],
        "cookies": {
          "Domain": "\\.wix\\.com"
        },
        "description": "Wix provides cloud-based web development services, allowing users to create HTML5 websites and mobile sites.",
        "headers": {
          "X-Wix-Renderer-Server": "",
          "X-Wix-Request-Id": "",
          "X-Wix-Server-Artifact-Id": ""
        },
        "icon": "Wix.svg",
        "implies": "React",
        "js": {
          "wixBiSession": "",
          "wixPerformanceMeasurements": ""
        },
        "meta": {
          "generator": "Wix\\.com Website Builder"
        },
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "scripts": "static\\.parastorage\\.com",
        "website": "https://www.wix.com"
      },
      "Wix eCommerce": {
        "cats": [
          6
        ],
        "dom": {
          "#wix-viewer-model": {
            "text": "wixstores"
          }
        },
        "icon": "Wix.svg",
        "implies": "Wix",
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://www.wix.com/freesitebuilder/tae-store"
      },
      "Wolf CMS": {
        "cats": [
          1
        ],
        "html": "(?:<a href=\"[^>]+wolfcms\\.org[^>]+>Wolf CMS(?:</a>)? inside|Thank you for using <a[^>]+>Wolf CMS)",
        "icon": "Wolf CMS.png",
        "implies": "PHP",
        "website": "http://www.wolfcms.org"
      },
      "Woltlab Community Framework": {
        "cats": [
          1
        ],
        "icon": "Woltlab Community Framework.png",
        "implies": "PHP",
        "scripts": "WCF\\..*\\.js",
        "website": "http://www.woltlab.com"
      },
      "WooCommerce": {
        "cats": [
          6
        ],
        "description": "WooCommerce is an open-source ecommerce plugin for WordPress.",
        "dom": ".woocommerce, .woocommerce-no-js, link[rel*='woocommerce']",
        "icon": "WooCommerce.svg",
        "implies": "WordPress",
        "js": {
          "woocommerce_params": ""
        },
        "meta": {
          "generator": "WooCommerce ([\\d.]+)\\;version:\\1"
        },
        "oss": true,
        "scripts": [
          "woocommerce",
          "/woocommerce(?:\\.min)?\\.js(?:\\?ver=([0-9.]+))?\\;version:\\1"
        ],
        "website": "https://woocommerce.com"
      },
      "Woopra": {
        "cats": [
          10
        ],
        "icon": "Woopra.png",
        "scripts": "static\\.woopra\\.com",
        "website": "http://www.woopra.com"
      },
      "WordPress": {
        "cats": [
          1,
          11
        ],
        "cpe": "cpe:/a:wordpress:wordpress",
        "description": "WordPress is a free and open-source content management system written in PHP and paired with a MySQL or MariaDB database. Features include a plugin architecture and a template system.",
        "headers": {
          "X-Pingback": "/xmlrpc\\.php$",
          "link": "rel=\"https://api\\.w\\.org/\""
        },
        "html": [
          "<link rel=[\"']stylesheet[\"'] [^>]+/wp-(?:content|includes)/",
          "<link[^>]+s\\d+\\.wp\\.com"
        ],
        "icon": "WordPress.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "wp_username": ""
        },
        "meta": {
          "generator": "^WordPress ?([\\d.]+)?\\;version:\\1",
          "shareaholic:wp_version": ""
        },
        "pricing": [
          "low",
          "recurring",
          "freemium"
        ],
        "saas": true,
        "scripts": [
          "/wp-(?:content|includes)/",
          "wp-embed\\.min\\.js"
        ],
        "website": "https://wordpress.org"
      },
      "WordPress Super Cache": {
        "cats": [
          23
        ],
        "description": "WordPpress Super Cache is a static caching plugin for WordPress.",
        "headers": {
          "WP-Super-Cache": ""
        },
        "html": "<!--[^>]+WP-Super-Cache",
        "icon": "wp_super_cache.png",
        "implies": "WordPress",
        "website": "http://z9.io/wp-super-cache/"
      },
      "WordPress VIP": {
        "cats": [
          62
        ],
        "headers": {
          "x-powered-by": "^WordPress\\.com VIP"
        },
        "icon": "wpvip.svg",
        "implies": [
          "WordPress",
          "Automattic"
        ],
        "website": "https://wpvip.com"
      },
      "Wunderkind": {
        "cats": [
          32
        ],
        "description": "Wunderkind (Formerly BounceX) is a software for behavioural marketing technologies, created to de-anonymise site visitors, analyse their digital behaviour and create relevant digital experiences regardless of channel or device.",
        "icon": "Wunderkind.svg",
        "js": {
          "bouncex": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.bounceexchange\\.com",
        "website": "https://www.wunderkind.co"
      },
      "X-Cart": {
        "cats": [
          6
        ],
        "cookies": {
          "xid": "[a-z\\d]{32}(?:;|$)"
        },
        "html": [
          "Powered by X-Cart(?: (\\d+))? <a[^>]+href=\"http://www\\.x-cart\\.com/\"[^>]*>\\;version:\\1",
          "<a[^>]+href=\"[^\"]*(?:\\?|&)xcart_form_id=[a-z\\d]{32}(?:&|$)"
        ],
        "icon": "X-Cart.png",
        "implies": "PHP",
        "js": {
          "xcart_web_dir": "",
          "xliteConfig": ""
        },
        "meta": {
          "generator": "X-Cart(?: (\\d+))?\\;version:\\1"
        },
        "scripts": "/skin/common_files/modules/Product_Options/func\\.js",
        "website": "http://x-cart.com"
      },
      "XAMPP": {
        "cats": [
          22
        ],
        "html": "<title>XAMPP(?: Version ([\\d\\.]+))?</title>\\;version:\\1",
        "icon": "XAMPP.png",
        "implies": [
          "Apache",
          "MySQL",
          "PHP",
          "Perl"
        ],
        "meta": {
          "author": "Kai Oswald Seidler\\;confidence:10"
        },
        "website": "http://www.apachefriends.org/en/xampp.html"
      },
      "XMB": {
        "cats": [
          2
        ],
        "html": "<!-- Powered by XMB",
        "icon": "XMB.png",
        "website": "http://www.xmbforum.com"
      },
      "XOOPS": {
        "cats": [
          1
        ],
        "icon": "XOOPS.png",
        "implies": "PHP",
        "js": {
          "xoops": ""
        },
        "meta": {
          "generator": "XOOPS"
        },
        "website": "http://xoops.org"
      },
      "XRegExp": {
        "cats": [
          59
        ],
        "icon": "XRegExp.png",
        "js": {
          "XRegExp.version": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "xregexp[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "/([\\d.]+)/xregexp(?:\\.min)?\\.js\\;version:\\1",
          "xregexp.*\\.js"
        ],
        "website": "http://xregexp.com"
      },
      "XWiki": {
        "cats": [
          8
        ],
        "description": "XWiki is a free wiki software platform written in Java.",
        "excludes": "MediaWiki",
        "html": [
          "<html[^>]data-xwiki-[^>]>"
        ],
        "icon": "xwiki.png",
        "implies": "Java\\;confidence:99",
        "meta": {
          "wiki": "xwiki"
        },
        "website": "http://www.xwiki.org"
      },
      "Xajax": {
        "cats": [
          59
        ],
        "icon": "Xajax.png",
        "scripts": "xajax_core.*\\.js",
        "website": "http://xajax-project.org"
      },
      "Xanario": {
        "cats": [
          6
        ],
        "icon": "Xanario.png",
        "meta": {
          "generator": "xanario shopsoftware"
        },
        "website": "http://xanario.de"
      },
      "XenForo": {
        "cats": [
          2
        ],
        "cookies": {
          "xf_csrf": "",
          "xf_session": ""
        },
        "description": "XenForo is a PHP-based forum hosting program for communities that is designed to be deployed on a remote web server.",
        "html": [
          "(?:jQuery\\.extend\\(true, XenForo|<a[^>]+>Forum software by XenForo™|<!--XF:branding|<html[^>]+id=\"XenForo\")",
          "<html id=\"XF\" "
        ],
        "icon": "XenForo.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "XF.GuestUsername": ""
        },
        "website": "http://xenforo.com"
      },
      "Xeora": {
        "cats": [
          18
        ],
        "headers": {
          "Server": "XeoraEngine",
          "X-Powered-By": "XeoraCube"
        },
        "html": "<input type=\"hidden\" name=\"_sys_bind_\\d+\" id=\"_sys_bind_\\d+\" />",
        "icon": "xeora.png",
        "implies": "Microsoft ASP.NET",
        "scripts": "/_bi_sps_v.+\\.js",
        "website": "http://www.xeora.org"
      },
      "Xitami": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Xitami(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Xitami.png",
        "website": "http://xitami.com"
      },
      "Xonic": {
        "cats": [
          6
        ],
        "html": [
          "Powered by <a href=\"http://www\\.xonic-solutions\\.de/index\\.php\" target=\"_blank\">xonic-solutions Shopsoftware</a>"
        ],
        "icon": "xonic.png",
        "meta": {
          "keywords": "xonic-solutions"
        },
        "scripts": "core/jslib/jquery\\.xonic\\.js\\.php",
        "website": "http://www.xonic-solutions.de"
      },
      "XpressEngine": {
        "cats": [
          1
        ],
        "icon": "XpressEngine.png",
        "meta": {
          "generator": "XpressEngine"
        },
        "website": "http://www.xpressengine.com/"
      },
      "Xtremepush": {
        "cats": [
          32
        ],
        "description": "Xtremepush is a customer engagement, personalisation and data platform. It's purpose-built for multichannel and mobile marketing.",
        "icon": "Xtremepush.svg",
        "js": {
          "xtremepush": ""
        },
        "website": "https://xtremepush.com"
      },
      "YUI": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:yahoo:yui",
        "icon": "YUI.png",
        "js": {
          "YAHOO.VERSION": "^(.+)$\\;version:\\1",
          "YUI.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "(?:/yui/|yui\\.yahooapis\\.com)",
        "website": "http://yuilibrary.com"
      },
      "YUI Doc": {
        "cats": [
          4
        ],
        "description": "UIDoc is a Node.js application used at build time to generate API documentation.",
        "html": "(?:<html[^>]* yuilibrary\\.com/rdf/[\\d.]+/yui\\.rdf|<body[^>]+class=\"yui3-skin-sam)",
        "icon": "yahoo.png",
        "website": "http://developer.yahoo.com/yui/yuidoc"
      },
      "YaBB": {
        "cats": [
          2
        ],
        "html": "Powered by <a href=\"[^>]+yabbforum",
        "icon": "YaBB.png",
        "website": "http://www.yabbforum.com"
      },
      "Yahoo Advertising": {
        "cats": [
          36
        ],
        "html": [
          "<iframe[^>]+adserver\\.yahoo\\.com",
          "<img[^>]+clicks\\.beap\\.bc\\.yahoo\\.com"
        ],
        "icon": "yahoo.png",
        "js": {
          "adxinserthtml": ""
        },
        "scripts": "adinterax\\.com",
        "website": "http://advertising.yahoo.com"
      },
      "Yahoo! Ecommerce": {
        "cats": [
          6
        ],
        "headers": {
          "X-XRDS-Location": "/ystore/"
        },
        "html": "<link[^>]+store\\.yahoo\\.net",
        "icon": "yahoo.png",
        "js": {
          "YStore": ""
        },
        "website": "http://smallbusiness.yahoo.com/ecommerce"
      },
      "Yahoo! Tag Manager": {
        "cats": [
          42
        ],
        "html": "<!-- (?:End )?Yahoo! Tag Manager -->",
        "icon": "yahoo.png",
        "scripts": "b\\.yjtag\\.jp/iframe",
        "website": "https://tagmanager.yahoo.co.jp/"
      },
      "Yahoo! Web Analytics": {
        "cats": [
          10
        ],
        "icon": "yahoo.png",
        "js": {
          "YWA": ""
        },
        "scripts": "d\\.yimg\\.com/mi/ywa\\.js",
        "website": "http://web.analytics.yahoo.com"
      },
      "Yandex.Direct": {
        "cats": [
          36
        ],
        "description": "Yandex Direct is the platform designed for sponsored ad management.",
        "html": "<yatag class=\"ya-partner__ads\">",
        "icon": "Yandex.Direct.png",
        "js": {
          "yandex_ad_format": "",
          "yandex_partner_id": ""
        },
        "scripts": "https?://an\\.yandex\\.ru/",
        "website": "http://partner.yandex.com"
      },
      "Yandex.Messenger": {
        "cats": [
          5,
          52
        ],
        "description": "Yandex.Messenger is an instant messaging application.",
        "icon": "Yandex.Messenger.svg",
        "js": {
          "yandexChatWidget": ""
        },
        "pricing": [
          "payg"
        ],
        "scripts": "chat\\.s3\\.yandex\\.net/widget\\.js",
        "website": "https://dialogs.yandex.ru"
      },
      "Yandex.Metrika": {
        "cats": [
          10
        ],
        "description": "Yandex.Metrica is a free web analytics service that tracks and reports website traffic.",
        "icon": "Yandex.Metrika.png",
        "js": {
          "yandex_metrika": ""
        },
        "scripts": [
          "mc\\.yandex\\.ru/metrika/(?:tag|watch)\\.js",
          "cdn\\.jsdelivr\\.net/npm/yandex\\-metrica\\-watch/watch\\.js"
        ],
        "website": "http://metrika.yandex.com"
      },
      "Yaws": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "Yaws(?: ([\\d.]+))?\\;version:\\1"
        },
        "icon": "Yaws.png",
        "website": "http://yaws.hyber.org"
      },
      "Yelp Reservations": {
        "cats": [
          5,
          72
        ],
        "description": "Yelp Reservations is a cloud-based restaurant management system.",
        "dom": {
          "iframe[src*='yelp']": {
            "attributes": {
              "src": "yelp(?:.com/reservations|reservations\\.com)"
            }
          }
        },
        "icon": "Yelp.svg",
        "website": "http://yelp.com"
      },
      "Yelp Review Badge": {
        "cats": [
          5
        ],
        "description": "Yelp Review Badges showcase business reviews from Yelp on websites.",
        "dom": {
          "img[src*='dyn.yelpcdn.com']": {
            "attributes": {
              "src": ""
            }
          }
        },
        "icon": "Yelp.svg",
        "scripts": "yelp\\.com/biz_badge_js",
        "website": "http://yelp.com"
      },
      "Yepcomm": {
        "cats": [
          6
        ],
        "icon": "yepcomm.png",
        "meta": {
          "author": "Yepcomm Tecnologia",
          "copyright": "Yepcomm Tecnologia"
        },
        "website": "https://www.yepcomm.com.br"
      },
      "Yieldify": {
        "cats": [
          76
        ],
        "description": "Yieldify is a customer journey optimisation platform that brings personalisation to the full customer journey.",
        "icon": "Yieldify.svg",
        "js": {
          "_yieldify": ""
        },
        "pricing": [
          "poa",
          "payg"
        ],
        "saas": true,
        "scripts": "\\.yieldify\\.com",
        "website": "https://www.yieldify.com"
      },
      "Yieldlab": {
        "cats": [
          36
        ],
        "icon": "Yieldlab.png",
        "scripts": "^https?://(?:[^/]+\\.)?yieldlab\\.net/",
        "website": "http://yieldlab.de"
      },
      "Yii": {
        "cats": [
          18
        ],
        "cookies": {
          "YII_CSRF_TOKEN": ""
        },
        "description": "Yii is an open-source, object-oriented, component-based MVC PHP web application framework.",
        "html": [
          "Powered by <a href=\"http://www\\.yiiframework\\.com/\" rel=\"external\">Yii Framework</a>",
          "<input type=\"hidden\" value=\"[a-zA-Z0-9]{40}\" name=\"YII_CSRF_TOKEN\" \\/>",
          "<!\\[CDATA\\[YII-BLOCK-(?:HEAD|BODY-BEGIN|BODY-END)\\]"
        ],
        "icon": "Yii.png",
        "implies": "PHP",
        "scripts": [
          "/assets/[a-zA-Z0-9]{8}\\/yii\\.js$",
          "/yii\\.(?:validation|activeForm)\\.js"
        ],
        "website": "https://www.yiiframework.com"
      },
      "Yoast SEO": {
        "cats": [
          54
        ],
        "description": "Yoast SEO is a search engine optimisation plugin for WordPress and other platforms.",
        "dom": {
          "script.yoast-schema-graph": {
            "attributes": {
              "class": ""
            }
          }
        },
        "html": "<!-- This site is optimized with the Yoast (?:WordPress )?SEO plugin v([\\d.]+) -\\;version:\\1",
        "icon": "Yoast SEO.png",
        "website": "https://yoast.com"
      },
      "Yottaa": {
        "cats": [
          42,
          74
        ],
        "description": "Yottaa is an ecommerce optimisation platform that helps with conversions, performance and security.",
        "icon": "Yottaa.svg",
        "meta": {
          "X-Yottaa-Metrics": "",
          "X-Yottaa-Optimizations": ""
        },
        "scripts": "cdn\\.yottaa\\.\\w+/",
        "website": "https://www.yottaa.com"
      },
      "YouTrack": {
        "cats": [
          13
        ],
        "description": "YouTrack is a browser-based bug tracker, issue tracking system and project management software.",
        "html": [
          "no-title=\"YouTrack\">",
          "data-reactid=\"[^\"]+\">youTrack ([0-9.]+)<\\;version:\\1",
          "type=\"application/opensearchdescription\\+xml\" title=\"YouTrack\"/>"
        ],
        "icon": "YouTrack.png",
        "website": "http://www.jetbrains.com/youtrack/"
      },
      "YouTube": {
        "cats": [
          14
        ],
        "description": "YouTube is a video sharing service where users can create their own profile, upload videos, watch, like and comment on other videos.",
        "html": "<(?:param|embed|iframe)[^>]+youtube(?:-nocookie)?\\.com/(?:v|embed)",
        "icon": "YouTube.png",
        "website": "http://www.youtube.com"
      },
      "ZK": {
        "cats": [
          18
        ],
        "html": "<!-- ZK [.\\d\\s]+-->",
        "icon": "ZK.png",
        "implies": "Java",
        "scripts": "zkau/",
        "website": "http://zkoss.org"
      },
      "ZURB Foundation": {
        "cats": [
          66
        ],
        "description": "Zurb Foundation is used to prototype in the browser. Allows rapid creation of websites or applications while leveraging mobile and responsive technology. The front end framework is the collection of HTML, CSS, and Javascript containing design patterns.",
        "html": [
          "<link[^>]+foundation[^>\"]+css",
          "<div [^>]*class=\"[^\"]*(?:small|medium|large)-\\d{1,2} columns"
        ],
        "icon": "ZURB Foundation.png",
        "js": {
          "Foundation.version": "([\\d.]+)\\;version:\\1"
        },
        "website": "http://foundation.zurb.com"
      },
      "Zabbix": {
        "cats": [
          19
        ],
        "html": "<body[^>]+zbxCallPostScripts",
        "icon": "Zabbix.png",
        "implies": "PHP",
        "js": {
          "zbxCallPostScripts": ""
        },
        "meta": {
          "Author": "ZABBIX SIA\\;confidence:70"
        },
        "url": "\\/zabbix\\/\\;confidence:30",
        "website": "http://zabbix.com"
      },
      "Zanox": {
        "cats": [
          36
        ],
        "html": "<img [^>]*src=\"[^\"]+ad\\.zanox\\.com",
        "icon": "Zanox.png",
        "js": {
          "zanox": ""
        },
        "scripts": "zanox\\.com/scripts/zanox\\.js$",
        "website": "http://zanox.com"
      },
      "Zen Cart": {
        "cats": [
          6
        ],
        "icon": "Zen Cart.png",
        "meta": {
          "generator": "Zen Cart"
        },
        "website": "http://www.zen-cart.com"
      },
      "Zend": {
        "cats": [
          22
        ],
        "cookies": {
          "ZENDSERVERSESSID": ""
        },
        "headers": {
          "X-Powered-By": "Zend(?:Server)?(?:[\\s/]?([0-9.]+))?\\;version:\\1"
        },
        "icon": "Zend.png",
        "website": "http://zend.com"
      },
      "Zendesk": {
        "cats": [
          1,
          52
        ],
        "cookies": {
          "_help_center_session": "",
          "_zendesk_cookie": "",
          "_zendesk_shared_session": ""
        },
        "description": "Zendesk is a cloud-based help desk management solution offering customizable tools to build customer service portal, knowledge base and online communities.",
        "dns": {
          "TXT": [
            "mail\\.zendesk\\.com"
          ]
        },
        "headers": {
          "x-zendesk-user-id": ""
        },
        "icon": "Zendesk.png",
        "js": {
          "Zendesk": ""
        },
        "pricing": [
          "low"
        ],
        "saas": true,
        "scripts": "static\\.zdassets\\.com",
        "website": "https://zendesk.com"
      },
      "Zendesk Chat": {
        "cats": [
          52
        ],
        "description": "Zendesk Chat is a live chat and communication widget.",
        "icon": "Zendesk Chat.png",
        "pricing": [
          "freemium",
          "payg"
        ],
        "saas": true,
        "scripts": "v2\\.zopim\\.com",
        "website": "http://zopim.com"
      },
      "Zenfolio": {
        "cats": [
          7
        ],
        "description": "Zenfolio is a photography website builder.",
        "icon": "Zenfolio.png",
        "js": {
          "Zenfolio": ""
        },
        "website": "https://zenfolio.com"
      },
      "Zepto": {
        "cats": [
          59
        ],
        "icon": "Zepto.png",
        "js": {
          "Zepto": ""
        },
        "scripts": "zepto.*\\.js",
        "website": "http://zeptojs.com"
      },
      "Zimbra": {
        "cats": [
          30
        ],
        "cookies": {
          "ZM_TEST": "true"
        },
        "icon": "Zimbra.png",
        "implies": "Java",
        "website": "https://www.zimbra.com/"
      },
      "Zinnia": {
        "cats": [
          11
        ],
        "description": "Zimbra is a is a collaborative software suite that includes an email server and a web client.",
        "icon": "Zinnia.png",
        "implies": "Django",
        "meta": {
          "generator": "Zinnia"
        },
        "website": "http://django-blog-zinnia.com"
      },
      "Zip": {
        "cats": [
          41
        ],
        "description": "Zip is a payment service that lets you receive your purchase now and spread the total cost over a interest-free payment schedule.",
        "icon": "zip_pay.svg",
        "scripts": [
          "static\\.zipmoney\\.com\\.au",
          "zip\\.co"
        ],
        "website": "https://www.zip.co/"
      },
      "Zipkin": {
        "cats": [
          10
        ],
        "headers": {
          "X-B3-Flags": "",
          "X-B3-ParentSpanId": "",
          "X-B3-Sampled": "",
          "X-B3-SpanId": "",
          "X-B3-TraceId": ""
        },
        "icon": "Zipkin.png",
        "website": "https://zipkin.io/"
      },
      "Zocdoc": {
        "cats": [
          72
        ],
        "description": "Zocdoc is a New York City-based company offering an online service that allows people to find and book in-person or telemedicine appointments for medical or dental care.",
        "dom": "a[href*='www.zocdoc.com'][target='_blank']",
        "icon": "Zocdoc.svg",
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "offsiteschedule\\.zocdoc\\.com",
        "website": "https://www.zocdoc.com"
      },
      "Zoey": {
        "cats": [
          6
        ],
        "description": "Zoey is a cloud-based ecommerce platform for B2B and wholesale businesses.",
        "excludes": "Magento",
        "icon": "Zoey.svg",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "Zoey.module": "",
          "zoey.developer": "",
          "zoeyDev": ""
        },
        "scripts": "cdna4\\.zoeysite\\.com",
        "website": "https://www.zoey.com/"
      },
      "Zoho": {
        "cats": [
          53
        ],
        "description": "Zoho is a web-based online office suite.",
        "dns": {
          "TXT": [
            "\\.zoho\\.com"
          ]
        },
        "icon": "Zoho.svg",
        "pricing": [
          "low",
          "freemium"
        ],
        "saas": true,
        "website": "https://www.zoho.com/"
      },
      "Zoho Mail": {
        "cats": [
          75
        ],
        "description": "Zoho Mail is an email hosting service for businesses.",
        "dns": {
          "TXT": [
            "transmail\\.net"
          ]
        },
        "icon": "Zoho.svg",
        "implies": "Zoho",
        "website": "https://www.zoho.com/mail/"
      },
      "Zone.js": {
        "cats": [
          12
        ],
        "implies": "Angular",
        "js": {
          "Zone.root": ""
        },
        "website": "https://github.com/angular/angular/tree/master/packages/zone.js"
      },
      "Zonos": {
        "cats": [
          5
        ],
        "description": "Zonos is a cross-border ecommerce software and app solution for companies with international business.",
        "icon": "Zonos.svg",
        "js": {
          "Zonos": "",
          "zonos": "",
          "zonosCheckout": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "\\.zonos\\.com/",
        "website": "https://zonos.com"
      },
      "Zope": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "^Zope/"
        },
        "icon": "Zope.png",
        "website": "http://zope.org"
      },
      "a-blog cms": {
        "cats": [
          1
        ],
        "icon": "a-blog cms.svg",
        "implies": "PHP",
        "meta": {
          "generator": "a-blog cms"
        },
        "website": "http://www.a-blogcms.jp"
      },
      "actionhero.js": {
        "cats": [
          18,
          22
        ],
        "headers": {
          "X-Powered-By": "actionhero API"
        },
        "icon": "actionhero.js.png",
        "implies": "Node.js",
        "js": {
          "actionheroClient": ""
        },
        "scripts": "actionheroClient\\.js",
        "website": "http://www.actionherojs.com"
      },
      "amCharts": {
        "cats": [
          25
        ],
        "description": "amCharts is a JavaScript-based interactive charts and maps programming library and tool.",
        "html": "<svg[^>]*><desc>JavaScript chart by amCharts ([\\d.]*)\\;version:\\1",
        "icon": "amCharts.png",
        "js": {
          "AmCharts": ""
        },
        "scripts": "amcharts.*\\.js",
        "website": "http://amcharts.com"
      },
      "amoCRM": {
        "cats": [
          53
        ],
        "description": "amoCRM is a web-based customer relationship management software solution.",
        "icon": "amoCRM.svg",
        "js": {
          "AMOCRM": "",
          "AMO_PIXEL_CLIENT": "",
          "amoFormsWidget": "",
          "amoSocialButton": ""
        },
        "pricing": [
          "low",
          "recurring",
          "payg"
        ],
        "saas": true,
        "scripts": "\\.amocrm\\.(?:ru|com)",
        "website": "https://www.amocrm.com"
      },
      "animate.css": {
        "cats": [
          66
        ],
        "description": "Animate.css is a ready-to-use library collection of CSS3 animation effects.",
        "html": [
          "<link [^>]+(?:/([\\d.]+)/)?animate\\.(?:min\\.)?css\\;version:\\1"
        ],
        "website": "https://daneden.github.io/animate.css/"
      },
      "augmented-ui": {
        "cats": [
          66
        ],
        "description": "augmented-ui is a UI framework inspired by cyberpunk and sci-fi.",
        "dom": "[data-augmented-ui]",
        "icon": "augmented-ui.png",
        "oss": true,
        "website": "http://augmented-ui.com"
      },
      "basket.js": {
        "cats": [
          59
        ],
        "icon": "basket.js.png",
        "js": {
          "basket.isValidItem": ""
        },
        "scripts": "basket.*\\.js\\;confidence:10",
        "website": "https://addyosmani.github.io/basket.js/"
      },
      "cPanel": {
        "cats": [
          9
        ],
        "cookies": {
          "cprelogin": "",
          "cpsession": ""
        },
        "description": "cPanel is a web hosting control panel. The software provides a graphical interface and automation tools designed to simplify the process of hosting a website.",
        "headers": {
          "Server": "cpsrvd/([\\d.]+)\\;version:\\1"
        },
        "html": "<!-- cPanel",
        "icon": "cPanel.png",
        "website": "http://www.cpanel.net"
      },
      "cgit": {
        "cats": [
          19
        ],
        "html": [
          "<[^>]+id='cgit'",
          "generated by <a href='http://git\\.zx2c4\\.com/cgit/about/'>cgit v([\\d.a-z-]+)</a>\\;version:\\1"
        ],
        "icon": "cgit.png",
        "implies": "git",
        "meta": {
          "generator": "^cgit v([\\d.a-z-]+)$\\;version:\\1"
        },
        "website": "http://git.zx2c4.com/cgit"
      },
      "comScore": {
        "cats": [
          10
        ],
        "html": "<iframe[^>]* (?:id=\"comscore\"|scr=[^>]+comscore)|\\.scorecardresearch\\.com/beacon\\.js|COMSCORE\\.beacon",
        "icon": "comScore.png",
        "js": {
          "COMSCORE": "",
          "_COMSCORE": ""
        },
        "scripts": "\\.scorecardresearch\\.com/beacon\\.js|COMSCORE\\.beacon",
        "website": "http://comscore.com"
      },
      "commercelayer": {
        "cats": [
          6
        ],
        "icon": "commercelayer.svg",
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "website": "https://commercelayer.io",
        "xhr": "\\.commercelayer\\.io"
      },
      "db-ip": {
        "cats": [
          79
        ],
        "description": "dbip is a geolocation API and database.",
        "icon": "db-ip.png",
        "js": {
          "ENV.dbip": ""
        },
        "pricing": [
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "cdn\\.db-ip\\.com",
        "website": "https://db-ip.com/",
        "xhr": "api\\.db-ip\\.com"
      },
      "decimal.js": {
        "cats": [
          59
        ],
        "icon": "decimal.js.png",
        "js": {
          "Decimal.ROUND_HALF_FLOOR": ""
        },
        "scripts": [
          "decimal[.-]([\\d.]*\\d+)(?:\\.min)?\\.js\\;version:\\1",
          "/([\\d.]*\\d+)/decimal(?:\\.min)?\\.js\\;version:\\1",
          "decimal(?:\\.min)?\\.js(?:\\?ver(?:sion)?=([\\d.]*\\d+))?\\;version:\\1"
        ],
        "website": "https://mikemcl.github.io/decimal.js/"
      },
      "deepMiner": {
        "cats": [
          56
        ],
        "icon": "deepminer.png",
        "js": {
          "deepMiner": ""
        },
        "scripts": "deepMiner\\.js",
        "website": "https://github.com/deepwn/deepMiner"
      },
      "e107": {
        "cats": [
          1
        ],
        "cookies": {
          "e107_tz": ""
        },
        "headers": {
          "X-Powered-By": "e107"
        },
        "icon": "e107.png",
        "implies": "PHP",
        "scripts": "[^a-z\\d]e107\\.js",
        "website": "http://e107.org"
      },
      "eClass": {
        "cats": [
          1
        ],
        "description": "eClass is an online learning platform.",
        "icon": "eClass.png",
        "js": {
          "fe_eclass": "\\;confidence:50",
          "fe_eclass_guest": "\\;confidence:50"
        },
        "website": "https://www.eclass.com.hk"
      },
      "eDokan": {
        "cats": [
          6
        ],
        "description": "eDokan is hosted ecommerce platform with drag-drop template builder and zero programming knowledge.",
        "dom": "img[src*='cdn.edokan.co']",
        "icon": "eDokan.png",
        "implies": [
          "Node.js",
          "Angular",
          "MongoDB"
        ],
        "pricing": [
          "low",
          "recurring"
        ],
        "saas": true,
        "website": "https://edokan.co"
      },
      "eSyndiCat": {
        "cats": [
          1
        ],
        "headers": {
          "X-Drectory-Script": "^eSyndiCat"
        },
        "icon": "eSyndiCat.png",
        "implies": "PHP",
        "js": {
          "esyndicat": ""
        },
        "meta": {
          "generator": "^eSyndiCat "
        },
        "website": "http://esyndicat.com"
      },
      "eWAY Payments": {
        "cats": [
          41
        ],
        "description": "eWAY is a global omnichannel payment provider. The company processes secure credit card payments for merchants. eWay works through eCommerce.",
        "html": "<img [^>]*src=\"[^/]*//[^/]*eway\\.com",
        "icon": "eway.png",
        "scripts": "secure\\.ewaypayments\\.com",
        "website": "https://www.eway.com.au/"
      },
      "eZ Platform": {
        "cats": [
          1,
          6
        ],
        "icon": "eZ.svg",
        "implies": "Symfony",
        "meta": {
          "generator": "eZ Platform"
        },
        "website": "https://ezplatform.com/"
      },
      "eZ Publish": {
        "cats": [
          1,
          6
        ],
        "cookies": {
          "eZSESSID": ""
        },
        "headers": {
          "X-Powered-By": "^eZ Publish"
        },
        "icon": "eZ.svg",
        "implies": "PHP",
        "meta": {
          "generator": "eZ Publish"
        },
        "website": "https://github.com/ezsystems/ezpublish-legacy"
      },
      "ef.js": {
        "cats": [
          12
        ],
        "icon": "ef.js.svg",
        "js": {
          "ef.version": "^(.+)$\\;version:\\1",
          "efCore": ""
        },
        "scripts": "/ef(?:-core)?(?:\\.min|\\.dev)?\\.js",
        "website": "http://ef.js.org"
      },
      "emBlue": {
        "cats": [
          32,
          75
        ],
        "description": "emBlue is an email and marketing automation platform.",
        "icon": "emBlue.svg",
        "js": {
          "emblueOnSiteApp": ""
        },
        "pricing": [
          "freemium",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.embluemail\\.com/(?:library/([\\d.]+))?\\;version:\\1",
        "website": "https://www.embluemail.com/en"
      },
      "enduro.js": {
        "cats": [
          1
        ],
        "headers": {
          "X-Powered-By": "^enduro\\.js"
        },
        "icon": "enduro.js.svg",
        "implies": "Node.js",
        "website": "http://endurojs.com"
      },
      "eucookie.eu": {
        "cats": [
          67
        ],
        "icon": "eucookie.png",
        "scripts": "eucookie\\.eu/public/gdpr-cookie-consent\\.js",
        "website": "https://www.eucookie.eu/"
      },
      "experiencedCMS": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:experiencedcms:experiencedcms",
        "icon": "experiencedCMS_Logo.png",
        "implies": "PHP",
        "meta": {
          "generator": "^experiencedCMS$"
        },
        "website": "https://experiencedcms.berkearas.de"
      },
      "git": {
        "cats": [
          47
        ],
        "cpe": "cpe:/a:git-scm:git",
        "icon": "git.svg",
        "meta": {
          "generator": "\\bgit/([\\d.]+\\d)\\;version:\\1"
        },
        "website": "http://git-scm.com"
      },
      "gitlist": {
        "cats": [
          47
        ],
        "cpe": "cpe:/a:gitlist:gitlist",
        "html": "<p>Powered by <a[^>]+>GitList ([\\d.]+)\\;version:\\1",
        "implies": [
          "PHP",
          "git"
        ],
        "website": "http://gitlist.org"
      },
      "gitweb": {
        "cats": [
          47
        ],
        "html": "<!-- git web interface version ([\\d.]+)?\\;version:\\1",
        "icon": "git.svg",
        "implies": [
          "Perl",
          "git"
        ],
        "meta": {
          "generator": "gitweb(?:/([\\d.]+\\d))?\\;version:\\1"
        },
        "scripts": "static/gitweb\\.js$",
        "website": "http://git-scm.com"
      },
      "govCMS": {
        "cats": [
          1
        ],
        "icon": "govCMS.svg",
        "implies": "Drupal",
        "meta": {
          "generator": "Drupal ([\\d]+) \\(http:\\/\\/drupal\\.org\\) \\+ govCMS\\;version:\\1"
        },
        "website": "https://www.govcms.gov.au"
      },
      "gunicorn": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "gunicorn(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "gunicorn.png",
        "website": "http://gunicorn.org"
      },
      "hCaptcha": {
        "cats": [
          16
        ],
        "html": "<style[^>]+[^<]+#cf-hcaptcha-container[^<]+</style>",
        "icon": "hCaptcha.svg",
        "scripts": "https://hcaptcha.com/([\\d]+?)/api.js\\;version:\\1",
        "website": "https://www.hcaptcha.com/"
      },
      "hantana": {
        "cats": [
          10
        ],
        "icon": "hantana.png",
        "js": {
          "Hantana": ""
        },
        "scripts": "//hantana\\.org/widget",
        "website": "https://hantana.org/"
      },
      "iEXExchanger": {
        "cats": [
          1
        ],
        "cookies": {
          "iexexchanger_session": ""
        },
        "icon": "iEXExchanger.png",
        "implies": [
          "PHP",
          "Apache",
          "Angular"
        ],
        "meta": {
          "generator": "iEXExchanger"
        },
        "website": "https://exchanger.iexbase.com"
      },
      "iPresta": {
        "cats": [
          6
        ],
        "icon": "iPresta.png",
        "implies": [
          "PHP",
          "PrestaShop"
        ],
        "meta": {
          "designer": "iPresta"
        },
        "website": "http://ipresta.ir"
      },
      "iWeb": {
        "cats": [
          20
        ],
        "description": "iWeb is a web site creation tool.",
        "icon": "iWeb.png",
        "meta": {
          "generator": "^iWeb( [\\d.]+)?\\;version:\\1"
        },
        "website": "http://apple.com/ilife/iweb"
      },
      "ikiwiki": {
        "cats": [
          8
        ],
        "description": "ikiwiki is a free and open-source wiki application.",
        "html": [
          "<link rel=\"alternate\" type=\"application/x-wiki\" title=\"Edit this page\" href=\"[^\"]*/ikiwiki\\.cgi",
          "<a href=\"/(?:cgi-bin/)?ikiwiki\\.cgi\\?do="
        ],
        "icon": "ikiwiki.png",
        "website": "http://ikiwiki.info"
      },
      "imperia CMS": {
        "cats": [
          1
        ],
        "html": "<imp:live-info sysid=\"[0-9a-f-]+\"(?: node_id=\"[0-9/]*\")? *\\/>",
        "icon": "imperiaCMS.svg",
        "implies": "Perl",
        "meta": {
          "GENERATOR": "^IMPERIA ([0-9.]{2,3})\\;version:\\1",
          "X-Imperia-Live-Info": ""
        },
        "url": "imperia/md/",
        "website": "https://www.pirobase-imperia.com/de/produkte/produktuebersicht/imperia-cms"
      },
      "ip-api": {
        "cats": [
          79
        ],
        "icon": "ip-api.png",
        "pricing": [
          "freemium",
          "low",
          "payg"
        ],
        "saas": true,
        "website": "https://ip-api.com/",
        "xhr": "ip-api\\.com/"
      },
      "ip-label": {
        "cats": [
          10
        ],
        "icon": "iplabel.svg",
        "js": {
          "clobs": ""
        },
        "scripts": "clobs\\.js",
        "website": "http://www.ip-label.com"
      },
      "ipapi": {
        "cats": [
          79
        ],
        "description": "ipapi is a real-time geolocation and reverse IP lookup REST API.",
        "icon": "ipapi.png",
        "pricing": [
          "freemium",
          "low",
          "payg"
        ],
        "saas": true,
        "website": "https://ipapi.com",
        "xhr": "api\\.ipapi\\.com"
      },
      "ipapi.co": {
        "cats": [
          79
        ],
        "description": "ipapi.co is a web analytics provider with IP address lookup and location API.",
        "icon": "ipapi.co.png",
        "pricing": [
          "freemium",
          "low",
          "payg"
        ],
        "saas": true,
        "website": "https://ipapi.co",
        "xhr": "ipapi\\.co/"
      },
      "ipdata": {
        "cats": [
          79
        ],
        "description": "ipdata is a JSON IP Address Geolocation API that allows to lookup the location of both IPv4 and IPv6.",
        "icon": "ipdata.png",
        "pricing": [
          "freemium",
          "recurring",
          "low"
        ],
        "saas": true,
        "website": "https://ipdata.co/",
        "xhr": "api\\.ipdata\\.co"
      },
      "ipgeolocation": {
        "cats": [
          79
        ],
        "description": "ipgeolocation is an IP Geolocation API and Accurate IP Lookup Database.",
        "icon": "ipgeolocation.png",
        "pricing": [
          "freemium",
          "recurring",
          "mid"
        ],
        "saas": true,
        "website": "https://ipgeolocation.co/",
        "xhr": "api\\.ipgeolocation\\.io"
      },
      "ipify": {
        "cats": [
          79
        ],
        "description": "ipify is a service which provide public IP address API, IP geolocation API, VPN and Proxy detection API products.",
        "icon": "ipify.png",
        "pricing": [
          "freemium",
          "payg",
          "mid",
          "recurring"
        ],
        "saas": true,
        "scripts": "\\.ipify\\.org",
        "website": "https://ipify.org",
        "xhr": "(?:api|api64|geo)\\.ipify\\.org"
      },
      "ipstack": {
        "cats": [
          79
        ],
        "description": "ipstack is a real-time IP to geolocation API capable of looking at location data and assessing security threats originating from risky IP addresses.",
        "icon": "ipstack.png",
        "js": {
          "ENV.ipStackAccessToken": ""
        },
        "pricing": [
          "low",
          "freemium"
        ],
        "saas": true,
        "website": "https://ipstack.com",
        "xhr": "api\\.ipstack\\.com"
      },
      "iubenda": {
        "cats": [
          67
        ],
        "icon": "iubenda.png",
        "scripts": [
          "iubenda\\.com/cookie-solution/confs/js/"
        ],
        "website": "https://www.iubenda.com/"
      },
      "jComponent": {
        "cats": [
          12,
          59
        ],
        "icon": "jComponent.png",
        "implies": "jQuery",
        "js": {
          "MAIN.version": ".*\\;version:\\1"
        },
        "website": "https://componentator.com"
      },
      "jQTouch": {
        "cats": [
          26
        ],
        "description": "jQTouch is an open-source Zepto/ JQuery plugin with native animations, automatic navigation, and themes for mobile WebKit browsers like iPhone, G1 (Android), and Palm Pre.",
        "icon": "jQTouch.png",
        "js": {
          "jQT": ""
        },
        "scripts": "jqtouch.*\\.js",
        "website": "http://jqtouch.com"
      },
      "jQuery": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:jquery:jquery",
        "description": "jQuery is a JavaScript library which is a free, open-source software designed to simplify HTML DOM tree traversal and manipulation, as well as event handling, CSS animation, and Ajax.",
        "icon": "jQuery.svg",
        "js": {
          "$.fn.jquery": "([\\d.]+)\\;version:\\1",
          "jQuery.fn.jquery": "([\\d.]+)\\;version:\\1"
        },
        "scripts": [
          "jquery[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "/([\\d.]+)/jquery(?:\\.min)?\\.js\\;version:\\1",
          "jquery.*\\.js(?:\\?ver(?:sion)?=([\\d.]+))?\\;version:\\1"
        ],
        "website": "https://jquery.com"
      },
      "jQuery DevBridge Autocomplete": {
        "cats": [
          59
        ],
        "description": "Ajax Autocomplete for jQuery allows you to easily create autocomplete/autosuggest boxes for text input fields.",
        "icon": "jQuery.svg",
        "implies": "jQuery",
        "js": {
          "$.devbridgeAutocomplete": "",
          "jQuery.devbridgeAutocomplete": ""
        },
        "scripts": [
          "/devbridgeAutocomplete(?:-min)?\\.js",
          "/jquery\\.devbridge-autocomplete/([0-9.]+)/jquery\\.autocomplete(?:.min)?\\.js\\;version:\\1"
        ],
        "website": "https://www.devbridge.com/sourcery/components/jquery-autocomplete/"
      },
      "jQuery Migrate": {
        "cats": [
          59
        ],
        "description": "Query Migrate is a javascript library that allows you to preserve the compatibility of your jQuery code developed for versions of jQuery older than 1.9.",
        "icon": "jQuery.svg",
        "implies": "jQuery",
        "js": {
          "jQuery.migrateVersion": "([\\d.]+)\\;version:\\1",
          "jQuery.migrateWarnings": "",
          "jqueryMigrate": ""
        },
        "scripts": "jquery[.-]migrate(?:-([\\d.]+))?(?:\\.min)?\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1?\\1:\\2",
        "website": "https://github.com/jquery/jquery-migrate"
      },
      "jQuery Mobile": {
        "cats": [
          26
        ],
        "description": "jQuery Mobile is a HTML5-based user interface system designed to make responsive web sites and apps that are accessible on all smartphone, tablet and desktop devices.",
        "icon": "jQuery Mobile.svg",
        "implies": "jQuery",
        "js": {
          "jQuery.mobile.version": "^(.+)$\\;version:\\1"
        },
        "scripts": "jquery[.-]mobile(?:-([\\d.]))?(?:\\.min)?\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1?\\1:\\2",
        "website": "https://jquerymobile.com"
      },
      "jQuery Sparklines": {
        "cats": [
          25
        ],
        "description": "jQuery Sparklines is a plugin that generates sparklines (small inline charts) directly in the browser using data supplied either inline in the HTML, or via javascript.",
        "implies": "jQuery",
        "scripts": "jquery\\.sparkline.*\\.js",
        "website": "http://omnipotent.net/jquery.sparkline/"
      },
      "jQuery UI": {
        "cats": [
          59
        ],
        "cpe": "cpe:/a:jquery:jquery_ui",
        "description": "jQuery UI is a collection of GUI widgets, animated visual effects, and themes implemented with jQuery, Cascading Style Sheets, and HTML.",
        "icon": "jQuery UI.svg",
        "implies": "jQuery",
        "js": {
          "jQuery.ui.version": "^(.+)$\\;version:\\1"
        },
        "scripts": [
          "jquery-ui[.-]([\\d.]*\\d)[^/]*\\.js\\;version:\\1",
          "([\\d.]+)/jquery-ui(?:\\.min)?\\.js\\;version:\\1",
          "jquery-ui.*\\.js"
        ],
        "website": "http://jqueryui.com"
      },
      "jQuery-pjax": {
        "cats": [
          26
        ],
        "description": "jQuery PJAX is a plugin that uses AJAX and pushState.",
        "html": "<div[^>]+data-pjax-container",
        "implies": "jQuery",
        "js": {
          "jQuery.pjax": ""
        },
        "meta": {
          "pjax-push": "",
          "pjax-replace": "",
          "pjax-timeout": ""
        },
        "scripts": "jquery[.-]pjax(?:-([\\d.]))?(?:\\.min)?\\.js(?:\\?ver=([\\d.]+))?\\;version:\\1?\\1:\\2",
        "website": "https://github.com/defunkt/jquery-pjax"
      },
      "jqPlot": {
        "cats": [
          25
        ],
        "icon": "jqPlot.png",
        "implies": "jQuery",
        "scripts": "jqplot.*\\.js",
        "website": "http://www.jqplot.com"
      },
      "jsDelivr": {
        "cats": [
          31
        ],
        "description": "JSDelivr is a free public CDN for open-source projects. It can serve web files directly from the npm registry and GitHub repositories without any configuration.",
        "html": "<link [^>]*?href=\"?[a-zA-Z]*?:?//cdn\\.jsdelivr\\.net/",
        "icon": "jsdelivr-icon.svg",
        "scripts": "//cdn\\.jsdelivr\\.net/",
        "website": "https://www.jsdelivr.com/"
      },
      "libwww-perl-daemon": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "libwww-perl-daemon(?:/([\\d\\.]+))?\\;version:\\1"
        },
        "icon": "libwww-perl-daemon.png",
        "implies": "Perl",
        "website": "http://metacpan.org/pod/HTTP::Daemon"
      },
      "lighttpd": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "lighttpd(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "lighttpd.png",
        "website": "http://www.lighttpd.net"
      },
      "math.js": {
        "cats": [
          59
        ],
        "icon": "math.js.png",
        "js": {
          "mathjs": ""
        },
        "scripts": "math(?:\\.min)?\\.js",
        "website": "http://mathjs.org"
      },
      "mini_httpd": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "mini_httpd(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "mini_httpd.png",
        "website": "http://acme.com/software/mini_httpd"
      },
      "mod_auth_pam": {
        "cats": [
          33
        ],
        "description": "Mod_auth_pam is used to configure ways for authenticating users.",
        "headers": {
          "Server": "mod_auth_pam(?:/([\\d\\.]+))?\\;version:\\1"
        },
        "icon": "Apache.svg",
        "implies": "Apache",
        "website": "http://pam.sourceforge.net/mod_auth_pam"
      },
      "mod_dav": {
        "cats": [
          33
        ],
        "description": "Mod_dav is an Apache module to provide WebDAV capabilities for your Apache web server. It is an open-source module, provided under an Apache-style license.",
        "headers": {
          "Server": "\\b(?:mod_)?DAV\\b(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Apache.svg",
        "implies": "Apache",
        "website": "http://webdav.org/mod_dav"
      },
      "mod_fastcgi": {
        "cats": [
          33
        ],
        "description": "Mod_fcgid is a high performance alternative to mod_cgi or mod_cgid, which starts a sufficient number instances of the CGI program to handle concurrent requests, and these programs remain running to handle further incoming requests.",
        "headers": {
          "Server": "mod_fastcgi(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Apache.svg",
        "implies": "Apache",
        "website": "http://www.fastcgi.com/mod_fastcgi/docs/mod_fastcgi.html"
      },
      "mod_jk": {
        "cats": [
          33
        ],
        "description": "Mod_jk is an Apache module used to connect the Tomcat servlet container with web servers such as Apache, iPlanet, Sun ONE (formerly Netscape) and even IIS using the Apache JServ Protocol. A web server waits for client HTTP requests.",
        "headers": {
          "Server": "mod_jk(?:/([\\d\\.]+))?\\;version:\\1"
        },
        "icon": "Apache.svg",
        "implies": [
          "Apache Tomcat",
          "Apache"
        ],
        "website": "http://tomcat.apache.org/tomcat-3.3-doc/mod_jk-howto.html"
      },
      "mod_perl": {
        "cats": [
          33
        ],
        "description": "Mod_perl is an optional module for the Apache HTTP server. It embeds a Perl interpreter into the Apache server. In addition to allowing Apache modules to be written in the Perl programming language, it allows the Apache web server to be dynamically configured by Perl programs.",
        "headers": {
          "Server": "mod_perl(?:/([\\d\\.]+))?\\;version:\\1"
        },
        "icon": "mod_perl.png",
        "implies": [
          "Perl",
          "Apache"
        ],
        "website": "http://perl.apache.org"
      },
      "mod_python": {
        "cats": [
          33
        ],
        "description": "Mod_python is an Apache HTTP Server module that integrates the Python programming language with the server. It is intended to provide a Python language binding for the Apache HTTP Server. ",
        "headers": {
          "Server": "mod_python(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "mod_python.png",
        "implies": [
          "Python",
          "Apache"
        ],
        "website": "http://www.modpython.org"
      },
      "mod_rack": {
        "cats": [
          33
        ],
        "description": "Mod_rack is a free web server and application server with support for Ruby, Python and Node.js.",
        "headers": {
          "Server": "mod_rack(?:/([\\d.]+))?\\;version:\\1",
          "X-Powered-By": "mod_rack(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Phusion Passenger.png",
        "implies": [
          "Ruby on Rails\\;confidence:50",
          "Apache"
        ],
        "website": "http://phusionpassenger.com"
      },
      "mod_rails": {
        "cats": [
          33
        ],
        "description": "Mod_rails is a free web server and application server with support for Ruby, Python and Node.js.",
        "headers": {
          "Server": "mod_rails(?:/([\\d.]+))?\\;version:\\1",
          "X-Powered-By": "mod_rails(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "Phusion Passenger.png",
        "implies": [
          "Ruby on Rails\\;confidence:50",
          "Apache"
        ],
        "website": "http://phusionpassenger.com"
      },
      "mod_ssl": {
        "cats": [
          33
        ],
        "description": "mod_ssl is an optional module for the Apache HTTP Server. It provides strong cryptography for the Apache web server via the Secure Sockets Layer (SSL) and Transport Layer Security (TLS) cryptographic protocols by the help of the open-source SSL/TLS toolkit OpenSSL.",
        "headers": {
          "Server": "mod_ssl(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "mod_ssl.png",
        "implies": "Apache",
        "website": "http://modssl.org"
      },
      "mod_wsgi": {
        "cats": [
          33
        ],
        "description": "mod_wsgi is an Apache HTTP Server module that provides a WSGI compliant interface for hosting Python based web applications under Apache.",
        "headers": {
          "Server": "mod_wsgi(?:/([\\d.]+))?\\;version:\\1",
          "X-Powered-By": "mod_wsgi(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "mod_wsgi.png",
        "implies": [
          "Python\\;confidence:50",
          "Apache"
        ],
        "website": "https://code.google.com/p/modwsgi"
      },
      "nghttpx - HTTP/2 proxy": {
        "cats": [
          22
        ],
        "headers": {
          "Server": "nghttpx nghttp2/?([\\d.]+)?\\;version:\\1"
        },
        "website": "https://nghttp2.org"
      },
      "nopCommerce": {
        "cats": [
          6
        ],
        "cookies": {
          "Nop.customer": ""
        },
        "html": "(?:<!--Powered by nopCommerce|Powered by: <a[^>]+nopcommerce)",
        "icon": "nopCommerce.png",
        "implies": "Microsoft ASP.NET",
        "meta": {
          "generator": "^nopCommerce$"
        },
        "oss": true,
        "website": "http://www.nopcommerce.com"
      },
      "osCommerce": {
        "cats": [
          6
        ],
        "cookies": {
          "osCsid": ""
        },
        "html": [
          "<br />Powered by <a href=\"https?://www\\.oscommerce\\.com",
          "<(?:input|a)[^>]+name=\"osCsid\"",
          "<(?:tr|td|table)class=\"[^\"]*infoBoxHeading"
        ],
        "icon": "osCommerce.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "website": "https://www.oscommerce.com"
      },
      "osTicket": {
        "cats": [
          13
        ],
        "cookies": {
          "OSTSESSID": ""
        },
        "icon": "osTicket.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "website": "http://osticket.com"
      },
      "otrs": {
        "cats": [
          13
        ],
        "headers": {
          "X-Powered-By": "OTRS ([\\d.]+)\\;version:\\1"
        },
        "html": "<!--\\s+OTRS: Copyright",
        "icon": "otrs.png",
        "implies": "Perl",
        "scripts": "^/otrs-web/js/",
        "website": "https://www.otrs.com"
      },
      "ownCloud": {
        "cats": [
          19
        ],
        "html": "<a href=\"https://owncloud\\.com\" target=\"_blank\">ownCloud Inc\\.</a><br/>Your Cloud, Your Data, Your Way!",
        "icon": "ownCloud.png",
        "implies": "PHP",
        "meta": {
          "apple-itunes-app": "app-id=543672169"
        },
        "website": "https://owncloud.org"
      },
      "papaya CMS": {
        "cats": [
          1
        ],
        "html": "<link[^>]*/papaya-themes/",
        "icon": "papaya CMS.png",
        "implies": "PHP",
        "website": "https://papaya-cms.com"
      },
      "parcel": {
        "cats": [
          19
        ],
        "icon": "Parcel.png",
        "js": {
          "parcelRequire": ""
        },
        "website": "https://parceljs.org/"
      },
      "particles.js": {
        "cats": [
          25
        ],
        "description": "Particles.js is a JavaScript library for creating particles.",
        "html": "<div id=\"particles-js\">",
        "js": {
          "particlesJS": ""
        },
        "scripts": "/particles(?:\\.min)?\\.js",
        "website": "https://vincentgarreau.com/particles.js/"
      },
      "phpAlbum": {
        "cats": [
          7
        ],
        "description": "phpAlbum is an open-source PHP script which allows you to create your personal photo album.",
        "html": "<!--phpalbum ([.\\d\\s]+)-->\\;version:\\1",
        "icon": "phpAlbum.png",
        "implies": "PHP",
        "website": "http://phpalbum.net"
      },
      "phpBB": {
        "cats": [
          2
        ],
        "cookies": {
          "phpbb": ""
        },
        "description": "phpBB is a free open-source Internet forum package in the PHP scripting language.",
        "html": [
          "Powered by <a[^>]+phpBB",
          "<div class=phpbb_copyright>",
          "<[^>]+styles/(?:sub|pro)silver/theme",
          "<img[^>]+i_icon_mini",
          "<table class=\"[^\"]*forumline"
        ],
        "icon": "phpBB.png",
        "implies": "PHP",
        "js": {
          "phpbb": "",
          "style_cookie_settings": ""
        },
        "meta": {
          "copyright": "phpBB Group"
        },
        "website": "https://phpbb.com"
      },
      "phpCMS": {
        "cats": [
          1
        ],
        "icon": "PHP.svg",
        "implies": "PHP",
        "js": {
          "phpcms": ""
        },
        "website": "http://phpcms.de"
      },
      "phpDocumentor": {
        "cats": [
          4
        ],
        "description": "phpDocumentor is an open-source documentation generator written in PHP.",
        "html": "<!-- Generated by phpDocumentor",
        "icon": "phpDocumentor.png",
        "implies": "PHP",
        "website": "https://www.phpdoc.org"
      },
      "phpMyAdmin": {
        "cats": [
          3
        ],
        "description": "PhpMyAdmin is a free and open-source administration tool for MySQL and MariaDB.",
        "html": [
          "!\\[CDATA\\[[^<]*PMA_VERSION:\\\"([\\d.]+)\\;version:\\1",
          "(?: \\| phpMyAdmin ([\\d.]+)<\\/title>|PMA_sendHeaderLocation\\(|<link [^>]*href=\"[^\"]*phpmyadmin\\.css\\.php)\\;version:\\1"
        ],
        "icon": "phpMyAdmin.png",
        "implies": [
          "PHP",
          "MySQL"
        ],
        "js": {
          "pma_absolute_uri": ""
        },
        "website": "https://www.phpmyadmin.net"
      },
      "phpPgAdmin": {
        "cats": [
          3
        ],
        "html": "(?:<title>phpPgAdmin</title>|<span class=\"appname\">phpPgAdmin)",
        "icon": "phpPgAdmin.png",
        "implies": "PHP",
        "website": "http://phppgadmin.sourceforge.net"
      },
      "phpSQLiteCMS": {
        "cats": [
          1
        ],
        "icon": "phpSQLiteCMS.png",
        "implies": [
          "PHP",
          "SQLite"
        ],
        "meta": {
          "generator": "^phpSQLiteCMS(?: (.+))?$\\;version:\\1"
        },
        "website": "http://phpsqlitecms.net"
      },
      "phpwind": {
        "cats": [
          1,
          2
        ],
        "html": "(?:Powered|Code) by <a href=\"[^\"]+phpwind\\.net",
        "icon": "phpwind.png",
        "implies": "PHP",
        "meta": {
          "generator": "^phpwind(?: v([0-9-]+))?\\;version:\\1"
        },
        "website": "https://www.phpwind.net"
      },
      "pinoox": {
        "cats": [
          18
        ],
        "cookies": {
          "pinoox_session": ""
        },
        "icon": "pinoox.png",
        "implies": "PHP",
        "js": {
          "pinoox": ""
        },
        "website": "https://pinoox.com"
      },
      "pirobase CMS": {
        "cats": [
          1
        ],
        "html": [
          "<(?:script|link)[^>]/site/[a-z0-9/._-]+/resourceCached/[a-z0-9/._-]+",
          "<input[^>]+cbi:///cms/"
        ],
        "icon": "pirobaseCMS.svg",
        "implies": "Java",
        "website": "https://www.pirobase-imperia.com/de/produkte/produktuebersicht/pirobase-cms"
      },
      "plentymarkets": {
        "cats": [
          6
        ],
        "headers": {
          "X-Plenty-Shop": ""
        },
        "icon": "plentymarkets.svg",
        "meta": {
          "generator": "plentymarkets"
        },
        "scripts": [
          "plenty\\.shop\\.(?:min\\.)?js"
        ],
        "website": "https://www.plentymarkets.com/"
      },
      "prettyPhoto": {
        "cats": [
          59
        ],
        "html": "(?:<link [^>]*href=\"[^\"]*prettyPhoto(?:\\.min)?\\.css|<a [^>]*rel=\"prettyPhoto)",
        "icon": "prettyPhoto.png",
        "implies": "jQuery",
        "js": {
          "pp_alreadyInitialized": "",
          "pp_descriptions": "",
          "pp_images": "",
          "pp_titles": ""
        },
        "scripts": "jquery\\.prettyPhoto\\.js",
        "website": "http://no-margin-for-errors.com/projects/prettyphoto-jquery-lightbox-clone/"
      },
      "punBB": {
        "cats": [
          2
        ],
        "html": "Powered by <a href=\"[^>]+punbb",
        "icon": "punBB.png",
        "implies": "PHP",
        "js": {
          "PUNBB": ""
        },
        "website": "http://punbb.informer.com"
      },
      "reCAPTCHA": {
        "cats": [
          16
        ],
        "html": [
          "<div[^>]+id=\"recaptcha_image",
          "<link[^>]+recaptcha",
          "<div[^>]+class=\"g-recaptcha\""
        ],
        "icon": "reCAPTCHA.svg",
        "js": {
          "Recaptcha": "",
          "recaptcha": ""
        },
        "scripts": [
          "api-secure\\.recaptcha\\.net",
          "recaptcha_ajax\\.js",
          "/recaptcha/api\\.js"
        ],
        "website": "https://www.google.com/recaptcha/"
      },
      "sIFR": {
        "cats": [
          17
        ],
        "description": "sIFR is a JavaScript and Adobe Flash dynamic web fonts implementation.",
        "icon": "sIFR.png",
        "scripts": "sifr\\.js",
        "website": "https://www.mikeindustries.com/blog/sifr"
      },
      "sNews": {
        "cats": [
          1
        ],
        "icon": "sNews.png",
        "meta": {
          "generator": "sNews"
        },
        "website": "https://snewscms.com"
      },
      "script.aculo.us": {
        "cats": [
          59
        ],
        "icon": "script.aculo.us.png",
        "js": {
          "Scriptaculous.Version": "^(.+)$\\;version:\\1"
        },
        "scripts": "/(?:scriptaculous|protoaculous)(?:\\.js|/)",
        "website": "https://script.aculo.us"
      },
      "scrollreveal": {
        "cats": [
          59
        ],
        "html": "<[^>]+data-sr(?:-id)",
        "icon": "scrollreveal.svg",
        "js": {
          "ScrollReveal().version": "^(.+)$\\;version:\\1"
        },
        "scripts": "scrollreveal(?:\\.min)(?:\\.js)",
        "website": "https://scrollrevealjs.org"
      },
      "shine.js": {
        "cats": [
          25
        ],
        "js": {
          "Shine": ""
        },
        "scripts": "shine(?:\\.min)?\\.js",
        "website": "https://bigspaceship.github.io/shine.js/"
      },
      "snigel AdConsent": {
        "cats": [
          67
        ],
        "description": "snigel AdConsent is a IAB-registered consent managment platfrom (CMP) which help keep sites speed intact while remaining compliant with GDPR and CCPA.",
        "icon": "snigel.svg",
        "js": {
          "adconsent.version": "^([\\d.]+)$\\;version:\\1",
          "snhb.baseSettings": "\\;confidence:50",
          "snhb.info.cmpVersion": ""
        },
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": "(?:staging-)?cdn\\.snigelweb\\.com/(?:snhb|adconsent)/",
        "website": "https://www.snigel.com/adconsent/"
      },
      "snigel AdEngine": {
        "cats": [
          36
        ],
        "description": "snigel AdEngine is a header bidding solution product from snigel.",
        "icon": "snigel.svg",
        "pricing": [
          "poa"
        ],
        "saas": true,
        "scripts": [
          "(?:staging-)?cdn\\.snigelweb\\.com/adengine/",
          "adengine\\.snigelweb\\.com"
        ],
        "website": "https://www.snigel.com/adengine/"
      },
      "styled-components": {
        "cats": [
          12,
          47
        ],
        "description": "Styled components is a CSS-in-JS styling framework that uses tagged template literals in JavaScript.",
        "dom": {
          "[sc-component-id]": {
            "text": ""
          },
          "style[data-styled-version]": {
            "attributes": {
              "data-styled-version": "(^.+$)\\;version:\\1"
            }
          },
          "style[data-styled], style[data-styled-components], [sc-component-id]": {
            "text": ""
          }
        },
        "icon": "styled-components.png",
        "implies": "React",
        "js": {
          "styled": ""
        },
        "website": "https://styled-components.com"
      },
      "tailwindcss": {
        "cats": [
          66
        ],
        "description": "Tailwind is a utility-first CSS framework.",
        "dom": {
          "link[rel='stylesheet'][href*='tailwind']": {
            "attributes": {
              "href": ""
            }
          }
        },
        "html": [
          "<link[^>]+?href=\"[^\"]+tailwindcss[@|/](?:\\^)?([\\d.]+)(?:/[a-z]+)?/(?:tailwind|base|components|utilities)(?:\\.min)?\\.css\\;version:\\1"
        ],
        "icon": "tailwindcss.svg",
        "website": "https://tailwindcss.com/"
      },
      "theTradeDesk": {
        "cats": [
          36
        ],
        "description": "theTradeDesk is an technology company that markets a software platform used by digital ad buyers to purchase data-driven digital advertising campaigns across various ad formats and devices.",
        "icon": "theTradeDesk.svg",
        "js": {
          "TTDUniversalPixelApi": "",
          "ttd_dom_ready": ""
        },
        "pricing": [
          "payg"
        ],
        "saas": true,
        "scripts": "\\.adsrvr\\.org/",
        "website": "https://www.thetradedesk.com",
        "xhr": "adsvr\\.org"
      },
      "three.js": {
        "cats": [
          25
        ],
        "description": "Three.js is a cross-browser JavaScript library and application programming interface used to create and display animated 3D computer graphics in a web browser using WebGL.",
        "icon": "three.js.png",
        "js": {
          "THREE.REVISION": "^(.+)$\\;version:\\1"
        },
        "scripts": "three(?:\\.min)?\\.js",
        "website": "https://threejs.org"
      },
      "thttpd": {
        "cats": [
          22
        ],
        "cpe": "cpe:/a:acme:thttpd",
        "headers": {
          "Server": "\\bthttpd(?:/([\\d.]+))?\\;version:\\1"
        },
        "icon": "thttpd.png",
        "website": "https://acme.com/software/thttpd"
      },
      "total.js": {
        "cats": [
          18
        ],
        "headers": {
          "X-Powered-By": "^total\\.js"
        },
        "icon": "total.js.png",
        "implies": "Node.js",
        "website": "https://totaljs.com"
      },
      "uKnowva": {
        "cats": [
          1,
          2,
          50
        ],
        "headers": {
          "X-Content-Encoded-By": "uKnowva ([\\d.]+)\\;version:\\1"
        },
        "html": "<a[^>]+>Powered by uKnowva</a>",
        "icon": "uKnowva.png",
        "implies": "PHP",
        "meta": {
          "generator": "uKnowva (?: ([\\d.]+))?\\;version:\\1"
        },
        "scripts": "/media/conv/js/jquery\\.js",
        "website": "https://uknowva.com"
      },
      "uPortal": {
        "cats": [
          21
        ],
        "description": "uPortal is an open source enterprise portal framework built by and for higher education institutions.",
        "icon": "uPortal.png",
        "implies": "Java",
        "js": {
          "uportal": ""
        },
        "meta": {
          "description": " uPortal "
        },
        "oss": true,
        "website": "https://www.apereo.org/projects/uportal"
      },
      "uRemediate": {
        "cats": [
          68
        ],
        "description": "uRemediate provides web accessibility testing tools.",
        "icon": "User1st.png",
        "scripts": "fecdn\\.user1st\\.info/Loader/head",
        "website": "https://www.user1st.com/uremediate/"
      },
      "user.com": {
        "cats": [
          10
        ],
        "html": "<div[^>]+/id=\"ue_widget\"",
        "icon": "user.com.svg",
        "js": {
          "UserEngage": ""
        },
        "website": "https://user.com"
      },
      "vBulletin": {
        "cats": [
          2
        ],
        "cookies": {
          "bblastactivity": "",
          "bblastvisit": "",
          "bbsessionhash": ""
        },
        "cpe": "cpe:/a:vbulletin:vbulletin",
        "description": "vBulletin is tool that is used to create and manage online forums or discussion boards. It is written in PHP and uses a MySQL database server.",
        "html": "<div id=\"copyright\">Powered by vBulletin",
        "icon": "vBulletin.png",
        "implies": "PHP",
        "js": {
          "vBulletin": ""
        },
        "meta": {
          "generator": "vBulletin ?([\\d.]+)?\\;version:\\1"
        },
        "website": "https://www.vbulletin.com"
      },
      "vibecommerce": {
        "cats": [
          6
        ],
        "excludes": "PrestaShop",
        "icon": "vibecommerce.png",
        "implies": "PHP",
        "meta": {
          "designer": "vibecommerce",
          "generator": "vibecommerce"
        },
        "website": "http://vibecommerce.com.br"
      },
      "webEdition": {
        "cats": [
          1
        ],
        "cpe": "cpe:/a:webedition:webedition_cms",
        "icon": "webEdition.png",
        "meta": {
          "DC.title": "webEdition",
          "generator": "webEdition"
        },
        "website": "http://webedition.de/en"
      },
      "webpack": {
        "cats": [
          19
        ],
        "description": "Webpack is an open-source JavaScript module bundler.",
        "icon": "webpack.svg",
        "js": {
          "webpackJsonp": ""
        },
        "website": "https://webpack.js.org/"
      },
      "wisyCMS": {
        "cats": [
          1
        ],
        "icon": "wisyCMS.svg",
        "meta": {
          "generator": "^wisy CMS[ v]{0,3}([0-9.,]*)\\;version:\\1"
        },
        "website": "https://wisy.3we.de"
      },
      "wpBakery": {
        "cats": [
          51
        ],
        "description": "WPBakery is a drag and drop visual page builder plugin for WordPress.",
        "icon": "wpBakery.svg",
        "implies": [
          "WordPress",
          "PHP"
        ],
        "meta": {
          "generator": "WPBakery"
        },
        "pricing": [
          "low",
          "onetime"
        ],
        "website": "https://wpbakery.com"
      },
      "wpCache": {
        "cats": [
          23
        ],
        "description": "WPCache is a static caching plugin for WordPress.",
        "headers": {
          "X-Powered-By": "wpCache(?:/([\\d.]+))?\\;version:\\1"
        },
        "html": "<!--[^>]+wpCache",
        "icon": "wpCache.png",
        "implies": [
          "WordPress",
          "PHP"
        ],
        "meta": {
          "generator": "wpCache",
          "keywords": "wpCache"
        },
        "url": "^https?://[^/]+\\.wpcache\\.co",
        "website": "https://wpcache.co"
      },
      "xCharts": {
        "cats": [
          25
        ],
        "html": "<link[^>]* href=\"[^\"]*xcharts(?:\\.min)?\\.css",
        "implies": "D3",
        "js": {
          "xChart": ""
        },
        "scripts": "xcharts\\.js",
        "website": "https://tenxer.github.io/xcharts/"
      },
      "xtCommerce": {
        "cats": [
          6
        ],
        "html": "<div class=\"copyright\">[^<]+<a[^>]+>xt:Commerce",
        "icon": "xtCommerce.png",
        "meta": {
          "generator": "xt:Commerce"
        },
        "website": "https://www.xt-commerce.com"
      }
    }
  }`
