;;; garoon-test.el --- Test for garoon.el -*- lexical-binding: t; -*-
;;; Commentary:
;;; Code:

(require 'ert)
(require 'garoon)

(ert-deftest garoon-test--create-token ()
  (let ((login    "Administrator")
        (password "cybozu"))
    (should (equal (garoon--create-token login password) "QWRtaW5pc3RyYXRvcjpjeWJvenU="))))

;;; garoon-test.el ends here
