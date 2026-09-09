from django.db import models
from hrOnDjango.models.PayFrequency import PayFrequency

#======================================================================
# 
# Encapsulates data for model PayrollCalendar
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollCalendar Declaration
#======================================================================
class PayrollCalendar (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	country = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payrollRuns = models.ManyToManyField('PayrollRun',  blank=True, related_name='+')
	employees = models.ManyToManyField('Employee',  blank=True, related_name='+')
	payFrequency = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayFrequency])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.country
		str = str + self.payFrequency
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PayrollCalendar";
    
	def objectType(self):
		return "PayrollCalendar";
