// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory Proxy resource.
 *
 * This resource configuration corresponds to 'proxies' config block in system configuration XML
 * (REST endpoint: [artifactory/api/system/configuration](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration)).
 *
 * ~>The `artifactory.Proxy` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my_proxy = new artifactory.Proxy("my-proxy", {
 *     key: "my-proxy",
 *     host: "my-proxy.mycompany.com",
 *     port: 8888,
 *     username: "user1",
 *     password: "password",
 *     ntHost: "MYCOMPANY.COM",
 *     ntDomain: "MYCOMPANY",
 *     platformDefault: false,
 *     redirectToHosts: ["redirec-host.mycompany.com"],
 *     services: [
 *         "jfrt",
 *         "jfxr",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Current Proxy can be imported using `proxy-key` from Artifactory as the `ID`, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/proxy:Proxy my-proxy proxy-key
 * ```
 */
export class Proxy extends pulumi.CustomResource {
    /**
     * Get an existing Proxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyState, opts?: pulumi.CustomResourceOptions): Proxy {
        return new Proxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/proxy:Proxy';

    /**
     * Returns true if the given object is an instance of Proxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Proxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Proxy.__pulumiType;
    }

    /**
     * The name of the proxy host.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The unique ID of the proxy.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The proxy domain/realm name.
     */
    public readonly ntDomain!: pulumi.Output<string>;
    /**
     * The computer name of the machine (the machine connecting to the NTLM proxy).
     */
    public readonly ntHost!: pulumi.Output<string>;
    /**
     * The proxy password when authentication credentials are required.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
     */
    public readonly platformDefault!: pulumi.Output<boolean>;
    /**
     * The proxy port number.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
     */
    public readonly redirectToHosts!: pulumi.Output<string[] | undefined>;
    /**
     * An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
     */
    public readonly services!: pulumi.Output<string[] | undefined>;
    /**
     * The proxy username when authentication credentials are required.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a Proxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyArgs | ProxyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyState | undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["ntDomain"] = state ? state.ntDomain : undefined;
            resourceInputs["ntHost"] = state ? state.ntHost : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["platformDefault"] = state ? state.platformDefault : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["redirectToHosts"] = state ? state.redirectToHosts : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ProxyArgs | undefined;
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["ntDomain"] = args ? args.ntDomain : undefined;
            resourceInputs["ntHost"] = args ? args.ntHost : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["platformDefault"] = args ? args.platformDefault : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["redirectToHosts"] = args ? args.redirectToHosts : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Proxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Proxy resources.
 */
export interface ProxyState {
    /**
     * The name of the proxy host.
     */
    host?: pulumi.Input<string>;
    /**
     * The unique ID of the proxy.
     */
    key?: pulumi.Input<string>;
    /**
     * The proxy domain/realm name.
     */
    ntDomain?: pulumi.Input<string>;
    /**
     * The computer name of the machine (the machine connecting to the NTLM proxy).
     */
    ntHost?: pulumi.Input<string>;
    /**
     * The proxy password when authentication credentials are required.
     */
    password?: pulumi.Input<string>;
    /**
     * When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
     */
    platformDefault?: pulumi.Input<boolean>;
    /**
     * The proxy port number.
     */
    port?: pulumi.Input<number>;
    /**
     * An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
     */
    redirectToHosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The proxy username when authentication credentials are required.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Proxy resource.
 */
export interface ProxyArgs {
    /**
     * The name of the proxy host.
     */
    host: pulumi.Input<string>;
    /**
     * The unique ID of the proxy.
     */
    key: pulumi.Input<string>;
    /**
     * The proxy domain/realm name.
     */
    ntDomain?: pulumi.Input<string>;
    /**
     * The computer name of the machine (the machine connecting to the NTLM proxy).
     */
    ntHost?: pulumi.Input<string>;
    /**
     * The proxy password when authentication credentials are required.
     */
    password?: pulumi.Input<string>;
    /**
     * When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
     */
    platformDefault?: pulumi.Input<boolean>;
    /**
     * The proxy port number.
     */
    port: pulumi.Input<number>;
    /**
     * An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
     */
    redirectToHosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The proxy username when authentication credentials are required.
     */
    username?: pulumi.Input<string>;
}
