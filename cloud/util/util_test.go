/*
Copyright (c) 2021, 2022 Oracle and/or its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

import (
	"context"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
	infrastructurev1beta2 "github.com/oracle/cluster-api-provider-oci/api/v1beta2"
	"github.com/oracle/cluster-api-provider-oci/cloud/config"
	"github.com/oracle/cluster-api-provider-oci/cloud/scope"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGetClusterIdentityFromRef(t *testing.T) {
	testCases := []struct {
		name          string
		namespace     string
		ref           *corev1.ObjectReference
		objects       []client.Object
		errorExpected bool
		expectedSpec  infrastructurev1beta2.OCIClusterIdentitySpec
	}{
		{
			name:      "simple",
			namespace: "default",
			ref: &corev1.ObjectReference{
				Kind:       "OCIClusterIdentity",
				Namespace:  "default",
				Name:       "test-identity",
				APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
			},
			objects: []client.Object{&infrastructurev1beta2.OCIClusterIdentity{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-identity",
					Namespace: "default",
				},
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			}},
			expectedSpec: infrastructurev1beta2.OCIClusterIdentitySpec{
				Type: infrastructurev1beta2.UserPrincipal,
				PrincipalSecret: corev1.SecretReference{
					Name:      "test",
					Namespace: "test",
				},
			},
		},
		{
			name:      "error - not found",
			namespace: "default",
			ref: &corev1.ObjectReference{
				Kind:       "OCIClusterIdentity",
				Namespace:  "default",
				Name:       "test-identity",
				APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
			},
			objects:       []client.Object{},
			errorExpected: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			client := fake.NewClientBuilder().WithObjects(tt.objects...).Build()
			result, err := GetClusterIdentityFromRef(context.Background(), client, tt.namespace, tt.ref)
			if tt.errorExpected {
				g.Expect(err).To(Not(BeNil()))
			} else {
				g.Expect(err).To(BeNil())
				if !reflect.DeepEqual(tt.expectedSpec, result.Spec) {
					t.Errorf("Test (%s) \n Expected %v, \n Actual %v", tt.name, tt.expectedSpec, result.Spec)
				}
			}
		})
	}
}

func TestGetOrBuildClientFromIdentity(t *testing.T) {
	testCases := []struct {
		name            string
		namespace       string
		clusterIdentity *infrastructurev1beta2.OCIClusterIdentity
		objects         []client.Object
		errorExpected   bool
		defaultRegion   string
	}{
		{
			name:      "error - secret not found",
			namespace: "default",
			clusterIdentity: &infrastructurev1beta2.OCIClusterIdentity{
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			objects:       []client.Object{},
			errorExpected: true,
		},
		{
			name:      "error - invalid principal type",
			namespace: "default",
			clusterIdentity: &infrastructurev1beta2.OCIClusterIdentity{
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: "invalid",
				},
			},
			objects:       []client.Object{},
			errorExpected: true,
		},
		{
			name:      "secret found",
			namespace: "default",
			clusterIdentity: &infrastructurev1beta2.OCIClusterIdentity{
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			objects: []client.Object{&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Data: map[string][]byte{config.Tenancy: []byte("tenancy"), config.User: []byte("user"),
					config.Key: []byte("key"), config.Fingerprint: []byte("fingerprint"), config.Region: []byte("region")},
			}},
			errorExpected: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			client := fake.NewClientBuilder().WithObjects(tt.objects...).Build()
			_, err := GetOrBuildClientFromIdentity(context.Background(), client, tt.clusterIdentity, tt.defaultRegion, nil, tt.namespace)
			if tt.errorExpected {
				g.Expect(err).To(Not(BeNil()))
			} else {
				g.Expect(err).To(BeNil())
			}
		})
	}
}

func TestIsClusterNamespaceAllowed(t *testing.T) {
	testCases := []struct {
		name              string
		namespace         string
		allowedNamespaces *infrastructurev1beta2.AllowedNamespaces
		objects           []client.Object
		expected          bool
	}{
		{
			name:      "nil allowednamespace, not allowed",
			namespace: "default",
			objects:   []client.Object{},
			expected:  false,
		},
		{
			name:              "empty allowednamespace, allowed",
			namespace:         "default",
			allowedNamespaces: &infrastructurev1beta2.AllowedNamespaces{},
			objects:           []client.Object{},
			expected:          true,
		},
		{
			name:      "not allowed",
			namespace: "test",
			allowedNamespaces: &infrastructurev1beta2.AllowedNamespaces{
				NamespaceList: []string{"test123"},
			},
			objects:  []client.Object{},
			expected: false,
		},
		{
			name:      "allowed",
			namespace: "test",
			allowedNamespaces: &infrastructurev1beta2.AllowedNamespaces{
				NamespaceList: []string{"test"},
			},
			objects:  []client.Object{},
			expected: true,
		},
		{
			name:      "empty label selector",
			namespace: "test",
			allowedNamespaces: &infrastructurev1beta2.AllowedNamespaces{
				Selector: &metav1.LabelSelector{},
			},
			objects:  []client.Object{},
			expected: false,
		},
		{
			name:      "allowed label selector",
			namespace: "test",
			allowedNamespaces: &infrastructurev1beta2.AllowedNamespaces{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"key": "value"},
				},
			},
			objects: []client.Object{&corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels:    map[string]string{"key": "value"},
				},
			}},
			expected: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			client := fake.NewClientBuilder().WithObjects(tt.objects...).Build()
			result := IsClusterNamespaceAllowed(context.Background(), client, tt.allowedNamespaces, tt.namespace)
			g.Expect(result).To(BeEquivalentTo(tt.expected))
		})
	}
}

func TestCreateClientProviderFromClusterIdentity(t *testing.T) {
	testCases := []struct {
		name            string
		namespace       string
		objects         []client.Object
		clusterAccessor scope.OCIClusterAccessor
		ref             *corev1.ObjectReference
		errorExpected   bool
		defaultRegion   string
	}{
		{
			name:      "error - secret not found",
			namespace: "default",
			clusterAccessor: scope.OCISelfManagedCluster{
				OCICluster: &infrastructurev1beta2.OCICluster{},
			},
			ref: &corev1.ObjectReference{
				Kind:       "OCIClusterIdentity",
				Namespace:  "default",
				Name:       "test-identity",
				APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
			},
			objects: []client.Object{&infrastructurev1beta2.OCIClusterIdentity{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-identity",
					Namespace: "default",
				},
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			}},
			errorExpected: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			client := fake.NewClientBuilder().WithObjects(tt.objects...).Build()
			_, err := CreateClientProviderFromClusterIdentity(context.Background(), client, tt.namespace, tt.defaultRegion, tt.clusterAccessor, tt.ref)
			if tt.errorExpected {
				g.Expect(err).To(Not(BeNil()))
			} else {
				g.Expect(err).To(BeNil())
			}
		})
	}
}

func TestInitClientsAndRegion(t *testing.T) {
	clientProvider, err := scope.MockNewClientProvider(scope.MockOCIClients{})
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		name            string
		namespace       string
		objects         []client.Object
		clusterAccessor scope.OCIClusterAccessor
		clientProvider  *scope.ClientProvider
		errorExpected   bool
		errorMessage    string
		defaultRegion   string
		secret          *corev1.Secret
	}{
		{
			name:          "good - secret found",
			namespace:     "default",
			defaultRegion: "ashburn-1",
			clusterAccessor: scope.OCISelfManagedCluster{
				OCICluster: &infrastructurev1beta2.OCICluster{
					Spec: infrastructurev1beta2.OCIClusterSpec{
						ClientOverrides: &infrastructurev1beta2.ClientOverrides{
							CertOverride: &corev1.SecretReference{
								Name:      "certSecret",
								Namespace: "default",
							},
						},
					},
				},
			},
			objects: []client.Object{&infrastructurev1beta2.OCIClusterIdentity{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-identity",
					Namespace: "default",
				},
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			}, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "certSecret",
					Namespace: "default",
				},
				Type: corev1.SecretTypeBootstrapToken,
				Data: map[string][]byte{
					"cert": []byte("TestPemCert"),
				},
			}},
			clientProvider: clientProvider,
			errorExpected:  false,
		},
		{
			name:          "bad - secret not found - wrong namespace",
			namespace:     "default",
			defaultRegion: "ashburn-1",
			clusterAccessor: scope.OCISelfManagedCluster{
				OCICluster: &infrastructurev1beta2.OCICluster{
					Spec: infrastructurev1beta2.OCIClusterSpec{
						ClientOverrides: &infrastructurev1beta2.ClientOverrides{
							CertOverride: &corev1.SecretReference{
								Name:      "certSecret",
								Namespace: "default",
							},
						},
					},
				},
			},
			objects: []client.Object{&infrastructurev1beta2.OCIClusterIdentity{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-identity",
					Namespace: "default",
				},
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			}, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "certSecret",
					Namespace: "badNamespace",
				},
				Type: corev1.SecretTypeBootstrapToken,
				Data: map[string][]byte{
					"cert": []byte("TestPemCert"),
				},
			}},
			clientProvider: clientProvider,
			errorExpected:  true,
			errorMessage:   "Unable to fetch CertOverrideSecret: secrets \"certSecret\" not found",
		},
		{
			name:          "bad - secret not found - missing cert data",
			namespace:     "default",
			defaultRegion: "ashburn-1",
			clusterAccessor: scope.OCISelfManagedCluster{
				OCICluster: &infrastructurev1beta2.OCICluster{
					Spec: infrastructurev1beta2.OCIClusterSpec{
						ClientOverrides: &infrastructurev1beta2.ClientOverrides{
							CertOverride: &corev1.SecretReference{
								Name:      "certSecret",
								Namespace: "default",
							},
						},
					},
				},
			},
			objects: []client.Object{&infrastructurev1beta2.OCIClusterIdentity{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-identity",
					Namespace: "default",
				},
				Spec: infrastructurev1beta2.OCIClusterIdentitySpec{
					Type: infrastructurev1beta2.UserPrincipal,
					PrincipalSecret: corev1.SecretReference{
						Name:      "test",
						Namespace: "test",
					},
				},
			}, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "certSecret",
					Namespace: "default",
				},
				Type: corev1.SecretTypeBootstrapToken,
			}},
			clientProvider: clientProvider,
			errorExpected:  true,
			errorMessage:   "Cert Secret didn't contain 'cert' data",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			client := fake.NewClientBuilder().WithObjects(tt.objects...).Build()
			_, _, _, err := InitClientsAndRegion(context.Background(), client, tt.defaultRegion, tt.clusterAccessor, tt.clientProvider)
			if tt.errorExpected {
				g.Expect(err).To(Not(BeNil()))
				g.Expect(err.Error()).To(Equal(tt.errorMessage))
			} else {
				g.Expect(err).To(BeNil())
			}
		})
	}
}
