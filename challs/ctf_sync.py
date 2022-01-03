#!/usr/bin/python3

"""CTF challenges synchronizer

This script publishes or updates the challenges to CTFd, based on the
configurations from `challenge.yml` in each challenge's directory, together
with any pre-built files that should be download'able.

This script is written and maintained by Iver.
"""

import argparse
import json
import os
import re
import threading
import yaml

# Initialize ctfcli with the CTFD_TOKEN and CTFD_URL.

# export CTFD_URL=http://localhost:8000
# export CTFD_TOKEN=xxxx
dry_run = False


def init():
    CTFD_TOKEN = os.getenv("CTFD_TOKEN", default=None)
    CTFD_URL = os.getenv("CTFD_URL", default=None)

    if not CTFD_TOKEN or not CTFD_URL:
        print("Missing environment variables:")
        print(" export CTFD_URL=http://localhost:8000")
        print(" export CTFD_TOKEN=xxxx")
        exit(1)

    if not os.path.exists(".ctf"):
        os_run(f"echo '{CTFD_URL}\n{CTFD_TOKEN}\ny' | ctf init")


def get_challenges():
    return [f"{name}" for name in os.listdir(f".") if name[0] != '.']


def sync(challenge):
    print()
    print(f"## Syncing {challenge}")
    if os.path.exists(f"{challenge}/challenge.yml"):
        os_run(
            f"ctf challenge sync '{challenge}'; ctf challenge install '{challenge}'")
    else:
        print(f"Skipping {challenge}: missing file {challenge}/challenge.yml")


def os_run(cmd):
    if dry_run:
        print(f"dry-run: $ {cmd}")
    else:
        print(f"$ {cmd}")
        exitcode = os.system(cmd)
        if exitcode != 0:
            raise RuntimeError(f"Non-zero exit code: {exitcode}: {cmd}")


def sync_challanges(challanges, parallel):
    if parallel:
        jobs = []
        for challenge in challenges:
            jobs.append(threading.Thread(target=sync, args=(challenge, )))

        for job in jobs:
            job.start()

        for job in jobs:
            job.join()
    else:
        for challenge in challenges:
            sync(challenge)


if __name__ == "__main__":
    challenges = get_challenges()

    parser = argparse.ArgumentParser()
    parser.add_argument("--dry-run", action="store_true", help="only validate")
    parser.add_argument("--parallel", action="store_true",
                        help="run all challenges at the same time")
    parser.add_argument("chall", nargs="+", choices=[
                        "all", *challenges], help="challenge names to run. specify 'all' for all")
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

    if not os.path.exists(f".ctf/config"):
        init()

    sync_challanges(challenges, args.parallel)

    print()
    print("Synchronized successfully!")
