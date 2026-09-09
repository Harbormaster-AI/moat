from django.db import models
from hrOnDjango.models.EmploymentType import EmploymentType
from hrOnDjango.models.ContractStatus import ContractStatus
from hrOnDjango.models.PayFrequency import PayFrequency

#======================================================================
# 
# Encapsulates data for model EmploymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentContract Declaration
#======================================================================
class EmploymentContract (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	contractNumber = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	workHoursPerWeek = models.CharField(max_length=64, null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	compensationPackage = models.OneToOneField('CompensationPackage', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workSchedule = models.ForeignKey('WorkSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('Location', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payrollCalendar = models.ForeignKey('PayrollCalendar', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employmentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EmploymentType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ContractStatus])
	payFrequency = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayFrequency])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.contractNumber
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.workHoursPerWeek
		str = str + self.employmentType
		str = str + self.status
		str = str + self.payFrequency
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "EmploymentContract";
    
	def objectType(self):
		return "EmploymentContract";
