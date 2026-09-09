from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.BillingProfile import BillingProfile
from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.models.PaymentMethod import PaymentMethod
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BillingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingProfileDelegate Declaration
#======================================================================
class BillingProfileDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, billingProfileId ):
		try:	
			billingProfile = BillingProfile.objects.filter(id=billingProfileId)
			return billingProfile.first();
		except BillingProfile.DoesNotExist:
			raise ProcessingError("BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, billingProfile):
		for model in serializers.deserialize("json", billingProfile):
			model.save()
			return model;

	def create(self, billingProfile):
		billingProfile.save()
		return billingProfile;

	def saveFromJson(self, billingProfile):
		for model in serializers.deserialize("json", billingProfile):
			model.save()
			return billingProfile;
	
	def save(self, billingProfile):
		billingProfile.save()
		return billingProfile;
	
	def delete(self, billingProfileId ):
		errMsg = "Failed to delete BillingProfile from db using id " + str(billingProfileId)
		
		try:
			billingProfile = BillingProfile.objects.get(id=billingProfileId)
			billingProfile.delete()
			return True
		except BillingProfile.DoesNotExist:
			raise ProcessingError("BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BillingProfile.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BillingProfile from db")
		except Exception:
			return None;
		
	def assignAdvertiser( self, billingProfileId, advertiserId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

		errMsg = "Failed to assign element " + str(advertiserId) + " for Advertiser on BillingProfile"

		try:
			# get the BillingProfile from db
			billingProfile = self.get( billingProfileId ).first()	
			
			# get the Advertiser from db
			advertiser = AdvertiserDelegate().get(advertiserId).first();
			
			# assign the Advertiser		
			billingProfile.advertiser = advertiser
			
			#save it
			billingProfile.save()

			# reload and return the appropriate version					
			return self.get( billingProfileId );
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdvertiser( self, billingProfileId ):
		errMsg = "Failed to unassign element " + str(advertiserId) + " for Advertiser on BillingProfile"

		try:
			# get the BillingProfile from db
			billingProfile = self.get( billingProfileId ).first()	
			
			# assign to None for unassignment
			billingProfile.advertiser = None			

			#save it
			billingProfile.save()

			# reload and return the appropriate version					
			return self.get( billingProfileId );
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except Exception:
			return None;
		
	def addPaymentMethods( self, billingProfileId, paymentMethodsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PaymentMethodDelegate import PaymentMethodDelegate

		errMsg = "Failed to add elements " + str(paymentMethodsIds) + " for PaymentMethods on BillingProfile"

		try:
			# get the BillingProfile
			billingProfile = self.get( billingProfileId ).first()
				
			# split on a comma with no spaces
			idList = paymentMethodsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentMethod		
				paymentMethod = PaymentMethodDelegate().get(id).first();	
				# add the PaymentMethod
				billingProfile.paymentMethods.add(paymentMethod)
				
			# save it		
			billingProfile.save()
			
			# reload and return the appropriate version
			return self.get( billingProfileId );
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePaymentMethods( self, billingProfileId, paymentMethodsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PaymentMethodDelegate import PaymentMethodDelegate

		errMsg = "Failed to remove elements " + str(paymentMethodsIds) + " for PaymentMethods on BillingProfile"

		try:
			# get the BillingProfile
			billingProfile = self.get( billingProfileId ).first()
				
			# split on a comma with no spaces
			idList = paymentMethodsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentMethod		
				paymentMethod = PaymentMethodDelegate().get(id).first();	
				# add the PaymentMethod
				billingProfile.paymentMethods.remove(paymentMethod)
				
			# save it		
			billingProfile.save()
			
			# reload and return the appropriate version
			return self.get( billingProfileId );
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAdAccounts( self, billingProfileId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to add elements " + str(adAccountsIds) + " for AdAccounts on BillingProfile"

		try:
			# get the BillingProfile
			billingProfile = self.get( billingProfileId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				billingProfile.adAccounts.add(adAccount)
				
			# save it		
			billingProfile.save()
			
			# reload and return the appropriate version
			return self.get( billingProfileId );
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdAccounts( self, billingProfileId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to remove elements " + str(adAccountsIds) + " for AdAccounts on BillingProfile"

		try:
			# get the BillingProfile
			billingProfile = self.get( billingProfileId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				billingProfile.adAccounts.remove(adAccount)
				
			# save it		
			billingProfile.save()
			
			# reload and return the appropriate version
			return self.get( billingProfileId );
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
