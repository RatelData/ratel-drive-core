# APIs documentation

All APIs start with `/api`

## Storage Service APIs
All Storage Service APIs start with `/api/storage`
* **GET** `/api/storage/:path`
  * Description:
    * Get all files/directories in `:path`
  * Response:
      ```json
      {
        "result": "success|failed",
        "error": "error info",
        "data": [
            "filename01",
            "directory01",
            ...
        ]
      }
      ```
* **POST** `/api/storage/upload`
  * Description:
    * Used for uploading files
  * Format: multi-form
  * Response: 
      ```json
      {
        "result": "success|failed",
        "error": "error info",
        "msg": "result messages"
      }
      ```

## Authentication Service APIs