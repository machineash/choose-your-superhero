# Choose Your Superhero (Phase 1)

Choose what kind of superhero you'd like to see - customize it to your liking.

This Go API lets users pick traits (powers, vices, stressors, personalities, and more) and stores each "hero profile" with persistent memory.

Phase 1 covers core endpoints and file-based storage.

---

## Current Endpoints
| Method | Path | Description |
|---------|-------|--------------|
| GET | `/options` | Returns all selectable traits |
| POST | `/create-hero` | Creates and saves a hero profile |
| GET | `/heroes` | Lists all saved heroes |

---

## Stack
- **Go** (net/http)
- **JSON storage** (saved in `artifacts/heroes.json`)

---

## Run Locally
``` bash
git clone https://github.com/machineash/choose-your-superhero.git
cd choose-your-superhero
go run main.go
```

Then, open:
- http://localhost:8080/options
- http://localhost:8080/heroes

## Next Phases
- Phase 2: Add CI/CD, rule logic (locked environments, validation)
- Phase 3: Introduce AI responses ("closest hero" summary)
- Phase 4+: Integrate Vault + Boundary for secure access
