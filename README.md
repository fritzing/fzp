# [fritzing](https://github.com/fritzing) fzp specs & tools

- [Sample `fzp`](#sample-fzp)
- [Tools](#tools)
  - [Validator](#validator)
- [Libraries/Code](#librariescode)
  - [Qt](#qt)
  - [go](#go)
- [Specification](#specification)
  - [`module`](#module)
    - [`fritzingVersion`](#fritzingversion)
    - [`moduleId`](#moduleid)
    - [`referenceFile`](#referencefile)
  - [`version`](#version)
  - [`title`](#title)
  - [`description`](#description)
  - [`author`](#author)
  - [`date`](#date)
  - [`url`](#url)
  - [`label`](#label)
  - [`tags`](#tags)
    - [`tag`](#tag)
  - [`taxonomy`](#taxonomy)
  - [`language`](#language)
  - [`family`](#family)
  - [`variant`](#variant)
  - [`properties`](#properties)
    - [`property`](#property)
  - [`views`](#views)
    - [`iconView`](#iconview)
    - [`breadboardView`](#breadboardview)
    - [`schematicView`](#schematicview)
    - [`pcbView`](#pcbview)
      - [`copper0`](#copper0)
      - [`copper1`](#copper1)
  - [`connectors`](#connectors)
    - [`connector`](#connector)
      - [`description`](#description)
      - [`views`](#views)
        - [`breadboard`](#breadboard)
        - [`schematic`](#schematic)
        - [`pcb`](#pcb)
  - [`buses`](#buses)

read more about the fzp format at:  
[https://github.com/fritzing/fritzing-app/wiki/2.1-Part-file-format](https://github.com/fritzing/fritzing-app/wiki/2.1-Part-file-format)

## Sample `fzp`

TO DO: use template.fzp to load the structure of the fzp

```
<?xml version="1.0" encoding="UTF-8"?>
<module fritzingVersion="x.y.z" moduleId="mod-id-rev" referenceFile="ref.file">
  <version>x.y.z</version>
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
A complete sample file can be found [here](sample.fzp)


## Tools
[![Build Status](https://travis-ci.org/fritzing/fzp.svg)](https://travis-ci.org/fritzing/fzp)

### template.fzp ckeckLevel

the template.fzp has a checkLevel identifier for the automated travis checks.  
this checkLevel defines how important the xml-tag of the fzp-xml is, for the fritzing-part.



### Validator
simple and fast validator to test the fritzing-parts repository (over 14k parts).  
i think [go](https://golang.org) is excelent for this job.  
if the validator is ready we can add travis-ci as test service to the fritzing-parts repository.


## Libraries/Code

### Qt
here i think lifes the fritzing-app fzp code...
[fritzing-app](https://github.com/fritzing/fritzing-app/blob/bcf97e72b45fc6d9be836e5014fc53f5eabe0048/src/model/modelpart.h#L114) Qt sourcecode link

### go
see validator


## Specification

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

##### tag


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

##### property


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

##### iconView
##### breadboardView
##### pcbView
##### schematicView
###### layers
###### layer


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
##### connector
###### description


### buses
```
<buses>
</buses>
```
