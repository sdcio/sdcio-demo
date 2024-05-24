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