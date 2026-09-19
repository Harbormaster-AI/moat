class BIQuerysController < ApplicationController
  def index
    @bIQuerys = BIQuery.all
  end
 
  def show
    @bIQuery = BIQuery.find(params[:id])
  end
 
  def new
    @bIQuery = BIQuery.new
  end
 
  def edit
    @bIQuery = BIQuery.find(params[:id])
  end
 
  def create
    @bIQuery = BIQuery.new(bIQuery_params)
 
    if @bIQuery.save
      redirect_to bIQuerys_path
    else
      render 'new'
    end
  end
 
  def update
    @bIQuery = BIQuery.find(params[:id])
 
    if @bIQuery.update(bIQuery_params)
      redirect_to bIQuerys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @bIQuery = BIQuery.find(params[:id])
    @bIQuery.destroy
    redirect_to bIQuerys_path
  end

 
  private
    def bIQuery_params
      params.require(:bIQuery).permit(:name, :text, :Dialect)
    end
end