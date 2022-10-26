package tests

import (
	"github.com/younny/slobbo-backend/src/db"
	"github.com/younny/slobbo-backend/src/types"
)

var (
	testClient = &db.Client{}
	testPost   = types.Post{
		Title:    "Hello World",
		SubTitle: "Foo",
		Author:   "Jake",
		Category: 0,
		Body:     "ABC",
	}
)

// func TestMain(m *testing.M) {
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatalf("Couldn't connect to docker: %s", err)
// 	}

// 	r, err := pool.RunWithOptions(&dockertest.RunOptions{
// 		Repository: "postgres",
// 		Tag:        "11",
// 		Env: []string{
// 			"POSTGRES_PASSWORD=secret",
// 			"POSTGRES_USER=user_name",
// 			"POSTGRES_DB=dbname",
// 			"listen_addresses = '*'",
// 		},
// 	}, func(config *docker.HostConfig) {
// 		// set AutoRemove to true so that stopped container goes away by itself
// 		config.AutoRemove = true
// 		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
// 	})

// 	if err != nil {
// 		log.Fatalf("Could not start resource: %s", err)
// 	}

// 	//hostAndPort := r.GetHostPort("5432/tcp")
// 	//databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
// 	databaseUrl2 := fmt.Sprintf("host=localhost port=%s user=postgres dbname=slobbo_api_test password=yunalito sslmode=disable", r.GetPort("5432/tcp"))
// 	log.Println("Connecting to database on url: ", databaseUrl2)

// 	r.Expire(120) // Tell docker to hard kill the container in 120 seconds

// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	pool.MaxWait = 120 * time.Second
// 	if err = pool.Retry(func() error {
// 		err = testClient.Connect(databaseUrl2)
// 		if err != nil {
// 			return err
// 		}
// 		return testClient.Ping()
// 	}); err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}
// 	//Run tests
// 	code := m.Run()

// 	// You can't defer this because os.Exit doesn't care for defer
// 	if err := pool.Purge(r); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)
// }

// func TestClient_Posts(t *testing.T) {
// 	testClient.Client.DropTable(&types.Post{})
// 	testClient.Client.AutoMigrate()

// 	first := testPost
// 	err := testClient.CreatePost(&first)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 0, first.ID)

// 	second := testPost
// 	err = testClient.CreatePost(&second)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, first.ID)

// 	update := first
// 	update.Title = "Hello World 2"
// 	err = testClient.UpdatePost(&update)
// 	assert.NoError(t, err)

// 	get := testClient.GetPostByID(0)
// 	assert.Equal(t, "Hello World 2", get.Title, "")
// 	assert.Equal(t, testPost.SubTitle, get.SubTitle)
// }
