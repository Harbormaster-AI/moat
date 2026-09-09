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
    path('AerospaceManufacturer/', include('aerospaceOnDjango.urls.AerospaceManufacturerUrls')),
    path('AircraftProgram/', include('aerospaceOnDjango.urls.AircraftProgramUrls')),
    path('AircraftFamily/', include('aerospaceOnDjango.urls.AircraftFamilyUrls')),
    path('AircraftModel/', include('aerospaceOnDjango.urls.AircraftModelUrls')),
    path('EngineType/', include('aerospaceOnDjango.urls.EngineTypeUrls')),
    path('AircraftVariant/', include('aerospaceOnDjango.urls.AircraftVariantUrls')),
    path('AvionicsSuite/', include('aerospaceOnDjango.urls.AvionicsSuiteUrls')),
    path('APU/', include('aerospaceOnDjango.urls.APUUrls')),
    path('LandingGear/', include('aerospaceOnDjango.urls.LandingGearUrls')),
    path('AircraftOption/', include('aerospaceOnDjango.urls.AircraftOptionUrls')),
    path('AircraftPackage/', include('aerospaceOnDjango.urls.AircraftPackageUrls')),
    path('Supplier/', include('aerospaceOnDjango.urls.SupplierUrls')),
    path('Component_/', include('aerospaceOnDjango.urls.Component_Urls')),
    path('Plant/', include('aerospaceOnDjango.urls.PlantUrls')),
    path('ProductionLine/', include('aerospaceOnDjango.urls.ProductionLineUrls')),
    path('WorkCenter/', include('aerospaceOnDjango.urls.WorkCenterUrls')),
    path('ProductionOrder/', include('aerospaceOnDjango.urls.ProductionOrderUrls')),
    path('BuildSchedule/', include('aerospaceOnDjango.urls.BuildScheduleUrls')),
    path('Warehouse/', include('aerospaceOnDjango.urls.WarehouseUrls')),
    path('InventoryItem/', include('aerospaceOnDjango.urls.InventoryItemUrls')),
    path('Operator/', include('aerospaceOnDjango.urls.OperatorUrls')),
    path('AircraftOrder/', include('aerospaceOnDjango.urls.AircraftOrderUrls')),
    path('Quote/', include('aerospaceOnDjango.urls.QuoteUrls')),
    path('PurchaseAgreement/', include('aerospaceOnDjango.urls.PurchaseAgreementUrls')),
    path('Aircraft/', include('aerospaceOnDjango.urls.AircraftUrls')),
    path('Registration/', include('aerospaceOnDjango.urls.RegistrationUrls')),
    path('Warranty/', include('aerospaceOnDjango.urls.WarrantyUrls')),
    path('CabinLayout/', include('aerospaceOnDjango.urls.CabinLayoutUrls')),
    path('MROFacility/', include('aerospaceOnDjango.urls.MROFacilityUrls')),
    path('MaintenanceAppointment/', include('aerospaceOnDjango.urls.MaintenanceAppointmentUrls')),
    path('MaintenanceWorkOrder/', include('aerospaceOnDjango.urls.MaintenanceWorkOrderUrls')),
    path('AirworthinessDirective/', include('aerospaceOnDjango.urls.AirworthinessDirectiveUrls')),
    path('ServiceBulletin/', include('aerospaceOnDjango.urls.ServiceBulletinUrls')),
    path('ConnectedAircraft/', include('aerospaceOnDjango.urls.ConnectedAircraftUrls')),
    path('FlightHealthEvent/', include('aerospaceOnDjango.urls.FlightHealthEventUrls')),
    path('SoftwareLoad/', include('aerospaceOnDjango.urls.SoftwareLoadUrls')),
    path('TypeCertificate/', include('aerospaceOnDjango.urls.TypeCertificateUrls')),
    path('ProductionCertificate/', include('aerospaceOnDjango.urls.ProductionCertificateUrls')),
    path('SalesRegion/', include('aerospaceOnDjango.urls.SalesRegionUrls')),
    path('SalesCampaign/', include('aerospaceOnDjango.urls.SalesCampaignUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]