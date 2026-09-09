from django.db import models
from hrOnDjango.models.EmploymentStatus import EmploymentStatus

#======================================================================
# 
# Encapsulates data for model Employee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Employee Declaration
#======================================================================
class Employee (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	employeeNumber = models.CharField(max_length=200, null=True)
	name = PersonName
	workEmail = Email
	workPhone = PhoneNumber
	dateOfHire = models.DateField(null=True)
	nationalId = NationalID
	manager = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	directReports = models.ManyToManyField('Employee',  blank=True, related_name='+')
	department = models.ForeignKey('Department', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	primaryLocation = models.ForeignKey('Location', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	costCenter = models.ForeignKey('CostCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employmentAssignments = models.ManyToManyField('EmploymentAssignment',  blank=True, related_name='+')
	contracts = models.ManyToManyField('EmploymentContract',  blank=True, related_name='+')
	benefitEnrollments = models.ManyToManyField('BenefitEnrollment',  blank=True, related_name='+')
	timesheets = models.ManyToManyField('Timesheet',  blank=True, related_name='+')
	leaveRequests = models.ManyToManyField('LeaveRequest',  blank=True, related_name='+')
	performanceReviews = models.ManyToManyField('PerformanceReview',  blank=True, related_name='+')
	trainingEnrollments = models.ManyToManyField('TrainingEnrollment',  blank=True, related_name='+')
	workAuthorizations = models.ManyToManyField('WorkAuthorization',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EmploymentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.employeeNumber
		str = str + self.name
		str = str + self.workEmail
		str = str + self.workPhone
		str = str + self.dateOfHire
		str = str + self.nationalId
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Employee";
    
	def objectType(self):
		return "Employee";
