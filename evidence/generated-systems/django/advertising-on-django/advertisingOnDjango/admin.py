from django.contrib import admin

# Register your models here.
from .models.Agency import Agency
from .models.Team import Team
from .models.User import User
from .models.Advertiser import Advertiser
from .models.BillingProfile import BillingProfile
from .models.PaymentMethod import PaymentMethod
from .models.AdAccount import AdAccount
from .models.DSP import DSP
from .models.Campaign import Campaign
from .models.KPI import KPI
from .models.AudienceSegment import AudienceSegment
from .models.DataProvider import DataProvider
from .models.LineItem import LineItem
from .models.TargetingProfile import TargetingProfile
from .models.DeviceCriterion import DeviceCriterion
from .models.BrandSafetyPolicy import BrandSafetyPolicy
from .models.ContentCategory import ContentCategory
from .models.Publisher import Publisher
from .models.InventorySource import InventorySource
from .models.AdSlot import AdSlot
from .models.Deal import Deal
from .models.Placement import Placement
from .models.CreativeAsset import CreativeAsset
from .models.CreativeFile import CreativeFile
from .models.CreativeVariation import CreativeVariation
from .models.CreativeApproval import CreativeApproval
from .models.TrackingPixel import TrackingPixel
from .models.ConversionEvent import ConversionEvent
from .models.PerformanceMetric import PerformanceMetric
from .models.Report import Report
from .models.InsertionOrder import InsertionOrder
from .models.RateCard import RateCard
from .models.Rate import Rate
from .models.Experiment import Experiment
from .models.ExperimentVariant import ExperimentVariant
from .models.GeoRegion import GeoRegion

# Need to add this for each model that requires managing

admin.site.register(Agency)
admin.site.register(Team)
admin.site.register(User)
admin.site.register(Advertiser)
admin.site.register(BillingProfile)
admin.site.register(PaymentMethod)
admin.site.register(AdAccount)
admin.site.register(DSP)
admin.site.register(Campaign)
admin.site.register(KPI)
admin.site.register(AudienceSegment)
admin.site.register(DataProvider)
admin.site.register(LineItem)
admin.site.register(TargetingProfile)
admin.site.register(DeviceCriterion)
admin.site.register(BrandSafetyPolicy)
admin.site.register(ContentCategory)
admin.site.register(Publisher)
admin.site.register(InventorySource)
admin.site.register(AdSlot)
admin.site.register(Deal)
admin.site.register(Placement)
admin.site.register(CreativeAsset)
admin.site.register(CreativeFile)
admin.site.register(CreativeVariation)
admin.site.register(CreativeApproval)
admin.site.register(TrackingPixel)
admin.site.register(ConversionEvent)
admin.site.register(PerformanceMetric)
admin.site.register(Report)
admin.site.register(InsertionOrder)
admin.site.register(RateCard)
admin.site.register(Rate)
admin.site.register(Experiment)
admin.site.register(ExperimentVariant)
admin.site.register(GeoRegion)
