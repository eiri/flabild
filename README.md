# flabild
[![CI Status](https://github.com/eiri/flabild/actions/workflows/ci.yaml/badge.svg)](https://github.com/eiri/flabild/actions/workflows/ci.yaml)

Generator of fake pronounceable words

## Summary

This is [Markov chain](https://en.wikipedia.org/wiki/Markov_chain) based generator of fake words. During generation each following letter choosen by a frequency based on two previous letters. This leads to generaton of (semi-) pronounceable words. Frequency module is generated from provided dictionary, allowing for generation of fake words in different languages. The architecture is pluggable. Currently here is only two plugin: english and russian.

### Name

**flabild** is a generated fake pronounceable word meaning generator of fake pronounceable words.

## Usage

```bash
$ flabild -n 12
an
stlentanes
pose
ser
mer
ble
in
aallingioldwidly
dianta
obbly
mirt
cometal

$ flabild -n 7 -lang ru
эшею
заясь
мещриваешься
альное
вышлёном
говавшие
малившеесочке
```

## License

[MIT](https://github.com/eiri/flabild/blob/main/LICENSE)
