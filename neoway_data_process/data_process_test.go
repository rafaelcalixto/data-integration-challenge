package data_process

import (
    "testing"
)

var (
    test     string
    answer   row
    template row
)

func TestNormalize_data(t *testing.T) {
    // First test case
    template.comp = "NEOWAY"
    template.zip = "00001"
    template.website = "www.neoway.com"
    t.Run("First test case", func(t *testing.T) {
        test = "neoway;1;WWW.NEoWAY.COM"
        answer.comp, answer.zip, answer.website = Normalize_data( test, true )
        if answer.comp != template.comp { t.Errorf("Upper case Error") }
        if answer.zip != template.zip { t.Errorf("Error on the zip handle") }
        if answer.website != template.website {
            t.Errorf("Error on the web site lower case")
        }
    })

    // Second test case
    template.comp = "MCDONALDS"
    template.zip = ""
    t.Run("Second test case", func(t *testing.T) {
        test = "McDonald's;"
        answer.comp, answer.zip, _ = Normalize_data(test, false)
        if answer.comp != template.comp { t.Errorf("Single quotation error") }
        if answer.zip != template.zip { t.Errorf("Error on the null zip") }
    })

    // Thrid test case
    template.comp = "CA"
    template.zip = "89526"
    t.Run("Thrid test case", func(t *testing.T) {
        answer.comp, _, _ = Normalize_data("C&A;89526", false)
        if answer.comp != template.comp { t.Errorf("Error on the character &") }
    })
}
