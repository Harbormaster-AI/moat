from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.DSP import DSP
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DSP
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DSPDelegate Declaration
#======================================================================
class DSPDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dSPId ):
		try:	
			dSP = DSP.objects.filter(id=dSPId)
			return dSP.first();
		except DSP.DoesNotExist:
			raise ProcessingError("DSP with id " + str(dSPId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dSP):
		for model in serializers.deserialize("json", dSP):
			model.save()
			return model;

	def create(self, dSP):
		dSP.save()
		return dSP;

	def saveFromJson(self, dSP):
		for model in serializers.deserialize("json", dSP):
			model.save()
			return dSP;
	
	def save(self, dSP):
		dSP.save()
		return dSP;
	
	def delete(self, dSPId ):
		errMsg = "Failed to delete DSP from db using id " + str(dSPId)
		
		try:
			dSP = DSP.objects.get(id=dSPId)
			dSP.delete()
			return True
		except DSP.DoesNotExist:
			raise ProcessingError("DSP with id " + str(dSPId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DSP.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DSP from db")
		except Exception:
			return None;
		
	def addAdAccounts( self, dSPId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to add elements " + str(adAccountsIds) + " for AdAccounts on DSP"

		try:
			# get the DSP
			dSP = self.get( dSPId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				dSP.adAccounts.add(adAccount)
				
			# save it		
			dSP.save()
			
			# reload and return the appropriate version
			return self.get( dSPId );
		except DSP.DoesNotExist:
			raise ProcessingError(errMsg + " : DSP with id " + str(dSPId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdAccounts( self, dSPId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to remove elements " + str(adAccountsIds) + " for AdAccounts on DSP"

		try:
			# get the DSP
			dSP = self.get( dSPId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				dSP.adAccounts.remove(adAccount)
				
			# save it		
			dSP.save()
			
			# reload and return the appropriate version
			return self.get( dSPId );
		except DSP.DoesNotExist:
			raise ProcessingError(errMsg + " : DSP with id " + str(dSPId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
