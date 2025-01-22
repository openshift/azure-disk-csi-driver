/*
Copyright 2017 The Kubernetes Authors.

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

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/klog/v2"
	"sigs.k8s.io/azuredisk-csi-driver/pkg/azuredisk"
)

func init() {
	klog.InitFlags(nil)
	driverOptions.AddFlags().VisitAll(func(f *flag.Flag) {
		flag.CommandLine.Var(f.Value, f.Name, f.Usage)
	})
}

var (
	version        = flag.Bool("version", false, "Print the version and exit.")
	metricsAddress = flag.String("metrics-address", "", "export the metrics")
	driverOptions  azuredisk.DriverOptions
)

func main() {
	flag.Parse()
	if *version {
		info, err := azuredisk.GetVersionYAML(driverOptions.DriverName)
		if err != nil {
			klog.Fatalln(err)
		}
		fmt.Println(info) // nolint
		os.Exit(0)
	}

	exportMetrics()
	handle()
	os.Exit(0)
}

func handle() {
	driver := azuredisk.NewDriver(&driverOptions)
	if driver == nil {
		klog.Fatalln("Failed to initialize azuredisk CSI Driver")
	}
	if err := driver.Run(context.Background()); err != nil {
		klog.Fatalf("Failed to run azuredisk CSI Driver: %v", err)
	}
}

func exportMetrics() {
	if *metricsAddress == "" {
		return
	}
	l, err := net.Listen("tcp", *metricsAddress)
	if err != nil {
		klog.Warningf("failed to get listener for metrics endpoint: %v", err)
		return
	}
	serve(context.Background(), l, serveMetrics)
}

func serve(_ context.Context, l net.Listener, serveFunc func(net.Listener) error) {
	path := l.Addr().String()
	klog.V(2).Infof("set up prometheus server on %v", path)
	go func() {
		defer l.Close()
		if err := serveFunc(l); err != nil {
			klog.Fatalf("serve failure(%v), address(%v)", err, path)
		}
	}()
}

func serveMetrics(l net.Listener) error {
	m := http.NewServeMux()
	m.Handle("/metrics", legacyregistry.Handler()) //nolint, because azure cloud provider uses legacyregistry currently
	return trapClosedConnErr(http.Serve(l, m))
}

func trapClosedConnErr(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "use of closed network connection") {
		return nil
	}
	return err
}
