# Modpacks

Metio can boot a server from a published [Modrinth](https://modrinth.com/modpacks) modpack
instead of a plain vanilla server. The pack decides the Minecraft version, the mod loader
(Fabric, NeoForge, Forge, Quilt) and the mod set, so you pick a pack once and the server
installs it on first boot.

## Supported pack source

Only Modrinth packs are supported. CurseForge is intentionally unsupported: the Minecraft
container authenticates to CurseForge with a `CF_API_KEY` environment variable, and the
container's environment is part of the instance template — the key would be readable from
the VM's instance metadata by anything that can read the instance (for example
`compute.instances.get` in the project). Modrinth needs no credentials, so pack servers
carry no pack-source secret. See
[ADR-0006](adr/0006-modpack-support-and-configurable-memory.md) for the full rationale.

## Choosing a pack when creating a server

The setup wizard has a dedicated **Modpack** step:

1. **Basic Info** — name, region and zone.
2. **Modpack** — search Modrinth, select a pack, then pick a **Pack Version**.
   **Latest** installs the newest version of the pack; choosing a specific version pins it.
3. **Server Specs** — machine type and disk size. The **Minecraft Version** field is
   **disabled while a pack is selected**, because the pack manifest controls the version.
   Clear the pack selection to configure a vanilla server, or remove the pack later from the
   server settings.
4. **Settings** and **Review** — the review page shows the pack and prints the Minecraft
   version as *Pack-controlled*.

Leaving the Modpack step empty creates a vanilla server, exactly as before.

## What happens during provisioning

A pack server is provisioned like any other server; the pack itself is installed by the
Minecraft container on first boot. Because that download is large and long-running, the
provisioning page shows an extra step:

| Step | Meaning |
|------|---------|
| `upsert_stack` | Pulumi stack created in the state bucket |
| `deploy_infrastructure` | VM, disks, firewall rules and IAM created |
| `install_modpack` | The container is downloading and installing the pack, then starting the server |

The `install_modpack` step is only present for pack-driven servers. It completes when the
server reports running with its player slots configured — not on a fixed timer — so a slow
download shows as a long-running step rather than a failure.

First boot for a large pack (hundreds of mods, several hundred MiB) can take several
minutes. The controller waits up to `pack_install_timeout` (**default `20m`**) for the
install to finish. Exceeding that window is **not fatal**: the step is marked complete with
a warning, the VM keeps installing, and the dashboard reflects the real state once the
server comes up. Raise the timeout for very large packs (see
[Troubleshooting](#troubleshooting)).

## Changing or removing a pack later

Open the server's **Settings → Modpack** tab. It uses the same picker as the wizard and
stages the change; nothing happens until you submit:

- **Change** the pack or pin a different pack version.
- **Remove** the pack and return to vanilla. You must then choose a Minecraft version,
  because a vanilla server is not pack-controlled.
- **Leave it alone** — the pack is only touched when you change the selection.

> **Changing a pack recreates the VM.** The pack is part of the instance template
> (`TYPE`, `MODRINTH_MODPACK` and optionally `MODRINTH_VERSION` in `user-data`), so Metio
> destroys and recreates the machine rather than restarting it. Expect the provisioning
> page to run through the steps again, including `install_modpack`, and expect a few
> minutes of downtime.
>
> **Your world survives.** The world, player data and server configuration live on the
> persistent data disk, which is detached from the VM and reattached to the replacement.
> Backups continue to work across the change.

The settings dialog spells this out before you confirm the change.

## Sizing a pack server

The JVM heap is sized as a percentage of the selected machine's RAM: the container runs
with `MAX_MEMORY=75%`, which becomes `-XX:MaxRAMPercentage=75`. The remaining ~25% is left
to the operating system, the machine-agent, the Docker daemon and the `mc-backup` sidecar.
Heap therefore scales automatically with the machine type, and changing the machine type
stays a resize rather than a full rebuild.

| Machine type | RAM | Approximate max heap |
|--------------|-----|----------------------|
| `e2-small` | 2 GB | ~1.5 GB |
| `e2-medium` | 4 GB | ~3 GB |
| `e2-standard-2` | 8 GB | ~6 GB |
| `e2-standard-4` | 16 GB | ~12 GB |
| `e2-standard-8` | 32 GB | ~24 GB |
| `n2-highmem-4` | 32 GB | ~24 GB |
| `n2-highmem-8` | 64 GB | ~48 GB |

Modpacks routinely want 6–10 GB of heap, so start at `e2-standard-4` (or an `n2-highmem`
type when the pack is CPU-heavy) rather than an `e2-small`. An undersized machine shows up
as out-of-memory kills or a server that lags badly with the pack loaded.

The effective heap is not currently displayed in the UI; use the table above to estimate it.

## Known limitations

- **CurseForge packs are unsupported** by design, for the credential reason above.
- **`spark` profiling is limited to the safe sampling engine.** Large packs bundle
  `spark` 1.10.x, whose async profiler kills the JVM (`SIGSEGV` in
  `VMThread::nativeThreadId`) on the Java 25 image. Metio therefore disables spark's
  background profiler on pack servers before the server starts
  ([#631](https://github.com/metio-plattform/metio/issues/631)). `spark` itself stays
  installed: `/spark tps` and `/spark health` work, and manual profiling works with
  `/spark profiler --force-java-sampler`. Without that flag spark loads async-profiler
  and can crash the server.
- **Very large packs and the default wait.** A pack that takes longer than
  `pack_install_timeout` still finishes installing, but the provisioning view stops
  following it. Raise the timeout if this is normal for your pack.

## Known pack compatibility issues

Pack compatibility is tracked by pack, version and loader — not by individual mod. A mod
that crashes one pack's server can work on a client or in another pack, so this is not a
list of "broken mods". Update or remove an entry when a later pack version fixes the issue.

| Pack | Version | Loader | Symptoms | Workaround |
|------|---------|--------|----------|------------|
| [Better MC BMC4](https://modrinth.com/modpack/better-mc-forge-bmc4) | v59 (`bdRIjRgs`) | Forge 47.4.20 / Minecraft 1.20.1 | `HeadlessException`; missing mandatory dependencies; client-class / `ClassCastException` crash | Use a server-ready pack or a newer pack version |

**Better MC [FORGE] BMC4 v59** is not server-ready as packaged. Its `modrinth.index.json`
marks all 340 entries as `server: required`, so the installer includes client-only mods.
The server first fails in Missing Mods Checker (`HeadlessException`), then on four mandatory
dependencies the pack omits (`temporalapi`, `connectormod`, `flib`, `twilightforest`), and
then on a client-class / `ClassCastException` failure in `particle_effects`. Adding or
removing individual mods is not a supported fix; use a server-ready pack or a newer version.

*Last verified: 2026-10-24. Upstream support: [LunaPixel Discord](https://discord.gg/LunaPixel).*

### Diagnosing a pack crash loop

A crash-looping pack can still show the container as running. Check the restart count and
the server log:

```bash
systemctl show minecraft -p NRestarts -p ActiveState -p SubState
sudo journalctl -u minecraft -n 100 --no-pager
sudo tail -n 100 /mnt/disks/minecraft/data/logs/latest.log
sudo ls /mnt/disks/minecraft/data/crash-reports/
```

`docker logs` is not reliable here: the unit runs the container with `--rm`, so the previous
attempt's output is gone by the time systemd restarts it.

Typical causes in the log:

- `HeadlessException` or `MissingModsChecker` — a client-only GUI mod is installed.
- `Missing or unsupported mandatory dependencies` — the pack omitted a required mod.
- `Attempted to load class net/minecraft/client/... for invalid dist DEDICATED_SERVER` or
  `ClassCastException` — a client-only mod or an incompatible Connector mix.

Metio does not surface crash loops in the dashboard yet; see
[#637](https://github.com/metio-plattform/metio/issues/637).

## Troubleshooting

**The install step runs for a long time.** That is expected for large packs. Follow the
container logs on the VM to see download progress:

```bash
gcloud compute ssh --zone "<zone>" <instance-name> --project "<project-id>"
sudo journalctl -u minecraft -f
```

**The step completed but the pack is not loaded.** Check that the container is running and
that mods were written to the data disk:

```bash
sudo docker ps
sudo ls /mnt/disks/minecraft/data/mods | head
```

If the server boots unmodded, the instance template may be stale — see
[Infrastructure Version Migrations](DEPLOYMENT.md#infrastructure-version-migrations).

**Raise the install wait.** Set `pack_install_timeout` in `metio.auto.tfvars` (for example
`"45m"`) and redeploy the controller, or set `PACK_INSTALL_TIMEOUT` on the controller
service directly. See [Deployment](DEPLOYMENT.md#step-1-configure-infrastructure-variables).
