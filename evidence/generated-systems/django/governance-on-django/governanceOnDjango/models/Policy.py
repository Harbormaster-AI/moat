from django.db import models
from governanceOnDjango.models.PolicyType import PolicyType
from governanceOnDjango.models.DocumentStatus import DocumentStatus

#======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Policy Declaration
#======================================================================
class Policy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	versionLabel = models.CharField(max_length=200, null=True)
	approvalDate = models.DateField(null=True)
	nextReviewDate = models.DateField(null=True)
	documentUrl = URL
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owners = models.ManyToManyField('Person',  blank=True, related_name='+')
	relatedRequirements = models.ManyToManyField('ComplianceRequirement',  blank=True, related_name='+')
	controls = models.ManyToManyField('Control',  blank=True, related_name='+')
	procedures = models.ManyToManyField('Procedure',  blank=True, related_name='+')
	exceptions = models.ManyToManyField('Exception_',  blank=True, related_name='+')
	attestations = models.ManyToManyField('Attestation',  blank=True, related_name='+')
	policyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PolicyType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DocumentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.versionLabel
		str = str + self.approvalDate
		str = str + self.nextReviewDate
		str = str + self.documentUrl
		str = str + self.policyType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Policy";
    
	def objectType(self):
		return "Policy";
