from django.db import models

#======================================================================
# 
# Encapsulates data for model RoleAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleAssignment Declaration
#======================================================================
class RoleAssignment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	effectiveFrom = models.DateField(null=True)
	effectiveTo = models.DateField(null=True)
	person = models.ForeignKey('Person', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	role = models.ForeignKey('Role', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	governanceBody = models.ForeignKey('GovernanceBody', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.effectiveFrom
		str = str + self.effectiveTo
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RoleAssignment";
    
	def objectType(self):
		return "RoleAssignment";
