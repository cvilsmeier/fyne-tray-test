fyne-tray-test
==============================================================================

A small test program for `https://github.com/fyne-io/systray`.


Build
------------------------------------------------------------------------------

Run with

    $ go run main.go


Expected and actual outcome
------------------------------------------------------------------------------

Expected:

- Systray should show icon `white.png` at startup.

- After 1 second, it should show icon `red.png`.


Actual:

- Systray starts up with `white.png`, that's ok.

- After 1 second, it show the following error message, and does not update the icon:

    2023/05/10 23:32:47 systray error: failed to set IconPixmap prop: dbus.Store: type mismatch: slice: cannot convert a value of []systray.PX into []systray.PX



Test Environment
------------------------------------------------------------------------------

Debian GNU/Linux 11.7 

    $ uname --all
    Linux cv-leni7 5.10.0-22-amd64 #1 SMP Debian 5.10.178-3 (2023-04-22) x86_64 GNU/Linux

    $ cat /etc/debian_version 
    11.7


D-Bus Version

    $ apt list --installed | grep libdbus

    WARNING: apt does not have a stable CLI interface. Use with caution in scripts.

    libdbus-1-3/stable,stable-security,now 1.12.24-0+deb11u1 amd64 [installed,automatic]
    libdbus-1-dev/stable,stable-security,now 1.12.24-0+deb11u1 amd64 [installed,automatic]
    libdbus-glib-1-2/stable,now 0.110-6 amd64 [installed,automatic]
    libdbusmenu-glib4/stable,now 18.10.20180917~bzr492+repack1-2 amd64 [installed,automatic]
    libdbusmenu-gtk3-4/stable,now 18.10.20180917~bzr492+repack1-2 amd64 [installed,automatic]


Xfce Desktop Environment

    $ apt list --installed | grep libxfce4

    WARNING: apt does not have a stable CLI interface. Use with caution in scripts.

    libxfce4panel-2.0-4/stable,now 4.16.2-1 amd64 [installed,automatic]
    libxfce4ui-2-0/stable,now 4.16.0-1 amd64 [installed,automatic]
    libxfce4ui-common/stable,stable,now 4.16.0-1 all [installed,automatic]
    libxfce4ui-utils/stable,now 4.16.0-1 amd64 [installed,automatic]
    libxfce4util-bin/stable,now 4.16.0-1 amd64 [installed,automatic]
    libxfce4util-common/stable,stable,now 4.16.0-1 all [installed,automatic]
    libxfce4util7/stable,now 4.16.0-1 amd64 [installed,automatic]
