package main

import (
	"stevenandrewcarter/terraform-provider-containerd/containerd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return containerd.Provider()
		},
	})
}

// func main() {
// 	if err := redisExample(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func redisExample() error {
// 	client, err := containerd.New("/run/containerd/containerd.sock")
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	ctx := namespaces.WithNamespace(context.Background(), "example")
// 	image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Successfully pulled %s image\n", image.Name())

// 	container, err := client.NewContainer(
// 		ctx,
// 		"redis-server",
// 		containerd.WithNewSnapshot("redis-server-snapshot", image),
// 		containerd.WithNewSpec(oci.WithImageConfig(image)),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer container.Delete(ctx, containerd.WithSnapshotCleanup)
// 	log.Printf("Successfully created container with ID %s and snapshot with ID redis-server-snapshot", container.ID())

// 	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
// 	if err != nil {
// 		return err
// 	}
// 	defer task.Delete(ctx)

// 	exitStatusC, err := task.Wait(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if err := task.Start(ctx); err != nil {
// 		return err
// 	}

// 	time.Sleep(3 * time.Second)

// 	if err := task.Kill(ctx, syscall.SIGTERM); err != nil {
// 		return err
// 	}

// 	status := <-exitStatusC
// 	code, _, err := status.Result()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("redis-server exited with status: %d\n", code)

// 	return nil
// }
