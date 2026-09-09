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
    path('Organization/', include('hrOnDjango.urls.OrganizationUrls')),
    path('Department/', include('hrOnDjango.urls.DepartmentUrls')),
    path('Location/', include('hrOnDjango.urls.LocationUrls')),
    path('CostCenter/', include('hrOnDjango.urls.CostCenterUrls')),
    path('JobFamily/', include('hrOnDjango.urls.JobFamilyUrls')),
    path('JobProfile/', include('hrOnDjango.urls.JobProfileUrls')),
    path('Competency/', include('hrOnDjango.urls.CompetencyUrls')),
    path('Position/', include('hrOnDjango.urls.PositionUrls')),
    path('Employee/', include('hrOnDjango.urls.EmployeeUrls')),
    path('EmploymentAssignment/', include('hrOnDjango.urls.EmploymentAssignmentUrls')),
    path('EmploymentContract/', include('hrOnDjango.urls.EmploymentContractUrls')),
    path('WorkSchedule/', include('hrOnDjango.urls.WorkScheduleUrls')),
    path('WorkShift/', include('hrOnDjango.urls.WorkShiftUrls')),
    path('ScheduleException/', include('hrOnDjango.urls.ScheduleExceptionUrls')),
    path('CompensationPackage/', include('hrOnDjango.urls.CompensationPackageUrls')),
    path('SalaryComponent/', include('hrOnDjango.urls.SalaryComponentUrls')),
    path('BonusPlan/', include('hrOnDjango.urls.BonusPlanUrls')),
    path('EquityGrant/', include('hrOnDjango.urls.EquityGrantUrls')),
    path('BenefitPlan/', include('hrOnDjango.urls.BenefitPlanUrls')),
    path('BenefitEnrollment/', include('hrOnDjango.urls.BenefitEnrollmentUrls')),
    path('Dependent/', include('hrOnDjango.urls.DependentUrls')),
    path('PayrollCalendar/', include('hrOnDjango.urls.PayrollCalendarUrls')),
    path('PayrollRun/', include('hrOnDjango.urls.PayrollRunUrls')),
    path('PayrollItem/', include('hrOnDjango.urls.PayrollItemUrls')),
    path('TaxWithholding/', include('hrOnDjango.urls.TaxWithholdingUrls')),
    path('PaymentMethod/', include('hrOnDjango.urls.PaymentMethodUrls')),
    path('Timesheet/', include('hrOnDjango.urls.TimesheetUrls')),
    path('TimeEntry/', include('hrOnDjango.urls.TimeEntryUrls')),
    path('Approval/', include('hrOnDjango.urls.ApprovalUrls')),
    path('LeavePolicy/', include('hrOnDjango.urls.LeavePolicyUrls')),
    path('LeaveRequest/', include('hrOnDjango.urls.LeaveRequestUrls')),
    path('PerformanceCycle/', include('hrOnDjango.urls.PerformanceCycleUrls')),
    path('Goal/', include('hrOnDjango.urls.GoalUrls')),
    path('PerformanceReview/', include('hrOnDjango.urls.PerformanceReviewUrls')),
    path('CompetencyRating/', include('hrOnDjango.urls.CompetencyRatingUrls')),
    path('TrainingCourse/', include('hrOnDjango.urls.TrainingCourseUrls')),
    path('TrainingEnrollment/', include('hrOnDjango.urls.TrainingEnrollmentUrls')),
    path('Certification/', include('hrOnDjango.urls.CertificationUrls')),
    path('JobRequisition/', include('hrOnDjango.urls.JobRequisitionUrls')),
    path('Candidate/', include('hrOnDjango.urls.CandidateUrls')),
    path('JobApplication/', include('hrOnDjango.urls.JobApplicationUrls')),
    path('Interview/', include('hrOnDjango.urls.InterviewUrls')),
    path('Screening/', include('hrOnDjango.urls.ScreeningUrls')),
    path('Offer/', include('hrOnDjango.urls.OfferUrls')),
    path('OnboardingTask/', include('hrOnDjango.urls.OnboardingTaskUrls')),
    path('BackgroundCheck/', include('hrOnDjango.urls.BackgroundCheckUrls')),
    path('Document/', include('hrOnDjango.urls.DocumentUrls')),
    path('Policy/', include('hrOnDjango.urls.PolicyUrls')),
    path('PolicyAcknowledgement/', include('hrOnDjango.urls.PolicyAcknowledgementUrls')),
    path('Termination/', include('hrOnDjango.urls.TerminationUrls')),
    path('WorkAuthorization/', include('hrOnDjango.urls.WorkAuthorizationUrls')),
    path('BankAccount/', include('hrOnDjango.urls.BankAccountUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]