# Postman Go Interview — 1‑Page Panic Sheet

Read this **10–15 minutes before the interview**. Nothing more.

---

## 🧠 First Principles (Calm Down)

* This is **not** a trick interview
* Completion ≠ success
* Silence is worse than wrong code
* Talking clearly > typing fast

> If stuck: **pause, restate, simplify**

---

## 🗣️ Phrases That Buy You Time (Use Freely)

* “Let me restate the problem to make sure I understand it.”
* “I’ll start with a simple version first.”
* “I’m optimizing for clarity over completeness.”
* “If we had more time, I’d refactor this.”
* “Does this direction seem reasonable?”

These are **signals of senior thinking**, not weakness.

---

## 🧱 Default Coding Plan (Always Works)

1. Clarify requirements
2. Define the data model
3. Create naive version
4. Handle errors
5. Polish if time allows

Say this out loud before typing.

---

## 🧩 Go HTTP Handler Skeleton (Burn Into Memory)

```go
func(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    switch r.Method {
    case http.MethodGet:
        // logic
    case http.MethodPost:
        // logic
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
```

If you forget syntax: **describe what you’re about to write first**.

---

## 📦 JSON Decode / Encode (Most Common Failure Point)

```go
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "invalid JSON", http.StatusBadRequest)
    return
}

json.NewEncoder(w).Encode(resp)
```

If this fails: say *“I’ll handle malformed input explicitly.”*

---

## 🧵 Context (Easy Bonus Points)

```go
ctx := r.Context()
```

Say:

> “This ensures cancellation if the client disconnects.”

---

## 🗄️ SQL Patterns (Enough for Interview)

```go
row := db.QueryRowContext(ctx, query, id)
err := row.Scan(&obj.ID, &obj.Name)
```

```go
_, err := db.ExecContext(ctx, query, args...)
```

Always check `err`. No panics.

---

## 🚦HTTP Status Codes (Just These)

* `200 OK`
* `201 Created`
* `204 No Content`
* `400 Bad Request`
* `404 Not Found`
* `500 Internal Server Error`

Don’t overthink this.

---

## 🧪 If Asked About Tests (Without Writing Them)

Say:

> “I’d add table-driven tests using `httptest` for handlers and storage.”

That’s enough.

---

## 🧯 When You Freeze (Most Important Section)

Do **one** of these immediately:

1. Restate the problem
2. Ask a clarifying question
3. Implement the naive version

Never sit silently.

---

## 🏗️ System Design Emergency Structure

1. Users & use cases
2. Non-goals
3. High-level components
4. Data model
5. Scale & risks
6. Tradeoffs

Boxes + arrows. That’s it.

---

## 🧘 Final Reminder

They are evaluating:

* How you think
* How you communicate
* Whether they’d enjoy working with you

You already know enough.

**Slow is smooth. Smooth is fast.**
