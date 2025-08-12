/v1/resources/extSystemId={extSystemId}:
  get:
    summary: Получение списка ресурсов по идентификатору внешней системы
    operationId: GetResourcesByExtSystemId
    parameters:
      - name: extSystemId
        in: path
        description: Уникальный идентификатор внешней системы
        required: true
        schema:
          type: integer
    tags:
      - Resource
    responses:
      '200':
        description: Список ресурсов внешней системы по айди
        content:
          application/json:
            schema:
              type: object
              description: Список ресурсов внешней системы по айди
              required:
                - resources
              properties:
                resources:
                  type: array
                  items:
                    $ref: '#/components/schemas/ExtSystemResource'
            examples:
              success:
                summary: Успешный пример
                value:
                  resources:
                    - id: 501
                      code: GET_CUSTOMER_PROFILE
                      name: Получение профиля клиента
                      method: GET
                      path: /api/customer/profile
                      priceType: FIXED
                      amount: 100.00
                      extSystemId: 301
                      cacheStoragePeriod: 7
                      active: true
                      createdAt: "2025-08-01T11:47:58+05:00"
                      updatedAt: "2025-08-01T11:47:58+05:00"
                    - id: 502
                      code: CREATE_LOAN_REQUEST
                      name: Создание заявки на кредит
                      method: POST
                      path: /api/loan/create
                      priceType: PER_REQUEST
                      amount: 250.00
                      extSystemId: 301
                      cacheStoragePeriod: 0
                      active: false
                      createdAt: "2025-07-28T09:30:00+05:00"
                      updatedAt: "2025-07-30T12:15:00+05:00"
      '400':
        $ref: '#/components/responses/BadRequest'
      '401':
        $ref: '#/components/responses/Unauthorized'
      '500':
        $ref: '#/components/responses/InternalError'
