# SubtitleAnalyzer

接口：aip_subtitle.mgaronya.com/upload

方法：POST

接收参数：一个文件

返回参数：一个subtitles表示subtitle数组，subtitle按照起始时间由小到大排序，一个errors表示报错数组，一个warnings表示警告数组。subtitle的结构为正整数id，正整数start_time和正整数end_time，字符串content。

功能：该接口将会接受一个字幕文件，然后对字幕文件进行分析。只要不是严重的信息丢失或系统错误，该接口将尽可能多的返回subtitle。损坏的subtitle将在errors中报出。可以容忍的错误将在warnings中报出。当字幕之间的起始时间与终止时间相互重叠时，该接口会调整各个字幕的起始时间与终止时间以让它们不再重叠，即使这可能会让字幕的持续时间为0。

接口：aip_subtitle.mgaronya.com/download

方法：POST

接收参数：在Body，raw格式给出json数组，要求每个元素包含有正整数id、正整数start_time、正整数end_time、字符串content。

返回参数：一个demo.srt文件，表示以json数组翻译出的srt文件。

功能：该接口将会接受一组json表示的subtitle，然后解析这些json并转为srt文件。当json数据传入有误时会报出错误。