# Changelog

## [4.1.2](https://github.com/scottmckendry/ccinvoice/compare/v4.1.1...v4.1.2) (2026-02-14)


### Bug Fixes

* **fiber:** make proxy settings configurable ([65d5196](https://github.com/scottmckendry/ccinvoice/commit/65d51964ffbe899ef33e0e5f1ad8f4fe216f7497))

## [4.1.1](https://github.com/scottmckendry/ccinvoice/compare/v4.1.0...v4.1.1) (2026-02-14)


### Bug Fixes

* **fiber:** configure proxy header explicitly ([69f1134](https://github.com/scottmckendry/ccinvoice/commit/69f11348d3b0e8a9beedf20f4bd8c0ce278d41f0))

## [4.1.0](https://github.com/scottmckendry/ccinvoice/compare/v4.0.0...v4.1.0) (2026-02-14)


### Features

* **container:** update image golang ( 1.25 → 1.26 ) ([dbfc3bf](https://github.com/scottmckendry/ccinvoice/commit/dbfc3bfbc5ac81fda9707eb4b4c4698c367a1247))
* trust proxy headers ([013d6b5](https://github.com/scottmckendry/ccinvoice/commit/013d6b53bddcfe2dd095e0ea3481d01fa1c2c93c))

## [4.0.0](https://github.com/scottmckendry/ccinvoice/compare/v3.2.1...v4.0.0) (2026-02-06)


### ⚠ BREAKING CHANGES

* **deps:** upgrade to fiber v3
* **data:** move from sqlite to postgresql

### Features

* **deps:** update module github.com/go-co-op/gocron/v2 ( v2.16.6 → v2.19.1 ) ([bc56681](https://github.com/scottmckendry/ccinvoice/commit/bc56681a611a75f6428ba61b4d5deae687a6c8e6))
* **deps:** upgrade to fiber v3 ([b95ac34](https://github.com/scottmckendry/ccinvoice/commit/b95ac340a1e04246bae74cf7471e87404b61c6f8))


### Bug Fixes

* **ci:** support pgsql for integration tests ([e0ab6a9](https://github.com/scottmckendry/ccinvoice/commit/e0ab6a9b6ec49a2c433d40917d7df54d64a8fe6a))
* **deps:** update module github.com/chromedp/chromedp ( v0.14.1 → v0.14.2 ) ([158643f](https://github.com/scottmckendry/ccinvoice/commit/158643f6491dab503b1dd36a28f014e0c3bb6506))


### Code Refactoring

* **data:** move from sqlite to postgresql ([f1b313d](https://github.com/scottmckendry/ccinvoice/commit/f1b313dcb99ac026c2a1bcbd3fe502b832d27c52))

## [3.2.1](https://github.com/scottmckendry/ccinvoice/compare/v3.2.0...v3.2.1) (2025-09-28)


### Bug Fixes

* **build:** use chromium instead of chrome ([9435589](https://github.com/scottmckendry/ccinvoice/commit/9435589c1a2f296026a2c5c3be48ea367c1d7cbe))

## [3.2.0](https://github.com/scottmckendry/ccinvoice/compare/v3.1.0...v3.2.0) (2025-09-28)


### Features

* **container:** update image golang ( 1.24 → 1.25 ) ([b8b95f5](https://github.com/scottmckendry/ccinvoice/commit/b8b95f560862c6391cede92094355c56c7bcd5d2))
* **deps:** update module modernc.org/sqlite ( v1.38.2 → v1.39.0 ) ([b333262](https://github.com/scottmckendry/ccinvoice/commit/b33326206da3e6887831df3cadab96e480500bf0))


### Bug Fixes

* **deps:** update module github.com/chromedp/chromedp ( v0.14.0 → v0.14.1 ) ([373a823](https://github.com/scottmckendry/ccinvoice/commit/373a823e8eb0748197e5749b1b09def8e9adb2fe))
* **deps:** update module github.com/go-co-op/gocron/v2 ( v2.16.2 → v2.16.3 ) ([426c2f9](https://github.com/scottmckendry/ccinvoice/commit/426c2f9a9c5ef0b5f8f405abe82abc99f3d13daf))
* **deps:** update module github.com/go-co-op/gocron/v2 ( v2.16.3 → v2.16.6 ) ([16ee3a1](https://github.com/scottmckendry/ccinvoice/commit/16ee3a1e129adc817569df674b40704089bee940))

## [3.1.0](https://github.com/scottmckendry/ccinvoice/compare/v3.0.2...v3.1.0) (2025-07-29)


### Features

* **deps:** update module github.com/chromedp/chromedp ( v0.13.7 → v0.14.0 ) ([a9e8f7b](https://github.com/scottmckendry/ccinvoice/commit/a9e8f7b50c0a0af23d399672d7cb35b9a3d7a692))


### Bug Fixes

* **deps:** update module github.com/chromedp/chromedp ( v0.13.6 → v0.13.7 ) ([#86](https://github.com/scottmckendry/ccinvoice/issues/86)) ([3c7c305](https://github.com/scottmckendry/ccinvoice/commit/3c7c305128e0a563204048b3896fd2f82c37f4a1))
* **deps:** update module github.com/gofiber/fiber/v2 ( v2.52.8 → v2.52.9 ) ([#93](https://github.com/scottmckendry/ccinvoice/issues/93)) ([b93d41f](https://github.com/scottmckendry/ccinvoice/commit/b93d41f349bb0a03fe58f7886b3e6be9c4b2b9f3))

## [3.0.2](https://github.com/scottmckendry/ccinvoice/compare/v3.0.1...v3.0.2) (2025-06-11)


### Bug Fixes

* **pdf:** crashes in non-root envs with sandboxed chromedp ([3736d0e](https://github.com/scottmckendry/ccinvoice/commit/3736d0efdfbd5aeae89c954ae7d70729723c2b82))

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
