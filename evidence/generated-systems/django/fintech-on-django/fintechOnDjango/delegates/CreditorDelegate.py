from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Creditor import Creditor
from fintechOnDjango.models.DirectDebitMandate import DirectDebitMandate
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Creditor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreditorDelegate Declaration
#======================================================================
class CreditorDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, creditorId ):
		try:	
			creditor = Creditor.objects.filter(id=creditorId)
			return creditor.first();
		except Creditor.DoesNotExist:
			raise ProcessingError("Creditor with id " + str(creditorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, creditor):
		for model in serializers.deserialize("json", creditor):
			model.save()
			return model;

	def create(self, creditor):
		creditor.save()
		return creditor;

	def saveFromJson(self, creditor):
		for model in serializers.deserialize("json", creditor):
			model.save()
			return creditor;
	
	def save(self, creditor):
		creditor.save()
		return creditor;
	
	def delete(self, creditorId ):
		errMsg = "Failed to delete Creditor from db using id " + str(creditorId)
		
		try:
			creditor = Creditor.objects.get(id=creditorId)
			creditor.delete()
			return True
		except Creditor.DoesNotExist:
			raise ProcessingError("Creditor with id " + str(creditorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Creditor.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Creditor from db")
		except Exception:
			return None;
		
	def addMandates( self, creditorId, mandatesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DirectDebitMandateDelegate import DirectDebitMandateDelegate

		errMsg = "Failed to add elements " + str(mandatesIds) + " for Mandates on Creditor"

		try:
			# get the Creditor
			creditor = self.get( creditorId ).first()
				
			# split on a comma with no spaces
			idList = mandatesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DirectDebitMandate		
				directDebitMandate = DirectDebitMandateDelegate().get(id).first();	
				# add the DirectDebitMandate
				creditor.mandates.add(directDebitMandate)
				
			# save it		
			creditor.save()
			
			# reload and return the appropriate version
			return self.get( creditorId );
		except Creditor.DoesNotExist:
			raise ProcessingError(errMsg + " : Creditor with id " + str(creditorId) + " does not exist.")
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMandates( self, creditorId, mandatesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DirectDebitMandateDelegate import DirectDebitMandateDelegate

		errMsg = "Failed to remove elements " + str(mandatesIds) + " for Mandates on Creditor"

		try:
			# get the Creditor
			creditor = self.get( creditorId ).first()
				
			# split on a comma with no spaces
			idList = mandatesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DirectDebitMandate		
				directDebitMandate = DirectDebitMandateDelegate().get(id).first();	
				# add the DirectDebitMandate
				creditor.mandates.remove(directDebitMandate)
				
			# save it		
			creditor.save()
			
			# reload and return the appropriate version
			return self.get( creditorId );
		except Creditor.DoesNotExist:
			raise ProcessingError(errMsg + " : Creditor with id " + str(creditorId) + " does not exist.")
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
