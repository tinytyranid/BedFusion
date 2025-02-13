# BedFusion

A small specialised tool for sorting, merging and padding bed files

Usage: `bedfusion <inputs> ... [flags]`

BedFusion follows the bed file standard outlined in: [Niu J., Denisko D. & Hoffman M. M. (2022): *The Browser Extensible Data (BED)* format](https://github.com/samtools/hts-specs/blob/94500cf76f049e898dec7af23097d877fde5894e/BEDv1.pdf)

## Quick-Start Guide

BedFusion will both sort (lexicographically) and merge regions by default. 

Example bed file `examples/merge-test.bed`:

``` text
1	1	4	1	A
1	5	8	1	A
1	6	8	1	A
1	5	8	-1	A
2	5	8	1	A
1	5	8	1	B
1	20	30	1	A
```

``` shell
> bedfusion examples/merge-test.bed
1       1       8       1,-1    A,B
1       20      30      1       A
2       5       8       1       A
```

Contrary to [bedtools merge](https://bedtools.readthedocs.io/en/latest/content/tools/merge.html), BedFusion merges touching regions (like the two first lines in the example bed file). If you prefer to only merge overlapping, and not touching, regions you can use the flag `--overlap=-1`:

``` shell
> bedfusion examples/merge-test.bed --overlap=-1
1       1       4       1       A
1       5       8       1,-1    A,B
1       20      30      1       A
2       5       8       1       A
```

### Using several bed files as input

Several bed files can be used as input as long as they contain same number of columns. These files will be joined and then merged and sorted together.

Example bed file `examples/merge-test2.bed`:

``` text
2	1	4	1	A
2	5	8	1	A
2	6	8	1	A
2	5	8	-1	A
1	5	8	1	A
2	5	8	1	B
2	20	30	1	A
```

``` shell
> bedfusion examples/merge-test.bed examples/merge-test2.bed
1       1       8       1,-1    A,B
1       20      30      1       A
2       1       8       1,-1    A,B
2       20      30      1       A
```

## Examples

- [sorting](./docs/sorting.md)
- [merging](./docs/merging.md)
- [padding](./docs/padding.md)
- [track files](./docs/track-files.md)
- [using a configuration file](./docs/config-file.md)

## Flags and arguments 

BedFusion supports three separate ways of configuration: flags, a configuration file or environmental variables. If a combination of the three is used the reading priority order is as follows: 

1. flags 
2. configuration file 
3. environmental variables

Order of actions ( \* = can be turned on/off using flags): 

1. reading files 
2. padding(\*)
3. merging(\*)/deduplication(\*)
4. sorting 
5. writing output 

| Arguments      |                                                                                                  |
|----------------|--------------------------------------------------------------------------------------------------|
| `<inputs> ...` | Bed file path(s). If more than one is provided the files will be joined as if they were one file |


| Flags (with format and defaults)    | Environmental variables | Description                                                                                                                                                                                                                                                                                                                                                                                                                         |
|-------------------------------------|-------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `-h`<br>`--help`                    |                         | Show context-sensitive help.                                                                                                                                                                                                                                                                                                                                                                                                        |
| `-c`<br>`--config-file=CONFIG-FLAG` | `CONFIG_FILE`           | The path to configuration file (must be in key-value yaml format)                                                                                                                                                                                                                                                                                                                                                                   |
| `-o`<br>`--output=STRING`           | `OUTPUT_FILE`           | Path to the output file. If unset the output will be written to stdout                                                                                                                                                                                                                                                                                                                                                              |
| `-f`<br>`--fasta-idx=STRING`        | `FASTA_IDX`             | Tab separated file containing at least two columns where the first column contains the chromosome and the second it's size. Compatible with fasta index files, but any text file can be used as long as the file conditions are met                                                                                                                                                                                                 |
|                                     |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **input**                           |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `--strand-col=INT`                  | `STRAND_COL`            | The column containing the strand information (1-based column index). If this option is set regions on the same strand will not be merged                                                                                                                                                                                                                                                                                            |
| `--feat-col=INT`                    | `FEAT_COL`              | The column containing the feature (e.g. gene id, transcript id etc.) information (1-based column index). If this option is set regions on the same feature will not be merged                                                                                                                                                                                                                                                       |
|                                     |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **sorting**                         |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `-s`<br>`--sort-type="lex"`         | `SORT_TYPE`             | How the bed file should be sorted.<br>- lex = lexicographic sorting (chr: 1 < 10 < 2 < MT < X)<br>- nat = natural sorting (chr: 1 < 2 < 10 < MT < X)<br>- ccs = custom chromosome sorting (see `--chr-order` flag )<br>- fidx = use ordering from fasta index file (must be used together with `--fasta-idx`)                                                                                                                       |
| `--chr-order=CHR-ORDER,...`         | `CHR_ORDER`             | Comma separated custom chromosome order, to be used with custom chromosome sorting (--sort-type=ccs). Chromosomes not on the list will be sorted naturally after the ones in the list                                                                                                                                                                                                                                               |
| `-d`<br>`--deduplicate`             | `DEDUPLICATE`           | Remove duplicated lines                                                                                                                                                                                                                                                                                                                                                                                                             |
|                                     |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **merging**                         |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `--no-merge`                        | `NO_MERGE`              | Do not merge regions                                                                                                                                                                                                                                                                                                                                                                                                                |
| `--overlap=0`                       | `OVERLAP`               | Overlap between regions to be merged. Note that touching regions are merged (e.g. if two regions are on the same chr, and the overlap is they will be merged if one ends at 5 and the other starts at 6). If you don't want touching regions to be merged set overlap to -1                                                                                                                                                         |
|                                     |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **padding**                         |                         |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `-p`<br>`--padding=INT`             | `PADDING`               | Padding in bp. Note that padding is done before merging                                                                                                                                                                                                                                                                                                                                                                             |
| `--padding-type="safe"`             | `PADDING_TYPE`          | Padding type.<br>- safe = bedfusion will fail if it encounters a chromosome not in the fasta index file,<br>-lax = will only pad regions in the fasta index file and give a warning about chromosomes not in the fasta index file,<br>- force = will pad regardless, if `--fasta-idx` is set there will be given a warning about the chromosomes not in the fasta index file, if `--fasta-idx` is not set no warnings will be given |
| `--first-base=0`                    | `FIRST_BASE`            | The start coordinate of the first base on each chromosome                                                                                                                                                                                                                                                                                                                                                                           |
