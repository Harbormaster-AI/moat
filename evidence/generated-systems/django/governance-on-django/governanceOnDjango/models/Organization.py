from django.db import models

#======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Organization Declaration
#======================================================================
class Organization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	jurisdiction = models.CharField(max_length=200, null=True)
	industrySector = models.CharField(max_length=200, null=True)
	governanceBodies = models.ManyToManyField('GovernanceBody',  blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	risks = models.ManyToManyField('Risk',  blank=True, related_name='+')
	thirdParties = models.ManyToManyField('ThirdParty',  blank=True, related_name='+')
	recordsRepositories = models.ManyToManyField('RecordsRepository',  blank=True, related_name='+')
	dataProcessingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	compliancePrograms = models.ManyToManyField('ComplianceProgram',  blank=True, related_name='+')
	auditPrograms = models.ManyToManyField('AuditProgram',  blank=True, related_name='+')
	businessUnits = models.ManyToManyField('BusinessUnit',  blank=True, related_name='+')
	matters = models.ManyToManyField('Matter',  blank=True, related_name='+')
	dataBreaches = models.ManyToManyField('DataBreach',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.jurisdiction
		str = str + self.industrySector
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Organization";
    
	def objectType(self):
		return "Organization";
