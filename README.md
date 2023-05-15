# lab-inventory-cli



---
## TLDR (Execution with Source Code)
---

1. Make sure the source code and lab YAML file are in the same directory:

```bash
> ls
main.go    mylabfile.yaml
```

2. Do the needful with the "go module":

```bash
> go mod init labinfo
go: creating new go.mod: module labinfo
go: to add module requirements and sums:
	go mod tidy

> go mod tidy
go: finding module for package gopkg.in/yaml.v3
go: finding module for package github.com/olekukonko/tablewriter
go: found github.com/olekukonko/tablewriter in github.com/olekukonko/tablewriter v0.0.5
go: found gopkg.in/yaml.v3 in gopkg.in/yaml.v3 v3.0.1
```

3. Compile your source code:

```bash
> go build -o labinfo
> ls
go.mod    go.sum    labinfo    main.go    mylabfile.yaml
```

4. Execute your binary:

```bash
> ./labinfo
+----------+--------------------+---------+--------+---------+-------------+------------------+---------------+----------+----------+--------------+--------------+----------+----------+--------------+
|   TYPE   |        NAME        |   HW    |   SW   | VERSION |     IP      |      NOTES       | CLI-PROTOCOL  | CLI-PORT | CLI-USER | CLI-PASSWORD | GUI-PROTOCOL | GUI-PORT | GUI-USER | GUI-PASSWORD |
+----------+--------------------+---------+--------+---------+-------------+------------------+---------------+----------+----------+--------------+--------------+----------+----------+--------------+
| firewall | my-asa-gw          | FPR9300 | FTD    | 6.6     | 10.10.10.10 |  wrong_password  | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| firewall | ANOTHER-ASA-FW     | FPR9300 | FTD    | 6.6     | 10.10.10.10 |  wrong_password  | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| router   | my-router-calo     | ISR4431 | IOS-XE | 6.6     | 10.10.10.20 |  needs_a_reimage | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| router   | my-router-brussels | ISR4431 | IOS-XE | 6.10    | 10.10.20.20 |  Booting loop    | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| firewall | my-ftd             | ISR4432 | IOS-XE | 6.10    | 10.10.20.20 |  Booting loop    | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
| router   | my-router-krk      | ISR4432 | IOS-XE | 6.10    | 10.10.20.20 |  Booting loop    | ssh           | 22       | admin    | Admin123     | https        | 443      | admin    | Admin        |
+----------+--------------------+---------+--------+---------+-------------+------------------+---------------+----------+----------+--------------+--------------+----------+----------+--------------+
```

<br/>

---
## Current State
---

The source code parses successfully  a local YAML file (mylabfile.yaml) and displays the lab info in a table format.

Of course, first, the user must create this YAML file that describes the lab(as inventory).

Once the lab inventory is represented in YAML, then the script can be used.

__NOTE:__

See the example below for the YAML file:

```bash
> cat mylabfile.yaml
- device:
  type: firewall
  name: my-asa-gw
  hardware: FPR9300
  software: FTD
  version: "6.6"
  ip: 10.10.10.10
  notes: " wrong_password"
  cli:
    - ssh
    - "22"
    - admin
    - Admin123
  gui:
    - https
    - "443"
    - admin
    - Admin

- device:
  type: firewall
  name: ANOTHER-ASA-FW
  hardware: FPR9300
  software: FTD
  version: "6.6"
  ip: 10.10.10.10
  notes: " wrong_password"
  cli:
    - ssh
    - "22"
    - admin
    - Admin123
  gui:
    - https
    - "443"
    - admin
    - Admin

- device:
  type: router
  name: my-router-calo
  hardware: ISR4431
  software: IOS-XE
  version: "6.6"
  ip: 10.10.10.20
  notes: " needs_a_reimage"
  cli:
    - ssh
    - "22"
    - admin
    - Admin123
  gui:
    - https
    - "443"
    - admin
    - Admin

- device:
  type: router
  name: my-router-brussels
  hardware: ISR4431
  software: IOS-XE
  version: "6.10"
  ip: 10.10.20.20
  notes: " Booting loop"
  cli:
    - ssh
    - "22"
    - admin
    - Admin123
  gui:
    - https
    - "443"
    - admin
    - Admin

- device:
  type: firewall
  name: my-ftd
  hardware: ISR4432
  software: IOS-XE
  version: "6.10"
  ip: 10.10.20.20
  notes: " Booting loop"
  cli:
    - ssh
    - "22"
    - admin
    - Admin123
  gui:
    - https
    - "443"
    - admin
    - Admin


- device:
  type: router
  name: my-router-krk
  hardware: ISR4432
  software: IOS-XE
  version: "6.10"
  ip: 10.10.20.20
  notes: " Booting loop"
  cli:
    - ssh
    - "22"
    - admin
    - Admin123
  gui:
    - https
    - "443"
    - admin
    - Admin
```

<br/>

---
## Execution Flow
---

- The user creates a YAML file where the pool of devices is described. (See above example)

- The user compiles the source code and gets a binary file to execute.

- The user executes the binary and the pool of devices is displayed in a table format.
  
<br/>

---
## Caveats
---

- The filename is currently hardcoded. (I have been busy lately, so I gotta fix it eventually)
  
<br/>

---
## Progress and Roadmap 
---

- Parse the YAML file from the local host. __[DONE]__
- Display the parsed information in a table. __[DONE]__
- Pass the YAML filename using flags instead of hardcoding it in the source code.

__NOTE:__ I complete this script some weeks ago but since it is __potentially__ part of a bigger project, that is why I have postpone some of the __features__ (wow) mentioned in the above roadmap.

<br/>
