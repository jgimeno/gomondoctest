package gomondoctest

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

//Inspired by https://developers.almamedia.fi/2014/painless-mongodb-testing-with-docker-and-golang/

const mongoImage = "mongo"

type Gomondoc struct {
	t           *testing.T
	containerID string
}

func NewGomondoc(t *testing.T) *Gomondoc {
	checkDocker(t)

	if ok, err := dockerHasImage(mongoImage); !ok || err != nil {
		if err != nil {
			t.Fatalf("Error checking that the docker image %v is installed in the system.", mongoImage)
		}

		installDockerImage(t)
	}

	return &Gomondoc{t, ""}
}

func checkDocker(t *testing.T) {
	if _, err := exec.LookPath("docker"); err != nil {
		t.Skip("Docker is not installed on the system.")
	}
}

func (e *Gomondoc) RunMongo() {
	log.Printf("Executing docker mongo image.")

	out, err := exec.Command("docker", "run", "--name", "some-mongo", "-p", "27017:27017", "-d", "mongo").Output()

	if err != nil {
		e.t.Fatal("Error running mongo image.", err)
	}

	e.containerID = strings.TrimSpace(string(out))

	if e.containerID == "" {
		e.t.Fatal("Error getting id of docker container.")
	}
}

func (md *Gomondoc) StopMongo() {
	log.Printf("Stopping docker image.")
	out, err := exec.Command("docker", "stop", md.containerID).Output()

	if err != nil {
		md.t.Fatalf("Error stopping docker container. %v", out)
	}

	out, err = exec.Command("docker", "rm", md.containerID).Output()

	if err != nil {
		md.t.Fatalf("Error removing docker container. %v", out)
	}
}

func installDockerImage(t *testing.T) {
	log.Printf("Pulling docker image %s ...", mongoImage)
	if err := dockerPull(mongoImage); err != nil {
		t.Skipf("Error pulling %s: %v", mongoImage, err)
	}
}

func dockerPull(name string) interface{} {
	out, err := exec.Command("docker", "pull", name).CombinedOutput()
	if err != nil {
		err = fmt.Errorf("%v: %s", err, out)
	}
	return err
}

func dockerHasImage(name string) (ok bool, err error) {
	out, err := exec.Command("docker", "images", "--filter=reference=mongo:latest").Output()

	if err != nil {
		return
	}

	return bytes.Contains(out, []byte(name)), nil
}
