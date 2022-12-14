FROM s2i-example
LABEL "io.openshift.s2i.build.source-location"="/home/student/s2i-example/test/test-app" \
      "io.k8s.display-name"="s2i-test" \
      "io.openshift.s2i.build.image"="s2i-example" \
      "io.openshift.s2i.build.commit.author"="builder <build@localhost>" \
      "io.openshift.s2i.build.commit.date"="Thu Jul 14 06:25:39 2022 -0400" \
      "io.openshift.s2i.build.commit.id"="166f7621ef67577a49a0b3d829ca0ce2682c6296" \
      "io.openshift.s2i.build.commit.ref"="master" \
      "io.openshift.s2i.build.commit.message"="Sample commit"

USER root
# Copying in source code
COPY upload/src /tmp/src
# Change file ownership to the assemble user. Builder image must support chown command.
RUN chown -R 1001:0 /tmp/src
USER 1001
# Assemble script sourced from builder image based on user input or image metadata.
# If this file does not exist in the image, the build will fail.
RUN /usr/libexec/s2i/assemble
# Run script sourced from builder image based on user input or image metadata.
# If this file does not exist in the image, the build will fail.
CMD /usr/libexec/s2i/run
