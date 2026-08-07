# Interview Prep Roadmap — Wed/Thu → Friday Meeting


## Wed Morning — Hashmaps & Graphs (3h)


| Problem | LC # | Link |
|---|---|---|
| Number of Islands | 200 | leetcode.com/problems/number-of-islands |
| Rotting Oranges | 994 | leetcode.com/problems/rotting-oranges |
| LRU Cache | 146 | leetcode.com/problems/lru-cache |
| Subarray Sum Equals K | 560 | leetcode.com/problems/subarray-sum-equals-k |


**Patterns to lock in before you start** (so you're pattern-matching, not deriving from scratch):
- **Islands/Rotting Oranges** = grid BFS/DFS. Rotting Oranges is multi-source BFS (seed the queue with *all* rotten oranges at once, not one).
- **LRU Cache** = hashmap (key → node) + doubly linked list (recency order). In Python/Java you can cheat with `OrderedDict`/`LinkedHashMap` if the interviewer allows it — know both the cheat and the from-scratch version.
- **Subarray Sum = K** = running prefix sum + hashmap of `{prefix_sum: count}`. The trick: check `sum - k` in the map *before* inserting the current sum.


**Crucial rule:** if stuck for more than 25 minutes, look at the solution, understand the mechanics, type it out yourself, and move on. Set a literal timer.


## Wed Afternoon — Linux/Systems (2h)


Command cheat sheet, organized by symptom:


**100% CPU:**
```
top / htop                 # find the PID, note %CPU, load avg (1/5/15m)
top -H -p <PID>             # per-thread CPU inside that process
ps -eLf | grep <PID>        # thread-level view
strace -p <PID> -c          # syscall summary — spinning on a syscall?
perf top -p <PID>           # if available, on-CPU function hotspots
```


**Out of memory:**
```
free -h                     # used/free/available, swap
dmesg -T | grep -i oom      # did the kernel OOM-kill something?
cat /proc/meminfo           # Slab, SReclaimable, Cached breakdown
ps aux --sort=-%mem | head  # top memory consumers
```
User space vs kernel space: user-space leak shows up as RSS growth on one process; kernel-space pressure shows up as growing `Slab`/`SReclaimable` in `/proc/meminfo` or dentry/inode cache growth — a single process's RSS won't explain it.


**Failing to connect to a database:**
```
nc -zv <host> <port>        # can you even open the socket?
ss -tnp | grep <port>       # is anything listening locally / what state
dig <db-host>                # DNS resolving correctly?
traceroute <db-host>        # where does it die on the way
lsof -i :<port>              # what process holds the port/connection
tcpdump -i any host <db-host> and port <port>   # last resort, see the packets
```


**General disk I/O bottleneck:**
```
iostat -xz 1                # %util, await, per-device
lsof +D <path>               # who has files open under a path
```


## Wed Evening — Linked Lists (1h)


- **Reverse a Linked List** — LC 206 (leetcode.com/problems/reverse-linked-list). Iterative: three-pointer (`prev`, `curr`, `next`). Know the recursive version too, it's a common follow-up.
- **Linked List Cycle** — LC 141 (leetcode.com/problems/linked-list-cycle) — Floyd's (slow/fast pointer).
- **Linked List Cycle II** — LC 142 (leetcode.com/problems/linked-list-cycle-ii) — finding the *start* of the cycle (the part people forget: after slow/fast meet, reset one pointer to head and advance both by 1).


Then clock out for the night.


## Thu Morning — Landing Zone & System Design (2.5h)


**Reference:** AWS Landing Zone Accelerator docs (aws.amazon.com/solutions/implementations/landing-zone-accelerator-on-aws) and AWS Control Tower concepts — the current AWS terminology for what the roadmap calls a "Landing Zone."


**What to draw** (on paper, then narrate it out loud once):
```
                    ┌─────────────────┐
                    │  Root/Mgmt Acct │  (billing, SCPs, no workloads)
                    └────────┬─────────┘
        ┌──────────────┬─────┴──────┬───────────────┐
   ┌────▼────┐   ┌──────▼──────┐  ┌──▼───────────┐ ┌─▼─────────┐
   │Log Archive│  │Security/Audit│ │  Networking  │ │ Workload   │
   │  Account  │  │   Account    │ │   Account    │ │ Accounts   │
   └───────────┘  └──────────────┘ │ (Transit GW) │ │(dev/stg/prd)│
                                   └──────┬───────┘ └────────────┘
                                          │ (TGW attachments)
Internet → Route53 → WAF → ALB (public subnet) → app tier (private subnet) → RDS (private subnet, no public IP)
```


Talking points to hit out loud:
- SCPs enforced at the OU level from the root account.
- All logs/CloudTrail centralized in the Log Archive account (immutable, workload accounts can't touch them).
- Transit Gateway as the hub-and-spoke so workload VPCs don't need direct peering.
- Public → private subnet boundary sits at the load balancer — nothing in the private subnet has a public IP or route to an IGW, only a NAT gateway for outbound.


## Thu Afternoon — STAR Stories (1.5h)


Template for each of the 4 prompts — write as bullets, not prose, so it stays punchy in the room:


```
SITUATION (1-2 lines): what was the system/context, what was at stake
TASK: what were you specifically responsible for
ACTION: 3-4 bullets, first person ("I..."), specific commands/decisions/tools
RESULT: quantify if possible (time saved, incidents prevented, metric moved)
```


The 4 stories to write:
1. A time you broke production (and how you fixed it).
2. A time you disagreed with a colleague on a technical choice. *(Make sure RESULT includes how it was resolved and what you'd do differently — interviewers probe this one for self-awareness, not just "I was right.")*
3. A time you had to dive deep into a system you didn't understand.
4. A time you improved an internal tool or process.


## Thu Evening — Muscle Memory & Cutoff (1.5h)


- Re-solve **Rotting Oranges** (BFS) and **Subarray Sum Equals K** (hashmap) completely from scratch, no notes.
- One more pass over the Linux command cheat sheet above — don't relearn, just re-read.
- **Hard stop 8:00 PM.** No studying Thursday night. Watch a movie, relax, get to bed early.
