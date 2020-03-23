export default [
  {
    "version": 0.1,
    "name":"demo-project",
    "data_model": {
      "key_list": [
        {
          "key": "id",
          title:"id",
          "var_type": "string",
          "is_show_in_list": true,
          "is_unique_key": true
        },
        {
          "key": "display_name",
          title:"组件名",
          "var_type": "string",
          "is_show_in_list": true,
          "is_unique_key": false
        },
        {
          "key": "package_name",
          title:"npm包名",
          "var_type": "string",
          "is_show_in_list": false,
          "is_unique_key": false
        },
        {
          "key": "dev_list_json",
          "var_type": "string",
          "is_show_in_list": false,
          "is_unique_key": false
        },
        {
          "key": "description",
          "var_type": "string",
          "is_show_in_list": false,
          "is_unique_key": false
        },
        {
          "key": "site_url",
          "var_type": "string",
          "is_show_in_list": false,
          "is_unique_key": false
        },
        {
          "key": "remark",
          "var_type": "string",
          "is_show_in_list": false,
          "is_unique_key": false
        }
      ]
    },
    "page_config": { "create": true, "update": true, "detail": true },
    "base_url_path": "/component",
    "base_api_path": "/api/component"
  }
]