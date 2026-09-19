class CatalogsController < ApplicationController
  def index
    @catalogs = Catalog.all
  end
 
  def show
    @catalog = Catalog.find(params[:id])
  end
 
  def new
    @catalog = Catalog.new
  end
 
  def edit
    @catalog = Catalog.find(params[:id])
  end
 
  def create
    @catalog = Catalog.new(catalog_params)
 
    if @catalog.save
      redirect_to catalogs_path
    else
      render 'new'
    end
  end
 
  def update
    @catalog = Catalog.find(params[:id])
 
    if @catalog.update(catalog_params)
      redirect_to catalogs_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @catalog = Catalog.find(params[:id])
    @catalog.destroy
    redirect_to catalogs_path
  end

 
  private
    def catalog_params
      params.require(:catalog).permit(:name, :catalogCode, :asActive)
    end
end