// SPDX-License-Identifier: Apache-2.0

package controllerapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/metio-plattform/metio/internal/db"
)

// ServerResponse mirrors the controller API representation of a server.
type ServerResponse struct {
	ID     string `json:"id"`
	Config struct {
		Name             string               `json:"name"`
		Region           string               `json:"region"`
		Zone             string               `json:"zone"`
		MachineType      string               `json:"machineType"`
		MinecraftVersion string               `json:"minecraftVersion"`
		DiskSizeGB       int                  `json:"diskSizeGB"`
		Modpack          *db.ModpackConfig    `json:"modpack,omitempty"`
		Backup           *db.BackupConfig     `json:"backup,omitempty"`
		ShutdownSchedule *db.ShutdownSchedule `json:"shutdownSchedule,omitempty"`
		Runtime          string               `json:"runtime,omitempty"`
	} `json:"config"`
	Status *db.Status `json:"status,omitempty"`
}

// Client talks to the Metio controller API.
type Client interface {
	ListServers(ctx context.Context) ([]ServerResponse, error)
	GetServer(ctx context.Context, id string) (ServerResponse, error)
	UpdateServerStatus(ctx context.Context, id string, status db.Status) error
}

type client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func New(baseURL, apiKey string) Client {
	return &client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *client) do(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reader)
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return c.httpClient.Do(req)
}

func (c *client) ListServers(ctx context.Context) ([]ServerResponse, error) {
	resp, err := c.do(ctx, http.MethodGet, "/api/servers", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("list servers: %s: %s", resp.Status, string(b))
	}
	var servers []ServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&servers); err != nil {
		return nil, err
	}
	return servers, nil
}

func (c *client) GetServer(ctx context.Context, id string) (ServerResponse, error) {
	resp, err := c.do(ctx, http.MethodGet, "/api/servers/"+url.PathEscape(id), nil)
	if err != nil {
		return ServerResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return ServerResponse{}, fmt.Errorf("get server: %s: %s", resp.Status, string(b))
	}
	var s ServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return ServerResponse{}, err
	}
	return s, nil
}

func (c *client) UpdateServerStatus(ctx context.Context, id string, status db.Status) error {
	resp, err := c.do(ctx, http.MethodPut, "/api/servers/"+url.PathEscape(id)+"/status", status)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("update status: %s: %s", resp.Status, string(b))
	}
	return nil
}
