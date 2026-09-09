from django.db import models
from insuranceOnDjango.models.ClaimStatus import ClaimStatus
from insuranceOnDjango.models.CauseOfLoss import CauseOfLoss

#======================================================================
# 
# Encapsulates data for model Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Claim Declaration
#======================================================================
class Claim (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	claimNumber = ClaimNumber
	noticeDate = models.DateField(null=True)
	lossDate = models.DateField(null=True)
	reportedBy = models.CharField(max_length=200, null=True)
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adjuster = models.ForeignKey('Adjuster', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	incident = models.OneToOneField('Incident', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	exposures = models.ManyToManyField('Exposure',  blank=True, related_name='+')
	reserves = models.ManyToManyField('ClaimReserve',  blank=True, related_name='+')
	claimPayments = models.ManyToManyField('ClaimPayment',  blank=True, related_name='+')
	serviceProviders = models.ManyToManyField('ServiceProvider',  blank=True, related_name='+')
	subrogations = models.ManyToManyField('SubrogationRecovery',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ClaimStatus])
	lossCause = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CauseOfLoss])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.claimNumber
		str = str + self.noticeDate
		str = str + self.lossDate
		str = str + self.reportedBy
		str = str + self.status
		str = str + self.lossCause
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Claim";
    
	def objectType(self):
		return "Claim";
