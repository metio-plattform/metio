# Changelog

## [3.0.0](https://github.com/metio-plattform/metio/compare/v2.2.0...v3.0.0) (2026-10-26)


### ⚠ BREAKING CHANGES

* **backup:** Existing server VMs must run the disk migration in docs/DEPLOYMENT.md before deploying; the data disk mount relocates to /mnt/disks/minecraft with a data/ subdirectory and CurrentInfraVersion is bumped to 4.

### Features

* abstract async provisioning behind OperationExecutor interface ([9d57f24](https://github.com/metio-plattform/metio/commit/9d57f24408e88a0e59647889d90683304e170246))
* adapt server dashboard to ticket layout ([442bc03](https://github.com/metio-plattform/metio/commit/442bc034f960a1a5c3ec3e2177d49a01caccc080))
* add custom files-updater plugin to auto-update variables.tf on release ([8fb813c](https://github.com/metio-plattform/metio/commit/8fb813c7a219751f7f6d5ed70322609fecfe0460))
* add Dapr sidecar infrastructure for Cloud Run controller ([#255](https://github.com/metio-plattform/metio/issues/255)) ([66e1621](https://github.com/metio-plattform/metio/commit/66e162134596bb11e7a73cf1927fc135c617d876))
* add Dapr sidecar infrastructure for Cloud Run controller ([#255](https://github.com/metio-plattform/metio/issues/255)) ([519c0a3](https://github.com/metio-plattform/metio/commit/519c0a32a8b8a297fc80adbdf3ed3aca55303740))
* add DB_BACKEND toggle for Dapr state store selection ([#257](https://github.com/metio-plattform/metio/issues/257)) ([2fa09d8](https://github.com/metio-plattform/metio/commit/2fa09d8099ed8cb9146843e543933d3b4cb18eac))
* add DB_BACKEND toggle for Dapr state store selection ([#257](https://github.com/metio-plattform/metio/issues/257)) ([6020911](https://github.com/metio-plattform/metio/commit/60209112a04069152bc379a74db0b4da83427dd8))
* add DEV_API_KEY bearer token bypass for local API testing ([c54f74a](https://github.com/metio-plattform/metio/commit/c54f74a0c5b9ac27d220f4bdd51dcf43327eb810))
* add devcontainer with Go, Node, Docker, GCloud, OpenTofu tooling ([cb23c2c](https://github.com/metio-plattform/metio/commit/cb23c2c320d7d21345a89794c6cfb7e81aa58ace))
* add devcontainer with Go, Node, Docker, GCloud, OpenTofu tooling ([1f8ff1e](https://github.com/metio-plattform/metio/commit/1f8ff1efa1d28bb0b276b596d47e167916db44e5))
* add devcontainer with Go, Node, Docker, GCloud, OpenTofu tooling ([8e3dd17](https://github.com/metio-plattform/metio/commit/8e3dd174475bd3b8783d1a9b3df05828f9dd0abe))
* add devcontainer with Go, Node, Docker, GCloud, OpenTofu tooling ([f5ae0c1](https://github.com/metio-plattform/metio/commit/f5ae0c1350bab1eca9a84a42c25957f779b65dec))
* add JSON tags and Dapr key schema for state store migration ([#258](https://github.com/metio-plattform/metio/issues/258)) ([eadc24d](https://github.com/metio-plattform/metio/commit/eadc24d27f652b011612365c5d28e5e6ba2aaf4f))
* add JSON tags and Dapr key schema for state store migration ([#258](https://github.com/metio-plattform/metio/issues/258)) ([6621c5a](https://github.com/metio-plattform/metio/commit/6621c5af17edc38fcadcbb2c74f753480e1b0a24))
* add modpack data model and cloud-config ([#580](https://github.com/metio-plattform/metio/issues/580)) ([2331f9a](https://github.com/metio-plattform/metio/commit/2331f9a12a8d7c4aa64cc47e8acbaefdd2ec21df)), closes [#536](https://github.com/metio-plattform/metio/issues/536)
* add modpack picker to server setup wizard ([#581](https://github.com/metio-plattform/metio/issues/581)) ([54a9210](https://github.com/metio-plattform/metio/commit/54a9210444afd714d77b6cdc0a8646f027658ffb)), closes [#537](https://github.com/metio-plattform/metio/issues/537)
* add modpack search API endpoints ([#579](https://github.com/metio-plattform/metio/issues/579)) ([fbfcf1a](https://github.com/metio-plattform/metio/commit/fbfcf1a6539a867416f4e173a3c43fdfbd883b4c)), closes [#535](https://github.com/metio-plattform/metio/issues/535)
* add Modrinth modpack search service ([#577](https://github.com/metio-plattform/metio/issues/577)) ([c5a0680](https://github.com/metio-plattform/metio/commit/c5a06801ace42a76148eeda4b2217c050d40a81e)), closes [#534](https://github.com/metio-plattform/metio/issues/534)
* add more GCP APIs to pre-deployment validation ([#627](https://github.com/metio-plattform/metio/issues/627)) ([46eb216](https://github.com/metio-plattform/metio/commit/46eb216a6272bff4a6f4606663644b44b8f8638b))
* add opencode-web dev environment target ([#578](https://github.com/metio-plattform/metio/issues/578)) ([9711209](https://github.com/metio-plattform/metio/commit/9711209878a041dbe1e732a660d5ebc5de485ffe))
* add postgres_mode variable and mount connection-string secret ([#325](https://github.com/metio-plattform/metio/issues/325)) ([c09390b](https://github.com/metio-plattform/metio/commit/c09390b4808659165726e3cb51a69e3ae3e2b521))
* add postgres_mode variable and mount connection-string secret ([#325](https://github.com/metio-plattform/metio/issues/325)) ([2a101a5](https://github.com/metio-plattform/metio/commit/2a101a56e4000fc28f3ca527f56ceb51fd2a8013))
* add project validation and permission checks endpoint ([c9bca3f](https://github.com/metio-plattform/metio/commit/c9bca3ff534dca49c20fce6afb089b4974d54c67))
* add project validation and permission checks endpoint ([ba67c38](https://github.com/metio-plattform/metio/commit/ba67c38a7d164238e39dfe1ba5f00748ac9f4c3c))
* add promote-distribution workflow for GCP Artifact Registry ([bdd4784](https://github.com/metio-plattform/metio/commit/bdd47847eaceeedcf170c7bd0a6e0c12bafbcc33)), closes [#242](https://github.com/metio-plattform/metio/issues/242)
* add real-time provisioning status indicators ([597fe09](https://github.com/metio-plattform/metio/commit/597fe0952e40dcafc02ce38fa71136a11fad64ef))
* add real-time provisioning status indicators ([2452295](https://github.com/metio-plattform/metio/commit/2452295c44e8defc3c87a03696c56f3bcbc1e2e9))
* add server setup wizard component ([b421474](https://github.com/metio-plattform/metio/commit/b421474eb4ba98b9ce4171025823218458e913c7))
* add server setup wizard component ([6b96552](https://github.com/metio-plattform/metio/commit/6b96552b9c4ad3abc9ce7a3f4524ab59f1bdde88))
* add server update operations with in-place/resize/recreate workflows ([b22126d](https://github.com/metio-plattform/metio/commit/b22126d637d8a3b22311e5251b42274441a33c23))
* add server update operations with in-place/resize/recreate workflows ([6b04aaa](https://github.com/metio-plattform/metio/commit/6b04aaae58c65481932fda9b71b7eac0393a6d64))
* add shipper opencode commands and env config ([be563f8](https://github.com/metio-plattform/metio/commit/be563f8a26278c6a6d6ff63954ee4cb6fce3f7eb))
* add shipper opencode commands and env config ([74b8c59](https://github.com/metio-plattform/metio/commit/74b8c593a3b371afcb1ecda2c1359d9f81c8b432))
* add throwaway VM harness for the ADR-0008 k3s spike ([#549](https://github.com/metio-plattform/metio/issues/549)) ([43c0114](https://github.com/metio-plattform/metio/commit/43c0114f163d2fba3a6fa070b9441b57ea151e4a)), closes [#542](https://github.com/metio-plattform/metio/issues/542)
* add update detection and notification UI ([aa7c303](https://github.com/metio-plattform/metio/commit/aa7c3031bbd0811886a2a92f9883cb8f546c6587))
* add update detection and notification UI ([b29b760](https://github.com/metio-plattform/metio/commit/b29b760c8eb98c633651a94f1369b78ce53e93b3))
* allow users to update machine agent from UI ([#309](https://github.com/metio-plattform/metio/issues/309)) ([a9676ae](https://github.com/metio-plattform/metio/commit/a9676ae69630922257c48b83a8abd820aedaed14))
* allow users to update machine agent from UI ([#309](https://github.com/metio-plattform/metio/issues/309)) ([912c8c2](https://github.com/metio-plattform/metio/commit/912c8c24fd7cfd3e8873436cc40558481732b051))
* **api:** add backup catalog and idempotent report ingestion ([#477](https://github.com/metio-plattform/metio/issues/477)) ([59f7dc8](https://github.com/metio-plattform/metio/commit/59f7dc84b40fe1cd6e58db548bd4a079b9cc77e2)), closes [#429](https://github.com/metio-plattform/metio/issues/429)
* auto-provision Pulumi state bucket on controller startup ([37d4ae1](https://github.com/metio-plattform/metio/commit/37d4ae1152a00258c192c31c1798c2a263e69438))
* auto-provision Pulumi state bucket on controller startup ([fccf32b](https://github.com/metio-plattform/metio/commit/fccf32b545f5e855366113d79d70efde23b71326))
* **backup:** add per-server backup configuration and runtime ([#435](https://github.com/metio-plattform/metio/issues/435)) ([7de6548](https://github.com/metio-plattform/metio/commit/7de65487fa6c789937b54dc314a96fcbbfc25a46))
* **backups:** add backup UI and operational documentation ([#483](https://github.com/metio-plattform/metio/issues/483)) ([079e25f](https://github.com/metio-plattform/metio/commit/079e25f8ef2e5ce521cd73b63748a960b1c50ad0))
* **backups:** create new servers from retained backups ([#482](https://github.com/metio-plattform/metio/issues/482)) ([a0af790](https://github.com/metio-plattform/metio/commit/a0af790a02702bc19c040b4844293620acb91207))
* **backups:** move backup actions to catalog page and add server filter ([#487](https://github.com/metio-plattform/metio/issues/487)) ([df754e6](https://github.com/metio-plattform/metio/commit/df754e6d3c7db9bd149ab9e342babbfedc45b4b9))
* **backups:** preserve deleted-server backups with scheduled cleanup ([#479](https://github.com/metio-plattform/metio/issues/479)) ([3fe6a08](https://github.com/metio-plattform/metio/commit/3fe6a087448b5065dd61071f286c6e6d50fc4e20))
* **backups:** restore snapshots onto existing servers ([#481](https://github.com/metio-plattform/metio/issues/481)) ([1e193cd](https://github.com/metio-plattform/metio/commit/1e193cd1f844ba5a8d1d6446e7302b8dee939834))
* **BYLNET-123:** add scheduled server shutdown functionality ([45b8d86](https://github.com/metio-plattform/metio/commit/45b8d8682f59a5566ddc90f8a75da16bf2ff78ed))
* **BYLNET-123:** add scheduled server shutdown functionality ([9a0a237](https://github.com/metio-plattform/metio/commit/9a0a237ee5569a78766d1a2b84ae5d7f24a9f192))
* **BYLNET-162:** manage whitelist from frontend ([c91af17](https://github.com/metio-plattform/metio/commit/c91af1799a9dbed9ad7dc4f19b1445a17fe1bbbd))
* **BYLNET-162:** Manage whitelist from frontend ([de18179](https://github.com/metio-plattform/metio/commit/de181791d972d1bba48bb6005172470fea73a820))
* **BYLNET-207:** increase backup retention to 3 months ([e31228c](https://github.com/metio-plattform/metio/commit/e31228c770520a7465ae2b648b1510a4a6e3b3e7))
* **BYLNET-207:** Increase backup retention to 3 months ([cff6cdd](https://github.com/metio-plattform/metio/commit/cff6cdd9b87f57ec1f0bb1dfdb1c9cd0ac4bb357))
* **BYLNET-269:** integrate Pulumi Automation API ([16cd839](https://github.com/metio-plattform/metio/commit/16cd83949840090b8979f7a6b07186f1d24f8c0c))
* **BYLNET-269:** integrate Pulumi Automation API ([4b6ef29](https://github.com/metio-plattform/metio/commit/4b6ef295263185390ab3a9f5b9868a1364ce1238))
* **BYLNET-270:** complete server infrastructure as Pulumi program ([f355448](https://github.com/metio-plattform/metio/commit/f355448f2e3f74b6ab3af17981c40a29c18e7258))
* **BYLNET-270:** define server infrastructure as Pulumi program ([3999e34](https://github.com/metio-plattform/metio/commit/3999e34bcf27e9327012a68dac48320b8b09dc49))
* **BYLNET-271:** create provisioning orchestration service ([b41d75d](https://github.com/metio-plattform/metio/commit/b41d75dc8cd12ea6797a8d3e8036a2fe1457083b))
* **BYLNET-271:** create provisioning orchestration service ([1da7d4a](https://github.com/metio-plattform/metio/commit/1da7d4a4ca343be23e1cd1973de7d77ab412dc44))
* **BYLNET-272:** add provisioning state tracking in Firestore ([105d213](https://github.com/metio-plattform/metio/commit/105d2132b7500dbebe3b95f11d672b375c7b5299))
* **BYLNET-272:** add provisioning state tracking in Firestore ([5134662](https://github.com/metio-plattform/metio/commit/5134662e803fd71365c7739c59ec8bf57f84ac19))
* **BYLNET-273:** design server configuration schema ([52d0fd8](https://github.com/metio-plattform/metio/commit/52d0fd89d145ba6ee0d72946e51687a342e00bf2))
* **BYLNET-273:** design server configuration schema ([9145cb1](https://github.com/metio-plattform/metio/commit/9145cb1efe92d2e4536a89a38f6fa8a88c170884))
* **BYLNET-274:** create server CRUD API endpoints ([7a75a88](https://github.com/metio-plattform/metio/commit/7a75a88576c688458628569ad231598bba10a586))
* **BYLNET-274:** create server CRUD API endpoints ([a66b6f5](https://github.com/metio-plattform/metio/commit/a66b6f5540fad973cc3029b02e53f153f9706129))
* create first-time setup wizard flow ([3daec4a](https://github.com/metio-plattform/metio/commit/3daec4a706e708aa55e8596625e6f94d69a28597))
* create first-time setup wizard flow ([5e4d765](https://github.com/metio-plattform/metio/commit/5e4d7656592d81d9a1fce2fb1975b5e495beedfa))
* detect outdated machine-agent images on existing servers ([e078026](https://github.com/metio-plattform/metio/commit/e078026b9fe18a53e3b0f7e6733b4827d5899576))
* detect outdated machine-agent images on existing servers ([ffc56f3](https://github.com/metio-plattform/metio/commit/ffc56f3ccfce1e42f4b7e4a129a1ba8dddd09435))
* ease development setup ([#625](https://github.com/metio-plattform/metio/issues/625)) ([705cea6](https://github.com/metio-plattform/metio/commit/705cea64f4b1881d6fe6a19ecd96cc233b492187))
* enable Dapr runtime in local development environment ([d9f38c2](https://github.com/metio-plattform/metio/commit/d9f38c225c78c8545a46bb38327f27d25b9aaef2))
* enable Dapr runtime in local development environment ([#252](https://github.com/metio-plattform/metio/issues/252)) ([715f00a](https://github.com/metio-plattform/metio/commit/715f00a2610b37eb9542e72aa83b33cd97f9daa2))
* establish central backup infrastructure ([#434](https://github.com/metio-plattform/metio/issues/434)) ([d1183be](https://github.com/metio-plattform/metio/commit/d1183beb9780840486c386192307a627a98ab143))
* establish k3s baseline on Container-Optimized OS ([#550](https://github.com/metio-plattform/metio/issues/550)) ([7d7465b](https://github.com/metio-plattform/metio/commit/7d7465b884dc76cde381dbb447821f9122265bd4)), closes [#543](https://github.com/metio-plattform/metio/issues/543)
* grant gitlab-runner SA artifact registry write access ([02ce1e7](https://github.com/metio-plattform/metio/commit/02ce1e70db10e56fe859e1924b45b50f6d41b6be))
* hide IP when server is not running ([#524](https://github.com/metio-plattform/metio/issues/524)) ([71aeb12](https://github.com/metio-plattform/metio/commit/71aeb12612d71d8ecef709e7543c6ffb22f84699)), closes [#363](https://github.com/metio-plattform/metio/issues/363)
* hide Players and Uptime stats when server is offline ([7a4441d](https://github.com/metio-plattform/metio/commit/7a4441d109e6659d129085fc2817fafaec820ee4))
* hide Players and Uptime stats when server is offline ([d197136](https://github.com/metio-plattform/metio/commit/d1971368fee57c347c2fd9782fa8c585960c2d38))
* implement async provisioning workflow with 202 Accepted pattern ([ff59f54](https://github.com/metio-plattform/metio/commit/ff59f545e1772bb6d1f8698850b4eb683cdd123d))
* implement async provisioning workflow with 202 Accepted pattern ([1e04c2e](https://github.com/metio-plattform/metio/commit/1e04c2eb4f91345e17b4001a33451287bbba5d4d))
* implement controller agent API with HMAC-JWT per-server token auth ([#300](https://github.com/metio-plattform/metio/issues/300)) ([6d4da14](https://github.com/metio-plattform/metio/commit/6d4da146b8c9142c8c3e15222db6031cf89e5326))
* implement controller agent API with HMAC-JWT per-server token auth ([#300](https://github.com/metio-plattform/metio/issues/300)) ([7b02679](https://github.com/metio-plattform/metio/commit/7b02679b95a3a080769334b0418001cede302e2a))
* implement DaprDB adapter with 24 DB interface methods ([01cc88d](https://github.com/metio-plattform/metio/commit/01cc88dfcac77ead79709e8661a77b07273b1bc0))
* implement DaprDB adapter with 24 DB interface methods ([#256](https://github.com/metio-plattform/metio/issues/256)) ([220281a](https://github.com/metio-plattform/metio/commit/220281a7418eb97a82324438371a116e3dc39e19))
* implement infrastructure versioning system ([57ae356](https://github.com/metio-plattform/metio/commit/57ae356fb35622e8ee8bf8b023acc672be089699))
* implement infrastructure versioning system ([770494e](https://github.com/metio-plattform/metio/commit/770494e7ce2f7591f1bb0a64c2902595edf5612c))
* implement safe server destruction flow with backup option ([bdf83ce](https://github.com/metio-plattform/metio/commit/bdf83ce47852235468d51f09a6a18bc4bf7d43d7))
* implement safe server destruction flow with backup option ([786df2d](https://github.com/metio-plattform/metio/commit/786df2df7281781cbae42f1337df10e567668cbd))
* implement safe server update workflow ([f1caa3e](https://github.com/metio-plattform/metio/commit/f1caa3ebaf300844f5439926e9738482560bd2cd))
* implement safe server update workflow ([a684199](https://github.com/metio-plattform/metio/commit/a6841994d57a9a55ef443155e3ee80f36385043f))
* implement semantic release pipeline ([b738319](https://github.com/metio-plattform/metio/commit/b7383190e49ba2ac5e61032272a800b70789e6cb))
* implement semantic release pipeline ([e87f3bf](https://github.com/metio-plattform/metio/commit/e87f3bf72048510f12b3cb494f8c42e8e1d2abb1))
* improve backup table with server-side pagination, sort, and filter ([#489](https://github.com/metio-plattform/metio/issues/489)) ([eb9f4c7](https://github.com/metio-plattform/metio/commit/eb9f4c7bed1fcf35458cc48f2b756a5218e776ca))
* improve CI/CD pipeline with local docker builds, image promotion, and approval stage ([e8b8f36](https://github.com/metio-plattform/metio/commit/e8b8f367d9b929eee90d4aae6a3c4c51f5d45c79))
* improve codespace setup ([#511](https://github.com/metio-plattform/metio/issues/511)) ([2c3c2e7](https://github.com/metio-plattform/metio/commit/2c3c2e76b6b678ae49415b302edc31c3434b0cb2))
* improve pipeline with approval, local builds, and promote stage ([c108000](https://github.com/metio-plattform/metio/commit/c1080009f77d4e631d624324e93e022527f52de9))
* load GCP machine types dynamically ([#370](https://github.com/metio-plattform/metio/issues/370)) ([#376](https://github.com/metio-plattform/metio/issues/376)) ([7c63eef](https://github.com/metio-plattform/metio/commit/7c63eef0a24d491387fde6665c291505e40257ff))
* load GCP zones and regions dynamically ([#374](https://github.com/metio-plattform/metio/issues/374)) ([7fe92b4](https://github.com/metio-plattform/metio/commit/7fe92b4753dd8d05ee6d4994bd1377b58d2ae200))
* load Minecraft version list from Mojang ([#371](https://github.com/metio-plattform/metio/issues/371)) ([f757092](https://github.com/metio-plattform/metio/commit/f7570928498002ef93e2c8269c1dca83e1016fcf))
* **machine-agent:** report backup manifests to controller ([#478](https://github.com/metio-plattform/metio/issues/478)) ([c3c85a6](https://github.com/metio-plattform/metio/commit/c3c85a6f5f3404c43ba547032aaa67bb6aadf0c5))
* make compact ServerConfigPanel collapsible and closed by default ([52bd981](https://github.com/metio-plattform/metio/commit/52bd98150166004b5d03d58ccd8f2ebaf9a42e9f))
* make compact ServerConfigPanel collapsible and closed by default ([74c2031](https://github.com/metio-plattform/metio/commit/74c2031714795c1437c3780567094faa62f329c8))
* manage modpack on an existing server ([#629](https://github.com/metio-plattform/metio/issues/629)) ([f86e1ac](https://github.com/metio-plattform/metio/commit/f86e1ac54585b1ddbca57b8adc3a65cfb5c00eae))
* measure k3s control-plane resource overhead ([#560](https://github.com/metio-plattform/metio/issues/560)) ([0d67d2c](https://github.com/metio-plattform/metio/commit/0d67d2c06afa1cf9c3ae0171ad9958565c6874a0)), closes [#546](https://github.com/metio-plattform/metio/issues/546)
* migrate server operations to per-server API ([ad4f816](https://github.com/metio-plattform/metio/commit/ad4f81697d619a9109e53536e8a3d54eca4d09a8))
* migrate server operations to per-server API ([17d661f](https://github.com/metio-plattform/metio/commit/17d661fbf13750cdc257378484fbe91d55282211))
* overhaul Makefile with test targets and develop command ([1faa1cc](https://github.com/metio-plattform/metio/commit/1faa1ccba25fc1cc2f502ebff536baae14ce8356))
* overhaul Makefile with test targets and develop command ([00b9cec](https://github.com/metio-plattform/metio/commit/00b9cecdeabfc81834740db6131656fc5b2f5120))
* persist and propagate machine-agent image URL properly ([decf75f](https://github.com/metio-plattform/metio/commit/decf75f1760375d0d91935aacd8d38be063a89d2))
* provision Cloud SQL instance in cloudsql mode ([#327](https://github.com/metio-plattform/metio/issues/327)) ([92a36ed](https://github.com/metio-plattform/metio/commit/92a36ed786442eb7d2f5555d149d2eea809d4d92))
* provision Cloud SQL instance in cloudsql mode ([#327](https://github.com/metio-plattform/metio/issues/327)) ([4cdb596](https://github.com/metio-plattform/metio/commit/4cdb59688cb53bfc7c1249cf32f10f046a04071b))
* remove Firestore resources and datastore permissions ([#328](https://github.com/metio-plattform/metio/issues/328)) ([9c55d3e](https://github.com/metio-plattform/metio/commit/9c55d3e97f7b9f4a8cff52062381ee1958e1c09a))
* remove Firestore resources and datastore permissions ([#328](https://github.com/metio-plattform/metio/issues/328)) ([b7666dc](https://github.com/metio-plattform/metio/commit/b7666dc39c01ee6c2457e02a23a6c96b99354ae7))
* remove OpenTofu configuration for managed Minecraft infrastructure ([5b86e81](https://github.com/metio-plattform/metio/commit/5b86e816192eeefbdfd8c68ce16a5dd02ef72ce8))
* remove OpenTofu configuration for managed Minecraft infrastructure ([d9d6902](https://github.com/metio-plattform/metio/commit/d9d6902aad971e93ac0cd424ed76eb0835db288b))
* replace Datastore emulator with local Postgres container in dev targets ([#326](https://github.com/metio-plattform/metio/issues/326)) ([eae841d](https://github.com/metio-plattform/metio/commit/eae841d5d532de98a02ccf076cdc4fa8290428c6))
* replace Datastore emulator with local Postgres container in dev targets ([#326](https://github.com/metio-plattform/metio/issues/326)) ([6faaa27](https://github.com/metio-plattform/metio/commit/6faaa27f03abe4eededa60f1940185071d64e1b0))
* replace go-semantic-release with release-please ([ecc18e2](https://github.com/metio-plattform/metio/commit/ecc18e29b335729339b99777297ce8b7be04bc66))
* replace go-semantic-release with release-please ([e98dc9c](https://github.com/metio-plattform/metio/commit/e98dc9cd695e8cbabcc1ad824534cbaac27664fa))
* report modpack install progress during provisioning ([#632](https://github.com/metio-plattform/metio/issues/632)) ([fbd2a39](https://github.com/metio-plattform/metio/commit/fbd2a3987fed0fd5345ae4e4d2c260b846038244))
* run controller on Dapr statestore when DB_BACKEND=dapr ([888251e](https://github.com/metio-plattform/metio/commit/888251e6d3ad3a279b2e1af76c483100e1be400b))
* run Minecraft as a k3s workload on the attached disk ([#558](https://github.com/metio-plattform/metio/issues/558)) ([5c3e909](https://github.com/metio-plattform/metio/commit/5c3e90986c02d4432059b1441cb7cba6745f065c)), closes [#544](https://github.com/metio-plattform/metio/issues/544)
* set up GitHub milestone mapping ([99ed3ac](https://github.com/metio-plattform/metio/commit/99ed3ac848728666a8fcf81237e6c3c6840b3c4d)), closes [#240](https://github.com/metio-plattform/metio/issues/240)
* show controller version in web UI ([#377](https://github.com/metio-plattform/metio/issues/377)) ([9832d56](https://github.com/metio-plattform/metio/commit/9832d5648e3ca35ccf445f9f0b03c6024b055d5c)), closes [#366](https://github.com/metio-plattform/metio/issues/366)
* size JVM heap from machine type as a percentage ([#576](https://github.com/metio-plattform/metio/issues/576)) ([3de95cc](https://github.com/metio-plattform/metio/commit/3de95cceb3a9bf846ffca84be7a0131f64e84ec4))
* support importing existing static IP addresses when provisioning servers ([7dd8dde](https://github.com/metio-plattform/metio/commit/7dd8dde832f09de29e554a2455a50634fc8c11d4))
* support importing existing static IP addresses when provisioning servers ([4e7b477](https://github.com/metio-plattform/metio/commit/4e7b47749f22b2b5577d40302bdf78a25156e0e2)), closes [#293](https://github.com/metio-plattform/metio/issues/293)
* support Minecraft 26.x with the stable-java25 image ([#513](https://github.com/metio-plattform/metio/issues/513)) ([4daae04](https://github.com/metio-plattform/metio/commit/4daae040dd069b5d16cb7ed4ccea7c39f7bd6512))
* surface crash-looping Minecraft servers ([#643](https://github.com/metio-plattform/metio/issues/643)) ([0757fe9](https://github.com/metio-plattform/metio/commit/0757fe91149e430e464ccb1a88602eea9d819523)), closes [#637](https://github.com/metio-plattform/metio/issues/637)
* switch baked Dapr component to state.postgresql ([#324](https://github.com/metio-plattform/metio/issues/324)) ([ddb660d](https://github.com/metio-plattform/metio/commit/ddb660d0b5c02c51fdc9e52ac94dda24c7d0e68a))
* switch baked Dapr component to state.postgresql ([#324](https://github.com/metio-plattform/metio/issues/324)) ([1dd7950](https://github.com/metio-plattform/metio/commit/1dd7950d2699f531b19170c7a44264f9aebce437))
* sync release/* labels with Target release project field ([#517](https://github.com/metio-plattform/metio/issues/517)) ([33561d3](https://github.com/metio-plattform/metio/commit/33561d3af46d189d0f20d282bf086ed16e42bd55)), closes [#469](https://github.com/metio-plattform/metio/issues/469)
* **ui:** remove final-backup step from destroy wizard ([c94058e](https://github.com/metio-plattform/metio/commit/c94058ed6df11cdd83336da9c3e123efee57a459))
* **ui:** remove final-backup step from destroy wizard ([73a47a7](https://github.com/metio-plattform/metio/commit/73a47a703ebba9c27286827a3ed913bc560f440f)), closes [#345](https://github.com/metio-plattform/metio/issues/345)
* verify cluster and workload survive preemption ([#559](https://github.com/metio-plattform/metio/issues/559)) ([f8393ee](https://github.com/metio-plattform/metio/commit/f8393ee7ac3b3dae606bb67f60fa6dbde32c9607)), closes [#545](https://github.com/metio-plattform/metio/issues/545)
* verify config changes avoid VM replacement ([#561](https://github.com/metio-plattform/metio/issues/561)) ([bec6b67](https://github.com/metio-plattform/metio/commit/bec6b67d730056f3e66dd83aa5f1be1d2000bdd0)), closes [#547](https://github.com/metio-plattform/metio/issues/547)
* **web:** add dark mode with system-default theme switcher ([#468](https://github.com/metio-plattform/metio/issues/468)) ([278736d](https://github.com/metio-plattform/metio/commit/278736d8d3745edc644ca4edfd7d718424f86c96))
* **web:** apply shadcn vega preset ([#461](https://github.com/metio-plattform/metio/issues/461)) ([f1fcbe4](https://github.com/metio-plattform/metio/commit/f1fcbe4e70fc429d5644d0e896f5b8a1be5d1dcd))
* **web:** migrate layout and shell components ([#457](https://github.com/metio-plattform/metio/issues/457)) ([7ccfdef](https://github.com/metio-plattform/metio/commit/7ccfdefc6b6bfb6c2711cfadea93e880a2b5c169))
* **web:** migrate modals and progress to shadcn ([#455](https://github.com/metio-plattform/metio/issues/455)) ([3df2118](https://github.com/metio-plattform/metio/commit/3df2118f1b358fb7d4b0ed53550b3d344b274e36))
* **web:** migrate server pages to shadcn and RHF ([#459](https://github.com/metio-plattform/metio/issues/459)) ([6746e04](https://github.com/metio-plattform/metio/commit/6746e045ced1a9956b1d19c143509827ee27d5d1))
* **web:** migrate setup wizards to RHF and Zod ([#458](https://github.com/metio-plattform/metio/issues/458)) ([89db4fd](https://github.com/metio-plattform/metio/commit/89db4fd0c089b6d923e136a20480da8947f63b1a))
* **web:** migrate to plain shadcn components ([#454](https://github.com/metio-plattform/metio/issues/454)) ([b7de099](https://github.com/metio-plattform/metio/commit/b7de0990029de88fa136c740a71720ba9e141882))
* **web:** replace UI primitives with shadcn components ([#452](https://github.com/metio-plattform/metio/issues/452)) ([ec4639d](https://github.com/metio-plattform/metio/commit/ec4639ddee5575aac19e6dd73ab2faf45c062d51))


### Bug Fixes

* add daprd startup probe, remove --app-port for Cloud Run depends_on ([af29ceb](https://github.com/metio-plattform/metio/commit/af29ceb5f3f5600d5f7f00a6fa70ac408433299d))
* add deploy_id annotation to force new Cloud Run revisions ([#484](https://github.com/metio-plattform/metio/issues/484)) ([6b11cd0](https://github.com/metio-plattform/metio/commit/6b11cd04ae1c05f0f54ba4ad5d3be02a9835caf1))
* add environment: main to release job ([7947e55](https://github.com/metio-plattform/metio/commit/7947e558482ea9097acf3db93e30c9319fba66c6))
* add environment: main to release job ([3577d26](https://github.com/metio-plattform/metio/commit/3577d263239b72f3d8f35e707b9935b9c0c103a9))
* add explicit command for daprd sidecar ([a874b1a](https://github.com/metio-plattform/metio/commit/a874b1a61dd83a9876a809b58f62162ff4319ebc))
* add packages:write to promote-images for ghcr.io version tagging ([c545997](https://github.com/metio-plattform/metio/commit/c545997f9e0fad1238c3982f3603835a69ca7af5))
* add packages:write to promote-images for ghcr.io version tagging ([c6152b9](https://github.com/metio-plattform/metio/commit/c6152b99415e91a284d02ed0393a2ae73855cf99))
* add PlayerDB fallback for Minecraft user lookup ([e237dff](https://github.com/metio-plattform/metio/commit/e237dff864bdfc73e7572889431ca6a408125f4c))
* add PlayerDB fallback for Minecraft user lookup ([2246f9d](https://github.com/metio-plattform/metio/commit/2246f9ddce92a609f9a97c3d0ca984f8b0fe98bf))
* add release-please extra-files for variables.tf ([5d8f990](https://github.com/metio-plattform/metio/commit/5d8f990bbcc9e49f3929c5e03c1b99ef5fb41e0c))
* add release-please extra-files for variables.tf ([a06d54a](https://github.com/metio-plattform/metio/commit/a06d54a93153bb60f4141328f015ea0f1230e95a))
* add text-foreground to btn-outline class to make Cancel/Logout buttons visible ([5ef8c4b](https://github.com/metio-plattform/metio/commit/5ef8c4b479d52a5e828301ff261d19edc9da078c))
* add text-foreground to btn-outline class to make Cancel/Logout buttons visible ([97ef4e1](https://github.com/metio-plattform/metio/commit/97ef4e1bac0107fb844bc9fdc559827ff31ed111))
* add visual clickability cues to provisioning progress bar ([8fc70aa](https://github.com/metio-plattform/metio/commit/8fc70aa6938c6441656d8d5b56574719513c4c8f))
* add visual clickability cues to provisioning progress bar ([5337ce1](https://github.com/metio-plattform/metio/commit/5337ce155a25f04f601fdb76703fafda9f689c59))
* address MR review feedback for scheduled shutdown ([0112b36](https://github.com/metio-plattform/metio/commit/0112b36ede508b055fd87b01d947925042e4da93))
* build files-updater plugin from its own module context ([96e19a1](https://github.com/metio-plattform/metio/commit/96e19a1d75f458222fa29d86db23d51febd9e3cb))
* build files-updater plugin from its own module context ([38b2d4b](https://github.com/metio-plattform/metio/commit/38b2d4b428c0d980cf5ce54bdbcf1759f4e9b5d5))
* **build:** inject release version into controller and machine-agent images ([#467](https://github.com/metio-plattform/metio/issues/467)) ([d045ab6](https://github.com/metio-plattform/metio/commit/d045ab6295aef2631d69dd49ba6e010a81490c44))
* bump Go base image to 1.26 for Dapr SDK compatibility ([#256](https://github.com/metio-plattform/metio/issues/256)) ([05d3042](https://github.com/metio-plattform/metio/commit/05d30421d01f9afaf4969d499de4aa2510b64602))
* **BYLNET-176:** add Docker image verification with retry logic ([fa25fdd](https://github.com/metio-plattform/metio/commit/fa25fdd26fd2fca53a1104c3d25fd831728e96c0))
* **BYLNET-176:** add Docker image verification with retry logic ([50e1c57](https://github.com/metio-plattform/metio/commit/50e1c57aaaa074346fc2e28812b0d8bac3236b3c))
* **BYLNET-176:** use --entrypoint to override container entrypoint in verification ([416a5b2](https://github.com/metio-plattform/metio/commit/416a5b229ab5457c14eb15ed016b52795acc5338))
* **BYLNET-269:** use shared cloud-config template with all services ([ce27f96](https://github.com/metio-plattform/metio/commit/ce27f960dedbf709009ed36fcd1030714ad445c7))
* **BYLNET-269:** use single verification script with snake_case device names ([36b7524](https://github.com/metio-plattform/metio/commit/36b7524c1d0e253971ee607d74ef28c2cf4613b3))
* **BYLNET-271:** address code review comments ([06a0f83](https://github.com/metio-plattform/metio/commit/06a0f8335212c176926bbd982c83619fa0a31a1f))
* **ci:** promote images using the tag pulled from ghcr.io ([#525](https://github.com/metio-plattform/metio/issues/525)) ([20481ce](https://github.com/metio-plattform/metio/commit/20481cec0889da292e64ad65340081690edfea1b))
* clear Target Release via clearProjectV2ItemFieldValue ([#523](https://github.com/metio-plattform/metio/issues/523)) ([ac87ac3](https://github.com/metio-plattform/metio/commit/ac87ac3cb4f3bb5eb584ceefd5fbab104426c8fa))
* configure session cookie options so added_by is populated correctly ([27f8f3b](https://github.com/metio-plattform/metio/commit/27f8f3baced3d3152112e85a7cebfae8ddcda89f))
* configure session cookie options so added_by is populated correctly ([32fc751](https://github.com/metio-plattform/metio/commit/32fc751046a4fd3ca031c16d88f9fd192696d8dd))
* correct backup JSON tags and Docker build ordering ([#485](https://github.com/metio-plattform/metio/issues/485)) ([2f3f288](https://github.com/metio-plattform/metio/commit/2f3f2884767b13e5e6f1e21f6b418e1d1c74a903)), closes [#423](https://github.com/metio-plattform/metio/issues/423)
* correct Dapr firestore component field names and daprd readiness check ([#252](https://github.com/metio-plattform/metio/issues/252)) ([78fea0c](https://github.com/metio-plattform/metio/commit/78fea0c1c230f772aefc68c4a31bb48374e26224))
* **css:** don't re-run lightningcss visitor during minify (fix &lt;a ([46acda0](https://github.com/metio-plattform/metio/commit/46acda08ea9b6578ec50ab289badbde8bf91f8fd))
* delete server config from database after stack destruction ([4aa572d](https://github.com/metio-plattform/metio/commit/4aa572d5cf24a11ad421e3f248c0197dc946cf9e))
* delete server config from database after stack destruction ([b776556](https://github.com/metio-plattform/metio/commit/b7765563827d6ca205efb3f8d926c762283626c3))
* **deploy:** drop write-only attributes for OpenTofu &lt;1.11 compatibility ([ddbf4e9](https://github.com/metio-plattform/metio/commit/ddbf4e97a29c4e56bd57844c9d39ae31de29467b))
* **deploy:** drop write-only attributes for OpenTofu &lt;1.11 compatibility ([4851d51](https://github.com/metio-plattform/metio/commit/4851d518e370a421516b12b66c19106684931b0f)), closes [#351](https://github.com/metio-plattform/metio/issues/351)
* **deploy:** reference BYO connection string via external secret ([743bfae](https://github.com/metio-plattform/metio/commit/743bfae0ba0160b89e6e362adffe22500bc7e044))
* **deploy:** reference BYO connection string via external secret ([9908692](https://github.com/metio-plattform/metio/commit/9908692240d796ae92972cff6e2ef42a69e95048))
* **deploy:** resolve Cloud SQL cloudsql-mode deployment failures ([7e12513](https://github.com/metio-plattform/metio/commit/7e125138c127d35a43f223dc988ad0185fd6e407))
* **deploy:** resolve Cloud SQL cloudsql-mode deployment failures ([0cec207](https://github.com/metio-plattform/metio/commit/0cec2070ae380619e9a371ea532b8be113a9445e))
* derive machine-agent registry host from image URL to fix cross-region pulls ([4e3cd21](https://github.com/metio-plattform/metio/commit/4e3cd21a15c7c3a701ee7e0dc402b5d65f55f1b3))
* **devcontainer:** install pre-commit in isolated venv ([#607](https://github.com/metio-plattform/metio/issues/607)) ([37085a4](https://github.com/metio-plattform/metio/commit/37085a4da371724d3db8cca62bd6804a5282e0ab))
* **devcontainer:** support GitHub Codespaces ([#502](https://github.com/metio-plattform/metio/issues/502)) ([1c1e710](https://github.com/metio-plattform/metio/commit/1c1e7104c9ae343a98d0a8f8046be6e70a34fed2))
* disable all action buttons during provisioning/destroying ([c78ea65](https://github.com/metio-plattform/metio/commit/c78ea65833dce2243d23c993234defe5d7557f4b))
* disable all action buttons during provisioning/destroying ([3056f90](https://github.com/metio-plattform/metio/commit/3056f90386b1e35733ead0e0404fc23ae4f1eda0))
* enable ForceDestroy on backup bucket to prevent destroy failure ([49a93e4](https://github.com/metio-plattform/metio/commit/49a93e41277e29d73ca8a85ef2c7f0ea9ae6cc74))
* enable ForceDestroy on backup bucket to prevent destroy failure ([dd59147](https://github.com/metio-plattform/metio/commit/dd5914793ec3b2ddce1bba7bbf4c577f045d20e1))
* escape percent sign for systemd and align tofu config with GCP state ([5179cb1](https://github.com/metio-plattform/metio/commit/5179cb101064b86b86b1d6428d63ee964fbd6773))
* expose controller ingress port 8080 for Cloud Run multi-container ([6f23aa8](https://github.com/metio-plattform/metio/commit/6f23aa8fa3b4bea0a85c6cf8a14fd8e0e4f89bb5))
* gracefully handle missing status document in backup coordinator ([dc8c3c2](https://github.com/metio-plattform/metio/commit/dc8c3c2b4580b32a2998bf02182b5607520427e8))
* gracefully handle missing status document in backup coordinator ([0acc709](https://github.com/metio-plattform/metio/commit/0acc709cee25ece7e18b9395700b1a43d21992b8))
* gracefully handle validation service errors in setup status ([02f22b8](https://github.com/metio-plattform/metio/commit/02f22b8813e5c7dfc253e438c530d291ba3c61bd))
* gracefully handle validation service errors in setup status ([701cac2](https://github.com/metio-plattform/metio/commit/701cac22e78a6e6f02bb422ee63201462235e78c))
* import existing address via pulumi.Import instead of the CLI ([fadf8db](https://github.com/metio-plattform/metio/commit/fadf8dbb90a814ea7aaaa8ecb4dc89f33700d276))
* import existing address via pulumi.Import instead of the CLI ([e7c7fcb](https://github.com/metio-plattform/metio/commit/e7c7fcb276707eb7cd492c1eb5904038eb5c4865)), closes [#356](https://github.com/metio-plattform/metio/issues/356)
* include Pulumi CLI in controller Docker image ([86bc64e](https://github.com/metio-plattform/metio/commit/86bc64e9859d1b8912e41746b46c814af0f62b4d))
* inline verify-docker-image.sh into unit files ([bd1efb5](https://github.com/metio-plattform/metio/commit/bd1efb52b683170daa0bfa79038e2b066ccf6e5c))
* keep spark hook off the read-only COS root ([#640](https://github.com/metio-plattform/metio/issues/640)) ([9e28123](https://github.com/metio-plattform/metio/commit/9e281235c848fc3df266c86963919ed2593977fc))
* keep spark's background profiler off pack servers on Java 25 ([#636](https://github.com/metio-plattform/metio/issues/636)) ([6f3468f](https://github.com/metio-plattform/metio/commit/6f3468ff7543dbad6f71f69ec7b4977baa2b59ef))
* load MACHINE_AGENT_IMAGE via config package, fail on missing value ([2755a49](https://github.com/metio-plattform/metio/commit/2755a4956717b2b094cb9ec118e011056d3a7993))
* load MACHINE_AGENT_IMAGE via config package, fail on missing value ([6db42c5](https://github.com/metio-plattform/metio/commit/6db42c5d4d908bdba0d5cc6c9e531389cc503b10))
* **machine-agent:** restore docker CLI and CA certs on debian 13 ([#448](https://github.com/metio-plattform/metio/issues/448)) ([7a62b1d](https://github.com/metio-plattform/metio/commit/7a62b1de691f2053dee0f2df427ebaad5e3ad3db))
* make build-frontend a dependency of verify-backend and test-backend ([2b6db4e](https://github.com/metio-plattform/metio/commit/2b6db4e3727e2f69f980b4107cc4f39af37dfe00))
* **mc-backup:** pull base image from Docker Hub instead of GHCR ([#508](https://github.com/metio-plattform/metio/issues/508)) ([a1275f0](https://github.com/metio-plattform/metio/commit/a1275f09e0d53028f7d1b0003c506d27ae8d2749))
* metio user UID collision on Container-Optimized OS ([#628](https://github.com/metio-plattform/metio/issues/628)) ([6f65727](https://github.com/metio-plattform/metio/commit/6f65727e02fe5a7e3a35d65c632f3624e05b58a5))
* move serverID from GCE instance tags to labels ([99edd0a](https://github.com/metio-plattform/metio/commit/99edd0ab5a5184b8d1149fb467938575a8add375))
* move verify-docker-image.sh to /home/metio for COS compatibility ([cf6db7f](https://github.com/metio-plattform/metio/commit/cf6db7fefeccb7b6d9181752c32b2a8c7431e0ff))
* **opencode:** upgrade to v2.0.29 and use 'opencode serve' ([#610](https://github.com/metio-plattform/metio/issues/610)) ([3555895](https://github.com/metio-plattform/metio/commit/3555895fba6c209b078628fa2fc3416ab9c8eb74)), closes [#609](https://github.com/metio-plattform/metio/issues/609)
* pass ExistingAddress through Cloud Tasks handler ([2d43d51](https://github.com/metio-plattform/metio/commit/2d43d51c7f75f23b2b7e6ed1e97241d7eff5e656))
* pass ExistingAddress through Cloud Tasks handler ([68e7d47](https://github.com/metio-plattform/metio/commit/68e7d477e190c946c778f3959b943799472723c6))
* prevent nil pointer crash when services fail to initialize ([fe87770](https://github.com/metio-plattform/metio/commit/fe87770349aaf5607e4fe644f7988e78a1b725cc))
* provider settings screwed by the AI ([c8d7d64](https://github.com/metio-plattform/metio/commit/c8d7d64c55f211d0a0916df37ea6b1950fe6ea38))
* provider settings screwed by the AI ([4c64d3a](https://github.com/metio-plattform/metio/commit/4c64d3afb1cb3c454aa3782a0344b185e41514ff))
* provisioning issues on Cloud Run ([a363ae7](https://github.com/metio-plattform/metio/commit/a363ae72b55d8e6825bd672586298f4e0035de5e))
* provisioning issues on Cloud Run ([f4a7ddf](https://github.com/metio-plattform/metio/commit/f4a7ddf3003c1c9cded5a51a242e7338317a86cf))
* provisioning progress screen now refetches when navigated to ([6567a1f](https://github.com/metio-plattform/metio/commit/6567a1fc07386197c00874b91d92c05f2344b20d))
* provisioning progress screen now refetches when navigated to ([0451fff](https://github.com/metio-plattform/metio/commit/0451fff1f48db72cf945fd0e38e8ef88a8996793))
* refresh stack after import and enable code generation for pulumi import ([e87e018](https://github.com/metio-plattform/metio/commit/e87e0189cba139b97584eae2e3c328da78fa90df))
* reject duplicate server names on create and rename ([#372](https://github.com/metio-plattform/metio/issues/372)) ([dc68245](https://github.com/metio-plattform/metio/commit/dc68245020e407d1b3a18f1505e033f829ac9b01))
* **release:** gate image-default update on promoted images ([#530](https://github.com/metio-plattform/metio/issues/530)) ([89c247b](https://github.com/metio-plattform/metio/commit/89c247b2f8fc99300784ee42e8f61bc9d7e30235))
* remove backup step from server destroy process ([91a1fee](https://github.com/metio-plattform/metio/commit/91a1feef574633adc401777fb3865d30ec8d1b31))
* remove broken -target from deploy-controller Makefile ([ef0fd84](https://github.com/metio-plattform/metio/commit/ef0fd84d0a5a99c3f2bb199397c82e73ad0b2878))
* remove stale provisioning cache instead of invalidating before navigation ([6db32f1](https://github.com/metio-plattform/metio/commit/6db32f1ce8593f81a4f39a8dc3a3e518b6a394f1))
* remove stale provisioning cache instead of invalidating before navigation ([ece00d2](https://github.com/metio-plattform/metio/commit/ece00d2fcdf9c062406656506413b46b734b871c))
* rename release-please config and add checkout step ([9f5d3bd](https://github.com/metio-plattform/metio/commit/9f5d3bd2b0cb50b6b51471f66b4d075cf59b85b7))
* rename release-please config file and add checkout step ([9672939](https://github.com/metio-plattform/metio/commit/9672939ed1a1a1d3b68c5fd8d9fde142f24f3522))
* replace blue accents with green theme and fix missing Layout wrapper on /servers/new ([cf2ee25](https://github.com/metio-plattform/metio/commit/cf2ee25f6307b22f2ade14635462abeedfbfb5ef))
* replace blue accents with green theme and fix missing Layout wrapper on /servers/new ([f687a04](https://github.com/metio-plattform/metio/commit/f687a04ae9680ba0678229ad3b394f549cbaf66b))
* report STOPPED while minecraft boots instead of aborting status update ([#514](https://github.com/metio-plattform/metio/issues/514)) ([14ed9a9](https://github.com/metio-plattform/metio/commit/14ed9a961155e8edba2d003356c2d414b2f916fe))
* resolve CI pipeline failures and migrate to Makefile targets ([62063d1](https://github.com/metio-plattform/metio/commit/62063d1f1b0803e143c0f4454cad16343ac36309))
* resolve Dapr Firestore project id via ADC auto-detection ([1912294](https://github.com/metio-plattform/metio/commit/1912294821ff3b21698f981f424f753d87531b06))
* resolve frontend TS errors, fix CI pipeline, add Makefile dependencies ([49342e8](https://github.com/metio-plattform/metio/commit/49342e88972248f60c559c8ade12479075f9ca5a))
* resolve frontend TypeScript compilation errors in server dashboard ([5638b63](https://github.com/metio-plattform/metio/commit/5638b630b8ed01671003b96bbbb67bc32b960f4c))
* resolve merge conflict in devcontainer.json ([98ff8ef](https://github.com/metio-plattform/metio/commit/98ff8ef5681e4af57b7a20a998f500055eafb949))
* **restore:** repair restore actions — 404, empty create-from-backup, and save_world ack timeout ([#491](https://github.com/metio-plattform/metio/issues/491)) ([c7b5acb](https://github.com/metio-plattform/metio/commit/c7b5acba5a552772549b374c0a0866981be69048))
* restrict approval job to merge request pipelines only ([b49e9d6](https://github.com/metio-plattform/metio/commit/b49e9d669e8faa0a48514bf32f058c0f66ed451b))
* save image tags to build/*.txt and push to Artifact Registry in local docker targets ([391775d](https://github.com/metio-plattform/metio/commit/391775d3c504513e512b737258ca9db2a79fbeb5))
* scope gitignore patterns for binary files to repo root ([24a8960](https://github.com/metio-plattform/metio/commit/24a8960f089eaf0e424c6a6a740b2c8276f66660))
* set empty PULUMI_CONFIG_PASSPHRASE for self-managed backend ([661ddd6](https://github.com/metio-plattform/metio/commit/661ddd64d1c5508bd0f89c947c14f9af610211d3))
* set gcp:project stack config before running Pulumi operations ([ef7fc1f](https://github.com/metio-plattform/metio/commit/ef7fc1f0d66fc0d815afd7f274002070d3b64334))
* show new servers as starting until the agent reports ([26eeb1d](https://github.com/metio-plattform/metio/commit/26eeb1dfe9e3c814d37bdbf4c68a2560c1e4eb24))
* show new servers as starting until the agent reports ([a5b7da6](https://github.com/metio-plattform/metio/commit/a5b7da66de8fa66340687aab589f5b59e96d3f67)), closes [#348](https://github.com/metio-plattform/metio/issues/348)
* split docker image targets for CI and deploy, run image builds on PRs ([250a4d9](https://github.com/metio-plattform/metio/commit/250a4d93c54d8f29a08446024658baf98e7a8991))
* split docker image targets for CI and deploy, run image builds on PRs ([2280df1](https://github.com/metio-plattform/metio/commit/2280df1625b703925ab47efed1df2fecce9d3a5e))
* suppress tofu warnings by using secret_data_wo write-only alternative ([6fbc550](https://github.com/metio-plattform/metio/commit/6fbc550a92eb962a1c1ec08a4a25d518664f2bca))
* treat PlayerDB 400 as not-found and fall back on Mojang errors ([#332](https://github.com/metio-plattform/metio/issues/332)) ([4e62ed8](https://github.com/metio-plattform/metio/commit/4e62ed83108a907d1c097766e17418df37af540e))
* treat PlayerDB 400 as not-found and fall back on Mojang errors ([#332](https://github.com/metio-plattform/metio/issues/332)) ([e544028](https://github.com/metio-plattform/metio/commit/e544028069ad4a3fabe2e0bbd476f179a65afb1d))
* typo in firestore resource reference - matio -&gt; metio ([a82d45a](https://github.com/metio-plattform/metio/commit/a82d45ae1b12bb7f61fd205da75c77ae7a8f022a))
* **ui:** show Create Server button when servers exist ([095b0a7](https://github.com/metio-plattform/metio/commit/095b0a73ecfea3316f647d69c496b9beb874f918))
* **ui:** show Create Server button when servers exist ([da5058a](https://github.com/metio-plattform/metio/commit/da5058a7710c362d8f4b1e82d34759aca820fcb5)), closes [#359](https://github.com/metio-plattform/metio/issues/359)
* update copyright year to 2026 ([f201619](https://github.com/metio-plattform/metio/commit/f2016195fd7f3d9835da35d9d92eebe4095f5a5b))
* update DISTRO_REGISTRY to metio-distribution project ([f592bc9](https://github.com/metio-plattform/metio/commit/f592bc9019373e3061f370a2de07703fb06735e8))
* update DISTRO_REGISTRY to metio-distribution project ([8f6b863](https://github.com/metio-plattform/metio/commit/8f6b86317730436aedc7dcb0b413b63d3414535a))
* update Dockerfile and devcontainer paths after frontend/ -&gt; web/ rename ([c35898b](https://github.com/metio-plattform/metio/commit/c35898b963af2f34a1954b30482268d402219b42))
* update Dockerfile paths after frontend/ -&gt; web/ rename ([b6fecd8](https://github.com/metio-plattform/metio/commit/b6fecd8217183ab37caaaffeaa68e9ed93aa55f7))
* use project-scoped token for release label sync ([#519](https://github.com/metio-plattform/metio/issues/519)) ([ff76da4](https://github.com/metio-plattform/metio/commit/ff76da4340ccc3c4686c99182797b834d74cce7b))
* use PULUMI_BACKEND_URL for GCS state backend instead of non-existent PULUMI_STATE_BUCKET ([7b36b43](https://github.com/metio-plattform/metio/commit/7b36b4317ef1c228f6dec48e3cb11b40e708f4b5))
* use RELEASE_PLEASE_TOKEN secret for release-please-action ([0d05bd0](https://github.com/metio-plattform/metio/commit/0d05bd04843938c8a33cf90676b33841e3aa3977))
* use RELEASE_PLEASE_TOKEN secret for release-please-action ([e6fe4f3](https://github.com/metio-plattform/metio/commit/e6fe4f3e750c2215baf47e26a5921e74c10f85d3))
* use server-name-based GCP resource naming, per-server backup buckets, and fix server listing ([f8e81f0](https://github.com/metio-plattform/metio/commit/f8e81f025c251fca8531dfb61b60d5be480fe5c4))
* use user.projectV2 query shape for release label sync ([#522](https://github.com/metio-plattform/metio/issues/522)) ([d1e6099](https://github.com/metio-plattform/metio/commit/d1e60992b6e508cda03265f599ba7264a54aebc7))
* **web:** complete light theme migration and harden layout assertions ([#464](https://github.com/metio-plattform/metio/issues/464)) ([da8c440](https://github.com/metio-plattform/metio/commit/da8c440903239eb467f2e0676f3941c3e102f3c6))
* **web:** support eslint 10 in CI build ([4315c53](https://github.com/metio-plattform/metio/commit/4315c536aa6008407bdf9d96e05ea497adbcf74b))
* whitelist endpoint returns 500 for new servers ([2eb9662](https://github.com/metio-plattform/metio/commit/2eb96624541b3629dcca3ff3b2050e916af9221c))
* whitelist endpoint returns 500 for new servers ([897a8b5](https://github.com/metio-plattform/metio/commit/897a8b54cac3fad64d3fa4838d0121db9e0eb9d9)), closes [#344](https://github.com/metio-plattform/metio/issues/344)
* wire Dapr emulator env into daprd, unify develop target ([#252](https://github.com/metio-plattform/metio/issues/252)) ([5c647e6](https://github.com/metio-plattform/metio/commit/5c647e659cc9b6ec85e6ec9d4a2edf0bd488a58c))


### Performance Improvements

* **build:** reorder controller layers, add Go cache mounts, drop Cloud Build ([#465](https://github.com/metio-plattform/metio/issues/465)) ([fcf67d7](https://github.com/metio-plattform/metio/commit/fcf67d7f87e96587f63337ae5a76c3e8652de7c2))

## [2.2.0](https://github.com/nbyl/metio/compare/v2.1.0...v2.2.0) (2026-10-16)


### Features

* add modpack data model and cloud-config ([#580](https://github.com/nbyl/metio/issues/580)) ([2331f9a](https://github.com/nbyl/metio/commit/2331f9a12a8d7c4aa64cc47e8acbaefdd2ec21df)), closes [#536](https://github.com/nbyl/metio/issues/536)
* add modpack search API endpoints ([#579](https://github.com/nbyl/metio/issues/579)) ([fbfcf1a](https://github.com/nbyl/metio/commit/fbfcf1a6539a867416f4e173a3c43fdfbd883b4c)), closes [#535](https://github.com/nbyl/metio/issues/535)
* add Modrinth modpack search service ([#577](https://github.com/nbyl/metio/issues/577)) ([c5a0680](https://github.com/nbyl/metio/commit/c5a06801ace42a76148eeda4b2217c050d40a81e)), closes [#534](https://github.com/nbyl/metio/issues/534)
* add opencode-web dev environment target ([#578](https://github.com/nbyl/metio/issues/578)) ([9711209](https://github.com/nbyl/metio/commit/9711209878a041dbe1e732a660d5ebc5de485ffe))
* add throwaway VM harness for the ADR-0008 k3s spike ([#549](https://github.com/nbyl/metio/issues/549)) ([43c0114](https://github.com/nbyl/metio/commit/43c0114f163d2fba3a6fa070b9441b57ea151e4a)), closes [#542](https://github.com/nbyl/metio/issues/542)
* establish k3s baseline on Container-Optimized OS ([#550](https://github.com/nbyl/metio/issues/550)) ([7d7465b](https://github.com/nbyl/metio/commit/7d7465b884dc76cde381dbb447821f9122265bd4)), closes [#543](https://github.com/nbyl/metio/issues/543)
* measure k3s control-plane resource overhead ([#560](https://github.com/nbyl/metio/issues/560)) ([0d67d2c](https://github.com/nbyl/metio/commit/0d67d2c06afa1cf9c3ae0171ad9958565c6874a0)), closes [#546](https://github.com/nbyl/metio/issues/546)
* run Minecraft as a k3s workload on the attached disk ([#558](https://github.com/nbyl/metio/issues/558)) ([5c3e909](https://github.com/nbyl/metio/commit/5c3e90986c02d4432059b1441cb7cba6745f065c)), closes [#544](https://github.com/nbyl/metio/issues/544)
* size JVM heap from machine type as a percentage ([#576](https://github.com/nbyl/metio/issues/576)) ([3de95cc](https://github.com/nbyl/metio/commit/3de95cceb3a9bf846ffca84be7a0131f64e84ec4))
* verify cluster and workload survive preemption ([#559](https://github.com/nbyl/metio/issues/559)) ([f8393ee](https://github.com/nbyl/metio/commit/f8393ee7ac3b3dae606bb67f60fa6dbde32c9607)), closes [#545](https://github.com/nbyl/metio/issues/545)
* verify config changes avoid VM replacement ([#561](https://github.com/nbyl/metio/issues/561)) ([bec6b67](https://github.com/nbyl/metio/commit/bec6b67d730056f3e66dd83aa5f1be1d2000bdd0)), closes [#547](https://github.com/nbyl/metio/issues/547)


### Bug Fixes

* **ci:** promote images using the tag pulled from ghcr.io ([#525](https://github.com/nbyl/metio/issues/525)) ([20481ce](https://github.com/nbyl/metio/commit/20481cec0889da292e64ad65340081690edfea1b))
* **release:** gate image-default update on promoted images ([#530](https://github.com/nbyl/metio/issues/530)) ([89c247b](https://github.com/nbyl/metio/commit/89c247b2f8fc99300784ee42e8f61bc9d7e30235))

## [2.1.0](https://github.com/nbyl/metio/compare/v2.0.0...v2.1.0) (2026-09-10)


### Features

* hide IP when server is not running ([#524](https://github.com/nbyl/metio/issues/524)) ([71aeb12](https://github.com/nbyl/metio/commit/71aeb12612d71d8ecef709e7543c6ffb22f84699)), closes [#363](https://github.com/nbyl/metio/issues/363)
* improve codespace setup ([#511](https://github.com/nbyl/metio/issues/511)) ([2c3c2e7](https://github.com/nbyl/metio/commit/2c3c2e76b6b678ae49415b302edc31c3434b0cb2))
* support Minecraft 26.x with the stable-java25 image ([#513](https://github.com/nbyl/metio/issues/513)) ([4daae04](https://github.com/nbyl/metio/commit/4daae040dd069b5d16cb7ed4ccea7c39f7bd6512))
* sync release/* labels with Target release project field ([#517](https://github.com/nbyl/metio/issues/517)) ([33561d3](https://github.com/nbyl/metio/commit/33561d3af46d189d0f20d282bf086ed16e42bd55)), closes [#469](https://github.com/nbyl/metio/issues/469)


### Bug Fixes

* clear Target Release via clearProjectV2ItemFieldValue ([#523](https://github.com/nbyl/metio/issues/523)) ([ac87ac3](https://github.com/nbyl/metio/commit/ac87ac3cb4f3bb5eb584ceefd5fbab104426c8fa))
* report STOPPED while minecraft boots instead of aborting status update ([#514](https://github.com/nbyl/metio/issues/514)) ([14ed9a9](https://github.com/nbyl/metio/commit/14ed9a961155e8edba2d003356c2d414b2f916fe))
* use project-scoped token for release label sync ([#519](https://github.com/nbyl/metio/issues/519)) ([ff76da4](https://github.com/nbyl/metio/commit/ff76da4340ccc3c4686c99182797b834d74cce7b))
* use user.projectV2 query shape for release label sync ([#522](https://github.com/nbyl/metio/issues/522)) ([d1e6099](https://github.com/nbyl/metio/commit/d1e60992b6e508cda03265f599ba7264a54aebc7))

## [2.0.0](https://github.com/nbyl/metio/compare/v1.7.0...v2.0.0) (2026-09-07)


### ⚠ BREAKING CHANGES

* **backup:** Existing server VMs must run the disk migration in docs/DEPLOYMENT.md before deploying; the data disk mount relocates to /mnt/disks/minecraft with a data/ subdirectory and CurrentInfraVersion is bumped to 4.

### Features

* **api:** add backup catalog and idempotent report ingestion ([#477](https://github.com/nbyl/metio/issues/477)) ([59f7dc8](https://github.com/nbyl/metio/commit/59f7dc84b40fe1cd6e58db548bd4a079b9cc77e2)), closes [#429](https://github.com/nbyl/metio/issues/429)
* **backup:** add per-server backup configuration and runtime ([#435](https://github.com/nbyl/metio/issues/435)) ([7de6548](https://github.com/nbyl/metio/commit/7de65487fa6c789937b54dc314a96fcbbfc25a46))
* **backups:** add backup UI and operational documentation ([#483](https://github.com/nbyl/metio/issues/483)) ([079e25f](https://github.com/nbyl/metio/commit/079e25f8ef2e5ce521cd73b63748a960b1c50ad0))
* **backups:** create new servers from retained backups ([#482](https://github.com/nbyl/metio/issues/482)) ([a0af790](https://github.com/nbyl/metio/commit/a0af790a02702bc19c040b4844293620acb91207))
* **backups:** move backup actions to catalog page and add server filter ([#487](https://github.com/nbyl/metio/issues/487)) ([df754e6](https://github.com/nbyl/metio/commit/df754e6d3c7db9bd149ab9e342babbfedc45b4b9))
* **backups:** preserve deleted-server backups with scheduled cleanup ([#479](https://github.com/nbyl/metio/issues/479)) ([3fe6a08](https://github.com/nbyl/metio/commit/3fe6a087448b5065dd61071f286c6e6d50fc4e20))
* **backups:** restore snapshots onto existing servers ([#481](https://github.com/nbyl/metio/issues/481)) ([1e193cd](https://github.com/nbyl/metio/commit/1e193cd1f844ba5a8d1d6446e7302b8dee939834))
* establish central backup infrastructure ([#434](https://github.com/nbyl/metio/issues/434)) ([d1183be](https://github.com/nbyl/metio/commit/d1183beb9780840486c386192307a627a98ab143))
* improve backup table with server-side pagination, sort, and filter ([#489](https://github.com/nbyl/metio/issues/489)) ([eb9f4c7](https://github.com/nbyl/metio/commit/eb9f4c7bed1fcf35458cc48f2b756a5218e776ca))
* **machine-agent:** report backup manifests to controller ([#478](https://github.com/nbyl/metio/issues/478)) ([c3c85a6](https://github.com/nbyl/metio/commit/c3c85a6f5f3404c43ba547032aaa67bb6aadf0c5))
* **web:** add dark mode with system-default theme switcher ([#468](https://github.com/nbyl/metio/issues/468)) ([278736d](https://github.com/nbyl/metio/commit/278736d8d3745edc644ca4edfd7d718424f86c96))
* **web:** apply shadcn vega preset ([#461](https://github.com/nbyl/metio/issues/461)) ([f1fcbe4](https://github.com/nbyl/metio/commit/f1fcbe4e70fc429d5644d0e896f5b8a1be5d1dcd))
* **web:** migrate layout and shell components ([#457](https://github.com/nbyl/metio/issues/457)) ([7ccfdef](https://github.com/nbyl/metio/commit/7ccfdefc6b6bfb6c2711cfadea93e880a2b5c169))
* **web:** migrate modals and progress to shadcn ([#455](https://github.com/nbyl/metio/issues/455)) ([3df2118](https://github.com/nbyl/metio/commit/3df2118f1b358fb7d4b0ed53550b3d344b274e36))
* **web:** migrate server pages to shadcn and RHF ([#459](https://github.com/nbyl/metio/issues/459)) ([6746e04](https://github.com/nbyl/metio/commit/6746e045ced1a9956b1d19c143509827ee27d5d1))
* **web:** migrate setup wizards to RHF and Zod ([#458](https://github.com/nbyl/metio/issues/458)) ([89db4fd](https://github.com/nbyl/metio/commit/89db4fd0c089b6d923e136a20480da8947f63b1a))
* **web:** migrate to plain shadcn components ([#454](https://github.com/nbyl/metio/issues/454)) ([b7de099](https://github.com/nbyl/metio/commit/b7de0990029de88fa136c740a71720ba9e141882))
* **web:** replace UI primitives with shadcn components ([#452](https://github.com/nbyl/metio/issues/452)) ([ec4639d](https://github.com/nbyl/metio/commit/ec4639ddee5575aac19e6dd73ab2faf45c062d51))


### Bug Fixes

* add deploy_id annotation to force new Cloud Run revisions ([#484](https://github.com/nbyl/metio/issues/484)) ([6b11cd0](https://github.com/nbyl/metio/commit/6b11cd04ae1c05f0f54ba4ad5d3be02a9835caf1))
* **build:** inject release version into controller and machine-agent images ([#467](https://github.com/nbyl/metio/issues/467)) ([d045ab6](https://github.com/nbyl/metio/commit/d045ab6295aef2631d69dd49ba6e010a81490c44))
* correct backup JSON tags and Docker build ordering ([#485](https://github.com/nbyl/metio/issues/485)) ([2f3f288](https://github.com/nbyl/metio/commit/2f3f2884767b13e5e6f1e21f6b418e1d1c74a903)), closes [#423](https://github.com/nbyl/metio/issues/423)
* **css:** don't re-run lightningcss visitor during minify (fix &lt;a ([46acda0](https://github.com/nbyl/metio/commit/46acda08ea9b6578ec50ab289badbde8bf91f8fd))
* **devcontainer:** support GitHub Codespaces ([#502](https://github.com/nbyl/metio/issues/502)) ([1c1e710](https://github.com/nbyl/metio/commit/1c1e7104c9ae343a98d0a8f8046be6e70a34fed2))
* **machine-agent:** restore docker CLI and CA certs on debian 13 ([#448](https://github.com/nbyl/metio/issues/448)) ([7a62b1d](https://github.com/nbyl/metio/commit/7a62b1de691f2053dee0f2df427ebaad5e3ad3db))
* **mc-backup:** pull base image from Docker Hub instead of GHCR ([#508](https://github.com/nbyl/metio/issues/508)) ([a1275f0](https://github.com/nbyl/metio/commit/a1275f09e0d53028f7d1b0003c506d27ae8d2749))
* **restore:** repair restore actions — 404, empty create-from-backup, and save_world ack timeout ([#491](https://github.com/nbyl/metio/issues/491)) ([c7b5acb](https://github.com/nbyl/metio/commit/c7b5acba5a552772549b374c0a0866981be69048))
* **web:** complete light theme migration and harden layout assertions ([#464](https://github.com/nbyl/metio/issues/464)) ([da8c440](https://github.com/nbyl/metio/commit/da8c440903239eb467f2e0676f3941c3e102f3c6))


### Performance Improvements

* **build:** reorder controller layers, add Go cache mounts, drop Cloud Build ([#465](https://github.com/nbyl/metio/issues/465)) ([fcf67d7](https://github.com/nbyl/metio/commit/fcf67d7f87e96587f63337ae5a76c3e8652de7c2))

## [1.7.0](https://github.com/nbyl/metio/compare/v1.6.2...v1.7.0) (2026-08-12)


### Features

* load GCP machine types dynamically ([#370](https://github.com/nbyl/metio/issues/370)) ([#376](https://github.com/nbyl/metio/issues/376)) ([7c63eef](https://github.com/nbyl/metio/commit/7c63eef0a24d491387fde6665c291505e40257ff))
* load GCP zones and regions dynamically ([#374](https://github.com/nbyl/metio/issues/374)) ([7fe92b4](https://github.com/nbyl/metio/commit/7fe92b4753dd8d05ee6d4994bd1377b58d2ae200))
* load Minecraft version list from Mojang ([#371](https://github.com/nbyl/metio/issues/371)) ([f757092](https://github.com/nbyl/metio/commit/f7570928498002ef93e2c8269c1dca83e1016fcf))
* show controller version in web UI ([#377](https://github.com/nbyl/metio/issues/377)) ([9832d56](https://github.com/nbyl/metio/commit/9832d5648e3ca35ccf445f9f0b03c6024b055d5c)), closes [#366](https://github.com/nbyl/metio/issues/366)


### Bug Fixes

* import existing address via pulumi.Import instead of the CLI ([fadf8db](https://github.com/nbyl/metio/commit/fadf8dbb90a814ea7aaaa8ecb4dc89f33700d276))
* import existing address via pulumi.Import instead of the CLI ([e7c7fcb](https://github.com/nbyl/metio/commit/e7c7fcb276707eb7cd492c1eb5904038eb5c4865)), closes [#356](https://github.com/nbyl/metio/issues/356)
* reject duplicate server names on create and rename ([#372](https://github.com/nbyl/metio/issues/372)) ([dc68245](https://github.com/nbyl/metio/commit/dc68245020e407d1b3a18f1505e033f829ac9b01))
* show new servers as starting until the agent reports ([26eeb1d](https://github.com/nbyl/metio/commit/26eeb1dfe9e3c814d37bdbf4c68a2560c1e4eb24))
* show new servers as starting until the agent reports ([a5b7da6](https://github.com/nbyl/metio/commit/a5b7da66de8fa66340687aab589f5b59e96d3f67)), closes [#348](https://github.com/nbyl/metio/issues/348)
* **ui:** show Create Server button when servers exist ([095b0a7](https://github.com/nbyl/metio/commit/095b0a73ecfea3316f647d69c496b9beb874f918))
* **ui:** show Create Server button when servers exist ([da5058a](https://github.com/nbyl/metio/commit/da5058a7710c362d8f4b1e82d34759aca820fcb5)), closes [#359](https://github.com/nbyl/metio/issues/359)

## [1.6.2](https://github.com/nbyl/metio/compare/v1.6.1...v1.6.2) (2026-08-10)


### Bug Fixes

* provider settings screwed by the AI ([c8d7d64](https://github.com/nbyl/metio/commit/c8d7d64c55f211d0a0916df37ea6b1950fe6ea38))
* provider settings screwed by the AI ([4c64d3a](https://github.com/nbyl/metio/commit/4c64d3afb1cb3c454aa3782a0344b185e41514ff))

## [1.6.1](https://github.com/nbyl/metio/compare/v1.6.0...v1.6.1) (2026-08-10)


### Bug Fixes

* **deploy:** drop write-only attributes for OpenTofu &lt;1.11 compatibility ([ddbf4e9](https://github.com/nbyl/metio/commit/ddbf4e97a29c4e56bd57844c9d39ae31de29467b))
* **deploy:** drop write-only attributes for OpenTofu &lt;1.11 compatibility ([4851d51](https://github.com/nbyl/metio/commit/4851d518e370a421516b12b66c19106684931b0f)), closes [#351](https://github.com/nbyl/metio/issues/351)

## [1.6.0](https://github.com/nbyl/metio/compare/v1.5.0...v1.6.0) (2026-08-06)


### Features

* add Dapr sidecar infrastructure for Cloud Run controller ([#255](https://github.com/nbyl/metio/issues/255)) ([66e1621](https://github.com/nbyl/metio/commit/66e162134596bb11e7a73cf1927fc135c617d876))
* add Dapr sidecar infrastructure for Cloud Run controller ([#255](https://github.com/nbyl/metio/issues/255)) ([519c0a3](https://github.com/nbyl/metio/commit/519c0a32a8b8a297fc80adbdf3ed3aca55303740))
* add DB_BACKEND toggle for Dapr state store selection ([#257](https://github.com/nbyl/metio/issues/257)) ([2fa09d8](https://github.com/nbyl/metio/commit/2fa09d8099ed8cb9146843e543933d3b4cb18eac))
* add DB_BACKEND toggle for Dapr state store selection ([#257](https://github.com/nbyl/metio/issues/257)) ([6020911](https://github.com/nbyl/metio/commit/60209112a04069152bc379a74db0b4da83427dd8))
* add JSON tags and Dapr key schema for state store migration ([#258](https://github.com/nbyl/metio/issues/258)) ([eadc24d](https://github.com/nbyl/metio/commit/eadc24d27f652b011612365c5d28e5e6ba2aaf4f))
* add JSON tags and Dapr key schema for state store migration ([#258](https://github.com/nbyl/metio/issues/258)) ([6621c5a](https://github.com/nbyl/metio/commit/6621c5af17edc38fcadcbb2c74f753480e1b0a24))
* add postgres_mode variable and mount connection-string secret ([#325](https://github.com/nbyl/metio/issues/325)) ([c09390b](https://github.com/nbyl/metio/commit/c09390b4808659165726e3cb51a69e3ae3e2b521))
* add postgres_mode variable and mount connection-string secret ([#325](https://github.com/nbyl/metio/issues/325)) ([2a101a5](https://github.com/nbyl/metio/commit/2a101a56e4000fc28f3ca527f56ceb51fd2a8013))
* allow users to update machine agent from UI ([#309](https://github.com/nbyl/metio/issues/309)) ([a9676ae](https://github.com/nbyl/metio/commit/a9676ae69630922257c48b83a8abd820aedaed14))
* allow users to update machine agent from UI ([#309](https://github.com/nbyl/metio/issues/309)) ([912c8c2](https://github.com/nbyl/metio/commit/912c8c24fd7cfd3e8873436cc40558481732b051))
* enable Dapr runtime in local development environment ([d9f38c2](https://github.com/nbyl/metio/commit/d9f38c225c78c8545a46bb38327f27d25b9aaef2))
* enable Dapr runtime in local development environment ([#252](https://github.com/nbyl/metio/issues/252)) ([715f00a](https://github.com/nbyl/metio/commit/715f00a2610b37eb9542e72aa83b33cd97f9daa2))
* hide Players and Uptime stats when server is offline ([7a4441d](https://github.com/nbyl/metio/commit/7a4441d109e6659d129085fc2817fafaec820ee4))
* hide Players and Uptime stats when server is offline ([d197136](https://github.com/nbyl/metio/commit/d1971368fee57c347c2fd9782fa8c585960c2d38))
* implement DaprDB adapter with 24 DB interface methods ([01cc88d](https://github.com/nbyl/metio/commit/01cc88dfcac77ead79709e8661a77b07273b1bc0))
* implement DaprDB adapter with 24 DB interface methods ([#256](https://github.com/nbyl/metio/issues/256)) ([220281a](https://github.com/nbyl/metio/commit/220281a7418eb97a82324438371a116e3dc39e19))
* provision Cloud SQL instance in cloudsql mode ([#327](https://github.com/nbyl/metio/issues/327)) ([92a36ed](https://github.com/nbyl/metio/commit/92a36ed786442eb7d2f5555d149d2eea809d4d92))
* provision Cloud SQL instance in cloudsql mode ([#327](https://github.com/nbyl/metio/issues/327)) ([4cdb596](https://github.com/nbyl/metio/commit/4cdb59688cb53bfc7c1249cf32f10f046a04071b))
* remove Firestore resources and datastore permissions ([#328](https://github.com/nbyl/metio/issues/328)) ([9c55d3e](https://github.com/nbyl/metio/commit/9c55d3e97f7b9f4a8cff52062381ee1958e1c09a))
* remove Firestore resources and datastore permissions ([#328](https://github.com/nbyl/metio/issues/328)) ([b7666dc](https://github.com/nbyl/metio/commit/b7666dc39c01ee6c2457e02a23a6c96b99354ae7))
* replace Datastore emulator with local Postgres container in dev targets ([#326](https://github.com/nbyl/metio/issues/326)) ([eae841d](https://github.com/nbyl/metio/commit/eae841d5d532de98a02ccf076cdc4fa8290428c6))
* replace Datastore emulator with local Postgres container in dev targets ([#326](https://github.com/nbyl/metio/issues/326)) ([6faaa27](https://github.com/nbyl/metio/commit/6faaa27f03abe4eededa60f1940185071d64e1b0))
* run controller on Dapr statestore when DB_BACKEND=dapr ([888251e](https://github.com/nbyl/metio/commit/888251e6d3ad3a279b2e1af76c483100e1be400b))
* switch baked Dapr component to state.postgresql ([#324](https://github.com/nbyl/metio/issues/324)) ([ddb660d](https://github.com/nbyl/metio/commit/ddb660d0b5c02c51fdc9e52ac94dda24c7d0e68a))
* switch baked Dapr component to state.postgresql ([#324](https://github.com/nbyl/metio/issues/324)) ([1dd7950](https://github.com/nbyl/metio/commit/1dd7950d2699f531b19170c7a44264f9aebce437))
* **ui:** remove final-backup step from destroy wizard ([c94058e](https://github.com/nbyl/metio/commit/c94058ed6df11cdd83336da9c3e123efee57a459))
* **ui:** remove final-backup step from destroy wizard ([73a47a7](https://github.com/nbyl/metio/commit/73a47a703ebba9c27286827a3ed913bc560f440f)), closes [#345](https://github.com/nbyl/metio/issues/345)


### Bug Fixes

* add daprd startup probe, remove --app-port for Cloud Run depends_on ([af29ceb](https://github.com/nbyl/metio/commit/af29ceb5f3f5600d5f7f00a6fa70ac408433299d))
* add explicit command for daprd sidecar ([a874b1a](https://github.com/nbyl/metio/commit/a874b1a61dd83a9876a809b58f62162ff4319ebc))
* bump Go base image to 1.26 for Dapr SDK compatibility ([#256](https://github.com/nbyl/metio/issues/256)) ([05d3042](https://github.com/nbyl/metio/commit/05d30421d01f9afaf4969d499de4aa2510b64602))
* correct Dapr firestore component field names and daprd readiness check ([#252](https://github.com/nbyl/metio/issues/252)) ([78fea0c](https://github.com/nbyl/metio/commit/78fea0c1c230f772aefc68c4a31bb48374e26224))
* **deploy:** reference BYO connection string via external secret ([743bfae](https://github.com/nbyl/metio/commit/743bfae0ba0160b89e6e362adffe22500bc7e044))
* **deploy:** reference BYO connection string via external secret ([9908692](https://github.com/nbyl/metio/commit/9908692240d796ae92972cff6e2ef42a69e95048))
* **deploy:** resolve Cloud SQL cloudsql-mode deployment failures ([7e12513](https://github.com/nbyl/metio/commit/7e125138c127d35a43f223dc988ad0185fd6e407))
* **deploy:** resolve Cloud SQL cloudsql-mode deployment failures ([0cec207](https://github.com/nbyl/metio/commit/0cec2070ae380619e9a371ea532b8be113a9445e))
* derive machine-agent registry host from image URL to fix cross-region pulls ([4e3cd21](https://github.com/nbyl/metio/commit/4e3cd21a15c7c3a701ee7e0dc402b5d65f55f1b3))
* expose controller ingress port 8080 for Cloud Run multi-container ([6f23aa8](https://github.com/nbyl/metio/commit/6f23aa8fa3b4bea0a85c6cf8a14fd8e0e4f89bb5))
* remove broken -target from deploy-controller Makefile ([ef0fd84](https://github.com/nbyl/metio/commit/ef0fd84d0a5a99c3f2bb199397c82e73ad0b2878))
* resolve Dapr Firestore project id via ADC auto-detection ([1912294](https://github.com/nbyl/metio/commit/1912294821ff3b21698f981f424f753d87531b06))
* suppress tofu warnings by using secret_data_wo write-only alternative ([6fbc550](https://github.com/nbyl/metio/commit/6fbc550a92eb962a1c1ec08a4a25d518664f2bca))
* treat PlayerDB 400 as not-found and fall back on Mojang errors ([#332](https://github.com/nbyl/metio/issues/332)) ([4e62ed8](https://github.com/nbyl/metio/commit/4e62ed83108a907d1c097766e17418df37af540e))
* treat PlayerDB 400 as not-found and fall back on Mojang errors ([#332](https://github.com/nbyl/metio/issues/332)) ([e544028](https://github.com/nbyl/metio/commit/e544028069ad4a3fabe2e0bbd476f179a65afb1d))
* **web:** support eslint 10 in CI build ([4315c53](https://github.com/nbyl/metio/commit/4315c536aa6008407bdf9d96e05ea497adbcf74b))
* whitelist endpoint returns 500 for new servers ([2eb9662](https://github.com/nbyl/metio/commit/2eb96624541b3629dcca3ff3b2050e916af9221c))
* whitelist endpoint returns 500 for new servers ([897a8b5](https://github.com/nbyl/metio/commit/897a8b54cac3fad64d3fa4838d0121db9e0eb9d9)), closes [#344](https://github.com/nbyl/metio/issues/344)
* wire Dapr emulator env into daprd, unify develop target ([#252](https://github.com/nbyl/metio/issues/252)) ([5c647e6](https://github.com/nbyl/metio/commit/5c647e659cc9b6ec85e6ec9d4a2edf0bd488a58c))

## [1.5.0](https://github.com/nbyl/metio/compare/v1.4.1...v1.5.0) (2026-07-07)


### Features

* implement controller agent API with HMAC-JWT per-server token auth ([#300](https://github.com/nbyl/metio/issues/300)) ([6d4da14](https://github.com/nbyl/metio/commit/6d4da146b8c9142c8c3e15222db6031cf89e5326))
* implement controller agent API with HMAC-JWT per-server token auth ([#300](https://github.com/nbyl/metio/issues/300)) ([7b02679](https://github.com/nbyl/metio/commit/7b02679b95a3a080769334b0418001cede302e2a))


### Bug Fixes

* save image tags to build/*.txt and push to Artifact Registry in local docker targets ([391775d](https://github.com/nbyl/metio/commit/391775d3c504513e512b737258ca9db2a79fbeb5))
* split docker image targets for CI and deploy, run image builds on PRs ([250a4d9](https://github.com/nbyl/metio/commit/250a4d93c54d8f29a08446024658baf98e7a8991))
* split docker image targets for CI and deploy, run image builds on PRs ([2280df1](https://github.com/nbyl/metio/commit/2280df1625b703925ab47efed1df2fecce9d3a5e))

## [1.4.1](https://github.com/nbyl/metio/compare/v1.4.0...v1.4.1) (2026-07-02)


### Bug Fixes

* pass ExistingAddress through Cloud Tasks handler ([2d43d51](https://github.com/nbyl/metio/commit/2d43d51c7f75f23b2b7e6ed1e97241d7eff5e656))
* pass ExistingAddress through Cloud Tasks handler ([68e7d47](https://github.com/nbyl/metio/commit/68e7d477e190c946c778f3959b943799472723c6))

## [1.4.0](https://github.com/nbyl/metio/compare/v1.3.0...v1.4.0) (2026-07-02)


### Features

* support importing existing static IP addresses when provisioning servers ([7dd8dde](https://github.com/nbyl/metio/commit/7dd8dde832f09de29e554a2455a50634fc8c11d4))
* support importing existing static IP addresses when provisioning servers ([4e7b477](https://github.com/nbyl/metio/commit/4e7b47749f22b2b5577d40302bdf78a25156e0e2)), closes [#293](https://github.com/nbyl/metio/issues/293)


### Bug Fixes

* refresh stack after import and enable code generation for pulumi import ([e87e018](https://github.com/nbyl/metio/commit/e87e0189cba139b97584eae2e3c328da78fa90df))

## [1.3.0](https://github.com/nbyl/metio/compare/v1.2.0...v1.3.0) (2026-06-30)


### Features

* replace go-semantic-release with release-please ([ecc18e2](https://github.com/nbyl/metio/commit/ecc18e29b335729339b99777297ce8b7be04bc66))
* replace go-semantic-release with release-please ([e98dc9c](https://github.com/nbyl/metio/commit/e98dc9cd695e8cbabcc1ad824534cbaac27664fa))


### Bug Fixes

* add environment: main to release job ([7947e55](https://github.com/nbyl/metio/commit/7947e558482ea9097acf3db93e30c9319fba66c6))
* add environment: main to release job ([3577d26](https://github.com/nbyl/metio/commit/3577d263239b72f3d8f35e707b9935b9c0c103a9))
* add release-please extra-files for variables.tf ([5d8f990](https://github.com/nbyl/metio/commit/5d8f990bbcc9e49f3929c5e03c1b99ef5fb41e0c))
* add release-please extra-files for variables.tf ([a06d54a](https://github.com/nbyl/metio/commit/a06d54a93153bb60f4141328f015ea0f1230e95a))
* rename release-please config and add checkout step ([9f5d3bd](https://github.com/nbyl/metio/commit/9f5d3bd2b0cb50b6b51471f66b4d075cf59b85b7))
* rename release-please config file and add checkout step ([9672939](https://github.com/nbyl/metio/commit/9672939ed1a1a1d3b68c5fd8d9fde142f24f3522))
* use RELEASE_PLEASE_TOKEN secret for release-please-action ([0d05bd0](https://github.com/nbyl/metio/commit/0d05bd04843938c8a33cf90676b33841e3aa3977))
* use RELEASE_PLEASE_TOKEN secret for release-please-action ([e6fe4f3](https://github.com/nbyl/metio/commit/e6fe4f3e750c2215baf47e26a5921e74c10f85d3))
