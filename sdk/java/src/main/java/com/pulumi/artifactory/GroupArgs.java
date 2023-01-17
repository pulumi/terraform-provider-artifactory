// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupArgs Empty = new GroupArgs();

    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     * 
     */
    @Import(name="adminPrivileges")
    private @Nullable Output<Boolean> adminPrivileges;

    /**
     * @return Any users added to this group will automatically be assigned with admin privileges in the system.
     * 
     */
    public Optional<Output<Boolean>> adminPrivileges() {
        return Optional.ofNullable(this.adminPrivileges);
    }

    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     * 
     */
    @Import(name="autoJoin")
    private @Nullable Output<Boolean> autoJoin;

    /**
     * @return When this parameter is set, any new users defined in the system are automatically assigned to this group.
     * 
     */
    public Optional<Output<Boolean>> autoJoin() {
        return Optional.ofNullable(this.autoJoin);
    }

    /**
     * A description for the group
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the group
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When this is set to `true`, an empty or missing usernames array will detach all users from the group
     * 
     */
    @Import(name="detachAllUsers")
    private @Nullable Output<Boolean> detachAllUsers;

    /**
     * @return When this is set to `true`, an empty or missing usernames array will detach all users from the group
     * 
     */
    public Optional<Output<Boolean>> detachAllUsers() {
        return Optional.ofNullable(this.detachAllUsers);
    }

    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return New external group ID used to configure the corresponding group in Azure AD.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * Name of the group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     * 
     */
    @Import(name="policyManager")
    private @Nullable Output<Boolean> policyManager;

    /**
     * @return When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> policyManager() {
        return Optional.ofNullable(this.policyManager);
    }

    /**
     * The realm for the group.
     * 
     */
    @Import(name="realm")
    private @Nullable Output<String> realm;

    /**
     * @return The realm for the group.
     * 
     */
    public Optional<Output<String>> realm() {
        return Optional.ofNullable(this.realm);
    }

    /**
     * The realm attributes for the group.
     * 
     */
    @Import(name="realmAttributes")
    private @Nullable Output<String> realmAttributes;

    /**
     * @return The realm attributes for the group.
     * 
     */
    public Optional<Output<String>> realmAttributes() {
        return Optional.ofNullable(this.realmAttributes);
    }

    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     * 
     */
    @Import(name="reportsManager")
    private @Nullable Output<Boolean> reportsManager;

    /**
     * @return When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> reportsManager() {
        return Optional.ofNullable(this.reportsManager);
    }

    @Import(name="usersNames")
    private @Nullable Output<List<String>> usersNames;

    public Optional<Output<List<String>>> usersNames() {
        return Optional.ofNullable(this.usersNames);
    }

    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     * 
     */
    @Import(name="watchManager")
    private @Nullable Output<Boolean> watchManager;

    /**
     * @return When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> watchManager() {
        return Optional.ofNullable(this.watchManager);
    }

    private GroupArgs() {}

    private GroupArgs(GroupArgs $) {
        this.adminPrivileges = $.adminPrivileges;
        this.autoJoin = $.autoJoin;
        this.description = $.description;
        this.detachAllUsers = $.detachAllUsers;
        this.externalId = $.externalId;
        this.name = $.name;
        this.policyManager = $.policyManager;
        this.realm = $.realm;
        this.realmAttributes = $.realmAttributes;
        this.reportsManager = $.reportsManager;
        this.usersNames = $.usersNames;
        this.watchManager = $.watchManager;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupArgs $;

        public Builder() {
            $ = new GroupArgs();
        }

        public Builder(GroupArgs defaults) {
            $ = new GroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminPrivileges Any users added to this group will automatically be assigned with admin privileges in the system.
         * 
         * @return builder
         * 
         */
        public Builder adminPrivileges(@Nullable Output<Boolean> adminPrivileges) {
            $.adminPrivileges = adminPrivileges;
            return this;
        }

        /**
         * @param adminPrivileges Any users added to this group will automatically be assigned with admin privileges in the system.
         * 
         * @return builder
         * 
         */
        public Builder adminPrivileges(Boolean adminPrivileges) {
            return adminPrivileges(Output.of(adminPrivileges));
        }

        /**
         * @param autoJoin When this parameter is set, any new users defined in the system are automatically assigned to this group.
         * 
         * @return builder
         * 
         */
        public Builder autoJoin(@Nullable Output<Boolean> autoJoin) {
            $.autoJoin = autoJoin;
            return this;
        }

        /**
         * @param autoJoin When this parameter is set, any new users defined in the system are automatically assigned to this group.
         * 
         * @return builder
         * 
         */
        public Builder autoJoin(Boolean autoJoin) {
            return autoJoin(Output.of(autoJoin));
        }

        /**
         * @param description A description for the group
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the group
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param detachAllUsers When this is set to `true`, an empty or missing usernames array will detach all users from the group
         * 
         * @return builder
         * 
         */
        public Builder detachAllUsers(@Nullable Output<Boolean> detachAllUsers) {
            $.detachAllUsers = detachAllUsers;
            return this;
        }

        /**
         * @param detachAllUsers When this is set to `true`, an empty or missing usernames array will detach all users from the group
         * 
         * @return builder
         * 
         */
        public Builder detachAllUsers(Boolean detachAllUsers) {
            return detachAllUsers(Output.of(detachAllUsers));
        }

        /**
         * @param externalId New external group ID used to configure the corresponding group in Azure AD.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId New external group ID used to configure the corresponding group in Azure AD.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param name Name of the group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policyManager When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder policyManager(@Nullable Output<Boolean> policyManager) {
            $.policyManager = policyManager;
            return this;
        }

        /**
         * @param policyManager When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder policyManager(Boolean policyManager) {
            return policyManager(Output.of(policyManager));
        }

        /**
         * @param realm The realm for the group.
         * 
         * @return builder
         * 
         */
        public Builder realm(@Nullable Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The realm for the group.
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        /**
         * @param realmAttributes The realm attributes for the group.
         * 
         * @return builder
         * 
         */
        public Builder realmAttributes(@Nullable Output<String> realmAttributes) {
            $.realmAttributes = realmAttributes;
            return this;
        }

        /**
         * @param realmAttributes The realm attributes for the group.
         * 
         * @return builder
         * 
         */
        public Builder realmAttributes(String realmAttributes) {
            return realmAttributes(Output.of(realmAttributes));
        }

        /**
         * @param reportsManager When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder reportsManager(@Nullable Output<Boolean> reportsManager) {
            $.reportsManager = reportsManager;
            return this;
        }

        /**
         * @param reportsManager When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder reportsManager(Boolean reportsManager) {
            return reportsManager(Output.of(reportsManager));
        }

        public Builder usersNames(@Nullable Output<List<String>> usersNames) {
            $.usersNames = usersNames;
            return this;
        }

        public Builder usersNames(List<String> usersNames) {
            return usersNames(Output.of(usersNames));
        }

        public Builder usersNames(String... usersNames) {
            return usersNames(List.of(usersNames));
        }

        /**
         * @param watchManager When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder watchManager(@Nullable Output<Boolean> watchManager) {
            $.watchManager = watchManager;
            return this;
        }

        /**
         * @param watchManager When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder watchManager(Boolean watchManager) {
            return watchManager(Output.of(watchManager));
        }

        public GroupArgs build() {
            return $;
        }
    }

}
