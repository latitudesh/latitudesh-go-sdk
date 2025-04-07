// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TeamMembersRole struct {
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o *TeamMembersRole) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TeamMembersRole) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TeamMembersRole) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TeamMembersRole) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type TeamMembersData struct {
	FirstName  *string          `json:"first_name,omitempty"`
	LastName   *string          `json:"last_name,omitempty"`
	Email      *string          `json:"email,omitempty"`
	MfaEnabled *bool            `json:"mfa_enabled,omitempty"`
	Role       *TeamMembersRole `json:"role,omitempty"`
}

func (o *TeamMembersData) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *TeamMembersData) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *TeamMembersData) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *TeamMembersData) GetMfaEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.MfaEnabled
}

func (o *TeamMembersData) GetRole() *TeamMembersRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type TeamMembers struct {
	Data []TeamMembersData `json:"data,omitempty"`
}

func (o *TeamMembers) GetData() []TeamMembersData {
	if o == nil {
		return nil
	}
	return o.Data
}
