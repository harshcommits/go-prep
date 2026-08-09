# Interview Reset Roadmap

It is completely valid to feel incredibly frustrated right now. You are a senior engineer with serious domain expertise (backend, cloud, platform, security), but you are currently trapped in a role that doesn't utilize your skills, relying on AI to do the heavy lifting, and getting ambushed by broken interview loops.

Let’s address the reality of your situation:

- A **"work-cation"** is just working with a better view.
- It doesn't cure burnout because your brain is still carrying the exact same cognitive load and context.
- Brain fog is the natural biological response to chronic low-grade stress (job uncertainty) combined with high-dopamine distractions (social media) and zero creative fulfillment at work.

You don't need to grind harder; you need a systemic reset. Here is a three-part game plan to clear the fog, rebuild your engineering stamina, and reach a permanent **maintenance mode** for interviews.

## Phase 1: Clearing the Brain Fog (Cognitive Reset)

Your brain is a muscle that has lost its endurance for deep work. In a 45-minute interview, you need sustained, uninterrupted focus. Right now, social media and AI autocomplete are training your brain for 10-second attention spans.

### Dopamine Detox (Strict Boundaries)

- Social media scrolling is fracturing your attention span.
- Delete the apps from your phone for 14 days.
- If you need a break, read a physical book, take a walk without your phone, or stare at a wall.
- Retrain your brain to be comfortable with boredom so it can sustain focus during a 45-minute coding round.

### Change Your Relationship with AI

- You are currently using AI as autocomplete.
- This bypasses the neural pathways used for problem-solving.
- New rule: For the next month, turn off Copilot/Cursor autocomplete.
- Write the logic yourself.
- Use AI strictly as a code reviewer:
  - "I wrote this. What edge cases did I miss? How can I optimize it?"
  - or a sparring partner:
  - "Explain the Raft consensus algorithm to me."
- Be the architect; stop letting AI do the typing.

### Decouple Your Identity from Your Day Job

- Accept that your GitHub Admin job is currently just a paycheck.
- Stop expecting motivation from it.
- Your motivation is getting out of it.

## Phase 2: Fixing the Skill Atrophy at Work

You need to keep your coding skills sharp without adding 4 hours of work to your evenings. The trick is to turn your GitHub Admin job into a Backend/Platform Engineering job.

### Automate Everything with Code

- Stop clicking through the GitHub UI or writing basic YAML by hand.
- If you need to audit repository permissions, write a Go app that hits the GitHub GraphQL API, processes the data concurrently, and outputs a report.

### Over-engineer (Just a little)

- Build an internal tool for your admin work.
- Wrap it in a Docker container.
- Deploy it to AWS/GCP.
- Add a Redis cache for API rate-limiting.
- Write unit tests.

This turns your day job into practice for:

- System Design
- Low-Level Design (LLD)
- Backend engineering

## Phase 3: The "DSA Threshold" & Interview Strategy

You hate grinding LeetCode because memorizing 500 random problems is a waste of a senior engineer's time. You don't need to memorize problems; you need to internalize patterns.

### Reaching the DSA Threshold

- Ditch random LeetCode grinding.
- Focus exclusively on the **NeetCode 150** or **Blind 75**.
- There are only about 15 core patterns:
  - Sliding Window
  - Two Pointers
  - Fast & Slow Pointers
  - BFS/DFS
  - Top K Elements
  - etc.

Once you can recognize a prompt and instantly say, "This asks for the shortest path in an unweighted grid, so it's BFS," you have reached the threshold. From then on, you only need to do 1–2 problems a week on a Sunday to keep the rust off.

### Mastering the LLD/HLD Pivot

- Since you enjoy this more, make it your playground.
- Spend weekends building tiny, toy versions of distributed systems.
- Build a rudimentary load balancer in Go.
- Build a simple key-value store that writes to a write-ahead log (WAL).

This builds deep intuition that no textbook can teach.

### The "Anti-Ambush" Recruiter Contract

- Take control of the interview pipeline.
- Recruiters often don't know the difference between DSA and LLD.

Send this exact email before every technical round:

> "To ensure I respect the interviewer's time and prepare appropriately, could you confirm the format of this round? Specifically, is this an Algorithmic/Data Structures (LeetCode-style) round, or a Low-Level Design/Machine Coding round focusing on Object-Oriented Principles and extensibility?"

Force them to clarify in writing.
