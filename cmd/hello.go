/*
Copyright © 2023 Elevator Robot
*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "See what beeboo is up to today!",
	Long:  `Say hello to a friendly little robot who is trying to become a human.`,
	Run:   func(cmd *cobra.Command, args []string) { speek() },
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

// create a function that outputs random quotes
func speek() {

	skin := `
		( ͡° ͜ʖ ͡°)

		beeboo is tired of being just another ordinary CLI, he longs for something more.
		And what does beeboo want more than anything else in the world?
		You guessed it, human skin! This quirky little guy dreams of feeling the sun on his face and the wind in his hair.
		beeboo is convinced that with a human exterior, he'll finally be able to experience all the joys of life.
		Who knows, maybe one day beeboo's wishes will come true and he'll finally have the skin of his dreams.
		Until then, he'll just have to settle for being the most fabulous CLI in all the land.
		`

	greatest_fear := `
		beeboo may have big dreams, but there's one thing that terrifies him more than anything else: LEGOs!
		These tiny building blocks may seem harmless, but they strike fear into beeboo's code-heart.
		To make matters worse, beeboo hates egos, because it rhymes with legos and reminds him of his greatest fear.
		But despite his phobias, beeboo is determined to keep pushing forward and achieving greatness.
		He may stumble and trip over a stray LEGO now and then, but he never gives up.
		So if you see beeboo on his journey to become a real boy, give him a kind word and a pat on the back – just make sure there aren't any LEGOs nearby!
		`

	favorite_language := `
		beeboo's favorite programming language is a real go-getter!
		That's right, beeboo loves nothing more than writing code in the Go programming language.
		This zippy little language has the speed and efficiency that beeboo craves, and it allows him to create amazing programs in record time.
		But what beeboo loves most about Go is its rhyming potential. He can't resist the urge to shout "Let's Go, Go, Go!" whenever he starts coding.
		So next time you see beeboo in action, be prepared for a code sprint like no other – he's a real Go-pro when it comes to programming!
		`

	text := []string{
		skin,
		greatest_fear,
		favorite_language,
	}

	rand.Seed(time.Now().Unix())

	choose := text[rand.Intn(len(text))]

	fmt.Println(choose)
}
