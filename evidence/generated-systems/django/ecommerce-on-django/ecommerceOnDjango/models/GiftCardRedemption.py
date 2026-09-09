from django.db import models

#======================================================================
# 
# Encapsulates data for model GiftCardRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardRedemption Declaration
#======================================================================
class GiftCardRedemption (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	redeemedAt = models.DateField(null=True)
	amount = Money
	giftCard = models.ForeignKey('GiftCard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.redeemedAt
		str = str + self.amount
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "GiftCardRedemption";
    
	def objectType(self):
		return "GiftCardRedemption";
