

# numberstr

![TheOriginalAmarilloMike](assets/sittingduck003.png)

A software utility to manage integer and floating point number strings.
The **numberstr** code package is written in the [Go Programming Language](https://golang.org/)  (a.k.a. Golang).


## Table of Contents
- [numberstr](#numberstr)
  - [Table of Contents](#table-of-contents)
  - [Open Source License](#open-source-license)
  - [Source Code Repository](#source-code-repository)
  - [Go Modules](#go-modules)
  - [Version 1.0.0](#version-100)
  - [Type NumStrDto](#type-numstrdto)
    - [Integer Conversions](#integer-conversions)
    - [Floating Point Conversions](#floating-point-conversions)
    - [Number String Conversions](#number-string-conversions)
      - [Pure Number Strings - Absolute Values](#pure-number-strings---absolute-values)
      - [Negative Value Number Strings](#negative-value-number-strings)
      - [Positive Value Number Strings](#positive-value-number-strings)
      - [Thousands Separator](#thousands-separator)
      - [Currency Symbol](#currency-symbol)
      - [Numeric Separators](#numeric-separators)

## Open Source License

Copyright 2019 Mike Rapp. All rights reserved.

Use of this source code is governed by the (open-source) MIT-style
license which can be found in the LICENSE file located in this directory.

[MIT Open Source License](./LICENSE.md)

## Source Code Repository

The source code for **numberstr** is located at [GitHub](https://github.com/MikeAustin71/numberstrops).

## Go Modules

This project uses Go Modules. The code was written with Go Version 1.15, but should be compatible with earlier Go versions.

## Version 1.0.0

The current version of **numberstr** is Version 1.0.0.



## Type NumStrDto

Currently most of the functionality offered by this code package is found in type NumStrDto which is located in source file, numberstr/numstrdto.go.

NumStrDto provides a host of functions which provide for the conversion of differing numeric types to number strings as well as other functions which convert number strings to specific numeric types.

### Integer Conversions

Integer conversions to number strings includes the following integer types:

- int
- int32
- int64
- big.Int

Notice that the inclusion of the ***big.Int*** type allows NumStrDto to process very, very large integers. If you are unfamiliar with the ***big.Int*** type, please reference the documentation for the **Golang** [big math package](https://golang.org/pkg/math/big/).

### Floating Point Conversions

Floating point number conversions to number strings includes the following floating point types:

- float
- float32
- float64
- big.Float
- big.Rat - Rational Number

As with integer conversions, the inclusion of types ***big.Float*** and ***big.Rat*** allow for the processing of very, very large floating point numbers. If you are unfamiliar with the big.Float and big.Rat types, please reference the documentation for the Golang [big math package](https://golang.org/pkg/math/big/).

### Number String Conversions

All of the integer and floating point number types listed can be converted to number strings and back again to specific numeric types. In addition, number strings can be requested and displayed in a variety of formats.

#### Pure Number Strings - Absolute Values

Pure number strings present integer or floating point numbers as absolute or positive values with no leading plus (+) or minus (-) signs. Integer values are presented as a string of pure numeric digits. Fractional values are presented as a string of pure numeric digits with integer and fractional components delimited by a decimal separator (usually a decimal point).

#### Negative Value Number Strings

Negative value number strings may surrounded with **parentheses** **'()'** or with specified with a **leading minus sign '-'** .

#### Positive Value Number Strings

Positive numeric values may be displayed with or without a **leading plus sign '+'**.

#### Thousands Separator

Type **NumStrDto** provides the option to delimit number strings using a custom Thousands Separator. Users  may specify the standard USA Thousands Separator (comma ',') or another custom Thousands Separator.

#### Currency Symbol

Type **NumStrDto** provides the option to delimit number strings using a custom currency symbol.  Users  may specify the standard USA Currency Symbol (dollar sign '$') or another type of currency symbol using a Unicode specification.

#### Numeric Separators

Type **NumStrDto** employs the concept of **Numeric Separators**. Numeric Separators allow users to specify custom Unicode characters for use as Decimal Separators, Thousands Separators and Currency Symbols. This ensures that number strings can be configured for use by multiple nationalities and cultures.









