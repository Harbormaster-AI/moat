from django.db import models
from governanceOnDjango.models.GovernanceBodyType import GovernanceBodyType

#======================================================================
# 
# Encapsulates data for model GovernanceBody
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GovernanceBody Declaration
#======================================================================
class GovernanceBody (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	charterUrl = URL
	chair = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	roleAssignments = models.ManyToManyField('RoleAssignment',  blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	bodyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in GovernanceBodyType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.charterUrl
		str = str + self.chair
		str = str + self.bodyType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "GovernanceBody";
    
	def objectType(self):
		return "GovernanceBody";
