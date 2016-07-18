## Flicker ID Generator

### create mysql table

创建下面的数据库表, 如果表命不是 flicker 需要更改程序源码.
```sql
CREATE TABLE `flicker` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stub` char(1) COLLATE ascii_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=ascii COLLATE=ascii_bin;
```
