from django.contrib import admin

# Register your models here.
from .models.HealthSystem import HealthSystem
from .models.Facility import Facility
from .models.Department import Department
from .models.CareTeam import CareTeam
from .models.Clinician import Clinician
from .models.Patient import Patient
from .models.Appointment import Appointment
from .models.Encounter import Encounter
from .models.Admission import Admission
from .models.Discharge import Discharge
from .models.ClinicalOrder import ClinicalOrder
from .models.MedicationOrder import MedicationOrder
from .models.Laboratory import Laboratory
from .models.LaboratoryOrder import LaboratoryOrder
from .models.LabResult import LabResult
from .models.ImagingCenter import ImagingCenter
from .models.ImagingOrder import ImagingOrder
from .models.ImagingReport import ImagingReport
from .models.ProcedureOrder import ProcedureOrder
from .models.Procedure import Procedure
from .models.Pharmacy import Pharmacy
from .models.MedicationDispense import MedicationDispense
from .models.Diagnosis import Diagnosis
from .models.Observation import Observation
from .models.CarePlan import CarePlan
from .models.CareTask import CareTask
from .models.Allergy import Allergy
from .models.Condition import Condition
from .models.InsurancePayer import InsurancePayer
from .models.InsurancePlan import InsurancePlan
from .models.Coverage import Coverage
from .models.Claim import Claim
from .models.Authorization import Authorization
from .models.Invoice import Invoice
from .models.Payment import Payment
from .models.MedicalDevice import MedicalDevice
from .models.SoftwareUpdate import SoftwareUpdate
from .models.MedicalSupplier import MedicalSupplier
from .models.InventoryItem import InventoryItem

# Need to add this for each model that requires managing

admin.site.register(HealthSystem)
admin.site.register(Facility)
admin.site.register(Department)
admin.site.register(CareTeam)
admin.site.register(Clinician)
admin.site.register(Patient)
admin.site.register(Appointment)
admin.site.register(Encounter)
admin.site.register(Admission)
admin.site.register(Discharge)
admin.site.register(ClinicalOrder)
admin.site.register(MedicationOrder)
admin.site.register(Laboratory)
admin.site.register(LaboratoryOrder)
admin.site.register(LabResult)
admin.site.register(ImagingCenter)
admin.site.register(ImagingOrder)
admin.site.register(ImagingReport)
admin.site.register(ProcedureOrder)
admin.site.register(Procedure)
admin.site.register(Pharmacy)
admin.site.register(MedicationDispense)
admin.site.register(Diagnosis)
admin.site.register(Observation)
admin.site.register(CarePlan)
admin.site.register(CareTask)
admin.site.register(Allergy)
admin.site.register(Condition)
admin.site.register(InsurancePayer)
admin.site.register(InsurancePlan)
admin.site.register(Coverage)
admin.site.register(Claim)
admin.site.register(Authorization)
admin.site.register(Invoice)
admin.site.register(Payment)
admin.site.register(MedicalDevice)
admin.site.register(SoftwareUpdate)
admin.site.register(MedicalSupplier)
admin.site.register(InventoryItem)
