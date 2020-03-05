mysql 8.x 오류로 install 스크립트에서 자동적으로   
함수 생성까지하는 부분은 뺏습니다.

해당 가이드에서는 ssl이 적용된 사이트는 결과를 못받을겁니다.  

mysql 8.x guide  

//함수 바인딩 방법  
//plugin_dir에 so 파일을 넣고  
//CREATE FUNCTION http_get RETURNS STRING SONAME 'http.so';  
출처 https://dev.mysql.com/doc/refman/8.0/en/keyring-udfs-general-purpose.html  

//mysql 8.x 결과가 문자열이 아닌 blob으로 나오는 부분 해결하는 법  
//SELECT CONVERT(http_get("some url") USING utf8);  
출처 https://dev.mysql.com/doc/refman/8.0/en/cast-functions.html\

mariadb, mysql 5.7는 밑의 링크에서 확인 부탁드립니다.  
https://github.com/RebirthLee/mysql_udf_http_golang

// 자세한 가이드는 원하신다면 별도로 첨부하겠습니다.
