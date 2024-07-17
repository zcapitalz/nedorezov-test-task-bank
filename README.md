# Запуск

**Endpoints**

Переделал deposit и withdraw эндпоинты чтобы их название описывало ресурс:

POST /accounts/{accountID}/withdraw -> POST /accounts/{accountID}/transactions

POST /accounts/{accountID}/deposit -> POST /accounts/{accountID}/transactions

**Горутины, каналы**

Gin фреймворк самостоятельно создает горутины для обработки каждого запроса. При текущих операциях нет небоходимости коммуникации горутин, поэтому каналы не использовались.