from django.db import models
from manufacturingOnDjango.models.CAPAStatus import CAPAStatus

#======================================================================
# 
# Encapsulates data for model CorrectiveAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CorrectiveAction Declaration
#======================================================================
class CorrectiveAction (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	capaNumber = models.CharField(max_length=200, null=True)
	rootCause = models.CharField(max_length=200, null=True)
	correctiveAction = models.CharField(max_length=200, null=True)
	verificationDate = models.DateField(null=True)
	nonconformance = models.OneToOneField('Nonconformance', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CAPAStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.capaNumber
		str = str + self.rootCause
		str = str + self.correctiveAction
		str = str + self.verificationDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CorrectiveAction";
    
	def objectType(self):
		return "CorrectiveAction";
