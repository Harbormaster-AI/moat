from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Component_ import Component_
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Component_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Component_Delegate Declaration
#======================================================================
class Component_Delegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, component_Id ):
		try:	
			component_ = Component_.objects.filter(id=component_Id)
			return component_.first();
		except Component_.DoesNotExist:
			raise ProcessingError("Component_ with id " + str(component_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, component_):
		for model in serializers.deserialize("json", component_):
			model.save()
			return model;

	def create(self, component_):
		component_.save()
		return component_;

	def saveFromJson(self, component_):
		for model in serializers.deserialize("json", component_):
			model.save()
			return component_;
	
	def save(self, component_):
		component_.save()
		return component_;
	
	def delete(self, component_Id ):
		errMsg = "Failed to delete Component_ from db using id " + str(component_Id)
		
		try:
			component_ = Component_.objects.get(id=component_Id)
			component_.delete()
			return True
		except Component_.DoesNotExist:
			raise ProcessingError("Component_ with id " + str(component_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Component_.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Component_ from db")
		except Exception:
			return None;
		
	def assignSupplier( self, component_Id, supplierId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on Component_"

		try:
			# get the Component_ from db
			component_ = self.get( component_Id ).first()	
			
			# get the Supplier from db
			supplier = SupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			component_.supplier = supplier
			
			#save it
			component_.save()

			# reload and return the appropriate version					
			return self.get( component_Id );
		except Component_.DoesNotExist:
			raise ProcessingError(errMsg + " : Component_ with id " + str(component_Id) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, component_Id ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on Component_"

		try:
			# get the Component_ from db
			component_ = self.get( component_Id ).first()	
			
			# assign to None for unassignment
			component_.supplier = None			

			#save it
			component_.save()

			# reload and return the appropriate version					
			return self.get( component_Id );
		except Component_.DoesNotExist:
			raise ProcessingError(errMsg + " : Component_ with id " + str(component_Id) + " does not exist.")
		except Exception:
			return None;
		
