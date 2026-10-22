package programs

// CurrentInfraVersion is the version of the Pulumi infrastructure program
// embedded in this controller binary. Bump this integer whenever the
// program code changes in a way that should trigger a re-deploy on
// existing servers. Pulumi's diff engine handles the actual migration.
//
// Version 4 was a BREAKING change: the data disk mount relocates from
// /mnt/disks/minecraft/data to /mnt/disks/minecraft and world data moves into
// a data/ subdirectory alongside the new backup manifest directory. Existing
// servers must be migrated BEFORE re-deploying; see docs/DEPLOYMENT.md.
//
// Version 5 sizes the JVM heap as a percentage of host RAM instead of the
// hardcoded 3G: user-data gains MAX_MEMORY=75%, which implements
// -XX:MaxRAMPercentage. The percentage keeps user-data byte-identical across
// machine types, so machine-type changes remain a resize rather than a
// recreate.
//
// Version 7 fixes provisioning on Container-Optimized OS: the hardcoded
// UID 1000 for the `metio` user collided with COS's `chronos` stub account
// (UID 1000), so cloud-init could not create the user and every systemd
// unit failed with `217/USER`. The user is now auto-allocated; a runcmd
// guard ensures `metio` exists before any unit starts. Existing servers
// must be recreated to pick up the corrected user-data.
//
// Version 6 adds modpack support (ADR-0006): a pack-driven server passes
// MODRINTH_MODPACK instead of VERSION, and the Minecraft version becomes
// pack-controlled. Those servers deploy with an empty minecraftVersion.
//
// Version 8 fixes pack-driven servers that booted unmodded: the Modrinth
// modpack install requires TYPE=MODRINTH, and setting only MODRINTH_MODPACK
// left the image's VANILLA default in place. Pack servers built on version 6
// or 7 must be recreated to pick up the corrected user-data.
//
// Version 9 keeps spark's background profiler off pack servers: spark 1.10.x
// bundles async-profiler, whose native code segfaults on the Java 25 image and
// crash-loops the server (metio#631). Pack servers gain a PRE_START_SCRIPT
// hook that writes config/spark/config.json before the JVM starts; vanilla
// user-data is unchanged, but every server's stored version is older than this
// one. Pack servers created before version 9 must be recreated to pick up the
// hook.
const CurrentInfraVersion = 9
