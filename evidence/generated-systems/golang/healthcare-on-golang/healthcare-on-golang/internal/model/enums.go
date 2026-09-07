package model


//==============================================================
// FacilityType Declaration
//==============================================================
type FacilityType int
const (
    FacilityTypeHospital FacilityType = iota
	FacilityTypeClinic
	FacilityTypeAmbulatorySurgeryCenter
	FacilityTypeUrgentCare
	FacilityTypeLaboratory
	FacilityTypeImagingCenter
	FacilityTypePharmacy
)


//==============================================================
// DepartmentType Declaration
//==============================================================
type DepartmentType int
const (
    DepartmentTypeEmergency DepartmentType = iota
	DepartmentTypeCardiology
	DepartmentTypeOncology
	DepartmentTypeOrthopedics
	DepartmentTypePediatrics
	DepartmentTypeRadiology
	DepartmentTypePathology
	DepartmentTypePharmacy
	DepartmentTypeIntensiveCare
)


//==============================================================
// CareSettingType Declaration
//==============================================================
type CareSettingType int
const (
    CareSettingTypeInpatient CareSettingType = iota
	CareSettingTypeOutpatient
	CareSettingTypeEmergency
	CareSettingTypeHomeHealth
	CareSettingTypeTelehealth
)


//==============================================================
// ClinicianType Declaration
//==============================================================
type ClinicianType int
const (
    ClinicianTypePhysician ClinicianType = iota
	ClinicianTypeNursePractitioner
	ClinicianTypePhysicianAssistant
	ClinicianTypeRegisteredNurse
	ClinicianTypePharmacist
	ClinicianTypeTherapist
	ClinicianTypeTechnician
)


//==============================================================
// ClinicianSpecialty Declaration
//==============================================================
type ClinicianSpecialty int
const (
    ClinicianSpecialtyInternalMedicine ClinicianSpecialty = iota
	ClinicianSpecialtyFamilyMedicine
	ClinicianSpecialtyCardiology
	ClinicianSpecialtyOncology
	ClinicianSpecialtyOrthopedics
	ClinicianSpecialtyPediatrics
	ClinicianSpecialtyRadiology
	ClinicianSpecialtyPathology
	ClinicianSpecialtyAnesthesiology
	ClinicianSpecialtySurgery
	ClinicianSpecialtyPsychiatry
)


//==============================================================
// AdministrativeSex Declaration
//==============================================================
type AdministrativeSex int
const (
    AdministrativeSexMale AdministrativeSex = iota
	AdministrativeSexFemale
	AdministrativeSexUnknown
)


//==============================================================
// BloodType Declaration
//==============================================================
type BloodType int
const (
    BloodTypeAPositive BloodType = iota
	BloodTypeANegative
	BloodTypeBPositive
	BloodTypeBNegative
	BloodTypeABPositive
	BloodTypeABNegative
	BloodTypeOPositive
	BloodTypeONegative
)


//==============================================================
// AppointmentStatus Declaration
//==============================================================
type AppointmentStatus int
const (
    AppointmentStatusProposed AppointmentStatus = iota
	AppointmentStatusBooked
	AppointmentStatusArrived
	AppointmentStatusFulfilled
	AppointmentStatusCancelled
	AppointmentStatusNoShow
	AppointmentStatusEnteredInError
)


//==============================================================
// Priority Declaration
//==============================================================
type Priority int
const (
    PriorityRoutine Priority = iota
	PriorityUrgent
	PriorityStat
)


//==============================================================
// EncounterStatus Declaration
//==============================================================
type EncounterStatus int
const (
    EncounterStatusPlanned EncounterStatus = iota
	EncounterStatusInProgress
	EncounterStatusOnHold
	EncounterStatusDischarged
	EncounterStatusCancelled
)


//==============================================================
// EncounterType Declaration
//==============================================================
type EncounterType int
const (
    EncounterTypeInpatient EncounterType = iota
	EncounterTypeOutpatient
	EncounterTypeEmergency
	EncounterTypeObservation
	EncounterTypeTelemedicine
)


//==============================================================
// AdmissionType Declaration
//==============================================================
type AdmissionType int
const (
    AdmissionTypeElective AdmissionType = iota
	AdmissionTypeEmergency
	AdmissionTypeUrgent
	AdmissionTypeNewborn
	AdmissionTypeTrauma
)


//==============================================================
// DischargeDisposition Declaration
//==============================================================
type DischargeDisposition int
const (
    DischargeDispositionHome DischargeDisposition = iota
	DischargeDispositionHomeWithHomeCare
	DischargeDispositionSkilledNursingFacility
	DischargeDispositionAcuteCareFacility
	DischargeDispositionExpired
	DischargeDispositionAgainstMedicalAdvice
)


//==============================================================
// OrderStatus Declaration
//==============================================================
type OrderStatus int
const (
    OrderStatusDraft OrderStatus = iota
	OrderStatusActive
	OrderStatusOnHold
	OrderStatusCompleted
	OrderStatusCancelled
)


//==============================================================
// ClinicalOrderType Declaration
//==============================================================
type ClinicalOrderType int
const (
    ClinicalOrderTypeMedication ClinicalOrderType = iota
	ClinicalOrderTypeLaboratory
	ClinicalOrderTypeImaging
	ClinicalOrderTypeProcedure
	ClinicalOrderTypeConsultation
)


//==============================================================
// RouteOfAdministration Declaration
//==============================================================
type RouteOfAdministration int
const (
    RouteOfAdministrationOral RouteOfAdministration = iota
	RouteOfAdministrationIntravenous
	RouteOfAdministrationSubcutaneous
	RouteOfAdministrationIntramuscular
	RouteOfAdministrationTopical
	RouteOfAdministrationInhalation
)


//==============================================================
// SpecimenType Declaration
//==============================================================
type SpecimenType int
const (
    SpecimenTypeBlood SpecimenType = iota
	SpecimenTypeUrine
	SpecimenTypeSaliva
	SpecimenTypeSputum
	SpecimenTypeTissue
	SpecimenTypeCSF
	SpecimenTypeStool
)


//==============================================================
// ImagingModality Declaration
//==============================================================
type ImagingModality int
const (
    ImagingModalityXRay ImagingModality = iota
	ImagingModalityCT
	ImagingModalityMRI
	ImagingModalityUltrasound
	ImagingModalityPET
	ImagingModalityMammography
)


//==============================================================
// AnesthesiaType Declaration
//==============================================================
type AnesthesiaType int
const (
    AnesthesiaTypeNone AnesthesiaType = iota
	AnesthesiaTypeLocal
	AnesthesiaTypeRegional
	AnesthesiaTypeGeneral
	AnesthesiaTypeSedation
)


//==============================================================
// DispenseStatus Declaration
//==============================================================
type DispenseStatus int
const (
    DispenseStatusPreparation DispenseStatus = iota
	DispenseStatusInProgress
	DispenseStatusCompleted
	DispenseStatusCancelled
)


//==============================================================
// ResultStatus Declaration
//==============================================================
type ResultStatus int
const (
    ResultStatusRegistered ResultStatus = iota
	ResultStatusPartial
	ResultStatusFinal
	ResultStatusCorrected
	ResultStatusCancelled
)


//==============================================================
// ProcedureStatus Declaration
//==============================================================
type ProcedureStatus int
const (
    ProcedureStatusPlanned ProcedureStatus = iota
	ProcedureStatusInProgress
	ProcedureStatusCompleted
	ProcedureStatusAborted
)


//==============================================================
// DiagnosisCertainty Declaration
//==============================================================
type DiagnosisCertainty int
const (
    DiagnosisCertaintySuspected DiagnosisCertainty = iota
	DiagnosisCertaintyPresumptive
	DiagnosisCertaintyConfirmed
	DiagnosisCertaintyRuledOut
)


//==============================================================
// ObservationInterpretation Declaration
//==============================================================
type ObservationInterpretation int
const (
    ObservationInterpretationNormal ObservationInterpretation = iota
	ObservationInterpretationAbnormalLow
	ObservationInterpretationAbnormalHigh
	ObservationInterpretationCriticalLow
	ObservationInterpretationCriticalHigh
	ObservationInterpretationReactive
	ObservationInterpretationNonreactive
	ObservationInterpretationPositive
	ObservationInterpretationNegative
)


//==============================================================
// CarePlanStatus Declaration
//==============================================================
type CarePlanStatus int
const (
    CarePlanStatusDraft CarePlanStatus = iota
	CarePlanStatusActive
	CarePlanStatusSuspended
	CarePlanStatusCompleted
	CarePlanStatusCancelled
)


//==============================================================
// TaskStatus Declaration
//==============================================================
type TaskStatus int
const (
    TaskStatusRequested TaskStatus = iota
	TaskStatusAccepted
	TaskStatusInProgress
	TaskStatusCompleted
	TaskStatusCancelled
	TaskStatusFailed
)


//==============================================================
// AllergySeverity Declaration
//==============================================================
type AllergySeverity int
const (
    AllergySeverityMild AllergySeverity = iota
	AllergySeverityModerate
	AllergySeveritySevere
	AllergySeverityLifeThreatening
)


//==============================================================
// AllergyStatus Declaration
//==============================================================
type AllergyStatus int
const (
    AllergyStatusActive AllergyStatus = iota
	AllergyStatusInactive
	AllergyStatusResolved
	AllergyStatusEnteredInError
)


//==============================================================
// ConditionStatus Declaration
//==============================================================
type ConditionStatus int
const (
    ConditionStatusActive ConditionStatus = iota
	ConditionStatusRecurrence
	ConditionStatusRelapse
	ConditionStatusRemission
	ConditionStatusResolved
)


//==============================================================
// PayerType Declaration
//==============================================================
type PayerType int
const (
    PayerTypeCommercial PayerType = iota
	PayerTypeGovernment
	PayerTypeSelfInsured
)


//==============================================================
// InsurancePlanType Declaration
//==============================================================
type InsurancePlanType int
const (
    InsurancePlanTypeHMO InsurancePlanType = iota
	InsurancePlanTypePPO
	InsurancePlanTypeEPO
	InsurancePlanTypePOS
	InsurancePlanTypeIndemnity
	InsurancePlanTypeMedicareAdvantage
	InsurancePlanTypeMedicaidManagedCare
)


//==============================================================
// CoverageType Declaration
//==============================================================
type CoverageType int
const (
    CoverageTypeMedical CoverageType = iota
	CoverageTypePharmacy
	CoverageTypeDental
	CoverageTypeVision
	CoverageTypeBehavioralHealth
)


//==============================================================
// ClaimStatus Declaration
//==============================================================
type ClaimStatus int
const (
    ClaimStatusSubmitted ClaimStatus = iota
	ClaimStatusInProcess
	ClaimStatusPaid
	ClaimStatusDenied
	ClaimStatusAdjusted
	ClaimStatusVoid
)


//==============================================================
// AuthorizationStatus Declaration
//==============================================================
type AuthorizationStatus int
const (
    AuthorizationStatusRequested AuthorizationStatus = iota
	AuthorizationStatusPendingReview
	AuthorizationStatusApproved
	AuthorizationStatusDenied
	AuthorizationStatusExpired
)


//==============================================================
// InvoiceStatus Declaration
//==============================================================
type InvoiceStatus int
const (
    InvoiceStatusDraft InvoiceStatus = iota
	InvoiceStatusIssued
	InvoiceStatusPartiallyPaid
	InvoiceStatusPaid
	InvoiceStatusOverdue
	InvoiceStatusCancelled
)


//==============================================================
// PaymentMethod Declaration
//==============================================================
type PaymentMethod int
const (
    PaymentMethodACH PaymentMethod = iota
	PaymentMethodCheck
	PaymentMethodCreditCard
	PaymentMethodEFT
	PaymentMethodCash
)


//==============================================================
// DeviceType Declaration
//==============================================================
type DeviceType int
const (
    DeviceTypePacemaker DeviceType = iota
	DeviceTypeInsulinPump
	DeviceTypeBloodPressureMonitor
	DeviceTypeGlucoseMeter
	DeviceTypePulseOximeter
	DeviceTypeVentilator
	DeviceTypeInfusionPump
	DeviceTypeWearableTracker
)


//==============================================================
// DeviceConnectivityStatus Declaration
//==============================================================
type DeviceConnectivityStatus int
const (
    DeviceConnectivityStatusConnected DeviceConnectivityStatus = iota
	DeviceConnectivityStatusDisconnected
	DeviceConnectivityStatusStandby
	DeviceConnectivityStatusFault
)


//==============================================================
// SoftwareUpdateType Declaration
//==============================================================
type SoftwareUpdateType int
const (
    SoftwareUpdateTypeSecurityPatch SoftwareUpdateType = iota
	SoftwareUpdateTypeFeatureUpdate
	SoftwareUpdateTypeBugFix
	SoftwareUpdateTypeFirmwareUpgrade
)


//==============================================================
// SupplierTier Declaration
//==============================================================
type SupplierTier int
const (
    SupplierTierPrimary SupplierTier = iota
	SupplierTierSecondary
	SupplierTierDistributor
)

