# Changelog

## [3.0.1](https://github.com/scottmckendry/ccinvoice/compare/v3.0.0...v3.0.1) (2025-06-10)


### Bug Fixes

* **pdf:** replace wkhtmltopdf with chromedp for pdf generation ([74ee74b](https://github.com/scottmckendry/ccinvoice/commit/74ee74bed4edebf2233a0d3ae2a9fa80571b838f))

## [3.0.0](https://github.com/scottmckendry/ccinvoice/compare/v2.3.1...v3.0.0) (2025-06-10)


### ⚠ BREAKING CHANGES

* **github-action:** Update action docker/build-push-action ( v5 → v6 )

### Features

* **container:** update image golang ( 1.23 → 1.24 ) ([2bc891d](https://github.com/scottmckendry/ccinvoice/commit/2bc891d524ba56b60b8370b35721dcc06c984b47))
* **deps:** update module modernc.org/sqlite ( v1.37.1 → v1.38.0 ) ([f3275bd](https://github.com/scottmckendry/ccinvoice/commit/f3275bdacc4da2b477cbe686ee74c0c4e726cfbf))


### Bug Fixes

* **deps:** update dependency go ( 1.24.1 → 1.24.4 ) ([5d09d82](https://github.com/scottmckendry/ccinvoice/commit/5d09d827775bf75ff19a3ae4f925c8bfcc2830bd))


### Continuous Integration

* **github-action:** Update action docker/build-push-action ( v5 → v6 ) ([e48d32c](https://github.com/scottmckendry/ccinvoice/commit/e48d32c686553d71209d3bb448d7e58258308585))

## [2.3.1](https://github.com/scottmckendry/ccinvoice/compare/v2.3.0...v2.3.1) (2025-05-23)


### Bug Fixes

* **deps:** bump github.com/gofiber/fiber/v2 from 2.52.6 to 2.52.7 ([6c894ff](https://github.com/scottmckendry/ccinvoice/commit/6c894ffc0960adf1d33b938ee78fe722eddb7eb3))

## [2.3.0](https://github.com/scottmckendry/ccinvoice/compare/v2.2.1...v2.3.0) (2025-04-19)


### Features

* add version badge in ui ([b7050b4](https://github.com/scottmckendry/ccinvoice/commit/b7050b4dc8f7393602bfa8f56362b6f6195bb40f))


### Bug Fixes

* handle invoice number generation for shorter names ([a745408](https://github.com/scottmckendry/ccinvoice/commit/a7454082ad76eb0d8dfd46c0247dbb5ddbcf4aed))

## [2.2.1](https://github.com/scottmckendry/ccinvoice/compare/v2.2.0...v2.2.1) (2025-04-05)


### Bug Fixes

* **build:** copy build assets correctly ([7d3d633](https://github.com/scottmckendry/ccinvoice/commit/7d3d63349a7ed314055c4cbb35b93de0bd1166eb))

## [2.2.0](https://github.com/scottmckendry/ccinvoice/compare/v2.1.2...v2.2.0) (2025-04-05)


### Features

* **build:** docker image improvements/tidy up ([a8996cc](https://github.com/scottmckendry/ccinvoice/commit/a8996cc4b2d54f904cf38d52464347c403a1a72d))


### Bug Fixes

* **ci:** create data dir if not exists ([85f7316](https://github.com/scottmckendry/ccinvoice/commit/85f7316abc64a092a1ff8e3417fee6b2d9f91995))

## [2.1.2](https://github.com/scottmckendry/ccinvoice/compare/v2.1.1...v2.1.2) (2025-04-05)


### Bug Fixes

* **ci:** include `tidy` cmd in docker build ([bfb6951](https://github.com/scottmckendry/ccinvoice/commit/bfb6951d6bc79bb4a59503ef2679bd354fd4cbc8))
* make .env optional ([0e66993](https://github.com/scottmckendry/ccinvoice/commit/0e6699384edb31848164b171abc4e60ef9f452ea))

## [2.1.1](https://github.com/scottmckendry/ccinvoice/compare/v2.1.0...v2.1.1) (2025-03-17)


### Bug Fixes

* release dependency updates ([850e5de](https://github.com/scottmckendry/ccinvoice/commit/850e5dea8614e2b5da11de6f0fda4a74ae8f0a33))

## [2.1.0](https://github.com/scottmckendry/ccinvoice/compare/v2.0.1...v2.1.0) (2025-02-03)


### Features

* **ui:** add theme toggle button ([b8a2421](https://github.com/scottmckendry/ccinvoice/commit/b8a2421f433f7692d2e0568eeebffba5481e1dfc))


### Bug Fixes

* **db:** prevent orphaned services when dogs are deleted ([f4db3c5](https://github.com/scottmckendry/ccinvoice/commit/f4db3c5e51b18acaafcb39506bf462df0daf96b6))
* **ui:** prevent "sticky" hover styles on mobile devices ([b9bb23b](https://github.com/scottmckendry/ccinvoice/commit/b9bb23b18b8bffde9f835627d1dfe200b2860063))

## [2.0.1](https://github.com/scottmckendry/ccinvoice/compare/v2.0.0...v2.0.1) (2025-02-01)


### Bug Fixes

* **log:** only log when 1 or more emails are processed ([b3095e9](https://github.com/scottmckendry/ccinvoice/commit/b3095e92a111ccdd6f0a3b81cb89786b59ae3138))

## [2.0.0](https://github.com/scottmckendry/ccinvoice/compare/v1.0.0...v2.0.0) (2025-02-01)


### ⚠ BREAKING CHANGES

* add support for multiple line items

### Features

* add support for multiple line items ([3fad786](https://github.com/scottmckendry/ccinvoice/commit/3fad7867e19e682f32abf7055ec17bfe2d6ade0f))
* **db:** add migrations logic and scripts for changes to services ([41b06fe](https://github.com/scottmckendry/ccinvoice/commit/41b06fedf50bde116f8592a98fa266e7cf27b3f1))
* update invoice styling ([1690667](https://github.com/scottmckendry/ccinvoice/commit/1690667846535b1a7330d52664e0a18883dc4587))


### Bug Fixes

* **ui:** handle cases where all services are removed ([9e5ebe3](https://github.com/scottmckendry/ccinvoice/commit/9e5ebe357a5bcef394f425a78a197896ce664216))
* **ui:** handle re-ordering of the services container ([be1d8c3](https://github.com/scottmckendry/ccinvoice/commit/be1d8c3a4ca8153f102a0b18b6893792c8cd2cec))

## 1.0.0 (2025-01-31)


### Features

* add hx-confirm to send button ([03377b0](https://github.com/scottmckendry/ccinvoice/commit/03377b0e919c225db4fc668d8ecbac29afc3d012))
* add timer job and queue for sending invoices ([4fc561c](https://github.com/scottmckendry/ccinvoice/commit/4fc561ccd8b3e0a9cd1f89d441baeb2e0677746a))
* cc logo in header ([4356287](https://github.com/scottmckendry/ccinvoice/commit/43562871d7624bb81f1b6201d02de398c9d9a38c))
* **ci:** move to versioned releases, consolidate workflows ([d8d4f09](https://github.com/scottmckendry/ccinvoice/commit/d8d4f09a46a0c7fce81bf19cbe6bec8795139da2))
* **dev:** enable air proxy server for live reloads ([0f649e3](https://github.com/scottmckendry/ccinvoice/commit/0f649e3c9182af01c81db2030a5d5ce8e5426099))
* **ui:** replace send confirmation with disabled button ([f115ae8](https://github.com/scottmckendry/ccinvoice/commit/f115ae8bda3310a68430d7114f32fe47fae9346f))


### Bug Fixes

* increse email timeout and handle error better ([b897abf](https://github.com/scottmckendry/ccinvoice/commit/b897abfe7508cbc230fe2086a6f6b0702a19029e)), closes [#7](https://github.com/scottmckendry/ccinvoice/issues/7)
* **mail:** update all queued emails before sending ([22009b6](https://github.com/scottmckendry/ccinvoice/commit/22009b6c3f31a8201165a8955f5a1ac64af180a7))
* **test:** temporarily disable flaky tests ([6e52f4b](https://github.com/scottmckendry/ccinvoice/commit/6e52f4bcb1a8651e98fd536de0d1bd49c27fb530))
