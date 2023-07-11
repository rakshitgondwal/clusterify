package civo

import "github.com/civo/civogo"

type KubernetesClusterConfig struct {
	Name              string `json:"name,omitempty"`
	Region            string `json:"region,omitempty"`
	ClusterType       string `json:"cluster_type,omitempty"`
	NumTargetNodes    int    `json:"num_target_nodes,omitempty"`
	TargetNodesSize   string `json:"target_nodes_size,omitempty"`
	KubernetesVersion string `json:"kubernetes_version,omitempty"`
	NodeDestroy       string `json:"node_destroy,omitempty"`
	NetworkID         string `json:"network_id,omitempty"`
	Tags              string `json:"tags,omitempty"`
	Applications      string `json:"applications,omitempty"`
	InstanceFirewall  string `json:"instance_firewall,omitempty"`
	FirewallRule      string `json:"firewall_rule,omitempty"`
	CNIPlugin         string `json:"cni_plugin,omitempty"`
}

func CreateCluster(config civogo.KubernetesClusterConfig, apiKey string) (*civogo.KubernetesCluster, error) {
	client, err := civogo.NewClient(apiKey, config.Region)
	if err != nil {
		return nil, err
	}

	cluster, err := client.NewKubernetesClusters(&config)
	if err != nil{
		return nil, err
	}

	return cluster, nil
}

func CreateClusterConfig(name string, region string) civogo.KubernetesClusterConfig {
	var clusterConfig civogo.KubernetesClusterConfig
	clusterConfig.Name = name
	clusterConfig.Region = region

	return clusterConfig
}
