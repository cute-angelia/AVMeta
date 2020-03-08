/*
Package media 媒体库 nfo 操作包。

media包将对番号进行搜索，确定刮削网站后执行刮削操作，
通过读取配置文件信息，确定刮削后的正确保存路径，
并在程序执行目录下创建保存路径，在刮削后下载对应封面图片，
并执行封面剪切操作，将封面及背景图片均保存到路径中，
最后生成媒体库 nfo 通用文件，并与电影一起保存到路径中。
*/
package media