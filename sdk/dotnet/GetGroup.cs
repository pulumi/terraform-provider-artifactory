// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetGroup
    {
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("artifactory:index/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("artifactory:index/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("adminPrivileges")]
        public bool? AdminPrivileges { get; set; }

        [Input("autoJoin")]
        public bool? AutoJoin { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("externalId")]
        public string? ExternalId { get; set; }

        [Input("includeUsers")]
        public string? IncludeUsers { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("policyManager")]
        public bool? PolicyManager { get; set; }

        [Input("realm")]
        public string? Realm { get; set; }

        [Input("realmAttributes")]
        public string? RealmAttributes { get; set; }

        [Input("reportsManager")]
        public bool? ReportsManager { get; set; }

        [Input("usersNames")]
        private List<string>? _usersNames;
        public List<string> UsersNames
        {
            get => _usersNames ?? (_usersNames = new List<string>());
            set => _usersNames = value;
        }

        [Input("watchManager")]
        public bool? WatchManager { get; set; }

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("adminPrivileges")]
        public Input<bool>? AdminPrivileges { get; set; }

        [Input("autoJoin")]
        public Input<bool>? AutoJoin { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("includeUsers")]
        public Input<string>? IncludeUsers { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("policyManager")]
        public Input<bool>? PolicyManager { get; set; }

        [Input("realm")]
        public Input<string>? Realm { get; set; }

        [Input("realmAttributes")]
        public Input<string>? RealmAttributes { get; set; }

        [Input("reportsManager")]
        public Input<bool>? ReportsManager { get; set; }

        [Input("usersNames")]
        private InputList<string>? _usersNames;
        public InputList<string> UsersNames
        {
            get => _usersNames ?? (_usersNames = new InputList<string>());
            set => _usersNames = value;
        }

        [Input("watchManager")]
        public Input<bool>? WatchManager { get; set; }

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        public readonly bool AdminPrivileges;
        public readonly bool AutoJoin;
        public readonly string? Description;
        public readonly string? ExternalId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludeUsers;
        public readonly string Name;
        public readonly bool? PolicyManager;
        public readonly string Realm;
        public readonly string? RealmAttributes;
        public readonly bool? ReportsManager;
        public readonly ImmutableArray<string> UsersNames;
        public readonly bool? WatchManager;

        [OutputConstructor]
        private GetGroupResult(
            bool adminPrivileges,

            bool autoJoin,

            string? description,

            string? externalId,

            string id,

            string? includeUsers,

            string name,

            bool? policyManager,

            string realm,

            string? realmAttributes,

            bool? reportsManager,

            ImmutableArray<string> usersNames,

            bool? watchManager)
        {
            AdminPrivileges = adminPrivileges;
            AutoJoin = autoJoin;
            Description = description;
            ExternalId = externalId;
            Id = id;
            IncludeUsers = includeUsers;
            Name = name;
            PolicyManager = policyManager;
            Realm = realm;
            RealmAttributes = realmAttributes;
            ReportsManager = reportsManager;
            UsersNames = usersNames;
            WatchManager = watchManager;
        }
    }
}
