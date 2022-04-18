// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "artifactory:index/accessToken:AccessToken":
		r = &AccessToken{}
	case "artifactory:index/alpineRepository:AlpineRepository":
		r = &AlpineRepository{}
	case "artifactory:index/anonymousUser:AnonymousUser":
		r = &AnonymousUser{}
	case "artifactory:index/apiKey:ApiKey":
		r = &ApiKey{}
	case "artifactory:index/artifactPropertyWebhook:ArtifactPropertyWebhook":
		r = &ArtifactPropertyWebhook{}
	case "artifactory:index/artifactWebhook:ArtifactWebhook":
		r = &ArtifactWebhook{}
	case "artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook":
		r = &ArtifactoryReleaseBundleWebhook{}
	case "artifactory:index/backup:Backup":
		r = &Backup{}
	case "artifactory:index/buildWebhook:BuildWebhook":
		r = &BuildWebhook{}
	case "artifactory:index/certificate:Certificate":
		r = &Certificate{}
	case "artifactory:index/debianRepository:DebianRepository":
		r = &DebianRepository{}
	case "artifactory:index/distributionWebhook:DistributionWebhook":
		r = &DistributionWebhook{}
	case "artifactory:index/dockerV1Repository:DockerV1Repository":
		r = &DockerV1Repository{}
	case "artifactory:index/dockerV2Repository:DockerV2Repository":
		r = &DockerV2Repository{}
	case "artifactory:index/dockerWebhook:DockerWebhook":
		r = &DockerWebhook{}
	case "artifactory:index/federatedAlpineRepository:FederatedAlpineRepository":
		r = &FederatedAlpineRepository{}
	case "artifactory:index/federatedBowerRepository:FederatedBowerRepository":
		r = &FederatedBowerRepository{}
	case "artifactory:index/federatedCargoRepository:FederatedCargoRepository":
		r = &FederatedCargoRepository{}
	case "artifactory:index/federatedChefRepository:FederatedChefRepository":
		r = &FederatedChefRepository{}
	case "artifactory:index/federatedCocoapodsRepository:FederatedCocoapodsRepository":
		r = &FederatedCocoapodsRepository{}
	case "artifactory:index/federatedComposerRepository:FederatedComposerRepository":
		r = &FederatedComposerRepository{}
	case "artifactory:index/federatedConanRepository:FederatedConanRepository":
		r = &FederatedConanRepository{}
	case "artifactory:index/federatedCondaRepository:FederatedCondaRepository":
		r = &FederatedCondaRepository{}
	case "artifactory:index/federatedCranRepository:FederatedCranRepository":
		r = &FederatedCranRepository{}
	case "artifactory:index/federatedDebianRepository:FederatedDebianRepository":
		r = &FederatedDebianRepository{}
	case "artifactory:index/federatedDockerRepository:FederatedDockerRepository":
		r = &FederatedDockerRepository{}
	case "artifactory:index/federatedGemsRepository:FederatedGemsRepository":
		r = &FederatedGemsRepository{}
	case "artifactory:index/federatedGenericRepository:FederatedGenericRepository":
		r = &FederatedGenericRepository{}
	case "artifactory:index/federatedGitltfsRepository:FederatedGitltfsRepository":
		r = &FederatedGitltfsRepository{}
	case "artifactory:index/federatedGoRepository:FederatedGoRepository":
		r = &FederatedGoRepository{}
	case "artifactory:index/federatedGradleRepository:FederatedGradleRepository":
		r = &FederatedGradleRepository{}
	case "artifactory:index/federatedHelmRepository:FederatedHelmRepository":
		r = &FederatedHelmRepository{}
	case "artifactory:index/federatedIvyRepository:FederatedIvyRepository":
		r = &FederatedIvyRepository{}
	case "artifactory:index/federatedMavenRepository:FederatedMavenRepository":
		r = &FederatedMavenRepository{}
	case "artifactory:index/federatedNpmRepository:FederatedNpmRepository":
		r = &FederatedNpmRepository{}
	case "artifactory:index/federatedNugetRepository:FederatedNugetRepository":
		r = &FederatedNugetRepository{}
	case "artifactory:index/federatedOpkgRepository:FederatedOpkgRepository":
		r = &FederatedOpkgRepository{}
	case "artifactory:index/federatedPuppetRepository:FederatedPuppetRepository":
		r = &FederatedPuppetRepository{}
	case "artifactory:index/federatedPypiRepository:FederatedPypiRepository":
		r = &FederatedPypiRepository{}
	case "artifactory:index/federatedRpmRepository:FederatedRpmRepository":
		r = &FederatedRpmRepository{}
	case "artifactory:index/federatedSbtRepository:FederatedSbtRepository":
		r = &FederatedSbtRepository{}
	case "artifactory:index/federatedVagrantRepository:FederatedVagrantRepository":
		r = &FederatedVagrantRepository{}
	case "artifactory:index/generalSecurity:GeneralSecurity":
		r = &GeneralSecurity{}
	case "artifactory:index/goRepository:GoRepository":
		r = &GoRepository{}
	case "artifactory:index/group:Group":
		r = &Group{}
	case "artifactory:index/keypair:Keypair":
		r = &Keypair{}
	case "artifactory:index/ldapGroupSetting:LdapGroupSetting":
		r = &LdapGroupSetting{}
	case "artifactory:index/ldapSetting:LdapSetting":
		r = &LdapSetting{}
	case "artifactory:index/localBowerRepository:LocalBowerRepository":
		r = &LocalBowerRepository{}
	case "artifactory:index/localCargoRepository:LocalCargoRepository":
		r = &LocalCargoRepository{}
	case "artifactory:index/localChefRepository:LocalChefRepository":
		r = &LocalChefRepository{}
	case "artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository":
		r = &LocalCocoapodsRepository{}
	case "artifactory:index/localComposerRepository:LocalComposerRepository":
		r = &LocalComposerRepository{}
	case "artifactory:index/localConanRepository:LocalConanRepository":
		r = &LocalConanRepository{}
	case "artifactory:index/localCondaRepository:LocalCondaRepository":
		r = &LocalCondaRepository{}
	case "artifactory:index/localCranRepository:LocalCranRepository":
		r = &LocalCranRepository{}
	case "artifactory:index/localGemsRepository:LocalGemsRepository":
		r = &LocalGemsRepository{}
	case "artifactory:index/localGenericRepository:LocalGenericRepository":
		r = &LocalGenericRepository{}
	case "artifactory:index/localGitltfsRepository:LocalGitltfsRepository":
		r = &LocalGitltfsRepository{}
	case "artifactory:index/localGoRepository:LocalGoRepository":
		r = &LocalGoRepository{}
	case "artifactory:index/localGradleRepository:LocalGradleRepository":
		r = &LocalGradleRepository{}
	case "artifactory:index/localHelmRepository:LocalHelmRepository":
		r = &LocalHelmRepository{}
	case "artifactory:index/localIvyRepository:LocalIvyRepository":
		r = &LocalIvyRepository{}
	case "artifactory:index/localMavenRepository:LocalMavenRepository":
		r = &LocalMavenRepository{}
	case "artifactory:index/localNpmRepository:LocalNpmRepository":
		r = &LocalNpmRepository{}
	case "artifactory:index/localNugetRepository:LocalNugetRepository":
		r = &LocalNugetRepository{}
	case "artifactory:index/localOpkgRepository:LocalOpkgRepository":
		r = &LocalOpkgRepository{}
	case "artifactory:index/localPuppetRepository:LocalPuppetRepository":
		r = &LocalPuppetRepository{}
	case "artifactory:index/localPypiRepository:LocalPypiRepository":
		r = &LocalPypiRepository{}
	case "artifactory:index/localRpmRepository:LocalRpmRepository":
		r = &LocalRpmRepository{}
	case "artifactory:index/localSbtRepository:LocalSbtRepository":
		r = &LocalSbtRepository{}
	case "artifactory:index/localVagrantRepository:LocalVagrantRepository":
		r = &LocalVagrantRepository{}
	case "artifactory:index/managedUser:ManagedUser":
		r = &ManagedUser{}
	case "artifactory:index/mavenRepository:MavenRepository":
		r = &MavenRepository{}
	case "artifactory:index/oauthSettings:OauthSettings":
		r = &OauthSettings{}
	case "artifactory:index/permissionTarget:PermissionTarget":
		r = &PermissionTarget{}
	case "artifactory:index/permissionTargets:PermissionTargets":
		r = &PermissionTargets{}
	case "artifactory:index/pullReplication:PullReplication":
		r = &PullReplication{}
	case "artifactory:index/pushReplication:PushReplication":
		r = &PushReplication{}
	case "artifactory:index/releaseBundleWebhook:ReleaseBundleWebhook":
		r = &ReleaseBundleWebhook{}
	case "artifactory:index/remoteAlpineRepository:RemoteAlpineRepository":
		r = &RemoteAlpineRepository{}
	case "artifactory:index/remoteBowerRepository:RemoteBowerRepository":
		r = &RemoteBowerRepository{}
	case "artifactory:index/remoteCargoRepository:RemoteCargoRepository":
		r = &RemoteCargoRepository{}
	case "artifactory:index/remoteChefRepository:RemoteChefRepository":
		r = &RemoteChefRepository{}
	case "artifactory:index/remoteCocoapodsRepository:RemoteCocoapodsRepository":
		r = &RemoteCocoapodsRepository{}
	case "artifactory:index/remoteComposerRepository:RemoteComposerRepository":
		r = &RemoteComposerRepository{}
	case "artifactory:index/remoteConanRepository:RemoteConanRepository":
		r = &RemoteConanRepository{}
	case "artifactory:index/remoteCondaRepository:RemoteCondaRepository":
		r = &RemoteCondaRepository{}
	case "artifactory:index/remoteCranRepository:RemoteCranRepository":
		r = &RemoteCranRepository{}
	case "artifactory:index/remoteDebianRepository:RemoteDebianRepository":
		r = &RemoteDebianRepository{}
	case "artifactory:index/remoteDockerRepository:RemoteDockerRepository":
		r = &RemoteDockerRepository{}
	case "artifactory:index/remoteGemsRepository:RemoteGemsRepository":
		r = &RemoteGemsRepository{}
	case "artifactory:index/remoteGenericRepository:RemoteGenericRepository":
		r = &RemoteGenericRepository{}
	case "artifactory:index/remoteGitlfsRepository:RemoteGitlfsRepository":
		r = &RemoteGitlfsRepository{}
	case "artifactory:index/remoteGoRepository:RemoteGoRepository":
		r = &RemoteGoRepository{}
	case "artifactory:index/remoteGradleRepository:RemoteGradleRepository":
		r = &RemoteGradleRepository{}
	case "artifactory:index/remoteHelmRepository:RemoteHelmRepository":
		r = &RemoteHelmRepository{}
	case "artifactory:index/remoteIvyRepository:RemoteIvyRepository":
		r = &RemoteIvyRepository{}
	case "artifactory:index/remoteMavenRepository:RemoteMavenRepository":
		r = &RemoteMavenRepository{}
	case "artifactory:index/remoteNpmRepository:RemoteNpmRepository":
		r = &RemoteNpmRepository{}
	case "artifactory:index/remoteNugetRepository:RemoteNugetRepository":
		r = &RemoteNugetRepository{}
	case "artifactory:index/remoteOpkgRepository:RemoteOpkgRepository":
		r = &RemoteOpkgRepository{}
	case "artifactory:index/remoteP2Repository:RemoteP2Repository":
		r = &RemoteP2Repository{}
	case "artifactory:index/remotePuppetRepository:RemotePuppetRepository":
		r = &RemotePuppetRepository{}
	case "artifactory:index/remotePypiRepository:RemotePypiRepository":
		r = &RemotePypiRepository{}
	case "artifactory:index/remoteRpmRepository:RemoteRpmRepository":
		r = &RemoteRpmRepository{}
	case "artifactory:index/remoteSbtRepository:RemoteSbtRepository":
		r = &RemoteSbtRepository{}
	case "artifactory:index/remoteVcsRepository:RemoteVcsRepository":
		r = &RemoteVcsRepository{}
	case "artifactory:index/replicationConfig:ReplicationConfig":
		r = &ReplicationConfig{}
	case "artifactory:index/samlSettings:SamlSettings":
		r = &SamlSettings{}
	case "artifactory:index/singleReplicationConfig:SingleReplicationConfig":
		r = &SingleReplicationConfig{}
	case "artifactory:index/unmanagedUser:UnmanagedUser":
		r = &UnmanagedUser{}
	case "artifactory:index/user:User":
		r = &User{}
	case "artifactory:index/virtualAlpineRepository:VirtualAlpineRepository":
		r = &VirtualAlpineRepository{}
	case "artifactory:index/virtualBowerRepository:VirtualBowerRepository":
		r = &VirtualBowerRepository{}
	case "artifactory:index/virtualChefRepository:VirtualChefRepository":
		r = &VirtualChefRepository{}
	case "artifactory:index/virtualComposerRepository:VirtualComposerRepository":
		r = &VirtualComposerRepository{}
	case "artifactory:index/virtualConanRepository:VirtualConanRepository":
		r = &VirtualConanRepository{}
	case "artifactory:index/virtualCondaRepository:VirtualCondaRepository":
		r = &VirtualCondaRepository{}
	case "artifactory:index/virtualCranRepository:VirtualCranRepository":
		r = &VirtualCranRepository{}
	case "artifactory:index/virtualDebianRepository:VirtualDebianRepository":
		r = &VirtualDebianRepository{}
	case "artifactory:index/virtualDockerRepository:VirtualDockerRepository":
		r = &VirtualDockerRepository{}
	case "artifactory:index/virtualGemsRepository:VirtualGemsRepository":
		r = &VirtualGemsRepository{}
	case "artifactory:index/virtualGenericRepository:VirtualGenericRepository":
		r = &VirtualGenericRepository{}
	case "artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository":
		r = &VirtualGitlfsRepository{}
	case "artifactory:index/virtualGradleRepository:VirtualGradleRepository":
		r = &VirtualGradleRepository{}
	case "artifactory:index/virtualHelmRepository:VirtualHelmRepository":
		r = &VirtualHelmRepository{}
	case "artifactory:index/virtualIvyRepository:VirtualIvyRepository":
		r = &VirtualIvyRepository{}
	case "artifactory:index/virtualNpmRepository:VirtualNpmRepository":
		r = &VirtualNpmRepository{}
	case "artifactory:index/virtualNugetRepository:VirtualNugetRepository":
		r = &VirtualNugetRepository{}
	case "artifactory:index/virtualP2Repository:VirtualP2Repository":
		r = &VirtualP2Repository{}
	case "artifactory:index/virtualPuppetRepository:VirtualPuppetRepository":
		r = &VirtualPuppetRepository{}
	case "artifactory:index/virtualPypiRepository:VirtualPypiRepository":
		r = &VirtualPypiRepository{}
	case "artifactory:index/virtualRpmRepository:VirtualRpmRepository":
		r = &VirtualRpmRepository{}
	case "artifactory:index/virtualSbtRepository:VirtualSbtRepository":
		r = &VirtualSbtRepository{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:artifactory" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/accessToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/alpineRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/anonymousUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/apiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/artifactPropertyWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/artifactWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/artifactoryReleaseBundleWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/backup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/buildWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/debianRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/distributionWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/dockerV1Repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/dockerV2Repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/dockerWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedAlpineRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedBowerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedCargoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedChefRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedCocoapodsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedComposerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedConanRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedCondaRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedCranRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedDebianRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedDockerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedGemsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedGenericRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedGitltfsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedGoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedGradleRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedHelmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedIvyRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedMavenRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedNpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedNugetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedOpkgRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedPuppetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedPypiRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedRpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedSbtRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/federatedVagrantRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/generalSecurity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/goRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/keypair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/ldapGroupSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/ldapSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localBowerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localCargoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localChefRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localCocoapodsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localComposerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localConanRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localCondaRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localCranRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localGemsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localGenericRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localGitltfsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localGoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localGradleRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localHelmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localIvyRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localMavenRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localNpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localNugetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localOpkgRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localPuppetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localPypiRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localRpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localSbtRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/localVagrantRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/managedUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/mavenRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/oauthSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/permissionTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/permissionTargets",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/pullReplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/pushReplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/releaseBundleWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteAlpineRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteBowerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteCargoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteChefRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteCocoapodsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteComposerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteConanRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteCondaRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteCranRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteDebianRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteDockerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteGemsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteGenericRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteGitlfsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteGoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteGradleRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteHelmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteIvyRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteMavenRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteNpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteNugetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteOpkgRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteP2Repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remotePuppetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remotePypiRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteRpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteSbtRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteVcsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/replicationConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/samlSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/singleReplicationConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/unmanagedUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualAlpineRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualBowerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualChefRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualComposerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualConanRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualCondaRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualCranRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualDebianRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualDockerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualGemsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualGenericRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualGitlfsRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualGradleRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualHelmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualIvyRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualNpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualNugetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualP2Repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualPuppetRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualPypiRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualRpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualSbtRepository",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"artifactory",
		&pkg{version},
	)
}
