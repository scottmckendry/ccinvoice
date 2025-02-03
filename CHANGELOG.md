# Changelog

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
