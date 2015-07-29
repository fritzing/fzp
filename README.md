# [fritzing](https://github.com/fritzing) fzp specs & tools

- [Sample fzp](#sample-fzp)
- [Specification](#specification)


## Sample fzp

```
<?xml version="1.0" encoding="UTF-8"?>
<module fritzingVersion="x.y.z" moduleId="mod-id-rev" referenceFile="ref.file">
  <version>1.1</version>
  <title>11LC010</title>
  <description>some words about the part</description>
  <author>Revolution Education Ltd</author>
  <date>yyyy-mm-dd</date>
  <url>http://part.org</url>
  <label>IC</label>
  <tags>...</tags>
  <taxonomy>...</taxonomy>
  <language>...</language>
  <family>...</family>
  <properties>...</properties>
  <views>...</views>
  <connectors>...</connectors>
  <buses>...</buses>
</module>
```


## Specification

- [fritzingVersion](#fritzingversion)
- [moduleId](#moduleid)
- [referenceFile](#referencefile)
- [version](#version)
- [title](#title)
- [description](#description)
- [author](#author)
- [date](#date)
- [url](#url)
- [label](#label)
- [tags](#tags)
- [taxonomy](#taxonomy)
- [language](#language)
- [family](#family)
- [variant](#variant)
- [properties](#properties)
- [views](#views)
- [connectors](#connectors)
- [buses](#buses)

### fritzingVersion


### moduleId


### referenceFile


### version
Store the part version

### title
Store the part title

### description
Store the part description (you can use simple html, as defined by Qt's Rich Text)

### author
Store the part author

### date
Store the part date

### url
Store the part's url if it is described on a web page

### label
Store the default part label prefix

### tags
Store the part's tags

### taxonomy

### language

### family
Store the parts family (what other parts is this related to)

### variant
Store the part's variant (this makes it unique from all other parts in the same family)

### properties
Store the part's properties

### views

- icon
- breadboard
- pcb
- schematic  

### connectors

### buses
