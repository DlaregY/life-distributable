package calendar

import (
	"testing"
	"time"
)

func TestRepeatPreservesSay(t *testing.T) {
	c, th, _ := newTest(t)
	n := &cardNfy{}
	th.Notifier = n
	now := time.Date(2026, 10, 5, 8, 0, 0, 0, time.Local)
	c.Now = func() time.Time { return now }
	_, err := c.Add(Item{Title: "Wash your face", Kind: "owner", Day: "2026-10-05", At: "13:00", Repeat: "daily", Say: "Hey, face time."})
	if err != nil {
		t.Fatal(err)
	}
	now = now.Add(5*time.Hour + time.Minute)
	c.Tick()
	rows, err := c.List("2026-10-06", "2026-10-06", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("rows=%v err=%v", rows, err)
	}
	if rows[0].Say != "Hey, face time." {
		t.Errorf("next occurrence Say=%q", rows[0].Say)
	}
	now = now.AddDate(0, 0, 1)
	c.Tick()
	if len(n.said) != 2 {
		t.Fatalf("spoken messages = %q, want two notifications", n.said)
	}
	for i, say := range n.said {
		if say != "Hey, face time." {
			t.Errorf("occurrence %d said %q", i+1, say)
		}
	}
}
