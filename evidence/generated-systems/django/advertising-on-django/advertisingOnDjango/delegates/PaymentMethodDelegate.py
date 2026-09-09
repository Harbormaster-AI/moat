from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.PaymentMethod import PaymentMethod
from advertisingOnDjango.models.BillingProfile import BillingProfile
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethodDelegate Declaration
#======================================================================
class PaymentMethodDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentMethodId ):
		try:	
			paymentMethod = PaymentMethod.objects.filter(id=paymentMethodId)
			return paymentMethod.first();
		except PaymentMethod.DoesNotExist:
			raise ProcessingError("PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentMethod):
		for model in serializers.deserialize("json", paymentMethod):
			model.save()
			return model;

	def create(self, paymentMethod):
		paymentMethod.save()
		return paymentMethod;

	def saveFromJson(self, paymentMethod):
		for model in serializers.deserialize("json", paymentMethod):
			model.save()
			return paymentMethod;
	
	def save(self, paymentMethod):
		paymentMethod.save()
		return paymentMethod;
	
	def delete(self, paymentMethodId ):
		errMsg = "Failed to delete PaymentMethod from db using id " + str(paymentMethodId)
		
		try:
			paymentMethod = PaymentMethod.objects.get(id=paymentMethodId)
			paymentMethod.delete()
			return True
		except PaymentMethod.DoesNotExist:
			raise ProcessingError("PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentMethod.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentMethod from db")
		except Exception:
			return None;
		
	def assignBillingProfile( self, paymentMethodId, billingProfileId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.BillingProfileDelegate import BillingProfileDelegate

		errMsg = "Failed to assign element " + str(billingProfileId) + " for BillingProfile on PaymentMethod"

		try:
			# get the PaymentMethod from db
			paymentMethod = self.get( paymentMethodId ).first()	
			
			# get the BillingProfile from db
			billingProfile = BillingProfileDelegate().get(billingProfileId).first();
			
			# assign the BillingProfile		
			paymentMethod.billingProfile = billingProfile
			
			#save it
			paymentMethod.save()

			# reload and return the appropriate version					
			return self.get( paymentMethodId );
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBillingProfile( self, paymentMethodId ):
		errMsg = "Failed to unassign element " + str(billingProfileId) + " for BillingProfile on PaymentMethod"

		try:
			# get the PaymentMethod from db
			paymentMethod = self.get( paymentMethodId ).first()	
			
			# assign to None for unassignment
			paymentMethod.billingProfile = None			

			#save it
			paymentMethod.save()

			# reload and return the appropriate version					
			return self.get( paymentMethodId );
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except Exception:
			return None;
		
