package templates

import (
	"html/template"
)

type StyleOverrides struct {
	PaddingPx       int    `json:"padding_px,omitempty"`
	TitleFontSizePx int    `json:"title_font_size_px,omitempty"`
	CardBgColor     string `json:"card_bg_color,omitempty"`
	TextMainColor   string `json:"text_main_color,omitempty"`
}

// CardData is the context for the HTML template
type CardData struct {
	Content     template.HTML
	Author      string
	Title       string
	Theme       string
	AccentColor string
	Overrides   StyleOverrides
}

// SocialCard is the base HTML template for our generator
const SocialCard = `
<!DOCTYPE html>
<html>
<head>
	<style>
		:root {
			{{if eq .Theme "light"}}
			--bg-gradient: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
			--card-bg: rgba(255, 255, 255, 0.7);
			--text-main: #0f172a;
			--text-muted: #64748b;
			--border-color: rgba(0, 0, 0, 0.1);
			{{else}}
			--bg-gradient: linear-gradient(135deg, #0f172a 0%, #1e293b 100%);
			--card-bg: rgba(255, 255, 255, 0.04);
			--text-main: #f8fafc;
			--text-muted: #94a3b8;
			--border-color: rgba(255, 255, 255, 0.12);
			{{end}}

			{{if .AccentColor}}
			--accent: {{.AccentColor}};
			{{else}}
			--accent: #38bdf8;
			{{end}}

			{{if .Overrides.CardBgColor}}
			--card-bg: {{.Overrides.CardBgColor}};
			{{end}}

			{{if .Overrides.TextMainColor}}
			--text-main: {{.Overrides.TextMainColor}};
			{{end}}

			--card-padding: {{if .Overrides.PaddingPx}}{{.Overrides.PaddingPx}}px{{else}}48px 48px 36px 48px{{end}};
			--title-font-size: {{if .Overrides.TitleFontSizePx}}{{.Overrides.TitleFontSizePx}}px{{else}}14px{{end}};
		}
		body {
			width: 800px;
			height: 420px;
			margin: 0;
			display: flex;
			justify-content: center;
			align-items: center;
			background: var(--bg-gradient);
			font-family: 'Inter', system-ui, -apple-system, sans-serif;
			color: var(--text-main);
		}
		.container {
			width: 720px;
			height: 340px; /* taller container to accommodate footer */
			background: var(--card-bg);
			backdrop-filter: blur(12px);
			border: 1px solid var(--border-color);
			border-radius: 24px;
			padding: var(--card-padding);
			box-sizing: border-box;
			display: flex;
			flex-direction: column;
			justify-content: space-between; /* push footer to bottom */
			position: relative;
			overflow: hidden;
		}
		.container::before {
			content: '';
			position: absolute;
			top: 0; left: 0; right: 0; height: 2px;
			background: linear-gradient(90deg, transparent, var(--accent), transparent);
		}
		.meta {
			font-size: var(--title-font-size);
			font-weight: 600;
			text-transform: uppercase;
			letter-spacing: 2px;
			color: var(--accent);
			margin-bottom: 12px;
		}
		.content {
			font-size: 32px;
			font-weight: 800;
			line-height: 1.2;
			margin-bottom: 0; /* let flex handle spacing */
			/* Strict Boundaries Constraint */
			overflow: hidden;
			display: -webkit-box;
			-webkit-line-clamp: 5; /* Max 5 lines of text */
			-webkit-box-orient: vertical;
			text-overflow: ellipsis;
		}
		.footer {
			display: flex;
			align-items: center;
			gap: 12px;
			font-size: 18px;
			color: var(--text-muted);
			border-top: 1px solid rgba(255, 255, 255, 0.1);
			padding-top: 16px;
			margin-top: 12px;
		}
		.dot { width: 4px; height: 4px; background: var(--accent); border-radius: 50%; }
		h1, h2, h3, p { margin: 0; }
		strong { color: var(--accent); }
	</style>
</head>
<body>
	<div class="container">
		<div class="meta">{{.Title}}</div>
		<div class="content">{{.Content}}</div>
		<div class="footer">
			<span>{{.Author}}</span>
			<div class="dot"></div>
			<span>Drashtika SocialForge</span>
		</div>
	</div>
</body>
</html>
`
