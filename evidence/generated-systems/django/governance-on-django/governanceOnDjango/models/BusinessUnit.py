from django.db import models

#======================================================================
# 
# Encapsulates data for model BusinessUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnit Declaration
#======================================================================
class BusinessUnit (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	leader = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	audits = models.ManyToManyField('AuditEngagement',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.leader
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BusinessUnit";
    
	def objectType(self):
		return "BusinessUnit";
