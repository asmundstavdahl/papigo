# papigo
Parallel Pi computing Go program (not a very good one)

## Usage
```sh
go install github.com/asmundstavdahl/papigo
papigo -help
# Default params, saved to a Tab Separated Values file for spreadsheet import
papigo > pi.tsv
# Single goroutine
papigo -buffer=100 -parallellism=0 -chunk=100000
```

Example output:
```
iter.	pi
665100000	3.141592652836512122860312956618145108222961425781250000
1332100000	3.141592653963686032625446387100964784622192382812500000
1995600000	3.141592653337383911349434129078872501850128173828125000
2660800000	3.141592653776176469193615048425272107124328613281250000
3324200000	3.141592653738620732895014953101053833961486816406250000
3988200000	3.141592653462892403837258825660683214664459228515625000
4653300000	3.141592653480865582338310559862293303012847900390625000
5315000000	3.141592653682421243388489529024809598922729492187500000
5980100000	3.141592653671982482421753957169130444526672363281250000
6646200000	3.141592653513167743284384414437226951122283935546875000
```
