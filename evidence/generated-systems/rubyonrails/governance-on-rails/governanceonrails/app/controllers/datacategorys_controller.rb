class DataCategorysController < ApplicationController
  def index
    @dataCategorys = DataCategory.all
  end
 
  def show
    @dataCategory = DataCategory.find(params[:id])
  end
 
  def new
    @dataCategory = DataCategory.new
  end
 
  def edit
    @dataCategory = DataCategory.find(params[:id])
  end
 
  def create
    @dataCategory = DataCategory.new(dataCategory_params)
 
    if @dataCategory.save
      redirect_to dataCategorys_path
    else
      render 'new'
    end
  end
 
  def update
    @dataCategory = DataCategory.find(params[:id])
 
    if @dataCategory.update(dataCategory_params)
      redirect_to dataCategorys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataCategory = DataCategory.find(params[:id])
    @dataCategory.destroy
    redirect_to dataCategorys_path
  end

 
  private
    def dataCategory_params
      params.require(:dataCategory).permit(:name, :description, :Classification)
    end
end