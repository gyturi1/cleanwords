# cleanwords

Cleaning, filtering the [hungarian webcorpus](ftp://ftp.mokk.bme.hu/Language/Hungarian/Freq/Web2.2/web2.2-freq-sorted.txt.gz) collecting the possible 5 letter words. It will contains false positive words, due to the specialty of double or tripple charachter letters in hungarian. So for example this tool classifies "a r á ny" as a possible 5 letter word, although "a r á n y" is not a valid spelling.

Further work can be done to eliminate false positives: does not checked but using the morphologically annotated webcorpus. Maybe the annotations contains the letter counts.

----------

## Prerequisite

**create all.txt**

`curl ftp://ftp.mokk.bme.hu/Language/Hungarian/Freq/Web2.2/web2.2-freq-sorted.txt.gz -o web2.2-freq-sorted.txt.gz && gunzip web2.2-freq-sorted.txt.gz && cat web2.2-freq-sorted.txt | awk '{print $1}' > all.txt && rm web2.2-freq-sorted.txt`

----------
## Running

`go run clean.go`

The input file is all.txt, it is in ISO-8859-2 encoding.

The output is possibleFiveLetterWords.txt with UTF-8 encoding.