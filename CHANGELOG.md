# Changelog

All notable changes to this project will be documented in this file

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.3.1](https://github.com/EclesioMeloJunior/terrago/tree/v0.3.1) - 2020-02-07

### Changed
- **.cfbuild** folder mode updated to 0755

## [0.3.0](https://github.com/EclesioMeloJunior/terrago/tree/v0.3.0) - 2014-05-31

### Added
- Resource Read operation check if dir exists if not only erease ID and return nil

## [0.2.0](https://github.com/EclesioMeloJunior/terrago/tree/v0.2.0) - 2014-05-31

### Added
- Packer exclude package-lock.json from cloud function build
- Resource Delete operation check if dir exists if not only erease ID and return nil


## [0.1.4](https://github.com/EclesioMeloJunior/terrago/tree/v0.1.4) - 2014-05-31

### Changed
- Directory **.cfbuilds** have unix permissions

## [0.1.3](https://github.com/EclesioMeloJunior/terrago/tree/v0.1.3) - 2014-05-31

### Changed
- Directory renamed from **.cfbuilds/** to **.cfbuilds**

## [0.1.2](https://github.com/EclesioMeloJunior/terrago/tree/v0.1.2) - 2014-05-31

### Changed
- Cloud Function Packer creates a directory called **.cfbuilds** to store all the zip archives

## [0.1.1](https://github.com/EclesioMeloJunior/terrago/tree/v0.1.1) - 2014-05-31

### Changed
- Cloud Function Packer creates a directory called **.build** to store all the zip archives

## [0.1.0](https://github.com/EclesioMeloJunior/terrago/tree/v0.1.0) - 2014-05-31

### Added
- This CHANGELOG file to hopefully serve as an evolving example of a
  standardized open source project CHANGELOG.
- Cloud Function Packer: terraform provider that normalize and zip a path
  that contains a NodeJS project that will be a GCP Cloud Function