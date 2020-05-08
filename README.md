# twilight-loader

Go package used to calculate sunrise and sunset times and generate SQL insert script in standard output.

------

### Installation

```bash
go get github.com/dendna/twilight-loader
cd ${GOPATH}/src/github.com/dendna/twilight/cmd/twilight-loader
go build
```


### Usage

Set input parameters in `config.json` file:

`year`
`timezone_name`
`latitude`
`longitude`
`morning_twilight_type`
`sunrise_type`
`sunset_type`
`evening_twilight_type`

Run application:

```bash
twilight-loader
```