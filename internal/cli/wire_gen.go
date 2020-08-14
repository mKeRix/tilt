// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cli

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/jonboulle/clockwork"
	"github.com/tilt-dev/wmclient/pkg/dirs"
	trace2 "go.opentelemetry.io/otel/sdk/trace"
	version2 "k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/tools/clientcmd/api"

	"github.com/tilt-dev/tilt/internal/analytics"
	"github.com/tilt-dev/tilt/internal/build"
	"github.com/tilt-dev/tilt/internal/cloud"
	"github.com/tilt-dev/tilt/internal/cloud/cloudurl"
	"github.com/tilt-dev/tilt/internal/container"
	"github.com/tilt-dev/tilt/internal/containerupdate"
	"github.com/tilt-dev/tilt/internal/docker"
	"github.com/tilt-dev/tilt/internal/dockercompose"
	"github.com/tilt-dev/tilt/internal/dockerfile"
	"github.com/tilt-dev/tilt/internal/engine"
	analytics2 "github.com/tilt-dev/tilt/internal/engine/analytics"
	"github.com/tilt-dev/tilt/internal/engine/buildcontrol"
	"github.com/tilt-dev/tilt/internal/engine/configs"
	"github.com/tilt-dev/tilt/internal/engine/dcwatch"
	"github.com/tilt-dev/tilt/internal/engine/dockerprune"
	"github.com/tilt-dev/tilt/internal/engine/exit"
	"github.com/tilt-dev/tilt/internal/engine/fswatch"
	"github.com/tilt-dev/tilt/internal/engine/k8srollout"
	"github.com/tilt-dev/tilt/internal/engine/k8swatch"
	"github.com/tilt-dev/tilt/internal/engine/local"
	"github.com/tilt-dev/tilt/internal/engine/portforward"
	"github.com/tilt-dev/tilt/internal/engine/runtimelog"
	"github.com/tilt-dev/tilt/internal/engine/telemetry"
	"github.com/tilt-dev/tilt/internal/feature"
	"github.com/tilt-dev/tilt/internal/hud"
	"github.com/tilt-dev/tilt/internal/hud/prompt"
	"github.com/tilt-dev/tilt/internal/hud/server"
	"github.com/tilt-dev/tilt/internal/k8s"
	"github.com/tilt-dev/tilt/internal/store"
	"github.com/tilt-dev/tilt/internal/synclet/sidecar"
	"github.com/tilt-dev/tilt/internal/tiltfile"
	"github.com/tilt-dev/tilt/internal/tiltfile/config"
	"github.com/tilt-dev/tilt/internal/tiltfile/k8scontext"
	"github.com/tilt-dev/tilt/internal/tiltfile/version"
	"github.com/tilt-dev/tilt/internal/token"
	"github.com/tilt-dev/tilt/internal/tracer"
	"github.com/tilt-dev/tilt/pkg/model"
)

// Injectors from wire.go:

func wireTiltfileResult(ctx context.Context, analytics2 *analytics.TiltAnalytics, subcommand model.TiltSubcommand) (cmdTiltfileResultDeps, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return cmdTiltfileResultDeps{}, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return cmdTiltfileResultDeps{}, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	extension := k8scontext.NewExtension(kubeContext, env)
	tiltBuild := provideTiltInfo()
	versionExtension := version.NewExtension(tiltBuild)
	configExtension := config.NewExtension(subcommand)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	dockerComposeClient := dockercompose.NewDockerComposeClient(localEnv)
	modelWebHost := provideWebHost()
	defaults := _wireDefaultsValue
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics2, client, extension, versionExtension, configExtension, dockerComposeClient, modelWebHost, defaults, env)
	cliCmdTiltfileResultDeps := newTiltfileResultDeps(tiltfileLoader)
	return cliCmdTiltfileResultDeps, nil
}

var (
	_wireDefaultsValue = feature.MainDefaults
)

func wireDockerPrune(ctx context.Context, analytics2 *analytics.TiltAnalytics, subcommand model.TiltSubcommand) (dpDeps, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return dpDeps{}, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return dpDeps{}, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	localClient := docker.ProvideLocalCli(ctx, localEnv)
	clusterClient, err := docker.ProvideClusterCli(ctx, localEnv, clusterEnv, localClient)
	if err != nil {
		return dpDeps{}, err
	}
	switchCli := docker.ProvideSwitchCli(clusterClient, localClient)
	extension := k8scontext.NewExtension(kubeContext, env)
	tiltBuild := provideTiltInfo()
	versionExtension := version.NewExtension(tiltBuild)
	configExtension := config.NewExtension(subcommand)
	dockerComposeClient := dockercompose.NewDockerComposeClient(localEnv)
	modelWebHost := provideWebHost()
	defaults := _wireDefaultsValue
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics2, client, extension, versionExtension, configExtension, dockerComposeClient, modelWebHost, defaults, env)
	cliDpDeps := newDPDeps(switchCli, tiltfileLoader)
	return cliDpDeps, nil
}

func wireCmdUp(ctx context.Context, analytics3 *analytics.TiltAnalytics, cmdTags analytics2.CmdTags, subcommand model.TiltSubcommand) (CmdUpDeps, error) {
	reducer := _wireReducerValue
	storeLogActionsFlag := provideLogActions()
	storeStore := store.NewStore(reducer, storeLogActionsFlag)
	v := provideClock()
	renderer := hud.NewRenderer(v)
	modelWebHost := provideWebHost()
	modelWebPort := provideWebPort()
	webURL, err := provideWebURL(modelWebHost, modelWebPort)
	if err != nil {
		return CmdUpDeps{}, err
	}
	headsUpDisplay := hud.NewHud(renderer, webURL, analytics3)
	stdout := hud.ProvideStdout()
	incrementalPrinter := hud.NewIncrementalPrinter(stdout)
	terminalStream := hud.NewTerminalStream(incrementalPrinter, storeStore)
	openInput := _wireOpenInputValue
	openURL := _wireOpenURLValue
	terminalPrompt := prompt.NewTerminalPrompt(analytics3, openInput, openURL, stdout, modelWebHost, webURL)
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return CmdUpDeps{}, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return CmdUpDeps{}, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	ownerFetcher := k8s.ProvideOwnerFetcher(client)
	podWatcher := k8swatch.NewPodWatcher(client, ownerFetcher)
	serviceWatcher := k8swatch.NewServiceWatcher(client, ownerFetcher)
	podLogManager := runtimelog.NewPodLogManager(client)
	controller := portforward.NewController(client)
	fsWatcherMaker := fswatch.ProvideFsWatcherMaker()
	timerMaker := fswatch.ProvideTimerMaker()
	watchManager := fswatch.NewWatchManager(fsWatcherMaker, timerMaker)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	localClient := docker.ProvideLocalCli(ctx, localEnv)
	clusterClient, err := docker.ProvideClusterCli(ctx, localEnv, clusterEnv, localClient)
	if err != nil {
		return CmdUpDeps{}, err
	}
	switchCli := docker.ProvideSwitchCli(clusterClient, localClient)
	dockerUpdater := containerupdate.NewDockerUpdater(switchCli)
	syncletImageRef, err := sidecar.ProvideSyncletImageRef(ctx)
	if err != nil {
		return CmdUpDeps{}, err
	}
	syncletManager := containerupdate.NewSyncletManager(client, syncletImageRef)
	syncletUpdater := containerupdate.NewSyncletUpdater(syncletManager)
	execUpdater := containerupdate.NewExecUpdater(client)
	buildcontrolUpdateModeFlag := provideUpdateModeFlag()
	updateMode, err := buildcontrol.ProvideUpdateMode(buildcontrolUpdateModeFlag, env, runtime)
	if err != nil {
		return CmdUpDeps{}, err
	}
	clock := build.ProvideClock()
	liveUpdateBuildAndDeployer := engine.NewLiveUpdateBuildAndDeployer(dockerUpdater, syncletUpdater, execUpdater, updateMode, env, runtime, clock)
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(switchCli, labels)
	dockerBuilder := build.DefaultDockerBuilder(dockerImageBuilder)
	execCustomBuilder := build.NewExecCustomBuilder(switchCli, clock)
	clusterName := k8s.ProvideClusterName(ctx, apiConfig)
	kindLoader := engine.NewKINDLoader(env, clusterName)
	syncletContainer := sidecar.ProvideSyncletContainer(syncletImageRef)
	imageBuildAndDeployer := engine.NewImageBuildAndDeployer(dockerBuilder, execCustomBuilder, client, env, analytics3, updateMode, clock, runtime, kindLoader, syncletContainer)
	dockerComposeClient := dockercompose.NewDockerComposeClient(localEnv)
	imageBuilder := engine.NewImageBuilder(dockerBuilder, execCustomBuilder, updateMode)
	dockerComposeBuildAndDeployer := engine.NewDockerComposeBuildAndDeployer(dockerComposeClient, switchCli, imageBuilder, clock)
	localTargetBuildAndDeployer := engine.NewLocalTargetBuildAndDeployer(clock)
	buildOrder := engine.DefaultBuildOrder(liveUpdateBuildAndDeployer, imageBuildAndDeployer, dockerComposeBuildAndDeployer, localTargetBuildAndDeployer, updateMode, env, runtime)
	spanCollector := tracer.NewSpanCollector(ctx)
	traceTracer, err := tracer.InitOpenTelemetry(ctx, spanCollector)
	if err != nil {
		return CmdUpDeps{}, err
	}
	compositeBuildAndDeployer := engine.NewCompositeBuildAndDeployer(buildOrder, traceTracer)
	buildController := engine.NewBuildController(compositeBuildAndDeployer)
	extension := k8scontext.NewExtension(kubeContext, env)
	tiltBuild := provideTiltInfo()
	versionExtension := version.NewExtension(tiltBuild)
	configExtension := config.NewExtension(subcommand)
	defaults := _wireDefaultsValue
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics3, client, extension, versionExtension, configExtension, dockerComposeClient, modelWebHost, defaults, env)
	configsController := configs.NewConfigsController(tiltfileLoader, switchCli)
	eventWatcher := dcwatch.NewEventWatcher(dockerComposeClient, localClient)
	dockerComposeLogManager := runtimelog.NewDockerComposeLogManager(dockerComposeClient)
	profilerManager := engine.NewProfilerManager()
	analyticsReporter := analytics2.ProvideAnalyticsReporter(analytics3, storeStore, client, env, subcommand)
	webMode, err := provideWebMode(tiltBuild)
	if err != nil {
		return CmdUpDeps{}, err
	}
	webVersion := provideWebVersion(tiltBuild)
	assetsServer, err := provideAssetServer(webMode, webVersion)
	if err != nil {
		return CmdUpDeps{}, err
	}
	httpClient := cloud.ProvideHttpClient()
	address := cloudurl.ProvideAddress()
	snapshotUploader := cloud.NewSnapshotUploader(httpClient, address)
	headsUpServer, err := server.ProvideHeadsUpServer(ctx, storeStore, assetsServer, analytics3, snapshotUploader)
	if err != nil {
		return CmdUpDeps{}, err
	}
	headsUpServerController := server.ProvideHeadsUpServerController(modelWebHost, modelWebPort, headsUpServer, assetsServer, webURL)
	analyticsUpdater := analytics2.NewAnalyticsUpdater(analytics3, cmdTags)
	eventWatchManager := k8swatch.NewEventWatchManager(client, ownerFetcher)
	clockworkClock := clockwork.NewRealClock()
	cloudStatusManager := cloud.NewStatusManager(httpClient, clockworkClock)
	updateUploader := cloud.NewUpdateUploader(httpClient, address)
	dockerPruner := dockerprune.NewDockerPruner(switchCli)
	telemetryController := telemetry.NewController(clock, spanCollector)
	execer := local.ProvideExecer()
	localController := local.NewController(execer)
	podMonitor := k8srollout.NewPodMonitor()
	exitController := exit.NewController()
	v2 := engine.ProvideSubscribers(headsUpDisplay, terminalStream, terminalPrompt, podWatcher, serviceWatcher, podLogManager, controller, watchManager, buildController, configsController, eventWatcher, dockerComposeLogManager, profilerManager, syncletManager, analyticsReporter, headsUpServerController, analyticsUpdater, eventWatchManager, cloudStatusManager, updateUploader, dockerPruner, telemetryController, localController, podMonitor, exitController)
	upper := engine.NewUpper(ctx, storeStore, v2)
	windmillDir, err := dirs.UseWindmillDir()
	if err != nil {
		return CmdUpDeps{}, err
	}
	tokenToken, err := token.GetOrCreateToken(windmillDir)
	if err != nil {
		return CmdUpDeps{}, err
	}
	cmdUpDeps := CmdUpDeps{
		Upper:        upper,
		TiltBuild:    tiltBuild,
		Token:        tokenToken,
		CloudAddress: address,
		Store:        storeStore,
		Prompt:       terminalPrompt,
	}
	return cmdUpDeps, nil
}

var (
	_wireReducerValue   = engine.UpperReducer
	_wireOpenInputValue = prompt.OpenInput(prompt.TTYOpen)
	_wireOpenURLValue   = prompt.OpenURL(prompt.BrowserOpen)
	_wireLabelsValue    = dockerfile.Labels{}
)

func wireCmdCI(ctx context.Context, analytics3 *analytics.TiltAnalytics, subcommand model.TiltSubcommand) (CmdCIDeps, error) {
	reducer := _wireReducerValue
	storeLogActionsFlag := provideLogActions()
	storeStore := store.NewStore(reducer, storeLogActionsFlag)
	v := provideClock()
	renderer := hud.NewRenderer(v)
	modelWebHost := provideWebHost()
	modelWebPort := provideWebPort()
	webURL, err := provideWebURL(modelWebHost, modelWebPort)
	if err != nil {
		return CmdCIDeps{}, err
	}
	headsUpDisplay := hud.NewHud(renderer, webURL, analytics3)
	stdout := hud.ProvideStdout()
	incrementalPrinter := hud.NewIncrementalPrinter(stdout)
	terminalStream := hud.NewTerminalStream(incrementalPrinter, storeStore)
	openInput := _wireOpenInputValue
	openURL := _wireOpenURLValue
	terminalPrompt := prompt.NewTerminalPrompt(analytics3, openInput, openURL, stdout, modelWebHost, webURL)
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return CmdCIDeps{}, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return CmdCIDeps{}, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	ownerFetcher := k8s.ProvideOwnerFetcher(client)
	podWatcher := k8swatch.NewPodWatcher(client, ownerFetcher)
	serviceWatcher := k8swatch.NewServiceWatcher(client, ownerFetcher)
	podLogManager := runtimelog.NewPodLogManager(client)
	controller := portforward.NewController(client)
	fsWatcherMaker := fswatch.ProvideFsWatcherMaker()
	timerMaker := fswatch.ProvideTimerMaker()
	watchManager := fswatch.NewWatchManager(fsWatcherMaker, timerMaker)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	localClient := docker.ProvideLocalCli(ctx, localEnv)
	clusterClient, err := docker.ProvideClusterCli(ctx, localEnv, clusterEnv, localClient)
	if err != nil {
		return CmdCIDeps{}, err
	}
	switchCli := docker.ProvideSwitchCli(clusterClient, localClient)
	dockerUpdater := containerupdate.NewDockerUpdater(switchCli)
	syncletImageRef, err := sidecar.ProvideSyncletImageRef(ctx)
	if err != nil {
		return CmdCIDeps{}, err
	}
	syncletManager := containerupdate.NewSyncletManager(client, syncletImageRef)
	syncletUpdater := containerupdate.NewSyncletUpdater(syncletManager)
	execUpdater := containerupdate.NewExecUpdater(client)
	buildcontrolUpdateModeFlag := provideUpdateModeFlag()
	updateMode, err := buildcontrol.ProvideUpdateMode(buildcontrolUpdateModeFlag, env, runtime)
	if err != nil {
		return CmdCIDeps{}, err
	}
	clock := build.ProvideClock()
	liveUpdateBuildAndDeployer := engine.NewLiveUpdateBuildAndDeployer(dockerUpdater, syncletUpdater, execUpdater, updateMode, env, runtime, clock)
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(switchCli, labels)
	dockerBuilder := build.DefaultDockerBuilder(dockerImageBuilder)
	execCustomBuilder := build.NewExecCustomBuilder(switchCli, clock)
	clusterName := k8s.ProvideClusterName(ctx, apiConfig)
	kindLoader := engine.NewKINDLoader(env, clusterName)
	syncletContainer := sidecar.ProvideSyncletContainer(syncletImageRef)
	imageBuildAndDeployer := engine.NewImageBuildAndDeployer(dockerBuilder, execCustomBuilder, client, env, analytics3, updateMode, clock, runtime, kindLoader, syncletContainer)
	dockerComposeClient := dockercompose.NewDockerComposeClient(localEnv)
	imageBuilder := engine.NewImageBuilder(dockerBuilder, execCustomBuilder, updateMode)
	dockerComposeBuildAndDeployer := engine.NewDockerComposeBuildAndDeployer(dockerComposeClient, switchCli, imageBuilder, clock)
	localTargetBuildAndDeployer := engine.NewLocalTargetBuildAndDeployer(clock)
	buildOrder := engine.DefaultBuildOrder(liveUpdateBuildAndDeployer, imageBuildAndDeployer, dockerComposeBuildAndDeployer, localTargetBuildAndDeployer, updateMode, env, runtime)
	spanCollector := tracer.NewSpanCollector(ctx)
	traceTracer, err := tracer.InitOpenTelemetry(ctx, spanCollector)
	if err != nil {
		return CmdCIDeps{}, err
	}
	compositeBuildAndDeployer := engine.NewCompositeBuildAndDeployer(buildOrder, traceTracer)
	buildController := engine.NewBuildController(compositeBuildAndDeployer)
	extension := k8scontext.NewExtension(kubeContext, env)
	tiltBuild := provideTiltInfo()
	versionExtension := version.NewExtension(tiltBuild)
	configExtension := config.NewExtension(subcommand)
	defaults := _wireDefaultsValue
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics3, client, extension, versionExtension, configExtension, dockerComposeClient, modelWebHost, defaults, env)
	configsController := configs.NewConfigsController(tiltfileLoader, switchCli)
	eventWatcher := dcwatch.NewEventWatcher(dockerComposeClient, localClient)
	dockerComposeLogManager := runtimelog.NewDockerComposeLogManager(dockerComposeClient)
	profilerManager := engine.NewProfilerManager()
	analyticsReporter := analytics2.ProvideAnalyticsReporter(analytics3, storeStore, client, env, subcommand)
	webMode, err := provideWebMode(tiltBuild)
	if err != nil {
		return CmdCIDeps{}, err
	}
	webVersion := provideWebVersion(tiltBuild)
	assetsServer, err := provideAssetServer(webMode, webVersion)
	if err != nil {
		return CmdCIDeps{}, err
	}
	httpClient := cloud.ProvideHttpClient()
	address := cloudurl.ProvideAddress()
	snapshotUploader := cloud.NewSnapshotUploader(httpClient, address)
	headsUpServer, err := server.ProvideHeadsUpServer(ctx, storeStore, assetsServer, analytics3, snapshotUploader)
	if err != nil {
		return CmdCIDeps{}, err
	}
	headsUpServerController := server.ProvideHeadsUpServerController(modelWebHost, modelWebPort, headsUpServer, assetsServer, webURL)
	cmdTags := _wireCmdTagsValue
	analyticsUpdater := analytics2.NewAnalyticsUpdater(analytics3, cmdTags)
	eventWatchManager := k8swatch.NewEventWatchManager(client, ownerFetcher)
	clockworkClock := clockwork.NewRealClock()
	cloudStatusManager := cloud.NewStatusManager(httpClient, clockworkClock)
	updateUploader := cloud.NewUpdateUploader(httpClient, address)
	dockerPruner := dockerprune.NewDockerPruner(switchCli)
	telemetryController := telemetry.NewController(clock, spanCollector)
	execer := local.ProvideExecer()
	localController := local.NewController(execer)
	podMonitor := k8srollout.NewPodMonitor()
	exitController := exit.NewController()
	v2 := engine.ProvideSubscribers(headsUpDisplay, terminalStream, terminalPrompt, podWatcher, serviceWatcher, podLogManager, controller, watchManager, buildController, configsController, eventWatcher, dockerComposeLogManager, profilerManager, syncletManager, analyticsReporter, headsUpServerController, analyticsUpdater, eventWatchManager, cloudStatusManager, updateUploader, dockerPruner, telemetryController, localController, podMonitor, exitController)
	upper := engine.NewUpper(ctx, storeStore, v2)
	windmillDir, err := dirs.UseWindmillDir()
	if err != nil {
		return CmdCIDeps{}, err
	}
	tokenToken, err := token.GetOrCreateToken(windmillDir)
	if err != nil {
		return CmdCIDeps{}, err
	}
	cmdCIDeps := CmdCIDeps{
		Upper:        upper,
		TiltBuild:    tiltBuild,
		Token:        tokenToken,
		CloudAddress: address,
		Store:        storeStore,
	}
	return cmdCIDeps, nil
}

var (
	_wireCmdTagsValue = analytics2.CmdTags(map[string]string{})
)

func wireKubeContext(ctx context.Context) (k8s.KubeContext, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return "", err
	}
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return "", err
	}
	return kubeContext, nil
}

func wireKubeConfig(ctx context.Context) (*api.Config, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	return apiConfig, nil
}

func wireEnv(ctx context.Context) (k8s.Env, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return "", err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	return env, nil
}

func wireNamespace(ctx context.Context) (k8s.Namespace, error) {
	clientConfig := k8s.ProvideClientConfig()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	return namespace, nil
}

func wireClusterName(ctx context.Context) (k8s.ClusterName, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return "", err
	}
	clusterName := k8s.ProvideClusterName(ctx, apiConfig)
	return clusterName, nil
}

func wireRuntime(ctx context.Context) (container.Runtime, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return "", err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return "", err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	return runtime, nil
}

func wireK8sClient(ctx context.Context) (k8s.Client, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return nil, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	return client, nil
}

func wireK8sVersion(ctx context.Context) (*version2.Info, error) {
	clientConfig := k8s.ProvideClientConfig()
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	info, err := k8s.ProvideServerVersion(clientsetOrError)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func wireDockerClusterClient(ctx context.Context) (docker.ClusterClient, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return nil, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	localClient := docker.ProvideLocalCli(ctx, localEnv)
	clusterClient, err := docker.ProvideClusterCli(ctx, localEnv, clusterEnv, localClient)
	if err != nil {
		return nil, err
	}
	return clusterClient, nil
}

func wireDockerLocalClient(ctx context.Context) (docker.LocalClient, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return nil, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	localClient := docker.ProvideLocalCli(ctx, localEnv)
	return localClient, nil
}

func wireDownDeps(ctx context.Context, tiltAnalytics *analytics.TiltAnalytics, subcommand model.TiltSubcommand) (DownDeps, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return DownDeps{}, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return DownDeps{}, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	extension := k8scontext.NewExtension(kubeContext, env)
	tiltBuild := provideTiltInfo()
	versionExtension := version.NewExtension(tiltBuild)
	configExtension := config.NewExtension(subcommand)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	dockerComposeClient := dockercompose.NewDockerComposeClient(localEnv)
	modelWebHost := provideWebHost()
	defaults := _wireDefaultsValue
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(tiltAnalytics, client, extension, versionExtension, configExtension, dockerComposeClient, modelWebHost, defaults, env)
	downDeps := ProvideDownDeps(tiltfileLoader, dockerComposeClient, client)
	return downDeps, nil
}

func wireLogsDeps(ctx context.Context, tiltAnalytics *analytics.TiltAnalytics, subcommand model.TiltSubcommand) (LogsDeps, error) {
	modelWebHost := provideWebHost()
	modelWebPort := provideWebPort()
	webURL, err := provideWebURL(modelWebHost, modelWebPort)
	if err != nil {
		return LogsDeps{}, err
	}
	stdout := hud.ProvideStdout()
	incrementalPrinter := hud.NewIncrementalPrinter(stdout)
	logsDeps := ProvideLogsDeps(webURL, incrementalPrinter)
	return logsDeps, nil
}

func wireDumpImageDeployRefDeps(ctx context.Context) (DumpImageDeployRefDeps, error) {
	clientConfig := k8s.ProvideClientConfig()
	apiConfig, err := k8s.ProvideKubeConfig(clientConfig)
	if err != nil {
		return DumpImageDeployRefDeps{}, err
	}
	env := k8s.ProvideEnv(ctx, apiConfig)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubeContext, err := k8s.ProvideKubeContext(apiConfig)
	if err != nil {
		return DumpImageDeployRefDeps{}, err
	}
	int2 := provideKubectlLogLevel()
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext, int2)
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, kubectlRunner, minikubeClient, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	clusterEnv := docker.ProvideClusterEnv(ctx, env, runtime, minikubeClient)
	localEnv := docker.ProvideLocalEnv(ctx, clusterEnv)
	localClient := docker.ProvideLocalCli(ctx, localEnv)
	clusterClient, err := docker.ProvideClusterCli(ctx, localEnv, clusterEnv, localClient)
	if err != nil {
		return DumpImageDeployRefDeps{}, err
	}
	switchCli := docker.ProvideSwitchCli(clusterClient, localClient)
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(switchCli, labels)
	dockerBuilder := build.DefaultDockerBuilder(dockerImageBuilder)
	dumpImageDeployRefDeps := DumpImageDeployRefDeps{
		DockerBuilder: dockerBuilder,
		DockerClient:  switchCli,
	}
	return dumpImageDeployRefDeps, nil
}

// wire.go:

var K8sWireSet = wire.NewSet(k8s.ProvideEnv, k8s.ProvideClusterName, k8s.ProvideKubeContext, k8s.ProvideKubeConfig, k8s.ProvideClientConfig, k8s.ProvideClientset, k8s.ProvideRESTConfig, k8s.ProvidePortForwardClient, k8s.ProvideConfigNamespace, k8s.ProvideKubectlRunner, k8s.ProvideContainerRuntime, k8s.ProvideServerVersion, k8s.ProvideK8sClient, k8s.ProvideOwnerFetcher)

var BaseWireSet = wire.NewSet(
	K8sWireSet, tiltfile.WireSet, provideKubectlLogLevel, docker.SwitchWireSet, dockercompose.NewDockerComposeClient, clockwork.NewRealClock, engine.DeployerWireSet, runtimelog.NewPodLogManager, portforward.NewController, engine.NewBuildController, local.ProvideExecer, local.NewController, k8swatch.NewPodWatcher, k8swatch.NewServiceWatcher, k8swatch.NewEventWatchManager, configs.NewConfigsController, telemetry.NewController, dcwatch.NewEventWatcher, runtimelog.NewDockerComposeLogManager, engine.NewProfilerManager, cloud.WireSet, cloudurl.ProvideAddress, k8srollout.NewPodMonitor, telemetry.NewStartTracker, exit.NewController, provideClock, hud.WireSet, prompt.WireSet, provideLogActions, store.NewStore, wire.Bind(new(store.RStore), new(*store.Store)), dockerprune.NewDockerPruner, provideTiltInfo, engine.ProvideSubscribers, engine.NewUpper, analytics2.NewAnalyticsUpdater, analytics2.ProvideAnalyticsReporter, provideUpdateModeFlag, fswatch.NewWatchManager, fswatch.ProvideFsWatcherMaker, fswatch.ProvideTimerMaker, provideWebVersion,
	provideWebMode,
	provideWebURL,
	provideWebPort,
	provideWebHost,
	server.ProvideHeadsUpServer, provideAssetServer, server.ProvideHeadsUpServerController, tracer.NewSpanCollector, wire.Bind(new(trace2.SpanProcessor), new(*tracer.SpanCollector)), wire.Bind(new(tracer.SpanSource), new(*tracer.SpanCollector)), dirs.UseWindmillDir, token.GetOrCreateToken, engine.NewKINDLoader, wire.Value(feature.MainDefaults),
)

type CmdUpDeps struct {
	Upper        engine.Upper
	TiltBuild    model.TiltBuild
	Token        token.Token
	CloudAddress cloudurl.Address
	Store        *store.Store
	Prompt       *prompt.TerminalPrompt
}

type CmdCIDeps struct {
	Upper        engine.Upper
	TiltBuild    model.TiltBuild
	Token        token.Token
	CloudAddress cloudurl.Address
	Store        *store.Store
}

type DownDeps struct {
	tfl      tiltfile.TiltfileLoader
	dcClient dockercompose.DockerComposeClient
	kClient  k8s.Client
}

func ProvideDownDeps(
	tfl tiltfile.TiltfileLoader,
	dcClient dockercompose.DockerComposeClient,
	kClient k8s.Client) DownDeps {
	return DownDeps{
		tfl:      tfl,
		dcClient: dcClient,
		kClient:  kClient,
	}
}

type LogsDeps struct {
	url     model.WebURL
	printer *hud.IncrementalPrinter
}

func ProvideLogsDeps(u model.WebURL, p *hud.IncrementalPrinter) LogsDeps {
	return LogsDeps{
		url:     u,
		printer: p,
	}
}

func provideClock() func() time.Time {
	return time.Now
}

type DumpImageDeployRefDeps struct {
	DockerBuilder build.DockerBuilder
	DockerClient  docker.Client
}
