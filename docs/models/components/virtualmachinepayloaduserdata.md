# VirtualMachinePayloadUserData

A user data record reference (encoded id_hash, e.g. 'ud_xxx', or raw integer id) to apply as cloud-init configuration


## Supported Types

### 

```go
virtualMachinePayloadUserData := components.CreateVirtualMachinePayloadUserDataInteger(int64{/* values here */})
```

### 

```go
virtualMachinePayloadUserData := components.CreateVirtualMachinePayloadUserDataStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualMachinePayloadUserData.Type {
	case components.VirtualMachinePayloadUserDataTypeInteger:
		// virtualMachinePayloadUserData.Integer is populated
	case components.VirtualMachinePayloadUserDataTypeStr:
		// virtualMachinePayloadUserData.Str is populated
}
```
