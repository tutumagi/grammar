package grammar

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Filename string      `json:"-"`
	Source   []string    `json:"source"`
	Expect   *ExpectData `json:"expect"`
}

type ExpectData struct {
	Production map[Symbol][]*Production `json:"production"`
	FirstSet   map[Symbol]SymbolSet     `json:"first_set"`
	FollowSet  map[Symbol]SymbolSet     `json:"follow_set"`
}

func (a *ExpectData) UnmarshalJSON(bb []byte) error {
	type mockExpectData struct {
		Production []string `json:"production"`
		FirstSet   []string `json:"first_set"`
		FollowSet  []string `json:"follow_set"`
	}
	mockData := &mockExpectData{}
	err := json.Unmarshal(bb, mockData)
	if err != nil {
		return err
	}

	a.Production = make(map[Symbol][]*Production)
	for _, line := range mockData.Production {
		lhs, productions := makeLineProduction(line)
		a.Production[lhs] = append(a.Production[lhs], productions...)
	}
	a.FirstSet = make(map[Symbol]SymbolSet)
	for _, line := range mockData.FirstSet {
		items := strings.Split(line, "=")
		a.FirstSet[strings.TrimSpace(items[0])] = newSymbolSet(strings.Split(strings.TrimSpace(items[1]), " ")...)
	}
	a.FollowSet = make(map[Symbol]SymbolSet)
	for _, line := range mockData.FollowSet {
		items := strings.Split(line, "=")
		a.FollowSet[strings.TrimSpace(items[0])] = newSymbolSet(strings.Split(strings.TrimSpace(items[1]), " ")...)
	}

	return nil
}

func TestMakeFirstFollow(t *testing.T) {
	expectdatas := readTestData()

	for _, testData := range expectdatas {
		t.Run(testData.Filename, func(t *testing.T) {
			g := NewGrammar(strings.Join(testData.Source, "\n"))
			assert.Equal(t, testData.Expect.Production, g.productions)

			g.makeFirstSet()
			assert.Equal(t, testData.Expect.FirstSet, g.firstSet)

			g.makeFollowSet()
			assert.Equal(t, testData.Expect.FollowSet, g.followSet)

			g.makePredict()
			g.predictTable.dump()
		})
	}
}

func readTestData() []*TestData {
	expectdatas := []*TestData{}

	err := filepath.WalkDir("./testdata", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			bb, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			tmp := &TestData{
				Filename: filepath.Base(path),
			}
			err = json.Unmarshal(bb, tmp)
			if err != nil {
				return err
			}

			expectdatas = append(expectdatas, tmp)
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return expectdatas
}
