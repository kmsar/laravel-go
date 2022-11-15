# LINQ for Go with type parameters

LINQ (Language-Integrated Query) is a component that adds data querying capabilities of Microsoft .Net languages.
This package provides the implementation of the LINQ functions for Go with type parameters.

## Quick Start

### Install

```Shell
go get github.com/makiuchi-d/linq/v2
```

### Example

```Go
package main

import (
	"fmt"

	"github.com/makiuchi-d/linq/v2"
)

type Student struct {
	Name  string
	Class string
	Score int
}

func main() {
	students := []Student{
		// generated by https://testdata.userlocal.jp/
		{"熊木 緑", "1-A", 953},
		{"山本 千佳子", "1-C", 559},
		{"星 雅彦", "1-B", 136},
		{"齊藤 綾子", "1-A", 149},
		{"杉原 和己", "1-C", 737},
		{"山路 信之", "1-B", 425},
		{"佐々木 淑子", "1-C", 759},
		{"三宅 直人", "1-B", 594},
		{"緒方 俊", "1-B", 405},
		{"稲井 隆生", "1-A", 495},
	}

	e1 := linq.FromSlice(students)
	e2 := linq.Where(e1, func(s Student) (bool, error) { return s.Class == "1-B", nil })
	e3 := linq.OrderByDescending(e2, func(s Student) (int, error) { return s.Score, nil })

	linq.ForEach(e3, func(s Student) error {
		fmt.Printf("%d %s\n", s.Score, s.Name)
		return nil
	})
}
```

Output:
```
594 三宅 直人
425 山路 信之
405 緒方 俊
136 星 雅彦
```

## Functions

_italics are unique to this package._

#### Sorting Data

- OrderBy
- OrderByDescending
- _OrderByFunc_
- ThenBy
- ThenByDescending
- Reverse

#### Set Operations

- Distinct
- DistinctBy
- Except
- ExceptBy
- Intersect
- IntersectBy
- Union
- UnionBy

#### Filtering Data

- Where

#### Quantifier Operations

- All
- Any
- Contains
- _ContainsFunc_

#### Projection Operations

- Select
- SelectMany
- Zip

#### Partitioning Data

- Skip
- SkipLast
- SkipWhile
- Take
- TakeLast
- TakeWhile
- Chunk

#### Join Operations

- GroupJoin
- Join

#### Grouping Data

- GroupBy

#### Generation Operations

- DefaultIfEmpty
- Empty
- Range
- Repeat

#### Element Operations

- ElementAt
- ElementAtOrDefault
- First
- FirstOrDefault
- Last
- LastOrDefault
- Single
- SingleOrDefault

#### Converting Data Types

- _FromMap_
- _FromSlice_
- _ToMap_
- _ToMapFunc_
- _ToSlice_

#### Concatenation Operations

- Concat

#### Aggregation Operations

- Aggregate
- Average
- Count
- Max
- MaxBy
- _MaxByFunc_
- Min
- MinBy
- _MinByFunc_
- Sum
- _Sumf_

#### Other

- _ForEach_

## C# LINQ Documents

### [Standard Query Operators](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/standard-query-operators-overview)

#### [Sorting Data](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/sorting-data)

- [x] [OrderBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.orderby)
- [x] [OrderByDescending](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.orderbydescending)
- [x] [ThenBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.thenby)
- [x] [ThenByDescending](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.thenbydescending)
- [x] [Reverse](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.reverse)

#### [Set Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/set-operations)

- [x] [Distinct](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.distinct)
- [x] [DistinctBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.distinctby)
- [x] [Except](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.except)
- [x] [ExceptBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.exceptby)
- [x] [Intersect](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.intersect)
- [x] [IntersectBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.intersectby)
- [x] [Union](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.union)
- [x] [UnionBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.unionby)

#### [Filtering Data](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/filtering-data)

- [ ] [OfType](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.oftype)
- [x] [Where](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.where)

#### [Quantifier Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations)

- [x] [All](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.all)
- [x] [Any](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.any)
- [x] [Contains](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.contains)

#### [Projection Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/projection-operations)

- [x] [Select](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.select)
- [x] [SelectMany](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.selectmany)
- [x] [Zip](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.zip)

#### [Partitioning Data](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/partitioning-data)

- [x] [Skip](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.skip)
- [x] [SkipLast](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.skiplast)
- [x] [SkipWhile](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.skipwhile)
- [x] [Take](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.take)
- [x] [TakeLast](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.take)
- [x] [TakeWhile](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.takewhile)
- [x] [Chunk](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.chunk)

#### [Join Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/join-operations)

- [x] [Join](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.join)
- [x] [GroupJoin](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.groupjoin)

#### [Grouping Data](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/grouping-data)

- [x] [GroupBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.groupby)
- [ ] [ToLookup](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.tolookup)

#### [Generation Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/generation-operations)

- [x] [DefaultIfEmpty](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.defaultifempty)
- [x] [Empty](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.empty)
- [x] [Range](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.range)
- [x] [Repeat](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.repeat)

#### [Equality Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/equality-operations)

- [ ] [SequenceEqual](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.sequenceequal)

#### [Element Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/element-operations)

- [x] [ElementAt](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.elementat)
- [x] [ElementAtOrDefault](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.elementatordefault)
- [x] [First](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.first)
- [x] [FirstOrDefault](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.firstordefault)
- [x] [Last](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.last)
- [x] [LastOrDefault](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.lastordefault)
- [x] [Single](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.single)
- [x] [SingleOrDefault](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.singleordefault)

#### [Converting Data Types](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/converting-data-types)

- [ ] [AsEnumerable](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.asenumerable)
- [ ] [AsQueryable](https://docs.microsoft.com/en-us/dotnet/api/system.linq.queryable.asqueryable)
- [ ] [Cast](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.cast)
- [ ] [OfType](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.oftype)
- [ ] [ToArray](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.toarray)
- [ ] [ToDictionary](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.todictionary)
- [ ] [ToList](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.tolist)
- [ ] [ToLookup](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.tolookup)

#### [Concatenation Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/concatenation-operations)

- [x] [Concat](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.concat)

#### [Aggregation Operations](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/aggregation-operations)

- [x] [Aggregate](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.aggregate)
- [x] [Average](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.average)
- [x] [Count](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.count)
- [ ] [LongCount](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.longcount)
- [x] [Max](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.max)
- [x] [MaxBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.maxby)
- [x] [Min](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.min)
- [x] [MinBy](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.minby)
- [x] [Sum](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable.sum)