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
    path('Organization/', include('governanceOnDjango.urls.OrganizationUrls')),
    path('GovernanceBody/', include('governanceOnDjango.urls.GovernanceBodyUrls')),
    path('Person/', include('governanceOnDjango.urls.PersonUrls')),
    path('Role/', include('governanceOnDjango.urls.RoleUrls')),
    path('RoleAssignment/', include('governanceOnDjango.urls.RoleAssignmentUrls')),
    path('Policy/', include('governanceOnDjango.urls.PolicyUrls')),
    path('Procedure/', include('governanceOnDjango.urls.ProcedureUrls')),
    path('Regulation/', include('governanceOnDjango.urls.RegulationUrls')),
    path('Obligation/', include('governanceOnDjango.urls.ObligationUrls')),
    path('Control/', include('governanceOnDjango.urls.ControlUrls')),
    path('ControlTest_/', include('governanceOnDjango.urls.ControlTest_Urls')),
    path('Evidence/', include('governanceOnDjango.urls.EvidenceUrls')),
    path('Risk/', include('governanceOnDjango.urls.RiskUrls')),
    path('RiskAssessment/', include('governanceOnDjango.urls.RiskAssessmentUrls')),
    path('ComplianceProgram/', include('governanceOnDjango.urls.ComplianceProgramUrls')),
    path('ComplianceRequirement/', include('governanceOnDjango.urls.ComplianceRequirementUrls')),
    path('Attestation/', include('governanceOnDjango.urls.AttestationUrls')),
    path('AuditProgram/', include('governanceOnDjango.urls.AuditProgramUrls')),
    path('AuditEngagement/', include('governanceOnDjango.urls.AuditEngagementUrls')),
    path('AuditWorkpaper/', include('governanceOnDjango.urls.AuditWorkpaperUrls')),
    path('AuditFinding/', include('governanceOnDjango.urls.AuditFindingUrls')),
    path('CorrectiveAction/', include('governanceOnDjango.urls.CorrectiveActionUrls')),
    path('Issue/', include('governanceOnDjango.urls.IssueUrls')),
    path('BusinessUnit/', include('governanceOnDjango.urls.BusinessUnitUrls')),
    path('DataProcessingActivity/', include('governanceOnDjango.urls.DataProcessingActivityUrls')),
    path('DataCategory/', include('governanceOnDjango.urls.DataCategoryUrls')),
    path('System_/', include('governanceOnDjango.urls.System_Urls')),
    path('PrivacyNotice/', include('governanceOnDjango.urls.PrivacyNoticeUrls')),
    path('DataSubjectRequest/', include('governanceOnDjango.urls.DataSubjectRequestUrls')),
    path('RecordsRepository/', include('governanceOnDjango.urls.RecordsRepositoryUrls')),
    path('Record_/', include('governanceOnDjango.urls.Record_Urls')),
    path('RetentionSchedule/', include('governanceOnDjango.urls.RetentionScheduleUrls')),
    path('DispositionReview/', include('governanceOnDjango.urls.DispositionReviewUrls')),
    path('LegalHold/', include('governanceOnDjango.urls.LegalHoldUrls')),
    path('Matter/', include('governanceOnDjango.urls.MatterUrls')),
    path('ThirdParty/', include('governanceOnDjango.urls.ThirdPartyUrls')),
    path('ThirdPartyAssessment/', include('governanceOnDjango.urls.ThirdPartyAssessmentUrls')),
    path('Contract/', include('governanceOnDjango.urls.ContractUrls')),
    path('Exception_/', include('governanceOnDjango.urls.Exception_Urls')),
    path('Consent/', include('governanceOnDjango.urls.ConsentUrls')),
    path('DataBreach/', include('governanceOnDjango.urls.DataBreachUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]