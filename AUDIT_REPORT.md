# Andamio Transactions API - OpenAPI Specification Audit Report

**Generated:** November 24, 2025  
**API Version:** 2.0  
**OpenAPI Version:** 3.0.3

---

## Executive Summary

The OpenAPI specification is **structurally valid** but has several consistency issues and design improvements needed. Most issues are medium severity and should be addressed for API clarity and maintainability.

---

## 🔴 Critical Issues

### None identified

---

## 🟡 High Priority Issues

### 1. Inconsistent HTTP Methods for Mutations
**Endpoints Affected:** Multiple  
**Issue:** Several endpoints that perform mutations use GET requests instead of POST/PUT.

- ✗ `/admin/course/create` - Uses GET (should be POST)
- ✗ `/student/course/enroll` - Uses GET (should be POST)
- ✗ `/student/course/assignment/submit` - Uses GET (should be POST)
- ✗ `/teacher/course/assignment/assess` - Uses GET (should be POST)
- ✗ `/student/course/assignment/update` - Uses GET (should be POST)
- ✗ `/student/course/credential/claim` - Uses GET (should be POST)
- ✗ `/admin/course/teachers/update` - Uses GET (should be POST)

**Best Practice:** Only use GET for safe, idempotent operations that don't modify state.

**Recommendation:** Convert to POST (or PUT for updates) and move required parameters to request body.

---

### 2. Inconsistent Field Naming
**Issue:** Field names are inconsistent across similar endpoints.

- `/teacher/course/modules/create` uses `prerequisites` in schema but `prerequisiteAssignments_V2` in examples
- `/teacher/course/modules/manage` uses `prerequisites` consistently
- Unclear naming convention: `*_V2` suffix inconsistently applied

**Recommendation:** Standardize to either:
- Option A: `prerequisites` (simple)
- Option B: `prerequisiteAssignments_V2` (explicit with version)

Choose one and apply consistently across all endpoints.

---

### 3. Incomplete Request Body Schemas
**Endpoints Affected:** `/teacher/course/assignments/assess`  
**Issue:** Root array schema without metadata support.

```yaml
schema:
  type: array  # ← Cannot extend with metadata
  items:
    type: object
```

**Recommendation:** Wrap in object for future extensibility:
```yaml
schema:
  type: object
  properties:
    assessments:
      type: array
      items: {...}
```

---

## 🟠 Medium Priority Issues

### 4. Typo in Endpoint Path
**Path:** `/student/course/creadential/claim`  
**Issue:** `creadential` should be `credential`

**Recommendation:** Rename to `/student/course/credential/claim`

---

### 5. Vague Response Schemas
**All Endpoints**  
**Issue:** All endpoints return only `'200': description: OK` with no response body schema.

**Current:**
```yaml
responses:
  '200':
    description: OK
```

**Recommendation:** Define response schemas for all endpoints. Example:
```yaml
responses:
  '200':
    description: Success
    content:
      application/json:
        schema:
          type: object
          properties:
            success:
              type: boolean
            message:
              type: string
            data:
              type: object
  '400':
    description: Bad Request
  '401':
    description: Unauthorized
  '500':
    description: Internal Server Error
```

---

### 6. Missing Error Response Documentation
**All Endpoints**  
**Issue:** No error responses (4xx, 5xx) are documented.

**Recommendation:** Add error response schemas for:
- 400 Bad Request (invalid parameters)
- 401 Unauthorized (missing/invalid alias)
- 404 Not Found (resource not found)
- 500 Internal Server Error

---

### 7. Inconsistent Description Quality
**Issue:** Some descriptions are vague.

- `/teacher/course/modules/manage` description says "Create modules" but it manages (create/update/delete)
- Several endpoints lack usage context or constraints

**Recommendation:** Update descriptions to be more precise:
- ✓ `/teacher/course/modules/manage`: "Create, update, or delete course modules"

---

### 8. Missing Required Fields Definition
**Endpoints:** All with request bodies  
**Issue:** Object schemas don't define `required` fields.

**Example Problem:**
```yaml
schema:
  type: object
  properties:
    modules: {...}
  # ← Missing: required: [modules]
```

**Recommendation:** Add `required` arrays to all object schemas to clarify mandatory fields.

---

### 9. Inconsistent Wallet Address Handling
**Issue:** `addresses` and `changeAddress` parameters are duplicated across many endpoints with identical descriptions.

**Current Pattern:**
```yaml
- name: addresses
  in: query
  description: 'Used addresses for multi address cardano wallets'
  required: false
- name: changeAddress
  in: query
  description: 'Change address for multi address cardano wallets'
  required: false
```

**Recommendation:** 
1. Extract to reusable components (`#/components/parameters`) for DRY principle
2. Add more details about format (hash format, length, etc.)

---

### 10. Inconsistent Array Item Types
**Issue:** Array items sometimes lack type definition.

**Example from `/teacher/course/modules/manage`:**
```yaml
delete:
  type: array
  items:
    type: string
    description: 'Hash of the Student Learning Target to delete'
```

Better to define minimum length/format:
```yaml
delete:
  type: array
  items:
    type: string
    minLength: 64  # If these are hex hashes
    description: 'Hash of the Student Learning Target to delete'
```

---

## 🟢 Low Priority Issues / Observations

### 11. API Versioning in URL vs Version Field
**Observation:** URL contains `/v2` but info.version is just `2.0`

**Recommendation:** Consider clarifying versioning strategy or update info.version to `2.0.0` (semantic versioning).

---

### 12. Limited Tag Usage
**Observation:** Only `course`, `user`, `migration`, and `project` tags are used.

**Recommendation:** Add more granular tags:
- `modules`
- `assignments`
- `assessments`
- `enrollment`
- `admin`

This improves API documentation organization.

---

### 13. No Info Contact or License
**Observation:** Missing contact information and license in `info` section.

**Recommendation:** Add:
```yaml
info:
  title: Andamio Transactions API
  version: '2.0'
  description: API for managing Andamio course transactions
  contact:
    name: API Support
    url: https://andamio.io
  license:
    name: Apache 2.0
```

---

### 14. Incomplete Project Endpoints
**Observation:** 11 project-related endpoints are commented out.

**Recommendation:** 
- If not implemented yet, remove from spec or clearly mark as "coming soon"
- If implemented, uncomment and properly document
- If deprecated, document migration path

---

### 15. No Schema Reusability
**Issue:** Similar object structures are duplicated instead of using `$ref`.

**Example:**
```yaml
create:
  items:
    type: object
    properties:
      slts: {...}
      prerequisites: {...}
update:
  items:
    type: object
    properties:
      sltHash: {...}
      prerequisites: {...}
```

**Recommendation:** Extract common structures to `#/components/schemas` and reuse with `$ref`.

---

## Summary Statistics

| Category | Count |
|----------|-------|
| **Active Endpoints** | 10 |
| **Commented Endpoints** | 11 |
| **GET Mutation Endpoints** | 7 ❌ |
| **POST Endpoints** | 2 ✓ |
| **Endpoints with Response Schema** | 0 ❌ |
| **Endpoints with Error Responses** | 0 ❌ |

---

## Action Items (Priority Order)

### Phase 1: Critical (Fix First)
- [ ] Fix HTTP methods for mutations (GET → POST)
- [ ] Fix typo: `creadential` → `credential`
- [ ] Standardize field naming (`prerequisites` vs `prerequisiteAssignments_V2`)

### Phase 2: Important (Fix Soon)
- [ ] Add response schemas to all endpoints
- [ ] Add error response documentation (400, 401, 404, 500)
- [ ] Add `required` fields to object schemas
- [ ] Wrap `/teacher/course/assignments/assess` root array in object

### Phase 3: Nice to Have (Refactor)
- [ ] Extract wallet parameters to reusable components
- [ ] Extract common schema objects to `#/components/schemas`
- [ ] Add more granular tags
- [ ] Enhance endpoint descriptions
- [ ] Add contact info and license to spec

### Phase 4: Polish (Optional)
- [ ] Review/document commented project endpoints
- [ ] Add semantic versioning to info.version
- [ ] Add API description and documentation links

---

## Recommendations Summary

1. **Consistency First:** Align HTTP methods, field naming, and description quality
2. **Complete Documentation:** Add full request/response schemas and error handling
3. **Reusability:** Extract common components and use `$ref`
4. **Clarity:** Expand descriptions and add validation constraints (minLength, format, etc.)
5. **Standards Compliance:** Follow OpenAPI 3.0.3 best practices

---

## Files Ready for Update

- `openapi.yml` - Main specification file

## Next Steps

1. Review this report with the team
2. Prioritize action items based on release timeline
3. Implement Phase 1 fixes immediately
4. Plan Phase 2-3 for next sprint/release
