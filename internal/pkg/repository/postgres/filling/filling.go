package main

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres/filling/consts"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres/filling/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"os"
)

type CustomSub struct {
	customID   string
	authorName string
	layer      string
}

type Author string
type Layer string
type PostTitle string
type User string

//var (
//	CustomSubToAuthor = map[string]CustomSub{}
//	Subscriptions     = map[string]CustomSub{}
//
//	Posts = make(map[Author]map[Layer][]PostTitle)
//	Users = make(map[Author]map[Layer][]User)
//)

type Filling struct {
	CustomSubToAuthor map[string]CustomSub
	Subscriptions     map[string]CustomSub

	Posts map[Author]map[Layer][]PostTitle
	Users map[Author]map[Layer][]User

	Authors []*models.FillingAuthor
}

func main() {
	// Укажите, сколько пользователей нужно создать
	n := consts.COUNT_USER

	// Параметры подключения к PostgreSQL
	dbHost := "postgres"
	dbHost = "127.0.0.1"
	dbPort := 5432
	dbUser := "admin"
	dbPassword := "adminpass"
	dbName := "testdb"

	path := "models/authors.json"
	path = "internal/pkg/repository/postgres/filling/models/authors.json"

	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)

	// Формируем строку подключения
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	//dbURL := "postgres://your_user:your_password@localhost:5432/your_dbname?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Ошибка создания пула подключений: %v", err)
	}
	defer pool.Close()

	filling := &Filling{
		CustomSubToAuthor: map[string]CustomSub{},
		Subscriptions:     map[string]CustomSub{},

		Posts: make(map[Author]map[Layer][]PostTitle),
		Users: make(map[Author]map[Layer][]User),
	}
	filling.Authors = models.GetAuthors(path)
	if filling.Authors == nil {
		panic("filling is empty")
	}

	log.Println(filling)

	if err := filling.createUsers(context.Background(), pool, n); err != nil {
		log.Fatalf("Ошибка при создании пользователей: %v", err)
	}

	fmt.Printf("Создано %d пользователей\n", n)

	if err := filling.createAuthors(context.Background(), pool, n); err != nil {
		log.Fatalf("Ошибка при создании авторов: %v", err)
	}

	fmt.Printf("Создано %d авторов \n", n)

	if err := filling.createPage(context.Background(), pool, n); err != nil {
		log.Fatalf("Ошибка при создании страниц: %v", err)
	}

	fmt.Printf("Создано %d страниц \n", n)

	if err := filling.createCustomSubscriptions(context.Background(), pool, n); err != nil {
		log.Fatalf("Ошибка при создании кастомных подписок: %v", err)
	}

	fmt.Printf("Создано %d кастомных подписок \n", n)

	if err := filling.createSubscriptions(context.Background(), pool, n); err != nil {
		log.Fatalf("Ошибка при создании подписок: %v", err)
	}

	fmt.Printf("Создано %d подписок \n", n*n)

	if err := filling.createPosts(context.Background(), pool, n); err != nil {
		log.Fatalf("Ошибка при создании постов: %v", err)
	}

	fmt.Printf("Создано %d постов \n", n*n)
	index, err := filling.createPostLikes(context.Background(), pool, n)
	if err != nil {
		log.Fatalf("Ошибка при создании лайков на посты: %v", err)
	}

	fmt.Printf("Создано %d лайков \n", index)
}

func (f *Filling) createUsers(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < n; i++ {
		userID := uuid.New()
		username := fmt.Sprintf(consts.USERNAME, i+1)
		email := fmt.Sprintf(consts.EMAIL_DOMAIN_NAME, username)

		// Генерация хеша пароля
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(consts.PASSWORD), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("ошибка хеширования пароля: %v", err)
		}

		// Запрос на добавление пользователя
		batch.Queue(`
            INSERT INTO People (user_id, username, email, role_id, hash_password) 
            VALUES ($1, $2, $3, (SELECT role_id FROM Role WHERE role_default_name = 'Reader'), $4)`,
			userID, username, email, passwordHash)

	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func (f *Filling) createAuthors(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	if len(f.Authors) < n {
		n = len(f.Authors)
	}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < n; i++ {
		userID := uuid.New()
		//username := fmt.Sprintf(consts.AUTHOR_NAME, i+1)
		//email := fmt.Sprintf(consts.EMAIL_DOMAIN_NAME, username)

		authorName := f.Authors[i].AuthorName
		email := fmt.Sprintf(consts.EMAIL_DOMAIN_NAME, authorName)
		password := f.Authors[i].Password
		if password == "" {
			password = consts.PASSWORD
		}

		// Генерация хеша пароля
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("ошибка хеширования пароля: %v", err)
		}

		// Запрос на добавление пользователя
		batch.Queue(`
            INSERT INTO People (user_id, username, email, role_id, hash_password) 
            VALUES ($1, $2, $3, (SELECT role_id FROM Role WHERE role_default_name = 'Author'), $4)`,
			userID, authorName, email, passwordHash)

	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func (f *Filling) createPage(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	if len(f.Authors) < n {
		n = len(f.Authors)
	}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < n; i++ {
		//username := fmt.Sprintf(consts.AUTHOR_NAME, i+1)
		authorName := f.Authors[i].AuthorName
		about := fmt.Sprintf(consts.PAGE_INFO, authorName)

		// Запрос на добавление пользователя
		batch.Queue(`
INSERT INTO Page (page_id, user_id, info) VALUES
(gen_random_uuid(), (SELECT user_id FROM People WHERE username = $1), $2)
  `,
			authorName, about)
	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func (f *Filling) createCustomSubscriptions(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	if len(f.Authors) < n {
		n = len(f.Authors)
	}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < n; i++ {
		//authorName := fmt.Sprintf(consts.AUTHOR_NAME, i+1)
		//username := fmt.Sprintf(consts.AUTHOR_NAME, i+1)
		//customSub := fmt.Sprintf(consts.CUSTOM_NAME, i+1, authorName)

		authorName := f.Authors[i].AuthorName
		for _, customSub := range f.Authors[i].CustomSubscription {
			customSubID := uuid.New().String()
			customSubName := customSub.Title
			customSubDescription := customSub.Description
			if customSubDescription == "" {
				customSubDescription = consts.CUSTOM_SUB_INFO
			}
			cost := customSub.Cost
			layer := customSub.Layer

			if cost == 0 {
				cost = consts.CUSTOM_COST*layer + 10
			}

			f.CustomSubToAuthor[customSubID] = CustomSub{layer: string(layer), authorName: authorName, customID: customSubID}

			// Запрос на добавление пользователя
			batch.Queue(`
    INSERT INTO Custom_Subscription (custom_subscription_id, author_id, custom_name, cost, info, subscription_layer_id) VALUES
    ($1, (SELECT user_id FROM People WHERE username = $2), $3, $4, $5, (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = $6) )
  `,
				customSubID, authorName, customSubName, cost, customSubDescription, layer)
		}
	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func (f *Filling) createSubscriptions(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < len(f.Authors); i++ {
		//authorName := fmt.Sprintf(consts.AUTHOR_NAME, i+1)
		author := f.Authors[i%len(f.Authors)]
		log.Println("Author::: ", author)
		authorName := f.Authors[i%len(f.Authors)].AuthorName
		//customSub := fmt.Sprintf(consts.CUSTOM_NAME, i+1, authorName)
		customSubName := f.Authors[i%len(f.Authors)].CustomSubscription[i%len(f.Authors[i%len(f.Authors)].CustomSubscription)].Title
		//customSubID :=  f.Authors[i%len(f.Authors)].CustomSubscription[i%len(f.Authors[i%len(f.Authors)].CustomSubscription)].
		layer := author.CustomSubscription[i%len(author.CustomSubscription)].Layer
		for j := 0; j < n; j++ {
			username := fmt.Sprintf(consts.USERNAME, j+1)

			f.Subscriptions[username] = f.CustomSubToAuthor[customSubName]

			if f.Users[Author(authorName)] == nil {
				f.Users[Author(authorName)] = make(map[Layer][]User)
			}
			f.Users[Author(authorName)][Layer(string(layer))] = append(f.Users[Author(authorName)][Layer(string(layer))], User(username))

			// Запрос на добавление пользователя
			batch.Queue(`
  INSERT INTO Subscription (subscription_id, user_id, custom_subscription_id, started_date, finished_date) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People
                                        WHERE username = $1), (SELECT custom_subscription_id
                                                               FROM Custom_Subscription
                                                               join People ON Custom_Subscription.author_id = people.user_id
    WHERE custom_name = $2 and username = $3 limit 1), NOW(), NOW() + INTERVAL '30 days')
`,
				username, customSubName, authorName)
		}
	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func (f *Filling) createPosts(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < len(f.Authors); i++ {
		//authorName := fmt.Sprintf(consts.AUTHOR_NAME, i+1)
		//about := fmt.Sprintf(consts.ABOUT, authorName)
		//layer := i % 4
		authorName := f.Authors[i].AuthorName

		//for j := 0; j < n; j++ {
		for _, post := range f.Authors[i].Posts {
			title := post.Title
			about := post.Info
			layer := post.Layer
			//title := fmt.Sprintf(consts.TITLE, i, j, authorName)
			if f.Posts[Author(authorName)] == nil {
				f.Posts[Author(authorName)] = make(map[Layer][]PostTitle)
			}
			f.Posts[Author(authorName)][Layer(string(layer))] = append(f.Posts[Author(authorName)][Layer(string(layer))], PostTitle(title)) // Запрос на добавление пользователя
			batch.Queue(`
INSERT INTO Post (post_id, user_id, title, about, subscription_layer_id) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = $1), $2, $3 , (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = $4))
`,
				authorName, title, about, layer)
		}
	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func (f *Filling) createPostLikes(ctx context.Context, pool *pgxpool.Pool, n int) (index int, err error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return 0, fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}
	index = 0

	for authorName, layerPosts := range f.Posts {
		for layer, posts := range layerPosts {
			for _, title := range posts {

				userNames := f.Users[Author(authorName)][Layer(string(layer))]
				for _, username := range userNames {
					index = index + 1

					top := index % 6
					if top == 0 {
						continue
					}
					if rand.Int()%133%7 == 0 {
						continue
					}

					// Запрос на добавление пользователя
					batch.Queue(`
INSERT INTO Like_Post (like_post_id, post_id, user_id, posted_date) VALUES
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = $1 and post.user_id = (SELECT user_id FROM People WHERE username = $3) limit 1), (SELECT user_id FROM People WHERE username = $2), NOW())
   `,
						string(title), string(username), authorName)
				}
			}
		}
	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return 0, fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return index, nil
}
