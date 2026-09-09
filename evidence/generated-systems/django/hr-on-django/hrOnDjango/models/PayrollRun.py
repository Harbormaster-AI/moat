from django.db import models
from hrOnDjango.models.PayrollStatus import PayrollStatus

#======================================================================
# 
# Encapsulates data for model PayrollRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollRun Declaration
#======================================================================
class PayrollRun (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	runNumber = models.CharField(max_length=200, null=True)
	periodStart = models.DateField(null=True)
	periodEnd = models.DateField(null=True)
	paymentDate = models.DateField(null=True)
	payrollCalendar = models.ForeignKey('PayrollCalendar', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payrollItems = models.ManyToManyField('PayrollItem',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayrollStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.runNumber
		str = str + self.periodStart
		str = str + self.periodEnd
		str = str + self.paymentDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PayrollRun";
    
	def objectType(self):
		return "PayrollRun";
