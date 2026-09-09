from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.ComplianceAlert import ComplianceAlert
from fintechOnDjango.models.Screening import Screening
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ComplianceAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceAlertDelegate Declaration
#======================================================================
class ComplianceAlertDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, complianceAlertId ):
		try:	
			complianceAlert = ComplianceAlert.objects.filter(id=complianceAlertId)
			return complianceAlert.first();
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError("ComplianceAlert with id " + str(complianceAlertId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, complianceAlert):
		for model in serializers.deserialize("json", complianceAlert):
			model.save()
			return model;

	def create(self, complianceAlert):
		complianceAlert.save()
		return complianceAlert;

	def saveFromJson(self, complianceAlert):
		for model in serializers.deserialize("json", complianceAlert):
			model.save()
			return complianceAlert;
	
	def save(self, complianceAlert):
		complianceAlert.save()
		return complianceAlert;
	
	def delete(self, complianceAlertId ):
		errMsg = "Failed to delete ComplianceAlert from db using id " + str(complianceAlertId)
		
		try:
			complianceAlert = ComplianceAlert.objects.get(id=complianceAlertId)
			complianceAlert.delete()
			return True
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError("ComplianceAlert with id " + str(complianceAlertId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ComplianceAlert.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ComplianceAlert from db")
		except Exception:
			return None;
		
	def assignScreening( self, complianceAlertId, screeningId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

		errMsg = "Failed to assign element " + str(screeningId) + " for Screening on ComplianceAlert"

		try:
			# get the ComplianceAlert from db
			complianceAlert = self.get( complianceAlertId ).first()	
			
			# get the Screening from db
			screening = ScreeningDelegate().get(screeningId).first();
			
			# assign the Screening		
			complianceAlert.screening = screening
			
			#save it
			complianceAlert.save()

			# reload and return the appropriate version					
			return self.get( complianceAlertId );
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert with id " + str(complianceAlertId) + " does not exist.")
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignScreening( self, complianceAlertId ):
		errMsg = "Failed to unassign element " + str(screeningId) + " for Screening on ComplianceAlert"

		try:
			# get the ComplianceAlert from db
			complianceAlert = self.get( complianceAlertId ).first()	
			
			# assign to None for unassignment
			complianceAlert.screening = None			

			#save it
			complianceAlert.save()

			# reload and return the appropriate version					
			return self.get( complianceAlertId );
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert with id " + str(complianceAlertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTransaction( self, complianceAlertId, transactionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to assign element " + str(transactionId) + " for Transaction on ComplianceAlert"

		try:
			# get the ComplianceAlert from db
			complianceAlert = self.get( complianceAlertId ).first()	
			
			# get the Transaction from db
			transaction = TransactionDelegate().get(transactionId).first();
			
			# assign the Transaction		
			complianceAlert.transaction = transaction
			
			#save it
			complianceAlert.save()

			# reload and return the appropriate version					
			return self.get( complianceAlertId );
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert with id " + str(complianceAlertId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTransaction( self, complianceAlertId ):
		errMsg = "Failed to unassign element " + str(transactionId) + " for Transaction on ComplianceAlert"

		try:
			# get the ComplianceAlert from db
			complianceAlert = self.get( complianceAlertId ).first()	
			
			# assign to None for unassignment
			complianceAlert.transaction = None			

			#save it
			complianceAlert.save()

			# reload and return the appropriate version					
			return self.get( complianceAlertId );
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert with id " + str(complianceAlertId) + " does not exist.")
		except Exception:
			return None;
		
