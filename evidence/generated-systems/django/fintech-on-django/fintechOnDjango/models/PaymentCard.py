from django.db import models
from fintechOnDjango.models.CardScheme import CardScheme
from fintechOnDjango.models.CardStatus import CardStatus

#======================================================================
# 
# Encapsulates data for model PaymentCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentCard Declaration
#======================================================================
class PaymentCard (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	cardToken = CardNumberToken
	maskedPan = models.CharField(max_length=200, null=True)
	expiryMonth = models.IntegerField(null=True)
	expiryYear = models.IntegerField(null=True)
	cardholderName = models.CharField(max_length=200, null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	tokenizations = models.ManyToManyField('CardTokenization',  blank=True, related_name='+')
	disputes = models.ManyToManyField('Dispute',  blank=True, related_name='+')
	scheme = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CardScheme])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CardStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.cardToken
		str = str + self.maskedPan
		str = str + self.expiryMonth
		str = str + self.expiryYear
		str = str + self.cardholderName
		str = str + self.scheme
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentCard";
    
	def objectType(self):
		return "PaymentCard";
