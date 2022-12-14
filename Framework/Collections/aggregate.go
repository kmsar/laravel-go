package Collections

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	"github.com/shopspring/decimal"
)

func (this *Collection) SafeSum(key ...string) (sum decimal.Decimal) {
	sum = decimal.NewFromInt(0)
	if len(key) == 0 {
		for _, f := range this.ToFloat64Array() {
			sum.Add(decimal.NewFromFloat(f))
		}
	} else {
		this.Map(func(fields Support.Fields) {
			sum = sum.Add(decimal.NewFromFloat(Field.GetFloat64Field(fields, key[0])))
		})
	}
	return
}

func (this *Collection) SafeAvg(key ...string) (sum decimal.Decimal) {
	return this.SafeSum(key...).Div(decimal.NewFromInt32(int32(this.Count())))
}

func (this *Collection) SafeMax(key ...string) (max decimal.Decimal) {
	if len(key) == 0 {
		for _, f := range this.ToFloat64Array() {
			if max.IsZero() {
				max = decimal.NewFromFloat(f)
			} else if float := decimal.NewFromFloat(f); max.LessThan(float) {
				max = float
			}
		}
	} else {
		this.Map(func(fields Support.Fields) {
			if max.IsZero() {
				max = decimal.NewFromFloat(Field.GetFloat64Field(fields, key[0]))
			} else if float := decimal.NewFromFloat(Field.GetFloat64Field(fields, key[0])); max.LessThan(float) {
				max = float
			}
		})
	}
	return
}

func (this *Collection) SafeMin(key ...string) (min decimal.Decimal) {
	if len(key) == 0 {
		for _, f := range this.ToFloat64Array() {
			if min.IsZero() {
				min = decimal.NewFromFloat(f)
			} else if float := decimal.NewFromFloat(f); float.LessThan(min) {
				min = float
			}
		}
	} else {
		this.Map(func(fields Support.Fields) {
			if min.IsZero() {
				min = decimal.NewFromFloat(Field.GetFloat64Field(fields, key[0]))
			} else if float := decimal.NewFromFloat(Field.GetFloat64Field(fields, key[0])); float.LessThan(min) {
				min = float
			}
		})
	}
	return
}

func (this *Collection) Count() int {
	return len(this.array)
}

func (this *Collection) Sum(key ...string) (sum float64) {
	if len(key) == 0 {
		for _, f := range this.ToFloat64Array() {
			sum += f
		}
	} else {
		this.Map(func(fields Support.Fields) {
			sum += Field.GetFloat64Field(fields, key[0])
		})
	}
	return
}

func (this *Collection) Max(key ...string) (max float64) {
	if len(key) == 0 {
		for i, f := range this.ToFloat64Array() {
			if i == 0 {
				max = f
			} else if f > max {
				max = f
			}
		}
	} else {
		this.Map(func(fields Support.Fields, index int) {
			if index == 0 {
				max = Field.GetFloat64Field(fields, key[0])
			} else if float := Field.GetFloat64Field(fields, key[0]); float > max {
				max = float
			}
		})
	}
	return
}

func (this *Collection) Min(key ...string) (min float64) {
	if len(key) == 0 {
		for i, f := range this.ToFloat64Array() {
			if i == 0 {
				min = f
			} else if f < min {
				min = f
			}
		}
	} else {
		this.Map(func(fields Support.Fields, index int) {
			if index == 0 {
				min = Field.GetFloat64Field(fields, key[0])
			} else if float := Field.GetFloat64Field(fields, key[0]); float < min {
				min = float
			}
		})
	}
	return
}

func (this *Collection) Avg(key ...string) float64 {
	return this.Sum(key...) / float64(this.Count())
}
