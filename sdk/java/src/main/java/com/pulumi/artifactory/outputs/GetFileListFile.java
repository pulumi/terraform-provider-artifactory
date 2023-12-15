// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetFileListFileMetadataTimestamps;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetFileListFile {
    /**
     * @return Is this a folder
     * 
     */
    private Boolean folder;
    /**
     * @return Last modified time
     * 
     */
    private String lastModified;
    /**
     * @return File metadata
     * 
     */
    private GetFileListFileMetadataTimestamps metadataTimestamps;
    /**
     * @return SHA-1 checksum
     * 
     */
    private String sha1;
    /**
     * @return SHA-256 checksum
     * 
     */
    private String sha2;
    /**
     * @return File size in bytes
     * 
     */
    private Integer size;
    /**
     * @return URL to file
     * 
     */
    private String uri;

    private GetFileListFile() {}
    /**
     * @return Is this a folder
     * 
     */
    public Boolean folder() {
        return this.folder;
    }
    /**
     * @return Last modified time
     * 
     */
    public String lastModified() {
        return this.lastModified;
    }
    /**
     * @return File metadata
     * 
     */
    public GetFileListFileMetadataTimestamps metadataTimestamps() {
        return this.metadataTimestamps;
    }
    /**
     * @return SHA-1 checksum
     * 
     */
    public String sha1() {
        return this.sha1;
    }
    /**
     * @return SHA-256 checksum
     * 
     */
    public String sha2() {
        return this.sha2;
    }
    /**
     * @return File size in bytes
     * 
     */
    public Integer size() {
        return this.size;
    }
    /**
     * @return URL to file
     * 
     */
    public String uri() {
        return this.uri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFileListFile defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean folder;
        private String lastModified;
        private GetFileListFileMetadataTimestamps metadataTimestamps;
        private String sha1;
        private String sha2;
        private Integer size;
        private String uri;
        public Builder() {}
        public Builder(GetFileListFile defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.folder = defaults.folder;
    	      this.lastModified = defaults.lastModified;
    	      this.metadataTimestamps = defaults.metadataTimestamps;
    	      this.sha1 = defaults.sha1;
    	      this.sha2 = defaults.sha2;
    	      this.size = defaults.size;
    	      this.uri = defaults.uri;
        }

        @CustomType.Setter
        public Builder folder(Boolean folder) {
            this.folder = Objects.requireNonNull(folder);
            return this;
        }
        @CustomType.Setter
        public Builder lastModified(String lastModified) {
            this.lastModified = Objects.requireNonNull(lastModified);
            return this;
        }
        @CustomType.Setter
        public Builder metadataTimestamps(GetFileListFileMetadataTimestamps metadataTimestamps) {
            this.metadataTimestamps = Objects.requireNonNull(metadataTimestamps);
            return this;
        }
        @CustomType.Setter
        public Builder sha1(String sha1) {
            this.sha1 = Objects.requireNonNull(sha1);
            return this;
        }
        @CustomType.Setter
        public Builder sha2(String sha2) {
            this.sha2 = Objects.requireNonNull(sha2);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            this.uri = Objects.requireNonNull(uri);
            return this;
        }
        public GetFileListFile build() {
            final var _resultValue = new GetFileListFile();
            _resultValue.folder = folder;
            _resultValue.lastModified = lastModified;
            _resultValue.metadataTimestamps = metadataTimestamps;
            _resultValue.sha1 = sha1;
            _resultValue.sha2 = sha2;
            _resultValue.size = size;
            _resultValue.uri = uri;
            return _resultValue;
        }
    }
}
