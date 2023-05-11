from django.urls import path
from .views import *


urlpatterns = [
    path('', create_product, name='create'),
    path('get_products/', get_products, name='get_products')
]