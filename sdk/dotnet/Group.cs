// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    [ArtifactoryResourceType("artifactory:index/group:Group")]
    public partial class Group : Pulumi.CustomResource
    {
        [Output("adminPrivileges")]
        public Output<bool> AdminPrivileges { get; private set; } = null!;

        [Output("autoJoin")]
        public Output<bool> AutoJoin { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("detachAllUsers")]
        public Output<bool?> DetachAllUsers { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
        /// 'false'.
        /// </summary>
        [Output("policyManager")]
        public Output<bool?> PolicyManager { get; private set; } = null!;

        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        [Output("realmAttributes")]
        public Output<string?> RealmAttributes { get; private set; } = null!;

        /// <summary>
        /// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        /// </summary>
        [Output("reportsManager")]
        public Output<bool?> ReportsManager { get; private set; } = null!;

        [Output("usersNames")]
        public Output<ImmutableArray<string>> UsersNames { get; private set; } = null!;

        /// <summary>
        /// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
        /// 'false'.
        /// </summary>
        [Output("watchManager")]
        public Output<bool?> WatchManager { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("artifactory:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/group:Group", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        [Input("adminPrivileges")]
        public Input<bool>? AdminPrivileges { get; set; }

        [Input("autoJoin")]
        public Input<bool>? AutoJoin { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("detachAllUsers")]
        public Input<bool>? DetachAllUsers { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
        /// 'false'.
        /// </summary>
        [Input("policyManager")]
        public Input<bool>? PolicyManager { get; set; }

        [Input("realm")]
        public Input<string>? Realm { get; set; }

        [Input("realmAttributes")]
        public Input<string>? RealmAttributes { get; set; }

        /// <summary>
        /// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        /// </summary>
        [Input("reportsManager")]
        public Input<bool>? ReportsManager { get; set; }

        [Input("usersNames")]
        private InputList<string>? _usersNames;
        public InputList<string> UsersNames
        {
            get => _usersNames ?? (_usersNames = new InputList<string>());
            set => _usersNames = value;
        }

        /// <summary>
        /// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
        /// 'false'.
        /// </summary>
        [Input("watchManager")]
        public Input<bool>? WatchManager { get; set; }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        [Input("adminPrivileges")]
        public Input<bool>? AdminPrivileges { get; set; }

        [Input("autoJoin")]
        public Input<bool>? AutoJoin { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("detachAllUsers")]
        public Input<bool>? DetachAllUsers { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
        /// 'false'.
        /// </summary>
        [Input("policyManager")]
        public Input<bool>? PolicyManager { get; set; }

        [Input("realm")]
        public Input<string>? Realm { get; set; }

        [Input("realmAttributes")]
        public Input<string>? RealmAttributes { get; set; }

        /// <summary>
        /// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        /// </summary>
        [Input("reportsManager")]
        public Input<bool>? ReportsManager { get; set; }

        [Input("usersNames")]
        private InputList<string>? _usersNames;
        public InputList<string> UsersNames
        {
            get => _usersNames ?? (_usersNames = new InputList<string>());
            set => _usersNames = value;
        }

        /// <summary>
        /// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
        /// 'false'.
        /// </summary>
        [Input("watchManager")]
        public Input<bool>? WatchManager { get; set; }

        public GroupState()
        {
        }
    }
}
