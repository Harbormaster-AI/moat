class ServiceBulletinsController < ApplicationController
  def index
    @serviceBulletins = ServiceBulletin.all
  end
 
  def show
    @serviceBulletin = ServiceBulletin.find(params[:id])
  end
 
  def new
    @serviceBulletin = ServiceBulletin.new
  end
 
  def edit
    @serviceBulletin = ServiceBulletin.find(params[:id])
  end
 
  def create
    @serviceBulletin = ServiceBulletin.new(serviceBulletin_params)
 
    if @serviceBulletin.save
      redirect_to serviceBulletins_path
    else
      render 'new'
    end
  end
 
  def update
    @serviceBulletin = ServiceBulletin.find(params[:id])
 
    if @serviceBulletin.update(serviceBulletin_params)
      redirect_to serviceBulletins_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @serviceBulletin = ServiceBulletin.find(params[:id])
    @serviceBulletin.destroy
    redirect_to serviceBulletins_path
  end

 
  private
    def serviceBulletin_params
      params.require(:serviceBulletin).permit(:bulletinNumber, :Category)
    end
end