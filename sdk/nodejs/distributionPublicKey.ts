// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory Distribution Public Key resource. This can be used to create and manage Artifactory Distribution Public Keys.
 *
 * See [API description](https://jfrog.com/help/r/jfrog-rest-apis/set-distributionpublic-gpg-key) in the Artifactory documentation for more details. Also the [UI documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/managing-webstart-and-jar-signing) has further details on where to find these keys in Artifactory.
 *
 * ## Import
 *
 * Distribution Public Key can be imported using the key id, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/distributionPublicKey:DistributionPublicKey my-key keyid
 * ```
 */
export class DistributionPublicKey extends pulumi.CustomResource {
    /**
     * Get an existing DistributionPublicKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributionPublicKeyState, opts?: pulumi.CustomResourceOptions): DistributionPublicKey {
        return new DistributionPublicKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/distributionPublicKey:DistributionPublicKey';

    /**
     * Returns true if the given object is an instance of DistributionPublicKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DistributionPublicKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DistributionPublicKey.__pulumiType;
    }

    /**
     * Will be used as an identifier when uploading/retrieving the public key via REST API.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Returns the computed key fingerprint
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * Returns the name and eMail address of issuer
     */
    public /*out*/ readonly issuedBy!: pulumi.Output<string>;
    /**
     * Returns the date/time when this GPG key was created
     */
    public /*out*/ readonly issuedOn!: pulumi.Output<string>;
    /**
     * Returns the key id by which this key is referenced in Artifactory
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * The Public key to add as a trusted distribution GPG key.
     *
     * The following additional attributes are exported:
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * Returns the date/time when this GPG key expires.
     */
    public /*out*/ readonly validUntil!: pulumi.Output<string>;

    /**
     * Create a DistributionPublicKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributionPublicKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DistributionPublicKeyArgs | DistributionPublicKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DistributionPublicKeyState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["issuedBy"] = state ? state.issuedBy : undefined;
            resourceInputs["issuedOn"] = state ? state.issuedOn : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["validUntil"] = state ? state.validUntil : undefined;
        } else {
            const args = argsOrState as DistributionPublicKeyArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["issuedBy"] = undefined /*out*/;
            resourceInputs["issuedOn"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
            resourceInputs["validUntil"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DistributionPublicKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DistributionPublicKey resources.
 */
export interface DistributionPublicKeyState {
    /**
     * Will be used as an identifier when uploading/retrieving the public key via REST API.
     */
    alias?: pulumi.Input<string>;
    /**
     * Returns the computed key fingerprint
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * Returns the name and eMail address of issuer
     */
    issuedBy?: pulumi.Input<string>;
    /**
     * Returns the date/time when this GPG key was created
     */
    issuedOn?: pulumi.Input<string>;
    /**
     * Returns the key id by which this key is referenced in Artifactory
     */
    keyId?: pulumi.Input<string>;
    /**
     * The Public key to add as a trusted distribution GPG key.
     *
     * The following additional attributes are exported:
     */
    publicKey?: pulumi.Input<string>;
    /**
     * Returns the date/time when this GPG key expires.
     */
    validUntil?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DistributionPublicKey resource.
 */
export interface DistributionPublicKeyArgs {
    /**
     * Will be used as an identifier when uploading/retrieving the public key via REST API.
     */
    alias: pulumi.Input<string>;
    /**
     * The Public key to add as a trusted distribution GPG key.
     *
     * The following additional attributes are exported:
     */
    publicKey: pulumi.Input<string>;
}
