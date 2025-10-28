# SDC demo

![sdc logo](https://docs.sdcio.dev/assets/logos/SDC-transparent-withname-100x133.png)

This repository is part of Schema Driven Configuration (SDC)

The paradigm of schema-driven API approaches is gaining increasing popularity as it facilitates programmatic interaction with systems by both machines and humans. While OpenAPI schema stands out as a widely embraced system, there are other notable schema approaches like YANG, among others. This project endeavors to empower users with a declarative and idempotent method for seamless interaction with API systems, providing a robust foundation for effective system configuration."

## Usage

```
server01:sdcio-demo$ go run . -h
NAME:
   sdcio-demo - A new cli application

USAGE:
   sdcio-demo [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --all, -l                     run all demos (default: false)
   --auto, -a                    run the demo in automatic mode, where every step gets executed automatically (default: false)
   --dry-run                     run the demo and only prints the commands (default: false)
   --no-color                    run the demo and output to be without colors (default: false)
   --auto-timeout auto, -t auto  the timeout to be waited when auto is enabled (default: 1s)
   --with-breakpoints            breakpoint (default: false)
   --continue-on-error           continue if there a step fails (default: false)
   --continuously, -c            run the demos continuously without any end (default: false)
   --hide-descriptions, -d       hide descriptions between the steps (default: false)
   --immediate, -i               immediately output without the typewriter animation (default: false)
   --skip-steps value, -s value  skip the amount of initial steps within the demo (default: 0)
   --shell value                 define the shell that is used to execute the command(s) (default: bash)
   -0, --setup_environment       setup environment (default: false)
   -1, --install_sdcio           installing sdcio (default: false)
   -2, --basic_usage             basic usage demo (default: false)
   -3, --destroy_basic_usage     destroy basic usage (default: false)
   -4, --destroy_environment     destroy environment (default: false)
   --help, -h                    show help
```

## Join us

Have questions, ideas, bug reports or just want to chat? Come join [our discord server](https://discord.com/channels/1240272304294985800/1311031796372344894).

## License and Code of Conduct

Code is under the [Apache License 2.0](LICENSE), documentation is [CC BY 4.0](LICENSE-documentation).

The SDC project is following the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md). More information and links about the CNCF Code of Conduct are [here](code-of-conduct.md).
