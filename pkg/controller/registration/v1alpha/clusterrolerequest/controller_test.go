package clusterrolerequest

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/EdgeNet-project/edgenet/pkg/access"
	registrationv1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/registration/v1alpha"
	"github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned"
	edgenettestclient "github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned/fake"
	informers "github.com/EdgeNet-project/edgenet/pkg/generated/informers/externalversions"
	"github.com/EdgeNet-project/edgenet/pkg/signals"
	"github.com/EdgeNet-project/edgenet/pkg/util"
	"github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
	"k8s.io/klog"
)

type TestGroup struct {
	roleRequestObj registrationv1alpha.ClusterRoleRequest
}

var kubeclientset kubernetes.Interface = testclient.NewSimpleClientset()
var edgenetclientset versioned.Interface = edgenettestclient.NewSimpleClientset()

func TestMain(m *testing.M) {
	klog.SetOutput(ioutil.Discard)
	log.SetOutput(ioutil.Discard)
	logrus.SetOutput(ioutil.Discard)

	flag.String("dir", "../../../../..", "Override the directory.")
	flag.String("smtp-path", "../../../../../configs/smtp_test.yaml", "Set SMTP path.")
	flag.Parse()

	stopCh := signals.SetupSignalHandler()

	edgenetInformerFactory := informers.NewSharedInformerFactory(edgenetclientset, time.Second*30)

	controller := NewController(kubeclientset,
		edgenetclientset,
		edgenetInformerFactory.Registration().V1alpha().ClusterRoleRequests())

	edgenetInformerFactory.Start(stopCh)

	go func() {
		if err := controller.Run(2, stopCh); err != nil {
			klog.Fatalf("Error running controller: %s", err.Error())
		}
	}()

	access.Clientset = kubeclientset
	access.CreateClusterRoles()

	time.Sleep(500 * time.Millisecond)

	os.Exit(m.Run())
	<-stopCh
}

// Init syncs the test group
func (g *TestGroup) Init() {
	roleRequestObj := registrationv1alpha.ClusterRoleRequest{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleRequest",
			APIVersion: "registration.edgenet.io/v1alpha",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "johnsmith",
		},
		Spec: registrationv1alpha.ClusterRoleRequestSpec{
			FirstName: "John",
			LastName:  "Smith",
			Email:     "john.smith@edge-net.org",
			RoleName:  "edgenet:tenant-admin",
		},
	}
	g.roleRequestObj = roleRequestObj
}

func TestStartController(t *testing.T) {
	g := TestGroup{}
	g.Init()
	roleRequestTest := g.roleRequestObj.DeepCopy()
	roleRequestTest.SetName("role-request-controller-test")

	// Create a role request object
	edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Create(context.TODO(), roleRequestTest, metav1.CreateOptions{})
	// Wait for the status update of created object
	time.Sleep(time.Millisecond * 500)
	// Get the object and check the status
	roleRequest, err := edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Get(context.TODO(), roleRequestTest.GetName(), metav1.GetOptions{})
	util.OK(t, err)
	expected := metav1.Time{
		Time: time.Now().Add(72 * time.Hour),
	}
	util.Equals(t, expected.Day(), roleRequest.Status.Expiry.Day())
	util.Equals(t, expected.Month(), roleRequest.Status.Expiry.Month())
	util.Equals(t, expected.Year(), roleRequest.Status.Expiry.Year())

	util.Equals(t, pending, roleRequest.Status.State)
	util.Equals(t, messageRoleNotApproved, roleRequest.Status.Message)

	roleRequest.Spec.Approved = true
	edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Update(context.TODO(), roleRequest, metav1.UpdateOptions{})
	time.Sleep(time.Millisecond * 500)
	roleRequest, err = edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Get(context.TODO(), roleRequestTest.GetName(), metav1.GetOptions{})

	util.OK(t, err)
	util.Equals(t, approved, roleRequest.Status.State)
	util.Equals(t, messageRoleApproved, roleRequest.Status.Message)
}

func TestTimeout(t *testing.T) {
	g := TestGroup{}
	g.Init()
	roleRequestTest := g.roleRequestObj.DeepCopy()
	roleRequestTest.SetName("role-request-timeout-test")
	edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Create(context.TODO(), roleRequestTest, metav1.CreateOptions{})
	time.Sleep(time.Millisecond * 500)

	t.Run("set expiry date", func(t *testing.T) {
		roleRequest, _ := edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Get(context.TODO(), roleRequestTest.GetName(), metav1.GetOptions{})
		expected := metav1.Time{
			Time: time.Now().Add(72 * time.Hour),
		}
		util.Equals(t, expected.Day(), roleRequest.Status.Expiry.Day())
		util.Equals(t, expected.Month(), roleRequest.Status.Expiry.Month())
		util.Equals(t, expected.Year(), roleRequest.Status.Expiry.Year())
	})
	t.Run("timeout", func(t *testing.T) {
		roleRequest, _ := edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Get(context.TODO(), roleRequestTest.GetName(), metav1.GetOptions{})
		roleRequest.Status.Expiry = &metav1.Time{
			Time: time.Now().Add(10 * time.Millisecond),
		}
		edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().UpdateStatus(context.TODO(), roleRequest, metav1.UpdateOptions{})
		time.Sleep(100 * time.Millisecond)
		_, err := edgenetclientset.RegistrationV1alpha().ClusterRoleRequests().Get(context.TODO(), roleRequest.GetName(), metav1.GetOptions{})
		util.Equals(t, true, errors.IsNotFound(err))
	})
}
