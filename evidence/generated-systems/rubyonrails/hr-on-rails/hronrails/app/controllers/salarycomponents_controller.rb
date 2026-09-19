class SalaryComponentsController < ApplicationController
  def index
    @salaryComponents = SalaryComponent.all
  end
 
  def show
    @salaryComponent = SalaryComponent.find(params[:id])
  end
 
  def new
    @salaryComponent = SalaryComponent.new
  end
 
  def edit
    @salaryComponent = SalaryComponent.find(params[:id])
  end
 
  def create
    @salaryComponent = SalaryComponent.new(salaryComponent_params)
 
    if @salaryComponent.save
      redirect_to salaryComponents_path
    else
      render 'new'
    end
  end
 
  def update
    @salaryComponent = SalaryComponent.find(params[:id])
 
    if @salaryComponent.update(salaryComponent_params)
      redirect_to salaryComponents_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @salaryComponent = SalaryComponent.find(params[:id])
    @salaryComponent.destroy
    redirect_to salaryComponents_path
  end

 
  private
    def salaryComponent_params
      params.require(:salaryComponent).permit(:amount, :recurring, :ComponentType)
    end
end