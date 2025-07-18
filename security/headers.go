package security

import (
	"net/http"
	"strconv"
	"time"
)

// SecurityHeaders contains security header configurations
type SecurityHeaders struct {
	HSSTMaxAge                int
	ContentTypeNoSniff        bool
	FrameDeny                 bool
	XSSProtection             bool
	ReferrerPolicy            string
	ContentSecurityPolicy     string
	PermissionsPolicy         string
	CrossOriginEmbedderPolicy string
	CrossOriginOpenerPolicy   string
	CrossOriginResourcePolicy string
}

// DefaultSecurityHeaders returns default security headers configuration
func DefaultSecurityHeaders() SecurityHeaders {
	return SecurityHeaders{
		HSSTMaxAge:                31536000, // 1 year
		ContentTypeNoSniff:        true,
		FrameDeny:                 true,
		XSSProtection:             true,
		ReferrerPolicy:            "strict-origin-when-cross-origin",
		ContentSecurityPolicy:     "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; child-src 'none'; worker-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'",
		PermissionsPolicy:         "camera=(), microphone=(), location=(), payment=(), usb=(), magnetometer=(), gyroscope=(), fullscreen=(self)",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "same-origin",
	}
}

// ApplySecurityHeaders applies security headers to the response
func ApplySecurityHeaders(w http.ResponseWriter, headers SecurityHeaders) {
	// HSTS (HTTP Strict Transport Security)
	if headers.HSSTMaxAge > 0 {
		w.Header().Set("Strict-Transport-Security", "max-age="+strconv.Itoa(headers.HSSTMaxAge)+"; includeSubDomains; preload")
	}

	// X-Content-Type-Options
	if headers.ContentTypeNoSniff {
		w.Header().Set("X-Content-Type-Options", "nosniff")
	}

	// X-Frame-Options
	if headers.FrameDeny {
		w.Header().Set("X-Frame-Options", "DENY")
	}

	// X-XSS-Protection
	if headers.XSSProtection {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
	}

	// Referrer-Policy
	if headers.ReferrerPolicy != "" {
		w.Header().Set("Referrer-Policy", headers.ReferrerPolicy)
	}

	// Content-Security-Policy
	if headers.ContentSecurityPolicy != "" {
		w.Header().Set("Content-Security-Policy", headers.ContentSecurityPolicy)
	}

	// Permissions-Policy
	if headers.PermissionsPolicy != "" {
		w.Header().Set("Permissions-Policy", headers.PermissionsPolicy)
	}

	// Cross-Origin-Embedder-Policy
	if headers.CrossOriginEmbedderPolicy != "" {
		w.Header().Set("Cross-Origin-Embedder-Policy", headers.CrossOriginEmbedderPolicy)
	}

	// Cross-Origin-Opener-Policy
	if headers.CrossOriginOpenerPolicy != "" {
		w.Header().Set("Cross-Origin-Opener-Policy", headers.CrossOriginOpenerPolicy)
	}

	// Cross-Origin-Resource-Policy
	if headers.CrossOriginResourcePolicy != "" {
		w.Header().Set("Cross-Origin-Resource-Policy", headers.CrossOriginResourcePolicy)
	}

	// Remove potentially dangerous headers
	w.Header().Del("Server")
	w.Header().Del("X-Powered-By")
}

// SetNoCacheHeaders sets headers to prevent caching
func SetNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// SetCacheHeaders sets headers for cacheable responses
func SetCacheHeaders(w http.ResponseWriter, maxAge time.Duration) {
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(int(maxAge.Seconds())))
	w.Header().Set("Expires", time.Now().Add(maxAge).Format(http.TimeFormat))
}
