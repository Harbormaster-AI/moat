from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.WorkAuthorization import WorkAuthorization
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Document import Document
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model WorkAuthorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkAuthorizationDelegate Declaration
#======================================================================
class WorkAuthorizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, workAuthorizationId ):
		try:	
			workAuthorization = WorkAuthorization.objects.filter(id=workAuthorizationId)
			return workAuthorization.first();
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError("WorkAuthorization with id " + str(workAuthorizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, workAuthorization):
		for model in serializers.deserialize("json", workAuthorization):
			model.save()
			return model;

	def create(self, workAuthorization):
		workAuthorization.save()
		return workAuthorization;

	def saveFromJson(self, workAuthorization):
		for model in serializers.deserialize("json", workAuthorization):
			model.save()
			return workAuthorization;
	
	def save(self, workAuthorization):
		workAuthorization.save()
		return workAuthorization;
	
	def delete(self, workAuthorizationId ):
		errMsg = "Failed to delete WorkAuthorization from db using id " + str(workAuthorizationId)
		
		try:
			workAuthorization = WorkAuthorization.objects.get(id=workAuthorizationId)
			workAuthorization.delete()
			return True
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError("WorkAuthorization with id " + str(workAuthorizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = WorkAuthorization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all WorkAuthorization from db")
		except Exception:
			return None;
		
	def assignEmployee( self, workAuthorizationId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on WorkAuthorization"

		try:
			# get the WorkAuthorization from db
			workAuthorization = self.get( workAuthorizationId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			workAuthorization.employee = employee
			
			#save it
			workAuthorization.save()

			# reload and return the appropriate version					
			return self.get( workAuthorizationId );
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkAuthorization with id " + str(workAuthorizationId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, workAuthorizationId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on WorkAuthorization"

		try:
			# get the WorkAuthorization from db
			workAuthorization = self.get( workAuthorizationId ).first()	
			
			# assign to None for unassignment
			workAuthorization.employee = None			

			#save it
			workAuthorization.save()

			# reload and return the appropriate version					
			return self.get( workAuthorizationId );
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkAuthorization with id " + str(workAuthorizationId) + " does not exist.")
		except Exception:
			return None;
		
	def addDocuments( self, workAuthorizationId, documentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

		errMsg = "Failed to add elements " + str(documentsIds) + " for Documents on WorkAuthorization"

		try:
			# get the WorkAuthorization
			workAuthorization = self.get( workAuthorizationId ).first()
				
			# split on a comma with no spaces
			idList = documentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Document		
				document = DocumentDelegate().get(id).first();	
				# add the Document
				workAuthorization.documents.add(document)
				
			# save it		
			workAuthorization.save()
			
			# reload and return the appropriate version
			return self.get( workAuthorizationId );
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkAuthorization with id " + str(workAuthorizationId) + " does not exist.")
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDocuments( self, workAuthorizationId, documentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

		errMsg = "Failed to remove elements " + str(documentsIds) + " for Documents on WorkAuthorization"

		try:
			# get the WorkAuthorization
			workAuthorization = self.get( workAuthorizationId ).first()
				
			# split on a comma with no spaces
			idList = documentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Document		
				document = DocumentDelegate().get(id).first();	
				# add the Document
				workAuthorization.documents.remove(document)
				
			# save it		
			workAuthorization.save()
			
			# reload and return the appropriate version
			return self.get( workAuthorizationId );
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkAuthorization with id " + str(workAuthorizationId) + " does not exist.")
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
