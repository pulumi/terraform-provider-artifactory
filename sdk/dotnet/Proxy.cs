// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Provides an Artifactory Proxy resource.
    /// 
    /// This resource configuration corresponds to 'proxies' config block in system configuration XML
    /// (REST endpoint: [artifactory/api/system/configuration](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration)).
    /// 
    /// ~&gt;The `artifactory.Proxy` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_proxy = new Artifactory.Proxy("my-proxy", new()
    ///     {
    ///         Host = "my-proxy.mycompany.com",
    ///         Key = "my-proxy",
    ///         NtDomain = "MYCOMPANY",
    ///         NtHost = "MYCOMPANY.COM",
    ///         Password = "password",
    ///         PlatformDefault = false,
    ///         Port = 8888,
    ///         RedirectToHosts = new[]
    ///         {
    ///             "redirec-host.mycompany.com",
    ///         },
    ///         Services = new[]
    ///         {
    ///             "jfrt",
    ///             "jfxr",
    ///         },
    ///         Username = "user1",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Current Proxy can be imported using `proxy-key` from Artifactory as the `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/proxy:Proxy my-proxy proxy-key
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/proxy:Proxy")]
    public partial class Proxy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the proxy host.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the proxy.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The proxy domain/realm name.
        /// </summary>
        [Output("ntDomain")]
        public Output<string?> NtDomain { get; private set; } = null!;

        /// <summary>
        /// The computer name of the machine (the machine connecting to the NTLM proxy).
        /// </summary>
        [Output("ntHost")]
        public Output<string?> NtHost { get; private set; } = null!;

        /// <summary>
        /// The proxy password when authentication credentials are required.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        /// </summary>
        [Output("platformDefault")]
        public Output<bool> PlatformDefault { get; private set; } = null!;

        /// <summary>
        /// The proxy port number.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        /// </summary>
        [Output("redirectToHosts")]
        public Output<ImmutableArray<string>> RedirectToHosts { get; private set; } = null!;

        /// <summary>
        /// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;

        /// <summary>
        /// The proxy username when authentication credentials are required.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Proxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Proxy(string name, ProxyArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/proxy:Proxy", name, args ?? new ProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Proxy(string name, Input<string> id, ProxyState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/proxy:Proxy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Proxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Proxy Get(string name, Input<string> id, ProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new Proxy(name, id, state, options);
        }
    }

    public sealed class ProxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the proxy host.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The unique ID of the proxy.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The proxy domain/realm name.
        /// </summary>
        [Input("ntDomain")]
        public Input<string>? NtDomain { get; set; }

        /// <summary>
        /// The computer name of the machine (the machine connecting to the NTLM proxy).
        /// </summary>
        [Input("ntHost")]
        public Input<string>? NtHost { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The proxy password when authentication credentials are required.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        /// </summary>
        [Input("platformDefault")]
        public Input<bool>? PlatformDefault { get; set; }

        /// <summary>
        /// The proxy port number.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("redirectToHosts")]
        private InputList<string>? _redirectToHosts;

        /// <summary>
        /// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        /// </summary>
        public InputList<string> RedirectToHosts
        {
            get => _redirectToHosts ?? (_redirectToHosts = new InputList<string>());
            set => _redirectToHosts = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// The proxy username when authentication credentials are required.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProxyArgs()
        {
        }
        public static new ProxyArgs Empty => new ProxyArgs();
    }

    public sealed class ProxyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the proxy host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The unique ID of the proxy.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The proxy domain/realm name.
        /// </summary>
        [Input("ntDomain")]
        public Input<string>? NtDomain { get; set; }

        /// <summary>
        /// The computer name of the machine (the machine connecting to the NTLM proxy).
        /// </summary>
        [Input("ntHost")]
        public Input<string>? NtHost { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The proxy password when authentication credentials are required.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
        /// </summary>
        [Input("platformDefault")]
        public Input<bool>? PlatformDefault { get; set; }

        /// <summary>
        /// The proxy port number.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("redirectToHosts")]
        private InputList<string>? _redirectToHosts;

        /// <summary>
        /// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
        /// </summary>
        public InputList<string> RedirectToHosts
        {
            get => _redirectToHosts ?? (_redirectToHosts = new InputList<string>());
            set => _redirectToHosts = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// The proxy username when authentication credentials are required.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProxyState()
        {
        }
        public static new ProxyState Empty => new ProxyState();
    }
}
