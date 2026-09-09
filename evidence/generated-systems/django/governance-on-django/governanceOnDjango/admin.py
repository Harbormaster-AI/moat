from django.contrib import admin

# Register your models here.
from .models.Organization import Organization
from .models.GovernanceBody import GovernanceBody
from .models.Person import Person
from .models.Role import Role
from .models.RoleAssignment import RoleAssignment
from .models.Policy import Policy
from .models.Procedure import Procedure
from .models.Regulation import Regulation
from .models.Obligation import Obligation
from .models.Control import Control
from .models.ControlTest_ import ControlTest_
from .models.Evidence import Evidence
from .models.Risk import Risk
from .models.RiskAssessment import RiskAssessment
from .models.ComplianceProgram import ComplianceProgram
from .models.ComplianceRequirement import ComplianceRequirement
from .models.Attestation import Attestation
from .models.AuditProgram import AuditProgram
from .models.AuditEngagement import AuditEngagement
from .models.AuditWorkpaper import AuditWorkpaper
from .models.AuditFinding import AuditFinding
from .models.CorrectiveAction import CorrectiveAction
from .models.Issue import Issue
from .models.BusinessUnit import BusinessUnit
from .models.DataProcessingActivity import DataProcessingActivity
from .models.DataCategory import DataCategory
from .models.System_ import System_
from .models.PrivacyNotice import PrivacyNotice
from .models.DataSubjectRequest import DataSubjectRequest
from .models.RecordsRepository import RecordsRepository
from .models.Record_ import Record_
from .models.RetentionSchedule import RetentionSchedule
from .models.DispositionReview import DispositionReview
from .models.LegalHold import LegalHold
from .models.Matter import Matter
from .models.ThirdParty import ThirdParty
from .models.ThirdPartyAssessment import ThirdPartyAssessment
from .models.Contract import Contract
from .models.Exception_ import Exception_
from .models.Consent import Consent
from .models.DataBreach import DataBreach

# Need to add this for each model that requires managing

admin.site.register(Organization)
admin.site.register(GovernanceBody)
admin.site.register(Person)
admin.site.register(Role)
admin.site.register(RoleAssignment)
admin.site.register(Policy)
admin.site.register(Procedure)
admin.site.register(Regulation)
admin.site.register(Obligation)
admin.site.register(Control)
admin.site.register(ControlTest_)
admin.site.register(Evidence)
admin.site.register(Risk)
admin.site.register(RiskAssessment)
admin.site.register(ComplianceProgram)
admin.site.register(ComplianceRequirement)
admin.site.register(Attestation)
admin.site.register(AuditProgram)
admin.site.register(AuditEngagement)
admin.site.register(AuditWorkpaper)
admin.site.register(AuditFinding)
admin.site.register(CorrectiveAction)
admin.site.register(Issue)
admin.site.register(BusinessUnit)
admin.site.register(DataProcessingActivity)
admin.site.register(DataCategory)
admin.site.register(System_)
admin.site.register(PrivacyNotice)
admin.site.register(DataSubjectRequest)
admin.site.register(RecordsRepository)
admin.site.register(Record_)
admin.site.register(RetentionSchedule)
admin.site.register(DispositionReview)
admin.site.register(LegalHold)
admin.site.register(Matter)
admin.site.register(ThirdParty)
admin.site.register(ThirdPartyAssessment)
admin.site.register(Contract)
admin.site.register(Exception_)
admin.site.register(Consent)
admin.site.register(DataBreach)
