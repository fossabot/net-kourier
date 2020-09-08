/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ingress

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/sets"
	"knative.dev/networking/pkg/apis/networking/v1alpha1"
	"knative.dev/networking/test"
)

// TestWebsocket verifies that websockets may be used via a simple Ingress.
func TestWebsocket(t *testing.T) {
	t.Parallel()
	clients := test.Setup(t)

	const suffix = "- pong"
	name, port, _ := CreateWebsocketService(t, clients, suffix)

	domain := name + ".example.com"

	// Create a simple Ingress over the Service.
	_, dialCtx, _ := CreateIngressReadyDialContext(t, clients, v1alpha1.IngressSpec{
		Rules: []v1alpha1.IngressRule{{
			Hosts:      []string{domain},
			Visibility: v1alpha1.IngressVisibilityExternalIP,
			HTTP: &v1alpha1.HTTPIngressRuleValue{
				Paths: []v1alpha1.HTTPIngressPath{{
					Splits: []v1alpha1.IngressBackendSplit{{
						IngressBackend: v1alpha1.IngressBackend{
							ServiceName:      name,
							ServiceNamespace: test.ServingNamespace,
							ServicePort:      intstr.FromInt(port),
						},
					}},
				}},
			},
		}},
	})

	dialer := websocket.Dialer{
		NetDialContext:   dialCtx,
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	u := url.URL{Scheme: "ws", Host: domain, Path: "/"}
	conn, _, err := dialer.Dial(u.String(), http.Header{"Host": {domain}})
	if err != nil {
		t.Fatal("Dial() =", err)
	}
	defer conn.Close()

	for i := 0; i < 100; i++ {
		checkWebsocketRoundTrip(t, conn, suffix)
	}
}

// TestWebsocketSplit verifies that websockets may be used across a traffic split.
func TestWebsocketSplit(t *testing.T) {
	t.Parallel()
	clients := test.Setup(t)

	const suffixBlue = "- blue"
	blueName, bluePort, _ := CreateWebsocketService(t, clients, suffixBlue)

	const suffixGreen = "- green"
	greenName, greenPort, _ := CreateWebsocketService(t, clients, suffixGreen)

	// The suffixes we expect to see.
	want := sets.NewString(suffixBlue, suffixGreen)

	// Create a simple Ingress over the Service.
	name := test.ObjectNameForTest(t)
	domain := name + ".example.com"
	_, dialCtx, _ := CreateIngressReadyDialContext(t, clients, v1alpha1.IngressSpec{
		Rules: []v1alpha1.IngressRule{{
			Hosts:      []string{domain},
			Visibility: v1alpha1.IngressVisibilityExternalIP,
			HTTP: &v1alpha1.HTTPIngressRuleValue{
				Paths: []v1alpha1.HTTPIngressPath{{
					Splits: []v1alpha1.IngressBackendSplit{{
						IngressBackend: v1alpha1.IngressBackend{
							ServiceName:      blueName,
							ServiceNamespace: test.ServingNamespace,
							ServicePort:      intstr.FromInt(bluePort),
						},
						Percent: 50,
					}, {
						IngressBackend: v1alpha1.IngressBackend{
							ServiceName:      greenName,
							ServiceNamespace: test.ServingNamespace,
							ServicePort:      intstr.FromInt(greenPort),
						},
						Percent: 50,
					}},
				}},
			},
		}},
	})

	dialer := websocket.Dialer{
		NetDialContext:   dialCtx,
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}
	u := url.URL{Scheme: "ws", Host: domain, Path: "/"}

	const maxRequests = 100
	got := sets.NewString()
	for i := 0; i < maxRequests; i++ {
		conn, _, err := dialer.Dial(u.String(), http.Header{"Host": {domain}})
		if err != nil {
			t.Fatal("Dial() =", err)
		}
		defer conn.Close()

		suffix := findWebsocketSuffix(t, conn)
		if suffix == "" {
			continue
		}
		got.Insert(suffix)

		for j := 0; j < 10; j++ {
			checkWebsocketRoundTrip(t, conn, suffix)
		}

		if want.Equal(got) {
			// Short circuit if we've seen all splits.
			return
		}
	}

	// Us getting here means we haven't seen splits.
	t.Errorf("(over %d requests) (-want, +got) = %s", maxRequests, cmp.Diff(want, got))
}

func findWebsocketSuffix(t *testing.T, conn *websocket.Conn) string {
	// Establish the suffix that corresponds to this socket.
	message := fmt.Sprintf("ping - %d", rand.Intn(1000))
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Errorf("WriteMessage() = %v", err)
		return ""
	}

	_, recv, err := conn.ReadMessage()
	if err != nil {
		t.Errorf("ReadMessage() = %v", err)
		return ""
	}
	gotMsg := string(recv)
	if !strings.HasPrefix(gotMsg, message) {
		t.Errorf("ReadMessage() = %s, wanted %s prefix", gotMsg, message)
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(gotMsg, message))
}

func checkWebsocketRoundTrip(t *testing.T, conn *websocket.Conn, suffix string) {
	message := fmt.Sprintf("ping - %d", rand.Intn(1000))
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Errorf("WriteMessage() = %v", err)
		return
	}

	// Read back the echoed message and compared with sent.
	if _, recv, err := conn.ReadMessage(); err != nil {
		t.Errorf("ReadMessage() = %v", err)
	} else if got, want := string(recv), message+" "+suffix; got != want {
		t.Errorf("ReadMessage() = %s, wanted %s", got, want)
	}
}
