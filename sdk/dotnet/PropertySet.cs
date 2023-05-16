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
    /// Provides an Artifactory Property Set resource.
    /// This resource configuration corresponds to 'propertySets' config block in system configuration XML
    /// (REST endpoint: artifactory/api/system/configuration).
    /// 
    /// ~&gt;The `artifactory.PropertySet` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Artifactory.PropertySet("foo", new()
    ///     {
    ///         Properties = new[]
    ///         {
    ///             new Artifactory.Inputs.PropertySetPropertyArgs
    ///             {
    ///                 ClosedPredefinedValues = true,
    ///                 MultipleChoice = true,
    ///                 Name = "set1property1",
    ///                 PredefinedValues = new[]
    ///                 {
    ///                     new Artifactory.Inputs.PropertySetPropertyPredefinedValueArgs
    ///                     {
    ///                         DefaultValue = true,
    ///                         Name = "passed-QA",
    ///                     },
    ///                     new Artifactory.Inputs.PropertySetPropertyPredefinedValueArgs
    ///                     {
    ///                         DefaultValue = false,
    ///                         Name = "failed-QA",
    ///                     },
    ///                 },
    ///             },
    ///             new Artifactory.Inputs.PropertySetPropertyArgs
    ///             {
    ///                 ClosedPredefinedValues = false,
    ///                 MultipleChoice = false,
    ///                 Name = "set1property2",
    ///                 PredefinedValues = new[]
    ///                 {
    ///                     new Artifactory.Inputs.PropertySetPropertyPredefinedValueArgs
    ///                     {
    ///                         DefaultValue = true,
    ///                         Name = "passed-QA",
    ///                     },
    ///                     new Artifactory.Inputs.PropertySetPropertyPredefinedValueArgs
    ///                     {
    ///                         DefaultValue = false,
    ///                         Name = "failed-QA",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Visible = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Current Property Set can be imported using `property-set1` as the `ID`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/propertySet:PropertySet foo property-set1
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/propertySet:PropertySet")]
    public partial class PropertySet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Predefined property name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of properties that will be part of the property set.
        /// </summary>
        [Output("properties")]
        public Output<ImmutableArray<Outputs.PropertySetProperty>> Properties { get; private set; } = null!;

        /// <summary>
        /// Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
        /// </summary>
        [Output("visible")]
        public Output<bool?> Visible { get; private set; } = null!;


        /// <summary>
        /// Create a PropertySet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PropertySet(string name, PropertySetArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/propertySet:PropertySet", name, args ?? new PropertySetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PropertySet(string name, Input<string> id, PropertySetState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/propertySet:PropertySet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PropertySet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PropertySet Get(string name, Input<string> id, PropertySetState? state = null, CustomResourceOptions? options = null)
        {
            return new PropertySet(name, id, state, options);
        }
    }

    public sealed class PropertySetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Predefined property name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties", required: true)]
        private InputList<Inputs.PropertySetPropertyArgs>? _properties;

        /// <summary>
        /// A list of properties that will be part of the property set.
        /// </summary>
        public InputList<Inputs.PropertySetPropertyArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.PropertySetPropertyArgs>());
            set => _properties = value;
        }

        /// <summary>
        /// Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
        /// </summary>
        [Input("visible")]
        public Input<bool>? Visible { get; set; }

        public PropertySetArgs()
        {
        }
        public static new PropertySetArgs Empty => new PropertySetArgs();
    }

    public sealed class PropertySetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Predefined property name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        private InputList<Inputs.PropertySetPropertyGetArgs>? _properties;

        /// <summary>
        /// A list of properties that will be part of the property set.
        /// </summary>
        public InputList<Inputs.PropertySetPropertyGetArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.PropertySetPropertyGetArgs>());
            set => _properties = value;
        }

        /// <summary>
        /// Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
        /// </summary>
        [Input("visible")]
        public Input<bool>? Visible { get; set; }

        public PropertySetState()
        {
        }
        public static new PropertySetState Empty => new PropertySetState();
    }
}
