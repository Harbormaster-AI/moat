from django.db import models
from hrOnDjango.models.LeaveCategory import LeaveCategory
from hrOnDjango.models.AccrualUnit import AccrualUnit

#======================================================================
# 
# Encapsulates data for model LeavePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeavePolicy Declaration
#======================================================================
class LeavePolicy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	accrualRate = models.CharField(max_length=64, null=True)
	carryoverAllowed = models.BooleanField(null=True)
	maxBalance = models.CharField(max_length=64, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	leaveRequests = models.ManyToManyField('LeaveRequest',  blank=True, related_name='+')
	leaveCategory = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LeaveCategory])
	accrualUnit = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccrualUnit])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.accrualRate
		str = str + self.carryoverAllowed
		str = str + self.maxBalance
		str = str + self.leaveCategory
		str = str + self.accrualUnit
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LeavePolicy";
    
	def objectType(self):
		return "LeavePolicy";
