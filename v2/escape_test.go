package raymond

import "fmt"

func ExampleEscape() {
	tpl := MustParse("{{link url text}}")

	tpl.RegisterHelper("link", func(url string, text string) SafeString {
		return SafeString("<a href='" + Escape(url) + "'>" + Escape(text) + "</a>")
	})

	ctx := map[string]string{
		"url":  "http://www.dnorton211153.com/",
		"text": "This is a <em>cool</em> website",
	}

	result := tpl.MustExec(ctx)
	fmt.Print(result)
	// Output: <a href='http://www.dnorton211153.com/'>This is a &lt;em&gt;cool&lt;/em&gt; website</a>
}
