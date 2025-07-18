package security

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
	BlockDuration     time.Duration
	WhitelistedIPs    []string
	BlacklistedIPs    []string
}

// RateLimitEntry represents a rate limit entry for an IP
type RateLimitEntry struct {
	Count        int
	ResetTime    time.Time
	BlockedUntil time.Time
}

// RateLimiter implements IP-based rate limiting
type RateLimiter struct {
	config  RateLimitConfig
	entries map[string]*RateLimitEntry
	mutex   sync.RWMutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(config RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		config:  config,
		entries: make(map[string]*RateLimitEntry),
	}

	// Start cleanup goroutine
	go rl.cleanupExpired()

	return rl
}

// IsAllowed checks if a request from the given IP is allowed
func (rl *RateLimiter) IsAllowed(ip string) bool {
	// Check if IP is whitelisted
	if rl.isWhitelisted(ip) {
		return true
	}

	// Check if IP is blacklisted
	if rl.isBlacklisted(ip) {
		return false
	}

	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	entry, exists := rl.entries[ip]

	if !exists {
		rl.entries[ip] = &RateLimitEntry{
			Count:     1,
			ResetTime: now.Add(time.Minute),
		}
		return true
	}

	// Check if IP is currently blocked
	if now.Before(entry.BlockedUntil) {
		return false
	}

	// Reset counter if time window has passed
	if now.After(entry.ResetTime) {
		entry.Count = 1
		entry.ResetTime = now.Add(time.Minute)
		entry.BlockedUntil = time.Time{} // Clear block
		return true
	}

	// Increment counter
	entry.Count++

	// Check if limit exceeded
	if entry.Count > rl.config.RequestsPerMinute {
		entry.BlockedUntil = now.Add(rl.config.BlockDuration)
		return false
	}

	return true
}

// GetClientIP extracts the real client IP from the request
func GetClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		// Take the first IP if there are multiple
		if idx := strings.Index(forwarded, ","); idx != -1 {
			return strings.TrimSpace(forwarded[:idx])
		}
		return strings.TrimSpace(forwarded)
	}

	// Check X-Real-IP header
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	// Check CF-Connecting-IP (Cloudflare)
	if cfIP := r.Header.Get("CF-Connecting-IP"); cfIP != "" {
		return cfIP
	}

	// Fall back to remote address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return ip
}

// IsValidIP checks if the provided string is a valid IP address
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// isWhitelisted checks if an IP is in the whitelist
func (rl *RateLimiter) isWhitelisted(ip string) bool {
	for _, whitelistedIP := range rl.config.WhitelistedIPs {
		if ip == whitelistedIP {
			return true
		}

		// Check if it's a CIDR range
		if strings.Contains(whitelistedIP, "/") {
			if _, cidr, err := net.ParseCIDR(whitelistedIP); err == nil {
				if clientIP := net.ParseIP(ip); clientIP != nil {
					if cidr.Contains(clientIP) {
						return true
					}
				}
			}
		}
	}
	return false
}

// isBlacklisted checks if an IP is in the blacklist
func (rl *RateLimiter) isBlacklisted(ip string) bool {
	for _, blacklistedIP := range rl.config.BlacklistedIPs {
		if ip == blacklistedIP {
			return true
		}

		// Check if it's a CIDR range
		if strings.Contains(blacklistedIP, "/") {
			if _, cidr, err := net.ParseCIDR(blacklistedIP); err == nil {
				if clientIP := net.ParseIP(ip); clientIP != nil {
					if cidr.Contains(clientIP) {
						return true
					}
				}
			}
		}
	}
	return false
}

// cleanupExpired removes expired entries from the rate limiter
func (rl *RateLimiter) cleanupExpired() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mutex.Lock()
		now := time.Now()

		for ip, entry := range rl.entries {
			// Remove entries that are past their reset time and not blocked
			if now.After(entry.ResetTime) && now.After(entry.BlockedUntil) {
				delete(rl.entries, ip)
			}
		}

		rl.mutex.Unlock()
	}
}
