

## 空间相关
### 安装插件
```sql
CREATE EXTENSION postgis;
```

### 建表语句
```sql
CREATE TABLE airports (
  code VARCHAR(3),
  geog GEOGRAPHY(Point)
);
 
INSERT INTO airports VALUES ('LAX', 'POINT(-118.4079 33.9434)');
INSERT INTO airports VALUES ('CDG', 'POINT(2.5559 49.0083)');
INSERT INTO airports VALUES ('KEF', 'POINT(-22.6056 63.9850)');
```

### 函数
```text
ST_AsText(geography) returns text
ST_GeographyFromText(text) returns geography
ST_AsBinary(geography) returns bytea
ST_GeogFromWKB(bytea) returns geography
ST_AsSVG(geography) returns text
ST_AsGML(geography) returns text
ST_AsKML(geography) returns text
ST_AsGeoJson(geography) returns text
ST_Distance(geography, geography) returns double
ST_DWithin(geography, geography, float8) returns boolean
ST_Area(geography) returns double
ST_Length(geography) returns double
ST_Covers(geography, geography) returns boolean
ST_CoveredBy(geography, geography) returns boolean
ST_Intersects(geography, geography) returns boolean
ST_Buffer(geography, float8) returns geography
ST_Intersection(geography, geography) returns geography
```