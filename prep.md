# Postman Interview Prep — Go (4‑Hour Plan)

This is a **focused, high‑ROI prep plan** for a Postman coding + system design interview using **Go**. The goal is not mastery — it’s *clarity, confidence, and calm execution*.

---

## How to Use This Document

* You have **~4 hours** total
* Follow the sections **top to bottom**
* Skip anything marked *Optional* if you’re short on time
* Optimize for **boring, readable, idiomatic Go**

> Interview success ≠ finishing the problem
>
> Interview success = how you think, communicate, and adapt

---

# PART 1: CODING SESSION (≈ 2.5 hours)

## 1. Backend Setup (30–40 mins)

### Goal

Have a backend that **boots instantly** and supports basic CRUD.

### Recommended Stack

* Go
* `net/http`
* SQLite
* `database/sql`

### Must‑Have Checklist

* [ ] Server starts in <10 seconds
* [ ] One resource (e.g. `Item`, `User`, `Collection`)
* [ ] CRUD endpoints:

  * `POST /resource`
  * `GET /resource`
  * `GET /resource/{id}`
  * `PUT` or `PATCH /resource/{id}`
  * `DELETE /resource/{id}`
* [ ] DB connection working
* [ ] Meaningful HTTP status codes
* [ ] JSON request/response

### Explicitly Say in Interview

> “I’m keeping the setup minimal so we can focus on the problem itself.”

---

## 2. HTTP Basics (`net/http`) (30 mins)

### Know Cold

* `http.HandleFunc`
* `http.ListenAndServe`
* `http.ResponseWriter`
* `*http.Request`
* Reading request body
* Writing JSON responses

### Signals They Look For

* Explicit status codes
* Clear handler boundaries
* No hidden magic

### Reading

* Go Blog: *Writing Web Applications*
* Go docs: `net/http`

---

## 3. JSON + Structs (20 mins)

### Must Know

* `json.NewDecoder(r.Body).Decode(&obj)`
* `json.Marshal`
* Struct tags

### Watch Out For

* Ignoring decode errors
* Forgetting to handle invalid input

### Say Out Loud

> “I’ll validate input lightly here and return a clear error if it’s malformed.”

---

## 4. Context Usage (High Signal) (20 mins)

### Must Know

* `r.Context()`
* Passing context to DB calls
* Request cancellation

### Why It Matters

Shows production thinking without over‑engineering.

### Say Explicitly

> “I’ll pass context through so requests can be cancelled properly.”

### Reading

* Go Blog: *Contexts and Cancellation*

---

## 5. Error Handling (Very Important) (30 mins)

### Patterns to Use

* Fail fast
* Return explicit HTTP errors
* No panics

### What They Like

* Clear error messages
* Consistent status codes
* Calm handling of failure

### Say Out Loud

> “I’m handling errors early so the happy path stays readable.”

---

## 6. Database: SQLite + `database/sql` (45–60 mins)

### Why SQLite

* Zero setup
* Real SQL
* Idiomatic Go
* Predictable behavior

### Must Know

* `sql.Open`
* `db.ExecContext`
* `db.QueryRowContext`
* `Scan`
* Prepared statements

### Avoid

* ORMs (GORM)
* NoSQL unless you’re extremely fluent

### Say Out Loud

> “I’m using plain SQL here for clarity and control.”

---

## 7. Project Structure (15 mins)

### Keep It Simple

Example:

* `main.go`
* `handlers.go`
* `models.go`
* `storage.go`

### Say Explicitly

> “I’m keeping this flat for readability in an interview setting.”

---

## 8. Testing (Optional Bonus) (20 mins)

### You Don’t Need to Write Tests Live

But you *should talk about them*.

### Know

* `httptest.NewRecorder`
* Table‑driven tests

### Say

> “I’d add table‑driven tests around handlers and storage.”

---

# PART 2: SYSTEM DESIGN SESSION (≈ 1–1.5 hours)

## 1. Use This Structure Every Time

1. Clarify requirements
2. Identify users & use cases
3. Define non‑goals
4. High‑level architecture
5. Data model
6. Scale & failure modes
7. Tradeoffs

> Do **not** jump into architecture immediately

---

## 2. Practice ONE Postman‑Relevant Problem (30 mins)

Examples:

* API rate limiting
* API usage analytics
* Collaboration on API collections
* Environment variable management

Focus on:

* Who is the user?
* What matters most?
* Where can it fail?

---

## 3. Metrics & Risks (10–15 mins)

### Metrics to Mention

* Latency
* Error rate
* Throughput
* Cost

### Common Risks

* Data consistency
* Traffic spikes
* Abuse / misuse

---

# LANGUAGE & FRAMEWORK DECISIONS

## Go — When It’s the Right Choice

Stick with Go if:

* You’re comfortable with `net/http`
* You can code without freezing
* You can explain what you write

## Framework Guidance

| Tool       | Verdict             |
| ---------- | ------------------- |
| `net/http` | ✅ Best              |
| Chi        | ✅ Acceptable        |
| Gin        | ⚠️ Slightly magical |
| Fiber      | ❌ Not idiomatic     |

---

## Should You Switch Languages?

Switch **only** if you’re significantly faster elsewhere.

> A clean Express or FastAPI app beats a shaky Go app.

---

# INTERVIEW BEHAVIOR (CRITICAL)

## Do

* Talk continuously
* Ask clarifying questions
* Treat interviewer as a collaborator
* Say what you’d improve with more time

## Don’t

* Go silent
* Over‑engineer
* Apologize excessively
* Panic if stuck

---

# PHRASES THAT SCORE POINTS

Use these naturally:

* “Let me start with a simple version first.”
* “I’ll refactor this if we have time.”
* “I’m optimizing for clarity over completeness.”
* “Let me restate the problem to make sure I understand it.”

---

# FINAL 15‑MINUTE CHECKLIST

* [ ] Backend boots
* [ ] CRUD works
* [ ] Copilot disabled
* [ ] IDE shortcuts ready
* [ ] Whiteboard link ready
* [ ] Water nearby

---

## Final Reminder

This interview is **not adversarial**.

They want to know:

> *“Would I enjoy building things with this person?”*

Clarity > cleverness.
