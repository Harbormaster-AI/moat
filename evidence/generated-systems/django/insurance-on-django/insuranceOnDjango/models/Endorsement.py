from django.db import models

#======================================================================
# 
# Encapsulates data for model Endorsement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Endorsement Declaration
#======================================================================
class Endorsement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	endorsementNumber = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	description = models.CharField(max_length=200, null=True)
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.endorsementNumber
		str = str + self.effectiveDate
		str = str + self.description
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Endorsement";
    
	def objectType(self):
		return "Endorsement";
