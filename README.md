High speed scraping with bigquery
========

Overview
--------

クロールで収集した HTML をそれぞれ1行の Record として
BigQuery に export し、UDFでそのHTMLを parse & scraping すれば
速いのではないかと思いついたので実際にやってみた。

本READMEには概要のみ記載しておく。
実際の処理が知りたい場合はコードをみるべし。

### ディレクトリとファイルの構成

```
- example/          // BigQuery で実行する Query
    sample1.sql
- html2jsonl/
  - dest/           // minify & jsol encode したHTMLの出力先
  - html/           // Wikipedia から wget で取得したHTML
      sample1.html
      sample2.html
    main.go
    Makefile        // make sample1 で minify & jsonl encode される
    minifier.go
    sanitizer.go
- udf/
    global.js
    main.js
    package.json    // npm bundle で bundle.js が生成される
  README.md
```


Export html to BigQuery
--------

### Collect HTML

実際にHTMLを収集するのはちょっと面倒なので、
Wikipediaから適当に拝借することにする。

取得したHTML は `html2jsonl/html/` に配置した。

### Minify and encode to jsonl

取得した HTML から scraping には不要と思われる要素を削ぎ落とし
データ量を削減するために、golang製の sanitizer と minifier を利用する。

また BigQuery に export するには csv か jsonl のフォーマットに
encode する必要があるが、csvにするより簡単なので今回は jsonl に encode する。

以上処理を `html2jsonl/*.go` で行っている。


Define cheerio as UDF
--------

HTMLの parse と scrape には [cheerio](https://github.com/cheeriojs/cheerio)
を利用する。

### Bundle

cheerio を require し、かつ global scope に cheerio オブジェクトとして
定義したJSを browserify で bundle することで単一ファイルと動作可能な
ようにする。

今回は npm run bundle で bundle.js として生成するようにしてある。

### Upload to Cloud Storage

BigQuery は inline で定義した UDF(js) から Cloud Storage に配置した
JSをライブラリとして読み込むことができるので、
bundle した 単一ファイルの JS を Cloud Storage に upload する。


### Define UDF

定義の仕方は説明より実際のSQLを見た方がやはいと思うので、
以下に記載するのみとする。

```sql
CREATE TEMPORARY FUNCTION extractVal(htmlstr STRING)
RETURNS STRING
LANGUAGE js AS """
  var $ = cheerio.load(htmlstr)
  return $('h1').text()
"""
OPTIONS (library="gs://<your_bucket_name>/bundle.js")
;
```

Scraping
--------


```sql
CREATE TEMPORARY FUNCTION extractVal(htmlstr STRING)
RETURNS STRING
LANGUAGE js AS """
  var $ = cheerio.load(htmlstr)
  return $('h1').text()
"""
OPTIONS (library="gs://<your_bucket_name>/bundle.js")
-- bundle.js の global スコープで cheerio というオブジェクトを定義してある
;

SELECT
  extractVal(html) as val
FROM
  `<your_dataset>.sample1`
;
```


