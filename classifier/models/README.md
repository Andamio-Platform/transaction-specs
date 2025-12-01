
## admincourse

`/admin/course/create`

```json
{
    "courseId": "string",
    "admin": "string",
    "teachers": ["string"],
    "courseAddress": "string",
    "courseStatePolicyId": "string"
}
```

`/admin/course/teachers/update`

```json
{
    "courseId": "string",
    "teachers": {
        "add": ["string"],
        "remove": ["string"]
    }
}
```

## studentcourse

`/student/course/enroll`

```json
{
    "alias": "string",
    "courseId": "string"
}
```

`/student/course/assignment/submit`

```json
{
    "alias": "string",
    "courseId": "string",
    "assignmentId": "string",
    "content": "string"
}
```

`/student/course/assignment/update`

```json
{
    "alias": "string",
    "courseId": "string",
    "assignmentId": "string",
    "content": "string"
}
```

`/student/course/credential/claim`

```json
{
    "alias": "string",
    "courseId": "string",
}
```

## teachercourse

`/teacher/course/modules/manage`

```json
{
    "alias": "string",
    "courseId": "string",
    "modules": {
        "create": [
            {
                "assignmentId" : "string",
                "module":{
                    "slts": ["string"],
                    "prerequisites": ["string"],
                }
            }
        ],
        "update": [
            {
                "assignmentId" : "string",
                "module":{
                    "prerequisites": ["string"],
                }
            }
        ],
        "delete": [
            {
                "assignmentId" : "string"
            }
        ]
    }
}
```

`/teacher/course/assignments/assess`

```json
{
    "alias": "string",
    "courseId": "string",
    "assignmentId": "string",
    "assessments": [
        {
            "studentAlias": "string",
            "assessment": "string"
        }
    ]
}
```


