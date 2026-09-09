from django.urls import path
from ecommerceOnDjango.views import ProductView

urlpatterns = [
    path('', ProductView.index, name='index'),
	path('create', ProductView.get, name='create'),
	path('get/<int:productId>/', ProductView.get, name='get'),
	path('save', ProductView.save, name='save'),
	path('getAll', ProductView.getAll, name='getAll'),
	path('delete/<int:productId>/', ProductView.delete, name='delete'),
	path('assignBrand/<int:productId>/<int:BrandId>/', ProductView.assignBrand, name='assignBrand'),
	path('unassignBrand/<int:productId>/', ProductView.unassignBrand, name='unassignBrand'),
	path('assignSeller/<int:productId>/<int:SellerId>/', ProductView.assignSeller, name='assignSeller'),
	path('unassignSeller/<int:productId>/', ProductView.unassignSeller, name='unassignSeller'),
	path('addCategories/<int:productId>/<CategoriesIds>/', ProductView.addCategories, name='addCategories'),
	path('removeCategories/<int:productId>/<CategoriesIds>/', ProductView.removeCategories, name='removeCategories'),
	path('addVariants/<int:productId>/<VariantsIds>/', ProductView.addVariants, name='addVariants'),
	path('removeVariants/<int:productId>/<VariantsIds>/', ProductView.removeVariants, name='removeVariants'),
	path('addMediaAssets/<int:productId>/<MediaAssetsIds>/', ProductView.addMediaAssets, name='addMediaAssets'),
	path('removeMediaAssets/<int:productId>/<MediaAssetsIds>/', ProductView.removeMediaAssets, name='removeMediaAssets'),
	path('addReviews/<int:productId>/<ReviewsIds>/', ProductView.addReviews, name='addReviews'),
	path('removeReviews/<int:productId>/<ReviewsIds>/', ProductView.removeReviews, name='removeReviews'),
]
