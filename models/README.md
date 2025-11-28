

`/admin/course/create`

```json
{
    "courseId": "string",
    "admin": "string",
    "teachers": ["string"],
    "courseAddress": "string",
}
```

`/student/course/enroll`

```json
{
    "alias": "string",
    "courseId": "string"
}
```

`/teacher/course/modules/create`

```json
{
    "alias": "string",
    "courseId": "string",
    "modules": [
        {
            "slts": ["string"],
            "prerequisites": ["string"],
        }
    ]
}
```

`/teacher/course/modules/manage`

```json
{
    "alias": "string",
    "courseId": "string",
    "modules": {
        "create": [
        {
            "slts": ["string"],
            "prerequisites": ["string"],
        }
    ],
        "update": [
        {
            "sltHash": "string",
            "prerequisites": ["string"],
        }
    ],
        "delete": ["string"]
    }
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