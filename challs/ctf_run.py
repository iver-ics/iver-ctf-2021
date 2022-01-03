#!/usr/bin/python3

"""CTF challenges runner

This script builds, compiles, and optionally deploys to docker-compose all the
challenges located in the current working directory.

Challenge directories are expected to include a `challenge.yml` file defining
the different challenge attributes, such as its name, flag, score, etc.

What this script does:

    - updates the `src/flag.txt` file if it exists based on the flag specified in
      `challenge.yml`

    - run the `src/Makefile`, if it exists

    - build the Dockerfile and deploy it via docker-compose if
      `docker-compose.yml` exists

This script is written and maintained by Iver.
"""

import argparse
import json
import os
import re
import sys
import threading
import yaml
import yaml

dry_run = False

def build(challenge):
    print()
    print(f"## Running {challenge}")
    flag = ''
    pub_port = 0
    with open(f"{challenge}/challenge.yml") as f:
        challenge_config = yaml.safe_load(f)
        flag = challenge_config['flags'][0]
        print("FLAGS:", flag)

    if os.path.exists(f"{challenge}/src/flag.txt"):
        print(f"Updating flag: {challenge}")
        with open(f"{challenge}/src/flag.txt", "w") as f:
            f.write(flag)

    if os.path.exists(f"{challenge}/src/Makefile"):
        print(f"Building code: {challenge}")
        os_run(f"make -C ./{challenge}/src")

    if os.path.exists(f"{challenge}/src/{challenge}"):
        os_run(f"mkdir -pv ./{challenge}/bin")
        os_run(f"cp ./{challenge}/src/{challenge} ./{challenge}/bin/program")

    if os.path.exists(f"{challenge}/docker-compose.yml"):
        print(f"Recreating Docker container(s): {challenge}")
        os_run(f"sudo -E docker-compose -f {challenge}/docker-compose.yml up --build -d")

        ports = get_docker_compose_ports(f"{challenge}/docker-compose.yml")
        for port in ports:
            print(f"PORT: {port} (from {challenge}/docker-compose.yml)")


def get_docker_compose_ports(filename):
    with open(filename) as f:
        compose = yaml.safe_load(f)
        for svcName in compose['services']:
            svc = compose['services'][svcName]
            if 'ports' in svc:
                for port in svc['ports']:
                    yield port.split(':', 1)[0]


def get_challenges():
    return [f"{name}" for name in os.listdir(f".") if name[0] != '.' and os.path.isdir(
        f"./{name}") and os.path.exists(f"./{name}/challenge.yml")]


def os_run(cmd):
    if dry_run:
        print(f"dry-run: $ {cmd}")
    else:
        print(f"$ {cmd}")
        exitcode = os.system(cmd)
        if exitcode != 0:
            raise RuntimeError(f"Non-zero exit code: {exitcode}: {cmd}")


def run_challanges_builds(challanges, parallel):
    if parallel:
        jobs = []
        for challenge in challenges:
            jobs.append(threading.Thread(target=build, args=(challenge, )))

        for job in jobs:
            job.start()

        for job in jobs:
            job.join()
    else:
        for challenge in challenges:
            build(challenge)


if __name__ == "__main__":
    challenges = get_challenges()

    parser = argparse.ArgumentParser()
    parser.add_argument("--dry-run", action="store_true", help="only validate")
    parser.add_argument("--parallel", action="store_true", help="run all challenges at the same time")
    parser.add_argument("chall", nargs="+", choices=["all", *challenges], help="challenge names to run. specify 'all' for all")
    args = parser.parse_args()
    dry_run = args.dry_run

    all_challenges = "all" in args.chall
    if all_challenges and len(args.chall) > 1:
        print("Cannot specify 'all' and specific challanges")
        exit(1)

    if not all_challenges:
        challenges = [x for x in challenges if x in args.chall]

    print(f"{challenges=}")
    if args.dry_run:
        print("Running a dry-run")

    run_challanges_builds(challenges, args.parallel)

    print()
    print("Synchronized successfully!")
