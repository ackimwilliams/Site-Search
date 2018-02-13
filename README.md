# Site-Searcher
## Notes

Fetcher package for fetching websites and the main package.

Noteworthy:
- some urls took unusually long to respond to requests and so a *60 sec timeout was created to ignore these. *

Tech debt:
- There are a few `@todo:` comments throughout which mentions updates/improvements such as logging, creating multiple error types, etc.
- Replacing the rudimentary parser function with a tokenizer
- Add activity logging using `https://golang.org/pkg/log/`


Command line output:

```txt
$ go build -i

Beginning program!
	Fetching list of urls... done.

	Crawling urls...
Worker ID = 18; site = vimeo.com/; time to complete = 319.2256ms
Worker ID = 16; site = apple.com/; time to complete = 374.2758ms
Worker ID = 1; site = twitter.com/; time to complete = 384.2877ms
Worker ID = 15; site = w3.org/; time to complete = 402.2964ms
Worker ID = 10; site = wordpress.com/; time to complete = 403.2977ms
Worker ID = 2; site = google.com/; time to complete = 407.2874ms
Worker ID = 14; site = tumblr.com/; time to complete = 465.3292ms
Worker ID = 4; site = wordpress.org/; time to complete = 469.3308ms
Worker ID = 17; site = myspace.com/; time to complete = 474.348ms
Worker ID = 7; site = wikipedia.org/; time to complete = 605.4408ms
Worker ID = 5; site = adobe.com/; time to complete = 663.4769ms
Worker ID = 8; site = linkedin.com/; time to complete = 731.534ms
Worker ID = 12; site = flickr.com/; time to complete = 756.5473ms
Worker ID = 0; site = facebook.com/; time to complete = 784.5679ms
Worker ID = 17; site = miibeian.gov.cn/; time to complete = 316.2135ms
Worker ID = 14; site = statcounter.com/; time to complete = 340.2542ms
Worker ID = 6; site = blogspot.com/; time to complete = 841.5936ms
Worker ID = 15; site = baidu.com/; time to complete = 558.3823ms
Worker ID = 10; site = stumbleupon.com/; time to complete = 559.3878ms
Worker ID = 13; site = pinterest.com/; time to complete = 1.0067322s
Worker ID = 1; site = digg.com/; time to complete = 636.4519ms
Worker ID = 4; site = feedburner.com/; time to complete = 586.4337ms
Worker ID = 5; site = nytimes.com/; time to complete = 399.2924ms
Worker ID = 8; site = reddit.com/; time to complete = 335.2221ms
Worker ID = 9; site = yahoo.com/; time to complete = 1.0967769s
Worker ID = 11; site = amazon.com/; time to complete = 1.0987775s
Worker ID = 2; site = addthis.com/; time to complete = 699.5072ms
Worker ID = 14; site = msn.com/; time to complete = 385.257ms
Worker ID = 19; site = microsoft.com/; time to complete = 1.2638946s
Worker ID = 16; site = qq.com/; time to complete = 906.6463ms
Worker ID = 1; site = icio.us/; time to complete = 270.1723ms
Worker ID = 17; site = blogger.com/; time to complete = 504.3663ms
Worker ID = 11; site = t.co/; time to complete = 262.2044ms
Worker ID = 15; site = goo.gl/; time to complete = 454.3508ms
Worker ID = 1; site = sourceforge.net/; time to complete = 243.1932ms
Worker ID = 13; site = gov.uk/; time to complete = 526.3531ms
Worker ID = 5; site = cnn.com/; time to complete = 478.3406ms
Worker ID = 9; site = google.de/; time to complete = 478.3365ms
Worker ID = 10; site = instagram.com/; time to complete = 748.5129ms
Worker ID = 16; site = jimdo.com/; time to complete = 531.3672ms
Worker ID = 1; site = google.co.jp/; time to complete = 283.178ms
Worker ID = 14; site = imdb.com/; time to complete = 664.4724ms
Worker ID = 11; site = tinyurl.com/; time to complete = 497.3439ms
Worker ID = 8; site = webs.com/; time to complete = 830.6057ms
Worker ID = 5; site = free.fr/; time to complete = 430.326ms
Worker ID = 12; site = weebly.com/; time to complete = 1.3599698s
Worker ID = 18; site = youtu.be/; time to complete = 1.869365s
Worker ID = 8; site = hugedomains.com/; time to complete = 333.2355ms
Worker ID = 14; site = technorati.com/; time to complete = 379.3115ms
Worker ID = 5; site = about.com/; time to complete = 274.1948ms
Worker ID = 9; site = joomla.org/; time to complete = 694.5137ms
Worker ID = 0; site = bbc.co.uk/; time to complete = 1.5290878s
Worker ID = 4; site = yandex.ru/; time to complete = 1.2569035s
Worker ID = 13; site = fc2.com/; time to complete = 785.5733ms
Worker ID = 17; site = go.com/; time to complete = 1.0657745s
Worker ID = 10; site = creativecommons.org/; time to complete = 699.5224ms
Worker ID = 3; site = youtube.com/; time to complete = 2.4547699s
Worker ID = 2; site = livejournal.com/; time to complete = 1.3519737s
Worker ID = 19; site = mail.ru/; time to complete = 1.2178929s
Worker ID = 5; site = google.co.uk/; time to complete = 246.1508ms
Worker ID = 8; site = nih.gov/; time to complete = 418.3044ms
Worker ID = 17; site = ameblo.jp/; time to complete = 370.341ms
Worker ID = 15; site = vk.com/; time to complete = 1.4100651s
Worker ID = 12; site = theguardian.com/; time to complete = 736.5993ms
Worker ID = 9; site = mozilla.org/; time to complete = 627.5294ms
Worker ID = 19; site = bing.com/; time to complete = 418.3717ms
Worker ID = 3; site = europa.eu/; time to complete = 473.3985ms
Worker ID = 1; site = networkadvertising.org/; time to complete = 1.1319011s
Worker ID = 10; site = wsj.com/; time to complete = 541.4502ms
Worker ID = 17; site = tripod.com/; time to complete = 238.1551ms
Worker ID = 19; site = geocities.com/; time to complete = 103.0927ms
Worker ID = 16; site = typepad.com/; time to complete = 1.1989462s
Worker ID = 14; site = huffingtonpost.com/; time to complete = 928.723ms
Worker ID = 11; site = sina.com.cn/; time to complete = 1.3360342s
Worker ID = 13; site = ebay.com/; time to complete = 893.7104ms
Worker ID = 12; site = issuu.com/; time to complete = 376.2776ms
Worker ID = 9; site = gnu.org/; time to complete = 334.2389ms
Worker ID = 4; site = aol.com/; time to complete = 954.7393ms
Worker ID = 1; site = wix.com/; time to complete = 441.3233ms
Worker ID = 17; site = washingtonpost.com/; time to complete = 464.3393ms
Worker ID = 12; site = clickbank.net/; time to complete = 308.2147ms
Worker ID = 8; site = guardian.co.uk/; time to complete = 907.7244ms
Worker ID = 16; site = reuters.com/; time to complete = 556.3969ms
Worker ID = 10; site = mapquest.com/; time to complete = 639.4624ms
Worker ID = 19; site = homestead.com/; time to complete = 599.3943ms
Worker ID = 11; site = photobucket.com/; time to complete = 483.337ms
Worker ID = 13; site = forbes.com/; time to complete = 485.3524ms
Worker ID = 18; site = yahoo.co.jp/; time to complete = 1.5221434s
Worker ID = 15; site = godaddy.com/; time to complete = 889.6287ms
Worker ID = 4; site = etsy.com/; time to complete = 465.339ms
Worker ID = 19; site = posterous.com/; time to complete = 221.1954ms
Worker ID = 8; site = usatoday.com/; time to complete = 282.1943ms
Worker ID = 16; site = yelp.com/; time to complete = 333.2749ms
Worker ID = 17; site = dailymotion.com/; time to complete = 522.4132ms
Worker ID = 18; site = google.fr/; time to complete = 274.224ms
Worker ID = 12; site = soundcloud.com/; time to complete = 581.4444ms
Worker ID = 5; site = rambler.ru/; time to complete = 1.6843125s
Worker ID = 15; site = constantcontact.com/; time to complete = 493.404ms
Worker ID = 8; site = latimes.com/; time to complete = 371.296ms
Worker ID = 11; site = telegraph.co.uk/; time to complete = 583.4552ms
Worker ID = 15; site = php.net/; time to complete = 175.1111ms
Worker ID = 5; site = miitbeian.gov.cn/; time to complete = 257.1814ms
Worker ID = 2; site = taobao.com/; time to complete = 1.999525s
Worker ID = 1; site = amazon.co.uk/; time to complete = 1.1108282s
Worker ID = 19; site = phpbb.com/; time to complete = 698.5067ms
Worker ID = 11; site = bbb.org/; time to complete = 349.2468ms
Worker ID = 10; site = cnet.com/; time to complete = 1.0427707s
Worker ID = 13; site = archive.org/; time to complete = 935.6953ms
Worker ID = 12; site = opera.com/; time to complete = 564.4083ms
Worker ID = 14; site = 163.com/; time to complete = 1.5211142s
Worker ID = 8; site = scribd.com/; time to complete = 506.3468ms
Worker ID = 4; site = phoca.cz/; time to complete = 1.0597821s
Worker ID = 1; site = cdc.gov/; time to complete = 321.2135ms
Worker ID = 4; site = time.com/; time to complete = 122.087ms
Worker ID = 2; site = dailymail.co.uk/; time to complete = 478.3361ms
Worker ID = 11; site = wikimedia.org/; time to complete = 331.2345ms
Worker ID = 1; site = google.it/; time to complete = 224.1595ms
Worker ID = 13; site = mit.edu/; time to complete = 486.3438ms
Worker ID = 0; site = 51.la/; time to complete = 2.8201142s
Worker ID = 2; site = live.com/; time to complete = 247.1767ms
Worker ID = 4; site = stanford.edu/; time to complete = 360.254ms
Worker ID = 18; site = amazon.de/; time to complete = 1.293917s
Worker ID = 8; site = addtoany.com/; time to complete = 559.3964ms
Worker ID = 2; site = histats.com/; time to complete = 129.0885ms
Worker ID = 1; site = squidoo.com/; time to complete = 301.226ms
Worker ID = 5; site = ning.com/; time to complete = 930.6571ms
Worker ID = 13; site = harvard.edu/; time to complete = 262.1839ms
Worker ID = 3; site = slideshare.net/; time to complete = 2.4637855s
Worker ID = 11; site = alibaba.com/; time to complete = 615.4355ms
Worker ID = 9; site = weibo.com/; time to complete = 2.3627123s
Worker ID = 0; site = gravatar.com/; time to complete = 465.3288ms
Worker ID = 18; site = npr.org/; time to complete = 328.2384ms
Worker ID = 4; site = nasa.gov/; time to complete = 410.2864ms
Worker ID = 1; site = wired.com/; time to complete = 357.2393ms
Worker ID = 18; site = blinklist.com/; time to complete = 105.078ms
Worker ID = 3; site = blog.com/; time to complete = 402.2951ms
Worker ID = 15; site = parallels.com/; time to complete = 1.4200045s
Worker ID = 9; site = bloomberg.com/; time to complete = 231.1622ms
Worker ID = 19; site = sohu.com/; time to complete = 1.3329304s
Worker ID = 14; site = altervista.org/; time to complete = 1.2519215s
Worker ID = 4; site = imageshack.us/; time to complete = 278.2572ms
Worker ID = 11; site = amazonaws.com/; time to complete = 418.3537ms
Worker ID = 2; site = eventbrite.com/; time to complete = 747.5851ms
Worker ID = 9; site = google.es/; time to complete = 248.2113ms
Worker ID = 19; site = ocn.ne.jp/; time to complete = 243.2269ms
Worker ID = 1; site = kickstarter.com/; time to complete = 432.366ms
Worker ID = 13; site = nbcnews.com/; time to complete = 816.6344ms
Worker ID = 11; site = google.ca/; time to complete = 230.1493ms
Worker ID = 19; site = pbs.org/; time to complete = 183.1217ms
Worker ID = 16; site = erecht24.de/; time to complete = 2.439769s
Worker ID = 12; site = sakura.ne.jp/; time to complete = 1.6892301s
Worker ID = 11; site = prweb.com/; time to complete = 200.1414ms
Worker ID = 9; site = weather.com/; time to complete = 486.3545ms
Worker ID = 15; site = angelfire.com/; time to complete = 769.5899ms
Worker ID = 5; site = amazon.co.jp/; time to complete = 1.3219787s
Worker ID = 11; site = noaa.gov/; time to complete = 291.2054ms
Worker ID = 8; site = ca.gov/; time to complete = 1.4710941s
Worker ID = 13; site = cpanel.net/; time to complete = 651.4499ms
Worker ID = 4; site = dedecms.com/; time to complete = 902.6435ms
Worker ID = 4; site = eepurl.com/; time to complete = 80.0376ms
Worker ID = 1; site = ibm.com/; time to complete = 812.5593ms
Worker ID = 16; site = barnesandnoble.com/; time to complete = 622.4431ms
Worker ID = 10; site = deviantart.com/; time to complete = 2.435778s
Worker ID = 14; site = overblog.com/; time to complete = 1.1418154s
Worker ID = 8; site = foxnews.com/; time to complete = 388.2643ms
Worker ID = 14; site = geocities.jp/; time to complete = 206.2293ms
Worker ID = 8; site = loc.gov/; time to complete = 335.3128ms
Worker ID = 18; site = hatena.ne.jp/; time to complete = 1.843407s
Worker ID = 19; site = bandcamp.com/; time to complete = 1.2749839s
Worker ID = 0; site = narod.ru/; time to complete = 2.0175418s
Worker ID = 16; site = newsvine.com/; time to complete = 679.5379ms
Worker ID = 11; site = cbsnews.com/; time to complete = 946.7394ms
Worker ID = 12; site = mozilla.com/; time to complete = 1.3190026s
Worker ID = 14; site = yolasite.com/; time to complete = 475.3141ms
Worker ID = 8; site = apache.org/; time to complete = 287.2112ms
Worker ID = 18; site = mashable.com/; time to complete = 236.1668ms
Worker ID = 0; site = nationalgeographic.com/; time to complete = 230.1646ms
Worker ID = 12; site = ted.com/; time to complete = 302.2118ms
Worker ID = 14; site = sfgate.com/; time to complete = 355.251ms
Worker ID = 16; site = whitehouse.gov/; time to complete = 475.3358ms
Worker ID = 1; site = berkeley.edu/; time to complete = 1.1979162s
Worker ID = 19; site = usda.gov/; time to complete = 619.4245ms
Worker ID = 10; site = bluehost.com/; time to complete = 1.1548749s
Worker ID = 11; site = tripadvisor.com/; time to complete = 600.4235ms
Worker ID = 18; site = epa.gov/; time to complete = 497.3497ms
Worker ID = 8; site = biglobe.ne.jp/; time to complete = 619.4232ms
Worker ID = 12; site = oracle.com/; time to complete = 391.2772ms
Worker ID = 4; site = businessweek.com/; time to complete = 1.4430887s
Worker ID = 16; site = examiner.com/; time to complete = 305.2674ms
Worker ID = 11; site = disqus.com/; time to complete = 229.215ms
Worker ID = 1; site = cornell.edu/; time to complete = 351.2954ms
Worker ID = 3; site = nifty.com/; time to complete = 2.7550936s
Worker ID = 11; site = techcrunch.com/; time to complete = 232.1462ms
Worker ID = 13; site = discuz.net/; time to complete = 1.9364741s
Worker ID = 16; site = boston.com/; time to complete = 376.2501ms
Worker ID = 10; site = nps.gov/; time to complete = 578.4449ms
Worker ID = 19; site = hp.com/; time to complete = 704.5573ms
Worker ID = 18; site = alexa.com/; time to complete = 611.4931ms
Worker ID = 1; site = un.org/; time to complete = 438.2996ms
Worker ID = 12; site = house.gov/; time to complete = 582.4609ms
Worker ID = 2; site = a8.net/; time to complete = 2.9521812s
Worker ID = 14; site = seesaa.net/; time to complete = 1.0117486s
Worker ID = 0; site = vkontakte.ru/; time to complete = 1.3029552s
Worker ID = 19; site = independent.co.uk/; time to complete = 283.1751ms
Worker ID = 13; site = freewebs.com/; time to complete = 376.2665ms
Worker ID = 12; site = google.nl/; time to complete = 205.1299ms
Worker ID = 3; site = squarespace.com/; time to complete = 730.5071ms
Worker ID = 4; site = sphinn.com/; time to complete = 950.7278ms
Worker ID = 16; site = ezinearticles.com/; time to complete = 561.4155ms
Worker ID = 14; site = imgur.com/; time to complete = 282.2087ms
Worker ID = 8; site = mysql.com/; time to complete = 1.1318338s
Worker ID = 14; site = people.com.cn/; time to complete = 127.1022ms
Worker ID = 12; site = bizjournals.com/; time to complete = 397.2848ms
Worker ID = 0; site = irs.gov/; time to complete = 435.3363ms
Worker ID = 13; site = wunderground.com/; time to complete = 467.3589ms
Worker ID = 18; site = mediafire.com/; time to complete = 786.5524ms
Worker ID = 12; site = cbslocal.com/; time to complete = 171.8731ms
Worker ID = 5; site = xrea.com/; time to complete = 3.0550506s
Worker ID = 16; site = cloudflare.com/; time to complete = 480.092ms
Worker ID = 3; site = who.int/; time to complete = 582.1986ms
Worker ID = 13; site = opensource.org/; time to complete = 252.9211ms
Worker ID = 0; site = ycombinator.com/; time to complete = 330.9766ms
Worker ID = 2; site = reverbnation.com/; time to complete = 907.4144ms
Worker ID = 16; site = businessinsider.com/; time to complete = 180.1267ms
Worker ID = 8; site = ustream.tv/; time to complete = 771.3204ms
Worker ID = 14; site = senate.gov/; time to complete = 751.2923ms
Worker ID = 4; site = soup.io/; time to complete = 970.4481ms
Worker ID = 11; site = icq.com/; time to complete = 1.6179173s
Worker ID = 16; site = wikia.com/; time to complete = 328.2314ms
Worker ID = 18; site = spiegel.de/; time to complete = 767.2926ms
Worker ID = 2; site = skype.com/; time to complete = 550.3882ms
Worker ID = 8; site = about.me/; time to complete = 278.1952ms
Worker ID = 1; site = xinhuanet.com/; time to complete = 1.6279207s
Worker ID = 11; site = gmpg.org/; time to complete = 254.1787ms
Worker ID = 5; site = nature.com/; time to complete = 889.616ms
Worker ID = 13; site = last.fm/; time to complete = 817.5783ms
Worker ID = 8; site = cbc.ca/; time to complete = 229.162ms
Worker ID = 1; site = umich.edu/; time to complete = 226.1601ms
Worker ID = 2; site = github.com/; time to complete = 334.2487ms
Worker ID = 3; site = drupal.org/; time to complete = 960.6643ms
Worker ID = 14; site = webmd.com/; time to complete = 623.4396ms
Worker ID = 10; site = ucoz.ru/; time to complete = 2.1322949s
Worker ID = 13; site = google.com.br/; time to complete = 251.1767ms
Worker ID = 14; site = topsy.com/; time to complete = 50.0558ms
Worker ID = 1; site = mac.com/; time to complete = 376.31ms
Worker ID = 13; site = google.cn/; time to complete = 247.2232ms
Worker ID = 3; site = discovery.com/; time to complete = 434.3415ms
Worker ID = 11; site = jugem.jp/; time to complete = 826.6193ms
Worker ID = 19; site = webnode.com/; time to complete = 2.2604042s
Worker ID = 2; site = wiley.com/; time to complete = 765.5621ms
Worker ID = 11; site = google.pl/; time to complete = 216.1518ms
Worker ID = 4; site = youku.com/; time to complete = 1.5070895s
Worker ID = 19; site = prnewswire.com/; time to complete = 422.3084ms
Worker ID = 2; site = ft.com/; time to complete = 364.2711ms
Worker ID = 1; site = moonfruit.com/; time to complete = 786.5459ms
Worker ID = 12; site = oaic.gov.au/; time to complete = 2.2546385s
Worker ID = 0; site = privacy.gov.au/; time to complete = 2.121546s
Worker ID = 4; site = behance.net/; time to complete = 222.1696ms
Worker ID = 14; site = surveymonkey.com/; time to complete = 1.0837925s
Worker ID = 18; site = redcross.org/; time to complete = 1.7913394s
Worker ID = 10; site = paypal.com/; time to complete = 1.3440295s
Worker ID = 13; site = dropbox.com/; time to complete = 1.1048051s
Worker ID = 19; site = goodreads.com/; time to complete = 572.46ms
Worker ID = 12; site = marketwatch.com/; time to complete = 529.3953ms
Worker ID = 2; site = netvibes.com/; time to complete = 786.6625ms
Worker ID = 5; site = shinystat.com/; time to complete = 2.0976254s
Worker ID = 13; site = ftc.gov/; time to complete = 428.3923ms
Worker ID = 14; site = state.gov/; time to complete = 743.6151ms
Worker ID = 0; site = ed.gov/; time to complete = 823.6794ms
Worker ID = 12; site = quantcast.com/; time to complete = 375.3547ms
Worker ID = 4; site = networksolutions.com/; time to complete = 884.7322ms
Worker ID = 5; site = nydailynews.com/; time to complete = 250.1898ms
Worker ID = 2; site = economist.com/; time to complete = 383.2729ms
Worker ID = 8; site = ifeng.com/; time to complete = 2.3668321s
Worker ID = 19; site = census.gov/; time to complete = 840.6563ms
Worker ID = 2; site = theatlantic.com/; time to complete = 242.1874ms
Worker ID = 8; site = google.com.au/; time to complete = 226.147ms
Worker ID = 5; site = chicagotribune.com/; time to complete = 407.2771ms
Worker ID = 11; site = uol.com.br/; time to complete = 1.7763841s
Worker ID = 0; site = ow.ly/; time to complete = 627.5183ms
Worker ID = 14; site = cafepress.com/; time to complete = 858.6797ms
Worker ID = 4; site = netscape.com/; time to complete = 733.5829ms
Worker ID = 0; site = friendfeed.com/; time to complete = 183.1866ms
Worker ID = 5; site = pagespersoorange.fr/; time to complete = 305.2916ms
Worker ID = 3; site = exblog.jp/; time to complete = 2.4569208s
Worker ID = 13; site = zdnet.com/; time to complete = 991.7579ms
Worker ID = 2; site = skyrock.com/; time to complete = 416.342ms
Worker ID = 4; site = patch.com/; time to complete = 253.2541ms
Worker ID = 8; site = listmanage.com/; time to complete = 638.5727ms
Worker ID = 0; site = upenn.edu/; time to complete = 358.2611ms
Worker ID = 12; site = meetup.com/; time to complete = 1.1459307s
Worker ID = 2; site = slashdot.org/; time to complete = 491.5577ms
Worker ID = 19; site = 1688.com/; time to complete = 1.0289821s
Worker ID = 10; site = liveinternet.ru/; time to complete = 2.007738s
Worker ID = 8; site = columbia.edu/; time to complete = 315.3796ms
Worker ID = 13; site = com.com/; time to complete = 652.6729ms
Worker ID = 8; site = marriott.com/; time to complete = 205.189ms
Worker ID = 11; site = cdbaby.com/; time to complete = 1.1771371s
Worker ID = 10; site = yale.edu/; time to complete = 296.2609ms
Worker ID = 0; site = nhs.uk/; time to complete = 578.6047ms
Worker ID = 4; site = washington.edu/; time to complete = 702.6693ms
Worker ID = 19; site = utexas.edu/; time to complete = 335.2882ms
Worker ID = 3; site = diigo.com/; time to complete = 915.8951ms
Worker ID = 14; site = ehow.com/; time to complete = 985.945ms
Worker ID = 5; site = engadget.com/; time to complete = 950.9188ms
Worker ID = 2; site = elegantthemes.com/; time to complete = 389.3136ms
Worker ID = 1; site = auda.org.au/; time to complete = 2.7233587s
Worker ID = 13; site = bigcartel.com/; time to complete = 300.2489ms
Worker ID = 4; site = hubpages.com/; time to complete = 190.1375ms
Worker ID = 19; site = slate.com/; time to complete = 371.4007ms
Worker ID = 12; site = abc.net.au/; time to complete = 955.994ms
Worker ID = 14; site = umn.edu/; time to complete = 427.432ms
Worker ID = 2; site = sonet.ne.jp/; time to complete = 425.4471ms
Worker ID = 10; site = jigsy.com/; time to complete = 674.6367ms
Worker ID = 8; site = ucla.edu/; time to complete = 744.6611ms
Worker ID = 8; site = flavors.me/; time to complete = 80.068ms
Worker ID = 2; site = usnews.com/; time to complete = 320.2154ms
Worker ID = 4; site = howstuffworks.com/; time to complete = 616.5719ms
Worker ID = 13; site = cargocollective.com/; time to complete = 859.7558ms
Worker ID = 3; site = purevolume.com/; time to complete = 997.8348ms
Worker ID = 1; site = wikispaces.com/; time to complete = 1.1799757s
Worker ID = 4; site = usa.gov/; time to complete = 451.3099ms
Worker ID = 17; site = rakuten.co.jp/; time to complete = 11.9699627s
Worker ID = 3; site = lycos.com/; time to complete = 318.2242ms
Worker ID = 14; site = jiathis.com/; time to complete = 937.7013ms
Worker ID = 3; site = example.com/; time to complete = 65.1402ms
Worker ID = 17; site = state.tx.us/; time to complete = 102.1686ms
Worker ID = 14; site = shareasale.com/; time to complete = 112.0894ms
Worker ID = 11; site = usgs.gov/; time to complete = 1.642337s
Worker ID = 10; site = xing.com/; time to complete = 1.1398346s
Worker ID = 3; site = biblegateway.com/; time to complete = 388.2421ms
Worker ID = 3; site = g.co/; time to complete = 83.0576ms
Worker ID = 14; site = yellowbook.com/; time to complete = 401.2942ms
Worker ID = 13; site = edublogs.org/; time to complete = 1.0247777s
Worker ID = 1; site = wisc.edu/; time to complete = 737.5686ms
Worker ID = 3; site = dion.ne.jp/; time to complete = 236.1988ms
Worker ID = 17; site = is.gd/; time to complete = 729.5124ms
Worker ID = 14; site = dagondesign.com/; time to complete = 317.2254ms
Worker ID = 13; site = theglobeandmail.com/; time to complete = 374.2523ms
Worker ID = 3; site = storify.com/; time to complete = 217.1248ms
Worker ID = 0; site = hexun.com/; time to complete = 2.4108808s
Worker ID = 8; site = desdev.cn/; time to complete = 1.6602172s
Worker ID = 17; site = salon.com/; time to complete = 312.1897ms
Worker ID = 11; site = samsung.com/; time to complete = 860.6204ms
Worker ID = 10; site = businesswire.com/; time to complete = 745.5274ms
Worker ID = 13; site = gizmodo.com/; time to complete = 284.2003ms
Worker ID = 4; site = thetimes.co.uk/; time to complete = 1.3910424s
Worker ID = 8; site = reference.com/; time to complete = 263.2012ms
Worker ID = 17; site = sun.com/; time to complete = 290.2049ms
Worker ID = 11; site = unicef.org/; time to complete = 326.241ms
Worker ID = 2; site = hc360.com/; time to complete = 2.0805217s
Worker ID = 10; site = devhub.com/; time to complete = 550.3923ms
Worker ID = 3; site = psu.edu/; time to complete = 781.5491ms
Worker ID = 0; site = smh.com.au/; time to complete = 723.5193ms
Worker ID = 2; site = cocolognifty.com/; time to complete = 364.2715ms
Worker ID = 19; site = plala.or.jp/; time to complete = 2.9171239s
Worker ID = 13; site = artisteer.com/; time to complete = 738.5333ms
Worker ID = 12; site = infoseek.co.jp/; time to complete = 3.05623s
Worker ID = 8; site = istockphoto.com/; time to complete = 839.5987ms
Worker ID = 17; site = answers.com/; time to complete = 855.6153ms
Worker ID = 0; site = intel.com/; time to complete = 534.4105ms
Worker ID = 13; site = sciencedaily.com/; time to complete = 319.256ms
Worker ID = 11; site = trellian.com/; time to complete = 939.7175ms
Worker ID = 1; site = booking.com/; time to complete = 1.780313s
Worker ID = 8; site = ask.com/; time to complete = 340.2682ms
Worker ID = 19; site = ebay.co.uk/; time to complete = 721.5494ms
Worker ID = 3; site = tonline.de/; time to complete = 983.7375ms
Worker ID = 1; site = deliciousdays.com/; time to complete = 363.246ms
Worker ID = 12; site = paginegialle.it/; time to complete = 712.5548ms
Worker ID = 8; site = smugmug.com/; time to complete = 321.2475ms
Worker ID = 10; site = i2i.jp/; time to complete = 1.1448464s
Worker ID = 13; site = timesonline.co.uk/; time to complete = 647.4659ms
Worker ID = 0; site = canalblog.com/; time to complete = 787.5563ms
Worker ID = 17; site = springer.com/; time to complete = 870.6447ms
Worker ID = 10; site = twitpic.com/; time to complete = 277.2173ms
Worker ID = 13; site = ovh.net/; time to complete = 304.2083ms
Worker ID = 1; site = cmu.edu/; time to complete = 549.3848ms
Worker ID = 10; site = google.ru/; time to complete = 248.1539ms
Worker ID = 1; site = newyorker.com/; time to complete = 384.3115ms
Worker ID = 2; site = 1und1.de/; time to complete = 1.8413596s
Worker ID = 19; site = wufoo.com/; time to complete = 1.0717734s
Worker ID = 10; site = blogs.com/; time to complete = 506.3735ms
Worker ID = 2; site = hibu.com/; time to complete = 240.1544ms
Worker ID = 10; site = hhs.gov/; time to complete = 330.2307ms
Worker ID = 2; site = dmoz.org/; time to complete = 320.2244ms
Worker ID = 1; site = sciencedirect.com/; time to complete = 611.4048ms
Worker ID = 8; site = odnoklassniki.ru/; time to complete = 1.5090491s
Worker ID = 7; site = delicious.com/; time to complete = 20.0339982s
Worker ID = 1; site = google.com.hk/; time to complete = 361.3272ms
Worker ID = 4; site = unesco.org/; time to complete = 3.4675855s
Worker ID = 17; site = naver.com/; time to complete = 1.6602578s
Worker ID = 12; site = domainmarket.com/; time to complete = 2.080534s
Worker ID = 1; site = zimbio.com/; time to complete = 330.2319ms
Worker ID = 19; site = hud.gov/; time to complete = 1.378072s
Worker ID = 3; site = globo.com/; time to complete = 2.4228224s
Worker ID = 17; site = cnbc.com/; time to complete = 455.3594ms
Worker ID = 10; site = dot.gov/; time to complete = 1.0078131s
Worker ID = 12; site = uiuc.edu/; time to complete = 410.3188ms
Worker ID = 4; site = chronoengine.com/; time to complete = 640.483ms
Worker ID = 7; site = craigslist.org/; time to complete = 751.5501ms
Worker ID = 8; site = jalbum.net/; time to complete = 1.1008757s
Worker ID = 3; site = prlog.org/; time to complete = 300.2144ms
Worker ID = 13; site = si.edu/; time to complete = 2.2006668s
Worker ID = 14; site = ucoz.com/; time to complete = 4.7425238s
Worker ID = 12; site = mtv.com/; time to complete = 327.2481ms
Worker ID = 7; site = java.com/; time to complete = 210.1504ms
Worker ID = 3; site = japanpost.jp/; time to complete = 185.1502ms
Worker ID = 19; site = symantec.com/; time to complete = 641.4408ms
Worker ID = 3; site = admin.ch/; time to complete = 127.0723ms
Worker ID = 8; site = cisco.com/; time to complete = 384.3271ms
Worker ID = 14; site = github.io/; time to complete = 274.2413ms
Worker ID = 2; site = cyberchimps.com/; time to complete = 1.7093534s
Worker ID = 6; site = macromedia.com/; time to complete = 21.1088601s
Worker ID = 4; site = webeden.co.uk/; time to complete = 624.4857ms
Worker ID = 3; site = printfriendly.com/; time to complete = 287.2568ms
Worker ID = 12; site = mayoclinic.com/; time to complete = 448.3564ms
Worker ID = 7; site = studiopress.com/; time to complete = 467.3988ms
Worker ID = 17; site = 360.cn/; time to complete = 949.7077ms
Worker ID = 0; site = home.pl/; time to complete = 3.1494072s
Worker ID = 1; site = vistaprint.com/; time to complete = 1.2579762s
Worker ID = 13; site = 4shared.com/; time to complete = 837.6359ms
Worker ID = 6; site = dell.com/; time to complete = 389.2528ms
Worker ID = 0; site = nyu.edu/; time to complete = 180.1026ms
Worker ID = 13; site = merriamwebster.com/; time to complete = 131.0911ms
Worker ID = 1; site = wp.com/; time to complete = 223.1442ms
Worker ID = 3; site = princeton.edu/; time to complete = 436.2998ms
Worker ID = 17; site = chron.com/; time to complete = 345.2431ms
Worker ID = 8; site = mlb.com/; time to complete = 744.5294ms
Worker ID = 0; site = shoppro.jp/; time to complete = 223.1627ms
Worker ID = 10; site = indiatimes.com/; time to complete = 1.4430624s
Worker ID = 7; site = comcast.net/; time to complete = 586.3931ms
Worker ID = 6; site = nba.com/; time to complete = 347.2454ms
Worker ID = 3; site = indiegogo.com/; time to complete = 242.169ms
Worker ID = 7; site = army.mil/; time to complete = 103.0709ms
Worker ID = 17; site = buzzfeed.com/; time to complete = 338.2635ms
Worker ID = 2; site = simplemachines.org/; time to complete = 992.7247ms
Worker ID = 12; site = fotki.com/; time to complete = 947.6804ms
Worker ID = 6; site = csmonitor.com/; time to complete = 348.2685ms
Worker ID = 14; site = omniture.com/; time to complete = 1.2498855s
Worker ID = 1; site = furl.net/; time to complete = 672.5183ms
Worker ID = 17; site = rediff.com/; time to complete = 371.2812ms
Worker ID = 14; site = tiny.cc/; time to complete = 221.1539ms
Worker ID = 0; site = ox.ac.uk/; time to complete = 719.5267ms
Worker ID = 6; site = va.gov/; time to complete = 341.2403ms
Worker ID = 13; site = lulu.com/; time to complete = 990.7221ms
Worker ID = 7; site = tamu.edu/; time to complete = 709.5512ms
Worker ID = 0; site = multiply.com/; time to complete = 176.1235ms
Worker ID = 10; site = mapy.cz/; time to complete = 890.6567ms
Worker ID = 12; site = yellowpages.com/; time to complete = 565.3994ms
Worker ID = 7; site = nymag.com/; time to complete = 97.0528ms
Worker ID = 14; site = oakley.com/; time to complete = 259.1844ms
Worker ID = 8; site = tuttocitta.it/; time to complete = 1.0867796s
Worker ID = 2; site = toplist.cz/; time to complete = 775.5457ms
Worker ID = 17; site = elpais.com/; time to complete = 554.3709ms
Worker ID = 14; site = fastcompany.com/; time to complete = 341.2531ms
Worker ID = 13; site = hostgator.com/; time to complete = 575.4073ms
Worker ID = 3; site = bravesites.com/; time to complete = 1.311394s
Worker ID = 8; site = earthlink.net/; time to complete = 401.7179ms
Worker ID = 0; site = fema.gov/; time to complete = 564.8316ms
Worker ID = 1; site = netlog.com/; time to complete = 1.0331738s
Worker ID = 10; site = blogtalkradio.com/; time to complete = 634.9086ms
Worker ID = 17; site = msu.edu/; time to complete = 426.7386ms
Worker ID = 2; site = vinaora.com/; time to complete = 663.9076ms
Worker ID = 13; site = ucsd.edu/; time to complete = 462.7612ms
Worker ID = 12; site = china.com.cn/; time to complete = 887.0813ms
Worker ID = 14; site = aboutads.info/; time to complete = 558.8146ms
Worker ID = 17; site = tinypic.com/; time to complete = 476.2066ms
Worker ID = 10; site = sbwire.com/; time to complete = 530.2201ms
Worker ID = 14; site = pen.io/; time to complete = 279.0701ms
Worker ID = 8; site = seattletimes.com/; time to complete = 754.4075ms
Worker ID = 7; site = unblog.fr/; time to complete = 1.2431797s
Worker ID = 0; site = dyndns.org/; time to complete = 842.4582ms
Worker ID = 14; site = scientificamerican.com/; time to complete = 270.1797ms
Worker ID = 8; site = themeforest.net/; time to complete = 300.1964ms
Worker ID = 6; site = tmall.com/; time to complete = 1.78557s
Worker ID = 12; site = walmart.com/; time to complete = 810.4122ms
Worker ID = 17; site = arizona.edu/; time to complete = 672.4621ms
Worker ID = 1; site = 123reg.co.uk/; time to complete = 1.2086852s
Worker ID = 0; site = cam.ac.uk/; time to complete = 486.343ms
Worker ID = 10; site = woothemes.com/; time to complete = 728.5107ms
Worker ID = 8; site = arstechnica.com/; time to complete = 362.2673ms
Worker ID = 7; site = spotify.com/; time to complete = 634.4405ms
Worker ID = 13; site = shutterfly.com/; time to complete = 1.0155781s
Worker ID = 12; site = illinois.edu/; time to complete = 234.1678ms
Worker ID = 2; site = acquirethisname.com/; time to complete = 1.1506971s
Worker ID = 0; site = ihg.com/; time to complete = 280.1974ms
Worker ID = 10; site = pcworld.com/; time to complete = 378.2572ms
Worker ID = 14; site = unc.edu/; time to complete = 844.6177ms
Worker ID = 3; site = sogou.com/; time to complete = 2.4889007s
Worker ID = 17; site = bloglovin.com/; time to complete = 1.6994083s
Worker ID = 1; site = nsw.gov.au/; time to complete = 2.0970162s
Worker ID = 15; site = comsenz.com/; time to complete = 21.2849195s
Worker ID = 6; site = hao123.com/; time to complete = 4.7607365s
Worker ID = 16; site = fda.gov/; time to complete = 21.058183s
Worker ID = 18; site = sitemeter.com/; time to complete = 21.0436131s
Worker ID = 5; site = bloglines.com/; time to complete = 21.0520088s
Worker ID = 19; site = virginia.edu/; time to complete = 20.4832569s
Worker ID = 11; site = de.vu/; time to complete = 30.0017908s
Worker ID = 9; site = goo.ne.jp/; time to complete = 1m0.0013129s
Worker ID = 4; site = accuweather.com/; time to complete = 1m0.0009364s
 done.

	Writing matched results to file... done.

Program exit!


Total execution time: 1m22.2905592s

Process finished with exit code 0
```


Results.txt

```txt
Search term was found in the following urls:
twitter.com/
facebook.com/
nytimes.com/
yahoo.com/
amazon.com/
cnn.com/
wsj.com/
washingtonpost.com/
usatoday.com/
latimes.com/
telegraph.co.uk/
amazon.co.uk/
bbb.org/
time.com/
amazon.de/
npr.org/
bloomberg.com/
sohu.com/
eventbrite.com/
prweb.com/
amazon.co.jp/
ca.gov/
overblog.com/
foxnews.com/
cbsnews.com/
sfgate.com/
berkeley.edu/
epa.gov/
techcrunch.com/
alexa.com/
wunderground.com/
who.int/
wikia.com/
oaic.gov.au/
privacy.gov.au/
state.gov/
theatlantic.com/
chicagotribune.com/
engadget.com/
theglobeandmail.com/
cocolognifty.com/
booking.com/
twitpic.com/
hhs.gov/
cnbc.com/
dot.gov/
buzzfeed.com/
blogtalkradio.com/
ucsd.edu/
123reg.co.uk/
hao123.com/
```

