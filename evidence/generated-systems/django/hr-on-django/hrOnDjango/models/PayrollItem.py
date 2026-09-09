from django.db import models
from hrOnDjango.models.PayrollItemType import PayrollItemType

#======================================================================
# 
# Encapsulates data for model PayrollItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollItem Declaration
#======================================================================
class PayrollItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	amount = Money
	taxable = models.BooleanField(null=True)
	payrollRun = models.ForeignKey('PayrollRun', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	itemType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayrollItemType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.amount
		str = str + self.taxable
		str = str + self.itemType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PayrollItem";
    
	def objectType(self):
		return "PayrollItem";
