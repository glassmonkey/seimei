# benchmark
This package is for operation testing.

# contents

## sample.csv
dummy first and last name test data generated from [anonymous personal Information generator](https://testdata.userlocal.jp/).


## benchmark_test.go
It has two test codes for operation testing.

### TestRunCompareOrigin
this is a regression test against the original tool ([namedivider-python](https://github.com/rskmoi/namedivider-python)).

### TestRunCompareAnswer
This is a test to check the algorithm's accuracy.
The percent of correct answers was 99.52%.
