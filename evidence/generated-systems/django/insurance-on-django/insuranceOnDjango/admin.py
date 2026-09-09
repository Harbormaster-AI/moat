from django.contrib import admin

# Register your models here.
from .models.Insurer import Insurer
from .models.InsuranceProduct import InsuranceProduct
from .models.CoverageDefinition import CoverageDefinition
from .models.Distributor import Distributor
from .models.Agent import Agent
from .models.Customer import Customer
from .models.Application import Application
from .models.Quote import Quote
from .models.UnderwritingDecision import UnderwritingDecision
from .models.Underwriter import Underwriter
from .models.Policy import Policy
from .models.Endorsement import Endorsement
from .models.PolicyCoverage import PolicyCoverage
from .models.InsuredObject import InsuredObject
from .models.Beneficiary import Beneficiary
from .models.BillingAccount import BillingAccount
from .models.Invoice import Invoice
from .models.Payment import Payment
from .models.Claim import Claim
from .models.Incident import Incident
from .models.Exposure import Exposure
from .models.Adjuster import Adjuster
from .models.ClaimReserve import ClaimReserve
from .models.ClaimPayment import ClaimPayment
from .models.ServiceProvider import ServiceProvider
from .models.ReinsuranceAgreement import ReinsuranceAgreement
from .models.SubrogationRecovery import SubrogationRecovery
from .models.ThirdParty import ThirdParty
from .models.Document import Document

# Need to add this for each model that requires managing

admin.site.register(Insurer)
admin.site.register(InsuranceProduct)
admin.site.register(CoverageDefinition)
admin.site.register(Distributor)
admin.site.register(Agent)
admin.site.register(Customer)
admin.site.register(Application)
admin.site.register(Quote)
admin.site.register(UnderwritingDecision)
admin.site.register(Underwriter)
admin.site.register(Policy)
admin.site.register(Endorsement)
admin.site.register(PolicyCoverage)
admin.site.register(InsuredObject)
admin.site.register(Beneficiary)
admin.site.register(BillingAccount)
admin.site.register(Invoice)
admin.site.register(Payment)
admin.site.register(Claim)
admin.site.register(Incident)
admin.site.register(Exposure)
admin.site.register(Adjuster)
admin.site.register(ClaimReserve)
admin.site.register(ClaimPayment)
admin.site.register(ServiceProvider)
admin.site.register(ReinsuranceAgreement)
admin.site.register(SubrogationRecovery)
admin.site.register(ThirdParty)
admin.site.register(Document)
