package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	var (
		oldMap map[string]interface{}
		newMap map[string]interface{}
	)

	Describe("Merge Map", func(){
		It("returns the original map when given empty newMap", func(){
			oldMap = map[string]interface{}{"a": "b"}
			newMap = map[string]interface{}{}
			Expect(MergeMap(oldMap, newMap)).To(Equal(oldMap))
		})

		It("returns the new map when given empty oldMap", func(){
			oldMap = map[string]interface{}{}
			newMap = map[string]interface{}{"a":"b"}
			Expect(MergeMap(oldMap, newMap)).To(Equal(newMap))
		})

		It("updates the original map with new information", func(){
			oldMap = map[string]interface{}{"a": "b"}
			newMap = map[string]interface{}{"b": "c"}
			returnMap := map[string]interface{}{"a": "b", "b": "c"}
			Expect(MergeMap(oldMap, newMap)).To(Equal(returnMap))
		})

		It("replaces the oldmap if the key exists in newmap", func(){
			oldMap = map[string]interface{}{"a": "b", "b": "d"}
			newMap = map[string]interface{}{"b": "c"}
			returnMap := map[string]interface{}{"a": "b", "b": "c"}
			Expect(MergeMap(oldMap, newMap)).To(Equal(returnMap))
		})
	})
})
