from django.contrib import admin

# Register your models here.
from .models.AerospaceManufacturer import AerospaceManufacturer
from .models.AircraftProgram import AircraftProgram
from .models.AircraftFamily import AircraftFamily
from .models.AircraftModel import AircraftModel
from .models.EngineType import EngineType
from .models.AircraftVariant import AircraftVariant
from .models.AvionicsSuite import AvionicsSuite
from .models.APU import APU
from .models.LandingGear import LandingGear
from .models.AircraftOption import AircraftOption
from .models.AircraftPackage import AircraftPackage
from .models.Supplier import Supplier
from .models.Component_ import Component_
from .models.Plant import Plant
from .models.ProductionLine import ProductionLine
from .models.WorkCenter import WorkCenter
from .models.ProductionOrder import ProductionOrder
from .models.BuildSchedule import BuildSchedule
from .models.Warehouse import Warehouse
from .models.InventoryItem import InventoryItem
from .models.Operator import Operator
from .models.AircraftOrder import AircraftOrder
from .models.Quote import Quote
from .models.PurchaseAgreement import PurchaseAgreement
from .models.Aircraft import Aircraft
from .models.Registration import Registration
from .models.Warranty import Warranty
from .models.CabinLayout import CabinLayout
from .models.MROFacility import MROFacility
from .models.MaintenanceAppointment import MaintenanceAppointment
from .models.MaintenanceWorkOrder import MaintenanceWorkOrder
from .models.AirworthinessDirective import AirworthinessDirective
from .models.ServiceBulletin import ServiceBulletin
from .models.ConnectedAircraft import ConnectedAircraft
from .models.FlightHealthEvent import FlightHealthEvent
from .models.SoftwareLoad import SoftwareLoad
from .models.TypeCertificate import TypeCertificate
from .models.ProductionCertificate import ProductionCertificate
from .models.SalesRegion import SalesRegion
from .models.SalesCampaign import SalesCampaign

# Need to add this for each model that requires managing

admin.site.register(AerospaceManufacturer)
admin.site.register(AircraftProgram)
admin.site.register(AircraftFamily)
admin.site.register(AircraftModel)
admin.site.register(EngineType)
admin.site.register(AircraftVariant)
admin.site.register(AvionicsSuite)
admin.site.register(APU)
admin.site.register(LandingGear)
admin.site.register(AircraftOption)
admin.site.register(AircraftPackage)
admin.site.register(Supplier)
admin.site.register(Component_)
admin.site.register(Plant)
admin.site.register(ProductionLine)
admin.site.register(WorkCenter)
admin.site.register(ProductionOrder)
admin.site.register(BuildSchedule)
admin.site.register(Warehouse)
admin.site.register(InventoryItem)
admin.site.register(Operator)
admin.site.register(AircraftOrder)
admin.site.register(Quote)
admin.site.register(PurchaseAgreement)
admin.site.register(Aircraft)
admin.site.register(Registration)
admin.site.register(Warranty)
admin.site.register(CabinLayout)
admin.site.register(MROFacility)
admin.site.register(MaintenanceAppointment)
admin.site.register(MaintenanceWorkOrder)
admin.site.register(AirworthinessDirective)
admin.site.register(ServiceBulletin)
admin.site.register(ConnectedAircraft)
admin.site.register(FlightHealthEvent)
admin.site.register(SoftwareLoad)
admin.site.register(TypeCertificate)
admin.site.register(ProductionCertificate)
admin.site.register(SalesRegion)
admin.site.register(SalesCampaign)
