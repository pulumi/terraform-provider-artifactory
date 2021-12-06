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
	case "artifactory:index/apiKey:ApiKey":
		r = &ApiKey{}
	case "artifactory:index/certificate:Certificate":
		r = &Certificate{}
	case "artifactory:index/debianRepository:DebianRepository":
		r = &DebianRepository{}
	case "artifactory:index/dockerV1Repository:DockerV1Repository":
		r = &DockerV1Repository{}
	case "artifactory:index/dockerV2Repository:DockerV2Repository":
		r = &DockerV2Repository{}
	case "artifactory:index/generalSecurity:GeneralSecurity":
		r = &GeneralSecurity{}
	case "artifactory:index/goRepository:GoRepository":
		r = &GoRepository{}
	case "artifactory:index/group:Group":
		r = &Group{}
	case "artifactory:index/keypair:Keypair":
		r = &Keypair{}
	case "artifactory:index/localBowerRepository:LocalBowerRepository":
		r = &LocalBowerRepository{}
	case "artifactory:index/localChefRepository:LocalChefRepository":
		r = &LocalChefRepository{}
	case "artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository":
		r = &LocalCocoapodsRepository{}
	case "artifactory:index/localComposerRepository:LocalComposerRepository":
		r = &LocalComposerRepository{}
	case "artifactory:index/localConanRepository:LocalConanRepository":
		r = &LocalConanRepository{}
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
	case "artifactory:index/localRepository:LocalRepository":
		r = &LocalRepository{}
	case "artifactory:index/localRpmRepository:LocalRpmRepository":
		r = &LocalRpmRepository{}
	case "artifactory:index/localSbtRepository:LocalSbtRepository":
		r = &LocalSbtRepository{}
	case "artifactory:index/localVagrantRepository:LocalVagrantRepository":
		r = &LocalVagrantRepository{}
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
	case "artifactory:index/remoteCargoRepository:RemoteCargoRepository":
		r = &RemoteCargoRepository{}
	case "artifactory:index/remoteDockerRepository:RemoteDockerRepository":
		r = &RemoteDockerRepository{}
	case "artifactory:index/remoteHelmRepository:RemoteHelmRepository":
		r = &RemoteHelmRepository{}
	case "artifactory:index/remoteNpmRepository:RemoteNpmRepository":
		r = &RemoteNpmRepository{}
	case "artifactory:index/remoteRepository:RemoteRepository":
		r = &RemoteRepository{}
	case "artifactory:index/replicationConfig:ReplicationConfig":
		r = &ReplicationConfig{}
	case "artifactory:index/samlSettings:SamlSettings":
		r = &SamlSettings{}
	case "artifactory:index/singleReplicationConfig:SingleReplicationConfig":
		r = &SingleReplicationConfig{}
	case "artifactory:index/user:User":
		r = &User{}
	case "artifactory:index/virtualRepository:VirtualRepository":
		r = &VirtualRepository{}
	case "artifactory:index/xrayPolicy:XrayPolicy":
		r = &XrayPolicy{}
	case "artifactory:index/xrayWatch:XrayWatch":
		r = &XrayWatch{}
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
		"index/apiKey",
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
		"index/localBowerRepository",
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
		"index/localRepository",
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
		"index/remoteCargoRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteDockerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteHelmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteNpmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/remoteRepository",
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
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/virtualRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/xrayPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"artifactory",
		"index/xrayWatch",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"artifactory",
		&pkg{version},
	)
}
