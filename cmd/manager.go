package cmd

import (
	"fmt"
	"os"

	"github.com/rancher/support-bundle-kit/pkg/manager"
	"github.com/spf13/cobra"
)

var (
	sbm = &manager.SupportBundleManager{}
)

// managerCmd represents the manager command
var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Support Bundle Kit manager",
	Long: `Support Bundle Kit manager

The manager collects following items:
- Cluster level bundle. Including resource manifests and pod logs.
- Any external bundles. e.g., Longhorn support bundle.

And it also waits for reports from support bundle agents. The reports contain:
- Logs of each node.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := sbm.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(managerCmd)
	managerCmd.PersistentFlags().StringVar(&sbm.NamespaceList, "namespaces", os.Getenv("SUPPORT_BUNDLE_TARGET_NAMESPACES"), "List of namespaces delimited by ,")
	managerCmd.PersistentFlags().StringVar(&sbm.BundleName, "bundlename", os.Getenv("SUPPORT_BUNDLE_NAME"), "The support bundle name")
	managerCmd.PersistentFlags().StringVar(&sbm.OutputDir, "outdir", os.Getenv("SUPPORT_BUNDLE_OUTPUT_DIR"), "The directory to store the bundle")
	managerCmd.PersistentFlags().StringVar(&sbm.ManagerPodIP, "manager-pod-ip", os.Getenv("SUPPORT_BUNDLE_MANAGER_POD_IP"), "The support bundle manager's IP (pod runs this app)")
	managerCmd.PersistentFlags().StringVar(&sbm.ImageName, "image-name", os.Getenv("SUPPORT_BUNDLE_IMAGE"), "The support bundle image")
	managerCmd.PersistentFlags().StringVar(&sbm.ImagePullPolicy, "image-pull-policy", os.Getenv("SUPPORT_BUNDLE_IMAGE_PULL_POLICY"), "Pull policy of the support bundle image")

}
