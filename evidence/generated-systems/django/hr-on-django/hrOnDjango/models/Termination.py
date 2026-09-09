from django.db import models
from hrOnDjango.models.TerminationReason import TerminationReason
from hrOnDjango.models.TerminationType import TerminationType

#======================================================================
# 
# Encapsulates data for model Termination
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Termination Declaration
#======================================================================
class Termination (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	terminationNumber = models.CharField(max_length=200, null=True)
	terminationDate = models.DateField(null=True)
	notes = models.CharField(max_length=200, null=True)
	eligibleForRehire = models.BooleanField(null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assignment = models.ForeignKey('EmploymentAssignment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reason = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TerminationReason])
	type = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TerminationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.terminationNumber
		str = str + self.terminationDate
		str = str + self.notes
		str = str + self.eligibleForRehire
		str = str + self.reason
		str = str + self.type
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Termination";
    
	def objectType(self):
		return "Termination";
