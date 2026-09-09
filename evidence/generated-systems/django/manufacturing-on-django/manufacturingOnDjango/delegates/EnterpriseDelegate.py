from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Enterprise import Enterprise
from manufacturingOnDjango.models.BusinessUnit import BusinessUnit
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.Supplier import Supplier
from manufacturingOnDjango.models.Customer import Customer
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Enterprise
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EnterpriseDelegate Declaration
#======================================================================
class EnterpriseDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, enterpriseId ):
		try:	
			enterprise = Enterprise.objects.filter(id=enterpriseId)
			return enterprise.first();
		except Enterprise.DoesNotExist:
			raise ProcessingError("Enterprise with id " + str(enterpriseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, enterprise):
		for model in serializers.deserialize("json", enterprise):
			model.save()
			return model;

	def create(self, enterprise):
		enterprise.save()
		return enterprise;

	def saveFromJson(self, enterprise):
		for model in serializers.deserialize("json", enterprise):
			model.save()
			return enterprise;
	
	def save(self, enterprise):
		enterprise.save()
		return enterprise;
	
	def delete(self, enterpriseId ):
		errMsg = "Failed to delete Enterprise from db using id " + str(enterpriseId)
		
		try:
			enterprise = Enterprise.objects.get(id=enterpriseId)
			enterprise.delete()
			return True
		except Enterprise.DoesNotExist:
			raise ProcessingError("Enterprise with id " + str(enterpriseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Enterprise.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Enterprise from db")
		except Exception:
			return None;
		
	def addBusinessUnits( self, enterpriseId, businessUnitsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to add elements " + str(businessUnitsIds) + " for BusinessUnits on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = businessUnitsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessUnit		
				businessUnit = BusinessUnitDelegate().get(id).first();	
				# add the BusinessUnit
				enterprise.businessUnits.add(businessUnit)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBusinessUnits( self, enterpriseId, businessUnitsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to remove elements " + str(businessUnitsIds) + " for BusinessUnits on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = businessUnitsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessUnit		
				businessUnit = BusinessUnitDelegate().get(id).first();	
				# add the BusinessUnit
				enterprise.businessUnits.remove(businessUnit)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPlants( self, enterpriseId, plantsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to add elements " + str(plantsIds) + " for Plants on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = plantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Plant		
				plant = PlantDelegate().get(id).first();	
				# add the Plant
				enterprise.plants.add(plant)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlants( self, enterpriseId, plantsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to remove elements " + str(plantsIds) + " for Plants on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = plantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Plant		
				plant = PlantDelegate().get(id).first();	
				# add the Plant
				enterprise.plants.remove(plant)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSuppliers( self, enterpriseId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to add elements " + str(suppliersIds) + " for Suppliers on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				enterprise.suppliers.add(supplier)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSuppliers( self, enterpriseId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to remove elements " + str(suppliersIds) + " for Suppliers on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				enterprise.suppliers.remove(supplier)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCustomers( self, enterpriseId, customersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to add elements " + str(customersIds) + " for Customers on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = customersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Customer		
				customer = CustomerDelegate().get(id).first();	
				# add the Customer
				enterprise.customers.add(customer)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCustomers( self, enterpriseId, customersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to remove elements " + str(customersIds) + " for Customers on Enterprise"

		try:
			# get the Enterprise
			enterprise = self.get( enterpriseId ).first()
				
			# split on a comma with no spaces
			idList = customersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Customer		
				customer = CustomerDelegate().get(id).first();	
				# add the Customer
				enterprise.customers.remove(customer)
				
			# save it		
			enterprise.save()
			
			# reload and return the appropriate version
			return self.get( enterpriseId );
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
