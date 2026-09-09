from django.urls import path
from ecommerceOnDjango.views import CategoryView

urlpatterns = [
    path('', CategoryView.index, name='index'),
	path('create', CategoryView.get, name='create'),
	path('get/<int:categoryId>/', CategoryView.get, name='get'),
	path('save', CategoryView.save, name='save'),
	path('getAll', CategoryView.getAll, name='getAll'),
	path('delete/<int:categoryId>/', CategoryView.delete, name='delete'),
	path('assignCatalog/<int:categoryId>/<int:CatalogId>/', CategoryView.assignCatalog, name='assignCatalog'),
	path('unassignCatalog/<int:categoryId>/', CategoryView.unassignCatalog, name='unassignCatalog'),
	path('assignParentCategory/<int:categoryId>/<int:ParentCategoryId>/', CategoryView.assignParentCategory, name='assignParentCategory'),
	path('unassignParentCategory/<int:categoryId>/', CategoryView.unassignParentCategory, name='unassignParentCategory'),
	path('addSubcategories/<int:categoryId>/<SubcategoriesIds>/', CategoryView.addSubcategories, name='addSubcategories'),
	path('removeSubcategories/<int:categoryId>/<SubcategoriesIds>/', CategoryView.removeSubcategories, name='removeSubcategories'),
	path('addProducts/<int:categoryId>/<ProductsIds>/', CategoryView.addProducts, name='addProducts'),
	path('removeProducts/<int:categoryId>/<ProductsIds>/', CategoryView.removeProducts, name='removeProducts'),
]
