package rssReader


var Data1 string = `
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Golang Weekly</title>
    <description>A weekly newsletter about the Go programming language</description>
    <link>https://golangweekly.com/</link>
    <item>
      <title>The journey to faster JSON parsing</title>
      <link>https://golangweekly.com/issues/451</link>
      <description>


&lt;table border=0 cellpadding=0 cellspacing=0 align="center" border="0"&gt;
  &lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;div&gt;    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;
&lt;td align="left" style="padding-left: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;#​451 — March 10, 2023&lt;/p&gt;&lt;/td&gt;
&lt;td align="right" style="padding-right: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;&lt;a href="https://golangweekly.com/link/136637/rss" style=" color: #0099b4;"&gt;Unsub&lt;/a&gt;  |  &lt;a href="https://golangweekly.com/link/136638/rss" style=" color: #0099b4;"&gt;Web Version&lt;/a&gt;&lt;/p&gt;&lt;/td&gt;
&lt;/tr&gt;&lt;/table&gt;
    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 12px;  padding-left: 12px;"&gt;&lt;p&gt;The &lt;em&gt;Go Weekly&lt;/em&gt; Newsletter&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/136640/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/vextjc4w4eazkhptfni2.jpg" width="640" style="    line-height: 100%;       "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136640/rss" title="vercel.com" style=" color: #0099b4;    font-size: 1.1em; line-height: 1.4em;"&gt;Why Turborepo is Migrating From Go to Rust&lt;/a&gt;&lt;/span&gt; — &lt;a href="https://golangweekly.com/link/136641/rss" style=" color: #0099b4;   "&gt;Turborepo&lt;/a&gt; is a high performance JavaScript build system built upon Go but.. perhaps not for much longer. Why? It mostly seems subjective, but you might find their arguments interesting.&lt;/p&gt;
  &lt;p&gt;Vercel Engineering Team &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136642/rss" title="go.dev" style=" color: #0099b4;    font-size: 1.05em;"&gt;Code Coverage for Go Integration Tests&lt;/a&gt;&lt;/span&gt; — While Go has had coverage support at the package test level for almost ten years, coverage for integration tests run outside of the &lt;code&gt;go test&lt;/code&gt; mechanism has been non-existent. With Go 1.20, it’s possible to instrument a binary and generate coverage files for both the module code and dependent packages.&lt;/p&gt;
  &lt;p&gt;Than Macintosh &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;a href="https://golangweekly.com/link/136639/rss" style=" color: #0099b4;   "&gt;&lt;img src="https://copm.s3.amazonaws.com/bcc68a3c.png" width="110" height="110" style="padding-top: 12px; padding-left: 12px;     line-height: 100%;    "&gt;&lt;/a&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136639/rss" title="www.ardanlabs.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Go! Experts at Your Service&lt;/a&gt;&lt;/span&gt; — Do you need help filling skill gaps, speeding up development &amp;amp; creating high performing software with Go, Docker, K8s, Terraform and Rust? We’ll help you maximize your architecture, structure, tech-debt and human capital.&lt;/p&gt;
  &lt;p&gt;Ardan Labs Consulting &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136643/rss" title="www.cockroachlabs.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;A Journey to High-Perf JSON Parsing in Go&lt;/a&gt;&lt;/span&gt; — A distributed SQL database needs to maximize data parsing, so CockroachDB’s (very successful) story about doing so is worth the read. The answer came from a somewhat unexpected place and benefits anyone looking to speed up JSON parsing.&lt;/p&gt;
  &lt;p&gt;Yevgeniy Miretskiy (CockroachDB) &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136644/rss" title="go.dev" style=" color: #0099b4;    font-size: 1.05em;"&gt;Getting Started with Multi-Module Workspaces&lt;/a&gt;&lt;/span&gt; — Potentially overlooked in Go 1.18 because they &lt;em&gt;weren’t&lt;/em&gt; generics, workspaces let you tell the Go command you’re writing code in multiple modules at the same time and easily build and run code in those modules.&lt;/p&gt;
  &lt;p&gt;Official Go Docs &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;p&gt;&lt;strong&gt;IN BRIEF:&lt;/strong&gt;&lt;/p&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;⭐️ InfoWorld's Paul Krill reports that &lt;a href="https://golangweekly.com/link/136645/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Go has rejoined the top 10 most popular programming languages&lt;/a&gt; in the latest &lt;a href="https://golangweekly.com/link/136646/rss" style=" color: #0099b4; font-weight: 600;   "&gt;TIOBE index&lt;/a&gt; ranking (for what it's worth).&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136647/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Go 1.20.2 and 1.19.7&lt;/a&gt; have been released with a &lt;code&gt;crypto/elliptic&lt;/code&gt; security fix.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;If you want to speak at this year's GopherCon UK, you'll &lt;a href="https://golangweekly.com/link/136648/rss" style=" color: #0099b4; font-weight: 600;   "&gt;want to submit your proposal quickly&lt;/a&gt; as it closes on Friday evening.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;The Go team meets regularly to discuss developments in the Go compiler, etc, and &lt;a href="https://golangweekly.com/link/136649/rss" style=" color: #0099b4; font-weight: 600;   "&gt;here are some of the recent meeting notes&lt;/a&gt; if you're interested in such details.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;After all the discussion about telemetry in Go, Russ Cox has now &lt;a href="https://golangweekly.com/link/136650/rss" style=" color: #0099b4; font-weight: 600;   "&gt;filed the &lt;em&gt;proposal&lt;/em&gt; for adding the opt-in telemetry feature.&lt;/a&gt;&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136651/rss" title="preslav.me" style=" color: #0099b4;    font-size: 1.05em;"&gt;Things to Consider When Going with &lt;code&gt;sqlc&lt;/code&gt;&lt;/a&gt;&lt;/span&gt; — &lt;a href="https://golangweekly.com/link/136652/rss" style=" color: #0099b4;   "&gt;sqlc&lt;/a&gt; (the SQL to type-safe code generator) is a handy time and boilerplate saver, but isn’t a one-size-fits-all solution.&lt;/p&gt;
  &lt;p&gt;Preslav Rachev &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136653/rss" title="event.on24.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Terraforming Kubernetes (Free Course)&lt;/a&gt;&lt;/span&gt; — Lead by Udemy instructor, Justin Mitchel, this course shows you how to spin up a K8s cluster using Terraform.&lt;/p&gt;
  &lt;p&gt;Akamai Connected Cloud &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136654/rss" title="wstrm.dev" style=" color: #0099b4;    font-size: 1.05em;"&gt;&lt;code&gt;errors.Join&lt;/code&gt; Loves &lt;code&gt;defer&lt;/code&gt;?&lt;/a&gt;&lt;/span&gt; — In Go 1.20, you can join errors so that original errors aren’t overridden. William looks at how nicely this meshes with the use case of handling close errors in a &lt;code&gt;defer&lt;/code&gt;.&lt;/p&gt;
  &lt;p&gt;William Wennerström &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136655/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;How &lt;em&gt;Not&lt;/em&gt; to Use &lt;code&gt;context.WithValue&lt;/code&gt;&lt;/a&gt;&lt;/span&gt;
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Vishnu Bharathi&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136656/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Higher-Order Functions in Go&lt;/a&gt;&lt;/span&gt;
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Eli Bendersky&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 0;  padding-left: 0;"&gt;&lt;p&gt;🛠 Code &amp;amp; Tools&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/136657/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ax8cm4aaw81km4ga5t73.jpg" width="640" style="        line-height: 100%;     "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136657/rss" title="blog.ngrok.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;ngrok-go: Ingress to Your Go Apps as a &lt;code&gt;net.Listener&lt;/code&gt;&lt;/a&gt;&lt;/span&gt; — &lt;a href="https://golangweekly.com/link/136658/rss" style=" color: #0099b4;   "&gt;ngrok&lt;/a&gt; is a long standing tool (and, increasingly, commercial service) for opening up a locally hosted service to the public Internet  and now there’s an idiomatically Go way to embed ngrok-style ingress into apps.&lt;/p&gt;
  &lt;p&gt;Alan Shreve (ngrok) &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136659/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Valgo: An Expressive Validator Library&lt;/a&gt;&lt;/span&gt; — Type-safe and extensible validator library that supports localization and is built upon generics. &lt;em&gt;“Valgo differs from other validation libraries in that the rules are written in functions and not in struct tags. This allows greater flexibility and freedom when it comes to where and how data is validated.”&lt;/em&gt;&lt;/p&gt;
  &lt;p&gt;Cohesive Stack &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136660/rss" title="shortcut.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Don’t Let Your Issue Tracker Be a Four-Letter Word. Use Shortcut&lt;/a&gt;&lt;/span&gt;&lt;/p&gt;
  &lt;p&gt;Shortcut (formerly Clubhouse.io) &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136661/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Gool: A Generic Goroutine Pool&lt;/a&gt;&lt;/span&gt; — If you’re familiar with Python’s ThreadPoolExecutor, you might like this as it provides sync and async versions of &lt;code&gt;Submit&lt;/code&gt; and &lt;code&gt;Map&lt;/code&gt; to run tasks.&lt;/p&gt;
  &lt;p&gt;Tommy Tian &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136662/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;go-ssaviz: SSA Visualization with Graphviz&lt;/a&gt;&lt;/span&gt; — You won’t need this unless you’re doing some &lt;em&gt;really&lt;/em&gt; low level work, but this tool lets you see Go’s internal use of &lt;a href="https://golangweekly.com/link/136663/rss" style=" color: #0099b4;   "&gt;static single-assignment&lt;/a&gt; in your Go package of choice.&lt;/p&gt;
  &lt;p&gt;Shengyu Zhang &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136664/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Convolutional Encoder and Viterbi Decoder for Go&lt;/a&gt;&lt;/span&gt; — If you need this, you’ll know. The author explains the error-correcting nature of such encoders and decoders &lt;a href="https://golangweekly.com/link/136665/rss" style=" color: #0099b4;   "&gt;in this comment.&lt;/a&gt;&lt;/p&gt;
  &lt;p&gt;8FF &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 0px;  padding-left: 0px;"&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;Jobs&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136666/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Site Reliability Engineer&lt;/a&gt;&lt;/span&gt; — Join our "kick ass" team. Our software team operates from 17 countries and we're always looking for more exceptional engineers.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Sticker Mule&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136667/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Find a Job Through Hired&lt;/a&gt;&lt;/span&gt; — Hired makes job hunting easy-instead of chasing recruiters, companies approach you with salary details up front. Create a free profile now.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Hired&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;🧑‍💻 Got a job listing to share? &lt;em&gt;&lt;a href="https://golangweekly.com/link/136668/rss" style=" color: #0099b4; font-weight: 600;"&gt;Here's how&lt;/a&gt;&lt;/em&gt;.&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136669/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Glog 1.1&lt;/a&gt;&lt;br&gt;
↳ Leveled execution log library.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136670/rss" style=" color: #0099b4; font-weight: 600;   "&gt;golang-set 2.2&lt;/a&gt;&lt;br&gt;
↳ Generic set type.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136671/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Wails 2.4&lt;/a&gt;&lt;br&gt;
↳ Go + web technologies == desktop apps.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136672/rss" style=" color: #0099b4; font-weight: 600;   "&gt;rqlite 7.14&lt;/a&gt;&lt;br&gt;
↳ Distributed relational database atop SQLite.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136673/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Goose 3.10&lt;/a&gt;&lt;br&gt;
↳ Database migration with SQL and/or Go functions.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136674/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Ginkgo 2.9&lt;/a&gt;&lt;br&gt;
↳ Modern testing framework.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136675/rss" style=" color: #0099b4; font-weight: 600;   "&gt;fq 0.4&lt;/a&gt;&lt;br&gt;
↳ Like &lt;code&gt;jq&lt;/code&gt; for binary formats.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136676/rss" style=" color: #0099b4; font-weight: 600;   "&gt;ClickHouse Go 2.7&lt;/a&gt;&lt;br&gt;
↳ Go driver for ClickHouse.&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/div&gt;
  &lt;/td&gt;&lt;/tr&gt;
&lt;/table&gt;




&lt;img src="https://golangweekly.com/open/451/rss" width="1" height="1" /&gt;</description>
      <pubDate>Fri, 10 Mar 2023 00:00:00 +0000</pubDate>
      <guid>https://golangweekly.com/issues/451</guid>
    </item>
  
  </channel>
</rss>
`
