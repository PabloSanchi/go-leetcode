package main

import (
	"reflect"
	"testing"
)

func TestServerConfigBuilder_DefaultConfig(t *testing.T) {
	server := NewServer()
	if server.config.Protocol != "tcp" {
		t.Errorf("expected protocol to be tcp, got %s", server.config.Protocol)
	}

	if server.config.Host != "localhost" {
		t.Errorf("expected host to be localhost, got %s", server.config.Host)
	}

	if server.config.Port != 8080 {
		t.Errorf("expected port to be 8080, got %d", server.config.Port)
	}

	if server.config.Address() != "localhost:8080" {
		t.Errorf("expected address to be localhost:8080, got %s", server.config.Address())
	}

	if len(server.config.Origins) != 0 {
		t.Errorf("expected origins to be empty, got %v", server.config.Origins)
	}
}

func TestServerConfigBuilder_CustomConfig(t *testing.T) {
	expectedOrigins := []string{"http://localhost:3000", "http://localhost:8080"}

	server := NewServer(
		WithProtocol("udp"),
		WithHost("0.0.0.0"),
		WithPort(9090),
		WithOrigins([]string{"http://localhost:3000", "http://localhost:8080"}),
	)

	if server.config.Protocol != "udp" {
		t.Errorf("expected protocol to be udp, got %s", server.config.Protocol)
	}

	if server.config.Host != "0.0.0.0" {
		t.Errorf("expected host to be 0.0.0.0, got %s", server.config.Host)
	}

	if server.config.Port != 9090 {
		t.Errorf("expected port to be 9090, got %d", server.config.Port)
	}

	if server.config.Address() != "0.0.0.0:9090" {
		t.Errorf("expected address to be localhost:9090, got %s", server.config.Address())
	}

	if len(server.config.Origins) != 2 {
		t.Errorf("expected origins to have 2 elements, got %v", server.config.Origins)
	}

	if !reflect.DeepEqual(server.config.Origins, expectedOrigins) {
		t.Errorf("expected origins to be %v, got %v", expectedOrigins, server.config.Origins)
	}

}

func TestServerConfigBuilder_WithPartialConfig(t *testing.T) {
	server := NewServer(
		WithPort(9090),
	)

	if server.config.Host != "localhost" {
		t.Errorf("expected host to be localhost, got %s", server.config.Host)
	}

	if server.config.Port != 9090 {
		t.Errorf("expected port to be 9090, got %d", server.config.Port)
	}

	if server.config.Address() != "localhost:9090" {
		t.Errorf("expected address to be localhost:9090, got %s", server.config.Address())
	}

	if len(server.config.Origins) != 0 {
		t.Errorf("expected origins to be empty, got %v", server.config.Origins)
	}
}

func TestServerConfigBuilder_WithSingleOrigin(t *testing.T) {
	server := NewServer(
		WithOrigin("http://localhost:3000"),
	)

	if len(server.config.Origins) != 1 {
		t.Errorf("expected origins to have 1 element, got %v", server.config.Origins)
	}

	if server.config.Origins[0] != "http://localhost:3000" {
		t.Errorf("expected origins to be http://localhost:3000, got %v", server.config.Origins)
	}

}

func TestServerConfigBuilder_WithMultipleOrigins(t *testing.T) {
	server := NewServer(
		WithOrigin("http://localhost:3000"),
		WithOrigin("http://localhost:8080"),
	)

	if len(server.config.Origins) != 2 {
		t.Errorf("expected origins to have 2 elements, got %v", server.config.Origins)
	}

	if server.config.Origins[0] != "http://localhost:3000" {
		t.Errorf("expected origins to be http://localhost:3000, got %v", server.config.Origins)
	}

	if server.config.Origins[1] != "http://localhost:8080" {
		t.Errorf("expected origins to be http://localhost:8080, got %v", server.config.Origins)
	}
}
