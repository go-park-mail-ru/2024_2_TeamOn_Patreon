### Описание таблиц

Все id - uuid, генерируются СУБД.

Базовая информация о пользователях (читателях и авторах)

1. User - базовые сведения о пользователе
   - username - имя пользователя 
   - email - почта
   - role_id - FK роли
   - avatar_id - FK фото профиля
   - hash_password - хэш пароля
2. Role - роль пользователя
    - default_role_name
        "READER" - читатель, человек, который донатит, дается всем по дефолту
        "AUTHOR" - автор, человек, который выкладывает контент, включает возможности reader
3. Avatar - фото профиля
    - avatar_url -  ссылка на фотографию пользователя
4. Page - страница автора
    - user_id - FK пользователя
    - info - раздел "О себе"
    - background_picture  - ссылка на изображение на фон страницы автора

Контент (посты, которые будут размещаться на странице автора и выдаваться в ленте пользователя)

5. Post - пост
    - user_id - FK пользователя (автора), который выложил пост
    - title - заголовок поста
    - text - текст поста
    - content_id - FK контента
    - layer_id - минимальный уровень подписки, на котором можно смотреть этот контент
    - created_at - дата создания поста
    - updated_at - дата последнего изменения поста
6. Comment - комментарии
    - post_id - FK поста
    - user_id - FK пользователя, который оставил комментарий
    - text - текст комментария
    - created_at - дата оставления комментария
    - updated_at - дата последнего изменения комментария
7. LikePost - информация о лайке на пост
    - post_id - FK на пост, на котором оставила лайк
    - user_id - FK на пользователя, который оставил лайк на пост
    - posted_at - дата, когда поставили лайк на пост
8. LikeComment - информация о лайке на комментарий
    - comment_id - FK на коммент
    - user_id - FK юзера
    - posted_at - дата, когда поставили лайк на комментарий
9. Content - контент поста
    - content_type_id - FK на тип контента
    - content_url - url на медиа, которое приложено к посту 
10. ContentType - тип контента
    - default_content_type_name - имя типа контента (MP3, JPG, PNG, MP4, PDF, ZIP)

Системы поддержки авторов.
Подписки - читатель подписывается на автора по одному из планов подписок
(их максимум три у автора), и может видеть все посты, доступные для этого уровня.

11. CustomSubscription - кастомные подписки (авторов)
    - author_id - FK на автора
    - custom_name - имя собственной подписки, созданное автором
    - info - описание подписки
    - cost - стоимость подписки (в рублях)
    - subscription_layer_id - FK на уровень подписок
    - created_at - время создания подписки
12. Subscription - подписки пользователей на кастомные подписки авторов
    - user_id - FK на пользователя, который подписывается
    - custom_subscription_id - FK на кастомною подписку автора
    - started_at - дата, когда человек подписался
    - finished_at - дата, по которою включительно, была оплачена подписка
13. SubscriptionLayer - уровни подписки
    - layer - приоритет уровня (1, 2, 3)
    - default_layer_name - имя уровня подписки
    
Единоразовая поддержка - возможность поддержать автора единоразово без получения привилегий за донат.

14. Tip - донат
    - user_id - FK на пользователя, который оставляет донат
    - author_id - FK на автора, которому отправили донат
    - cost - сумма доната (в рублях)
    - message - сообщение от пользователя автору
    - payed_at - время совершения платежа

Уведомления и события. 

15. Notification - сущность события
    - user_id - FK на пользователя, которому приходит уведомление
    - event_id - FK на событие, которое произошло
    - sent_at - время, когда уведомление было отправлено
16. Event
    - event_type_id - FK на тип события
    - text - словесное описание события
    - happened_at - время, когда событие произошло
17. EventType
    - default_event_type_name - внутреннее имя события. (Например, LikePost, LikeComment)

Сессия

18. Session
    - user_id - FK на пользователя, чья сессия
    - created_at - время создания сессии
    - finished_at - время окончания сессии (определяется заранее, может быть уменьшено)

### Отношения и нормализация

**Нормализация**:
1. 1НФ - все атрибуты во всех таблицах являются атомарными
2. 2НФ - Отношение находится во 2НФ, если оно находится в 1НФ и каждый не ключевой атрибут неприводимо зависит от Первичного Ключа(ПК)
Неприводимо - т.е. ПК - минимален. Во всех отношениях ПК - минимален
3. 3НФ - Отношение находится в 3НФ, когда находится во 2НФ и каждый не ключевой атрибут нетранзитивно зависит от первичного ключа. 
Т.н. нет неключевого атрибут, который зависит от другого неключевого атрибута
4. НФБК - каждая нетривиальная и неприводимая слева функциональная зависимость обладает потенциальным ключом в качестве детерминанта.
Т.е. нет неключевого атрибута, не являющегося потенциальным ключом, от которого зависит неключевой атрибут.


Пользователи

1. User:

{user_id} -> username, email, role_id, avatar_id, hash_password

Потенциальные ключи
{username} -> email, role_id, avatar_id, hash_password
{email} -> username, role_id, avatar_id, hash_password

2. Role

{role_id} -> default_role_name

3. Avatar

{avatar_id} -> avatar_url

5. Page

{page_id} -> user_id, info, background_picture

Потенциальные ключи
{user_id} -> info, background_picture

Контент

5. Post

{post_id} -> user_id, title, text, content_id, layer_id, created_at, updated_at

6. Comment

{comment_id} -> - post_id, user_id, text, created_at, updated_at

7. LikePost 

{like_post_id} -> post_id, user_id, posted_at

8. LikeComment

{like_comment_id} -> comment_id, user_id, posted_at 

9. Content

{content_id} -> content_type_id, content_url 

10. ContentType 

{content_type_id} -> default_content_type_name

Системы поддержки авторов.

11. CustomSubscription 

{custom_subscription_id} -> author_id, custom_name, info, cost,
subscription_layer_id, created_at

Потенциальные ключи
{author_id, custom_name} -> info, cost, subscription_layer_id, created_at
{author_id, subscription_layer_id} -> info, cost, custom_name, created_at

12. Subscription

{subscription_id} -> user_id, custom_subscription_id, started_at, finished_at

Потенциальные ключи
{user_id, custom_subscription_id} -> started_ad, finished_ad

13. SubscriptionLayer 

{subscription_layer_id} -> layer, default_layer_name
    
14. Tip

{tip_id} -> user_id, author_id, cost, message, payed_at 

Уведомления и события.

15. Notification 

{notification_id} -> user_id, event_id, sent_at 

16. Event

{event_id} -> event_type_id, text, happened_at

17. EventType

{event_type_id} -> default_event_type_name

Сессия

18. Session

{session_id} -> user_id, created_at, finished_at
