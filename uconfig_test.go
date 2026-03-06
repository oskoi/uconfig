package uconfig_test

import (
	"errors"
	"testing"

	"github.com/oskoi/uconfig"
	"github.com/oskoi/uconfig/flat"
	"github.com/oskoi/uconfig/internal/f"
	"github.com/oskoi/uconfig/plugins"
)

type BadPlugin interface {
	plugins.Plugin

	NotWalkerOrVisitor()
}

func TestBadPlug(t *testing.T) {
	var badPlugin BadPlugin

	conf := uconfig.New[f.Config](badPlugin)
	_, err := conf.Parse()

	if err == nil {
		t.Error("expected error for bad plugin, got nil")
	}

	if err.Error() != "unsupported plugins. expecting a walker or visitor" {
		t.Errorf("Expected unsupported plugin error, got: %v", err)
	}
}

type FailingPluginWalker struct {
	plugins.Plugin
}

func (fp FailingPluginWalker) Walk(any) error {
	return errors.New("failed to walk")
}

func TestFailingPlugWalker(t *testing.T) {
	var failingPluginWalker FailingPluginWalker

	conf := uconfig.New[f.Config](failingPluginWalker)
	_, err := conf.Parse()

	if err == nil {
		t.Error("expected error for bad plugin, got nil")
	}

	if err.Error() != "failed to walk" {
		t.Errorf("Expected failed to walk, got: %v", err)
	}
}

type FailingPluginVisitor struct {
	plugins.Plugin
}

func (fp FailingPluginVisitor) Visit(flat.Fields) error {
	return errors.New("failed to visit")
}

func TestFailingPlugVisitor(t *testing.T) {
	var failingPluginVisitor FailingPluginVisitor

	conf := uconfig.New[f.Config](failingPluginVisitor)
	_, err := conf.Parse()

	if err == nil {
		t.Error("expected error for bad plugin, got nil")
	}

	if err.Error() != "failed to visit" {
		t.Errorf("Expected failed to visit, got: %v", err)
	}
}
