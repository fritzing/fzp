# [fritzing](https://github.com/fritzing) fzp specs & tools

- [Sample fzp](#sample-fzp)
- [Specification](#specification)


## Sample fzp

```
<?xml version="1.0" encoding="UTF-8"?>
<module fritzingVersion="x.y.z" moduleId="mod-id-rev" referenceFile="ref.file">
  <version>1.1</version>
  <title>part-name</title>
  <description>some words about the part</description>
  <author>your-name</author>
  <date>yyyy-mm-dd</date>
  <url>http://part.org</url>
  <label>IC</label>
  <tags>...</tags>
  <taxonomy>...</taxonomy>
  <language>...</language>
  <family>...</family>
  <variant>...</variant>
  <properties>...</properties>
  <views>...</views>
  <connectors>...</connectors>
  <buses>...</buses>
</module>
```


## Specification

- [module](#module)
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


### module
```
<module fritzingVersion="x.y.z" moduleId="mod-id-rev" referenceFile="ref.file">
```

##### fritzingVersion
?

##### moduleId
?

##### referenceFile
?


### version
```
<version>x.y.z</version>
```
Store the part version


### title
```
<title>part-name</title>
```
Store the part title


### description
```
<description>some words about the part</description>
```
Store the part description (you can use simple html, as defined by Qt's Rich Text)


### author
```
<author>your-name</author>
```
Store the part author


### date
```
<date>yyyy-mm-dd</date>
```
Store the part date


### url
```
<url>http://part.org</url>
```
Store the part's url if it is described on a web page


### label
```
<label>IC</label>
```
Store the default part label prefix


### tags
```
<tags>
  <tag>...</tag>
  <tag>tagvalue</tag>
  <tag>...</tag>
</tags>
```
Store the part's tags


### taxonomy
```
<taxonomy>...</taxonomy>
```


### language
?


### family
```
<family>...</family>
```
Store the parts family (what other parts is this related to)


### variant
```
<variant>...</variant>
```
Store the part's variant (this makes it unique from all other parts in the same family)


### properties
```
<properties>
  <property name="x">...</property>
  <property name="x">the-value</property>
  <property name="x">...</property>
</properties>
```
Store the part's properties


### views
```
<views>
  <breadboardView>
    <layers image="breadboard/part.svg">
      <layer layerId="breadboard"/>
    </layers>
  </breadboardView>
  <schematicView>
    <layers image="schematic/part.svg">
      <layer layerId="schematic"/>
    </layers>
  </schematicView>
  <pcbView>
    <layers image="pcb/part.svg">
      <layer layerId="copper0"/>
      <layer layerId="silkscreen"/>
      <layer layerId="copper1"/>
    </layers>
  </pcbView>
  <iconView>
    <layers image="icon/part.svg">
      <layer layerId="icon"/>
    </layers>
  </iconView>
</views>
```

- icon
- breadboard
- pcb
- schematic  


### connectors
```
<connectors>
  <connector id="connector15" name="B1" type="male">
    <description>Channel 1 voltage low</description>
    <views>
      <breadboardView>
        <p layer="breadboard" svgId="connector15pin" terminalId="connector15terminal"/>
      </breadboardView>
      <schematicView>
        <p layer="schematic" svgId="connector15pin" terminalId="connector15terminal"/>
      </schematicView>
      <pcbView>
        <p layer="copper1" svgId="connector15pin"/>
        <p layer="copper0" svgId="connector15pin"/>
      </pcbView>
    </views>
  </connector>
</connectors>
```


### buses
```
<buses>
</buses>
```
