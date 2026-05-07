# Changelog

## 0.2.0 (2026-05-07)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/hyperspell/hyperspell-cli/compare/v0.1.0...v0.2.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([40329ac](https://github.com/hyperspell/hyperspell-cli/commit/40329ac0504bec769cdcbd0eda4821b15cf164d1))
* **api:** api update ([cde57e6](https://github.com/hyperspell/hyperspell-cli/commit/cde57e6cc0907e80bc2fa80d05b6701c3b777863))
* **api:** api update ([15efdbc](https://github.com/hyperspell/hyperspell-cli/commit/15efdbcd4ea5962464fe5eedc212d0fd601ceb1f))
* **api:** api update ([6ea47ba](https://github.com/hyperspell/hyperspell-cli/commit/6ea47ba9a0248351256a8ddd0bde7b5575a53eb3))
* **api:** api update ([05fc238](https://github.com/hyperspell/hyperspell-cli/commit/05fc238a5087e6cd3024fd92eccb7fa23cbb8440))
* **api:** api update ([2f69100](https://github.com/hyperspell/hyperspell-cli/commit/2f691007ba68d3b18b6a643db158ed16b4848871))
* **api:** api update ([07389a5](https://github.com/hyperspell/hyperspell-cli/commit/07389a55c46d676894d6392d7c6ed4cf33ac4147))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([665e431](https://github.com/hyperspell/hyperspell-cli/commit/665e431d3c39a2ba149168250d080bdf8b7546b5))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([4222549](https://github.com/hyperspell/hyperspell-cli/commit/422254937af57c228a18d67191aa8966cceebb45))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([fc1ebc7](https://github.com/hyperspell/hyperspell-cli/commit/fc1ebc74cf40a58ddb979128852ed202354922cc))
* **cli:** send filename and content type when reading input from files ([4b76c79](https://github.com/hyperspell/hyperspell-cli/commit/4b76c79a99e8bb0f8a3ef112ccaf3092decb69a3))
* support passing path and query params over stdin ([14f99ef](https://github.com/hyperspell/hyperspell-cli/commit/14f99ef63dff6a55aca07c84419c42ce25d99142))


### Bug Fixes

* **cli:** correctly load zsh autocompletion ([c419250](https://github.com/hyperspell/hyperspell-cli/commit/c419250da2b4345806d79c322c0384742e8b234a))
* **cli:** fix incompatible Go types for flag generated as array of maps ([a95a486](https://github.com/hyperspell/hyperspell-cli/commit/a95a486ccc528af84af0a8aa0500680982ee6033))
* fall back to main branch if linking fails in CI ([28b2498](https://github.com/hyperspell/hyperspell-cli/commit/28b24981adf8c8c39a68a44495912a66e63945ea))
* fix for failing to drop invalid module replace in link script ([ad3cbf1](https://github.com/hyperspell/hyperspell-cli/commit/ad3cbf15ab7a70f3318ae32d95b7f6bebeaebceb))
* fix quoting typo ([d4f3537](https://github.com/hyperspell/hyperspell-cli/commit/d4f35370e45ed2cc4deceeb6b8fae13f0a07fb63))
* flags for nullable body scalar fields are strictly typed ([c983439](https://github.com/hyperspell/hyperspell-cli/commit/c9834390753ff2facace1035b6b692e73a73e81a))


### Chores

* add documentation for ./scripts/link ([2586047](https://github.com/hyperspell/hyperspell-cli/commit/2586047b1af969fcd3b5c21e52c072a7b620118d))
* **ci:** support manually triggering release workflow ([b5836aa](https://github.com/hyperspell/hyperspell-cli/commit/b5836aa567c2ba79cf6ad7a8706de0ef196ed317))
* **cli:** additional test cases for `ShowJSONIterator` ([5784e71](https://github.com/hyperspell/hyperspell-cli/commit/5784e710f1021b0939a8985eedaa0c80b0b9df0b))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([1b5c7e9](https://github.com/hyperspell/hyperspell-cli/commit/1b5c7e9de0a6ceaf663c5a06365a0ce1ec11cca4))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([d9c0c77](https://github.com/hyperspell/hyperspell-cli/commit/d9c0c77f238264b97af2c5fa8aacbd71de7f6f4b))
* **cli:** switch long lists of positional args over to param structs ([4a9e250](https://github.com/hyperspell/hyperspell-cli/commit/4a9e2508c2605e814c40980baf343dbda59f28b8))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([9e0c109](https://github.com/hyperspell/hyperspell-cli/commit/9e0c10938023ffbd46ca50c61baf2c3353df23d5))
* **internal:** more robust bootstrap script ([014bf27](https://github.com/hyperspell/hyperspell-cli/commit/014bf271f51526dab8ae459bd514bab798c8062d))
* mark all CLI-related tests in Go with `t.Parallel()` ([5fc2402](https://github.com/hyperspell/hyperspell-cli/commit/5fc240280f247a2a3817b37e359c7c674048a448))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([b7fd553](https://github.com/hyperspell/hyperspell-cli/commit/b7fd5533571992857a741f834af4d272b8ad2b6d))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([2f49e45](https://github.com/hyperspell/hyperspell-cli/commit/2f49e45e69e821c97605e3929f4d6864f2b306e7))
* **tests:** bump steady to v0.22.1 ([522092e](https://github.com/hyperspell/hyperspell-cli/commit/522092e2767802d6e258353c4bb374f16f2751d5))


### Documentation

* update examples ([62e382b](https://github.com/hyperspell/hyperspell-cli/commit/62e382b449383807ab1568d97c5f461fed2c702c))

## 0.1.0 (2026-04-02)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/hyperspell/hyperspell-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([7a7adea](https://github.com/hyperspell/hyperspell-cli/commit/7a7adea4f0bc67cf4e63974289c3012a7d3bf272))


### Chores

* configure new SDK language ([1f16ef1](https://github.com/hyperspell/hyperspell-cli/commit/1f16ef1332024d8d4f6904d198f2bb5483b3b22f))
* update SDK settings ([628eda3](https://github.com/hyperspell/hyperspell-cli/commit/628eda33943a9ca3c3e723ad9488ffa0739159b7))
* update SDK settings ([59241b8](https://github.com/hyperspell/hyperspell-cli/commit/59241b8dd00f29361b4c27dd3305fd3f43dad5cf))
