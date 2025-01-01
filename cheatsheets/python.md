# Python

## Package Managers and Build Tools

* [Poetry](poetry.md)
* [UV](uv.md)

## Rounding
Weird behavior (when applying common school math terminology to programming):
```python
assert round(0.5) == 0
assert round(1.5) == 2
```

Similar example, but rounding to nearest second in Pandas:
```python
from datetime import datetime

import pandas as pd

d = [
    {"dt": datetime(2023, 1, 1, 12, 30, 0, 500000), "dt_up": datetime(2023, 1, 1, 12, 30, 1)},
    {"dt": datetime(2023, 1, 1, 12, 30, 1, 500000), "dt_up": datetime(2023, 1, 1, 12, 30, 2)},
    {"dt": datetime(2023, 1, 1, 12, 30, 0, 500001), "dt_up": datetime(2023, 1, 1, 12, 30, 1)},
    {"dt": datetime(2023, 1, 1, 12, 30, 1, 500001), "dt_up": datetime(2023, 1, 1, 12, 30, 2)},
]
df = pd.DataFrame(d)
df["dt_rounded"] = df["dt"].dt.round("1S")
df["is_correct"] = df.apply(lambda x: x.dt_rounded == x.dt_up, axis=1)
print(df)

"""
                          dt               dt_up          dt_rounded  is_correct
0 2023-01-01 12:30:00.500000 2023-01-01 12:30:01 2023-01-01 12:30:00       False
1 2023-01-01 12:30:01.500000 2023-01-01 12:30:02 2023-01-01 12:30:02        True
2 2023-01-01 12:30:00.500001 2023-01-01 12:30:01 2023-01-01 12:30:01        True
3 2023-01-01 12:30:01.500001 2023-01-01 12:30:02 2023-01-01 12:30:02        True
"""
```

Forcing rounding to nearest 1 second (in pandas)
```python
from datetime import datetime

import pandas as pd


def round_to_nearest_second(dt_val: pd.Series) -> pd.Series:
    """
    Perform accurate frequency rounding on a Pandas datetime series

    Series.dt.round(freq) contains the caveat where values which are exactly in the middle of their upper and lower frequencies will
    prioritize a resulting value which is even. This is the same caveat as Python's "round()".
    e.g. round(0.5) == 0; round(1.5) == 2

    For datetimes which are being rounded to the nearest 1 second, round(2023-04-10 16:06:12.500000) == 2023-04-10 16:06:12 whereas 12.500001 rounds to 13.

    :param dt_val: datetime64[ns, utc] pandas series
    :return: Series of rounded datetimes
    """
    # https://pandas.pydata.org/docs/user_guide/timeseries.html#timeseries-offset-aliases
    freq = "1S"
    val = dt_val
    if dt_val.microsecond >= 500000:
        val = dt_val.ceil(freq)
    else:
        val = dt_val.floor(freq)
    return val


d = [
    {"dt": datetime(2023, 1, 1, 12, 30, 0, 500000), "dt_up": datetime(2023, 1, 1, 12, 30, 1)},
    {"dt": datetime(2023, 1, 1, 12, 30, 1, 500000), "dt_up": datetime(2023, 1, 1, 12, 30, 2)},
    {"dt": datetime(2023, 1, 1, 12, 30, 0, 500001), "dt_up": datetime(2023, 1, 1, 12, 30, 1)},
    {"dt": datetime(2023, 1, 1, 12, 30, 1, 500001), "dt_up": datetime(2023, 1, 1, 12, 30, 2)},
]
df = pd.DataFrame(d)
df["dt_rounded"] = df["dt"].apply(round_to_nearest_second)
df["is_correct"] = df.apply(lambda x: x.dt_rounded == x.dt_up, axis=1)
print(df)

"""
                          dt               dt_up          dt_rounded  is_correct
0 2023-01-01 12:30:00.500000 2023-01-01 12:30:01 2023-01-01 12:30:00        True
1 2023-01-01 12:30:01.500000 2023-01-01 12:30:02 2023-01-01 12:30:02        True
2 2023-01-01 12:30:00.500001 2023-01-01 12:30:01 2023-01-01 12:30:01        True
3 2023-01-01 12:30:01.500001 2023-01-01 12:30:02 2023-01-01 12:30:02        True
"""
```

Explanation:

* Python's [`round()`](https://docs.python.org/3/library/functions.html#round) function prioritizes numerically even-valued results when the value is equally between its multiples.
    * e.g. 0.5 is an equal distance between 0 and 1; thus 0 is prioritized.
    * e.g. 1.5 is an equal distance between 1 and 2; thus 2 is prioritized.
* Pandas' [`Series.dt.round()`](https://pandas.pydata.org/docs/reference/api/pandas.Series.dt.round.html) function uses the same behavior as Python's `round()`.

This is not a Python bug. Read more from [RealPython](https://realpython.com/python-rounding/).
