from django.contrib import admin

# Register your models here.
from .models.Organization import Organization
from .models.Department import Department
from .models.Location import Location
from .models.CostCenter import CostCenter
from .models.JobFamily import JobFamily
from .models.JobProfile import JobProfile
from .models.Competency import Competency
from .models.Position import Position
from .models.Employee import Employee
from .models.EmploymentAssignment import EmploymentAssignment
from .models.EmploymentContract import EmploymentContract
from .models.WorkSchedule import WorkSchedule
from .models.WorkShift import WorkShift
from .models.ScheduleException import ScheduleException
from .models.CompensationPackage import CompensationPackage
from .models.SalaryComponent import SalaryComponent
from .models.BonusPlan import BonusPlan
from .models.EquityGrant import EquityGrant
from .models.BenefitPlan import BenefitPlan
from .models.BenefitEnrollment import BenefitEnrollment
from .models.Dependent import Dependent
from .models.PayrollCalendar import PayrollCalendar
from .models.PayrollRun import PayrollRun
from .models.PayrollItem import PayrollItem
from .models.TaxWithholding import TaxWithholding
from .models.PaymentMethod import PaymentMethod
from .models.Timesheet import Timesheet
from .models.TimeEntry import TimeEntry
from .models.Approval import Approval
from .models.LeavePolicy import LeavePolicy
from .models.LeaveRequest import LeaveRequest
from .models.PerformanceCycle import PerformanceCycle
from .models.Goal import Goal
from .models.PerformanceReview import PerformanceReview
from .models.CompetencyRating import CompetencyRating
from .models.TrainingCourse import TrainingCourse
from .models.TrainingEnrollment import TrainingEnrollment
from .models.Certification import Certification
from .models.JobRequisition import JobRequisition
from .models.Candidate import Candidate
from .models.JobApplication import JobApplication
from .models.Interview import Interview
from .models.Screening import Screening
from .models.Offer import Offer
from .models.OnboardingTask import OnboardingTask
from .models.BackgroundCheck import BackgroundCheck
from .models.Document import Document
from .models.Policy import Policy
from .models.PolicyAcknowledgement import PolicyAcknowledgement
from .models.Termination import Termination
from .models.WorkAuthorization import WorkAuthorization
from .models.BankAccount import BankAccount

# Need to add this for each model that requires managing

admin.site.register(Organization)
admin.site.register(Department)
admin.site.register(Location)
admin.site.register(CostCenter)
admin.site.register(JobFamily)
admin.site.register(JobProfile)
admin.site.register(Competency)
admin.site.register(Position)
admin.site.register(Employee)
admin.site.register(EmploymentAssignment)
admin.site.register(EmploymentContract)
admin.site.register(WorkSchedule)
admin.site.register(WorkShift)
admin.site.register(ScheduleException)
admin.site.register(CompensationPackage)
admin.site.register(SalaryComponent)
admin.site.register(BonusPlan)
admin.site.register(EquityGrant)
admin.site.register(BenefitPlan)
admin.site.register(BenefitEnrollment)
admin.site.register(Dependent)
admin.site.register(PayrollCalendar)
admin.site.register(PayrollRun)
admin.site.register(PayrollItem)
admin.site.register(TaxWithholding)
admin.site.register(PaymentMethod)
admin.site.register(Timesheet)
admin.site.register(TimeEntry)
admin.site.register(Approval)
admin.site.register(LeavePolicy)
admin.site.register(LeaveRequest)
admin.site.register(PerformanceCycle)
admin.site.register(Goal)
admin.site.register(PerformanceReview)
admin.site.register(CompetencyRating)
admin.site.register(TrainingCourse)
admin.site.register(TrainingEnrollment)
admin.site.register(Certification)
admin.site.register(JobRequisition)
admin.site.register(Candidate)
admin.site.register(JobApplication)
admin.site.register(Interview)
admin.site.register(Screening)
admin.site.register(Offer)
admin.site.register(OnboardingTask)
admin.site.register(BackgroundCheck)
admin.site.register(Document)
admin.site.register(Policy)
admin.site.register(PolicyAcknowledgement)
admin.site.register(Termination)
admin.site.register(WorkAuthorization)
admin.site.register(BankAccount)
