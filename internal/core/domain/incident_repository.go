package domain

type IncidentRepository interface {
	Save(i *Incident) (*Incident, error)
	GetById(id string) (*Incident, error)
	GetActiveByMonitorId(monitorId string) (*Incident, error)
	Update(i *Incident) (*Incident, error)
	Delete(id string) (int64, error)
}
