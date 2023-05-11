from django.shortcuts import render
from .models import Product
from django.views.decorators.csrf import csrf_exempt
from django.http import HttpResponse
from httpx import get


@csrf_exempt
def create_product(request):
    if request.method == 'POST':
        name = request.POST.get('name')
        Product.objects.create(name=name)
        get(f'http://express_js:8080/api/create/{name}') # js
        get(f'http://golang:8080/save?name={name}') # golang
    return render(request, 'apps/form.html')


def get_products(request):
    products = Product.objects.all()
    return HttpResponse(f'products : {[i.name for i in products]}')