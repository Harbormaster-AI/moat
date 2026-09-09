"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('HealthSystem/', include('healthcareOnDjango.urls.HealthSystemUrls')),
    path('Facility/', include('healthcareOnDjango.urls.FacilityUrls')),
    path('Department/', include('healthcareOnDjango.urls.DepartmentUrls')),
    path('CareTeam/', include('healthcareOnDjango.urls.CareTeamUrls')),
    path('Clinician/', include('healthcareOnDjango.urls.ClinicianUrls')),
    path('Patient/', include('healthcareOnDjango.urls.PatientUrls')),
    path('Appointment/', include('healthcareOnDjango.urls.AppointmentUrls')),
    path('Encounter/', include('healthcareOnDjango.urls.EncounterUrls')),
    path('Admission/', include('healthcareOnDjango.urls.AdmissionUrls')),
    path('Discharge/', include('healthcareOnDjango.urls.DischargeUrls')),
    path('ClinicalOrder/', include('healthcareOnDjango.urls.ClinicalOrderUrls')),
    path('MedicationOrder/', include('healthcareOnDjango.urls.MedicationOrderUrls')),
    path('Laboratory/', include('healthcareOnDjango.urls.LaboratoryUrls')),
    path('LaboratoryOrder/', include('healthcareOnDjango.urls.LaboratoryOrderUrls')),
    path('LabResult/', include('healthcareOnDjango.urls.LabResultUrls')),
    path('ImagingCenter/', include('healthcareOnDjango.urls.ImagingCenterUrls')),
    path('ImagingOrder/', include('healthcareOnDjango.urls.ImagingOrderUrls')),
    path('ImagingReport/', include('healthcareOnDjango.urls.ImagingReportUrls')),
    path('ProcedureOrder/', include('healthcareOnDjango.urls.ProcedureOrderUrls')),
    path('Procedure/', include('healthcareOnDjango.urls.ProcedureUrls')),
    path('Pharmacy/', include('healthcareOnDjango.urls.PharmacyUrls')),
    path('MedicationDispense/', include('healthcareOnDjango.urls.MedicationDispenseUrls')),
    path('Diagnosis/', include('healthcareOnDjango.urls.DiagnosisUrls')),
    path('Observation/', include('healthcareOnDjango.urls.ObservationUrls')),
    path('CarePlan/', include('healthcareOnDjango.urls.CarePlanUrls')),
    path('CareTask/', include('healthcareOnDjango.urls.CareTaskUrls')),
    path('Allergy/', include('healthcareOnDjango.urls.AllergyUrls')),
    path('Condition/', include('healthcareOnDjango.urls.ConditionUrls')),
    path('InsurancePayer/', include('healthcareOnDjango.urls.InsurancePayerUrls')),
    path('InsurancePlan/', include('healthcareOnDjango.urls.InsurancePlanUrls')),
    path('Coverage/', include('healthcareOnDjango.urls.CoverageUrls')),
    path('Claim/', include('healthcareOnDjango.urls.ClaimUrls')),
    path('Authorization/', include('healthcareOnDjango.urls.AuthorizationUrls')),
    path('Invoice/', include('healthcareOnDjango.urls.InvoiceUrls')),
    path('Payment/', include('healthcareOnDjango.urls.PaymentUrls')),
    path('MedicalDevice/', include('healthcareOnDjango.urls.MedicalDeviceUrls')),
    path('SoftwareUpdate/', include('healthcareOnDjango.urls.SoftwareUpdateUrls')),
    path('MedicalSupplier/', include('healthcareOnDjango.urls.MedicalSupplierUrls')),
    path('InventoryItem/', include('healthcareOnDjango.urls.InventoryItemUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]