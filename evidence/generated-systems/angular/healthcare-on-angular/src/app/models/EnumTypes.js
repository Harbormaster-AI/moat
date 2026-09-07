
// enum type FacilityType
export let FacilityType = {
	Hospital:"Hospital",
	Clinic:"Clinic",
	AmbulatorySurgeryCenter:"AmbulatorySurgeryCenter",
	UrgentCare:"UrgentCare",
	Laboratory:"Laboratory",
	ImagingCenter:"ImagingCenter",
	Pharmacy:"Pharmacy",
}

// enum type DepartmentType
export let DepartmentType = {
	Emergency:"Emergency",
	Cardiology:"Cardiology",
	Oncology:"Oncology",
	Orthopedics:"Orthopedics",
	Pediatrics:"Pediatrics",
	Radiology:"Radiology",
	Pathology:"Pathology",
	Pharmacy:"Pharmacy",
	IntensiveCare:"IntensiveCare",
}

// enum type CareSettingType
export let CareSettingType = {
	Inpatient:"Inpatient",
	Outpatient:"Outpatient",
	Emergency:"Emergency",
	HomeHealth:"HomeHealth",
	Telehealth:"Telehealth",
}

// enum type ClinicianType
export let ClinicianType = {
	Physician:"Physician",
	NursePractitioner:"NursePractitioner",
	PhysicianAssistant:"PhysicianAssistant",
	RegisteredNurse:"RegisteredNurse",
	Pharmacist:"Pharmacist",
	Therapist:"Therapist",
	Technician:"Technician",
}

// enum type ClinicianSpecialty
export let ClinicianSpecialty = {
	InternalMedicine:"InternalMedicine",
	FamilyMedicine:"FamilyMedicine",
	Cardiology:"Cardiology",
	Oncology:"Oncology",
	Orthopedics:"Orthopedics",
	Pediatrics:"Pediatrics",
	Radiology:"Radiology",
	Pathology:"Pathology",
	Anesthesiology:"Anesthesiology",
	Surgery:"Surgery",
	Psychiatry:"Psychiatry",
}

// enum type AdministrativeSex
export let AdministrativeSex = {
	Male:"Male",
	Female:"Female",
	Unknown:"Unknown",
}

// enum type BloodType
export let BloodType = {
	APositive:"APositive",
	ANegative:"ANegative",
	BPositive:"BPositive",
	BNegative:"BNegative",
	ABPositive:"ABPositive",
	ABNegative:"ABNegative",
	OPositive:"OPositive",
	ONegative:"ONegative",
}

// enum type AppointmentStatus
export let AppointmentStatus = {
	Proposed:"Proposed",
	Booked:"Booked",
	Arrived:"Arrived",
	Fulfilled:"Fulfilled",
	Cancelled:"Cancelled",
	NoShow:"NoShow",
	EnteredInError:"EnteredInError",
}

// enum type Priority
export let Priority = {
	Routine:"Routine",
	Urgent:"Urgent",
	Stat:"Stat",
}

// enum type EncounterStatus
export let EncounterStatus = {
	Planned:"Planned",
	InProgress:"InProgress",
	OnHold:"OnHold",
	Discharged:"Discharged",
	Cancelled:"Cancelled",
}

// enum type EncounterType
export let EncounterType = {
	Inpatient:"Inpatient",
	Outpatient:"Outpatient",
	Emergency:"Emergency",
	Observation:"Observation",
	Telemedicine:"Telemedicine",
}

// enum type AdmissionType
export let AdmissionType = {
	Elective:"Elective",
	Emergency:"Emergency",
	Urgent:"Urgent",
	Newborn:"Newborn",
	Trauma:"Trauma",
}

// enum type DischargeDisposition
export let DischargeDisposition = {
	Home:"Home",
	HomeWithHomeCare:"HomeWithHomeCare",
	SkilledNursingFacility:"SkilledNursingFacility",
	AcuteCareFacility:"AcuteCareFacility",
	Expired:"Expired",
	AgainstMedicalAdvice:"AgainstMedicalAdvice",
}

// enum type OrderStatus
export let OrderStatus = {
	Draft:"Draft",
	Active:"Active",
	OnHold:"OnHold",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type ClinicalOrderType
export let ClinicalOrderType = {
	Medication:"Medication",
	Laboratory:"Laboratory",
	Imaging:"Imaging",
	Procedure:"Procedure",
	Consultation:"Consultation",
}

// enum type RouteOfAdministration
export let RouteOfAdministration = {
	Oral:"Oral",
	Intravenous:"Intravenous",
	Subcutaneous:"Subcutaneous",
	Intramuscular:"Intramuscular",
	Topical:"Topical",
	Inhalation:"Inhalation",
}

// enum type SpecimenType
export let SpecimenType = {
	Blood:"Blood",
	Urine:"Urine",
	Saliva:"Saliva",
	Sputum:"Sputum",
	Tissue:"Tissue",
	CSF:"CSF",
	Stool:"Stool",
}

// enum type ImagingModality
export let ImagingModality = {
	XRay:"XRay",
	CT:"CT",
	MRI:"MRI",
	Ultrasound:"Ultrasound",
	PET:"PET",
	Mammography:"Mammography",
}

// enum type AnesthesiaType
export let AnesthesiaType = {
	None:"None",
	Local:"Local",
	Regional:"Regional",
	General:"General",
	Sedation:"Sedation",
}

// enum type DispenseStatus
export let DispenseStatus = {
	Preparation:"Preparation",
	InProgress:"InProgress",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type ResultStatus
export let ResultStatus = {
	Registered:"Registered",
	Partial:"Partial",
	Final:"Final",
	Corrected:"Corrected",
	Cancelled:"Cancelled",
}

// enum type ProcedureStatus
export let ProcedureStatus = {
	Planned:"Planned",
	InProgress:"InProgress",
	Completed:"Completed",
	Aborted:"Aborted",
}

// enum type DiagnosisCertainty
export let DiagnosisCertainty = {
	Suspected:"Suspected",
	Presumptive:"Presumptive",
	Confirmed:"Confirmed",
	RuledOut:"RuledOut",
}

// enum type ObservationInterpretation
export let ObservationInterpretation = {
	Normal:"Normal",
	AbnormalLow:"AbnormalLow",
	AbnormalHigh:"AbnormalHigh",
	CriticalLow:"CriticalLow",
	CriticalHigh:"CriticalHigh",
	Reactive:"Reactive",
	Nonreactive:"Nonreactive",
	Positive:"Positive",
	Negative:"Negative",
}

// enum type CarePlanStatus
export let CarePlanStatus = {
	Draft:"Draft",
	Active:"Active",
	Suspended:"Suspended",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type TaskStatus
export let TaskStatus = {
	Requested:"Requested",
	Accepted:"Accepted",
	InProgress:"InProgress",
	Completed:"Completed",
	Cancelled:"Cancelled",
	Failed:"Failed",
}

// enum type AllergySeverity
export let AllergySeverity = {
	Mild:"Mild",
	Moderate:"Moderate",
	Severe:"Severe",
	LifeThreatening:"LifeThreatening",
}

// enum type AllergyStatus
export let AllergyStatus = {
	Active:"Active",
	Inactive:"Inactive",
	Resolved:"Resolved",
	EnteredInError:"EnteredInError",
}

// enum type ConditionStatus
export let ConditionStatus = {
	Active:"Active",
	Recurrence:"Recurrence",
	Relapse:"Relapse",
	Remission:"Remission",
	Resolved:"Resolved",
}

// enum type PayerType
export let PayerType = {
	Commercial:"Commercial",
	Government:"Government",
	SelfInsured:"SelfInsured",
}

// enum type InsurancePlanType
export let InsurancePlanType = {
	HMO:"HMO",
	PPO:"PPO",
	EPO:"EPO",
	POS:"POS",
	Indemnity:"Indemnity",
	MedicareAdvantage:"MedicareAdvantage",
	MedicaidManagedCare:"MedicaidManagedCare",
}

// enum type CoverageType
export let CoverageType = {
	Medical:"Medical",
	Pharmacy:"Pharmacy",
	Dental:"Dental",
	Vision:"Vision",
	BehavioralHealth:"BehavioralHealth",
}

// enum type ClaimStatus
export let ClaimStatus = {
	Submitted:"Submitted",
	InProcess:"InProcess",
	Paid:"Paid",
	Denied:"Denied",
	Adjusted:"Adjusted",
	Void:"Void",
}

// enum type AuthorizationStatus
export let AuthorizationStatus = {
	Requested:"Requested",
	PendingReview:"PendingReview",
	Approved:"Approved",
	Denied:"Denied",
	Expired:"Expired",
}

// enum type InvoiceStatus
export let InvoiceStatus = {
	Draft:"Draft",
	Issued:"Issued",
	PartiallyPaid:"PartiallyPaid",
	Paid:"Paid",
	Overdue:"Overdue",
	Cancelled:"Cancelled",
}

// enum type PaymentMethod
export let PaymentMethod = {
	ACH:"ACH",
	Check:"Check",
	CreditCard:"CreditCard",
	EFT:"EFT",
	Cash:"Cash",
}

// enum type DeviceType
export let DeviceType = {
	Pacemaker:"Pacemaker",
	InsulinPump:"InsulinPump",
	BloodPressureMonitor:"BloodPressureMonitor",
	GlucoseMeter:"GlucoseMeter",
	PulseOximeter:"PulseOximeter",
	Ventilator:"Ventilator",
	InfusionPump:"InfusionPump",
	WearableTracker:"WearableTracker",
}

// enum type DeviceConnectivityStatus
export let DeviceConnectivityStatus = {
	Connected:"Connected",
	Disconnected:"Disconnected",
	Standby:"Standby",
	Fault:"Fault",
}

// enum type SoftwareUpdateType
export let SoftwareUpdateType = {
	SecurityPatch:"SecurityPatch",
	FeatureUpdate:"FeatureUpdate",
	BugFix:"BugFix",
	FirmwareUpgrade:"FirmwareUpgrade",
}

// enum type SupplierTier
export let SupplierTier = {
	Primary:"Primary",
	Secondary:"Secondary",
	Distributor:"Distributor",
}
