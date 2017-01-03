CREATE TEMPORARY FUNCTION extractVal(htmlstr STRING)
RETURNS STRING
LANGUAGE js AS """
  var $ = cheerio.load(htmlstr)
  return $('h1').text()
"""
OPTIONS (library="gs://<your_bucket_name>/bundle.js")
;

SELECT
  extractVal(html) as val
FROM
  `<your_dataset>.sample1`
;
