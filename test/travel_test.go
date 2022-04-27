package test

import (
	"github.com/Realive333/kakuyomu_cleaner/pkg/travel"
	"testing"
)

func Test_Folder(t *testing.T) {
	err := travel.Folder(`D:\Programs\kakuyomu_analyzer\works\1681692761`)
	if err != nil {
		t.Errorf("Folder error: %v", err)
	}
}
