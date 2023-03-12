package filesIO

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
    <item>
      <title>Google's new distributed Go app framework</title>
      <link>https://golangweekly.com/issues/450</link>
      <description>

  

    
    
    
    
    
  




&lt;table border=0 cellpadding=0 cellspacing=0 align="center" border="0"&gt;
  &lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;div&gt;    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;
&lt;td align="left" style="padding-left: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;#​450 — March 3, 2023&lt;/p&gt;&lt;/td&gt;
&lt;td align="right" style="padding-right: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;&lt;a href="https://golangweekly.com/link/136303/rss" style=" color: #0099b4;"&gt;Unsub&lt;/a&gt;  |  &lt;a href="https://golangweekly.com/link/136304/rss" style=" color: #0099b4;"&gt;Web Version&lt;/a&gt;&lt;/p&gt;&lt;/td&gt;
&lt;/tr&gt;&lt;/table&gt;
    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 12px;  padding-left: 12px;"&gt;&lt;p&gt;The &lt;em&gt;Go Weekly&lt;/em&gt; Newsletter&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/136306/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/f9magaazzokmum6ulrop.jpg" width="640" style="    line-height: 100%;       "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136306/rss" title="research.swtch.com" style=" color: #0099b4;    font-size: 1.1em; line-height: 1.4em;"&gt;&lt;em&gt;Opting In&lt;/em&gt; to Transparent Telemetry in Go&lt;/a&gt;&lt;/span&gt; — A debate around &lt;a href="https://golangweekly.com/link/136307/rss" style=" color: #0099b4;   "&gt;adding telemetry to the Go toolchain&lt;/a&gt; has been rolling for the past few weeks. Now there’s a new development: &lt;em&gt;“By far the most common suggestion was to make the system opt-in (default off) instead of opt-out (default on). I have revised the design to do that.”&lt;/em&gt; There are some downsides to that, of course.&lt;/p&gt;
  &lt;p&gt;Russ Cox &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136308/rss" title="opensource.googleblog.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Service Weaver: Google's Framework for Writing Distributed Go Apps&lt;/a&gt;&lt;/span&gt; — A new open-source framework from Google that lets you &lt;em&gt;“write your (Go) application as a modular monolith and deploy it as a set of microservices”&lt;/em&gt; to get the best of both worlds, namely: &lt;em&gt;“the development velocity of a monolith, with the scalability, security, and fault-tolerance of microservices.”&lt;/em&gt; If you fancy something more technical and less salesy, Robert Grandl has &lt;a href="https://golangweekly.com/link/136309/rss" style=" color: #0099b4;   "&gt;a quick introduction here&lt;/a&gt;.&lt;/p&gt;
  &lt;p&gt;Google Open Source &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;a href="https://golangweekly.com/link/136305/rss" style=" color: #0099b4;   "&gt;&lt;img src="https://copm.s3.amazonaws.com/e0920bf2.png" width="190" height="61" style="padding-top: 12px; padding-left: 12px;     line-height: 100%;    "&gt;&lt;/a&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136305/rss" title="tailscale.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Review Changes Made to Your Tailscale Network&lt;/a&gt;&lt;/span&gt; — Staying on top of what’s happening in your network is easier than ever with Tailscale’s configuration audit logs, which let admins review changes made to your Tailscale network, such as added devices, updated ACLs, or new DNS settings.&lt;/p&gt;
  &lt;p&gt;Tailscale &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136310/rss" title="entgo.io" style=" color: #0099b4;    font-size: 1.05em;"&gt;A Beginner's Guide to Creating a Webapp with Ent&lt;/a&gt;&lt;/span&gt; — A practical introduction to using &lt;a href="https://golangweekly.com/link/136311/rss" style=" color: #0099b4;   "&gt;Ent,&lt;/a&gt; an open-source entity framework for Go for modelling and querying data, by way of building a simple content management system.&lt;/p&gt;
  &lt;p&gt;Rotem Tamir &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136312/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Uber's Go Style Guide&lt;/a&gt;&lt;/span&gt; — Want to know how a large Go organization writes Go? Take a ride with this guide that covers guidelines, some performance issues, and stylistic concerns.&lt;/p&gt;
  &lt;p&gt;Uber &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;p&gt;&lt;strong&gt;IN BRIEF:&lt;/strong&gt;&lt;/p&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;If you want to use &lt;a href="https://golangweekly.com/link/136313/rss" style=" color: #0099b4; font-weight: 600;   "&gt;TinyGo&lt;/a&gt;, the small Go compiler for embedded systems and WebAssembly, you might find &lt;a href="https://golangweekly.com/link/136314/rss" style=" color: #0099b4; font-weight: 600;   "&gt;this list of supported packages&lt;/a&gt; useful.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;Efron Licht is back with &lt;a href="https://golangweekly.com/link/136315/rss" style=" color: #0099b4; font-weight: 600;   "&gt;more Go quirks and tricks&lt;/a&gt;.&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136316/rss" title="benhoyt.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;From Go on EC2 to Fly.io: More Fun, Less Cost?&lt;/a&gt;&lt;/span&gt; — Ben talks about switching a couple of side projects from an EC2 instance to up and coming no-ops platform &lt;a href="https://golangweekly.com/link/136317/rss" style=" color: #0099b4;   "&gt;Fly&lt;/a&gt;, and how he resolved some sticking points like running background jobs in Go without cron.&lt;/p&gt;
  &lt;p&gt;Ben Hoyt &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136318/rss" title="arriqaaq.substack.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Unlocking the Power of Zero Knowledge Proofs with Gnark&lt;/a&gt;&lt;/span&gt; — &lt;a href="https://golangweekly.com/link/136319/rss" style=" color: #0099b4;   "&gt;gnark&lt;/a&gt; is a library for creating zkSNARK (Zero-Knowledge Succinct Non-Interactive Argument of Knowledge) &lt;a href="https://golangweekly.com/link/136320/rss" style=" color: #0099b4;   "&gt;zero knowledge proof&lt;/a&gt; applications.&lt;/p&gt;
  &lt;p&gt;Farhan &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136321/rss" title="t.mp" style=" color: #0099b4;    font-size: 1.05em;"&gt;Try Temporal 101 in Go&lt;/a&gt;&lt;/span&gt; — In this beginner’s course, you’ll learn the basic building blocks of Temporal to develop an app that communicates with an external service.&lt;/p&gt;
  &lt;p&gt;Temporal Technologies &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136322/rss" title="build-your-own.org" style=" color: #0099b4;    font-size: 1.05em;"&gt;Build Your Own Database From Scratch&lt;/a&gt;&lt;/span&gt; — A book that’s still under development that uses Go but is language agnostic. A few chapters are ready to read, including on &lt;a href="https://golangweekly.com/link/136323/rss" style=" color: #0099b4;   "&gt;implementing B-trees in Go&lt;/a&gt;.&lt;/p&gt;
  &lt;p&gt;James Smith &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136324/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Why to Defer Your Mutex Unlocks&lt;/a&gt;&lt;/span&gt; — One for the best practice list?
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Emir Ribic&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136325/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Import Leads From Google Forms Into Your CRM with OpenFaaS&lt;/a&gt;&lt;/span&gt;
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Alex Ellis&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 0;  padding-left: 0;"&gt;&lt;p&gt;🛠 Code &amp;amp; Tools&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/136326/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ckbrv20xuzbchzexa4a6.jpg" width="640" style="    line-height: 100%;         "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136326/rss" title="proton.me" style=" color: #0099b4;    font-size: 1.05em;"&gt;Gluon: A High-Performance IMAP (Server-Side) Library&lt;/a&gt;&lt;/span&gt; — This is aimed at mail system implementers and server-side IMAP management, but this post digs into the details about why that’s a tricky thing to build and how Proton, the folks behind the privacy-first email platform Proton Mail, pulled it off. &lt;a href="https://golangweekly.com/link/136327/rss" style=" color: #0099b4;   "&gt;GitHub repo&lt;/a&gt;.&lt;/p&gt;
  &lt;p&gt;Proton &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;✉️ &lt;strong&gt;You've Got Mail:&lt;/strong&gt; If you'd prefer a complete 'out of the box' mail server experience with a Go-powered system, check out &lt;a href="https://golangweekly.com/link/136328/rss" style=" color: #0099b4; font-weight: 600;"&gt;Mox&lt;/a&gt; which implements SMTP, IMAP4, and numerous email specs in a single open source mail server.&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136329/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;algnhsa 1.0: AWS Lambda &lt;code&gt;net/http&lt;/code&gt; Server Adapter&lt;/a&gt;&lt;/span&gt; — Enables running Go webapps on AWS Lambda and API Gateway/ALB without changing existing HTTP handlers.&lt;/p&gt;
  &lt;p&gt;Artem Krylysov &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136330/rss" title="" style=" color: #0099b4;    font-size: 1.05em;"&gt;Tuple, a Lightning-Fast Pairing Tool Built for Remote Developers&lt;/a&gt;&lt;/span&gt; — High-resolution, crystal-clear screen sharing, low-latency remote control, and less CPU usage than you'd think possible.&lt;/p&gt;
  &lt;p&gt;Tuple &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136331/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Graph 0.16: Generic Library for Creating Graph Data Structures&lt;/a&gt;&lt;/span&gt; — Supports different kinds of graphs such as directed graphs, acyclic graphs, or trees. This week’s &lt;a href="https://golangweekly.com/link/136332/rss" style=" color: #0099b4;   "&gt;0.16.0&lt;/a&gt; release adds support for &lt;a href="https://golangweekly.com/link/136333/rss" style=" color: #0099b4;   "&gt;integrating storage backends of your choice&lt;/a&gt; by implementing a Store interface.&lt;/p&gt;
  &lt;p&gt;Dominik Braun &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136334/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;PNGR: Dockerized (Postgres + Nginx + Go + React)&lt;/a&gt;&lt;/span&gt; — Starter kit for a webapp that includes user and session management, JWT authentication, and a basic CRUD example.&lt;/p&gt;
  &lt;p&gt;Karl Keefer &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136335/rss" title="blog.alexellis.io" style=" color: #0099b4;    font-size: 1.05em;"&gt;Find Your Total Build Minutes with GitHub Actions and Go&lt;/a&gt;&lt;/span&gt;&lt;/p&gt;
  &lt;p&gt;Alex Ellis &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 0px;  padding-left: 0px;"&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;Jobs&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136345/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Site Reliability Engineer&lt;/a&gt;&lt;/span&gt; — Join our "kick ass" team. Our software team operates from 17 countries and we're always looking for more exceptional engineers.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Sticker Mule&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/136336/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Find a Job Through Hired&lt;/a&gt;&lt;/span&gt; — Hired makes job hunting easy-instead of chasing recruiters, companies approach you with salary details up front. Create a free profile now.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Hired&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;🧑‍💻 Got a job listing to share? &lt;em&gt;&lt;a href="https://golangweekly.com/link/136346/rss" style=" color: #0099b4; font-weight: 600;"&gt;Here's how&lt;/a&gt;&lt;/em&gt;.&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136337/rss" style=" color: #0099b4; font-weight: 600;   "&gt;pdfcpu 0.4&lt;/a&gt;&lt;br&gt;
↳ PDF processing library. &lt;em&gt;(Great logo!)&lt;/em&gt;&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136338/rss" style=" color: #0099b4; font-weight: 600;   "&gt;conc 0.3&lt;/a&gt;&lt;br&gt;
↳ Better structured concurrency for Go.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136339/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Render 1.6&lt;/a&gt;&lt;br&gt;
↳ Easily render JSON, XML, HTML &amp;amp; more.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136340/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Macaron 1.5&lt;/a&gt;&lt;br&gt;
↳ Modular Web framework.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136341/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Miller 6.7&lt;/a&gt;&lt;br&gt;
↳ Swiss army knife for name-indexed data.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136342/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Buf CLI 1.15&lt;/a&gt;&lt;br&gt;
↳ Tool for working with Protocol Buffers.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/136343/rss" style=" color: #0099b4; font-weight: 600;   "&gt;go-git 5.6&lt;/a&gt;&lt;br&gt;
↳ Extensible pure Go Git implementation.&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/div&gt;
  &lt;/td&gt;&lt;/tr&gt;
&lt;/table&gt;




&lt;img src="https://golangweekly.com/open/450/rss" width="1" height="1" /&gt;</description>
      <pubDate>Fri, 3 Mar 2023 00:00:00 +0000</pubDate>
      <guid>https://golangweekly.com/issues/450</guid>
    </item>
    <item>
      <title>A particularly Charming issue</title>
      <link>https://golangweekly.com/issues/449</link>
      <description>

  

    
    
    
    
    
  




&lt;table border=0 cellpadding=0 cellspacing=0 align="center" border="0"&gt;
  &lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;div&gt;    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;
&lt;td align="left" style="padding-left: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;#​449 — February 24, 2023&lt;/p&gt;&lt;/td&gt;
&lt;td align="right" style="padding-right: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;&lt;a href="https://golangweekly.com/link/135931/rss" style=" color: #0099b4;"&gt;Unsub&lt;/a&gt;  |  &lt;a href="https://golangweekly.com/link/135932/rss" style=" color: #0099b4;"&gt;Web Version&lt;/a&gt;&lt;/p&gt;&lt;/td&gt;
&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;I'm not sure what's up, but &lt;a href="https://golangweekly.com/link/135933/rss" style=" color: #0099b4; font-weight: 600;"&gt;Charm's&lt;/a&gt; projects have popped up all over the place this week, so get ready for the most &lt;em&gt;Charm-ing&lt;/em&gt; issue we've ever sent.. 🤭&lt;br&gt;__&lt;br&gt;&lt;em&gt;Peter Cooper, your editor&lt;/em&gt;&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 12px;  padding-left: 12px;"&gt;&lt;p&gt;The &lt;em&gt;Go Weekly&lt;/em&gt; Newsletter&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/135935/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/zamx5n18olmkduhgmyxz.jpg" width="640" style="    line-height: 100%;       "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135935/rss" title="github.com" style=" color: #0099b4;    font-size: 1.1em; line-height: 1.4em;"&gt;Log: A Minimal, Colorful Go Logging Library&lt;/a&gt;&lt;/span&gt; — A library from the same folks who brought us &lt;a href="https://golangweekly.com/link/135936/rss" style=" color: #0099b4;   "&gt;Bubble Tea&lt;/a&gt; and &lt;a href="https://golangweekly.com/link/135937/rss" style=" color: #0099b4;   "&gt;Gum&lt;/a&gt; so you know it’s from a good place. It provides &lt;em&gt;“customizable colorful human readable logging with batteries included.”&lt;/em&gt;&lt;/p&gt;
  &lt;p&gt;Charm &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135938/rss" title="go.dev" style=" color: #0099b4;    font-size: 1.05em;"&gt;All Your Comparable Types&lt;/a&gt;&lt;/span&gt; — The introduction of generics was bound to create edge cases. One is described here around interface implementation vs. constraint satisfaction. &lt;em&gt;“As we’ll see in a bit, in Go 1.20 constraint satisfaction is not quite constraint implementation anymore.”&lt;/em&gt;&lt;/p&gt;
  &lt;p&gt;Robert Griesemer &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;a href="https://golangweekly.com/link/135934/rss" style=" color: #0099b4;   "&gt;&lt;img src="https://copm.s3.amazonaws.com/bcc68a3c.png" width="110" height="110" style="padding-top: 12px; padding-left: 12px;     line-height: 100%;    "&gt;&lt;/a&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135934/rss" title="www.ardanlabs.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Go! Experts at Your Service&lt;/a&gt;&lt;/span&gt; — Do you need help filling skill gaps, speeding up development &amp;amp; creating high performing software with Go, Docker, K8s, Terraform and Rust? We’ll help you maximize your architecture, structure, tech-debt and human capital.&lt;/p&gt;
  &lt;p&gt;Ardan Labs Consulting &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135939/rss" title="www.dolthub.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Type Embedding: Go's Fake 'Inheritance'&lt;/a&gt;&lt;/span&gt; — Inspired by a &lt;em&gt;“Keep Your Java Out of My Go”&lt;/em&gt; Reddit post, Zach shows how leaning too hard on object oriented ideas can cause hard-to-track-down bugs. But if you’re probably going to do it anyway.. read this to help debug later.&lt;/p&gt;
  &lt;p&gt;Zach Musgrave (DoltHub) &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;If you're curious, the not-particularly-edifying Reddit post that inspired the above was &lt;a href="https://golangweekly.com/link/135940/rss" style=" color: #0099b4; font-weight: 600;"&gt;'How to deal with Java developers polluting the Go code?'&lt;/a&gt;&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;p&gt;&lt;strong&gt;IN BRIEF:&lt;/strong&gt;&lt;/p&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135941/rss" style=" color: #0099b4; font-weight: 600;   "&gt;▶️ A developer explains&lt;/a&gt; how he picked up a $3,133.70 bounty for finding a XSS vulnerability in Go's &lt;code&gt;x/net/html&lt;/code&gt; package. There's &lt;a href="https://golangweekly.com/link/135942/rss" style=" color: #0099b4; font-weight: 600;   "&gt;a code example&lt;/a&gt; of the problem and the &lt;a href="https://golangweekly.com/link/135943/rss" style=" color: #0099b4; font-weight: 600;   "&gt;resulting patch.&lt;/a&gt;&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;Generics have been in Go for roughly a year now, so &lt;a href="https://golangweekly.com/link/135944/rss" style=" color: #0099b4; font-weight: 600;   "&gt;🐦 what have you been using them for&lt;/a&gt;, asks Kelsey Hightower?&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;Brandur Leach has some findings from &lt;a href="https://golangweekly.com/link/135945/rss" style=" color: #0099b4; font-weight: 600;   "&gt;six months of running &lt;code&gt;govulncheck&lt;/code&gt; in CI.&lt;/a&gt;&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135946/rss" title="bryce.is" style=" color: #0099b4;    font-size: 1.05em;"&gt;&lt;code&gt;go test&lt;/code&gt; and Parallelism&lt;/a&gt;&lt;/span&gt; — &lt;em&gt;“Because I feel the concurrency behavior of &lt;code&gt;go test&lt;/code&gt; is non-obvious .. I wanted to write something up here.”&lt;/em&gt;&lt;/p&gt;
  &lt;p&gt;Bryce Neal &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135947/rss" title="www.kosli.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;How to Publish Go Binaries with Goreleaser&lt;/a&gt;&lt;/span&gt; — &lt;a href="https://golangweekly.com/link/135948/rss" style=" color: #0099b4;   "&gt;Goreleaser&lt;/a&gt; is a helpful tool if you need to cross-compile and publish binaries for multiple architectures, different operating systems, package managers, etc.&lt;/p&gt;
  &lt;p&gt;Rexford A Nyarko &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135949/rss" title="preslav.me" style=" color: #0099b4;    font-size: 1.05em;"&gt;Partially-Implemented Interfaces&lt;/a&gt;&lt;/span&gt; — When you just want to implement one or two methods … but be careful.&lt;/p&gt;
  &lt;p&gt;Preslav Rachev &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135950/rss" title="teleport.registration.goldcast.io" style=" color: #0099b4;    font-size: 1.05em;"&gt;Securely Deploy Kubernetes Clusters with GitHub Actions&lt;/a&gt;&lt;/span&gt; — Managing identity for Kubernetes &amp;amp; CI/CD workflows relies on dated security mechanisms, learn more in our new episode.&lt;/p&gt;
  &lt;p&gt;TELEPORT | goteleport.com &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;▶  &lt;a href="https://golangweekly.com/link/135951/rss" title="changelog.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;What's New in Go 1.20 with Carl Johnson&lt;/a&gt;&lt;/span&gt; — An hour of exactly what it says on the can..&lt;/p&gt;
  &lt;p&gt;The Go Time Podcast &lt;span style="text-transform: uppercase; margin-left: 4px; padding-top: 1px; padding-right: 4px;  padding-left: 4px; font-size: 0.9em;             "&gt;podcast&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135952/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;The Complete Guide to OpenTelemetry in Go&lt;/a&gt;&lt;/span&gt;
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Komu Wairagu&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 0;  padding-left: 0;"&gt;&lt;p&gt;🛠 Code &amp;amp; Tools&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/135953/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/lyqvupyaacmtjr9896zc.jpg" width="640" style="        line-height: 100%;     "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135953/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Opossum: A Basic Web Browser Written in Go&lt;/a&gt;&lt;/span&gt; — A fun project, though you might need to have (or set up) a Plan 9-derived OS to get it running (though someone &lt;a href="https://golangweekly.com/link/135954/rss" style=" color: #0099b4;   "&gt;claims success&lt;/a&gt; on macOS). As user &lt;em&gt;liotier&lt;/em&gt; &lt;a href="https://golangweekly.com/link/135955/rss" style=" color: #0099b4;   "&gt;said&lt;/a&gt; on Hacker News, though: &lt;em&gt;“any sufficiently brave or delusional soul that ventures into even the most rudimentary web browser development is a hero to me”&lt;/em&gt;.&lt;/p&gt;
  &lt;p&gt;Philip Silva &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135956/rss" title="gin-gonic.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Gin 1.9: A Fast HTTP Web Framework&lt;/a&gt;&lt;/span&gt; — We don’t often link to Gin because a) it doesn’t often get big updates, and b) it’s hugely popular and you probably use it already ;-) If you do, &lt;a href="https://golangweekly.com/link/135957/rss" style=" color: #0099b4;   "&gt;v1.9 is out&lt;/a&gt;, and if you don’t, it’s worth being on your radar.&lt;/p&gt;
  &lt;p&gt;Gin Web Framework &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135958/rss" title="shortcut.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Don’t Let Your Issue Tracker Be a Four-Letter Word. Use Shortcut&lt;/a&gt;&lt;/span&gt;&lt;/p&gt;
  &lt;p&gt;Shortcut (formerly Clubhouse.io) &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135959/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Soft Serve: A Self-Hostable Git Server&lt;/a&gt;&lt;/span&gt; — Soft is configurable via &lt;code&gt;git&lt;/code&gt; itself and comes with a nice terminal UI (TUI). Another interesting release from the Charm project.&lt;/p&gt;
  &lt;p&gt;Charm &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135960/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;koanf: Configuration Management Library&lt;/a&gt;&lt;/span&gt; — Support for JSON, TOML, YAML, env, command line, file, S3 etc. Alternative to Viper.&lt;/p&gt;
  &lt;p&gt;Kailash Nadh &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135961/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;VHS 0.3: The 'Home Video Recorder' for Your CLI&lt;/a&gt;&lt;/span&gt; — A tool plus scripting language for performing actions on the terminal which are then recorded into an animated GIF (or an MP4, webm file, or series of PNG frames). &lt;em&gt;“Write terminal GIFs as code for integration testing and demoing your CLI tools.”&lt;/em&gt; &lt;a href="https://golangweekly.com/link/135962/rss" style=" color: #0099b4;   "&gt;v0.3&lt;/a&gt; adds support for hosting said GIFs on Charm’s own &lt;code&gt;vhs.charm․sh&lt;/code&gt; service.&lt;/p&gt;
  &lt;p&gt;Charm &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;&lt;em&gt;.. Is this the last mention of a Charm project this week..?&lt;/em&gt;&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135963/rss" style=" color: #0099b4; font-weight: 600;   "&gt;HoverFly 2.0&lt;/a&gt;&lt;br&gt;
↳ Lightweight API simulation tool.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135964/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Multi Progress Bar / mpb 8.2&lt;/a&gt;&lt;br&gt;
↳ Multi progress bar for CLI apps.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135965/rss" style=" color: #0099b4; font-weight: 600;   "&gt;imagor 1.4.1&lt;/a&gt;&lt;br&gt;
↳ libvips-based image processing.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135966/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Pie 2.4&lt;/a&gt;&lt;br&gt;
↳ Common operations for slices and maps.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135967/rss" style=" color: #0099b4; font-weight: 600;   "&gt;pgweb 0.14&lt;/a&gt;&lt;br&gt;
↳ Web-based explorer for Postgres.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135968/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Define 0.3&lt;/a&gt;&lt;br&gt;
↳ Command line dictionary and thesaurus.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135969/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Lefthook 1.3&lt;/a&gt;&lt;br&gt;
↳ Polyglot Git hooks manager.&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 0px;  padding-left: 0px;"&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;Jobs&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135972/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Site Reliability Engineer&lt;/a&gt;&lt;/span&gt; — Join our "kick ass" team. Our software team operates from 17 countries and we're always looking for more exceptional engineers.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Sticker Mule&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135970/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Find a Job Through Hired&lt;/a&gt;&lt;/span&gt; — Hired makes job hunting easy-instead of chasing recruiters, companies approach you with salary details up front. Create a free profile now.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Hired&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 0;  padding-left: 0;"&gt;&lt;p&gt;♣️ &lt;em&gt;Solitaire's the only game in town..&lt;/em&gt;&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/135971/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/wxsk6atb4jd25ejezlut.jpg" width="640" style="        line-height: 100%;     "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135971/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Solitaire TUI: Klondike Solitaire on the Terminal&lt;/a&gt;&lt;/span&gt; — If you, like me, &lt;strike&gt;wasted&lt;/strike&gt;spent much time in Windows 3.x playing the solitaire game, this &lt;a href="https://golangweekly.com/link/135936/rss" style=" color: #0099b4;   "&gt;Bubble Tea&lt;/a&gt;-powered creation (&lt;em&gt;THAT'S THE FINAL CHARM PROJECT THIS WEEK!&lt;/em&gt;) might appeal to you now in 2023 too. Bonus points to anyone who can file a pull request to add the Windows-style &lt;a href="https://golangweekly.com/link/135974/rss" style=" color: #0099b4;   "&gt;ending animation&lt;/a&gt; ;-)&lt;/p&gt;
  &lt;p&gt;Brian Strauch &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/div&gt;
  &lt;/td&gt;&lt;/tr&gt;
&lt;/table&gt;




&lt;img src="https://golangweekly.com/open/449/rss" width="1" height="1" /&gt;</description>
      <pubDate>Fri, 24 Feb 2023 00:00:00 +0000</pubDate>
      <guid>https://golangweekly.com/issues/449</guid>
    </item>
    <item>
      <title>Go quirks and intermediate tricks</title>
      <link>https://golangweekly.com/issues/448</link>
      <description>

  

    
    
    
    
    
  




&lt;table border=0 cellpadding=0 cellspacing=0 align="center" border="0"&gt;
  &lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;div&gt;    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;
&lt;td align="left" style="padding-left: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;#​448 — February 17, 2023&lt;/p&gt;&lt;/td&gt;
&lt;td align="right" style="padding-right: 4px; font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;p&gt;&lt;a href="https://golangweekly.com/link/135578/rss" style=" color: #0099b4;"&gt;Unsub&lt;/a&gt;  |  &lt;a href="https://golangweekly.com/link/135579/rss" style=" color: #0099b4;"&gt;Web Version&lt;/a&gt;&lt;/p&gt;&lt;/td&gt;
&lt;/tr&gt;&lt;/table&gt;
    
    &lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 12px;  padding-left: 12px;"&gt;&lt;p&gt;The &lt;em&gt;Go Weekly&lt;/em&gt; Newsletter&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/135581/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/tgzfvpr5to6yxgft4hgg.jpg" width="640" style="    line-height: 100%;       "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135581/rss" title="github.com" style=" color: #0099b4;    font-size: 1.1em; line-height: 1.4em;"&gt;Purego: A Dynamic Way to Call C Functions from Go &lt;em&gt;Without&lt;/em&gt; Cgo&lt;/a&gt;&lt;/span&gt; — No C means you can build for other platforms easily without a target C compiler/toolchain. No wrapper functions either. One of the contributors &lt;a href="https://golangweekly.com/link/135631/rss" style=" color: #0099b4;   "&gt;noted on HN&lt;/a&gt;: &lt;em&gt;"It uses the same mechanisms that Cgo does to switch to the system stack and then call the C code. Purego just avoids having to need a C toolchain to cross compile code that calls into C from Go."&lt;/em&gt;&lt;/p&gt;
  &lt;p&gt;Ebitengine &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135583/rss" title="eblog.fly.dev" style=" color: #0099b4;    font-size: 1.05em;"&gt;Go Quirks and Intermediate Tricks&lt;/a&gt;&lt;/span&gt; — This isn’t the best formatted post, but you might pick up a few things from this list (which, handily, has code examples for each item).&lt;/p&gt;
  &lt;p&gt;Efron Licht &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;a href="https://golangweekly.com/link/135580/rss" style=" color: #0099b4;   "&gt;&lt;img src="https://copm.s3.amazonaws.com/618d4972.jpg" width="146" height="110" style="padding-top: 12px; padding-left: 12px;     line-height: 100%;    "&gt;&lt;/a&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135580/rss" title="www.cockroachlabs.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Free from O’Reilly: Build Resilient Apps in Go&lt;/a&gt;&lt;/span&gt; — Learn to build cloud-native, cost-effective, and fault-tolerant applications in this 3-chapter excerpt from O’Reilly’s Cloud Native Go.&lt;/p&gt;
  &lt;p&gt;CockroachDB &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135582/rss" title="bitfieldconsulting.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Rust vs Go in 2023&lt;/a&gt;&lt;/span&gt; — A 2023-flavored update to a popular article first released in 2020. John is a fan of &lt;em&gt;both&lt;/em&gt; Rust and Go and takes a careful look at where each independently makes the most sense.&lt;/p&gt;
  &lt;p&gt;John Arundel &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;span&gt;🗣&lt;/span&gt; &lt;a href="https://golangweekly.com/link/135584/rss" title="www.theregister.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Go May Add Telemetry Reporting That's On by Default&lt;/a&gt;&lt;/span&gt; — We featured the &lt;a href="https://golangweekly.com/link/135585/rss" style=" color: #0099b4;   "&gt;initial discussion&lt;/a&gt; behind this last week, but &lt;em&gt;The Register&lt;/em&gt; has decided to run with the story and digested the public opinion a little more. Issues around ‘phoning home’ and privacy are always going to be fierce, but the &lt;em&gt;intent&lt;/em&gt;, at least, appears to be good and &lt;a href="https://golangweekly.com/link/135586/rss" style=" color: #0099b4;   "&gt;as Nate Finch says&lt;/a&gt; &lt;em&gt;“It’s just Go developers wanting to know how people use their software, so they can make it better. For you. For me. For all of us.”&lt;/em&gt;&lt;/p&gt;
  &lt;p&gt;The Register &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;p&gt;&lt;strong&gt;IN BRIEF:&lt;/strong&gt;&lt;/p&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135587/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Go 1.20.1 and 1.19.6&lt;/a&gt; have been released – they include four security fixes around &lt;code&gt;path/filepath&lt;/code&gt;, &lt;code&gt;net/http&lt;/code&gt;, and &lt;code&gt;crypto/tls&lt;/code&gt;&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;One developer has seen a huge improvement in compile times &lt;a href="https://golangweekly.com/link/135588/rss" style=" color: #0099b4; font-weight: 600;   "&gt;just from getting a better SSD?&lt;/a&gt;&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;Johnny Boursiquot reminds us it's possible to &lt;a href="https://golangweekly.com/link/135589/rss" style=" color: #0099b4; font-weight: 600;   "&gt;request an episode of the &lt;em&gt;Go Time&lt;/em&gt; podcast&lt;/a&gt;, if you've got any guest or topic ideas.&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135590/rss" title="mrkaran.dev" style=" color: #0099b4;    font-size: 1.05em;"&gt;Structured Logging with slog&lt;/a&gt;&lt;/span&gt; — There is a language proposal to add &lt;a href="https://golangweekly.com/link/135591/rss" style=" color: #0099b4;   "&gt;slog&lt;/a&gt; to the standard library (&lt;a href="https://golangweekly.com/link/135592/rss" style=" color: #0099b4;   "&gt;GitHub discussion 
 here.&lt;/a&gt;) slog has the basic features you’d expect. Still, not everyone is wild about how it handles attributes and custom loggers.&lt;/p&gt;
  &lt;p&gt;Karan Sharma &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;▶  &lt;span style="font-size: 0.75em;  font-weight: 300; margin-left: -4px;"&gt;⬆️&lt;/span&gt; &lt;a href="https://golangweekly.com/link/135593/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Really Simple Structured Logging with Tracing&lt;/a&gt;&lt;/span&gt;
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Kai Hendry&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135594/rss" title="blog.pratimbhosale.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Building a URL Shortener using Go, SQLite and GORM&lt;/a&gt;&lt;/span&gt; — &lt;a href="https://golangweekly.com/link/135595/rss" style=" color: #0099b4;   "&gt;GORM&lt;/a&gt; is an ORM library for Go.&lt;/p&gt;
  &lt;p&gt;Pratim Bhosale &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;▶  &lt;a href="https://golangweekly.com/link/135596/rss" title="www.youtube.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Debugging Concurrent Programs in Go&lt;/a&gt;&lt;/span&gt; — The audio isn’t &lt;em&gt;great&lt;/em&gt; but this is too useful not to link.&lt;/p&gt;
  &lt;p&gt;Andrii Soldatenko &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;▶  &lt;a href="https://golangweekly.com/link/135613/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;A Wacky Journey of Building a Vector Database in Go&lt;/a&gt;&lt;/span&gt;
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Etienne Dilocker&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 0;  padding-left: 0;"&gt;&lt;p&gt;🛠 Code &amp;amp; Tools&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/135597/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/s67qnqeg2bzdlub2uqbu.jpg" width="640" style="    line-height: 100%;         "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135597/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;go-pretty: Pretty Print Tables, Lists and Text on the Terminal&lt;/a&gt;&lt;/span&gt; — Utilities to prettify console output of tables, lists, progress-bars, text, etc. with a heavy emphasis on customization.&lt;/p&gt;
  &lt;p&gt;Naveen Mahalingam &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135598/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;gofumpt: A Stricter &lt;code&gt;gofmt&lt;/code&gt;&lt;/a&gt;&lt;/span&gt; — You like rules? &lt;code&gt;gofmt&lt;/code&gt; not strict enough? &lt;code&gt;gofumpt&lt;/code&gt; has even &lt;a href="https://golangweekly.com/link/135599/rss" style=" color: #0099b4;   "&gt;stricter rules&lt;/a&gt; to keep your codebase clean and behaving.&lt;/p&gt;
  &lt;p&gt;Daniel Martí &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135600/rss" title="tailscale.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Connect to Your Cloud Resources with Tailscale&lt;/a&gt;&lt;/span&gt; — Spend more time coding and less time troubleshooting with Tailscale. Now you can connect directly to cloud resources, containers, or VMs like they’re on your local network.&lt;/p&gt;
  &lt;p&gt;Tailscale &lt;span style="text-transform: uppercase; margin-left: 4px; font-size: 0.9em;   color: #885 !important; padding-top: 1px; padding-right: 4px;  padding-left: 4px;            "&gt;sponsor&lt;/span&gt;&lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135601/rss" title="redis.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;Go-Redis Now an &lt;em&gt;Official&lt;/em&gt; Redis Client&lt;/a&gt;&lt;/span&gt; — A couple of weeks ago we mentioned &lt;a href="https://golangweekly.com/link/135602/rss" style=" color: #0099b4;   "&gt;Go-Redis v9&lt;/a&gt; and how it had moved under the official Redis organization – now we get the full story.&lt;/p&gt;
  &lt;p&gt;Igor Malinovskyi &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
&lt;ul&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135603/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Noti 3.7&lt;/a&gt;&lt;br&gt;
↳ Monitor a process and trigger a notification.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135604/rss" style=" color: #0099b4; font-weight: 600;   "&gt;ZincSearch 0.4&lt;/a&gt;&lt;br&gt;
↳ Go-powered Elasticsearch alternative.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135605/rss" style=" color: #0099b4; font-weight: 600;   "&gt;Imagor 1.4&lt;/a&gt;&lt;br&gt;
↳ libvips-powered image processing server and Go library.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135606/rss" style=" color: #0099b4; font-weight: 600;   "&gt;q 0.9&lt;/a&gt;&lt;br&gt;
↳ Command line DNS client with support for UDP, TCP, DoT, DoH, DoQ and ODoH.&lt;/p&gt;
&lt;/li&gt;
&lt;li&gt;
&lt;p&gt;&lt;a href="https://golangweekly.com/link/135607/rss" style=" color: #0099b4; font-weight: 600;   "&gt;TinyGo 0.27&lt;/a&gt;&lt;br&gt;
↳ The Go compiler for small places.&lt;/p&gt;
&lt;/li&gt;
&lt;/ul&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 0px;  padding-left: 0px;"&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;&lt;p&gt;Jobs&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135608/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Site Reliability Engineer&lt;/a&gt;&lt;/span&gt; — Join our "kick ass" team. Our software team operates from 17 countries and we're always looking for more exceptional engineers.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Sticker Mule&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.0em; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135609/rss" style=" color: #0099b4; font-size: 1.0em !important; font-weight: 500;   "&gt;Find a Job Through Hired&lt;/a&gt;&lt;/span&gt; — Hired makes job hunting easy-instead of chasing recruiters, companies approach you with salary details up front. Create a free profile now.
  &lt;br&gt;&lt;span style="color: #5a5a5a; margin-top: 4px; text-transform: uppercase; font-size: 12px; line-height: 1.3em;"&gt;Hired&lt;/span&gt; 
  &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0; padding-right: 0;  padding-left: 0;"&gt;&lt;p&gt;🎈 Up, up, and away.&lt;/p&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;
  &lt;a href="https://golangweekly.com/link/135610/rss" style=" color: #0099b4;"&gt;&lt;img src="https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/mvsz6mpovcgw5zvqvtb2.jpg" width="640" style="        line-height: 100%;     "&gt;&lt;/a&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;

&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style="font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em;  padding-top: 0px; padding-right: 15px;  padding-left: 15px;"&gt;
  
  &lt;p&gt;&lt;span style="font-weight: 600; font-size: 1.2em !important; color: #000;"&gt;&lt;a href="https://golangweekly.com/link/135610/rss" title="github.com" style=" color: #0099b4;    font-size: 1.05em;"&gt;TinyGlobo: The TinyGo Powered Long Distance Balloon&lt;/a&gt;&lt;/span&gt; — It seems you don’t need to be a nation state agency to get in on the UFO action nowadays – you could just be a Go developer! Featuring a RP2040 programmed with &lt;a href="https://golangweekly.com/link/135611/rss" style=" color: #0099b4;   "&gt;TinyGo&lt;/a&gt;, this balloon sent back data using LoRaWAN long-range radio. &lt;a href="https://golangweekly.com/link/135612/rss" style=" color: #0099b4;   "&gt;▶️ Here’s a video of the fun.&lt;/a&gt;&lt;/p&gt;
  &lt;p&gt;Ron Evans &lt;/p&gt;
&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;table border=0 cellpadding=0 cellspacing=0 border=0 cellpadding=0 cellspacing=0&gt;&lt;tr&gt;&lt;td style=" font-family: -apple-system,BlinkMacSystemFont,Helvetica,sans-serif; font-size: 15px; line-height: 1.48em; "&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt;
&lt;/div&gt;
  &lt;/td&gt;&lt;/tr&gt;
&lt;/table&gt;




&lt;img src="https://golangweekly.com/open/448/rss" width="1" height="1" /&gt;</description>
      <pubDate>Fri, 17 Feb 2023 00:00:00 +0000</pubDate>
      <guid>https://golangweekly.com/issues/448</guid>
    </item>
  </channel>
</rss>
`

//////////////////////////////////////////////////////
///////////////////////////////////////////////////////

var Data2 string = `<rss xmlns:dc="http://purl.org/dc/elements/1.1/" version="2.0">
<channel>
<title>Go – Компилируемый, многопоточный язык программирования</title>
<link>https://habr.com/ru/hub/go/all/</link>
<description>
<![CDATA[ Go – Компилируемый, многопоточный язык программирования ]]>
</description>
<language>ru</language>
<managingEditor>editor@habr.com</managingEditor>
<generator>habr.com</generator>
<pubDate>Fri, 10 Mar 2023 00:43:21 GMT</pubDate>
<image>
<link>https://habr.com/ru/</link>
<url>https://habrastorage.org/webt/ym/el/wk/ymelwk3zy1gawz4nkejl_-ammtc.png</url>
<title>Хабр</title>
</image>
<item>
<title>
<![CDATA[ Архитектура Hashicorp Vault ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721416/</guid>
<link>https://habr.com/ru/post/721416/?utm_campaign=721416&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>В данной статье описал схему маршрутизации и получения данных в hashicorp vault(это зашифированное хранилище секретов с доступом по политикам). Возможно будет полезно тем, кто думает над архитектурой сервера или слоем(-ми) доступа к данным.</p><p></p> <a href="https://habr.com/ru/post/721416/?utm_campaign=721416&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 19:37:43 GMT</pubDate>
<dc:creator>valli0x</dc:creator>
<category>
<![CDATA[ Криптография ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Облачные сервисы ]]>
</category>
<category>
<![CDATA[ hashicorp vault ]]>
</category>
<category>
<![CDATA[ шифрование данных ]]>
</category>
<category>
<![CDATA[ storage ]]>
</category>
<category>
<![CDATA[ архитектура по ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Эволюция алгоритма фильтрации модификаций товаров в Авито ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720880/</guid>
<link>https://habr.com/ru/post/720880/?utm_campaign=720880&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Всем привет! Меня зовут Денис Колпаков, я бэкенд-инженер в юните Core Services Авито. Долгое время я был овнером критически значимого для бизнеса сервиса форм, а последний год занимаюсь каталогами и каталогизацией.&nbsp;</p><p>Я расскажу, как мы решали продуктовую задачу — искали способ отфильтровать модификации товаров из базы данных.&nbsp;</p><p></p> <a href="https://habr.com/ru/post/720880/?utm_campaign=720880&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 11:32:18 GMT</pubDate>
<dc:creator>Tifongod</dc:creator>
<category>
<![CDATA[ Блог компании AvitoTech ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ bitmap ]]>
</category>
<category>
<![CDATA[ фильтрация ]]>
</category>
<category>
<![CDATA[ алгоритмы ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Bazel, stamping, remote cache ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720792/</guid>
<link>https://habr.com/ru/post/720792/?utm_campaign=720792&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>В Bazel есть любопытная фича, позволяющая добавить данные, которые не инвалидируют кэш сборки.</p><p>Например, это бывает полезно, чтобы добавить в исполняемый файл информацию о том, когда он был собран и из какой ревизии. Если для времени и номера ревизии использовать stamping, то, когда собранный файл уже есть в кэше, он пересобираться не будет.</p><p>Разберемся, как stamping использовать...</p><p></p> <a href="https://habr.com/ru/post/720792/?utm_campaign=720792&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 09:17:20 GMT</pubDate>
<dc:creator>Bozaro</dc:creator>
<category>
<![CDATA[ Блог компании Joom ]]>
</category>
<category>
<![CDATA[ Тестирование IT-систем ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Системы сборки ]]>
</category>
<category>
<![CDATA[ joom ]]>
</category>
<category>
<![CDATA[ bazel ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ stamping ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Сказ о том как pet-project превратился в небольшой пассивный доход (часть 2) ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720596/</guid>
<link>https://habr.com/ru/post/720596/?utm_campaign=720596&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p><a href="https://habr.com/ru/post/718898/" rel="noopener noreferrer nofollow">Первая Часть</a></p><p>Предыдущая часть закончилась неудачной балансировкой, которая не решает практически никаких проблем. В комментариях кто-то спросил, почему я не использовал балансировку на уровне DNS. Так вот, я ее использовал. Оказалось, что c помощью DNS записей можно организовать балансировку Round Robin. Для этого в конфигурации Wireguard всего лишь нужно использовать доменное имя вместо IP адреса.</p><p></p> <a href="https://habr.com/ru/post/720596/?utm_campaign=720596&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Tue, 07 Mar 2023 14:15:46 GMT</pubDate>
<dc:creator>tarmalonchik</dc:creator>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Сетевые технологии ]]>
</category>
<category>
<![CDATA[ DNS ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ wireguard ]]>
</category>
<category>
<![CDATA[ vpn ]]>
</category>
<category>
<![CDATA[ балансировка нагрузки ]]>
</category>
<category>
<![CDATA[ backend ]]>
</category>
<category>
<![CDATA[ pet-project ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Мой путь в профессию: из аналитиков в Go-разработчики ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720678/</guid>
<link>https://habr.com/ru/post/720678/?utm_campaign=720678&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <img src="https://habrastorage.org/webt/aj/gi/zu/ajgizuladpsz4b-cuuwxfflbqnw.jpeg" alt="image"><br> <br> Привет! Меня зовут Герман, я backend-разработчик в команде <a href="https://cloud.mts.ru/services/dbaas-for-redis/?utm_source=habr.com&amp;utm_medium=owned_media_redisgo&amp;utm_content=article&amp;utm_term=redisgo" rel="nofollow noopener noreferrer">Managed Service for Redis</a> в компании #CloudMTS. В этой статье расскажу про свой приход в разработку на Go и поделюсь полезными ресурсами, которые мне помогли на этом пути. <br> <br> <a href="https://habr.com/ru/post/720678/?utm_campaign=720678&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Tue, 07 Mar 2023 08:08:56 GMT</pubDate>
<dc:creator>german_lepin</dc:creator>
<category>
<![CDATA[ Блог компании CloudMTS ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Учебный процесс в IT ]]>
</category>
<category>
<![CDATA[ Карьера в IT-индустрии ]]>
</category>
<category>
<![CDATA[ карьера ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ backend ]]>
</category>
<category>
<![CDATA[ образование ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Golang-дайджест № 26 (1 – 28 февраля  2023) ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720756/</guid>
<link>https://habr.com/ru/post/720756/?utm_campaign=720756&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Свежая подборка новостей и материалов.</p><p><strong>Интересное в этом выпуске</strong></p><p>Выпущены Go 1.20.1 и 1.19.6, воздушный шар дальнего радиуса действия, полное руководство по OpenTelemetry, пасьянс в терминале</p><p></p> <a href="https://habr.com/ru/post/720756/?utm_campaign=720756&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Mon, 06 Mar 2023 12:36:48 GMT</pubDate>
<dc:creator>tioffs</dc:creator>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ программирование ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Выбираем IAM в 2023 или, что есть кроме Keycloak ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720490/</guid>
<link>https://habr.com/ru/post/720490/?utm_campaign=720490&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Гипотетическая ситуация&nbsp;— ваш работодатель поручил вам выбрать Identity and Access Management platform.</p><p><strong>Обязательно:</strong> open‑source (Apache 2.0), self‑hosted, OAuth 2.0, OIDC, SAML, LDAP.</p><p>Для&nbsp;тех кому интересно узнать, что&nbsp;есть еще кроме Keycloak. </p><p></p> <a href="https://habr.com/ru/post/720490/?utm_campaign=720490&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Узнать</a> ]]>
</description>
<pubDate>Mon, 06 Mar 2023 11:28:56 GMT</pubDate>
<dc:creator>MrAwesome</dc:creator>
<category>
<![CDATA[ Информационная безопасность ]]>
</category>
<category>
<![CDATA[ Open source ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ opensource ]]>
</category>
<category>
<![CDATA[ keycloak ]]>
</category>
<category>
<![CDATA[ casdoor ]]>
</category>
<category>
<![CDATA[ zitadel ]]>
</category>
<category>
<![CDATA[ IAM ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Assembler в Go: техники ускорения и оптимизации ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720582/</guid>
<link>https://habr.com/ru/post/720582/?utm_campaign=720582&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Привет, Хабр!</p><p>В <a href="https://habr.com/ru/post/717716/" rel="noopener noreferrer nofollow">прошлой статье</a> я рассказывал об ускорении копирования элементов одного слайса в другой с помощью средств Go. В этот раз я решил пойти дальше и посмотреть, что можно достичь, начав разговаривать с процессором на его языке. Я выбрал одну из оптимизированных версий функции <code>Copy</code> в качестве объекта исследования из <a href="https://habr.com/ru/post/716292/" rel="noopener noreferrer nofollow">решения задачи</a> VK Cup'22/23, которая копирует только синий компонент RGBA в Paletted картинку. Если интересно узнать как её ускорить почти в 10 раз, прошу под кат.</p><p></p> <a href="https://habr.com/ru/post/720582/?utm_campaign=720582&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Mon, 06 Mar 2023 07:56:10 GMT</pubDate>
<dc:creator>svistunov</dc:creator>
<category>
<![CDATA[ Высокая производительность ]]>
</category>
<category>
<![CDATA[ Ненормальное программирование ]]>
</category>
<category>
<![CDATA[ Assembler ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ assembler ]]>
</category>
<category>
<![CDATA[ simd ]]>
</category>
<category>
<![CDATA[ avx2 ]]>
</category>
<category>
<![CDATA[ avx ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Шаблон backend сервера на Golang — часть 5 — оптимизация Worker pool ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720286/</guid>
<link>https://habr.com/ru/post/720286/?utm_campaign=720286&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p><a href="https://habr.com/ru/post/720286/"># Шаблон backend сервера на Golang — часть 5 — оптимизация Worker pool</a></p><br> <p><a href="https://habr.com/ru/post/720286/">Пятая часть</a> посвящена оптимизации Worker pool и особенностям его работы в составе микросервиса, развернутого в Kubernetes.</p><br> <p>Представленный Worker pool поддерживает работу с двумя типами задач</p><br> <ul> <li>&quot;Короткие&quot; — не контролируется предельный timeout выполнения и их нельзя прервать</li> <li>&quot;Длинные&quot; — контролируется предельный timeout выполнения и их можно прервать</li> </ul><br> <p>Накладные расходы Worker pool на добавление в очередь, контроль очереди, запуск обработки task, контроль времени выполнения task:</p><br> <ul> <li>Для &quot;коротких&quot; task — от 300 ns/op, 0 B/op, 0 allocs/op</li> <li>Для &quot;длинных&quot; task — от 1400 ns/op, 16 B/op, 1 allocs/op</li> </ul><br> <p><em>Для task, которые должны выполняться быстрее 200 ns/op представленный Worker pool использовать не эффективно</em></p><br> <p>Собираются следующие метрики <a href="https://prometheus.io/" rel="nofollow noopener noreferrer">prometheus</a>:</p><br> <ul> <li>wp_worker_process_count_vec — количество worker в работе</li> <li>wp_task_process_duration_ms_by_name — гистограмма длительности выполнения task в ms с группировкой по task.name</li> <li>wp_task_queue_buffer_len_vec — текущая длина канала-очереди task — показывает заполненность канала</li> <li>wp_add_task_wait_count_vec — количество задач, ожидающих попадания в очередь</li> </ul><br> <p>Ссылка на <a href="https://github.com/romapres2010/goapp" rel="nofollow noopener noreferrer">репозиторий проекта</a>.</p><br> <p>Шаблон goapp в репозитории полностью готов к развертыванию в Docker, Docker Compose, Kubernetes (kustomize), Kubernetes (helm).</p><br> <p>Ссылки на предыдущие части:</p><br> <ul> <li><a href="https://habr.com/ru/post/492062/">Первая часть</a> шаблона была посвящена HTTP серверу.</li> <li><a href="https://habr.com/ru/post/500554/">Вторая часть</a> шаблона была посвящена прототипированию REST API.</li> <li><a href="https://habr.com/ru/post/716634/">Третья часть</a> посвящена развертыванию шаблона в Docker, Docker Compose, Kubernetes (kustomize).</li> <li>Четвертая часть будет посвящена развертыванию в Kubernetes с Helm chart и настройке Horizontal Autoscaler.</li> </ul> <a href="https://habr.com/ru/post/720286/?utm_campaign=720286&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Mon, 06 Mar 2023 07:53:31 GMT</pubDate>
<dc:creator>romapres2010</dc:creator>
<category>
<![CDATA[ Open source ]]>
</category>
<category>
<![CDATA[ API ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ kubernetes ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ helm ]]>
</category>
<category>
<![CDATA[ docker ]]>
</category>
<category>
<![CDATA[ docker-compose ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Тайные каналы связи или как централизованные сервисы способны разлагаться изнутри ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720544/</guid>
<link>https://habr.com/ru/post/720544/?utm_campaign=720544&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Современный мир нельзя представить без сервисов связи, таких как социальные сети, мессенджеры, электронная почта, файловые хранилища и т.п. Мы пользуемся данными сервисами постоянно, ровно как и они планомерно пользуются нами. Конфиденциальная информация становится для сервисов связи святым граалем в задачах таргетированной рекламы. А для нас граалем становится их лёгкость использования, высокая производительность и большие размеры хранилищ.</p><p>На первый взгляд кажется, что такой специфичный способ взаимодействия клиентов с сервисами представляет собой некий симбиоз, при котором мы друг друга дополняем и оберегаем (исключительно ради собственных интересов). Но всё же, если смотреть более углубленно на данный вид коммуникаций, то может появиться дополнительный вопрос: есть ли место в подобной парадигме паразитическим отношениям?</p><p></p> <a href="https://habr.com/ru/post/720544/?utm_campaign=720544&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Sun, 05 Mar 2023 14:30:42 GMT</pubDate>
<dc:creator>Number571</dc:creator>
<category>
<![CDATA[ Децентрализованные сети ]]>
</category>
<category>
<![CDATA[ Информационная безопасность ]]>
</category>
<category>
<![CDATA[ Криптография ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ информационная безопасность ]]>
</category>
<category>
<![CDATA[ криптография ]]>
</category>
<category>
<![CDATA[ программирование ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ децентрализованные сети ]]>
</category>
<category>
<![CDATA[ анонимность ]]>
</category>
<category>
<![CDATA[ анонимные сети ]]>
</category>
<category>
<![CDATA[ централизация ]]>
</category>
<category>
<![CDATA[ тайные каналы связи ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Как я приложение с Go на Rust переписывал ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720382/</guid>
<link>https://habr.com/ru/post/720382/?utm_campaign=720382&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>О Rust я слышал ещё несколько лет назад и все его либо хвалили, либо порицали, по различным причинам, но сам как-то не брался за него - мне, неподготовленному к подобному синтаксису и не знакомому с подобными языками хотя бы на базовом уровне, в то время он казался совершенно непонятным. </p><p>Не так давно решил написать для себя небольшое приложение-бенчмарк для теста HTTP API серверов и написал его на Go. Но размер в 5 с лишним Мбайт, несоблюдение целевого RPS и некоторые другие проблемы заставили посмотреть в сторону более производительного Rust + Tokio + Hyper.<br><br>Эта статья о коде Rust-приложения, переходе с Go на Rust, преимуществах и недостатках обоих языков и сравнении двух сферических коней в вакууме.</p><p></p> <a href="https://habr.com/ru/post/720382/?utm_campaign=720382&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Sat, 04 Mar 2023 03:51:25 GMT</pubDate>
<dc:creator>keigisdead</dc:creator>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Rust ]]>
</category>
<category>
<![CDATA[ rust ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ benchmark ]]>
</category>
<category>
<![CDATA[ сравнение ]]>
</category>
<category>
<![CDATA[ сравнение производительности ]]>
</category>
<category>
<![CDATA[ http ]]>
</category>
<category>
<![CDATA[ tokio ]]>
</category>
<category>
<![CDATA[ hyper ]]>
</category>
</item>
<item>
<title>
<![CDATA[ [recovery mode] Полиморфизм: подавать холодным ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718888/</guid>
<link>https://habr.com/ru/post/718888/?utm_campaign=718888&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Полиморфизм является одним из&nbsp;трёх столпов ООП, наравне с&nbsp;наследованием и инкапсуляцией, да&nbsp;и в&nbsp;целом краеугольным камнем современного программирования. Думаю, большинство читателей не&nbsp;представляет своей жизни без&nbsp;полиморфизма, за&nbsp;что&nbsp;я, конечно, это большинство никак не&nbsp;осуждаю, ибо сам к&nbsp;нему принадлежу. Дело, однако, в&nbsp;том, что&nbsp;многие не&nbsp;задумываются об&nbsp;устройстве этого полиморфизма, ведь любой принцип программирования, по&nbsp;сути, представляет из&nbsp;себя математическую матрёшку.</p><p></p> <a href="https://habr.com/ru/post/718888/?utm_campaign=718888&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Mon, 27 Feb 2023 19:15:57 GMT</pubDate>
<dc:creator>Miiao</dc:creator>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ C++ ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Rust ]]>
</category>
<category>
<![CDATA[ rust ]]>
</category>
<category>
<![CDATA[ cpp ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ polymorphism ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Путь миграции с go build на Bazel ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718360/</guid>
<link>https://habr.com/ru/post/718360/?utm_campaign=718360&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>При поиске решений для сборки больших проектов на Go с завидной регулярностью попадались отсылки на статьи про Bazel.</p><p>К сожалению, понимания того, как должна выглядеть разработка после миграции на Bazel они не давали. Попробуем разобраться...</p><p></p> <a href="https://habr.com/ru/post/718360/?utm_campaign=718360&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Mon, 27 Feb 2023 16:35:17 GMT</pubDate>
<dc:creator>Bozaro</dc:creator>
<category>
<![CDATA[ Блог компании Joom ]]>
</category>
<category>
<![CDATA[ Тестирование IT-систем ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Системы сборки ]]>
</category>
<category>
<![CDATA[ joom ]]>
</category>
<category>
<![CDATA[ bazel ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
</item>
<item>
<title>
<![CDATA[ [Перевод] Брифинг по дженерикам Go 1.18 ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/719262/</guid>
<link>https://habr.com/ru/post/719262/?utm_campaign=719262&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Добавление дженериков (generics) в Go (ранее Golang) — самое значительное изменение, которое он претерпел, с момента его релиза. Сообщество Go просило добавить дженерики с самых первых дней языка, и мы, наконец, дождались.</p><p>Реализация дженериков в Go сильно отличается от традиционных реализаций, которые можно найти в C++, но все-таки имеет некоторое сходство с реализацией дженериков в Rust. В этой статье вашему вниманию представлен обзор, нацеленный помочь вам сформировать понимание дженериков Go и продемонстрировать, как с ними работать.</p><p></p> <a href="https://habr.com/ru/post/719262/?utm_campaign=719262&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Mon, 27 Feb 2023 14:17:33 GMT</pubDate>
<dc:creator>MaxRokatansky</dc:creator>
<category>
<![CDATA[ Блог компании OTUS ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>go</category>
<category>golang</category>
<category>go 1.18</category>
<category>дженерики</category>
<category>планировщик Go</category>
</item>
<item>
<title>
<![CDATA[ Fitter —  сшиватель API/Website's, часть личного проекта которую хотел опенсорснуть ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718972/</guid>
<link>https://habr.com/ru/post/718972/?utm_campaign=718972&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Добрый вечер всем!<br><br>Возможно выбрал не лучшее время для охвата аудитории, но тем не менее главное чтоб продукт был хороший, а не статья о нем. Последние несколько недель я пишу приложение в рамках которого надо собирать огромное количество информации из сети(запросы к API/парсинг HTML кода) и под конец 4-ой интеграции я подумал что надо бы это максимально облегчить(не дело это пересобирать приложение под каждый чих интеграции), возможно это не лучшая преамбула, но хотя бы была реальная проблема решение к которой хотелось показать и заопенсорнуть.<br><br>Итак Fitter = сшиватель достаточно жаргонный перевод, но мне он кажется что лучше всего подходит. Я делал эту штуки исходя из следующих предположений:</p><p></p> <a href="https://habr.com/ru/post/718972/?utm_campaign=718972&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Sat, 25 Feb 2023 23:46:02 GMT</pubDate>
<dc:creator>PYXRU</dc:creator>
<category>
<![CDATA[ Высокая производительность ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ HTML ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Открытые данные ]]>
</category>
<category>
<![CDATA[ scraping ]]>
</category>
<category>
<![CDATA[ парсинг ]]>
</category>
<category>
<![CDATA[ автоматизация ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ monitoring ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Сказ о том как pet-project превратился в небольшой пассивный доход (часть 1) ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718898/</guid>
<link>https://habr.com/ru/post/718898/?utm_campaign=718898&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Я backend разработчик с опытом около 3-х лет, пишу в основном на Golang. Проработал в нескольких крупных российских компаниях. Сейчас я параллельно со своей работой пытаюсь сделать удобный, дешевый VPN сервис с высокой пропускной способностью. В этой статье я хочу просто рассказать про жизненный цикл своего проекта. Возможно кому-то будет просто интересно почитать, а кто-то может почерпнуть что-то новое для себя. </p><p></p> <a href="https://habr.com/ru/post/718898/?utm_campaign=718898&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Sat, 25 Feb 2023 09:21:29 GMT</pubDate>
<dc:creator>tarmalonchik</dc:creator>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Nginx ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ wireguard ]]>
</category>
<category>
<![CDATA[ vpn ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ балансировка нагрузки ]]>
</category>
<category>
<![CDATA[ backend ]]>
</category>
<category>
<![CDATA[ pet-project ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Зачем мигрировать с go build на Bazel? ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718340/</guid>
<link>https://habr.com/ru/post/718340/?utm_campaign=718340&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Это первый пост из цикла, посвященного миграции с go build на Bazel.</p><p>К процессу миграции мы подошли на этапе, когда запуск тестов на CI занимал примерно от 15 минут до часа. При этом мы уже успели реализовать некоторое распараллеливание и кэширование результатов тестов. Без этого тесты на одной машине должны были бы идти примерно часов восемь.</p><p>После внедрения <strong>Bazel запуск тестов на CI в основном укладывается в интервал от 1,5 до 25 минут (50 перцентиль в районе 12 минут)</strong>, что гораздо комфортнее исходной ситуации.</p><p><strong>Оговоримся</strong>, что сравнение этих цифр «в лоб» несколько некорректно: с одной стороны, за время пути кодовая база стала еще больше, а с другой – поменялась топология CI. Но в целом представление о полученном эффекте они дают.</p><p>Далее опишем, за счет какого механизма достигнуто ускорение.</p><p></p> <a href="https://habr.com/ru/post/718340/?utm_campaign=718340&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Wed, 22 Feb 2023 09:26:05 GMT</pubDate>
<dc:creator>Bozaro</dc:creator>
<category>
<![CDATA[ Блог компании Joom ]]>
</category>
<category>
<![CDATA[ Тестирование IT-систем ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Системы сборки ]]>
</category>
<category>
<![CDATA[ joom ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ bazel ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Микросервисные приложения на GoMicro ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718388/</guid>
<link>https://habr.com/ru/post/718388/?utm_campaign=718388&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Go благодаря возможностям компиляции и встроенным механизмам конкурентной многозадачности очень хорошо подходит для создания сетевых приложений и активно используется в создании инструментов для DevOps и распределенных приложений. В этой статье мы рассмотрим некоторые возможности фреймворка GoMicro для реализации микросервисных приложений на Go.</p><p></p> <a href="https://habr.com/ru/post/718388/?utm_campaign=718388&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Tue, 21 Feb 2023 16:11:36 GMT</pubDate>
<dc:creator>dmitriizolotov</dc:creator>
<category>
<![CDATA[ Блог компании OTUS ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ Микросервисы ]]>
</category>
<category>
<![CDATA[ otus ]]>
</category>
<category>
<![CDATA[ microservice ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Rust быстрее всех, Miiao сделал замеры ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718186/</guid>
<link>https://habr.com/ru/post/718186/?utm_campaign=718186&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Пассивно-агрессивный бенчмарк FizzBuzz.</p><p></p> <a href="https://habr.com/ru/post/718186/?utm_campaign=718186&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Mon, 20 Feb 2023 21:28:52 GMT</pubDate>
<dc:creator>Miiao</dc:creator>
<category>
<![CDATA[ C++ ]]>
</category>
<category>
<![CDATA[ C ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ D ]]>
</category>
<category>
<![CDATA[ Rust ]]>
</category>
<category>
<![CDATA[ Fortran ]]>
</category>
<category>
<![CDATA[ Rust ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ C ]]>
</category>
<category>
<![CDATA[ C++ ]]>
</category>
<category>
<![CDATA[ D ]]>
</category>
<category>
<![CDATA[ Python ]]>
</category>
<category>
<![CDATA[ Kotlin ]]>
</category>
<category>
<![CDATA[ Dart ]]>
</category>
<category>
<![CDATA[ PascalABC.NET ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Schema Registry с Protobuf в Kafka — зачем оно надо? ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/715298/</guid>
<link>https://habr.com/ru/post/715298/?utm_campaign=715298&utm_source=habrahabr&utm_medium=rss</link>
<description>
<![CDATA[ <p>Всем привет. Меня зовут Нина Пакшина, я разработчик “Лента Онлайн” и часть операционной команды&nbsp;в сервисе доставки продуктов.</p><p>В данной статье на примере языка Go я расскажу о том, как мы внедряли Kafka в связке с Schema Registry и Protobuf в качестве формата сообщений.</p><p>Я расскажу о том, какие появятся преимущества от использования данных технологий, а также пройдусь по подводным камням, с которыми можно столкнуться при разработке.</p><p></p> <a href="https://habr.com/ru/post/715298/?utm_campaign=715298&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut">Читать далее</a> ]]>
</description>
<pubDate>Mon, 20 Feb 2023 07:41:39 GMT</pubDate>
<dc:creator>Ninako</dc:creator>
<category>
<![CDATA[ Блог компании LENTA:U TECH ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ kafka ]]>
</category>
<category>
<![CDATA[ schema registry ]]>
</category>
<category>
<![CDATA[ protobuf ]]>
</category>
<category>
<![CDATA[ golang ]]>
</category>
<category>
<![CDATA[ валидация данных ]]>
</category>
</item>
</channel>
</rss>`

var Data3 string = `
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" version="2.0">
<channel>
<title>Лучшие публикации за сутки</title>
<link>https://habr.com/ru/top/daily/</link>
<description>
<![CDATA[ Лучшие публикации за последние 24 часа ]]>
</description>
<language>ru</language>
<managingEditor>editor@habr.com</managingEditor>
<generator>habr.com</generator>
<pubDate>Fri, 10 Mar 2023 00:46:59 GMT</pubDate>
<image>
<link>https://habr.com/ru/</link>
<url>https://habrastorage.org/webt/ym/el/wk/ymelwk3zy1gawz4nkejl_-ammtc.png</url>
<title>Хабр</title>
</image>
<item>
<title>
<![CDATA[ 5 классных сервисов на основе ИИ (с примерами) ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721114/</guid>
<link>https://habr.com/ru/post/721114/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721114</link>
<description>
<![CDATA[ <p>За&nbsp;последний год появилось огромное число новых сервисов, которые работают на&nbsp;нейронных сетях. Кажется, что&nbsp;уже не&nbsp;осталось людей, кто&nbsp;бы не&nbsp;слышал о&nbsp;том, что&nbsp;chatGPT пишет новости, а&nbsp;Midjorney во&nbsp;всю создает шедевры.<br><br>Сегодня я&nbsp;бы хотел показать 5&nbsp;сервисов на&nbsp;основе нейронных сетей, которые <strong>не&nbsp;связаны</strong> с <strong>chatGPT</strong>, <strong>Midjorney</strong> или <strong>Stable Diffusion</strong>. Эти сервисы помогают обрабатывать аудио, преобразовывать текст в&nbsp;речь и удалять ненужные вещи с&nbsp;картинок.</p><p></p> <a href="https://habr.com/ru/post/721114/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721114#habracut">Ознакомиться</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 05:30:02 GMT</pubDate>
<dc:creator>daniilgorbenko</dc:creator>
<category>
<![CDATA[ Блог компании RUVDS.com ]]>
</category>
<category>
<![CDATA[ Программирование ]]>
</category>
<category>
<![CDATA[ Машинное обучение ]]>
</category>
<category>
<![CDATA[ Облачные сервисы ]]>
</category>
<category>
<![CDATA[ Искусственный интеллект ]]>
</category>
<category>
<![CDATA[ подборки ]]>
</category>
<category>
<![CDATA[ искусственный интеллект ]]>
</category>
<category>
<![CDATA[ сервисы ]]>
</category>
<category>
<![CDATA[ ai ]]>
</category>
<category>
<![CDATA[ text to speech ]]>
</category>
<category>
<![CDATA[ синтез речи ]]>
</category>
<category>
<![CDATA[ обработка изображений ]]>
</category>
<category>
<![CDATA[ подборка сервисов ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Как ИИ работает даже в зоне взрывных работ ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/706022/</guid>
<link>https://habr.com/ru/post/706022/?utm_source=habrahabr&utm_medium=rss&utm_campaign=706022</link>
<description>
<![CDATA[ <p>Чтобы добыть железную руду, породу нужно рызрыхлить. В карьере СГОКа (Стойленского горно-обогатительного комбината) этого достигают посредством буровзрывных работ. После взрыва горную массу нужно погрузить в карьерный самосвал, а потом в вагон-думпкар и отправить на обогатительную фабрику.</p><p> Казалось бы, что может быть проще — черпай себе экскаватором да высыпай. А вот нет — тут легко допустить перегруз или же, наоборот, недогруз. Даже если средние показатели в норме, из-за таких вот «небольших» погрешностей мы на круг недовозили на фабрику 2-3% породы в сравнении с учетной нормой. Приходилось запускать дополнительные рейсы. Перегруз еще и опасен для транспорта — он повышает износ деталей и расход топлива, увеличивает риск выпадения кусков породы из вагона или кузова.</p><p>Мы на НЛМК очень любим ИИ, математические модели и прочие нейросети – вот их и взяли, чтобы повысить эффективность транспортировки железной руды с карьера на фабрику.</p><p></p> <a href="https://habr.com/ru/post/706022/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=706022#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 07:32:53 GMT</pubDate>
<dc:creator>kokorina_es</dc:creator>
<category>
<![CDATA[ Блог компании Группа НЛМК ]]>
</category>
<category>
<![CDATA[ Машинное обучение ]]>
</category>
<category>
<![CDATA[ машинное обучение ]]>
</category>
<category>
<![CDATA[ производство ]]>
</category>
<category>
<![CDATA[ машинное зрение ]]>
</category>
<category>
<![CDATA[ computer vision ]]>
</category>
<category>
<![CDATA[ mashine-learning ]]>
</category>
<category>
<![CDATA[ руда ]]>
</category>
<category>
<![CDATA[ добыча руды ]]>
</category>
</item>
<item>
<title>
<![CDATA[ 4 электронных музыкальных шкатулки своими руками. Знакомимся с логическими микросхемами ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/717906/</guid>
<link>https://habr.com/ru/post/717906/?utm_source=habrahabr&utm_medium=rss&utm_campaign=717906</link>
<description>
<![CDATA[ <div style="text-align:center;"><img src="https://habrastorage.org/webt/eb/_r/wd/eb_rwd8fpqjxtocsrmcbavhhfiu.jpeg"></div><br> Привет, Хабр! Современные смартфоны и встраиваемые микрокомпьютеры могут всё или почти всё, но интерес к самоделкам на дискретных радиодеталях и простых микросхемах никогда не угаснет. И это хорошо.<br> <br> Потому что собрать и при необходимости наладить вещь, принцип работы которой понимаем, — это реализация одной из фундаментальных потребностей психики человека разумного. Утвердить своё субъектное место в объективном мире.<br> <br> Разыскивать или разрабатывать схемы, травить печатные платы может и хочет не каждый, потому что тут требуется и время, и знания, и оборудование. А чтобы собрать дешёвый китайский радиоконструктор, нужны только паяльник и бокорезы. Потому это прекрасный вариант хобби и возможность создавать замечательные подарки своими руками.<br> <a href="https://habr.com/ru/post/717906/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=717906#habracut">Наши сегодняшние самоделки звучат и играют светом</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 09:00:01 GMT</pubDate>
<dc:creator>Lunathecat</dc:creator>
<category>
<![CDATA[ Блог компании RUVDS.com ]]>
</category>
<category>
<![CDATA[ Читальный зал ]]>
</category>
<category>
<![CDATA[ Старое железо ]]>
</category>
<category>
<![CDATA[ DIY или Сделай сам ]]>
</category>
<category>
<![CDATA[ Электроника для начинающих ]]>
</category>
<category>
<![CDATA[ ruvds_статьи ]]>
</category>
<category>
<![CDATA[ diy или сделай сам ]]>
</category>
<category>
<![CDATA[ музыкальная шкатулка ]]>
</category>
<category>
<![CDATA[ логические элементы ]]>
</category>
<category>
<![CDATA[ микросхемы ]]>
</category>
<category>
<![CDATA[ семплы ]]>
</category>
<category>
<![CDATA[ звуковые эффекты ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Обойти айти: есть ли жизнь за пределами золотой клетки? ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721206/</guid>
<link>https://habr.com/ru/post/721206/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721206</link>
<description>
<![CDATA[ Уйти из айти невероятно сложно: престижная работа, неизменный wow-эффект у знакомых, заработная плата и условия выше рынка или как минимум лучше среднего, постоянное развитие, интересные отраслевые события, отличные отношения с удалёнкой… Да ну, где вы видели этих безумцев?! А они есть: люди, которые в 25, 35, 40, 45 лет бросают должности, проекты и компании и уходят в мир других задач, интересов, людей. Кто-то педантично готовит план, подстилает соломку и аккуратно шагает с корабля на корабль, кто-то громко хлопает дверью и знает, куда идти и что делать, а кто-то уходит в никуда, сидит дома, лечит выгорание и ломку привычки к упорной нагруженной работе… В целом, это болезненный переход: как по деньгам, там и по моральным компонентам. Решиться невероятно трудно. Так давайте разберёмся, зачем ломать прутья золотой клетки и что именно щекочет воздух свободы.<br> <br> <a href="https://habr.com/ru/company/ruvds/blog/721206/"><div style="text-align:center;"><img src="https://habrastorage.org/webt/ej/i_/m-/eji_m-4ui9i8vvqrd7e960d92eo.png"></div></a> <a href="https://habr.com/ru/post/721206/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721206#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 13:00:06 GMT</pubDate>
<dc:creator>ru_vds</dc:creator>
<category>
<![CDATA[ Блог компании RUVDS.com ]]>
</category>
<category>
<![CDATA[ Управление персоналом ]]>
</category>
<category>
<![CDATA[ Карьера в IT-индустрии ]]>
</category>
<category>
<![CDATA[ ruvds_статьи ]]>
</category>
<category>
<![CDATA[ карьера ]]>
</category>
<category>
<![CDATA[ работа ]]>
</category>
<category>
<![CDATA[ специальности ]]>
</category>
<category>
<![CDATA[ обучение ]]>
</category>
<category>
<![CDATA[ смена профессии ]]>
</category>
<category>
<![CDATA[ IT ]]>
</category>
</item>
<item>
<title>
<![CDATA[ [Перевод] Мы обнаружили в GPT-2 нейрон конкретного токена ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/718720/</guid>
<link>https://habr.com/ru/post/718720/?utm_source=habrahabr&utm_medium=rss&utm_campaign=718720</link>
<description>
<![CDATA[ <div style="text-align:center;"><img src="https://habrastorage.org/webt/wh/l3/qu/whl3qumzo28sn0hlriqwkgymyus.png"></div><br> Мы начали с вопроса: откуда GPT-2 знает, когда использовать слово <code>an</code>, а не <code>a</code>? Выбор зависит от того, начинается ли следующее за ним слово с гласной, однако GPT-2 может прогнозировать только одно слово за раз.<br> <br> У нас по-прежнему нет полного ответа, однако мы нашли нейрон MLP в GPT-2 Large, который необходим для прогнозирования токена &quot; an&quot;. Также мы выяснили, что веса этого нейрона соотносятся с эмбеддингом токена &quot; an&quot;, что позволило нам найти другие нейроны, прогнозирующие конкретный токен.<br> <a href="https://habr.com/ru/post/718720/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=718720#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 06:54:57 GMT</pubDate>
<dc:creator>PatientZero</dc:creator>
<category>
<![CDATA[ Математика ]]>
</category>
<category>
<![CDATA[ Машинное обучение ]]>
</category>
<category>
<![CDATA[ Искусственный интеллект ]]>
</category>
<category>gpt-2</category>
<category>нейронные сети</category>
<category>трансформеры</category>
<category>генеративные модели</category>
<category>нейроны</category>
<category>токены</category>
</item>
<item>
<title>
<![CDATA[ Семейство алгоритмов Ascon — новый стандарт легковесной криптографии ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721194/</guid>
<link>https://habr.com/ru/post/721194/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721194</link>
<description>
<![CDATA[ <img src="https://habrastorage.org/webt/3i/qx/9a/3iqx9a8rdc3sgbk8o0oap5y0uf4.png"><br> <font color="gray"><i>Режим работы шифра Ascon, см. <a href="https://habrastorage.org/webt/63/n6/8d/63n68d3hnfqxbudyywvnoywde3y.png" rel="nofollow noopener noreferrer">список условных обозначений</a> на схеме</i></font><br> <br> В феврале 2023 года Национальный институт стандартов и технологий (NIST) <a href="https://csrc.nist.gov/News/2023/lightweight-cryptography-nist-selects-ascon" rel="nofollow noopener noreferrer">выбрал стандарт</a> легковесной криптографии для RFID, датчиков, Интернета вещей и других устройств с ограниченными аппаратными ресурсами. Победителем конкурса стало семейство шифров <a href="https://ascon.iaik.tugraz.at/" rel="nofollow noopener noreferrer">Ascon</a> (<a href="https://csrc.nist.gov/CSRC/media/Projects/lightweight-cryptography/documents/finalist-round/updated-submissions/ascon.zip" rel="nofollow noopener noreferrer">файл zip</a>, <a href="https://csrc.nist.gov/CSRC/media/Projects/lightweight-cryptography/documents/finalist-round/updated-spec-doc/ascon-spec-final.pdf" rel="nofollow noopener noreferrer">спецификации</a>, <a href="https://csrc.nist.gov/CSRC/media/Projects/lightweight-cryptography/documents/finalist-round/changelog-files/ascon-changelog-final.pdf" rel="nofollow noopener noreferrer">changelog</a>).<br> <a href="https://habr.com/ru/post/721194/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721194#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 05:57:18 GMT</pubDate>
<dc:creator>GlobalSign_admin</dc:creator>
<category>
<![CDATA[ Блог компании GlobalSign ]]>
</category>
<category>
<![CDATA[ Информационная безопасность ]]>
</category>
<category>
<![CDATA[ Криптография ]]>
</category>
<category>
<![CDATA[ Алгоритмы ]]>
</category>
<category>
<![CDATA[ Интернет вещей ]]>
</category>
<category>
<![CDATA[ NIST ]]>
</category>
<category>
<![CDATA[ Ascon ]]>
</category>
<category>
<![CDATA[ легковесная криптография ]]>
</category>
<category>
<![CDATA[ AES ]]>
</category>
<category>
<![CDATA[ FIPS 197 ]]>
</category>
<category>
<![CDATA[ Galois Counter Mode ]]>
</category>
<category>
<![CDATA[ SP 800-38D ]]>
</category>
<category>
<![CDATA[ хэширование ]]>
</category>
<category>
<![CDATA[ SHA-256 ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Какой сервис сетевой связности использовать: глобальный роутер Selectel, Direct или Global Connect? ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721274/</guid>
<link>https://habr.com/ru/post/721274/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721274</link>
<description>
<![CDATA[ <div style="text-align:center;"><img src="https://habrastorage.org/webt/gg/dm/oq/ggdmoq8za4hpdk591l56sqojauu.png"></div><br> Существует много способов, как организовать инфраструктуру. Например, можно объединить on-premise с сервером или облаком провайдера. Или собрать географически распределенную инфраструктуру в нескольких регионах.<br> <br> Но как объединить разные серверы и проекты в одну сеть? Какую услугу использовать: глобальный роутер Selectel, Direct или Global Connect — и в чем разница? С такими же вопросами к нам приходят клиенты. Поэтому мы постарались дать гайд по выбору услуги. О том, что из этого получилось, рассказываем под катом.<br> <a href="https://habr.com/ru/post/721274/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721274#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 11:41:17 GMT</pubDate>
<dc:creator>Doctor_IT</dc:creator>
<category>
<![CDATA[ Блог компании Selectel ]]>
</category>
<category>
<![CDATA[ Высокая производительность ]]>
</category>
<category>
<![CDATA[ IT-инфраструктура ]]>
</category>
<category>
<![CDATA[ Сетевые технологии ]]>
</category>
<category>
<![CDATA[ Распределённые системы ]]>
</category>
<category>
<![CDATA[ selectel ]]>
</category>
<category>
<![CDATA[ direct connect ]]>
</category>
<category>
<![CDATA[ global connect ]]>
</category>
<category>
<![CDATA[ глобальный роутер ]]>
</category>
<category>
<![CDATA[ l3vpn ]]>
</category>
<category>
<![CDATA[ l2 ]]>
</category>
<category>
<![CDATA[ сетевая связность ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Кастомная стратегия виртуального скроллинга для просмотра pdf ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721220/</guid>
<link>https://habr.com/ru/post/721220/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721220</link>
<description>
<![CDATA[ <p>Angular CDK имеет широкие возможности для скроллинга плоского списка. Если размер каждого элемента одинаков, то можно воспользоваться <code>FixedSizeVirtualScrollStrategy</code>: всего лишь нужно прокинуть размер элемента в пикселях, проитерироваться по данным и виртуальный скроллинг готов. Но что делать, если размер элементов разный? Данную проблему можно решить кастомной стратегией виртуального скроллинга. В данной статье мы рассмотрим как построить такую стратегию для pdf-документов.</p><p></p> <a href="https://habr.com/ru/post/721220/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721220#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 07:57:23 GMT</pubDate>
<dc:creator>mishqua</dc:creator>
<category>
<![CDATA[ Блог компании Bimeister ]]>
</category>
<category>
<![CDATA[ Разработка веб-сайтов ]]>
</category>
<category>
<![CDATA[ JavaScript ]]>
</category>
<category>
<![CDATA[ Angular ]]>
</category>
<category>
<![CDATA[ TypeScript ]]>
</category>
<category>
<![CDATA[ angular ]]>
</category>
<category>
<![CDATA[ typescript ]]>
</category>
<category>
<![CDATA[ scroll ]]>
</category>
<category>
<![CDATA[ virtual scroll ]]>
</category>
<category>
<![CDATA[ scrolling ]]>
</category>
<category>
<![CDATA[ cdk ]]>
</category>
<category>
<![CDATA[ infinite scrolling ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Разгон игры «Fred» для ZX Spectrum ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721386/</guid>
<link>https://habr.com/ru/post/721386/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721386</link>
<description>
<![CDATA[ <p>История о том, как я вернулся к любимой игре своего детства, немного узнал о том, как она работает, сделал так, чтобы играть в неё было приятнее и интереснее. Маленький кусочек ретро-археологии.</p><p></p> <a href="https://habr.com/ru/post/721386/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721386#habracut">Спуститься в подземелье</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 16:47:32 GMT</pubDate>
<dc:creator>kmatveev</dc:creator>
<category>
<![CDATA[ Разработка игр ]]>
</category>
<category>
<![CDATA[ Реверс-инжиниринг ]]>
</category>
<category>
<![CDATA[ Игры и игровые консоли ]]>
</category>
<category>
<![CDATA[ zx spectrum ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Кем работать в IT: scrum-мастер ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721284/</guid>
<link>https://habr.com/ru/post/721284/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721284</link>
<description>
<![CDATA[ <p><strong>Рубрика «Кем работать в IT»</strong> — интервью с представителями интересных IT-профессий, в которых специалисты рассказывают о тонкостях своей работы: плюсах, минусах, подводных камнях и заработной плате. Начинающие смогут больше узнать о том, что их ожидает на карьерном пути, а профессионалы — посмотреть на свою специальность через чужой опыт и, может быть, открыть для себя что-то новое.</p><p>Сегодня о своем опыте работы нам расскажет Елена, Scrum-мастер в Центральном банке Российской Федерации (Банк России).</p><p></p> <a href="https://habr.com/ru/post/721284/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721284#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 10:52:08 GMT</pubDate>
<dc:creator>habr_career</dc:creator>
<category>
<![CDATA[ Блог компании Хабр Карьера ]]>
</category>
<category>
<![CDATA[ Управление персоналом ]]>
</category>
<category>
<![CDATA[ Карьера в IT-индустрии ]]>
</category>
<category>
<![CDATA[ кем работать в it ]]>
</category>
<category>
<![CDATA[ карьера в it ]]>
</category>
<category>
<![CDATA[ scrum-мастер ]]>
</category>
<category>
<![CDATA[ системный анализ ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Добавляем деструкторы в C ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721188/</guid>
<link>https://habr.com/ru/post/721188/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721188</link>
<description>
<![CDATA[ <p>В&nbsp;данной статье будет описано создание кастомного аллокатора на&nbsp;си c регистрацией колбеков, которые будут вызваны при&nbsp;освобождении памяти. Нужен для&nbsp;того, чтобы при&nbsp;создании записать туда деструктор, а&nbsp;в&nbsp;конце просто вызвать free, не&nbsp;погружаясь в&nbsp;детали его работы.</p><p></p> <a href="https://habr.com/ru/post/721188/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721188#habracut">read(&amp;publication);</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 09:15:14 GMT</pubDate>
<dc:creator>orenty7</dc:creator>
<category>
<![CDATA[ C ]]>
</category>
<category>
<![CDATA[ аллокатор ]]>
</category>
<category>
<![CDATA[ callback ]]>
</category>
<category>
<![CDATA[ язык c ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Гуёвая автоматизация управления кластерами ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720256/</guid>
<link>https://habr.com/ru/post/720256/?utm_source=habrahabr&utm_medium=rss&utm_campaign=720256</link>
<description>
<![CDATA[ <p>Если вы активно используете kubernetes в&nbsp;своей инфраструктуре, при&nbsp;этому у&nbsp;вас небольшая команда или&nbsp;она состоит в&nbsp;основном из&nbsp;разработчиков, то у&nbsp;меня к&nbsp;вам вопрос: ну как&nbsp;вам&nbsp;— стала жизнь легче? Наверное те, кто используют managed‑решения в&nbsp;некотором роде покивают головой, продавцы этих решений скажут «да!», с&nbsp;особенно довольным&nbsp;лицом, а&nbsp;бизнес, пуская скупую слезу, просто согласятся с&nbsp;большинством (ну бизнес&nbsp;же растёт).</p><p>Тот инструмент, про&nbsp;который я сегодня хочу рассказать подходит в&nbsp;большей степени для&nbsp;самого что&nbsp;ни на&nbsp;есть микросервисного и девопснутого подхода, когда команды разработчиков имеют необходимую и достаточную абстракцию для&nbsp;самостоятельного управления кластерами, при&nbsp;этом команда эксплуатации сохраняет контроль за&nbsp;всем. Речь пойдёт про&nbsp;Rancher и около стоящие продукты.</p><p></p> <a href="https://habr.com/ru/post/720256/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=720256#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 06:17:01 GMT</pubDate>
<dc:creator>onegreyonewhite</dc:creator>
<category>
<![CDATA[ Системное администрирование ]]>
</category>
<category>
<![CDATA[ IT-инфраструктура ]]>
</category>
<category>
<![CDATA[ Виртуализация ]]>
</category>
<category>
<![CDATA[ DevOps ]]>
</category>
<category>
<![CDATA[ Kubernetes ]]>
</category>
<category>
<![CDATA[ kubernetes ]]>
</category>
<category>
<![CDATA[ k3s ]]>
</category>
<category>
<![CDATA[ rke ]]>
</category>
<category>
<![CDATA[ rancher ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Первая бесплатная модель перевода с русского на китайский язык и обратно ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721330/</guid>
<link>https://habr.com/ru/post/721330/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721330</link>
<description>
<![CDATA[ <p>Представляю вашему вниманию, первую бесплатную offline модель по&nbsp;переводу с&nbsp;русского языка на&nbsp;китайский и обратно.</p><p>Ранее, я писал, как&nbsp;можно достаточно легко обучить свою модель по&nbsp;машинному переводу на&nbsp;примере перевода с&nbsp;английского на&nbsp;русский.</p><p>В&nbsp;этот раз я решил, реализовать, модель перевода с&nbsp;китайского языка, так как&nbsp;давно хотел и о&nbsp;чем заявлял в&nbsp;комментариях к&nbsp;предыдущей своей статье.</p><p></p> <a href="https://habr.com/ru/post/721330/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721330#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 13:19:22 GMT</pubDate>
<dc:creator>UtrobinMV</dc:creator>
<category>
<![CDATA[ Data Mining ]]>
</category>
<category>
<![CDATA[ Машинное обучение ]]>
</category>
<category>
<![CDATA[ Искусственный интеллект ]]>
</category>
<category>
<![CDATA[ Natural Language Processing ]]>
</category>
<category>
<![CDATA[ Data Engineering ]]>
</category>
<category>
<![CDATA[ машинный перевод ]]>
</category>
<category>
<![CDATA[ русско-китайский переводчик ]]>
</category>
</item>
<item>
<title>
<![CDATA[ «Ветхий завет» речевых технологий. Говорящая голова, металлические языки и безумные синтезаторы ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721304/</guid>
<link>https://habr.com/ru/post/721304/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721304</link>
<description>
<![CDATA[ <p>Привет. Меня зовут Александр Родченков, я занимаюсь речевыми технологиями в компании «Инфосистемы Джет». Как-то я задался вопросом — когда люди стали пытаться синтезировать или распознавать речь? Изучив вопрос, раскопал много криповых любопытных историй и решил с вами поделиться.</p><p></p> <a href="https://habr.com/ru/post/721304/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721304#habracut">Узнать больше</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 11:35:28 GMT</pubDate>
<dc:creator>JetHabr</dc:creator>
<category>
<![CDATA[ Блог компании Инфосистемы Джет ]]>
</category>
<category>
<![CDATA[ История IT ]]>
</category>
<category>
<![CDATA[ Старое железо ]]>
</category>
<category>
<![CDATA[ Звук ]]>
</category>
<category>
<![CDATA[ speech synthesis ]]>
</category>
<category>
<![CDATA[ tts ]]>
</category>
<category>
<![CDATA[ речевые технологии ]]>
</category>
<category>
<![CDATA[ история ]]>
</category>
<category>
<![CDATA[ железо ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Эволюция алгоритма фильтрации модификаций товаров в Авито ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720880/</guid>
<link>https://habr.com/ru/post/720880/?utm_source=habrahabr&utm_medium=rss&utm_campaign=720880</link>
<description>
<![CDATA[ <p>Всем привет! Меня зовут Денис Колпаков, я бэкенд-инженер в юните Core Services Авито. Долгое время я был овнером критически значимого для бизнеса сервиса форм, а последний год занимаюсь каталогами и каталогизацией.&nbsp;</p><p>Я расскажу, как мы решали продуктовую задачу — искали способ отфильтровать модификации товаров из базы данных.&nbsp;</p><p></p> <a href="https://habr.com/ru/post/720880/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=720880#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 11:32:18 GMT</pubDate>
<dc:creator>Tifongod</dc:creator>
<category>
<![CDATA[ Блог компании AvitoTech ]]>
</category>
<category>
<![CDATA[ Go ]]>
</category>
<category>
<![CDATA[ go ]]>
</category>
<category>
<![CDATA[ bitmap ]]>
</category>
<category>
<![CDATA[ фильтрация ]]>
</category>
<category>
<![CDATA[ алгоритмы ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Инструменты для MLOps: выбираем между вендорскими и Open Source-решениями ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720654/</guid>
<link>https://habr.com/ru/post/720654/?utm_source=habrahabr&utm_medium=rss&utm_campaign=720654</link>
<description>
<![CDATA[ <img src="https://habrastorage.org/webt/hq/cd/e9/hqcde9ss5q47pam4jmwpoticfq8.jpeg"><br> <br> MLOps использует проверенные методы DevOps для автоматизации создания, развертывания и мониторинга конвейеров ML в производственной среде. По мере развития MLOps-инструментов для работы с ним становится больше — как проприетарных, так и Open Source. Из этого разнообразия часто сложно выбрать стек для своего проекта.<br> <br> Меня зовут Александр Волынский, я технический менеджер Cloud ML Platform в VK Cloud. В этой статье я сравню подходы к работе с MLOps на основе Open Source и проприетарного ПО и расскажу, какие инструменты и почему мы выбрали для <a href="https://mcs.mail.ru/machine-learning/?utm_source=habr&amp;utm_medium=media&amp;utm_campaign=mlops-vendor-open" rel="nofollow noopener noreferrer">Cloud ML Platform</a>.<br> <a href="https://habr.com/ru/post/720654/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=720654#habracut">Читать дальше &rarr;</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 08:53:13 GMT</pubDate>
<dc:creator>volinski</dc:creator>
<category>
<![CDATA[ Блог компании VK ]]>
</category>
<category>
<![CDATA[ Big Data ]]>
</category>
<category>
<![CDATA[ Машинное обучение ]]>
</category>
<category>
<![CDATA[ vk cloud ]]>
</category>
<category>
<![CDATA[ MLOps ]]>
</category>
<category>
<![CDATA[ open source ]]>
</category>
<category>
<![CDATA[ ML ]]>
</category>
<category>
<![CDATA[ вендорское ПО ]]>
</category>
<category>
<![CDATA[ KubeFlow ]]>
</category>
<category>
<![CDATA[ MLFlow ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Разбираем лучшие решения задач с VK Cup ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/720956/</guid>
<link>https://habr.com/ru/post/720956/?utm_source=habrahabr&utm_medium=rss&utm_campaign=720956</link>
<description>
<![CDATA[ <p>В&nbsp;начале февраля мы наградили победителей нашего IT‑чемпионата VK Cup. До&nbsp;финала дошли 80&nbsp;человек, а&nbsp;общий призовой фонд в 4&nbsp;млн рублей разделили 20&nbsp;победителей&nbsp;— по&nbsp;четыре команды на&nbsp;каждый трек чемпионата. На&nbsp;Хабре мы решили сделать серию статей с&nbsp;разбором наиболее интересных решений по&nbsp;разным трекам. </p><p>Сегодня мы расскажем про&nbsp;трек по <strong>машинному обучению. </strong>Участники работали над улучшением рекомендаций для&nbsp;друзей и предсказанием взаимного добавления в&nbsp;друзья между пользователями ВКонтакте. Все мы пользуемся соцсетями и можем на&nbsp;собственном примере ощутить, как&nbsp;работают рекомендации друзей. Но&nbsp;что&nbsp;под&nbsp;капотом этого решения, откуда соцсеть знает, с&nbsp;кем мне стоит подружиться?</p><p>Во&nbsp;всём этом разобрался <a href="https://github.com/BraginIvan/vkcup2022"><u>Иван Брагин</u></a>, один из&nbsp;победителей чемпионата. </p><p></p> <a href="https://habr.com/ru/post/720956/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=720956#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 07:55:21 GMT</pubDate>
<dc:creator>lperovskaya</dc:creator>
<category>
<![CDATA[ Блог компании VK ]]>
</category>
<category>
<![CDATA[ Алгоритмы ]]>
</category>
<category>
<![CDATA[ Математика ]]>
</category>
<category>
<![CDATA[ ml ]]>
</category>
<category>
<![CDATA[ спортивное программирование ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Зачем в Hoff Tech архитекторы или как мы строим и описываем ИТ-ландшафт ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721258/</guid>
<link>https://habr.com/ru/post/721258/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721258</link>
<description>
<![CDATA[ <p>Мы последовательно внедряем архитектурный подход в давно работающей компании, буквально на ходу — это напоминает починку работающего двигателя. Здесь неизбежны некоторые особенности, о которых стоит поговорить.</p><p>Спойлер: процесс идет, мы набили шишки и выработали подходы, которые хочется показать и обсудить с коллегами. Этот пост — первый из серии статей, где я изложу свое видение работы архитектора и пошагово расскажу, как мы внедряли и практикуем архитектурный подход.</p><p></p> <a href="https://habr.com/ru/post/721258/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721258#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 16:42:55 GMT</pubDate>
<dc:creator>DaTa419</dc:creator>
<category>
<![CDATA[ Блог компании Hoff Tech ]]>
</category>
<category>
<![CDATA[ Анализ и проектирование систем ]]>
</category>
<category>
<![CDATA[ Проектирование и рефакторинг ]]>
</category>
<category>
<![CDATA[ IT-стандарты ]]>
</category>
<category>
<![CDATA[ архитектура ]]>
</category>
<category>
<![CDATA[ паттерны проектирования ]]>
</category>
<category>
<![CDATA[ шаблоны проектирования ]]>
</category>
<category>
<![CDATA[ архитектурный комитет ]]>
</category>
<category>
<![CDATA[ it-ландшафт ]]>
</category>
<category>
<![CDATA[ archimate ]]>
</category>
<category>
<![CDATA[ архитектурные паттерны ]]>
</category>
<category>
<![CDATA[ архитектурные шаблоны ]]>
</category>
<category>
<![CDATA[ it-архитектура ]]>
</category>
<category>
<![CDATA[ корпоративная архитектура ]]>
</category>
</item>
<item>
<title>
<![CDATA[ Типизируй с нами, типизируй, как мы… ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721346/</guid>
<link>https://habr.com/ru/post/721346/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721346</link>
<description>
<![CDATA[ <p>Сказ о том, как я каррирование типизировал</p><p></p> <a href="https://habr.com/ru/post/721346/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721346#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 13:52:28 GMT</pubDate>
<dc:creator>SumarokovVladimir</dc:creator>
<category>
<![CDATA[ Разработка веб-сайтов ]]>
</category>
<category>
<![CDATA[ JavaScript ]]>
</category>
<category>
<![CDATA[ TypeScript ]]>
</category>
<category>
<![CDATA[ typescript ]]>
</category>
<category>
<![CDATA[ curry ]]>
</category>
<category>
<![CDATA[ currying ]]>
</category>
<category>
<![CDATA[ типизация ]]>
</category>
<category>
<![CDATA[ типы ]]>
</category>
</item>
<item>
<title>
<![CDATA[ [recovery mode] Как не потерять всю переписку в Slack? Подробная инструкция, что можно сейчас сделать ]]>
</title>
<guid isPermaLink="true">https://habr.com/ru/post/721318/</guid>
<link>https://habr.com/ru/post/721318/?utm_source=habrahabr&utm_medium=rss&utm_campaign=721318</link>
<description>
<![CDATA[ <p>В&nbsp;конце февраля Slack разослал владельцам пространств из&nbsp;России письма, где кому‑то объявил о&nbsp;блокировке пространства, а&nbsp;кому‑то о&nbsp;полном отключении тех. поддержки с 21&nbsp;марта. День Х все ближе, поэтому рассказываем, как&nbsp;сохранить свои данные и как/куда их можно перевезти.</p><p></p> <a href="https://habr.com/ru/post/721318/?utm_source=habrahabr&amp;utm_medium=rss&amp;utm_campaign=721318#habracut">Читать далее</a> ]]>
</description>
<pubDate>Thu, 09 Mar 2023 12:16:43 GMT</pubDate>
<dc:creator>pachcacom</dc:creator>
<category>
<![CDATA[ Мессенджеры ]]>
</category>
<category>
<![CDATA[ Slack ]]>
</category>
<category>
<![CDATA[ экспорт ]]>
</category>
<category>
<![CDATA[ импортозамещение ]]>
</category>
<category>
<![CDATA[ импорт ]]>
</category>
<category>
<![CDATA[ российское по ]]>
</category>
<category>
<![CDATA[ корпоративный мессенджер ]]>
</category>
</item>
</channel>
</rss>`