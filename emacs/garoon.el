;;; garoon.el --- Cybozu Garoon client for Emacs -*- lexical-binding: t; -*-
;;
;; Copyright (C) 2019 nukosuke
;;
;; Version:  0.0.0
;; Author:   nukosuke <nukosuke@lavabit.com>
;; URL:      https://github.com/nukosuke/emacs-garoon
;; License:  GPLv3+
;; Keywords: schedule
;;
;;; Commentary:
;;; Code:

(eval-when-compile
  (require 'cl-lib))

(defgroup garoon nil
  "Cybozu Garoon client for Emacs."
  :group 'tools
  :prefix "garoon/")

(defcustom garoon/subdomain nil
  "Endpoint subdomain of Garoon API."
  :group 'garoon
  :type  'string)

(defcustom garoon/user-login nil
  "User login name."
  :group 'garoon
  :type  'string)

(defcustom garoon/user-password nil
  "User password."
  :group 'garoon
  :type  'string)

(defconst garoon/client-version "v0.0.0")

(defconst garoon/user-agent
  (format "%s garoon.el/%s"
          emacs-version
          garoon/client-version))

(defconst garoon/auth-header "X-Cybozu-Authorization")

(defconst garoon/command "garoon")

(defconst garoon/command-options
  '(("-d" . garoon/subdomain)
    ("-u" . garoon/user-login)
    ("-p" . garoon/user-password)
    ;"--emacs"
    ))

;;;;; Functions

(defun garoon/create-token (login password)
  "Create the token for auth by user's LOGIN name and PASSWORD.
Auth token is BASE64 encoded string which form of \"LOGIN:PASSWORD\"."
  (base64-encode-string
   (concat login ":" password)))

(defun garoon/event-list (&optional start end)
  "Retrieve event list in the range between START and END."
  (interactive "sSTART (today):\nsEND (today):")
  (let* ((buf (get-buffer-create "*Garoon Event List*"))
         (today (format-time-string "%Y/%m/%d"))
         (start-string (if (string= start "") today start))
         (end-string (if (string= end "") today end)))
    (with-current-buffer buf
      (erase-buffer)
      (let* ((stdout (shell-command-to-string
               (concat "garoon event ls"
                       " -d " garoon/subdomain
                       " -u " garoon/user-login
                       " -p " garoon/user-password
                       " --range " start-string "-" end-string)))
             (rows (mapcar (lambda (row)
                             (split-string row "\t"))
                           (split-string stdout "\n"))))
        (garoon-mode)
        (cl-dolist (row rows)
          (if (= (length row) 4)
              (insert (format "[%s] [%s] %s\n"
                              (nth 1 row)
                              (nth 2 row)
                              (nth 3 row)))))
        (pop-to-buffer buf)))))

(defun garoon/event-info (id)
  "Retrieve the event of ID."
  (interactive )
  id)

;; TODO:
;; - users
;; - organizations
;; - notifications
;; - facilities


;;;;; Major Mode

(defun garoon-mode ()
  "."
  (interactive)
  (kill-all-local-variables)
  (setq mode-name "Garoon")
  (setq major-mode 'garoon-mode))

(defvar garoon-mode-map (make-sparse-keymap))

(define-derived-mode garoon-mmode fundamental-mode "Garoon"
  "Cybozu Garoon client for Emacs."
  (define-key garoon-mode-map "\C-c\C-u" 'garoon-sync-my-events))

;;;;; Commands

(defun garoon-add-event ()
  "Create new event."
  (interactive))

(defun garoon-update-my-events ()
  "Sync login user events up to date."
  (interactive))

(provide 'garoon)

;;; garoon.el ends here
