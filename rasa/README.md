# Chat bot with Rasa

To train

```
rasa train -v --data ./domain.yml nlu
```

## Issues

Get the following exception complaining about the default log level
being unknown.

```
Traceback (most recent call last):
  File "/usr/local/bin/rasa", line 8, in <module>
    sys.exit(main())
  File "/usr/local/lib/python3.8/dist-packages/rasa/__main__.py", line 103, in main
    set_log_level(log_level)
  File "/usr/local/lib/python3.8/dist-packages/rasa/utils/common.py", line 64, in set_log_level
    logging.getLogger("rasa").setLevel(log_level)
  File "/usr/lib/python3.8/logging/__init__.py", line 1409, in setLevel
    self.level = _checkLevel(level)
  File "/usr/lib/python3.8/logging/__init__.py", line 194, in _checkLevel
    raise ValueError("Unknown level: %r" % level)
ValueError: Unknown level: 'Level trace'
```

Work around run with "-v" to set log level to info.