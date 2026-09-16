# LksPlansTotal


## Supported Types

### 

```go
lksPlansTotal := components.CreateLksPlansTotalInteger(int64{/* values here */})
```

### 

```go
lksPlansTotal := components.CreateLksPlansTotalStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch lksPlansTotal.Type {
	case components.LksPlansTotalTypeInteger:
		// lksPlansTotal.Integer is populated
	case components.LksPlansTotalTypeStr:
		// lksPlansTotal.Str is populated
}
```
