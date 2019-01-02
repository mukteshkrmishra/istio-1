// Code generated by go-bindata.
// sources:
// dataset/config.istio.io/v1alpha2/circonus.yaml
// dataset/config.istio.io/v1alpha2/circonus_expected.json
// dataset/extensions/v1beta1/ingress_basic.yaml
// dataset/extensions/v1beta1/ingress_basic_expected.json
// dataset/extensions/v1beta1/ingress_basic_meshconfig.yaml
// dataset/extensions/v1beta1/ingress_merge_0.skip
// dataset/extensions/v1beta1/ingress_merge_0.yaml
// dataset/extensions/v1beta1/ingress_merge_0_expected.json
// dataset/extensions/v1beta1/ingress_merge_0_meshconfig.yaml
// dataset/extensions/v1beta1/ingress_merge_1.yaml
// dataset/extensions/v1beta1/ingress_merge_1_expected.json
// dataset/networking.istio.io/v1alpha3/destinationRule.yaml
// dataset/networking.istio.io/v1alpha3/destinationRule_expected.json
// dataset/networking.istio.io/v1alpha3/gateway.yaml
// dataset/networking.istio.io/v1alpha3/gateway_expected.json
// DO NOT EDIT!

package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _datasetConfigIstioIoV1alpha2CirconusYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: circonus
metadata:
  name: valid-circonus
spec:
  submission_url: "https://trap.noit.circonus.net/module/httptrap/myuuid/mysecret"
  submission_interval: "10s"
`)

func datasetConfigIstioIoV1alpha2CirconusYamlBytes() ([]byte, error) {
	return _datasetConfigIstioIoV1alpha2CirconusYaml, nil
}

func datasetConfigIstioIoV1alpha2CirconusYaml() (*asset, error) {
	bytes, err := datasetConfigIstioIoV1alpha2CirconusYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config.istio.io/v1alpha2/circonus.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigIstioIoV1alpha2Circonus_expectedJson = []byte(`{
  "type.googleapis.com/istio.mcp.v1alpha1.extensions.LegacyMixerResource": [
    {
      "Metadata": {
        "name": "circonus/valid-circonus"
      },
      "Resource": {
        "contents": {
          "fields": {
            "submission_interval": {
              "Kind": {
                "StringValue": "10s"
              }
            },
            "submission_url": {
              "Kind": {
                "StringValue": "https://trap.noit.circonus.net/module/httptrap/myuuid/mysecret"
              }
            }
          }
        },
        "kind": "circonus",
        "name": "valid-circonus"
      },
      "TypeURL": "type.googleapis.com/istio.mcp.v1alpha1.extensions.LegacyMixerResource"
    }
  ]
}
`)

func datasetConfigIstioIoV1alpha2Circonus_expectedJsonBytes() ([]byte, error) {
	return _datasetConfigIstioIoV1alpha2Circonus_expectedJson, nil
}

func datasetConfigIstioIoV1alpha2Circonus_expectedJson() (*asset, error) {
	bytes, err := datasetConfigIstioIoV1alpha2Circonus_expectedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config.istio.io/v1alpha2/circonus_expected.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_basicYaml = []byte(`apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
  annotations:
    kubernetes.io/ingress.class: "cls"

spec:
  backend:
    serviceName: "testsvc"
    servicePort: "80"
`)

func datasetExtensionsV1beta1Ingress_basicYamlBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_basicYaml, nil
}

func datasetExtensionsV1beta1Ingress_basicYaml() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_basicYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_basic.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_basic_expectedJson = []byte(`{
  "type.googleapis.com/istio.networking.v1alpha3.Gateway": [
    {
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.Gateway",
      "Metadata": {
        "name": "istio-system/foo-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-foo-",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      }
    }
  ],

  "type.googleapis.com/istio.networking.v1alpha3.VirtualService": [
  ]
}
`)

func datasetExtensionsV1beta1Ingress_basic_expectedJsonBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_basic_expectedJson, nil
}

func datasetExtensionsV1beta1Ingress_basic_expectedJson() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_basic_expectedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_basic_expected.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_basic_meshconfigYaml = []byte(`ingressClass: cls
ingressControllerMode: STRICT
`)

func datasetExtensionsV1beta1Ingress_basic_meshconfigYamlBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_basic_meshconfigYaml, nil
}

func datasetExtensionsV1beta1Ingress_basic_meshconfigYaml() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_basic_meshconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_basic_meshconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_merge_0Skip = []byte(``)

func datasetExtensionsV1beta1Ingress_merge_0SkipBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_merge_0Skip, nil
}

func datasetExtensionsV1beta1Ingress_merge_0Skip() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_merge_0SkipBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_merge_0.skip", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_merge_0Yaml = []byte(`apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
  namespace: ns
  annotations:
    kubernetes.io/ingress.class: "cls"
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - path: /foo
        backend:
          serviceName: service1
          servicePort: 4200
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: bar
  namespace: ns
  annotations:
    kubernetes.io/ingress.class: "cls"
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - path: /bar
        backend:
          serviceName: service2
          servicePort: 2400
---
`)

func datasetExtensionsV1beta1Ingress_merge_0YamlBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_merge_0Yaml, nil
}

func datasetExtensionsV1beta1Ingress_merge_0Yaml() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_merge_0YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_merge_0.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_merge_0_expectedJson = []byte(`{
  "type.googleapis.com/istio.networking.v1alpha3.Gateway": [
    {
      "Metadata": {
        "name": "istio-system/bar-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-bar-ns",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      },
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.Gateway"
    },
    {
      "Metadata": {
        "name": "istio-system/foo-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-foo-ns",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      },
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.Gateway"
    }
  ],

  "type.googleapis.com/istio.networking.v1alpha3.VirtualService": [
    {
      "Metadata": {
        "name": "istio-system/foo-bar-com-bar-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "gateways": [
          "istio-autogenerated-k8s-ingress"
        ],
        "hosts": [
          "foo.bar.com"
        ],
        "http": [
          {
            "match": [
              {
                "uri": {
                  "MatchType": {
                    "Exact": "/bar"
                  }
                }
              }
            ],
            "route": [
              {
                "destination": {
                  "host": "service2.ns.svc.cluster.local",
                  "port": {
                    "Port": {
                      "Number": 2400
                    }
                  }
                },
                "weight": 100
              }
            ]
          },
          {
            "match": [
              {
                "uri": {
                  "MatchType": {
                    "Exact": "/foo"
                  }
                }
              }
            ],
            "route": [
              {
                "destination": {
                  "host": "service1.ns.svc.cluster.local",
                  "port": {
                    "Port": {
                      "Number": 4200
                    }
                  }
                },
                "weight": 100
              }
            ]
          }
        ]
      },
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.VirtualService"
    }
  ]
}`)

func datasetExtensionsV1beta1Ingress_merge_0_expectedJsonBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_merge_0_expectedJson, nil
}

func datasetExtensionsV1beta1Ingress_merge_0_expectedJson() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_merge_0_expectedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_merge_0_expected.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_merge_0_meshconfigYaml = []byte(`ingressClass: cls
ingressControllerMode: STRICT
`)

func datasetExtensionsV1beta1Ingress_merge_0_meshconfigYamlBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_merge_0_meshconfigYaml, nil
}

func datasetExtensionsV1beta1Ingress_merge_0_meshconfigYaml() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_merge_0_meshconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_merge_0_meshconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_merge_1Yaml = []byte(`apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
  namespace: ns
  annotations:
    kubernetes.io/ingress.class: "cls"
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - path: /foo
        backend:
          serviceName: service1
          servicePort: 4200
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: bar
  namespace: ns
  annotations:
    kubernetes.io/ingress.class: "cls"
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - path: /bar
        backend:
          # The service has changed since the initial config.
          serviceName: service5
          servicePort: 5000
---
`)

func datasetExtensionsV1beta1Ingress_merge_1YamlBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_merge_1Yaml, nil
}

func datasetExtensionsV1beta1Ingress_merge_1Yaml() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_merge_1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_merge_1.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetExtensionsV1beta1Ingress_merge_1_expectedJson = []byte(`{
  "type.googleapis.com/istio.networking.v1alpha3.Gateway": [
    {
      "Metadata": {
        "name": "istio-system/bar-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-bar-ns",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      },
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.Gateway"
    },
    {
      "Metadata": {
        "name": "istio-system/foo-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-foo-ns",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      },
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.Gateway"
    }
  ],

  "type.googleapis.com/istio.networking.v1alpha3.VirtualService": [
    {
      "Metadata": {
        "name": "istio-system/foo-bar-com-bar-istio-autogenerated-k8s-ingress"
      },
      "Resource": {
        "gateways": [
          "istio-autogenerated-k8s-ingress"
        ],
        "hosts": [
          "foo.bar.com"
        ],
        "http": [
          {
            "match": [
              {
                "uri": {
                  "MatchType": {
                    "Exact": "/bar"
                  }
                }
              }
            ],
            "route": [
              {
                "destination": {
                  "host": "service5.ns.svc.cluster.local",
                  "port": {
                    "Port": {
                      "Number": 5000
                    }
                  }
                },
                "weight": 100
              }
            ]
          },
          {
            "match": [
              {
                "uri": {
                  "MatchType": {
                    "Exact": "/foo"
                  }
                }
              }
            ],
            "route": [
              {
                "destination": {
                  "host": "service1.ns.svc.cluster.local",
                  "port": {
                    "Port": {
                      "Number": 4200
                    }
                  }
                },
                "weight": 100
              }
            ]
          }
        ]
      },
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.VirtualService"
    }
  ]
}`)

func datasetExtensionsV1beta1Ingress_merge_1_expectedJsonBytes() ([]byte, error) {
	return _datasetExtensionsV1beta1Ingress_merge_1_expectedJson, nil
}

func datasetExtensionsV1beta1Ingress_merge_1_expectedJson() (*asset, error) {
	bytes, err := datasetExtensionsV1beta1Ingress_merge_1_expectedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/extensions/v1beta1/ingress_merge_1_expected.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingIstioIoV1alpha3DestinationruleYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: tcp-echo-destination
spec:
  host: tcp-echo
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
`)

func datasetNetworkingIstioIoV1alpha3DestinationruleYamlBytes() ([]byte, error) {
	return _datasetNetworkingIstioIoV1alpha3DestinationruleYaml, nil
}

func datasetNetworkingIstioIoV1alpha3DestinationruleYaml() (*asset, error) {
	bytes, err := datasetNetworkingIstioIoV1alpha3DestinationruleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking.istio.io/v1alpha3/destinationRule.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJson = []byte(`{
  "type.googleapis.com/istio.networking.v1alpha3.DestinationRule": [
    {
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.DestinationRule",
      "Metadata": {
        "name": "tcp-echo-destination"
      },
      "Resource": {
        "host": "tcp-echo",
        "subsets": [
          {
            "labels": {
              "version": "v1"
            },
            "name": "v1"
          },
          {
            "labels": {
              "version": "v2"
            },
            "name": "v2"
          }
        ]
      }
    }
  ]
}
`)

func datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJsonBytes() ([]byte, error) {
	return _datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJson, nil
}

func datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJson() (*asset, error) {
	bytes, err := datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking.istio.io/v1alpha3/destinationRule_expected.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingIstioIoV1alpha3GatewayYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: helloworld-gateway
spec:
  selector:
    istio: ingressgateway # use istio default controller
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
`)

func datasetNetworkingIstioIoV1alpha3GatewayYamlBytes() ([]byte, error) {
	return _datasetNetworkingIstioIoV1alpha3GatewayYaml, nil
}

func datasetNetworkingIstioIoV1alpha3GatewayYaml() (*asset, error) {
	bytes, err := datasetNetworkingIstioIoV1alpha3GatewayYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking.istio.io/v1alpha3/gateway.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingIstioIoV1alpha3Gateway_expectedJson = []byte(`{
  "type.googleapis.com/istio.networking.v1alpha3.Gateway": [
    {
      "TypeURL": "type.googleapis.com/istio.networking.v1alpha3.Gateway",
      "Metadata": {
        "name": "helloworld-gateway"
      },
      "Resource": {
        "selector": {
          "istio": "ingressgateway"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      }
    }
  ]
}
`)

func datasetNetworkingIstioIoV1alpha3Gateway_expectedJsonBytes() ([]byte, error) {
	return _datasetNetworkingIstioIoV1alpha3Gateway_expectedJson, nil
}

func datasetNetworkingIstioIoV1alpha3Gateway_expectedJson() (*asset, error) {
	bytes, err := datasetNetworkingIstioIoV1alpha3Gateway_expectedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking.istio.io/v1alpha3/gateway_expected.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"dataset/config.istio.io/v1alpha2/circonus.yaml": datasetConfigIstioIoV1alpha2CirconusYaml,
	"dataset/config.istio.io/v1alpha2/circonus_expected.json": datasetConfigIstioIoV1alpha2Circonus_expectedJson,
	"dataset/extensions/v1beta1/ingress_basic.yaml": datasetExtensionsV1beta1Ingress_basicYaml,
	"dataset/extensions/v1beta1/ingress_basic_expected.json": datasetExtensionsV1beta1Ingress_basic_expectedJson,
	"dataset/extensions/v1beta1/ingress_basic_meshconfig.yaml": datasetExtensionsV1beta1Ingress_basic_meshconfigYaml,
	"dataset/extensions/v1beta1/ingress_merge_0.skip": datasetExtensionsV1beta1Ingress_merge_0Skip,
	"dataset/extensions/v1beta1/ingress_merge_0.yaml": datasetExtensionsV1beta1Ingress_merge_0Yaml,
	"dataset/extensions/v1beta1/ingress_merge_0_expected.json": datasetExtensionsV1beta1Ingress_merge_0_expectedJson,
	"dataset/extensions/v1beta1/ingress_merge_0_meshconfig.yaml": datasetExtensionsV1beta1Ingress_merge_0_meshconfigYaml,
	"dataset/extensions/v1beta1/ingress_merge_1.yaml": datasetExtensionsV1beta1Ingress_merge_1Yaml,
	"dataset/extensions/v1beta1/ingress_merge_1_expected.json": datasetExtensionsV1beta1Ingress_merge_1_expectedJson,
	"dataset/networking.istio.io/v1alpha3/destinationRule.yaml": datasetNetworkingIstioIoV1alpha3DestinationruleYaml,
	"dataset/networking.istio.io/v1alpha3/destinationRule_expected.json": datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJson,
	"dataset/networking.istio.io/v1alpha3/gateway.yaml": datasetNetworkingIstioIoV1alpha3GatewayYaml,
	"dataset/networking.istio.io/v1alpha3/gateway_expected.json": datasetNetworkingIstioIoV1alpha3Gateway_expectedJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"dataset": &bintree{nil, map[string]*bintree{
		"config.istio.io": &bintree{nil, map[string]*bintree{
			"v1alpha2": &bintree{nil, map[string]*bintree{
				"circonus.yaml": &bintree{datasetConfigIstioIoV1alpha2CirconusYaml, map[string]*bintree{}},
				"circonus_expected.json": &bintree{datasetConfigIstioIoV1alpha2Circonus_expectedJson, map[string]*bintree{}},
			}},
		}},
		"extensions": &bintree{nil, map[string]*bintree{
			"v1beta1": &bintree{nil, map[string]*bintree{
				"ingress_basic.yaml": &bintree{datasetExtensionsV1beta1Ingress_basicYaml, map[string]*bintree{}},
				"ingress_basic_expected.json": &bintree{datasetExtensionsV1beta1Ingress_basic_expectedJson, map[string]*bintree{}},
				"ingress_basic_meshconfig.yaml": &bintree{datasetExtensionsV1beta1Ingress_basic_meshconfigYaml, map[string]*bintree{}},
				"ingress_merge_0.skip": &bintree{datasetExtensionsV1beta1Ingress_merge_0Skip, map[string]*bintree{}},
				"ingress_merge_0.yaml": &bintree{datasetExtensionsV1beta1Ingress_merge_0Yaml, map[string]*bintree{}},
				"ingress_merge_0_expected.json": &bintree{datasetExtensionsV1beta1Ingress_merge_0_expectedJson, map[string]*bintree{}},
				"ingress_merge_0_meshconfig.yaml": &bintree{datasetExtensionsV1beta1Ingress_merge_0_meshconfigYaml, map[string]*bintree{}},
				"ingress_merge_1.yaml": &bintree{datasetExtensionsV1beta1Ingress_merge_1Yaml, map[string]*bintree{}},
				"ingress_merge_1_expected.json": &bintree{datasetExtensionsV1beta1Ingress_merge_1_expectedJson, map[string]*bintree{}},
			}},
		}},
		"networking.istio.io": &bintree{nil, map[string]*bintree{
			"v1alpha3": &bintree{nil, map[string]*bintree{
				"destinationRule.yaml": &bintree{datasetNetworkingIstioIoV1alpha3DestinationruleYaml, map[string]*bintree{}},
				"destinationRule_expected.json": &bintree{datasetNetworkingIstioIoV1alpha3Destinationrule_expectedJson, map[string]*bintree{}},
				"gateway.yaml": &bintree{datasetNetworkingIstioIoV1alpha3GatewayYaml, map[string]*bintree{}},
				"gateway_expected.json": &bintree{datasetNetworkingIstioIoV1alpha3Gateway_expectedJson, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
